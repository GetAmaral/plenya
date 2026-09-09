package services

import (
	"testing"

	"github.com/google/uuid"

	"github.com/plenya/api/internal/models"
)

func res(lote uuid.UUID, v float64, unidade string) models.LabResult {
	return models.LabResult{ID: uuid.New(), LabResultBatchID: lote, ResultNumeric: &v, Unit: &unidade}
}

// O caso comum: colesterol total e HDL do mesmo lipidograma viram a razão.
func TestRazoes_CalculaQuandoOsDoisVemDoMesmoLote(t *testing.T) {
	lote := uuid.New()
	m := map[string]models.LabResult{
		"PLN919303A4": res(lote, 234, "mg/dL"), // colesterol total
		"PLN53A449CA": res(lote, 89, "mg/dL"),  // HDL
		"PLN543993C6": res(lote, 86, "mg/dL"),  // apoB
		"PLNA8451657": res(lote, 193, "mg/dL"), // apoA1
	}
	aplicaRazoesDerivadas(m)

	for _, c := range []struct {
		codigo string
		quer   float64
	}{
		{"PLN09DBBB62", 234.0 / 89.0},
		{"PLNAPOBA1", 86.0 / 193.0},
	} {
		r, ok := m[c.codigo]
		if !ok {
			t.Fatalf("%s não foi calculada", c.codigo)
		}
		if got := *r.ResultNumeric; got < c.quer-1e-9 || got > c.quer+1e-9 {
			t.Errorf("%s: esperava %.4f, veio %.4f", c.codigo, c.quer, got)
		}
		if *r.Unit != "ratio" {
			t.Errorf("%s: unidade %q, e o item do escore está em ratio", c.codigo, *r.Unit)
		}
		if !resultadoDerivado(r) {
			t.Errorf("%s: precisa ser reconhecível como derivada, senão vira FK para lab_results", c.codigo)
		}
	}
}

// A outra guarda: unidades diferentes não viram razão. A apoB é reportada em g/L por muitos
// laboratórios, e 0,86 g/L sobre 193 mg/dL daria 0,0045, que cai no melhor nível e entrega os 18
// pontos sem que nada acuse. O par inteiro em g/L, esse sim, tem de passar: a razão é adimensional.
func TestRazoes_ExigeMesmaUnidadeNosDoisComponentes(t *testing.T) {
	lote := uuid.New()
	m := map[string]models.LabResult{
		"PLN543993C6": res(lote, 0.86, "g/L"),
		"PLNA8451657": res(lote, 193, "mg/dL"),
	}
	aplicaRazoesDerivadas(m)
	if r, ok := m["PLNAPOBA1"]; ok {
		t.Errorf("dividiu g/L por mg/dL e produziu %.4f", *r.ResultNumeric)
	}

	m = map[string]models.LabResult{
		"PLN543993C6": res(lote, 0.86, "g/L"),
		"PLNA8451657": res(lote, 1.93, "g/L"),
	}
	aplicaRazoesDerivadas(m)
	r, ok := m["PLNAPOBA1"]
	if !ok {
		t.Fatal("o par inteiro em g/L tem de calcular: a razão é adimensional")
	}
	if quer := 0.86 / 1.93; *r.ResultNumeric < quer-1e-9 || *r.ResultNumeric > quer+1e-9 {
		t.Errorf("esperava %.4f, veio %.4f", quer, *r.ResultNumeric)
	}
}

// A guarda que importa: componentes de coletas diferentes não viram razão. Dividir o colesterol de
// hoje pelo HDL de dois anos atrás produz um número plausível que nunca existiu na paciente.
func TestRazoes_NaoMisturaColetas(t *testing.T) {
	m := map[string]models.LabResult{
		"PLN919303A4": res(uuid.New(), 234, "mg/dL"),
		"PLN53A449CA": res(uuid.New(), 89, "mg/dL"),
	}
	aplicaRazoesDerivadas(m)
	if _, ok := m["PLN09DBBB62"]; ok {
		t.Error("calculou a razão com numerador e denominador de lotes diferentes")
	}
}

// O que está no prontuário vale mais que a conta: laudo que imprime a razão não é sobrescrito.
func TestRazoes_NaoSobrescreveOQueFoiLancado(t *testing.T) {
	lote := uuid.New()
	lancada := res(lote, 9.99, "ratio")
	m := map[string]models.LabResult{
		"PLN09DBBB62": lancada,
		"PLN919303A4": res(lote, 234, "mg/dL"),
		"PLN53A449CA": res(lote, 89, "mg/dL"),
	}
	aplicaRazoesDerivadas(m)
	if *m["PLN09DBBB62"].ResultNumeric != 9.99 {
		t.Errorf("sobrescreveu o resultado lançado, veio %.2f", *m["PLN09DBBB62"].ResultNumeric)
	}
}

// Faltando um componente, ou com denominador zero, não há razão nenhuma para inventar.
func TestRazoes_ComponenteFaltandoOuDenominadorZero(t *testing.T) {
	lote := uuid.New()
	m := map[string]models.LabResult{"PLN919303A4": res(lote, 234, "mg/dL")}
	aplicaRazoesDerivadas(m)
	if _, ok := m["PLN09DBBB62"]; ok {
		t.Error("calculou a razão sem o denominador")
	}

	m = map[string]models.LabResult{
		"PLN919303A4": res(lote, 234, "mg/dL"),
		"PLN53A449CA": res(lote, 0, "mg/dL"),
	}
	aplicaRazoesDerivadas(m)
	if _, ok := m["PLN09DBBB62"]; ok {
		t.Error("dividiu por zero")
	}
}
