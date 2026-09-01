package pdfdoc

import (
	"os"
	"strings"
	"testing"
)

func fdeck(v float64) *float64 { return &v }

// ferritinaRuler — uma régua realista (escala do catálogo + três resultados subindo).
func ferritinaRuler() Ruler {
	return Ruler{
		Code: "PLNCEFB97FD", Display: "Ferritina", Sub: "estoque de ferro", Unit: "ng/mL",
		Axis: [2]float64{0, 520},
		Segments: []RulerSegment{
			{Level: 0, A: 0, B: 15}, {Level: 2, A: 15, B: 30}, {Level: 3, A: 30, B: 50},
			{Level: 5, A: 50, B: 200}, {Level: 4, A: 200, B: 300}, {Level: 1, A: 300, B: 520},
		},
		History: []RulerPoint{
			{Value: 239.1, Text: "239,1", Date: "2024-05-27"},
			{Value: 432, Text: "432", Date: "2025-05-19"},
			{Value: 500, Text: "500", Date: "2026-04-29"},
		},
	}
}

func TestRulerSVGDesenhaAEscalaInteiraEOsDoisUltimosPontos(t *testing.T) {
	svg := rulerSVG(ferritinaRuler())

	if !strings.HasPrefix(svg, "<svg") || !strings.HasSuffix(svg, "</svg>") {
		t.Fatal("saída não é um SVG fechado")
	}
	// Um retângulo por faixa, mais o de fundo: a rampa inteira é a escala.
	if n := strings.Count(svg, "<rect"); n < 7 {
		t.Errorf("retângulos = %d, quer ao menos 7 (fundo + 6 faixas)", n)
	}
	// Só os dois últimos resultados aparecem; o de 2024 não polui a régua.
	if strings.Contains(svg, "239,1") {
		t.Error("a régua mostra no máximo os dois últimos pontos")
	}
	if !strings.Contains(svg, "432") || !strings.Contains(svg, "500") {
		t.Error("faltou o valor anterior ou o atual")
	}
	// A seta de direção é o que o leitor mais precisa e mais erra.
	if !strings.Contains(svg, "<path") {
		t.Error("faltou a seta do valor anterior para o atual")
	}
	if !strings.Contains(svg, "ng/mL") {
		t.Error("faltou a unidade")
	}
}

func TestRulerSVGUmPontoSoNaoDesenhaSeta(t *testing.T) {
	r := ferritinaRuler()
	r.History = r.History[:1]
	svg := rulerSVG(r)
	if strings.Contains(svg, "<path") {
		t.Error("com um resultado só não há direção para mostrar")
	}
}

func TestRulerSVGIdDeRecorteEUnicoPorRegua(t *testing.T) {
	// Dois SVGs na mesma página não podem compartilhar id de clipPath: um recorta o outro.
	a := ferritinaRuler()
	b := ferritinaRuler()
	b.Code, b.Display = "PLND2C05835", "TSH"
	if idOf(t, rulerSVG(a)) == idOf(t, rulerSVG(b)) {
		t.Error("réguas diferentes têm que ter clipPath diferente")
	}
}

func idOf(t *testing.T, svg string) string {
	t.Helper()
	i := strings.Index(svg, `<clipPath id="`)
	if i < 0 {
		t.Fatal("SVG sem clipPath")
	}
	rest := svg[i+len(`<clipPath id="`):]
	return rest[:strings.Index(rest, `"`)]
}

func TestMiniRulerMantemAOrdemDasFaixas(t *testing.T) {
	svg := miniRulerSVG(ferritinaRuler(), 500, true)
	if svg == "" {
		t.Fatal("mini-régua vazia")
	}
	// A mini comprime a escala para caber em 264px, mas a ORDEM das faixas no eixo é a mesma:
	// alguns exames pioram para a direita, outros para a esquerda.
	first := strings.Index(svg, rulerRamp[0])
	last := strings.Index(svg, rulerRamp[1])
	if first < 0 || last < 0 || first > last {
		t.Error("a ordem das faixas na mini tem que espelhar a da régua grande")
	}
	if strings.Count(svg, "<circle") != 2 {
		t.Error("faltou o ponto do paciente (anel + miolo)")
	}
}

