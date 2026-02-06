package services

import (
	"fmt"
	"os"
	"strings"

	"github.com/jung-kurt/gofpdf/v2"
	"github.com/skip2/go-qrcode"
	"github.com/plenya/api/internal/crypto"
	"github.com/plenya/api/internal/models"
)

// PDFService handles PDF generation
type PDFService struct{}

// NewPDFService creates a new PDF service
func NewPDFService() *PDFService {
	return &PDFService{}
}

// getInitials extracts initials from a full name (e.g., "João da Silva" -> "JS")
func getInitials(fullName string) string {
	words := strings.Fields(fullName)
	if len(words) == 0 {
		return "XX"
	}

	initials := ""
	for _, word := range words {
		// Skip common Portuguese prepositions/articles
		lower := strings.ToLower(word)
		if lower == "da" || lower == "de" || lower == "do" || lower == "dos" || lower == "das" {
			continue
		}
		if len(word) > 0 {
			initials += strings.ToUpper(string(word[0]))
		}
	}

	if initials == "" {
		return "XX"
	}
	return initials
}

// drawLetterhead draws the Plenya letterhead (header + footer + watermark) using vector graphics
func drawLetterhead(pdf *gofpdf.Fpdf) {
	// Page border (subtle)
	pdf.SetDrawColor(208, 208, 208) // #D0D0D0
	pdf.SetLineWidth(0.1)
	pdf.Rect(2, 2, 206, 293, "D")

	// === HEADER SECTION ===

	// Logo image (small PNG - top center)
	pdf.ImageOptions(
		"/app/assets/logo-only.png",
		80, 8,    // x, y position (centered)
		50, 0,    // width, height (0 = auto maintain aspect)
		false,
		gofpdf.ImageOptions{ImageType: "PNG"},
		0, "",
	)

	// "PLENYA" text (serif font)
	pdf.SetFont("OpenSans", "B", 18)
	pdf.SetTextColor(107, 107, 107) // #6B6B6B
	pdf.SetXY(0, 22)
	pdf.CellFormat(210, 8, "PLENYA", "", 0, "C", false, 0, "")

	// Tagline "Saúde, Performance e Longevidade" (italic, smaller)
	pdf.SetFont("OpenSans", "", 8)
	pdf.SetTextColor(153, 153, 153) // #999999
	pdf.SetXY(0, 30)
	pdf.CellFormat(210, 4, "Saúde, Performance e Longevidade", "", 0, "C", false, 0, "")

	// Header line (decorative)
	pdf.SetDrawColor(184, 151, 124) // #B8977C (gold/beige)
	pdf.SetLineWidth(0.15)
	pdf.SetAlpha(0.6, "Normal")
	pdf.Line(15, 38, 195, 38)
	pdf.SetAlpha(1.0, "Normal")

	// === WATERMARK (center) ===
	pdf.SetAlpha(0.08, "Normal")
	pdf.ImageOptions(
		"/app/assets/watermark-only.png",
		55, 120,  // x, y position (centered vertically)
		100, 0,   // width, height (0 = auto)
		false,
		gofpdf.ImageOptions{ImageType: "PNG"},
		0, "",
	)
	pdf.SetAlpha(1.0, "Normal")

	// === FOOTER SECTION ===

	// Footer line (decorative)
	pdf.SetDrawColor(184, 151, 124) // #B8977C
	pdf.SetLineWidth(0.15)
	pdf.SetAlpha(0.6, "Normal")
	pdf.Line(15, 273, 195, 273)
	pdf.SetAlpha(1.0, "Normal")

	// Footer text (address)
	pdf.SetFont("OpenSans", "", 7)
	pdf.SetTextColor(170, 170, 170) // #AAAAAA
	pdf.SetXY(0, 278)
	pdf.CellFormat(210, 3, "Av. Duque de Caxias, 1371 – Londrina – PR - (43) 99638-0044", "", 0, "C", false, 0, "")

	// Footer text (website and email)
	pdf.SetXY(0, 283)
	pdf.CellFormat(210, 3, "www.plenya.com.br - contato@plenya.com.br", "", 0, "C", false, 0, "")

	// Reset colors for content
	pdf.SetTextColor(0, 0, 0)
	pdf.SetDrawColor(0, 0, 0)
}

