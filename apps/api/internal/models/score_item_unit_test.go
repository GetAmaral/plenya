package models

import "testing"

// Os pares abaixo são os que aparecem DE VERDADE em produção e no dev, colhidos cruzando
// score_items.unit com lab_results.unit nos snapshots mais recentes. A guarda de unidade só
// presta se recusar os que de fato medem outra coisa e deixar passar os que são a mesma
// grandeza escrita de outro jeito: recusar demais silencia item correto.
func TestUnitMatches_ParesReaisDoCatalogo(t *testing.T) {
	casos := []struct {
		escala, resultado string
		casa              bool
		porque            string
	}{
		// Mesma grandeza, grafia diferente: tem que passar.
		{"IU/mL", "UI/mL", true, "UI é a grafia portuguesa de IU"},
		{"IU/L", "UI/L", true, "idem"},
		{"mIU/mL", "mUI/mL", true, "idem"},
		{"µUI/mL", "µU/mL", true, "U e UI no mesmo ensaio"},
		{"µg/dL", "mcg/dL", true, "mc é a grafia ASCII de µ"},
		{"µg/dL", "ug/dL", true, "u ASCII no lugar do sinal de micro"},
		{"μg/L", "ug/L", true, "mu grego U+03BC no lugar do micro U+00B5"},
		{"mg/g", "mg/g de creatinina", true, "o laudo só qualifica o denominador"},
		{"k/µL", "mil/mm3", true, "1 mm³ = 1 µL e mil = k"},
		{"k/µL", "x 10³/mm3", true, "idem, com o multiplicador em notação"},
		{"k/µL", "x10³/mm3", true, "idem sem espaço"},
		{"M/µL", "milhões/mm3", true, "milhões = M, mm³ = µL"},
		{"mIU/L", "µUI/mL", true, "1 mIU/L = 1 µIU/mL, exato"},
		{"mm/h", "mm", true, "VHS com o /h omitido no laudo"},
		{"mm/hr", "mm", true, "o laudo escreve hr no lugar de h"},
		{"mm/hr", "mm/h", true, "idem"},
		{"mm/hr", "mm/Hora", true, "e por extenso também"},
		{"mL/min/1.73m²", "mL/min/1,73m2", true, "a mesma superfície corporal, escrita de três jeitos"},
		{"mL/min/1.73m²", "mL/min/1,73 m²", true, "idem"},
		{"µg/dL", "µg/100 mL", true, "1 dL = 100 mL, por definição"},
		{"mIU/mL", "UI/L", true, "1 mIU/mL = 1 IU/L, exato"},
		{"µUI/mL", "µIU/mL", true, "as duas grafias da mesma unidade têm que convergir"},

		// Grandezas diferentes: tem que recusar, é o motivo da guarda existir.
		{"k/µL", "%", false, "contagem absoluta de linfócitos não é o percentual"},
		{"células/campo", "/µL", false, "contagem por campo não é concentração"},
		{"células/campo", "/uL", false, "idem"},
		{"células/campo", "/mL", false, "idem"},
		{"g/dL", "mg/dL", false, "mil vezes de diferença"},
		{"ng/mL", "pg/mL", false, "mil vezes"},
		{"ng/dL", "pg/mL", false, "dez vezes"},
		{"µg/dL", "mg/L", false, "cem vezes"},
		{"mg/dL", "U/L", false, "concentração não é atividade enzimática"},
		{"mg/g", "mg/L", false, "razão por creatinina não é concentração"},
		{"M/µL", "/mm3", false, "milhões por µL não é contagem crua por mm³"},
		{"mEq/L", "mmol/L", false, "só é igual em íon monovalente; o catálogo é corrigido no dado"},
		{"U/mL", "UI/mL", true, "sorologia escreve U e UI para o mesmo ensaio"},
		// Aritmética de prefixo (1 U/mL = 1 kU/L) a normalização não faz de propósito: sai
		// caro em risco e o catálogo já registra esses pares por exame.
		{"U/mL", "kU/L", false, "resolvido pelos sinônimos curados, não pela regra de string"},
	}

	for _, c := range casos {
		si := &ScoreItem{Unit: &c.escala}
		if got := si.UnitMatches(c.resultado, nil); got != c.casa {
			t.Errorf("UnitMatches(%q, %q, nil) = %v, queria %v — %s",
				c.escala, c.resultado, got, c.casa, c.porque)
		}
	}
}

