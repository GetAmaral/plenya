package dto

import (
	"github.com/plenya/api/internal/models"
)

// CreateLabResultBatchRequest representa a requisição para criar um lote de resultados
type CreateLabResultBatchRequest struct {
	// patientId vem do selectedPatient (context)
	LabRequestID       *string                         `json:"labRequestId,omitempty"`
	RequestingDoctorID *string                         `json:"requestingDoctorId,omitempty"`
	LaboratoryName     string                          `json:"laboratoryName" validate:"required,max=200"`
	CollectionDate     string                          `json:"collectionDate" validate:"required"` // RFC3339
	ResultDate         *string                         `json:"resultDate,omitempty"`               // RFC3339
	Status             models.LabResultBatchStatus     `json:"status" validate:"required"`
	Observations       *string                         `json:"observations,omitempty"`
	Attachments        *string                         `json:"attachments,omitempty"`
	LabResults         []CreateLabResultInBatchRequest `json:"labResults" validate:"required,min=1"`
}

// CreateLabResultInBatchRequest representa um resultado dentro de um batch
type CreateLabResultInBatchRequest struct {
	LabTestDefinitionID *string  `json:"labTestDefinitionId,omitempty" validate:"omitempty,uuid"`
	TestName            string   `json:"testName" validate:"required,max=200"`
	TestType            string   `json:"testType" validate:"required,max=100"`
	ResultText          *string  `json:"resultText,omitempty"`
	ResultNumeric       *float64 `json:"resultNumeric,omitempty"`
	Unit                *string  `json:"unit,omitempty"`
	Interpretation      *string  `json:"interpretation,omitempty"`
	Level               *int     `json:"level,omitempty"`
	Matched             *bool    `json:"matched,omitempty"` // true se matched com definição catalogada
	Source              *string  `json:"source,omitempty"`  // "pdf" | "manual" (default manual)
	MatchReason         *string  `json:"matchReason,omitempty"`
}

// UpdateLabResultBatchRequest representa a requisição para atualizar um lote
type UpdateLabResultBatchRequest struct {
	LabRequestID       *string                         `json:"labRequestId,omitempty"`
	RequestingDoctorID *string                         `json:"requestingDoctorId,omitempty"`
	LaboratoryName     *string                         `json:"laboratoryName,omitempty" validate:"omitempty,max=200"`
	CollectionDate     *string                         `json:"collectionDate,omitempty"` // RFC3339
	ResultDate         *string                         `json:"resultDate,omitempty"`     // RFC3339
	Status             *models.LabResultBatchStatus    `json:"status,omitempty"`
	Observations       *string                         `json:"observations,omitempty"`
	Attachments        *string                         `json:"attachments,omitempty"`
	LabResults         []UpdateLabResultInBatchRequest `json:"labResults,omitempty"` // Nested results sync
}

// UpdateLabResultInBatchRequest - pode ter ID (update) ou não (create)
type UpdateLabResultInBatchRequest struct {
	ID                  *string  `json:"id,omitempty"` // Se presente, faz update; se ausente, cria novo
	LabTestDefinitionID *string  `json:"labTestDefinitionId,omitempty" validate:"omitempty,uuid"`
	TestName            *string  `json:"testName,omitempty" validate:"omitempty,max=200"`
	TestType            *string  `json:"testType,omitempty" validate:"omitempty,max=100"`
	ResultText          *string  `json:"resultText,omitempty"`
	ResultNumeric       *float64 `json:"resultNumeric,omitempty"`
	Unit                *string  `json:"unit,omitempty" validate:"omitempty,max=50"`
	Interpretation      *string  `json:"interpretation,omitempty"`
	Level               *int     `json:"level,omitempty"`
}

// LabResultBatchResponse representa a resposta de um lote (sem resultados)
type LabResultBatchResponse struct {
	ID                 string                      `json:"id"`
	PatientID          string                      `json:"patientId"`
	LabRequestID       *string                     `json:"labRequestId,omitempty"`
	RequestingDoctorID *string                     `json:"requestingDoctorId,omitempty"`
	LaboratoryName     string                      `json:"laboratoryName"`
	CollectionDate     string                      `json:"collectionDate"`
	ResultDate         *string                     `json:"resultDate,omitempty"`
	Status             models.LabResultBatchStatus `json:"status"`
	Observations       *string                     `json:"observations,omitempty"`
	Attachments        *string                     `json:"attachments,omitempty"`
	PDFContentJSON     *string                     `json:"pdfContentJson,omitempty"`
	ResultCount        int                         `json:"resultCount"`
	IsCritical         bool                        `json:"isCritical"`
	WorstLevel         *int                        `json:"worstLevel,omitempty"`
	ReviewedAt         *string                     `json:"reviewedAt,omitempty"`
	CreatedAt          string                      `json:"createdAt"`
	UpdatedAt          string                      `json:"updatedAt"`
}

