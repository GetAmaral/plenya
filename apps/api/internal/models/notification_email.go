package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/plenya/api/internal/crypto"
	"gorm.io/gorm"
)

// NotificationEmail é um e-mail automático (no-reply, newsletter, notificações de
// plataformas) que chega na caixa mas NÃO vira Lead. Fica no bucket "Notificações" da
// caixa de e-mail — acessível, fora do caminho dos leads/pacientes reais.
//
// Corpo cifrado em repouso (LGPD), mesmo padrão de LeadActivity. Assunto fica em claro
// pra preview da lista (e-mail automático tem baixa sensibilidade no assunto).
type NotificationEmail struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	FromEmail  string    `gorm:"type:varchar(320);index" json:"fromEmail"`
	FromName   *string   `gorm:"type:varchar(255)" json:"fromName,omitempty"`
	Subject    string    `gorm:"type:text;not null;default:''" json:"subject"`
	BodyText   string    `gorm:"type:text;not null;default:''" json:"bodyText"` // cifrado em repouso
	MessageID  *string   `gorm:"type:varchar(998);uniqueIndex:idx_notification_emails_message_id" json:"messageId,omitempty"`
	ReceivedAt time.Time `gorm:"type:timestamptz;index;not null;default:now()" json:"receivedAt"`
	IsRead     bool      `gorm:"not null;default:false;index" json:"isRead"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func (NotificationEmail) TableName() string { return "notification_emails" }

// BeforeCreate gera UUID v7 (ordenável por tempo, padrão do projeto).
func (n *NotificationEmail) BeforeCreate(tx *gorm.DB) error {
	if n.ID == uuid.Nil {
		n.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// BeforeSave cifra o corpo em repouso (idempotente via isEncrypted).
func (n *NotificationEmail) BeforeSave(tx *gorm.DB) error {
	if n.BodyText != "" && !isEncrypted(n.BodyText) {
		enc, err := crypto.EncryptWithDefaultKey(n.BodyText)
		if err != nil {
			return err
		}
		n.BodyText = enc
	}
	return nil
}

// AfterFind descriptografa o corpo. Falha silenciosa (conteúdo legado/plain não crasha).
func (n *NotificationEmail) AfterFind(tx *gorm.DB) error {
	if n.BodyText != "" && isEncrypted(n.BodyText) {
		if dec, err := crypto.DecryptWithDefaultKey(n.BodyText); err == nil {
			n.BodyText = dec
		}
	}
	return nil
}
