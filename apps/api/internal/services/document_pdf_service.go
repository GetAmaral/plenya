package services

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	gofpdf "codeberg.org/go-pdf/fpdf"
	"github.com/skip2/go-qrcode"

	"github.com/plenya/api/internal/models"
)

// DocumentPDFService gera o PDF (gofpdf) de documentos clínicos emitidos
// (atestado/declaração/laudo). Molde: prescription_pdf_service + pdf_service.addSignatureSection,
// com bloco de assinatura generalizado (digital ICP-Brasil ou manual, conforme cert do médico).
type DocumentPDFService struct{}

func NewDocumentPDFService() *DocumentPDFService { return &DocumentPDFService{} }

func docTypeTitle(t models.IssuedDocumentType) string {
	switch t {
	case models.IssuedDocCertificate:
		return "ATESTADO MÉDICO"
	case models.IssuedDocDeclaration:
		return "DECLARAÇÃO"
	case models.IssuedDocReport:
		return "RELATÓRIO MÉDICO"
	default:
		return "DOCUMENTO MÉDICO"
	}
}

// BuildDocumentPDF monta o PDF não-assinado (com bloco de assinatura digital ou manual).
// validationURL é o destino do QR (página pública de validação por hash).
func (s *DocumentPDFService) BuildDocumentPDF(
	doc *models.IssuedDocument,
	patient *models.Patient,
	doctor *models.User,
	hasDigitalSignature bool,
	validationURL string,
) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetAuthor("Plenya EMR - Documentos Clínicos", true)
	pdf.SetCreator("Plenya EMR", true)
	pdf.SetTitle(doc.Title, true)

	regularFont, err := os.ReadFile("/usr/share/fonts/opensans/OpenSans-Regular.ttf")
	if err != nil {
		return nil, err
	}
	boldFont, err := os.ReadFile("/usr/share/fonts/opensans/OpenSans-Bold.ttf")
	if err != nil {
		return nil, err
	}
	pdf.AddUTF8FontFromBytes("OpenSans", "", regularFont)
	pdf.AddUTF8FontFromBytes("OpenSans", "B", boldFont)
	pdf.SetFont("OpenSans", "", 10)
	pdf.SetAutoPageBreak(true, 40)
	pdf.AddPage()

	// Letterhead
	pdf.ImageOptions("/app/PlenyaA4-150dpi.png", 0, 0, 210, 297, false,
		gofpdf.ImageOptions{ImageType: "PNG"}, 0, "")

	y := 50.0

	// Título
	pdf.SetFont("OpenSans", "B", 16)
	pdf.SetXY(20, y)
	pdf.Cell(170, 10, docTypeTitle(doc.Type))
	y += 14

	// Data de emissão
	pdf.SetFont("OpenSans", "", 10)
	pdf.SetXY(20, y)
	pdf.Cell(170, 6, fmt.Sprintf("Emissão: %s", doc.IssuedAt.Format("02/01/2006")))
	y += 12

	// Médico
	pdf.SetFont("OpenSans", "B", 11)
	pdf.SetXY(20, y)
	pdf.Cell(170, 6, "MÉDICO")
	y += 6
	pdf.SetFont("OpenSans", "", 10)
	doctorInfo := doctor.Name
	if doctor.CRM != nil && doctor.CRMUF != nil {
		doctorInfo += fmt.Sprintf("\nCRM-%s %s", *doctor.CRMUF, *doctor.CRM)
	}
	if doctor.Specialty != nil && *doctor.Specialty != "" {
		doctorInfo += " - " + *doctor.Specialty
	}
	pdf.SetXY(20, y)
	pdf.MultiCell(170, 5, doctorInfo, "", "", false)
	y += 16

	// Paciente
	pdf.SetFont("OpenSans", "B", 11)
	pdf.SetXY(20, y)
	pdf.Cell(170, 6, "PACIENTE")
	y += 6
	pdf.SetFont("OpenSans", "", 10)
	patientInfo := patient.Name
	if patient.CPF != nil && *patient.CPF != "" {
		patientInfo += fmt.Sprintf("\nCPF: %s", formatCPF(*patient.CPF))
	}
	pdf.SetXY(20, y)
	pdf.MultiCell(170, 5, patientInfo, "", "", false)
	y += 16

	// Corpo do documento
	if strings.TrimSpace(doc.Body) != "" {
		pdf.SetFont("OpenSans", "", 11)
		pdf.SetXY(20, y)
		pdf.MultiCell(170, 6, doc.Body, "", "", false)
		y = pdf.GetY() + 6
	}

	// Afastamento (atestado)
	if doc.DaysOff != nil && *doc.DaysOff > 0 {
		dias := "dia"
		if *doc.DaysOff != 1 {
			dias = "dias"
		}
		pdf.SetFont("OpenSans", "B", 11)
		pdf.SetXY(20, y)
		pdf.Cell(170, 6, fmt.Sprintf("Recomenda-se afastamento de %d %s.", *doc.DaysOff, dias))
		y += 10
	}

	// CID (só com consentimento)
	if doc.IncludesCID && doc.CIDConsent && doc.CIDCode != nil && *doc.CIDCode != "" {
		pdf.SetFont("OpenSans", "", 10)
		pdf.SetXY(20, y)
		pdf.Cell(170, 6, fmt.Sprintf("CID-10: %s (incluído com consentimento do paciente)", *doc.CIDCode))
		y += 10
	}

	// Finalidade
	if doc.Purpose != nil && strings.TrimSpace(*doc.Purpose) != "" {
		pdf.SetFont("OpenSans", "", 10)
		pdf.SetXY(20, y)
		pdf.MultiCell(170, 5, "Finalidade: "+*doc.Purpose, "", "", false)
		y = pdf.GetY() + 8
	}

	// Bloco de assinatura
	if y > 230 {
		y = 230
	}
	addDocSignatureSection(pdf, doc, doctor, y, doc.IssuedAt.Format("02/01/2006"), hasDigitalSignature, validationURL)

	var buf bytes.Buffer
	if err := pdf.Output(&buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// addDocSignatureSection — bloco de assinatura generalizado (de pdf_service.addSignatureSection),
// recebendo Doctor + URL de validação em vez de *models.LabRequest.
func addDocSignatureSection(
	pdf *gofpdf.Fpdf,
	doc *models.IssuedDocument,
	doctor *models.User,
	startY float64,
	formattedDate string,
	hasDigitalSignature bool,
	validationURL string,
) {
	currentY := startY

	pdf.SetFont("OpenSans", "", 9)
	pdf.SetXY(25, currentY)
	pdf.Cell(125, 4, formattedDate)
	currentY += 6

	if hasDigitalSignature {
		imageSize := 30.0
		sealX := 130.0
		qrX := sealX + imageSize + 3
		imagesY := startY

		pdf.ImageOptions("/app/icp-brasil-seal.png", sealX, imagesY, imageSize, imageSize, false,
			gofpdf.ImageOptions{ImageType: "PNG"}, 0, "")

		qrCodePath := fmt.Sprintf("/tmp/qr_doc_%s.png", doc.ID)
		qrCode, qrErr := qrcode.Encode(validationURL, qrcode.Medium, 256)
		if qrErr == nil {
			_ = os.WriteFile(qrCodePath, qrCode, 0644)
			pdf.ImageOptions(qrCodePath, qrX, imagesY, imageSize, imageSize, false,
				gofpdf.ImageOptions{ImageType: "PNG"}, 0, "")
			pdf.SetFont("OpenSans", "", 6)
			tw := pdf.GetStringWidth("Escaneie para validar")
			pdf.SetXY(qrX+(imageSize-tw)/2, imagesY+imageSize+1)
			pdf.Cell(tw, 2, "Escaneie para validar")
		}

		textX := 20.0
		textWidth := 105.0

		pdf.SetFont("OpenSans", "B", 9)
		pdf.SetXY(textX, currentY)
		pdf.SetTextColor(0, 120, 0)
		pdf.Cell(textWidth, 3.5, "DOCUMENTO ASSINADO DIGITALMENTE")
		pdf.SetTextColor(0, 0, 0)
		currentY += 4

		pdf.SetFont("OpenSans", "", 7)
		pdf.SetXY(textX, currentY)
		pdf.Cell(textWidth, 2.5, "Certificado ICP-Brasil (PAdES) - Plenya")
		currentY += 3

		if doctor.CertificateName != nil && *doctor.CertificateName != "" {
			pdf.SetXY(textX, currentY)
			pdf.Cell(textWidth, 2.5, "Titular: "+*doctor.CertificateName)
			currentY += 2.5
		}
		if doctor.CRM != nil && doctor.CRMUF != nil {
			pdf.SetXY(textX, currentY)
			pdf.Cell(textWidth, 2.5, "CRM-"+*doctor.CRMUF+": "+*doctor.CRM)
			currentY += 2.5
		}
		certInfoLine := ""
		if doctor.CertificateSerial != nil && *doctor.CertificateSerial != "" {
			certInfoLine = "Serie: " + *doctor.CertificateSerial
		}
		if doctor.CertificateCPF != nil && *doctor.CertificateCPF != "" {
			if certInfoLine != "" {
				certInfoLine += " | "
			}
			certInfoLine += "CPF: " + *doctor.CertificateCPF
		}
		if certInfoLine != "" {
			pdf.SetXY(textX, currentY)
			pdf.Cell(textWidth, 2.5, certInfoLine)
			currentY += 2.5
		}

		currentY += 1
		pdf.SetFont("OpenSans", "", 7)
		pdf.SetXY(textX, currentY)
		pdf.Cell(textWidth, 2.5, "Validar documento:")
		currentY += 2.5
		pdf.SetXY(textX, currentY)
		pdf.Cell(textWidth, 2.5, validationURL)
		currentY += 3
		pdf.SetXY(textX, currentY)
		pdf.Cell(textWidth, 2.5, "Verificar assinatura: https://validar.iti.gov.br")
		currentY += 3

		pdf.SetFont("OpenSans", "", 6.5)
		pdf.SetTextColor(60, 60, 60)
		pdf.SetXY(textX, currentY)
		pdf.Cell(textWidth, 2.5, "Documento assinado digitalmente (ICP-Brasil) - CFM 2.299/2021.")
		currentY += 2.5
		pdf.SetXY(textX, currentY)
		pdf.Cell(textWidth, 2.5, "Validade juridica equivalente ao documento fisico.")
		pdf.SetTextColor(0, 0, 0)

		if qrErr == nil {
			os.Remove(qrCodePath)
		}
	} else {
		currentY += 12
		pdf.Line(60, currentY, 150, currentY)
		pdf.SetFont("OpenSans", "", 9)
		pdf.SetXY(25, currentY+2)
		pdf.Cell(120, 5, "Assinatura e Carimbo do Médico")
	}
}
