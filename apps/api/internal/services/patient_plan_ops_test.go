package services

import (
	"errors"
	"testing"

	"github.com/plenya/api/internal/pdfdoc"
)

func slideDeRegua() *pdfdoc.DeckSlide {
	return &pdfdoc.DeckSlide{
		ID: "s-reguas", Kind: pdfdoc.DeckRulers, Title: "A ferritina dobrou",
		Rulers: []pdfdoc.DeckRulerBlock{{Ruler: pdfdoc.Ruler{
			Code: "PLNFERR", Display: "Ferritina", Unit: "ng/mL",
			Axis:     [2]float64{0, 520},
			Segments: []pdfdoc.RulerSegment{{Level: 5, A: 50, B: 200}},
			History:  []pdfdoc.RulerPoint{{Value: 48, Text: "48"}, {Value: 500, Text: "500"}},
		}}},
	}
}

func slideDeTabela() *pdfdoc.DeckSlide {
	return &pdfdoc.DeckSlide{
		ID: "s-tabela", Kind: pdfdoc.DeckTableKind,
		Table: &pdfdoc.DeckTable{
			Columns: []pdfdoc.DeckTableCol{
				{Label: "O quê"},
				{Label: "Quanto", Style: pdfdoc.DeckColDose},
				{Label: "Por quê", Style: pdfdoc.DeckColWhy},
			},
			Rows: []pdfdoc.DeckTableRow{{Cells: []string{"Proteína", "112 g", "protege a massa magra"}}},
		},
	}
}

// O vocabulário é fechado: o que não está nele é recusado, e não permitido por omissão. Uma feature
// que escreve num documento lido por paciente erra melhor negando.
func TestClassifyPath(t *testing.T) {
	regua := slideDeRegua()
	casos := []struct {
		slide  *pdfdoc.DeckSlide
		path   string
		classe FieldClass
		erro   error
		porque string
	}{
		{regua, "title", FieldAuthoredText, nil, "prosa do envelope"},
		{regua, "punch", FieldAuthoredText, nil, "o fecho é prosa, mesmo carregando número"},
		{regua, "rulers[0].display", FieldAuthoredText, nil, "nome que o paciente lê é autoral"},
		{regua, "rulers[0].note", FieldAuthoredText, nil, "a leitura clínica é autoral"},
		{regua, "rulers[0].axis", FieldNumericAuthored, nil, "o eixo é o único número afinado à mão"},
		{regua, "rulers[0].unit", FieldNumeric, nil, "unidade é número"},
		{regua, "rulers[0].segments", FieldDossierOwned, ErrCaminhoDoDossie, "faixa vem do catálogo do escore"},
		{regua, "rulers[0].history", FieldDossierOwned, ErrCaminhoDoDossie, "histórico é medida do paciente"},
		{regua, "rulers[0].code", FieldDossierOwned, ErrCaminhoDoDossie, "o código é a chave do exame"},
		{regua, "rulers[0].segments[2].b", FieldUnknown, ErrCaminhoDesconhecido, "caminho fundo dentro da faixa nem existe no vocabulário"},
		{regua, "qualquerCoisa", FieldUnknown, ErrCaminhoDesconhecido, "campo inventado"},
		{regua, "", FieldUnknown, ErrCaminhoDesconhecido, "caminho vazio"},
		{regua, "table.dense", FieldUnknown, ErrCaminhoDesconhecido, "slide de régua não tem tabela"},
		{slideDeTabela(), "table.dense", FieldAuthoredText, nil, "aqui tem"},
	}
	for _, c := range casos {
		classe, err := ClassifyPath(c.slide, c.path)
		if classe != c.classe {
			t.Errorf("%q: classe %v, esperava %v (%s)", c.path, classe, c.classe, c.porque)
		}
		if c.erro != nil && !errors.Is(err, c.erro) {
			t.Errorf("%q: erro %v, esperava %v (%s)", c.path, err, c.erro, c.porque)
		}
		if c.erro == nil && err != nil {
			t.Errorf("%q: erro inesperado %v (%s)", c.path, err, c.porque)
		}
	}
}

// A célula de tabela não tem classe fixa: numa coluna de dose ela é número, numa coluna de prosa
// não é. Sem olhar o estilo da coluna, "112 g" viraria texto livre editável sem aceite.
func TestClassifyPath_CelulaDependeDaColuna(t *testing.T) {
	s := slideDeTabela()
	if c, _ := ClassifyPath(s, "table.rows[0].cells[0]"); c != FieldAuthoredText {
		t.Errorf("coluna de prosa deu %v, esperava texto", c)
	}
	if c, _ := ClassifyPath(s, "table.rows[0].cells[1]"); c != FieldNumeric {
		t.Errorf("coluna de DOSE deu %v, esperava numérico", c)
	}
	if c, _ := ClassifyPath(s, "table.rows[0].cells[2]"); c != FieldAuthoredText {
		t.Errorf("coluna de explicação deu %v, esperava texto", c)
	}
}

