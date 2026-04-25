package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TelemedLobbyToken é o token público colocado no link enviado por email/WhatsApp
// pra paciente entrar direto na sala Daily.co (sem precisar logar no portal).
//
// Gerado automaticamente quando appointment.type=telemedicine é criado.
// Expira 4h após o ScheduledAt (cobre atraso natural). Single-use opcional —
// V1 deixa ser usado múltiplas vezes (paciente pode dar reload na página).
//
// O token é distinto do JWT do portal: paciente que já está logado entra via
// /consultas/[id] (com auth); paciente vindo só do link cai aqui.
type TelemedLobbyToken struct {
	ID            uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	AppointmentID uuid.UUID  `gorm:"type:uuid;not null;index" json:"appointmentId"`
	Token         string     `gorm:"type:varchar(128);uniqueIndex;not null" json:"-"`
	ExpiresAt     time.Time  `gorm:"not null" json:"expiresAt"`
	UsedAt        *time.Time `json:"usedAt,omitempty"`
	CreatedAt     time.Time  `gorm:"autoCreateTime" json:"createdAt"`
}

func (TelemedLobbyToken) TableName() string {
	return "telemed_lobby_tokens"
}

func (t *TelemedLobbyToken) BeforeCreate(tx *gorm.DB) error {
	if t.ID == uuid.Nil {
		t.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

func (t *TelemedLobbyToken) IsValid() bool {
	return time.Now().UTC().Before(t.ExpiresAt)
}
