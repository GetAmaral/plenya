package services

// O deck montado SÓ COM CÓDIGO, custo zero de modelo.
//
// A geração por LLM custava ~US$ 0,20 por plano e refazia, toda vez, um trabalho que é montagem:
// escolher os exames por ranking, agrupá-los de três em três, hidratar a régua do escore, formatar
// a mini-série do histórico, transformar as condutas registradas em slides, virar a receita em
// "para levar". Nada disso é julgamento — é o dado, arrumado.
//
// O que este montador NÃO escreve, e fica para a conversa com o médico:
//
//   - o TÍTULO como afirmação ("A ferritina dobrou em dois anos"). Aqui ele sai descritivo, a
//     partir dos próprios exames do slide ("Ferritina, saturação e ferro sérico"): o deck já é
//     legível e imprimível, e a discussão troca por afirmação.
//   - o PUNCH, que é a consequência clínica. Sai vazio; inventá-lo seria escrever leitura médica.
//   - o enquadramento temático da seção ("O que o cigarro está cobrando").
//
// A divisão é essa: o código põe o que é verdade do prontuário, a conversa põe o que é leitura.

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/pdfdoc"
)

// MontaDeckLocal devolve o deck completo a partir do dossiê, sem chamar modelo nenhum.
func MontaDeckLocal(d *dto.PlanDossierResponse) []pdfdoc.DeckSlide {
	if d == nil {
		return nil
	}
	secoes := montaArco(d)
	porSecao := make([][]pdfdoc.DeckSlide, len(secoes))
	for i, sec := range secoes {
		porSecao[i] = montaSecao(d, sec, secoes)
	}
	// As mesmas regras do servidor que já valiam para o caminho com LLM: eyebrow numerado,
	// variante escura só em capa e fecho, nome do paciente na capa.
	slides := aplicaRegrasDoDeck(secoes, porSecao,
		primeiroNomeDe(d.Patient.Name), formataDataPT(d.GeneratedAt))
	slides, _ = EnsureSlideIDs(slides)
	slides, _ = hidrataReguas(slides, d)
	poeLegenda(slides)
	return slides
}

func montaSecao(d *dto.PlanDossierResponse, sec PlanArcSection, todas []PlanArcSection) []pdfdoc.DeckSlide {
	switch sec.Key {
	case SecCapa:
		return []pdfdoc.DeckSlide{{Kind: pdfdoc.DeckCover}}
	case SecResumo:
		return []pdfdoc.DeckSlide{montaResumo(d)}
	case SecBem, SecMovendo:
		return montaSlidesDeRegua(d, sec)
	case SecFaltando:
		return montaExamesQueFaltam(d)
	case SecPlano:
		return montaCondutas(d, sec)
	case SecSequencia:
		return montaSequencia(d)
	case SecLevar:
		return montaParaLevar(d)
	case SecFecho:
		return []pdfdoc.DeckSlide{montaFecho(d, todas)}
	}
	return nil
}

// montaFecho — o fecho com o retrato em NÚMEROS, não em conclusão.
//
// O fecho dos decks aprovados é o parágrafo mais escrito de todos, e é leitura pura: não sai daqui.
// O que sai é a contagem, que é verdade e evita a última página do relatório ser um título sozinho
// num fundo escuro.
func montaFecho(d *dto.PlanDossierResponse, secoes []PlanArcSection) pdfdoc.DeckSlide {
	// Conta o que o DECK mostra, não o que o dossiê tem.
	//
	// `d.Strong` traz 50 achados nesta paciente, dos quais o deck desenha nove: dizer "50 exames nos
	// melhores patamares" descreveria um relatório que o paciente não está lendo. E a lista inteira
	// carrega o checklist de ausência, que infla o lado bom do escore.
	bem, movendo, condutas := 0, 0, 0
	for _, s := range secoes {
		switch s.Key {
		case SecBem:
			bem = len(s.Exames)
		case SecMovendo:
			movendo = len(s.Exames)
		case SecPlano:
			// Os slides do plano, não as 26 linhas do plano de cuidado: o deck mostra seis.
			condutas = s.Slides
		}
	}
	var partes []string
	if bem > 0 {
		partes = append(partes, fmt.Sprintf("%d exames nos melhores patamares da escala", bem))
	}
	if movendo > 0 {
		partes = append(partes, fmt.Sprintf("%d que se movem e merecem acompanhamento", movendo))
	}
	if condutas > 0 {
		partes = append(partes, fmt.Sprintf("%d condutas para os próximos três meses", condutas))
	}
	s := pdfdoc.DeckSlide{Kind: pdfdoc.DeckClosing, Title: "O que este relatório mostra, em uma página"}
	if len(partes) > 0 {
		s.Lede = "Este relatório reúne " + enumeraSimples(partes) + "."
	}
	return s
}

