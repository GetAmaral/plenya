package services

import (
	"strings"
	"testing"
	"time"

	"github.com/plenya/api/internal/models"
)

func f(v float64) *float64 { return &v }
func s(v string) *string   { return &v }

func tplComp(substance string, qty float64, unit string) models.MagistralFormulaTemplateComponent {
	return models.MagistralFormulaTemplateComponent{Substance: substance, Quantity: qty, Unit: unit}
}

var agora = time.Date(2026, 8, 19, 12, 0, 0, 0, time.UTC)

func TestSuggestDoseSemRegra(t *testing.T) {
	got := SuggestDose(SuggestInput{Component: tplComp("Melatonina", 3, "mg"), Now: agora})
	if got.Suggested != nil {
		t.Error("sem regra não pode haver sugestão")
	}
	if got.BaseDose != 3 {
		t.Errorf("a dose da fórmula-base tem que vir junto, veio %v", got.BaseDose)
	}
}

func TestSuggestDosePorPeso(t *testing.T) {
	peso := time.Date(2026, 7, 20, 0, 0, 0, 0, time.UTC)
	rule := &models.MagistralFormulaTemplateRule{
		Kind: models.DoseRulePerKg, PerKg: f(20), MinDose: 500, MaxDose: 3000, MaxDataAgeDays: 180,
	}

	got := SuggestDose(SuggestInput{
		Component: tplComp("Creatina", 3000, "mg"),
		Rule:      rule,
		Weight:    &MeasurementInput{Value: 80, Date: &peso, Source: "consulta", Unit: "kg"},
		Now:       agora,
	})
	if got.Suggested == nil || *got.Suggested != 1600 {
		t.Fatalf("80 kg × 20 mg/kg = 1600 mg, veio %v", got.Suggested)
	}
	if got.Clamped {
		t.Error("1600 está dentro da faixa, não deveria travar")
	}
	if !strings.Contains(got.Basis, "80 kg") || !strings.Contains(got.Basis, "consulta") {
		t.Errorf("a base precisa dizer o peso e de onde veio: %q", got.Basis)
	}

	// trava superior: peso alto não vira dose absurda
	pesado := SuggestDose(SuggestInput{
		Component: tplComp("Creatina", 3000, "mg"),
		Rule:      rule,
		Weight:    &MeasurementInput{Value: 300, Date: &peso, Source: "consulta", Unit: "kg"},
		Now:       agora,
	})
	if pesado.Suggested == nil || *pesado.Suggested != 3000 {
		t.Fatalf("a trava deveria cortar em 3000, veio %v", pesado.Suggested)
	}
	if !pesado.Clamped || !strings.Contains(pesado.Basis, "trava") {
		t.Errorf("a resposta precisa DIZER que travou: %+v", pesado)
	}

	// peso velho não sugere nada
	velho := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	antigo := SuggestDose(SuggestInput{
		Component: tplComp("Creatina", 3000, "mg"),
		Rule:      rule,
		Weight:    &MeasurementInput{Value: 80, Date: &velho, Source: "consulta", Unit: "kg"},
		Now:       agora,
	})
	if antigo.Suggested != nil {
		t.Error("peso de 2 anos atrás não pode virar dose")
	}
	if !strings.Contains(antigo.Reason, "acima do limite") {
		t.Errorf("o motivo precisa ser dito: %q", antigo.Reason)
	}

	// sem peso, sem sugestão
	semPeso := SuggestDose(SuggestInput{Component: tplComp("Creatina", 3000, "mg"), Rule: rule, Now: agora})
	if semPeso.Suggested != nil || semPeso.Reason == "" {
		t.Errorf("sem peso: nenhuma sugestão e um motivo, veio %+v", semPeso)
	}
}

