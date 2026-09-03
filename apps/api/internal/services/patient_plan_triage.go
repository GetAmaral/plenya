package services

import (
	"fmt"

	"github.com/plenya/api/internal/pdfdoc"
)

// A triagem: o que entra direto e o que precisa de aceite.
//
// A regra que o usuário escolheu é "texto aplica, número vira sugestão". Ela não pode ser
// implementada por CAMPO, e essa é a descoberta que molda este arquivo: `punch`, `title`, `lede` e
// célula de tabela carregam número o tempo todo. "A ferritina subiu de 48 para 96" é um `punch`, e
// aplicar isso direto seria pôr um número no documento do paciente sem ninguém conferir.
//
// Então a classe do CAMPO diz o que é permitido, e o DIFF diz o que é seguro:
//
//	campo do dossiê        → recusa, nem vira sugestão
//	campo numérico          → sugestão, sempre
//	eixo                    → sugestão, e só se não cortar dado
//	campo de texto          → aplica SE não introduziu numeral novo; senão, sugestão
//	add / remove / reorder  → sugestão, sempre
//
// Estrutural é sempre sugestão porque perder um slide em silêncio é pior que um número errado: o
// número errado alguém lê, o slide que sumiu ninguém procura.

// TriageDecision — o destino de uma operação.
type TriageDecision string

const (
	TriageApply   TriageDecision = "aplicar"
	TriageSuggest TriageDecision = "sugerir"
	TriageReject  TriageDecision = "recusar"
)

// TriagedOp — uma operação com o destino decidido e o porquê.
type TriagedOp struct {
	Op       PlanOp         `json:"op"`
	Decision TriageDecision `json:"decision"`
	Class    string         `json:"class"`
	// Reason é escrito para o MÉDICO ler no cartão de sugestão ou na lista de recusadas.
	Reason string `json:"reason"`
	// Provenance — para cada numeral novo, de onde ele pode ter vindo no dossiê. Vazio com
	// `found:false` significa que o número não existe no dossiê, o que é o sinal mais forte.
	Provenance []NumeralProof `json:"provenance,omitempty"`
}

// NumeralProof — um número que a operação escreve, e as origens candidatas.
type NumeralProof struct {
	Numeral string        `json:"numeral"`
	Found   bool          `json:"found"`
	Matches []NumeralFact `json:"matches,omitempty"`
}

// Triage decide o destino de cada operação contra os slides atuais e o dossiê congelado.
//
// Não aplica nada: devolve o julgamento. Quem aplica é o chamador, que precisa gravar revisão e
// montar as sugestões na mesma transação.
func Triage(slides []pdfdoc.DeckSlide, ops []PlanOp, ix *NumericIndex) []TriagedOp {
	out := make([]TriagedOp, 0, len(ops))
	for _, op := range ops {
		out = append(out, triageUma(slides, op, ix))
	}
	return out
}

func triageUma(slides []pdfdoc.DeckSlide, op PlanOp, ix *NumericIndex) TriagedOp {
	t := TriagedOp{Op: op}

	switch op.Op {
	case OpAdd, OpRemove, OpReorder:
		t.Decision = TriageSuggest
		t.Class = "estrutural"
		t.Reason = "muda quais slides existem ou em que ordem"
		if op.Op == OpAdd && op.Slide != nil {
			// Slide novo passa pelo mesmo escrutínio: todo número dentro dele precisa de origem.
			t.Provenance = provaDoSlide(*op.Slide, ix)
			if temNumeroSemOrigem(t.Provenance) {
				t.Reason = "slide novo com número que não está no dossiê"
			}
		}
		return t

	case OpEdit:
		i := indiceDoSlide(slides, op.SlideID)
		if i < 0 {
			t.Decision = TriageReject
			t.Reason = "slide não encontrado"
			return t
		}
		classe, err := ClassifyPath(&slides[i], op.Path)
		t.Class = classe.String()
		if err != nil {
			t.Decision = TriageReject
			t.Reason = err.Error()
			return t
		}

		switch classe {
		case FieldNumericAuthored:
			t.Decision = TriageSuggest
			t.Reason = "o eixo da régua é ajuste de escala, e precisa de conferência"
			if eixo, ok := comoEixo(op.Value); ok {
				if idx, ok := indiceDaRegua(op.Path); ok && idx < len(slides[i].Rulers) {
					if err := ValidateAxis(eixo, slides[i].Rulers[idx]); err != nil {
						t.Decision = TriageReject
						t.Reason = err.Error()
					}
				}
			} else {
				t.Decision = TriageReject
				t.Reason = "eixo precisa ser um par de números"
			}
			return t

		case FieldNumeric:
			t.Decision = TriageSuggest
			t.Reason = "toca número, unidade ou dose"
			t.Provenance = provaDoTexto(fmt.Sprint(op.Value), ix)
			return t

		case FieldAuthoredText:
			antes := valorAtualComoTexto(&slides[i], op.Path)
			depois := fmt.Sprint(op.Value)
			novos := NumeralDelta(antes, depois)
			if len(novos) == 0 {
				t.Decision = TriageApply
				t.Reason = "reescrita sem número novo"
				return t
			}
			t.Decision = TriageSuggest
			t.Reason = "a reescrita introduz número"
			for _, n := range novos {
				t.Provenance = append(t.Provenance, provaDoNumeral(n, ix))
			}
			return t
		}
	}

	t.Decision = TriageReject
	t.Reason = "operação desconhecida"
	return t
}