// enumeraSimples junta com vírgula e "e", sem teto de tamanho.
func enumeraSimples(ps []string) string {
	switch len(ps) {
	case 0:
		return ""
	case 1:
		return ps[0]
	}
	return strings.Join(ps[:len(ps)-1], ", ") + " e " + ps[len(ps)-1]
}

// ---------------------------------------------------------------------------
// resumo

func montaResumo(d *dto.PlanDossierResponse) pdfdoc.DeckSlide {
	s := pdfdoc.DeckSlide{Kind: pdfdoc.DeckSummary, Title: "Onde você está, em uma página"}
	sum := &pdfdoc.DeckSummaryBlock{StepsTitle: "O que vamos fazer"}
	if c := cartaoDeResumo(d, d.Strong, "O que está forte", "bom"); c != nil {
		sum.Cards = append(sum.Cards, *c)
	}
	if c := cartaoDeResumo(d, d.Moving, "O que está se movendo", "ruim"); c != nil {
		sum.Cards = append(sum.Cards, *c)
	}
	// Os passos são as condutas de maior prioridade, com as palavras que o médico já escreveu. Sai
	// da mesma seleção da seção "O plano", com teto menor: o resumo prometia prioridade e entregava
	// as quatro primeiras da ordem alfabética do dossiê, que podem não ser nenhuma das urgentes.
	for _, c := range condutasPorPrioridade(d.CarePlan, passosNoResumo) {
		if t := strings.TrimSpace(c.Recommendation); t != "" {
			sum.Steps = append(sum.Steps, encurta(t, 90))
		}
	}
	s.Summary = sum
	return s
}

// cartaoDeResumo — quatro linhas, que é o teto do layout.
func cartaoDeResumo(d *dto.PlanDossierResponse, fs []dto.PlanFinding, titulo, tom string) *pdfdoc.DeckSummaryCard {
	c := pdfdoc.DeckSummaryCard{Title: titulo, Tone: tom}
	for _, f := range fs {
		if len(c.Lines) >= 4 {
			break
		}
		// `Stale` fora, pela mesma razão que o arco o exclui das réguas (patient_plan_arco_local.go):
		// medida de dois anos atrás não é o retrato de hoje. Sem isto, o slide 2 podia liderar com
		// um achado que os slides de régua deixaram de fora de propósito.
		if f.Source != dto.PlanSourceLab || f.Code == "" || f.Stale {
			continue
		}
		r, ok := d.Rulers[f.Code]
		if !ok {
			continue
		}
		c.Lines = append(c.Lines, pdfdoc.DeckSummaryLine{
			Name: nomeParaPaciente(r), Sub: r.Gloss, Code: f.Code,
			Value: f.Text, Unit: f.Unit,
		})
	}
	if len(c.Lines) == 0 {
		return nil
	}
	return &c
}

// ---------------------------------------------------------------------------
// réguas

