package services

import (
	html2 "html"
	"strings"
	"testing"

	"github.com/google/uuid"

	"github.com/plenya/api/internal/models"
)

// testScoreGroups monta um grupo sintético (sem DB) para exercitar o render do pôster.
func testScoreGroups() []models.ScoreGroup {
	return []models.ScoreGroup{{
		ID:    uuid.Must(uuid.NewV7()),
		Name:  "Hemograma Completo",
		Order: 1,
	}}
}

// O pôster já falhou em prod ("open /app/internal/templates/score_poster.html: no such file")
// porque lia template e logo do disco: existiam só no dev, via bind-mount de ./apps/api em /app.
// A imagem de prod copia apenas o binário. Estes testes prendem os dois assets no go:embed —
// rodam a partir de qualquer cwd, exatamente como o binário de prod.
func TestRenderPosterHTMLUsesEmbeddedTemplate(t *testing.T) {
	s := NewScorePDFService()

	html, err := s.renderPosterHTML(testScoreGroups())
	if err != nil {
		t.Fatalf("renderPosterHTML: %v", err)
	}
	if !strings.Contains(html, "Hemograma Completo") {
		t.Error("template renderizou sem os dados do grupo")
	}
	if strings.Contains(html, "/app/internal/templates/") {
		t.Error("HTML ainda referencia caminho em disco; em prod esse caminho não existe")
	}
}

// O HTML é renderizado a partir de um arquivo temporário (file://), então o logo precisa ir
// embutido como data-URI. E html/template troca data: URI cru em src por "#ZgotmplZ" —
// daí o template.URL em logoDataURI(); sem ele o pôster sai sem logo e nada falha em voz alta.
func TestRenderPosterHTMLInlinesLogo(t *testing.T) {
	s := NewScorePDFService()

	html, err := s.renderPosterHTML(testScoreGroups())
	if err != nil {
		t.Fatalf("renderPosterHTML: %v", err)
	}
	// Unescape antes de comparar: no atributo, html/template escapa o "+" de svg+xml como
	// &#43; (o parser do browser desfaz isso). A asserção é sobre o valor que o Chromium vê.
	if !strings.Contains(html2.UnescapeString(html), "src=\"data:image/svg+xml;base64,") {
		t.Error("logo não foi embutido como data-URI")
	}
	if strings.Contains(html, "ZgotmplZ") {
		t.Error("html/template sanitizou o data-URI do logo (falta template.URL)")
	}
}
