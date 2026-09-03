package services

import (
	"testing"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/pdfdoc"
)

// A régua do deck precisa sair com a escala e o histórico do dossiê. O modelo só escolhe o exame e
// escreve o rótulo; se a hidratação falhar, o slide desenha uma barra vazia ao lado do nome de um
// exame real, que é pior que não desenhar nada.
func TestHidrataReguas(t *testing.T) {
	dossie := &dto.PlanDossierResponse{
		Rulers: map[string]dto.PlanRuler{
			"PLNCREA": {
				Code: "PLNCREA", Name: "Creatinina", Unit: "mg/dL",
				Axis:     []float64{0.17, 190.4},
				Segments: []dto.PlanRulerSegment{{Level: 5, A: 0.6, B: 1.1}, {Level: 2, A: 1.1, B: 190.4}},
				History:  []dto.PlanRulerPoint{{Date: "2026-02-07", Value: 0.96, Plot: 0.96, Text: "0,96"}},
			},
		},
	}
	slides := []pdfdoc.DeckSlide{{
		Kind: pdfdoc.DeckRulers,
		Rulers: []pdfdoc.DeckRulerBlock{
			{Ruler: pdfdoc.Ruler{Code: "PLNCREA", Display: "Creatinina", Sub: "o rim filtrando"}},
			// O modelo erra o código com frequência e acerta o nome: o fallback tem que pegar.
			{Ruler: pdfdoc.Ruler{Code: "ERRADO", Display: "Creatinina"}},
		},
	}}
	out, avisos := hidrataReguas(slides, dossie)
	if len(out[0].Rulers) != 2 {
		t.Fatalf("esperava 2 réguas, veio %d (avisos: %v)", len(out[0].Rulers), avisos)
	}
	for i, b := range out[0].Rulers {
		if len(b.Segments) != 2 {
			t.Errorf("régua %d: esperava 2 faixas do dossiê, veio %d", i, len(b.Segments))
		}
		if b.Axis[1] == 0 {
			t.Errorf("régua %d: eixo não veio do dossiê (%v)", i, b.Axis)
		}
		if len(b.History) != 1 {
			t.Errorf("régua %d: histórico não veio do dossiê", i)
		}
		if b.Unit != "mg/dL" {
			t.Errorf("régua %d: unidade não veio do dossiê (%q)", i, b.Unit)
		}
	}
	// O texto autoral é do MODELO, não do catálogo.
	if out[0].Rulers[0].Sub != "o rim filtrando" {
		t.Errorf("o sub do autor foi sobrescrito: %q", out[0].Rulers[0].Sub)
	}
}
