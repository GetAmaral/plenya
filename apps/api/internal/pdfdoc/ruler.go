package pdfdoc

import (
	"fmt"
	"hash/fnv"
	"math"
	"sort"
	"strconv"
	"strings"
)

// Régua de faixa — o átomo visual da devolutiva ao paciente.
//
// A régua inteira é a escala Plenya daquele exame: uma rampa contínua do pior nível ao ótimo, com
// o ponto marcando onde o paciente está e, quando há valor anterior, uma seta do antigo para o
// atual. Ela substitui a tabela de resultado com faixa de referência do laboratório, e as duas
// escolhas de desenho são deliberadas:
//
//   - mostra a faixa-meta NO LUGAR da faixa de referência, não somada a ela;
//   - mostra a DIREÇÃO da mudança com uma seta, porque direção é o que o leitor mais erra e o que
//     mais importa clinicamente.
//
// Port de `pacs/<paciente>/deck/ruler.py`, que era copiado de paciente para paciente.

// rulerRamp — rampa contínua do nível 0 (pior, terracota) ao 5 (ótimo, petróleo).
var rulerRamp = [6]string{"#B3503C", "#CD8674", "#E1C6B9", "#AFC9D5", "#5D93AC", "#0E4C6B"}

const (
	rulerInk   = "#0A1F26"
	rulerMuted = "#5A6B70"
	rulerCream = "#F5F1E8"

	rulerDarkText  = "#EAE7DA"
	rulerDarkMuted = "#92B8B4"
	rulerDarkSurf  = "#063B4F"
)

// Geometria, em px do slide de 1920×1080.
const (
	rulerNameW  = 450.0  // coluna do nome, alinhada à direita
	rulerTrackX = 490.0  // onde a barra começa
	rulerTrackW = 1000.0 // largura da barra
	rulerValueW = 190.0  // coluna do valor, à direita
	rulerNoteX  = 200.0  // a nota recua para caber na largura do slide
	rulerCY     = 62.0   // centro vertical da barra
	rulerH      = 36.0   // altura da barra
	rulerRowH   = 132.0  // altura da linha inteira
	rulerNoteH  = 34.0   // acréscimo quando há nota
)

// RulerPoint — um resultado do paciente na régua.
type RulerPoint struct {
	Value float64 `json:"value"`
	Text  string  `json:"text"` // valor já formatado ("1,023")
	Date  string  `json:"date,omitempty"`
}

// RulerSegment — uma faixa de nível.
type RulerSegment struct {
	Level int     `json:"level"`
	A     float64 `json:"a"`
	B     float64 `json:"b"`
}

// Ruler — tudo que a régua precisa para se desenhar.
// As tags json não são decoração: esta struct é serializada dentro do `patient_plans.content` e
// aparece nos tipos TS gerados. Sem tag, o Go grava "Display"/"Segments" e o front tipado lê
// `display`/`segments` — e recebe undefined em silêncio.
type Ruler struct {
	Code     string         `json:"code"`
	Display  string         `json:"display"` // nome que o paciente lê ("Ferritina"), não o do catálogo
	Sub      string         `json:"sub,omitempty"`
	Unit     string         `json:"unit,omitempty"`
	Axis     [2]float64     `json:"axis"`
	Segments []RulerSegment `json:"segments"`
	History  []RulerPoint   `json:"history,omitempty"`

	// Note — uma linha de leitura clínica embaixo da régua. É também onde mora o RÓTULO
	// AVALIATIVO quando o slide não o traz no título: barra colorida sem rótulo tem desempenho
	// pior do que barra colorida com rótulo.
	Note string `json:"note,omitempty"`
	// Dark — derivado da variante do slide na hora de desenhar, nunca gravado: se fosse
	// persistido, um slide que muda de fundo deixaria réguas com a cor do fundo antigo.
	Dark bool `json:"-"`
}

