package services

import (
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/plenya/api/internal/dto"
)

func diaCob(t *testing.T, s string) time.Time {
	t.Helper()
	d, err := time.Parse("2006-01-02", s)
	if err != nil {
		t.Fatalf("data %q: %v", s, err)
	}
	return d
}

// O painel do catálogo quase nunca tem resultado próprio: o laboratório não reporta "hemograma
// completo", reporta hemoglobina e plaquetas. Medido neste banco, o hemograma tem 0 resultados
// próprios e 336 nos filhos. Resolver só pelo painel diria "nunca feito" e mandaria repetir o
// exame de quem acabou de fazer.
func TestCoberturaResolveOPainelPelosFilhos(t *testing.T) {
	quando, via := resolveCoverage(time.Time{}, false, diaCob(t, "2026-02-09"))
	if via != dto.LabCoverageChildren {
		t.Errorf("origem = %q, quer children", via)
	}
	if quando.IsZero() {
		t.Error("painel com filho feito não pode sair sem data")
	}
}

// O inverso existe: "Gasometria venosa" tem 16 resultados próprios e nenhum filho. Olhar só os
// filhos erraria esse lado do catálogo.
func TestCoberturaUsaOResultadoProprioQuandoNaoHaFilho(t *testing.T) {
	_, via := resolveCoverage(diaCob(t, "2026-02-07"), true, time.Time{})
	if via != dto.LabCoverageOwn {
		t.Errorf("origem = %q, quer own", via)
	}
}

// "Rotina de urina" tem os dois (25 próprios, 93 nos filhos): vale a coleta mais recente, e a
// origem tem que ser quem deu a data.
func TestCoberturaComOsDoisFicaComAColetaMaisRecente(t *testing.T) {
	quando, via := resolveCoverage(diaCob(t, "2025-01-01"), true, diaCob(t, "2026-02-09"))
	if via != dto.LabCoverageChildren || !quando.Equal(diaCob(t, "2026-02-09")) {
		t.Errorf("got %v/%q, quer a data do filho", quando, via)
	}
	quando2, via2 := resolveCoverage(diaCob(t, "2026-05-01"), true, diaCob(t, "2026-02-09"))
	if via2 != dto.LabCoverageOwn || !quando2.Equal(diaCob(t, "2026-05-01")) {
		t.Errorf("got %v/%q, quer a data própria", quando2, via2)
	}
}

func TestCoberturaSemNadaEhNunca(t *testing.T) {
	quando, via := resolveCoverage(time.Time{}, false, time.Time{})
	if via != dto.LabCoverageNever || !quando.IsZero() {
		t.Errorf("got %v/%q, quer never sem data", quando, via)
	}
}

func TestDescendentesPegaNetoENaoTravaEmCiclo(t *testing.T) {
	pai := uuid.Must(uuid.NewV7())
	filho := uuid.Must(uuid.NewV7())
	neto := uuid.Must(uuid.NewV7())
	// Catálogo de três níveis: olhar só um nível resolveria o painel como "nunca feito".
	arvore := map[uuid.UUID][]uuid.UUID{pai: {filho}, filho: {neto}}
	got := descendentes(arvore, pai)
	if len(got) != 2 {
		t.Fatalf("descendentes = %d, quer 2 (filho e neto)", len(got))
	}

	// parent_test_id é campo livre: um ciclo no catálogo não pode travar o serviço.
	a, b := uuid.Must(uuid.NewV7()), uuid.Must(uuid.NewV7())
	ciclo := map[uuid.UUID][]uuid.UUID{a: {b}, b: {a}}
	if n := len(descendentes(ciclo, a)); n != 1 {
		t.Errorf("com ciclo, descendentes = %d, quer 1 sem laço infinito", n)
	}
}
