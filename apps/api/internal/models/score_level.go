package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ScoreLevel represents a risk stratification level for a ScoreItem (e.g., "Nível 0: <40% FEVE")
// @Description Nível de estratificação de risco para um item de escore
type ScoreLevel struct {
	// @example 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// @minimum 0
	// @maximum 6
	// @example 5
	Level int `gorm:"type:integer;not null;check:level >= 0 AND level <= 6;index:idx_score_level_level" json:"level" validate:"required,gte=0,lte=6"`

	// @minLength 1
	// @maxLength 500
	// @example 55 a 70 (Ótimo)
	Name string `gorm:"type:varchar(500);not null" json:"name" validate:"required,min=1,max=500"`

	// @example <40
	LowerLimit *string `gorm:"type:varchar(50)" json:"lowerLimit,omitempty" validate:"omitempty,max=50"`

	// @example 49.9
	UpperLimit *string `gorm:"type:varchar(50)" json:"upperLimit,omitempty" validate:"omitempty,max=50"`

	// @enum =,>,>=,<,<=,between
	// @example between
	Operator string `gorm:"type:varchar(10);not null;check:operator IN ('=','>','>=','<','<=','between');default:'between'" json:"operator" validate:"required,oneof== > >= < <= between"`

	// Relevância clínica - explicação técnica para profissionais de saúde
	// @example FEVE entre 55-70% indica função cardíaca normal. Redução abaixo de 50% sugere disfunção sistólica leve, com risco aumentado de insuficiência cardíaca
	ClinicalRelevance *string `gorm:"type:text" json:"clinicalRelevance,omitempty"`

	// Explicação para o paciente - linguagem simples e acessível
	// @example Seu coração está bombeando sangue de forma eficiente. Este é um resultado normal e saudável
	PatientExplanation *string `gorm:"type:text" json:"patientExplanation,omitempty"`

	// Conduta clínica recomendada
	// @example Manter acompanhamento regular. Otimizar controle de fatores de risco cardiovascular (pressão arterial, diabetes, colesterol). Considerar ecocardiograma de controle em 12 meses
	Conduct *string `gorm:"type:text" json:"conduct,omitempty"`

	// Data da última revisão dos campos clínicos
	// @example 2026-01-25T10:30:00Z
	LastReview *time.Time `gorm:"type:timestamp" json:"lastReview,omitempty"`

	// Foreign Keys
	// @example 550e8400-e29b-41d4-a716-446655440000
	ItemID uuid.UUID `gorm:"type:uuid;not null;index:idx_score_level_item" json:"itemId" validate:"required"`

	// Relationships
	Item *ScoreItem `gorm:"foreignKey:ItemID;constraint:OnDelete:CASCADE" json:"item,omitempty"`

	// Timestamps
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for ScoreLevel
func (ScoreLevel) TableName() string {
	return "score_levels"
}

// BeforeCreate hook to generate UUID v7
func (sl *ScoreLevel) BeforeCreate(tx *gorm.DB) error {
	if sl.ID == uuid.Nil {
		sl.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// BeforeUpdate hook to update LastReview when clinical fields change
func (sl *ScoreLevel) BeforeUpdate(tx *gorm.DB) error {
	// Check if any clinical field was changed
	if tx.Statement.Changed("ClinicalRelevance") ||
		tx.Statement.Changed("PatientExplanation") ||
		tx.Statement.Changed("Conduct") {
		now := time.Now()
		sl.LastReview = &now
	}
	return nil
}

// EvaluatesTrue verifica se um valor numérico satisfaz o condicional deste level
func (sl *ScoreLevel) EvaluatesTrue(value float64) bool {
	parseFloat := func(s *string) float64 {
		if s == nil {
			return 0
		}
		var result float64
		// Ignora erro, retorna 0 se parsing falhar
		_, _ = fmt.Sscanf(*s, "%f", &result)
		return result
	}

	switch sl.Operator {
	case "=":
		if sl.LowerLimit == nil {
			return false
		}
		return value == parseFloat(sl.LowerLimit)

	case ">":
		if sl.LowerLimit == nil {
			return false
		}
		return value > parseFloat(sl.LowerLimit)

	case ">=":
		if sl.LowerLimit == nil {
			return false
		}
		return value >= parseFloat(sl.LowerLimit)

	case "<":
		if sl.LowerLimit == nil {
			return false
		}
		return value < parseFloat(sl.LowerLimit)

	case "<=":
		if sl.LowerLimit == nil {
			return false
		}
		return value <= parseFloat(sl.LowerLimit)

	case "between":
		if sl.LowerLimit == nil || sl.UpperLimit == nil {
			return false
		}
		lower := parseFloat(sl.LowerLimit)
		upper := parseFloat(sl.UpperLimit)
		return value >= lower && value <= upper

	default:
		return false
	}
}
