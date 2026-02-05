package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ScoreItem represents a specific clinical parameter (e.g., "Hemoglobina - Homens", "FEVE")
// @Description Item de escore - parâmetro clínico específico
type ScoreItem struct {
	// @example 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Código para identificação programática (pode ser compartilhado entre múltiplos score_items)
	// @example HGB_M, GLUCOSE_FASTING, IMC, BREASTFEEDING_DURATION
	Code *string `gorm:"type:varchar(100);index" json:"code,omitempty" validate:"omitempty,max=100"`

	// @minLength 2
	// @maxLength 300
	// @example Hemoglobina - Homens
	Name string `gorm:"type:varchar(300);not null" json:"name" validate:"required,min=2,max=300"`

	// @example g/dL
	Unit *string `gorm:"type:varchar(50)" json:"unit,omitempty" validate:"omitempty,max=50"`

	// @example 1 g/dL = 10 g/L
	UnitConversion *string `gorm:"type:text" json:"unitConversion,omitempty"`

	// Relevância clínica - explicação técnica para profissionais de saúde
	// @example Valores baixos de hemoglobina indicam anemia, que pode estar associada a fadiga, redução da capacidade funcional e aumento do risco cardiovascular
	ClinicalRelevance *string `gorm:"type:text" json:"clinicalRelevance,omitempty"`

	// Explicação para o paciente - linguagem simples e acessível
	// @example Hemoglobina é a proteína que transporta oxigênio no sangue. Valores baixos podem causar cansaço e falta de ar
	PatientExplanation *string `gorm:"type:text" json:"patientExplanation,omitempty"`

	// Conduta clínica recomendada
	// @example Investigar causa da anemia (deficiência de ferro, B12, folato, doença crônica). Suplementação conforme indicação. Encaminhar ao hematologista se Hb < 10 g/dL ou causa não esclarecida
	Conduct *string `gorm:"type:text" json:"conduct,omitempty"`

	// Data da última revisão dos campos clínicos ou artigos associados
	// @example 2026-01-25T10:30:00Z
	LastReview *time.Time `gorm:"type:timestamp" json:"lastReview,omitempty"`

	// @minimum 0
	// @maximum 100
	// @example 20
	Points *float64 `gorm:"type:double precision" json:"points,omitempty" validate:"omitempty,gte=0,lte=100"`

	// @minimum 0
	// @maximum 9999
	// @example 1
	Order int `gorm:"type:integer;not null;default:0;index:idx_score_item_order" json:"order" validate:"gte=0,lte=9999"`

	// Foreign Keys
	// @example 550e8400-e29b-41d4-a716-446655440000
	SubgroupID uuid.UUID `gorm:"type:uuid;not null;index:idx_score_item_subgroup" json:"subgroupId" validate:"required"`

	// Self-referencing for hierarchical items (optional)
	// @example 550e8400-e29b-41d4-a716-446655440000
	ParentItemID *uuid.UUID `gorm:"type:uuid;index:idx_score_item_parent" json:"parentItemId,omitempty"`

	// Relationships
	Subgroup   *ScoreSubgroup `gorm:"foreignKey:SubgroupID;constraint:OnDelete:CASCADE" json:"subgroup,omitempty"`
	ParentItem *ScoreItem     `gorm:"foreignKey:ParentItemID;constraint:OnDelete:SET NULL" json:"parentItem,omitempty"`
	ChildItems []ScoreItem    `gorm:"foreignKey:ParentItemID;constraint:OnDelete:SET NULL" json:"childItems,omitempty"`
	Levels     []ScoreLevel   `gorm:"foreignKey:ItemID;constraint:OnDelete:CASCADE" json:"levels,omitempty"`

	// Many-to-many relationship with Articles
	// @items.type object
	Articles []Article `gorm:"many2many:article_score_items;" json:"articles,omitempty"`

	// Timestamps
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for ScoreItem
func (ScoreItem) TableName() string {
	return "score_items"
}

// BeforeCreate hook to generate UUID v7
func (si *ScoreItem) BeforeCreate(tx *gorm.DB) error {
	if si.ID == uuid.Nil {
		si.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// BeforeUpdate hook to update LastReview when clinical fields change
func (si *ScoreItem) BeforeUpdate(tx *gorm.DB) error {
	// Check if any clinical field was changed
	if tx.Statement.Changed("ClinicalRelevance") ||
		tx.Statement.Changed("PatientExplanation") ||
		tx.Statement.Changed("Conduct") {
		now := time.Now()
		si.LastReview = &now
	}
	return nil
}
