package dto

import "github.com/plenya/api/internal/models"

// DTOs do catálogo magistral (componentes, incompatibilidades e a checagem de fórmula).

// MagistralComponentRequest — criação/edição de um componente. Campos clínicos são ponteiros:
// ausente significa "não cadastrado", que a tela mostra como silêncio e não como zero.
type MagistralComponentRequest struct {
	Name        string  `json:"name" validate:"required,min=2,max=200"`
	Synonyms    string  `json:"synonyms,omitempty" validate:"max=2000"`
	CAS         *string `json:"cas,omitempty" validate:"omitempty,max=20"`
	DCBCode     *string `json:"dcbCode,omitempty" validate:"omitempty,max=20"`
	DefaultUnit string  `json:"defaultUnit,omitempty" validate:"max=20"`

	UsualDose *float64 `json:"usualDose,omitempty" validate:"omitempty,gt=0"`
	MinDose   *float64 `json:"minDose,omitempty" validate:"omitempty,gt=0"`
	MaxDose   *float64 `json:"maxDose,omitempty" validate:"omitempty,gt=0"`

	// Densidade aparente (g/mL) — sem ela a calculadora de cápsula não opina.
	BulkDensity      *float64 `json:"bulkDensity,omitempty" validate:"omitempty,gt=0"`
	DensitySource    *string  `json:"densitySource,omitempty" validate:"omitempty,max=60"`
	Brand            *string  `json:"brand,omitempty" validate:"omitempty,max=60"`
	ElementalPercent *float64 `json:"elementalPercent,omitempty" validate:"omitempty,gt=0,lte=100"`
	DoseAsElemental  *bool    `json:"doseAsElemental,omitempty"`
	CorrectionNote   *string  `json:"correctionNote,omitempty" validate:"omitempty,max=300"`

	EutecticFormer     bool `json:"eutecticFormer"`
	Hygroscopic        bool `json:"hygroscopic"`
	Oxidizing          bool `json:"oxidizing"`
	OxidationSensitive bool `json:"oxidationSensitive"`
	Photosensitive     bool `json:"photosensitive"`

	Bitterness        *int    `json:"bitterness,omitempty" validate:"omitempty,min=0,max=3"`
	SachetOK          *bool   `json:"sachetOk,omitempty"`
	Notes             *string `json:"notes,omitempty"`
	Indications       *string `json:"indications,omitempty"`
	DoseReference     *string `json:"doseReference,omitempty"`
	IndicationBullets *string `json:"indicationBullets,omitempty"`
	DoseBullets       *string `json:"doseBullets,omitempty"`
	Source            string  `json:"source,omitempty" validate:"max=30"`
}

// MagistralDefaultDoseRequest — "salvar como padrão desta substância" (curadoria oportunista).
type MagistralDefaultDoseRequest struct {
	Substance string  `json:"substance" validate:"required,min=2,max=200"`
	Dose      float64 `json:"dose" validate:"required,gt=0"`
	Unit      string  `json:"unit" validate:"required,max=20"`
}

// MagistralIncompatibilityRequest — par curado de substâncias que não convivem bem.
type MagistralIncompatibilityRequest struct {
	ComponentAID string  `json:"componentAId" validate:"required,uuid"`
	ComponentBID string  `json:"componentBId" validate:"required,uuid"`
	Severity     string  `json:"severity" validate:"required,oneof=info warn avoid"`
	Mechanism    string  `json:"mechanism" validate:"required,max=200"`
	Note         *string `json:"note,omitempty"`
	Source       string  `json:"source,omitempty" validate:"max=200"`
}

// MagistralCheckComponent — um componente da fórmula que está sendo montada na tela.
type MagistralCheckComponent struct {
	MagistralComponentID *string `json:"magistralComponentId,omitempty"`
	Substance            string  `json:"substance" validate:"required"`
	Quantity             float64 `json:"quantity"`
	Unit                 string  `json:"unit"`
	// A dose escrita é do elemento (ou do ativo puro), não do insumo.
	AsElemental bool `json:"asElemental"`
}

// MagistralCheckRequest — a fórmula como está na tela.
type MagistralCheckRequest struct {
	PharmaceuticalForm string                    `json:"pharmaceuticalForm" validate:"required,max=60"`
	Components         []MagistralCheckComponent `json:"components" validate:"required,min=1,max=20,dive"`
	// Posologia da fórmula, de onde sai o número de tomadas por dia. A conta é do backend: a
	// mesma regra escrita duas vezes diverge com o tempo.
	Posology string `json:"posology,omitempty" validate:"max=300"`
	// Override explícito das tomadas por dia. 0 = deduzir da posologia.
	DosesPerDay float64 `json:"dosesPerDay,omitempty" validate:"omitempty,gte=0,lte=12"`
	// Veículo/base da fórmula. É contra ele que rodam as regras de incompatibilidade de base.
	Vehicle string `json:"vehicle,omitempty" validate:"max=200"`
}

