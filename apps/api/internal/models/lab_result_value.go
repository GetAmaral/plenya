package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// LabResultValue armazena valores estruturados de exames laboratoriais
// @Description Valor estruturado de um parâmetro de exame laboratorial
type LabResultValue struct {
	// ID único do valor
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`

	// ID do resultado laboratorial (LabResult)
	LabResultID uuid.UUID `gorm:"type:uuid;not null;index:idx_lab_result_value_result" json:"labResultId" validate:"required"`

	// ID da definição do teste (LabTestDefinition)
	LabTestDefinitionID uuid.UUID `gorm:"type:uuid;not null;index:idx_lab_result_value_test" json:"labTestDefinitionId" validate:"required"`

	// Valor numérico (para resultados quantitativos)
	NumericValue *float64 `gorm:"type:double precision" json:"numericValue,omitempty"`

	// Valor textual (para resultados qualitativos)
	// @example Negativo, Positivo, Reagente, Não Reagente
	TextValue *string `gorm:"type:text" json:"textValue,omitempty"`

	// Valor booleano (para resultados sim/não)
	BooleanValue *bool `gorm:"type:boolean" json:"booleanValue,omitempty"`

	// Unidade de medida do valor (pode sobrescrever a unidade padrão)
	Unit *string `gorm:"type:varchar(50)" json:"unit,omitempty"`

	// Faixa de referência aplicável (texto)
	ReferenceRange *string `gorm:"type:varchar(200)" json:"referenceRange,omitempty"`

	// Indica se o valor está fora da faixa de referência
	IsAbnormal bool `gorm:"type:boolean;not null;default:false" json:"isAbnormal"`

	// Indica se o valor é crítico (requer atenção imediata)
	IsCritical bool `gorm:"type:boolean;not null;default:false" json:"isCritical"`

	// Observações sobre o valor
	Notes *string `gorm:"type:text" json:"notes,omitempty"`

	// Timestamps
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relacionamentos
	LabResult         LabResult         `gorm:"foreignKey:LabResultID;constraint:OnDelete:CASCADE" json:"labResult,omitempty"`
	LabTestDefinition LabTestDefinition `gorm:"foreignKey:LabTestDefinitionID;constraint:OnDelete:RESTRICT" json:"labTestDefinition,omitempty"`
}

// TableName especifica o nome da tabela
func (LabResultValue) TableName() string {
	return "lab_result_values"
}

// BeforeCreate hook to generate UUID v7
func (lrv *LabResultValue) BeforeCreate(tx *gorm.DB) error {
	if lrv.ID == uuid.Nil {
		lrv.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
