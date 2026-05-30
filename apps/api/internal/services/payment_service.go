package services

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	gofpdf "codeberg.org/go-pdf/fpdf"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

var (
	ErrPaymentNotFound = errors.New("payment not found")
	ErrPriceNotFound   = errors.New("consultation price not found")
)

// Dados fiscais da clínica para o recibo (nome legal + CNPJ + endereço fiscal).
// Fonte: docs/seo/01-NAP-master.md. Endereço fiscal (CNPJ), não o de atendimento.
const (
	clinicLegalName = "Plenya Saúde"
	clinicCNPJ      = "66.991.259/0001-50"
	clinicAddress   = "Av. Gil de Abreu e Souza, 2335, Casa 634 - Bairro Esperança, Londrina/PR, 86058-100"
	clinicPhone     = "(43) 99974-8899"
	clinicEmail     = "contato@plenyasaude.com.br"
	clinicSite      = "plenyasaude.com.br"
)

type PaymentService struct {
	db *gorm.DB
}

func NewPaymentService(db *gorm.DB) *PaymentService {
	return &PaymentService{db: db}
}

// ============================================================
// Pagamentos
// ============================================================

// Create registra um pagamento e gera o número de recibo sequencial do ano.
func (s *PaymentService) Create(createdBy uuid.UUID, req *dto.CreatePaymentRequest) (*dto.PaymentResponse, error) {
	patientID, err := uuid.Parse(req.PatientID)
	if err != nil {
		return nil, errors.New("invalid patient id")
	}

	paidAt := time.Now().UTC()
	if req.PaidAt != nil && *req.PaidAt != "" {
		t, err := time.Parse(time.RFC3339, *req.PaidAt)
		if err != nil {
			return nil, errors.New("invalid paidAt format, expected RFC3339")
		}
		paidAt = t
	}

	payment := models.AppointmentPayment{
		PatientID:       patientID,
		AmountCents:     req.AmountCents,
		Method:          req.Method,
		Status:          models.PaymentStatusPaid,
		PaidAt:          paidAt,
		Notes:           req.Notes,
		CreatedByUserID: createdBy,
	}
	if req.AppointmentID != nil && *req.AppointmentID != "" {
		aid, err := uuid.Parse(*req.AppointmentID)
		if err != nil {
			return nil, errors.New("invalid appointment id")
		}
		payment.AppointmentID = &aid
	}

	// Transação: numeração de recibo atômica (UPSERT com RETURNING) + insert.
	err = s.db.Transaction(func(tx *gorm.DB) error {
		year := paidAt.Year()
		var seq int
		if err := tx.Raw(
			`INSERT INTO payment_receipt_counters (year, last_seq) VALUES (?, 1)
			 ON CONFLICT (year) DO UPDATE SET last_seq = payment_receipt_counters.last_seq + 1
			 RETURNING last_seq`, year,
		).Scan(&seq).Error; err != nil {
			return err
		}
		payment.ReceiptNumber = fmt.Sprintf("%d-%06d", year, seq)
		return tx.Create(&payment).Error
	})
	if err != nil {
		return nil, err
	}

	return s.reload(payment.ID)
}

// List lista pagamentos, opcionalmente filtrando por paciente e período.
func (s *PaymentService) List(patientID *uuid.UUID, from, to *time.Time) ([]dto.PaymentResponse, error) {
	var payments []models.AppointmentPayment
	query := s.db.Preload("Patient").Order("paid_at DESC").Limit(500)
	if patientID != nil {
		query = query.Where("patient_id = ?", *patientID)
	}
	if from != nil {
		query = query.Where("paid_at >= ?", *from)
	}
	if to != nil {
		query = query.Where("paid_at < ?", *to)
	}
	if err := query.Find(&payments).Error; err != nil {
		return nil, err
	}
	result := make([]dto.PaymentResponse, len(payments))
	for i := range payments {
		result[i] = *s.paymentDTO(&payments[i])
	}
	return result, nil
}

// Refund estorna um pagamento.
func (s *PaymentService) Refund(id uuid.UUID, req *dto.RefundPaymentRequest) (*dto.PaymentResponse, error) {
	var p models.AppointmentPayment
	if err := s.db.Where("id = ?", id).First(&p).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPaymentNotFound
		}
		return nil, err
	}
	now := time.Now().UTC()
	p.Status = models.PaymentStatusRefunded
	p.RefundedAt = &now
	p.RefundReason = &req.Reason
	if err := s.db.Save(&p).Error; err != nil {
		return nil, err
	}
	return s.reload(p.ID)
}

