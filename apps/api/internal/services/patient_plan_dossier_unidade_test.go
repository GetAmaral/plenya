package services

import (
	"testing"

	"github.com/plenya/api/internal/models"
)

// A guarda da régua tem de comparar a unidade do RESULTADO com a escala do item.
//
// Antes ela comparava a unidade do CATÁLOGO, que é a mesma do item de escore por construção, então
// nunca reprovava. O caso real: lipoproteína(a) com escala em nmol/L, laudo brasileiro em mg/dL,
// conversor marcando "revisar" e motor do escore recusando classificar — e a régua desenhando
// "24,1 nmol/L" sobre um valor que ninguém converteu.
func TestUnidadeDoResultadoPrefereALaudada(t *testing.T) {
	casos := []struct {
		nome     string
		linha    labRow
		esperado string
	}{
		{"laudo em mg/dL contra catálogo em nmol/L", labRow{Unit: "mg/dL", DefUnit: "nmol/L"}, "mg/dL"},
		{"laudo sem unidade cai para a do catálogo", labRow{Unit: "", DefUnit: "nmol/L"}, "nmol/L"},
		{"espaço em branco conta como ausente", labRow{Unit: "   ", DefUnit: "ng/mL"}, "ng/mL"},
		{"sem nenhuma das duas não bloqueia", labRow{}, ""},
	}
	for _, c := range casos {
		if got := unidadeDoResultado(c.linha); got != c.esperado {
			t.Errorf("%s: esperava %q, veio %q", c.nome, c.esperado, got)
		}
	}
}

// E o efeito prático: com a unidade do laudo, MesmaGrandeza reprova o par que o motor do escore já
// reprova, então a régua deixa de existir em vez de existir errada.
func TestGrandezaDiferenteReprovaComAUnidadeDoLaudo(t *testing.T) {
	linha := labRow{Unit: "mg/dL", DefUnit: "nmol/L"}
	if models.MesmaGrandeza("nmol/L", unidadeDoResultado(linha), nil) {
		t.Error("mg/dL e nmol/L não são a mesma grandeza; a régua não deveria ser desenhada")
	}
	// O caminho que continua valendo: laudo e escala na mesma unidade.
	if !models.MesmaGrandeza("nmol/L", unidadeDoResultado(labRow{Unit: "nmol/L", DefUnit: "nmol/L"}), nil) {
		t.Error("mesma unidade tem de passar")
	}
}
