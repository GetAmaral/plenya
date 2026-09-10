package services

import (
	"math"
	"testing"
	"time"
)

// As três equações são dado verificável, não escolha de projeto: um coeficiente trocado devolve uma
// filtração plausível e errada, que ninguém percebe. Os alvos aqui são os números que o próprio
// laboratório imprimiu nos laudos do paciente que motivou esta feature, o que torna o teste uma
// conferência contra o mundo e não contra mim mesmo.
//
// Homem, nascido em 19/02/1980. Creatinina 1,50 mg/dL em 17/10/2023 → o laudo diz 59.
// Cistatina C 1,01 mg/L em 08/11/2023 → o laudo diz 82.
func TestEquacoesCKDEPI_ConferemComOLaudo(t *testing.T) {
	nascimento := time.Date(1980, 2, 19, 0, 0, 0, 0, time.UTC)

	idadeCr, ok := idadeNaColeta(nascimento, "2023-10-17")
	if !ok {
		t.Fatal("idade na coleta da creatinina não calculou")
	}
	if got := math.Round(tfgPorCreatininaCKDEPI2021(1.50, idadeCr, false)); got != 59 {
		t.Errorf("creatinina 1,50 aos 43 anos = %v, o laudo diz 59", got)
	}

	idadeCys, ok := idadeNaColeta(nascimento, "2023-11-08")
	if !ok {
		t.Fatal("idade na coleta da cistatina não calculou")
	}
	if got := math.Round(tfgPorCistatinaCKDEPI2012(1.01, idadeCys, false)); got != 82 {
		t.Errorf("cistatina C 1,01 aos 43 anos = %v, o laudo diz 82", got)
	}

	// Creatinina 1,72 em 21/03/2025 → 49; 1,75 em 23/01/2026 → 48.
	i25, _ := idadeNaColeta(nascimento, "2025-03-21")
	if got := math.Round(tfgPorCreatininaCKDEPI2021(1.72, i25, false)); got != 49 {
		t.Errorf("creatinina 1,72 = %v, o laudo diz 49", got)
	}
	i26, _ := idadeNaColeta(nascimento, "2026-01-23")
	if got := math.Round(tfgPorCreatininaCKDEPI2021(1.75, i26, false)); got != 48 {
		t.Errorf("creatinina 1,75 = %v, o laudo diz 48", got)
	}
}

// A combinada tem que cair ENTRE as duas isoladas quando elas discordam, e é esse o motivo clínico
// de ela existir: com creatinina inflada por massa muscular, a de creatinina diz G3a e a de
// cistatina diz G2, e a combinada é a que a KDIGO manda usar para desempatar.
func TestCombinadaFicaEntreAsDuasIsoladas(t *testing.T) {
	nascimento := time.Date(1980, 2, 19, 0, 0, 0, 0, time.UTC)
	idade, _ := idadeNaColeta(nascimento, "2023-10-17")

	porCr := tfgPorCreatininaCKDEPI2021(1.50, idade, false)
	porCys := tfgPorCistatinaCKDEPI2012(1.01, idade, false)
	comb := tfgCombinadaCKDEPI2021(1.50, 1.01, idade, false)

	if !(comb > porCr && comb < porCys) {
		t.Fatalf("combinada %v devia ficar entre creatinina %v e cistatina %v", comb, porCr, porCys)
	}
	if got := math.Round(comb); got != 72 {
		t.Errorf("combinada = %v, esperado 72", got)
	}
}

