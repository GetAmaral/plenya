package services

import (
	"strings"
	"testing"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/pdfdoc"
)

// conduta monta um item do plano com a letra e a prioridade que interessam ao corte.
func conduta(letra, prioridade, recomendacao string) dto.CarePlanItemResponse {
	return dto.CarePlanItemResponse{
		LetterCode:     letra,
		Priority:       prioridade,
		Recommendation: recomendacao,
	}
}

// O caso que motivou a função: o dossiê chega ordenado por LETRA antes da prioridade, porque é
// assim que a tela do plano agrupa. Com mais de 6 condutas, cortar nessa ordem deixa de fora
// justamente as urgentes das letras finais.
func TestCondutasPorPrioridadeCortaPelaUrgencia(t *testing.T) {
	// Ordem de entrada = a de ListByPatient: letra ASC, depois prioridade. As três altas estão nas
	// letras do fim, que é exatamente onde o corte alfabético as perdia.
	dossie := []dto.CarePlanItemResponse{
		conduta("A", "low", "A-baixa-1"),
		conduta("A", "low", "A-baixa-2"),
		conduta("A", "low", "A-baixa-3"),
		conduta("G", "low", "G-baixa-1"),
		conduta("G", "low", "G-baixa-2"),
		conduta("G", "low", "G-baixa-3"),
		conduta("I", "high", "I-alta"),
		conduta("R", "high", "R-alta"),
		conduta("R", "medium", "R-media"),
	}

	got := recomendacoes(condutasPorPrioridade(dossie, 6))

	if len(got) != 6 {
		t.Fatalf("esperava 6 condutas, veio %d", len(got))
	}
	// Antes da correção nenhuma das três chegava ao deck: eram a 7ª, a 8ª e a 9ª da lista.
	esperadoNoTopo := []string{"I-alta", "R-alta", "R-media"}
	for i, c := range esperadoNoTopo {
		if got[i] != c {
			t.Fatalf("as urgentes têm de liderar o corte; esperava %v no topo, veio %v", esperadoNoTopo, got)
		}
	}
	// E as três que sobraram são as de prioridade baixa que vinham primeiro no dossiê.
	if got[3] != "A-baixa-1" || got[5] != "A-baixa-3" {
		t.Errorf("o resto do corte deveria seguir a ordem do dossiê: %v", got)
	}
}

// Dentro da mesma prioridade a ordem do dossiê (letra, depois recência) tem de sobreviver, senão o
// deck deixa de sair agrupado por pilar.
func TestCondutasPorPrioridadeEstavelDentroDaFaixa(t *testing.T) {
	dossie := []dto.CarePlanItemResponse{
		conduta("A", "high", "A-alta"),
		conduta("G", "high", "G-alta"),
		conduta("I", "high", "I-alta"),
		conduta("R", "high", "R-alta"),
	}

	got := recomendacoes(condutasPorPrioridade(dossie, 6))

	esperado := []string{"A-alta", "G-alta", "I-alta", "R-alta"}
	for i := range esperado {
		if got[i] != esperado[i] {
			t.Fatalf("ordem dentro da mesma prioridade mudou: %v", got)
		}
	}
}

// Prioridade fora de high/medium cai no fim da fila, como o ELSE de `carePlanPriorityOrder`. O
// CHECK da coluna hoje impede o caso, mas o dia em que uma prioridade nova entrar as duas ordens
// têm de andar juntas, senão o item fica no meio de uma listagem e no fim da outra.
func TestCondutasPorPrioridadeEspelhaOElseDoSQL(t *testing.T) {
	dossie := []dto.CarePlanItemResponse{
		conduta("A", "low", "baixa"),
		conduta("A", "", "sem-prioridade"),
		conduta("A", "medium", "media"),
		conduta("A", "high", "alta"),
	}

	got := recomendacoes(condutasPorPrioridade(dossie, 6))

	// low e vazio empatam em 2, e o desempate é a ordem do dossiê.
	esperado := []string{"alta", "media", "baixa", "sem-prioridade"}
	for i := range esperado {
		if got[i] != esperado[i] {
			t.Fatalf("esperava %v, veio %v", esperado, got)
		}
	}
}