// GenerateReceipt produz o PDF do recibo de um pagamento.
func (s *PaymentService) GenerateReceipt(id uuid.UUID) ([]byte, string, error) {
	var p models.AppointmentPayment
	if err := s.db.Preload("Patient").Where("id = ?", id).First(&p).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", ErrPaymentNotFound
		}
		return nil, "", err
	}
	pdf, err := s.buildReceiptPDF(&p)
	if err != nil {
		return nil, "", err
	}
	filename := fmt.Sprintf("recibo-%s.pdf", strings.ReplaceAll(p.ReceiptNumber, "/", "-"))
	return pdf, filename, nil
}

func (s *PaymentService) reload(id uuid.UUID) (*dto.PaymentResponse, error) {
	var p models.AppointmentPayment
	if err := s.db.Preload("Patient").Where("id = ?", id).First(&p).Error; err != nil {
		return nil, err
	}
	return s.paymentDTO(&p), nil
}

func (s *PaymentService) paymentDTO(p *models.AppointmentPayment) *dto.PaymentResponse {
	resp := &dto.PaymentResponse{
		ID:              p.ID.String(),
		PatientID:       p.PatientID.String(),
		AmountCents:     p.AmountCents,
		Method:          p.Method,
		Status:          p.Status,
		PaidAt:          p.PaidAt.Format(time.RFC3339),
		ReceiptNumber:   p.ReceiptNumber,
		Notes:           p.Notes,
		RefundReason:    p.RefundReason,
		CreatedByUserID: p.CreatedByUserID.String(),
		CreatedAt:       p.CreatedAt.Format(time.RFC3339),
	}
	if p.AppointmentID != nil {
		aid := p.AppointmentID.String()
		resp.AppointmentID = &aid
	}
	if p.RefundedAt != nil {
		r := p.RefundedAt.Format(time.RFC3339)
		resp.RefundedAt = &r
	}
	if p.Patient.ID != uuid.Nil {
		resp.Patient = &dto.AppointmentParty{
			ID:    p.Patient.ID.String(),
			Name:  p.Patient.Name,
			Email: p.Patient.Email,
			Phone: p.Patient.Phone,
		}
	}
	return resp
}

// ============================================================
// Preços por tipo de consulta
// ============================================================

// UpsertPrice cria ou atualiza o preço de um tipo de consulta.
func (s *PaymentService) UpsertPrice(req *dto.ConsultationPriceRequest) (*dto.ConsultationPriceResponse, error) {
	var price models.ConsultationPrice
	err := s.db.Where("type = ?", req.Type).First(&price).Error
	active := true
	if req.Active != nil {
		active = *req.Active
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		price = models.ConsultationPrice{Type: req.Type, AmountCents: req.AmountCents, Active: active}
		if err := s.db.Create(&price).Error; err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	} else {
		price.AmountCents = req.AmountCents
		price.Active = active
		if err := s.db.Save(&price).Error; err != nil {
			return nil, err
		}
	}
	return s.priceDTO(&price), nil
}

// ListPrices lista os preços cadastrados.
func (s *PaymentService) ListPrices() ([]dto.ConsultationPriceResponse, error) {
	var prices []models.ConsultationPrice
	if err := s.db.Order("type ASC").Find(&prices).Error; err != nil {
		return nil, err
	}
	result := make([]dto.ConsultationPriceResponse, len(prices))
	for i := range prices {
		result[i] = *s.priceDTO(&prices[i])
	}
	return result, nil
}

func (s *PaymentService) priceDTO(p *models.ConsultationPrice) *dto.ConsultationPriceResponse {
	return &dto.ConsultationPriceResponse{
		ID:          p.ID.String(),
		Type:        p.Type,
		AmountCents: p.AmountCents,
		Active:      p.Active,
		CreatedAt:   p.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   p.UpdatedAt.Format(time.RFC3339),
	}
}

// ============================================================
// Recibo PDF
// ============================================================

