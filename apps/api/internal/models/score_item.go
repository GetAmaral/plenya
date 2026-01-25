package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ScoreItem represents a specific clinical parameter (e.g., "Hemoglobina - Homens", "FEVE")
// @Description Item de escore - parâmetro clínico específico
type ScoreItem struct {
	// @example 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`

	// @minLength 2
	// @maxLength 300
	// @example Hemoglobina - Homens
	Name string `gorm:"type:varchar(300);not null" json:"name" validate:"required,min=2,max=300"`

	// @example g/dL
	Unit *string `gorm:"type:varchar(50)" json:"unit,omitempty" validate:"omitempty,max=50"`

	// @example 1 g/dL = 10 g/L
	UnitConversion *string `gorm:"type:text" json:"unitConversion,omitempty"`

	// @minimum 0
	// @maximum 100
	// @example 20
	Points float64 `gorm:"type:double precision;not null;default:0" json:"points" validate:"gte=0,lte=100"`

	// @minimum 0
	// @maximum 9999
	// @example 1
	Order int `gorm:"type:integer;not null;default:0;index:idx_score_item_order" json:"order" validate:"gte=0,lte=9999"`

	// Foreign Keys
	// @example 550e8400-e29b-41d4-a716-446655440000
	SubgroupID uuid.UUID `gorm:"type:uuid;not null;index:idx_score_item_subgroup" json:"subgroupId" validate:"required"`

	// Self-referencing for hierarchical items (optional)
	// @example 550e8400-e29b-41d4-a716-446655440000
	ParentItemID *uuid.UUID `gorm:"type:uuid;index:idx_score_item_parent" json:"parentItemId,omitempty"`

	// Relationships
	Subgroup   *ScoreSubgroup `gorm:"foreignKey:SubgroupID;constraint:OnDelete:CASCADE" json:"subgroup,omitempty"`
	ParentItem *ScoreItem     `gorm:"foreignKey:ParentItemID;constraint:OnDelete:SET NULL" json:"parentItem,omitempty"`
	ChildItems []ScoreItem    `gorm:"foreignKey:ParentItemID;constraint:OnDelete:SET NULL" json:"childItems,omitempty"`
	Levels     []ScoreLevel   `gorm:"foreignKey:ItemID;constraint:OnDelete:CASCADE" json:"levels,omitempty"`

	// Timestamps
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for ScoreItem
func (ScoreItem) TableName() string {
	return "score_items"
}

// BeforeCreate hook to generate UUID v7
func (si *ScoreItem) BeforeCreate(tx *gorm.DB) error {
	if si.ID == uuid.Nil {
		si.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
