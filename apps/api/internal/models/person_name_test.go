package models

import "testing"

func TestNormalizePersonName(t *testing.T) {
	casos := map[string]string{
		// o que a recepção digita de verdade
		"JOÃO DA SILVA":     "João da Silva",
		"joão da silva":     "João da Silva",
		"JoÃo Da SiLvA":     "João da Silva",
		"  maria   clara  ": "Maria Clara",
		"joão":              "João",
		"":                  "",
		"   ":               "",

		// acento e cedilha na caixa certa
		"JOSÉ GONÇALVES DE ASSUNÇÃO": "José Gonçalves de Assunção",
		"ÂNGELA MÜLLER":              "Ângela Müller",

		// partículas
		"maria dos santos e silva": "Maria dos Santos e Silva",
		"pedro van der berg":       "Pedro van der Berg",
		// partícula que abre o nome continua maiúscula
		"DA SILVA JUNIOR": "Da Silva Junior",

		// sobrenomes curtos que a regra de "1 ou 2 letras" estragaria
		"maria de sá": "Maria de Sá",
		"ana pó":      "Ana Pó",
		"li wei":      "Li Wei",

		// inicial, hífen e apóstrofo
		"joão p. da silva":  "João P. da Silva",
		"ana-maria d'ávila": "Ana-Maria D'Ávila",
		"o'brien mcdonald":  "O'Brien Mcdonald",

		// e-mail no campo do nome acontece, e o ponto não pode virar fronteira de palavra
		"getfilho@yahoo.com.br": "Getfilho@yahoo.com.br",
		"j.p. morgan":           "J.P. Morgan",

		// sufixo dinástico
		"joão paulo ii": "João Paulo II",
		// acima de XII a lista para: sufixo dinástico nesse tamanho não existe em prontuário,
		// e reconhecer "XVI" abriria a porta para tratar sobrenome como numeral.
		"luiz xvi": "Luiz Xvi",
	}
	for entrada, esperado := range casos {
		if got := NormalizePersonName(entrada); got != esperado {
			t.Errorf("NormalizePersonName(%q) = %q, esperado %q", entrada, got, esperado)
		}
	}
}

// Normalizar duas vezes tem que dar o mesmo resultado: o hook roda em todo save, e nome já
// arrumado não pode mudar de novo.
func TestNormalizePersonNameIdempotente(t *testing.T) {
	for _, n := range []string{"João da Silva", "Maria de Sá", "Ana-Maria D'Ávila", "João Paulo II"} {
		if got := NormalizePersonName(NormalizePersonName(n)); got != n {
			t.Errorf("segunda passada mudou %q para %q", n, got)
		}
	}
}
