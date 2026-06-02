package dto

// CreatePatientProblemRequest — Description é texto livre obrigatório; CIDCode opcional.
type CreatePatientProblemRequest struct {
	Description        string  `json:"description" validate:"required,min=2,max=300"`
	CIDCode            *string `json:"cidCode,omitempty"`
	CIDVersion         string  `json:"cidVersion,omitempty"`
	OnsetDate          *string `json:"onsetDate,omitempty"` // YYYY-MM-DD ou RFC3339
	OnsetAppointmentID *string `json:"onsetAppointmentId,omitempty" validate:"omitempty,uuid"`
	Notes              *string `json:"notes,omitempty"`
}

// UpdatePatientProblemRequest — inclui resolver/inativar via status + resolvedDate.
type UpdatePatientProblemRequest struct {
	Description  *string `json:"description,omitempty"`
	CIDCode      *string `json:"cidCode,omitempty"`
	Status       *string `json:"status,omitempty" validate:"omitempty,oneof=active resolved inactive"`
	OnsetDate    *string `json:"onsetDate,omitempty"`
	ResolvedDate *string `json:"resolvedDate,omitempty"`
	Notes        *string `json:"notes,omitempty"`
}

// PatientProblemResponse
type PatientProblemResponse struct {
	ID                 string  `json:"id"`
	PatientID          string  `json:"patientId"`
	Description        string  `json:"description"`
	CIDCode            *string `json:"cidCode,omitempty"`
	CIDVersion         string  `json:"cidVersion"`
	Status             string  `json:"status"`
	OnsetDate          *string `json:"onsetDate,omitempty"`
	ResolvedDate       *string `json:"resolvedDate,omitempty"`
	OnsetAppointmentID *string `json:"onsetAppointmentId,omitempty"`
	Notes              *string `json:"notes,omitempty"`
	RecordedByUserID   string  `json:"recordedByUserId"`
	RecordedByName     string  `json:"recordedByName,omitempty"`
	CreatedAt          string  `json:"createdAt"`
	UpdatedAt          string  `json:"updatedAt"`
}
