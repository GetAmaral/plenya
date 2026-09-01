package services

import (
	"strings"
	"testing"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/pdfdoc"
)

// Um deck e um dossiê que combinam: ferritina 48 -> 500, e uma tabela com dose.
func cenario() ([]pdfdoc.DeckSlide, *NumericIndex) {
	slides := []pdfdoc.DeckSlide{
		*slideDeRegua(),
		*slideDeTabela(),
	}
	slides[0].Punch = "A ferritina subiu"
	d := &dto.PlanDossierResponse{
		Rulers: map[string]dto.PlanRuler{
			"PLNFERR": {
				Code: "PLNFERR", Name: "Ferritina", Unit: "ng/mL",
				History: []dto.PlanRulerPoint{
					{Date: "2025-08-14", Value: 48, Text: "48"},
					{Date: "2026-02-06", Value: 500, Text: "500"},
				},
			},
		},
	}
	return slides, BuildNumericIndex(d)
}

// A regra escolhida foi "texto aplica, número vira sugestão". O ponto deste teste é que ela NÃO
// pode ser por campo: o mesmo `punch` aplica ou vira sugestão dependendo do que a reescrita traz.
func TestTriage_TextoDependeDoNumeroNovo(t *testing.T) {
	slides, ix := cenario()

	semNumero := Triage(slides, []PlanOp{
		{Op: OpEdit, SlideID: "s-reguas", Path: "punch", Value: "A ferritina subiu bastante"},
	}, ix)
	if semNumero[0].Decision != TriageApply {
		t.Errorf("reescrita sem número deu %q (%s), esperava aplicar",
			semNumero[0].Decision, semNumero[0].Reason)
	}

	comNumero := Triage(slides, []PlanOp{
		{Op: OpEdit, SlideID: "s-reguas", Path: "punch", Value: "A ferritina subiu de 48 para 500"},
	}, ix)
	if comNumero[0].Decision != TriageSuggest {
		t.Fatalf("punch com número novo deu %q, esperava sugerir", comNumero[0].Decision)
	}
	if len(comNumero[0].Provenance) != 2 {
		t.Fatalf("esperava origem para os dois números, veio %d", len(comNumero[0].Provenance))
	}
	for _, p := range comNumero[0].Provenance {
		if !p.Found {
			t.Errorf("o número %q está no dossiê e foi dado como sem origem", p.Numeral)
		}
	}

	// E o caso que mais importa: número que NÃO existe no dossiê.
	inventado := Triage(slides, []PlanOp{
		{Op: OpEdit, SlideID: "s-reguas", Path: "punch", Value: "A ferritina chegou a 777"},
	}, ix)
	if inventado[0].Decision != TriageSuggest {
		t.Fatalf("deu %q", inventado[0].Decision)
	}
	if !temNumeroSemOrigem(inventado[0].Provenance) {
		t.Error("777 não existe no dossiê e passou como se tivesse origem")
	}
}

// Campo numérico é sugestão sempre, mesmo quando o valor confere: quem decide se o número certo
// está na frase certa é o médico, e a verificação só prova que ele existe em algum lugar.
func TestTriage_NumericoSempreSugere(t *testing.T) {
	slides, ix := cenario()
	got := Triage(slides, []PlanOp{
		{Op: OpEdit, SlideID: "s-tabela", Path: "table.rows[0].cells[1]", Value: "500 mg"},
	}, ix)
	if got[0].Decision != TriageSuggest {
		t.Errorf("célula de coluna DOSE deu %q, esperava sugerir", got[0].Decision)
	}
	// A mesma linha, na coluna de prosa, é texto.
	got = Triage(slides, []PlanOp{
		{Op: OpEdit, SlideID: "s-tabela", Path: "table.rows[0].cells[2]", Value: "protege o músculo"},
	}, ix)
	if got[0].Decision != TriageApply {
		t.Errorf("célula de prosa deu %q, esperava aplicar", got[0].Decision)
	}
}

