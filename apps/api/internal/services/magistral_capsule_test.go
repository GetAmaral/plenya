package services

import (
	"strings"
	"testing"
)

// Teste de mesa da calculadora. É a função deste módulo cujo erro tem consequência clínica mais
// direta, então os casos são conferidos na mão:
//
//	500 mg a 0,5 g/mL  = 1,0 mL de pó → com +25% = 1,25 mL → cabe em 000 (1,37 mL), 1 cápsula.
//	100 mg a 0,5 g/mL  = 0,2 mL       → com +25% = 0,25 mL → cabe em 4 (0,21 mL)? não; em 3 (0,30).
func TestCalculateCapsule(t *testing.T) {
	t.Run("sem densidade a calculadora se cala", func(t *testing.T) {
		got := CalculateCapsule([]CapsuleInput{
			{Substance: "Melatonina", MassMg: 3, BulkDensity: 0.5},
			{Substance: "Magnésio dimalato", MassMg: 300},
		})
		if got.Decided {
			t.Fatal("não deveria opinar sem densidade de todos os componentes")
		}
		if len(got.Missing) != 1 || got.Missing[0] != "Magnésio dimalato" {
			t.Errorf("deveria apontar quem falta, veio %v", got.Missing)
		}
		if !strings.Contains(got.Explanation, "Magnésio dimalato") {
			t.Errorf("a explicação precisa nomear quem falta: %q", got.Explanation)
		}
		if got.Size != "" || got.CapsulesPerDose != 0 {
			t.Error("não pode sugerir tamanho sem dado")
		}
	})

	t.Run("dose pequena cabe em cápsula pequena", func(t *testing.T) {
		got := CalculateCapsule([]CapsuleInput{{Substance: "X", MassMg: 100, BulkDensity: 0.5}})
		if !got.Decided {
			t.Fatal("deveria decidir")
		}
		if got.VolumeML != 0.2 {
			t.Errorf("volume esperado 0,2 mL, veio %v", got.VolumeML)
		}
		if got.Size != "3" || got.CapsulesPerDose != 1 {
			t.Errorf("esperava 1 cápsula tamanho 3, veio %d tamanho %s", got.CapsulesPerDose, got.Size)
		}
		if got.SachetRecommended {
			t.Error("não é caso de sachê")
		}
		if !strings.Contains(got.Explanation, "incerteza") {
			t.Errorf("a incerteza precisa estar dita no texto: %q", got.Explanation)
		}
	})

	t.Run("volume grande vira sachê", func(t *testing.T) {
		// 4 g a 0,5 g/mL = 8 mL: 000 comporta 1,37 mL, precisaria de 6 cápsulas por dose.
		got := CalculateCapsule([]CapsuleInput{{Substance: "Creatina", MassMg: 4000, BulkDensity: 0.5}})
		if !got.SachetRecommended {
			t.Errorf("8 mL por dose deveria recomendar sachê, veio tamanho %s × %d", got.Size, got.CapsulesPerDose)
		}
		if !strings.Contains(got.Explanation, "sachê") {
			t.Errorf("explicação deveria recomendar sachê: %q", got.Explanation)
		}
	})

	t.Run("soma de componentes", func(t *testing.T) {
		got := CalculateCapsule([]CapsuleInput{
			{Substance: "A", MassMg: 250, BulkDensity: 0.5}, // 0,5 mL
			{Substance: "B", MassMg: 250, BulkDensity: 0.5}, // 0,5 mL
		})
		if got.VolumeML != 1 {
			t.Errorf("volume somado deveria ser 1 mL, veio %v", got.VolumeML)
		}
		// 1 mL +25% = 1,25 mL → 000 (1,37) em 1 cápsula
		if got.Size != "000" || got.CapsulesPerDose != 1 {
			t.Errorf("esperava 1 cápsula 000, veio %d tamanho %s", got.CapsulesPerDose, got.Size)
		}
	})

	t.Run("faixa aparece quando o cenário compactado muda o tamanho", func(t *testing.T) {
		// volume 0,6 mL: +25% = 0,75 mL → 00 (0,95); -25% = 0,45 mL → 1 (0,50)
		got := CalculateCapsule([]CapsuleInput{{Substance: "A", MassMg: 300, BulkDensity: 0.5}})
		if got.Size != "00" {
			t.Errorf("cenário conservador deveria mandar (00), veio %s", got.Size)
		}
		if got.SizeIfCompacted != "1" {
			t.Errorf("cenário compactado deveria sugerir 1, veio %q", got.SizeIfCompacted)
		}
	})
}

func TestMassToMg(t *testing.T) {
	cases := []struct {
		qty  float64
		unit string
		want float64
		ok   bool
	}{
		{300, "mg", 300, true},
		{1.5, "g", 1500, true},
		{500, "mcg", 0.5, true},
		{1000, "UI", 0, false},
		{10, "%", 0, false},
		{5, "mL", 0, false},
	}
	for _, c := range cases {
		got, ok := MassToMg(c.qty, c.unit)
		if ok != c.ok || got != c.want {
			t.Errorf("%v %s → (%v, %v), esperado (%v, %v)", c.qty, c.unit, got, ok, c.want, c.ok)
		}
	}
}

// Densidade aproximada por classe não muda o cálculo, mas o texto tem que dizer que é estimativa
// por classe — senão o médico lê ordem de grandeza como se fosse medida do insumo.
func TestCalculateCapsuleAproximada(t *testing.T) {
	got := CalculateCapsule([]CapsuleInput{
		{Substance: "Magnésio quelato", MassMg: 300, BulkDensity: 0.65, DensityApprox: true},
		{Substance: "Creatina", MassMg: 100, BulkDensity: 0.55},
	})
	if !got.Decided || !got.Approximate {
		t.Fatalf("deveria decidir e marcar como aproximada: %+v", got)
	}
	if !strings.Contains(got.Explanation, "aproximada por classe") {
		t.Errorf("o texto precisa declarar a aproximação: %q", got.Explanation)
	}

	medida := CalculateCapsule([]CapsuleInput{{Substance: "Creatina", MassMg: 500, BulkDensity: 0.55}})
	if medida.Approximate || strings.Contains(medida.Explanation, "aproximada por classe") {
		t.Error("sem componente aproximado, o texto não deve falar em aproximação")
	}
}
