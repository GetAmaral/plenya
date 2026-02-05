package dto

import "github.com/plenya/api/internal/models"

// MedicationRequest representa um medicamento na prescrição
type MedicationRequest struct {
	MedicationDefinitionID *string                   `json:"medicationDefinitionId,omitempty" validate:"omitempty,uuid"`
	MedicationName         string                    `json:"medicationName" validate:"required,min=3,max=200"`
	ActiveIngredient       string                    `json:"activeIngredient" validate:"required,min=3,max=200"`
	Category               models.MedicationCategory `json:"category" validate:"required,oneof=simple c1 c5 antibiotic glp1"`
	Concentration          string                    `json:"concentration" validate:"required"`
	Dosage                 string                    `json:"dosage" validate:"required"`
	Frequency              string                    `json:"frequency" validate:"required"`
	Route                  string                    `json:"route" validate:"required"`
	Duration               int                       `json:"duration" validate:"required,min=1,max=365"`
	Quantity               int                       `json:"quantity" validate:"required,min=1"`
	QuantityInWords        string                    `json:"quantityInWords" validate:"required,min=3,max=200"`
	Instructions           *string                   `json:"instructions,omitempty"`
}

// CreatePrescriptionRequest representa o payload de criação de prescrição digital
type CreatePrescriptionRequest struct {
	PatientID            string              `json:"patientId" validate:"required,uuid"`
	Medications          []MedicationRequest `json:"medications" validate:"required,min=1,max=10,dive"`
	GeneralInstructions  *string             `json:"generalInstructions,omitempty"`
	PrescriptionDate     string              `json:"prescriptionDate" validate:"required"` // formato: RFC3339
}

// UpdatePrescriptionRequest representa o payload de atualização de prescrição
type UpdatePrescriptionRequest struct {
	Medications         *[]MedicationRequest       `json:"medications,omitempty" validate:"omitempty,min=1,max=10,dive"`
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

// PrescriptionResponse representa uma prescrição digital na resposta
type PrescriptionResponse struct {
	ID                  string                    `json:"id"`
	PatientID           string                    `json:"patientId"`
	DoctorID            string                    `json:"doctorId"`
	Medications         []MedicationResponse      `json:"medications"`
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
	IsUsed              bool                      `json:"isUsed"`
	DispensedAt         *string                   `json:"dispensedAt,omitempty"`
	CreatedAt           string                    `json:"createdAt"`
	UpdatedAt           string                    `json:"updatedAt"`
}
