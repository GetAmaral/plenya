package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ConversationOwnerType identifica o tipo de dono de conversa para ConversationRead.
// Espelha a enumeração interna usada nos handlers (/conversations/:type/...).
type ConversationOwnerType string

const (
	ConversationOwnerLead    ConversationOwnerType = "lead"
	ConversationOwnerPatient ConversationOwnerType = "patient"
)

// ConversationRead registra a última posição de leitura de um usuário (atendente)
// numa conversa específica (Lead ou Patient).
//
// Modelo:
//   - 1 linha por (UserID, OwnerType, OwnerID) — UNIQUE composto.
//   - LastReadAt = timestamp em que o usuário "viu" a conversa pela última vez.
//   - O cálculo de unread_count na lista é: COUNT(activities WHERE created_at > LastReadAt
//     AND actor_user_id IS NULL) — só inbound conta como não-lido.
//
// Não confundir com Notification.Read (que é por notificação avulsa). Aqui rastreamos
// um cursor por conversa, agnóstico ao número de mensagens.
//
// @Description Marcador de leitura por (usuário, conversa) — usado pra unread badges
type ConversationRead struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Usuário (atendente) que leu a conversa.
	UserID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_conv_reads_user_owner,priority:1" json:"userId"`

	// Tipo do dono da conversa: "lead" | "patient".
	OwnerType string `gorm:"type:varchar(20);not null;uniqueIndex:idx_conv_reads_user_owner,priority:2" json:"ownerType"`

	// ID do dono (Lead.ID ou Patient.ID).
	OwnerID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_conv_reads_user_owner,priority:3" json:"ownerId"`

	// Última posição de leitura — todas activities anteriores estão lidas.
	LastReadAt time.Time `gorm:"not null" json:"lastReadAt"`

	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

func (ConversationRead) TableName() string {
	return "conversation_reads"
}

// BeforeCreate hook gera UUID v7
func (c *ConversationRead) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.Must(uuid.NewV7())
	}
	if c.LastReadAt.IsZero() {
		c.LastReadAt = time.Now().UTC()
	}
	return nil
}