func montaSlidesDeRegua(d *dto.PlanDossierResponse, sec PlanArcSection) []pdfdoc.DeckSlide {
	if len(sec.Exames) == 0 || sec.Slides == 0 {
		return nil
	}
	// Distribui os exames pelos slides que o arco reservou, o mais parelho possível: 9 exames em
	// 3 slides são 3+3+3; 8 em 3 são 3+3+2. Nunca sobra slide com uma régua sozinha.
	porSlide := repartir(len(sec.Exames), sec.Slides)
	var out []pdfdoc.DeckSlide
	pos := 0
	for _, n := range porSlide {
		var blocos []pdfdoc.DeckRulerBlock
		var nomes []string
		for i := 0; i < n && pos < len(sec.Exames); i++ {
			code := sec.Exames[pos]
			pos++
			r, ok := d.Rulers[code]
			if !ok {
				continue
			}
			nome := nomeParaPaciente(r)
			nomes = append(nomes, nome)
			blocos = append(blocos, pdfdoc.DeckRulerBlock{Ruler: pdfdoc.Ruler{
				Code: code, Display: nome, Sub: r.Gloss, Note: notaDaRegua(r),
			}})
		}
		if len(blocos) == 0 {
			continue
		}
		out = append(out, pdfdoc.DeckSlide{
			Kind: pdfdoc.DeckRulers, Rulers: blocos, Title: enumera(nomes),
		})
	}
	return out
}

// notaDaRegua — o rótulo avaliativo, e depois a série.
//
// A régua não pode entrar num slide sem um rótulo avaliativo visível no MESMO slide, e neste deck o
// punch está vazio de propósito: sobra o `note`. O rótulo sai do NÍVEL do escore, que é dado, não
// leitura — é a mesma classificação que pinta a barra, escrita por extenso.
//
// A série ("239 em 2024, 432 em 2025, 500 agora") é a assinatura do `note` nos decks aprovados: a
// barra desenha só os dois últimos pontos, então é o texto que mostra a direção no tempo.
func notaDaRegua(r dto.PlanRuler) string {
	partes := []string{}
	if r := rotuloDoNivel(nivelAtual(r)); r != "" {
		partes = append(partes, r)
	}
	if s := serieDoHistorico(r); s != "" {
		partes = append(partes, s)
	}
	return strings.Join(partes, " ")
}

// nivelAtual — o nível da ÚLTIMA medida, que é a que o texto imprime como "agora".
//
// Andar para trás procurando um ponto classificado (a primeira versão) dava um rótulo que não era
// do número ao lado: "No melhor patamar da escala. … 500 agora", com o veredicto pertencendo a uma
// medida antiga. Quando o último ponto não traz nível, o servidor classifica pelo valor contra os
// próprios segmentos da régua, que é a mesma conta que pinta a barra.
func nivelAtual(r dto.PlanRuler) int {
	if len(r.History) == 0 {
		return -1
	}
	ultimo := r.History[len(r.History)-1]
	if ultimo.Level != nil {
		return *ultimo.Level
	}
	for _, seg := range r.Segments {
		if ultimo.Value >= seg.A && ultimo.Value <= seg.B {
			return seg.Level
		}
	}
	return -1
}

// rotuloDoNivel traduz o nível do escore para a mesma frase, sempre.
//
// Vocabulário fixo de propósito: dizer "no melhor patamar" para nível 5 é ler a escala em voz alta,
// não emitir juízo clínico. O juízo é o punch, e o punch fica em branco.
//
// E fala de POSIÇÃO na escala, nunca de direção do valor. Muitos itens são de faixa ótima no meio
// (a hemoglobina glicada de 5,9 cai no nível 1 por estar ALTA), e "abaixo do ideal" mandava ao
// paciente exatamente a leitura contrária da que a barra desenha.
func rotuloDoNivel(n int) string {
	switch n {
	case 5:
		return "No melhor patamar da escala."
	case 4:
		return "Perto do melhor patamar."
	case 3:
		return "No meio da escala."
	case 2:
		return "Longe do ideal."
	case 1:
		return "Quase no pior patamar."
	case 0:
		return "No pior patamar da escala."
	}
	return ""
}

