package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// LabTestCategory define categorias de exames laboratoriais
type LabTestCategory string

const (
	LabTestCategoryHematology   LabTestCategory = "hematology"    // Hemograma, coagulação
	LabTestCategoryBiochemistry LabTestCategory = "biochemistry"  // Glicose, ureia, creatinina
	LabTestCategoryHormones     LabTestCategory = "hormones"      // TSH, T4, testosterona
	LabTestCategoryImmunology   LabTestCategory = "immunology"    // Sorologias, autoimunes
	LabTestCategoryMicrobiology LabTestCategory = "microbiology"  // Culturas, antibiogramas
	LabTestCategoryUrine        LabTestCategory = "urine"         // EAS, urocultura
	LabTestCategoryImaging      LabTestCategory = "imaging"       // Raio-X, TC, RM
	LabTestCategoryFunctional   LabTestCategory = "functional"    // Medicina funcional
	LabTestCategoryGenetics     LabTestCategory = "genetics"      // Testes genéticos
	LabTestCategoryOther        LabTestCategory = "other"         // Outros
)

// LabTestDefinition representa a definição de um exame laboratorial ou parâmetro
// @Description Definição estruturada de exames e seus parâmetros
type LabTestDefinition struct {
	// ID único da definição
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`

	// Código interno único (usado para referência programática)
	// @example HEMOGRAMA_COMPLETO, HGB, GLUCOSE_FASTING
	Code string `gorm:"type:varchar(100);not null;unique;index" json:"code" validate:"required,max=100"`

	// Nome do exame/parâmetro
	// @example Hemograma Completo, Hemoglobina, Glicemia de Jejum
	Name string `gorm:"type:varchar(300);not null" json:"name" validate:"required,max=300"`

	// Nome curto/abreviação (opcional)
	// @example Hemograma, Hb, Gli
	ShortName *string `gorm:"type:varchar(50)" json:"shortName,omitempty"`

	// Código TUSS (Terminologia Unificada da Saúde Suplementar)
	// Usado para faturamento e solicitação no Brasil
	// @example 40304485
	TussCode *string `gorm:"type:varchar(20);index" json:"tussCode,omitempty"`

	// Código LOINC (Logical Observation Identifiers Names and Codes)
	// Padrão internacional para identificação de exames
	// @example 718-7 (Hemoglobin)
	LoincCode *string `gorm:"type:varchar(20);index" json:"loincCode,omitempty"`

	// Categoria do exame
	// @enum hematology,biochemistry,hormones,immunology,microbiology,urine,imaging,functional,genetics,other
	Category LabTestCategory `gorm:"type:varchar(30);not null;index" json:"category" validate:"required"`

	// Indica se o exame pode ser solicitado individualmente
	// true: pode ser solicitado (ex: Hemograma Completo, Glicemia)
	// false: só aparece como resultado de outro exame (ex: Hemoglobina, Bilirrubina Indireta)
	IsRequestable bool `gorm:"type:boolean;not null;default:true;index" json:"isRequestable"`

	// ID do exame pai (hierarquia)
	// Ex: Hemoglobina tem parentTestId = Hemograma Completo
	// Ex: Bilirrubina Indireta tem parentTestId = Bilirrubina Total e Frações
	ParentTestID *uuid.UUID `gorm:"type:uuid;index" json:"parentTestId,omitempty"`

	// Unidade de medida padrão
	// @example g/dL, mg/dL, mU/L, %
	Unit *string `gorm:"type:varchar(50)" json:"unit,omitempty"`

	// Fórmula de conversão entre unidades (se aplicável)
	// @example 1 g/dL = 10 g/L
	UnitConversion *string `gorm:"type:text" json:"unitConversion,omitempty"`

	// Tipo de resultado
	// @enum numeric, text, boolean, categorical
	ResultType string `gorm:"type:varchar(20);not null;default:'numeric'" json:"resultType" validate:"required"`

	// Método de coleta/realização
	// @example Sangue venoso com jejum de 8-12h
	CollectionMethod *string `gorm:"type:text" json:"collectionMethod,omitempty"`

	// Tempo de jejum necessário (em horas)
	FastingHours *int `gorm:"type:integer" json:"fastingHours,omitempty"`

	// Material biológico
	// @example Sangue total, Soro, Urina, Fezes
	SpecimenType *string `gorm:"type:varchar(100)" json:"specimenType,omitempty"`

	// Descrição/observações sobre o exame
	Description *string `gorm:"type:text" json:"description,omitempty"`

	// Indicações clínicas principais
	ClinicalIndications *string `gorm:"type:text" json:"clinicalIndications,omitempty"`

	// Ordem de exibição (para organizar parâmetros de um exame)
	DisplayOrder int `gorm:"type:integer;not null;default:0" json:"displayOrder"`

	// Status (ativo/inativo)
	IsActive bool `gorm:"type:boolean;not null;default:true;index" json:"isActive"`

	// Timestamps
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relacionamentos

	// Exame pai (se for um parâmetro/subexame)
	ParentTest *LabTestDefinition `gorm:"foreignKey:ParentTestID;constraint:OnDelete:SET NULL" json:"parentTest,omitempty"`

	// Subexames/parâmetros (se for um exame composto)
	SubTests []LabTestDefinition `gorm:"foreignKey:ParentTestID;constraint:OnDelete:SET NULL" json:"subTests,omitempty"`

	// Mapeamentos para itens do escore
	ScoreMappings []LabTestScoreMapping `gorm:"foreignKey:LabTestID;constraint:OnDelete:CASCADE" json:"scoreMappings,omitempty"`
}

// TableName especifica o nome da tabela
func (LabTestDefinition) TableName() string {
	return "lab_test_definitions"
}

// BeforeCreate hook to generate UUID v7
func (ltd *LabTestDefinition) BeforeCreate(tx *gorm.DB) error {
	if ltd.ID == uuid.Nil {
		ltd.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// LabTestScoreMapping mapeia um exame laboratorial para um item do escore
// @Description Mapeamento entre exame e item do escore Plenya
type LabTestScoreMapping struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`

	// ID do exame laboratorial
	LabTestID uuid.UUID `gorm:"type:uuid;not null;index:idx_lab_test_score" json:"labTestId" validate:"required"`

	// ID do item do escore
	ScoreItemID uuid.UUID `gorm:"type:uuid;not null;index:idx_lab_test_score" json:"scoreItemId" validate:"required"`

	// Especificação de gênero (se o mapeamento for específico)
	// @enum male, female, null (ambos)
	Gender *Gender `gorm:"type:varchar(10)" json:"gender,omitempty"`

	// Faixa etária mínima (em anos)
	MinAge *int `gorm:"type:integer" json:"minAge,omitempty"`

	// Faixa etária máxima (em anos)
	MaxAge *int `gorm:"type:integer" json:"maxAge,omitempty"`

	// Observações sobre o mapeamento
	Notes *string `gorm:"type:text" json:"notes,omitempty"`

	// Status
	IsActive bool `gorm:"type:boolean;not null;default:true" json:"isActive"`

	// Timestamps
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relacionamentos
	LabTest   LabTestDefinition `gorm:"foreignKey:LabTestID" json:"labTest,omitempty"`
	ScoreItem ScoreItem         `gorm:"foreignKey:ScoreItemID" json:"scoreItem,omitempty"`
}

// TableName especifica o nome da tabela
func (LabTestScoreMapping) TableName() string {
	return "lab_test_score_mappings"
}
