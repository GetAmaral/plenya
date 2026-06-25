package dto

// CreateIssuedDocumentRequest — emissão de documento clínico (atestado/declaração/laudo).
// patientId vem do path (/patients/:id/issued-documents).
type CreateIssuedDocumentRequest struct {
	AppointmentID *string `json:"appointmentId,omitempty" validate:"omitempty,uuid"`
	Type          string  `json:"type" validate:"required,oneof=certificate declaration report orientation"`
	Title         string  `json:"title" validate:"required,max=200"`
	Body          string  `json:"body" validate:"max=10000"`
	BodyHtml      *string `json:"bodyHtml,omitempty" validate:"omitempty,max=50000"`
	Purpose       *string `json:"purpose,omitempty" validate:"omitempty,max=300"`
	DaysOff       *int    `json:"daysOff,omitempty" validate:"omitempty,min=0,max=365"`
	IncludesCid   bool    `json:"includesCid"`
	CidCode       *string `json:"cidCode,omitempty" validate:"omitempty,max=10"`
	CidConsent    bool    `json:"cidConsent"`
}

// IssuedDocumentResponse — documento emitido (metadados + estado de assinatura).
type IssuedDocumentResponse struct {
	ID                  string  `json:"id"`
	PatientID           string  `json:"patientId"`
	AppointmentID       *string `json:"appointmentId,omitempty"`
	DoctorID            string  `json:"doctorId"`
	DoctorName          string  `json:"doctorName,omitempty"`
	Type                string  `json:"type"`
	Title               string  `json:"title"`
	Body                string  `json:"body"`
	BodyHtml            *string `json:"bodyHtml,omitempty"`
	Purpose             *string `json:"purpose,omitempty"`
	DaysOff             *int    `json:"daysOff,omitempty"`
	IncludesCid         bool    `json:"includesCid"`
	CidCode             *string `json:"cidCode,omitempty"`
	CidConsent          bool    `json:"cidConsent"`
	Status              string  `json:"status"`
	HasDigitalSignature bool    `json:"hasDigitalSignature"`
	SignedAt            *string `json:"signedAt,omitempty"`
	SignedPdfHash       *string `json:"signedPdfHash,omitempty"`
	CertificateSerial   *string `json:"certificateSerial,omitempty"`
	QrCodeData          *string `json:"qrCodeData,omitempty"`
	PatientDocumentID   *string `json:"patientDocumentId,omitempty"`
	IssuedByUserID      string  `json:"issuedByUserId"`
	IssuedAt            string  `json:"issuedAt"`
	CreatedAt           string  `json:"createdAt"`
}
