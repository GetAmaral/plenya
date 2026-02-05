package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PrescriptionMedication representa um medicamento em uma prescrição
// Uma prescrição pode ter múltiplos medicamentos
// @Description Medicamento prescrito em uma prescrição
type PrescriptionMedication struct {
	// ID único do medicamento na prescrição
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// ID da prescrição (FK)
	// @required
	PrescriptionID uuid.UUID `gorm:"type:uuid;not null;index" json:"prescriptionId" validate:"required"`

	// ID da definição de medicamento (opcional - se foi selecionado do catálogo)
	MedicationDefinitionID *uuid.UUID `gorm:"type:uuid;index" json:"medicationDefinitionId,omitempty"`

	// Nome comercial do medicamento
	// @example Losartana Potássica
	// @minLength 3
	// @maxLength 200
	MedicationName string `gorm:"type:varchar(200);not null" json:"medicationName" validate:"required,min=3,max=200"`

	// Princípio ativo (DCB - Denominação Comum Brasileira)
	// @example Losartana
	// @minLength 3
	// @maxLength 200
	ActiveIngredient string `gorm:"type:varchar(200);not null" json:"activeIngredient" validate:"required,min=3,max=200"`

	// Categoria do medicamento
	// @enum simple,c1,c5,antibiotic,glp1
	Category MedicationCategory `gorm:"type:varchar(20);not null;default:'simple';index" json:"category"`

	// Concentração
	// @example 50mg
	Concentration string `gorm:"type:varchar(100);not null" json:"concentration" validate:"required"`

	// Dosagem
	// @example 1 comprimido
	Dosage string `gorm:"type:varchar(100);not null" json:"dosage" validate:"required"`

	// Frequência
	// @example 1x ao dia
	Frequency string `gorm:"type:varchar(100);not null" json:"frequency" validate:"required"`

	// Via de administração
	// @example oral
	Route string `gorm:"type:varchar(50);not null" json:"route" validate:"required"`

	// Duração do tratamento (dias)
	// @example 30
	Duration int `gorm:"type:integer;not null" json:"duration" validate:"required,min=1"`

	// Quantidade prescrita
	// @example 30
	Quantity int `gorm:"type:integer;not null" json:"quantity" validate:"required,min=1"`

	// Quantidade por extenso (obrigatório para controlados)
	// @example trinta comprimidos
	QuantityInWords string `gorm:"type:varchar(200);not null" json:"quantityInWords" validate:"required"`

	// Instruções específicas deste medicamento (opcional)
	Instructions *string `gorm:"type:text" json:"instructions,omitempty"`

	// Timestamps
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relações
	Prescription         *Prescription         `gorm:"foreignKey:PrescriptionID;constraint:OnDelete:CASCADE" json:"prescription,omitempty"`
	MedicationDefinition *MedicationDefinition `gorm:"foreignKey:MedicationDefinitionID" json:"medicationDefinition,omitempty"`
}

// TableName especifica o nome da tabela
func (PrescriptionMedication) TableName() string {
	return "prescription_medications"
}

// BeforeCreate hook to generate UUID v7
func (pm *PrescriptionMedication) BeforeCreate(tx *gorm.DB) error {
	if pm.ID == uuid.Nil {
		pm.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
