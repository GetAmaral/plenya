package services

import (
	"math"
	"regexp"
	"strings"
)

// DosesPorDia lê da posologia quantas tomadas o dia tem.
//
// Existe porque tudo neste módulo raciocina em dose DIÁRIA — a faixa do catálogo, o teto da IN 28,
// a regra de dose dinâmica — enquanto a fórmula é escrita por cápsula. Sem essa conta, uma regra
// que manda 5.000 UI/dia numa fórmula tomada duas vezes ao dia entrega 10.000.
//
// Só conta o que está escrito de forma inequívoca e cai para 1 no resto: o número aparece na frase
// que o médico lê, então erro aqui fica visível em vez de silencioso. Quando a posologia dá um
// intervalo ("1 a 2 vezes ao dia"), vale o maior — para a faixa de dose, errar para mais vira
// "acima da faixa, confirme", que pede conferência, em vez de um "abaixo da faixa" falso.
func DosesPorDia(posologia string) float64 {
	t := strings.ToLower(strings.TrimSpace(posologia))
	if t == "" {
		return 1
	}

	if m := reVezesAoDia.FindAllStringSubmatch(t, -1); len(m) > 0 {
		maior := 0.0
		for _, g := range m {
			v := paraNumero(g[2])
			if v == 0 {
				v = paraNumero(g[1])
			}
			if v > maior && v <= 12 {
				maior = v
			}
		}
		if maior > 0 {
			return maior
		}
	}

	if m := reCadaHoras.FindStringSubmatch(t); m != nil {
		h := paraNumero(m[1])
		if h == 0 {
			h = paraNumero(m[2])
		}
		if h > 0 && h <= 24 {
			return math.Max(1, math.Round(24/h))
		}
	}

	// Sem número: conta as ocasiões nomeadas ("1 dose após o almoço e 1 após o jantar" = 2).
	ocasioes := []string{"manhã", "manha", "almoço", "almoco", "tarde", "jantar", "noite", "deitar"}
	achadas := map[string]bool{}
	for _, o := range ocasioes {
		if strings.Contains(t, o) {
			achadas[o] = true
		}
	}
	// "noite" e "deitar" são a mesma ocasião.
	if achadas["noite"] && achadas["deitar"] {
		delete(achadas, "deitar")
	}
	if n := float64(len(achadas)); n >= 1 && n <= 6 {
		return n
	}
	return 1
}

var (
	reVezesAoDia = regexp.MustCompile(`(\d+)\s*(?:a\s*(\d+)\s*)?(?:vezes|vez|x)\s*(?:ao|por|/)\s*dia`)
	reCadaHoras  = regexp.MustCompile(`(?:de\s*(\d+)\s*em\s*\d+|a\s*cada\s*(\d+))\s*horas?`)
)

func paraNumero(s string) float64 {
	if s == "" {
		return 0
	}
	var v float64
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0
		}
		v = v*10 + float64(r-'0')
	}
	return v
}
