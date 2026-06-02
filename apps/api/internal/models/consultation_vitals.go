package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ConsultationVitals são os sinais vitais aferidos DURANTE a consulta.
//
// Distinto de PhysicalAssessment (módulo Treino, estratificação ACSM do
// educador físico): aqui é o vital leve, por visita, do atendimento clínico.
// Ancorado ao Appointment.
//
// @Description Sinais vitais de uma consulta (PA, FC, peso, etc.)
type ConsultationVitals struct {
	ID            uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	AppointmentID *uuid.UUID `gorm:"type:uuid;index" json:"appointmentId,omitempty"`
	PatientID     uuid.UUID  `gorm:"type:uuid;not null;index" json:"patientId"`

	SystolicBP        *int     `gorm:"type:int" json:"systolicBp,omitempty"`
	DiastolicBP       *int     `gorm:"type:int" json:"diastolicBp,omitempty"`
	HeartRate         *int     `gorm:"type:int" json:"heartRate,omitempty"`
	RespRate          *int     `gorm:"type:int" json:"respRate,omitempty"`
	Temperature       *float64 `gorm:"type:decimal(4,1)" json:"temperature,omitempty"`
	SpO2              *int     `gorm:"type:int;column:spo2" json:"spo2,omitempty"`
	Weight            *float64 `gorm:"type:decimal(5,2)" json:"weight,omitempty"`
	Height            *float64 `gorm:"type:decimal(5,2)" json:"height,omitempty"`
	WaistCircumference *float64 `gorm:"type:decimal(5,2)" json:"waistCircumference,omitempty"`

	// BMI computado em BeforeSave a partir de Weight (kg) e Height (cm).
	BMI *float64 `gorm:"type:decimal(4,1)" json:"bmi,omitempty"`

	MeasuredByUserID uuid.UUID `gorm:"type:uuid;not null" json:"measuredByUserId"`
	MeasuredAt       time.Time `gorm:"type:timestamptz;not null" json:"measuredAt"`

	CreatedAt time.Time      `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Patient        Patient      `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE" json:"-"`
	Appointment    *Appointment `gorm:"foreignKey:AppointmentID" json:"-"`
	MeasuredByUser User         `gorm:"foreignKey:MeasuredByUserID;constraint:OnDelete:RESTRICT" json:"measuredByUser,omitempty"`
}

func (ConsultationVitals) TableName() string { return "consultation_vitals" }

func (v *ConsultationVitals) BeforeCreate(tx *gorm.DB) error {
	if v.ID == uuid.Nil {
		v.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// BeforeSave calcula o IMC quando há peso e altura.
func (v *ConsultationVitals) BeforeSave(tx *gorm.DB) error {
	if v.Weight != nil && v.Height != nil && *v.Height > 0 {
		m := *v.Height / 100.0
		bmi := *v.Weight / (m * m)
		bmi = float64(int(bmi*10+0.5)) / 10 // 1 casa
		v.BMI = &bmi
	} else {
		v.BMI = nil
	}
	return nil
}
