package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ScoreLevel represents a risk stratification level for a ScoreItem (e.g., "Nível 0: <40% FEVE")
// @Description Nível de estratificação de risco para um item de escore
type ScoreLevel struct {
	// @example 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`

	// @minimum 0
	// @maximum 6
	// @example 5
	Level int `gorm:"type:integer;not null;check:level >= 0 AND level <= 6;index:idx_score_level_level" json:"level" validate:"required,gte=0,lte=6"`

	// @minLength 1
	// @maxLength 500
	// @example 55 a 70 (Ótimo)
	Name string `gorm:"type:varchar(500);not null" json:"name" validate:"required,min=1,max=500"`

	// @example <40
	LowerLimit *string `gorm:"type:varchar(50)" json:"lowerLimit,omitempty" validate:"omitempty,max=50"`

	// @example 49.9
	UpperLimit *string `gorm:"type:varchar(50)" json:"upperLimit,omitempty" validate:"omitempty,max=50"`

	// @enum =,>,>=,<,<=,between
	// @example between
	Operator string `gorm:"type:varchar(10);not null;check:operator IN ('=','>','>=','<','<=','between');default:'between'" json:"operator" validate:"required,oneof== > >= < <= between"`

	// Foreign Keys
	// @example 550e8400-e29b-41d4-a716-446655440000
	ItemID uuid.UUID `gorm:"type:uuid;not null;index:idx_score_level_item" json:"itemId" validate:"required"`

	// Relationships
	Item *ScoreItem `gorm:"foreignKey:ItemID;constraint:OnDelete:CASCADE" json:"item,omitempty"`

	// Timestamps
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for ScoreLevel
func (ScoreLevel) TableName() string {
	return "score_levels"
}

// BeforeCreate hook to generate UUID v7
func (sl *ScoreLevel) BeforeCreate(tx *gorm.DB) error {
	if sl.ID == uuid.Nil {
		sl.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
