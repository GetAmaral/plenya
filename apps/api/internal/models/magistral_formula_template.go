package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// MagistralFormulaTemplate é uma fórmula-base: a fórmula pronta que o médico reusa entre
// pacientes, com indicação própria. A receita emitida guarda de qual base saiu.
// @Description Fórmula magistral base (template) reutilizável
type MagistralFormulaTemplate struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Nome de uso ("Fórmula do sono")
	Name string `gorm:"type:varchar(200);not null" json:"name" validate:"required,min=2,max=200"`

	// Para que serve esta fórmula — em texto e em tópicos, mesma dupla do catálogo de
	// substâncias: o tópico é para decidir, o texto é para conferir.
	Indication        *string `gorm:"type:text" json:"indication,omitempty"`
	IndicationBullets *string `gorm:"type:text" json:"indicationBullets,omitempty"`

	PharmaceuticalForm string       `gorm:"type:varchar(60);not null" json:"pharmaceuticalForm" validate:"required"`
	UsageType          FormulaUsage `gorm:"type:varchar(10);not null;default:'internal'" json:"usageType"`
	Route              string       `gorm:"type:varchar(40);not null;default:''" json:"route"`
	Vehicle            string       `gorm:"type:varchar(200);not null;default:''" json:"vehicle"`
	QuantityToDispense float64      `gorm:"type:numeric(12,3);not null;default:60" json:"quantityToDispense"`
	QuantityUnit       string       `gorm:"type:varchar(30);not null;default:'cápsulas'" json:"quantityUnit"`
	Posology           string       `gorm:"type:varchar(300);not null;default:''" json:"posology"`
	Duration           int          `gorm:"type:integer;not null;default:0" json:"duration"`
	Instructions       *string      `gorm:"type:text" json:"instructions,omitempty"`
	Notes              *string      `gorm:"type:text" json:"notes,omitempty"`

	// Quantas receitas nasceram desta base — ordena a lista pelo repertório real.
	UsageCount int  `gorm:"type:integer;not null;default:0" json:"usageCount"`
	IsActive   bool `gorm:"not null;default:true" json:"isActive"`

	CreatedBy  *uuid.UUID `gorm:"type:uuid" json:"createdBy,omitempty"`
	ReviewedBy *uuid.UUID `gorm:"type:uuid" json:"reviewedBy,omitempty"`
	LastReview *time.Time `gorm:"type:timestamptz" json:"lastReview,omitempty"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Components []MagistralFormulaTemplateComponent `gorm:"foreignKey:TemplateID;constraint:OnDelete:CASCADE" json:"components,omitempty"`
}

func (MagistralFormulaTemplate) TableName() string { return "magistral_formula_templates" }

func (t *MagistralFormulaTemplate) BeforeCreate(tx *gorm.DB) error {
	if t.ID == uuid.Nil {
		t.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// MagistralFormulaTemplateComponent é uma substância da fórmula-base, com a dose padrão.
type MagistralFormulaTemplateComponent struct {
	ID                   uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	TemplateID           uuid.UUID  `gorm:"type:uuid;not null;index" json:"templateId"`
	MagistralComponentID *uuid.UUID `gorm:"type:uuid" json:"magistralComponentId,omitempty"`
	DisplayOrder         int        `gorm:"type:integer;not null;default:0" json:"displayOrder"`

	Substance string             `gorm:"type:varchar(200);not null" json:"substance" validate:"required"`
	Quantity  float64            `gorm:"type:numeric(14,4);not null" json:"quantity" validate:"required,gt=0"`
	Unit      string             `gorm:"type:varchar(20);not null" json:"unit" validate:"required"`
	Category  MedicationCategory `gorm:"type:varchar(20);not null;default:'simple'" json:"category"`
	Note      string             `gorm:"type:varchar(200);not null;default:''" json:"note"`
	// A dose da base é do elemento, não do insumo (ver fator de correção).
	AsElemental bool `gorm:"not null;default:false" json:"asElemental"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Rule *MagistralFormulaTemplateRule `gorm:"foreignKey:TemplateComponentID;constraint:OnDelete:CASCADE" json:"rule,omitempty"`
}

func (MagistralFormulaTemplateComponent) TableName() string {
	return "magistral_formula_template_components"
}

func (c *MagistralFormulaTemplateComponent) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// DoseRuleKind — como a dose do componente é calculada quando há regra.
type DoseRuleKind string

const (
	DoseRuleFixed        DoseRuleKind = "fixed"
	DoseRulePerKg        DoseRuleKind = "per_kg"
	DoseRuleLabThreshold DoseRuleKind = "lab_threshold"
	// lab_band é o formato que a clínica realmente usa: conduta por faixa do exame, não por
	// um único corte. Vitamina D, ferritina e B12 não têm um limiar, têm três.
	DoseRuleLabBand DoseRuleKind = "lab_band"
)

