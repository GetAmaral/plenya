package pdfdoc

import "strings"

// Med — um medicamento já formatado para exibição.
type Med struct {
	Name            string // "Losartana potássica"
	Concentration   string // "50 mg"
	ActiveIngredient string // princípio ativo (DCB), quando difere do nome comercial
	Posology        string // "1 comprimido · de 12/12h · via oral · por 30 dias"
	Quantity        string // "Quantidade: 60 (sessenta) comprimidos"
	Instructions    string // orientação específica (opcional)
}

// Prescription — dados para o Receituário (PDF).
type Prescription struct {
	Title               string // default "Receituário"
	ControlLabel        string // "" ou "Receituário de Controle Especial" (c1/c5)
	Patient             Patient
	Meds                []Med
	GeneralInstructions string // orientações gerais (opcional)
	ValidUntil          string // "10/07/2026" (opcional)
	Doctor              Doctor
	Signature           Signature
	Clinic              Clinic
}

// medsHTML — cada medicamento é um bloco .med independente (para paginar pelo motor único).
func medsHTML(meds []Med) string {
	var b strings.Builder
	for i, m := range meds {
		b.WriteString(`<div class="med"><div class="medhead">`)
		b.WriteString(itoa(i + 1))
		b.WriteString(`. `)
		b.WriteString(esc(m.Name))
		if m.Concentration != "" {
			b.WriteString(` <span class="medconc">` + esc(m.Concentration) + `</span>`)
		}
		b.WriteString(`</div>`)
		if m.ActiveIngredient != "" && m.ActiveIngredient != m.Name {
			b.WriteString(`<div class="medsub">` + esc(m.ActiveIngredient) + `</div>`)
		}
		if m.Posology != "" {
			b.WriteString(`<div class="medpos">` + esc(m.Posology) + `</div>`)
		}
		if m.Quantity != "" {
			b.WriteString(`<div class="medqty">` + esc(m.Quantity) + `</div>`)
		}
		if m.Instructions != "" {
			b.WriteString(`<div class="medinstr">` + esc(m.Instructions) + `</div>`)
		}
		b.WriteString(`</div>`)
	}
	return b.String()
}

// RenderPrescription gera o PDF do Receituário pelo motor único.
func RenderPrescription(in Prescription) ([]byte, error) {
	if in.Title == "" {
		in.Title = "Receituário"
	}
	var body strings.Builder
	if in.ControlLabel != "" {
		body.WriteString(`<div class="ctrltag">` + esc(in.ControlLabel) + `</div>`)
	}
	body.WriteString(`<div class="sec"><span class="eyebrow">Prescrição</span></div>`)
	body.WriteString(medsHTML(in.Meds))
	if in.GeneralInstructions != "" {
		body.WriteString(`<div class="indic"><span class="eyebrow">Orientações gerais</span>` + esc(in.GeneralInstructions) + `</div>`)
	}
	if in.ValidUntil != "" {
		body.WriteString(`<div class="validity">Válida até <b>` + esc(in.ValidUntil) + `</b>.</div>`)
	}
	return renderDocument(Doc{
		Title:     in.Title,
		Patient:   &in.Patient,
		Body:      body.String(),
		Signature: signatureBlock(in.Doctor, in.Signature),
		Clinic:    in.Clinic,
	})
}
