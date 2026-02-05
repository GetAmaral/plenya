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
	ReferenceRange      *string  `json:"referenceRange,omitempty"`
	Interpretation      *string  `json:"interpretation,omitempty"`
	Level               *int     `json:"level,omitempty"`
	Matched             *bool    `json:"matched,omitempty"` // true se matched com definição catalogada
}

// UpdateLabResultBatchRequest representa a requisição para atualizar um lote
type UpdateLabResultBatchRequest struct {
	LabRequestID       *string                          `json:"labRequestId,omitempty"`
	RequestingDoctorID *string                          `json:"requestingDoctorId,omitempty"`
	LaboratoryName     *string                          `json:"laboratoryName,omitempty" validate:"omitempty,max=200"`
	CollectionDate     *string                          `json:"collectionDate,omitempty"` // RFC3339
	ResultDate         *string                          `json:"resultDate,omitempty"`     // RFC3339
	Status             *models.LabResultBatchStatus     `json:"status,omitempty"`
	Observations       *string                          `json:"observations,omitempty"`
	Attachments        *string                          `json:"attachments,omitempty"`
	LabResults         []UpdateLabResultInBatchRequest  `json:"labResults,omitempty"` // Nested results sync
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
	ReferenceRange      *string  `json:"referenceRange,omitempty"`
	Interpretation      *string  `json:"interpretation,omitempty"`
	Level               *int     `json:"level,omitempty"`
}

// LabResultBatchResponse representa a resposta de um lote (sem resultados)
type LabResultBatchResponse struct {
	ID                 string                       `json:"id"`
	PatientID          string                       `json:"patientId"`
	LabRequestID       *string                      `json:"labRequestId,omitempty"`
	RequestingDoctorID *string                      `json:"requestingDoctorId,omitempty"`
	LaboratoryName     string                       `json:"laboratoryName"`
	CollectionDate     string                       `json:"collectionDate"`
	ResultDate         *string                      `json:"resultDate,omitempty"`
	Status             models.LabResultBatchStatus  `json:"status"`
	Observations       *string                      `json:"observations,omitempty"`
	Attachments        *string                      `json:"attachments,omitempty"`
	ResultCount        int                          `json:"resultCount"`
	CreatedAt          string                       `json:"createdAt"`
	UpdatedAt          string                       `json:"updatedAt"`
}

// LabResultBatchDetailResponse representa a resposta detalhada (com resultados)
type LabResultBatchDetailResponse struct {
	ID                 string                       `json:"id"`
	PatientID          string                       `json:"patientId"`
	LabRequestID       *string                      `json:"labRequestId,omitempty"`
	RequestingDoctorID *string                      `json:"requestingDoctorId,omitempty"`
	LaboratoryName     string                       `json:"laboratoryName"`
	CollectionDate     string                       `json:"collectionDate"`
	ResultDate         *string                      `json:"resultDate,omitempty"`
	Status             models.LabResultBatchStatus  `json:"status"`
	Observations       *string                      `json:"observations,omitempty"`
	Attachments        *string                      `json:"attachments,omitempty"`
	ResultCount        int                          `json:"resultCount"`
	LabResults         []LabResultInBatchResponse   `json:"labResults"`
	CreatedAt          string                       `json:"createdAt"`
	UpdatedAt          string                       `json:"updatedAt"`
}

// LabResultInBatchResponse representa um resultado individual dentro de um batch
type LabResultInBatchResponse struct {
	ID                  string   `json:"id"`
	LabResultBatchID    string   `json:"labResultBatchId"`
	LabTestDefinitionID *string  `json:"labTestDefinitionId,omitempty"`
	TestName            string   `json:"testName"`
	TestType            string   `json:"testType"`
	ResultText          *string  `json:"resultText,omitempty"`
	ResultNumeric       *float64 `json:"resultNumeric,omitempty"`
	Unit                *string  `json:"unit,omitempty"`
	ReferenceRange      *string  `json:"referenceRange,omitempty"`
	Interpretation      *string  `json:"interpretation,omitempty"`
	Level               *int     `json:"level,omitempty"`
	CreatedAt           string   `json:"createdAt"`
	UpdatedAt           string   `json:"updatedAt"`
}
