package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ScoreVersionItem — M2M entre ScoreVersion e ScoreItem, com ordem de exibição NA version.
// O item é a fonte de verdade (grupo/subgrupo/pilares/levels/campos de site); aqui só registramos
// que ele pertence a esta version e em que ordem aparece.
// @Description Item de uma versão de escore (com ordem de exibição)
type ScoreVersionItem struct {
	// @example 5c012e51-1000-7000-8000-000000000001
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// @example 5c012e51-0000-7000-8000-000000000002
	VersionID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_score_version_item_uniq,priority:1;index:idx_score_version_item_version" json:"versionId" validate:"required"`

	// @example 550e8400-e29b-41d4-a716-446655440000
	ScoreItemID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_score_version_item_uniq,priority:2;index:idx_score_version_item_item" json:"scoreItemId" validate:"required"`

	// Ordem de exibição do item dentro desta version (independente de Order/LightOrder do item).
	// @minimum 0
	// @maximum 9999
	// @example 1
	DisplayOrder int `gorm:"type:integer;not null;default:0;index:idx_score_version_item_order" json:"displayOrder" validate:"gte=0,lte=9999"`

	// Relationships
	Version   *ScoreVersion `gorm:"foreignKey:VersionID;constraint:OnDelete:CASCADE" json:"-"`
	ScoreItem *ScoreItem    `gorm:"foreignKey:ScoreItemID;constraint:OnDelete:CASCADE" json:"scoreItem,omitempty"`

	// Timestamps — tabela de junção SEM soft-delete (hard-delete na remoção; índice único limpo).
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

// TableName specifies the table name for ScoreVersionItem
func (ScoreVersionItem) TableName() string {
	return "score_version_items"
}

// BeforeCreate hook to generate UUID v7
func (svi *ScoreVersionItem) BeforeCreate(tx *gorm.DB) error {
	if svi.ID == uuid.Nil {
		svi.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
