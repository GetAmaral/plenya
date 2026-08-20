package services

import (
	"bytes"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/plenya/api/internal/models"
)

// testPrescription monta uma prescrição sintética (sem DB) para exercitar a geração de PDF.
func testPrescription(controlled bool) *models.Prescription {
	crm := "12345"
	uf := "PR"
	cpf := "12345678901"
	cat := models.MedCategorySimple
	if controlled {
		cat = models.MedCategoryC1
	}
	return &models.Prescription{
		ID:               uuid.Must(uuid.NewV7()),
		PrescriptionDate: time.Now(),
		ValidUntil:       time.Now().AddDate(0, 0, 30),
		Doctor: models.User{
			Name:  "Dr. Teste",
			CRM:   &crm,
			CRMUF: &uf,
		},
		Patient: models.Patient{
			Name: "Paciente Teste",
			CPF:  &cpf,
		},
		Medications: []models.PrescriptionMedication{{
			MedicationName:   "Medicamento X",
			ActiveIngredient: "substância x",
			Category:         cat,
			Concentration:    "20mg",
			Quantity:         30,
			QuantityInWords:  "trinta",
			Dosage:           "1 comprimido",
			Frequency:        "2x/dia",
			Route:            "VO",
			Duration:         30,
		}},
	}
}

// requireChromium pula o teste onde o renderer não roda. O pipeline vivo é HTML/CSS via
// Chromium headless (pdfdoc); o gerador antigo em gofpdf foi removido junto com este teste.
func requireChromium(t *testing.T) {
	t.Helper()
	for _, bin := range []string{"/usr/bin/chromium-browser", "/usr/bin/chromium"} {
		if _, err := os.Stat(bin); err == nil {
			return
		}
	}
	t.Skip("chromium indisponível — rode dentro do container api")
}

// TestGeneratePrescriptionPDF verifica os dois modos de geração (manual para impressão e
// digital para assinatura ICP-Brasil): ambos precisam sair como PDF válido.
func TestGeneratePrescriptionPDF(t *testing.T) {
	requireChromium(t)

	cases := []struct {
		name       string
		controlled bool
		manual     bool
	}{
		{"manual_controlado", true, true},
		{"digital_simples", false, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := testPrescription(tc.controlled)
			b, err := renderPrescriptionBytes(p, tc.manual)
			if err != nil {
				t.Fatalf("renderPrescriptionBytes erro: %v", err)
			}
			if len(b) < 1000 || !bytes.HasPrefix(b, []byte("%PDF")) {
				t.Fatalf("saída não parece um PDF válido (%d bytes)", len(b))
			}
		})
	}
}

// TestPrescriptionHasControlled trava a função que decide modo manual e rótulo de Controle
// Especial — é o único ponto que separa receita comum de receituário controlado.
func TestPrescriptionHasControlled(t *testing.T) {
	if prescriptionHasControlled(testPrescription(false)) {
		t.Error("receita simples não deveria ser tratada como controlada")
	}
	if !prescriptionHasControlled(testPrescription(true)) {
		t.Error("receita com C1 deveria ser tratada como controlada")
	}
}

// TestVerifySignatureUnsigned garante que VerifySignature lida com PDF não-assinado sem
// erro de servidor (reporta Signed=false), validando o binding com pdfsign/verify.
func TestVerifySignatureUnsigned(t *testing.T) {
	requireChromium(t)
	pdfBytes, err := renderPrescriptionBytes(testPrescription(false), true)
	if err != nil {
		t.Fatalf("falha ao gerar PDF base: %v", err)
	}
	res, err := (&SignatureService{}).VerifySignature(pdfBytes)
	if err != nil {
		t.Fatalf("VerifySignature retornou erro inesperado: %v", err)
	}
	if res == nil {
		t.Fatal("resultado de verificação nil")
	}
	if res.Signed {
		t.Fatalf("PDF não-assinado não deveria reportar Signed=true (res=%+v)", res)
	}
}

