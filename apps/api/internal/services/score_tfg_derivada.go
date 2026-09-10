package services

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// TAXA DE FILTRAÇÃO GLOMERULAR: o número renal que decide conduta, e que nada calculava.
//
// O item existe no catálogo (`ETFG`) e nunca teve item de escore, porque eTFG não é medida: é conta
// sobre creatinina e/ou cistatina C. Quando vinha, vinha porque alguém digitou o número impresso no
// rodapé do laudo — e o número do laudo é SEMPRE o de creatinina isolada, mesmo quando a cistatina
// C foi dosada na mesma coleta.
//
// A diferença entre as duas não é acadêmica. Num paciente que treina força e suplementa creatina, a
// creatinina sobe sem que a filtração tenha mudado: 1,50 mg/dL dá eTFG 59 (estágio G3a) enquanto a
// cistatina C de 1,01 mg/L da mesma janela dá 82, e a equação combinada dá 72 (G2). Estágio G3a e
// G2 mudam meta de pressão, mudam dose de droga e mudam a conversa com o paciente.
//
// Por isso a preferência é combinada > cistatina > creatinina: a KDIGO recomenda a combinada
// justamente quando a creatinina é suspeita, e ela é a mais exata das três sempre que os dois
// insumos existem NA MESMA COLETA.
//
// Como as razões lipídicas, a eTFG é GRAVADA como resultado no lote de onde os insumos saíram, e
// não calculada em memória na hora de pontuar: assim ela herda a data, acumula série e desenha
// régua na devolutiva. Ver a mesma decisão, com mais detalhe, em score_razoes_derivadas.go.
const (
	codigoTFG  = "ETFG"
	nomeTFG    = "Taxa de Filtração Glomerular Estimada"
	unidadeTFG = "mL/min/1.73m²"

	codigoCreatinina = "PLN357DC859"
	codigoCistatinaC = "PLN42D34D96"
)

// equacaoTFG diz qual das três produziu o número, e vai para a interpretação da linha. Sem isso a
// régua mostraria 59 e 82 como se fossem a mesma medida piorando, quando são equações diferentes.
type equacaoTFG string

const (
	tfgCombinada  equacaoTFG = "creatinina + cistatina C (CKD-EPI 2021)"
	tfgCistatinaC equacaoTFG = "cistatina C (CKD-EPI 2012)"
	tfgCreatinina equacaoTFG = "creatinina (CKD-EPI 2021)"
)

