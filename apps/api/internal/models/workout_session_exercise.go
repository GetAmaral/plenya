package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ExercisePhase define a fase do exercício dentro da sessão
type ExercisePhase string

const (
	ExercisePhaseWarmup   ExercisePhase = "warmup"
	ExercisePhaseMain     ExercisePhase = "main"
	ExercisePhaseCooldown ExercisePhase = "cooldown"
)

// ExerciseCadence define a cadência do exercício
type ExerciseCadence string

const (
	ExerciseCadenceNormal    ExerciseCadence = "normal"
	ExerciseCadenceSlow      ExerciseCadence = "slow"
	ExerciseCadencePaused    ExerciseCadence = "paused"
	ExerciseCadenceExplosive ExerciseCadence = "explosive"
	ExerciseCadenceFree      ExerciseCadence = "free"
)

// WorkoutSessionExercise representa um exercício dentro de uma sessão de treino
// @Description Exercício com configuração de séries, repetições, cadência e descanso
type WorkoutSessionExercise struct {
	// ID único
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// ID da sessão
	SessionID uuid.UUID `gorm:"type:uuid;not null;index" json:"sessionId"`

	// ID do exercício
	ExerciseID uuid.UUID `gorm:"type:uuid;not null" json:"exerciseId"`

	// Fase do exercício
	Phase ExercisePhase `gorm:"type:varchar(10);not null;check:phase IN ('warmup','main','cooldown')" json:"phase" validate:"required,oneof=warmup main cooldown"`

	// Ordem de exibição dentro da fase
	Order int `gorm:"type:int;not null;default:0" json:"order"`

	// Número de séries
	Sets int `gorm:"type:int;not null;default:3" json:"sets"`

	// Repetições (pode ser "8-12", "45s", "15")
	Reps string `gorm:"type:varchar(20);not null;default:'10'" json:"reps"`

	// Cadência do exercício
	Cadence ExerciseCadence `gorm:"type:varchar(15);not null;default:'normal';check:cadence IN ('normal','slow','paused','explosive','free')" json:"cadence"`

	// Descanso entre séries (segundos)
	RestBetweenSetsSec int `gorm:"type:int;not null;default:60" json:"restBetweenSetsSec"`

	// Descanso entre exercícios (segundos)
	RestBetweenExercisesSec int `gorm:"type:int;not null;default:90" json:"restBetweenExercisesSec"`

	// Observações
	Notes *string `gorm:"type:text" json:"notes,omitempty"`

	// Timestamps
	CreatedAt time.Time      `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relações
	Session  WorkoutPlanSession `gorm:"foreignKey:SessionID;constraint:OnDelete:CASCADE" json:"session,omitempty"`
	Exercise Exercise           `gorm:"foreignKey:ExerciseID;constraint:OnDelete:RESTRICT" json:"exercise,omitempty"`
}

// TableName especifica o nome da tabela
func (WorkoutSessionExercise) TableName() string {
	return "workout_session_exercises"
}

// BeforeCreate hook to generate UUID v7
func (e *WorkoutSessionExercise) BeforeCreate(tx *gorm.DB) error {
	if e.ID == uuid.Nil {
		e.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
