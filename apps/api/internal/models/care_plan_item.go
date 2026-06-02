package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CarePlanItem — recomendação do plano de cuidado, ancorada a um pilar AGIR (letra A/G/I/R)
// e, opcionalmente, a um biomarcador (lab_test_code) ou score item. O "plano" do paciente é
// o conjunto de itens ativos, agrupados por letra. Multidisciplinar: cada clínico (médico,
// nutrólogo, educador físico) adiciona recomendações à sua letra.
//
// Ancorar por LetterCode (e não pela M2M score_item_method_pillars, hoje esparsa) torna o
// plano robusto independentemente do mapeamento de pilares estar completo.
type CarePlanPriority string

const (
	CarePlanPriorityHigh   CarePlanPriority = "high"
	CarePlanPriorityMedium CarePlanPriority = "medium"
	CarePlanPriorityLow    CarePlanPriority = "low"
)

type CarePlanItemStatus string

const (
	CarePlanStatusActive    CarePlanItemStatus = "active"
	CarePlanStatusAchieved  CarePlanItemStatus = "achieved"
	CarePlanStatusSuspended CarePlanItemStatus = "suspended"
)

// @Description Recomendação do plano de cuidado AGIR
type CarePlanItem struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	PatientID uuid.UUID `gorm:"type:uuid;not null;index" json:"patientId"`

	// LetterCode — pilar AGIR: 'A' (Atividade), 'G' (Gestão), 'I' (Integração), 'R' (Ritmo).
	LetterCode string `gorm:"type:varchar(1);not null;check:letter_code IN ('A','G','I','R')" json:"letterCode"`

	Recommendation string  `gorm:"type:text;not null" json:"recommendation"`
	Rationale      *string `gorm:"type:text" json:"rationale,omitempty"`
	Target         *string `gorm:"type:varchar(300)" json:"target,omitempty"`

	// Âncora a um biomarcador (código do lab_test_definitions) e/ou a um score item.
	LabTestCode *string    `gorm:"type:varchar(20);column:lab_test_code" json:"labTestCode,omitempty"`
	ScoreItemID *uuid.UUID `gorm:"type:uuid" json:"scoreItemId,omitempty"`

	// Snapshot que motivou a recomendação (rastreabilidade do "porquê").
	SourceSnapshotID *uuid.UUID `gorm:"type:uuid" json:"sourceSnapshotId,omitempty"`

	Priority CarePlanPriority   `gorm:"type:varchar(10);not null;default:'medium';check:priority IN ('high','medium','low')" json:"priority"`
	Status   CarePlanItemStatus `gorm:"type:varchar(10);not null;default:'active';check:status IN ('active','achieved','suspended')" json:"status"`

	AuthorUserID uuid.UUID `gorm:"type:uuid;not null" json:"authorUserId"`
	AuthorRole   string    `gorm:"type:varchar(30)" json:"authorRole"`

	CreatedAt time.Time      `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Patient    Patient `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE" json:"-"`
	AuthorUser User    `gorm:"foreignKey:AuthorUserID;constraint:OnDelete:RESTRICT" json:"authorUser,omitempty"`
}

func (CarePlanItem) TableName() string { return "care_plan_items" }

func (c *CarePlanItem) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
