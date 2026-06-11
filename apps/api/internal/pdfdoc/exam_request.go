package pdfdoc

import (
	"strconv"
	"strings"
)

// ExamItem — um exame solicitado já resolvido (Name livre + TUSS opcional, anexado no render).
type ExamItem struct {
	Name string
	Tuss string // "" quando a linha não casa com o catálogo
}

// ExamRequest — dados para a Solicitação de Exames (PDF). O título do PDF pode diferir do
// menu do EMR ("Pedido de Exames" no EMR, "Solicitação de Exames" no documento).
type ExamRequest struct {
	Title      string // default "Solicitação de Exames"
	Patient    Patient
	Indication string       // opcional: vazio não renderiza
	ExamPages  [][]ExamItem // páginas (quebra por linha em branco / >40); TUSS resolvido por quem monta
	Doctor     Doctor
	Signature  Signature
	Clinic     Clinic // zero value => DefaultClinic()
}

const maxPerPage = 40 // ≤20 = 1 col; 21–40 = 2 col (20/col); >40 = nova página

// pagesFromExams quebra por linha em branco E a cada 40 exames (2 col × 20).
func pagesFromExams(text string) [][]string {
	var blocks [][]string
	var cur []string
	for _, ln := range strings.Split(text, "\n") {
		t := strings.TrimSpace(ln)
		if t == "" {
			if len(cur) > 0 {
				blocks = append(blocks, cur)
				cur = nil
			}
			continue
		}
		cur = append(cur, t)
	}
	if len(cur) > 0 {
		blocks = append(blocks, cur)
	}
	var pages [][]string
	for _, b := range blocks {
		for i := 0; i < len(b); i += maxPerPage {
			end := i + maxPerPage
			if end > len(b) {
				end = len(b)
			}
			pages = append(pages, b[i:end])
		}
	}
	return pages
}

// ExamPagesFromText — conveniência: texto livre (1 exame/linha, linha em branco = nova página)
// para páginas de ExamItem SEM TUSS. O adaptador resolve o TUSS depois (LabTestMatcher),
// preservando a caixa de texto 100% livre.
func ExamPagesFromText(text string) [][]ExamItem {
	src := pagesFromExams(text)
	out := make([][]ExamItem, len(src))
	for i, pg := range src {
		items := make([]ExamItem, len(pg))
		for j, name := range pg {
			items[j] = ExamItem{Name: name}
		}
		out[i] = items
	}
	return out
}

// examPadding — espaçamento dinâmico: generoso com poucos exames, mínimo (1.5px) com 20+ na coluna.
func examPadding(items []ExamItem) string {
	const minPad, maxPad = 1.5, 7.0
	ctrl := len(items)
	if ctrl > 20 {
		ctrl = 20
	}
	pad := minPad + (maxPad-minPad)*float64(20-ctrl)/19.0
	if pad < minPad {
		pad = minPad
	}
	if pad > maxPad {
		pad = maxPad
	}
	return strconv.FormatFloat(pad, 'f', 2, 64)
}

func examColHTML(items []ExamItem) string {
	var b strings.Builder
	b.WriteString(`<ul class="excol">`)
	for _, it := range items {
		b.WriteString(`<li class="exitem"><span class="bullet"></span><span class="exname">`)
		b.WriteString(esc(it.Name))
		if it.Tuss != "" {
			b.WriteString(`<span class="extuss"> (` + esc(it.Tuss) + `)</span>`)
		}
		b.WriteString(`</span></li>`)
	}
	b.WriteString(`</ul>`)
	return b.String()
}

// examListHTML — ≤20 = 1 coluna; 21–40 = 2 colunas com 20 por coluna (preenche a 1ª até 20).
func examListHTML(items []ExamItem) string {
	pad := examPadding(items)
	var cols string
	if len(items) <= 20 {
		cols = examColHTML(items)
	} else {
		cols = examColHTML(items[:20]) + examColHTML(items[20:])
	}
	return `<div class="exwrap" style="--expad:` + pad + `px">` + cols + `</div>`
}

// RenderExamRequest gera o PDF vetorial da Solicitação de Exames.
func RenderExamRequest(in ExamRequest) ([]byte, error) {
	if in.Title == "" {
		in.Title = "Solicitação de Exames"
	}
	if (in.Clinic == Clinic{}) {
		in.Clinic = DefaultClinic()
	}
	if len(in.ExamPages) == 0 {
		in.ExamPages = [][]ExamItem{{}}
	}
	indic := indicationHTML(in.Indication)
	foot := signatureHTML(in.Doctor, in.Signature) + footerNAPHTML(in.Clinic)

	var body strings.Builder
	for _, items := range in.ExamPages {
		top := headerHTML() + titleHTML(in.Title) + patientHTML(in.Patient) + indic +
			`<div class="sec"><span class="eyebrow">Exames solicitados</span>` + examListHTML(items) + `</div>`
		body.WriteString(pageHTML(top, foot))
	}
	return renderHTMLToPDF(documentHTML(body.String()), a4Options())
}
