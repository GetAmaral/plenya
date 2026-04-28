package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// RefreshToken — tabela de tokens de refresh emitidos. Armazena APENAS o hash
// SHA-256 do token (nunca o token plano), pra que vazamento da tabela não
// permita uso direto. Refresh-token rotation: cada uso emite um novo refresh
// e revoga o anterior (set RevokedAt=now).
//
// Tipos suportados:
//   - "refresh"      — refresh token padrão (login/oauth)
//   - "magic_link"   — magic link do Score Light (single-use, jti tracking)
//
// Logout revoga o(s) refresh token(s) ativos do user.
type RefreshToken struct {
	ID         uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	UserID     uuid.UUID  `gorm:"type:uuid;not null;index" json:"userId"`
	TokenHash  string     `gorm:"type:varchar(64);not null;uniqueIndex" json:"-"` // SHA-256 hex
	Type       string     `gorm:"type:varchar(20);not null;default:'refresh';index" json:"type"`
	ExpiresAt  time.Time  `gorm:"not null;index" json:"expiresAt"`
	RevokedAt  *time.Time `gorm:"index" json:"revokedAt,omitempty"`
	UsedAt     *time.Time `gorm:"index" json:"usedAt,omitempty"` // pra magic_link single-use
	UserAgent  *string    `gorm:"type:text" json:"userAgent,omitempty"`
	IPAddress  *string    `gorm:"type:varchar(45)" json:"ipAddress,omitempty"`
	CreatedAt  time.Time  `gorm:"not null;autoCreateTime" json:"createdAt"`

	// Relação opcional pra audit/debug
	User *User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
}

func (RefreshToken) TableName() string { return "refresh_tokens" }

func (r *RefreshToken) BeforeCreate(tx *gorm.DB) error {
	if r.ID == uuid.Nil {
		r.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// IsActive — token usável (não revogado, não usado se single-use, não expirado).
func (r *RefreshToken) IsActive() bool {
	now := time.Now().UTC()
	if r.RevokedAt != nil {
		return false
	}
	if r.UsedAt != nil {
		return false
	}
	if r.ExpiresAt.Before(now) {
		return false
	}
	return true
}
