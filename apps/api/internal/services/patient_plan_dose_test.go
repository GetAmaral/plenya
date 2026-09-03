package services

import "testing"

// A dose vem de texto livre do receituário, e o ponto é ambíguo em PT-BR. A primeira versão
// apagava todo ponto achando que era separador de milhar, e "2.5 mg" entrava no índice como 25:
// a dose CERTA virava alarme falso e a errada passava na conferência.
func TestNumerosEmTexto(t *testing.T) {
	casos := []struct {
		txt       string
		contem    []float64
		naoContem []float64
	}{
		// Ponto sem vírgula é ambíguo, e `leituras` devolve as DUAS leituras de propósito. O que
		// importa aqui é que 2,5 ESTEJA presente: a versão anterior apagava o ponto e indexava só
		// 25, então a dose certa do receituário virava "número não encontrado". Incluir as duas é
		// o lado certo de errar num índice cuja função é não dar alarme falso.
		{"2.5 mg", []float64{2.5, 25}, nil},
		{"2,5 mg", []float64{2.5}, []float64{25}},
		{"50 mg", []float64{50}, nil},
		{"Tomar 1 cápsula de 12/12 horas", []float64{1, 12, 12}, nil},
		{"60 cápsulas", []float64{60}, nil},
		// Ambíguo de verdade: as duas leituras entram, porque não dá para saber qual é.
		{"1.023", []float64{1023, 1.023}, nil},
		{"sem número", nil, nil},
	}
	for _, c := range casos {
		got := numerosEmTexto(c.txt)
		tem := func(v float64) bool {
			for _, g := range got {
				if g == v {
					return true
				}
			}
			return false
		}
		for _, v := range c.contem {
			if !tem(v) {
				t.Errorf("%q: esperava conter %v, veio %v", c.txt, v, got)
			}
		}
		for _, v := range c.naoContem {
			if tem(v) {
				t.Errorf("%q: NÃO podia conter %v, veio %v", c.txt, v, got)
			}
		}
	}
}
