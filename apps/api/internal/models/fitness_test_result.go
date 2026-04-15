package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// FitnessTestResult stores raw test results + computed classifications
// @Description Resultado de teste de fitness com classificacoes ACSM
type FitnessTestResult struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	PatientID   uuid.UUID `gorm:"type:uuid;not null;index" json:"patientId"`
	CreatedByID uuid.UUID `gorm:"type:uuid;not null" json:"createdById"`
	AssessmentDate time.Time `gorm:"type:date;not null" json:"assessmentDate"`

	// Raw inputs
	AbdominalReps *int `gorm:"type:int" json:"abdominalReps"`   // rep/min
	PushupReps    *int `gorm:"type:int" json:"pushupReps"`      // rep/min
	PlankSeconds  *int `gorm:"type:int" json:"plankSeconds"`    // seconds
	BurpeeCycles  *int `gorm:"type:int" json:"burpeeCycles"`    // cycles/3min
	FRTReps       *int `gorm:"type:int" json:"frtReps"`         // rep/90s (Functional Resistance Test)

	// Computed (by service, NOT input)
	AbdominalLevel *string `gorm:"type:varchar(20)" json:"abdominalLevel"`
	PushupLevel    *string `gorm:"type:varchar(20)" json:"pushupLevel"`
	PlankLevel     *string `gorm:"type:varchar(20)" json:"plankLevel"`
	BurpeeLevel    *string `gorm:"type:varchar(20)" json:"burpeeLevel"`
	FRTLevel       *string `gorm:"type:varchar(20)" json:"frtLevel"`

	OverallScore          int    `gorm:"type:int;not null;default:0" json:"overallScore"`
	OverallClassification string `gorm:"type:varchar(30);not null;default:''" json:"overallClassification"`

	Notes *string `gorm:"type:text" json:"notes"`

	// Timestamps
	CreatedAt time.Time      `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	Patient   Patient `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE" json:"patient,omitempty"`
	CreatedBy User    `gorm:"foreignKey:CreatedByID;constraint:OnDelete:RESTRICT" json:"createdBy,omitempty"`
}

func (FitnessTestResult) TableName() string {
	return "fitness_test_results"
}

func (f *FitnessTestResult) BeforeCreate(tx *gorm.DB) error {
	if f.ID == uuid.Nil {
		f.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
