package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// RecallStatus define os estados de um retorno previsto.
type RecallStatus string

const (
	RecallPending   RecallStatus = "pending"   // aguardando contato/agendamento
	RecallScheduled RecallStatus = "scheduled" // virou consulta
	RecallDismissed RecallStatus = "dismissed" // dispensado (não aplicável / paciente não quer)
)

// AppointmentRecall — fila interna de retornos previstos da recepção.
//
// Ao concluir uma consulta, a recepção (ou o médico) registra um retorno previsto
// (ex.: reavaliação anual de biomarcadores, retorno em 3 meses). A secretária
// trabalha a fila manualmente: contata o paciente e agenda. SEM aviso automático
// ao paciente nesta fase. É um artefato operacional do balcão, não clínico.
type AppointmentRecall struct {
	// ID único do retorno
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Paciente a retornar
	PatientID uuid.UUID `gorm:"type:uuid;not null;index" json:"patientId"`

	// Médico de preferência do retorno (opcional)
	DoctorID *uuid.UUID `gorm:"type:uuid;index" json:"doctorId,omitempty"`

	// Consulta que originou o retorno (opcional)
	SourceAppointmentID *uuid.UUID `gorm:"type:uuid;index" json:"sourceAppointmentId,omitempty"`

	// Data-alvo do retorno (quando contatar/agendar)
	DueDate time.Time `gorm:"type:date;not null;index" json:"dueDate"`

	// Motivo do retorno (ex.: reavaliação anual de biomarcadores)
	Reason *string `gorm:"type:text" json:"reason,omitempty"`

	// Observações da recepção
	Notes *string `gorm:"type:text" json:"notes,omitempty"`

	// Status do retorno
	Status RecallStatus `gorm:"type:varchar(20);not null;default:'pending';check:status IN ('pending','scheduled','dismissed')" json:"status"`

	// Consulta gerada ao agendar o retorno (preenchido quando status=scheduled)
	ScheduledAppointmentID *uuid.UUID `gorm:"type:uuid;index" json:"scheduledAppointmentId,omitempty"`

	// Quem registrou o retorno (staff)
	CreatedByUserID uuid.UUID `gorm:"type:uuid;not null;index" json:"createdByUserId"`

	CreatedAt time.Time      `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relações
	Patient Patient `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE" json:"patient,omitempty"`
	Doctor  *User   `gorm:"foreignKey:DoctorID;constraint:OnDelete:SET NULL" json:"doctor,omitempty"`
}

// TableName especifica o nome da tabela
func (AppointmentRecall) TableName() string {
	return "appointment_recalls"
}

// BeforeCreate gera UUID v7.
func (r *AppointmentRecall) BeforeCreate(tx *gorm.DB) error {
	if r.ID == uuid.Nil {
		r.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
