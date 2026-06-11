package pdfdoc

import (
	"embed"
	"encoding/base64"
	"fmt"
	"sync"
)

//go:embed assets/fonts/*.woff2 assets/fonts/*.ttf assets/svg/*.svg
var assetFS embed.FS

var (
	assetOnce   sync.Once
	fontFaceCSS string
)

// dataURI lê um asset embutido e devolve um data-URI base64 com o mime informado.
func dataURI(path, mime string) string {
	b, err := assetFS.ReadFile(path)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("data:%s;base64,%s", mime, base64.StdEncoding.EncodeToString(b))
}

// imgSVG devolve uma tag <img> com o SVG embutido como data-URI (preserva object-fit/clip do CSS).
func imgSVG(name, class string) string {
	uri := dataURI("assets/svg/"+name, "image/svg+xml")
	return fmt.Sprintf(`<img class="%s" src="%s">`, class, uri)
}

// fontFaces monta o bloco @font-face com as fontes embutidas como data-URI (HTML autossuficiente).
// Cormorant Garamond (livre, OFL) no título; Inter no corpo. Nalieta (paga) nunca é embutida.
func fontFaces() string {
	assetOnce.Do(func() {
		c5 := dataURI("assets/fonts/cormorant-500.woff2", "font/woff2")
		c6 := dataURI("assets/fonts/cormorant-600.woff2", "font/woff2")
		ir := dataURI("assets/fonts/Inter-Regular.ttf", "font/ttf")
		im := dataURI("assets/fonts/Inter-Medium.ttf", "font/ttf")
		is := dataURI("assets/fonts/Inter-SemiBold.ttf", "font/ttf")
		ib := dataURI("assets/fonts/Inter-Bold.ttf", "font/ttf")
		fontFaceCSS = fmt.Sprintf(`
@font-face { font-family:'Cormorant'; src:url(%s) format('woff2'); font-weight:500; }
@font-face { font-family:'Cormorant'; src:url(%s) format('woff2'); font-weight:600; }
@font-face { font-family:'Inter'; src:url(%s) format('truetype'); font-weight:400; }
@font-face { font-family:'Inter'; src:url(%s) format('truetype'); font-weight:500; }
@font-face { font-family:'Inter'; src:url(%s) format('truetype'); font-weight:600; }
@font-face { font-family:'Inter'; src:url(%s) format('truetype'); font-weight:700; }
`, c5, c6, ir, im, is, ib)
	})
	return fontFaceCSS
}
