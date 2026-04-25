package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// DoctorAbsence representa um período de indisponibilidade de um médico
// (férias, congresso, feriado, atestado).
//
// Slot picker exclui slots que caem dentro de qualquer DoctorAbsence ativa.
// Diferente de WorkingHours (recorrente semanal), DoctorAbsence é pontual
// — datas absolutas StartAt/EndAt.
//
// @Description Período de ausência pontual de um médico
type DoctorAbsence struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Médico ausente (FK users.id, role=doctor).
	DoctorID uuid.UUID `gorm:"type:uuid;not null;index:idx_absence_doctor" json:"doctorId"`

	// Início da ausência (timestamp completo).
	StartAt time.Time `gorm:"not null;index:idx_absence_range,priority:1" json:"startAt"`

	// Fim da ausência (timestamp completo). Slot picker considera
	// [StartAt, EndAt) — exclusivo no fim.
	EndAt time.Time `gorm:"not null;index:idx_absence_range,priority:2" json:"endAt"`

	// Motivo livre da ausência (mostrado na UI da agenda — não é PHI).
	// Ex: "Férias", "Congresso SBEM", "Feriado nacional".
	Reason string `gorm:"type:varchar(100)" json:"reason"`

	CreatedAt time.Time `gorm:"not null;autoCreateTime" json:"createdAt"`

	// Relação
	Doctor *User `gorm:"foreignKey:DoctorID;constraint:OnDelete:CASCADE" json:"-"`
}

// TableName especifica o nome da tabela
func (DoctorAbsence) TableName() string {
	return "doctor_absences"
}

// BeforeCreate hook gera UUID v7
func (d *DoctorAbsence) BeforeCreate(tx *gorm.DB) error {
	if d.ID == uuid.Nil {
		d.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// BeforeSave valida que EndAt > StartAt.
func (d *DoctorAbsence) BeforeSave(tx *gorm.DB) error {
	if !d.EndAt.After(d.StartAt) {
		return errors.New("doctor_absence: end_at deve ser posterior a start_at")
	}
	return nil
}