func serieDoHistorico(r dto.PlanRuler) string {
	if len(r.History) < 2 {
		return ""
	}
	h := r.History
	if len(h) > 3 {
		h = h[len(h)-3:]
	}
	// Duas coletas no MESMO ano viram "0,37 em 2026, 63,10 agora", que lê como se o ano explicasse
	// a diferença: quando o ano se repete, quem marca o tempo é o mês.
	//
	// Mas a troca é POR PONTO. Trocar todos pelo mês porque DOIS coincidem produzia "X em novembro,
	// Y em março, Z agora" para coletas de 2024, 2026 e 2026 — a cronologia lida ao contrário, que
	// é exatamente o que se queria evitar. Só quando a série INTEIRA cabe num ano o mês vai sozinho.
	umAnoSo := true
	for _, p := range h {
		if anoDe(p.Date) != anoDe(h[len(h)-1].Date) {
			umAnoSo = false
		}
	}
	repetido := map[string]int{}
	for _, p := range h {
		repetido[anoDe(p.Date)]++
	}
	var partes []string
	for i, p := range h {
		if i == len(h)-1 {
			partes = append(partes, p.Text+" agora")
			continue
		}
		quando := anoDe(p.Date)
		switch {
		case umAnoSo:
			quando = mesDe(p.Date)
		case repetido[quando] > 1:
			quando = mesDe(p.Date) + " de " + quando
		}
		partes = append(partes, p.Text+" em "+quando)
	}
	return strings.Join(partes, ", ") + "."
}

// ---------------------------------------------------------------------------
// exames que faltam

func montaExamesQueFaltam(d *dto.PlanDossierResponse) []pdfdoc.DeckSlide {
	if d.LabRequest == nil || strings.TrimSpace(d.LabRequest.Exams) == "" {
		return nil
	}
	var linhas []pdfdoc.DeckTableRow
	total, inicioDeBloco := 0, true
	for _, l := range strings.Split(d.LabRequest.Exams, "\n") {
		t := strings.TrimSpace(l)
		// Linha em branco separa blocos no pedido, e é o que marca onde um cabeçalho pode estar.
		if t == "" {
			inicioDeBloco = true
			continue
		}
		// Linha de justificativa começa com "#": adere ao exame de cima, como no PDF do pedido.
		// Consecutivas concatenam, igual a `parseExamBlocks` — sobrescrever perderia a segunda.
		if strings.HasPrefix(t, "#") {
			j := strings.TrimSpace(strings.TrimPrefix(t, "#"))
			if j != "" && len(linhas) > 0 {
				if atual := linhas[len(linhas)-1].Cells[1]; atual != "" {
					j = atual + " " + j
				}
				linhas[len(linhas)-1].Cells[1] = j
			}
			continue
		}
		if ehCabecalhoDeGrupo(t, inicioDeBloco) {
			inicioDeBloco = false
			continue
		}
		inicioDeBloco = false
		total++
		if len(linhas) >= 8 {
			continue
		}
		linhas = append(linhas, pdfdoc.DeckTableRow{Cells: []string{t, ""}})
	}
	if len(linhas) == 0 {
		return nil
	}
	// A coluna "Por quê" só existe se o pedido trouxe justificativa. Cabeçalho sobre oito células
	// vazias é pior que uma coluna a menos: promete uma explicação que a página não dá.
	temPorque := false
	for _, l := range linhas {
		if l.Cells[1] != "" {
			temPorque = true
		}
	}
	colunas := []pdfdoc.DeckTableCol{{Label: "Exame", Width: "560px"}}
	if temPorque {
		colunas = append(colunas, pdfdoc.DeckTableCol{Label: "Por quê", Style: "why"})
	} else {
		for i := range linhas {
			linhas[i].Cells = linhas[i].Cells[:1]
		}
	}
	// O corte se ANUNCIA. Oito linhas é o que cabe no slide; sumir com as outras em silêncio faria
	// o paciente ler a lista como se fosse o pedido inteiro.
	if resto := total - len(linhas); resto > 0 {
		celulas := []string{fmt.Sprintf("e mais %d exames no pedido", resto)}
		if len(colunas) > 1 {
			celulas = append(celulas, "")
		}
		linhas = append(linhas, pdfdoc.DeckTableRow{Cells: celulas})
	}
	return []pdfdoc.DeckSlide{{
		Kind:  pdfdoc.DeckTableKind,
		Title: "Os exames que ainda não voltaram",
		Table: &pdfdoc.DeckTable{Dense: len(linhas) > 4, Columns: colunas, Rows: linhas},
	}}
}

