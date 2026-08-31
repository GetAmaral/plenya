package handlers

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/plenya/api/internal/models"
)

// O PDF sai do EMR e vai parar no WhatsApp pessoal do médico: o nome do arquivo precisa
// identificar o documento sozinho, sem o sistema em volta.
func TestLabRequestFileName(t *testing.T) {
	id := uuid.MustParse("01a016a2-1a88-74e6-a23f-04a75aebfba3")
	req := &models.LabRequest{
		ID:      id,
		Date:    time.Date(2026, 8, 19, 0, 0, 0, 0, time.UTC),
		Patient: &models.Patient{Name: "Luiz Gustavo José Carvalho"},
	}

	got := labRequestFileName(req)
	want := "Luiz-Gustavo-José-Carvalho_PedidoExame_2026-08-19_01a016a2.pdf"
	if got != want {
		t.Fatalf("labRequestFileName = %q, esperado %q", got, want)
	}
}

func TestLabRequestFileName_SemPaciente(t *testing.T) {
	req := &models.LabRequest{
		ID:        uuid.MustParse("01a016a2-1a88-74e6-a23f-04a75aebfba3"),
		CreatedAt: time.Date(2026, 1, 2, 15, 4, 5, 0, time.UTC),
	}
	if got := labRequestFileName(req); got != "Paciente_PedidoExame_2026-01-02_01a016a2.pdf" {
		t.Fatalf("sem paciente carregado o nome deveria cair no genérico, veio %q", got)
	}
}
