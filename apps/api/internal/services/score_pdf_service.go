package services

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"os"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/plenya/api/internal/models"
)

// ScorePDFService handles PDF generation for scores
type ScorePDFService struct{}

// NewScorePDFService creates a new score PDF service
func NewScorePDFService() *ScorePDFService {
	return &ScorePDFService{}
}

// GeneratePDFFromHTML converte HTML em PDF A4 (Chromium headless via go-rod). Wrapper exportado
// reusado pelo relatório longitudinal AGIR (P3 frente 3b). Usa página A4 (NÃO o pôster 60×200cm).
func (s *ScorePDFService) GeneratePDFFromHTML(html string) ([]byte, error) {
	return s.generatePDFFromHTML(html, a4PDFOptions())
}

// posterPDFOptions — papel do pôster (60cm × 200cm = 600×2000mm).
func posterPDFOptions() *proto.PagePrintToPDF {
	w := 23.62 // 600mm em polegadas
	h := 78.74 // 2000mm em polegadas
	zero := 0.0
	scale := 1.0
	return &proto.PagePrintToPDF{
		PrintBackground:     true,
		PreferCSSPageSize:   false,
		PaperWidth:          &w,
		PaperHeight:         &h,
		MarginTop:           &zero,
		MarginBottom:        &zero,
		MarginLeft:          &zero,
		MarginRight:         &zero,
		Scale:               &scale,
		DisplayHeaderFooter: false,
	}
}

// a4PDFOptions — papel A4 com margens por página (documentos clínicos paginam em múltiplas folhas).
func a4PDFOptions() *proto.PagePrintToPDF {
	w := 8.27  // 210mm
	h := 11.69 // 297mm
	margin := 0.5
	scale := 1.0
	return &proto.PagePrintToPDF{
		PrintBackground:     true,
		PreferCSSPageSize:   false,
		PaperWidth:          &w,
		PaperHeight:         &h,
		MarginTop:           &margin,
		MarginBottom:        &margin,
		MarginLeft:          &margin,
		MarginRight:         &margin,
		Scale:               &scale,
		DisplayHeaderFooter: false,
	}
}

// GeneratePosterPDF generates a PDF poster (60cm x 300cm) for all scores
func (s *ScorePDFService) GeneratePosterPDF(groups []models.ScoreGroup) ([]byte, error) {
	fmt.Printf("[PDF] Starting PDF generation for %d groups\n", len(groups))

	// Render HTML template
	html, err := s.renderPosterHTML(groups)
	if err != nil {
		fmt.Printf("[PDF] Failed to render HTML: %v\n", err)
		return nil, fmt.Errorf("failed to render HTML: %v", err)
	}
	fmt.Printf("[PDF] HTML rendered successfully (%d bytes)\n", len(html))

	// Generate PDF using Rod (Headless Chrome)
	pdf, err := s.generatePDFFromHTML(html, posterPDFOptions())
	if err != nil {
		fmt.Printf("[PDF] Failed to generate PDF: %v\n", err)
		return nil, fmt.Errorf("failed to generate PDF: %v", err)
	}
	fmt.Printf("[PDF] PDF generated successfully (%d bytes)\n", len(pdf))

	return pdf, nil
}

