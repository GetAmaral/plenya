package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Modos de automação do recepcionista virtual por conversa.
const (
	ConversationAutomationOff     = "off"     // IA não age (nem sugere proativamente)
	ConversationAutomationCopilot = "copilot" // IA só sugere; humano envia
	ConversationAutomationAuto    = "auto"    // IA atende sozinha após o fallback de tempo
)

// ConversationAutomation guarda o modo do recepcionista virtual para uma conversa
// (lead ou patient). Linha existe só quando alguém define um modo; sem linha = default
// global (config RECEPTION_BOT_DEFAULT_MODE). Chave única por (owner_type, owner_id).
//
// @Description Configuração de automação (recepcionista virtual) de uma conversa
type ConversationAutomation struct {
	// @example 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Tipo do dono da conversa: "lead" ou "patient"
	OwnerType string `gorm:"type:varchar(10);not null;uniqueIndex:uq_conversation_automation_owner,priority:1" json:"ownerType"`
	// ID do lead ou patient
	OwnerID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:uq_conversation_automation_owner,priority:2" json:"ownerId"`

	// Modo: off | copilot | auto
	// @enum off,copilot,auto
	Mode string `gorm:"type:varchar(10);not null;default:'off';check:mode IN ('off','copilot','auto')" json:"mode"`

	// Minutos sem resposta humana antes de a IA assumir (modo auto). 0 = usar default global.
	FallbackMinutes int `gorm:"type:integer;not null;default:0" json:"fallbackMinutes"`

	// Pausa temporária (handoff): enquanto > now, a IA não age. Limpo ao reativar.
	PausedUntil *time.Time `gorm:"type:timestamptz" json:"pausedUntil,omitempty"`

	// Última vez que a IA enviou nesta conversa (anti-spam + telemetria).
	LastBotAt *time.Time `gorm:"type:timestamptz" json:"lastBotAt,omitempty"`

	// Quem alterou o modo por último (auditoria).
	UpdatedBy *uuid.UUID `gorm:"type:uuid" json:"updatedBy,omitempty"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name
func (ConversationAutomation) TableName() string {
	return "conversation_automations"
}

// BeforeCreate hook generates UUID v7
func (a *ConversationAutomation) BeforeCreate(tx *gorm.DB) error {
	if a.ID == uuid.Nil {
		a.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
