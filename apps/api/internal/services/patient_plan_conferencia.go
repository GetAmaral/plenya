package services

// As conferências MECÂNICAS do deck gerado.
//
// Tudo aqui é regra medida nos dois decks aprovados que dá para verificar sem julgamento clínico.
// Não bloqueia nada e não reescreve nada: vira aviso no slide exato, porque a geração varia de uma
// execução para outra e o médico precisa ver ONDE ela variou, em vez de reler treze slides
// procurando o que está fora do padrão.
//
// O que NÃO está aqui, de propósito: se o número significa o que a frase diz, se o achado merecia
// slide, se a conduta está certa. Isso é leitura clínica e não tem verificação mecânica.

import (
	"fmt"
	"regexp"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/pdfdoc"
)

var (
	reEm        = regexp.MustCompile(`(?i)<em>`)
	reTravessao = regexp.MustCompile(`[—–]`)
	// "Não é X. É Y." — a construção de coach que a regra editorial proíbe.
	reNaoEMasE = regexp.MustCompile(`(?i)\bnão é\b[^.]{1,60}\.\s*é\b`)
)

// confereDeck devolve os desvios do padrão medido. Ordem: pelo slide.
func confereDeck(slides []pdfdoc.DeckSlide) []dto.PlanGenWarning {
	var out []dto.PlanGenWarning
	avisa := func(i int, motivo string) {
		out = append(out, dto.PlanGenWarning{
			Kind: dto.PlanGenWarningEstilo, SlideIndex: i + 1, SlideID: slides[i].ID,
			Title: slides[i].Title, Reason: motivo,
		})
	}

	for i := range slides {
		s := slides[i]

		// O punch é o fecho do slide, e nos dois decks ele está em 85% deles com EXATAMENTE um
		// <em>. Ausente só em capa, para-levar e fecho.
		semPunch := s.Kind == pdfdoc.DeckCover || s.Kind == pdfdoc.DeckClosing || s.Kind == pdfdoc.DeckTakeaway
		switch {
		case s.Punch == "" && !semPunch:
			avisa(i, "slide sem punch: nos decks aprovados só capa, para-levar e fecho não têm")
		case s.Punch != "":
			if n := len(reEm.FindAllString(s.Punch, -1)); n != 1 {
				avisa(i, fmt.Sprintf("punch com %d <em>: o padrão é exatamente um, no fecho da frase", n))
			}
			if c := len([]rune(s.Punch)); c < 45 || c > 130 {
				avisa(i, fmt.Sprintf("punch com %d caracteres: a faixa medida é 55 a 110", c))
			}
		}

		// Régua: 2 a 4 por slide. Uma só desperdiça o slide; cinco não cabem.
		if n := len(s.Rulers); n > 0 && (n < 2 || n > 4) {
			avisa(i, fmt.Sprintf("%d régua(s) no slide: os decks aprovados usam de 2 a 4, nunca uma só", n))
		}

		// Título: uma linha. O fecho é a exceção, e é longo de propósito.
		if s.Kind != pdfdoc.DeckClosing && s.Kind != pdfdoc.DeckCover {
			if c := len([]rune(s.Title)); c > 60 {
				avisa(i, fmt.Sprintf("título com %d caracteres: passa de uma linha (a faixa é 16 a 53)", c))
			}
		}

		// As regras editoriais invariantes, nos campos que o paciente lê.
		// Slice e não map: a ordem de um map em Go é aleatória, e a caixa da tela mostra só os
		// primeiros oito avisos — com map, um deck com dois defeitos mostrava um subconjunto
		// diferente a cada execução.
		for _, c := range []struct{ campo, txt string }{
			{"título", s.Title}, {"punch", s.Punch}, {"lede", s.Lede}, {"kicker", s.Kicker},
		} {
			campo, txt := c.campo, c.txt
			if reTravessao.MatchString(txt) {
				avisa(i, "travessão no "+campo+": use vírgula, ponto ou dois-pontos")
			}
			if reNaoEMasE.MatchString(txt) {
				avisa(i, `construção "Não é X. É Y." no `+campo+": é maneirismo de IA, reescreva")
			}
		}

		// Coluna de dose não quebra linha. Prosa nela vaza o slide para fora da moldura, e o
		// estouro nem sempre acusa porque o excesso é horizontal.
		if s.Table != nil {
			for ci, col := range s.Table.Columns {
				if col.Style != "dose" {
					continue
				}
				for _, r := range s.Table.Rows {
					if ci < len(r.Cells) && len([]rune(r.Cells[ci])) > 28 {
						avisa(i, fmt.Sprintf("coluna %q é de dose e tem célula com %d caracteres: dose não quebra linha",
							col.Label, len([]rune(r.Cells[ci]))))
						break
					}
				}
			}
		}
	}

	// A legenda da rampa é uma vez por deck. Como o servidor a põe, mais de uma é defeito nosso.
	legendas := 0
	for _, s := range slides {
		if s.Legend {
			legendas++
		}
	}
	temRegua := false
	for _, s := range slides {
		if len(s.Rulers) > 0 {
			temRegua = true
			break
		}
	}
	if temRegua && legendas != 1 {
		out = append(out, dto.PlanGenWarning{
			Kind:   dto.PlanGenWarningEstilo,
			Reason: fmt.Sprintf("a legenda da rampa aparece %d vez(es): o padrão é exatamente uma, no primeiro slide de régua", legendas),
		})
	}

	// A tabela é o bloco que carrega o plano: 9 de 21 e 8 de 20 nos aprovados. Deck sem tabela
	// nenhuma quase sempre quer dizer que a seção de conduta não existiu.
	comTabela := 0
	for _, s := range slides {
		if s.Table != nil && len(s.Table.Rows) > 0 {
			comTabela++
		}
	}
	if len(slides) >= 10 && comTabela == 0 {
		out = append(out, dto.PlanGenWarning{
			Kind:   dto.PlanGenWarningEstilo,
			Reason: "nenhum slide com tabela: nos decks aprovados a tabela carrega 40% do deck, e é como o plano se explica",
		})
	}
	return out
}

// semEstilo tira os avisos de estilo da lista, para a reconferência depois do reparo não duplicar.
func semEstilo(ws []dto.PlanGenWarning) []dto.PlanGenWarning {
	out := ws[:0:0]
	for _, w := range ws {
		if w.Kind != dto.PlanGenWarningEstilo {
			out = append(out, w)
		}
	}
	return out
}

// avisosNumericos confere cada número escrito contra o índice do dossiê.
//
// Extraído para poder rodar DE NOVO depois do reparo: o reparo reescreve prosa, e é na prosa que os
// números moram.
func avisosNumericos(slides []pdfdoc.DeckSlide, ix *NumericIndex) []dto.PlanGenWarning {
	var out []dto.PlanGenWarning
	for i := range slides {
		for _, p := range provaDoSlide(slides[i], ix) {
			if p.Found {
				continue
			}
			out = append(out, dto.PlanGenWarning{
				Kind: dto.PlanGenWarningNumeral, SlideIndex: i + 1, SlideID: slides[i].ID,
				Title: slides[i].Title, Numeral: p.Numeral,
				Reason: "este número não foi encontrado no prontuário compilado",
			})
		}
	}
	return out
}
