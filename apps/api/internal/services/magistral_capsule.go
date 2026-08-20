package services

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// Calculadora de tamanho de cápsula.
//
// A matemática é simples: volume dos ativos = Σ(massa / densidade aparente); número de cápsulas
// por dose = teto(volume / volume da cápsula). O que é difícil é a HONESTIDADE do resultado.
//
// A densidade aparente de um pó varia com lote, granulometria e compactação da farmácia. O erro
// real desse cálculo é da ordem de ±25%. Consequências de projeto, deliberadas:
//
//   1. o resultado é uma FAIXA, nunca um número seco;
//   2. faltando densidade de QUALQUER componente, a calculadora NÃO OPINA — devolve quem está
//      sem densidade cadastrada. Silêncio ganha de número inventado;
//   3. quem decide é a farmácia: isto é auxílio de consultório, não especificação de manipulação.

// CapsuleSize é um tamanho de cápsula gelatinosa dura. Volumes em mL — padrão de indústria.
type CapsuleSize struct {
	Name   string
	Volume float64 // mL
}

// capsuleSizes, do maior para o menor.
var capsuleSizes = []CapsuleSize{
	{"000", 1.37},
	{"00", 0.95},
	{"0", 0.68},
	{"1", 0.50},
	{"2", 0.37},
	{"3", 0.30},
	{"4", 0.21},
	{"5", 0.13},
}

// densityErrorMargin — incerteza assumida da densidade aparente. Define a faixa do resultado.
const densityErrorMargin = 0.25

// maxCapsulesPerDose — acima disto a recomendação deixa de ser cápsula: ninguém toma 3 cápsulas
// por dose de bom grado, e a farmácia vai sugerir sachê de qualquer forma.
const maxCapsulesPerDose = 2

// CapsuleInput — um componente com massa já convertida para miligramas.
type CapsuleInput struct {
	Substance   string
	MassMg      float64
	BulkDensity float64 // g/mL; 0 = não cadastrada
	// DensityApprox: a densidade veio de aproximação por classe de pó, não de medida do insumo.
	// Muda o texto, não o cálculo: quem lê precisa saber o que está lendo.
	DensityApprox bool
}

// CapsuleAdvice é o que a tela mostra.
type CapsuleAdvice struct {
	// Decidido diz se dá para opinar. Falso => Missing lista quem falta.
	Decided bool     `json:"decided"`
	Missing []string `json:"missing,omitempty"`

	// Volume estimado dos ativos, em mL, e a faixa (± margem de densidade).
	VolumeML    float64 `json:"volumeMl,omitempty"`
	VolumeMinML float64 `json:"volumeMinMl,omitempty"`
	VolumeMaxML float64 `json:"volumeMaxMl,omitempty"`

	// Tamanho recomendado e quantas cápsulas por dose. Vazio quando a recomendação é sachê.
	Size            string `json:"size,omitempty"`
	CapsulesPerDose int    `json:"capsulesPerDose,omitempty"`
	// SizeIfCompacted é o tamanho que serve no cenário otimista (pó mais compactado que o
	// esperado). Quando difere do recomendado, a tela mostra os dois — é a faixa honesta.
	SizeIfCompacted string `json:"sizeIfCompacted,omitempty"`

	// SachetRecommended: nem a maior cápsula resolve em número aceitável de doses.
	SachetRecommended bool `json:"sachetRecommended"`

	// Explanation é o texto pronto, com a incerteza declarada.
	Explanation string `json:"explanation"`
	// Approximate: pelo menos um componente usou densidade aproximada por classe.
	Approximate bool `json:"approximate"`
}

