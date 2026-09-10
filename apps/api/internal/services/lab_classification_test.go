package services

import (
	"strings"
	"testing"

	"github.com/plenya/api/internal/models"
)

func intPtr(v int) *int { return &v }
func strPtr(v string) *string { return &v }

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
// A lipoproteína(a) tem DUAS escalas no catálogo, em nmol/L e em mg/dL, porque não existe fator de
// conversão válido entre elas. Nenhuma tem recorte de idade, então o desempate por faixa etária não
// separa as duas e a escolha caía na ordem que o banco devolveu. Quando vinha a de mg/dL contra um
// laudo em nmol/L, a guarda de unidade recusava classificar — e a Lp(a) de 12 nmol/L ficava sem
// nível com a escala certa ali do lado.
//
// O filtro já existia em patient_plan_dossier_service.go, que monta a régua, e não no
// classificador: régua certa, nível ausente, sobre o mesmo exame.
func TestFiltraPelaUnidade_EscolheAEscalaDaGrandezaDoLaudo(t *testing.T) {
	semSinonimos := func(*string) [][2]string { return nil }
	nmol := models.ScoreItem{
		Name:        "Lipoproteína A",
		Unit:        strPtr("nmol/L"),
		LabTestCode: strPtr("PLNA31F0501"),
		Levels:      []models.ScoreLevel{{Level: 5, Name: "≤30", Operator: "<=", UpperLimit: strPtr("30")}},
	}
	mgdl := models.ScoreItem{
		Name:        "Lipoproteína A (mg/dL)",
		Unit:        strPtr("mg/dL"),
		LabTestCode: strPtr("PLNA31F0501"),
		Levels:      []models.ScoreLevel{{Level: 5, Name: "≤14", Operator: "<=", UpperLimit: strPtr("14")}},
	}

	// A de mg/dL vem PRIMEIRO de propósito: é a ordem que reproduzia o defeito.
	got := pickScoringItem(filtraPelaUnidade([]models.ScoreItem{mgdl, nmol}, "nmol/L", semSinonimos))
	if got == nil || got.Name != nmol.Name {
		t.Fatalf("laudo em nmol/L devia escolher a escala em nmol/L, veio %v", got)
	}
	if !got.UnitMatches("nmol/L", nil) {
		t.Fatal("a escala escolhida tem que passar na guarda de unidade")
	}

	got = pickScoringItem(filtraPelaUnidade([]models.ScoreItem{nmol, mgdl}, "mg/dL", semSinonimos))
	if got == nil || got.Name != mgdl.Name {
		t.Fatalf("laudo em mg/dL devia escolher a escala em mg/dL, veio %v", got)
	}
}

// Nenhuma escala casando, devolve a lista inteira: a guarda adiante recusa e grava o MOTIVO, que é
// melhor do que classificar contra a grandeza errada em silêncio. É o caso do sedimento urinário,
// com escala em células/campo e laudo em /µL.
func TestFiltraPelaUnidade_SemCasarDevolveTudo(t *testing.T) {
	semSinonimos := func(*string) [][2]string { return nil }
	campo := models.ScoreItem{
		Name:        "Hemácias (RBC) - Sedimento",
		Unit:        strPtr("células/campo"),
		LabTestCode: strPtr("PLN6B5C27A4"),
		Levels:      []models.ScoreLevel{{Level: 5, Name: "≤5", Operator: "<=", UpperLimit: strPtr("5")}},
	}

	got := filtraPelaUnidade([]models.ScoreItem{campo}, "/µL", semSinonimos)
	if len(got) != 1 {
		t.Fatalf("esperava a lista original, veio %d itens", len(got))
	}
	if got[0].UnitMatches("/µL", nil) {
		t.Fatal("células/campo e /µL não são a mesma grandeza: a guarda tem que recusar")
	}
}