// Unidade ausente de qualquer um dos lados não pode bloquear: item categórico não tem unidade,
// e laudo sem unidade declarada não é motivo para deixar de avaliar.
func TestUnitMatches_VazioNaoBloqueia(t *testing.T) {
	vazio := ""
	gdl := "g/dL"
	if !(&ScoreItem{}).UnitMatches("g/dL", nil) {
		t.Error("item sem unidade deveria passar")
	}
	if !(&ScoreItem{Unit: &vazio}).UnitMatches("g/dL", nil) {
		t.Error("unidade vazia no item deveria passar")
	}
	if !(&ScoreItem{Unit: &gdl}).UnitMatches("", nil) {
		t.Error("exame sem unidade deveria passar")
	}
	if !(&ScoreItem{Unit: &gdl}).UnitMatches("   ", nil) {
		t.Error("unidade só com espaço deveria passar")
	}
}

// A comparação é simétrica: a ordem dos operandos não pode mudar a resposta.
func TestUnitMatches_Simetrica(t *testing.T) {
	pares := [][2]string{{"mIU/L", "µUI/mL"}, {"mm/h", "mm"}, {"k/µL", "mil/mm3"}, {"g/dL", "mg/dL"}}
	for _, p := range pares {
		a, b := p[0], p[1]
		ida := (&ScoreItem{Unit: &a}).UnitMatches(b, nil)
		volta := (&ScoreItem{Unit: &b}).UnitMatches(a, nil)
		if ida != volta {
			t.Errorf("assimétrico em %q/%q: %v vs %v", a, b, ida, volta)
		}
	}
}

// A tabela `lab_test_unit_conversions` guarda equivalências que dependem do analito e que
// nenhuma regra de string pode adivinhar: `mEq/L` só é igual a `mmol/L` em íon monovalente.
// Sem os sinônimos do exame, a guarda recusa — e é isso que protege um divalente.
func TestUnitMatches_SinonimosCuradosDoExame(t *testing.T) {
	meq := "mEq/L"
	sodio := &ScoreItem{Unit: &meq}

	if sodio.UnitMatches("mmol/L", nil) {
		t.Error("sem sinônimo do catálogo, mEq/L não pode virar mmol/L sozinho: em cálcio seria o dobro")
	}
	if !sodio.UnitMatches("mmol/L", [][2]string{{"mEq/L", "mmol/L"}}) {
		t.Error("com o par curado do sódio, tem que aceitar")
	}
	// A ordem em que o catálogo gravou o par não pode importar.
	if !sodio.UnitMatches("mmol/L", [][2]string{{"mmol/L", "mEq/L"}}) {
		t.Error("par curado invertido também tem que aceitar")
	}
	// Sinônimo de outro par não vale de graça.
	if sodio.UnitMatches("mmol/L", [][2]string{{"ng/mL", "µg/L"}}) {
		t.Error("par curado de outro exame não pode liberar mEq/L")
	}
	// Aritmética de prefixo, que a normalização não faz: vem do catálogo (CA-125, IgE).
	uml := "U/mL"
	ca125 := &ScoreItem{Unit: &uml}
	if !ca125.UnitMatches("kU/L", [][2]string{{"U/mL", "kU/L"}}) {
		t.Error("1 U/mL = 1 kU/L está no catálogo e tem que passar")
	}

	// O sinônimo curado também passa pela normalização dos dois lados.
	ngml := "ng/mL"
	ferritina := &ScoreItem{Unit: &ngml}
	if !ferritina.UnitMatches("ug/L", [][2]string{{"ng/mL", "µg/L"}}) {
		t.Error("o par curado tem que casar mesmo com o laudo escrevendo u no lugar de µ")
	}
}
