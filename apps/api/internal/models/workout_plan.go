package models

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// WorkoutObjective define os objetivos de treino
type WorkoutObjective string

const (
	WorkoutObjectiveHypertrophy    WorkoutObjective = "hypertrophy"
	WorkoutObjectiveStrength       WorkoutObjective = "strength"
	WorkoutObjectiveEndurance      WorkoutObjective = "endurance"
	WorkoutObjectiveWeightLoss     WorkoutObjective = "weight_loss"
	WorkoutObjectiveConditioning   WorkoutObjective = "conditioning"
	WorkoutObjectiveRehabilitation WorkoutObjective = "rehabilitation"
)

// WorkoutIntensity define os níveis de intensidade
type WorkoutIntensity string

const (
	WorkoutIntensityVeryLight WorkoutIntensity = "very_light"
	WorkoutIntensityLight     WorkoutIntensity = "light"
	WorkoutIntensityModerate  WorkoutIntensity = "moderate"
	WorkoutIntensityHigh      WorkoutIntensity = "high"
	WorkoutIntensityVeryHigh  WorkoutIntensity = "very_high"
)

// WorkoutPlan representa um plano de treino
// @Description Plano de treino com sessões e exercícios
type WorkoutPlan struct {
	// ID único do plano
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// ID do paciente
	PatientID uuid.UUID `gorm:"type:uuid;not null;index" json:"patientId"`

	// ID do profissional que criou
	CreatedByID uuid.UUID `gorm:"type:uuid;not null" json:"createdById"`

	// Nome do plano
	Name string `gorm:"type:varchar(200);not null" json:"name" validate:"required,min=2,max=200"`

	// Objetivo do treino
	Objective WorkoutObjective `gorm:"type:varchar(30);not null;check:objective IN ('hypertrophy','strength','endurance','weight_loss','conditioning','rehabilitation')" json:"objective" validate:"required,oneof=hypertrophy strength endurance weight_loss conditioning rehabilitation"`

	// Intensidade
	Intensity WorkoutIntensity `gorm:"type:varchar(20);not null;check:intensity IN ('very_light','light','moderate','high','very_high')" json:"intensity" validate:"required,oneof=very_light light moderate high very_high"`

	// Duração em minutos
	DurationMinutes int `gorm:"type:int;not null;default:60" json:"durationMinutes"`

	// Frequência semanal
	WeeklyFrequency int `gorm:"type:int;not null;default:3" json:"weeklyFrequency"`

	// Código público para compartilhamento (8 chars alfanuméricos, case-sensitive)
	PublicCode string `gorm:"type:varchar(8);uniqueIndex;not null" json:"publicCode"`

	// HTML pré-renderizado do treino (com GIFs base64 inline)
	HtmlContent *string `gorm:"type:text" json:"htmlContent,omitempty"`

	// Ativo
	IsActive bool `gorm:"not null;default:true" json:"isActive"`

	// Título computado para exibição no frontend (não persistido)
	DisplayTitle string `gorm:"-" json:"displayTitle"`

	// Timestamps
	CreatedAt time.Time      `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relações
	Patient   Patient              `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE" json:"patient,omitempty"`
	CreatedBy User                 `gorm:"foreignKey:CreatedByID;constraint:OnDelete:RESTRICT" json:"createdBy,omitempty"`
	Sessions  []WorkoutPlanSession `gorm:"foreignKey:PlanID" json:"sessions,omitempty"`
}

// TableName especifica o nome da tabela
func (WorkoutPlan) TableName() string {
	return "workout_plans"
}

// GetTitle retorna um título legível para o plano
func (w *WorkoutPlan) GetTitle() string {
	objectiveLabels := map[WorkoutObjective]string{
		WorkoutObjectiveHypertrophy:    "Hipertrofia",
		WorkoutObjectiveStrength:       "Força",
		WorkoutObjectiveEndurance:      "Resistência",
		WorkoutObjectiveWeightLoss:     "Emagrecimento",
		WorkoutObjectiveConditioning:   "Condicionamento",
		WorkoutObjectiveRehabilitation: "Reabilitação",
	}
	label, ok := objectiveLabels[w.Objective]
	if !ok {
		label = string(w.Objective)
	}
	return fmt.Sprintf("%s - %s", w.Name, label)
}

// AfterFind popula DisplayTitle após carregar do banco
func (w *WorkoutPlan) AfterFind(tx *gorm.DB) error {
	w.DisplayTitle = w.GetTitle()
	return nil
}

// BeforeCreate hook to generate UUID v7 and PublicCode
func (w *WorkoutPlan) BeforeCreate(tx *gorm.DB) error {
	if w.ID == uuid.Nil {
		w.ID = uuid.Must(uuid.NewV7())
	}
	if w.PublicCode == "" {
		code, err := generatePublicCode(8)
		if err != nil {
			return err
		}
		w.PublicCode = code
	}
	return nil
}

// generatePublicCode gera um código alfanumérico aleatório (case-sensitive)
func generatePublicCode(length int) (string, error) {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, length)
	for i := range result {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[n.Int64()]
	}
	return string(result), nil
}