func provaDoNumeral(n Numeral, ix *NumericIndex) NumeralProof {
	fatos := ix.Match(n)
	return NumeralProof{Numeral: n.Raw, Found: len(fatos) > 0, Matches: fatos}
}

func provaDoTexto(s string, ix *NumericIndex) []NumeralProof {
	var out []NumeralProof
	for _, n := range ExtractNumerals(s) {
		out = append(out, provaDoNumeral(n, ix))
	}
	return out
}

// provaDoSlide varre todo texto de um slide novo atrás de número sem origem.
func provaDoSlide(s pdfdoc.DeckSlide, ix *NumericIndex) []NumeralProof {
	var out []NumeralProof
	visita := func(texto string) {
		out = append(out, provaDoTexto(texto, ix)...)
	}
	visita(s.Title)
	visita(s.Eyebrow)
	visita(s.Lede)
	visita(s.Kicker)
	visita(s.Source)
	visita(s.Punch)
	for _, c := range s.Cards {
		visita(c.Kicker)
		visita(c.Body)
		// O veredicto é o MAIOR texto do slide de decisão. Sem esta linha, um número escrito ali
		// escapava inteiro da conferência de proveniência.
		visita(c.Verdict)
	}
	for _, r := range s.Rulers {
		visita(r.Note)
		visita(r.Sub)
		// Display carrega número em régua rotulada pelo valor ("27 U/L").
		visita(r.Display)
	}
	if s.Table != nil {
		for _, col := range s.Table.Columns {
			visita(col.Label)
		}
		for _, row := range s.Table.Rows {
			for _, cel := range row.Cells {
				visita(cel)
			}
		}
	}
	if s.Take != nil {
		visita(s.Take.Note)
		if s.Take.Highlight != nil {
			visita(s.Take.Highlight.Dose)
			visita(s.Take.Highlight.Obs)
			// Name, When e Unit carregam número tanto quanto Dose: "Vitamina D 5.000 UI",
			// "uma vez por semana", "mg por semana". Uma posologia inventada num slide novo
			// passava sem acender o aviso de número sem origem.
			visita(s.Take.Highlight.Name)
			visita(s.Take.Highlight.When)
			visita(s.Take.Highlight.Unit)
		}
		for _, g := range s.Take.Groups {
			visita(g.Title)
			for _, it := range g.Items {
				visita(it.Dose)
				visita(it.Sub)
				visita(it.Name)
			}
		}
	}
	if s.Summary != nil {
		visita(s.Summary.Legend)
		visita(s.Summary.StepsTitle)
		for _, p := range s.Summary.Steps {
			visita(p)
		}
		for _, c := range s.Summary.Cards {
			visita(c.Title)
			for _, l := range c.Lines {
				visita(l.Value)
				visita(l.Unit)
				visita(l.Sub)
			}
		}
	}
	for _, st := range s.Steps {
		visita(st.When)
		visita(st.What)
		visita(st.Detail)
	}
	return out
}

func temNumeroSemOrigem(ps []NumeralProof) bool {
	for _, p := range ps {
		if !p.Found {
			return true
		}
	}
	return false
}

// valorAtualComoTexto lê o valor que está hoje no caminho, para comparar os numerais.
func valorAtualComoTexto(s *pdfdoc.DeckSlide, path string) string {
	copia := *s
	// setPath vai e volta por JSON; ler é o mesmo caminho ao contrário, e reusar a travessia evita
	// duas implementações que divergem.
	v, err := getPath(&copia, path)
	if err != nil || v == nil {
		return ""
	}
	return fmt.Sprint(v)
}

func comoEixo(v any) ([]float64, bool) {
	lista, ok := v.([]any)
	if !ok || len(lista) != 2 {
		if f, ok := v.([]float64); ok && len(f) == 2 {
			return f, true
		}
		return nil, false
	}
	out := make([]float64, 0, 2)
	for _, x := range lista {
		f, ok := x.(float64)
		if !ok {
			return nil, false
		}
		out = append(out, f)
	}
	return out, true
}

func indiceDaRegua(path string) (int, bool) {
	m := qualquerIndiceNum.FindStringSubmatch(path)
	if m == nil {
		return 0, false
	}
	var i int
	fmt.Sscanf(m[1], "%d", &i)
	return i, true
}