// SincronizaTFGDerivada calcula a eTFG do paciente e a grava, lote a lote.
//
// Idempotente pelo mesmo desenho das razões: recalcula tudo e só escreve o que mudou.
func SincronizaTFGDerivada(db *gorm.DB, patientID uuid.UUID) error {
	var def models.LabTestDefinition
	err := db.Where("code = ? AND deleted_at IS NULL", codigoTFG).First(&def).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil // catálogo sem eTFG: nada a gravar
	}
	if err != nil {
		return err
	}

	var paciente models.Patient
	if err := db.First(&paciente, patientID).Error; err != nil {
		return err
	}
	// Sem data de nascimento não há idade, e sem idade nenhuma das três equações existe. O cadastro
	// rápido da recepção grava data zerada (ver CreatePatientRequest), e isso não é erro: é só
	// motivo para não calcular.
	if paciente.BirthDate.IsZero() {
		// Retira o que já houver: se a data de nascimento foi limpa depois, a eTFG antiga não pode
		// continuar valendo 22 pontos sobre uma conta que o código agora se recusa a fazer.
		return retiraTFGSemInsumo(db, patientID, def.ID, nil)
	}

	creatininas, err := componentesPorLote(db, patientID, codigoCreatinina)
	if err != nil {
		return err
	}
	cistatinas, err := componentesPorLote(db, patientID, codigoCistatinaC)
	if err != nil {
		return err
	}

	// Sexo indefinido não vira conta. As três equações têm coeficiente de sexo, e a diferença não é
	// de arredondamento: a mesma creatinina de 0,9 mg/dL aos 43 anos dá ~109 em homem e ~80 em
	// mulher, que é G1 contra G2. O cadastro rápido da recepção grava `other` quando o campo é
	// omitido (patient_service.go), e é caminho normal, não erro — mas chutar masculino ali
	// escreveria um estágio renal inventado no prontuário, com 22 pontos de escore atrás dele.
	if paciente.Gender != models.GenderMale && paciente.Gender != models.GenderFemale {
		return retiraTFGSemInsumo(db, patientID, def.ID, nil)
	}
	feminino := paciente.Gender == models.GenderFemale

	// A união dos lotes: um lote pode ter só um dos dois insumos e ainda assim render eTFG.
	lotes := make([]uuid.UUID, 0, len(creatininas)+len(cistatinas))
	vistos := make(map[uuid.UUID]struct{}, len(creatininas)+len(cistatinas))
	for l := range creatininas {
		lotes, vistos[l] = append(lotes, l), struct{}{}
	}
	for l := range cistatinas {
		if _, ja := vistos[l]; !ja {
			lotes = append(lotes, l)
		}
	}
	// Ordem fixa. A escolha de qual lote fica com a eTFG do dia não pode depender da ordem de
	// iteração do mapa do Go: com uma coleta partida em dois lotes — um só com creatinina e outro
	// com creatinina e cistatina — quem chegasse primeiro ficava dono da linha, e a guarda do DIA
	// impedia o outro de escrever. A preferência declarada (combinada > cistatina > creatinina)
	// virava sorteio, e mudava de execução para execução numa base recém-criada.
	sort.Slice(lotes, func(i, j int) bool { return lotes[i].String() < lotes[j].String() })

	// Melhor candidato POR DIA, e não por lote, porque a eTFG é uma por coleta.
	melhorDoDia := make(map[string]candidatoTFG, len(lotes))
	for _, lote := range lotes {
		cr, temCr := creatininas[lote]
		cys, temCys := cistatinas[lote]

		// Unidade errada não vira conta. A creatinina sai em mg/dL e a cistatina C em mg/L; um
		// laudo em µmol/L entraria na fórmula como se fosse mg/dL e devolveria uma filtração
		// plausível e falsa, que nenhuma rede de plausibilidade pegaria depois.
		if temCr && !mesmaUnidade(cr.Unid, "mg/dL") {
			temCr = false
		}
		if temCys && !mesmaUnidade(cys.Unid, "mg/L") {
			temCys = false
		}

		dia := cr.Dia
		if !temCr {
			dia = cys.Dia
		}
		idade, ok := idadeNaColeta(paciente.BirthDate, dia)
		if !ok || idade < 18 {
			// Abaixo de 18 as equações do adulto não valem (a criança tem a Schwartz), e é o que o
			// próprio laudo avisa. Melhor não ter eTFG do que ter a errada.
			continue
		}

		var c candidatoTFG
		switch {
		case temCr && temCys:
			c = candidatoTFG{lote, tfgCombinadaCKDEPI2021(cr.Valor, cys.Valor, idade, feminino), tfgCombinada, 3}
		case temCys:
			c = candidatoTFG{lote, tfgPorCistatinaCKDEPI2012(cys.Valor, idade, feminino), tfgCistatinaC, 2}
		case temCr:
			c = candidatoTFG{lote, tfgPorCreatininaCKDEPI2021(cr.Valor, idade, feminino), tfgCreatinina, 1}
		default:
			continue
		}
		if c.valor <= 0 || math.IsNaN(c.valor) || math.IsInf(c.valor, 0) {
			continue
		}
		// Inteiro, que é como laboratório e diretriz reportam filtração. Uma casa decimal sugeriria
		// uma precisão que a equação não tem: o intervalo de confiança dela é de dezenas de pontos.
		c.valor = math.Round(c.valor)

		if atual, ja := melhorDoDia[dia]; !ja || c.prioridade > atual.prioridade {
			melhorDoDia[dia] = c
		}
	}

	for dia, c := range melhorDoDia {
		if err := gravaTFG(db, patientID, def.ID, c.lote, dia, c.valor, c.equacao); err != nil {
			return err
		}
	}
	return retiraTFGSemInsumo(db, patientID, def.ID, melhorDoDia)
}

