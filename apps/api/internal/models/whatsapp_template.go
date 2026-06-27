package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// WhatsAppTemplateVar — metadado NOSSO de uma variável {{n}} do template (rótulo + exemplo),
// pra UI e pra documentar como cabear cada parâmetro.
type WhatsAppTemplateVar struct {
	Index   int    `json:"index"`
	Label   string `json:"label"`
	Example string `json:"example"`
}

// WhatsAppTemplate — catálogo de templates WhatsApp no EMR. Campos de conteúdo/status são CACHE
// sincronizado da Meta (fonte da verdade); os demais (purpose/variables/wiringNotes/enabled) são
// nossos e preservados no sync. Ver services/whatsapp_template_service.go.
type WhatsAppTemplate struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	MetaID   string    `gorm:"type:varchar(64)" json:"metaId"`
	Name     string    `gorm:"type:varchar(120);not null;uniqueIndex:uq_whatsapp_templates_name_lang" json:"name"`
	Language string    `gorm:"type:varchar(12);not null;default:'pt_BR';uniqueIndex:uq_whatsapp_templates_name_lang" json:"language"`
	Category string    `gorm:"type:varchar(20)" json:"category"`
	Status   string    `gorm:"type:varchar(20);not null;default:'UNKNOWN'" json:"status"`
	BodyText string    `gorm:"type:text" json:"bodyText"`

	VariableCount int                   `gorm:"not null;default:0" json:"variableCount"`
	Variables     []WhatsAppTemplateVar `gorm:"type:jsonb;serializer:json" json:"variables"`
	Purpose       string                `gorm:"type:text" json:"purpose"`
	WiringNotes   string                `gorm:"type:text" json:"wiringNotes"`
	Enabled       bool                  `gorm:"not null;default:true" json:"enabled"`
	// ChatSelectable: atendente pode escolher este template no compositor de conversas. Independe
	// de Enabled (que é "o sistema pode enviar"). false p/ internos/automáticos.
	ChatSelectable bool `gorm:"not null;default:true" json:"chatSelectable"`

	LastSyncedAt *time.Time `gorm:"type:timestamptz" json:"lastSyncedAt,omitempty"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
}

// TableName fixa o nome (o GORM pluralizaria "WhatsAppTemplate" como "whats_app_templates").
func (WhatsAppTemplate) TableName() string { return "whatsapp_templates" }

func (t *WhatsAppTemplate) BeforeCreate(tx *gorm.DB) error {
	if t.ID == uuid.Nil {
		t.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
