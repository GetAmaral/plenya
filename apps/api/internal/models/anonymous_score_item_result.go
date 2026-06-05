package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AnonymousScoreItemResult — resultado da avaliação de um item Light numa sessão anônima.
// Análogo mínimo a PatientScoreItemResult (sem PatientID/DataSource/Lab/Anamnesis): o
// suficiente pra renderizar o radar por pilar (buildAgir) via Item.MethodPillars. Persistido
// apenas para itens efetivamente avaliados (status=evaluated). Fase 3 de
// docs/emr/plano-score-versions-unificacao.md.
type AnonymousScoreItemResult struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Snapshot anônimo a que pertence
	SnapshotID uuid.UUID `gorm:"type:uuid;not null;index:idx_anon_item_result_snapshot" json:"snapshotId" validate:"required"`

	// Item de score avaliado
	ItemID uuid.UUID `gorm:"type:uuid;not null;index:idx_anon_item_result_item" json:"itemId" validate:"required"`

	// Grupo do item (denormalizado)
	GroupID uuid.UUID `gorm:"type:uuid;not null;index:idx_anon_item_result_group" json:"groupId" validate:"required"`

	// Status da avaliação
	Status EvaluationStatus `gorm:"type:varchar(30);not null" json:"status" validate:"required,oneof=evaluated not_applicable no_data_available"`

	// Nível atingido (0-6, denormalizado) + FK opcional ao ScoreLevel
	LevelNumber    *int       `gorm:"type:integer" json:"levelNumber,omitempty" validate:"omitempty,gte=0,lte=6"`
	LevelMatchedID *uuid.UUID `gorm:"type:uuid" json:"levelMatchedId,omitempty"`

	// Valor numérico usado (quando a resposta foi numérica)
	ValueUsed *float64 `gorm:"type:double precision" json:"valueUsed,omitempty"`

	// Pontuação
	MaxPoints    float64 `gorm:"type:double precision;not null;default:0" json:"maxPoints" validate:"gte=0"`
	ActualPoints float64 `gorm:"type:double precision;not null;default:0" json:"actualPoints" validate:"gte=0"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Snapshot *AnonymousScoreSnapshot `gorm:"foreignKey:SnapshotID;constraint:OnDelete:CASCADE" json:"-"`
	Item     *ScoreItem              `gorm:"foreignKey:ItemID" json:"item,omitempty"`
}

// TableName specifies the table name
func (AnonymousScoreItemResult) TableName() string {
	return "anonymous_score_item_results"
}

// BeforeCreate hook generates UUID v7
func (r *AnonymousScoreItemResult) BeforeCreate(tx *gorm.DB) error {
	if r.ID == uuid.Nil {
		r.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
