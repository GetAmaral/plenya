package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// NotificationType define os tipos de notificação
type NotificationType string

const (
	NotificationTrialExpiring       NotificationType = "trial_expiring"        // Trial expirando em breve
	NotificationRenewalUpcoming     NotificationType = "renewal_upcoming"      // Renovação próxima
	NotificationSubscriptionExpired NotificationType = "subscription_expired"  // Subscription expirado
	NotificationPaymentPending      NotificationType = "payment_pending"       // Pagamento pendente
	NotificationGeneral             NotificationType = "general"               // Notificação geral
)

// Notification representa uma notificação do sistema
// @Description Notificação para usuário
type Notification struct {
	// Primary Key
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// ID do usuário que receberá a notificação
	// @example 550e8400-e29b-41d4-a716-446655440000
	UserID uuid.UUID `gorm:"type:uuid;not null;index:idx_notifications_user" json:"userId" validate:"required"`

	// ID do paciente relacionado (opcional)
	// @example 550e8400-e29b-41d4-a716-446655440000
	PatientID *uuid.UUID `gorm:"type:uuid;index:idx_notifications_patient" json:"patientId,omitempty"`

	// ID do subscription relacionado (opcional)
	// @example 550e8400-e29b-41d4-a716-446655440000
	SubscriptionID *uuid.UUID `gorm:"type:uuid;index:idx_notifications_subscription" json:"subscriptionId,omitempty"`

	// Tipo da notificação
	// @enum trial_expiring,renewal_upcoming,subscription_expired,payment_pending,general
	// @example trial_expiring
	Type NotificationType `gorm:"type:varchar(50);not null;index:idx_notifications_type;check:type IN ('trial_expiring','renewal_upcoming','subscription_expired','payment_pending','general')" json:"type" validate:"required,oneof=trial_expiring renewal_upcoming subscription_expired payment_pending general"`

	// Título da notificação
	// @minLength 1
	// @maxLength 200
	// @example Trial expirando em 7 dias
	Title string `gorm:"type:varchar(200);not null" json:"title" validate:"required,min=1,max=200"`

	// Mensagem da notificação
	// @example Seu período trial expira em 7 dias. Renove agora para continuar aproveitando os benefícios.
	Message string `gorm:"type:text;not null" json:"message" validate:"required"`

	// URL de ação (opcional)
	// @example /patients/123/subscriptions
	ActionURL *string `gorm:"type:varchar(500)" json:"actionUrl,omitempty"`

	// Texto do botão de ação (opcional)
	// @example Ver Subscription
	ActionText *string `gorm:"type:varchar(50)" json:"actionText,omitempty"`

	// Se foi lida
	// @example false
	Read bool `gorm:"type:boolean;not null;default:false;index:idx_notifications_read" json:"read"`

	// Data de leitura
	ReadAt *time.Time `gorm:"type:timestamp" json:"readAt,omitempty"`

	// Timestamps
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index:idx_notifications_deleted_at" json:"-"`

	// Relationships
	User         User                  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
	Patient      *Patient              `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE" json:"patient,omitempty"`
	Subscription *PatientSubscription  `gorm:"foreignKey:SubscriptionID;constraint:OnDelete:CASCADE" json:"subscription,omitempty"`
}

func (Notification) TableName() string {
	return "notifications"
}

// BeforeCreate hook - gera UUID v7
func (n *Notification) BeforeCreate(tx *gorm.DB) error {
	if n.ID == uuid.Nil {
		n.ID = uuid.Must(uuid.NewV7())
	}
	return n.validate()
}

// BeforeSave hook - validações
func (n *Notification) BeforeSave(tx *gorm.DB) error {
	return n.validate()
}

// validate valida os campos
func (n *Notification) validate() error {
	// Validar tipo
	validTypes := map[NotificationType]bool{
		NotificationTrialExpiring:       true,
		NotificationRenewalUpcoming:     true,
		NotificationSubscriptionExpired: true,
		NotificationPaymentPending:      true,
		NotificationGeneral:             true,
	}
	if !validTypes[n.Type] {
		return gorm.ErrInvalidData
	}

	return nil
}

// MarkAsRead marca a notificação como lida
func (n *Notification) MarkAsRead(tx *gorm.DB) error {
	now := time.Now()
	n.Read = true
	n.ReadAt = &now
	return tx.Save(n).Error
}
