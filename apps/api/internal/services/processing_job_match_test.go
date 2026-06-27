package services

import (
	"testing"

	"github.com/google/uuid"
)

// TestMatchTestDefinition_Fuzzy valida o fallback fuzzy (Levenshtein) do matcher:
// tolera typo do laudo ("Triglicerídios" com "i" a mais) e respeita as salvaguardas
// (nomes curtos não fazem fuzzy; empate na menor distância não casa).
func TestMatchTestDefinition_Fuzzy(t *testing.T) {
	s := &ProcessingJobService{}
	tg := uuid.New()
	chol := uuid.New()

	// defMap SEM o alt_name "trigliceridios" — só o fuzzy pode resolver.
	defMap := map[string]uuid.UUID{
		"trigliceridos":    tg, // alt_name oficial (sem o 2º "i")
		"colesterol total": chol,
		"creatinina":       uuid.New(),
	}

	// 1) typo do laudo deve casar via fuzzy (dist 1)
	if got := s.matchTestDefinition("Triglicerídios", defMap); got == nil || *got != tg {
		t.Fatalf("esperava casar Triglicerídios->Triglicerídeos via fuzzy, got=%v", got)
	}
	// 2) grafia exata continua casando
	if got := s.matchTestDefinition("trigliceridos", defMap); got == nil || *got != tg {
		t.Fatalf("match exato quebrou, got=%v", got)
	}

	// 3) salvaguarda: nome curto (<8) NÃO faz fuzzy (evita sodio/iodio)
	short := map[string]uuid.UUID{"sodio": uuid.New()}
	if got := s.matchTestDefinition("iodio", short); got != nil {
		t.Fatalf("sodio/iodio (5 chars) NÃO deveria casar por fuzzy, got=%v", got)
	}

	// 4) salvaguarda: empate na menor distância NÃO casa (ambíguo)
	ambi := map[string]uuid.UUID{"glicemiaa": uuid.New(), "glicemiab": uuid.New()}
	if got := s.matchTestDefinition("glicemiax", ambi); got != nil {
		t.Fatalf("empate de distância deveria abortar o match, got=%v", got)
	}
}

func TestLevenshtein(t *testing.T) {
	cases := []struct {
		a, b string
		want int
	}{
		{"trigliceridios", "trigliceridos", 1},
		{"", "abc", 3},
		{"abc", "abc", 0},
		{"kitten", "sitting", 3},
	}
	for _, c := range cases {
		if got := levenshtein(c.a, c.b); got != c.want {
			t.Errorf("levenshtein(%q,%q)=%d want %d", c.a, c.b, got, c.want)
		}
	}
}
