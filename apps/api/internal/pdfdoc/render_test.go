package pdfdoc

import (
	"os"
	"strings"
	"testing"
)

func sampleExamsText() string {
	return strings.Join([]string{
		"Hemograma completo com plaquetas", "Glicemia de jejum", "Insulina de jejum",
		"Hemoglobina glicada (HbA1c)", "Colesterol total e frações (HDL, LDL)", "Triglicérides",
		"Apolipoproteína B", "Lipoproteína (a)", "TSH", "T4 livre", "T3 livre",
		"Creatinina com taxa de filtração glomerular estimada", "Ureia", "Ácido úrico",
		"Cistatina C", "TGO (AST)", "TGP (ALT)", "Gama-GT", "Fosfatase alcalina",
		"Bilirrubinas total e frações", "Proteínas totais e frações", "Albumina",
		"Vitamina D (25-hidroxivitamina D)", "Vitamina B12", "Ácido fólico", "Ferritina",
		"Ferro sérico e saturação de transferrina", "Zinco sérico", "Magnésio", "Cálcio",
		"Fósforo", "Sódio", "Potássio", "Proteína C reativa ultrassensível", "Homocisteína",
		"VHS (velocidade de hemossedimentação)", "Cortisol matinal", "Testosterona total e livre",
		"Microalbuminúria em amostra isolada", "Urina tipo I (EAS)",
		"", // quebra de página: laboratório | imagem
		"Ultrassonografia de abdome total", "Ecocardiograma transtorácico com Doppler",
		"Densitometria óssea (coluna lombar e fêmur)", "Tomografia computadorizada de tórax de baixa dose",
	}, "\n")
}

func sampleExamRequest() ExamRequest {
	return ExamRequest{
		Patient:    Patient{Name: "Maria Helena Soares", BirthInfo: "12/03/1979 · 47 anos", CPFMasked: "***.456.789-**"},
		Indication: "Avaliação metabólica e de longevidade, rastreio do pilar Gestão do Método AGIR. Paciente assintomática, em programa de acompanhamento contínuo.",
		ExamPages:  ExamPagesFromText(sampleExamsText()),
		Doctor:     Doctor{Name: "Dr. Getúlio José Mattos do Amaral Filho", Credentials: "CRM-PR 21.876 · RQE 16.038 · Nefrologia"},
		Signature: Signature{
			Digital:     true,
			SignedAt:    "10/06/2026, 14:32 (horário de Brasília)",
			ValidateURL: "https://app.plenyasaude.com.br/lab-requests/validate/019eb4a2-7c10-7f3a-9c21-8d4e5b60a1b2",
		},
	}
}

func chromiumAvailable() bool {
	if _, err := os.Stat("/usr/bin/chromium-browser"); err == nil {
		return true
	}
	if _, err := os.Stat("/usr/bin/chromium"); err == nil {
		return true
	}
	return false
}

func TestRenderExamRequest(t *testing.T) {
	if !chromiumAvailable() {
		t.Skip("chromium ausente — pulando render")
	}
	pdf, err := RenderExamRequest(sampleExamRequest())
	if err != nil {
		t.Fatalf("render: %v", err)
	}
	if len(pdf) < 2000 || string(pdf[:5]) != "%PDF-" {
		t.Fatalf("saída não parece PDF (len=%d)", len(pdf))
	}
	_ = os.WriteFile("/tmp/exam-request-test.pdf", pdf, 0o644)
	t.Logf("PDF gerado: %d bytes -> /tmp/exam-request-test.pdf", len(pdf))
}

// TestExamPaginationLogic valida paginação/colunas sem precisar de Chromium.
func TestExamPaginationLogic(t *testing.T) {
	pages := ExamPagesFromText(sampleExamsText())
	if len(pages) != 2 {
		t.Fatalf("esperava 2 páginas (40 lab + bloco imagem), veio %d", len(pages))
	}
	if len(pages[0]) != 40 || len(pages[1]) != 4 {
		t.Fatalf("contagem por página errada: %d, %d", len(pages[0]), len(pages[1]))
	}
	if examPadding(pages[0]) != "1.50" {
		t.Fatalf("40 exames deveriam dar padding mínimo 1.50, veio %s", examPadding(pages[0]))
	}
}
