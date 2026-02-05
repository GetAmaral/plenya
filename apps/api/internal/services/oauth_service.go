package services

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

var (
	ErrInvalidIDToken = errors.New("invalid ID token")
	ErrEmailConflict  = errors.New("email já cadastrado com senha - faça login tradicional")
)

type OAuthService struct {
	db          *gorm.DB
	cfg         *config.Config
	authService *AuthService
}

func NewOAuthService(db *gorm.DB, cfg *config.Config, authService *AuthService) *OAuthService {
	return &OAuthService{
		db:          db,
		cfg:         cfg,
		authService: authService,
	}
}

// ValidateGoogleIDToken valida token Google via API oficial
func (s *OAuthService) ValidateGoogleIDToken(ctx context.Context, idToken string) (*oauth2.Tokeninfo, error) {
	oauth2Service, err := oauth2.NewService(ctx, option.WithoutAuthentication())
	if err != nil {
		return nil, fmt.Errorf("failed to create oauth2 service: %w", err)
	}

	tokenInfo, err := oauth2Service.Tokeninfo().IdToken(idToken).Do()
	if err != nil {
		return nil, ErrInvalidIDToken
	}

	// Validar audience (client ID)
	if tokenInfo.Audience != s.cfg.OAuth.GoogleClientID {
		return nil, ErrInvalidIDToken
	}

	// Audience já valida a origem do token (Google OAuth2)
	return tokenInfo, nil
}

// ValidateAppleIDToken valida token Apple (placeholder - implementação completa requer JWT parsing)
func (s *OAuthService) ValidateAppleIDToken(idToken string) (email string, sub string, err error) {
	// TODO: Implementar validação completa de JWT da Apple
	// Por enquanto, retorna erro para forçar uso de Google
	return "", "", errors.New("Apple Sign In não implementado completamente")
}

// AuthenticateOrCreateOAuthUser lógica principal de autenticação OAuth
func (s *OAuthService) AuthenticateOrCreateOAuthUser(email, name, picture, provider, providerID string) (*dto.AuthResponse, error) {
	// 1. Buscar user por OAuth provider + provider ID
	var user models.User
	err := s.db.Preload("SelectedPatient").
		Where("oauth_provider = ? AND oauth_provider_id = ?", provider, providerID).
		First(&user).Error

	if err == nil {
		// Usuário OAuth já existe, atualizar picture se mudou
		if picture != "" && (user.OAuthPictureURL == nil || *user.OAuthPictureURL != picture) {
			user.OAuthPictureURL = &picture
			s.db.Save(&user)
		}

		// Gerar tokens JWT
		return s.authService.GenerateAuthResponse(&user)
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// 2. Não encontrou por OAuth, verificar se email existe
	err = s.db.Where("email = ?", email).First(&user).Error
	if err == nil {
		// Email existe
		if user.PasswordHash != nil {
			// Tem senha - conflito (não pode linkar automaticamente)
			return nil, ErrEmailConflict
		}

		// Email existe mas é OAuth de outro provider - por segurança, não linkar automaticamente
		return nil, ErrEmailConflict
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// 3. Email não existe - criar novo User + Patient
	return s.createOAuthUser(email, name, picture, provider, providerID)
}

// createOAuthUser cria novo usuário OAuth com Patient automático
func (s *OAuthService) createOAuthUser(email, name, picture, provider, providerID string) (*dto.AuthResponse, error) {
	// Iniciar transação
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Criar User
	user := models.User{
		Name:            name,
		Email:           email,
		PasswordHash:    nil, // OAuth users não têm senha
		OAuthProvider:   &provider,
		OAuthProviderID: &providerID,
		OAuthPictureURL: nil,
	}

	if picture != "" {
		user.OAuthPictureURL = &picture
	}

	// Set role patient
	if err := user.SetRoles([]string{string(models.RolePatient)}); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Criar Patient automaticamente
	// CPF nullable - pode preencher depois
	// BirthDate placeholder - deve preencher depois
	// Gender "other" - deve preencher depois
	birthDate := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)

	patient := models.Patient{
		UserID:    user.ID,
		Name:      name,
		CPF:       nil, // Nullable agora
		BirthDate: birthDate,
		Gender:    models.GenderOther,
	}

	if err := tx.Create(&patient).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Auto-selecionar próprio patient
	user.SelectedPatientID = &patient.ID
	user.SelectedPatient = &patient // Para retornar no response

	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Commit transação
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	// Gerar tokens JWT
	return s.authService.GenerateAuthResponse(&user)
}

// GoogleCallback processa callback do Google OAuth
func (s *OAuthService) GoogleCallback(ctx context.Context, idToken string) (*dto.AuthResponse, error) {
	// Validar ID token
	tokenInfo, err := s.ValidateGoogleIDToken(ctx, idToken)
	if err != nil {
		return nil, err
	}

	// Extrair dados (campos disponíveis: Email, UserId, VerifiedEmail, Audience)
	email := tokenInfo.Email
	name := tokenInfo.Email // Default para email (API v2 não retorna Name na tokeninfo)
	picture := ""           // API v2 não retorna Picture na tokeninfo
	providerID := tokenInfo.UserId

	// Autenticar ou criar usuário
	return s.AuthenticateOrCreateOAuthUser(email, name, picture, "google", providerID)
}

// AppleCallback processa callback da Apple Sign In
func (s *OAuthService) AppleCallback(ctx context.Context, idToken string, userName *dto.AppleUserName) (*dto.AuthResponse, error) {
	// Validar ID token (placeholder)
	email, sub, err := s.ValidateAppleIDToken(idToken)
	if err != nil {
		return nil, err
	}

	// Montar nome
	name := email // Default
	if userName != nil && userName.FirstName != "" {
		name = strings.TrimSpace(userName.FirstName + " " + userName.LastName)
	}

	// Autenticar ou criar usuário
	return s.AuthenticateOrCreateOAuthUser(email, name, "", "apple", sub)
}
