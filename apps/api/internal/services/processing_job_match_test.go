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
	if got := s.matchTestDefinition("Triglicerídios", nil, defMap, nil); got == nil || *got != tg {
		t.Fatalf("esperava casar Triglicerídios->Triglicerídeos via fuzzy, got=%v", got)
	}
	// 2) grafia exata continua casando
	if got := s.matchTestDefinition("trigliceridos", nil, defMap, nil); got == nil || *got != tg {
		t.Fatalf("match exato quebrou, got=%v", got)
	}

	// 3) salvaguarda: nome curto (<8) NÃO faz fuzzy (evita sodio/iodio)
	short := map[string]uuid.UUID{"sodio": uuid.New()}
	if got := s.matchTestDefinition("iodio", nil, short, nil); got != nil {
		t.Fatalf("sodio/iodio (5 chars) NÃO deveria casar por fuzzy, got=%v", got)
	}

	// 4) salvaguarda: empate na menor distância NÃO casa (ambíguo)
	ambi := map[string]uuid.UUID{"glicemiaa": uuid.New(), "glicemiab": uuid.New()}
	if got := s.matchTestDefinition("glicemiax", nil, ambi, nil); got != nil {
		t.Fatalf("empate de distância deveria abortar o match, got=%v", got)
	}
}

// TestMatchTestDefinition_Specimen valida a desambiguação por espécime: um nome que casa por
// substring DOIS candidatos (sangue e urina) e o espécime do resultado decide.
func TestMatchTestDefinition_Specimen(t *testing.T) {
	s := &ProcessingJobService{}
	sangue := uuid.New()
	urina := uuid.New()
	defMap := map[string]uuid.UUID{
		"proteinas totais":    sangue,
		"proteinas urinarias": urina,
	}
	specByID := map[uuid.UUID]string{sangue: "Sangue", urina: "Urina"}

	// "proteinas" casa por substring AMBOS; espécime decide.
	u := "Urina"
	if got := s.matchTestDefinition("Proteínas", &u, defMap, specByID); got == nil || *got != urina {
		t.Fatalf("Proteínas+Urina deveria casar o item de urina por espécime, got=%v", got)
	}
	sg := "Sangue"
	if got := s.matchTestDefinition("Proteínas", &sg, defMap, specByID); got == nil || *got != sangue {
		t.Fatalf("Proteínas+Sangue deveria casar o item de sangue por espécime, got=%v", got)
	}
	// Sem espécime e ambíguo: mantém o 1º (não quebra; comportamento atual).
	if got := s.matchTestDefinition("Proteínas", nil, defMap, specByID); got == nil {
		t.Fatalf("sem espécime deveria cair no 1º candidato, got nil")
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
