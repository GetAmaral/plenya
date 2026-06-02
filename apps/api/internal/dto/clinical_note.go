package dto

// CreateClinicalNoteRequest — payload de criação de nota de evolução.
// PatientID é opcional (usa selectedPatient se ausente, como na anamnese).
// AppointmentID ancora a nota à visita. Sign=true cria já assinada (atalho do
// "Finalizar consulta"); senão nasce como rascunho (draft).
type CreateClinicalNoteRequest struct {
	AppointmentID  *string `json:"appointmentId,omitempty" validate:"omitempty,uuid"`
	PatientID      string  `json:"patientId,omitempty" validate:"omitempty,uuid"`
	Layout         string  `json:"layout,omitempty" validate:"omitempty,oneof=soap apso"`
	Subjective     *string `json:"subjective,omitempty"`
	SubjectiveHtml *string `json:"subjectiveHtml,omitempty"`
	Objective      *string `json:"objective,omitempty"`
	ObjectiveHtml  *string `json:"objectiveHtml,omitempty"`
	Assessment     *string `json:"assessment,omitempty"`
	AssessmentHtml *string `json:"assessmentHtml,omitempty"`
	Plan           *string `json:"plan,omitempty"`
	PlanHtml       *string `json:"planHtml,omitempty"`
	Visibility     string  `json:"visibility,omitempty" validate:"omitempty,oneof=all medicalOnly psychOnly authorOnly"`
	Sign           bool    `json:"sign,omitempty"`
}

// UpdateClinicalNoteRequest — atualização (só permitida em nota draft).
type UpdateClinicalNoteRequest struct {
	Layout         *string `json:"layout,omitempty" validate:"omitempty,oneof=soap apso"`
	Subjective     *string `json:"subjective,omitempty"`
	SubjectiveHtml *string `json:"subjectiveHtml,omitempty"`
	Objective      *string `json:"objective,omitempty"`
	ObjectiveHtml  *string `json:"objectiveHtml,omitempty"`
	Assessment     *string `json:"assessment,omitempty"`
	AssessmentHtml *string `json:"assessmentHtml,omitempty"`
	Plan           *string `json:"plan,omitempty"`
	PlanHtml       *string `json:"planHtml,omitempty"`
	Visibility     *string `json:"visibility,omitempty" validate:"omitempty,oneof=all medicalOnly psychOnly authorOnly"`
}

// ClinicalNoteResponse — projeção da nota na resposta.
type ClinicalNoteResponse struct {
	ID             string       `json:"id"`
	AppointmentID  *string      `json:"appointmentId,omitempty"`
	PatientID      string       `json:"patientId"`
	AuthorID       string       `json:"authorId"`
	Layout         string       `json:"layout"`
	Subjective     *string      `json:"subjective,omitempty"`
	SubjectiveHtml *string      `json:"subjectiveHtml,omitempty"`
	Objective      *string      `json:"objective,omitempty"`
	ObjectiveHtml  *string      `json:"objectiveHtml,omitempty"`
	Assessment     *string      `json:"assessment,omitempty"`
	AssessmentHtml *string      `json:"assessmentHtml,omitempty"`
	Plan           *string      `json:"plan,omitempty"`
	PlanHtml       *string      `json:"planHtml,omitempty"`
	Status         string       `json:"status"`
	SignedAt       *string      `json:"signedAt,omitempty"`
	AmendmentOfID  *string      `json:"amendmentOfId,omitempty"`
	Visibility     string       `json:"visibility"`
	DisplayTitle   string       `json:"displayTitle"`
	Author         *AuthorBrief `json:"author,omitempty"`
	CreatedAt      string       `json:"createdAt"`
	UpdatedAt      string       `json:"updatedAt"`
}
