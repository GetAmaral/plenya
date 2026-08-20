package services

import (
	"strings"
	"testing"
)

// As fórmulas do formulário escrevem o nome com parêntese e com alternativa. Todas essas
// substâncias existem no catálogo, e nenhuma casava antes — quatro fórmulas ficavam sem cálculo
// de cápsula só por causa da forma de escrever.
func TestNameCandidates(t *testing.T) {
	casos := []struct {
		nome    string
		esperar string
	}{
		{"Vitamina B6 (piridoxal-5-fosfato)", "piridoxal-5-fosfato"},
		{"Vitamina B2 (riboflavina)", "riboflavina"},
		{"PQQ (pirroloquinolina quinona)", "PQQ"},
		{"Coenzima Q10 (ubiquinona) ou ubiquinol", "Coenzima Q10"},
		{"Curcumina padronizada (95% curcuminoides)", "Curcumina"},
		{"Alfa-GPC (L-alfa-glicerofosfocolina)", "Alfa-GPC"},
	}
	for _, c := range casos {
		got := nameCandidates(c.nome)
		achou := false
		for _, g := range got {
			if strings.EqualFold(strings.TrimSpace(g), c.esperar) {
				achou = true
			}
		}
		if !achou {
			t.Errorf("%q devia gerar o candidato %q, gerou %v", c.nome, c.esperar, got)
		}
		if !strings.EqualFold(got[0], c.nome) {
			t.Errorf("o nome escrito tem que ser a primeira tentativa, veio %q", got[0])
		}
	}
}

// Qualificador que MUDA o insumo não pode ser descartado: "betaína anidra" não é "betaína", e
// "magnésio quelato" não é "magnésio" — densidade e fator de correção mudam junto.
func TestNameCandidatesNaoDescartaOQueMudaOInsumo(t *testing.T) {
	for _, nome := range []string{"Betaína anidra", "Magnésio quelato", "Magnésio L-treonato"} {
		if got := nameCandidates(nome); len(got) != 1 {
			t.Errorf("%q não devia gerar variante, gerou %v", nome, got)
		}
	}
}
