package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ScoreVersion — uma versão PÚBLICA do escore (ex.: "Triagem", "Light"). Seleciona e ordena
// um subconjunto de ScoreItems (via ScoreVersionItem). Grupos/subgrupos/pilares vêm do próprio
// ScoreItem — a version NÃO duplica taxonomia, só escolhe e ordena. Fonte do gerador de configs
// estáticas do site. Substitui o flag is_light_version (removido na próxima migration).
// @Description Versão pública do escore (subset de itens, ordenado)
type ScoreVersion struct {
	// @example 5c012e51-0000-7000-8000-000000000002
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// @minLength 2
	// @maxLength 120
	// @example Triagem
	Name string `gorm:"type:varchar(120);not null" json:"name" validate:"required,min=2,max=120"`

	// Identificador estável usado no site (arquivo <slug>-config.json e rota /score-versions/:slug)
	// @example triagem
	Slug string `gorm:"type:varchar(60);not null;uniqueIndex:idx_score_version_slug" json:"slug" validate:"required,min=2,max=60"`

	// @example Versão pública e gratuita do Escore Plenya.
	Description *string `gorm:"type:text" json:"description,omitempty"`

	// Texto de introdução exibido na página do site desta versão.
	SiteIntro *string `gorm:"type:text" json:"siteIntro,omitempty"`

	// @minimum 0
	// @maximum 9999
	// @example 1
	Order int `gorm:"type:integer;not null;default:0;index:idx_score_version_order" json:"order" validate:"gte=0,lte=9999"`

	// @example true
	Active bool `gorm:"type:boolean;not null;default:true" json:"active"`

	// Relationships
	Items []ScoreVersionItem `gorm:"foreignKey:VersionID;constraint:OnDelete:CASCADE" json:"items,omitempty"`

	// Timestamps
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for ScoreVersion
func (ScoreVersion) TableName() string {
	return "score_versions"
}

// BeforeCreate hook to generate UUID v7
func (sv *ScoreVersion) BeforeCreate(tx *gorm.DB) error {
	if sv.ID == uuid.Nil {
		sv.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
