package services

import (
	"errors"
	"fmt"
	"math"

	"github.com/google/uuid"
	"gorm.io/gorm"

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
// A razão é GRAVADA como resultado, no lote de onde os componentes saíram, e não calculada em
// memória na hora de pontuar. A diferença é a devolutiva: um número que só existe dentro do motor
// pontua sem aparecer em lugar nenhum, sem data, sem histórico e sem régua no deck. O escore do
// paciente andaria 18 pontos e a página que explica o escore não teria como mostrar por quê.
// Gravada no lote, a razão herda a data da coleta, acumula série ao longo dos exames e é desenhada
// como qualquer outro exame.
//
// `source = "derived"` é o que a separa de um valor lançado, e é por ele que a sincronização sabe o
// que pode reescrever. Resultado que uma pessoa digitou nunca é tocado.
const sourceRazaoDerivada = "derived"

type razaoDerivada struct {
	Codigo      string // o exame do catálogo que recebe o resultado
	Nome        string // test_name gravado na linha
	Numerador   string
	Denominador string
	Unidade     string // tem de bater com a unidade do item de escore, senão UnitMatches descarta
}

var razoesDerivadas = []razaoDerivada{
	{Codigo: "PLN09DBBB62", Nome: "Relação Colesterol Total/HDL",
		Numerador: "PLN919303A4", Denominador: "PLN53A449CA", Unidade: "ratio"},
	{Codigo: "PLN73ED669D", Nome: "Relação Triglicerídeos/HDL",
		Numerador: "PLNA0C5545F", Denominador: "PLN53A449CA", Unidade: "ratio"},
	{Codigo: "PLNAPOBA1", Nome: "Relação Apolipoproteína B/A1",
		Numerador: "PLN543993C6", Denominador: "PLNA8451657", Unidade: "ratio"},
}

// componenteDoLote — o valor de um componente dentro de um lote, com o dia da coleta.
type componenteDoLote struct {
	Valor float64
	Unid  string
	Dia   string // AAAA-MM-DD da coleta do lote; é por ele que se procura a razão já lançada
}

// SincronizaRazoesDerivadas calcula as razões do paciente e as grava, lote a lote.
//
// Idempotente: recalcula tudo e só escreve o que mudou. Roda antes de cada snapshot, e é barata
// porque um paciente tem dezenas de lotes, não milhares.
func SincronizaRazoesDerivadas(db *gorm.DB, patientID uuid.UUID) error {
	defs, err := definicoesDasRazoes(db)
	if err != nil || len(defs) == 0 {
		return err
	}
	for _, r := range razoesDerivadas {
		alvo, existe := defs[r.Codigo]
		if !existe {
			continue // catálogo sem esta razão: nada a gravar
		}
		nums, err := componentesPorLote(db, patientID, r.Numerador)
		if err != nil {
			return err
		}
		if len(nums) == 0 {
			continue
		}
		dens, err := componentesPorLote(db, patientID, r.Denominador)
		if err != nil {
			return err
		}
		for lote, num := range nums {
			den, temDen := dens[lote]
			if !temDen {
				continue
			}
			valor, ok := razaoDe(num, den)
			if !ok {
				continue
			}
			if err := gravaRazao(db, patientID, alvo, r, lote, num.Dia, valor); err != nil {
				return err
			}
		}
	}
	return nil
}

// razaoDe divide um componente pelo outro, com as guardas que impedem um número plausível e falso.
//
// A unidade dos dois tem de ser a MESMA, ou o quociente não é adimensional. Todo item comum é
// protegido por `UnitMatches`, e a razão escaparia dessa rede: o que é carimbado como "ratio" é a
// unidade DELA, não a dos componentes. A apolipoproteína B é reportada em g/L por muitos
// laboratórios, e 0,86 g/L sobre 193 mg/dL dá 0,0045, que cai no melhor nível e entrega 18 pontos
// sem que nada acuse; ao contrário, 86 mg/dL sobre 1,93 g/L dá 44,5 e custa os 18. Como a razão é
// adimensional, exigir igualdade também aceita o par inteiro em g/L, que dá o mesmo número.
func razaoDe(num, den componenteDoLote) (float64, bool) {
	if den.Valor == 0 || !mesmaUnidade(num.Unid, den.Unid) {
		return 0, false
	}
	// Duas casas, que é como o laudo imprime razão lipídica e como o paciente lê na régua. Sem
	// isto o deck mostrava "0,5269" ao lado de "3,05", com a precisão sugerindo uma medição que
	// não houve: a razão não é medida, é conta sobre dois números de duas casas.
	return math.Round(num.Valor/den.Valor*100) / 100, true
}

// definicoesDasRazoes traz o id de catálogo de cada razão que existe.
func definicoesDasRazoes(db *gorm.DB) (map[string]uuid.UUID, error) {
	codigos := make([]string, 0, len(razoesDerivadas))
	for _, r := range razoesDerivadas {
		codigos = append(codigos, r.Codigo)
	}
	var defs []models.LabTestDefinition
	if err := db.Where("code IN ? AND deleted_at IS NULL", codigos).Find(&defs).Error; err != nil {
		return nil, err
	}
	out := make(map[string]uuid.UUID, len(defs))
	for i := range defs {
		out[defs[i].Code] = defs[i].ID
	}
	return out, nil
}

// componentesPorLote devolve, por lote do paciente, o valor mais recente daquele exame.
//
// Chaveado por LOTE, e não "o mais recente do paciente", porque é o lote que garante que numerador
// e denominador descrevem a MESMA coleta. Sem isso a razão sairia do colesterol de março sobre o
// HDL de dois anos atrás: um número plausível que nunca existiu no paciente, e que nenhuma rede de
// plausibilidade pegaria.
func componentesPorLote(db *gorm.DB, patientID uuid.UUID, codigo string) (map[uuid.UUID]componenteDoLote, error) {
	var linhas []struct {
		LoteID uuid.UUID
		Valor  float64
		Unid   string
		Dia    string
	}
	err := db.
		Table("lab_results r").
		Select("r.lab_result_batch_id AS lote_id, r.result_numeric AS valor, " +
			"coalesce(r.unit, '') AS unid, to_char(b.collection_date, 'YYYY-MM-DD') AS dia").
		Joins("JOIN lab_result_batches b ON b.id = r.lab_result_batch_id AND b.deleted_at IS NULL").
		Joins("JOIN lab_test_definitions d ON d.id = r.lab_test_definition_id").
		Where("b.patient_id = ? AND d.code = ?", patientID, codigo).
		Where("r.deleted_at IS NULL AND r.result_numeric IS NOT NULL").
		Order("r.lab_result_batch_id, r.created_at DESC").
		Scan(&linhas).Error
	if err != nil {
		return nil, err
	}
	out := make(map[uuid.UUID]componenteDoLote, len(linhas))
	for _, l := range linhas {
		if _, visto := out[l.LoteID]; visto {
			continue // a ordenação já pôs o mais recente do lote na frente
		}
		out[l.LoteID] = componenteDoLote{Valor: l.Valor, Unid: l.Unid, Dia: l.Dia}
	}
	return out, nil
}

// gravaRazao insere ou atualiza a linha da razão naquele lote.
//
// Nunca toca num resultado que não seja `derived`: se o laudo imprimiu a razão e alguém a lançou, o
// que está no prontuário vale mais que a conta.
//
// A busca pelo já lançado é por DIA DE COLETA, e não por lote, porque a mesma coleta se parte em
// lotes: numa paciente do banco de desenvolvimento a razão digitada estava num lote e o colesterol
// e o HDL dela em outro, os dois de 23/07. Procurando só dentro do lote dos componentes, a
// sincronização não via a linha digitada e gravava uma segunda ao lado, deixando o mesmo exame
// duas vezes no mesmo dia com valores levemente diferentes (3,05 e 3,0526).
func gravaRazao(db *gorm.DB, patientID, defID uuid.UUID, r razaoDerivada, lote uuid.UUID, dia string, valor float64) error {
	var lancada int64
	if err := db.Table("lab_results r").
		Joins("JOIN lab_result_batches b ON b.id = r.lab_result_batch_id AND b.deleted_at IS NULL").
		Where("b.patient_id = ? AND r.lab_test_definition_id = ?", patientID, defID).
		Where("to_char(b.collection_date, 'YYYY-MM-DD') = ?", dia).
		Where("r.deleted_at IS NULL AND r.source <> ?", sourceRazaoDerivada).
		Count(&lancada).Error; err != nil {
		return err
	}
	if lancada > 0 {
		return nil // alguém lançou a razão nesta coleta: o prontuário vence a conta
	}

	var existente models.LabResult
	err := db.Where("lab_result_batch_id = ? AND lab_test_definition_id = ? AND deleted_at IS NULL",
		lote, defID).First(&existente).Error

	switch {
	case err == nil:
		if existente.Source != sourceRazaoDerivada {
			return nil // lançado por gente: não se mexe
		}
		if existente.ResultNumeric != nil && quaseIgual(*existente.ResultNumeric, valor) {
			return nil
		}
		return db.Model(&existente).Updates(map[string]any{
			"result_numeric": valor,
			"unit":           r.Unidade,
		}).Error

	case errors.Is(err, gorm.ErrRecordNotFound):
		unidade := r.Unidade
		nota := fmt.Sprintf("Calculada pelo escore a partir de %s e %s da mesma coleta.",
			r.Numerador, r.Denominador)
		return db.Create(&models.LabResult{
			LabResultBatchID:    lote,
			LabTestDefinitionID: &defID,
			TestName:            r.Nome,
			TestType:            "biochemistry",
			ResultNumeric:       &valor,
			Unit:                &unidade,
			UnitOriginal:        &unidade,
			Matched:             true,
			Source:              sourceRazaoDerivada,
			Interpretation:      &nota,
		}).Error

	default:
		return err
	}
}

// quaseIgual evita reescrever a linha a cada snapshot por causa do último bit do float.
func quaseIgual(a, b float64) bool {
	d := a - b
	return d > -1e-9 && d < 1e-9
}
