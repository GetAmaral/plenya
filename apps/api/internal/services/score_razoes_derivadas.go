package services

import (
	"github.com/google/uuid"

	"github.com/plenya/api/internal/models"
)

// RAZÕES DO ESCORE: itens que existem no catálogo e que nada alimentava.
//
// `Relação Colesterol Total/HDL` e `Relação Triglicerídeos/HDL` estão no escore desde sempre, com
// 18 e 14 pontos. Medido em produção antes desta mudança: a primeira tinha sido avaliada 3 vezes
// contra 16 `no_data_available`, e a segunda 2 contra 17. O motivo não é dado faltando, é que razão
// não vem em laudo: o laboratório reporta colesterol total e HDL, não o quociente. Só pontuava
// quando alguém, por acaso, digitava a razão impressa no rodapé de algum laudo.
//
// Então o cálculo é do motor. Ele acontece aqui, sobre o mapa que o snapshot já carregou, e não
// como INSERT em `lab_results`, por duas razões: a razão é conta e não medida, e um resultado
// gravado ficaria congelado enquanto os componentes dele mudam a cada coleta.
//
// A guarda que importa é o LOTE. Os componentes vêm do mapa "mais recente por código", que mistura
// datas de propósito: o colesterol pode ser de março e o HDL de 2024. Dividir um pelo outro
// produziria um número que nunca existiu em paciente nenhum, dentro da faixa plausível e portanto
// invisível. Exigir que os dois venham do MESMO lote é o que garante que o quociente descreve uma
// coleta real. Custa falso negativo quando o laboratório separa o lipidograma em dois lotes do
// mesmo dia, e esse é o lado certo para errar.
type razaoDerivada struct {
	Codigo      string // o código do item de escore que recebe o resultado
	Numerador   string
	Denominador string
	Unidade     string // tem de bater com a unidade do item, senão UnitMatches descarta
}

var razoesDerivadas = []razaoDerivada{
	// Colesterol Total / HDL
	{Codigo: "PLN09DBBB62", Numerador: "PLN919303A4", Denominador: "PLN53A449CA", Unidade: "ratio"},
	// Triglicerídeos / HDL
	{Codigo: "PLN73ED669D", Numerador: "PLNA0C5545F", Denominador: "PLN53A449CA", Unidade: "ratio"},
	// Apolipoproteína B / Apolipoproteína A1
	{Codigo: "PLNAPOBA1", Numerador: "PLN543993C6", Denominador: "PLNA8451657", Unidade: "ratio"},
}

// aplicaRazoesDerivadas acrescenta ao mapa as razões que dá para calcular.
//
// Nunca sobrescreve: se a razão foi lançada como resultado (um laudo que a imprime, uma carga
// manual), o que está no prontuário vale mais que a conta.
func aplicaRazoesDerivadas(porCodigo map[string]models.LabResult) {
	for _, r := range razoesDerivadas {
		if _, existe := porCodigo[r.Codigo]; existe {
			continue
		}
		num, okN := porCodigo[r.Numerador]
		den, okD := porCodigo[r.Denominador]
		if !okN || !okD {
			continue
		}
		if num.ResultNumeric == nil || den.ResultNumeric == nil || *den.ResultNumeric == 0 {
			continue
		}
		// Mesma unidade nos dois, ou o quociente não é adimensional.
		//
		// Todo item comum é protegido por `UnitMatches`, e a razão escaparia dessa rede: o que é
		// carimbado como "ratio" é a unidade DELA, não a dos componentes. A apolipoproteína B é
		// reportada em g/L por muitos laboratórios, e 0,86 g/L dividido por 193 mg/dL dá 0,0045,
		// que cai no melhor nível e entrega 18 pontos sem que nada acuse. Ao contrário, 86 mg/dL
		// sobre 1,93 g/L dá 44,5 e custa os 18. Como a razão é adimensional, exigir unidades
		// iguais também aceita o par inteiro em g/L, que dá exatamente o mesmo número.
		if !mesmaUnidadeDeResultado(num.Unit, den.Unit) {
			continue
		}
		// Mesma coleta, ou o quociente descreve um paciente que não existe.
		if num.LabResultBatchID != den.LabResultBatchID {
			continue
		}
		valor := *num.ResultNumeric / *den.ResultNumeric
		unidade := r.Unidade
		porCodigo[r.Codigo] = models.LabResult{
			// ID zerado de propósito. `patient_score_item_results.lab_result_id` tem chave
			// estrangeira para `lab_results`, e este resultado não existe lá: quem grava tem de
			// olhar `IsZero()` antes de referenciá-lo. É o que `resultadoDerivado` responde.
			ID:               uuid.Nil,
			LabResultBatchID: num.LabResultBatchID,
			TestName:         "Razão calculada pelo escore",
			ResultNumeric:    &valor,
			Unit:             &unidade,
		}
	}
}

// mesmaUnidadeDeResultado é `mesmaUnidade` para os ponteiros que `LabResult.Unit` usa. Aproveita a
// normalização de lá, que já resolve "µg" e "mcg" como a mesma coisa. Unidade ausente conta como
// diferente: sem saber a grandeza dos dois lados, não há divisão defensável.
func mesmaUnidadeDeResultado(a, b *string) bool {
	return a != nil && b != nil && mesmaUnidade(*a, *b)
}

// resultadoDerivado diz se um resultado veio da conta acima, e não da tabela.
//
// Existe porque o snapshot guarda `lab_result_id` com chave estrangeira: referenciar um resultado
// que só existe em memória gravaria um UUID zerado e a inserção falharia, ou pior, apontaria para
// nada se a constraint mudasse.
func resultadoDerivado(r models.LabResult) bool { return r.ID == uuid.Nil }
