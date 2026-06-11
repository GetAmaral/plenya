package pdfdoc

// Receipt — recibo de pagamento (documento fiscal; identidade fiscal no rodapé, assinatura
// da clínica — NÃO leva selo de assinatura do médico).
type Receipt struct {
	Number      string // "2026-000123"
	PayerName   string // pagador (paciente)
	AmountBRL   string // "R$ 850,00"
	Description string // "serviços de saúde" / "consulta médica"
	Method      string // "PIX" / "Cartão" ...
	PaidAt      string // "10/06/2026"
	Notes       string
	Refunded    bool
	RefundNote  string
	PlaceDate   string // "Londrina, 10 de junho de 2026"
	Clinic      Clinic // identidade FISCAL (razão social + CNPJ + endereço fiscal)
}

// RenderReceipt gera o PDF vetorial do recibo.
func RenderReceipt(in Receipt) ([]byte, error) {
	if (in.Clinic == Clinic{}) {
		in.Clinic = DefaultClinic()
	}
	top := headerHTML() +
		`<div class="titlerow"><div class="title">Recibo</div>` +
		`<div class="docref">Nº ` + esc(in.Number) + `</div></div><div class="title-rule"></div>`

	top += `<div class="rcpt-amount">` + esc(in.AmountBRL) + `</div>`

	body := "Recebemos de <b>" + esc(in.PayerName) + "</b> a importância de <b>" + esc(in.AmountBRL) + "</b>"
	if in.Description != "" {
		body += ", referente a " + esc(in.Description)
	}
	if in.Method != "" {
		body += ", paga via " + esc(in.Method)
	}
	if in.PaidAt != "" {
		body += " em " + esc(in.PaidAt)
	}
	body += "."
	top += `<div class="rcpt-body"><p>` + body + `</p>`
	if in.Notes != "" {
		top += `<p class="rcpt-notes">Observações: ` + esc(in.Notes) + `</p>`
	}
	top += `</div>`
	if in.Refunded {
		r := "Pagamento estornado"
		if in.RefundNote != "" {
			r += " — " + esc(in.RefundNote)
		}
		top += `<div class="rcpt-refund">` + r + `</div>`
	}

	foot := `<div class="rcpt-sign"><div class="rcpt-date">` + esc(in.PlaceDate) + `.</div>` +
		`<div class="rcpt-signline">` + esc(in.Clinic.LegalName) + `<br>` +
		`<span class="rcpt-signsub">CNPJ ` + esc(in.Clinic.CNPJ) + `</span></div></div>` +
		footerNAPHTML(in.Clinic)
	return renderHTMLToPDF(documentHTML(pageHTML(top, foot)), a4Options())
}
