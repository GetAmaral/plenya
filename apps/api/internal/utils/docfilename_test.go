package utils

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

// O PDF sai do EMR e vai parar no WhatsApp pessoal do médico: o nome do arquivo precisa
// identificar o documento sozinho, sem o sistema em volta.
func TestDocumentFileName(t *testing.T) {
	id := uuid.MustParse("01a016a2-1a88-74e6-a23f-04a75aebfba3")
	date := time.Date(2026, 8, 19, 0, 0, 0, 0, time.UTC)

	got := DocumentFileName("Luiz Gustavo José Carvalho", "PedidoExame", date, id)
	if want := "Luiz-Gustavo-José-Carvalho_PedidoExame_2026-08-19_01a016a2.pdf"; got != want {
		t.Fatalf("DocumentFileName = %q, esperado %q", got, want)
	}

	got = DocumentFileName("Ana Claudia Correa Zuin", "ReceitaManipulado", date, id)
	if want := "Ana-Claudia-Correa-Zuin_ReceitaManipulado_2026-08-19_01a016a2.pdf"; got != want {
		t.Fatalf("receita manipulada = %q, esperado %q", got, want)
	}
}

func TestDocumentFileName_SemPaciente(t *testing.T) {
	id := uuid.MustParse("01a016a2-1a88-74e6-a23f-04a75aebfba3")
	date := time.Date(2026, 1, 2, 15, 4, 5, 0, time.UTC)
	if got := DocumentFileName("   ", "Receita", date, id); got != "Paciente_Receita_2026-01-02_01a016a2.pdf" {
		t.Fatalf("sem nome o arquivo deveria cair no genérico, veio %q", got)
	}
}

func TestCompactName(t *testing.T) {
	cases := map[string]string{
		"maria da silva":     "Maria-Da-Silva",
		"  ana  clara  ":     "Ana-Clara",
		"José-Maria D'Ávila": "José-Maria-D-Ávila",
		"":                   "",
	}
	for in, want := range cases {
		if got := CompactName(in); got != want {
			t.Fatalf("CompactName(%q) = %q, esperado %q", in, got, want)
		}
	}
}

// filename= não aceita acento; o nome acentuado vai no filename*=UTF-8.
func TestASCIIFallback(t *testing.T) {
	in := "Luiz-Gustavo-José-Carvalho_PedidoExame_2026-08-19_01a016a2.pdf"
	if got := ASCIIFallback(in, "pedido-exame.pdf"); got != "Luiz-Gustavo-Jose-Carvalho_PedidoExame_2026-08-19_01a016a2.pdf" {
		t.Fatalf("ASCIIFallback não removeu o acento: %q", got)
	}
	if got := ASCIIFallback("氏名", "receita.pdf"); got != "receita.pdf" {
		t.Fatalf("nome sem ASCII deveria cair no genérico, veio %q", got)
	}
}

// As duas formas saem no mesmo cabeçalho: a ASCII para cliente antigo, a UTF-8 percent-encoded
// para o resto. Espaço e acento no filename= sem aspas quebravam o download no Safari.
func TestContentDisposition(t *testing.T) {
	got := ContentDisposition("Ana-Cláudia_Receita_2026-08-31_01a0592b.pdf", "receita.pdf")
	want := `inline; filename="Ana-Claudia_Receita_2026-08-31_01a0592b.pdf"; ` +
		`filename*=UTF-8''Ana-Cl%C3%A1udia_Receita_2026-08-31_01a0592b.pdf`
	if got != want {
		t.Fatalf("ContentDisposition =\n  %q\nesperado\n  %q", got, want)
	}
}
