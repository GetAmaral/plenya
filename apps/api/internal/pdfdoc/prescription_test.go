package pdfdoc

import (
	"bytes"
	"os"
	"strings"
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

// formulaLonga — o teto da validação (20 componentes), que é o caso que saía quebrado.
func formulaLonga(n int) Prescription {
	comps := make([]FormulaComponent, 0, n)
	for i := 0; i < n; i++ {
		comps = append(comps, FormulaComponent{
			Substance: "Substância de teste com nome longo " + itoa(i+1),
			Quantity:  itoa((i+1)*25) + " mg",
		})
	}
	return Prescription{
		Compounded: true,
		Patient:    Patient{Name: "Paciente Teste"},
		Formulas: []Formula{{
			Name: "Fórmula longa", Form: "cápsula", UsageLabel: "USO INTERNO",
			Components: comps, Vehicle: "Excipiente qsp 1 cápsula",
			Dispense: "60 (sessenta) cápsulas", Posology: "1 cápsula ao deitar",
		}},
		Doctor:    Doctor{Name: "Dr. Teste", Credentials: "CRM-PR 12345"},
		Signature: Signature{Digital: true, ValidateURL: "https://app.plenyasaude.com.br/x"},
	}
}

// TestFormulaEhContainerDivisivel — a fórmula precisa ser .split (container divisível) e emitir um
// bloco por componente. Como bloco atômico ela transbordava por cima da assinatura em silêncio.
func TestFormulaEhContainerDivisivel(t *testing.T) {
	html := formulasHTML(formulaLonga(20).Formulas)
	if !strings.Contains(html, `class="formula split"`) {
		t.Error("fórmula não está marcada como container divisível (.split): volta a ser bloco atômico")
	}
	if !strings.Contains(html, `class="fstart"`) {
		t.Error("cabeçalho não está grudado no primeiro componente: pode ficar órfão no pé da página")
	}
	// 20 componentes + o veículo, cada um no seu .fcomps (bloco do paginador).
	if got := strings.Count(html, `class="fcomps"`); got != 21 {
		t.Errorf("esperava 21 blocos .fcomps (20 componentes + veículo), obtive %d", got)
	}
}

// TestRenderFormulaLongaQuebraEmVezDeTransbordar — o teste antigo afirmava no comentário que 20
// componentes cabiam numa página e só checava se a saída era um PDF válido; não cabiam, e a caixa
// de aviamento saía por cima da assinatura. Agora tem de quebrar em 2 páginas, sem erro.
func TestRenderFormulaLongaQuebraEmVezDeTransbordar(t *testing.T) {
	if !chromiumAvailable() {
		t.Skip("chromium ausente — pulando render")
	}
	b, err := RenderPrescription(formulaLonga(20))
	if err != nil {
		t.Fatalf("fórmula de 20 componentes não renderizou: %v", err)
	}
	// Não conto páginas por bytes do PDF: "/Type /Page\n" depende da serialização exata do
	// Chromium, e um upgrade que emita object streams faria o teste falhar num documento correto.
	// Que a fórmula DIVIDE está travado estruturalmente em TestFormulaEhContainerDivisivel; aqui o
	// que importa é que 20 componentes não são mais recusados pela rede de transbordo.
	if len(b) < 1000 || !bytes.HasPrefix(b, []byte("%PDF")) {
		t.Fatalf("saída não parece um PDF válido (%d bytes)", len(b))
	}
}
