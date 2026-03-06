package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PeriodizationFramework define os frameworks de periodização
type PeriodizationFramework string

const (
	PeriodizationBompa      PeriodizationFramework = "bompa"
	PeriodizationLinear     PeriodizationFramework = "linear"
	PeriodizationUndulating PeriodizationFramework = "undulating"
	PeriodizationBlock      PeriodizationFramework = "block"
)

// WorkoutPeriodization representa uma periodização de treino (macrociclo)
// @Description Periodização de treino com mesociclos
type WorkoutPeriodization struct {
	// ID único da periodização
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// ID do paciente
	PatientID uuid.UUID `gorm:"type:uuid;not null;index" json:"patientId"`

	// ID do profissional que criou
	CreatedByID uuid.UUID `gorm:"type:uuid;not null" json:"createdById"`

	// Framework de periodização
	Framework PeriodizationFramework `gorm:"type:varchar(20);not null;check:framework IN ('bompa','linear','undulating','block')" json:"framework" validate:"required,oneof=bompa linear undulating block"`

	// Data de início
	StartDate time.Time `gorm:"type:date;not null" json:"startDate"`

	// Total de semanas
	TotalWeeks int `gorm:"type:int;not null" json:"totalWeeks" validate:"required,min=1,max=104"`

	// Objetivo
	Objective string `gorm:"type:varchar(200);not null" json:"objective" validate:"required,min=2,max=200"`

	// Justificativa científica (gerada por IA/RAG)
	ScientificJustification *string `gorm:"type:text" json:"scientificJustification"`

	// Título computado para exibição no frontend (não persistido)
	DisplayTitle string `gorm:"-" json:"displayTitle"`

	// Timestamps
	CreatedAt time.Time      `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relações
	Patient    Patient            `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE" json:"patient,omitempty"`
	CreatedBy  User               `gorm:"foreignKey:CreatedByID;constraint:OnDelete:RESTRICT" json:"createdBy,omitempty"`
	Mesocycles []WorkoutMesocycle `gorm:"foreignKey:PeriodizationID" json:"mesocycles,omitempty"`
}

// TableName especifica o nome da tabela
func (WorkoutPeriodization) TableName() string {
	return "workout_periodizations"
}

// GetTitle retorna um título legível para a periodização
func (wp *WorkoutPeriodization) GetTitle() string {
	frameworkLabels := map[PeriodizationFramework]string{
		PeriodizationBompa:      "Bompa",
		PeriodizationLinear:     "Linear",
		PeriodizationUndulating: "Ondulatória",
		PeriodizationBlock:      "Bloco",
	}
	label, ok := frameworkLabels[wp.Framework]
	if !ok {
		label = string(wp.Framework)
	}
	return fmt.Sprintf("Periodização %s - %s (%d semanas)", label, wp.Objective, wp.TotalWeeks)
}

// AfterFind popula DisplayTitle após carregar do banco
func (wp *WorkoutPeriodization) AfterFind(tx *gorm.DB) error {
	wp.DisplayTitle = wp.GetTitle()
	return nil
}

// BeforeCreate hook to generate UUID v7
func (wp *WorkoutPeriodization) BeforeCreate(tx *gorm.DB) error {
	if wp.ID == uuid.Nil {
		wp.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
