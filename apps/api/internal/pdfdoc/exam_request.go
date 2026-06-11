package pdfdoc

import (
	"strconv"
	"strings"
)

// ExamRequest — dados para a Solicitação de Exames (PDF). O título do PDF pode diferir do
// menu do EMR ("Pedido de Exames" no EMR, "Solicitação de Exames" no documento).
type ExamRequest struct {
	Title      string  // default "Solicitação de Exames"
	Patient    Patient
	Indication string  // opcional: vazio não renderiza
	Exams      string  // 1 exame por linha, na ordem registrada; LINHA EM BRANCO = nova página
	Doctor     Doctor
	Signature  Signature
	Clinic     Clinic  // zero value => DefaultClinic()
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

// examPadding — espaçamento dinâmico: generoso com poucos exames, comprime até o mínimo (1.5px)
// com 20+ na coluna mais cheia.
func examPadding(items []string) string {
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

func examColHTML(items []string) string {
	var b strings.Builder
	b.WriteString(`<ul class="excol">`)
	for _, name := range items {
		b.WriteString(`<li class="exitem"><span class="bullet"></span><span class="exname">`)
		b.WriteString(esc(name))
		b.WriteString(`</span></li>`)
	}
	b.WriteString(`</ul>`)
	return b.String()
}

// examListHTML — ≤20 = 1 coluna; 21–40 = 2 colunas com 20 por coluna (preenche a 1ª até 20).
func examListHTML(items []string) string {
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
	pages := pagesFromExams(in.Exams)
	if len(pages) == 0 {
		pages = [][]string{{}}
	}
	indic := indicationHTML(in.Indication)
	foot := signatureHTML(in.Doctor, in.Signature) + footerNAPHTML(in.Clinic)

	var body strings.Builder
	for _, items := range pages {
		top := headerHTML() + titleHTML(in.Title) + patientHTML(in.Patient) + indic +
			`<div class="sec"><span class="eyebrow">Exames solicitados</span>` + examListHTML(items) + `</div>`
		body.WriteString(pageHTML(top, foot))
	}
	return renderHTMLToPDF(documentHTML(body.String()), a4Options())
}
