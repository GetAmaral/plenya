package dto

import "github.com/plenya/api/internal/models"

// MedicationRequest representa um medicamento na prescrição
type MedicationRequest struct {
	MedicationDefinitionID *string `json:"medicationDefinitionId,omitempty" validate:"omitempty,uuid"`
	MedicationName         string  `json:"medicationName" validate:"required,min=3,max=200"`
	// Princípio ativo e concentração são OPCIONAIS: a tela deixou de exigi-los (o service cai
	// para o nome comercial quando vêm vazios) e a validação aqui continuava recusando, com 400
	// e sem erro de campo na tela — o médico só via "validation failed".
	ActiveIngredient string `json:"activeIngredient" validate:"omitempty,max=200"`
	// a_b (tarja preta) entra na lista: a migration 00080 permite a categoria no catálogo e a
	// tela preenche a partir dele; sem isto, escolher a substância dava 400 no salvar.
	Category      models.MedicationCategory `json:"category" validate:"required,oneof=simple c1 c5 antibiotic glp1 a_b"`
	Concentration string                    `json:"concentration" validate:"omitempty,max=100"`
	Dosage        string                    `json:"dosage" validate:"required"`
	Frequency     string                    `json:"frequency" validate:"required"`
	Route         string                    `json:"route" validate:"required"`
	Duration      int                       `json:"duration" validate:"required,min=1,max=365"`
	Quantity      int                       `json:"quantity" validate:"required,min=1"`
	// Extenso só é exigido em controlado, e é o service que confere isso. Aqui era "required" e
	// derrubava toda receita cuja posologia não dá para calcular quantidade ("se necessário").
	QuantityInWords string  `json:"quantityInWords" validate:"omitempty,max=200"`
	Instructions    *string `json:"instructions,omitempty"`
}

// FormulaComponentRequest — uma substância dentro de uma fórmula magistral.
type FormulaComponentRequest struct {
	MagistralComponentID   *string `json:"magistralComponentId,omitempty" validate:"omitempty,uuid"`
	MedicationDefinitionID *string `json:"medicationDefinitionId,omitempty" validate:"omitempty,uuid"`
	Substance              string  `json:"substance" validate:"required,min=2,max=200"`
	Quantity               float64 `json:"quantity" validate:"required,gt=0"`
	Unit                   string  `json:"unit" validate:"required,max=20"`
	// a_b (tarja preta) entra na lista: a migration 00080 permite a categoria no catálogo e a
	// tela preenche a partir dele; sem isto, escolher a substância dava 400 no salvar.
	Category models.MedicationCategory `json:"category" validate:"required,oneof=simple c1 c5 antibiotic glp1 a_b"`
	Note     string                    `json:"note,omitempty" validate:"max=200"`
	// O que o motor de regras sugeriu, quando houve sugestão. Gravado AO LADO da dose prescrita:
	// é o que permite saber depois quando o médico discordou da sugestão.
	SuggestedQuantity *float64 `json:"suggestedQuantity,omitempty"`
	// A dose escrita é do elemento (ou ativo puro), não do insumo.
	AsElemental bool `json:"asElemental,omitempty"`
}

// FormulaRequest — uma fórmula magistral da receita de manipulado.
type FormulaRequest struct {
	TemplateID         *string             `json:"templateId,omitempty" validate:"omitempty,uuid"`
	Name               string              `json:"name,omitempty" validate:"max=200"`
	PharmaceuticalForm string              `json:"pharmaceuticalForm" validate:"required,max=60"`
	UsageType          models.FormulaUsage `json:"usageType" validate:"required,oneof=internal external"`
	Route              string              `json:"route,omitempty" validate:"max=40"`
	Vehicle            string              `json:"vehicle,omitempty" validate:"max=200"`
	QuantityToDispense float64             `json:"quantityToDispense" validate:"required,gt=0"`
	QuantityUnit       string              `json:"quantityUnit" validate:"required,max=30"`
	QuantityInWords    string              `json:"quantityInWords,omitempty" validate:"max=200"`
	Posology           string              `json:"posology" validate:"required,max=300"`
	Duration           int                 `json:"duration,omitempty" validate:"omitempty,min=1,max=365"`
	Instructions       *string             `json:"instructions,omitempty"`
	// Teto de 20 componentes por fórmula: acima disso o bloco não cabe numa página do
	// receituário e o paginador do PDF não quebra bloco (transbordaria em silêncio).
	Components []FormulaComponentRequest `json:"components" validate:"required,min=1,max=20,dive"`
}

// CreatePrescriptionRequest representa o payload de criação de prescrição digital
type CreatePrescriptionRequest struct {
	PatientID string `json:"patientId" validate:"required,uuid"`
	// Tipo da receita. Ausente = commercial (todo cliente antigo continua funcionando).
	Type models.PrescriptionType `json:"type,omitempty" validate:"omitempty,oneof=commercial compounded"`
	// Medicamentos industrializados (receita commercial).
	Medications []MedicationRequest `json:"medications,omitempty" validate:"omitempty,max=10,dive"`
	// Fórmulas magistrais (receita compounded).
	Formulas            []FormulaRequest `json:"formulas,omitempty" validate:"omitempty,max=10,dive"`
	GeneralInstructions *string          `json:"generalInstructions,omitempty"`
	PrescriptionDate    string           `json:"prescriptionDate" validate:"required"` // formato: RFC3339
}

