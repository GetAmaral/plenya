package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PatientMagicLink é o token enviado quando paciente já tem User+Patient
// e pede pra entrar via "magic link" ou "esqueci a senha". Single-use.
//
// Difere do magic link do Escore Light: aquele cria User+Patient
// no consume; este só autentica um User existente com role=patient.
//
// M3 — Token é deprecated: novos registros gravam apenas TokenHash
// (sha256 hex). Plain token só existe em memória durante o request que
// gera o link. Coluna Token mantida pra retrocompat com rows antigas
// (AutoMigrate não dropa coluna). Lookups novos sempre por TokenHash.
type PatientMagicLink struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	UserID    uuid.UUID  `gorm:"type:uuid;not null;index" json:"userId"`
	// Deprecated: M3 — vazia em registros novos. Use TokenHash.
	Token     string     `gorm:"type:varchar(128);uniqueIndex" json:"-"`
	TokenHash string     `gorm:"type:varchar(64);uniqueIndex:idx_patient_magic_links_hash" json:"-"`
	ExpiresAt time.Time  `gorm:"not null" json:"expiresAt"`
	UsedAt    *time.Time `json:"usedAt,omitempty"`
	IPHash    *string    `gorm:"type:varchar(64)" json:"-"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"createdAt"`
}

func (PatientMagicLink) TableName() string {
	return "patient_magic_links"
}

func (m *PatientMagicLink) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.Must(uuid.NewV7())
	}
	if m.ExpiresAt.IsZero() {
		m.ExpiresAt = time.Now().UTC().Add(15 * time.Minute)
	}
	return nil
}

// IsValid retorna true se ainda pode ser consumido.
func (m *PatientMagicLink) IsValid() bool {
	return m.UsedAt == nil && time.Now().UTC().Before(m.ExpiresAt)
}