// gruposDoPedido — as palavras que o médico usa para separar o pedido em seções.
//
// Vocabulário FECHADO, e não uma regra de forma. A primeira versão pulava qualquer linha toda em
// maiúscula com até 24 caracteres, e isso apagava exame de verdade: no pedido desta paciente,
// `ACTH`, `DHEA-S`, `HIV 1+2` e `VDRL` sumiriam da tabela sem aviso nenhum, e a justificativa
// seguinte grudaria no exame errado. Mostrar "SANGUE" como linha é feio; sumir com o HIV não é.
var gruposDoPedido = []string{
	"SANGUE", "URINA", "FEZES", "SOROLOGIA", "SOROLOGIAS", "HORMÔNIO", "HORMÔNIOS",
	"IMAGEM", "IMAGENS", "GENÉTICA", "MICROBIOMA", "EXAMES", "OUTROS",
}

// ehCabecalhoDeGrupo — só no INÍCIO de um bloco, só em maiúsculas, e só se a primeira palavra é uma
// das do vocabulário. "HORMÔNIOS — COLETA SEM DEXAMETASONA" entra por ela; "ACTH", que vem logo
// abaixo, não.
func ehCabecalhoDeGrupo(t string, inicioDeBloco bool) bool {
	if !inicioDeBloco || t != strings.ToUpper(t) {
		return false
	}
	primeira := strings.FieldsFunc(t, func(r rune) bool { return r == ' ' || r == '—' || r == '-' || r == ':' })
	if len(primeira) == 0 {
		return false
	}
	for _, g := range gruposDoPedido {
		if primeira[0] == g {
			return true
		}
	}
	return false
}

// ---------------------------------------------------------------------------
// o plano, uma conduta por slide

func montaCondutas(d *dto.PlanDossierResponse, sec PlanArcSection) []pdfdoc.DeckSlide {
	var out []pdfdoc.DeckSlide
	// A MESMA seleção que `montaArco` fez, e não `d.CarePlan` cru: o dossiê vem ordenado por letra
	// AGIR antes da prioridade, então iterá-lo aqui montava seis slides alfabéticos enquanto o arco
	// (e o prompt que ele gera) prometia os seis mais urgentes. A função é pura, então chamá-la de
	// novo com o mesmo teto devolve exatamente a mesma escolha.
	for _, c := range condutasPorPrioridade(d.CarePlan, sec.Slides) {
		s := pdfdoc.DeckSlide{Kind: pdfdoc.DeckPlanStep, Title: tituloDaConduta(c.Recommendation)}
		var cards []pdfdoc.DeckCard
		// `rationale` e `target` são texto que o MÉDICO já escreveu no plano de cuidado. Copiar é
		// mais fiel que reescrever, e não custa nada.
		if t := texto(c.Rationale); t != "" {
			cards = append(cards, pdfdoc.DeckCard{Kicker: "Por que entra agora", Body: t})
		}
		if t := texto(c.Target); t != "" {
			cards = append(cards, pdfdoc.DeckCard{Kicker: "O que se espera", Body: t})
		}
		s.Cards = cards
		out = append(out, s)
	}
	return out
}

// tituloDaConduta — a conduta vira título cortando na CLÁUSULA, não no caractere.
//
// "Musculação três vezes por semana, começando já na semana 1, junto com a tirzepatida" cortada por
// tamanho virava "Musculação três vezes por semana, começando já na…", que termina no meio de uma
// ideia. A primeira oração já é o título ("Musculação três vezes por semana"), e o resto continua
// inteiro nos cartões abaixo.
func tituloDaConduta(rec string) string {
	t := strings.TrimSpace(rec)
	if i := strings.IndexAny(t, ",;:"); i >= 16 && i <= 53 {
		t = t[:i]
	}
	return strings.TrimRight(encurta(t, 53), ".")
}

// ---------------------------------------------------------------------------
// a sequência