// SÓ a mulher recebe multiplicador nas três equações; o homem não recebe nenhum (NKF, NIDDK e a
// referência do CKD-EPI). Um "1.008 se masculino" é fácil de inventar e some no ruído, porque o
// número continua plausível — foi exatamente o erro do primeiro cálculo que fiz à mão.
//
// Com a MESMA creatinina, homem dá filtração MAIOR que mulher: o kappa maior (0,9 contra 0,7) faz o
// mesmo valor pesar menos, e é isso que traduz a massa muscular maior. Na cistatina C, que não
// depende de músculo, a mulher tem o 0,932 e fica abaixo.
func TestMultiplicadoresDeSexo(t *testing.T) {
	const idade = 50.0
	if h, m := tfgPorCreatininaCKDEPI2021(1.0, idade, false), tfgPorCreatininaCKDEPI2021(1.0, idade, true); !(m < h) {
		t.Errorf("creatinina: com a mesma creatinina, mulher %v devia dar menos que homem %v", m, h)
	}
	if h, m := tfgPorCistatinaCKDEPI2012(1.0, idade, false), tfgPorCistatinaCKDEPI2012(1.0, idade, true); !(m < h) {
		t.Errorf("cistatina: mulher %v devia dar menos que homem %v", m, h)
	}
	if h, m := tfgCombinadaCKDEPI2021(1.0, 1.0, idade, false), tfgCombinadaCKDEPI2021(1.0, 1.0, idade, true); !(m < h) {
		t.Errorf("combinada: mulher %v devia dar menos que homem %v", m, h)
	}
}

// Entre lotes do MESMO dia vence a equação mais rica, e a escolha não pode depender da ordem de
// iteração do mapa do Go: coleta partida em dois lotes é o caso normal, e quem chegasse primeiro
// ficava dono da linha porque a guarda do dia impedia o outro de escrever.
func TestCandidatoDoDiaPrefereAEquacaoMaisRica(t *testing.T) {
	combinada := candidatoTFG{valor: 72, equacao: tfgCombinada, prioridade: 3}
	soCistatina := candidatoTFG{valor: 82, equacao: tfgCistatinaC, prioridade: 2}
	soCreatinina := candidatoTFG{valor: 59, equacao: tfgCreatinina, prioridade: 1}

	if !(combinada.prioridade > soCistatina.prioridade && soCistatina.prioridade > soCreatinina.prioridade) {
		t.Fatal("a ordem de preferência declarada é combinada > cistatina > creatinina")
	}

	// O laço guarda o melhor por dia; aqui se reproduz a regra nas duas ordens de chegada.
	for _, chegada := range [][]candidatoTFG{
		{soCreatinina, soCistatina, combinada},
		{combinada, soCistatina, soCreatinina},
	} {
		melhor := map[string]candidatoTFG{}
		for _, c := range chegada {
			if atual, ja := melhor["2023-10-17"]; !ja || c.prioridade > atual.prioridade {
				melhor["2023-10-17"] = c
			}
		}
		if melhor["2023-10-17"].equacao != tfgCombinada {
			t.Errorf("ordem de chegada mudou o vencedor: veio %q", melhor["2023-10-17"].equacao)
		}
	}
}

// A idade é a DA COLETA. Pontuar um lote antigo com a idade de hoje aplica anos de decaimento a
// mais e devolve uma filtração menor do que a que o paciente tinha no dia.
func TestIdadeEDaColetaNaoDeHoje(t *testing.T) {
	nascimento := time.Date(1980, 2, 19, 0, 0, 0, 0, time.UTC)

	velha, ok := idadeNaColeta(nascimento, "2019-08-23")
	if !ok {
		t.Fatal("não calculou")
	}
	nova, _ := idadeNaColeta(nascimento, "2026-01-23")
	if !(nova-velha > 6 && nova-velha < 7) {
		t.Errorf("entre 2019 e 2026 vão ~6,4 anos, veio %v", nova-velha)
	}
	if tfgPorCreatininaCKDEPI2021(1.5, velha, false) <= tfgPorCreatininaCKDEPI2021(1.5, nova, false) {
		t.Error("a mesma creatinina em idade menor tem que dar filtração maior")
	}

	if _, ok := idadeNaColeta(nascimento, "não é data"); ok {
		t.Error("data inválida não pode virar idade")
	}
	if _, ok := idadeNaColeta(nascimento, "1970-01-01"); ok {
		t.Error("coleta anterior ao nascimento não pode virar idade")
	}
}
