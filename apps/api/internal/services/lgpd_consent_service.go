package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// LGPDConsentService gerencia o aceite do termo de uso/LGPD por usuário.
// Reutiliza ErrUserNotFound de user_service.go.
type LGPDConsentService struct {
	db *gorm.DB
}

// NewLGPDConsentService instancia o service.
func NewLGPDConsentService(db *gorm.DB) *LGPDConsentService {
	return &LGPDConsentService{db: db}
}

// LGPDConsentStatus reporta se o usuário aceitou o termo e quando.
type LGPDConsentStatus struct {
	Accepted    bool       `json:"accepted"`
	ConsentedAt *time.Time `json:"consentedAt,omitempty"`
}

// Status retorna o estado do consentimento do usuário.
func (s *LGPDConsentService) Status(userID uuid.UUID) (*LGPDConsentStatus, error) {
	var user models.User
	if err := s.db.Select("id", "lgpd_consented_at").
		Where("id = ?", userID).
		First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &LGPDConsentStatus{
		Accepted:    user.LGPDConsentedAt != nil,
		ConsentedAt: user.LGPDConsentedAt,
	}, nil
}

// Accept registra o aceite do termo (idempotente — se já aceitou, mantém timestamp original).
func (s *LGPDConsentService) Accept(userID uuid.UUID) (*LGPDConsentStatus, error) {
	var user models.User
	if err := s.db.Where("id = ?", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	if user.LGPDConsentedAt == nil {
		now := time.Now()
		user.LGPDConsentedAt = &now
		if err := s.db.Save(&user).Error; err != nil {
			return nil, err
		}
	}

	return &LGPDConsentStatus{
		Accepted:    true,
		ConsentedAt: user.LGPDConsentedAt,
	}, nil
}
