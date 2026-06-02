package dto

// CreatePatientAllergyRequest — payload de criação de alergia.
// Se NoKnownAllergies=true, Substance pode vir vazio (asserção NKA).
type CreatePatientAllergyRequest struct {
	Substance        string  `json:"substance,omitempty"`
	SubstanceType    string  `json:"substanceType,omitempty" validate:"omitempty,oneof=drug food environmental other"`
	Reaction         *string `json:"reaction,omitempty"`
	Severity         string  `json:"severity,omitempty" validate:"omitempty,oneof=mild moderate severe anaphylaxis"`
	NoKnownAllergies bool    `json:"noKnownAllergies,omitempty"`
	Notes            *string `json:"notes,omitempty"`
}

// UpdatePatientAllergyRequest — atualização parcial (inclui inativar via status).
type UpdatePatientAllergyRequest struct {
	Substance     *string `json:"substance,omitempty"`
	SubstanceType *string `json:"substanceType,omitempty" validate:"omitempty,oneof=drug food environmental other"`
	Reaction      *string `json:"reaction,omitempty"`
	Severity      *string `json:"severity,omitempty" validate:"omitempty,oneof=mild moderate severe anaphylaxis"`
	Status        *string `json:"status,omitempty" validate:"omitempty,oneof=active inactive"`
	Notes         *string `json:"notes,omitempty"`
}

// PatientAllergyResponse
type PatientAllergyResponse struct {
	ID               string  `json:"id"`
	PatientID        string  `json:"patientId"`
	Substance        string  `json:"substance"`
	SubstanceType    string  `json:"substanceType"`
	Reaction         *string `json:"reaction,omitempty"`
	Severity         string  `json:"severity"`
	Status           string  `json:"status"`
	NoKnownAllergies bool    `json:"noKnownAllergies"`
	Notes            *string `json:"notes,omitempty"`
	RecordedByUserID string  `json:"recordedByUserId"`
	RecordedByName   string  `json:"recordedByName,omitempty"`
	CreatedAt        string  `json:"createdAt"`
	UpdatedAt        string  `json:"updatedAt"`
}
