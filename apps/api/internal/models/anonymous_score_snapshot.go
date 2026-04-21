package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AnonymousScoreSnapshot é o resultado calculado (radar) de uma AnonymousScoreSession.
// Análogo a PatientScoreSnapshot, mas sem PatientID/CalculatedByUserID.
//
// @Description Snapshot do Escore Light (resultado calculado)
type AnonymousScoreSnapshot struct {
	// @example 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Sessão a que pertence
	// @example 550e8400-e29b-41d4-a716-446655440000
	SessionID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_anon_snapshot_session" json:"sessionId" validate:"required"`

	// Pontos reais obtidos
	TotalActualPoints float64 `gorm:"type:double precision;not null;default:0" json:"totalActualPoints" validate:"gte=0"`

	// Pontos máximos possíveis (apenas dos itens avaliados)
	TotalPossiblePoints float64 `gorm:"type:double precision;not null;default:0" json:"totalPossiblePoints" validate:"gte=0"`

	// Score final em percentual
	TotalScorePercentage float64 `gorm:"type:double precision;not null;default:0" json:"totalScorePercentage" validate:"gte=0,lte=100"`

	// Quantidade de itens efetivamente avaliados
	ItemsEvaluatedCount int `gorm:"type:integer;not null;default:0" json:"itemsEvaluatedCount" validate:"gte=0"`

	// Quantidade de itens que não puderam ser avaliados (sem dado ou não aplicável)
	ItemsNotEvaluatedCount int `gorm:"type:integer;not null;default:0" json:"itemsNotEvaluatedCount" validate:"gte=0"`

	// Timestamps
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Session      *AnonymousScoreSession           `gorm:"foreignKey:SessionID;constraint:OnDelete:CASCADE" json:"-"`
	GroupResults []AnonymousScoreGroupResult      `gorm:"foreignKey:SnapshotID;constraint:OnDelete:CASCADE" json:"groupResults,omitempty"`
}

// TableName specifies the table name
func (AnonymousScoreSnapshot) TableName() string {
	return "anonymous_score_snapshots"
}

// BeforeCreate hook generates UUID v7
func (s *AnonymousScoreSnapshot) BeforeCreate(tx *gorm.DB) error {
	if s.ID == uuid.Nil {
		s.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
