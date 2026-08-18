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

// Casos reais do laudo que ficavam sem nível: rótulos com barra, plural, parêntese e
// frases inteiras ("Amostra NEGATIVA").
func TestMatchQualitativeLevel_RotulosDeLaudo(t *testing.T) {
	urina := []models.ScoreLevel{
		{Level: 0, Name: "Turvo intenso/Purulento"},
		{Level: 4, Name: "Translúcido"},
		{Level: 5, Name: "Límpido/Cristalino"},
	}
	if l := matchQualitativeLevel(urina, "Límpido"); l == nil || *l != 5 {
		t.Fatalf("aspecto 'Límpido' deveria casar com 'Límpido/Cristalino'")
	}

	sedimento := []models.ScoreLevel{
		{Level: 0, Name: "Abundantes (4+)"},
		{Level: 4, Name: "Raras (1+)"},
		{Level: 5, Name: "Ausentes"},
	}
	if l := matchQualitativeLevel(sedimento, "Ausente"); l == nil || *l != 5 {
		t.Fatalf("singular do laudo deveria casar com o plural do nível")
	}

	glicose := []models.ScoreLevel{
		{Level: 3, Name: "100 a 249 (1+)"},
		{Level: 5, Name: "Negativo (<15)"},
	}
	if l := matchQualitativeLevel(glicose, "Negativo"); l == nil || *l != 5 {
		t.Fatalf("parêntese no nome do nível não pode atrapalhar o match")
	}

	cultura := []models.ScoreLevel{
		{Level: 0, Name: ">=10^7"},
		{Level: 5, Name: "Negativa"},
	}
	if l := matchQualitativeLevel(cultura, "Amostra NEGATIVA"); l == nil || *l != 5 {
		t.Fatalf("frase do laudo deveria casar com o nível pelo termo que importa")
	}

	sorologia := []models.ScoreLevel{
		{Level: 0, Name: "Reagente"},
		{Level: 5, Name: "Não-reagente"},
	}
	if l := matchQualitativeLevel(sorologia, "Amostra não reagente para HIV"); l == nil || *l != 5 {
		t.Fatalf("'não reagente' dentro da frase deveria dar nível 5")
	}
	// A armadilha: "reagente" é substring de "não reagente". Não pode virar nível 0.
	if l := matchQualitativeLevel(sorologia, "nao reagente"); l == nil || *l != 5 {
		t.Fatalf("'nao reagente' não pode ser lido como 'reagente'")
	}
}

func TestNumericFromComparativeText(t *testing.T) {
	cases := map[string]float64{
		"Superior a 1.000,0":        1000,
		"superior a 1.000,0 mUI/mL": 1000,
		"< 5":                       5,
		"maior que 100":             100,
		"inferior a 0,90":           0.9,
	}
	for text, want := range cases {
		got, ok := numericFromComparativeText(text)
		if !ok || got != want {
			t.Fatalf("numericFromComparativeText(%q) = %v/%v, esperado %v", text, got, ok, want)
		}
	}
	if _, ok := numericFromComparativeText("Não reagente"); ok {
		t.Fatalf("texto sem número não pode virar número")
	}
}

func TestIsPendingLabText(t *testing.T) {
	for _, text := range []string{"Em Andamento", "AGUARDANDO liberação", "Material insuficiente"} {
		if !isPendingLabText(text) {
			t.Fatalf("%q deveria ser reconhecido como resultado ainda não liberado", text)
		}
	}
	if isPendingLabText("Não reagente") {
		t.Fatalf("resultado liberado não pode ser marcado como pendente do laboratório")
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
		"Não reagente":   intPtr(5),
		"nao-reagente":   intPtr(5),
		"NÃO-REAGENTE":   intPtr(5),
		"Negativo":       intPtr(5),
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

// "1.000,0" (milhar + decimal brasileiro) virava "1.000.0" e não parseava — o valor caía
// como texto e ficava sem nível.
func TestParseNumericResult_FormatoBrasileiro(t *testing.T) {
	cases := map[string]float64{
		"1.000,0":  1000,
		"2.548,42": 2548.42,
		"7,429":    7.429,
		"7.429":    7.429,
		"98":       98,
		"1,6":      1.6,
		"-2,0":     -2,
	}
	for text, want := range cases {
		got, err := parseNumericResult(text)
		if err != nil || got != want {
			t.Fatalf("parseNumericResult(%q) = %v (err=%v), esperado %v", text, got, err, want)
		}
	}
}