func TestInlineHTMLLiberaEnfaseEEscapaOResto(t *testing.T) {
	// O <em> do punch é o que muda de cor na frase de fechamento.
	if got := inlineHTML("é cedo para <em>mudar a direção</em>"); got != "é cedo para <em>mudar a direção</em>" {
		t.Errorf("got %q", got)
	}
	// Conteúdo vem do JSONB: qualquer outra tag tem que virar texto.
	got := inlineHTML(`<script>alert(1)</script>`)
	if strings.Contains(got, "<script") {
		t.Errorf("tag não permitida passou: %q", got)
	}
	if got := inlineHTML("linha<br>outra"); got != "linha<br>outra" {
		t.Errorf("br devia passar: %q", got)
	}
}

func TestDeckHTMLNumeraEIdentificaCadaSlide(t *testing.T) {
	d := Deck{Title: "Teste", Slides: []DeckSlide{
		{Kind: DeckCover, Variant: DeckVariantDeep, Title: "José Ricardo"},
		{Kind: DeckRulers, Title: "A ferritina dobrou", Rulers: []DeckRulerBlock{{ferritinaRuler()}}, Legend: true},
	}}
	html, err := DeckHTML(d)
	if err != nil {
		t.Fatal(err)
	}
	if n := strings.Count(html, `<section class="slide`); n != 2 {
		t.Errorf("slides = %d, quer 2", n)
	}
	if !strings.Contains(html, `<span class="pagenum">01</span>`) ||
		!strings.Contains(html, `<span class="pagenum">02</span>`) {
		t.Error("numeração das páginas ausente ou fora do padrão 01/02")
	}
	// A capa usa h1; os demais slides usam h2.
	if !strings.Contains(html, "<h1>José Ricardo</h1>") {
		t.Error("a capa tem que sair em h1")
	}
	if !strings.Contains(html, "<h2>A ferritina dobrou</h2>") {
		t.Error("slide comum tem que sair em h2")
	}
	if !strings.Contains(html, "slide deep") {
		t.Error("a variante de fundo não foi aplicada")
	}
	// As fontes vão embutidas: o container não tem rede para o Google Fonts.
	if !strings.Contains(html, "@font-face") || !strings.Contains(html, "Fraunces") {
		t.Error("as fontes do deck têm que ir embutidas no HTML")
	}
	if strings.Contains(html, "fonts.googleapis.com") {
		t.Error("o deck não pode depender de fonte remota")
	}
}

func TestDeckHTMLRecusaDeckVazio(t *testing.T) {
	if _, err := DeckHTML(Deck{Title: "vazio"}); err == nil {
		t.Error("deck sem slides tem que dar erro, não PDF em branco")
	}
}

func TestA4SlideScaleEncaixaNaLargura(t *testing.T) {
	// O slide é escalado para a LARGURA da folha (A4 paisagem é mais "quadrada" que 16:9),
	// então nada é cortado e sobra faixa em cima e embaixo.
	got := a4SlideScale()
	if got <= 0 || got >= 1 {
		t.Fatalf("escala = %v, quer entre 0 e 1", got)
	}
	if w := SlideW * got; w < a4LandW-0.5 || w > a4LandW+0.5 {
		t.Errorf("slide escalado = %v px de largura, quer %v", w, a4LandW)
	}
	if h := SlideH * got; h > a4LandH {
		t.Errorf("slide escalado = %v px de altura, não cabe em %v", h, a4LandH)
	}
}

