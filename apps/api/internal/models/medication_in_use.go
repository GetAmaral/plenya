package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// MedicationSource — origem da medicação em uso.
type MedicationSource string

const (
	MedSourcePrescribedHere MedicationSource = "prescribed_here"
	MedSourceExternal       MedicationSource = "external"
	MedSourcePatientReported MedicationSource = "patient_reported"
)

// MedicationInUseStatus — estado atual da medicação em uso.
type MedicationInUseStatus string

const (
	MedInUseActive    MedicationInUseStatus = "active"
	MedInUseSuspended MedicationInUseStatus = "suspended"
	MedInUseStopped   MedicationInUseStatus = "stopped"
)

// MedicationInUse é o ESTADO atual de medicação do paciente (reconciliação),
// distinto de Prescription (o ATO de emitir receita, com validade/assinatura).
// Inclui medicações externas e auto-reportadas. SourcePrescriptionID liga à
// receita quando originou aqui.
//
// @Description Medicação em uso pelo paciente (reconciliação)
type MedicationInUse struct {
	ID               uuid.UUID             `gorm:"type:uuid;primaryKey" json:"id"`
	PatientID        uuid.UUID             `gorm:"type:uuid;not null;index" json:"patientId"`
	MedicationName   string                `gorm:"type:varchar(200);not null" json:"medicationName"`
	ActiveIngredient string                `gorm:"type:varchar(200);not null;default:''" json:"activeIngredient"`
	Dosage           string                `gorm:"type:varchar(100);not null;default:''" json:"dosage"`
	Frequency        string                `gorm:"type:varchar(100);not null;default:''" json:"frequency"`
	Route            string                `gorm:"type:varchar(50);not null;default:''" json:"route"`
	Source           MedicationSource      `gorm:"type:varchar(20);not null;default:'patient_reported';check:source IN ('prescribed_here','external','patient_reported')" json:"source"`
	Status           MedicationInUseStatus `gorm:"type:varchar(12);not null;default:'active';check:status IN ('active','suspended','stopped')" json:"status"`

	SourcePrescriptionID    *uuid.UUID `gorm:"type:uuid;index" json:"sourcePrescriptionId,omitempty"`
	ReconciledAppointmentID *uuid.UUID `gorm:"type:uuid;index" json:"reconciledAppointmentId,omitempty"`

	StartDate *time.Time `gorm:"type:date" json:"startDate,omitempty"`
	EndDate   *time.Time `gorm:"type:date" json:"endDate,omitempty"`
	Notes     *string    `gorm:"type:text" json:"notes,omitempty"`

	RecordedByUserID uuid.UUID `gorm:"type:uuid;not null" json:"recordedByUserId"`

	CreatedAt time.Time      `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Patient        Patient `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE" json:"-"`
	RecordedByUser User    `gorm:"foreignKey:RecordedByUserID;constraint:OnDelete:RESTRICT" json:"recordedByUser,omitempty"`
}

func (MedicationInUse) TableName() string { return "medications_in_use" }

func (m *MedicationInUse) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
