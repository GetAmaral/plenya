package pdfdoc

// Receipt — recibo de pagamento (documento fiscal; identidade fiscal no rodapé, assinatura
// da clínica — NÃO leva selo de assinatura do médico).
type Receipt struct {
	Number      string // "2026-000123"
	PayerName   string // pagador (paciente)
	AmountBRL   string // "R$ 850,00"
	AmountWords string // "oitocentos e cinquenta reais" (valor por extenso)
	Description string // "consulta médica" / "serviços de saúde"
	Method      string // "PIX" / "Cartão" ...
	PaidAt      string // "10/06/2026"
	Notes       string
	Refunded    bool
	RefundNote  string
	PlaceDate   string // "Londrina, 10 de junho de 2026"
	Clinic      Clinic // identidade FISCAL (razão social + CNPJ + endereço fiscal)
}

// RenderReceipt gera o PDF do recibo pelo motor único. O recibo não leva selo do médico: a
// assinatura é a linha da clínica (razão social + CNPJ) — passada como Signature do Doc.
func RenderReceipt(in Receipt) ([]byte, error) {
	if (in.Clinic == Clinic{}) {
		in.Clinic = DefaultClinic()
	}

	// Valor sóbrio à direita do título (não um preço gigante).
	titleRight := `<div class="rcpt-head"><div class="rcpt-eyebrow">Recibo Nº ` + esc(in.Number) + `</div>` +
		`<div class="rcpt-val">` + esc(in.AmountBRL) + `</div></div>`

	// Narrativa clássica do recibo, com valor por extenso.
	amount := "<b>" + esc(in.AmountBRL) + "</b>"
	if in.AmountWords != "" {
		amount += " (" + esc(in.AmountWords) + ")"
	}
	narr := "Recebemos de <b>" + esc(in.PayerName) + "</b> a importância de " + amount
	if in.Description != "" {
		narr += ", referente a " + esc(in.Description)
	}
	if in.Method != "" {
		narr += ", paga via " + esc(in.Method)
	}
	if in.PaidAt != "" {
		narr += " em " + esc(in.PaidAt)
	}
	narr += "."

	body := `<div class="rcpt-body"><p>` + narr + `</p>`
	if in.Notes != "" {
		body += `<p class="rcpt-notes">Observações: ` + esc(in.Notes) + `</p>`
	}
	body += `</div>`
	if in.Refunded {
		r := "Pagamento estornado"
		if in.RefundNote != "" {
			r += " — " + esc(in.RefundNote)
		}
		body += `<div class="rcpt-refund">` + r + `</div>`
	}

	sign := `<div class="rcpt-sign"><div class="rcpt-date">` + esc(in.PlaceDate) + `.</div>` +
		`<div class="rcpt-signline">` + esc(in.Clinic.LegalName) + `<br>` +
		`<span class="rcpt-signsub">CNPJ ` + esc(in.Clinic.CNPJ) + `</span></div></div>`

	return renderDocument(Doc{
		Title:      "Recibo",
		TitleRight: titleRight,
		Body:       body,
		Signature:  sign,
		Clinic:     in.Clinic,
	})
}