// addSignatureSection adiciona a seção de assinatura digital ou manual em uma posição Y específica
func addSignatureSection(pdf *gofpdf.Fpdf, req *models.LabRequest, startY float64, formattedDate string, hasDigitalSignature bool) {
	currentY := startY

	// Data
	pdf.SetFont("OpenSans", "", 9)
	pdf.SetXY(25, currentY)
	pdf.Cell(125, 4, formattedDate)
	currentY += 6

	if hasDigitalSignature {
		// ============================================
		// ASSINATURA DIGITAL (CFM 2.299/2021)
		// ============================================

		// Largura máxima para textos: 125mm (deixa espaço para QR/selo à direita)
		textWidth := 125.0

		// Título verde principal
		pdf.SetFont("OpenSans", "B", 9)
		pdf.SetXY(25, currentY)
		pdf.SetTextColor(0, 120, 0)
		pdf.Cell(textWidth, 4, "DOCUMENTO ASSINADO DIGITALMENTE")
		pdf.SetTextColor(0, 0, 0)
		currentY += 5

		// Plataforma emissora (legislação)
		pdf.SetFont("OpenSans", "", 8)
		pdf.SetXY(25, currentY)
		pdf.Cell(textWidth, 3, "Certificado ICP-Brasil (PAdES) - Plenya Sistema de Prescricao Digital")
		currentY += 4

		// Informações do certificado
		if req.Doctor != nil {
			// Titular
			if req.Doctor.CertificateName != nil && *req.Doctor.CertificateName != "" {
				pdf.SetXY(25, currentY)
				pdf.SetFont("OpenSans", "", 8)
				pdf.Cell(textWidth, 3, "Titular: "+*req.Doctor.CertificateName)
				currentY += 3
			}

			// Conselho profissional (CRM-UF: número)
			if req.Doctor.CRM != nil && req.Doctor.CRMUF != nil {
				pdf.SetXY(25, currentY)
				pdf.SetFont("OpenSans", "", 8)
				crmLine := "CRM-" + *req.Doctor.CRMUF + ": " + *req.Doctor.CRM
				pdf.Cell(textWidth, 3, crmLine)
				currentY += 3
			}

			// Serie e CPF na mesma linha separados por |
			certInfoLine := ""
			if req.Doctor.CertificateSerial != nil && *req.Doctor.CertificateSerial != "" {
				certInfoLine = "Serie: " + *req.Doctor.CertificateSerial
			}
			if req.Doctor.CertificateCPF != nil && *req.Doctor.CertificateCPF != "" {
				if certInfoLine != "" {
					certInfoLine += " | "
				}
				certInfoLine += "CPF: " + *req.Doctor.CertificateCPF
			}
			if certInfoLine != "" {
				pdf.SetXY(25, currentY)
				pdf.SetFont("OpenSans", "", 8)
				pdf.Cell(textWidth, 3, certInfoLine)
				currentY += 3
			}
		}

		currentY += 2

		// Links de validação (em 2 linhas para não encavalar)
		pdf.SetFont("OpenSans", "", 8)
		pdf.SetXY(25, currentY)
		pdf.Cell(textWidth, 3, "Validar documento:")
		currentY += 3

		pdf.SetXY(25, currentY)
		fullValidationURL := "https://plenya.com.br/lab-requests/validate/" + req.ID.String()
		pdf.Cell(textWidth, 3, fullValidationURL)
		currentY += 4

		pdf.SetXY(25, currentY)
		pdf.Cell(textWidth, 3, "Verificar assinatura: https://validar.iti.gov.br")
		currentY += 4

		// Nota sobre validade legal
		pdf.SetFont("OpenSans", "", 8)
		pdf.SetTextColor(60, 60, 60)
		pdf.SetXY(25, currentY)
		pdf.MultiCell(textWidth, 3, "Este documento foi assinado digitalmente com certificado ICP-Brasil, "+
			"conforme Resolucao CFM 2.299/2021 e RDC ANVISA 1.000/2025. "+
			"Possui validade juridica equivalente ao documento fisico.", "", "L", false)
		pdf.SetTextColor(0, 0, 0)

		// QR CODE E SELO
		qrX := 155.0
		qrSize := 30.0
		qrStartY := startY

		// Selo ICP-Brasil
		sealWidth := 20.0
		sealHeight := sealWidth * (247.0 / 204.0)

		pdf.ImageOptions(
			"/app/icp-brasil-seal.png",
			qrX, qrStartY, sealWidth, sealHeight,
			false,
			gofpdf.ImageOptions{ImageType: "PNG"},
			0, "",
		)
		qrStartY += sealHeight + 2

		// QR Code
		qrCodeURL := fmt.Sprintf("https://plenya.com.br/lab-requests/validate/%s", req.ID)
		qrCodePath := fmt.Sprintf("/tmp/qr_lab_%s.png", req.ID)

		qrCode, qrErr := qrcode.Encode(qrCodeURL, qrcode.Medium, 256)
		if qrErr == nil {
			os.WriteFile(qrCodePath, qrCode, 0644)

			pdf.ImageOptions(
				qrCodePath,
				qrX, qrStartY, qrSize, qrSize,
				false,
				gofpdf.ImageOptions{ImageType: "PNG"},
				0, "",
			)

			pdf.SetFont("OpenSans", "", 7)
			pdf.SetXY(qrX-2, qrStartY+qrSize+1)
			pdf.MultiCell(qrSize+4, 2.5, "Escaneie para\nvalidar", "", "C", false)

			os.Remove(qrCodePath)
		}

	} else {
		// ============================================
		// ASSINATURA MANUAL (sem certificado digital)
		// ============================================

		currentY += 8

		// Linha para assinatura manual
		pdf.Line(60, currentY, 150, currentY)

		pdf.SetFont("OpenSans", "", 9)
		pdf.SetXY(25, currentY+2)
		pdf.Cell(120, 5, "Assinatura e Carimbo do Medico Solicitante")
	}
}

