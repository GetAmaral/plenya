package pdfdoc

import "strings"

// Med — um medicamento já formatado para exibição.
type Med struct {
	Name             string // "Losartana potássica"
	Concentration    string // "50 mg"
	ActiveIngredient string // princípio ativo (DCB), quando difere do nome comercial
	Posology         string // "Tomar 1 comprimido uma vez ao dia" — a frase indentada da 2ª linha
	Quantity         string // o campo à DIREITA do cabeçalho: "uso contínuo" ou "30 comprimidos"
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
	Name       string // "Fórmula do sono" (opcional)
	Form       string // "cápsula"
	UsageLabel string // "USO INTERNO" / "USO EXTERNO"
	Components []FormulaComponent
	Vehicle    string // "Excipiente qsp 1 cápsula" (opcional)
	// Dispense e Posology são VALORES, sem rótulo: o layout imprime "AVIAR" e "POSOLOGIA" como
	// etiqueta na coluna da esquerda. Mandar "Aviar 60 cápsulas" aqui duplicaria a palavra.
	Dispense     string // "60 (sessenta) cápsulas"
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
		// Cabeçalho em três partes: nome à esquerda, guia pontilhada esticando, e o campo à
		// direita ("uso contínuo" ou a quantidade). A guia é o que impede a leitura ambígua numa
		// receita com vários itens — sem ela, o olho não sabe qual quantidade pertence a qual
		// medicamento.
		// O número fica numa COLUNA própria, e não colado ao nome: as linhas de baixo (princípio
		// ativo, posologia, orientação) se alinham por ela sozinhas. Indentar por padding não
		// resolvia — o padding em `em` vale sobre a fonte de cada linha, e a partir do item 10 o
		// número fica mais largo e as linhas de baixo ficam à esquerda do nome.
		b.WriteString(`<div class="med"><div class="medrow"><span class="mednum">`)
		b.WriteString(itoa(i + 1))
		b.WriteString(`.</span><div class="medbody"><div class="medhead">`)
		b.WriteString(`<span class="mednome">`)
		b.WriteString(esc(m.Name))
		if m.Concentration != "" {
			b.WriteString(` <span class="medconc">` + esc(m.Concentration) + `</span>`)
		}
		b.WriteString(`</span>`)
		if m.Quantity != "" {
			b.WriteString(`<span class="medguia"></span><span class="medqty">` + esc(m.Quantity) + `</span>`)
		}
		b.WriteString(`</div>`)
		if m.ActiveIngredient != "" && m.ActiveIngredient != m.Name {
			b.WriteString(`<div class="medsub">` + esc(m.ActiveIngredient) + `</div>`)
		}
		if m.Posology != "" {
			b.WriteString(`<div class="medpos">` + esc(m.Posology) + `</div>`)
		}
		if m.Instructions != "" {
			b.WriteString(`<div class="medinstr">` + esc(m.Instructions) + `</div>`)
		}
		b.WriteString(`</div></div></div>`) // medbody, medrow, med
	}
	return b.String()
}

// compHTML — uma linha da composição: substância à esquerda, quantidade à direita, pontilhado
// ligando as duas, e a observação EMBAIXO. Sai envolvida no próprio .fcomps para que cada
// componente seja um bloco do paginador sem perder o recuo da coluna do nome.
func compHTML(c FormulaComponent) string {
	qty := esc(c.Quantity)
	if c.AsElemental {
		qty += ` <span class="compelem">(do elemento)</span>`
	}
	s := `<div class="fcomps"><div class="comp"><span class="compname">` + esc(c.Substance) +
		`</span><span class="dots"></span><span class="compqty">` + qty + `</span></div>`
	// A observação vai EMBAIXO, não no meio da linha. Inline, ela empurrava o pontilhado e a
	// quantidade para a direita — "Palmitato de ascorbila <nota de duas linhas> 100 mg" deixava
	// de parecer receituário.
	if c.Note != "" {
		s += `<div class="compnote">` + esc(c.Note) + `</div>`
	}
	return s + `</div>`
}

// formulasHTML — layout clássico do receituário magistral: substância à esquerda, quantidade à
// direita, pontilhado ligando as duas.
//
// Cada fórmula é um container DIVISÍVEL (.split): os filhos é que paginam, e o .formula é reaberto
// na página seguinte com as mesmas classes. Antes ela era um bloco atômico, e por isso a validação
// limitava o número de componentes — mitigação que não bastava: uma fórmula de 20 componentes, o
// próprio teto, saía com a caixa de aviamento por cima da assinatura, em silêncio.
func formulasHTML(formulas []Formula) string {
	var b strings.Builder
	for i, f := range formulas {
		b.WriteString(`<div class="formula split">`)

		// Cabeçalho: número em dourado, nome em serifa, forma farmacêutica embaixo e a tarja de
		// uso à direita. A régua dourada fecha o cabeçalho e abre a composição.
		//
		// O cabeçalho sai GRUDADO no primeiro componente, num bloco só: cabeçalho de fórmula sozinho
		// no pé da página é órfão, e num manipulado a composição da página seguinte passaria a ler
		// como se fosse de outra fórmula.
		b.WriteString(`<div class="fstart">`)
		b.WriteString(`<div class="fhead"><div class="ftitle"><span class="fnum">` + itoa(i+1) + `</span>`)
		name := f.Name
		if name == "" {
			name = f.Form
		}
		b.WriteString(`<span class="fname">` + esc(name) + `</span>`)
		if f.Name != "" && f.Form != "" {
			b.WriteString(`<span class="fform">` + esc(f.Form) + `</span>`)
		}
		b.WriteString(`</div>`)
		if f.UsageLabel != "" {
			b.WriteString(`<div class="fuse">` + esc(f.UsageLabel) + `</div>`)
		}
		b.WriteString(`</div><div class="frule"></div>`)

		// Composição: substância à esquerda, quantidade à direita, pontilhado ligando as duas.
		// O primeiro componente fecha o bloco do cabeçalho; os demais são blocos irmãos, e é entre
		// eles que a fórmula longa quebra de página.
		if len(f.Components) > 0 {
			b.WriteString(compHTML(f.Components[0]))
		}
		b.WriteString(`</div>`) // fecha .fstart
		for _, c := range f.Components[min(1, len(f.Components)):] {
			b.WriteString(compHTML(c))
		}
		if f.Vehicle != "" {
			b.WriteString(`<div class="fcomps"><div class="fveh"><span class="compname">` +
				esc(f.Vehicle) + `</span><span class="dots"></span></div></div>`)
		}

		// Aviamento e posologia num painel próprio: é o que a farmácia e a paciente procuram, e
		// solto no meio do texto virava mais três linhas iguais às outras.
		if f.Dispense != "" || f.Posology != "" || f.Instructions != "" {
			b.WriteString(`<div class="fbox">`)
			if f.Dispense != "" {
				b.WriteString(`<div class="frow"><span class="flabel">Aviar</span><span class="fvalue">` + esc(f.Dispense) + `</span></div>`)
			}
			if f.Posology != "" {
				b.WriteString(`<div class="frow"><span class="flabel">Posologia</span><span class="fvalue">` + esc(f.Posology) + `</span></div>`)
			}
			if f.Instructions != "" {
				b.WriteString(`<div class="finstr">` + esc(f.Instructions) + `</div>`)
			}
			b.WriteString(`</div>`)
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
