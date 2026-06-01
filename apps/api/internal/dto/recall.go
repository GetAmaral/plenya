package dto

import "github.com/plenya/api/internal/models"

// CreateRecallRequest — payload de criação de um retorno previsto.
type CreateRecallRequest struct {
	PatientID           string  `json:"patientId" validate:"required,uuid"`
	DoctorID            *string `json:"doctorId,omitempty" validate:"omitempty,uuid"`
	SourceAppointmentID *string `json:"sourceAppointmentId,omitempty" validate:"omitempty,uuid"`
	DueDate             string  `json:"dueDate" validate:"required"` // YYYY-MM-DD
	Reason              *string `json:"reason,omitempty"`
	Notes               *string `json:"notes,omitempty"`
}

// UpdateRecallRequest — atualização de status/data/observações.
type UpdateRecallRequest struct {
	Status                 *models.RecallStatus `json:"status,omitempty" validate:"omitempty,oneof=pending scheduled dismissed"`
	DueDate                *string              `json:"dueDate,omitempty"` // YYYY-MM-DD
	Reason                 *string              `json:"reason,omitempty"`
	Notes                  *string              `json:"notes,omitempty"`
	ScheduledAppointmentID *string              `json:"scheduledAppointmentId,omitempty" validate:"omitempty,uuid"`
}

// RecallResponse — retorno previsto na resposta.
type RecallResponse struct {
	ID                     string              `json:"id"`
	PatientID              string              `json:"patientId"`
	DoctorID               *string             `json:"doctorId,omitempty"`
	SourceAppointmentID    *string             `json:"sourceAppointmentId,omitempty"`
	DueDate                string              `json:"dueDate"`
	Reason                 *string             `json:"reason,omitempty"`
	Notes                  *string             `json:"notes,omitempty"`
	Status                 models.RecallStatus `json:"status"`
	ScheduledAppointmentID *string             `json:"scheduledAppointmentId,omitempty"`
	CreatedByUserID        string              `json:"createdByUserId"`
	CreatedAt              string              `json:"createdAt"`
	UpdatedAt              string              `json:"updatedAt"`
	Patient                *AppointmentParty   `json:"patient,omitempty"`
	Doctor                 *AppointmentParty   `json:"doctor,omitempty"`
}
