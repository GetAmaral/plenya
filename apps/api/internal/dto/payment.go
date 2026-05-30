package dto

import "github.com/plenya/api/internal/models"

// CreatePaymentRequest — registra um pagamento de consulta.
type CreatePaymentRequest struct {
	PatientID     string               `json:"patientId" validate:"required,uuid"`
	AppointmentID *string              `json:"appointmentId,omitempty" validate:"omitempty,uuid"`
	AmountCents   int64                `json:"amountCents" validate:"required,gt=0"`
	Method        models.PaymentMethod `json:"method" validate:"required,oneof=cash pix card transfer other"`
	// PaidAt opcional (RFC3339). Ausente => agora.
	PaidAt *string `json:"paidAt,omitempty"`
	Notes  *string `json:"notes,omitempty"`
}

// RefundPaymentRequest — estorna um pagamento.
type RefundPaymentRequest struct {
	Reason string `json:"reason" validate:"required"`
}

// PaymentResponse — pagamento na resposta.
type PaymentResponse struct {
	ID              string               `json:"id"`
	PatientID       string               `json:"patientId"`
	AppointmentID   *string              `json:"appointmentId,omitempty"`
	AmountCents     int64                `json:"amountCents"`
	Method          models.PaymentMethod `json:"method"`
	Status          models.PaymentStatus `json:"status"`
	PaidAt          string               `json:"paidAt"`
	ReceiptNumber   string               `json:"receiptNumber"`
	Notes           *string              `json:"notes,omitempty"`
	RefundedAt      *string              `json:"refundedAt,omitempty"`
	RefundReason    *string              `json:"refundReason,omitempty"`
	CreatedByUserID string               `json:"createdByUserId"`
	CreatedAt       string               `json:"createdAt"`
	Patient         *AppointmentParty    `json:"patient,omitempty"`
}

// ConsultationPriceRequest — cria/atualiza preço por tipo.
type ConsultationPriceRequest struct {
	Type        models.AppointmentType `json:"type" validate:"required,oneof=initial_assessment follow_up telemedicine procedure results_review"`
	AmountCents int64                  `json:"amountCents" validate:"gte=0"`
	Active      *bool                  `json:"active,omitempty"`
}

// ConsultationPriceResponse — preço na resposta.
type ConsultationPriceResponse struct {
	ID          string                 `json:"id"`
	Type        models.AppointmentType `json:"type"`
	AmountCents int64                  `json:"amountCents"`
	Active      bool                   `json:"active"`
	CreatedAt   string                 `json:"createdAt"`
	UpdatedAt   string                 `json:"updatedAt"`
}
