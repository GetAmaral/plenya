package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// MedicationCategory define as categorias regulatórias de medicamentos
type MedicationCategory string

const (
	MedCategorySimple     MedicationCategory = "simple"      // Receita simples
	MedCategoryC1         MedicationCategory = "c1"          // Controle Especial (Lista C1)
	MedCategoryC5         MedicationCategory = "c5"          // Psicotrópicos (Lista C5)
	MedCategoryAntibiotic MedicationCategory = "antibiotic"  // Antimicrobianos (RDC 471)
	MedCategoryGLP1       MedicationCategory = "glp1"        // GLP-1 agonistas
)

// MedicationDefinition representa a definição de um medicamento no catálogo
// @Description Definição de medicamento com regras regulatórias (ANVISA)
type MedicationDefinition struct {
	// ID único da definição
	// @example 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Nome comercial comum do medicamento
	// @minLength 3
	// @maxLength 500
	// @example Fluoxetina 20mg
	CommonName string `gorm:"type:varchar(500);not null;index" json:"commonName" validate:"required,min=3,max=500"`

	// Princípio ativo (DCB - Denominação Comum Brasileira)
	// @minLength 3
	// @maxLength 500
	// @example Cloridrato de Fluoxetina
	ActiveIngredient string `gorm:"type:varchar(500);not null" json:"activeIngredient" validate:"required,min=3,max=500"`

	// Categoria regulatória do medicamento
	// @enum simple,c1,c5,antibiotic,glp1
	Category MedicationCategory `gorm:"type:varchar(20);not null;index;check:category IN ('simple','c1','c5','antibiotic','glp1')" json:"category" validate:"required,oneof=simple c1 c5 antibiotic glp1"`

	// Regras de validação baseadas na categoria

	// Dias de validade da prescrição
	// @example 30
	// C1: 30 dias, Antibióticos: 10 dias, GLP1: 90 dias, Simples: 30 dias
	ValidityDays int `gorm:"type:integer;not null;default:30" json:"validityDays" validate:"required,min=1,max=365"`

	// Máximo de substâncias por prescrição
	// @example 3
	// C1: máximo 3 substâncias
	MaxPerPrescription int `gorm:"type:integer;not null;default:10" json:"maxPerPrescription" validate:"required,min=1,max=100"`

	// Duração máxima do tratamento em dias
	// @example 60
	// C1: 60 dias, Anticonvulsivantes: 180 dias
	MaxTreatmentDays int `gorm:"type:integer;not null;default:60" json:"maxTreatmentDays" validate:"required,min=1,max=365"`

	// Requer assinatura digital ICP-Brasil
	// true para medicamentos controlados (C1, C5)
	RequiresDigitalSignature bool `gorm:"type:boolean;not null;default:false" json:"requiresDigitalSignature"`

	// Requer registro no SNCR (Sistema Nacional de Controle de Receitas)
	// true para medicamentos controlados
	RequiresSNCR bool `gorm:"type:boolean;not null;default:false" json:"requiresSNCR"`

	// Metadata

	// Código ANVISA do medicamento
	// @example 1234567890123
	ANVISACode *string `gorm:"type:varchar(50)" json:"anvisaCode,omitempty"`

	// Timestamps
	CreatedAt time.Time      `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName especifica o nome da tabela
func (MedicationDefinition) TableName() string {
	return "medication_definitions"
}

// BeforeCreate hook to generate UUID v7
func (m *MedicationDefinition) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