// O mapa de sinônimos traduzia só o TEXTO DO LAUDO. Quando é o NOME DO NÍVEL que está na outra
// forma do mesmo vocabulário, nada casava: FAN "Não reagente" contra nível "Negativo", proteinúria
// "Negativa" contra nível "Negativo (<10)". Os dois lados dizem a mesma coisa e o resultado saía do
// escore em silêncio — a proteinúria inclusive, que é o dado que decide se há doença glomerular.
func TestMatchQualitativeLevel_CanonicalizaOsDoisLados(t *testing.T) {
	fan := []models.ScoreLevel{
		{Level: 5, Name: "Negativo"},
		{Level: 3, Name: "1:80"},
		{Level: 0, Name: "≥1:640"},
	}
	if l := matchQualitativeLevel(fan, "Não reagente (AC-0), título não reagente"); l == nil || *l != 5 {
		t.Fatalf("FAN não reagente devia cair no nível 5 (Negativo), veio %v", l)
	}

	proteinas := []models.ScoreLevel{
		{Level: 5, Name: "Negativo (<10)"},
		{Level: 3, Name: "10 a 29 (Traços)"},
		{Level: 0, Name: "≥300 (3+ a 4+)"},
	}
	if l := matchQualitativeLevel(proteinas, "Negativa"); l == nil || *l != 5 {
		t.Fatalf("proteinúria negativa devia cair no nível 5, veio %v", l)
	}

	// O caminho antigo continua valendo: texto na forma canônica, nível na forma coloquial.
	sorologia := []models.ScoreLevel{{Level: 5, Name: "Não-reagente"}, {Level: 0, Name: "Reagente"}}
	if l := matchQualitativeLevel(sorologia, "Negativo"); l == nil || *l != 5 {
		t.Fatalf("texto negativo contra nível não-reagente devia continuar casando, veio %v", l)
	}
}

// Texto com dois termos de canonicalização OPOSTA não pode depender da ordem de iteração do mapa:
// "Não detectado" (→ não reagente) e "Reagente" (→ reagente) na mesma frase davam nível 5 ou nível 0
// conforme a execução. A frase mais longa e mais específica ganha, sempre.
func TestCanonicalizaQualitativo_Determinista(t *testing.T) {
	// "detectado" (-> reagente) e "ausente" (-> nao reagente) na mesma frase: com varredura do mapa,
	// o vencedor mudava de execução para execução.
	ambiguo := strings.Fields(normalizeQualitative("Anticorpo detectado antigeno ausente"))
	primeiro := canonicalizaQualitativo(ambiguo)
	if primeiro == "" {
		t.Fatal("esperava alguma forma canônica para um texto com dois termos do vocabulário")
	}
	for i := 0; i < 300; i++ {
		if got := canonicalizaQualitativo(ambiguo); got != primeiro {
			t.Fatalf("iteração %d devolveu %q, antes era %q: resultado dependente da ordem do mapa", i, got, primeiro)
		}
	}

	// Especificidade: a frase mais longa ganha da palavra contida nela.
	if got := canonicalizaQualitativo(strings.Fields(normalizeQualitative("Não detectado"))); got != "nao reagente" {
		t.Fatalf("\"não detectado\" tem que ganhar de \"detectado\", veio %q", got)
	}

	// E a lista tem que estar mesmo do mais longo para o mais curto.
	for i := 1; i < len(sinonimosQualitativosOrdenados); i++ {
		if len(sinonimosQualitativosOrdenados[i-1]) < len(sinonimosQualitativosOrdenados[i]) {
			t.Fatalf("ordem quebrada em %d: %q antes de %q",
				i, sinonimosQualitativosOrdenados[i-1], sinonimosQualitativosOrdenados[i])
		}
	}
}

// Laudo que escreve o qualitativo como frase, com ponto final, é comum — e "Negativo." não casava
// com o nível "Negativo" porque o token ficava "negativo.". O resultado saía do escore em silêncio.
func TestMatchQualitativeLevel_PontuacaoDeFimDeFrase(t *testing.T) {
	niveis := []models.ScoreLevel{{Level: 5, Name: "Não-reagente"}, {Level: 0, Name: "Reagente"}}
	for _, txt := range []string{"Negativo", "Negativo.", "Não reagente.", "Negativo;", "Negativo!"} {
		l := matchQualitativeLevel(niveis, txt)
		if l == nil || *l != 5 {
			t.Errorf("%q devia cair no nível 5, veio %v", txt, l)
		}
	}

	// E o ponto que separa MILHAR tem que sobreviver: tratá-lo como pontuação transformava
	// "Superior a 1.000,0" em "1 000,0" e o valor lido virava 1.
	if v, ok := numericFromComparativeText("Superior a 1.000,0"); !ok || v != 1000 {
		t.Errorf("separador de milhar corrompido: %v %v", v, ok)
	}
}

// Nível cujo nome é faixa numérica continua exigindo número: "Normal" do urobilinogênio não pode
// ser empurrado para dentro de "0,6-1,1" só porque soa saudável.
func TestMatchQualitativeLevel_NaoInventaFaixaNumerica(t *testing.T) {
	urobilinogenio := []models.ScoreLevel{
		{Level: 5, Name: "0,6-1,1"},
		{Level: 3, Name: "1,1-2,1 (1+)"},
		{Level: 0, Name: "≤0,1 (Ausente)"},
	}
	if l := matchQualitativeLevel(urobilinogenio, "Normal"); l != nil {
		t.Fatalf("\"Normal\" não é uma faixa: esperava sem nível, veio %v", *l)
	}
}

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