// renderPosterHTML renders the poster HTML template with score data
func (s *ScorePDFService) renderPosterHTML(groups []models.ScoreGroup) (string, error) {
	// Read template file
	templatePath := "/app/internal/templates/score_poster.html"
	fmt.Printf("[PDF] Reading template from: %s\n", templatePath)
	tmplContent, err := os.ReadFile(templatePath)
	if err != nil {
		fmt.Printf("[PDF] Failed to read template: %v\n", err)
		return "", fmt.Errorf("failed to read template: %v", err)
	}
	fmt.Printf("[PDF] Template read successfully (%d bytes)\n", len(tmplContent))

	// Parse template with custom functions
	funcMap := template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"gtInt": func(a, b int) bool {
			return a > b
		},
		"dict": func(values ...interface{}) (map[string]interface{}, error) {
			if len(values)%2 != 0 {
				return nil, fmt.Errorf("dict requires an even number of arguments")
			}
			dict := make(map[string]interface{}, len(values)/2)
			for i := 0; i < len(values); i += 2 {
				key, ok := values[i].(string)
				if !ok {
					return nil, fmt.Errorf("dict keys must be strings")
				}
				dict[key] = values[i+1]
			}
			return dict, nil
		},
		"formatLevelRange": func(level models.ScoreLevel) string {
			if level.Operator == "between" && level.LowerLimit != nil && level.UpperLimit != nil {
				return fmt.Sprintf("%s - %s", *level.LowerLimit, *level.UpperLimit)
			}
			if level.Operator == "=" {
				if level.LowerLimit != nil {
					return fmt.Sprintf("= %s", *level.LowerLimit)
				}
				if level.UpperLimit != nil {
					return fmt.Sprintf("= %s", *level.UpperLimit)
				}
			}
			if level.Operator == ">" {
				if level.LowerLimit != nil {
					return fmt.Sprintf("> %s", *level.LowerLimit)
				}
				if level.UpperLimit != nil {
					return fmt.Sprintf("> %s", *level.UpperLimit)
				}
			}
			if level.Operator == ">=" {
				if level.LowerLimit != nil {
					return fmt.Sprintf("≥ %s", *level.LowerLimit)
				}
				if level.UpperLimit != nil {
					return fmt.Sprintf("≥ %s", *level.UpperLimit)
				}
			}
			if level.Operator == "<" {
				if level.UpperLimit != nil {
					return fmt.Sprintf("< %s", *level.UpperLimit)
				}
				if level.LowerLimit != nil {
					return fmt.Sprintf("< %s", *level.LowerLimit)
				}
			}
			if level.Operator == "<=" {
				if level.UpperLimit != nil {
					return fmt.Sprintf("≤ %s", *level.UpperLimit)
				}
				if level.LowerLimit != nil {
					return fmt.Sprintf("≤ %s", *level.LowerLimit)
				}
			}
			return ""
		},
		"gt": func(a, b float64) bool {
			return a > b
		},
	}

	tmpl, err := template.New("poster").Funcs(funcMap).Parse(string(tmplContent))
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %v", err)
	}

	// Prepare data
	data := map[string]interface{}{
		"Groups":      groups,
		"GeneratedAt": time.Now().Format("02/01/2006 às 15:04"),
		"Year":        time.Now().Year(),
	}

	// Execute template
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %v", err)
	}

	return buf.String(), nil
}

// generatePDFFromHTML converts HTML to PDF using Rod (Headless Chrome).
// opts define o tamanho de papel/margens (pôster vs A4).
func (s *ScorePDFService) generatePDFFromHTML(html string, opts *proto.PagePrintToPDF) ([]byte, error) {
	fmt.Println("[PDF] Launching browser...")

	// Use Chromium from Alpine Linux if available, otherwise auto-detect
	chromiumPath := "/usr/bin/chromium-browser"
	if _, err := os.Stat(chromiumPath); os.IsNotExist(err) {
		// Fallback to auto-detection
		chromiumPath, _ = launcher.LookPath()
	}
	fmt.Printf("[PDF] Browser path: %s\n", chromiumPath)

	// Launch browser with explicit path and headless mode
	u := launcher.New().
		Bin(chromiumPath).
		Headless(true).
		NoSandbox(true). // Required for Docker/Alpine
		MustLaunch()
	fmt.Printf("[PDF] Browser launched at: %s\n", u)
	browser := rod.New().ControlURL(u).MustConnect()
	defer browser.MustClose()
	fmt.Println("[PDF] Browser connected")

	// Create page
	fmt.Println("[PDF] Creating page...")
	page := browser.MustPage()
	defer page.MustClose()
	fmt.Println("[PDF] Page created")

	// Save HTML to temp file (data URL too large for 1.5MB HTML)
	tmpFile, err := os.CreateTemp("", "poster-*.html")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(html); err != nil {
		return nil, fmt.Errorf("failed to write temp file: %v", err)
	}
	tmpFile.Close()

	// Navigate to file
	fileURL := "file://" + tmpFile.Name()
	fmt.Printf("[PDF] Navigating to: %s\n", fileURL)
	page.MustNavigate(fileURL).MustWaitLoad()
	fmt.Println("[PDF] Page loaded")

	// Generate PDF stream com as opções de papel fornecidas (pôster ou A4)
	pdfStream, err := page.PDF(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to generate PDF: %v", err)
	}

	// Read stream to bytes
	pdfBytes, err := io.ReadAll(pdfStream)
	if err != nil {
		return nil, fmt.Errorf("failed to read PDF stream: %v", err)
	}

	return pdfBytes, nil
}