// MagistralAlertResponse — aviso sobre a fórmula. Avisa, não bloqueia.
type MagistralAlertResponse struct {
	Level      string   `json:"level"`
	Kind       string   `json:"kind"`
	Substances []string `json:"substances"`
	Message    string   `json:"message"`
}

// MagistralComponentMatch — o que o catálogo sabe de cada componente digitado. A tela usa para
// oferecer "salvar como padrão" exatamente onde falta dado.
type MagistralComponentMatch struct {
	ID          *string  `json:"id,omitempty"`
	Substance   string   `json:"substance"`
	Known       bool     `json:"known"`
	HasDose     bool     `json:"hasDose"`
	HasDensity  bool     `json:"hasDensity"`
	UsualDose   *float64 `json:"usualDose,omitempty"`
	DefaultUnit string   `json:"defaultUnit,omitempty"`
	// Indicação e posologia de referência do catálogo (curadas ou extraídas do material), para o
	// médico ver enquanto monta a fórmula sem precisar abrir outra tela.
	Indications       *string `json:"indications,omitempty"`
	DoseReference     *string `json:"doseReference,omitempty"`
	IndicationBullets *string `json:"indicationBullets,omitempty"`
	DoseBullets       *string `json:"doseBullets,omitempty"`
	// suggested = veio do RAG e ainda não foi conferido. A tela precisa dizer isso.
	EvidenceStatus string `json:"evidenceStatus,omitempty"`

	// Conversão elemento ↔ insumo, quando o insumo é diluído ou quelado.
	ElementalPercent *float64 `json:"elementalPercent,omitempty"`
	DoseAsElemental  bool     `json:"doseAsElemental"`
	RawMaterialMg    *float64 `json:"rawMaterialMg,omitempty"`
	CorrectionNote   *string  `json:"correctionNote,omitempty"`

	// Forma que o prescritor usa no lugar desta.
	PreferredName  *string `json:"preferredName,omitempty"`
	PreferenceNote *string `json:"preferenceNote,omitempty"`
	// Categoria de receita que a substância carrega, para a tela preencher o componente.
	DefaultCategory string `json:"defaultCategory,omitempty"`
	// Restrição de finalidade ou exigência regulatória.
	RegulatoryNote string `json:"regulatoryNote,omitempty"`
}

// MagistralCheckResponse — avisos + cálculo de cápsula (quando a forma é encapsulada).
type MagistralCheckResponse struct {
	Alerts     []MagistralAlertResponse  `json:"alerts"`
	Capsule    interface{}               `json:"capsule,omitempty"`
	Components []MagistralComponentMatch `json:"components"`
}

// --- Fórmulas-base (templates) e regras de dose ---

// DoseRuleRequest — regra de dose de um componente da fórmula-base.
// Piso e teto são obrigatórios: é a trava que impede dado errado no prontuário de virar dose
// absurda, e ela vale para os três tipos de regra.
type DoseRuleRequest struct {
	Kind models.DoseRuleKind `json:"kind" validate:"required,oneof=fixed per_kg lab_threshold lab_band"`

	PerKg *float64 `json:"perKg,omitempty" validate:"omitempty,gt=0"`

	// O exame é o `code` de lab_test_definitions, escolhido na tela.
	LabCode      *string  `json:"labCode,omitempty" validate:"omitempty,max=64"`
	LabOperator  *string  `json:"labOperator,omitempty" validate:"omitempty,oneof=lt lte gt gte"`
	LabThreshold *float64 `json:"labThreshold,omitempty"`
	DoseIfTrue   *float64 `json:"doseIfTrue,omitempty" validate:"omitempty,gt=0"`
	DoseIfFalse  *float64 `json:"doseIfFalse,omitempty" validate:"omitempty,gt=0"`

	// Conduta por faixa do exame — como a clínica de fato dosa vitamina D, ferritina e B12.
	Bands []DoseBandRequest `json:"bands,omitempty" validate:"omitempty,max=8,dive"`

	// Unidade em que o limiar/faixa foi escrito. Resultado em outra unidade não vira sugestão.
	LabUnit *string `json:"labUnit,omitempty" validate:"omitempty,max=50"`

	FixedDose *float64 `json:"fixedDose,omitempty" validate:"omitempty,gt=0"`

	// Passo prático de arredondamento da dose sugerida (100 UI, 50 mg).
	RoundTo *float64 `json:"roundTo,omitempty" validate:"omitempty,gt=0"`

	MinDose float64 `json:"minDose" validate:"required,gt=0"`
	MaxDose float64 `json:"maxDose" validate:"required,gt=0"`

	MaxDataAgeDays int    `json:"maxDataAgeDays,omitempty" validate:"omitempty,min=1,max=3650"`
	Note           string `json:"note,omitempty" validate:"max=300"`
}

