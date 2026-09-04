package services

// O que o SERVIDOR decide no deck, e não o modelo.
//
// Tudo aqui é regra observada nos dois decks aprovados que não depende de julgamento clínico. Deixar
// no prompt significa pedir ao modelo que acerte por memória a cada geração, e ele erra: o próprio
// deck aprovado do Ricardo tem "· 1 de 3" repetido em dois slides seguidos, porque foi escrito de
// uma vez só. Aqui isso é aritmética.

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/pdfdoc"
)

// aplicaRegrasDoDeck compõe os eyebrows numerados, a variante escura, a legenda da rampa e a capa.
func aplicaRegrasDoDeck(secoes []PlanArcSection, porSecao [][]pdfdoc.DeckSlide, primeiroNome, data string) []pdfdoc.DeckSlide {
	var out []pdfdoc.DeckSlide

	for si, sec := range secoes {
		if si >= len(porSecao) {
			break
		}
		slides := porSecao[si]
		for i := range slides {
			s := slides[i]

			// 1. o eyebrow numerado. O contador vem do ARCO, não da contagem do que voltou: se o
			// modelo devolveu menos slides do que prometeu, o "de N" continua sendo o N combinado,
			// e a diferença aparece como seção incompleta em vez de numeração remendada.
			rotulo := strings.TrimSpace(sec.Label)
			switch {
			case sec.Key == SecCapa:
				// A capa não leva rótulo de seção: leva a data, como nos dois decks.
				s.Eyebrow = "Seus exames"
				if data != "" {
					s.Eyebrow += " · " + data
				}
			case rotulo == "":
			case sec.Slides > 1:
				// O denominador vem do arco. Se a seção devolveu MAIS slides do que prometeu, o
				// numerador passaria dele ("· 4 de 3"); o total impresso passa a ser o maior dos
				// dois, que é a única leitura que não mente.
				total := sec.Slides
				if len(slides) > total {
					total = len(slides)
				}
				s.Eyebrow = fmt.Sprintf("%s · %d de %d", rotulo, i+1, total)
			default:
				s.Eyebrow = rotulo
			}

			// 2. a variante. Nos dois decks aprovados, "deep" existe em exatamente dois slides,
			// capa e fecho, e "dark" nunca foi usado. O docstring do deck da Ana diz por quê: um
			// tom só nas páginas de conteúdo.
			if sec.Key == SecCapa || sec.Key == SecFecho {
				s.Variant = "deep"
			} else {
				s.Variant = ""
			}

			// 3. a legenda é decidida DEPOIS, em `poeLegenda`: aqui as réguas ainda não passaram
			// pela hidratação, que remove as que não existem no dossiê. Marcar agora podia deixar
			// a legenda num slide que ficaria sem régua nenhuma.
			s.Legend = false

			// 4. o nome do paciente na capa. Vem daqui e não do modelo: o nome nunca sai do
			// prontuário para a API, e mesmo assim a capa precisa dele, como nos dois decks.
			if sec.Key == SecCapa && primeiroNome != "" {
				s.Title = primeiroNome
			}

			out = append(out, s)
		}
	}
	return out
}

// primeiroNomeDe devolve só o primeiro nome, que é o que a capa dos decks aprovados usa.
func primeiroNomeDe(nome string) string {
	campos := strings.Fields(strings.TrimSpace(nome))
	if len(campos) == 0 {
		return ""
	}
	return campos[0]
}

// formataDataPT devolve "27 de agosto de 2026" a partir de um ISO. É o formato do eyebrow da capa
// nos dois decks aprovados.
func formataDataPT(iso string) string {
	t, err := time.Parse(time.RFC3339, iso)
	if err != nil {
		if t, err = time.Parse("2006-01-02", iso); err != nil {
			return ""
		}
	}
	meses := []string{"janeiro", "fevereiro", "março", "abril", "maio", "junho",
		"julho", "agosto", "setembro", "outubro", "novembro", "dezembro"}
	return fmt.Sprintf("%d de %s de %d", t.Day(), meses[int(t.Month())-1], t.Year())
}

