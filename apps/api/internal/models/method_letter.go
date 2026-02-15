package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// MethodLetter represents a letter/phase within a method (e.g., A, G, I, R in AGIR)
// @Description Letra/fase de uma metodologia (ex: A, G, I, R no AGIR)
type MethodLetter struct {
	// Primary Key
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Código da letra (ex: "A", "G", "I", "R")
	// @minLength 1
	// @maxLength 10
	// @example A
	Code string `gorm:"type:varchar(10);not null;index:idx_method_letter_code" json:"code" validate:"required,min=1,max=10"`

	// Nome completo (ex: "Alimentação e Atividade Física")
	// @minLength 3
	// @maxLength 300
	// @example Alimentação e Atividade Física
	Name string `gorm:"type:varchar(300);not null" json:"name" validate:"required,min=3,max=300"`

	// Descrição
	Description *string `gorm:"type:text" json:"description,omitempty"`

	// Relevância clínica
	ClinicalRelevance *string `gorm:"type:text" json:"clinicalRelevance,omitempty"`

	// Explicação para paciente
	PatientExplanation *string `gorm:"type:text" json:"patientExplanation,omitempty"`

	// Conduta clínica
	Conduct *string `gorm:"type:text" json:"conduct,omitempty"`

	// Última revisão dos campos clínicos
	LastReview *time.Time `gorm:"type:timestamp" json:"lastReview,omitempty"`

	// Cor (para UI)
	// @example #10B981
	Color *string `gorm:"type:varchar(7)" json:"color,omitempty" validate:"omitempty,hexcolor"`

	// Ícone/emoji
	// @example 🥗
	// @maxLength 50
	Icon *string `gorm:"type:varchar(50)" json:"icon,omitempty" validate:"omitempty,max=50"`

	// Ordem de exibição
	Order int `gorm:"type:integer;not null;default:0;index:idx_method_letter_order" json:"order" validate:"gte=0,lte=9999"`

	// Foreign Key (required - CASCADE delete)
	MethodID uuid.UUID `gorm:"type:uuid;not null;index:idx_method_letter_method" json:"methodId" validate:"required"`

	// Relationships
	Method  *Method        `gorm:"foreignKey:MethodID;constraint:OnDelete:CASCADE" json:"method,omitempty"`
	Pillars []MethodPillar `gorm:"foreignKey:LetterID;constraint:OnDelete:CASCADE" json:"pillars,omitempty"`

	// Timestamps
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (MethodLetter) TableName() string {
	return "method_letters"
}

func (ml *MethodLetter) BeforeCreate(tx *gorm.DB) error {
	if ml.ID == uuid.Nil {
		ml.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

func (ml *MethodLetter) BeforeUpdate(tx *gorm.DB) error {
	if tx.Statement.Changed("ClinicalRelevance") ||
		tx.Statement.Changed("PatientExplanation") ||
		tx.Statement.Changed("Conduct") {
		now := time.Now()
		ml.LastReview = &now
	}
	return nil
}
