package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// WorkoutPlanSession representa uma sessão dentro de um plano de treino
// @Description Sessão de treino (ex: "Treino A", "Treino B")
type WorkoutPlanSession struct {
	// ID único da sessão
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// ID do plano de treino
	PlanID uuid.UUID `gorm:"type:uuid;not null;index" json:"planId"`

	// Nome da sessão
	Name string `gorm:"type:varchar(100);not null" json:"name" validate:"required,min=1,max=100"`

	// Ordem de exibição
	Order int `gorm:"type:int;not null;default:0" json:"order"`

	// Timestamps
	CreatedAt time.Time      `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relações
	Plan      WorkoutPlan              `gorm:"foreignKey:PlanID;constraint:OnDelete:CASCADE" json:"plan,omitempty"`
	Exercises []WorkoutSessionExercise `gorm:"foreignKey:SessionID" json:"exercises,omitempty"`
}

// TableName especifica o nome da tabela
func (WorkoutPlanSession) TableName() string {
	return "workout_plan_sessions"
}

// BeforeCreate hook to generate UUID v7
func (s *WorkoutPlanSession) BeforeCreate(tx *gorm.DB) error {
	if s.ID == uuid.Nil {
		s.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
