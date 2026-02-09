package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// LabTestUnitConversion represents a unit conversion rule for a lab test
// @Description Conversão de unidades para exames laboratoriais
type LabTestUnitConversion struct {
	// UUID v7 (time-ordered) - SEM default PostgreSQL
	// @example 019c1a1e-0579-7f3b-a1bd-4767008e844c
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// ID da definição de exame
	// @example 019c1a1e-0579-7f3b-a1bd-4767008e844c
	LabTestDefinitionID uuid.UUID `gorm:"type:uuid;not null;index:idx_unit_conv_test" json:"labTestDefinitionId" validate:"required"`

	// Unidade principal/padrão (ex: "g/dL")
	// @example g/dL
	// @maxLength 50
	MainUnit string `gorm:"type:varchar(50);not null" json:"mainUnit" validate:"required,max=50"`

	// Unidade alternativa/secundária (ex: "g/L")
	// @example g/L
	// @maxLength 50
	SecondaryUnit string `gorm:"type:varchar(50);not null;uniqueIndex:idx_unique_test_secondary_unit" json:"secondaryUnit" validate:"required,max=50"`

	// Fator de conversão: secondaryValue = mainValue * factor
	// @example 10
	ConversionFactor float64 `gorm:"type:double precision;not null;check:conversion_factor > 0" json:"conversionFactor" validate:"required,gt=0,lte=1000000"`

	// Relacionamento com LabTestDefinition
	LabTestDefinition *LabTestDefinition `gorm:"foreignKey:LabTestDefinitionID;constraint:OnDelete:CASCADE" json:"labTestDefinition,omitempty"`

	// Timestamps
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// BeforeCreate hook para gerar UUID v7
func (luc *LabTestUnitConversion) BeforeCreate(tx *gorm.DB) error {
	if luc.ID == uuid.Nil {
		luc.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// BeforeSave hook para validar conversion factor
func (luc *LabTestUnitConversion) BeforeSave(tx *gorm.DB) error {
	if luc.ConversionFactor <= 0 {
		return gorm.ErrInvalidValue
	}
	return nil
}

// ConvertToMain converte um valor da unidade secundária para a principal
func (luc *LabTestUnitConversion) ConvertToMain(secondaryValue float64) float64 {
	return secondaryValue / luc.ConversionFactor
}

// ConvertToSecondary converte um valor da unidade principal para a secundária
func (luc *LabTestUnitConversion) ConvertToSecondary(mainValue float64) float64 {
	return mainValue * luc.ConversionFactor
}