// reparaEstouro manda de volta ao modelo só os slides que não couberam, com o excesso medido, e
// devolve o deck com eles trocados. Devolve `false` quando não deu para reparar: nesse caso o deck
// original segue, com o aviso na tela.
//
// Troca POR ID e não por índice: o modelo pode devolver os slides em qualquer ordem, e casar por
// posição colocaria o conteúdo de um slide dentro de outro sem erro nenhum.
func (s *PatientPlanAssistantService) reparaEstouro(
	plan *models.PatientPlan, estouro []pdfdoc.DeckOverflow,
	estilo []dto.PlanGenWarning, meta *AICallMeta,
) ([]pdfdoc.DeckSlide, bool) {
	// Junta os dois tipos de defeito no MESMO reparo: geometria (o slide não cabe) e forma (punch
	// fora da faixa, régua sozinha, travessão). São a mesma correção — encurtar e ajustar prosa —
	// e mandar duas rodadas para o modelo seria pagar duas vezes pelo mesmo trabalho.
	problemas := map[string][]string{}
	porIndice := map[int]pdfdoc.DeckOverflow{}
	for _, o := range estouro {
		porIndice[o.Slide] = o
	}
	for i, sl := range plan.Content {
		if o, ok := porIndice[i+1]; ok { // CheckDeckOverflow numera a partir de 1
			// Altura e LARGURA são defeitos diferentes e pedem cortes diferentes.
			//
			// A medição empurra o slide quando `right > 1` OU `bottom > 1`, e num slide que só
			// vaza para o lado o `bottom` costuma ser bem negativo. Formatar só ele mandava ao
			// modelo "sobraram -480px de altura", que não quer dizer nada, e o defeito real (quase
			// sempre prosa numa coluna de dose, que não quebra linha) nunca era comunicado.
			//
			// %.0f e não %d: a medida é float, o Chromium devolve subpixel, e o `go vet` pega.
			if o.Bottom > 1 {
				problemas[sl.ID] = append(problemas[sl.ID],
					fmt.Sprintf("sobraram %.0fpx de ALTURA para fora da moldura", o.Bottom))
			}
			if o.Right > 1 {
				problemas[sl.ID] = append(problemas[sl.ID],
					fmt.Sprintf("sobraram %.0fpx de LARGURA para fora: quase sempre é prosa numa coluna "+
						"de dose, que não quebra linha", o.Right))
			}
		}
	}
	for _, w := range estilo {
		if w.SlideID != "" {
			problemas[w.SlideID] = append(problemas[w.SlideID], w.Reason)
		}
	}
	var quebrados []pdfdoc.DeckSlide
	var excessos []string
	for _, sl := range plan.Content {
		ps, ok := problemas[sl.ID]
		if !ok {
			continue
		}
		quebrados = append(quebrados, sl)
		excessos = append(excessos, "slide id="+sl.ID+": "+strings.Join(ps, "; "))
	}
	if len(quebrados) == 0 {
		return nil, false
	}
	bruto, err := json.Marshal(quebrados)
	if err != nil {
		return nil, false
	}
	corrigidos, mr, err := s.ai.RepairOverflow(PlanRepairRequest{
		SlidesJSON: string(bruto),
		Excessos:   strings.Join(excessos, "\n"), Model: s.model,
	})
	if meta != nil {
		meta.InputTokens += mr.InputTokens
		meta.OutputTokens += mr.OutputTokens
		meta.CacheReadTokens += mr.CacheReadTokens
		meta.CacheWriteTokens += mr.CacheWriteTokens
		meta.LatencyMs += mr.LatencyMs
	}
	if err != nil || len(corrigidos) == 0 {
		return nil, false
	}
	porID := map[string]pdfdoc.DeckSlide{}
	for _, c := range corrigidos {
		if c.ID != "" {
			porID[c.ID] = c
		}
	}
	out := make([]pdfdoc.DeckSlide, len(plan.Content))
	copy(out, plan.Content)
	trocados := 0
	for i := range out {
		novo, ok := porID[out[i].ID]
		if !ok {
			continue
		}
		// O id, as réguas e TUDO QUE O SERVIDOR DECIDE ficam: o reparo é para encurtar PROSA.
		//
		// Se o modelo tirar uma régua inteira para caber, isso muda o que o paciente vê do próprio
		// exame, e não é dele a decisão. E o eyebrow numerado, a variante e a legenda vieram do
		// arco, não do modelo: sem preservá-los, a primeira execução do reparo apagou a legenda da
		// rampa do deck inteiro, porque o slide que a carregava era justamente um dos que
		// estouravam.
		// O `kind` também fica: o schema do reparo exige `kind` e oferece o enum inteiro, então
		// um `rulers-cards` podia voltar como `rulers`. As réguas continuariam desenhando, e os
		// cartões sumiriam do PDF sem aviso nenhum — a conferência de bloco não dispara, porque
		// `rulers` com réguas está correto.
		novo.Kind = out[i].Kind
		novo.ID = out[i].ID
		novo.Eyebrow = out[i].Eyebrow
		novo.Variant = out[i].Variant
		novo.Legend = out[i].Legend
		// As réguas ficam SEMPRE as do original, não só quando a contagem muda.
		//
		// O schema do slide manda o modelo NÃO escrever axis, segments nem history (o servidor
		// hidrata a partir do dossiê), e o reparo reusa esse mesmo schema. Então tudo que volta do
		// reparo tem régua com eixo [0,0] e sem faixa. Comparar a CONTAGEM não pegava isso, porque
		// o prompt do reparo proíbe cortar régua e a contagem batia: o slide era salvo com a barra
		// vazia ao lado do nome de um exame real, e a remedição do estouro ainda MELHORAVA, porque
		// a régua sem dado ocupa menos. É o mesmo modo de falha que a hidratação existe para
		// impedir, entrando pela porta dos fundos.
		novo.Rulers = out[i].Rulers
		if out[i].Summary != nil {
			novo.Summary = out[i].Summary
		}
		out[i] = novo
		trocados++
	}
	return out, trocados > 0
}

// poeLegenda marca a legenda da rampa no primeiro slide que DE FATO tem régua.
//
// Roda depois da hidratação de propósito: ela remove as réguas cujo exame não está no dossiê, e um
// slide pode terminar sem nenhuma. Marcando antes, a legenda ficava num slide vazio e os slides de
// régua de verdade não ganhavam nenhuma — e a conferência ficava quieta, porque contava uma.
func poeLegenda(slides []pdfdoc.DeckSlide) {
	posta := false
	for i := range slides {
		slides[i].Legend = false
		if posta || len(slides[i].Rulers) == 0 {
			continue
		}
		if slides[i].Kind == pdfdoc.DeckRulers || slides[i].Kind == pdfdoc.DeckRulersCards {
			slides[i].Legend = true
			posta = true
		}
	}
}
