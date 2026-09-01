package models

import (
	"math"
	"testing"
)

func TestFatorEntreUnidades(t *testing.T) {
	casos := []struct {
		de, para string
		fator    float64
		ok       bool
		porque   string
	}{
		// Prefixo no numerador.
		{"pg/mL", "ng/mL", 1e-3, true, "mil pg em um ng"},
		{"ng/mL", "pg/mL", 1e3, true, "e a volta"},
		{"mg/dL", "g/dL", 1e-3, true, "o erro de 10x que estava na linha da Alfa-1 Globulina"},
		{"µg/dL", "mg/L", 1e-2, true, "1 µg/dL = 0,01 mg/L"},

		// Prefixo no denominador.
		{"ng/dL", "ng/mL", 1e-2, true, "um dL tem 100 mL, então o valor por mL é cem vezes menor"},
		{"ng/mL", "ng/dL", 1e2, true, "o erro de 10x que estava na linha do T3 Reverso"},
		{"mg/L", "mg/dL", 1e-1, true, ""},

		// Identidades de definição.
		{"/mm3", "/µL", 1, true, "1 mm³ é 1 µL, exato"},
		{"µg/100 mL", "µg/dL", 1, true, "1 dL é 100 mL"},
		{"mcg/dL", "µg/dL", 1, true, "mc é a grafia ASCII de µ"},
		{"UI/mL", "IU/mL", 1, true, "UI é a grafia portuguesa de IU"},
		{"U/mL", "kU/L", 1, true, "aritmética resolve o que a normalização não faz"},

		// Contagens com multiplicador.
		{"mil/mm3", "k/µL", 1, true, "mil por mm³ é k por µL"},
		{"milhões/mm3", "M/µL", 1, true, ""},
		{"x10³/mm3", "k/µL", 1, true, ""},
		{"/mm3", "k/µL", 1e-3, true, "contagem crua contra milhares"},

		// A caixa do prefixo carrega significado e não pode ser perdida.
		{"M/µL", "m/µL", 1e9, true, "M é mega e m é mili: nove ordens de grandeza"},

		// O que NÃO é aritmética de prefixo e tem que ser recusado.
		{"mEq/L", "mmol/L", 0, false, "depende da valência do íon; é decisão de catálogo"},
		{"mmol/L", "mg/dL", 0, false, "molar para massa exige o peso molecular"},
		{"%", "k/µL", 0, false, "percentual não é contagem"},
		{"células/campo", "/µL", 0, false, "contagem por campo não é concentração"},
		{"specific gravity", "g/mL", 0, false, "adimensional sem prefixo; precisa de linha curada"},
		{"mmHg", "kPa", 0, false, "pressão não está na tabela de bases"},
		{"mg/dL", "U/L", 0, false, "concentração não é atividade enzimática"},
		{"mg", "mg/dL", 0, false, "massa solta não se compara com concentração"},
	}

	for _, c := range casos {
		f, ok := FatorEntreUnidades(c.de, c.para)
		if ok != c.ok {
			t.Errorf("FatorEntreUnidades(%q,%q) ok=%v, queria %v — %s", c.de, c.para, ok, c.ok, c.porque)
			continue
		}
		if !c.ok {
			continue
		}
		if math.Abs(f-c.fator) > 1e-9*math.Max(1, math.Abs(c.fator)) {
			t.Errorf("FatorEntreUnidades(%q,%q) = %g, queria %g — %s", c.de, c.para, f, c.fator, c.porque)
		}
	}
}

// Ida e volta tem que fechar: o produto dos dois fatores é 1.
func TestFatorEntreUnidades_IdaEVolta(t *testing.T) {
	pares := [][2]string{{"pg/mL", "ng/mL"}, {"mg/dL", "g/dL"}, {"ng/dL", "ng/mL"}, {"/mm3", "k/µL"}, {"U/mL", "kU/L"}}
	for _, p := range pares {
		ida, ok1 := FatorEntreUnidades(p[0], p[1])
		volta, ok2 := FatorEntreUnidades(p[1], p[0])
		if !ok1 || !ok2 {
			t.Fatalf("%v: esperava parsear os dois sentidos", p)
		}
		if math.Abs(ida*volta-1) > 1e-9 {
			t.Errorf("%v: ida %g e volta %g não fecham em 1", p, ida, volta)
		}
	}
}
