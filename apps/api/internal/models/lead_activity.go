package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// LeadActivityType categoriza eventos no histórico do lead.
type LeadActivityType string

const (
	LeadActivityCreated         LeadActivityType = "created"
	LeadActivityMessageSent     LeadActivityType = "message_sent"     // outbound (Plenya → cliente)
	LeadActivityMessageReceived LeadActivityType = "message_received" // inbound (cliente → Plenya)
	LeadActivityStatusChanged   LeadActivityType = "status_changed"
	LeadActivityNoteAdded       LeadActivityType = "note_added"
	LeadActivityConverted       LeadActivityType = "converted"
	LeadActivityAssigned        LeadActivityType = "assigned"
	LeadActivityUnsubscribed    LeadActivityType = "unsubscribed"
)

// LeadActivityChannel indica por qual canal a atividade ocorreu.
type LeadActivityChannel string

const (
	LeadChannelEmail    LeadActivityChannel = "email"
	LeadChannelWhatsApp LeadActivityChannel = "whatsapp"
	LeadChannelInternal LeadActivityChannel = "internal" // notas, mudanças de status feitas no admin
)

// LeadActivity é o log imutável de eventos relacionados a um Lead.
//
// @Description Evento no histórico de um lead (mensagem, mudança de status, nota)
type LeadActivity struct {
	// @example 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Lead a que esta atividade pertence
	LeadID uuid.UUID `gorm:"type:uuid;not null;index:idx_lead_activities_lead" json:"leadId" validate:"required"`

	// Tipo do evento
	// @enum created,message_sent,message_received,status_changed,note_added,converted,assigned,unsubscribed
	Type LeadActivityType `gorm:"type:varchar(40);not null;index:idx_lead_activities_type" json:"type" validate:"required"`

	// Canal (email, whatsapp, internal)
	// @enum email,whatsapp,internal
	Channel LeadActivityChannel `gorm:"type:varchar(20)" json:"channel,omitempty"`

	// Conteúdo livre — corpo da mensagem, nota, descrição da mudança
	Content *string `gorm:"type:text" json:"content,omitempty"`

	// Metadados estruturados (ex: { "template": "magic_link", "wa_message_id": "wamid.X..." })
	Metadata datatypes.JSON `gorm:"type:jsonb" json:"metadata,omitempty"`

	// Usuário responsável pela ação (null se inbound do cliente ou ação do sistema)
	ActorUserID *uuid.UUID `gorm:"type:uuid" json:"actorUserId,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime;index:idx_lead_activities_created" json:"createdAt"`

	// Relationships
	Lead  *Lead `gorm:"foreignKey:LeadID;constraint:OnDelete:CASCADE" json:"-"`
	Actor *User `gorm:"foreignKey:ActorUserID;constraint:OnDelete:SET NULL" json:"actor,omitempty"`
}

// TableName especifica o nome da tabela
func (LeadActivity) TableName() string {
	return "lead_activities"
}

// BeforeCreate hook gera UUID v7
func (a *LeadActivity) BeforeCreate(tx *gorm.DB) error {
	if a.ID == uuid.Nil {
		a.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
