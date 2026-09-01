package pdfdoc

import (
	"strings"
	"testing"
)

func deckDeTeste() Deck {
	return Deck{Title: "t", Slides: []DeckSlide{{Kind: DeckCover, Title: "Capa"}}}
}

// Os dois modos existem por razões opostas e trocá-los é silencioso nos dois sentidos: o PDF sem
// fonte cai na fallback e muda toda a métrica do slide (sem erro), e a prévia com fonte embutida
// volta a custar 1,9 MB por render.
func TestDeckHTMLFontes(t *testing.T) {
	embutido, err := DeckHTML(deckDeTeste())
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(embutido, "src:url(data:font/") {
		t.Error("o HTML do PDF tem que embutir as fontes: o Chromium renderiza de string, sem rede")
	}
	if strings.Contains(embutido, "/api/v1/deck-fonts/") {
		t.Error("o HTML do PDF não pode depender de rede")
	}

	linkado, err := DeckHTMLForBrowser(deckDeTeste(), "https://api.exemplo/api/v1/deck-fonts")
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(linkado, "data:font/") {
		t.Error("a prévia não pode embutir fonte: são 96% do payload")
	}
	for _, f := range deckFonts {
		if !strings.Contains(linkado, "https://api.exemplo/api/v1/deck-fonts/"+f.Nome) {
			t.Errorf("a fonte %s não foi declarada por link", f.Nome)
		}
	}

	// O ganho é o motivo da mudança; se ele sumir, a prévia ao vivo deixa de ser viável.
	if len(linkado) > len(embutido)/10 {
		t.Errorf("prévia com %d bytes contra %d do PDF: esperava pelo menos 10x menor",
			len(linkado), len(embutido))
	}

	// Base vazia não pode produzir página sem fonte: cai no modo embutido.
	semBase, err := DeckHTMLForBrowser(deckDeTeste(), "  ")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(semBase, "src:url(data:font/") {
		t.Error("sem base, tem que voltar a embutir — página sem fonte mente sobre o encaixe")
	}
}

// A rota serve por nome, contra uma lista fechada. Caminho não existe como conceito aqui, então
// travessia de diretório não é possível por construção.
func TestDeckFontFile(t *testing.T) {
	for _, f := range deckFonts {
		b, mime, ok := DeckFontFile(f.Nome)
		if !ok || len(b) == 0 {
			t.Errorf("%s não foi servida", f.Nome)
		}
		if mime != f.MIME {
			t.Errorf("%s veio como %q, esperava %q", f.Nome, mime, f.MIME)
		}
	}
	for _, ruim := range []string{"", "../../etc/passwd", "assets/fonts/Inter-Regular.ttf", "Inter-Regular.TTF"} {
		if _, _, ok := DeckFontFile(ruim); ok {
			t.Errorf("%q não deveria ser servido", ruim)
		}
	}
}
