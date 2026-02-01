package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AppointmentStatus define os status possíveis de uma consulta
type AppointmentStatus string

const (
	AppointmentScheduled AppointmentStatus = "scheduled" // Agendada
	AppointmentConfirmed AppointmentStatus = "confirmed" // Confirmada
	AppointmentCompleted AppointmentStatus = "completed" // Realizada
	AppointmentCancelled AppointmentStatus = "cancelled" // Cancelada
	AppointmentNoShow    AppointmentStatus = "no_show"   // Paciente não compareceu
)

// AppointmentType define os tipos de consulta
type AppointmentType string

const (
	AppointmentRoutine   AppointmentType = "routine"    // Consulta de rotina
	AppointmentFollowUp  AppointmentType = "follow_up"  // Retorno
	AppointmentUrgent    AppointmentType = "urgent"     // Urgência
	AppointmentEmergency AppointmentType = "emergency"  // Emergência
)

// Appointment representa uma consulta/agendamento
// @Description Consulta médica com data, hora, médico e paciente
type Appointment struct {
	// ID único da consulta
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// ID do paciente
	PatientID uuid.UUID `gorm:"type:uuid;not null;index" json:"patientId"`

	// ID do médico
	DoctorID uuid.UUID `gorm:"type:uuid;not null;index" json:"doctorId"`

	// Data e hora da consulta
	ScheduledAt time.Time `gorm:"type:timestamp;not null;index" json:"scheduledAt"`

	// Duração estimada em minutos
	DurationMinutes int `gorm:"type:int;not null;default:30" json:"durationMinutes"`

	// Tipo de consulta
	Type AppointmentType `gorm:"type:varchar(20);not null;check:type IN ('routine','follow_up','urgent','emergency')" json:"type"`

	// Status da consulta
	Status AppointmentStatus `gorm:"type:varchar(20);not null;default:'scheduled';check:status IN ('scheduled','confirmed','completed','cancelled','no_show')" json:"status"`

	// Motivo da consulta
	Reason string `gorm:"type:text;not null" json:"reason" validate:"required"`

	// Observações do paciente
	PatientNotes *string `gorm:"type:text" json:"patientNotes,omitempty"`

	// Observações do médico (preenchido após a consulta)
	DoctorNotes *string `gorm:"type:text" json:"doctorNotes,omitempty"`

	// Diagnóstico (preenchido após a consulta)
	Diagnosis *string `gorm:"type:text" json:"diagnosis,omitempty"`

	// Link para anamnese (se criada durante a consulta)
	AnamnesisID *uuid.UUID `gorm:"type:uuid;index" json:"anamnesisId,omitempty"`

	// Data de confirmação
	ConfirmedAt *time.Time `gorm:"type:timestamp" json:"confirmedAt,omitempty"`

	// Data de conclusão
	CompletedAt *time.Time `gorm:"type:timestamp" json:"completedAt,omitempty"`

	// Data de cancelamento
	CancelledAt *time.Time `gorm:"type:timestamp" json:"cancelledAt,omitempty"`

	// Motivo do cancelamento
	CancellationReason *string `gorm:"type:text" json:"cancellationReason,omitempty"`

	// Data de criação
	CreatedAt time.Time `gorm:"not null;autoCreateTime" json:"createdAt"`

	// Data de atualização
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime" json:"updatedAt"`

	// Data de deleção (soft delete)
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relações
	Patient   Patient    `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE" json:"patient,omitempty"`
	Doctor    User       `gorm:"foreignKey:DoctorID;constraint:OnDelete:RESTRICT" json:"doctor,omitempty"`
	Anamnesis *Anamnesis `gorm:"foreignKey:AnamnesisID" json:"anamnesis,omitempty"`
}

// TableName especifica o nome da tabela
func (Appointment) TableName() string {
	return "appointments"
}

// BeforeCreate hook to generate UUID v7
func (a *Appointment) BeforeCreate(tx *gorm.DB) error {
	if a.ID == uuid.Nil {
		a.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
