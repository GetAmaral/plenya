package cmed

import "testing"

// Linhas REAIS da Lista de Preços da CMED. O texto de apresentação é o único lugar onde a
// concentração e a forma existem — se a derivação erra aqui, o médico digita tudo à mão.
func TestDeriveFromPresentation(t *testing.T) {
	tests := []struct {
		presentation string
		wantConc     string
		wantForm     string
		wantRoute    string
		wantQty      int // 0 = espera nil
		wantConf     string
	}{
		{
			presentation: "500 MG COM REV CT 2 BL AL PLAS TRANS X 07",
			wantConc:     "500 MG", wantForm: "comprimido_revestido", wantRoute: "oral",
			wantQty: 7, wantConf: ConfidenceHigh,
		},
		{
			presentation: "500 MG CAP GEL DURA CT BL AL PVC/PE/PVDC TRANS X 100",
			wantConc:     "500 MG", wantForm: "capsula_dura", wantRoute: "oral",
			wantQty: 100, wantConf: ConfidenceHigh,
		},
		{
			presentation: "2 MG COM CT BL AL PLAS PVDC TRANS X 30",
			wantConc:     "2 MG", wantForm: "comprimido", wantRoute: "oral",
			wantQty: 30, wantConf: ConfidenceHigh,
		},
		{
			presentation: "12,5 MG COM CT BL AL AL X 30",
			wantConc:     "12,5 MG", wantForm: "comprimido", wantRoute: "oral",
			wantQty: 30, wantConf: ConfidenceHigh,
		},
		{
			// Termina em volume, não em contagem: quantidade tem que ficar nil.
			presentation: "30 MG/ML SOL INJ SC CT FA VD TRANS X 0,5 ML",
			wantConc:     "30 MG/ML", wantForm: "solucao_injetavel_sc", wantRoute: "subcutanea",
			wantQty: 0, wantConf: ConfidenceHigh,
		},
		{
			// Associação em parênteses no início: devolve a associação inteira.
			presentation: "(100 + 100)MG/ML SOL INJ IM CX CAMA 3 AMP VD AMB X 1 ML",
			wantConc:     "100 + 100", wantForm: "solucao_injetavel_im", wantRoute: "intramuscular",
			wantQty: 0, wantConf: ConfidenceHigh,
		},
		{
			// Sem concentração no texto (vacina em seringa preenchida) — não inventar.
			presentation: "SUS INJ CT 1 SER PREENC VD TRANS X 0,5 ML",
			wantConc:     "", wantForm: "suspensao_injetavel", wantRoute: "injetavel",
			wantQty: 0, wantConf: ConfidenceMedium,
		},
		{
			// Associação em maiúsculas: a concentração é o par inteiro, não só o 2º ativo.
			presentation: "10 MG/G + 0,443 MG/G CREM DERM CT BG AL X 40 G",
			wantConc:     "10 MG/G + 0,443 MG/G", wantForm: "creme_dermatologico", wantRoute: "topica",
			wantQty: 0, wantConf: ConfidenceHigh,
		},
	}

	for _, tc := range tests {
		t.Run(tc.presentation, func(t *testing.T) {
			got := DeriveFromPresentation(tc.presentation)

			if got.Concentration != tc.wantConc {
				t.Errorf("concentração = %q, esperado %q", got.Concentration, tc.wantConc)
			}
			if got.Form != tc.wantForm {
				t.Errorf("forma = %q, esperado %q", got.Form, tc.wantForm)
			}
			if got.Route != tc.wantRoute {
				t.Errorf("via = %q, esperado %q", got.Route, tc.wantRoute)
			}
			if tc.wantQty == 0 && got.PackageQty != nil {
				t.Errorf("quantidade da embalagem = %d, esperado nenhuma", *got.PackageQty)
			}
			if tc.wantQty != 0 && (got.PackageQty == nil || *got.PackageQty != tc.wantQty) {
				t.Errorf("quantidade da embalagem = %v, esperado %d", got.PackageQty, tc.wantQty)
			}
			if got.Confidence != tc.wantConf {
				t.Errorf("confiança = %q, esperado %q", got.Confidence, tc.wantConf)
			}
		})
	}
}

// Texto vazio não pode virar concentração vazia com confiança alta.
func TestDeriveFromPresentation_Vazio(t *testing.T) {
	got := DeriveFromPresentation("")
	if got.Confidence != ConfidenceNone {
		t.Fatalf("apresentação vazia deveria ter confiança %q, veio %q", ConfidenceNone, got.Confidence)
	}
}
