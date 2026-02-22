package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PatientScoreSnapshot represents a complete snapshot of a patient's health score at a specific point in time
// @Description Snapshot completo do escore de saúde do paciente
type PatientScoreSnapshot struct {
	// @example 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// ID do paciente
	// @example 550e8400-e29b-41d4-a716-446655440000
	PatientID uuid.UUID `gorm:"type:uuid;not null;index:idx_snapshot_patient" json:"patientId" validate:"required"`

	// ID do usuário que calculou o snapshot
	// @example 550e8400-e29b-41d4-a716-446655440000
	CalculatedByUserID uuid.UUID `gorm:"type:uuid;not null;index:idx_snapshot_user" json:"calculatedByUserId" validate:"required"`

	// Data e hora do cálculo
	// @example 2026-02-12T10:30:00Z
	CalculatedAt time.Time `gorm:"type:timestamp;not null;index:idx_snapshot_date" json:"calculatedAt" validate:"required"`

	// Pontos reais obtidos (soma de todos itens avaliados)
	// @minimum 0
	// @example 450.5
	TotalActualPoints float64 `gorm:"type:double precision;not null;default:0" json:"totalActualPoints" validate:"gte=0"`

	// Pontos máximos possíveis (soma apenas dos itens avaliados)
	// @minimum 0
	// @example 800.0
	TotalPossiblePoints float64 `gorm:"type:double precision;not null;default:0" json:"totalPossiblePoints" validate:"gte=0"`

	// Score final em percentual (actualPoints/possiblePoints × 100)
	// @minimum 0
	// @maximum 100
	// @example 56.31
	TotalScorePercentage float64 `gorm:"type:double precision;not null;default:0" json:"totalScorePercentage" validate:"gte=0,lte=100"`

	// Quantidade de itens que foram avaliados (tinham dados disponíveis e eram aplicáveis)
	// @minimum 0
	// @example 42
	ItemsEvaluatedCount int `gorm:"type:integer;not null;default:0" json:"itemsEvaluatedCount" validate:"gte=0"`

	// Quantidade de itens que não foram avaliados (sem dados ou não aplicáveis)
	// @minimum 0
	// @example 8
	ItemsNotEvaluatedCount int `gorm:"type:integer;not null;default:0" json:"itemsNotEvaluatedCount" validate:"gte=0"`

	// Observações opcionais sobre o cálculo
	// @example Primeira avaliação após diagnóstico
	Notes *string `gorm:"type:text" json:"notes,omitempty"`

	// Título computado para exibição no frontend (não persistido)
	DisplayTitle string `gorm:"-" json:"displayTitle"`

	// Timestamps
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Patient           *Patient                   `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE" json:"patient,omitempty"`
	CalculatedByUser  *User                      `gorm:"foreignKey:CalculatedByUserID;constraint:OnDelete:RESTRICT" json:"calculatedByUser,omitempty"`
	GroupResults      []PatientScoreGroupResult  `gorm:"foreignKey:SnapshotID;constraint:OnDelete:CASCADE" json:"groupResults,omitempty"`
	ItemResults       []PatientScoreItemResult   `gorm:"foreignKey:SnapshotID;constraint:OnDelete:CASCADE" json:"itemResults,omitempty"`
}

// TableName specifies the table name for PatientScoreSnapshot
func (PatientScoreSnapshot) TableName() string {
	return "patient_score_snapshots"
}

// GetTitle retorna um título legível para o snapshot
func (pss *PatientScoreSnapshot) GetTitle() string {
	return fmt.Sprintf("Snapshot %s - %.1f%%",
		pss.CalculatedAt.Format("02/01/2006 15:04"),
		pss.TotalScorePercentage)
}

// AfterFind popula DisplayTitle após carregar do banco
func (pss *PatientScoreSnapshot) AfterFind(tx *gorm.DB) error {
	pss.DisplayTitle = pss.GetTitle()
	return nil
}

// BeforeCreate hook to generate UUID v7
func (pss *PatientScoreSnapshot) BeforeCreate(tx *gorm.DB) error {
	if pss.ID == uuid.Nil {
		pss.ID = uuid.Must(uuid.NewV7())
	}
	// Set CalculatedAt to now if not set
	if pss.CalculatedAt.IsZero() {
		pss.CalculatedAt = time.Now()
	}
	return nil
}
