package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AnamnesisItem represents a single item/field in an anamnesis record
// @Description Item de anamnese vinculado a um ScoreItem (parâmetro clínico)
type AnamnesisItem struct {
	// @example 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`

	// Foreign key to Anamnesis
	// @example 550e8400-e29b-41d4-a716-446655440000
	AnamnesisID uuid.UUID `gorm:"type:uuid;not null;index:idx_anamnesis_item_anamnesis" json:"anamnesisId" validate:"required"`

	// Foreign key to ScoreItem (defines what is being recorded)
	// @example 550e8400-e29b-41d4-a716-446655440000
	ScoreItemID uuid.UUID `gorm:"type:uuid;not null;index:idx_anamnesis_item_score_item" json:"scoreItemId" validate:"required"`

	// Text value for qualitative/descriptive data
	// @example Paciente relata dor torácica intensa há 2 horas
	TextValue *string `gorm:"type:text" json:"textValue,omitempty"`

	// Numeric value for quantitative data
	// @example 120.5
	NumericValue *float64 `gorm:"type:double precision" json:"numericValue,omitempty" validate:"omitempty"`

	// Order for display (within the anamnesis)
	// @minimum 0
	// @maximum 9999
	// @example 1
	Order int `gorm:"type:integer;not null;default:0;index:idx_anamnesis_item_order" json:"order" validate:"gte=0,lte=9999"`

	// Relationships
	Anamnesis *Anamnesis `gorm:"foreignKey:AnamnesisID;constraint:OnDelete:CASCADE" json:"anamnesis,omitempty"`
	ScoreItem *ScoreItem `gorm:"foreignKey:ScoreItemID;constraint:OnDelete:RESTRICT" json:"scoreItem,omitempty"`

	// Timestamps
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for AnamnesisItem
func (AnamnesisItem) TableName() string {
	return "anamnesis_items"
}

// BeforeCreate hook to generate UUID v7
func (ai *AnamnesisItem) BeforeCreate(tx *gorm.DB) error {
	if ai.ID == uuid.Nil {
		ai.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
