package dto

import "github.com/plenya/api/internal/models"

// CreatePrescriptionRequest representa o payload de criação de prescrição
type CreatePrescriptionRequest struct {
	PatientID        string  `json:"patientId" validate:"required,uuid"`
	MedicationName   string  `json:"medicationName" validate:"required"`
	ActiveIngredient *string `json:"activeIngredient,omitempty"`
	Dosage           string  `json:"dosage" validate:"required"`
	Frequency        string  `json:"frequency" validate:"required"`
	Route            string  `json:"route" validate:"required"`
	Duration         *string `json:"duration,omitempty"`
	Quantity         *string `json:"quantity,omitempty"`
	Instructions     *string `json:"instructions,omitempty"`
	PrescriptionDate string  `json:"prescriptionDate" validate:"required"` // formato: RFC3339
	StartDate        *string `json:"startDate,omitempty"`                  // formato: YYYY-MM-DD
	EndDate          *string `json:"endDate,omitempty"`                    // formato: YYYY-MM-DD
	Notes            *string `json:"notes,omitempty"`
}

// UpdatePrescriptionRequest representa o payload de atualização de prescrição
type UpdatePrescriptionRequest struct {
	Status       *models.PrescriptionStatus `json:"status,omitempty" validate:"omitempty,oneof=active completed cancelled expired"`
	Instructions *string                    `json:"instructions,omitempty"`
	Notes        *string                    `json:"notes,omitempty"`
}

// PrescriptionResponse representa uma prescrição na resposta
type PrescriptionResponse struct {
	ID               string                     `json:"id"`
	PatientID        string                     `json:"patientId"`
	DoctorID         string                     `json:"doctorId"`
	MedicationName   string                     `json:"medicationName"`
	ActiveIngredient *string                    `json:"activeIngredient,omitempty"`
	Dosage           string                     `json:"dosage"`
	Frequency        string                     `json:"frequency"`
	Route            string                     `json:"route"`
	Duration         *string                    `json:"duration,omitempty"`
	Quantity         *string                    `json:"quantity,omitempty"`
	Instructions     *string                    `json:"instructions,omitempty"`
	Status           models.PrescriptionStatus  `json:"status"`
	PrescriptionDate string                     `json:"prescriptionDate"`
	StartDate        *string                    `json:"startDate,omitempty"`
	EndDate          *string                    `json:"endDate,omitempty"`
	Notes            *string                    `json:"notes,omitempty"`
	CreatedAt        string                     `json:"createdAt"`
	UpdatedAt        string                     `json:"updatedAt"`
}
