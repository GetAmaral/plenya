package models

import (
	"strings"
)

// Conversão de unidade por aritmética de prefixo.
//
// A tabela `lab_test_unit_conversions` é curada à mão e cobre 145 pares. O laboratório manda
// muito mais do que isso, e a maioria do que falta não é decisão clínica nenhuma: `pg/mL` para
// `ng/mL` é dividir por mil. Quando o par não está na tabela, o conversor desistia em silêncio e
// o resultado ficava gravado na unidade do laudo — para depois ser comparado, lá na frente,
// contra uma escala em outra unidade.
//
// O que ESTA aritmética faz é só o que é verdade por definição de prefixo SI. O que depende do
// analito continua na tabela curada: `mEq/L` só é igual a `mmol/L` em íon monovalente, e por isso
// `Eq` e `mol` são bases DIFERENTES aqui — a aritmética se recusa a relacionar as duas.

// prefixosSI é sensível a maiúscula de propósito: `M` é mega e `m` é mili. Num hemograma isso é a
// diferença entre 4,5 M/µL (hemácias) e 4,5 m/µL, que não existe.
var prefixosSI = map[string]float64{
	"T": 1e12, "G": 1e9, "M": 1e6, "k": 1e3, "h": 1e2,
	"":  1,
	"d": 1e-1, "c": 1e-2, "m": 1e-3, "u": 1e-6, "n": 1e-9, "p": 1e-12, "f": 1e-15,
}

// basesConhecidas mapeia a grafia da base para uma chave canônica. Duas unidades só se comparam
// se a base bater dos dois lados. A chave vazia é contagem pura (`/µL`, `k/µL`, `M/µL`).
var basesConhecidas = map[string]string{
	"":        "contagem",
	"g":       "massa",
	"l":       "volume",
	"mol":     "mol",
	"eq":      "eq",
	"u":       "ui",
	"iu":      "ui",
	"ui":      "ui",
	"cel":     "contagem",
	"cels":    "contagem",
	"células": "contagem",
	"celulas": "contagem",
}

// preparaParaAritmetica arruma a grafia SEM mexer na caixa, que aqui carrega significado.
func preparaParaAritmetica(u string) string {
	s := strings.TrimSpace(u)
	if i := strings.Index(s, " de "); i > 0 {
		s = s[:i]
	}
	s = strings.ReplaceAll(s, " ", "")

	// Micro em todas as grafias vira o prefixo `u`.
	s = strings.ReplaceAll(s, "µ", "u")
	s = strings.ReplaceAll(s, "μ", "u")
	s = strings.ReplaceAll(s, "mcg", "ug")
	s = strings.ReplaceAll(s, "mcmol", "umol")

	// 1 mm³ = 1 µL e 1 dL = 100 mL, por definição.
	s = strings.ReplaceAll(s, "mm³", "uL")
	s = strings.ReplaceAll(s, "mm3", "uL")
	s = strings.ReplaceAll(s, "/100mL", "/dL")
	s = strings.ReplaceAll(s, "/100ml", "/dL")

	// Multiplicadores escritos por extenso, como o hemograma faz.
	for _, mil := range []string{"x10³", "x10^3", "10³", "mil"} {
		s = strings.ReplaceAll(s, mil+"/", "k/")
	}
	for _, milhao := range []string{"milhões", "milhoes"} {
		s = strings.ReplaceAll(s, milhao+"/", "M/")
	}
	return s
}

// separaPrefixo quebra `mg` em (1e-3, "massa"). Devolve ok=false se não reconhecer a base, que é
// o caso de `%`, `células/campo`, `Descritivo`, `specific gravity` e afins: essas não são
// grandezas com prefixo, e inventar aritmética em cima delas é o erro que se quer evitar.
func separaPrefixo(tok string) (float64, string, bool) {
	if base, ok := basesConhecidas[strings.ToLower(tok)]; ok {
		return 1, base, true
	}
	for p, mult := range prefixosSI {
		if p == "" || !strings.HasPrefix(tok, p) {
			continue
		}
		resto := tok[len(p):]
		if base, ok := basesConhecidas[strings.ToLower(resto)]; ok {
			return mult, base, true
		}
	}
	return 0, "", false
}

// FatorEntreUnidades devolve por quanto multiplicar um valor em `de` para obtê-lo em `para`.
// ok=false quando as unidades não são relacionáveis por prefixo — aí a decisão não é aritmética.
func FatorEntreUnidades(de, para string) (float64, bool) {
	numDe, denDe, okDe := parteAParte(de)
	numPara, denPara, okPara := parteAParte(para)
	if !okDe || !okPara {
		return 0, false
	}
	if numDe.base != numPara.base || denDe.base != denPara.base {
		return 0, false
	}
	return (numDe.mult / numPara.mult) * (denPara.mult / denDe.mult), true
}

type parteDeUnidade struct {
	mult float64
	base string
}

// parteAParte separa numerador e denominador. Sem barra, o denominador é neutro — assim `mg` e
// `g` se comparam, e `mg/dL` nunca se compara com `mg` solto.
func parteAParte(u string) (num, den parteDeUnidade, ok bool) {
	s := preparaParaAritmetica(u)
	if s == "" {
		return num, den, false
	}

	numTok, denTok := s, ""
	if i := strings.Index(s, "/"); i >= 0 {
		numTok, denTok = s[:i], s[i+1:]
	}

	m, b, ok := separaPrefixo(numTok)
	if !ok {
		return num, den, false
	}
	num = parteDeUnidade{mult: m, base: b}

	if denTok == "" {
		return num, parteDeUnidade{mult: 1, base: "adimensional"}, true
	}
	m, b, ok = separaPrefixo(denTok)
	if !ok {
		return num, den, false
	}
	return num, parteDeUnidade{mult: m, base: b}, true
}
