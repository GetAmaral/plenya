package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AnonymousScoreItem representa a resposta do paciente a uma única pergunta do Escore Light.
// Análogo a AnamnesisItem, mas vinculado a AnonymousScoreSession (não a Anamnesis).
//
// @Description Resposta a um item do Escore Light
type AnonymousScoreItem struct {
	// @example 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Sessão a que esta resposta pertence
	// @example 550e8400-e29b-41d4-a716-446655440000
	SessionID uuid.UUID `gorm:"type:uuid;not null;index:idx_anon_item_session" json:"sessionId" validate:"required"`

	// ID do ScoreItem ao qual esta resposta se refere
	// @example 550e8400-e29b-41d4-a716-446655440000
	ScoreItemID uuid.UUID `gorm:"type:uuid;not null;index:idx_anon_item_score" json:"scoreItemId" validate:"required"`

	// Valor numérico (para itens de medida objetiva ou níveis 0-6 de anamnese)
	// @example 4
	NumericValue *float64 `gorm:"type:double precision" json:"numericValue,omitempty"`

	// Valor textual (raramente usado — itens de free-text)
	TextValue *string `gorm:"type:text" json:"textValue,omitempty"`

	// Nível selecionado diretamente (0-6) — atalho para anamnese de múltipla escolha
	// @minimum 0
	// @maximum 6
	// @example 3
	SelectedLevel *int `gorm:"type:integer;check:selected_level >= 0 AND selected_level <= 6" json:"selectedLevel,omitempty" validate:"omitempty,gte=0,lte=6"`

	// Timestamps
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Session   *AnonymousScoreSession `gorm:"foreignKey:SessionID;constraint:OnDelete:CASCADE" json:"-"`
	ScoreItem *ScoreItem             `gorm:"foreignKey:ScoreItemID;constraint:OnDelete:RESTRICT" json:"scoreItem,omitempty"`
}

// TableName specifies the table name
func (AnonymousScoreItem) TableName() string {
	return "anonymous_score_items"
}

// BeforeCreate hook generates UUID v7
func (i *AnonymousScoreItem) BeforeCreate(tx *gorm.DB) error {
	if i.ID == uuid.Nil {
		i.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