func TestSuggestDosePorExame(t *testing.T) {
	colheita := time.Date(2026, 7, 12, 0, 0, 0, 0, time.UTC)
	rule := &models.MagistralFormulaTemplateRule{
		Kind: models.DoseRuleLabThreshold,
		// o exame é referência ao catálogo, escolhida na tela
		LabCode: s("PLN1BF562ED"), LabOperator: s("lt"), LabThreshold: f(30),
		DoseIfTrue: f(5000), DoseIfFalse: f(2000),
		MinDose: 1000, MaxDose: 10000, MaxDataAgeDays: 365,
	}
	lab := &MeasurementInput{Value: 22, Date: &colheita, Source: "25-OH-vitamina D", Unit: "ng/mL"}

	abaixo := SuggestDose(SuggestInput{Component: tplComp("Vitamina D3", 2000, "UI"), Rule: rule, Lab: lab, Now: agora})
	if abaixo.Suggested == nil || *abaixo.Suggested != 5000 {
		t.Fatalf("22 < 30 ⇒ 5000 UI, veio %v", abaixo.Suggested)
	}
	// a base é a frase que o médico lê antes de aceitar
	for _, esperado := range []string{"25-OH-vitamina D", "22", "ng/mL", "de 12/07", "<", "30", "5.000"} {
		if !strings.Contains(abaixo.Basis, esperado) {
			t.Errorf("a base precisa conter %q: %q", esperado, abaixo.Basis)
		}
	}

	acima := *lab
	acima.Value = 45
	fora := SuggestDose(SuggestInput{Component: tplComp("Vitamina D3", 2000, "UI"), Rule: rule, Lab: &acima, Now: agora})
	if fora.Suggested == nil || *fora.Suggested != 2000 {
		t.Fatalf("45 não é < 30 ⇒ dose alternativa 2000, veio %v", fora.Suggested)
	}
	// a frase tem que deixar claro que a condição NÃO foi atendida
	if !strings.Contains(fora.Basis, "não atinge") {
		t.Errorf("a base do ramo falso precisa dizer que não atingiu o limiar: %q", fora.Basis)
	}

	semExame := SuggestDose(SuggestInput{Component: tplComp("Vitamina D3", 2000, "UI"), Rule: rule, Now: agora})
	if semExame.Suggested != nil || !strings.Contains(semExame.Reason, "exame") {
		t.Errorf("sem exame: nenhuma sugestão e motivo claro, veio %+v", semExame)
	}
}

func TestSuggestDoseTravaInferior(t *testing.T) {
	peso := agora.AddDate(0, 0, -10)
	got := SuggestDose(SuggestInput{
		Component: tplComp("Magnésio", 300, "mg"),
		Rule:      &models.MagistralFormulaTemplateRule{Kind: models.DoseRulePerKg, PerKg: f(1), MinDose: 200, MaxDose: 600, MaxDataAgeDays: 365},
		Weight:    &MeasurementInput{Value: 50, Date: &peso, Source: "consulta", Unit: "kg"},
		Now:       agora,
	})
	if got.Suggested == nil || *got.Suggested != 200 || !got.Clamped {
		t.Fatalf("50 mg está abaixo do piso 200, deveria subir e avisar: %+v", got)
	}
}

func TestFmtDose(t *testing.T) {
	cases := map[float64]string{5000: "5.000", 0.25: "0,25", 3: "3", 1600: "1.600", 12.5: "12,5"}
	for in, want := range cases {
		if got := fmtDose(in); got != want {
			t.Errorf("fmtDose(%v) = %q, esperado %q", in, got, want)
		}
	}
}

// --- faixas, guarda de unidade e arredondamento ---

func bandsVitD() []models.MagistralFormulaTemplateRuleBand {
	return []models.MagistralFormulaTemplateRuleBand{
		{UpperBound: f(10), Dose: 7000, Label: "deficiência grave", DisplayOrder: 0},
		{LowerBound: f(10), UpperBound: f(20), Dose: 5000, Label: "deficiência", DisplayOrder: 1},
		{LowerBound: f(20), UpperBound: f(30), Dose: 2000, Label: "insuficiência", DisplayOrder: 2},
		{LowerBound: f(30), Dose: 1000, Label: "suficiente", DisplayOrder: 3},
	}
}

func ruleVitD() *models.MagistralFormulaTemplateRule {
	return &models.MagistralFormulaTemplateRule{
		Kind: models.DoseRuleLabBand, LabCode: s("PLN1BF562ED"), LabUnit: s("ng/mL"),
		Bands: bandsVitD(), MinDose: 1000, MaxDose: 10000, MaxDataAgeDays: 365,
	}
}

func labVitD(value float64, unit string) *MeasurementInput {
	colheita := time.Date(2026, 7, 10, 0, 0, 0, 0, time.UTC)
	return &MeasurementInput{Value: value, Date: &colheita, Source: "25-hidroxivitamina D", Unit: unit}
}

