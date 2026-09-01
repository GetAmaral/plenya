package services

import (
	"testing"

	"github.com/plenya/api/internal/dto"
)

func temValor(n Numeral, v float64) bool {
	for _, c := range n.Values {
		if c == v {
			return true
		}
	}
	return false
}

// O laudo escreve número em PT-BR e em inglês, às vezes no mesmo lote. Errar a leitura aqui não dá
// erro: dá um número que não casa com o dossiê e vira sugestão à toa, ou pior, casa com o errado.
func TestExtractNumerals(t *testing.T) {
	casos := []struct {
		texto    string
		quer     []float64 // um valor por numeral, na ordem; ambíguo confere com temValor
		unidades []string
		porque   string
	}{
		{"a glicose está em 139", []float64{139}, []string{""}, "inteiro solto"},
		{"HOMA-IR de 13,5", []float64{13.5}, []string{""}, "decimal PT-BR"},
		{"PCR de 63,10 mg/L", []float64{63.10}, []string{"mg/L"}, "decimal com unidade composta"},
		{"proteína de 112 g por dia", []float64{112}, []string{"g"}, "unidade separada por espaço"},
		{"saturação de 12%", []float64{12}, []string{"%"}, "percentual colado"},
		{"T-score de -2,5", []float64{-2.5}, []string{""}, "negativo (densitometria)"},
		{"contagem de 1.234,5 células", []float64{1234.5}, []string{"células"}, "milhar e decimal juntos"},
		{"vitamina D de 18 ng/mL", []float64{18}, []string{"ng/mL"}, "unidade com barra"},
	}
	for _, c := range casos {
		got := ExtractNumerals(c.texto)
		if len(got) != len(c.quer) {
			t.Errorf("%q: achou %d numerais, esperava %d (%s)", c.texto, len(got), len(c.quer), c.porque)
			continue
		}
		for i := range c.quer {
			if !temValor(got[i], c.quer[i]) {
				t.Errorf("%q: numeral %d = %v, esperava conter %v (%s)", c.texto, i, got[i].Values, c.quer[i], c.porque)
			}
			if got[i].Unit != c.unidades[i] {
				t.Errorf("%q: unidade %d = %q, esperava %q", c.texto, i, got[i].Unit, c.unidades[i])
			}
		}
	}
}

// `1.023` é mil e vinte e três em PT-BR e 1,023 em inglês, e a densidade urinária vale 1,023 de
// verdade. Escolher uma leitura seria chutar; as duas entram, e casar com qualquer uma basta,
// porque a pergunta é "este número existe no dossiê", não "quanto ele vale".
func TestExtractNumerals_PontoAmbiguo(t *testing.T) {
	n := ExtractNumerals("densidade de 1.023")
	if len(n) != 1 {
		t.Fatalf("achou %d numerais", len(n))
	}
	if !temValor(n[0], 1.023) {
		t.Errorf("faltou a leitura decimal (1,023) em %v", n[0].Values)
	}
	if !temValor(n[0], 1023) {
		t.Errorf("faltou a leitura de milhar (1023) em %v", n[0].Values)
	}
}

func dossieDeTeste() *dto.PlanDossierResponse {
	i := func(v int) *int { return &v }
	f := func(v float64) *float64 { return &v }
	return &dto.PlanDossierResponse{
		Patient: dto.PlanDossierPatient{Age: 63},
		Rulers: map[string]dto.PlanRuler{
			"PLNFERR": {
				Code: "PLNFERR", Name: "Ferritina", Unit: "ng/mL", Points: 18,
				Axis:     []float64{0, 520},
				Segments: []dto.PlanRulerSegment{{Level: 5, A: 50, B: 200}},
				History: []dto.PlanRulerPoint{
					{Date: "2025-08-14", Value: 48, Text: "48"},
					{Date: "2026-02-06", Value: 96.8, Text: "96,8"},
				},
			},
			"PLNUSG": {
				Code: "PLNUSG", Name: "Densidade urinária", Unit: "",
				History: []dto.PlanRulerPoint{{Date: "2026-02-06", Value: 1.023, Text: "1,023"}},
			},
		},
		Moving: []dto.PlanFinding{
			{Code: "PLNGLI", Name: "Glicose de jejum", Unit: "mg/dL", Value: 139, PointsLost: 35},
		},
		Vitals: []dto.PlanDossierVitals{
			{SystolicBP: i(140), DiastolicBP: i(90), Weight: f(75), Height: f(165)},
		},
		Snapshot: &dto.PlanDossierSnapshot{TotalPercentage: 60.2},
	}
}

func TestNumericIndex_Match(t *testing.T) {
	ix := BuildNumericIndex(dossieDeTeste())
	acha := func(texto string) []NumeralFact {
		ns := ExtractNumerals(texto)
		if len(ns) == 0 {
			t.Fatalf("nenhum numeral em %q", texto)
		}
		return ix.Match(ns[0])
	}

	if len(acha("48")) == 0 {
		t.Error("48 está no histórico da ferritina e não foi encontrado")
	}
	if len(acha("139")) == 0 {
		t.Error("139 é a glicose e não foi encontrado")
	}
	if len(acha("140")) == 0 {
		t.Error("140 é a pressão sistólica e não foi encontrado")
	}
	if len(acha("63")) == 0 {
		t.Error("63 é a idade da paciente e não foi encontrado")
	}
	if len(acha("1.023")) == 0 {
		t.Error("1,023 é a densidade e não foi encontrado (leitura ambígua)")
	}

	// Arredondamento: quem escreve "97" está arredondando 96,8, e isso é legítimo.
	if len(acha("97")) == 0 {
		t.Error("97 deveria casar com 96,8 por arredondamento de quem escreve sem casa decimal")
	}
	// Mas quem escreve duas casas afirmou duas casas.
	if len(acha("96,80")) == 0 {
		t.Error("96,80 é o próprio valor")
	}
	if len(acha("96,5")) != 0 {
		t.Error("96,5 não existe no dossiê e não pode casar com 96,8")
	}

	if len(acha("777")) != 0 {
		t.Error("777 não existe em lugar nenhum e casou — a verificação não estaria valendo nada")
	}

	// A origem tem que ser legível: é o que vai ao lado do botão de aceitar.
	fatos := acha("48")
	if fatos[0].Label == "" || fatos[0].Source == "" {
		t.Errorf("fato sem origem legível: %+v", fatos[0])
	}
}

