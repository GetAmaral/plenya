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

		// SELO E QR CODE LADO A LADO (invertidos: selo à esquerda, QR à direita)
		// A4 = 210mm, margem direita 15mm, então limite = 195mm
		imageSize := 30.0 // QR e selo com mesmo tamanho (30x30mm)
		sealX := 130.0    // Selo à esquerda
		qrX := sealX + imageSize + 3 // QR ao lado do selo com 3mm de espaço (= 163mm)
		// Fim do QR = 163 + 30 = 193mm (dentro da margem de 195mm)
		imagesY := startY

		// Selo ICP-Brasil (à esquerda)
		pdf.ImageOptions(
			"/app/icp-brasil-seal.png",
			sealX, imagesY, imageSize, imageSize,
			false,
			gofpdf.ImageOptions{ImageType: "PNG"},
			0, "",
		)

		// QR Code (à direita)
		qrCodeURL := fmt.Sprintf("https://plenya.com.br/lab-requests/validate/%s", req.ID)
		qrCodePath := fmt.Sprintf("/tmp/qr_lab_%s.png", req.ID)

		qrCode, qrErr := qrcode.Encode(qrCodeURL, qrcode.Medium, 256)
		if qrErr == nil {
			os.WriteFile(qrCodePath, qrCode, 0644)

			pdf.ImageOptions(
				qrCodePath,
				qrX, imagesY, imageSize, imageSize,
				false,
				gofpdf.ImageOptions{ImageType: "PNG"},
				0, "",
			)

			// Texto "Escaneie para validar" centralizado com o QR code
			pdf.SetFont("OpenSans", "", 6)
			textWidth := pdf.GetStringWidth("Escaneie para validar")
			textX := qrX + (imageSize-textWidth)/2 // Centralizar com o QR
			pdf.SetXY(textX, imagesY+imageSize+1)
			pdf.Cell(textWidth, 2, "Escaneie para validar")
		}

		// TEXTOS DA ASSINATURA (largura reduzida para não conflitar com imagens)
		textX := 20.0              // Recuo menor
		textWidth := 105.0         // Largura reduzida (era 125) para dar espaço às imagens

		// Título verde principal
		pdf.SetFont("OpenSans", "B", 9)
		pdf.SetXY(textX, currentY)
		pdf.SetTextColor(0, 120, 0)
		pdf.Cell(textWidth, 3.5, "DOCUMENTO ASSINADO DIGITALMENTE")
		pdf.SetTextColor(0, 0, 0)
		currentY += 4

		// Plataforma emissora (legislação) - linha mais compacta
		pdf.SetFont("OpenSans", "", 7)
		pdf.SetXY(textX, currentY)
		pdf.Cell(textWidth, 2.5, "Certificado ICP-Brasil (PAdES) - Plenya Sistema de Prescricao Digital")
		currentY += 3

		// Informações do certificado - mais compactas
		if req.Doctor != nil {
			pdf.SetFont("OpenSans", "", 7)

			// Titular
			if req.Doctor.CertificateName != nil && *req.Doctor.CertificateName != "" {
				pdf.SetXY(textX, currentY)
				pdf.Cell(textWidth, 2.5, "Titular: "+*req.Doctor.CertificateName)
				currentY += 2.5
			}

			// Conselho profissional (CRM-UF: número)
			if req.Doctor.CRM != nil && req.Doctor.CRMUF != nil {
				pdf.SetXY(textX, currentY)
				crmLine := "CRM-" + *req.Doctor.CRMUF + ": " + *req.Doctor.CRM
				pdf.Cell(textWidth, 2.5, crmLine)
				currentY += 2.5
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
				pdf.SetXY(textX, currentY)
				pdf.Cell(textWidth, 2.5, certInfoLine)
				currentY += 2.5
			}
		}

		currentY += 1

		// Links de validação - mais compactos
		pdf.SetFont("OpenSans", "", 7)
		pdf.SetXY(textX, currentY)
		pdf.Cell(textWidth, 2.5, "Validar documento:")
		currentY += 2.5

		pdf.SetXY(textX, currentY)
		fullValidationURL := "https://plenya.com.br/lab-requests/validate/" + req.ID.String()
		pdf.Cell(textWidth, 2.5, fullValidationURL)
		currentY += 3

		pdf.SetXY(textX, currentY)
		pdf.Cell(textWidth, 2.5, "Verificar assinatura: https://validar.iti.gov.br")
		currentY += 3

		// Nota sobre validade legal - compacta e controlada (2 linhas fixas)
		pdf.SetFont("OpenSans", "", 6.5)
		pdf.SetTextColor(60, 60, 60)
		pdf.SetXY(textX, currentY)
		pdf.Cell(textWidth, 2.5, "Documento assinado digitalmente (ICP-Brasil) - CFM 2.299/2021 e RDC ANVISA 1.000/2025.")
		currentY += 2.5
		pdf.SetXY(textX, currentY)
		pdf.Cell(textWidth, 2.5, "Validade juridica equivalente ao documento fisico.")
		pdf.SetTextColor(0, 0, 0)

		// Limpeza do arquivo temporário do QR code
		if qrErr == nil {
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

	// Desabilitar quebra automática de página para controlar manualmente
	pdf.SetAutoPageBreak(false, 0)

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

	// PACIENTE (uma linha com bold nos rótulos)
	xPos := 25.0
	pdf.SetXY(xPos, currentY)

	// "Paciente:" em bold 11pt (mesmo tamanho de "Exames Solicitados")
	pdf.SetFont("OpenSans", "B", 11)
	text := "Paciente: "
	pdf.Cell(pdf.GetStringWidth(text), 5, text)
	xPos += pdf.GetStringWidth(text) + 3 // +3mm de espaço extra

	// Nome em bold 13pt (11 + 2)
	pdf.SetXY(xPos, currentY)
	pdf.SetFont("OpenSans", "B", 13)
	pdf.Cell(pdf.GetStringWidth(patientName), 5, patientName)
	xPos += pdf.GetStringWidth(patientName)

	if patientCPF != "" {
		// Separador
		pdf.SetXY(xPos, currentY)
		pdf.SetFont("OpenSans", "", 9)
		text = "  |  "
		pdf.Cell(pdf.GetStringWidth(text), 5, text)
		xPos += pdf.GetStringWidth(text)

		// "CPF:" em bold
		pdf.SetXY(xPos, currentY)
		pdf.SetFont("OpenSans", "B", 9)
		text = "CPF: "
		pdf.Cell(pdf.GetStringWidth(text), 5, text)
		xPos += pdf.GetStringWidth(text)

		// Valor do CPF normal
		pdf.SetXY(xPos, currentY)
		pdf.SetFont("OpenSans", "", 9)
		pdf.Cell(pdf.GetStringWidth(patientCPF), 5, patientCPF)
		xPos += pdf.GetStringWidth(patientCPF)
	}

	if patientBirthDate != "" {
		// Separador
		pdf.SetXY(xPos, currentY)
		pdf.SetFont("OpenSans", "", 9)
		text = "  |  "
		pdf.Cell(pdf.GetStringWidth(text), 5, text)
		xPos += pdf.GetStringWidth(text)

		// "Nasc.:" em bold
		pdf.SetXY(xPos, currentY)
		pdf.SetFont("OpenSans", "B", 9)
		text = "Nasc.: "
		pdf.Cell(pdf.GetStringWidth(text), 5, text)
		xPos += pdf.GetStringWidth(text)

		// Valor da data normal
		pdf.SetXY(xPos, currentY)
		pdf.SetFont("OpenSans", "", 9)
		pdf.Cell(pdf.GetStringWidth(patientBirthDate), 5, patientBirthDate)
	}

	currentY += 7

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

	// MÉDICO SOLICITANTE (uma linha com bold nos rótulos)
	if req.Doctor != nil {
		xPos := 25.0
		pdf.SetXY(xPos, currentY)

		// "Médico:" em bold 11pt (mesmo tamanho de "Exames Solicitados")
		pdf.SetFont("OpenSans", "B", 11)
		text := "Médico: "
		pdf.Cell(pdf.GetStringWidth(text), 5, text)
		xPos += pdf.GetStringWidth(text) + 3 // +3mm de espaço extra

		// Nome do médico em bold 12pt (11 + 1)
		if req.Doctor.Name != "" {
			doctorName := req.Doctor.Name
			if !strings.HasPrefix(doctorName, "Dr.") && !strings.HasPrefix(doctorName, "Dra.") {
				doctorName = "Dr. " + doctorName
			}

			pdf.SetXY(xPos, currentY)
			pdf.SetFont("OpenSans", "B", 12)
			pdf.Cell(pdf.GetStringWidth(doctorName), 5, doctorName)
			xPos += pdf.GetStringWidth(doctorName)
		}

		// CRM-UF
		if req.Doctor.CRM != nil && req.Doctor.CRMUF != nil {
			// Separador
			pdf.SetXY(xPos, currentY)
			pdf.SetFont("OpenSans", "", 9)
			text = "  |  "
			pdf.Cell(pdf.GetStringWidth(text), 5, text)
			xPos += pdf.GetStringWidth(text)

			// "CRM-UF:" em bold
			pdf.SetXY(xPos, currentY)
			pdf.SetFont("OpenSans", "B", 9)
			text = "CRM-" + *req.Doctor.CRMUF + ": "
			pdf.Cell(pdf.GetStringWidth(text), 5, text)
			xPos += pdf.GetStringWidth(text)

			// Número do CRM normal
			pdf.SetXY(xPos, currentY)
			pdf.SetFont("OpenSans", "", 9)
			pdf.Cell(pdf.GetStringWidth(*req.Doctor.CRM), 5, *req.Doctor.CRM)
		}

		currentY += 6
		currentY += 4
	}

	// Section: Exames Solicitados - bold font
	pdf.SetFont("OpenSans", "B", 11)
	pdf.SetXY(25, currentY)
	pdf.Cell(160, 6, "Exames Solicitados")
	currentY += 8

	// Exams list - adaptive layout based on ALTURA (não número de itens)
	pdf.SetFont("OpenSans", "", 10)

	// Preparar variável para assinatura (usada no loop e no final)
	hasDigitalSignature := req.SignedAt != nil

	// Altura máxima para container de exames (assinatura em Y=230mm)
	maxExamContainerY := 215.0 // Container termina em 215mm (deixa 15mm de margem segura para assinatura)

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
				addSignatureSection(pdf, req, 230.0, formattedDate, hasDigitalSignature)

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

		// Usar MultiCell para quebrar linha automaticamente em nomes longos
		// Salvar posição Y antes do MultiCell
		startY := columnY
		pdf.MultiCell(colWidth-5, 5, exam, "", "L", false)

		// Calcular quantas linhas o MultiCell ocupou
		endY := pdf.GetY()
		linesUsed := (endY - startY) / 5

		// Ajustar columnY baseado nas linhas usadas
		columnY = startY + (linesUsed * 6) // 6mm por linha (5mm texto + 1mm espaço)
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
	// Assinatura subida 5mm (0,5cm) - agora em Y=230mm
	// Assinatura tem ~32mm de altura, termina em ~262mm (seguro, antes do rodapé em 277mm)
	if currentY < 230.0 {
		currentY = 230.0
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