func TestSuggestDosePorFaixa(t *testing.T) {
	casos := []struct {
		valor float64
		dose  float64
		faixa string
	}{
		{8, 7000, "deficiência grave"},
		{10, 7000, "deficiência grave"}, // (lower, upper]: o 10 fecha a primeira faixa
		{10.1, 5000, "deficiência"},
		{22, 2000, "insuficiência"},
		{30, 2000, "insuficiência"},
		{31, 1000, "suficiente"},
	}
	for _, c := range casos {
		got := SuggestDose(SuggestInput{
			Component: tplComp("Vitamina D3", 2000, "UI"),
			Rule:      ruleVitD(), Lab: labVitD(c.valor, "ng/mL"), Now: agora,
		})
		if got.Suggested == nil || *got.Suggested != c.dose {
			t.Fatalf("%v ng/mL devia sugerir %v UI, veio %v (%s)", c.valor, c.dose, got.Suggested, got.Reason)
		}
		if !strings.Contains(got.Basis, c.faixa) {
			t.Errorf("a frase precisa nomear a faixa %q, veio %q", c.faixa, got.Basis)
		}
	}
}

// A guarda de unidade só serve se não for barulhenta: variante cosmética da MESMA unidade passa.
func TestUnidadeEquivalenteNaoAtrapalha(t *testing.T) {
	iguais := [][2]string{
		{"µg/L", "mcg/L"}, {"µg/L", "ug/l"}, {"ng/mL", "µg/L"},
		{"µmol/L", "umol/L"}, {"IU/mL", "UI/mL"}, {"mIU/L", "µUI/mL"}, {"µUI/mL", "µU/mL"},
	}
	for _, par := range iguais {
		if msg := unitMismatch(s(par[0]), par[1]); msg != "" {
			t.Errorf("%s e %s são a mesma unidade, e a regra recusou: %s", par[0], par[1], msg)
		}
	}
}

// E precisa recusar de verdade quando a unidade muda o número: é o caso do cortisol em nmol/L
// sobre definição em µg/dL (fator 27,6) e o da vitamina D gravada em pg/mL, ambos no banco.
func TestUnidadeDiferenteNaoSugere(t *testing.T) {
	got := SuggestDose(SuggestInput{
		Component: tplComp("Vitamina D3", 2000, "UI"),
		Rule:      ruleVitD(), Lab: labVitD(46.2, "pg/mL"), Now: agora,
	})
	if got.Suggested != nil {
		t.Fatalf("resultado em pg/mL contra regra em ng/mL não pode virar dose, veio %v", *got.Suggested)
	}
	if !strings.Contains(got.Reason, "pg/mL") || !strings.Contains(got.Reason, "ng/mL") {
		t.Errorf("o motivo tem que nomear as duas unidades, veio %q", got.Reason)
	}

	diferentes := [][2]string{{"ng/mL", "µg/dL"}, {"mg/dL", "mg/L"}, {"pg/mL", "ng/mL"}}
	for _, par := range diferentes {
		if msg := unitMismatch(s(par[0]), par[1]); msg == "" {
			t.Errorf("%s e %s não são a mesma unidade e a regra deixou passar", par[0], par[1])
		}
	}
}

func TestFaixaDescobertaNaoInventaDose(t *testing.T) {
	rule := ruleVitD()
	// Só a faixa de deficiência: acima de 20 não há conduta cadastrada.
	rule.Bands = []models.MagistralFormulaTemplateRuleBand{{UpperBound: f(20), Dose: 5000}}

	got := SuggestDose(SuggestInput{
		Component: tplComp("Vitamina D3", 2000, "UI"),
		Rule:      rule, Lab: labVitD(35, "ng/mL"), Now: agora,
	})
	if got.Suggested != nil {
		t.Fatalf("valor fora das faixas não pode virar dose, veio %v", *got.Suggested)
	}
	if !strings.Contains(got.Reason, "faixa") {
		t.Errorf("o motivo devia falar de faixa, veio %q", got.Reason)
	}
}

