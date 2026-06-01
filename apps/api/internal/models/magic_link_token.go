package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// MagicLinkToken — registro single-use dos magic links do Score Light (claim
// anônimo por email/WhatsApp). Armazena APENAS o hash SHA-256 do jti, nunca o
// token plano.
//
// Tabela própria (sem FK de usuário) de propósito: o magic link é emitido no
// RequestClaim, ANTES de o User existir — o User/Patient só nascem no
// ConfirmClaim. Antes esse jti era gravado em refresh_tokens com UserID=uuid.Nil,
// o que violava a FK refresh_tokens→users (SQLSTATE 23503) e derrubava o envio.
type MagicLinkToken struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	TokenHash string     `gorm:"type:varchar(64);not null;uniqueIndex" json:"-"` // SHA-256 hex do jti
	ExpiresAt time.Time  `gorm:"not null;index" json:"expiresAt"`
	UsedAt    *time.Time `gorm:"index" json:"usedAt,omitempty"` // single-use
	CreatedAt time.Time  `gorm:"not null;autoCreateTime" json:"createdAt"`
}

func (MagicLinkToken) TableName() string { return "magic_link_tokens" }

func (m *MagicLinkToken) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// IsActive — token usável (não usado, não expirado).
func (m *MagicLinkToken) IsActive() bool {
	now := time.Now().UTC()
	if m.UsedAt != nil {
		return false
	}
	if m.ExpiresAt.Before(now) {
		return false
	}
	return true
}
