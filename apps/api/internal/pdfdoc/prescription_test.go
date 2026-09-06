package pdfdoc

import (
	"os"
	"testing"
)

func TestRenderPrescription(t *testing.T) {
	if !chromiumAvailable() {
		t.Skip("chromium ausente")
	}
	in := Prescription{
		Patient: Patient{Name: "Maria Helena Soares", BirthInfo: "12/03/1979 · 47 anos", CPFMasked: "***.456.789-**"},
		Meds: []Med{
			{Name: "Losartana potássica", Concentration: "50 mg", ActiveIngredient: "Losartana potássica",
				Posology: "Tomar 1 comprimido de 12 em 12 horas",
				Quantity: "uso contínuo"},
			{Name: "Rosuvastatina cálcica", Concentration: "10 mg", ActiveIngredient: "Rosuvastatina cálcica",
				Posology:     "Tomar 1 comprimido uma vez ao dia, por 30 dias",
				Quantity:     "30 (trinta comprimidos)",
				Instructions: "Tomar à noite, com ou sem alimento."},
			{Name: "Clexane", Concentration: "40 mg/0,4 mL", ActiveIngredient: "Enoxaparina sódica",
				Posology: "Aplicar 1 seringa uma vez ao dia, via subcutânea",
				Quantity: "por 7 dias"},
			// Nome longo: o item onde a guia pontilhada pode sumir e o campo da direita colar no
			// medicamento. Se este quebrar, quebra na receita real — "Redoxon Zinco vitamina C 1 g
			// + zinco 10 mg" é um caso que existe em produção.
			{Name: "Redoxon Zinco", Concentration: "vitamina C 1 g + zinco 10 mg",
				ActiveIngredient: "Ácido ascórbico + sulfato de zinco mono-hidratado",
				Posology:         "Tomar 1 comprimido efervescente uma vez ao dia, pela manhã",
				Quantity:         "2 (duas caixas)"},
		},
		GeneralInstructions: "Manter dieta com restrição de sódio e atividade física regular. Retorno em 30 dias com novos exames.",
		ValidUntil:          "10/07/2026",
		Doctor:              Doctor{Name: "Dr. Getúlio José Mattos do Amaral Filho", Credentials: "CRM-PR 21.876 · RQE 16.038 · Nefrologia"},
		Signature: Signature{
			Digital:     true,
			SignedAt:    "10/06/2026, 14:32 (horário de Brasília)",
			ValidateURL: "https://app.plenyasaude.com.br/prescriptions/validate/019eb4a2-7c10-7f3a-9c21-8d4e5b60a1b2",
		},
	}
	// Item 10: onde a indentação por padding quebrava — o número mais largo empurrava o nome e as
	// linhas de baixo ficavam à esquerda dele.
	for len(in.Meds) < 10 {
		in.Meds = append(in.Meds, Med{Name: "Ácido fólico", Concentration: "5 mg",
			Posology: "Tomar 1 comprimido uma vez ao dia", Quantity: "uso contínuo"})
	}

	pdf, err := RenderPrescription(in)
	if err != nil {
		t.Fatalf("render comum: %v", err)
	}
	if len(pdf) < 2000 || string(pdf[:5]) != "%PDF-" {
		t.Fatalf("comum não é PDF (len=%d)", len(pdf))
	}
	_ = os.WriteFile("/tmp/prescription-comum.pdf", pdf, 0o644)

	// Controlado: rótulo + CPF do médico (RDC 1.000/2025) + assinatura manual.
	ctrl := in
	ctrl.ControlLabel = "Receituário de Controle Especial"
	ctrl.Doctor.Credentials = "CRM-PR 21.876 · RQE 16.038 · Nefrologia · CPF 123.456.789-01"
	ctrl.Meds = []Med{{Name: "Clonazepam", Concentration: "2 mg", Posology: "Tomar 1 comprimido ao deitar", Quantity: "30 (trinta) comprimidos"}}
	ctrl.Signature = Signature{Digital: false, PlaceDate: "Londrina, 10 de junho de 2026"}
	pdf2, err := RenderPrescription(ctrl)
	if err != nil {
		t.Fatalf("render controlado: %v", err)
	}
	_ = os.WriteFile("/tmp/prescription-controlada.pdf", pdf2, 0o644)
	t.Logf("OK comum=%d controlada=%d", len(pdf), len(pdf2))
}