func TestTriage_RecusaCampoDoDossie(t *testing.T) {
	slides, ix := cenario()
	got := Triage(slides, []PlanOp{
		{Op: OpEdit, SlideID: "s-reguas", Path: "rulers[0].history", Value: []any{}},
	}, ix)
	if got[0].Decision != TriageReject {
		t.Errorf("editar o histórico da régua deu %q, esperava recusar", got[0].Decision)
	}
	if !strings.Contains(got[0].Reason, "dossiê") {
		t.Errorf("motivo %q não explica que o campo vem do dossiê", got[0].Reason)
	}
}

// Estrutural é sempre sugestão: o número errado alguém lê, o slide que sumiu ninguém procura.
func TestTriage_EstruturalSempreSugere(t *testing.T) {
	slides, ix := cenario()
	for _, op := range []PlanOp{
		{Op: OpRemove, SlideID: "s-tabela"},
		{Op: OpReorder, Order: []string{"s-tabela", "s-reguas"}},
		{Op: OpAdd, Slide: &pdfdoc.DeckSlide{Kind: pdfdoc.DeckCover, Title: "novo"}},
	} {
		got := Triage(slides, []PlanOp{op}, ix)
		if got[0].Decision != TriageSuggest {
			t.Errorf("%s deu %q, esperava sugerir", op.Op, got[0].Decision)
		}
	}
}

// Slide novo passa pelo mesmo escrutínio: todo número dentro dele precisa de origem no dossiê.
func TestTriage_SlideNovoComNumeroInventado(t *testing.T) {
	slides, ix := cenario()
	got := Triage(slides, []PlanOp{{
		Op: OpAdd,
		Slide: &pdfdoc.DeckSlide{
			Kind: pdfdoc.DeckRulers, Title: "Sua ferritina está em 999",
			Punch: "subiu de 48 para 999",
		},
	}}, ix)
	if !temNumeroSemOrigem(got[0].Provenance) {
		t.Error("999 não existe no dossiê e o slide novo passou sem apontar isso")
	}
	if !strings.Contains(got[0].Reason, "não está no dossiê") {
		t.Errorf("motivo %q não avisa do número sem origem", got[0].Reason)
	}
}

// O eixo é o único número do dossiê legitimamente ajustado à mão, e o limite é não cortar dado.
func TestTriage_Eixo(t *testing.T) {
	slides, ix := cenario()

	folgado := Triage(slides, []PlanOp{
		{Op: OpEdit, SlideID: "s-reguas", Path: "rulers[0].axis", Value: []any{0.0, 600.0}},
	}, ix)
	if folgado[0].Decision != TriageSuggest {
		t.Errorf("eixo válido deu %q, esperava sugerir", folgado[0].Decision)
	}

	cortando := Triage(slides, []PlanOp{
		{Op: OpEdit, SlideID: "s-reguas", Path: "rulers[0].axis", Value: []any{0.0, 300.0}},
	}, ix)
	if cortando[0].Decision != TriageReject {
		t.Errorf("eixo que corta o ponto de 500 deu %q, esperava recusar", cortando[0].Decision)
	}
}

// O limite honesto, fixado em teste para ninguém prometer mais do que a verificação entrega:
// ela prova que o número EXISTE no dossiê, nunca que ele significa o que a frase diz.
func TestTriage_NaoValidaSignificado(t *testing.T) {
	slides, ix := cenario()
	// 48 é ferritina. Escrever que 48 é colesterol passa igual, porque o número existe.
	got := Triage(slides, []PlanOp{
		{Op: OpEdit, SlideID: "s-reguas", Path: "punch", Value: "Seu colesterol está em 48"},
	}, ix)
	if got[0].Decision != TriageSuggest {
		t.Fatalf("deu %q", got[0].Decision)
	}
	if temNumeroSemOrigem(got[0].Provenance) {
		t.Fatal("48 existe no dossiê; a verificação não deveria dizer que não")
	}
	// A origem candidata é o que o médico lê para perceber a troca. Se ela não vier legível, a
	// sugestão vira prosa e o aceite vira carimbo.
	if len(got[0].Provenance[0].Matches) == 0 || got[0].Provenance[0].Matches[0].Label == "" {
		t.Error("sem origem legível, o médico não tem como perceber que 48 é ferritina, não colesterol")
	}
}
