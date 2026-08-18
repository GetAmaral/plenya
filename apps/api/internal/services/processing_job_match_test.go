package services

import (
	"testing"

	"github.com/google/uuid"
)

// Laudo escreve os analitos de um painel como "Painel - Analito". O match por nome cai no
// painel (o nome dele é substring), e antes disso jogava os oito analitos da gasometria na
// mesma definição-container, sem faixa que classificasse nenhum.
func TestResolvePanelAnalyte_Gasometria(t *testing.T) {
	panelID := uuid.New()
	phID := uuid.New()
	pco2ID := uuid.New()
	po2ID := uuid.New()
	hco3ID := uuid.New()

	idx := &testDefIndex{
		byName:   map[string]uuid.UUID{},
		specByID: map[uuid.UUID]string{},
		isPanel:  map[uuid.UUID]bool{panelID: true},
		children: map[uuid.UUID]map[string]uuid.UUID{
			panelID: {
				"ph venoso":            phID,
				"ph gasometria venosa": phID,
				"pco2 venoso":          pco2ID,
				"po2 venoso":           po2ID,
				"hco3 bicarbonato":     hco3ID,
				"bicarbonato":          hco3ID,
			},
		},
	}

	tests := []struct {
		name string
		raw  string
		want *uuid.UUID
	}{
		{"pH", "Gasometria Venosa - pH", &phID},
		{"pCO2", "Gasometria Venosa - pCO2", &pco2ID},
		{"pO2", "Gasometria Venosa - pO2", &po2ID},
		{"HCO3", "Gasometria Venosa - HCO3-", &hco3ID},
		{"bicarbonato por extenso", "Gasometria Venosa - Bicarbonato", &hco3ID},
		// Analitos sem definição própria: melhor "fora do catálogo" do que pendurados no painel.
		{"CO2 total sem definição", "Gasometria Venosa - CO2 Total", nil},
		{"base excess sem definição", "Gasometria Venosa - Base Excess", nil},
		{"saturação sem definição", "Gasometria Venosa - Saturação de O2", nil},
		// Nome do painel puro não é um resultado: não pode grudar valor no container.
		{"painel sem analito", "Gasometria Venosa", nil},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := idx.resolvePanelAnalyte(tc.raw, &panelID)
			switch {
			case tc.want == nil && got != nil:
				t.Fatalf("%q: esperava nenhum match, veio %s", tc.raw, got)
			case tc.want != nil && got == nil:
				t.Fatalf("%q: esperava %s, não casou", tc.raw, tc.want)
			case tc.want != nil && *got != *tc.want:
				t.Fatalf("%q: casou com o analito errado", tc.raw)
			}
		})
	}
}

// Match que não é painel passa direto — a resolução só age sobre containers.
func TestResolvePanelAnalyte_NaoPainelPassaDireto(t *testing.T) {
	plainID := uuid.New()
	idx := &testDefIndex{
		isPanel:  map[uuid.UUID]bool{},
		children: map[uuid.UUID]map[string]uuid.UUID{},
	}

	got := idx.resolvePanelAnalyte("Glicose - jejum", &plainID)
	if got == nil || *got != plainID {
		t.Fatalf("definição comum deveria passar intacta pela resolução de painel")
	}

	if idx.resolvePanelAnalyte("Glicose", nil) != nil {
		t.Fatalf("sem match de entrada, a saída tem que continuar sem match")
	}
}

func TestAnalyteSuffix(t *testing.T) {
	cases := map[string]string{
		"Gasometria Venosa - pH":                    "pH",
		"Gasometria Venosa – Base Excess":           "Base Excess", // en dash
		"Ultrassonografia de Abdome Total - Fígado": "Fígado",
		"Glicose": "",
		"Hemoglobina glicada - HbA1c - resultado": "resultado",
	}
	for raw, want := range cases {
		if got := analyteSuffix(raw); got != want {
			t.Fatalf("analyteSuffix(%q) = %q, esperado %q", raw, got, want)
		}
	}
}
