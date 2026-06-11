package pdfdoc

import (
	"strings"

	qrcode "github.com/skip2/go-qrcode"
)

// qrSVG gera um QR Code VETORIAL (SVG) a partir do conteúdo, em petrol.
// Vetor em vez do PNG raster do gofpdf legado — nada serrilha.
func qrSVG(content, fill string) string {
	q, err := qrcode.New(content, qrcode.Medium)
	if err != nil {
		return ""
	}
	bm := q.Bitmap() // [][]bool, quadrado, já com quiet zone
	n := len(bm)
	var b strings.Builder
	b.Grow(n * n * 24)
	b.WriteString(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 `)
	b.WriteString(itoa(n))
	b.WriteString(" ")
	b.WriteString(itoa(n))
	b.WriteString(`" shape-rendering="crispEdges"><path fill="`)
	b.WriteString(fill)
	b.WriteString(`" d="`)
	// Agrupa módulos contíguos por linha em traços horizontais (SVG menor).
	for y := 0; y < n; y++ {
		x := 0
		for x < n {
			if !bm[y][x] {
				x++
				continue
			}
			start := x
			for x < n && bm[y][x] {
				x++
			}
			b.WriteString("M")
			b.WriteString(itoa(start))
			b.WriteString(" ")
			b.WriteString(itoa(y))
			b.WriteString("h")
			b.WriteString(itoa(x - start))
			b.WriteString("v1h-")
			b.WriteString(itoa(x - start))
			b.WriteString("z")
		}
	}
	b.WriteString(`"/></svg>`)
	return b.String()
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	neg := n < 0
	if neg {
		n = -n
	}
	var buf [20]byte
	i := len(buf)
	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	if neg {
		i--
		buf[i] = '-'
	}
	return string(buf[i:])
}
