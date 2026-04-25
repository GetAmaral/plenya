package dto

// === Templates ===

type ContinuumTemplateItemPayload struct {
	ID                 *string `json:"id,omitempty"`
	Type               string  `json:"type" validate:"required,oneof=appointment box reassessment milestone custom"`
	Specialty          *string `json:"specialty,omitempty" validate:"omitempty,oneof=doctor nutritionist psychologist physicalEducator"`
	Title              string  `json:"title" validate:"required,max=160"`
	Description        string  `json:"description"`
	WeekOffset         int     `json:"weekOffset" validate:"min=0"`
	ExpectedOffsetDays int     `json:"expectedOffsetDays" validate:"min=0,max=6"`
	LateAfterDays      int     `json:"lateAfterDays" validate:"min=0"`
	BoxTemplateID      *string `json:"boxTemplateId,omitempty"`
	Position           int     `json:"position"`
}

type ContinuumTemplatePayload struct {
	Name          string                         `json:"name" validate:"required,max=120"`
	Description   string                         `json:"description"`
	DurationWeeks int                            `json:"durationWeeks" validate:"required,min=1,max=520"`
	Status        string                         `json:"status,omitempty" validate:"omitempty,oneof=active archived"`
	Items         []ContinuumTemplateItemPayload `json:"items"`
}

type CloneContinuumTemplatePayload struct {
	Name string `json:"name" validate:"required,max=120"`
}

// === Box Templates ===

type ContinuumBoxTemplatePayload struct {
	Name        string `json:"name" validate:"required,max=120"`
	Description string `json:"description"`
	Contents    string `json:"contents"`
	Notes       string `json:"notes"`
	Status      string `json:"status,omitempty" validate:"omitempty,oneof=active archived"`
}

type CloneContinuumBoxTemplatePayload struct {
	Name string `json:"name" validate:"required,max=120"`
}

// === Inscrição de paciente ===

type EnrollPatientContinuumPayload struct {
	TemplateID          string  `json:"templateId" validate:"required,uuid"`
	StartDate           string  `json:"startDate" validate:"required"` // ISO YYYY-MM-DD
	CoordinatorDoctorID *string `json:"coordinatorDoctorId,omitempty" validate:"omitempty,uuid"`
	Notes               string  `json:"notes,omitempty"`
}

type UpdateContinuumItemPayload struct {
	Status           *string `json:"status,omitempty" validate:"omitempty,oneof=pending scheduled completed missed cancelled skipped"`
	AppointmentID    *string `json:"appointmentId,omitempty" validate:"omitempty,uuid"`
	CompletedAt      *string `json:"completedAt,omitempty"` // ISO RFC3339
	CompletedRefType *string `json:"completedRefType,omitempty"`
	CompletedRefID   *string `json:"completedRefId,omitempty" validate:"omitempty,uuid"`
}

// === Plano integrado ===

type UpdateIntegratedPlanPayload struct {
	Content string `json:"content"`
}

// === Box logístico ===

type UpdateBoxPayload struct {
	Status       *string `json:"status,omitempty" validate:"omitempty,oneof=planned preparing shipped delivered cancelled"`
	TrackingCode *string `json:"trackingCode,omitempty"`
	Carrier      *string `json:"carrier,omitempty"`
	Notes        *string `json:"notes,omitempty"`
	Contents     *string `json:"contents,omitempty"`
	Address      *string `json:"address,omitempty"`
}
