package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ContinuumTemplate é o esqueleto de um programa Continuum (Semestral, Anual,
// Trimestral...). Admin/manager edita; cada inscrição de paciente faz snapshot
// imutável (TemplateSnapshot em PatientContinuum) — mudanças posteriores no
// template NÃO afetam inscrições já abertas.
type ContinuumTemplate struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name            string    `gorm:"type:varchar(120);not null" json:"name" validate:"required"`
	Description     string    `gorm:"type:text" json:"description"`
	DurationWeeks   int       `gorm:"not null;check:duration_weeks > 0" json:"durationWeeks" validate:"required,min=1,max=520"`
	Status          string    `gorm:"type:varchar(20);not null;default:'active';check:status IN ('active','archived')" json:"status"`
	CreatedByUserID uuid.UUID `gorm:"type:uuid;not null" json:"createdByUserId"`

	CreatedAt time.Time      `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Items []ContinuumTemplateItem `gorm:"foreignKey:TemplateID;constraint:OnDelete:CASCADE" json:"items,omitempty"`
}

func (ContinuumTemplate) TableName() string { return "continuum_templates" }

func (t *ContinuumTemplate) BeforeCreate(tx *gorm.DB) error {
	if t.ID == uuid.Nil {
		t.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// ContinuumItemType enumera os tipos de marco/entrega que compõem um Continuum.
type ContinuumItemType string

const (
	ContinuumItemAppointment  ContinuumItemType = "appointment"  // consulta com profissional
	ContinuumItemBox          ContinuumItemType = "box"          // entrega de Box Plenya
	ContinuumItemReassessment ContinuumItemType = "reassessment" // reavaliação trimestral (escore + exames)
	ContinuumItemMilestone    ContinuumItemType = "milestone"    // marco genérico (reunião equipe, plano integrado, fechamento)
	ContinuumItemCustom       ContinuumItemType = "custom"       // item livre
)

// ContinuumItemSpecialty quando Type=appointment indica qual especialista.
type ContinuumItemSpecialty string

const (
	ContinuumSpecialtyDoctor           ContinuumItemSpecialty = "doctor"
	ContinuumSpecialtyNutritionist     ContinuumItemSpecialty = "nutritionist"
	ContinuumSpecialtyPsychologist     ContinuumItemSpecialty = "psychologist"
	ContinuumSpecialtyPhysicalEducator ContinuumItemSpecialty = "physicalEducator"
)

// ContinuumTemplateItem é uma linha do template: representa um marco/entrega
// que vai virar PatientContinuumItem na inscrição. Ordenado por (WeekOffset,
// Position).
type ContinuumTemplateItem struct {
	ID         uuid.UUID         `gorm:"type:uuid;primaryKey" json:"id"`
	TemplateID uuid.UUID         `gorm:"type:uuid;not null;index" json:"templateId"`
	Type       ContinuumItemType `gorm:"type:varchar(30);not null;check:type IN ('appointment','box','reassessment','milestone','custom')" json:"type" validate:"required"`

	// Specialty só aplicável quando Type=appointment. Pode ser nil pra outros tipos.
	Specialty *ContinuumItemSpecialty `gorm:"type:varchar(30)" json:"specialty,omitempty"`

	Title       string `gorm:"type:varchar(160);not null" json:"title" validate:"required"`
	Description string `gorm:"type:text" json:"description"`

	// WeekOffset 0-based — semana relativa ao startDate da inscrição.
	WeekOffset int `gorm:"not null;check:week_offset >= 0" json:"weekOffset" validate:"min=0"`

	// ExpectedOffsetDays — dia dentro da semana (0..6) em que o marco é esperado.
	ExpectedOffsetDays int `gorm:"not null;default:0;check:expected_offset_days >= 0 AND expected_offset_days <= 6" json:"expectedOffsetDays"`

	// LateAfterDays — dias após ExpectedDate em que o item vira "missed"
	// se ainda estiver pending. Cron ContinuumLateJob aplica essa regra.
	LateAfterDays int `gorm:"not null;default:7;check:late_after_days >= 0" json:"lateAfterDays"`

	// BoxTemplateID referencia ContinuumBoxTemplate quando Type=box.
	BoxTemplateID *uuid.UUID `gorm:"type:uuid" json:"boxTemplateId,omitempty"`

	Position int `gorm:"not null;default:0" json:"position"`

	CreatedAt time.Time `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime" json:"updatedAt"`
}

func (ContinuumTemplateItem) TableName() string { return "continuum_template_items" }

func (i *ContinuumTemplateItem) BeforeCreate(tx *gorm.DB) error {
	if i.ID == uuid.Nil {
		i.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
