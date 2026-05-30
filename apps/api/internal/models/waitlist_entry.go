package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// WaitlistStatus define os estados de uma entrada na lista de espera.
type WaitlistStatus string

const (
	WaitlistWaiting   WaitlistStatus = "waiting"   // aguardando encaixe
	WaitlistScheduled WaitlistStatus = "scheduled" // virou consulta
	WaitlistCancelled WaitlistStatus = "cancelled" // desistiu / não cabe mais
)

// WaitlistEntry — lista de espera de encaixe da recepção.
//
// Quando não há vaga no horário desejado, a secretária registra o paciente aqui
// (opcionalmente com médico e tipo de consulta preferidos). Quando uma vaga
// abre, a entrada é convertida em agendamento manualmente (sem aviso automático
// nesta fase). É um artefato operacional do balcão, não clínico.
type WaitlistEntry struct {
	// ID único da entrada
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Paciente que aguarda encaixe
	PatientID uuid.UUID `gorm:"type:uuid;not null;index" json:"patientId"`

	// Médico preferido (opcional — pode aguardar qualquer médico)
	DoctorID *uuid.UUID `gorm:"type:uuid;index" json:"doctorId,omitempty"`

	// Tipo de consulta desejado (opcional)
	PreferredType *AppointmentType `gorm:"type:varchar(20)" json:"preferredType,omitempty"`

	// Observações da recepção (janela de horário preferida, urgência, etc.)
	Notes *string `gorm:"type:text" json:"notes,omitempty"`

	// Status da entrada
	Status WaitlistStatus `gorm:"type:varchar(20);not null;default:'waiting';check:status IN ('waiting','scheduled','cancelled')" json:"status"`

	// Consulta gerada ao converter a entrada (preenchido quando status=scheduled)
	ScheduledAppointmentID *uuid.UUID `gorm:"type:uuid;index" json:"scheduledAppointmentId,omitempty"`

	// Quem criou a entrada (staff)
	CreatedByUserID uuid.UUID `gorm:"type:uuid;not null;index" json:"createdByUserId"`

	CreatedAt time.Time      `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relações
	Patient Patient `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE" json:"patient,omitempty"`
	Doctor  *User   `gorm:"foreignKey:DoctorID;constraint:OnDelete:SET NULL" json:"doctor,omitempty"`
}

// TableName especifica o nome da tabela
func (WaitlistEntry) TableName() string {
	return "waitlist_entries"
}

// BeforeCreate gera UUID v7.
func (w *WaitlistEntry) BeforeCreate(tx *gorm.DB) error {
	if w.ID == uuid.Nil {
		w.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
