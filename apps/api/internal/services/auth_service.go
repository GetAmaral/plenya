package services

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidToken       = errors.New("invalid token")
	ErrTokenRevoked       = errors.New("token revoked or already used")
)

// Token types — distinção crítica pra middleware de auth aceitar APENAS access tokens.
const (
	TokenTypeAccess  = "access"
	TokenTypeRefresh = "refresh"
	// TokenTypeMFAChallenge é emitido após login bem-sucedido quando o user tem 2FA
	// habilitado. TTL curto (5min). Validado em /auth/login/verify-2fa.
	TokenTypeMFAChallenge = "mfa_challenge"
)

// hashToken — SHA-256 hex do token. Nunca armazenamos o token plano em DB.
func hashToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}

type AuthService struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewAuthService(db *gorm.DB, cfg *config.Config) *AuthService {
	return &AuthService{db: db, cfg: cfg}
}

// JWTClaims representa os claims do JWT.
// Campo Type ("access"|"refresh"|"mfa_challenge") evita que refresh tokens sejam
// usados como access tokens (CRITICAL — bug clássico permitia escalada).
type JWTClaims struct {
	UserID string   `json:"userId"`
	Email  string   `json:"email"`
	Roles  []string `json:"roles"`
	Type   string   `json:"type,omitempty"`
	jwt.RegisteredClaims
}

// Register cria um novo usuário
func (s *AuthService) Register(req *dto.RegisterRequest) (*dto.AuthResponse, error) {
	// Verificar se usuário já existe
	var existingUser models.User
	if err := s.db.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return nil, ErrUserAlreadyExists
	}

	// M5 — força mínima de senha (12+ chars, 1 letra + 1 número).
	// dto.RegisterRequest validator marca min=8 (legado); reforçamos aqui.
	if err := validatePasswordStrength(req.Password); err != nil {
		return nil, err
	}

	// Hash da senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	passwordStr := string(hashedPassword)

	// Criar usuário
	user := models.User{
		Email:        req.Email,
		PasswordHash: &passwordStr,
	}

	// Set roles
	if err := user.SetRoles(req.Roles); err != nil {
		return nil, err
	}

	if err := s.db.Create(&user).Error; err != nil {
		return nil, err
	}

	// Gerar tokens
	return s.generateAuthResponse(&user)
}

// LoginAttempt — resultado do Login com possível MFA challenge.
// Quando MFARequired=true, o caller NÃO recebe tokens; precisa chamar
// VerifyMFAChallenge(MFAToken, code).
type LoginAttempt struct {
	Auth                 *dto.AuthResponse
	MFARequired          bool
	MFAToken             string // JWT short-lived com Type="mfa_challenge"
	MFAEnrollmentMissing bool   // role exige 2FA mas user ainda não habilitou
}