// CalculateCapsule decide o tamanho de cápsula para UMA dose.
func CalculateCapsule(components []CapsuleInput) CapsuleAdvice {
	if len(components) == 0 {
		return CapsuleAdvice{Explanation: "Sem componentes para calcular."}
	}

	var missing []string
	volume := 0.0
	aproximada := false
	for _, c := range components {
		if c.DensityApprox {
			aproximada = true
		}
		if c.BulkDensity <= 0 {
			missing = append(missing, c.Substance)
			continue
		}
		if c.MassMg <= 0 {
			continue
		}
		// massa em mg → g; densidade em g/mL ⇒ volume em mL
		volume += (c.MassMg / 1000.0) / c.BulkDensity
	}

	if len(missing) > 0 {
		sort.Strings(missing)
		return CapsuleAdvice{
			Decided: false,
			Missing: missing,
			Explanation: "Sem densidade cadastrada para " + strings.Join(missing, ", ") +
				". Cadastre a densidade aparente para o tamanho da cápsula ser calculado.",
		}
	}

	if volume <= 0 {
		return CapsuleAdvice{Explanation: "Sem massa para calcular."}
	}

	volMin := volume * (1 - densityErrorMargin)
	volMax := volume * (1 + densityErrorMargin)

	advice := CapsuleAdvice{
		Decided:     true,
		Approximate: aproximada,
		VolumeML:    round3(volume),
		VolumeMinML: round3(volMin),
		VolumeMaxML: round3(volMax),
	}

	// Cenário conservador (pó menos compactado que o esperado) manda na recomendação.
	size, n, ok := smallestCapsule(volMax)
	if !ok {
		advice.SachetRecommended = true
		advice.Explanation = fmt.Sprintf(
			"Volume estimado de %s mL por dose: nem a cápsula 000 resolve em até %d cápsulas. Recomendado sachê.",
			fmtML(volume), maxCapsulesPerDose,
		)
		return advice
	}

	advice.Size = size
	advice.CapsulesPerDose = n
	if optimistic, on, ook := smallestCapsule(volMin); ook && optimistic != size && on == n {
		advice.SizeIfCompacted = optimistic
	}

	plural := "cápsula"
	if n > 1 {
		plural = "cápsulas"
	}
	advice.Explanation = fmt.Sprintf(
		"Volume estimado de %s mL por dose (faixa %s a %s): %d %s tamanho %s. "+
			"Estimativa com ±%.0f%% de incerteza — a densidade real varia com o lote e a compactação da farmácia.",
		fmtML(volume), fmtML(volMin), fmtML(volMax), n, plural, size, densityErrorMargin*100,
	)
	if advice.SizeIfCompacted != "" {
		advice.Explanation += fmt.Sprintf(" Com pó mais compactado, o tamanho %s serve.", advice.SizeIfCompacted)
	}
	if aproximada {
		advice.Explanation += " Densidade aproximada por classe do pó: use como ordem de grandeza e confirme com a farmácia."
	}
	return advice
}

// smallestCapsule devolve o MENOR tamanho que comporta o volume.
//
// Duas passadas, nesta ordem: primeiro procura o menor tamanho que resolve em UMA cápsula —
// engolir uma cápsula pequena é sempre melhor que engolir duas menores ainda. Só quando nem a
// 000 cabe numa dose é que aceita duas cápsulas.
func smallestCapsule(volume float64) (string, int, bool) {
	for maxN := 1; maxN <= maxCapsulesPerDose; maxN++ {
		for i := len(capsuleSizes) - 1; i >= 0; i-- {
			c := capsuleSizes[i]
			n := int(math.Ceil(volume / c.Volume))
			if n <= maxN {
				return c.Name, n, true
			}
		}
	}
	return "", 0, false
}

func round3(v float64) float64 { return math.Round(v*1000) / 1000 }

// fmtML formata em pt-BR com 2 casas, sem zeros à toa: 0.68 → "0,68"; 1.5 → "1,5".
// fmtML formata em pt-BR. Acima de mil delega para fmtDose, que agrupa o milhar: "10000 mcg" no
// meio de um alerta clínico se lê errado.
func fmtML(v float64) string {
	if v >= 1000 {
		return fmtDose(v)
	}
	s := strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.2f", v), "0"), ".")
	if s == "" || s == "-" {
		s = "0"
	}
	return strings.Replace(s, ".", ",", 1)
}

// MassToMg converte a quantidade do componente para miligramas quando dá. Unidades sem massa
// definida (UI, %, mL) devolvem ok=false — não entram no cálculo de volume.
func MassToMg(quantity float64, unit string) (float64, bool) {
	switch strings.ToLower(strings.TrimSpace(unit)) {
	case "mg":
		return quantity, true
	case "g":
		return quantity * 1000, true
	case "mcg", "µg", "ug":
		return quantity / 1000, true
	default:
		return 0, false
	}
}