func (s *PaymentService) buildReceiptPDF(p *models.AppointmentPayment) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetAuthor(clinicLegalName, true)
	pdf.SetCreator("Plenya EMR", true)
	pdf.SetTitle("Recibo "+p.ReceiptNumber, true)

	// Fontes UTF-8 (mesmas usadas na prescrição).
	if reg, err := os.ReadFile("/usr/share/fonts/opensans/OpenSans-Regular.ttf"); err == nil {
		pdf.AddUTF8FontFromBytes("OpenSans", "", reg)
	}
	if bold, err := os.ReadFile("/usr/share/fonts/opensans/OpenSans-Bold.ttf"); err == nil {
		pdf.AddUTF8FontFromBytes("OpenSans", "B", bold)
	}
	pdf.SetFont("OpenSans", "", 11)
	pdf.AddPage()

	const left = 20.0
	y := 25.0

	// Cabeçalho — emitente.
	pdf.SetFont("OpenSans", "B", 16)
	pdf.SetXY(left, y)
	pdf.Cell(170, 9, "RECIBO")
	pdf.SetFont("OpenSans", "", 10)
	pdf.SetXY(left, y)
	pdf.CellFormat(170, 9, "Nº "+p.ReceiptNumber, "", 0, "R", false, 0, "")
	y += 14

	pdf.SetFont("OpenSans", "B", 12)
	pdf.SetXY(left, y)
	pdf.Cell(170, 6, clinicLegalName)
	y += 6
	pdf.SetFont("OpenSans", "", 9)
	for _, line := range []string{
		"CNPJ: " + clinicCNPJ,
		clinicAddress,
		"Telefone: " + clinicPhone + "   Email: " + clinicEmail,
		"Site: " + clinicSite,
	} {
		pdf.SetXY(left, y)
		pdf.Cell(170, 5, line)
		y += 5
	}
	y += 6

	// Linha divisória.
	pdf.SetDrawColor(200, 200, 200)
	pdf.Line(left, y, left+170, y)
	y += 10

	// Corpo — valor por extenso e descrição.
	valor := formatBRL(p.AmountCents)
	pdf.SetFont("OpenSans", "B", 13)
	pdf.SetXY(left, y)
	pdf.Cell(170, 8, "Valor: "+valor)
	y += 12

	pacienteNome := ""
	if p.Patient.ID != uuid.Nil {
		pacienteNome = p.Patient.Name
	}
	corpo := fmt.Sprintf(
		"Recebi de %s a importância de %s, referente a serviços de saúde prestados, paga via %s em %s.",
		pacienteNome, valor, paymentMethodLabel(p.Method), p.PaidAt.Format("02/01/2006"),
	)
	pdf.SetFont("OpenSans", "", 11)
	pdf.SetXY(left, y)
	pdf.MultiCell(170, 6, corpo, "", "L", false)
	y = pdf.GetY() + 4

	if p.Notes != nil && *p.Notes != "" {
		pdf.SetFont("OpenSans", "", 9)
		pdf.SetXY(left, y)
		pdf.MultiCell(170, 5, "Observações: "+*p.Notes, "", "L", false)
		y = pdf.GetY() + 4
	}

	if p.Status == models.PaymentStatusRefunded {
		pdf.SetFont("OpenSans", "B", 11)
		pdf.SetTextColor(180, 30, 30)
		pdf.SetXY(left, y)
		reason := ""
		if p.RefundReason != nil {
			reason = " - " + *p.RefundReason
		}
		pdf.MultiCell(170, 6, "PAGAMENTO ESTORNADO"+reason, "", "L", false)
		pdf.SetTextColor(0, 0, 0)
		y = pdf.GetY() + 4
	}

	// Rodapé — local, data e assinatura.
	y += 20
	pdf.SetFont("OpenSans", "", 10)
	pdf.SetXY(left, y)
	pdf.Cell(170, 6, "Londrina, "+p.PaidAt.Format("02/01/2006"))
	y += 18
	pdf.Line(left, y, left+90, y)
	y += 5
	pdf.SetXY(left, y)
	pdf.Cell(90, 5, clinicLegalName)

	var buf bytes.Buffer
	if err := pdf.Output(&buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func paymentMethodLabel(m models.PaymentMethod) string {
	switch m {
	case models.PaymentMethodCash:
		return "dinheiro"
	case models.PaymentMethodPix:
		return "Pix"
	case models.PaymentMethodCard:
		return "cartão"
	case models.PaymentMethodTransfer:
		return "transferência"
	default:
		return "outro"
	}
}

// formatBRL converte centavos em "R$ 1.234,56".
func formatBRL(cents int64) string {
	neg := cents < 0
	if neg {
		cents = -cents
	}
	reais := cents / 100
	cent := cents % 100
	// Separador de milhar.
	s := fmt.Sprintf("%d", reais)
	var parts []string
	for len(s) > 3 {
		parts = append([]string{s[len(s)-3:]}, parts...)
		s = s[:len(s)-3]
	}
	parts = append([]string{s}, parts...)
	out := fmt.Sprintf("R$ %s,%02d", strings.Join(parts, "."), cent)
	if neg {
		out = "-" + out
	}
	return out
}
