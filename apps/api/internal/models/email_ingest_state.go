package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// EmailIngestState rastreia o último UID processado por (account, folder).
// Permite que o worker IMAP retome de onde parou após restart sem reprocessar emails.
//
// @Description Estado de ingestão IMAP por conta + pasta (último UID processado)
type EmailIngestState struct {
	// @example 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Conta IMAP (ex: contato@plenyasaude.com.br)
	Account string `gorm:"type:varchar(255);not null;uniqueIndex:idx_eis_account_folder" json:"account"`

	// Pasta IMAP (ex: INBOX, Sent Items)
	Folder string `gorm:"type:varchar(100);not null;uniqueIndex:idx_eis_account_folder" json:"folder"`

	// Último UID processado nesta pasta (0 = nada processado ainda)
	LastUID uint32 `gorm:"not null;default:0" json:"lastUid"`

	// Última vez que o worker tocou neste registro (heartbeat)
	LastSeenAt time.Time `gorm:"autoUpdateTime" json:"lastSeenAt"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
}

// TableName especifica o nome da tabela
func (EmailIngestState) TableName() string {
	return "email_ingest_states"
}

// BeforeCreate gera UUID v7
func (e *EmailIngestState) BeforeCreate(tx *gorm.DB) error {
	if e.ID == uuid.Nil {
		e.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
