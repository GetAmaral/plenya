package services

// O arco calculado em CÓDIGO, sem chamada ao modelo.
//
// A primeira versão gastava uma chamada inteira para o modelo decidir quais seções existem, quantos
// slides cada uma tem e quais exames vão em cada. Nada disso é julgamento clínico: é ranking (o
// dossiê já entrega `strong` ordenado por peso e `moving` por pontos perdidos) mais uma divisão por
// três (cada slide leva de 2 a 4 réguas). O modelo acertava por memória o que uma conta acerta
// sempre — e a conta não erra o contador do eyebrow, que era o defeito recorrente.
//
// O que o código NÃO decide, e continua com o modelo: o rótulo temático da seção. "O que o cigarro
// está cobrando" é leitura clínica do conjunto, e é o que separa os slides bons dos genéricos. Cada
// seção pode devolver um `label` melhor que o padrão, e custa poucos tokens.

import (
	"fmt"

	"github.com/plenya/api/internal/dto"
)

// reguasPorSlide é o divisor. 2 a 4 é a faixa medida nos dois decks aprovados, e 3 é a média (3,14).
const reguasPorSlide = 3

// montaArco devolve as seções na ordem do arco, já com contagem e material.
func montaArco(d *dto.PlanDossierResponse) []PlanArcSection {
	if d == nil {
		return nil
	}
	var arco []PlanArcSection

	arco = append(arco,
		PlanArcSection{Key: SecCapa, Label: "Seus exames", Slides: 1},
		PlanArcSection{Key: SecResumo, Label: "Resumo", Slides: 1},
	)

	// Boa notícia primeiro, e nunca em menos de 2 slides: regra dos dois decks.
	bem := codigosDe(d.Strong, 9)
	if len(bem) >= 2 {
		arco = append(arco, PlanArcSection{
			Key: SecBem, Label: "O que está bem", Slides: emSlides(len(bem)),
			Exames: bem, Porque: "os achados de maior peso que estão no ótimo",
		})
	}

	// O que está se movendo. Três slides foi o número que funcionou nos dois decks.
	movendo := codigosDe(d.Moving, 9)
	// Dois é o mínimo, como em `bem`: uma seção de régua com um exame só produz um slide com uma
	// régua sozinha, que a conferência sempre acusa e o reparo NÃO consegue corrigir (ele tem ordem
	// de não cortar régua e não sabe acrescentar uma). Aviso garantido e rodada desperdiçada.
	if len(movendo) >= 2 {
		arco = append(arco, PlanArcSection{
			Key: SecMovendo, Label: "O que está se movendo", Slides: emSlides(len(movendo)),
			Exames: movendo, Porque: "os achados que mais tiram ponto, do topo do ranking",
		})
	}

	// Os exames que faltam só existem se houver pedido.
	if d.LabRequest != nil && d.LabRequest.Exams != "" {
		arco = append(arco, PlanArcSection{
			Key: SecFaltando, Label: "O que ainda falta", Slides: 1,
			Porque: "o último pedido de exames",
		})
	}

	// Uma conduta por slide. Sem conduta registrada a seção NÃO existe: inventar conduta é o erro
	// mais caro que esta feature pode cometer.
	if n := len(d.CarePlan); n > 0 {
		if n > 6 {
			n = 6
		}
		var cond []string
		for _, c := range d.CarePlan[:n] {
			cond = append(cond, c.Recommendation)
		}
		arco = append(arco, PlanArcSection{
			Key: SecPlano, Label: "O plano", Slides: n, Condutas: cond,
			Porque: "uma conduta registrada por slide",
		})
	}

	arco = append(arco, PlanArcSection{
		Key: SecSequencia, Label: "A sequência", Slides: 1,
		Porque: "os próximos meses, em ordem, como tabela de duas colunas",
	})
	if len(d.Medications) > 0 {
		arco = append(arco, PlanArcSection{
			Key: SecLevar, Label: "Para levar", Slides: 1,
			Porque: "as receitas vigentes, com dose",
		})
	}
	arco = append(arco, PlanArcSection{Key: SecFecho, Label: "Em uma página", Slides: 1})
	return arco
}

// emSlides divide os exames em slides de até `reguasPorSlide`, nunca deixando um slide com uma
// régua sozinha — que é o defeito que o modelo cometia quando decidia isto sozinho.
func emSlides(n int) int {
	if n <= 0 {
		return 0
	}
	s := (n + reguasPorSlide - 1) / reguasPorSlide
	// Sobrou 1 para o último slide: melhor um slide a menos, com 4 réguas em algum deles.
	if s > 1 && n%reguasPorSlide == 1 {
		s--
	}
	return s
}

