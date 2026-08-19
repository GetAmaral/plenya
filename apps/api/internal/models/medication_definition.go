package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// MedicationCategory define as categorias regulatórias de medicamentos
type MedicationCategory string

const (
	MedCategorySimple     MedicationCategory = "simple"     // Receita simples
	MedCategoryC1         MedicationCategory = "c1"         // Controle Especial (Lista C1)
	MedCategoryC5         MedicationCategory = "c5"         // Psicotrópicos (Lista C5)
	MedCategoryAntibiotic MedicationCategory = "antibiotic" // Antimicrobianos (RDC 471)
	MedCategoryGLP1       MedicationCategory = "glp1"       // GLP-1 agonistas
	// MedCategoryAB — tarja preta: Notificação de Receita A (amarela) ou B (azul) da Portaria
	// 344/98. O EMR NÃO emite Notificação de Receita, então estes entram no catálogo com
	// IsPrescribable=false: servem para reconciliar medicação em uso, não para prescrever.
	MedCategoryAB MedicationCategory = "a_b"
)

// Procedência da linha do catálogo.
const (
	MedSourceManual = "manual" // criada à mão pelo admin
	MedSourceCMED   = "cmed"   // importada da Lista de Preços de Medicamentos (ANVISA/CMED)
)

// Como a categoria regulatória foi definida — o que separa fato de palpite.
const (
	MedCategorySourceManual   = "manual"        // curada por humano
	MedCategorySourceDerived  = "cmed_derived"  // derivada com regra defensável
	MedCategorySourceFallback = "cmed_fallback" // chute conservador: a fonte não permitia afirmar
)

// Tarja como a CMED publica. NÃO é a Portaria 344 — é o proxy mais próximo que a fonte oferece.
const (
	MedStripeRed           = "vermelha"
	MedStripeRedRestricted = "vermelha_restrita"
	MedStripeBlack         = "preta"
	MedStripeNone          = "isento"
)

