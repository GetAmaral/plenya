package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PatientScoreGroupResult represents the aggregated score for a ScoreGroup within a snapshot
// @Description Resultado agregado de score por grupo clínico
type PatientScoreGroupResult struct {
	// @example 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// ID do snapshot pai
	// @example 550e8400-e29b-41d4-a716-446655440000
	SnapshotID uuid.UUID `gorm:"type:uuid;not null;index:idx_group_result_snapshot" json:"snapshotId" validate:"required"`

	// ID do grupo de score
	// @example 550e8400-e29b-41d4-a716-446655440000
	GroupID uuid.UUID `gorm:"type:uuid;not null;index:idx_group_result_group" json:"groupId" validate:"required"`

	// Pontos reais obtidos no grupo
	// @minimum 0
	// @example 85.5
	ActualPoints float64 `gorm:"type:double precision;not null;default:0" json:"actualPoints" validate:"gte=0"`

	// Pontos máximos possíveis no grupo (apenas itens avaliados)
	// @minimum 0
	// @example 150.0
	PossiblePoints float64 `gorm:"type:double precision;not null;default:0" json:"possiblePoints" validate:"gte=0"`

	// Score do grupo em percentual
	// @minimum 0
	// @maximum 100
	// @example 57.0
	ScorePercentage float64 `gorm:"type:double precision;not null;default:0" json:"scorePercentage" validate:"gte=0,lte=100"`

	// Quantidade de itens avaliados neste grupo
	// @minimum 0
	// @example 12
	ItemsEvaluatedCount int `gorm:"type:integer;not null;default:0" json:"itemsEvaluatedCount" validate:"gte=0"`

	// Timestamps
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Snapshot *PatientScoreSnapshot `gorm:"foreignKey:SnapshotID;constraint:OnDelete:CASCADE" json:"snapshot,omitempty"`
	Group    *ScoreGroup           `gorm:"foreignKey:GroupID;constraint:OnDelete:RESTRICT" json:"group,omitempty"`
}

// TableName specifies the table name for PatientScoreGroupResult
func (PatientScoreGroupResult) TableName() string {
	return "patient_score_group_results"
}

// BeforeCreate hook to generate UUID v7
func (psgr *PatientScoreGroupResult) BeforeCreate(tx *gorm.DB) error {
	if psgr.ID == uuid.Nil {
		psgr.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
