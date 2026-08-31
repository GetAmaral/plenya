package services

import (
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

func comp(cat models.MedicationCategory, substance string, qty float64, unit string) dto.FormulaComponentRequest {
	return dto.FormulaComponentRequest{Substance: substance, Quantity: qty, Unit: unit, Category: cat}
}

func formulaReq(components ...dto.FormulaComponentRequest) dto.FormulaRequest {
	return dto.FormulaRequest{
		PharmaceuticalForm: "cápsula",
		UsageType:          models.FormulaUsageInternal,
		QuantityToDispense: 60,
		QuantityUnit:       "cápsulas",
		Posology:           "1 cápsula ao deitar",
		Components:         components,
	}
}

func TestValidateFormulas(t *testing.T) {
	ok := formulaReq(comp(models.MedCategorySimple, "Melatonina", 3, "mg"))

	if err := validateFormulas(nil); err == nil {
		t.Error("receita de manipulado sem fórmula deveria falhar")
	}
	if err := validateFormulas([]dto.FormulaRequest{ok}); err != nil {
		t.Errorf("fórmula simples deveria passar, veio: %v", err)
	}

	// 21 componentes: limite de LAYOUT (o bloco não cabe na página e o paginador não quebra bloco)
	big := ok
	big.Components = nil
	for i := 0; i < 21; i++ {
		big.Components = append(big.Components, comp(models.MedCategorySimple, "Substância", 1, "mg"))
	}
	if err := validateFormulas([]dto.FormulaRequest{big}); err == nil {
		t.Error("21 componentes deveriam ser recusados")
	}
	big.Components = big.Components[:20]
	if err := validateFormulas([]dto.FormulaRequest{big}); err != nil {
		t.Errorf("20 componentes são permitidos, veio: %v", err)
	}

	// controlado exige extenso
	ctrl := formulaReq(comp(models.MedCategoryC1, "Clonazepam", 0.25, "mg"))
	if err := validateFormulas([]dto.FormulaRequest{ctrl}); err == nil {
		t.Error("fórmula com controlado sem extenso deveria falhar")
	}
	ctrl.QuantityInWords = "sessenta cápsulas"
	if err := validateFormulas([]dto.FormulaRequest{ctrl}); err != nil {
		t.Errorf("com extenso deveria passar, veio: %v", err)
	}

	// 3 substâncias C1 distintas espalhadas em fórmulas diferentes passam; a 4ª não
	mk := func(nome string) dto.FormulaRequest {
		f := formulaReq(comp(models.MedCategoryC1, nome, 1, "mg"))
		f.QuantityInWords = "sessenta cápsulas"
		return f
	}
	tres := []dto.FormulaRequest{mk("Clonazepam"), mk("Zolpidem"), mk("Alprazolam")}
	if err := validateFormulas(tres); err != nil {
		t.Errorf("3 substâncias controladas deveriam passar, veio: %v", err)
	}
	// a mesma substância repetida em duas fórmulas conta uma vez
	if err := validateFormulas(append(tres, mk("clonazepam"))); err != nil {
		t.Errorf("substância repetida não deveria estourar o limite, veio: %v", err)
	}
	if err := validateFormulas(append(tres, mk("Bromazepam"))); err == nil {
		t.Error("4 substâncias controladas distintas deveriam ser recusadas")
	}
}

func TestBuildFormulasHighestCategory(t *testing.T) {
	req := []dto.FormulaRequest{formulaReq(
		comp(models.MedCategorySimple, "Magnésio", 300, "mg"),
		comp(models.MedCategoryC1, "Clonazepam", 0.25, "mg"),
		comp(models.MedCategorySimple, "Melatonina", 3, "mg"),
	)}
	formulas, err := buildFormulas(uuid.Nil, req)
	if err != nil {
		t.Fatalf("buildFormulas: %v", err)
	}
	if formulas[0].HighestCategory != models.MedCategoryC1 {
		t.Errorf("categoria da fórmula deveria ser a mais restritiva (c1), veio %q", formulas[0].HighestCategory)
	}
	if formulas[0].Components[2].DisplayOrder != 2 {
		t.Error("ordem dos componentes precisa ser preservada — é a ordem impressa na receita")
	}
}

func TestCalculateValidUntilFormulas(t *testing.T) {
	base := time.Date(2026, 8, 1, 12, 0, 0, 0, time.UTC)
	days := func(tt time.Time) int { return int(tt.Sub(base).Hours() / 24) }

	formulas := []models.PrescriptionFormula{
		{HighestCategory: models.MedCategorySimple},
		{HighestCategory: models.MedCategoryAntibiotic},
	}
	if got := days(calculateValidUntilFormulas(base, formulas)); got != 10 {
		t.Errorf("a fórmula mais restritiva (antimicrobiano) deveria puxar para 10 dias, veio %d", got)
	}
}

// O receituário magistral tem layout próprio; este teste trava o conteúdo do HTML sem depender
// do Chromium.
func TestBuildFormulasPDF(t *testing.T) {
	instr := "Manter refrigerado."
	f := models.PrescriptionFormula{
		Name:               "Fórmula do sono",
		PharmaceuticalForm: "cápsula",
		UsageType:          models.FormulaUsageInternal,
		Vehicle:            "Excipiente qsp 1 cápsula",
		QuantityToDispense: 60,
		QuantityUnit:       "cápsulas",
		QuantityInWords:    "sessenta",
		Posology:           "1 cápsula ao deitar",
		Duration:           60,
		Instructions:       &instr,
		Components: []models.PrescriptionFormulaComponent{
			{Substance: "Melatonina", Quantity: 0.25, Unit: "mg", Note: "liberação prolongada"},
			{Substance: "Magnésio dimalato", Quantity: 300, Unit: "mg"},
		},
	}

	out := buildFormulasPDF([]models.PrescriptionFormula{f})[0]
	// Sem o "Aviar": a palavra virou etiqueta do layout, impressa pelo pdfdoc. Mandá-la daqui
	// imprimiria "AVIAR  Aviar 60 (sessenta) cápsulas".
	if out.Dispense != "60 (sessenta) cápsulas" {
		t.Errorf("linha de aviamento errada: %q", out.Dispense)
	}
	if out.Posology != "1 cápsula ao deitar · por 60 dias" {
		t.Errorf("posologia errada: %q", out.Posology)
	}
	if out.UsageLabel != "Uso interno" {
		t.Errorf("rótulo de uso errado: %q", out.UsageLabel)
	}
	// decimal em pt-BR e sem zeros à toa
	if out.Components[0].Quantity != "0,25 mg" {
		t.Errorf("quantidade do componente deveria sair como 0,25 mg, veio %q", out.Components[0].Quantity)
	}
	if out.Components[1].Quantity != "300 mg" {
		t.Errorf("quantidade inteira não deveria ganhar casas decimais, veio %q", out.Components[1].Quantity)
	}
	if !strings.Contains(out.Components[0].Note, "liberação") {
		t.Error("nota do componente sumiu")
	}
	// porcentagem cola no número (convenção brasileira); as demais unidades levam espaço
	if got := componentQuantity(10, "%"); got != "10%" {
		t.Errorf("porcentagem deveria sair como 10%%, veio %q", got)
	}
}
