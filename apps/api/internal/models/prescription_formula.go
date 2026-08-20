package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PrescriptionType distingue a receita de industrializado da receita de manipulado.
type PrescriptionType string

const (
	PrescriptionCommercial PrescriptionType = "commercial"
	PrescriptionCompounded PrescriptionType = "compounded"
)

// FormulaUsage — uso interno (ingerido) ou externo (aplicado na pele/mucosa). Vai em destaque no
// receituário magistral porque é a informação que evita o erro de administração mais grave.
type FormulaUsage string

const (
	FormulaUsageInternal FormulaUsage = "internal"
	FormulaUsageExternal FormulaUsage = "external"
)

// PrescriptionFormula representa UMA fórmula magistral dentro de uma prescrição de manipulado.
// Uma receita pode ter várias (ex.: cápsula do dia + cápsula da noite + creme).
// @Description Fórmula magistral de uma prescrição de manipulado
type PrescriptionFormula struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// ID da prescrição (FK)
	// @required
	PrescriptionID uuid.UUID `gorm:"type:uuid;not null;index" json:"prescriptionId" validate:"required"`

	// Fórmula-base que originou esta fórmula, quando veio de um template curado.
	TemplateID *uuid.UUID `gorm:"type:uuid" json:"templateId,omitempty"`

	// Ordem de impressão no receituário
	DisplayOrder int `gorm:"type:integer;not null;default:0" json:"displayOrder"`

	// Nome de uso da fórmula (opcional): "Fórmula do sono", "Cápsula da manhã"
	Name string `gorm:"type:varchar(200);not null;default:''" json:"name"`

	// Forma farmacêutica a manipular
	// @example cápsula
	PharmaceuticalForm string `gorm:"type:varchar(60);not null" json:"pharmaceuticalForm" validate:"required"`

	// Uso interno ou externo
	// @enum internal,external
	UsageType FormulaUsage `gorm:"type:varchar(10);not null;default:'internal'" json:"usageType"`

	// Via de administração (opcional; o uso já diz o essencial em manipulado tópico)
	Route string `gorm:"type:varchar(40);not null;default:''" json:"route"`

	// Veículo / excipiente ("qsp 1 cápsula", "creme não iônico qsp 60 g")
	Vehicle string `gorm:"type:varchar(200);not null;default:''" json:"vehicle"`

	// Quantidade a aviar. Decimal porque manipulado se avia em gramas e mililitros, não só em
	// unidades inteiras ("60 g de creme", "0,5 mL").
	QuantityToDispense float64 `gorm:"type:numeric(12,3);not null" json:"quantityToDispense" validate:"required,gt=0"`

	// Unidade do aviamento: cápsulas, sachês, g, mL
	QuantityUnit string `gorm:"type:varchar(30);not null" json:"quantityUnit" validate:"required"`

	// Quantidade por extenso (obrigatória quando há componente controlado)
	QuantityInWords string `gorm:"type:varchar(200);not null;default:''" json:"quantityInWords"`

	// Posologia ("1 cápsula ao deitar")
	Posology string `gorm:"type:varchar(300);not null;default:''" json:"posology"`

	// Duração do tratamento em dias (0 = não informada)
	Duration int `gorm:"type:integer;not null;default:0" json:"duration"`

	// Orientações específicas desta fórmula
	Instructions *string `gorm:"type:text" json:"instructions,omitempty"`

	// Categoria mais restritiva entre os componentes — decide modo de assinatura e rótulo de
	// controle especial. Denormalizada de propósito (ver migration 00068).
	HighestCategory MedicationCategory `gorm:"type:varchar(20);not null;default:'simple'" json:"highestCategory"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Prescription *Prescription                  `gorm:"foreignKey:PrescriptionID;constraint:OnDelete:CASCADE" json:"prescription,omitempty"`
	Components   []PrescriptionFormulaComponent `gorm:"foreignKey:FormulaID;constraint:OnDelete:CASCADE" json:"components,omitempty"`
}

func (PrescriptionFormula) TableName() string { return "prescription_formulas" }

func (f *PrescriptionFormula) BeforeCreate(tx *gorm.DB) error {
	if f.ID == uuid.Nil {
		f.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// PrescriptionFormulaComponent é uma substância dentro de uma fórmula magistral.
// @Description Componente (substância) de uma fórmula magistral
type PrescriptionFormulaComponent struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// ID da fórmula (FK)
	// @required
	FormulaID uuid.UUID `gorm:"type:uuid;not null;index" json:"formulaId" validate:"required"`

	// Vínculo com o catálogo magistral (fase seguinte) ou com o catálogo de industrializados.
	MagistralComponentID   *uuid.UUID `gorm:"type:uuid" json:"magistralComponentId,omitempty"`
	MedicationDefinitionID *uuid.UUID `gorm:"type:uuid" json:"medicationDefinitionId,omitempty"`

	DisplayOrder int `gorm:"type:integer;not null;default:0" json:"displayOrder"`

	// Substância (DCB quando existir)
	// @example Melatonina
	Substance string `gorm:"type:varchar(200);not null" json:"substance" validate:"required,min=2,max=200"`

	// Quantidade por dose/unidade aviada. Decimal: manipulado usa 0,25 mg e 12,5 mcg.
	Quantity float64 `gorm:"type:numeric(14,4);not null" json:"quantity" validate:"required,gt=0"`

	// Unidade: mg, g, mcg, UI, mL, %
	Unit string `gorm:"type:varchar(20);not null" json:"unit" validate:"required"`

	// Categoria de controle do componente
	Category MedicationCategory `gorm:"type:varchar(20);not null;default:'simple'" json:"category"`

	// Observação impressa ao lado do componente ("liberação prolongada", "microencapsulado")
	Note string `gorm:"type:varchar(200);not null;default:''" json:"note"`

	// O que o motor de regras sugeriu, quando houve sugestão. Fica ao lado do que foi prescrito
	// para revelar, depois, quando o médico discordou. Não sai impresso.
	SuggestedQuantity *float64 `gorm:"type:numeric(14,4)" json:"suggestedQuantity,omitempty"`

	// A dose escrita é do ELEMENTO (ou do ativo puro) e não do insumo. Muda o que a farmácia pesa
	// e o volume que o pó ocupa na cápsula.
	AsElemental bool `gorm:"not null;default:false" json:"asElemental"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Formula *PrescriptionFormula `gorm:"foreignKey:FormulaID;constraint:OnDelete:CASCADE" json:"formula,omitempty"`
}

func (PrescriptionFormulaComponent) TableName() string { return "prescription_formula_components" }

func (c *PrescriptionFormulaComponent) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// categoryRank ordena as categorias por restrição: quem manda na receita é a mais restritiva.
var categoryRank = map[MedicationCategory]int{
	MedCategorySimple:     0,
	MedCategoryGLP1:       1,
	MedCategoryAntibiotic: 2,
	MedCategoryC5:         3,
	MedCategoryC1:         4,
	MedCategoryAB:         5,
}

// HighestCategoryOf devolve a categoria mais restritiva de uma lista de componentes.
func HighestCategoryOf(components []PrescriptionFormulaComponent) MedicationCategory {
	highest := MedCategorySimple
	for _, c := range components {
		if categoryRank[c.Category] > categoryRank[highest] {
			highest = c.Category
		}
	}
	return highest
}

// IsControlled diz se a categoria exige receituário de controle especial (e, portanto, receita
// impressa e assinada à mão enquanto o SNCR não abre).
func IsControlled(category MedicationCategory) bool {
	return category == MedCategoryC1 || category == MedCategoryC5 || category == MedCategoryAB
}
