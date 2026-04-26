package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// WorkoutSession representa a EXECUÇÃO real de uma sessão de treino feita por
// um paciente. Difere de WorkoutPlanSession (template definido pelo profissional):
// uma plan-session pode gerar N workout-sessions ao longo do tempo.
//
// @Description Sessão de treino executada pelo paciente
type WorkoutSession struct {
	// ID único
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Paciente que executou
	PatientID uuid.UUID `gorm:"type:uuid;not null;index" json:"patientId"`

	// Plano de treino origem (FK)
	PlanID uuid.UUID `gorm:"type:uuid;not null;index" json:"planId"`

	// Sessão template do plano (FK) — qual "Treino A/B" foi feito
	PlanSessionID uuid.UUID `gorm:"type:uuid;not null;index" json:"planSessionId"`

	// Data planejada (pode ser hoje ou retroativa quando o paciente loga atrasado)
	ScheduledDate time.Time `gorm:"type:date;not null;index" json:"scheduledDate"`

	// Quando o paciente terminou (nil = ainda em andamento)
	CompletedAt *time.Time `json:"completedAt,omitempty"`

	// Notas do paciente
	Notes *string `gorm:"type:text" json:"notes,omitempty"`

	// Timestamps
	CreatedAt time.Time      `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relações
	Patient     Patient                      `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE" json:"-"`
	Plan        WorkoutPlan                  `gorm:"foreignKey:PlanID;constraint:OnDelete:CASCADE" json:"plan,omitempty"`
	PlanSession WorkoutPlanSession           `gorm:"foreignKey:PlanSessionID;constraint:OnDelete:CASCADE" json:"planSession,omitempty"`
	Logs        []WorkoutSessionExerciseLog  `gorm:"foreignKey:SessionID" json:"logs,omitempty"`
}

func (WorkoutSession) TableName() string { return "workout_sessions" }

func (s *WorkoutSession) BeforeCreate(tx *gorm.DB) error {
	if s.ID == uuid.Nil {
		s.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
