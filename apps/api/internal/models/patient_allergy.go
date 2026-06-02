package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AllergySubstanceType classifica a natureza da alergia.
type AllergySubstanceType string

const (
	AllergyTypeDrug          AllergySubstanceType = "drug"
	AllergyTypeFood          AllergySubstanceType = "food"
	AllergyTypeEnvironmental AllergySubstanceType = "environmental"
	AllergyTypeOther         AllergySubstanceType = "other"
)

// AllergySeverity gradua a gravidade da reação.
type AllergySeverity string

const (
	AllergySeverityMild        AllergySeverity = "mild"
	AllergySeverityModerate    AllergySeverity = "moderate"
	AllergySeveritySevere      AllergySeverity = "severe"
	AllergySeverityAnaphylaxis AllergySeverity = "anaphylaxis"
)

// AllergyStatus indica se a alergia está ativa no quadro atual.
type AllergyStatus string

const (
	AllergyStatusActive   AllergyStatus = "active"
	AllergyStatusInactive AllergyStatus = "inactive"
)

// PatientAllergy registra uma alergia do paciente. Base do CDS de prescrição
// (cruza Substance/princípio ativo com o ActiveIngredient da medicação).
//
// NoKnownAllergies=true representa a asserção explícita "sem alergias
// conhecidas" (registro ativo, sem substância) — distinto de "nunca perguntado".
//
// @Description Alergia do paciente (medicamentosa, alimentar, ambiental)
type PatientAllergy struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	PatientID uuid.UUID `gorm:"type:uuid;not null;index" json:"patientId"`

	// Substância/agente (ex.: "Dipirona", "Penicilina", "Frutos do mar").
	// Vazio quando NoKnownAllergies=true.
	Substance     string               `gorm:"type:varchar(200);not null;default:''" json:"substance"`
	SubstanceType AllergySubstanceType `gorm:"type:varchar(15);not null;default:'drug';check:substance_type IN ('drug','food','environmental','other')" json:"substanceType"`
	Reaction      *string              `gorm:"type:text" json:"reaction,omitempty"`
	Severity      AllergySeverity      `gorm:"type:varchar(12);not null;default:'moderate';check:severity IN ('mild','moderate','severe','anaphylaxis')" json:"severity"`
	Status        AllergyStatus        `gorm:"type:varchar(10);not null;default:'active';check:status IN ('active','inactive')" json:"status"`

	// Asserção "sem alergias conhecidas" (NKA). Quando true, esta é a única
	// linha relevante e Substance fica vazio.
	NoKnownAllergies bool `gorm:"type:boolean;not null;default:false" json:"noKnownAllergies"`

	Notes            *string   `gorm:"type:text" json:"notes,omitempty"`
	RecordedByUserID uuid.UUID `gorm:"type:uuid;not null" json:"recordedByUserId"`

	CreatedAt time.Time      `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Patient        Patient `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE" json:"-"`
	RecordedByUser User    `gorm:"foreignKey:RecordedByUserID;constraint:OnDelete:RESTRICT" json:"recordedByUser,omitempty"`
}

func (PatientAllergy) TableName() string { return "patient_allergies" }

func (a *PatientAllergy) BeforeCreate(tx *gorm.DB) error {
	if a.ID == uuid.Nil {
		a.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
