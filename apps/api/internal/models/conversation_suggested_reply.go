package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ConversationSuggestedReply guarda o rascunho ATUAL gerado pela IA para uma conversa
// (lead/patient), produzido pelo job após o debounce X. No Copiloto é mostrado no compositor
// e aguarda envio humano; no Auto vira envio (com preview X/3 quando a conversa está aberta).
// Um por conversa (upsert por owner). Plano: docs/emr/plano-atendimento-ia-global.md.
//
// @Description Rascunho de resposta da IA para uma conversa
type ConversationSuggestedReply struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	OwnerType string    `gorm:"type:varchar(10);not null;uniqueIndex:uq_suggested_reply_owner,priority:1" json:"ownerType"`
	OwnerID   uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:uq_suggested_reply_owner,priority:2" json:"ownerId"`

	// Corpo do rascunho (já sanitizado na voz Plenya).
	Reply string `gorm:"type:text;not null" json:"reply"`

	// Ação sugerida pelo cérebro: ask|answer|handle_objection|propose_schedule|handoff.
	Action string `gorm:"type:varchar(20);not null;default:'answer'" json:"action"`

	// Modelo que gerou (telemetria).
	Model string `gorm:"type:varchar(60)" json:"model,omitempty"`

	// Activity (inbound) que este rascunho responde — para invalidar quando chega msg nova.
	BasedOnActivityID *uuid.UUID `gorm:"type:uuid" json:"basedOnActivityId,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

// TableName specifies the table name
func (ConversationSuggestedReply) TableName() string { return "conversation_suggested_replies" }

// BeforeCreate hook generates UUID v7
func (r *ConversationSuggestedReply) BeforeCreate(tx *gorm.DB) error {
	if r.ID == uuid.Nil {
		r.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