// candidatoTFG é a eTFG que um lote consegue produzir, com a prioridade da equação usada:
// 3 combinada, 2 cistatina C, 1 creatinina. Entre lotes do MESMO dia vence a maior prioridade.
type candidatoTFG struct {
	lote       uuid.UUID
	valor      float64
	equacao    equacaoTFG
	prioridade int
}

// retiraTFGSemInsumo apaga a eTFG calculada de um dia que deixou de produzir valor.
//
// Sem isto a linha derivada nunca é retirada: se a creatinina que a gerou for apagada, ou tiver a
// unidade corrigida para µmol/L (e aí `mesmaUnidade` passa a recusá-la), o dia some do cálculo e a
// eTFG antiga continua no prontuário, pontuando 22 pontos sobre um número que nada mais sustenta.
// Só toca no que o motor escreveu: resultado lançado por gente nunca é retirado.
func retiraTFGSemInsumo(db *gorm.DB, patientID, defID uuid.UUID, melhorDoDia map[string]candidatoTFG) error {
	var derivadas []struct {
		ID  uuid.UUID
		Dia string
	}
	if err := db.Table("lab_results r").
		Select("r.id AS id, to_char(b.collection_date, 'YYYY-MM-DD') AS dia").
		Joins("JOIN lab_result_batches b ON b.id = r.lab_result_batch_id AND b.deleted_at IS NULL").
		Where("b.patient_id = ? AND r.lab_test_definition_id = ?", patientID, defID).
		Where("r.source = ? AND r.deleted_at IS NULL", sourceRazaoDerivada).
		Scan(&derivadas).Error; err != nil {
		return err
	}
	for _, d := range derivadas {
		if _, aindaVale := melhorDoDia[d.Dia]; aindaVale {
			continue
		}
		if err := db.Model(&models.LabResult{}).Where("id = ?", d.ID).
			Update("deleted_at", time.Now()).Error; err != nil {
			return err
		}
	}
	return nil
}

// idadeNaColeta devolve a idade em anos na data da coleta, e não a idade de hoje. Um lote de 2019
// pontuado com a idade de 2026 usaria sete anos a mais de decaimento e devolveria uma filtração
// menor do que a que o paciente tinha no dia.
func idadeNaColeta(nascimento time.Time, dia string) (float64, bool) {
	coleta, err := time.Parse("2006-01-02", dia)
	if err != nil {
		return 0, false
	}
	anos := coleta.Sub(nascimento).Hours() / 24 / 365.2425
	if anos <= 0 || anos > 120 {
		return 0, false
	}
	return anos, true
}

// As três equações do CKD-EPI, sem termo de raça, como publicadas.
//
// Fontes: Inker LA et al., N Engl J Med 2021;385:1737-49 (creatinina 2021 e combinada 2021) e
// Inker LA et al., N Engl J Med 2012;367:20-29 (cistatina C). Coeficientes conferidos contra o
// NIDDK e a National Kidney Foundation, e validados contra os números impressos pelo laboratório:
// creatinina 1,50 aos 43 anos devolve 59, e cistatina C 1,01 na mesma idade devolve 82, que é
// exatamente o que o laudo trouxe.

func tfgPorCreatininaCKDEPI2021(scr, idade float64, feminino bool) float64 {
	kappa, alfa, sexo := 0.9, -0.302, 1.0
	if feminino {
		kappa, alfa, sexo = 0.7, -0.241, 1.012
	}
	r := scr / kappa
	return 142 * math.Pow(math.Min(r, 1), alfa) * math.Pow(math.Max(r, 1), -1.200) *
		math.Pow(0.9938, idade) * sexo
}

func tfgPorCistatinaCKDEPI2012(scys, idade float64, feminino bool) float64 {
	sexo := 1.0
	if feminino {
		sexo = 0.932
	}
	r := scys / 0.8
	return 133 * math.Pow(math.Min(r, 1), -0.499) * math.Pow(math.Max(r, 1), -1.328) *
		math.Pow(0.996, idade) * sexo
}

