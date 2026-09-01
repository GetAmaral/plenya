package services

import (
	"math"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

func ptrS(s string) *string { return &s }
func ptrI(i int) *int       { return &i }

func mustDay(t *testing.T, s string) time.Time {
	t.Helper()
	d, err := time.Parse("2006-01-02", s)
	if err != nil {
		t.Fatalf("data inválida %q: %v", s, err)
	}
	return d
}

// levelsFerritina reproduz a escala real da ferritina masculina do catálogo: um nível aberto para
// baixo, um aberto para cima e quatro faixas fechadas fora de ordem no banco.
func levelsFerritina() []models.ScoreLevel {
	return []models.ScoreLevel{
		{Level: 0, Operator: "<=", UpperLimit: ptrS("15")},
		{Level: 1, Operator: ">", LowerLimit: ptrS("300")},
		{Level: 2, Operator: "between", LowerLimit: ptrS("15"), UpperLimit: ptrS("30")},
		{Level: 3, Operator: "between", LowerLimit: ptrS("30"), UpperLimit: ptrS("50")},
		{Level: 4, Operator: "between", LowerLimit: ptrS("200"), UpperLimit: ptrS("300")},
		{Level: 5, Operator: "between", LowerLimit: ptrS("50"), UpperLimit: ptrS("200")},
	}
}

func TestLevelEdgesOrdenaESemRepetir(t *testing.T) {
	got := levelEdges(levelsFerritina())
	want := []float64{15, 30, 50, 200, 300}
	if len(got) != len(want) {
		t.Fatalf("edges = %v, quer %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("edges = %v, quer %v", got, want)
		}
	}
}

func TestRulerAxisSemHistoricoUsaSoOPadding(t *testing.T) {
	// vão 15→300 = 285; padding 22% = 62,7. O piso é zero: exame não tem valor negativo.
	axis := rulerAxis([]float64{15, 30, 50, 200, 300}, nil)
	if axis[0] != 0 {
		t.Errorf("piso = %v, quer 0 (15 - 62,7 é negativo)", axis[0])
	}
	if math.Abs(axis[1]-362.7) > 1e-9 {
		t.Errorf("teto = %v, quer 362,7", axis[1])
	}
}

func TestRulerAxisEsticaParaCaberOValorDoPaciente(t *testing.T) {
	// Valor de 500 estoura o teto de 362,7: o eixo tem que crescer para o ponto não colar na borda.
	hist := []dto.PlanRulerPoint{{Value: 239.1}, {Value: 432}, {Value: 500}}
	axis := rulerAxis([]float64{15, 30, 50, 200, 300}, hist)
	if math.Abs(axis[1]-520) > 1e-9 {
		t.Errorf("teto = %v, quer 520 (500 × 1,04)", axis[1])
	}
	if axis[0] != 0 {
		t.Errorf("piso = %v, quer 0", axis[0])
	}
}

func TestRulerAxisEsticaParaBaixoQuandoOValorFicaAbaixoDaEscala(t *testing.T) {
	// RDW: fronteiras 13→15,6 (vão 2,6; padding 0,572) e resultado 12,4 abaixo do piso 12,428.
	axis := rulerAxis([]float64{13, 14.5, 15.6}, []dto.PlanRulerPoint{{Value: 12.4}})
	if math.Abs(axis[0]-11.904) > 1e-9 {
		t.Errorf("piso = %v, quer 11,904 (12,4 × 0,96)", axis[0])
	}
	if math.Abs(axis[1]-16.172) > 1e-9 {
		t.Errorf("teto = %v, quer 16,172 (padding, sem influência do histórico)", axis[1])
	}
}

func TestLevelSegmentsFechaAsPontasNoEixo(t *testing.T) {
	axis := []float64{0, 520}
	segs := levelSegments(levelsFerritina(), axis)
	if len(segs) != 6 {
		t.Fatalf("segmentos = %d, quer 6", len(segs))
	}
	// Ordenados pela posição no eixo, não pelo número do nível.
	if segs[0].Level != 0 || segs[0].A != 0 || segs[0].B != 15 {
		t.Errorf("primeiro segmento = %+v, quer nível 0 de 0 a 15", segs[0])
	}
	if segs[0].Label != "≤15" {
		t.Errorf("rótulo do nível aberto para baixo = %q, quer \"≤15\"", segs[0].Label)
	}
	last := segs[len(segs)-1]
	if last.Level != 1 || last.A != 300 || last.B != 520 {
		t.Errorf("último segmento = %+v, quer nível 1 de 300 até o teto do eixo", last)
	}
	if last.Label != ">300" {
		t.Errorf("rótulo do nível aberto para cima = %q, quer \">300\"", last.Label)
	}
	if segs[1].Label != "15-30" {
		t.Errorf("rótulo de faixa fechada = %q, quer \"15-30\"", segs[1].Label)
	}
}

func TestLevelSegmentsIgnoraNivelCategorico(t *testing.T) {
	// Item de anamnese usa operador "=": não tem faixa, e não pode virar uma régua torta.
	segs := levelSegments([]models.ScoreLevel{
		{Level: 5, Operator: "=", LowerLimit: ptrS("1")},
		{Level: 0, Operator: "=", LowerLimit: ptrS("0")},
	}, []float64{0, 1})
	if len(segs) != 0 {
		t.Errorf("segmentos = %+v, quer nenhum", segs)
	}
}

func TestBuildRulerRecusaEscalaSemFaixa(t *testing.T) {
	item := models.ScoreItem{Name: "Tabagismo atual", Levels: []models.ScoreLevel{
		{Level: 5, Operator: "=", LowerLimit: ptrS("0")},
	}}
	if _, ok := buildRuler("X", &item, nil); ok {
		t.Error("item categórico não deveria virar régua")
	}
}

func TestFormatNumberPTUsaVirgulaESomeComCasaInutil(t *testing.T) {
	// 1,06392 sai com 4 CASAS (o que o banco guarda em decimal(12,4)), não com 4 dígitos
	// significativos: a regra antiga vinha do %.4g que também produzia notação científica.
	cases := map[float64]string{15: "15", 1.003: "1,003", 0.9: "0,9", 300: "300", 1.06392: "1,0639"}
	for in, want := range cases {
		if got := formatNumberPT(in); got != want {
			t.Errorf("formatNumberPT(%v) = %q, quer %q", in, got, want)
		}
	}
}

func TestResultDisplayTextPreservaAGrafiaDoLaudo(t *testing.T) {
	// O laudo imprimiu "3,00"; mostrar "3" faria o paciente achar que é outro exame.
	if got := resultDisplayText("3.00", 3.0); got != "3,00" {
		t.Errorf("got %q, quer \"3,00\"", got)
	}
	// Texto que não bate com o número gravado não é confiável: cai na formatação do número.
	if got := resultDisplayText("999", 3.0); got != "3" {
		t.Errorf("got %q, quer \"3\"", got)
	}
	// Sem texto no laudo, formata o número.
	if got := resultDisplayText("", 1.023); got != "1,023" {
		t.Errorf("got %q, quer \"1,023\"", got)
	}
}

func TestBuildHistoryFicaComAUltimaLeituraDoDia(t *testing.T) {
	day := func(s string) labRow {
		d := mustDay(t, s)
		return labRow{Collected: d, Day: collectionDay(d)}
	}
	a, b := day("2026-01-10"), day("2026-01-10")
	a.Numeric, b.Numeric = 10, 12
	c := day("2026-03-01")
	c.Numeric = 8
	got := buildHistory([]labRow{a, b, c}, nil)
	if len(got) != 2 {
		t.Fatalf("pontos = %d, quer 2 (um por dia)", len(got))
	}
	if got[0].Value != 12 {
		t.Errorf("valor do dia repetido = %v, quer 12 (a última leitura)", got[0].Value)
	}
	if got[0].Date > got[1].Date {
		t.Error("histórico tem que sair do mais antigo para o mais recente")
	}
}

func TestClassifyFindingsSeparaEOrdenaPeloQueMaisPesa(t *testing.T) {
	rulers := map[string]dto.PlanRuler{
		"LEVE":  {Code: "LEVE", Name: "Leve", Points: 6, History: []dto.PlanRulerPoint{{Level: ptrI(1), Value: 1}}},
		"PESA":  {Code: "PESA", Name: "Pesa", Points: 35, History: []dto.PlanRulerPoint{{Level: ptrI(2), Value: 2}}},
		"OTIMO": {Code: "OTIMO", Name: "Ótimo", Points: 10, History: []dto.PlanRulerPoint{{Level: ptrI(5), Value: 3}}},
		"CAIU": {Code: "CAIU", Name: "Caiu", Points: 18, History: []dto.PlanRulerPoint{
			{Level: ptrI(5), Value: 4}, {Level: ptrI(4), Value: 5},
		}},
	}
	strong, moving := classifyFindings(rulers, nil)

	if len(strong) != 1 || strong[0].Code != "OTIMO" {
		t.Errorf("strong = %+v, quer só OTIMO", strong)
	}
	// PESA (nível 2, 35 pontos) perde 21; LEVE (nível 1, 6 pontos) perde 4,8; CAIU perde 3,6.
	wantOrder := []string{"PESA", "LEVE", "CAIU"}
	if len(moving) != len(wantOrder) {
		t.Fatalf("moving = %+v, quer %v", moving, wantOrder)
	}
	for i, code := range wantOrder {
		if moving[i].Code != code {
			t.Errorf("moving[%d] = %s, quer %s", i, moving[i].Code, code)
		}
	}
	// Nível bom que PIOROU entra em "se movendo", não em "está bem": a direção é o sinal.
	if moving[2].Trend != dto.PlanTrendWorsening {
		t.Errorf("tendência de CAIU = %q, quer worsening", moving[2].Trend)
	}
}

func TestClassifyFindingsPrefereOsPontosDoEscoreQuandoExistem(t *testing.T) {
	id := uuid.Must(uuid.NewV7())
	rulers := map[string]dto.PlanRuler{
		"A": {Code: "A", ScoreItemID: id.String(), Points: 10, History: []dto.PlanRulerPoint{{Level: ptrI(1)}}},
	}
	_, moving := classifyFindings(rulers, map[uuid.UUID]float64{id: 7.5})
	if len(moving) != 1 || moving[0].PointsLost != 7.5 {
		t.Errorf("pointsLost = %+v, quer 7,5 vindo do snapshot", moving)
	}
}

func TestClassifyFindingsPegaOsPontosDaVarianteCERTA(t *testing.T) {
	// Caso IGF-1: o mesmo código de exame tem um item guarda-chuva e variantes por faixa etária, e
	// o motor avalia TODAS. Os pontos têm que vir da variante que desenhou a régua; pegar a outra
	// troca a ordem dos achados, que é ordenada justamente por eles.
	usada := uuid.Must(uuid.NewV7())
	outraVariante := uuid.Must(uuid.NewV7())
	rulers := map[string]dto.PlanRuler{
		"IGF1": {Code: "IGF1", ScoreItemID: usada.String(), Points: 20, History: []dto.PlanRulerPoint{{Level: ptrI(2)}}},
	}
	_, moving := classifyFindings(rulers, map[uuid.UUID]float64{
		usada:         4.0,
		outraVariante: 19.0,
	})
	if len(moving) != 1 {
		t.Fatalf("moving = %+v, quer 1 achado", moving)
	}
	if moving[0].PointsLost != 4.0 {
		t.Errorf("pointsLost = %v, quer 4 (da variante usada), não 19 (da outra)", moving[0].PointsLost)
	}
}

func TestClassifyFindingsEstimaQuandoOEscoreNaoTemOItem(t *testing.T) {
	// Sem escore calculado a ordenação ainda precisa funcionar: estima pelo peso e pelo nível.
	rulers := map[string]dto.PlanRuler{
		"A": {Code: "A", ScoreItemID: uuid.Must(uuid.NewV7()).String(), Points: 10,
			History: []dto.PlanRulerPoint{{Level: ptrI(0)}}},
	}
	_, moving := classifyFindings(rulers, map[uuid.UUID]float64{})
	if len(moving) != 1 || moving[0].PointsLost != 10 {
		t.Errorf("pointsLost = %+v, quer 10 (nível 0 perde o peso inteiro)", moving)
	}
}

func TestClassifyFindingsDevolveListasVaziasNuncaNil(t *testing.T) {
	// Slice nil vira `null` em JSON e quebra o .map() do front. Paciente sem exame é caso normal.
	strong, moving := classifyFindings(map[string]dto.PlanRuler{}, nil)
	if strong == nil || moving == nil {
		t.Fatalf("strong=%v moving=%v; as duas têm que ser [] e não nil", strong, moving)
	}
	if len(strong) != 0 || len(moving) != 0 {
		t.Errorf("sem régua não pode haver achado: strong=%v moving=%v", strong, moving)
	}
}

func TestClassifyFindingsIgnoraReguaSemNivelClassificado(t *testing.T) {
	// Resultado que o classificador não conseguiu pontuar não vira achado nem para um lado
	// nem para o outro — entrar como nível 0 seria inventar um alarme.
	rulers := map[string]dto.PlanRuler{
		"SEMNIVEL": {Code: "SEMNIVEL", Points: 10, History: []dto.PlanRulerPoint{{Level: nil, Value: 1}}},
		"SEMHIST":  {Code: "SEMHIST", Points: 10},
	}
	strong, moving := classifyFindings(rulers, nil)
	if len(strong) != 0 || len(moving) != 0 {
		t.Errorf("strong=%v moving=%v; quer nenhum achado", strong, moving)
	}
}

func TestCollectionDayNaoRepartUmaColetaEmDoisDias(t *testing.T) {
	sp := saoPaulo()

	// Dia-calendário digitado: gravado como meia-noite UTC. No fuso local é 21:00 do dia ANTERIOR,
	// então formatar em São Paulo volta um dia e reparte a coleta.
	diaCalendario := time.Date(2024, 11, 5, 0, 0, 0, 0, time.UTC)
	if got := collectionDay(diaCalendario); got != "2024-11-05" {
		t.Errorf("dia-calendário = %q, quer \"2024-11-05\"", got)
	}

	// Instante real de coleta, de dia: os dois fusos concordam.
	coletaDeDia := time.Date(2024, 11, 5, 9, 0, 0, 0, sp)
	if got := collectionDay(coletaDeDia); got != "2024-11-05" {
		t.Errorf("coleta de dia = %q, quer \"2024-11-05\"", got)
	}
	if collectionDay(diaCalendario) != collectionDay(coletaDeDia) {
		t.Error("as duas formas de gravar a mesma coleta têm que cair no mesmo dia")
	}

	// Instante real à noite: em UTC já virou o dia seguinte, mas o dia da coleta é o local.
	importacaoDaNoite := time.Date(2026, 2, 6, 23, 12, 29, 900000000, sp)
	if got := collectionDay(importacaoDaNoite); got != "2026-02-06" {
		t.Errorf("importação da noite = %q, quer \"2026-02-06\" (em UTC seria 02-07)", got)
	}
	tarde := time.Date(2026, 2, 7, 21, 7, 38, 0, sp)
	if got := collectionDay(tarde); got != "2026-02-07" {
		t.Errorf("noite de 07 = %q, quer \"2026-02-07\" (em UTC seria 02-08)", got)
	}
}

func TestLevelForValueUsaAMesmaRegraDoMotor(t *testing.T) {
	lv := levelsFerritina()
	cases := map[float64]int{10: 0, 20: 2, 40: 3, 100: 5, 250: 4, 400: 1}
	for value, want := range cases {
		got := levelForValue(lv, value)
		if got == nil {
			t.Errorf("valor %v ficou sem nível, quer %d", value, want)
			continue
		}
		if *got != want {
			t.Errorf("valor %v caiu no nível %d, quer %d", value, *got, want)
		}
	}
}

func TestLevelForValueNaoInventaNivelForaDaEscala(t *testing.T) {
	// Melhor sem nível do que classificado errado: régua sem bolinha colorida é honesta.
	lv := []models.ScoreLevel{{Level: 5, Operator: "between", LowerLimit: ptrS("10"), UpperLimit: ptrS("20")}}
	if got := levelForValue(lv, 999); got != nil {
		t.Errorf("valor fora de todas as faixas recebeu nível %d, quer nenhum", *got)
	}
}

func TestRulerAxisNaoDestroiEscalaNegativa(t *testing.T) {
	// T-score de densitometria vive INTEIRO no negativo. Um piso em zero apagava a escala: o eixo
	// virava [0,1], todo segmento caía fora e a barra saía pintada de "ótimo" — um paciente com
	// osteoporose receberia um PDF dizendo o contrário do exame.
	edges := []float64{-2.5, -2, -1.5, -1}
	axis := rulerAxis(edges, []dto.PlanRulerPoint{{Value: -2.8}})

	if axis[0] >= edges[0] {
		t.Errorf("piso = %v, tem que ficar abaixo de %v para a escala caber", axis[0], edges[0])
	}
	if axis[1] <= edges[len(edges)-1] {
		t.Errorf("teto = %v, tem que ficar acima de %v", axis[1], edges[len(edges)-1])
	}
	// O valor do paciente tem que caber DENTRO do eixo, não colar na borda.
	if -2.8 < axis[0] || -2.8 > axis[1] {
		t.Errorf("valor -2,8 ficou fora do eixo %v", axis)
	}
}

func TestRulerAxisMantemOPisoEmZeroQuandoAEscalaEPositiva(t *testing.T) {
	// Exame de sangue não tem valor negativo: aqui o piso continua valendo, senão metade da barra
	// seria desperdiçada com valores impossíveis.
	axis := rulerAxis([]float64{15, 30, 50, 200, 300}, nil)
	if axis[0] != 0 {
		t.Errorf("piso = %v, quer 0", axis[0])
	}
}

func TestLevelSegmentsRotulaFaixaNegativaDeFormaLegivel(t *testing.T) {
	segs := levelSegments([]models.ScoreLevel{
		{Level: 1, Operator: "between", LowerLimit: ptrS("-2.5"), UpperLimit: ptrS("-2.0")},
	}, []float64{-3, 0})
	if len(segs) != 1 {
		t.Fatalf("segmentos = %d, quer 1", len(segs))
	}
	// "-2,5--2" não se lê.
	if segs[0].Label != "-2,5 a -2" {
		t.Errorf("rótulo = %q, quer \"-2,5 a -2\"", segs[0].Label)
	}
}

func TestFormatNumberPTNuncaUsaNotacaoCientifica(t *testing.T) {
	// CK, triglicerídeo e ferritina passam de 1000. Com %.4g, 12345,6 saía "1,235e+04" no PDF do
	// paciente e 1234,5 perdia a casa decimal em silêncio.
	cases := map[float64]string{
		1234.5:   "1234,5",
		12345.6:  "12345,6",
		253000.4: "253000,4",
		500:      "500",
		1.003:    "1,003",
		0.11:     "0,11",
	}
	for in, want := range cases {
		if got := formatNumberPT(in); got != want {
			t.Errorf("formatNumberPT(%v) = %q, quer %q", in, got, want)
		}
	}
}

func TestRulerAxisEsticaDeVerdadeEmValorNegativo(t *testing.T) {
	// Multiplicar valor negativo por 0,96 o APROXIMA do zero: a folga encolhia em vez de crescer e
	// o eixo nunca continha o ponto do paciente. T-score de -3,5 numa escala de -2,5 a -1.
	axis := rulerAxis([]float64{-2.5, -2, -1.5, -1}, []dto.PlanRulerPoint{{Value: -3.5}})
	if axis[0] > -3.5 {
		t.Errorf("piso = %v, tem que ficar abaixo de -3,5 para o ponto caber", axis[0])
	}
}

func TestAnamnesisFindingsUsaSelectedLevelETemTendencia(t *testing.T) {
	// O nível vem de selected_level, não de numeric_value: ler o número cru no lugar do nível já
	// fez o escore sair zerado antes. E uma resposta que muda entre consultas tem direção.
	id := uuid.Must(uuid.NewV7())
	rows := []anamnesisRow{
		{ScoreItemID: id, Code: "SONO", Name: "Sono", Points: 20, Level: ptrI(5), Day: "2026-02-01"},
		{ScoreItemID: id, Code: "SONO", Name: "Sono", Points: 20, Level: ptrI(2), Day: "2026-02-11"},
	}
	got := anamnesisFindings(rows, nil)
	if len(got) != 1 {
		t.Fatalf("achados = %d, quer 1 (um por item, com histórico)", len(got))
	}
	f := got[0]
	if f.Source != dto.PlanSourceAnamnesis {
		t.Errorf("origem = %q, quer anamnesis", f.Source)
	}
	if f.Level != 2 {
		t.Errorf("nível = %d, quer 2 (a resposta mais recente)", f.Level)
	}
	if f.Trend != dto.PlanTrendWorsening {
		t.Errorf("tendência = %q, quer worsening", f.Trend)
	}
	if f.Kind != dto.PlanFindingMoving {
		t.Errorf("nível 2 tem que entrar em 'se movendo', veio %q", f.Kind)
	}
}

func TestAnamnesisFindingsIgnoraRespostaSemNivel(t *testing.T) {
	rows := []anamnesisRow{{ScoreItemID: uuid.Must(uuid.NewV7()), Name: "x", Level: nil}}
	if got := anamnesisFindings(rows, nil); len(got) != 0 {
		t.Errorf("resposta sem nível não pode virar achado: %+v", got)
	}
}

func TestStrongOrdenaPeloPesoDoItemENaoPelosPontosPerdidos(t *testing.T) {
	// Em nível 4-5 quase ninguém perdeu ponto, então ordenar por pontos perdidos deixa tudo
	// empatado em zero — e o checklist de ausência da anamnese ("Adrenalectomia: não", nível 5)
	// afoga o marcador pesado que está de fato no ótimo.
	rulers := map[string]dto.PlanRuler{
		"AUSENCIA": {Code: "AUSENCIA", Name: "Adrenalectomia", Points: 9, History: []dto.PlanRulerPoint{{Level: ptrI(5)}}},
		"PESADO":   {Code: "PESADO", Name: "Glicose", Points: 35, History: []dto.PlanRulerPoint{{Level: ptrI(5)}}},
	}
	strong, _ := classifyFindings(rulers, nil)
	if len(strong) != 2 {
		t.Fatalf("strong = %+v, quer 2", strong)
	}
	if strong[0].Code != "PESADO" {
		t.Errorf("primeiro = %s, quer PESADO (35 pontos antes de 9)", strong[0].Code)
	}
}

func TestLevelSegmentsRespeitaOSinalDoOperador(t *testing.T) {
	// "< 15" rotulado como "≤15" diz ao paciente que 15 pertence a esta faixa, quando o motor o
	// classifica na vizinha.
	segs := levelSegments([]models.ScoreLevel{
		{Level: 0, Operator: "<", UpperLimit: ptrS("15")},
		{Level: 5, Operator: ">=", LowerLimit: ptrS("50")},
	}, []float64{0, 100})
	if len(segs) != 2 {
		t.Fatalf("segmentos = %d, quer 2", len(segs))
	}
	if segs[0].Label != "<15" {
		t.Errorf("operador '<' rotulado como %q, quer \"<15\"", segs[0].Label)
	}
	if segs[1].Label != "≥50" {
		t.Errorf("operador '>=' rotulado como %q, quer \"≥50\"", segs[1].Label)
	}
}