// rulerSVG desenha a régua. Devolve um <svg> autossuficiente, sem depender de JS.
func rulerSVG(r Ruler) string {
	lo, hi := r.Axis[0], r.Axis[1]
	span := hi - lo
	if span == 0 {
		span = 1
	}

	text, muted, surf := rulerInk, rulerMuted, rulerCream
	if r.Dark {
		text, muted, surf = rulerDarkText, rulerDarkMuted, rulerDarkSurf
	}

	// x() posiciona um valor no eixo, preso às bordas.
	x := func(v float64) float64 {
		return rulerTrackX + (math.Min(math.Max(v, lo), hi)-lo)/span*rulerTrackW
	}

	top, bot := rulerCY-rulerH/2, rulerCY+rulerH/2
	rowH := rulerRowH
	if r.Note != "" {
		rowH += rulerNoteH
	}
	width := rulerTrackX + rulerTrackW + 34 + rulerValueW
	clipID := "rgc" + shortHash(r.Code+r.Display)

	var b strings.Builder
	fmt.Fprintf(&b, `<svg class="rg" width="%s" height="%s" viewBox="0 0 %s %s" xmlns="http://www.w3.org/2000/svg">`,
		num(width), num(rowH), num(width), num(rowH))
	fmt.Fprintf(&b, `<defs><clipPath id="%s"><rect x="%s" y="%s" width="%s" height="%s" rx="%s"/></clipPath></defs>`,
		clipID, num(rulerTrackX), num(top), num(rulerTrackW), num(rulerH), num(rulerH/2))

	// nome e subtítulo, alinhados à direita da barra
	fmt.Fprintf(&b, `<text x="%s" y="%s" text-anchor="end" class="rg-name" fill="%s">%s</text>`,
		num(rulerNameW), num(rulerCY+2), text, esc(r.Display))
	if r.Sub != "" {
		fmt.Fprintf(&b, `<text x="%s" y="%s" text-anchor="end" class="rg-sub" fill="%s">%s</text>`,
			num(rulerNameW), num(rulerCY+30), muted, esc(r.Sub))
	}

	// a rampa inteira, do pior ao ótimo
	fmt.Fprintf(&b, `<g clip-path="url(#%s)">`, clipID)
	fmt.Fprintf(&b, `<rect x="%s" y="%s" width="%s" height="%s" fill="%s"/>`,
		num(rulerTrackX), num(top), num(rulerTrackW), num(rulerH), rulerRamp[0])
	segs := append([]RulerSegment(nil), r.Segments...)
	sort.Slice(segs, func(i, j int) bool { return segs[i].A < segs[j].A })
	for _, s := range segs {
		a, bb := math.Max(s.A, lo), math.Min(s.B, hi)
		if bb <= a {
			continue
		}
		xa, xb := x(a), x(bb)
		fmt.Fprintf(&b, `<rect x="%s" y="%s" width="%s" height="%s" fill="%s"/>`,
			num(xa), num(top), num(math.Max(xb-xa, 2)), num(rulerH), rampColor(s.Level))
	}
	b.WriteString(`</g>`)

	// fronteiras, discretas: só as que cabem sem colidir
	last := -999.0
	for _, e := range segmentEdges(segs) {
		if e <= lo || e >= hi {
			continue
		}
		xe := x(e)
		if xe < rulerTrackX+26 || xe > rulerTrackX+rulerTrackW-26 {
			continue
		}
		if xe-last < 70 {
			continue
		}
		last = xe
		fmt.Fprintf(&b, `<text x="%s" y="%s" text-anchor="middle" class="rg-tick" fill="%s">%s</text>`,
			num(xe), num(bot+30), muted, esc(formatRulerNumber(e)))
	}

	// os pontos: no máximo os dois últimos, com seta do antigo para o atual
	pts := r.History
	if len(pts) > 2 {
		pts = pts[len(pts)-2:]
	}
	var curX float64
	var cur *RulerPoint

	if len(pts) == 2 {
		x0, x1 := x(pts[0].Value), x(pts[1].Value)
		d := 1.0
		if x1 < x0 {
			d = -1
		}
		if math.Abs(x1-x0) > 46 {
			body := fmt.Sprintf("M%s %s H%s", num(x0+d*15), num(rulerCY), num(x1-d*30))
			head := fmt.Sprintf("M%s %s l%s -10 v20 z", num(x1-d*18), num(rulerCY), num(-d*17))
			// contorno na cor da superfície: a seta atravessa a rampa e precisa se destacar
			// de qualquer faixa por baixo.
			fmt.Fprintf(&b, `<path d="%s" stroke="%s" stroke-width="11" fill="none" stroke-linecap="round"/>`, body, surf)
			fmt.Fprintf(&b, `<path d="%s" fill="%s" stroke="%s" stroke-width="6" stroke-linejoin="round"/>`, head, surf, surf)
			fmt.Fprintf(&b, `<path d="%s" stroke="%s" stroke-width="4" fill="none" stroke-linecap="round"/>`, body, text)
			fmt.Fprintf(&b, `<path d="%s" fill="%s"/>`, head, text)
		}
		fmt.Fprintf(&b, `<circle cx="%s" cy="%s" r="12" fill="%s" stroke="%s" stroke-width="4"/>`,
			num(x0), num(rulerCY), surf, text)
		fmt.Fprintf(&b, `<text x="%s" y="%s" text-anchor="middle" class="rg-old" fill="%s">%s</text>`,
			num(x0), num(top-14), text, esc(pts[0].Text))
		curX, cur = x1, &pts[1]
	} else if len(pts) == 1 {
		curX, cur = x(pts[0].Value), &pts[0]
	}

	if cur != nil {
		fmt.Fprintf(&b, `<circle cx="%s" cy="%s" r="16.5" fill="%s"/>`, num(curX), num(rulerCY), surf)
		fmt.Fprintf(&b, `<circle cx="%s" cy="%s" r="12.5" fill="%s"/>`, num(curX), num(rulerCY), text)
		vx := rulerTrackX + rulerTrackW + 34
		fmt.Fprintf(&b, `<text x="%s" y="%s" class="rg-val" fill="%s">%s</text>`,
			num(vx), num(rulerCY+4), text, esc(cur.Text))
		fmt.Fprintf(&b, `<text x="%s" y="%s" class="rg-unit" fill="%s">%s</text>`,
			num(vx), num(rulerCY+32), muted, esc(r.Unit))
	}

	if r.Note != "" {
		fmt.Fprintf(&b, `<text x="%s" y="%s" class="rg-note" fill="%s">%s</text>`,
			num(rulerNoteX), num(rowH-12), muted, esc(r.Note))
	}
	b.WriteString(`</svg>`)
	return b.String()
}

