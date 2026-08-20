package services

import "strings"

// ConverteDose leva uma quantidade da unidade escrita na fórmula para a unidade do catálogo.
//
// Existe porque três conferências deste módulo — faixa de dose, teto da IN 28 e interferência em
// exame — comparam número da receita com número do catálogo, e a receita pode estar escrita em
// outra unidade da mesma família. Sem conversão, "biotina 500 mcg" contra um limiar de 5 mg
// disparava alarme (500 > 5) numa dose rotineira, e "biotina 10 mg" contra teto de 45 mcg passava
// batido. Os dois erros no mesmo lugar, em direções opostas.
//
// Só converte dentro da mesma família de massa. Fora disso devolve false, e quem chama pula a
// conferência em vez de comparar número com número de grandezas diferentes.
func ConverteDose(quantidade float64, de, para string) (float64, bool) {
	fatorDe, ok1 := fatorParaMcg[normalizaUnidadeDose(de)]
	fatorPara, ok2 := fatorParaMcg[normalizaUnidadeDose(para)]
	if !ok1 || !ok2 || fatorPara == 0 {
		return 0, false
	}
	return quantidade * fatorDe / fatorPara, true
}

// mesmaUnidade responde se as duas unidades são a mesma coisa escrita de jeitos diferentes.
func mesmaUnidade(a, b string) bool {
	na, nb := normalizaUnidadeDose(a), normalizaUnidadeDose(b)
	return na != "" && na == nb
}

func normalizaUnidadeDose(u string) string {
	s := strings.ToLower(strings.TrimSpace(u))
	s = strings.NewReplacer("µ", "u", "μ", "u", "mcg", "ug").Replace(s)
	return s
}

// Só massa: UI e porcentagem não convertem para miligrama sem saber de qual substância se trata.
var fatorParaMcg = map[string]float64{
	"ug": 1,
	"mg": 1000,
	"g":  1_000_000,
}