// semMaterial diz se o arco ficou só com as seções fixas, sem exame nem achado de anamnese.
func semMaterial(arco []PlanArcSection, d *dto.PlanDossierResponse) bool {
	for _, s := range arco {
		if len(s.Exames) > 0 || len(s.Condutas) > 0 {
			return false
		}
	}
	return len(achadosDaAnamnese(d, 1)) == 0
}

// achadosDaAnamnese — os achados que NÃO têm régua, e por isso não entram no mapa de exames.
// São eles que viram tabela e cartão nos decks aprovados.
func achadosDaAnamnese(d *dto.PlanDossierResponse, max int) []achadoEnxuto {
	var fs []dto.PlanFinding
	for _, f := range append(append([]dto.PlanFinding{}, d.Moving...), d.Strong...) {
		if f.Source == dto.PlanSourceAnamnesis && !f.Stale {
			fs = append(fs, f)
		}
		if len(fs) >= max {
			break
		}
	}
	return enxugaAchados(fs, max)
}

// codigosDe tira os códigos de exame de uma lista de achados, sem repetir e sem os de anamnese
// (que não têm régua).
func codigosDe(fs []dto.PlanFinding, max int) []string {
	visto := map[string]bool{}
	var out []string
	for _, f := range fs {
		if f.Code == "" || visto[f.Code] || f.Source != dto.PlanSourceLab {
			continue
		}
		// Achado velho não lidera seção: é exame a refazer, não retrato de hoje.
		if f.Stale {
			continue
		}
		visto[f.Code] = true
		out = append(out, f.Code)
		if len(out) >= max {
			break
		}
	}
	return out
}

// descreveArco monta o texto do arco para o prompt. Fica no bloco CACHEADO junto do dossiê, e não
// no pedido de cada seção: ele é idêntico nas oito chamadas, e repeti-lo custava ~1,5 mil tokens
// de entrada não cacheada por chamada.
// kindDaSecao — o bloco que cada seção usa nos decks aprovados.
//
// Vai no arco porque é decisão de forma, não de conteúdo, e deixá-la com o modelo custou duas
// regressões medidas: a sequência saiu com o kind "sequence", que NENHUM dos dois decks aprovados
// usou (lá ela é sempre tabela de duas colunas), e a contagem de tabelas caiu de 2-5 para 1.
var kindDaSecao = map[string]string{
	SecCapa:      "cover",
	SecResumo:    "summary",
	SecBem:       "rulers",
	SecMovendo:   "rulers (ou rulers-cards, se um exame que FALTA contrasta com os que há)",
	SecFaltando:  "table",
	SecPlano:     "plan-step ou table",
	SecSequencia: "table de 2 colunas (a primeira com style dose), NUNCA o kind sequence",
	SecLevar:     "takeaway",
	SecFecho:     "closing",
}

func descreveArco(arco []PlanArcSection) string {
	out := "ARCO DECIDIDO (o servidor calculou; escreva exatamente estes slides):\n"
	for i, s := range arco {
		out += fmt.Sprintf("%d. [%s] %s — %d slide(s)", i+1, s.Key, s.Label, s.Slides)
		if k := kindDaSecao[s.Key]; k != "" {
			out += ", kind " + k
		}
		if len(s.Exames) > 0 {
			out += fmt.Sprintf(", exames: %v", s.Exames)
		}
		if len(s.Condutas) > 0 {
			out += fmt.Sprintf(", %d conduta(s)", len(s.Condutas))
		}
		out += "\n"
	}
	return out
}

// ---------------------------------------------------------------------------
// O dossiê enxuto.

// reguaEnxuta é a régua sem o que o modelo não pode usar.
//
// A régua completa tem 9 campos e ~160 tokens. Quatro deles o modelo NÃO pode escrever nem citar:
// `scoreItemId` (id interno), `edges` (as mesmas fronteiras que já estão em `segments`), `points`
// (peso no escore, que não vai para o deck) e o histórico inteiro, do qual o desenho usa os dois
// últimos pontos e o texto da régua cita no máximo três. Mandá-los custa em cache escrito (1,25x)
// em toda geração.
type reguaEnxuta struct {
	Nome      string         `json:"nome"`
	Unidade   string         `json:"unidade,omitempty"`
	Eixo      []float64      `json:"eixo,omitempty"`
	Faixas    []faixaEnxuta  `json:"faixas,omitempty"`
	Historico []pontoDoExame `json:"historico,omitempty"`
}

type faixaEnxuta struct {
	Nivel int    `json:"n"`
	Faixa string `json:"faixa"`
}

type pontoDoExame struct {
	Data  string `json:"data"`
	Valor string `json:"valor"`
	Nivel *int   `json:"nivel,omitempty"`
}

// historicoNoTexto é quantos pontos o `note` da régua chega a citar ("239 em 2024, 432 em 2025,
// 500 agora"). Além disso é peso morto: o desenho usa os dois últimos.
const historicoNoTexto = 3

