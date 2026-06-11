package services

import (
	"github.com/google/uuid"

	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/pdfdoc"
)

// renderReceiptBytes gera os bytes do recibo pelo pipeline pdfdoc (substitui buildReceiptPDF/gofpdf).
// Usa a identidade FISCAL (razão social + CNPJ + endereço fiscal), não a de atendimento.
func (s *PaymentService) renderReceiptBytes(p *models.AppointmentPayment) ([]byte, error) {
	payer := "Cliente"
	if p.Patient.ID != uuid.Nil && p.Patient.Name != "" {
		payer = p.Patient.Name
	}
	notes := ""
	if p.Notes != nil {
		notes = *p.Notes
	}
	refundNote := ""
	if p.RefundReason != nil {
		refundNote = *p.RefundReason
	}
	fiscal := pdfdoc.Clinic{
		LegalName:   clinicLegalName,
		CNPJ:        clinicCNPJ,
		AddressLine: "Av. Gil de Abreu e Souza, 2335, Casa 634",
		AddressCont: "Bairro Esperança · Londrina/PR · 86058-100",
		Phone:       clinicPhone,
		Email:       clinicEmail,
		Site:        clinicSite,
	}
	return pdfdoc.RenderReceipt(pdfdoc.Receipt{
		Number:      p.ReceiptNumber,
		PayerName:   payer,
		AmountBRL:   formatBRL(p.AmountCents),
		Description: "serviços de saúde",
		Method:      paymentMethodLabel(p.Method),
		PaidAt:      p.PaidAt.In(saoPaulo()).Format("02/01/2006"),
		Notes:       notes,
		Refunded:    p.Status == models.PaymentStatusRefunded,
		RefundNote:  refundNote,
		PlaceDate:   placeDatePT(p.PaidAt),
		Clinic:      fiscal,
	})
}
