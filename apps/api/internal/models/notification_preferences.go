package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// NotificationPreferences governa quais pushes o usuário recebe.
// PK = UserID (1:1 com User). Default-on quando o registro não existir
// — endpoint PATCH cria on-demand.
//
// @Description Preferências de notificação push do usuário
type NotificationPreferences struct {
	// Usuário (PK)
	UserID uuid.UUID `gorm:"type:uuid;primaryKey" json:"userId"`

	// Lembrete de consulta (1h antes)
	AppointmentReminder bool `gorm:"not null;default:true" json:"appointmentReminder"`

	// Mensagem nova da clínica
	MessageAlert bool `gorm:"not null;default:true" json:"messageAlert"`

	// Lembrete diário de treino
	WorkoutReminder bool `gorm:"not null;default:true" json:"workoutReminder"`

	// Horário do lembrete de treino (HH:MM em string — TIME causa parse mess
	// entre Go/Postgres, string mantém simples e fuso-aware na app)
	WorkoutReminderTime string `gorm:"type:varchar(5);not null;default:'07:00'" json:"workoutReminderTime" validate:"omitempty,len=5"`

	// Timestamps
	CreatedAt time.Time      `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relação
	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
}

func (NotificationPreferences) TableName() string { return "notification_preferences" }

func (n *NotificationPreferences) BeforeCreate(tx *gorm.DB) error {
	// UserID é a PK — não geramos UUID aqui.
	return nil
}
