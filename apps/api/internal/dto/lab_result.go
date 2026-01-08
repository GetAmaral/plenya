package dto

import "github.com/plenya/api/internal/models"

// CreateLabResultRequest representa o payload de criação de resultado de exame
type CreateLabResultRequest struct {
	PatientID      string                 `json:"patientId" validate:"required,uuid"`
	TestName       string                 `json:"testName" validate:"required"`
	TestType       string                 `json:"testType" validate:"required"`
	OrderedAt      string                 `json:"orderedAt" validate:"required"` // formato: RFC3339
	PerformedAt    *string                `json:"performedAt,omitempty"`         // formato: RFC3339
	Status         models.LabResultStatus `json:"status" validate:"required,oneof=pending in_progress completed cancelled"`
	Result         *string                `json:"result,omitempty"`
	Unit           *string                `json:"unit,omitempty"`
	ReferenceRange *string                `json:"referenceRange,omitempty"`
	Interpretation *string                `json:"interpretation,omitempty"`
}

// UpdateLabResultRequest representa o payload de atualização de resultado de exame
type UpdateLabResultRequest struct {
	PerformedAt    *string                 `json:"performedAt,omitempty"` // formato: RFC3339
	Status         *models.LabResultStatus `json:"status,omitempty" validate:"omitempty,oneof=pending in_progress completed cancelled"`
	Result         *string                 `json:"result,omitempty"`
	Unit           *string                 `json:"unit,omitempty"`
	ReferenceRange *string                 `json:"referenceRange,omitempty"`
	Interpretation *string                 `json:"interpretation,omitempty"`
}

// LabResultResponse representa um resultado de exame na resposta
type LabResultResponse struct {
	ID             string                 `json:"id"`
	PatientID      string                 `json:"patientId"`
	DoctorID       string                 `json:"doctorId"`
	TestName       string                 `json:"testName"`
	TestType       string                 `json:"testType"`
	OrderedAt      string                 `json:"orderedAt"`
	PerformedAt    *string                `json:"performedAt,omitempty"`
	Status         models.LabResultStatus `json:"status"`
	Result         *string                `json:"result,omitempty"`
	Unit           *string                `json:"unit,omitempty"`
	ReferenceRange *string                `json:"referenceRange,omitempty"`
	Interpretation *string                `json:"interpretation,omitempty"`
	CreatedAt      string                 `json:"createdAt"`
	UpdatedAt      string                 `json:"updatedAt"`
}
