package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// LabResult representa um resultado de exame laboratorial individual
// @Description Resultado individual de um exame dentro de um lote (batch)
type LabResult struct {
	// ID único do resultado
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// ID do lote de resultados (obrigatório)
	LabResultBatchID uuid.UUID `gorm:"type:uuid;not null;index:idx_result_batch" json:"labResultBatchId" validate:"required"`

	// ID da definição do teste (opcional - FK para LabTestDefinition)
	LabTestDefinitionID *uuid.UUID `gorm:"type:uuid;index:idx_result_test_def" json:"labTestDefinitionId,omitempty"`

	// Nome do exame
	TestName string `gorm:"type:varchar(200);not null" json:"testName" validate:"required"`

	// Tipo de exame (ex: hemograma, glicemia, etc)
	TestType string `gorm:"type:varchar(100);not null;index" json:"testType" validate:"required"`

	// Resultado em texto (para exames qualitativos)
	ResultText *string `gorm:"type:text" json:"resultText,omitempty"`

	// Resultado numérico (para exames quantitativos) - CONVERTIDO para unidade padrão
	ResultNumeric *float64 `gorm:"type:decimal(12,4)" json:"resultNumeric,omitempty"`

	// Unidade de medida - CONVERTIDA para unidade padrão
	Unit *string `gorm:"type:varchar(50)" json:"unit,omitempty"`

	// Resultado numérico ORIGINAL (antes da conversão de unidade)
	ResultNumericOriginal *float64 `gorm:"type:decimal(12,4)" json:"resultNumericOriginal,omitempty"`

	// Unidade ORIGINAL (antes da conversão)
	UnitOriginal *string `gorm:"type:varchar(50)" json:"unitOriginal,omitempty"`

	// O que aconteceu com a unidade na ingestão: "ok" (já veio certa), "convertido" ou
	// "revisar". Sem este campo um resultado que não converteu é indistinguível de um que
	// chegou correto, e foi assim que resultados ficaram gravados na unidade do laudo sendo
	// comparados, lá na frente, contra uma escala em outra grandeza.
	UnitConversionStatus *string `gorm:"type:varchar(20);index" json:"unitConversionStatus,omitempty"`

	// Por que não converteu, em português, para a fila de curadoria se explicar sozinha.
	UnitConversionNote *string `gorm:"type:text" json:"unitConversionNote,omitempty"`

	// Interpretação/Observações específicas deste teste
	Interpretation *string `gorm:"type:text" json:"interpretation,omitempty"`

	// Nível/Prioridade do resultado (opcional)
	Level *int `gorm:"type:integer" json:"level,omitempty"`

	// Indica se o exame foi matched com uma definição catalogada
	// true = matched com LabTestDefinition, false = extraído mas não catalogado.
	// SEM `default:true` de propósito: com default, o GORM omite o valor `false`
	// (zero-value) no INSERT e o banco aplicava true, anulando o "Não catalogado".
	Matched bool `gorm:"type:boolean;not null;index" json:"matched"`

	// Origem do resultado: "pdf" (importado de laudo) ou "manual" (digitado).
	Source string `gorm:"type:varchar(20);not null;default:'manual'" json:"source"`

	// Motivo quando o exame não casou com o catálogo (ex: "Não encontrado no catálogo
	// de exames"). NULL quando casou. Usado para mostrar ao usuário por que ficou de fora.
	MatchReason *string `gorm:"type:text" json:"matchReason,omitempty"`

	// Motivo quando o exame NÃO recebeu nível de classificação (Level nulo): sem item de
	// score, não se aplica ao paciente, valor fora das faixas, não numérico, etc. NULL
	// quando classificado. Eixo diferente de MatchReason (catálogo vs score).
	ClassifyReason *string `gorm:"type:text" json:"classifyReason,omitempty"`

	// Material/espécime do exame (ex: "Sangue", "Soro", "Urina", "Fezes"). Extraído do laudo
	// (campo `material`), normalizado. Desambigua exames de mesmo nome em espécimes diferentes
	// (glicose sangue vs urina) e enriquece o prontuário. Opcional.
	Specimen *string `gorm:"type:varchar(50)" json:"specimen,omitempty"`

	// Método analítico do exame (ex: "ECLIA", "HPLC", "Enzimático"). Extraído do laudo. Opcional.
	Method *string `gorm:"type:varchar(150)" json:"method,omitempty"`

	// Faixa de referência impressa pelo laboratório (texto livre, ex: "70 a 99 mg/dL"). Para
	// exibição e fallback de normal/alterado quando não há item de score. Opcional.
	ReferenceRange *string `gorm:"type:text" json:"referenceRange,omitempty"`

	// Data de coleta DESTE exame (os laudos trazem "Coletado em:" por exame, que pode diferir
	// da data do lote). A tabela de resultados renderiza por esta data; se NULL, herda a do
	// lote (coalesce no read). Opcional.
	CollectionDate *time.Time `gorm:"type:timestamptz" json:"collectionDate,omitempty"`

	// Data de criação
	CreatedAt time.Time `gorm:"not null;autoCreateTime" json:"createdAt"`

	// Data de atualização
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime" json:"updatedAt"`

	// Data de deleção (soft delete)
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relações
	LabResultBatch    LabResultBatch     `gorm:"foreignKey:LabResultBatchID;constraint:OnDelete:CASCADE" json:"labResultBatch,omitempty"`
	LabTestDefinition *LabTestDefinition `gorm:"foreignKey:LabTestDefinitionID;constraint:OnDelete:SET NULL" json:"labTestDefinition,omitempty"`
	LabResultValues   []LabResultValue   `gorm:"foreignKey:LabResultID;constraint:OnDelete:CASCADE" json:"labResultValues,omitempty"`
}

// TableName especifica o nome da tabela
func (LabResult) TableName() string {
	return "lab_results"
}

// BeforeCreate hook to generate UUID v7
func (lr *LabResult) BeforeCreate(tx *gorm.DB) error {
	if lr.ID == uuid.Nil {
		lr.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