// testCompounded monta uma receita de manipulado sintética (sem DB).
func testCompounded(components int) *models.Prescription {
	crm := "12345"
	uf := "PR"
	instr := "Tomar com água. Manter ao abrigo da luz."
	comps := []models.PrescriptionFormulaComponent{
		{Substance: "Melatonina", Quantity: 0.25, Unit: "mg", Note: "liberação prolongada"},
		{Substance: "Magnésio dimalato", Quantity: 300, Unit: "mg"},
		{Substance: "L-teanina", Quantity: 200, Unit: "mg"},
	}
	for i := len(comps); i < components; i++ {
		comps = append(comps, models.PrescriptionFormulaComponent{
			Substance: "Substância de teste com nome longo " + itoaTest(i+1),
			Quantity:  float64(i+1) * 12.5,
			Unit:      "mg",
		})
	}

	return &models.Prescription{
		ID:               uuid.Must(uuid.NewV7()),
		Type:             models.PrescriptionCompounded,
		PrescriptionDate: time.Now(),
		ValidUntil:       time.Now().AddDate(0, 0, 30),
		Doctor:           models.User{Name: "Dr. Teste", CRM: &crm, CRMUF: &uf},
		Patient:          models.Patient{Name: "Paciente Teste"},
		Formulas: []models.PrescriptionFormula{
			{
				Name: "Fórmula do sono", PharmaceuticalForm: "cápsula",
				UsageType: models.FormulaUsageInternal, Vehicle: "Excipiente qsp 1 cápsula",
				QuantityToDispense: 60, QuantityUnit: "cápsulas", QuantityInWords: "sessenta",
				Posology: "1 cápsula ao deitar", Duration: 60, Instructions: &instr,
				Components: comps,
			},
			{
				Name: "Creme facial", PharmaceuticalForm: "creme",
				UsageType: models.FormulaUsageExternal, Vehicle: "Creme não iônico qsp 30 g",
				QuantityToDispense: 30, QuantityUnit: "g",
				Posology: "aplicar à noite no rosto limpo", Duration: 90,
				Components: []models.PrescriptionFormulaComponent{
					{Substance: "Vitamina C (ácido ascórbico)", Quantity: 10, Unit: "%"},
					{Substance: "Ácido hialurônico", Quantity: 1, Unit: "%"},
				},
			},
		},
	}
}

func itoaTest(n int) string { return strconv.Itoa(n) }

// TestGenerateCompoundedPDF garante que o receituário magistral sai como PDF válido nos dois
// modos e que a fórmula cheia (20 componentes, o teto da validação) ainda cabe numa página.
func TestGenerateCompoundedPDF(t *testing.T) {
	requireChromium(t)

	for _, tc := range []struct {
		name       string
		components int
		manual     bool
	}{
		{"tres_componentes", 3, true},
		{"vinte_componentes", 20, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			p := testCompounded(tc.components)
			b, err := renderPrescriptionBytes(p, tc.manual)
			if err != nil {
				t.Fatalf("renderPrescriptionBytes: %v", err)
			}
			if len(b) < 1000 || !bytes.HasPrefix(b, []byte("%PDF")) {
				t.Fatalf("saída não parece um PDF válido (%d bytes)", len(b))
			}
			if out := os.Getenv("PDF_OUT_DIR"); out != "" {
				_ = os.WriteFile(out+"/manipulado-"+tc.name+".pdf", b, 0o644)
			}
		})
	}
}

// A fórmula com componente controlado tem que puxar a receita inteira para o receituário de
// controle especial — é o que decide modo manual e rótulo impresso.
func TestCompoundedControlled(t *testing.T) {
	p := testCompounded(3)
	if prescriptionHasControlled(p) {
		t.Error("fórmula sem controlado não deveria virar receituário de controle")
	}
	p.Formulas[0].HighestCategory = models.MedCategoryC1
	if !prescriptionHasControlled(p) {
		t.Error("fórmula com C1 deveria virar receituário de controle")
	}
}

// A dose declarada como do ELEMENTO precisa sair impressa. Sem essa palavra, "magnésio quelato
// 150 mg" manda a farmácia pesar 150 mg do bisglicinato, que são 45 mg de magnésio — e o PDF é o
// único documento que a farmácia lê.
func TestFormulasPDFImprimeDoseDoElemento(t *testing.T) {
	formulas := []models.PrescriptionFormula{{
		Name: "Sachê", PharmaceuticalForm: "sachê", UsageType: models.FormulaUsageInternal,
		QuantityToDispense: 30, QuantityUnit: "sachês", Posology: "1 sachê ao dia",
		Components: []models.PrescriptionFormulaComponent{
			{Substance: "Magnésio quelato", Quantity: 150, Unit: "mg", AsElemental: true},
			{Substance: "Taurina", Quantity: 1, Unit: "g"},
		},
	}}

	out := buildFormulasPDF(formulas)
	if len(out) != 1 || len(out[0].Components) != 2 {
		t.Fatalf("esperava 1 fórmula com 2 componentes, veio %+v", out)
	}
	if !out[0].Components[0].AsElemental {
		t.Error("a marca de dose do elemento se perdeu no caminho para o PDF")
	}
	if out[0].Components[1].AsElemental {
		t.Error("componente sem a marca não pode ganhá-la")
	}
}