// Login autentica um usuário.
// Retorna LoginAttempt — caller decide o que fazer com MFAToken vs AuthResponse.
func (s *AuthService) Login(req *dto.LoginRequest) (*LoginAttempt, error) {
	// Buscar usuário com paciente selecionado
	var user models.User
	if err := s.db.Preload("SelectedPatient").Where("email = ?", req.Email).First(&user).Error; err != nil {
		return nil, ErrInvalidCredentials
	}

	// Verificar se é usuário OAuth (sem senha)
	if user.PasswordHash == nil {
		provider := "OAuth"
		if user.OAuthProvider != nil {
			provider = *user.OAuthProvider
		}
		return nil, errors.New("esta conta usa login " + provider + " - use o botão OAuth")
	}

	// Verificar senha
	if err := bcrypt.CompareHashAndPassword([]byte(*user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	// 2FA habilitado → emite challenge token, NÃO emite access/refresh.
	if user.TwoFactorEnabled {
		challenge, err := s.generateMFAChallengeToken(&user)
		if err != nil {
			return nil, err
		}
		return &LoginAttempt{
			MFARequired: true,
			MFAToken:    challenge,
		}, nil
	}

	// 2FA não habilitado, mas role exige (admin/doctor/manager) → flag pro
	// frontend redirecionar pra setup. Login segue (não bloqueia logins existentes).
	resp, err := s.generateAuthResponse(&user)
	if err != nil {
		return nil, err
	}
	enrollmentRequired := false
	for _, r := range user.GetRoles() {
		if r == string(models.RoleAdmin) || r == string(models.RoleDoctor) || r == string(models.RoleManager) {
			enrollmentRequired = true
			break
		}
	}
	return &LoginAttempt{
		Auth:                 resp,
		MFAEnrollmentMissing: enrollmentRequired,
	}, nil
}

// generateMFAChallengeToken — emite JWT curto (5min) com Type=mfa_challenge.
// O caller troca esse token + código TOTP pelo par access/refresh em
// /auth/login/verify-2fa.
func (s *AuthService) generateMFAChallengeToken(user *models.User) (string, error) {
	claims := JWTClaims{
		UserID: user.ID.String(),
		Email:  user.Email,
		Type:   TokenTypeMFAChallenge,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "plenya-emr",
		},
	}
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tok.SignedString([]byte(s.cfg.JWT.Secret))
}

// VerifyMFAChallenge — consome MFAToken + TOTP code, emite par access/refresh.
// Usa pquerna/otp/totp pra validar (caller injeta validador via parâmetro pra
// não acoplar dependência aqui — ver auth_2fa.go).
func (s *AuthService) VerifyMFAChallenge(challengeToken, totpCode string) (*dto.AuthResponse, error) {
	claims, err := s.validateToken(challengeToken)
	if err != nil {
		return nil, ErrInvalidToken
	}
	if claims.Type != TokenTypeMFAChallenge {
		return nil, ErrInvalidToken
	}
	userID, err := uuid.Parse(claims.UserID)
	if err != nil {
		return nil, ErrInvalidToken
	}
	var user models.User
	if err := s.db.Preload("SelectedPatient").First(&user, userID).Error; err != nil {
		return nil, ErrInvalidCredentials
	}
	if !user.TwoFactorEnabled || user.TwoFactorSecret == "" {
		return nil, ErrInvalidToken
	}
	if !ValidateTOTP(user.TwoFactorSecret, totpCode) {
		return nil, ErrInvalidCredentials
	}
	return s.generateAuthResponse(&user)
}

// RefreshToken gera um novo par (access + refresh) a partir de um refresh token.
// Política de rotação: revoga o refresh atual antes de emitir o novo, garantindo
// single-use. Se o token recebido já estiver revogado/usado, retorna ErrTokenRevoked
// (sinal pra forçar re-login + alerta de possível roubo).
func (s *AuthService) RefreshToken(refreshToken string) (*dto.AuthResponse, error) {
	// Validar JWT formato/assinatura
	claims, err := s.validateToken(refreshToken)
	if err != nil {
		return nil, err
	}
	// CRITICAL: refresh token DEVE ser de tipo "refresh" — bloqueia uso de
	// access token no /refresh (que rotaciona e dá tokens novos).
	if claims.Type != TokenTypeRefresh {
		return nil, ErrInvalidToken
	}

	// Validar contra tabela refresh_tokens (rotação + revogação)
	hash := hashToken(refreshToken)
	var rt models.RefreshToken
	if err := s.db.Where("token_hash = ? AND type = ?", hash, "refresh").First(&rt).Error; err != nil {
		// Token JWT válido mas não está na tabela — possivelmente revogado
		// num logout anterior. Tratamos como invalid pra não vazar info.
		return nil, ErrInvalidToken
	}
	if !rt.IsActive() {
		return nil, ErrTokenRevoked
	}

	// Buscar usuário
	userID, err := uuid.Parse(claims.UserID)
	if err != nil {
		return nil, ErrInvalidToken
	}
	var user models.User
	if err := s.db.Preload("SelectedPatient").First(&user, userID).Error; err != nil {
		return nil, ErrInvalidCredentials
	}

	// Revoga o refresh atual ANTES de emitir o novo (rotação atômica)
	now := time.Now().UTC()
	if err := s.db.Model(&rt).Update("revoked_at", now).Error; err != nil {
		return nil, err
	}

	// Gerar e persistir novo par
	return s.generateAuthResponse(&user)
}

// Logout — revoga TODOS os refresh tokens ativos do user. Implementação
// simples (sem need do client mandar o refresh): qualquer device com sessão
// ativa precisará re-logar. Trade-off é aceitável pra logout explícito.
func (s *AuthService) Logout(userID uuid.UUID) error {
	now := time.Now().UTC()
	return s.db.Model(&models.RefreshToken{}).
		Where("user_id = ? AND type = ? AND revoked_at IS NULL", userID, "refresh").
		Update("revoked_at", now).Error
}

// GenerateTokensForUser é o entry-point público para serviços externos (ex: magic link
// do Score Light) gerarem AuthResponse válida sem exigir senha.
func (s *AuthService) GenerateTokensForUser(user *models.User) (*dto.AuthResponse, error) {
	return s.generateAuthResponse(user)
}

// generateAuthResponse gera access token + refresh token, persiste o hash
// do refresh em refresh_tokens, e retorna o par. AccessToken é stateless
// (não persistido). RefreshToken é stateful (validado contra DB no /refresh).
func (s *AuthService) generateAuthResponse(user *models.User) (*dto.AuthResponse, error) {
	// Parse access expiry
	accessExpiry, err := time.ParseDuration(s.cfg.JWT.AccessExpiry)
	if err != nil {
		return nil, errors.New("invalid access expiry configuration")
	}

	// Parse refresh expiry
	refreshExpiry, err := time.ParseDuration(s.cfg.JWT.RefreshExpiry)
	if err != nil {
		return nil, errors.New("invalid refresh expiry configuration")
	}

	// Gerar access token (Type=access)
	accessToken, err := s.generateTypedToken(user, TokenTypeAccess, accessExpiry)
	if err != nil {
		return nil, err
	}

	// Gerar refresh token (Type=refresh) e persistir hash
	refreshToken, err := s.generateTypedToken(user, TokenTypeRefresh, refreshExpiry)
	if err != nil {
		return nil, err
	}
	rt := models.RefreshToken{
		UserID:    user.ID,
		TokenHash: hashToken(refreshToken),
		Type:      "refresh",
		ExpiresAt: time.Now().UTC().Add(refreshExpiry),
	}
	if err := s.db.Create(&rt).Error; err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         *s.userToDTO(user),
	}, nil
}

