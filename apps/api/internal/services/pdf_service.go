package services

import (
	"fmt"
	"os"
	"strings"

	"github.com/jung-kurt/gofpdf/v2"
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
		if req.Patient.CPF != "" {
			// Format CPF: 000.000.000-00
			cpf := req.Patient.CPF
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

	// Patient name - HIGHLIGHTED with larger font and more prominence
	pdf.SetFont("OpenSans", "B", 16)
	pdf.SetXY(25, currentY)
	pdf.Cell(160, 8, patientName)
	currentY += 11

	// CPF and Birth Date on same line
	pdf.SetFont("OpenSans", "", 10)
	pdf.SetXY(25, currentY)

	infoLine := ""
	if patientCPF != "" {
		infoLine = "CPF: " + patientCPF
	}
	if patientBirthDate != "" {
		if infoLine != "" {
			infoLine += "    "
		}
		infoLine += "Data de Nascimento: " + patientBirthDate
	}

	if infoLine != "" {
		pdf.Cell(160, 6, infoLine)
		currentY += 8
	}

	currentY += 4

	// Section: Exames Solicitados - bold font
	pdf.SetFont("OpenSans", "B", 11)
	pdf.SetXY(25, currentY)
	pdf.Cell(160, 6, "Exames Solicitados")
	currentY += 8

	// Exams list - adaptive layout based on number of items
	pdf.SetFont("OpenSans", "", 10)
	totalExams := len(exams)
	useTwoColumns := totalExams > 20
	itemsPerPage := 20
	if useTwoColumns {
		itemsPerPage = 40 // 20 per column
	}

	// Column dimensions
	leftColX := 25.0
	rightColX := 110.0
	colWidth := 75.0
	if !useTwoColumns {
		colWidth = 153.0 // Full width for single column
	}

	columnY := currentY
	startY := currentY // Save starting Y position for right column
	leftColumnEndY := currentY // Track where left column ends

	for i, exam := range exams {
		// Check if we need a new page
		if i > 0 && i%itemsPerPage == 0 {
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

			// Add section title again on new page - Bold
			pdf.SetFont("OpenSans", "B", 11)
			pdf.SetXY(25, columnY)
			if useTwoColumns {
				pdf.Cell(160, 6, "Exames Solicitados (continuação)")
			} else {
				pdf.Cell(160, 6, "Exames Solicitados (continuação)")
			}
			columnY += 8
			pdf.SetFont("OpenSans", "", 10)
		}

		// Determine X position based on column and item index
		var xPos float64
		if useTwoColumns {
			// Calculate which column this item should be in
			pageItemIndex := i % itemsPerPage // Position within current page
			if pageItemIndex < 20 {
				// First 20 items: left column
				xPos = leftColX
			} else {
				// Next 20 items: right column
				xPos = rightColX
			}
		} else {
			xPos = leftColX
		}

		// Add exam number and text
		pdf.SetXY(xPos, columnY)
		pdf.Cell(5, 5, fmt.Sprintf("%d.", i+1))
		pdf.SetXY(xPos+7, columnY)

		// Use Cell instead of MultiCell for two-column layout to prevent overflow
		if useTwoColumns {
			// Truncate long names if needed
			examText := exam
			if len(examText) > 45 {
				examText = examText[:42] + "..."
			}
			pdf.Cell(colWidth-7, 5, examText)
			// Move to next row (fixed height for single-line)
			columnY += 6
		} else {
			// Full width with MultiCell for single column
			pdf.MultiCell(colWidth-7, 5, exam, "", "L", false)
			// Capture actual Y position after MultiCell (handles multi-line text)
			columnY = pdf.GetY() + 1 // Add 1mm spacing between items
		}

		// Reset to top for right column when switching from item 19 to 20
		if useTwoColumns {
			pageItemIndex := i % itemsPerPage
			if pageItemIndex == 19 {
				// Save left column end position before switching
				leftColumnEndY = columnY
				// Moving to right column, reset to top
				columnY = startY
			}
		}
	}

	// Update currentY to the last position used
	// For two columns, use the lower (higher Y value) of both columns
	if useTwoColumns && leftColumnEndY > columnY {
		currentY = leftColumnEndY
	} else {
		currentY = columnY
	}

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

	// Date above signature (aligned left) - add spacing
	currentY += 18 // Add margin (3 lines)

	// Check if we need a new page for signature area (need ~40mm for date + signature)
	if currentY > 250.0 {
		pdf.AddPage()

		// Draw letterhead on new page
		pdf.ImageOptions(
			"/app/PlenyaA4-150dpi.png",
			0, 0, 210, 297,
			false,
			gofpdf.ImageOptions{ImageType: "PNG"},
			0, "",
		)

		currentY = 235.0 // Start signature area on new page
	}

	// Ensure minimum position for signature area
	if currentY < 235.0 {
		currentY = 235.0
	}

	pdf.SetFont("OpenSans", "", 10)
	pdf.SetXY(60, currentY) // Add generous left indent
	pdf.Cell(160, 5, formattedDate)

	// Signature line (at bottom)
	signatureY := currentY + 10
	pdf.Line(60, signatureY, 150, signatureY)
	pdf.SetFont("OpenSans", "", 9)
	pdf.SetXY(25, signatureY+2)
	pdf.CellFormat(160, 5, "Assinatura do Médico Solicitante", "", 0, "C", false, 0, "")

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
