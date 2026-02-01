package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AnamnesisTemplate representa um template de anamnese
// @Description Template de anamnese com itens pré-configurados para uma área específica
type AnamnesisTemplate struct {
	// @example 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Nome do template
	// @minLength 2
	// @maxLength 200
	// @example Anamnese Clínica Geral
	Name string `gorm:"type:varchar(200);not null" json:"name" validate:"required,min=2,max=200"`

	// Área de aplicação do template
	// @example Clínica Médica, Psicologia, Pediatria, Cardiologia
	Area string `gorm:"type:varchar(100);not null;index:idx_anamnesis_template_area" json:"area" validate:"required,max=100"`

	// Timestamps
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Items []AnamnesisTemplateItem `gorm:"foreignKey:AnamnesisTemplateID;constraint:OnDelete:CASCADE" json:"items,omitempty"`
}

// TableName specifies the table name for AnamnesisTemplate
func (AnamnesisTemplate) TableName() string {
	return "anamnesis_templates"
}

// BeforeCreate hook to generate UUID v7
func (at *AnamnesisTemplate) BeforeCreate(tx *gorm.DB) error {
	if at.ID == uuid.Nil {
		at.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