// DoseBandRequest — uma faixa do exame. Intervalo meio-aberto (lowerBound, upperBound], mesma
// convenção das faixas do escore; nulo em lowerBound é -infinito, nulo em upperBound é +infinito.
type DoseBandRequest struct {
	LowerBound *float64 `json:"lowerBound,omitempty"`
	UpperBound *float64 `json:"upperBound,omitempty"`
	Dose       float64  `json:"dose" validate:"required,gt=0"`
	Label      string   `json:"label,omitempty" validate:"max=120"`
}

// FormulaTemplateComponentRequest — substância da fórmula-base, com regra opcional.
type FormulaTemplateComponentRequest struct {
	MagistralComponentID *string                   `json:"magistralComponentId,omitempty" validate:"omitempty,uuid"`
	Substance            string                    `json:"substance" validate:"required,min=2,max=200"`
	Quantity             float64                   `json:"quantity" validate:"required,gt=0"`
	Unit                 string                    `json:"unit" validate:"required,max=20"`
	Category             models.MedicationCategory `json:"category,omitempty" validate:"omitempty,oneof=simple c1 c5 antibiotic glp1"`
	Note                 string                    `json:"note,omitempty" validate:"max=200"`
	// A dose da base é do elemento, não do insumo (fator de correção).
	AsElemental bool             `json:"asElemental,omitempty"`
	Rule        *DoseRuleRequest `json:"rule,omitempty"`
}

// FormulaTemplateRequest — a fórmula-base inteira.
type FormulaTemplateRequest struct {
	Name               string              `json:"name" validate:"required,min=2,max=200"`
	Indication         *string             `json:"indication,omitempty"`
	IndicationBullets  *string             `json:"indicationBullets,omitempty"`
	PharmaceuticalForm string              `json:"pharmaceuticalForm" validate:"required,max=60"`
	UsageType          models.FormulaUsage `json:"usageType,omitempty" validate:"omitempty,oneof=internal external"`
	Route              string              `json:"route,omitempty" validate:"max=40"`
	Vehicle            string              `json:"vehicle,omitempty" validate:"max=200"`
	QuantityToDispense float64             `json:"quantityToDispense,omitempty" validate:"omitempty,gt=0"`
	QuantityUnit       string              `json:"quantityUnit,omitempty" validate:"max=30"`
	Posology           string              `json:"posology,omitempty" validate:"max=300"`
	Duration           int                 `json:"duration,omitempty" validate:"omitempty,min=1,max=365"`
	Instructions       *string             `json:"instructions,omitempty"`
	Notes              *string             `json:"notes,omitempty"`

	Components []FormulaTemplateComponentRequest `json:"components" validate:"required,min=1,max=20,dive"`
}

// DoseSuggestionItem — sugestão para UM componente.
//
// `suggested` ausente com `reason` preenchido é resposta legítima e frequente: sem peso, sem
// exame ou com dado velho, o sistema diz por que não sugeriu em vez de inventar número.
type DoseSuggestionItem struct {
	TemplateComponentID string   `json:"templateComponentId"`
	Substance           string   `json:"substance"`
	Unit                string   `json:"unit"`
	BaseDose            float64  `json:"baseDose"`
	Suggested           *float64 `json:"suggested,omitempty"`
	Basis               string   `json:"basis,omitempty"`
	Clamped             bool     `json:"clamped"`
	Reason              string   `json:"reason,omitempty"`
	// Última dose que o médico ASSINOU desta substância nesta fórmula-base.
	LastPrescribed *float64 `json:"lastPrescribed,omitempty"`
	// A escada de faixas da regra, com a do paciente marcada.
	Bands []DoseBandView `json:"bands,omitempty"`
}

// DoseBandView — uma faixa da regra como a tela mostra.
type DoseBandView struct {
	LowerBound *float64 `json:"lowerBound,omitempty"`
	UpperBound *float64 `json:"upperBound,omitempty"`
	Dose       float64  `json:"dose"`
	Label      string   `json:"label,omitempty"`
	Active     bool     `json:"active"`
}

// DoseSuggestionResponse — sugestões da fórmula-base para um paciente.
type DoseSuggestionResponse struct {
	// Tomadas por dia lidas da posologia da fórmula-base. As regras são escritas em dose diária.
	DosesPerDay  float64              `json:"dosesPerDay,omitempty"`
	TemplateID   string               `json:"templateId"`
	TemplateName string               `json:"templateName"`
	WeightKg     *float64             `json:"weightKg,omitempty"`
	WeightSource string               `json:"weightSource,omitempty"`
	WeightDate   *string              `json:"weightDate,omitempty"`
	Items        []DoseSuggestionItem `json:"items"`
}