// MagistralFormulaTemplateRule é a regra de dose de UM componente da fórmula-base.
//
// MinDose e MaxDose não são opcionais: toda regra é travada nos dois extremos. Peso errado no
// prontuário ou exame em unidade trocada não conseguem produzir sugestão absurda — a trava
// corta, e a resposta diz que cortou.
type MagistralFormulaTemplateRule struct {
	ID                  uuid.UUID    `gorm:"type:uuid;primaryKey" json:"id"`
	TemplateComponentID uuid.UUID    `gorm:"type:uuid;not null;index" json:"templateComponentId"`
	Kind                DoseRuleKind `gorm:"type:varchar(16);not null" json:"kind"`

	// per_kg
	PerKg *float64 `gorm:"type:numeric(12,4)" json:"perKg,omitempty"`

	// lab_threshold — o exame é uma referência ao catálogo (lab_test_definitions.code),
	// escolhida na tela. Nenhum código clínico entra hardcoded no Go.
	LabCode      *string  `gorm:"type:varchar(64)" json:"labCode,omitempty"`
	LabOperator  *string  `gorm:"type:varchar(4)" json:"labOperator,omitempty"`
	LabThreshold *float64 `gorm:"type:numeric(14,4)" json:"labThreshold,omitempty"`
	DoseIfTrue   *float64 `gorm:"type:numeric(14,4)" json:"doseIfTrue,omitempty"`
	DoseIfFalse  *float64 `gorm:"type:numeric(14,4)" json:"doseIfFalse,omitempty"`

	// A unidade em que o limiar (ou a faixa) foi escrito. Resultado gravado em OUTRA unidade
	// não vira sugestão: 31% dos resultados numéricos do banco estão numa unidade diferente da
	// definição do exame, incluindo cortisol em nmol/L sobre definição em µg/dL.
	LabUnit *string `gorm:"type:varchar(50)" json:"labUnit,omitempty"`

	// fixed
	FixedDose *float64 `gorm:"type:numeric(14,4)" json:"fixedDose,omitempty"`

	// Passo prático de arredondamento: 1.234,5678 mg não se manipula. NULL = não arredonda.
	RoundTo *float64 `gorm:"type:numeric(14,4)" json:"roundTo,omitempty"`

	MinDose float64 `gorm:"type:numeric(14,4);not null" json:"minDose" validate:"required,gt=0"`
	MaxDose float64 `gorm:"type:numeric(14,4);not null" json:"maxDose" validate:"required,gt=0"`

	// Dado mais velho que isto não sugere nada. Dose calculada sobre peso de três anos atrás é
	// risco, não conveniência.
	MaxDataAgeDays int    `gorm:"type:integer;not null;default:365" json:"maxDataAgeDays"`
	Note           string `gorm:"type:varchar(300);not null;default:''" json:"note"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Bands []MagistralFormulaTemplateRuleBand `gorm:"foreignKey:RuleID;constraint:OnDelete:CASCADE" json:"bands,omitempty"`
}

func (MagistralFormulaTemplateRule) TableName() string {
	return "magistral_formula_template_rules"
}

// MagistralFormulaTemplateRuleBand é uma faixa do exame com a dose correspondente.
//
// Intervalo MEIO-ABERTO (LowerBound, UpperBound], a mesma convenção já canônica nas faixas do
// escore — nil em Lower é -infinito, nil em Upper é +infinito. Duas convenções de faixa no mesmo
// sistema seria pedir para alguém ler a errada.
type MagistralFormulaTemplateRuleBand struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	RuleID       uuid.UUID `gorm:"type:uuid;not null;index" json:"ruleId"`
	DisplayOrder int       `gorm:"type:integer;not null;default:0" json:"displayOrder"`

	LowerBound *float64 `gorm:"type:numeric(14,4)" json:"lowerBound,omitempty"`
	UpperBound *float64 `gorm:"type:numeric(14,4)" json:"upperBound,omitempty"`

	Dose float64 `gorm:"type:numeric(14,4);not null" json:"dose" validate:"required,gt=0"`
	// Como a faixa aparece na frase que o médico lê: "deficiência grave", "insuficiência".
	Label string `gorm:"type:varchar(120);not null;default:''" json:"label"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (MagistralFormulaTemplateRuleBand) TableName() string {
	return "magistral_formula_template_rule_bands"
}

func (b *MagistralFormulaTemplateRuleBand) BeforeCreate(tx *gorm.DB) error {
	if b.ID == uuid.Nil {
		b.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

func (r *MagistralFormulaTemplateRule) BeforeCreate(tx *gorm.DB) error {
	if r.ID == uuid.Nil {
		r.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