// sampleDeck — um deck curto que exercita os oito tipos de slide, no formato dos decks reais.
func sampleDeck() Deck {
	tsh := Ruler{
		Code: "PLND2C05835", Display: "TSH", Sub: "o comando da tireoide", Unit: "µUI/mL",
		Axis:     [2]float64{0, 10},
		Segments: []RulerSegment{{Level: 5, A: 0, B: 2}, {Level: 4, A: 2, B: 3}, {Level: 2, A: 3, B: 4.5}, {Level: 1, A: 4.5, B: 10}},
		History:  []RulerPoint{{Value: 3.9, Text: "3,9", Date: "2025-05-19"}, {Value: 4.98, Text: "4,98", Date: "2026-04-29"}},
		Note:     "alto há dois anos, e subindo",
	}
	fer := ferritinaRuler()
	return Deck{Title: "José Ricardo · Seus exames", Slides: []DeckSlide{
		{Kind: DeckCover, Variant: DeckVariantDeep, Eyebrow: "Seus exames · 27 de agosto de 2026",
			Title: "José Ricardo",
			Lede:  "Três anos de exames lidos juntos. A boa notícia primeiro, depois as coisas que estão se movendo."},
		{Kind: DeckSummary, Eyebrow: "Resumo", Title: "Onde você está, em uma página",
			Summary: &DeckSummaryBlock{
				Cards: []DeckSummaryCard{
					{Title: "O que está forte", Tone: "bom", Lines: []DeckSummaryLine{
						{Name: "Rim", Sub: "sem perda de proteína", Value: "2,7", Unit: "mg/g", Ruler: &fer, Plot: fdeck(120)},
					}},
					{Title: "O que está se movendo", Tone: "ruim", Lines: []DeckSummaryLine{
						{Name: "Ferritina", Sub: "dobrou em dois anos", Value: "500", Unit: "ng/mL", Ruler: &fer},
						{Name: "TSH", Sub: "alto há dois anos", Value: "4,98", Unit: "µUI/mL", Ruler: &tsh},
					}},
				},
				StepsTitle: "O que vamos fazer",
				Steps:      []string{"Tirzepatida uma vez por semana", "Treino de força com carga", "Creatina e ômega-3", "Os exames que faltam"},
			},
			Punch: "Nenhum dos quatro é doença ainda. <em>É cedo o bastante para mudar a direção.</em>"},
		{Kind: DeckRulers, Eyebrow: "O que está se movendo · 1 de 3", Title: "A ferritina dobrou em dois anos",
			Rulers: []DeckRulerBlock{{fer}, {tsh}}, Legend: true,
			Punch: "Subir a ferritina e <em>não</em> subir a saturação é a chave."},
		{Kind: DeckTwoCards, Eyebrow: "O que está se movendo · 1 de 3", Title: "Os dois caminhos, e qual é o seu",
			Cards: []DeckCard{
				{Kicker: "Caminho 1 · sobrecarga de origem genética", Dim: true,
					Body: "A marca é a saturação de transferrina alta, acima de 45%, junto com a ferritina alta."},
				{Kicker: "Caminho 2 · causa metabólica", Focus: true,
					Body: "A sua saturação é 30%, e caiu de 36% no ano passado."},
			}},
		{Kind: DeckPlanStep, Eyebrow: "O plano · 1 de 5", Title: "A tirzepatida: como começa",
			Lede:  "Uma aplicação por semana, sempre no mesmo dia.",
			Cards: []DeckCard{{Kicker: "Semanas 1 a 4", Body: "1,25 mg por semana."}, {Kicker: "A partir da 5", Body: "2,5 mg por semana."}}},
		{Kind: DeckSequence, Eyebrow: "A sequência", Title: "Os próximos três meses, em ordem",
			Steps: []DeckSeqStep{
				{When: "Esta semana", What: "Começa a tirzepatida", Detail: "Junto com creatina e ômega-3."},
				{When: "Em 4 semanas", What: "Consulta de ajuste", Detail: "Com os exames que faltam na mão."},
				{When: "Em 3 meses", What: "Reavaliação completa"},
			}},
		{Kind: DeckTakeaway, Eyebrow: "Para levar", Title: "O que você começa a tomar agora",
			Take: &DeckTakeawayBox{
				Highlight: &DeckHighlight{When: "Uma vez por semana", Name: "Tirzepatida",
					Obs: "1,25 mg nas quatro primeiras semanas", Dose: "1,25", Unit: "mg por semana"},
				Groups: []DeckDoseGroup{
					{Title: "De manhã", Items: []DeckDose{{Name: "Creatina", Dose: "5 g"}, {Name: "Ômega-3", Dose: "1 a 2 g"}}},
					{Title: "Todo dia", Items: []DeckDose{{Name: "Água", Sub: "mais nos dias de treino", Dose: "3,0 a 3,5 L"}}},
					{Title: "Na semana", Items: []DeckDose{{Name: "Força com carga", Dose: "2 a 3x"}}},
				},
				Note: "A fórmula manipulada ainda não entra: ela é montada na consulta de quatro semanas."}},
		{Kind: DeckClosing, Variant: DeckVariantDeep, Eyebrow: "Em uma página",
			Title: "Você está bem, e três marcadores estão se movendo na mesma direção.",
			Lede:  "Começamos pela tirzepatida, junto com proteína e treino de carga."},
	}}
}

