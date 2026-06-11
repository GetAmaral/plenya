package pdfdoc

import (
	"os"
	"testing"
)

func TestRenderCarePlanReport(t *testing.T) {
	if !chromiumAvailable() {
		t.Skip("chromium ausente")
	}
	in := CarePlanReport{
		Patient:     Patient{Name: "Maria Helena Soares"},
		EmittedAt:   "10/06/2026", ScoreCalcAt: "08/06/2026", ScorePct: "78%",
		Attention: []ReportBiomarker{
			{Name: "Vitamina D (25-OH)", Label: "Subótimo", Kind: "attention"},
			{Name: "Insulina de jejum", Label: "Muito alterado", Kind: "critical"},
			{Name: "Ferritina", Label: "Crítico", Kind: "critical"},
			{Name: "PCR ultrassensível", Label: "Subótimo", Kind: "attention"},
		},
		Optimal: []ReportBiomarker{
			{Name: "HDL", Label: "Ótimo", Kind: "optimal"},
			{Name: "TSH", Label: "Bom", Kind: "good"},
			{Name: "Função renal (TFG)", Label: "Ótimo", Kind: "optimal"},
		},
		Pillars: []ReportPillar{
			{Letter: "A", Name: "Atividade", Recs: []ReportRec{
				{Text: "Treino de força 3x/semana, progressivo.", Target: "≥150 min/sem"},
				{Text: "Caminhada diária pós-refeições."},
			}},
			{Letter: "G", Name: "Gestão", Recs: []ReportRec{
				{Text: "Repor vitamina D com colecalciferol.", Target: "25-OH > 40 ng/mL"},
				{Text: "Reavaliar insulina e HOMA-IR em 90 dias."},
			}},
			{Letter: "I", Name: "Integração", Recs: []ReportRec{
				{Text: "Prática de respiração/meditação 10 min/dia."},
			}},
			{Letter: "R", Name: "Ritmo", Recs: []ReportRec{
				{Text: "Higiene do sono: janela fixa, sem telas 1h antes.", Target: "7-8h/noite"},
			}},
		},
		Doctor: Doctor{Name: "Dr. Getúlio José Mattos do Amaral Filho", Credentials: "CRM-PR 21.876 · RQE 16.038 · Nefrologia"},
		Signature: Signature{Digital: true, SignedAt: "10/06/2026, 14:32 (horário de Brasília)",
			ValidateURL: "https://app.plenyasaude.com.br/documents/validate/019eb4a2-7c10-7f3a-9c21-8d4e5b60a1b2"},
	}
	pdf, err := RenderCarePlanReport(in)
	if err != nil {
		t.Fatalf("render: %v", err)
	}
	if len(pdf) < 2000 || string(pdf[:5]) != "%PDF-" {
		t.Fatalf("não é PDF (len=%d)", len(pdf))
	}
	_ = os.WriteFile("/tmp/relatorio-agir.pdf", pdf, 0o644)
	t.Logf("OK: %d bytes", len(pdf))
}