func montaSequencia(d *dto.PlanDossierResponse) []pdfdoc.DeckSlide {
	// Tabela de 2 colunas, com a primeira em `dose` e valores relativos: é assim nos DOIS decks
	// aprovados, e o kind `sequence` nunca foi usado em nenhum.
	linhas := []pdfdoc.DeckTableRow{}
	// "Agora" é a conduta mais urgente, não a primeira do alfabeto: mesma seleção da seção do
	// plano, com teto 1. Antes, a linha que abre a linha do tempo podia ser um item de prioridade
	// baixa que o próprio plano nem lidera.
	if primeira := condutasPorPrioridade(d.CarePlan, primeiroPasso); len(primeira) > 0 {
		linhas = append(linhas, pdfdoc.DeckTableRow{
			Cells: []string{"Agora", encurta(primeira[0].Recommendation, 90)}})
	}
	if d.LabRequest != nil && strings.TrimSpace(d.LabRequest.Exams) != "" {
		linhas = append(linhas, pdfdoc.DeckTableRow{
			Cells: []string{"Nas próximas semanas", "Coletar os exames do pedido."}})
	}
	linhas = append(linhas, pdfdoc.DeckTableRow{
		Cells: []string{"Em 3 meses", "Refazer os exames que estão se movendo e rever o plano."}})
	if len(linhas) < 2 {
		return nil
	}
	return []pdfdoc.DeckSlide{{
		Kind:  pdfdoc.DeckTableKind,
		Title: "Os próximos três meses, em ordem",
		Table: &pdfdoc.DeckTable{
			Columns: []pdfdoc.DeckTableCol{
				{Label: "Quando", Style: "dose", Width: "300px"},
				{Label: "O que acontece"},
			},
			Rows: linhas,
		},
	}}
}

// ---------------------------------------------------------------------------
// para levar

func montaParaLevar(d *dto.PlanDossierResponse) []pdfdoc.DeckSlide {
	if len(d.Medications) == 0 {
		return nil
	}
	take := &pdfdoc.DeckTakeawayBox{}
	grupos := 0
	for _, p := range d.Medications {
		grupos += len(p.Formulas) + len(p.Medications)
	}
	for _, p := range d.Medications {
		for _, f := range p.Formulas {
			if len(take.Groups) >= 3 {
				break
			}
			g := pdfdoc.DeckDoseGroup{Title: tituloDaFormula(f)}
			for _, c := range f.Components {
				g.Items = append(g.Items, pdfdoc.DeckDose{
					Name: c.Substance,
					Dose: strings.TrimSpace(fmt.Sprintf("%s %s", formataNumero(c.Quantity), c.Unit)),
				})
			}
			take.Groups = append(take.Groups, g)
		}
		for _, m := range p.Medications {
			if len(take.Groups) >= 3 {
				break
			}
			take.Groups = append(take.Groups, pdfdoc.DeckDoseGroup{
				Title: m.Name,
				Items: []pdfdoc.DeckDose{{Name: m.Concentration, Sub: m.Frequency, Dose: m.Dosage}},
			})
		}
	}
	if len(take.Groups) == 0 {
		return nil
	}
	take.Note = "A receita assinada sai à parte, com a validade e o registro."
	// Três grupos é o que cabe na largura. O que ficou de fora existe na receita, e dizer quantos
	// evita o paciente ler a página como se fosse tudo o que ele vai tomar.
	if resto := grupos - len(take.Groups); resto > 0 {
		take.Note = fmt.Sprintf("Mais %d itens vão na receita assinada, que sai à parte com a validade e o registro.", resto)
	}
	return []pdfdoc.DeckSlide{{
		Kind: pdfdoc.DeckTakeaway, Title: "O que você começa a tomar agora", Take: take,
	}}
}

// tituloDaFormula — o nome, e a posologia só quando ela cabe INTEIRA.
//
// Posologia cortada no meio ("Tomar 1 dose no…") é pior que posologia ausente: o paciente lê meia
// instrução de como tomar. A instrução completa sai na receita assinada, que vai à parte.
func tituloDaFormula(f dto.PlanDossierFormula) string {
	nome := encurta(f.Name, 46)
	if t := strings.TrimSpace(f.Posology); t != "" {
		if junto := f.Name + " · " + t; len([]rune(junto)) <= 46 {
			return junto
		}
	}
	return nome
}

// ---------------------------------------------------------------------------
// utilitários