func TestRenderDeck169(t *testing.T) {
	if !chromiumAvailable() {
		t.Skip("chromium ausente — pulando render")
	}
	pdf, err := RenderDeck(sampleDeck(), DeckPaper169)
	if err != nil {
		t.Fatalf("render 16:9: %v", err)
	}
	if len(pdf) < 2000 || string(pdf[:5]) != "%PDF-" {
		t.Fatalf("saída não parece PDF (len=%d)", len(pdf))
	}
	_ = os.WriteFile("/tmp/deck-16x9-test.pdf", pdf, 0o644)
	t.Logf("PDF 16:9: %d bytes -> /tmp/deck-16x9-test.pdf", len(pdf))
}

func TestRenderDeckA4(t *testing.T) {
	if !chromiumAvailable() {
		t.Skip("chromium ausente — pulando render")
	}
	pdf, err := RenderDeck(sampleDeck(), DeckPaperA4)
	if err != nil {
		t.Fatalf("render A4: %v", err)
	}
	if len(pdf) < 2000 || string(pdf[:5]) != "%PDF-" {
		t.Fatalf("saída não parece PDF (len=%d)", len(pdf))
	}
	_ = os.WriteFile("/tmp/deck-a4-test.pdf", pdf, 0o644)
	t.Logf("PDF A4: %d bytes -> /tmp/deck-a4-test.pdf", len(pdf))
}

// TestDeckNaoTransborda é a armadilha nº 7 do processo antigo virando teste: o slide tem altura
// fixa e overflow:hidden, então conteúdo demais SOME na impressão sem erro nenhum.
func TestDeckNaoTransborda(t *testing.T) {
	if !chromiumAvailable() {
		t.Skip("chromium ausente — pulando render")
	}
	over, err := CheckDeckOverflow(sampleDeck())
	if err != nil {
		t.Fatalf("medição: %v", err)
	}
	for _, o := range over {
		t.Errorf("slide %02d (%s) transborda: %.0fpx à direita, %.0fpx embaixo", o.Slide, o.Title, o.Right, o.Bottom)
	}
}

