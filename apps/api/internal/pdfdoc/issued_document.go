package pdfdoc

import (
	"regexp"
	"strings"
)

// IssuedDoc — atestado, declaração ou laudo/relatório (corpo em texto livre).
type IssuedDoc struct {
	Title     string // "Atestado Médico" / "Declaração" / "Laudo / Relatório"
	Patient   Patient
	Body      string // texto livre (parágrafos separados por linha em branco)
	CID       string // "" ou "CID-10: J45" (apenas com consentimento — LGPD)
	Doctor    Doctor
	Signature Signature
	Clinic    Clinic
}

var paraSplit = regexp.MustCompile(`\n\s*\n`)

// bodyHTML transforma texto livre em parágrafos (linha em branco = novo parágrafo; \n = <br>).
func bodyHTML(body string) string {
	body = strings.TrimSpace(body)
	if body == "" {
		return ""
	}
	var b strings.Builder
	b.WriteString(`<div class="docbody">`)
	for _, para := range paraSplit.Split(body, -1) {
		para = strings.TrimSpace(para)
		if para == "" {
			continue
		}
		lines := strings.Split(para, "\n")
		for i := range lines {
			lines[i] = esc(strings.TrimSpace(lines[i]))
		}
		b.WriteString(`<p>` + strings.Join(lines, "<br>") + `</p>`)
	}
	b.WriteString(`</div>`)
	return b.String()
}

// RenderIssuedDocument gera o PDF vetorial de atestado/declaração/laudo.
func RenderIssuedDocument(in IssuedDoc) ([]byte, error) {
	if in.Title == "" {
		in.Title = "Documento"
	}
	if (in.Clinic == Clinic{}) {
		in.Clinic = DefaultClinic()
	}
	top := headerHTML() + titleHTML(in.Title) + patientHTML(in.Patient) + bodyHTML(in.Body)
	if in.CID != "" {
		top += `<div class="doccid">` + esc(in.CID) + `</div>`
	}
	foot := signatureHTML(in.Doctor, in.Signature) + footerNAPHTML(in.Clinic)
	return renderHTMLToPDF(documentHTML(pageHTML(top, foot)), a4Options())
}
