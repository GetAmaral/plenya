package dto

// CreateMedicationInUseRequest
type CreateMedicationInUseRequest struct {
	MedicationName       string  `json:"medicationName" validate:"required,min=2,max=200"`
	ActiveIngredient     string  `json:"activeIngredient,omitempty"`
	Dosage               string  `json:"dosage,omitempty"`
	Frequency            string  `json:"frequency,omitempty"`
	Route                string  `json:"route,omitempty"`
	Source               string  `json:"source,omitempty" validate:"omitempty,oneof=prescribed_here external patient_reported"`
	SourcePrescriptionID *string `json:"sourcePrescriptionId,omitempty" validate:"omitempty,uuid"`
	StartDate            *string `json:"startDate,omitempty"` // YYYY-MM-DD ou RFC3339
	Notes                *string `json:"notes,omitempty"`
}

// UpdateMedicationInUseRequest — inclui suspender/parar via status + endDate.
type UpdateMedicationInUseRequest struct {
	MedicationName   *string `json:"medicationName,omitempty"`
	ActiveIngredient *string `json:"activeIngredient,omitempty"`
	Dosage           *string `json:"dosage,omitempty"`
	Frequency        *string `json:"frequency,omitempty"`
	Route            *string `json:"route,omitempty"`
	Status           *string `json:"status,omitempty" validate:"omitempty,oneof=active suspended stopped"`
	EndDate          *string `json:"endDate,omitempty"`
	Notes            *string `json:"notes,omitempty"`
}

// MedicationInUseResponse
type MedicationInUseResponse struct {
	ID                      string  `json:"id"`
	PatientID               string  `json:"patientId"`
	MedicationName          string  `json:"medicationName"`
	ActiveIngredient        string  `json:"activeIngredient"`
	Dosage                  string  `json:"dosage"`
	Frequency               string  `json:"frequency"`
	Route                   string  `json:"route"`
	Source                  string  `json:"source"`
	Status                  string  `json:"status"`
	SourcePrescriptionID    *string `json:"sourcePrescriptionId,omitempty"`
	ReconciledAppointmentID *string `json:"reconciledAppointmentId,omitempty"`
	StartDate               *string `json:"startDate,omitempty"`
	EndDate                 *string `json:"endDate,omitempty"`
	Notes                   *string `json:"notes,omitempty"`
	RecordedByUserID        string  `json:"recordedByUserId"`
	RecordedByName          string  `json:"recordedByName,omitempty"`
	CreatedAt               string  `json:"createdAt"`
	UpdatedAt               string  `json:"updatedAt"`
}