func TestCheckDeckOverflowAcusaSlideCheioDemais(t *testing.T) {
	if !chromiumAvailable() {
		t.Skip("chromium ausente — pulando render")
	}
	// Oito réguas não cabem num slide; o verificador precisa dizer isso em vez de deixar sumir.
	var rs []DeckRulerBlock
	for i := 0; i < 8; i++ {
		rs = append(rs, DeckRulerBlock{ferritinaRuler()})
	}
	over, err := CheckDeckOverflow(Deck{Title: "t", Slides: []DeckSlide{
		{Kind: DeckRulers, Title: "Réguas demais", Rulers: rs, Legend: true},
	}})
	if err != nil {
		t.Fatalf("medição: %v", err)
	}
	if len(over) == 0 {
		t.Fatal("oito réguas num slide deveriam ser acusadas como estouro")
	}
	if over[0].Slide != 1 || over[0].Bottom <= 0 {
		t.Errorf("estouro mal descrito: %+v", over[0])
	}
}

func TestRenderPlanReportAchataOsSlidesNoA4Assinado(t *testing.T) {
	if !chromiumAvailable() {
		t.Skip("chromium ausente — pulando render")
	}
	pdf, err := RenderPlanReport(PlanReport{
		Title:     "Relatório de devolutiva",
		Patient:   Patient{Name: "José Ricardo Mattos do Amaral", BirthInfo: "64 anos"},
		Slides:    sampleDeck().Slides,
		EmittedAt: "01/09/2026",
		Doctor:    Doctor{Name: "Dr. Getúlio José Mattos do Amaral Filho", Credentials: "CRM-PR 21.876 · Nefrologia"},
		Signature: Signature{Digital: true, SignedAt: "01/09/2026, 10:12 (horário de Brasília)",
			ValidateURL: "https://app.plenyasaude.com.br/documentos/validar/abc"},
	})
	if err != nil {
		t.Fatalf("render relatório: %v", err)
	}
	if len(pdf) < 2000 || string(pdf[:5]) != "%PDF-" {
		t.Fatalf("saída não parece PDF (len=%d)", len(pdf))
	}
	_ = os.WriteFile("/tmp/deck-report-test.pdf", pdf, 0o644)
	t.Logf("relatório: %d bytes -> /tmp/deck-report-test.pdf", len(pdf))
}

func TestRenderPlanReportRecusaPlanoVazio(t *testing.T) {
	if _, err := RenderPlanReport(PlanReport{Title: "x"}); err == nil {
		t.Error("plano sem slides tem que dar erro, não um relatório em branco")
	}
}

func TestRenderPlanReportNaoRepeteONomeDoPacienteDaCapa(t *testing.T) {
	// A papelaria já imprime o nome no bloco de identificação; a capa do deck o traria de novo
	// dois centímetros abaixo.
	html := flattenSlide(DeckSlide{Kind: DeckCover, Eyebrow: "Seus exames · 27/08", Title: "José Ricardo", Lede: "Três anos lidos juntos."})
	if strings.Contains(html, "José Ricardo") {
		t.Errorf("o título da capa não deveria sair no relatório: %s", html)
	}
	if !strings.Contains(html, "Seus exames") || !strings.Contains(html, "Três anos lidos juntos.") {
		t.Errorf("data e frase de abertura têm que continuar: %s", html)
	}
	outro := flattenSlide(DeckSlide{Kind: DeckRulers, Title: "A ferritina dobrou"})
	if !strings.Contains(outro, "A ferritina dobrou") {
		t.Errorf("slide comum perdeu o título: %s", outro)
	}
}

func TestRelatorioTrazOsTamanhosDeTextoDaRegua(t *testing.T) {
	// As classes .rg-* vivem no deckCSS, que o relatório não carrega. Sem repeti-las, o <text> herda
	// o corpo do documento e, como o valor é lido em unidades do viewBox, o nome do exame sai com
	// menos de 5px no PDF assinado.
	for _, cls := range []string{".rg-name", ".rg-sub", ".rg-tick", ".rg-val", ".rg-unit", ".rg-old", ".rg-note"} {
		if !strings.Contains(reportDeckCSS, cls) {
			t.Errorf("reportDeckCSS não define %s — a régua sairia ilegível no relatório", cls)
		}
	}
}