// generateTypedToken — gera JWT marcando o Type (access|refresh).
func (s *AuthService) generateTypedToken(user *models.User, tokenType string, expiry time.Duration) (string, error) {
	claims := JWTClaims{
		UserID: user.ID.String(),
		Email:  user.Email,
		Roles:  user.GetRoles(),
		Type:   tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "plenya-emr",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.cfg.JWT.Secret))
}

// generateToken — DEPRECATED. Mantido pra retrocompat (alguns serviços externos
// chamam direto). Emite token com Type=access.
func (s *AuthService) generateToken(user *models.User, expiry time.Duration) (string, error) {
	return s.generateTypedToken(user, TokenTypeAccess, expiry)
}

// validateToken valida um JWT token e retorna os claims
func (s *AuthService) validateToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.JWT.Secret), nil
	})

	if err != nil {
		return nil, ErrInvalidToken
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}

// GetUserByID busca um usuário por ID incluindo paciente selecionado
func (s *AuthService) GetUserByID(userID uuid.UUID) (*dto.UserDTO, error) {
	var user models.User
	if err := s.db.Preload("SelectedPatient").First(&user, userID).Error; err != nil {
		return nil, err
	}

	return s.userToDTO(&user), nil
}

// UpdateSelectedPatient atualiza o paciente selecionado do usuário
func (s *AuthService) UpdateSelectedPatient(userID, patientID uuid.UUID) (*dto.UserDTO, error) {
	// Verificar se paciente existe (Select apenas ID para evitar carregar relações)
	var exists bool
	if err := s.db.Model(&models.Patient{}).
		Select("1").
		Where("id = ?", patientID).
		Where("deleted_at IS NULL").
		Limit(1).
		Find(&exists).Error; err != nil {
		return nil, err
	}

	// Se não encontrou o paciente
	var count int64
	if err := s.db.Model(&models.Patient{}).
		Where("id = ?", patientID).
		Where("deleted_at IS NULL").
		Count(&count).Error; err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, ErrPatientNotFound
	}

	// Atualizar usuário usando UpdateColumn para evitar hooks de validação
	if err := s.db.Model(&models.User{}).Where("id = ?", userID).UpdateColumn("selected_patient_id", patientID).Error; err != nil {
		return nil, err
	}

	// Buscar usuário atualizado com relação
	return s.GetUserByID(userID)
}