// Mini-régua do resumo executivo: mesma escala, mesmas cores, MESMA ordem no eixo.
const (
	miniW = 264.0
	miniH = 22.0
	// Numa barra de 264px uma faixa que ocupa 3% do eixo some. Cada faixa recebe largura
	// proporcional ao seu vão, presa entre 8% e 30% do total, e depois renormalizada: ordem,
	// direção e formato ficam intactos, só a escala interna é comprimida para caber.
	miniMinFrac = 0.08
	miniMaxFrac = 0.30
)

// miniRulerSVG desenha a régua compacta usada nos cartões do resumo.
func miniRulerSVG(r Ruler, value float64, hasValue bool) string {
	lo, hi := r.Axis[0], r.Axis[1]
	span := hi - lo
	if span == 0 {
		span = 1
	}
	segs := make([]RulerSegment, 0, len(r.Segments))
	for _, s := range r.Segments {
		a, b := math.Max(s.A, lo), math.Min(s.B, hi)
		if b > a {
			segs = append(segs, RulerSegment{Level: s.Level, A: a, B: b})
		}
	}
	if len(segs) == 0 {
		return ""
	}
	sort.Slice(segs, func(i, j int) bool { return segs[i].A < segs[j].A })

	fracs := make([]float64, len(segs))
	total := 0.0
	for i, s := range segs {
		f := math.Min(math.Max((s.B-s.A)/span, miniMinFrac), miniMaxFrac)
		fracs[i] = f
		total += f
	}
	if total == 0 {
		total = 1
	}

	clipID := "mnc" + shortHash(r.Code+"mini")
	var b strings.Builder
	fmt.Fprintf(&b, `<svg class="mini" width="%s" height="%s" viewBox="0 0 %s %s" xmlns="http://www.w3.org/2000/svg">`,
		num(miniW), num(miniH), num(miniW), num(miniH))
	fmt.Fprintf(&b, `<defs><clipPath id="%s"><rect x="0" y="0" width="%s" height="%s" rx="%s"/></clipPath></defs><g clip-path="url(#%s)">`,
		clipID, num(miniW), num(miniH), num(miniH/2), clipID)

	pos := 0.0
	markX, marked := 0.0, false
	for i, s := range segs {
		w := fracs[i] / total * miniW
		fmt.Fprintf(&b, `<rect x="%s" y="0" width="%s" height="%s" fill="%s"/>`,
			num(pos), num(w+0.6), num(miniH), rampColor(s.Level))
		if hasValue && !marked {
			inside := (value > s.A || i == 0) && (value <= s.B || i == len(segs)-1)
			if inside {
				t := 0.5
				if vao := s.B - s.A; vao > 0 {
					t = math.Min(math.Max((value-s.A)/vao, 0), 1)
				}
				markX, marked = pos+t*w, true
			}
		}
		pos += w
	}
	b.WriteString(`</g>`)

	if hasValue {
		if !marked {
			markX = miniW / 2
		}
		markX = math.Min(math.Max(markX, 10), miniW-10)
		fmt.Fprintf(&b, `<circle cx="%s" cy="%s" r="10" fill="%s"/>`, num(markX), num(miniH/2), rulerCream)
		fmt.Fprintf(&b, `<circle cx="%s" cy="%s" r="7" fill="%s"/>`, num(markX), num(miniH/2), rulerInk)
	}
	b.WriteString(`</svg>`)
	return b.String()
}

