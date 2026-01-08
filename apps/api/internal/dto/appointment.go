package dto

import "github.com/plenya/api/internal/models"

// CreateAppointmentRequest representa o payload de criação de consulta
type CreateAppointmentRequest struct {
	PatientID       string                 `json:"patientId" validate:"required,uuid"`
	DoctorID        string                 `json:"doctorId" validate:"required,uuid"`
	ScheduledAt     string                 `json:"scheduledAt" validate:"required"` // formato: RFC3339
	DurationMinutes int                    `json:"durationMinutes" validate:"required,min=15,max=480"`
	Type            models.AppointmentType `json:"type" validate:"required,oneof=routine follow_up urgent emergency"`
	Reason          string                 `json:"reason" validate:"required"`
	PatientNotes    *string                `json:"patientNotes,omitempty"`
}

// UpdateAppointmentRequest representa o payload de atualização de consulta
type UpdateAppointmentRequest struct {
	ScheduledAt *string                  `json:"scheduledAt,omitempty"` // formato: RFC3339
	Status      *models.AppointmentStatus `json:"status,omitempty" validate:"omitempty,oneof=scheduled confirmed completed cancelled no_show"`
	DoctorNotes *string                  `json:"doctorNotes,omitempty"`
	Diagnosis   *string                  `json:"diagnosis,omitempty"`
}

// CancelAppointmentRequest representa o payload de cancelamento
type CancelAppointmentRequest struct {
	Reason string `json:"reason" validate:"required"`
}

// AppointmentResponse representa uma consulta na resposta
type AppointmentResponse struct {
	ID                 string                    `json:"id"`
	PatientID          string                    `json:"patientId"`
	DoctorID           string                    `json:"doctorId"`
	ScheduledAt        string                    `json:"scheduledAt"`
	DurationMinutes    int                       `json:"durationMinutes"`
	Type               models.AppointmentType    `json:"type"`
	Status             models.AppointmentStatus  `json:"status"`
	Reason             string                    `json:"reason"`
	PatientNotes       *string                   `json:"patientNotes,omitempty"`
	DoctorNotes        *string                   `json:"doctorNotes,omitempty"`
	Diagnosis          *string                   `json:"diagnosis,omitempty"`
	AnamnesisID        *string                   `json:"anamnesisId,omitempty"`
	ConfirmedAt        *string                   `json:"confirmedAt,omitempty"`
	CompletedAt        *string                   `json:"completedAt,omitempty"`
	CancelledAt        *string                   `json:"cancelledAt,omitempty"`
	CancellationReason *string                   `json:"cancellationReason,omitempty"`
	CreatedAt          string                    `json:"createdAt"`
	UpdatedAt          string                    `json:"updatedAt"`
}
