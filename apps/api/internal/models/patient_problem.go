package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ProblemStatus — ciclo de um problema na lista.
type ProblemStatus string

const (
	ProblemStatusActive   ProblemStatus = "active"
	ProblemStatusResolved ProblemStatus = "resolved"
	ProblemStatusInactive ProblemStatus = "inactive"
)

// PatientProblem é um item da lista de problemas do paciente.
//
// Description é texto livre (obrigatório); CIDCode é opcional (preenchido pelo
// autocomplete do catálogo CIDCode). CIDVersion guarda a versão pra acomodar
// CID-11 (jan/2027) sem migrar dados.
//
// @Description Problema ativo/resolvido da lista de problemas
type PatientProblem struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	PatientID   uuid.UUID `gorm:"type:uuid;not null;index" json:"patientId"`
	Description string    `gorm:"type:varchar(300);not null" json:"description"`

	CIDCode    *string `gorm:"type:varchar(10);index;column:cid_code" json:"cidCode,omitempty"`
	CIDVersion string  `gorm:"type:varchar(8);not null;default:'CID-10';column:cid_version" json:"cidVersion"`

	Status       ProblemStatus `gorm:"type:varchar(10);not null;default:'active';check:status IN ('active','resolved','inactive')" json:"status"`
	OnsetDate    *time.Time    `gorm:"type:date" json:"onsetDate,omitempty"`
	ResolvedDate *time.Time    `gorm:"type:date" json:"resolvedDate,omitempty"`

	// Consulta em que o problema foi identificado (opcional).
	OnsetAppointmentID *uuid.UUID `gorm:"type:uuid;index" json:"onsetAppointmentId,omitempty"`

	Notes            *string   `gorm:"type:text" json:"notes,omitempty"`
	RecordedByUserID uuid.UUID `gorm:"type:uuid;not null" json:"recordedByUserId"`

	CreatedAt time.Time      `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Patient        Patient `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE" json:"-"`
	RecordedByUser User    `gorm:"foreignKey:RecordedByUserID;constraint:OnDelete:RESTRICT" json:"recordedByUser,omitempty"`
}

func (PatientProblem) TableName() string { return "problem_list_items" }

func (p *PatientProblem) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