// exameParaPrompt — UM objeto por exame, fundindo a régua e o achado.
//
// Antes iam as duas coisas separadas: `reguasCitadas` com nome/unidade/eixo/faixas/histórico, e
// `estaBem`/`seMovendo` com nome/valor/unidade/nível/data do MESMO exame. O dossiê era 56% do
// prefixo cacheado e metade dele era repetição.
//
// As seis faixas também saíram: o modelo escreve o rótulo avaliativo ("dentro da faixa ótima"), e
// para isso precisa saber em que faixa o paciente ESTÁ, não o desenho inteiro da escala — que quem
// desenha é o servidor, a partir do dossiê congelado.
type exameParaPrompt struct {
	Nome      string         `json:"nome"`
	Unidade   string         `json:"un,omitempty"`
	Atual     string         `json:"atual,omitempty"`
	Data      string         `json:"data,omitempty"`
	Nivel     *int           `json:"nivel,omitempty"`
	Faixa     string         `json:"faixaDoNivel,omitempty"`
	Tendencia string         `json:"tendencia,omitempty"`
	Porque    string         `json:"porque,omitempty"`
	Antes     []pontoDoExame `json:"antes,omitempty"`
}

// examesParaPrompt monta o mapa dos exames que o arco escolheu.
func examesParaPrompt(d *dto.PlanDossierResponse, codigos map[string]bool) map[string]exameParaPrompt {
	achado := map[string]dto.PlanFinding{}
	for _, f := range append(append([]dto.PlanFinding{}, d.Strong...), d.Moving...) {
		if _, ja := achado[f.Code]; !ja {
			achado[f.Code] = f
		}
	}
	out := map[string]exameParaPrompt{}
	for code, r := range d.Rulers {
		if !codigos[code] {
			continue
		}
		e := exameParaPrompt{Nome: r.Name, Unidade: r.Unit}
		if n := len(r.History); n > 0 {
			ult := r.History[n-1]
			e.Atual, e.Data, e.Nivel = ult.Text, ult.Date, ult.Level
			// Só a faixa em que o paciente está: é o que sustenta o rótulo avaliativo.
			if ult.Level != nil {
				for _, sg := range r.Segments {
					if sg.Level == *ult.Level {
						e.Faixa = sg.Label
						break
					}
				}
			}
			// Os anteriores, para a mini-série do `note` ("239 em 2024, 432 em 2025, 500 agora").
			ini := n - historicoNoTexto
			if ini < 0 {
				ini = 0
			}
			for _, p := range r.History[ini : n-1] {
				e.Antes = append(e.Antes, pontoDoExame{Data: p.Date, Valor: p.Text, Nivel: p.Level})
			}
		}
		if f, ok := achado[code]; ok {
			e.Tendencia, e.Porque = string(f.Trend), f.Reason
		}
		out[code] = e
	}
	return out
}

func enxuga(r dto.PlanRuler) reguaEnxuta {
	out := reguaEnxuta{Nome: r.Name, Unidade: r.Unit, Eixo: r.Axis}
	for _, sg := range r.Segments {
		// `label` já é a faixa legível ("≤15", ">300") que o dossiê calculou: dizer o mesmo com
		// dois números seria mandar o dobro para o modelo ler pior.
		out.Faixas = append(out.Faixas, faixaEnxuta{Nivel: sg.Level, Faixa: sg.Label})
	}
	h := r.History
	if len(h) > historicoNoTexto {
		h = h[len(h)-historicoNoTexto:]
	}
	for _, p := range h {
		out.Historico = append(out.Historico, pontoDoExame{Data: p.Date, Valor: p.Text, Nivel: p.Level})
	}
	return out
}

// achadoEnxuto — o achado sem os campos que só servem ao ranking, que já foi feito.
type achadoEnxuto struct {
	Codigo string `json:"code"`
	Nome   string `json:"nome"`
	Valor  string `json:"valor,omitempty"`
	Unid   string `json:"unidade,omitempty"`
	Nivel  int    `json:"nivel"`
	Data   string `json:"data,omitempty"`
	Tend   string `json:"tendencia,omitempty"`
	Porque string `json:"porque,omitempty"`
}

func enxugaAchados(fs []dto.PlanFinding, n int) []achadoEnxuto {
	if len(fs) > n {
		fs = fs[:n]
	}
	out := make([]achadoEnxuto, 0, len(fs))
	for _, f := range fs {
		out = append(out, achadoEnxuto{
			Codigo: f.Code, Nome: f.Name, Valor: f.Text, Unid: f.Unit,
			Nivel: f.Level, Data: f.Date, Tend: string(f.Trend), Porque: f.Reason,
		})
	}
	return out
}