// GenerateLabRequestPDF generates a PDF for a lab request
func (s *PDFService) GenerateLabRequestPDF(req *models.LabRequest) (string, error) {
	// Create PDF (A4 portrait) with UTF-8 encoding support
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetAuthor("Plenya", true)
	pdf.SetCreator("Plenya EMR", true)

	// Add Open Sans fonts (modern, elegant) - reading from bytes to avoid path issues
	regularFont, err := os.ReadFile("/usr/share/fonts/opensans/OpenSans-Regular.ttf")
	if err != nil {
		return "", fmt.Errorf("failed to read regular font: %v", err)
	}
	boldFont, err := os.ReadFile("/usr/share/fonts/opensans/OpenSans-Bold.ttf")
	if err != nil {
		return "", fmt.Errorf("failed to read bold font: %v", err)
	}
	semiBoldFont, err := os.ReadFile("/usr/share/fonts/opensans/OpenSans-SemiBold.ttf")
	if err != nil {
		return "", fmt.Errorf("failed to read semibold font: %v", err)
	}

	pdf.AddUTF8FontFromBytes("OpenSans", "", regularFont)
	pdf.AddUTF8FontFromBytes("OpenSans", "B", boldFont)
	pdf.AddUTF8FontFromBytes("OpenSans", "SB", semiBoldFont)

	// UTF-8 fonts handle Portuguese characters directly - no translator needed
	pdf.SetFont("OpenSans", "", 10)

	pdf.AddPage()

	// Add letterhead background image (150 DPI - optimized quality/size)
	pdf.ImageOptions(
		"/app/PlenyaA4-150dpi.png",
		0, 0, 210, 297,
		false,
		gofpdf.ImageOptions{ImageType: "PNG"},
		0, "",
	)

	// Get patient info
	patientName := "Não informado"
	patientCPF := ""
	patientBirthDate := ""

	if req.Patient != nil {
		patientName = req.Patient.Name
		if req.Patient.CPF != nil && *req.Patient.CPF != "" {
			// Decrypt CPF if encrypted (base64)
			cpf := *req.Patient.CPF
			if len(cpf) > 11 {
				decrypted, err := crypto.DecryptWithDefaultKey(cpf)
				if err == nil {
					cpf = decrypted
				}
			}

			// Format CPF: 000.000.000-00
			if len(cpf) == 11 {
				patientCPF = fmt.Sprintf("%s.%s.%s-%s", cpf[0:3], cpf[3:6], cpf[6:9], cpf[9:11])
			} else {
				patientCPF = cpf
			}
		}
		if !req.Patient.BirthDate.IsZero() {
			patientBirthDate = req.Patient.BirthDate.Format("02/01/2006")
		}
	}

	// Format date in Portuguese (will be used near signature)
	formattedDate := req.Date.Format("02/01/2006")

	// Parse exams
	exams := []string{}
	for _, exam := range strings.Split(req.Exams, "\n") {
		exam = strings.TrimSpace(exam)
		if exam != "" {
			exams = append(exams, exam)
		}
	}

	// Content starts at Y=50mm (closer to letterhead)
	currentY := 50.0

	// Title - bold font
	pdf.SetFont("OpenSans", "B", 15)
	pdf.SetXY(25, currentY)
	pdf.CellFormat(160, 8, "Solicitação de Exames", "", 0, "C", false, 0, "")
	currentY += 14

	// PACIENTE
	pdf.SetFont("OpenSans", "B", 10)
	pdf.SetXY(25, currentY)
	pdf.Cell(160, 5, "PACIENTE")
	currentY += 6

	// Patient name
	pdf.SetFont("OpenSans", "B", 12)
	pdf.SetXY(25, currentY)
	pdf.Cell(160, 6, patientName)
	currentY += 7

	// CPF and Birth Date on same line
	pdf.SetFont("OpenSans", "", 9)
	pdf.SetXY(25, currentY)

	infoLine := ""
	if patientCPF != "" {
		infoLine = "CPF: " + patientCPF
	}
	if patientBirthDate != "" {
		if infoLine != "" {
			infoLine += "    "
		}
		infoLine += "Nascimento: " + patientBirthDate
	}

	if infoLine != "" {
		pdf.Cell(160, 5, infoLine)
		currentY += 6
	}

	// Endereço do paciente (obrigatório pela legislação)
	if req.Patient != nil && req.Patient.Address != nil && *req.Patient.Address != "" {
		pdf.SetXY(25, currentY)
		addressLine := "Endereço: " + *req.Patient.Address
		// Adiciona município e estado se disponível
		if req.Patient.Municipality != nil && *req.Patient.Municipality != "" {
			addressLine += " - " + *req.Patient.Municipality
			if req.Patient.State != nil && *req.Patient.State != "" {
				addressLine += "/" + *req.Patient.State
			}
		}
		pdf.Cell(160, 5, addressLine)
		currentY += 6
	}

	currentY += 4

	// Section: Exames Solicitados - bold font
	pdf.SetFont("OpenSans", "B", 11)
	pdf.SetXY(25, currentY)
	pdf.Cell(160, 6, "Exames Solicitados")
	currentY += 8

	// Exams list - adaptive layout based on ALTURA (não número de itens)
	pdf.SetFont("OpenSans", "", 10)

	// Preparar variável para assinatura (usada no loop e no final)
	hasDigitalSignature := req.SignedAt != nil

	// Altura máxima para container de exames (deixa espaço para assinatura em Y=200mm)
	maxExamContainerY := 185.0 // Container termina em 185mm

	// Column dimensions
	leftColX := 25.0
	rightColX := 110.0
	colWidth := 75.0

	columnY := currentY
	startY := currentY // Save starting Y position for columns
	currentColumn := 1 // 1 = esquerda, 2 = direita

	for _, exam := range exams {
		// Verificar se precisa trocar de coluna ou página ANTES de adicionar item
		estimatedHeight := 6.0 // Altura estimada por item

		if columnY + estimatedHeight > maxExamContainerY {
			if currentColumn == 1 {
				// Coluna 1 cheia -> passar para coluna 2
				columnY = startY
				currentColumn = 2
			} else {
				// Coluna 2 cheia -> Adicionar assinatura na página ATUAL antes de sair
				addSignatureSection(pdf, req, 200.0, formattedDate, hasDigitalSignature)

				// Agora criar nova página
				pdf.AddPage()

				// Draw letterhead on new page
				pdf.ImageOptions(
					"/app/PlenyaA4-150dpi.png",
					0, 0, 210, 297,
					false,
					gofpdf.ImageOptions{ImageType: "PNG"},
					0, "",
				)

				// Reset position for new page
				columnY = 50.0
				startY = 50.0
				currentColumn = 1

				// Add section title again on new page
				pdf.SetFont("OpenSans", "B", 11)
				pdf.SetXY(25, columnY)
				pdf.Cell(160, 6, "Exames Solicitados (continuação)")
				columnY += 8
				startY = columnY
				pdf.SetFont("OpenSans", "", 10)
			}
		}

		// Determine X position based on current column
		xPos := leftColX
		if currentColumn == 2 {
			xPos = rightColX
		}

		// Add bullet and exam text
		pdf.SetXY(xPos, columnY)
		pdf.Cell(5, 5, "•")
		pdf.SetXY(xPos+5, columnY)

		// Truncate long names if needed
		examText := exam
		if len(examText) > 45 {
			examText = examText[:42] + "..."
		}
		pdf.Cell(colWidth-5, 5, examText)

		// Move to next row
		columnY += 6
	}

	// Update currentY to fixed position (após container de exames)
	currentY = maxExamContainerY + 5 // Adiciona pequeno espaçamento

	// Notes section
	if req.Notes != nil && *req.Notes != "" {
		currentY += 6
		pdf.SetFont("OpenSans", "B", 10)
		pdf.SetXY(25, currentY)
		pdf.Cell(160, 6, "Observações:")
		currentY += 6

		pdf.SetFont("OpenSans", "", 9)
		pdf.SetXY(25, currentY)
		pdf.MultiCell(160, 5, *req.Notes, "", "L", false)
	}

	// ASSINATURA: Verificar se precisa nova página (precisa de ~70mm de espaço)
	if currentY > 220.0 {
		pdf.AddPage()

		// Draw letterhead on new page
		pdf.ImageOptions(
			"/app/PlenyaA4-150dpi.png",
			0, 0, 210, 297,
			false,
			gofpdf.ImageOptions{ImageType: "PNG"},
			0, "",
		)

		currentY = 50.0 // Começa do topo se for nova página
	}

	// Adiciona espaçamento antes da assinatura
	currentY += 10

	// Garantir posição mínima para assinatura
	if currentY < 200.0 {
		currentY = 200.0
	}

	// ========================================
	// SEÇÃO DE ASSINATURA (CFM 2.299/2021 e RDC 1.000/2025)
	// ========================================
	addSignatureSection(pdf, req, currentY, formattedDate, hasDigitalSignature)

	// Generate filename: InitialsUppercase_labRequest_YYYY-MM-DD_hash.pdf
	initials := getInitials(patientName)
	dateStr := req.Date.Format("2006-01-02")
	hash := req.ID.String()[:8] // First 8 chars of UUID as hash
	filename := fmt.Sprintf("%s_labRequest_%s_%s.pdf", initials, dateStr, hash)
	pdfPath := fmt.Sprintf("/app/uploads/lab-requests/%s", filename)

	// Check for any errors that occurred during PDF generation
	if pdf.Error() != nil {
		return "", fmt.Errorf("PDF generation error: %v", pdf.Error())
	}

	err = pdf.OutputFileAndClose(pdfPath)
	if err != nil {
		return "", fmt.Errorf("failed to save PDF file: %v", err)
	}

	// Return URL
	return fmt.Sprintf("/uploads/lab-requests/%s", filename), nil
}