// O corte não pode devolver mais slides do que o teto, e com menos condutas que o teto devolve
// todas: `montaArco` usa o tamanho da volta como número de slides da seção.
func TestCondutasPorPrioridadeRespeitaOTeto(t *testing.T) {
	var muitas []dto.CarePlanItemResponse
	for i := 0; i < 9; i++ {
		muitas = append(muitas, conduta("G", "medium", "conduta"))
	}
	if got := condutasPorPrioridade(muitas, 6); len(got) != 6 {
		t.Errorf("teto de 6 não respeitado: %d", len(got))
	}
	if got := condutasPorPrioridade(muitas[:2], 6); len(got) != 2 {
		t.Errorf("com 2 condutas esperava 2, veio %d", len(got))
	}
	if got := condutasPorPrioridade(nil, 6); len(got) != 0 {
		t.Errorf("sem conduta a seção não existe, veio %d", len(got))
	}
}

// A entrada não pode ser reordenada no lugar: o dossiê é compartilhado com o resto da montagem,
// que conta com a ordem por pilar.
func TestCondutasPorPrioridadeNaoMexeNaEntrada(t *testing.T) {
	dossie := []dto.CarePlanItemResponse{
		conduta("A", "low", "primeira"),
		conduta("G", "high", "segunda"),
	}

	condutasPorPrioridade(dossie, 6)

	if dossie[0].Recommendation != "primeira" || dossie[1].Recommendation != "segunda" {
		t.Errorf("a ordem do dossiê foi alterada: %v", dossie)
	}
}

// O deck montado tem de trazer as MESMAS condutas que o arco escolheu.
//
// É a regressão que motivou este teste: `montaArco` passou a ordenar por prioridade, mas
// `montaCondutas` continuava iterando `d.CarePlan` cru e só usava `sec.Slides` como contagem.
// O arco (e o prompt que ele gera) prometia as seis urgentes, e o deck montava as seis
// alfabéticas — dois caminhos de geração divergindo a partir do mesmo dossiê. Nenhum teste de
// unidade da função de corte pega isso; só olhar o deck pronto pega.
func TestDeckMontaAsCondutasQueOArcoEscolheu(t *testing.T) {
	d := &dto.PlanDossierResponse{
		Patient: dto.PlanDossierPatient{Name: "Fulana de Tal"},
		CarePlan: []dto.CarePlanItemResponse{
			conduta("A", "low", "A-baixa-1"),
			conduta("A", "low", "A-baixa-2"),
			conduta("A", "low", "A-baixa-3"),
			conduta("G", "low", "G-baixa-1"),
			conduta("G", "low", "G-baixa-2"),
			conduta("G", "low", "G-baixa-3"),
			conduta("I", "high", "I-alta"),
			conduta("R", "high", "R-alta"),
		},
	}

	slides := MontaDeckLocal(d)

	var doPlano []string
	for _, s := range slides {
		if s.Kind == pdfdoc.DeckPlanStep {
			doPlano = append(doPlano, s.Title)
		}
	}
	if len(doPlano) != maxCondutasNoPlano {
		t.Fatalf("esperava %d slides de conduta, vieram %d: %v", maxCondutasNoPlano, len(doPlano), doPlano)
	}
	// As duas urgentes estão nas letras do fim do alfabeto, então eram as duas últimas do dossiê.
	for _, urgente := range []string{"I-alta", "R-alta"} {
		if !contemTitulo(doPlano, urgente) {
			t.Errorf("a conduta urgente %q não chegou ao deck: %v", urgente, doPlano)
		}
	}

	// O resumo e a linha "Agora" da sequência mostram recortes do MESMO plano, e não podem abrir
	// com uma conduta de prioridade baixa enquanto a seção do plano lidera com as urgentes.
	for _, s := range slides {
		if s.Kind == pdfdoc.DeckSummary && s.Summary != nil {
			if len(s.Summary.Steps) == 0 || !contemTitulo(s.Summary.Steps, "I-alta") {
				t.Errorf("o resumo não abriu pela conduta mais urgente: %v", s.Summary.Steps)
			}
		}
		if s.Kind == pdfdoc.DeckTableKind && s.Table != nil && len(s.Table.Rows) > 0 {
			primeira := s.Table.Rows[0].Cells
			if len(primeira) == 2 && primeira[0] == "Agora" && !contemTitulo([]string{primeira[1]}, "I-alta") {
				t.Errorf(`a linha "Agora" não é a conduta mais urgente: %q`, primeira[1])
			}
		}
	}
}

// contemTitulo aceita o título encurtado que o deck aplica.
func contemTitulo(lista []string, alvo string) bool {
	for _, s := range lista {
		if strings.Contains(s, alvo) {
			return true
		}
	}
	return false
}
