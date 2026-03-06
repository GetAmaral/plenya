package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// MesocyclePhase define as fases de um mesociclo
type MesocyclePhase string

const (
	MesocycleAccumulation   MesocyclePhase = "accumulation"
	MesocycleTransformation MesocyclePhase = "transformation"
	MesocycleRealization    MesocyclePhase = "realization"
	MesocycleHypertrophy    MesocyclePhase = "hypertrophy"
	MesocycleStrength       MesocyclePhase = "strength"
	MesocycleEndurance      MesocyclePhase = "endurance"
	MesocyclePower          MesocyclePhase = "power"
)

// WorkoutMesocycle representa um mesociclo dentro de uma periodização
// @Description Mesociclo com fase, duração, volume e intensidade
type WorkoutMesocycle struct {
	// ID único do mesociclo
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// ID da periodização
	PeriodizationID uuid.UUID `gorm:"type:uuid;not null;index" json:"periodizationId"`

	// Ordem dentro da periodização
	Order int `gorm:"type:int;not null" json:"order"`

	// Fase do mesociclo
	Phase MesocyclePhase `gorm:"type:varchar(20);not null;check:phase IN ('accumulation','transformation','realization','hypertrophy','strength','endurance','power')" json:"phase" validate:"required,oneof=accumulation transformation realization hypertrophy strength endurance power"`

	// Duração em semanas
	DurationWeeks int `gorm:"type:int;not null" json:"durationWeeks" validate:"required,min=1,max=52"`

	// Percentual de volume
	VolumePercent int `gorm:"type:int;not null" json:"volumePercent"`

	// Percentual de intensidade
	IntensityPercent int `gorm:"type:int;not null" json:"intensityPercent"`

	// Foco fisiológico
	PhysiologicalFocus string `gorm:"type:varchar(100);not null" json:"physiologicalFocus"`

	// Data de início
	StartDate time.Time `gorm:"type:date;not null" json:"startDate"`

	// Data de término
	EndDate time.Time `gorm:"type:date;not null" json:"endDate"`

	// Timestamps
	CreatedAt time.Time      `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relações
	Periodization WorkoutPeriodization `gorm:"foreignKey:PeriodizationID;constraint:OnDelete:CASCADE" json:"periodization,omitempty"`
}

// TableName especifica o nome da tabela
func (WorkoutMesocycle) TableName() string {
	return "workout_mesocycles"
}

// BeforeCreate hook to generate UUID v7
func (m *WorkoutMesocycle) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
