package services

import "strings"

var extUnidades = []string{"zero", "um", "dois", "três", "quatro", "cinco", "seis", "sete", "oito", "nove",
	"dez", "onze", "doze", "treze", "quatorze", "quinze", "dezesseis", "dezessete", "dezoito", "dezenove"}
var extDezenas = []string{"", "", "vinte", "trinta", "quarenta", "cinquenta", "sessenta", "setenta", "oitenta", "noventa"}
var extCentenas = []string{"", "cento", "duzentos", "trezentos", "quatrocentos", "quinhentos", "seiscentos", "setecentos", "oitocentos", "novecentos"}

func extAte999(n int) string {
	if n == 0 {
		return ""
	}
	if n == 100 {
		return "cem"
	}
	var p []string
	if c := n / 100; c > 0 {
		p = append(p, extCentenas[c])
	}
	if r := n % 100; r > 0 {
		if r < 20 {
			p = append(p, extUnidades[r])
		} else if u := r % 10; u == 0 {
			p = append(p, extDezenas[r/10])
		} else {
			p = append(p, extDezenas[r/10]+" e "+extUnidades[u])
		}
	}
	return strings.Join(p, " e ")
}

// numeroExtenso — inteiro por extenso em pt-BR (0..999.999; suficiente para recibos).
func numeroExtenso(n int) string {
	if n == 0 {
		return "zero"
	}
	if n < 1000 {
		return extAte999(n)
	}
	milhares, resto := n/1000, n%1000
	head := "mil"
	if milhares > 1 {
		head = extAte999(milhares) + " mil"
	}
	if resto == 0 {
		return head
	}
	sep := " e "
	if resto >= 100 && resto%100 != 0 {
		sep = " "
	}
	return head + sep + extAte999(resto)
}

// valorPorExtenso — "R$ 850,00" => "oitocentos e cinquenta reais".
func valorPorExtenso(cents int64) string {
	reais, cent := int(cents/100), int(cents%100)
	var p []string
	if reais > 0 {
		u := "reais"
		if reais == 1 {
			u = "real"
		}
		p = append(p, numeroExtenso(reais)+" "+u)
	}
	if cent > 0 {
		u := "centavos"
		if cent == 1 {
			u = "centavo"
		}
		p = append(p, numeroExtenso(cent)+" "+u)
	}
	if len(p) == 0 {
		return "zero real"
	}
	return strings.Join(p, " e ")
}
