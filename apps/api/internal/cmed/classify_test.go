package cmed

import (
	"testing"

	"github.com/plenya/api/internal/models"
)

// Casos-âncora medidos no dado real da CMED. Cada linha aqui é uma decisão regulatória que,
// se mudar sem querer, muda a validade de uma receita de verdade.
func TestClassify(t *testing.T) {
	tests := []struct {
		nome           string
		substance      string
		classCode      string
		stripe         string
		wantCategory   models.MedicationCategory
		wantSource     string
		wantReview     bool
		wantPrescrible bool
	}{
		{
			nome: "GLP-1 pela substância, mesmo em classe de antiobesidade",
			// É o caso que a derivação por classe terapêutica perderia: a CMED põe Wegovy
			// e Saxenda em "preparações antiobesidade", não em antidiabéticos.
			substance: "SEMAGLUTIDA", classCode: "A8A", stripe: models.MedStripeRedRestricted,
			wantCategory: models.MedCategoryGLP1, wantSource: models.MedCategorySourceDerived,
			wantReview: false, wantPrescrible: true,
		},
		{
			nome:      "antibacteriano pela classe J1, mesmo sem tarja publicada",
			substance: "CEFALEXINA MONOIDRATADA", classCode: "J1D1", stripe: "",
			wantCategory: models.MedCategoryAntibiotic, wantSource: models.MedCategorySourceDerived,
			wantReview: false, wantPrescrible: true,
		},
		{
			nome:      "antifúngico (J2) NÃO vira antibiótico — RDC 471 não é 'tudo que começa com J'",
			substance: "FLUCONAZOL", classCode: "J2A1", stripe: models.MedStripeRed,
			wantCategory: models.MedCategorySimple, wantSource: models.MedCategorySourceDerived,
			wantReview: false, wantPrescrible: true,
		},
		{
			nome:      "anabolizante entra como C5 e pede conferência",
			substance: "UNDECILATO DE TESTOSTERONA", classCode: "G3B", stripe: models.MedStripeRedRestricted,
			wantCategory: models.MedCategoryC5, wantSource: models.MedCategorySourceDerived,
			wantReview: true, wantPrescrible: true,
		},
		{
			nome:      "tarja preta fica fora do receituário (o EMR não emite Notificação A/B)",
			substance: "CLONAZEPAM", classCode: "N3A0", stripe: models.MedStripeBlack,
			wantCategory: models.MedCategoryAB, wantSource: models.MedCategorySourceDerived,
			wantReview: true, wantPrescrible: false,
		},
		{
			nome:      "retido + classe do SNC = controle especial C1",
			substance: "CLORIDRATO DE DONEPEZILA", classCode: "N7D1", stripe: models.MedStripeRedRestricted,
			wantCategory: models.MedCategoryC1, wantSource: models.MedCategorySourceDerived,
			wantReview: false, wantPrescrible: true,
		},
		{
			nome:      "retido fora do SNC: assume C1 mas admite que é palpite",
			substance: "ISOTRETINOINA", classCode: "D5B1", stripe: models.MedStripeRedRestricted,
			wantCategory: models.MedCategoryC1, wantSource: models.MedCategorySourceFallback,
			wantReview: true, wantPrescrible: true,
		},
		{
			nome:      "tarja vermelha comum = receita simples",
			substance: "LOSARTANA POTASSICA", classCode: "C9C0", stripe: models.MedStripeRed,
			wantCategory: models.MedCategorySimple, wantSource: models.MedCategorySourceDerived,
			wantReview: false, wantPrescrible: true,
		},
		{
			nome:      "isento de prescrição também é simples",
			substance: "ACIDO ACETILSALICILICO", classCode: "N2B2", stripe: models.MedStripeNone,
			wantCategory: models.MedCategorySimple, wantSource: models.MedCategorySourceDerived,
			wantReview: false, wantPrescrible: true,
		},
		{
			nome:      "sem tarja publicada: simples, marcado para revisão",
			substance: "CARVEDILOL", classCode: "C7A0", stripe: "",
			wantCategory: models.MedCategorySimple, wantSource: models.MedCategorySourceFallback,
			wantReview: true, wantPrescrible: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.nome, func(t *testing.T) {
			got := Classify(tc.substance, tc.classCode, tc.stripe)

			if got.Category != tc.wantCategory {
				t.Errorf("categoria = %q, esperado %q", got.Category, tc.wantCategory)
			}
			if got.CategorySource != tc.wantSource {
				t.Errorf("origem da categoria = %q, esperado %q", got.CategorySource, tc.wantSource)
			}
			if got.NeedsReview != tc.wantReview {
				t.Errorf("needsReview = %v, esperado %v", got.NeedsReview, tc.wantReview)
			}
			if got.IsPrescribable != tc.wantPrescrible {
				t.Errorf("isPrescribable = %v, esperado %v", got.IsPrescribable, tc.wantPrescrible)
			}
		})
	}
}

func TestNormalizeStripe(t *testing.T) {
	cases := map[string]string{
		"Tarja Vermelha":               models.MedStripeRed,
		"Tarja Vermelha sob restrição": models.MedStripeRedRestricted,
		"Tarja Preta":                  models.MedStripeBlack,
		"Tarja Sem Tarja":              models.MedStripeNone,
		"- (*) ":                       "", // tarja desconhecida não se afirma
		"":                             "",
	}
	for raw, want := range cases {
		if got := NormalizeStripe(raw); got != want {
			t.Errorf("NormalizeStripe(%q) = %q, esperado %q", raw, got, want)
		}
	}
}

func TestSplitTherapeuticClass(t *testing.T) {
	code, desc := SplitTherapeuticClass("J1G1 - FLUORQUINOLONAS ORAIS")
	if code != "J1G1" || desc != "FLUORQUINOLONAS ORAIS" {
		t.Fatalf("split = (%q, %q), esperado (J1G1, FLUORQUINOLONAS ORAIS)", code, desc)
	}
	if code, desc := SplitTherapeuticClass("SEM CODIGO"); code != "" || desc != "SEM CODIGO" {
		t.Fatalf("classe sem código deveria virar descrição pura, veio (%q, %q)", code, desc)
	}
}

// A categoria manda nas regras da receita: antibiótico vale 10 dias, GLP-1 vale 90.
func TestRulesFor(t *testing.T) {
	if r := RulesFor(models.MedCategoryAntibiotic); r.ValidityDays != 10 {
		t.Errorf("antibiótico deveria valer 10 dias, veio %d", r.ValidityDays)
	}
	if r := RulesFor(models.MedCategoryGLP1); r.ValidityDays != 90 {
		t.Errorf("GLP-1 deveria valer 90 dias, veio %d", r.ValidityDays)
	}
	if r := RulesFor(models.MedCategoryC1); !r.RequiresDigitalSignature || r.MaxPerPrescription != 3 {
		t.Errorf("C1 deveria exigir assinatura e limitar a 3 itens, veio %+v", r)
	}
	if r := RulesFor("categoria_inexistente"); r.ValidityDays != 30 {
		t.Errorf("categoria desconhecida deveria cair no perfil simples, veio %+v", r)
	}
}
