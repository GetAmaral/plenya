package repository

import (
	"github.com/google/uuid"
	"github.com/plenya/api/internal/models"
	"gorm.io/gorm"
)

type NotificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
	return &NotificationRepository{db: db}
}

// GetByID busca uma notificação por ID
func (r *NotificationRepository) GetByID(id string) (*models.Notification, error) {
	var notification models.Notification
	if err := r.db.Preload("User").Preload("Patient").Preload("Subscription").Where("id = ?", id).First(&notification).Error; err != nil {
		return nil, err
	}
	return &notification, nil
}

// GetByUserID busca todas as notificações de um usuário
func (r *NotificationRepository) GetByUserID(userID string, limit int) ([]models.Notification, error) {
	var notifications []models.Notification
	query := r.db.Where("user_id = ?", userID).Order("created_at DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}

	if err := query.Find(&notifications).Error; err != nil {
		return nil, err
	}
	return notifications, nil
}

// GetUnreadByUserID busca notificações não lidas de um usuário
func (r *NotificationRepository) GetUnreadByUserID(userID string) ([]models.Notification, error) {
	var notifications []models.Notification
	if err := r.db.Where("user_id = ? AND read = ?", userID, false).Order("created_at DESC").Find(&notifications).Error; err != nil {
		return nil, err
	}
	return notifications, nil
}

// CountUnreadByUserID conta notificações não lidas de um usuário
func (r *NotificationRepository) CountUnreadByUserID(userID string) (int64, error) {
	var count int64
	if err := r.db.Model(&models.Notification{}).Where("user_id = ? AND read = ?", userID, false).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// Create cria uma nova notificação
func (r *NotificationRepository) Create(notification *models.Notification) error {
	return r.db.Create(notification).Error
}

// Update atualiza uma notificação
func (r *NotificationRepository) Update(notification *models.Notification) error {
	return r.db.Save(notification).Error
}

// MarkAsRead marca uma notificação como lida
func (r *NotificationRepository) MarkAsRead(id string) error {
	notification, err := r.GetByID(id)
	if err != nil {
		return err
	}
	return notification.MarkAsRead(r.db)
}

// MarkAllAsReadByUserID marca todas as notificações de um usuário como lidas
func (r *NotificationRepository) MarkAllAsReadByUserID(userID string) error {
	return r.db.Model(&models.Notification{}).
		Where("user_id = ? AND read = ?", userID, false).
		Updates(map[string]interface{}{
			"read":    true,
			"read_at": gorm.Expr("NOW()"),
		}).Error
}

// MarkReadByTypes marca como lidas as notificações não lidas de um usuário cujos tipos estão
// na lista (ex.: ao entrar na tela de WhatsApp, marca todas as de WhatsApp inbound).
func (r *NotificationRepository) MarkReadByTypes(userID string, types []string) error {
	if len(types) == 0 {
		return nil
	}
	return r.db.Model(&models.Notification{}).
		Where("user_id = ? AND read = ? AND type IN ?", userID, false, types).
		Updates(map[string]interface{}{
			"read":    true,
			"read_at": gorm.Expr("NOW()"),
		}).Error
}

// FindLatestConversationNotification busca a notificação mais recente (lida ou não) de um
// usuário para um owner (lead OU patient) e tipo específico. Usada para colapsar as
// notificações de conversa: em vez de criar uma por mensagem, atualizamos a existente do
// mesmo contato/canal. Retorna (nil, nil) quando não há.
func (r *NotificationRepository) FindLatestConversationNotification(
	userID uuid.UUID, leadID, patientID *uuid.UUID, notifType models.NotificationType,
) (*models.Notification, error) {
	q := r.db.Where("user_id = ? AND type = ?", userID, notifType)
	switch {
	case leadID != nil && *leadID != uuid.Nil:
		q = q.Where("lead_id = ?", *leadID)
	case patientID != nil && *patientID != uuid.Nil:
		q = q.Where("patient_id = ?", *patientID)
	default:
		return nil, nil
	}
	var n models.Notification
	err := q.Order("created_at DESC").First(&n).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &n, nil
}

// Delete soft delete de uma notificação
func (r *NotificationRepository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&models.Notification{}).Error
}

// DeleteAllByUserID deleta todas as notificações de um usuário
func (r *NotificationRepository) DeleteAllByUserID(userID string) error {
	return r.db.Where("user_id = ?", userID).Delete(&models.Notification{}).Error
}
