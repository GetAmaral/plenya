package pdfdoc

import (
	"regexp"
	"strings"
)

// IssuedDoc — atestado, declaração, laudo ou orientações.
type IssuedDoc struct {
	Kind      string // categoria (tarja): "ATESTADO MÉDICO" / "ORIENTAÇÕES" / …
	Title     string // título do documento (digitado pelo médico) — destaque principal
	Patient   Patient
	Body      string // texto livre (fallback dos documentos antigos)
	BodyHTML  string // corpo rich-text JÁ SANITIZADO; quando presente, vence o Body
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

// RenderIssuedDocument gera o PDF vetorial de atestado/declaração/laudo/orientações.
func RenderIssuedDocument(in IssuedDoc) ([]byte, error) {
	if in.Title == "" {
		in.Title = "Documento"
	}
	if (in.Clinic == Clinic{}) {
		in.Clinic = DefaultClinic()
	}
	// Corpo: rich-text sanitizado tem prioridade; senão, texto puro com parágrafos.
	body := bodyHTML(in.Body)
	if strings.TrimSpace(in.BodyHTML) != "" {
		body = `<div class="docbody">` + in.BodyHTML + `</div>`
	}
	main := titleKindHTML(in.Kind, in.Title) + patientHTML(in.Patient) + body
	if in.CID != "" {
		main += `<div class="doccid">` + esc(in.CID) + `</div>`
	}
	// A assinatura é um bloco indivisível (break-inside:avoid): flui após o corpo e,
	// se não couber, abre nova página — em vez de ser cortada como no modelo antigo.
	main += `<div class="sig-block">` + signatureHTML(in.Doctor, in.Signature) + `</div>`
	return renderHTMLToPDF(documentHTML(flowDocHTML(headerHTML(), main, footerNAPHTML(in.Clinic))), a4Options())
}

// flowDocHTML monta um documento de texto livre que pagina sozinho: cabeçalho corrente (thead)
// e rodapé NAP corrente (tfoot) repetidos em toda página pelo Chromium, miolo fluido no tbody,
// e marca d'água/claim em camada fixa repetida por página.
func flowDocHTML(header, main, footer string) string {
	bg := `<div class="run-bg">` + imgSVG("pattern.svg", "wm") +
		`<div class="claim-v">` + esc(claimText) + `</div></div>`
	return bg + `<table class="doc">` +
		`<thead><tr><td>` + header + `</td></tr></thead>` +
		`<tfoot><tr><td>` + footer + `</td></tr></tfoot>` +
		`<tbody><tr><td>` + main + `</td></tr></tbody>` +
		`</table>`
}