func rampColor(level int) string {
	if level < 0 {
		level = 0
	}
	if level > 5 {
		level = 5
	}
	return rulerRamp[level]
}

// segmentEdges — as fronteiras internas entre faixas, ordenadas e sem repetição.
func segmentEdges(segs []RulerSegment) []float64 {
	seen := map[float64]bool{}
	var out []float64
	for _, s := range segs {
		for _, v := range []float64{s.A, s.B} {
			v = math.Round(v*1e6) / 1e6
			if !seen[v] {
				seen[v] = true
				out = append(out, v)
			}
		}
	}
	sort.Float64s(out)
	return out
}

// formatRulerNumber — número curto, com vírgula decimal. Mesma regra do dossiê.
func formatRulerNumber(v float64) string {
	if math.Abs(v-math.Round(v)) < 1e-9 {
		return strconv.FormatInt(int64(math.Round(v)), 10)
	}
	// 'f' e nunca 'g': com 'g' um CK de 12345,6 sairia como "1,235e+04" no PDF do paciente, e um
	// triglicerídeo de 1234,5 perderia a casa decimal em silêncio. Arredonda em 4 casas (o banco
	// guarda decimal(12,4)) e deixa o Go escolher a representação mais curta que volta ao mesmo
	// número, sem zero à toa.
	r := math.Round(v*1e4) / 1e4
	return strings.Replace(strconv.FormatFloat(r, 'f', -1, 64), ".", ",", 1)
}

// num formata coordenada SVG sem zeros à toa ("490" em vez de "490.000000").
func num(v float64) string {
	return strconv.FormatFloat(math.Round(v*10)/10, 'f', -1, 64)
}

// shortHash dá um id estável e curto para o clipPath. Dois SVGs na mesma página não podem
// compartilhar id de clipPath, senão um recorta o outro.
func shortHash(s string) string {
	h := fnv.New32a()
	_, _ = h.Write([]byte(s))
	return strconv.FormatUint(uint64(h.Sum32()), 36)
}
