package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// EvaluationStatus defines the possible evaluation statuses for a score item
type EvaluationStatus string

const (
	// Item was successfully evaluated with available data
	EvaluationStatusEvaluated EvaluationStatus = "evaluated"
	// Item does not apply to this patient (e.g., gender/age filters)
	EvaluationStatusNotApplicable EvaluationStatus = "not_applicable"
	// No data available for this item in patient's history
	EvaluationStatusNoDataAvailable EvaluationStatus = "no_data_available"
)

// DataSource defines where the evaluation data came from
type DataSource string

const (
	DataSourceLabResult     DataSource = "lab_result"
	DataSourceAnamnesisItem DataSource = "anamnesis_item"
)

// PatientScoreItemResult represents the evaluation of a single ScoreItem for a patient
// @Description Resultado da avaliação individual de um item de score
type PatientScoreItemResult struct {
	// @example 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// ID do snapshot pai
	// @example 550e8400-e29b-41d4-a716-446655440000
	SnapshotID uuid.UUID `gorm:"type:uuid;not null;index:idx_item_result_snapshot" json:"snapshotId" validate:"required"`

	// ID do item de score avaliado
	// @example 550e8400-e29b-41d4-a716-446655440000
	ItemID uuid.UUID `gorm:"type:uuid;not null;index:idx_item_result_item" json:"itemId" validate:"required"`

	// ID do grupo (denormalizado para agregação rápida)
	// @example 550e8400-e29b-41d4-a716-446655440000
	GroupID uuid.UUID `gorm:"type:uuid;not null;index:idx_item_result_group" json:"groupId" validate:"required"`

	// Status da avaliação
	// @enum evaluated,not_applicable,no_data_available
	// @example evaluated
	Status EvaluationStatus `gorm:"type:varchar(30);not null;check:status IN ('evaluated','not_applicable','no_data_available')" json:"status" validate:"required,oneof=evaluated not_applicable no_data_available"`

	// Fonte dos dados usados na avaliação
	// @enum lab_result,anamnesis_item
	// @example lab_result
	DataSource *DataSource `gorm:"type:varchar(20);check:data_source IN ('lab_result','anamnesis_item')" json:"dataSource,omitempty" validate:"omitempty,oneof=lab_result anamnesis_item"`

	// ID do resultado de exame laboratorial usado (se aplicável)
	// @example 550e8400-e29b-41d4-a716-446655440000
	LabResultID *uuid.UUID `gorm:"type:uuid;index:idx_item_result_lab_result" json:"labResultId,omitempty"`

	// ID do item de anamnese usado (se aplicável)
	// @example 550e8400-e29b-41d4-a716-446655440000
	AnamnesisItemID *uuid.UUID `gorm:"type:uuid;index:idx_item_result_anamnesis_item" json:"anamnesisItemId,omitempty"`

	// Valor numérico usado na avaliação
	// @example 13.5
	ValueUsed *float64 `gorm:"type:double precision" json:"valueUsed,omitempty"`

	// ID do nível que foi atingido
	// @example 550e8400-e29b-41d4-a716-446655440000
	LevelMatchedID *uuid.UUID `gorm:"type:uuid;index:idx_item_result_level" json:"levelMatchedId,omitempty"`

	// Número do nível (0-6, denormalizado para acesso rápido)
	// @minimum 0
	// @maximum 6
	// @example 3
	LevelNumber *int `gorm:"type:integer;check:level_number >= 0 AND level_number <= 6" json:"levelNumber,omitempty" validate:"omitempty,gte=0,lte=6"`

	// Pontos máximos do item
	// @minimum 0
	// @maximum 100
	// @example 20.0
	MaxPoints float64 `gorm:"type:double precision;not null;default:0" json:"maxPoints" validate:"gte=0,lte=100"`

	// Pontos reais obtidos (proporcional ao nível atingido)
	// @minimum 0
	// @maximum 100
	// @example 12.0
	ActualPoints float64 `gorm:"type:double precision;not null;default:0" json:"actualPoints" validate:"gte=0,lte=100"`

	// Motivo da não avaliação (se status != evaluated)
	// @example Item não aplicável: sexo masculino
	NotEvaluatedReason *string `gorm:"type:text" json:"notEvaluatedReason,omitempty"`

	// Timestamps
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Snapshot     *PatientScoreSnapshot `gorm:"foreignKey:SnapshotID;constraint:OnDelete:CASCADE" json:"snapshot,omitempty"`
	Item         *ScoreItem            `gorm:"foreignKey:ItemID;constraint:OnDelete:RESTRICT" json:"item,omitempty"`
	Group        *ScoreGroup           `gorm:"foreignKey:GroupID;constraint:OnDelete:RESTRICT" json:"group,omitempty"`
	LevelMatched *ScoreLevel           `gorm:"foreignKey:LevelMatchedID;constraint:OnDelete:SET NULL" json:"levelMatched,omitempty"`
	LabResult    *LabResult            `gorm:"foreignKey:LabResultID;constraint:OnDelete:SET NULL" json:"labResult,omitempty"`
	AnamnesisItem *AnamnesisItem       `gorm:"foreignKey:AnamnesisItemID;constraint:OnDelete:SET NULL" json:"anamnesisItem,omitempty"`
}

// TableName specifies the table name for PatientScoreItemResult
func (PatientScoreItemResult) TableName() string {
	return "patient_score_item_results"
}

// BeforeCreate hook to generate UUID v7
func (psir *PatientScoreItemResult) BeforeCreate(tx *gorm.DB) error {
	if psir.ID == uuid.Nil {
		psir.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