// UpdatePrescriptionRequest representa o payload de atualização de prescrição
type UpdatePrescriptionRequest struct {
	Medications         *[]MedicationRequest       `json:"medications,omitempty" validate:"omitempty,min=1,max=10,dive"`
	Formulas            *[]FormulaRequest          `json:"formulas,omitempty" validate:"omitempty,min=1,max=10,dive"`
	GeneralInstructions *string                    `json:"generalInstructions,omitempty"`
	Status              *models.PrescriptionStatus `json:"status,omitempty" validate:"omitempty,oneof=active completed cancelled expired"`
}

// MedicationResponse representa um medicamento na resposta
type MedicationResponse struct {
	ID                     string                    `json:"id"`
	MedicationDefinitionID *string                   `json:"medicationDefinitionId,omitempty"`
	MedicationName         string                    `json:"medicationName"`
	ActiveIngredient       string                    `json:"activeIngredient"`
	Category               models.MedicationCategory `json:"category"`
	Concentration          string                    `json:"concentration"`
	Dosage                 string                    `json:"dosage"`
	Frequency              string                    `json:"frequency"`
	Route                  string                    `json:"route"`
	Duration               int                       `json:"duration"`
	Quantity               int                       `json:"quantity"`
	QuantityInWords        string                    `json:"quantityInWords"`
	Instructions           *string                   `json:"instructions,omitempty"`
}

// FormulaComponentResponse — componente de fórmula na resposta.
type FormulaComponentResponse struct {
	ID                     string                    `json:"id"`
	MagistralComponentID   *string                   `json:"magistralComponentId,omitempty"`
	MedicationDefinitionID *string                   `json:"medicationDefinitionId,omitempty"`
	Substance              string                    `json:"substance"`
	Quantity               float64                   `json:"quantity"`
	Unit                   string                    `json:"unit"`
	Category               models.MedicationCategory `json:"category"`
	Note                   string                    `json:"note,omitempty"`
	SuggestedQuantity      *float64                  `json:"suggestedQuantity,omitempty"`
	AsElemental            bool                      `json:"asElemental,omitempty"`
}

// FormulaResponse — fórmula magistral na resposta.
type FormulaResponse struct {
	ID                 string                     `json:"id"`
	TemplateID         *string                    `json:"templateId,omitempty"`
	Name               string                     `json:"name,omitempty"`
	PharmaceuticalForm string                     `json:"pharmaceuticalForm"`
	UsageType          models.FormulaUsage        `json:"usageType"`
	Route              string                     `json:"route,omitempty"`
	Vehicle            string                     `json:"vehicle,omitempty"`
	QuantityToDispense float64                    `json:"quantityToDispense"`
	QuantityUnit       string                     `json:"quantityUnit"`
	QuantityInWords    string                     `json:"quantityInWords,omitempty"`
	Posology           string                     `json:"posology"`
	Duration           int                        `json:"duration,omitempty"`
	Instructions       *string                    `json:"instructions,omitempty"`
	HighestCategory    models.MedicationCategory  `json:"highestCategory"`
	Components         []FormulaComponentResponse `json:"components"`
}

// PrescriptionResponse representa uma prescrição digital na resposta
type PrescriptionResponse struct {
	ID                  string                    `json:"id"`
	PatientID           string                    `json:"patientId"`
	DoctorID            string                    `json:"doctorId"`
	Type                models.PrescriptionType   `json:"type"`
	Medications         []MedicationResponse      `json:"medications"`
	Formulas            []FormulaResponse         `json:"formulas"`
	GeneralInstructions *string                   `json:"generalInstructions,omitempty"`
	Status              models.PrescriptionStatus `json:"status"`
	PrescriptionDate    string                    `json:"prescriptionDate"`
	ValidUntil          string                    `json:"validUntil"`
	SNCRNumber          *string                   `json:"sncrNumber,omitempty"`
	SNCRStatus          *string                   `json:"sncrStatus,omitempty"`
	SignedPDFPath       *string                   `json:"signedPdfPath,omitempty"`
	SignedPDFHash       *string                   `json:"signedPdfHash,omitempty"`
	QRCodeData          *string                   `json:"qrCodeData,omitempty"`
	SignedAt            *string                   `json:"signedAt,omitempty"`
	CertificateSerial   *string                   `json:"certificateSerial,omitempty"`
	// SignatureMode: "digital" (ICP-Brasil) ou "manual" (impressão + assinatura/carimbo à mão).
	SignatureMode *string `json:"signatureMode,omitempty"`
	IsUsed        bool    `json:"isUsed"`
	DispensedAt   *string `json:"dispensedAt,omitempty"`
	CreatedAt     string  `json:"createdAt"`
	UpdatedAt     string  `json:"updatedAt"`
}
