package services

import (
	"os"
	"testing"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/database"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/pdfdoc"
)

// TestLabRequestE2E carrega um lab request REAL do banco dev (decripta CPF via AfterFind) e
// renderiza pelo pipeline pdfdoc — valida o mapeamento model→documento com dados reais.
// Rodar com: PLENYA_E2E_LAB_ID=<uuid> go test -run TestLabRequestE2E -v
func TestLabRequestE2E(t *testing.T) {
	id := os.Getenv("PLENYA_E2E_LAB_ID")
	if id == "" {
		t.Skip("defina PLENYA_E2E_LAB_ID para rodar o e2e")
	}
	cfg, err := config.Load()
	if err != nil {
		t.Skipf("config: %v", err)
	}
	if err := database.Connect(cfg); err != nil {
		t.Skipf("db: %v", err)
	}
	var lr models.LabRequest
	if err := database.DB.Preload("Patient").Preload("Doctor").First(&lr, "id = ?", id).Error; err != nil {
		t.Fatalf("load lab request: %v", err)
	}
	in := buildExamRequest(&lr)
	t.Logf("paciente=%q cpf=%q nasc=%q | médico=%q cred=%q | digital=%v",
		in.Patient.Name, in.Patient.CPFMasked, in.Patient.BirthInfo,
		in.Doctor.Name, in.Doctor.Credentials, in.Signature.Digital)
	pdf, err := pdfdoc.RenderExamRequest(in)
	if err != nil {
		t.Fatalf("render: %v", err)
	}
	if len(pdf) < 2000 || string(pdf[:5]) != "%PDF-" {
		t.Fatalf("saída não é PDF (len=%d)", len(pdf))
	}
	_ = os.WriteFile("/tmp/lab-e2e.pdf", pdf, 0o644)
	t.Logf("OK: %d bytes -> /tmp/lab-e2e.pdf", len(pdf))
}
