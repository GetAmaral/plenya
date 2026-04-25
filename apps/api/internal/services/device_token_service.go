package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

var (
	// ErrDeviceTokenNotFound é retornado quando o token não existe ou pertence a outro usuário.
	ErrDeviceTokenNotFound = errors.New("device token not found")
)

// DeviceTokenService gerencia tokens de push de devices autenticados.
type DeviceTokenService struct {
	db *gorm.DB
}

// NewDeviceTokenService instancia o service.
func NewDeviceTokenService(db *gorm.DB) *DeviceTokenService {
	return &DeviceTokenService{db: db}
}

// RegisterTokenInput agrupa os campos necessários para registrar/atualizar um device token.
type RegisterTokenInput struct {
	UserID      uuid.UUID
	Platform    models.DevicePlatform
	AppVariant  models.AppVariant
	Token       string
	DeviceLabel *string
	AppVersion  *string
}

// Register cria um novo device token ou atualiza o existente (idempotente por token+app_variant).
func (s *DeviceTokenService) Register(input RegisterTokenInput) (*models.DeviceToken, error) {
	var existing models.DeviceToken
	err := s.db.Where("token = ? AND app_variant = ?", input.Token, input.AppVariant).
		First(&existing).Error

	if err == nil {
		existing.UserID = input.UserID
		existing.Platform = input.Platform
		existing.DeviceLabel = input.DeviceLabel
		existing.AppVersion = input.AppVersion
		existing.LastSeenAt = time.Now()
		if err := s.db.Save(&existing).Error; err != nil {
			return nil, err
		}
		return &existing, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	dt := models.DeviceToken{
		UserID:      input.UserID,
		Platform:    input.Platform,
		AppVariant:  input.AppVariant,
		Token:       input.Token,
		DeviceLabel: input.DeviceLabel,
		AppVersion:  input.AppVersion,
		LastSeenAt:  time.Now(),
	}
	if err := s.db.Create(&dt).Error; err != nil {
		return nil, err
	}
	return &dt, nil
}

// ListByUser retorna todos os tokens ativos de um usuário, ordenados por LastSeenAt desc.
func (s *DeviceTokenService) ListByUser(userID uuid.UUID) ([]models.DeviceToken, error) {
	var tokens []models.DeviceToken
	if err := s.db.Where("user_id = ?", userID).
		Order("last_seen_at desc").
		Find(&tokens).Error; err != nil {
		return nil, err
	}
	return tokens, nil
}

// Revoke faz soft delete de um token específico do usuário.
func (s *DeviceTokenService) Revoke(userID, tokenID uuid.UUID) error {
	res := s.db.Where("id = ? AND user_id = ?", tokenID, userID).
		Delete(&models.DeviceToken{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return ErrDeviceTokenNotFound
	}
	return nil
}

// Touch atualiza o LastSeenAt do token (chamado por middleware quando o app faz request).
func (s *DeviceTokenService) Touch(tokenID uuid.UUID) error {
	return s.db.Model(&models.DeviceToken{}).
		Where("id = ?", tokenID).
		Update("last_seen_at", time.Now()).Error
}
