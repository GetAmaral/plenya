package dto

// PrepResponseInput é uma resposta do paciente a um item do formulário de preparação.
type PrepResponseInput struct {
	ScoreItemID   string   `json:"scoreItemId" validate:"required,uuid"`
	NumericValue  *float64 `json:"numericValue,omitempty"`
	SelectedLevel *int     `json:"selectedLevel,omitempty" validate:"omitempty,gte=0,lte=6"`
	TextValue     *string  `json:"textValue,omitempty"`
}

// SubmitPrepRequest é o payload de salvar/enviar a preparação pré-consulta.
// Submit=false salva rascunho; Submit=true finaliza (status submitted).
type SubmitPrepRequest struct {
	AppointmentID  *string             `json:"appointmentId,omitempty" validate:"omitempty,uuid"`
	ChiefComplaint *string             `json:"chiefComplaint,omitempty"`
	Responses      []PrepResponseInput `json:"responses" validate:"omitempty,dive"`
	ConsentVersion *string             `json:"consentVersion,omitempty" validate:"omitempty,max=20"`
	Submit         bool                `json:"submit,omitempty"`
}

// PrepResponseView é uma resposta no formato de leitura.
type PrepResponseView struct {
	ScoreItemID   string   `json:"scoreItemId"`
	NumericValue  *float64 `json:"numericValue,omitempty"`
	SelectedLevel *int     `json:"selectedLevel,omitempty"`
	TextValue     *string  `json:"textValue,omitempty"`
}

// ConsultationPrepView é a preparação no formato de leitura (portal do paciente).
type ConsultationPrepView struct {
	ID             string             `json:"id"`
	PatientID      string             `json:"patientId"`
	AppointmentID  *string            `json:"appointmentId,omitempty"`
	Status         string             `json:"status"`
	ChiefComplaint *string            `json:"chiefComplaint,omitempty"`
	SubmittedAt    *string            `json:"submittedAt,omitempty"`
	Responses      []PrepResponseView `json:"responses"`
	CreatedAt      string             `json:"createdAt"`
	UpdatedAt      string             `json:"updatedAt"`
}
