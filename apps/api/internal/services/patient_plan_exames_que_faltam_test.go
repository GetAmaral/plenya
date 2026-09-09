package services

import (
	"strings"
	"testing"

	"github.com/plenya/api/internal/dto"
)

// O slide "os exames que ainda não voltaram" errava de duas maneiras ao mesmo tempo, e as duas
// saíam impressas para o paciente: listava o pedido inteiro como pendente mesmo depois de o laudo
// chegar, e colava a justificativa de um exame descartado na última linha que tinha sobrado.
//
// O pedido usado aqui é o formato real: linhas em branco separando blocos, cabeçalho de grupo em
// caixa alta e justificativa em "#" depois do exame que ela explica.
const pedidoDeExemplo = `Ácido fólico eritrocitário
ACTH
Alumínio
Apolipoproteína A1
Apolipoproteína B
Bilirrubinas totais e frações
Cobre
DHEA-S
Ferro
Selênio

Radiografia de tórax PA + perfil

TC coração para escore de cálcio coronariano
# Refina a estratificação de risco em quem está na faixa intermediária.`

func dossieComPedido(voltaram []string) *dto.PlanDossierResponse {
	return &dto.PlanDossierResponse{
		LabRequest: &dto.PlanDossierLabRequest{
			ID: "req", Date: "2026-07-09", Exams: pedidoDeExemplo, Returned: voltaram,
		},
	}
}

func linhasDoSlide(t *testing.T, d *dto.PlanDossierResponse) [][]string {
	t.Helper()
	slides := montaExamesQueFaltam(d)
	if len(slides) != 1 {
		t.Fatalf("esperava 1 slide, veio %d", len(slides))
	}
	var out [][]string
	for _, r := range slides[0].Table.Rows {
		out = append(out, r.Cells)
	}
	return out
}

// Sem nada devolvido o slide lista o pedido, e o cabeçalho de grupo não vira exame.
func TestExamesQueFaltam_SemRetornoListaOPedido(t *testing.T) {
	linhas := linhasDoSlide(t, dossieComPedido(nil))
	// Oito cabem; a nona linha é o aviso de corte.
	if len(linhas) != 9 {
		t.Fatalf("esperava 8 exames + aviso de corte, veio %d linhas", len(linhas))
	}
	if got := linhas[0][0]; got != "Ácido fólico eritrocitário" {
		t.Errorf("primeira linha: %q", got)
	}
	if aviso := linhas[8][0]; !strings.Contains(aviso, "ainda não voltaram") {
		t.Errorf("o aviso de corte tem de contar o que a página conta, veio %q", aviso)
	}
}

// O que já voltou sai da lista. É o defeito que fazia o slide dizer que faltavam os 38 exames de
// uma paciente que já tinha 33 lançados.
func TestExamesQueFaltam_DescontaOQueJaVoltou(t *testing.T) {
	voltaram := []string{
		"Ácido fólico eritrocitário", "ACTH", "Alumínio", "Apolipoproteína A1",
		"Apolipoproteína B", "Bilirrubinas totais e frações", "Cobre", "Ferro", "Selênio",
	}
	linhas := linhasDoSlide(t, dossieComPedido(voltaram))
	var nomes []string
	for _, c := range linhas {
		nomes = append(nomes, c[0])
	}
	esperado := []string{"DHEA-S", "Radiografia de tórax PA + perfil", "TC coração para escore de cálcio coronariano"}
	if len(nomes) != len(esperado) {
		t.Fatalf("esperava %v, veio %v", esperado, nomes)
	}
	for i := range esperado {
		if nomes[i] != esperado[i] {
			t.Errorf("linha %d: esperava %q, veio %q", i, esperado[i], nomes[i])
		}
	}
}

// A justificativa adere ao exame que ela explica, e não à última linha que sobrou. Aqui o "#" vem
// depois do escore de cálcio, que continua na lista: ele fica com ela, e o DHEA-S não.
func TestExamesQueFaltam_JustificativaAdereAoExameCerto(t *testing.T) {
	voltaram := []string{
		"Ácido fólico eritrocitário", "ACTH", "Alumínio", "Apolipoproteína A1",
		"Apolipoproteína B", "Bilirrubinas totais e frações", "Cobre", "Ferro", "Selênio",
	}
	linhas := linhasDoSlide(t, dossieComPedido(voltaram))
	for _, c := range linhas {
		temPorque := len(c) > 1 && c[1] != ""
		if c[0] == "TC coração para escore de cálcio coronariano" && !temPorque {
			t.Error("o escore de cálcio perdeu a justificativa que era dele")
		}
		if c[0] != "TC coração para escore de cálcio coronariano" && temPorque {
			t.Errorf("%q ficou com uma justificativa que não é dele: %q", c[0], c[1])
		}
	}
}

// O caso que saiu impresso: com o corte de oito linhas, toda justificativa posterior ao oitavo
// exame aderia à oitava linha. No pedido real desta paciente isso fez o slide dizer que o cobre
// serve para quantificar escore de cálcio coronariano.
func TestExamesQueFaltam_JustificativaDeExameCortadoNaoGrudaNaOitava(t *testing.T) {
	linhas := linhasDoSlide(t, dossieComPedido(nil))
	for _, c := range linhas {
		if len(c) > 1 && c[1] != "" {
			t.Errorf("%q ficou com a justificativa de um exame que nem está na página: %q", c[0], c[1])
		}
	}
}
