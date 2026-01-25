package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ScoreSubgroup represents a subcategory within a ScoreGroup (e.g., "Série Vermelha", "Série Branca")
// @Description Subgrupo de escores - subcategoria dentro de um grupo
type ScoreSubgroup struct {
	// @example 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`

	// @minLength 2
	// @maxLength 200
	// @example Série Vermelha
	Name string `gorm:"type:varchar(200);not null" json:"name" validate:"required,min=2,max=200"`

	// @minimum 0
	// @maximum 9999
	// @example 1
	Order int `gorm:"type:integer;not null;default:0;index:idx_score_subgroup_order" json:"order" validate:"gte=0,lte=9999"`

	// Foreign Keys
	// @example 550e8400-e29b-41d4-a716-446655440000
	GroupID uuid.UUID `gorm:"type:uuid;not null;index:idx_score_subgroup_group" json:"groupId" validate:"required"`

	// Relationships
	Group *ScoreGroup  `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE" json:"group,omitempty"`
	Items []ScoreItem  `gorm:"foreignKey:SubgroupID;constraint:OnDelete:CASCADE" json:"items,omitempty"`

	// Timestamps
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for ScoreSubgroup
func (ScoreSubgroup) TableName() string {
	return "score_subgroups"
}

// BeforeCreate hook to generate UUID v7
func (ss *ScoreSubgroup) BeforeCreate(tx *gorm.DB) error {
	if ss.ID == uuid.Nil {
		ss.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
