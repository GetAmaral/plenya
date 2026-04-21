package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AnonymousScoreGroupResult agrega o score de um ScoreGroup dentro de um AnonymousScoreSnapshot.
//
// @Description Resultado por grupo do Escore Light (radar)
type AnonymousScoreGroupResult struct {
	// @example 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	SnapshotID uuid.UUID `gorm:"type:uuid;not null;index:idx_anon_group_snapshot" json:"snapshotId" validate:"required"`
	GroupID    uuid.UUID `gorm:"type:uuid;not null;index:idx_anon_group_group" json:"groupId" validate:"required"`

	ActualPoints        float64 `gorm:"type:double precision;not null;default:0" json:"actualPoints" validate:"gte=0"`
	PossiblePoints      float64 `gorm:"type:double precision;not null;default:0" json:"possiblePoints" validate:"gte=0"`
	ScorePercentage     float64 `gorm:"type:double precision;not null;default:0" json:"scorePercentage" validate:"gte=0,lte=100"`
	ItemsEvaluatedCount int     `gorm:"type:integer;not null;default:0" json:"itemsEvaluatedCount" validate:"gte=0"`

	// Timestamps
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Snapshot *AnonymousScoreSnapshot `gorm:"foreignKey:SnapshotID;constraint:OnDelete:CASCADE" json:"-"`
	Group    *ScoreGroup             `gorm:"foreignKey:GroupID;constraint:OnDelete:RESTRICT" json:"group,omitempty"`
}

// TableName specifies the table name
func (AnonymousScoreGroupResult) TableName() string {
	return "anonymous_score_group_results"
}

// BeforeCreate hook generates UUID v7
func (r *AnonymousScoreGroupResult) BeforeCreate(tx *gorm.DB) error {
	if r.ID == uuid.Nil {
		r.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
