package services

import (
	"testing"

	"github.com/plenya/api/internal/models"
)

func intPtr(v int) *int { return &v }

// IGF-1 no catálogo tem itens por faixa etária MAIS um item guarda-chuva sem faixa nenhuma.
// Pegar o primeiro da lista deixava o resultado sem nível quando o guarda-chuva vinha antes.
func TestPickScoringItem_IgnoraGuardaChuvaSemFaixas(t *testing.T) {
	umbrella := models.ScoreItem{Name: "IGF-1 (Somatomedina C)"}
	faixa := models.ScoreItem{
		Name:        "IGF-1 (51-70 anos)",
		AgeRangeMin: intPtr(51),
		AgeRangeMax: intPtr(70),
		Levels:      []models.ScoreLevel{{Level: 3, Name: "Baixo"}},
	}

	got := pickScoringItem([]models.ScoreItem{umbrella, faixa})
	if got == nil || got.Name != faixa.Name {
		t.Fatalf("esperava o item com faixas, veio %v", got)
	}
}

// Entre dois itens com faixas, o recorte etário mais estreito é o mais específico.
func TestPickScoringItem_PrefereMaisEspecifico(t *testing.T) {
	amplo := models.ScoreItem{
		Name:   "Ferritina (geral)",
		Levels: []models.ScoreLevel{{Level: 5, Name: "Ótimo"}},
	}
	estreito := models.ScoreItem{
		Name:        "Ferritina (31-50 anos)",
		AgeRangeMin: intPtr(31),
		AgeRangeMax: intPtr(50),
		Levels:      []models.ScoreLevel{{Level: 5, Name: "Ótimo"}},
	}

	if got := pickScoringItem([]models.ScoreItem{amplo, estreito}); got == nil || got.Name != estreito.Name {
		t.Fatalf("esperava o recorte etário mais estreito, veio %v", got)
	}
}

func TestPickScoringItem_TodosSemFaixas(t *testing.T) {
	if got := pickScoringItem([]models.ScoreItem{{Name: "Sem faixas"}}); got != nil {
		t.Fatalf("item sem faixas não pode ser escolhido para classificar")
	}
}

// Sorologia é qualitativa: os níveis são "Reagente" (0) e "Não-reagente" (5), sem limite
// numérico. Antes, qualquer resultado em texto era descartado sem sequer olhar os níveis.
func TestMatchQualitativeLevel(t *testing.T) {
	levels := []models.ScoreLevel{
		{Level: 0, Name: "Reagente", Operator: "="},
		{Level: 5, Name: "Não-reagente", Operator: "="},
	}

	cases := map[string]*int{
		"Não reagente": intPtr(5),
		"nao-reagente": intPtr(5),
		"NÃO-REAGENTE": intPtr(5),
		"Negativo":     intPtr(5),
		"não detectável": intPtr(5),
		"Reagente":       intPtr(0),
		"positivo":       intPtr(0),
		"Detectável":     intPtr(0),
		"inconclusivo":   nil,
		"":               nil,
	}

	for text, want := range cases {
		got := matchQualitativeLevel(levels, text)
		switch {
		case want == nil && got != nil:
			t.Fatalf("%q: esperava nenhum nível, veio %d", text, *got)
		case want != nil && got == nil:
			t.Fatalf("%q: esperava nível %d, não classificou", text, *want)
		case want != nil && *got != *want:
			t.Fatalf("%q: esperava nível %d, veio %d", text, *want, *got)
		}
	}
}