// ChangePassword troca a senha do usuário após validar a senha atual.
// Retorna ErrInvalidCredentials se a senha atual estiver errada (mesmo
// erro do Login pra não vazar info).
func (s *AuthService) ChangePassword(userID uuid.UUID, current, next string) error {
	var user models.User
	if err := s.db.Select("id", "password_hash").First(&user, userID).Error; err != nil {
		return ErrInvalidCredentials
	}
	if user.PasswordHash == nil {
		// OAuth-only user — não tem senha pra trocar
		return ErrInvalidCredentials
	}
	if err := bcrypt.CompareHashAndPassword([]byte(*user.PasswordHash), []byte(current)); err != nil {
		return ErrInvalidCredentials
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(next), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	hashStr := string(hash)
	return s.db.Model(&models.User{}).Where("id = ?", userID).
		UpdateColumn("password_hash", hashStr).Error
}

// Enable2FAResult — payload retornado por StartEnable2FA pra o frontend
// montar o QR code. Secret é exibido uma única vez (caso QR scan falhe).
type Enable2FAResult struct {
	Secret      string `json:"secret"`
	OTPAuthURL  string `json:"otpAuthUrl"`
}

// StartEnable2FA — gera secret + URL de provisioning. Persiste o secret
// como pendente (TwoFactorSecret setado, TwoFactorEnabled=false) — só vira
// true após o usuário validar o primeiro código em ConfirmEnable2FA.
func (s *AuthService) StartEnable2FA(userID uuid.UUID) (*Enable2FAResult, error) {
	var user models.User
	if err := s.db.First(&user, userID).Error; err != nil {
		return nil, ErrInvalidCredentials
	}
	secret, err := GenerateTOTPSecret()
	if err != nil {
		return nil, err
	}
	if err := s.db.Model(&user).Update("two_factor_secret", secret).Error; err != nil {
		return nil, err
	}
	return &Enable2FAResult{
		Secret:     secret,
		OTPAuthURL: BuildOTPAuthURL("Plenya EMR", user.Email, secret),
	}, nil
}

// ConfirmEnable2FA — valida o primeiro código TOTP e habilita 2FA.
// Retorna ErrInvalidCredentials se o código não bater (NÃO habilita).
func (s *AuthService) ConfirmEnable2FA(userID uuid.UUID, code string) error {
	var user models.User
	if err := s.db.First(&user, userID).Error; err != nil {
		return ErrInvalidCredentials
	}
	if user.TwoFactorSecret == "" {
		return errors.New("2FA enrollment not started — call /2fa/enable first")
	}
	if !ValidateTOTP(user.TwoFactorSecret, code) {
		return ErrInvalidCredentials
	}
	return s.db.Model(&user).Update("two_factor_enabled", true).Error
}

// Disable2FA remove o segredo TOTP do usuário após validar a senha.
// Idempotente: se 2FA já está desabilitado, zera secret novamente sem erro.
func (s *AuthService) Disable2FA(userID uuid.UUID, password string) error {
	var user models.User
	if err := s.db.Select("id", "password_hash").First(&user, userID).Error; err != nil {
		return ErrInvalidCredentials
	}
	if user.PasswordHash == nil {
		return ErrInvalidCredentials
	}
	if err := bcrypt.CompareHashAndPassword([]byte(*user.PasswordHash), []byte(password)); err != nil {
		return ErrInvalidCredentials
	}
	return s.db.Model(&models.User{}).Where("id = ?", userID).
		Updates(map[string]interface{}{
			"two_factor_enabled": false,
			"two_factor_secret":  "",
		}).Error
}

// UpdatePreferences atualiza as preferências do usuário
func (s *AuthService) UpdatePreferences(userID uuid.UUID, preferences map[string]interface{}) (*dto.UserDTO, error) {
	// Atualizar preferências usando UpdateColumn
	if err := s.db.Model(&models.User{}).Where("id = ?", userID).UpdateColumn("preferences", preferences).Error; err != nil {
		return nil, err
	}

	// Buscar usuário atualizado
	return s.GetUserByID(userID)
}

// userToDTO converte User para UserDTO incluindo selectedPatient
func (s *AuthService) userToDTO(user *models.User) *dto.UserDTO {
	userDTO := &dto.UserDTO{
		ID:               user.ID.String(),
		Name:             user.Name,
		Email:            user.Email,
		CPF:              user.CPF,
		Roles:            user.GetRoles(),
		TwoFactorEnabled: user.TwoFactorEnabled,
		CreatedAt:        user.CreatedAt.Format(time.RFC3339),
	}

	// Adicionar selectedPatientId se existir
	if user.SelectedPatientID != nil {
		selectedPatientIDStr := user.SelectedPatientID.String()
		userDTO.SelectedPatientID = &selectedPatientIDStr
	}

	// Adicionar selectedPatient se foi carregado
	if user.SelectedPatient != nil {
		userDTO.SelectedPatient = &dto.PatientResponse{
			ID:           user.SelectedPatient.ID.String(),
			UserID:       user.SelectedPatient.UserID.String(),
			Name:         user.SelectedPatient.Name,
			BirthDate:    user.SelectedPatient.BirthDate.Format("2006-01-02"),
			Gender:       user.SelectedPatient.Gender,
			Age:          user.SelectedPatient.Age,
			AgeText:      user.SelectedPatient.AgeText,
			Menopause:    user.SelectedPatient.Menopause,
			Phone:        user.SelectedPatient.Phone,
			Address:      user.SelectedPatient.Address,
			Municipality: user.SelectedPatient.Municipality,
			State:        user.SelectedPatient.State,
			MotherName:   user.SelectedPatient.MotherName,
			FatherName:   user.SelectedPatient.FatherName,
			Height:       user.SelectedPatient.Height,
			Weight:       user.SelectedPatient.Weight,
			CreatedAt:    user.SelectedPatient.CreatedAt.Format(time.RFC3339),
			UpdatedAt:    user.SelectedPatient.UpdatedAt.Format(time.RFC3339),
		}
	}

	// Adicionar preferences se existir
	if len(user.Preferences) > 0 {
		var prefs map[string]interface{}
		if err := json.Unmarshal(user.Preferences, &prefs); err == nil {
			userDTO.Preferences = prefs
		}
	}

	return userDTO
}

// GenerateAuthResponse gera JWT tokens para um user (público para uso em OAuthService)
func (s *AuthService) GenerateAuthResponse(user *models.User) (*dto.AuthResponse, error) {
	return s.generateAuthResponse(user)
}