// nomeParaPaciente devolve o nome do CATÁLOGO, nunca o do score.
//
// O do score traz a variante ("Ferritina - Mulheres Pós-Menopausa"), que a gramática proíbe pôr na
// régua: o paciente não tem que ler a mecânica do escore para entender o próprio exame.
func nomeParaPaciente(r dto.PlanRuler) string {
	if t := strings.TrimSpace(r.PatientName); t != "" {
		return semTravessao(t)
	}
	// Sem o nome do catálogo, corta a variante do nome do score no primeiro " - ".
	if i := strings.Index(r.Name, " - "); i > 0 {
		return semTravessao(r.Name[:i])
	}
	return semTravessao(r.Name)
}

// semTravessao corta o nome NO travessão, em vez de trocá-lo por vírgula.
//
// O que vem depois dele no catálogo é qualificação da amostra ("Microalbuminúria/creatininúria —
// urina isolada"), e o paciente reconhece o exame pelo que vem antes. Trocar por vírgula só mudava
// a pontuação de um nome que já não cabia na linha.
var reTravessaoNome = regexp.MustCompile(`\s*[—–].*$`)

func semTravessao(s string) string {
	return strings.TrimSpace(reTravessaoNome.ReplaceAllString(s, ""))
}

// repartir divide n itens em k grupos o mais parelho possível, do maior para o menor.
func repartir(n, k int) []int {
	if k <= 0 {
		return nil
	}
	base, resto := n/k, n%k
	out := make([]int, k)
	for i := range out {
		out[i] = base
		if i < resto {
			out[i]++
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(out)))
	return out
}

// enumera monta "A, B e C" — o título descritivo do slide de régua.
//
// Com teto de 53 caracteres, que é o topo da faixa medida: nome comprido de catálogo
// ("Microalbuminúria/creatininúria, relação") estourava a linha e o título vazava a moldura. O que
// não cabe vira contagem, e a régua abaixo diz o nome inteiro de qualquer jeito.
func enumera(nomes []string) string {
	junta := func(ns []string) string {
		switch len(ns) {
		case 0:
			return ""
		case 1:
			return ns[0]
		case 2:
			return ns[0] + " e " + ns[1]
		}
		return strings.Join(ns[:len(ns)-1], ", ") + " e " + ns[len(ns)-1]
	}
	for n := len(nomes); n >= 2; n-- {
		t := junta(nomes[:n])
		if n < len(nomes) {
			// Vírgula, e não `junta`: "Alfafetoproteína e Sódio e mais 1" põe dois "e" na mesma
			// frase. O "e" fica sendo o da contagem.
			resto := "e mais um"
			if d := len(nomes) - n; d > 1 {
				resto = fmt.Sprintf("e mais %d", d)
			}
			t = strings.Join(nomes[:n], ", ") + " " + resto
		}
		if len([]rune(t)) <= 53 {
			return t
		}
	}
	// Um nome só, e ainda comprido: corta, porque a alternativa é o título sair da moldura.
	return encurta(nomes[0], 53)
}

func encurta(s string, n int) string {
	s = strings.TrimSpace(s)
	r := []rune(s)
	if len(r) <= n {
		return s
	}
	// Corta na palavra, não no meio dela.
	corte := string(r[:n])
	if i := strings.LastIndex(corte, " "); i > n/2 {
		corte = corte[:i]
	}
	return strings.TrimRight(corte, " ,.;:") + "…"
}

// texto lê um campo opcional sem explodir em nil.
func texto(p *string) string {
	if p == nil {
		return ""
	}
	return strings.TrimSpace(*p)
}

// mesDe devolve o mês por extenso de uma data AAAA-MM-DD.
func mesDe(data string) string {
	meses := []string{"janeiro", "fevereiro", "março", "abril", "maio", "junho",
		"julho", "agosto", "setembro", "outubro", "novembro", "dezembro"}
	if len(data) >= 7 {
		if m, err := strconv.Atoi(data[5:7]); err == nil && m >= 1 && m <= 12 {
			return meses[m-1]
		}
	}
	return anoDe(data)
}

func anoDe(data string) string {
	if len(data) >= 4 {
		return data[:4]
	}
	return data
}

func formataNumero(v float64) string {
	s := fmt.Sprintf("%.4f", v)
	s = strings.TrimRight(strings.TrimRight(s, "0"), ".")
	return strings.ReplaceAll(s, ".", ",")
}
