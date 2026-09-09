package services

import "testing"

func compRazao(v float64, unidade string) componenteDoLote {
	return componenteDoLote{Valor: v, Unid: unidade}
}

// O caso comum: colesterol total e HDL do mesmo lipidograma, e apoB e apoA1 do mesmo lote.
func TestRazaoDe_DivideQuandoAsUnidadesBatem(t *testing.T) {
	for _, c := range []struct {
		nome     string
		num, den componenteDoLote
		quer     float64
	}{
		// Duas casas, como o laudo imprime e como a régua mostra.
		{"colesterol total sobre HDL", compRazao(234, "mg/dL"), compRazao(89, "mg/dL"), 2.63},
		{"apoB sobre apoA1", compRazao(86, "mg/dL"), compRazao(193, "mg/dL"), 0.45},
		// A razão é adimensional, então o par inteiro em g/L dá exatamente o mesmo número.
		{"o par inteiro em g/L", compRazao(0.86, "g/L"), compRazao(1.93, "g/L"), 0.45},
		// A normalização de unidade vem de `mesmaUnidade`, que trata µ e mc como a mesma coisa.
		{"mesma unidade escrita diferente", compRazao(10, "µg/dL"), compRazao(5, "mcg/dL"), 2},
	} {
		got, ok := razaoDe(c.num, c.den)
		if !ok {
			t.Errorf("%s: não calculou", c.nome)
			continue
		}
		if got < c.quer-1e-9 || got > c.quer+1e-9 {
			t.Errorf("%s: esperava %.4f, veio %.4f", c.nome, c.quer, got)
		}
	}
}

// A guarda que evita o número plausível e falso: a apolipoproteína B é reportada em g/L por muitos
// laboratórios, e 0,86 g/L sobre 193 mg/dL dá 0,0045, que cai no melhor nível e entrega 18 pontos
// sem que nada acuse. O inverso, 86 mg/dL sobre 1,93 g/L, dá 44,5 e custa os 18.
func TestRazaoDe_RecusaUnidadesDiferentes(t *testing.T) {
	for _, c := range []struct {
		nome     string
		num, den componenteDoLote
	}{
		{"apoB em g/L sobre apoA1 em mg/dL", compRazao(0.86, "g/L"), compRazao(193, "mg/dL")},
		{"o inverso", compRazao(86, "mg/dL"), compRazao(1.93, "g/L")},
		{"unidade ausente no numerador", compRazao(86, ""), compRazao(193, "mg/dL")},
		{"unidade ausente nos dois", compRazao(86, ""), compRazao(193, "")},
	} {
		if v, ok := razaoDe(c.num, c.den); ok {
			t.Errorf("%s: dividiu assim mesmo e produziu %.4f", c.nome, v)
		}
	}
}

func TestRazaoDe_RecusaDenominadorZero(t *testing.T) {
	if v, ok := razaoDe(compRazao(234, "mg/dL"), compRazao(0, "mg/dL")); ok {
		t.Errorf("dividiu por zero e produziu %v", v)
	}
}

// A tabela é o contrato com o catálogo: código de razão que não existe em `lab_test_definitions`
// nunca é gravado, e componente trocado produz a conta errada em todo paciente. Vale conferir que
// ninguém mexeu nela sem querer.
func TestRazoesDerivadas_TabelaEstaCoerente(t *testing.T) {
	vistos := map[string]bool{}
	for _, r := range razoesDerivadas {
		if r.Codigo == "" || r.Numerador == "" || r.Denominador == "" {
			t.Errorf("%+v: código, numerador e denominador são obrigatórios", r)
		}
		if r.Numerador == r.Denominador {
			t.Errorf("%s: numerador e denominador iguais dariam sempre 1", r.Codigo)
		}
		if r.Unidade != "ratio" {
			t.Errorf("%s: unidade %q, e os itens de escore de razão estão em ratio", r.Codigo, r.Unidade)
		}
		if vistos[r.Codigo] {
			t.Errorf("%s: aparece duas vezes, e a segunda sobrescreveria a primeira", r.Codigo)
		}
		vistos[r.Codigo] = true
	}
}
