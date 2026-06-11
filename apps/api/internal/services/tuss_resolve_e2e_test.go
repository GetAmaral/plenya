package services

import (
	"os"
	"testing"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/database"
	"github.com/plenya/api/internal/pdfdoc"
)

// TestTussResolveE2E: nomes reais -> TUSS resolvido pelo catálogo + render. Rodar com PLENYA_E2E=1.
func TestTussResolveE2E(t *testing.T) {
	if os.Getenv("PLENYA_E2E") == "" {
		t.Skip("set PLENYA_E2E=1")
	}
	cfg, err := config.Load()
	if err != nil { t.Skipf("config: %v", err) }
	if err := database.Connect(cfg); err != nil { t.Skipf("db: %v", err) }

	exams := "Hemograma completo\nCreatinina\nTSH\nFerritina\nVitamina D\nPSA Total\nUm exame inventado sem match"
	pages := resolveExamPages(exams)
	for _, pg := range pages {
		for _, it := range pg {
			t.Logf("  %-45s TUSS=%q", it.Name, it.Tuss)
		}
	}
	pdf, err := pdfdoc.RenderExamRequest(pdfdoc.ExamRequest{
		Patient:   pdfdoc.Patient{Name: "Teste TUSS"},
		ExamPages: pages,
		Doctor:    pdfdoc.Doctor{Name: "Dr. Getúlio", Credentials: "CRM-PR 21.876"},
		Signature: pdfdoc.Signature{Digital: false, PlaceDate: "Londrina, 11 de junho de 2026"},
	})
	if err != nil { t.Fatalf("render: %v", err) }
	_ = os.WriteFile("/tmp/tuss.pdf", pdf, 0o644)
}
