package pdfdoc

import "strings"

// Med — um medicamento já formatado para exibição.
type Med struct {
	Name             string // "Losartana potássica"
	Concentration    string // "50 mg"
	ActiveIngredient string // princípio ativo (DCB), quando difere do nome comercial
	Posology         string // "1 comprimido · de 12/12h · via oral · por 30 dias"
	Quantity         string // "Quantidade: 60 (sessenta) comprimidos"
	Instructions     string // orientação específica (opcional)
}

// FormulaComponent — uma substância da fórmula, já formatada.
type FormulaComponent struct {
	Substance string // "Melatonina"
	Quantity  string // "3 mg"
	Note      string // "liberação prolongada" (opcional)
	// AsElemental — a dose escrita é do ELEMENTO, não do insumo. Precisa sair impresso: sem essa
	// palavra, "magnésio quelato 150 mg" manda a farmácia pesar 150 mg do bisglicinato, que são
	// 45 mg de magnésio. O documento é assinado e é por ele que a farmácia manipula.
	AsElemental bool
}

// Formula — uma fórmula magistral já formatada para impressão.
type Formula struct {
	Name         string // "Fórmula do sono" (opcional)
	Form         string // "cápsula"
	UsageLabel   string // "USO INTERNO" / "USO EXTERNO"
	Components   []FormulaComponent
	Vehicle      string // "Excipiente qsp 1 cápsula" (opcional)
	Dispense     string // "Aviar 60 (sessenta) cápsulas"
	Posology     string // "1 cápsula ao deitar · por 60 dias"
	Instructions string // orientação específica (opcional)
}

// Prescription — dados para o Receituário (PDF).
type Prescription struct {
	// Compounded — receituário magistral. É por AQUI que o layout decide, não por "tem fórmula":
	// uma receita que chegasse com as duas listas renderizava só as fórmulas e descartava os
	// industrializados em silêncio, num documento assinado.
	Compounded bool

	Title               string // default "Receituário"
	ControlLabel        string // "" ou "Receituário de Controle Especial" (c1/c5)
	Patient             Patient
	Meds                []Med
	Formulas            []Formula
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

// formulasHTML — layout clássico do receituário magistral: substância à esquerda, quantidade à
// direita, pontilhado ligando as duas. Cada fórmula é UM bloco (unidade atômica de paginação).
func formulasHTML(formulas []Formula) string {
	var b strings.Builder
	for i, f := range formulas {
		b.WriteString(`<div class="formula"><div class="fhead"><div class="fname">`)
		b.WriteString(itoa(i + 1))
		b.WriteString(`. `)
		if f.Name != "" {
			b.WriteString(esc(f.Name))
			if f.Form != "" {
				b.WriteString(` <span class="fform">` + esc(f.Form) + `</span>`)
			}
		} else if f.Form != "" {
			b.WriteString(esc(f.Form))
		}
		b.WriteString(`</div>`)
		if f.UsageLabel != "" {
			b.WriteString(`<div class="fuse">` + esc(f.UsageLabel) + `</div>`)
		}
		b.WriteString(`</div>`)

		for _, c := range f.Components {
			qty := esc(c.Quantity)
			if c.AsElemental {
				qty += ` <span class="compelem">(do elemento)</span>`
			}
			b.WriteString(`<div class="comp"><span class="compname">` + esc(c.Substance) +
				`</span><span class="dots"></span><span class="compqty">` + qty + `</span></div>`)
			// A observação vai EMBAIXO, não no meio da linha. Inline, ela empurrava o pontilhado e
			// a quantidade para a direita — "Palmitato de ascorbila <nota de duas linhas> 100 mg"
			// deixava de parecer receituário.
			if c.Note != "" {
				b.WriteString(`<div class="compnote">` + esc(c.Note) + `</div>`)
			}
		}
		if f.Vehicle != "" {
			b.WriteString(`<div class="fveh"><span class="compname">` + esc(f.Vehicle) + `</span><span class="dots"></span></div>`)
		}
		if f.Dispense != "" {
			b.WriteString(`<div class="fdisp">` + esc(f.Dispense) + `</div>`)
		}
		if f.Posology != "" {
			b.WriteString(`<div class="fpos">` + esc(f.Posology) + `</div>`)
		}
		if f.Instructions != "" {
			b.WriteString(`<div class="finstr">` + esc(f.Instructions) + `</div>`)
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
	if in.Compounded || (len(in.Formulas) > 0 && len(in.Meds) == 0) {
		body.WriteString(`<div class="sec"><span class="eyebrow">Prescrição magistral</span></div>`)
		body.WriteString(formulasHTML(in.Formulas))
	} else {
		body.WriteString(`<div class="sec"><span class="eyebrow">Prescrição</span></div>`)
		body.WriteString(medsHTML(in.Meds))
	}
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