// MedicationDefinition representa a definição de um medicamento no catálogo
// @Description Definição de medicamento com regras regulatórias (ANVISA)
type MedicationDefinition struct {
	// ID único da definição
	// @example 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Nome comercial comum do medicamento
	// @minLength 3
	// @maxLength 500
	// @example Fluoxetina 20mg
	CommonName string `gorm:"type:varchar(500);not null;index" json:"commonName" validate:"required,min=3,max=500"`

	// Princípio ativo (DCB - Denominação Comum Brasileira)
	// @minLength 3
	// @maxLength 500
	// @example Cloridrato de Fluoxetina
	ActiveIngredient string `gorm:"type:text;not null" json:"activeIngredient" validate:"required,min=3"`

	// Categoria regulatória do medicamento
	// @enum simple,c1,c5,antibiotic,glp1,a_b
	Category MedicationCategory `gorm:"type:varchar(20);not null;index" json:"category" validate:"required,oneof=simple c1 c5 antibiotic glp1 a_b"`

	// Regras de validação baseadas na categoria

	// Dias de validade da prescrição
	// @example 30
	// C1: 30 dias, Antibióticos: 10 dias, GLP1: 90 dias, Simples: 30 dias
	ValidityDays int `gorm:"type:integer;not null;default:30" json:"validityDays" validate:"required,min=1,max=365"`

	// Máximo de substâncias por prescrição
	// @example 3
	// C1: máximo 3 substâncias
	MaxPerPrescription int `gorm:"type:integer;not null;default:10" json:"maxPerPrescription" validate:"required,min=1,max=100"`

	// Duração máxima do tratamento em dias
	// @example 60
	// C1: 60 dias, Anticonvulsivantes: 180 dias
	MaxTreatmentDays int `gorm:"type:integer;not null;default:60" json:"maxTreatmentDays" validate:"required,min=1,max=365"`

	// Requer assinatura digital ICP-Brasil
	// true para medicamentos controlados (C1, C5)
	RequiresDigitalSignature bool `gorm:"type:boolean;not null;default:false" json:"requiresDigitalSignature"`

	// Requer registro no SNCR (Sistema Nacional de Controle de Receitas)
	// true para medicamentos controlados
	RequiresSNCR bool `gorm:"type:boolean;not null;default:false" json:"requiresSNCR"`

	// Metadata

	// Código ANVISA do medicamento
	// @example 1234567890123
	ANVISACode *string `gorm:"type:varchar(50)" json:"anvisaCode,omitempty"`

	// ── Proveniência (import CMED) ──────────────────────────────────────────────────

	// GGREM — chave natural da CMED, usada na idempotência do reimport mensal.
	// NULL nas linhas criadas à mão, que por isso ficam imunes ao import.
	GGREM *string `gorm:"type:varchar(20)" json:"ggrem,omitempty"`

	// @enum manual,cmed
	Source string `gorm:"type:varchar(10);not null;default:'manual'" json:"source"`

	// Edição da CMED que produziu a linha (YYYYMM). Sem ela, o preço mente sobre a data.
	SourceVersion  *string    `gorm:"type:varchar(8)" json:"sourceVersion,omitempty"`
	LastImportedAt *time.Time `json:"lastImportedAt,omitempty"`

	// false = sumiu da lista publicada. Nunca apagamos (há FK de prescrições antigas).
	IsActive bool `gorm:"type:boolean;not null;default:true" json:"isActive"`

	// ── Identidade comercial ────────────────────────────────────────────────────────

	// Texto da ANVISA com embalagem + forma + concentração ("500 MG COM REV CT BL X 30").
	Presentation         *string  `gorm:"type:text" json:"presentation,omitempty"`
	Laboratory           *string  `gorm:"type:varchar(200)" json:"laboratory,omitempty"`
	ProductType          *string  `gorm:"type:varchar(60)" json:"productType,omitempty"` // Genérico|Similar|Novo|...
	EAN13                *string  `gorm:"type:varchar(14)" json:"ean13,omitempty"`
	TherapeuticClass     *string  `gorm:"type:varchar(200)" json:"therapeuticClass,omitempty"`
	TherapeuticClassCode *string  `gorm:"type:varchar(10)" json:"therapeuticClassCode,omitempty"`
	Stripe               *string  `gorm:"type:varchar(20)" json:"stripe,omitempty"`
	PMCPrice             *float64 `gorm:"type:numeric(12,2)" json:"pmcPrice,omitempty"`

	// ── Derivados do texto de apresentação ──────────────────────────────────────────

	Concentration      *string `gorm:"type:varchar(120)" json:"concentration,omitempty"`
	PharmaceuticalForm *string `gorm:"type:varchar(60)" json:"pharmaceuticalForm,omitempty"`
	Route              *string `gorm:"type:varchar(40)" json:"route,omitempty"`
	PackageQuantity    *int    `json:"packageQuantity,omitempty"`

	// @enum high,medium,none
	DerivationConfidence *string `gorm:"type:varchar(10)" json:"derivationConfidence,omitempty"`

	// ── Honestidade da classificação e curadoria ────────────────────────────────────

	// Lista da Portaria 344/98 (A1..C5). A CMED NÃO traz: só curadoria preenche. Quando
	// preenchida, manda sobre Category.
	ControlList *string `gorm:"type:varchar(4)" json:"controlList,omitempty"`

	// @enum manual,cmed_derived,cmed_fallback
	CategorySource string `gorm:"type:varchar(14);not null;default:'manual'" json:"categorySource"`

	// Classificação derivada de forma imperfeita — a UI avisa antes de prescrever.
	NeedsReview bool `gorm:"type:boolean;not null;default:false" json:"needsReview"`

	// false = fora do autocomplete de receita, mas presente no catálogo.
	IsPrescribable bool `gorm:"type:boolean;not null;default:true" json:"isPrescribable"`

	// Enquanto CuratedAt != nil, o reimport mensal só atualiza campos de FONTE e não encosta
	// nos clínicos — é o que faz a correção do médico sobreviver.
	CuratedAt *time.Time `json:"curatedAt,omitempty"`
	CuratedBy *uuid.UUID `gorm:"type:uuid" json:"curatedBy,omitempty"`

	// Timestamps
	CreatedAt time.Time      `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName especifica o nome da tabela
func (MedicationDefinition) TableName() string {
	return "medication_definitions"
}

// BeforeCreate hook to generate UUID v7
func (m *MedicationDefinition) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
