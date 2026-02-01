package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ScoreGroup represents a top-level category for clinical scores (e.g., "Hemograma", "Lipídeos")
// @Description Grupo de escores clínicos - categoria principal
type ScoreGroup struct {
	// @example 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// @minLength 2
	// @maxLength 200
	// @example Hemograma Completo
	Name string `gorm:"type:varchar(200);not null;uniqueIndex:idx_score_group_name" json:"name" validate:"required,min=2,max=200"`

	// @minimum 0
	// @maximum 9999
	// @example 1
	Order int `gorm:"type:integer;not null;default:0;index:idx_score_group_order" json:"order" validate:"gte=0,lte=9999"`

	// Relationships
	Subgroups []ScoreSubgroup `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE" json:"subgroups,omitempty"`

	// Timestamps
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for ScoreGroup
func (ScoreGroup) TableName() string {
	return "score_groups"
}

// BeforeCreate hook to generate UUID v7
func (sg *ScoreGroup) BeforeCreate(tx *gorm.DB) error {
	if sg.ID == uuid.Nil {
		sg.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
