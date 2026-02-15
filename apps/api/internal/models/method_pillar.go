package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// MethodPillar represents a specific pillar/category within a letter
// @Description Pilar específico dentro de uma letra (ex: "Avaliação Nutricional" na letra A)
type MethodPillar struct {
	// Primary Key
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Nome do pilar
	// @minLength 3
	// @maxLength 300
	// @example Avaliação Nutricional
	Name string `gorm:"type:varchar(300);not null" json:"name" validate:"required,min=3,max=300"`

	// Descrição
	Description *string `gorm:"type:text" json:"description,omitempty"`

	// Relevância clínica
	ClinicalRelevance *string `gorm:"type:text" json:"clinicalRelevance,omitempty"`

	// Explicação para paciente
	PatientExplanation *string `gorm:"type:text" json:"patientExplanation,omitempty"`

	// Conduta clínica
	Conduct *string `gorm:"type:text" json:"conduct,omitempty"`

	// Última revisão
	LastReview *time.Time `gorm:"type:timestamp" json:"lastReview,omitempty"`

	// Ordem de exibição
	Order int `gorm:"type:integer;not null;default:0;index:idx_method_pillar_order" json:"order" validate:"gte=0,lte=9999"`

	// Foreign Key (required - CASCADE delete)
	LetterID uuid.UUID `gorm:"type:uuid;not null;index:idx_method_pillar_letter" json:"letterId" validate:"required"`

	// Relationships
	Letter     *MethodLetter `gorm:"foreignKey:LetterID;constraint:OnDelete:CASCADE" json:"letter,omitempty"`
	ScoreItems []ScoreItem   `gorm:"many2many:score_item_method_pillars;" json:"scoreItems,omitempty"`

	// Timestamps
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (MethodPillar) TableName() string {
	return "method_pillars"
}

func (mp *MethodPillar) BeforeCreate(tx *gorm.DB) error {
	if mp.ID == uuid.Nil {
		mp.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

func (mp *MethodPillar) BeforeUpdate(tx *gorm.DB) error {
	if tx.Statement.Changed("ClinicalRelevance") ||
		tx.Statement.Changed("PatientExplanation") ||
		tx.Statement.Changed("Conduct") {
		now := time.Now()
		mp.LastReview = &now
	}
	return nil
}