// Classificar pelo CAMPO não funciona: `punch` e `title` carregam número o tempo todo. O que
// separa "encurta o título" de "escreve que subiu para 96" é o numeral NOVO.
func TestNumeralDelta(t *testing.T) {
	casos := []struct {
		antes, depois string
		novos         int
		porque        string
	}{
		{"A ferritina subiu", "A ferritina subiu bastante", 0, "reescrita sem número"},
		// Ganhar unidade conta como mudança, e é decisão deliberada: a MESMA regra que pega isto é
		// a que pega "112 g" virar "112 mg". Unidade é onde os erros de escore desta base moraram
		// (creatinina em U/L, sedimento em células/campo), e passar unidade em silêncio num
		// documento que o paciente lê é pior do que pedir um aceite a mais.
		{"A ferritina está em 48", "A ferritina está em 48 ng/mL", 1, "unidade nova é mudança"},
		{"A ferritina subiu", "A ferritina subiu de 48 para 96", 2, "dois números novos"},
		{"Ferritina em 48", "Ferritina em 96", 1, "trocou o número"},
		{"Proteína de 112 g", "Proteína de 112 mg", 1, "mesmo número, unidade diferente: é mudança"},
		{"Três meses", "Três meses de acompanhamento", 0, "número por extenso não é numeral"},
		{"", "Comece com 1,25 mg", 1, "campo vazio ganhando número"},
	}
	for _, c := range casos {
		got := NumeralDelta(c.antes, c.depois)
		if len(got) != c.novos {
			t.Errorf("%q -> %q: %d numerais novos, esperava %d (%s) — %v",
				c.antes, c.depois, len(got), c.novos, c.porque, got)
		}
	}
}

// Dossiê nulo não pode explodir: um plano recém-criado pode não ter congelamento ainda.
func TestNumericIndex_Vazio(t *testing.T) {
	ix := BuildNumericIndex(nil)
	if ix == nil {
		t.Fatal("índice nulo")
	}
	if len(ix.Match(Numeral{Raw: "1", Values: []float64{1}})) != 0 {
		t.Error("índice vazio devolveu fato")
	}
	var nulo *NumericIndex
	if len(nulo.Match(Numeral{Values: []float64{1}})) != 0 {
		t.Error("índice nil deveria devolver vazio em vez de estourar")
	}
}

// Os onze formatos que aparecem de verdade no histórico de um paciente com 83 réguas, colhidos do
// dossiê de produção antes de escrever isto: `9,99` (41 ocorrências), `99,9` (38), `99` (37),
// `999` (25), `9,9` (24), `9,999` (10), `99,99` (9), `9` (6), `999,9` (4), `99,999` (2) e `9999`.
//
// A verificação contra o dossiê real rodou uma vez, fora do repositório: 197 de 197 valores do
// histórico foram reencontrados pelo próprio índice, e nenhum número inventado casou. O fixture
// real NÃO fica versionado — dado clínico mora em `pacs/`, que é gitignored. O que fica é a forma.
func TestExtractNumerals_FormatosReaisDoLaboratorio(t *testing.T) {
	formatos := []struct {
		texto string
		valor float64
	}{
		{"9,99", 9.99},
		{"99,9", 99.9},
		{"99", 99},
		{"999", 999},
		{"9,9", 9.9},
		{"9,999", 9.999},
		{"99,99", 99.99},
		{"9", 9},
		{"999,9", 999.9},
		{"99,999", 99.999},
		{"9999", 9999},
	}
	for _, f := range formatos {
		ns := ExtractNumerals(f.texto)
		if len(ns) != 1 {
			t.Errorf("%q: extraiu %d numerais", f.texto, len(ns))
			continue
		}
		if !temValor(ns[0], f.valor) {
			t.Errorf("%q: leu %v, esperava %v", f.texto, ns[0].Values, f.valor)
		}
	}
}

// O ciclo completo: o valor formatado que o dossiê guarda tem que ser reencontrado pelo índice
// construído do mesmo dossiê. É o caminho exato que uma frase do deck percorre ao citar um número,
// e foi assim que os 197 valores reais foram conferidos.
func TestNumericIndex_RoundTripDoHistorico(t *testing.T) {
	d := dossieDeTeste()
	ix := BuildNumericIndex(d)
	for code, r := range d.Rulers {
		for _, p := range r.History {
			ns := ExtractNumerals(p.Text)
			if len(ns) == 0 {
				t.Errorf("%s: não extraiu numeral de %q", code, p.Text)
				continue
			}
			if len(ix.Match(ns[0])) == 0 {
				t.Errorf("%s: o valor %q não foi reencontrado no índice construído do mesmo dossiê", code, p.Text)
			}
		}
	}
}
