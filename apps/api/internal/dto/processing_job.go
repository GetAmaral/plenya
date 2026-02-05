package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/plenya/api/internal/models"
)

// ProcessingJobResponse - status de job para frontend
// @Description Resposta com status do processamento assíncrono
type ProcessingJobResponse struct {
	ID               uuid.UUID                    `json:"id" example:"019c1a1e-0579-7f3b-a1bd-4767008e844c"`
	LabResultBatchID uuid.UUID                    `json:"labResultBatchId" example:"019c1a1e-0579-7f3b-a1bd-4767008e844c"`
	Type             models.ProcessingJobType     `json:"type" example:"pdf_extraction"`
	Status           models.ProcessingJobStatus   `json:"status" example:"processing" enums:"pending,processing,completed,failed"`
	ErrorMessage     *string                      `json:"errorMessage,omitempty" example:"OCR extraction failed"`
	Attempts         int                          `json:"attempts" example:"1"`
	CreatedAt        time.Time                    `json:"createdAt" example:"2024-01-20T15:04:05Z"`
	StartedAt        *time.Time                   `json:"startedAt,omitempty" example:"2024-01-20T15:04:10Z"`
	CompletedAt      *time.Time                   `json:"completedAt,omitempty" example:"2024-01-20T15:05:00Z"`
}

// AILabResultExtractionResponse - resposta estruturada da IA
// @Description Dados extraídos do PDF via IA (Claude)
type AILabResultExtractionResponse struct {
	LaboratoryName string                  `json:"laboratoryName" example:"Laboratório Fleury"`
	CollectionDate string                  `json:"collectionDate" example:"2024-01-20"` // ISO8601 date
	ResultDate     string                  `json:"resultDate" example:"2024-01-21"`      // ISO8601 date
	LabResults     []AIExtractedLabResult  `json:"labResults"`
}

// AIExtractedLabResult - resultado individual extraído pela IA
// @Description Resultado de exame laboratorial interpretado pela IA
type AIExtractedLabResult struct {
	LabTestDefinitionID *uuid.UUID `json:"labTestDefinitionId,omitempty" example:"019c1a1e-0579-7f3b-a1bd-4767008e844c"` // null se não matched
	TestName            string     `json:"testName" example:"Glicemia de Jejum"`
	TestType            string     `json:"testType" example:"biochemistry"`
	ResultNumeric       *float64   `json:"resultNumeric,omitempty" example:"95.5"`
	ResultText          *string    `json:"resultText,omitempty" example:"Normal"`
	Unit                *string    `json:"unit,omitempty" example:"mg/dL"`
	ReferenceRange      *string    `json:"referenceRange,omitempty" example:"70-100 mg/dL"`
	Interpretation      *string    `json:"interpretation,omitempty" example:"Dentro dos valores normais"`
	Matched             bool       `json:"matched" example:"true"` // true se encontrou match com definição catalogada
}

// LabTestSummary - resumo de definição de exame para enviar no prompt da IA
// @Description Resumo compacto de LabTestDefinition para matching pela IA
type LabTestSummary struct {
	ID        uuid.UUID `json:"id" example:"019c1a1e-0579-7f3b-a1bd-4767008e844c"`
	Code      string    `json:"code" example:"GLICEMIA"`
	Name      string    `json:"name" example:"Glicemia de Jejum"`
	ShortName *string   `json:"shortName,omitempty" example:"Glicemia"`
	TussCode  *string   `json:"tussCode,omitempty" example:"40301010"`
	LoincCode *string   `json:"loincCode,omitempty" example:"1558-6"`
	Category  string    `json:"category" example:"biochemistry"`
	Unit      *string   `json:"unit,omitempty" example:"mg/dL"`
}
