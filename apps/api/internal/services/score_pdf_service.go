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
	pdf, err := s.generatePDFFromHTML(html)
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

// generatePDFFromHTML converts HTML to PDF using Rod (Headless Chrome)
func (s *ScorePDFService) generatePDFFromHTML(html string) ([]byte, error) {
	fmt.Println("[PDF] Launching browser...")
	// Launch browser
	path, _ := launcher.LookPath()
	fmt.Printf("[PDF] Browser path: %s\n", path)
	u := launcher.New().Bin(path).MustLaunch()
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

	// Configure PDF options (60cm x 300cm = 600mm x 3000mm)
	paperWidth := 23.62  // 600mm in inches (600 / 25.4)
	paperHeight := 118.11 // 3000mm in inches (3000 / 25.4)
	marginZero := 0.0
	scale := 1.0

	pdfOptions := &proto.PagePrintToPDF{
		PrintBackground:     true,
		PreferCSSPageSize:   false,
		PaperWidth:          &paperWidth,
		PaperHeight:         &paperHeight,
		MarginTop:           &marginZero,
		MarginBottom:        &marginZero,
		MarginLeft:          &marginZero,
		MarginRight:         &marginZero,
		Scale:               &scale,
		DisplayHeaderFooter: false,
	}

	// Generate PDF stream
	pdfStream, err := page.PDF(pdfOptions)
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