func tfgCombinadaCKDEPI2021(scr, scys, idade float64, feminino bool) float64 {
	kappa, alfa, sexo := 0.9, -0.144, 1.0
	if feminino {
		kappa, alfa, sexo = 0.7, -0.219, 0.963
	}
	rc, ry := scr/kappa, scys/0.8
	return 135 *
		math.Pow(math.Min(rc, 1), alfa) * math.Pow(math.Max(rc, 1), -0.544) *
		math.Pow(math.Min(ry, 1), -0.323) * math.Pow(math.Max(ry, 1), -0.778) *
		math.Pow(0.9961, idade) * sexo
}

// gravaTFG grava a eTFG do DIA, e não do lote.
//
// A busca do que já existe é pelo dia de coleta inteiro, porque a mesma coleta se parte em lotes:
// procurando só dentro do lote escolhido, a sincronização não via a linha que estava no lote irmão,
// contava "já tem uma no dia" e desistia — a eTFG nascia com o primeiro insumo que chegasse e nunca
// mais era atualizada, mesmo quando a cistatina C aparecia depois e permitia a equação melhor.
//
// QUANDO A CONTA SUPERA O QUE ESTÁ NO PRONTUÁRIO
//
// A regra geral do projeto é que valor lançado por gente vence a conta, e ela vale para medida. A
// eTFG não é medida: a que vem impressa no laudo é uma CONTA do laboratório, e sempre a de
// creatinina isolada. Tratá-la como intocável fazia a feature ser inerte justamente no caso que a
// motiva — coleta com creatinina, cistatina C e a eTFG impressa ao lado, em que a combinada é a
// certa e nunca era gravada.
//
// Então a linha impressa é superada SÓ quando a nossa equação usa cistatina C, que é informação que
// o laboratório comprovadamente não usou. Se o melhor que temos é creatinina isolada, é a mesma
// conta que ele fez, e não se mexe. O valor original nunca some: fica escrito na interpretação da
// linha, com a data em que foi substituído.
func gravaTFG(db *gorm.DB, patientID, defID, lote uuid.UUID, dia string, valor float64, eq equacaoTFG) error {
	nota := fmt.Sprintf("Calculada pelo escore por %s, a partir dos insumos da mesma coleta.", eq)
	usaCistatina := eq == tfgCombinada || eq == tfgCistatinaC

	var existente models.LabResult
	err := db.Table("lab_results r").
		Select("r.*").
		Joins("JOIN lab_result_batches b ON b.id = r.lab_result_batch_id AND b.deleted_at IS NULL").
		Where("b.patient_id = ? AND r.lab_test_definition_id = ?", patientID, defID).
		Where("to_char(b.collection_date, 'YYYY-MM-DD') = ? AND r.deleted_at IS NULL", dia).
		Order("r.created_at ASC").
		First(&existente).Error

	switch {
	case err == nil:
		if existente.Source != sourceRazaoDerivada {
			if !usaCistatina {
				return nil // mesma conta que o laboratório fez: não se mexe
			}
			original := "sem valor"
			if existente.ResultNumeric != nil {
				original = formatNumberPT(*existente.ResultNumeric)
			}
			nota = fmt.Sprintf("%s Substitui a eTFG impressa no laudo (%s), que é calculada só com "+
				"creatinina.", nota, original)
		}
		if existente.ResultNumeric != nil && quaseIgual(*existente.ResultNumeric, valor) &&
			existente.Interpretation != nil && *existente.Interpretation == nota {
			return nil
		}
		unidade := unidadeTFG
		return db.Model(&models.LabResult{}).Where("id = ?", existente.ID).Updates(map[string]any{
			"result_numeric": valor,
			"unit":           unidade,
			"unit_original":  unidade,
			"source":         sourceRazaoDerivada,
			"interpretation": nota,
		}).Error

	case errors.Is(err, gorm.ErrRecordNotFound):
		unidade := unidadeTFG
		return db.Create(&models.LabResult{
			LabResultBatchID:    lote,
			LabTestDefinitionID: &defID,
			TestName:            nomeTFG,
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
