package dto

// CreateClinicalNoteRequest — payload de criação de nota de evolução.
// PatientID é opcional (usa selectedPatient se ausente, como na anamnese).
// AppointmentID ancora a nota à visita. Sign=true cria já assinada (atalho do
// "Finalizar consulta"); senão nasce como rascunho (draft).
type CreateClinicalNoteRequest struct {
	AppointmentID       *string `json:"appointmentId,omitempty" validate:"omitempty,uuid"`
	PatientID           string  `json:"patientId,omitempty" validate:"omitempty,uuid"`
	ClinicalHistory     *string `json:"clinicalHistory,omitempty"`
	ClinicalHistoryHtml *string `json:"clinicalHistoryHtml,omitempty"`
	Conduct             *string `json:"conduct,omitempty"`
	ConductHtml         *string `json:"conductHtml,omitempty"`
	Visibility          string  `json:"visibility,omitempty" validate:"omitempty,oneof=all medicalOnly psychOnly authorOnly"`
	Sign                bool    `json:"sign,omitempty"`
}

// UpdateClinicalNoteRequest — atualização (só permitida em nota draft).
type UpdateClinicalNoteRequest struct {
	ClinicalHistory     *string `json:"clinicalHistory,omitempty"`
	ClinicalHistoryHtml *string `json:"clinicalHistoryHtml,omitempty"`
	Conduct             *string `json:"conduct,omitempty"`
	ConductHtml         *string `json:"conductHtml,omitempty"`
	Visibility          *string `json:"visibility,omitempty" validate:"omitempty,oneof=all medicalOnly psychOnly authorOnly"`
}

// ClinicalNoteResponse — projeção da nota na resposta.
type ClinicalNoteResponse struct {
	ID                  string       `json:"id"`
	AppointmentID       *string      `json:"appointmentId,omitempty"`
	PatientID           string       `json:"patientId"`
	AuthorID            string       `json:"authorId"`
	ClinicalHistory     *string      `json:"clinicalHistory,omitempty"`
	ClinicalHistoryHtml *string      `json:"clinicalHistoryHtml,omitempty"`
	Conduct             *string      `json:"conduct,omitempty"`
	ConductHtml         *string      `json:"conductHtml,omitempty"`
	Status              string       `json:"status"`
	SignedAt       *string      `json:"signedAt,omitempty"`
	AmendmentOfID  *string      `json:"amendmentOfId,omitempty"`
	Visibility     string       `json:"visibility"`
	DisplayTitle   string       `json:"displayTitle"`
	Author         *AuthorBrief `json:"author,omitempty"`
	CreatedAt      string       `json:"createdAt"`
	UpdatedAt      string       `json:"updatedAt"`
}
