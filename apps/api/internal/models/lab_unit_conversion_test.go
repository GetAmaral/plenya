package models

import (
	"strings"
	"testing"
)

func exame(unidade string) *LabTestDefinition { return &LabTestDefinition{Unit: &unidade} }

// A rede de plausibilidade existe contra RÓTULO errado, não contra paciente doente. Aritmética
// correta em cima de rótulo errado produz número confiantemente errado, que é pior que recusar.
func TestConverte_PlausibilidadeBloqueiaRotuloErrado(t *testing.T) {
	// Hemácias: o laudo diz `/mm3` mas o valor 4,17 está claramente em M/µL. A aritmética
	// converteria para 4,17 milionésimos e o paciente sairia com anemia catastrófica.
	hemacias := exame("M/µL")
	faixa := func(v float64) bool { return v >= 0.35 && v <= 60 } // faixa do escore, com folga

	got := hemacias.ConverteParaUnidadePrincipal(nil, 4.17, "/mm3", faixa)

	if got.Status != ConversaoPendente {
		t.Fatalf("status = %q, queria revisar (a conversão daria %v)", got.Status, got.Valor)
	}
	if got.Valor != 4.17 || got.Unidade != "/mm3" {
		t.Errorf("valor mexido: %v %s — deveria ficar exatamente como veio", got.Valor, got.Unidade)
	}
	if !strings.Contains(got.Motivo, "rótulo") {
		t.Errorf("motivo %q não explica que o problema é o rótulo", got.Motivo)
	}
}

// E não pode bloquear conversão legítima: um valor ruim mas possível tem que converter.
func TestConverte_PlausibilidadeDeixaPassarValorRuimPorémPossivel(t *testing.T) {
	dht := exame("ng/dL")
	faixa := func(v float64) bool { return v >= 3 && v <= 850 }

	got := dht.ConverteParaUnidadePrincipal(nil, 400, "pg/mL", faixa)

	if got.Status != ConversaoAplicada {
		t.Fatalf("status = %q (%s), queria convertido", got.Status, got.Motivo)
	}
	if got.Valor != 40 || got.Unidade != "ng/dL" {
		t.Errorf("= %v %s, queria 40 ng/dL", got.Valor, got.Unidade)
	}
}

// Sem faixa cadastrada não há como julgar plausibilidade, e aritmética correta é aritmética
// correta: converte.
func TestConverte_SemFaixaConverteAssimMesmo(t *testing.T) {
	got := exame("ng/mL").ConverteParaUnidadePrincipal(nil, 2000, "pg/mL", nil)
	if got.Status != ConversaoAplicada || got.Valor != 2 {
		t.Fatalf("= %v %s (%s), queria 2 ng/mL convertido", got.Valor, got.Unidade, got.Status)
	}
}

// Grafia diferente da mesma unidade não é conversão nenhuma, e o VHS é o caso que já passava na
// guarda do escore e era recusado aqui antes de a lógica ser compartilhada.
func TestConverte_MesmaUnidadeEscritaDeOutroJeito(t *testing.T) {
	for _, c := range []struct{ exameUnidade, laudo string }{
		{"µg/dL", "mcg/dL"},
		{"IU/mL", "UI/mL"},
		{"mm/hr", "mm"},
		{"mIU/mL", "mUI/mL"},
	} {
		got := exame(c.exameUnidade).ConverteParaUnidadePrincipal(nil, 7, c.laudo, nil)
		if got.Status != ConversaoDesnecessaria {
			t.Errorf("%s vs %s: status = %q (%s), queria ok", c.exameUnidade, c.laudo, got.Status, got.Motivo)
		}
		if got.Valor != 7 {
			t.Errorf("%s vs %s: valor virou %v; não havia nada para converter", c.exameUnidade, c.laudo, got.Valor)
		}
	}
}

// Grandeza diferente sem regra: mantém o valor e AVISA. Antes desistia calado.
func TestConverte_GrandezaDiferenteAvisa(t *testing.T) {
	got := exame("k/µL").ConverteParaUnidadePrincipal(nil, 43.5, "%", nil)

	if got.Status != ConversaoPendente {
		t.Fatalf("status = %q, queria revisar", got.Status)
	}
	if got.Valor != 43.5 || got.Unidade != "%" {
		t.Errorf("valor mexido: %v %s", got.Valor, got.Unidade)
	}
	if got.Motivo == "" {
		t.Error("ficou sem motivo: é justamente o silêncio que causou o problema")
	}
}

// Unidade ausente de um dos lados não é problema a resolver aqui.
func TestConverte_SemUnidade(t *testing.T) {
	if got := exame("").ConverteParaUnidadePrincipal(nil, 5, "mg/dL", nil); got.Status != ConversaoDesnecessaria {
		t.Errorf("exame sem unidade: status = %q", got.Status)
	}
	if got := exame("mg/dL").ConverteParaUnidadePrincipal(nil, 5, "", nil); got.Status != ConversaoDesnecessaria {
		t.Errorf("laudo sem unidade: status = %q", got.Status)
	}
}
