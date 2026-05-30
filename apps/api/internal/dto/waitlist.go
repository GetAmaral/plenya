package dto

import "github.com/plenya/api/internal/models"

// CreateWaitlistEntryRequest — payload de criação de entrada na lista de espera.
type CreateWaitlistEntryRequest struct {
	PatientID     string                  `json:"patientId" validate:"required,uuid"`
	DoctorID      *string                 `json:"doctorId,omitempty" validate:"omitempty,uuid"`
	PreferredType *models.AppointmentType `json:"preferredType,omitempty" validate:"omitempty,oneof=initial_assessment follow_up telemedicine procedure results_review"`
	Notes         *string                 `json:"notes,omitempty"`
}

// UpdateWaitlistEntryRequest — atualização de status/observações.
type UpdateWaitlistEntryRequest struct {
	Status                 *models.WaitlistStatus `json:"status,omitempty" validate:"omitempty,oneof=waiting scheduled cancelled"`
	Notes                  *string                `json:"notes,omitempty"`
	ScheduledAppointmentID *string                `json:"scheduledAppointmentId,omitempty" validate:"omitempty,uuid"`
}

// WaitlistEntryResponse — entrada na resposta.
type WaitlistEntryResponse struct {
	ID                     string                  `json:"id"`
	PatientID              string                  `json:"patientId"`
	DoctorID               *string                 `json:"doctorId,omitempty"`
	PreferredType          *models.AppointmentType `json:"preferredType,omitempty"`
	Notes                  *string                 `json:"notes,omitempty"`
	Status                 models.WaitlistStatus   `json:"status"`
	ScheduledAppointmentID *string                 `json:"scheduledAppointmentId,omitempty"`
	CreatedByUserID        string                  `json:"createdByUserId"`
	CreatedAt              string                  `json:"createdAt"`
	UpdatedAt              string                  `json:"updatedAt"`
	Patient                *AppointmentParty       `json:"patient,omitempty"`
	Doctor                 *AppointmentParty       `json:"doctor,omitempty"`
}
