package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// DataExportService implementa LGPD Art. 18, V — direito à portabilidade.
// Retorna JSON com tudo que armazenamos vinculado ao userID. Inclui dados
// derivados que o user produziu (anamneses redigidas, prescrições assinadas)
// e dados pessoais do User. Não inclui dados de outros usuários (mesmo que
// tenham relação — médico não exporta paciente alheio).
type DataExportService struct {
	db *gorm.DB
}

func NewDataExportService(db *gorm.DB) *DataExportService {
	return &DataExportService{db: db}
}

// DataExport é o payload retornado em /me/export.
type DataExport struct {
	ExportedAt    time.Time              `json:"exportedAt"`
	UserID        uuid.UUID              `json:"userId"`
	User          map[string]interface{} `json:"user"`
	Sessions      []map[string]any       `json:"sessions"`
	Notifications []map[string]any       `json:"notifications"`
	Preferences   any                    `json:"preferences,omitempty"`
}

// Export coleta os dados pessoais do usuário.
func (s *DataExportService) Export(_ context.Context, userID uuid.UUID) (*DataExport, error) {
	var user models.User
	if err := s.db.First(&user, userID).Error; err != nil {
		return nil, err
	}

	var deviceTokens []models.DeviceToken
	if err := s.db.Where("user_id = ?", userID).Find(&deviceTokens).Error; err != nil {
		return nil, err
	}

	var notifications []models.Notification
	if err := s.db.Where("user_id = ?", userID).
		Order("created_at desc").
		Limit(500).
		Find(&notifications).Error; err != nil {
		return nil, err
	}

	sessions := make([]map[string]any, 0, len(deviceTokens))
	for _, d := range deviceTokens {
		sessions = append(sessions, map[string]any{
			"id":         d.ID,
			"platform":   d.Platform,
			"appVariant": d.AppVariant,
			"appVersion": d.AppVersion,
			"lastSeenAt": d.LastSeenAt,
			"createdAt":  d.CreatedAt,
		})
	}

	notifPayload := make([]map[string]any, 0, len(notifications))
	for _, n := range notifications {
		notifPayload = append(notifPayload, map[string]any{
			"id":        n.ID,
			"type":      n.Type,
			"title":     n.Title,
			"message":   n.Message,
			"read":      n.Read,
			"createdAt": n.CreatedAt,
		})
	}

	userPayload := map[string]interface{}{
		"id":               user.ID,
		"name":             user.Name,
		"email":            user.Email,
		"roles":            user.GetRoles(),
		"twoFactorEnabled": user.TwoFactorEnabled,
		"crm":              user.CRM,
		"crmUF":            user.CRMUF,
		"specialty":        user.Specialty,
		"professionalPhone": user.ProfessionalPhone,
		"createdAt":        user.CreatedAt,
		"lgpdConsentedAt":  user.LGPDConsentedAt,
	}
	if user.CPF != nil {
		userPayload["cpf"] = *user.CPF
	}

	return &DataExport{
		ExportedAt:    time.Now().UTC(),
		UserID:        user.ID,
		User:          userPayload,
		Sessions:      sessions,
		Notifications: notifPayload,
		Preferences:   user.Preferences,
	}, nil
}
