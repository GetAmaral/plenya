package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// WorkingHours representa um bloco de horário de atendimento de um médico
// em um dia da semana específico.
//
// Modelo:
//   - Cada médico tem N rows. Ex: seg 8-12 + seg 14-18 = 2 rows.
//   - Quarta sem atendimento = 0 rows pra Weekday=3.
//   - SlotDuration controla a granularidade do slot picker (default 30min).
//
// Slot picker (CalendarSlotService) intersecciona:
//   working_hours ∩ ¬DoctorAbsence ∩ ¬GoogleFreeBusy ∩ ¬Appointment_existente
//
// @Description Bloco semanal de horário de atendimento de um médico
type WorkingHours struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Médico dono do bloco (FK users.id, role=doctor).
	DoctorID uuid.UUID `gorm:"type:uuid;not null;index:idx_wh_doctor" json:"doctorId"`

	// Dia da semana: 0=Domingo, 1=Segunda, ..., 6=Sábado.
	Weekday int `gorm:"not null;check:weekday >= 0 AND weekday <= 6" json:"weekday" validate:"min=0,max=6"`

	// Início do bloco em minutos desde meia-noite (0..1439).
	// Ex: 8h00 = 480, 14h30 = 870.
	StartMinute int `gorm:"not null;check:start_minute >= 0 AND start_minute < 1440" json:"startMinute" validate:"min=0,max=1439"`

	// Fim do bloco em minutos desde meia-noite (1..1440).
	// 1440 representa fim do dia (00:00 do dia seguinte conceitualmente).
	EndMinute int `gorm:"not null;check:end_minute > 0 AND end_minute <= 1440" json:"endMinute" validate:"min=1,max=1440"`

	// Duração de cada slot oferecido dentro do bloco (default 30min).
	// Slot picker oferece slots [StartMinute, StartMinute+SlotDuration, ...].
	SlotDuration int `gorm:"not null;default:30;check:slot_duration > 0 AND slot_duration <= 480" json:"slotDuration" validate:"min=1,max=480"`

	CreatedAt time.Time `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime" json:"updatedAt"`

	// Relação
	Doctor *User `gorm:"foreignKey:DoctorID;constraint:OnDelete:CASCADE" json:"-"`
}

// TableName especifica o nome da tabela
func (WorkingHours) TableName() string {
	return "working_hours"
}

// BeforeCreate hook gera UUID v7
func (w *WorkingHours) BeforeCreate(tx *gorm.DB) error {
	if w.ID == uuid.Nil {
		w.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// BeforeSave valida que EndMinute > StartMinute. CHECK constraints do DB
// já garantem ranges individuais (0..1439 / 1..1440) — esse hook cobre a
// relação entre os dois (que CHECK Postgres puro não pode expressar
// portavelmente em todos os GORMs sem cláusula multi-coluna).
func (w *WorkingHours) BeforeSave(tx *gorm.DB) error {
	if w.EndMinute <= w.StartMinute {
		return errors.New("working_hours: end_minute deve ser maior que start_minute")
	}
	if w.SlotDuration <= 0 {
		return errors.New("working_hours: slot_duration deve ser positivo")
	}
	if w.SlotDuration > (w.EndMinute - w.StartMinute) {
		return errors.New("working_hours: slot_duration não pode ser maior que o bloco")
	}
	return nil
}
