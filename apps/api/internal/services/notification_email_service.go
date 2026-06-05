package services

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/plenya/api/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// NotificationEmailService gerencia o bucket de e-mails automáticos (no-reply, newsletters).
type NotificationEmailService struct {
	db *gorm.DB
}

func NewNotificationEmailService(db *gorm.DB) *NotificationEmailService {
	return &NotificationEmailService{db: db}
}

// NotificationEmailInput é o payload de gravação a partir do ingest de e-mail.
type NotificationEmailInput struct {
	FromEmail  string
	FromName   *string
	Subject    string
	BodyText   string
	MessageID  *string
	ReceivedAt time.Time
}

// Record grava um e-mail automático. Idempotente por message_id (ON CONFLICT DO NOTHING),
// pra reprocessamento do ingest não duplicar.
func (s *NotificationEmailService) Record(in NotificationEmailInput) error {
	if in.ReceivedAt.IsZero() {
		in.ReceivedAt = time.Now().UTC()
	}
	rec := models.NotificationEmail{
		FromEmail:  in.FromEmail,
		FromName:   in.FromName,
		Subject:    in.Subject,
		BodyText:   in.BodyText,
		MessageID:  in.MessageID,
		ReceivedAt: in.ReceivedAt,
	}
	if err := s.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&rec).Error; err != nil {
		return fmt.Errorf("notification email: record: %w", err)
	}
	return nil
}

// List retorna e-mails automáticos (mais recentes primeiro). unreadOnly filtra não-lidos.
func (s *NotificationEmailService) List(limit int, unreadOnly bool) ([]models.NotificationEmail, error) {
	if limit <= 0 {
		limit = 100
	}
	if limit > 500 {
		limit = 500
	}
	q := s.db.Model(&models.NotificationEmail{})
	if unreadOnly {
		q = q.Where("is_read = ?", false)
	}
	var items []models.NotificationEmail
	if err := q.Order("received_at DESC").Limit(limit).Find(&items).Error; err != nil {
		return nil, fmt.Errorf("notification email: list: %w", err)
	}
	return items, nil
}

// UnreadCount conta não-lidos (badge).
func (s *NotificationEmailService) UnreadCount() (int64, error) {
	var n int64
	if err := s.db.Model(&models.NotificationEmail{}).Where("is_read = ?", false).Count(&n).Error; err != nil {
		return 0, fmt.Errorf("notification email: unread count: %w", err)
	}
	return n, nil
}

// MarkRead marca um e-mail automático como lido.
func (s *NotificationEmailService) MarkRead(id uuid.UUID) error {
	if err := s.db.Model(&models.NotificationEmail{}).
		Where("id = ?", id).
		Updates(map[string]any{"is_read": true, "updated_at": time.Now().UTC()}).Error; err != nil {
		return fmt.Errorf("notification email: mark read: %w", err)
	}
	return nil
}