// LabResultBatchDetailResponse representa a resposta detalhada (com resultados)
type LabResultBatchDetailResponse struct {
	ID                 string                      `json:"id"`
	PatientID          string                      `json:"patientId"`
	LabRequestID       *string                     `json:"labRequestId,omitempty"`
	RequestingDoctorID *string                     `json:"requestingDoctorId,omitempty"`
	LaboratoryName     string                      `json:"laboratoryName"`
	CollectionDate     string                      `json:"collectionDate"`
	ResultDate         *string                     `json:"resultDate,omitempty"`
	Status             models.LabResultBatchStatus `json:"status"`
	Observations       *string                     `json:"observations,omitempty"`
	Attachments        *string                     `json:"attachments,omitempty"`
	PDFContentJSON     *string                     `json:"pdfContentJson,omitempty"`
	ResultCount        int                         `json:"resultCount"`
	IsCritical         bool                        `json:"isCritical"`
	WorstLevel         *int                        `json:"worstLevel,omitempty"`
	ReviewedAt         *string                     `json:"reviewedAt,omitempty"`
	HasPDF             bool                        `json:"hasPdf"` // tem PDF original p/ baixar
	LabResults         []LabResultInBatchResponse  `json:"labResults"`
	CreatedAt          string                      `json:"createdAt"`
	UpdatedAt          string                      `json:"updatedAt"`
}

// LabInboxItemResponse representa um lote na fila "Exames a revisar" (cross-patient).
type LabInboxItemResponse struct {
	ID             string                      `json:"id"`
	PatientID      string                      `json:"patientId"`
	PatientName    string                      `json:"patientName"`
	LaboratoryName string                      `json:"laboratoryName"`
	CollectionDate string                      `json:"collectionDate"`
	ResultDate     *string                     `json:"resultDate,omitempty"`
	Status         models.LabResultBatchStatus `json:"status"`
	IsCritical     bool                        `json:"isCritical"`
	WorstLevel     *int                        `json:"worstLevel,omitempty"`
	AbnormalCount  int                         `json:"abnormalCount"` // results com level <= 2
	TotalResults   int                         `json:"totalResults"`
	ReviewedAt     *string                     `json:"reviewedAt,omitempty"`
}

// LabInboxCountResponse — contadores para o badge do sidebar.
type LabInboxCountResponse struct {
	Total    int64 `json:"total"`
	Critical int64 `json:"critical"`
}

// LabTestDefinitionResponse representa os dados essenciais de uma definição de exame
type LabTestDefinitionResponse struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Code     string  `json:"code"`
	Category string  `json:"category"`
	Unit     *string `json:"unit,omitempty"`
}

// LabResultInBatchResponse representa um resultado individual dentro de um batch
type LabResultInBatchResponse struct {
	ID                    string                     `json:"id"`
	LabResultBatchID      string                     `json:"labResultBatchId"`
	LabTestDefinitionID   *string                    `json:"labTestDefinitionId,omitempty"`
	LabTestDefinition     *LabTestDefinitionResponse `json:"labTestDefinition,omitempty"` // Objeto preloaded
	TestName              string                     `json:"testName"`
	TestType              string                     `json:"testType"`
	ResultText            *string                    `json:"resultText,omitempty"`
	ResultNumeric         *float64                   `json:"resultNumeric,omitempty"`         // Valor CONVERTIDO
	Unit                  *string                    `json:"unit,omitempty"`                  // Unidade CONVERTIDA
	ResultNumericOriginal *float64                   `json:"resultNumericOriginal,omitempty"` // Valor ORIGINAL
	UnitOriginal          *string                    `json:"unitOriginal,omitempty"`          // Unidade ORIGINAL
	Interpretation        *string                    `json:"interpretation,omitempty"`
	Level                 *int                       `json:"level,omitempty"`
	Matched               bool                       `json:"matched"`               // casou com o catálogo
	Source                string                     `json:"source"`                // "pdf" | "manual"
	MatchReason           *string                    `json:"matchReason,omitempty"` // por que não casou
	CreatedAt             string                     `json:"createdAt"`
	UpdatedAt             string                     `json:"updatedAt"`
}
