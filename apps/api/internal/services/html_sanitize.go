package services

import (
	"strings"

	"golang.org/x/net/html"
)

// Sanitização do HTML rich-text do editor (Tiptap) ANTES de gravar/renderizar.
// O HTML vem do cliente e é injetado num Chromium server-side pra gerar o PDF —
// sem allowlist, um <img src="file://...">/<iframe>/<script> viraria leitura de
// arquivo/SSRF. Aqui reescrevemos mantendo só as tags/estilos que o editor produz.

var allowedRichTags = map[string]bool{
	"p": true, "br": true, "strong": true, "b": true, "em": true, "i": true,
	"u": true, "s": true, "strike": true, "del": true,
	"h1": true, "h2": true, "h3": true,
	"ul": true, "ol": true, "li": true, "span": true, "mark": true,
}

// dropSubtreeRichTags — a subárvore inteira (inclusive texto) é descartada.
// dropSubtreeRichTags — containers cujo texto interno também é descartado. NÃO
// inclui void tags (img) nem âncoras: essas caem no strip simples abaixo (a tag
// some, o texto fica), evitando skipDepth preso por falta de fechamento.
var dropSubtreeRichTags = map[string]bool{
	"script": true, "style": true, "iframe": true, "object": true,
	"embed": true, "noscript": true, "template": true, "svg": true,
}

// sanitizeRichHTMLPtr sanitiza um *string de HTML rich-text in place (nil-safe).
// Conveniência para os campos opcionais (ClinicalNote, Anamnesis, …).
func sanitizeRichHTMLPtr(p *string) *string {
	if p == nil {
		return nil
	}
	s := sanitizeRichHTML(*p)
	return &s
}

// safeStyle mantém só color/background-color/text-align, sem url()/expression().
func safeStyle(v string) string {
	var out []string
	for _, decl := range strings.Split(v, ";") {
		decl = strings.TrimSpace(decl)
		i := strings.IndexByte(decl, ':')
		if i < 0 {
			continue
		}
		prop := strings.ToLower(strings.TrimSpace(decl[:i]))
		val := strings.TrimSpace(decl[i+1:])
		switch prop {
		case "color", "background-color", "text-align":
			if strings.ContainsAny(val, "(){}<>\"'") {
				continue // bloqueia url()/expression() e quebra de atributo
			}
			out = append(out, prop+": "+val)
		}
	}
	return strings.Join(out, "; ")
}

// sanitizeRichHTML reescreve o HTML mantendo só a allowlist. Texto é escapado;
// tags desconhecidas somem (conteúdo preservado); subárvores perigosas são
// descartadas por inteiro. Saída vazia → "".
func sanitizeRichHTML(in string) string {
	in = strings.TrimSpace(in)
	if in == "" {
		return ""
	}
	z := html.NewTokenizer(strings.NewReader(in))
	var b strings.Builder
	skipDepth := 0
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return strings.TrimSpace(b.String())
		case html.TextToken:
			if skipDepth == 0 {
				b.WriteString(html.EscapeString(string(z.Text())))
			}
		case html.StartTagToken, html.SelfClosingTagToken:
			nameB, hasAttr := z.TagName()
			tag := string(nameB)
			if dropSubtreeRichTags[tag] {
				// Só abre janela de skip em start tag de container; self-closing
				// (<svg/>) não tem fechamento e prenderia o skipDepth.
				if tt == html.StartTagToken {
					skipDepth++
				}
				continue
			}
			if skipDepth > 0 || !allowedRichTags[tag] {
				continue
			}
			style := ""
			for hasAttr {
				var k, v []byte
				k, v, hasAttr = z.TagAttr()
				if strings.EqualFold(string(k), "style") {
					style = safeStyle(string(v))
				}
			}
			if tag == "br" {
				b.WriteString("<br>")
				continue
			}
			if style != "" {
				b.WriteString("<" + tag + ` style="` + style + `">`)
			} else {
				b.WriteString("<" + tag + ">")
			}
		case html.EndTagToken:
			nameB, _ := z.TagName()
			tag := string(nameB)
			if dropSubtreeRichTags[tag] {
				if skipDepth > 0 {
					skipDepth--
				}
				continue
			}
			if skipDepth > 0 || tag == "br" || !allowedRichTags[tag] {
				continue
			}
			b.WriteString("</" + tag + ">")
		}
	}
}
