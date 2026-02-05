package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// LabResult representa um resultado de exame laboratorial individual
// @Description Resultado individual de um exame dentro de um lote (batch)
type LabResult struct {
	// ID único do resultado
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// ID do lote de resultados (obrigatório)
	LabResultBatchID uuid.UUID `gorm:"type:uuid;not null;index:idx_result_batch" json:"labResultBatchId" validate:"required"`

	// ID da definição do teste (opcional - FK para LabTestDefinition)
	LabTestDefinitionID *uuid.UUID `gorm:"type:uuid;index:idx_result_test_def" json:"labTestDefinitionId,omitempty"`

	// Nome do exame
	TestName string `gorm:"type:varchar(200);not null" json:"testName" validate:"required"`

	// Tipo de exame (ex: hemograma, glicemia, etc)
	TestType string `gorm:"type:varchar(100);not null;index" json:"testType" validate:"required"`

	// Resultado em texto (para exames qualitativos)
	ResultText *string `gorm:"type:text" json:"resultText,omitempty"`

	// Resultado numérico (para exames quantitativos)
	ResultNumeric *float64 `gorm:"type:decimal(12,4)" json:"resultNumeric,omitempty"`

	// Unidade de medida
	Unit *string `gorm:"type:varchar(50)" json:"unit,omitempty"`

	// Valores de referência
	ReferenceRange *string `gorm:"type:text" json:"referenceRange,omitempty"`

	// Interpretação/Observações específicas deste teste
	Interpretation *string `gorm:"type:text" json:"interpretation,omitempty"`

	// Nível/Prioridade do resultado (opcional)
	Level *int `gorm:"type:integer" json:"level,omitempty"`

	// Indica se o exame foi matched com uma definição catalogada
	// true = matched com LabTestDefinition, false = extraído mas não catalogado
	Matched bool `gorm:"type:boolean;not null;default:true;index" json:"matched"`

	// Data de criação
	CreatedAt time.Time `gorm:"not null;autoCreateTime" json:"createdAt"`

	// Data de atualização
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime" json:"updatedAt"`

	// Data de deleção (soft delete)
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relações
	LabResultBatch    LabResultBatch     `gorm:"foreignKey:LabResultBatchID;constraint:OnDelete:CASCADE" json:"labResultBatch,omitempty"`
	LabTestDefinition *LabTestDefinition `gorm:"foreignKey:LabTestDefinitionID;constraint:OnDelete:SET NULL" json:"labTestDefinition,omitempty"`
	LabResultValues   []LabResultValue   `gorm:"foreignKey:LabResultID;constraint:OnDelete:CASCADE" json:"labResultValues,omitempty"`
}

// TableName especifica o nome da tabela
func (LabResult) TableName() string {
	return "lab_results"
}

// BeforeCreate hook to generate UUID v7
func (lr *LabResult) BeforeCreate(tx *gorm.DB) error {
	if lr.ID == uuid.Nil {
		lr.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