func TestArredondamentoAntesDaTrava(t *testing.T) {
	peso := time.Date(2026, 8, 1, 0, 0, 0, 0, time.UTC)
	rule := &models.MagistralFormulaTemplateRule{
		Kind: models.DoseRulePerKg, PerKg: f(4.5), RoundTo: f(50),
		MinDose: 200, MaxDose: 400, MaxDataAgeDays: 365,
	}
	// 73,4 kg × 4,5 = 330,3 mg → arredonda para 350 mg, dentro da trava.
	got := SuggestDose(SuggestInput{
		Component: tplComp("Magnésio quelato", 300, "mg"),
		Rule:      rule,
		Weight:    &MeasurementInput{Value: 73.4, Date: &peso, Source: "peso da consulta", Unit: "kg"},
		Now:       agora,
	})
	if got.Suggested == nil || *got.Suggested != 350 {
		t.Fatalf("esperava 350 mg arredondados, veio %v", got.Suggested)
	}
	if got.Clamped {
		t.Error("350 está dentro de 200 a 400: não devia travar")
	}
	if !strings.Contains(got.Basis, "arredondada") {
		t.Errorf("a frase precisa dizer que arredondou, veio %q", got.Basis)
	}

	// 110 kg × 4,5 = 495 → arredonda para 500 → a trava corta em 400, e a frase diz isso.
	got = SuggestDose(SuggestInput{
		Component: tplComp("Magnésio quelato", 300, "mg"),
		Rule:      rule,
		Weight:    &MeasurementInput{Value: 110, Date: &peso, Source: "peso da consulta", Unit: "kg"},
		Now:       agora,
	})
	if got.Suggested == nil || *got.Suggested != 400 {
		t.Fatalf("a trava devia cortar em 400, veio %v", got.Suggested)
	}
	if !got.Clamped || !strings.Contains(got.Basis, "trava") {
		t.Errorf("a resposta precisa assumir que travou, veio %q", got.Basis)
	}
}

// A regra é escrita em dose DIÁRIA e o campo que a tela preenche é a dose de UMA cápsula. Sem a
// divisão pelas tomadas, uma regra de 5.000 UI/dia numa fórmula tomada duas vezes ao dia
// entregaria 10.000 UI por dia.
func TestSuggestDoseDivideNasTomadas(t *testing.T) {
	rule := ruleVitD()
	rule.RoundTo = f(250)

	uma := SuggestDose(SuggestInput{
		Component: tplComp("Vitamina D3", 2000, "UI"), Rule: rule,
		Lab: labVitD(15, "ng/mL"), Now: agora, DosesPerDay: 1,
	})
	if uma.Suggested == nil || *uma.Suggested != 5000 {
		t.Fatalf("uma tomada ao dia recebe a dose diária inteira, veio %v", uma.Suggested)
	}

	duas := SuggestDose(SuggestInput{
		Component: tplComp("Vitamina D3", 2000, "UI"), Rule: rule,
		Lab: labVitD(15, "ng/mL"), Now: agora, DosesPerDay: 2,
	})
	if duas.Suggested == nil || *duas.Suggested != 2500 {
		t.Fatalf("5.000 UI/dia em duas tomadas são 2.500 por cápsula, veio %v", duas.Suggested)
	}
	if !strings.Contains(duas.Basis, "÷ 2 tomadas") {
		t.Errorf("a conta precisa aparecer na frase, veio %q", duas.Basis)
	}
}

// A trava é diária: corta antes de dividir, senão o teto viraria teto por cápsula.
func TestSuggestDoseTravaEhDiaria(t *testing.T) {
	peso := time.Date(2026, 8, 1, 0, 0, 0, 0, time.UTC)
	rule := &models.MagistralFormulaTemplateRule{
		Kind: models.DoseRulePerKg, PerKg: f(5), MinDose: 200, MaxDose: 350, MaxDataAgeDays: 365,
	}
	got := SuggestDose(SuggestInput{
		Component: tplComp("Magnésio", 150, "mg"), Rule: rule,
		Weight: &MeasurementInput{Value: 110, Date: &peso, Source: "consulta", Unit: "kg"},
		Now:    agora, DosesPerDay: 2,
	})
	// 110 × 5 = 550 → trava diária corta em 350 → duas tomadas de 175.
	if got.Suggested == nil || *got.Suggested != 175 {
		t.Fatalf("teto de 350 mg/dia em duas tomadas são 175 por dose, veio %v", got.Suggested)
	}
	if !got.Clamped {
		t.Error("a resposta precisa assumir que travou")
	}
}
