package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// WorkoutSessionExerciseLog é o log set-a-set de UM exercício dentro de uma
// WorkoutSession. Granularidade fina pra gerar histórico de progressão (peso
// por set, RPE, dor) que vira gráfico no app e no continuum do profissional.
//
// @Description Log de execução de um set de exercício
type WorkoutSessionExerciseLog struct {
	// ID único
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Sessão de execução
	SessionID uuid.UUID `gorm:"type:uuid;not null;index;uniqueIndex:uniq_session_exercise_set" json:"sessionId"`

	// Exercício planejado (FK pro WorkoutSessionExercise template)
	PlanExerciseID uuid.UUID `gorm:"type:uuid;not null;index;uniqueIndex:uniq_session_exercise_set" json:"planExerciseId"`

	// Exercício real executado (FK pro Exercise — denormalizado pra histórico
	// sobreviver remoção do plano)
	ExerciseID uuid.UUID `gorm:"type:uuid;not null;index" json:"exerciseId"`

	// Número do set (1-indexed)
	SetNumber int `gorm:"type:int;not null;uniqueIndex:uniq_session_exercise_set;check:set_number >= 1 AND set_number <= 20" json:"setNumber" validate:"required,min=1,max=20"`

	// Repetições executadas (pode ser 0 quando é tempo)
	Reps *int `gorm:"type:int" json:"reps,omitempty"`

	// Peso (kg) — nullable para exercícios com peso corporal
	Weight *float64 `gorm:"type:numeric(6,2)" json:"weight,omitempty"`

	// Duração (s) — pra exercícios isométricos / cardio
	DurationSec *int `gorm:"type:int" json:"durationSec,omitempty"`

	// RPE 1-10 (escala de esforço percebido)
	RPE *int `gorm:"type:int;check:rpe IS NULL OR (rpe >= 1 AND rpe <= 10)" json:"rpe,omitempty"`

	// Notas do paciente sobre o set
	Notes *string `gorm:"type:text" json:"notes,omitempty"`

	// Timestamps
	CreatedAt time.Time      `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relações
	Session      WorkoutSession         `gorm:"foreignKey:SessionID;constraint:OnDelete:CASCADE" json:"-"`
	PlanExercise WorkoutSessionExercise `gorm:"foreignKey:PlanExerciseID;constraint:OnDelete:CASCADE" json:"-"`
	Exercise     Exercise               `gorm:"foreignKey:ExerciseID;constraint:OnDelete:RESTRICT" json:"exercise,omitempty"`
}

func (WorkoutSessionExerciseLog) TableName() string { return "workout_session_exercise_logs" }

func (l *WorkoutSessionExerciseLog) BeforeCreate(tx *gorm.DB) error {
	if l.ID == uuid.Nil {
		l.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
