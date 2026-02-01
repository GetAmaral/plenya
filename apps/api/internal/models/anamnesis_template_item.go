package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AnamnesisTemplateItem represents a single item in an anamnesis template
// @Description Item de template de anamnese vinculado a um ScoreItem
type AnamnesisTemplateItem struct {
	// @example 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Foreign key to AnamnesisTemplate
	// @example 550e8400-e29b-41d4-a716-446655440000
	AnamnesisTemplateID uuid.UUID `gorm:"type:uuid;not null;index:idx_anamnesis_template_item_template" json:"anamnesisTemplateId" validate:"required"`

	// Foreign key to ScoreItem
	// @example 550e8400-e29b-41d4-a716-446655440000
	ScoreItemID uuid.UUID `gorm:"type:uuid;not null;index:idx_anamnesis_template_item_score_item" json:"scoreItemId" validate:"required"`

	// Order for display (within the template)
	// @minimum 0
	// @maximum 9999
	// @example 1
	Order int `gorm:"type:integer;not null;default:0;index:idx_anamnesis_template_item_order" json:"order" validate:"gte=0,lte=9999"`

	// Relationships
	AnamnesisTemplate *AnamnesisTemplate `gorm:"foreignKey:AnamnesisTemplateID;constraint:OnDelete:CASCADE" json:"anamnesisTemplate,omitempty"`
	ScoreItem         *ScoreItem         `gorm:"foreignKey:ScoreItemID;constraint:OnDelete:RESTRICT" json:"scoreItem,omitempty"`

	// Timestamps
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for AnamnesisTemplateItem
func (AnamnesisTemplateItem) TableName() string {
	return "anamnesis_template_items"
}

// BeforeCreate hook to generate UUID v7
func (ati *AnamnesisTemplateItem) BeforeCreate(tx *gorm.DB) error {
	if ati.ID == uuid.Nil {
		ati.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
