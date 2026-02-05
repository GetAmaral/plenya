package services

import (
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
)

type AuthService struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewAuthService(db *gorm.DB, cfg *config.Config) *AuthService {
	return &AuthService{db: db, cfg: cfg}
}

// JWTClaims representa os claims do JWT
type JWTClaims struct {
	UserID string   `json:"userId"`
	Email  string   `json:"email"`
	Roles  []string `json:"roles"`
	jwt.RegisteredClaims
}

// Register cria um novo usuário
func (s *AuthService) Register(req *dto.RegisterRequest) (*dto.AuthResponse, error) {
	// Verificar se usuário já existe
	var existingUser models.User
	if err := s.db.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return nil, ErrUserAlreadyExists
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

// Login autentica um usuário
func (s *AuthService) Login(req *dto.LoginRequest) (*dto.AuthResponse, error) {
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

	// Gerar tokens
	return s.generateAuthResponse(&user)
}

// RefreshToken gera um novo access token a partir do refresh token
func (s *AuthService) RefreshToken(refreshToken string) (*dto.AuthResponse, error) {
	// Validar refresh token
	claims, err := s.validateToken(refreshToken)
	if err != nil {
		return nil, err
	}

	// Buscar usuário com paciente selecionado
	userID, err := uuid.Parse(claims.UserID)
	if err != nil {
		return nil, ErrInvalidToken
	}

	var user models.User
	if err := s.db.Preload("SelectedPatient").First(&user, userID).Error; err != nil {
		return nil, ErrInvalidCredentials
	}

	// Gerar novos tokens
	return s.generateAuthResponse(&user)
}

// generateAuthResponse gera access token, refresh token e resposta
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

	// Gerar access token
	accessToken, err := s.generateToken(user, accessExpiry)
	if err != nil {
		return nil, err
	}

	// Gerar refresh token
	refreshToken, err := s.generateToken(user, refreshExpiry)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         *s.userToDTO(user),
	}, nil
}

// generateToken gera um JWT token
func (s *AuthService) generateToken(user *models.User, expiry time.Duration) (string, error) {
	claims := JWTClaims{
		UserID: user.ID.String(),
		Email:  user.Email,
		Roles:  user.GetRoles(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "plenya-emr",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.cfg.JWT.Secret))
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