func TestApplyOps_Edit(t *testing.T) {
	slides := []pdfdoc.DeckSlide{*slideDeRegua()}

	out, err := ApplyOps(slides, []PlanOp{
		{Op: OpEdit, SlideID: "s-reguas", Path: "title", Value: "Título novo"},
		{Op: OpEdit, SlideID: "s-reguas", Path: "rulers[0].note", Value: "48 em 2025, 500 agora"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if out[0].Title != "Título novo" {
		t.Errorf("título = %q", out[0].Title)
	}
	if out[0].Rulers[0].Note != "48 em 2025, 500 agora" {
		t.Errorf("note = %q", out[0].Rulers[0].Note)
	}
	// A entrada não pode ter sido mexida: quem chama precisa do estado anterior para a revisão.
	if slides[0].Title == "Título novo" {
		t.Error("ApplyOps mutou a lista de entrada")
	}
	// E o que veio do dossiê tem que continuar intacto depois de editar o irmão.
	if len(out[0].Rulers[0].History) != 2 || out[0].Rulers[0].History[0].Value != 48 {
		t.Error("o histórico da régua foi corrompido por uma edição de texto")
	}
}

func TestApplyOps_RecusaCampoDoDossie(t *testing.T) {
	slides := []pdfdoc.DeckSlide{*slideDeRegua()}
	_, err := ApplyOps(slides, []PlanOp{
		{Op: OpEdit, SlideID: "s-reguas", Path: "rulers[0].history", Value: []any{}},
	})
	if !errors.Is(err, ErrCaminhoDoDossie) {
		t.Fatalf("erro %v, esperava recusa de campo do dossiê", err)
	}
}

// O alvo é o ID, nunca o índice. É esta propriedade que impede uma sugestão criada antes de um
// reorder de escrever no slide errado depois.
func TestApplyOps_AlvoPorID(t *testing.T) {
	slides := []pdfdoc.DeckSlide{
		{ID: "a", Kind: pdfdoc.DeckCover, Title: "primeiro"},
		{ID: "b", Kind: pdfdoc.DeckClosing, Title: "segundo"},
	}
	// Reordena e depois edita "b": tem que acertar "b" onde quer que ele esteja.
	depois, err := ApplyOps(slides, []PlanOp{
		{Op: OpReorder, Order: []string{"b", "a"}},
		{Op: OpEdit, SlideID: "b", Path: "title", Value: "editado"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if depois[0].ID != "b" || depois[0].Title != "editado" {
		t.Errorf("editou o slide errado: %+v", depois[0])
	}
	if depois[1].Title != "primeiro" {
		t.Errorf("o outro slide foi alterado: %q", depois[1].Title)
	}

	if _, err := ApplyOps(slides, []PlanOp{{Op: OpEdit, SlideID: "nao-existe", Path: "title", Value: "x"}}); !errors.Is(err, ErrSlideNaoEncontrado) {
		t.Errorf("id inexistente deu %v", err)
	}
}

func TestApplyOps_AddRemoveReorder(t *testing.T) {
	slides := []pdfdoc.DeckSlide{
		{ID: "a", Kind: pdfdoc.DeckCover},
		{ID: "b", Kind: pdfdoc.DeckClosing},
	}

	comNovo, err := ApplyOps(slides, []PlanOp{
		{Op: OpAdd, AfterSlideID: "a", Slide: &pdfdoc.DeckSlide{Kind: pdfdoc.DeckRulers, Title: "meio"}},
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(comNovo) != 3 || comNovo[1].Title != "meio" {
		t.Fatalf("inserção fora de lugar: %+v", comNovo)
	}
	if comNovo[1].ID == "" {
		t.Error("slide novo entrou sem id: ficaria inendereçável")
	}

	semB, err := ApplyOps(slides, []PlanOp{{Op: OpRemove, SlideID: "b"}})
	if err != nil || len(semB) != 1 || semB[0].ID != "a" {
		t.Fatalf("remoção errada: %+v, %v", semB, err)
	}

	// Reorder tem que ser completo: uma lista parcial some com slide em silêncio.
	if _, err := ApplyOps(slides, []PlanOp{{Op: OpReorder, Order: []string{"a"}}}); err == nil {
		t.Error("reorder parcial passou; sumiria com um slide sem ninguém ver")
	}
	if _, err := ApplyOps(slides, []PlanOp{{Op: OpReorder, Order: []string{"a", "fantasma"}}}); !errors.Is(err, ErrSlideNaoEncontrado) {
		t.Error("reorder com id inventado passou")
	}
}

// O eixo pode ser apertado à mão quando um extremo esmaga a escala, mas não pode cortar dado: eixo
// que deixa um ponto de fora esconde o valor sem erro nenhum, igual ao overflow do slide.
func TestValidateAxis(t *testing.T) {
	r := slideDeRegua().Rulers[0] // faixa 50–200, pontos 48 e 500

	if err := ValidateAxis([]float64{0, 520}, r); err != nil {
		t.Errorf("o eixo original foi recusado: %v", err)
	}
	if err := ValidateAxis([]float64{0, 600}, r); err != nil {
		t.Errorf("eixo mais folgado foi recusado: %v", err)
	}
	if err := ValidateAxis([]float64{0, 300}, r); !errors.Is(err, ErrEixoCortaDado) {
		t.Errorf("eixo que corta o ponto de 500 passou: %v", err)
	}
	if err := ValidateAxis([]float64{60, 520}, r); !errors.Is(err, ErrEixoCortaDado) {
		t.Errorf("eixo que corta o ponto de 48 e o início da faixa passou: %v", err)
	}
	if err := ValidateAxis([]float64{520, 0}, r); err == nil {
		t.Error("eixo invertido passou")
	}
	if err := ValidateAxis([]float64{10}, r); err == nil {
		t.Error("eixo com um número só passou")
	}
}
