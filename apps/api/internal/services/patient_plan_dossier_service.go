package services

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/repository"
)

// Dossiê do plano de paciente: deriva do prontuário tudo que a devolutiva (o "deck" de resultados)
// consegue montar sozinha, para que preparar o plano do próximo paciente não seja recomeçar do zero.
//
// Até aqui, o `reguas.json` de cada paciente era montado à mão a partir do banco. Este serviço
// reproduz aquele arquivo a partir da fonte, e o formato de saída é o MESMO de propósito: os decks
// já feitos consomem sem adaptação.
//
// A divisão de trabalho é deliberada: aqui saem FATOS e CANDIDATOS ORDENADOS (a escala de cada
// exame, o histórico do paciente, o que está bem, o que está se movendo, e quanto cada achado pesa).
// O que é julgamento clínico — a leitura dos achados, o arco narrativo, os títulos em voz de
// paciente e as condutas — continua sendo escrito por quem atende.

// Padding do eixo da régua, em fração do vão entre a primeira e a última fronteira de nível. Dá
// respiro para o segmento aberto das pontas ("≤15", ">300") ter largura desenhável.
const rulerAxisPad = 0.22

// Folga extra quando o paciente tem valor FORA do vão das fronteiras: o eixo se estica para que o
// ponto dele caiba com margem, em vez de ficar colado na borda.
const (
	rulerHistoryPadLow  = 0.96
	rulerHistoryPadHigh = 1.04
)

// PatientPlanDossierService monta o dossiê. Só lê.
type PatientPlanDossierService struct {
	db           *gorm.DB
	snapshotRepo *repository.ScoreSnapshotRepository
	carePlan     *CarePlanService
}

func NewPatientPlanDossierService(
	db *gorm.DB,
	snapshotRepo *repository.ScoreSnapshotRepository,
	carePlan *CarePlanService,
) *PatientPlanDossierService {
	return &PatientPlanDossierService{db: db, snapshotRepo: snapshotRepo, carePlan: carePlan}
}

// labRow — uma linha crua de resultado, já com a data efetiva resolvida.
type labRow struct {
	Code      string
	TestName  string
	Numeric   float64
	Text      string
	Reference string
	// Day — o dia da coleta como AAAA-MM-DD. Ver collectionDay() para o porquê de não ser
	// simplesmente Collected.Format().
	Day        string
	Collected  time.Time
	ResultText string
	// DefName e DefGloss vêm do CATÁLOGO de exames, não do score: o nome que o paciente
	// reconhece e a glosa do que o exame mede.
	DefName  string
	DefGloss string
	// DefUnit — unidade do exame no catálogo. Serve para conferir se a escala do item de escore
	// fala da mesma grandeza.
	DefUnit string
}

// collectionDay resolve o dia-calendário da coleta, e é mais chato do que parece porque a coluna
// `timestamptz` guarda DUAS coisas diferentes:
//
//   - um dia-calendário digitado, gravado como meia-noite UTC — "2024-11-05 00:00:00Z", que em
//     America/Sao_Paulo é "2024-11-04 21:00". Formatar no fuso local volta um dia.
//   - um instante real (coleta ou importação) — "2026-02-06 23:12:29.9-03". Formatar em UTC
//     avança um dia, porque em UTC já é 02:12 do dia seguinte.
//
// As duas formas convivem no banco (em `lab_result_batches`, 8 de uma e 14 de outra; em
// `lab_results.collection_date`, 258 de 258 são instantes reais). Escolher um fuso fixo erra uma
// das duas, e o erro não é cosmético: o dia é a chave de deduplicação do histórico, então uma
// coleta só vira dois pontos na régua, com um dia de distância.
//
// A distinção é confiável: só um dia-calendário cai exatamente em 00:00:00 UTC.
func collectionDay(t time.Time) string {
	utc := t.UTC()
	if utc.Hour() == 0 && utc.Minute() == 0 && utc.Second() == 0 && utc.Nanosecond() == 0 {
		return utc.Format("2006-01-02")
	}
	return t.In(saoPaulo()).Format("2006-01-02")
}

// Build monta o dossiê do paciente.
func (s *PatientPlanDossierService) Build(patientID uuid.UUID) (*dto.PlanDossierResponse, error) {
	var patient models.Patient
	if err := s.db.First(&patient, patientID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPatientNotFound
		}
		return nil, err
	}

	rows, err := s.loadLabRows(patientID)
	if err != nil {
		return nil, err
	}

	rulers, err := s.buildRulers(&patient, rows)
	if err != nil {
		return nil, err
	}

	out := &dto.PlanDossierResponse{
		Patient: dto.PlanDossierPatient{
			ID:     patient.ID.String(),
			Name:   patient.Name,
			Gender: string(patient.Gender),
			Age:    patient.Age,
		},
		Rulers:      rulers,
		Vitals:      []dto.PlanDossierVitals{},
		CarePlan:    []dto.CarePlanItemResponse{},
		Medications: []dto.PlanDossierPrescription{},
		GeneratedAt: time.Now().In(saoPaulo()).Format(time.RFC3339),
	}

	// Escore vigente: quando existe, é ele quem diz quantos pontos o paciente deixou na mesa.
	//
	// A chave é o ItemID, NÃO o código do exame: um código pode ter várias variantes que se aplicam
	// ao mesmo paciente ao mesmo tempo (o IGF-1 tem um item guarda-chuva sem recorte etário mais
	// quatro por faixa de idade), e o motor avalia todas. Chaveando por código, a última do laço
	// sobrescrevia as outras e os pontos podiam acabar vindo de uma variante diferente da que
	// desenhou a régua — trocando a ORDEM dos achados que o médico vê, que é ordenada por eles.
	lostBySnapshot := map[uuid.UUID]float64{}
	if snap, sErr := s.snapshotRepo.GetLatestByPatientID(patientID); sErr == nil && snap != nil {
		out.Snapshot = &dto.PlanDossierSnapshot{
			ID:              snap.ID.String(),
			CalculatedAt:    snap.CalculatedAt.In(saoPaulo()).Format(time.RFC3339),
			TotalPercentage: snap.TotalScorePercentage,
		}
		for i := range snap.ItemResults {
			r := snap.ItemResults[i]
			if r.Status != models.EvaluationStatusEvaluated {
				continue
			}
			lostBySnapshot[r.ItemID] = math.Max(r.MaxPoints-r.ActualPoints, 0)
		}
	} else if sErr != nil && !errors.Is(sErr, repository.ErrNoSnapshots) {
		return nil, sErr
	}

	// O plano não se monta só com exame. A anamnese responde metade do que vira slide (sono,
	// tabagismo, histórico familiar, IMC), e a medida de consultório responde o resto (pressão,
	// cintura). As três fontes entram na MESMA lista de achados, ordenadas pelo mesmo critério:
	// o que mais pesa aparece primeiro, venha de onde vier.
	out.Strong, out.Moving = classifyFindings(rulers, lostBySnapshot)

	anamRows, aErr := s.loadAnamnesisRows(patientID)
	if aErr != nil {
		return nil, aErr
	}
	for _, f := range anamnesisFindings(anamRows, lostBySnapshot) {
		if f.Kind == dto.PlanFindingStrong {
			out.Strong = append(out.Strong, f)
		} else {
			out.Moving = append(out.Moving, f)
		}
	}
	sortStrong(out.Strong)
	sortMoving(out.Moving)

	vitals, vErr := s.loadVitals(patientID)
	if vErr != nil {
		return nil, vErr
	}
	out.Vitals = vitals

	if s.carePlan != nil {
		items, cErr := s.carePlan.ListByPatient(patientID, false)
		if cErr != nil {
			return nil, cErr
		}
		out.CarePlan = items
	}

	if lr, lErr := s.loadLastLabRequest(patientID); lErr != nil {
		return nil, lErr
	} else if lr != nil {
		out.LabRequest = lr
	}

	presc, pErr := s.loadPrescriptions(patientID)
	if pErr != nil {
		return nil, pErr
	}
	out.Medications = presc

	return out, nil
}

// loadLabRows traz os resultados numéricos do paciente já ligados ao código do catálogo. A data
// efetiva é a do exame quando o laudo traz uma ("Coletado em:" por exame) e a do lote quando não.
//
// NÃO filtra por status do lote, de propósito: o dossiê tem que enxergar o que já está no
// prontuário mesmo que ninguém tenha fechado a revisão. Não é detalhe — no dev os 22 lotes estão
// em `pending`, e um filtro por `completed` devolveria um dossiê vazio.
func (s *PatientPlanDossierService) loadLabRows(patientID uuid.UUID) ([]labRow, error) {
	var raw []struct {
		Code       string
		TestName   string
		Numeric    float64
		ResultText *string
		Reference  *string
		DefUnit    *string
		DefName    *string
		DefGloss   *string
		Collected  time.Time
	}
	err := s.db.
		Table("lab_results AS lr").
		Select(`ltd.code                                        AS code,
		        lr.test_name                                    AS test_name,
		        lr.result_numeric                               AS numeric,
		        lr.result_text                                  AS result_text,
		        lr.reference_range                              AS reference,
		        ltd.unit                                        AS def_unit,
		        ltd.name                                        AS def_name,
		        ltd.patient_gloss                               AS def_gloss,
		        COALESCE(lr.collection_date, b.collection_date) AS collected`).
		Joins("JOIN lab_result_batches b ON b.id = lr.lab_result_batch_id AND b.deleted_at IS NULL").
		Joins("JOIN lab_test_definitions ltd ON ltd.id = lr.lab_test_definition_id AND ltd.deleted_at IS NULL").
		Where("b.patient_id = ? AND lr.deleted_at IS NULL AND lr.result_numeric IS NOT NULL", patientID).
		Order("collected ASC").
		Scan(&raw).Error
	if err != nil {
		return nil, err
	}

	out := make([]labRow, 0, len(raw))
	for _, r := range raw {
		row := labRow{
			Code:      r.Code,
			TestName:  r.TestName,
			Numeric:   r.Numeric,
			Collected: r.Collected,
			Day:       collectionDay(r.Collected),
		}
		if r.Reference != nil {
			row.Reference = strings.TrimSpace(*r.Reference)
		}
		if r.ResultText != nil {
			row.ResultText = strings.TrimSpace(*r.ResultText)
		}
		if r.DefUnit != nil {
			row.DefUnit = strings.TrimSpace(*r.DefUnit)
		}
		if r.DefName != nil {
			row.DefName = strings.TrimSpace(*r.DefName)
		}
		if r.DefGloss != nil {
			row.DefGloss = strings.TrimSpace(*r.DefGloss)
		}
		row.Text = resultDisplayText(row.ResultText, r.Numeric)
		out = append(out, row)
	}
	return out, nil
}

// anamnesisRow — uma resposta de anamnese já ligada ao item de escore.
type anamnesisRow struct {
	ScoreItemID uuid.UUID
	Code        string
	Name        string
	Unit        string
	Points      float64
	Level       *int
	Numeric     *float64
	Text        string
	Day         string
}

// loadAnamnesisRows traz as respostas de anamnese do paciente.
//
// O nível vem de `selected_level`, NÃO de `numeric_value`: é o que o motor do escore usa, e ler o
// número cru no lugar do nível já fez o escore sair zerado antes. O `numeric_value` serve para
// desenhar a régua quando o item é de medida (IMC, razão cintura/altura), não para classificar.
//
// Também não há filtro de "finalizada": a tabela nem tem status, e a anamnese da consulta de hoje é
// justamente a que mais interessa para a devolutiva de hoje.
func (s *PatientPlanDossierService) loadAnamnesisRows(patientID uuid.UUID) ([]anamnesisRow, error) {
	var raw []struct {
		ScoreItemID uuid.UUID
		Code        *string
		Name        string
		Unit        *string
		Points      *float64
		Level       *int
		Numeric     *float64
		TextValue   *string
		Measured    time.Time
	}
	err := s.db.
		Table("anamnesis_items AS ai").
		Select(`ai.score_item_id      AS score_item_id,
		        si.anamnese_item_code AS code,
		        si.name               AS name,
		        si.unit               AS unit,
		        si.points             AS points,
		        ai.selected_level     AS level,
		        ai.numeric_value      AS numeric,
		        ai.text_value         AS text_value,
		        a.consultation_date   AS measured`).
		Joins("JOIN anamnesis a ON a.id = ai.anamnesis_id AND a.deleted_at IS NULL").
		Joins("JOIN score_items si ON si.id = ai.score_item_id AND si.deleted_at IS NULL").
		Where("a.patient_id = ? AND ai.deleted_at IS NULL AND ai.selected_level IS NOT NULL", patientID).
		Order("measured ASC").
		Scan(&raw).Error
	if err != nil {
		return nil, err
	}

	out := make([]anamnesisRow, 0, len(raw))
	for _, r := range raw {
		row := anamnesisRow{
			ScoreItemID: r.ScoreItemID,
			Name:        r.Name,
			Level:       r.Level,
			Numeric:     r.Numeric,
			Day:         collectionDay(r.Measured),
		}
		if r.Code != nil {
			row.Code = *r.Code
		}
		if r.Unit != nil {
			row.Unit = *r.Unit
		}
		if r.Points != nil {
			row.Points = *r.Points
		}
		switch {
		case r.Numeric != nil:
			row.Text = formatNumberPT(*r.Numeric)
		case r.TextValue != nil:
			row.Text = strings.TrimSpace(*r.TextValue)
		}
		out = append(out, row)
	}
	return out, nil
}

// anamnesisFindings transforma as respostas em achados, com a MESMA regra de classificação e
// ordenação dos exames. Uma resposta que muda entre consultas tem direção, igual a um exame que
// muda entre coletas.
func anamnesisFindings(rows []anamnesisRow, lostBySnapshot map[uuid.UUID]float64) []dto.PlanFinding {
	byItem := map[uuid.UUID][]anamnesisRow{}
	var ordem []uuid.UUID
	for _, r := range rows {
		if _, seen := byItem[r.ScoreItemID]; !seen {
			ordem = append(ordem, r.ScoreItemID)
		}
		byItem[r.ScoreItemID] = append(byItem[r.ScoreItemID], r)
	}

	out := make([]dto.PlanFinding, 0, len(ordem))
	for _, id := range ordem {
		hist := byItem[id]
		last := hist[len(hist)-1]
		if last.Level == nil {
			continue
		}
		level := *last.Level

		trend := dto.PlanTrendSingle
		if len(hist) >= 2 {
			prev := hist[len(hist)-2]
			switch {
			case prev.Level == nil:
			case level > *prev.Level:
				trend = dto.PlanTrendImproving
			case level < *prev.Level:
				trend = dto.PlanTrendWorsening
			default:
				trend = dto.PlanTrendStable
			}
		}

		lost, ok := lostBySnapshot[id]
		if !ok {
			lost = last.Points * float64(5-clampInt(level, 0, 5)) / 5
		}

		f := dto.PlanFinding{
			Source: dto.PlanSourceAnamnesis,
			Code:   last.Code, Name: last.Name, Unit: last.Unit,
			Level: level, Text: last.Text, Date: last.Day,
			Points: last.Points, Trend: trend, PointsLost: math.Max(lost, 0),
		}
		if last.Numeric != nil {
			f.Value = *last.Numeric
		}

		switch {
		case level <= 2:
			f.Kind = dto.PlanFindingMoving
			f.Reason = fmt.Sprintf("nível %d na escala do escore", level)
		case trend == dto.PlanTrendWorsening:
			f.Kind = dto.PlanFindingMoving
			f.Reason = "piorou em relação à consulta anterior"
		case level >= 4:
			f.Kind = dto.PlanFindingStrong
			f.Reason = fmt.Sprintf("nível %d na escala do escore", level)
		default:
			continue
		}
		out = append(out, f)
	}
	return out
}

// loadVitals traz as duas últimas medidas de consultório. Duas, e não uma, porque um número de
// pressão sozinho não diz se está subindo.
func (s *PatientPlanDossierService) loadVitals(patientID uuid.UUID) ([]dto.PlanDossierVitals, error) {
	var rows []models.ConsultationVitals
	if err := s.db.Where("patient_id = ?", patientID).
		Order("measured_at DESC").Limit(2).Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]dto.PlanDossierVitals, 0, len(rows))
	for i := range rows {
		v := rows[i]
		out = append(out, dto.PlanDossierVitals{
			MeasuredAt:  collectionDay(v.MeasuredAt),
			SystolicBP:  v.SystolicBP,
			DiastolicBP: v.DiastolicBP,
			HeartRate:   v.HeartRate,
			Weight:      v.Weight,
			Height:      v.Height,
			Waist:       v.WaistCircumference,
			BMI:         v.BMI,
		})
	}
	return out, nil
}

// buildRulers monta uma régua por código de exame que o paciente tenha feito E que tenha escala no
// catálogo do escore aplicável a ele.
func (s *PatientPlanDossierService) buildRulers(patient *models.Patient, rows []labRow) (map[string]dto.PlanRuler, error) {
	byCode := map[string][]labRow{}
	for _, r := range rows {
		byCode[r.Code] = append(byCode[r.Code], r)
	}
	if len(byCode) == 0 {
		return map[string]dto.PlanRuler{}, nil
	}

	// Sinônimos de unidade do catálogo, escopados por exame: sem eles a régua do sódio sumiria
	// só porque o item diz `mEq/L` e o laudo diz `mmol/L`.
	catalogo := carregaCatalogoDeExames(s.db)

	codes := make([]string, 0, len(byCode))
	for c := range byCode {
		codes = append(codes, c)
	}

	var items []models.ScoreItem
	if err := s.db.
		Preload("Levels", func(db *gorm.DB) *gorm.DB {
			return db.Where("deleted_at IS NULL").Order("level ASC")
		}).
		Where("lab_test_code IN ? AND deleted_at IS NULL", codes).
		Find(&items).Error; err != nil {
		return nil, err
	}

	// Valor mais recente por exame: é o contexto que decide se um item condicionado se aplica.
	ultimoPorCodigo := map[string]float64{}
	for _, r := range rows {
		ultimoPorCodigo[r.Code] = r.Numeric // rows vem ordenado por data crescente
	}

	// Um código pode ter várias variantes (por sexo, faixa etária, menopausa, TRH). O motor de
	// aplicabilidade do escore é quem escolhe a do paciente — o mesmo usado no cálculo, para a
	// régua nunca discordar do escore.
	//
	// `RequirementMet` é o segundo filtro: item cuja escala só vale dentro do contexto de outro
	// exame (a razão %Free PSA só discrimina com PSA total entre 4 e 10) some quando o contexto não
	// bate. Sem isso ele virava o achado número 1 de um paciente sem nada na próstata.
	applicable := map[string][]models.ScoreItem{}
	for i := range items {
		if items[i].AppliesToPatient(patient) && items[i].RequirementMet(ultimoPorCodigo) {
			code := *items[i].LabTestCode
			applicable[code] = append(applicable[code], items[i])
		}
	}

	// pickScoringItem é o MESMO desempate que o classificador de resultados usa
	// (lab_result_batch_service.go) para gravar `lab_results.level`: escolher outra variante aqui
	// faria a régua desenhar uma escala diferente da que o escore usa para o mesmo exame.
	//
	// Isso alinha a ESCALA, não garante que o nível gravado no resultado antigo bata: aquele foi
	// calculado com a idade/menopausa do dia da importação. Por isso o nível de cada ponto é
	// recalculado sobre a escala escolhida, em buildHistory.
	out := make(map[string]dto.PlanRuler, len(applicable))
	for code, cands := range applicable {
		item := pickScoringItem(cands)
		if item == nil {
			continue
		}
		// Escala e exame têm que falar da mesma grandeza. Três itens do sedimento urinário têm a
		// escala em células/campo enquanto o laboratório reporta /µL: classificar ali põe um
		// resultado de 0,5 na faixa "≤10" e o paciente lê "ótimo" sobre um número que ninguém
		// comparou. Melhor a régua não existir do que existir errada.
		if len(byCode[code]) > 0 && !item.UnitMatches(byCode[code][0].DefUnit, catalogo.sinonimosDe(item.LabTestCode)) {
			continue
		}
		ruler, ok := buildRuler(code, item, byCode[code])
		if !ok {
			continue
		}
		out[code] = ruler
	}
	return out, nil
}

// buildRuler converte a escala de um ScoreItem + o histórico do paciente numa régua desenhável.
// Devolve ok=false para item sem escala numérica utilizável (item categórico, de operador "=").
func buildRuler(code string, item *models.ScoreItem, rows []labRow) (dto.PlanRuler, bool) {
	edges := levelEdges(item.Levels)
	if len(edges) < 2 {
		return dto.PlanRuler{}, false
	}

	history := buildHistory(rows, item.Levels)
	axis := rulerAxis(edges, history)
	segments := levelSegments(item.Levels, axis)
	if len(segments) == 0 {
		return dto.PlanRuler{}, false
	}
	for i := range history {
		history[i].Plot = clampFloat(history[i].Value, axis[0], axis[1])
	}

	unit := ""
	if item.Unit != nil {
		unit = *item.Unit
	}
	points := 0.0
	if item.Points != nil {
		points = *item.Points
	}
	nomePaciente, glosa := "", ""
	if len(rows) > 0 {
		nomePaciente, glosa = rows[0].DefName, rows[0].DefGloss
	}
	return dto.PlanRuler{
		Code:        code,
		ScoreItemID: item.ID.String(),
		Name:        item.Name,
		PatientName: nomePaciente,
		Gloss:       glosa,
		Unit:        unit,
		Points:      points,
		Axis:        axis,
		Edges:       edges,
		Segments:    segments,
		History:     history,
	}, true
}

// levelEdges — as fronteiras finitas entre níveis, sem repetição e ordenadas.
func levelEdges(levels []models.ScoreLevel) []float64 {
	seen := map[float64]bool{}
	var out []float64
	for i := range levels {
		for _, p := range []*string{levels[i].LowerLimit, levels[i].UpperLimit} {
			v, ok := parseLimit(p)
			if !ok {
				continue
			}
			v = roundTo(v, 9)
			if !seen[v] {
				seen[v] = true
				out = append(out, v)
			}
		}
	}
	sort.Float64s(out)
	return out
}

// rulerAxis — o intervalo desenhado. Parte do vão entre a primeira e a última fronteira, com
// respiro nas pontas, e se estica quando o paciente tem valor fora dele. Nunca fica negativo:
// exame não tem valor negativo, e um eixo começando abaixo de zero desperdiça metade da barra.
func rulerAxis(edges []float64, history []dto.PlanRulerPoint) []float64 {
	span := edges[len(edges)-1] - edges[0]
	if span <= 0 {
		// Escala de fronteira única: sem vão para medir, o respiro vem do próprio valor.
		span = math.Abs(edges[0])
		if span == 0 {
			span = 1
		}
	}
	pad := rulerAxisPad * span
	lo, hi := edges[0]-pad, edges[len(edges)-1]+pad

	// O eixo é a ESCALA DO ESCORE, e não se estica para caber o paciente.
	//
	// Esticar parecia gentil e destruía a régua: a PCR da paciente estava em 63,1 com o escore
	// indo até 10, então o eixo virava [0, 65,6] e as seis faixas do escore ocupavam 15% da barra.
	// O HOMA-IR, 22%. O resto era uma extensão vazia de "nível 0", e o paciente via uma bolinha no
	// fim de uma barra que não dizia mais o que é bom. Medido nas 18 réguas de um deck real: 16
	// bem, essas 2 esmagadas.
	//
	// O deck aprovado corta exatamente aí: PCR até 15,7 e HOMA até 5,0, com a bolinha presa na
	// ponta. É a regra que a gramática registra — o `axis` é o único número afinado à mão quando um
	// valor extremo esmaga a escala. O gerador desfazia esse ajuste toda vez.
	//
	// Não some dado: `rulerSVG` prende o ponto às bordas (math.Min/math.Max) e o valor continua
	// impresso ao lado. Bolinha na borda É a informação de que o paciente está fora da escala.
	//
	// Consequência assumida: no T-score da densitometria, que vive todo no negativo, a bolinha de
	// um -3,5 volta a encostar na borda esquerda. A esticada tinha sido acrescentada para evitar
	// isso; encostar na borda é a leitura correta de estar fora da faixa.
	// Piso em zero SÓ quando a escala é não-negativa. Exame de sangue não tem valor negativo e um
	// eixo começando abaixo de zero desperdiçaria metade da barra — mas T-score de densitometria
	// vive INTEIRO no negativo (-2,5 a -1), e nele o piso apagava a escala: o eixo virava [0,1],
	// todo segmento caía fora e a barra saía pintada de "ótimo" com a bolinha presa na ponta
	// esquerda. Um paciente com osteoporose receberia um PDF dizendo o contrário do exame.
	if edges[0] >= 0 && lo < 0 {
		lo = 0
	}
	if hi <= lo {
		hi = lo + 1
	}
	return []float64{lo, hi}
}

// levelSegments — cada nível vira uma faixa fechada. Nível de ponta ("≤15", ">300") é aberto por
// definição e fecha na borda do eixo, senão não teria largura para desenhar.
func levelSegments(levels []models.ScoreLevel, axis []float64) []dto.PlanRulerSegment {
	var out []dto.PlanRulerSegment
	for i := range levels {
		lv := levels[i]
		lo, hasLo := parseLimit(lv.LowerLimit)
		hi, hasHi := parseLimit(lv.UpperLimit)

		var a, b float64
		var label string
		switch lv.Operator {
		case "<", "<=":
			// O limite de um nível "abaixo de X" pode vir gravado em qualquer um dos dois campos.
			edge, ok := hi, hasHi
			if !ok {
				edge, ok = lo, hasLo
			}
			if !ok {
				continue
			}
			a, b = axis[0], edge
			// O sinal tem que ser o do operador. Rotular um nível "< 15" como "≤15" diz ao paciente
			// que 15 pertence a esta faixa, quando o motor o classifica na faixa vizinha.
			if lv.Operator == "<" {
				label = "<" + formatNumberPT(edge)
			} else {
				label = "≤" + formatNumberPT(edge)
			}
		case ">", ">=":
			if !hasLo {
				continue
			}
			a, b = lo, axis[1]
			if lv.Operator == ">=" {
				label = "≥" + formatNumberPT(lo)
			} else {
				label = ">" + formatNumberPT(lo)
			}
		case "between":
			if !hasLo || !hasHi {
				continue
			}
			a, b = lo, hi
			// Faixa com número negativo não pode usar hífen de intervalo: "-2,5--2" não se lê.
			sep := "-"
			if lo < 0 || hi < 0 {
				sep = " a "
			}
			label = formatNumberPT(lo) + sep + formatNumberPT(hi)
		default:
			// Operador "=" é nível categórico (anamnese): não tem faixa para desenhar.
			continue
		}
		out = append(out, dto.PlanRulerSegment{Level: lv.Level, A: roundTo(a, 9), B: roundTo(b, 9), Label: label})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].A < out[j].A })
	return out
}

// buildHistory — os resultados do paciente na ordem da coleta. Um exame repetido no mesmo dia fica
// com a última leitura: o laudo mais recente do dia é o que vale.
//
// O nível de cada ponto é RECALCULADO sobre a escala do item escolhido, com a mesma função que o
// motor usa (ScoreLevel.EvaluatesTrue), em vez de reaproveitar o `lab_results.level` gravado na
// importação. O gravado foi calculado com a idade/menopausa que o paciente tinha NAQUELE dia e com
// a variante aplicável naquele momento; num item com recorte etário, o paciente troca de faixa e o
// nível antigo passa a descrever outra escala. Como a régua desenha ESTA escala, o ponto tem que
// ser classificado por ela — senão a bolinha cai num segmento e o rótulo diz outro.
func buildHistory(rows []labRow, levels []models.ScoreLevel) []dto.PlanRulerPoint {
	byDay := map[string]labRow{}
	var order []string
	for _, r := range rows {
		if _, seen := byDay[r.Day]; !seen {
			order = append(order, r.Day)
		}
		byDay[r.Day] = r
	}
	sort.Strings(order)

	out := make([]dto.PlanRulerPoint, 0, len(order))
	for _, day := range order {
		r := byDay[day]
		out = append(out, dto.PlanRulerPoint{
			Date:  day,
			Value: r.Numeric,
			Level: levelForValue(levels, r.Numeric),
			Text:  r.Text,
			Ref:   r.Reference,
			Name:  r.TestName,
			Over:  strings.HasPrefix(r.ResultText, ">"),
			Under: strings.HasPrefix(r.ResultText, "<"),
			Plot:  r.Numeric,
		})
	}
	return out
}

// levelForValue devolve o nível em que o valor cai, pela MESMA regra do motor de classificação
// (primeiro nível cujo condicional bate, na ordem do nível). Valor fora de todas as faixas fica sem
// nível: melhor não classificado do que classificado errado.
func levelForValue(levels []models.ScoreLevel, value float64) *int {
	for i := range levels {
		if levels[i].EvaluatesTrue(value) {
			l := levels[i].Level
			return &l
		}
	}
	return nil
}

// classifyFindings separa os achados em "o que está bem" e "o que está se movendo", e ordena cada
// lista pelo que mais pesa. É PROPOSTA: quem decide o que vira slide é o médico.
//
// As listas nascem vazias, nunca nil: em JSON, slice nil vira `null`, e quem consome no front
// quebra no .map(). Paciente sem exame nenhum é caso normal, não excepcional.
func classifyFindings(rulers map[string]dto.PlanRuler, lostBySnapshot map[uuid.UUID]float64) (strong, moving []dto.PlanFinding) {
	strong, moving = []dto.PlanFinding{}, []dto.PlanFinding{}
	for code, r := range rulers {
		if len(r.History) == 0 {
			continue
		}
		last := r.History[len(r.History)-1]
		if last.Level == nil {
			continue
		}
		level := *last.Level

		trend := trendOf(r)

		lost, ok := 0.0, false
		if itemID, pErr := uuid.Parse(r.ScoreItemID); pErr == nil {
			lost, ok = lostBySnapshot[itemID]
		}
		if !ok {
			// Sem escore calculado (ou item que o escore não avaliou): estima pelo peso do item e
			// pelo quanto falta para o nível ótimo. Serve só para ordenar.
			lost = r.Points * float64(5-clampInt(level, 0, 5)) / 5
		}

		f := dto.PlanFinding{
			Source: dto.PlanSourceLab,
			Code:   code, Name: r.Name, Unit: r.Unit,
			Level: level, Value: last.Value, Text: last.Text, Date: last.Date,
			Points: r.Points, Trend: trend, PointsLost: math.Max(lost, 0),
		}
		if dias, ok := diasDesde(last.Date); ok {
			f.DaysAgo = &dias
			f.Stale = dias > diasParaEnvelhecer
		}

		switch {
		case level <= 2:
			f.Kind = dto.PlanFindingMoving
			f.Reason = fmt.Sprintf("nível %d na escala do escore", level)
			moving = append(moving, f)
		case trend == dto.PlanTrendWorsening:
			f.Kind = dto.PlanFindingMoving
			f.Reason = "caiu de nível em relação ao exame anterior"
			moving = append(moving, f)
		case level >= 4:
			f.Kind = dto.PlanFindingStrong
			f.Reason = fmt.Sprintf("nível %d na escala do escore", level)
			strong = append(strong, f)
		}
	}

	sortMoving(moving)
	sortStrong(strong)
	return strong, moving
}

// diasParaEnvelhecer — acima disso a medida deixa de ser o retrato de hoje. Dezoito meses é o que
// separa "este é o seu exame" de "isto precisa ser refeito antes de decidir".
const diasParaEnvelhecer = 548

func diasDesde(dia string) (int, bool) {
	d, err := time.Parse("2006-01-02", dia)
	if err != nil {
		return 0, false
	}
	n := int(time.Now().In(saoPaulo()).Sub(d).Hours() / 24)
	if n < 0 {
		n = 0
	}
	return n, true
}

// trendOf decide a direção olhando a DISTÂNCIA ATÉ O ÓTIMO, não só o número do nível.
//
// Comparando com os decks aprovados: o LDL do Ricardo foi 151 → 115 → 115 → 127, e as três últimas
// medidas estão todas no nível 3. Por nível, "estável"; pelo valor, subiu 12 pontos e é um dos
// quatro achados de manchete do deck ("subindo, sem tratamento"). Movimento dentro da mesma faixa é
// movimento, e a direção é o sinal que mais importa.
func trendOf(r dto.PlanRuler) dto.PlanFindingTrend {
	if len(r.History) < 2 {
		return dto.PlanTrendSingle
	}
	atual, anterior := r.History[len(r.History)-1], r.History[len(r.History)-2]
	if atual.Level != nil && anterior.Level != nil && *atual.Level != *anterior.Level {
		if *atual.Level > *anterior.Level {
			return dto.PlanTrendImproving
		}
		return dto.PlanTrendWorsening
	}
	// Mesmo nível: desempata pela distância até a faixa ótima.
	dA, okA := distanciaAoOtimo(r, atual.Value)
	dB, okB := distanciaAoOtimo(r, anterior.Value)
	if !okA || !okB || math.Abs(dA-dB) < 1e-9 {
		return dto.PlanTrendStable
	}
	if dA > dB {
		return dto.PlanTrendWorsening
	}
	return dto.PlanTrendImproving
}

// distanciaAoOtimo — o quanto o valor está longe da melhor faixa da régua. Zero se está dentro.
// Funciona com escala que piora para a direita, para a esquerda ou nas duas pontas.
func distanciaAoOtimo(r dto.PlanRuler, v float64) (float64, bool) {
	melhor := -1
	for _, s := range r.Segments {
		if s.Level > melhor {
			melhor = s.Level
		}
	}
	if melhor < 0 {
		return 0, false
	}
	dist := math.Inf(1)
	for _, s := range r.Segments {
		if s.Level != melhor {
			continue
		}
		switch {
		case v >= s.A && v <= s.B:
			return 0, true
		case v < s.A:
			dist = math.Min(dist, s.A-v)
		default:
			dist = math.Min(dist, v-s.B)
		}
	}
	if math.IsInf(dist, 1) {
		return 0, false
	}
	return dist, true
}

// sortMoving ordena o que está se movendo por PONTOS PERDIDOS, mas joga para o fim o que está
// velho: uma medida de dois anos atrás não sustenta conduta hoje, ela sustenta um pedido de exame.
func sortMoving(f []dto.PlanFinding) {
	sort.Slice(f, func(i, j int) bool {
		if f[i].Stale != f[j].Stale {
			return !f[i].Stale
		}
		if f[i].PointsLost != f[j].PointsLost {
			return f[i].PointsLost > f[j].PointsLost
		}
		return f[i].Code < f[j].Code
	})
}

// sortStrong ordena o que está bem por PESO DO ITEM, não por pontos perdidos.
//
// Em nível 4-5 o paciente não perdeu quase nada por definição, então ordenar por pontos perdidos
// deixaria a lista inteira empatada em zero e sem sinal nenhum. Pior: boa parte da anamnese é
// checklist de ausência ("Adrenalectomia: não", "Amputação de membro: não"), tudo em nível 5, o que
// afoga os achados que valem alguma coisa — é a mesma inflação já conhecida no escore. Ordenar pelo
// peso do item faz o que de fato importa vir primeiro: o marcador pesado que está no ótimo.
func sortStrong(f []dto.PlanFinding) {
	sort.Slice(f, func(i, j int) bool {
		if f[i].Points != f[j].Points {
			return f[i].Points > f[j].Points
		}
		if f[i].Level != f[j].Level {
			return f[i].Level > f[j].Level
		}
		return f[i].Code < f[j].Code
	})
}

func (s *PatientPlanDossierService) loadLastLabRequest(patientID uuid.UUID) (*dto.PlanDossierLabRequest, error) {
	var lr models.LabRequest
	err := s.db.Where("patient_id = ?", patientID).Order("date DESC").First(&lr).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	out := &dto.PlanDossierLabRequest{
		ID:    lr.ID.String(),
		Date:  lr.Date.Format("2006-01-02"),
		Exams: lr.Exams,
	}
	if lr.SignedAt != nil {
		v := lr.SignedAt.In(saoPaulo()).Format(time.RFC3339)
		out.SignedAt = &v
	}
	return out, nil
}

func (s *PatientPlanDossierService) loadPrescriptions(patientID uuid.UUID) ([]dto.PlanDossierPrescription, error) {
	// Os preloads são a mesma cadeia de `PrescriptionService.reloadPrescription`, na mesma ordem de
	// impressão. Sem eles o dossiê devolvia a casca da receita e o "para levar" saía sem dose.
	var rows []models.Prescription
	if err := s.db.
		Preload("Medications").
		Preload("Formulas", func(db *gorm.DB) *gorm.DB { return db.Order("display_order") }).
		Preload("Formulas.Components", func(db *gorm.DB) *gorm.DB { return db.Order("display_order") }).
		Where("patient_id = ? AND status = ?", patientID, models.PrescriptionActive).
		Order("prescription_date DESC").
		Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]dto.PlanDossierPrescription, 0, len(rows))
	for i := range rows {
		p := dto.PlanDossierPrescription{
			ID:     rows[i].ID.String(),
			Type:   string(rows[i].Type),
			Date:   rows[i].PrescriptionDate.In(saoPaulo()).Format("2006-01-02"),
			Status: string(rows[i].Status),
		}
		if rows[i].SignedAt != nil {
			v := rows[i].SignedAt.In(saoPaulo()).Format(time.RFC3339)
			p.SignedAt = &v
		}
		for _, f := range rows[i].Formulas {
			fo := dto.PlanDossierFormula{
				Name: f.Name, Form: f.PharmaceuticalForm, Posology: f.Posology,
				Route: f.Route, Duration: f.Duration,
				Quantity: strings.TrimRight(strings.TrimRight(
					strconv.FormatFloat(f.QuantityToDispense, 'f', 3, 64), "0"), ".") + " " + f.QuantityUnit,
			}
			for _, c := range f.Components {
				fo.Components = append(fo.Components, dto.PlanDossierFormulaComponent{
					Substance: c.Substance, Quantity: c.Quantity, Unit: c.Unit,
				})
			}
			p.Formulas = append(p.Formulas, fo)
		}
		for _, m := range rows[i].Medications {
			p.Medications = append(p.Medications, dto.PlanDossierMedication{
				Name: m.MedicationName, Concentration: m.Concentration, Dosage: m.Dosage,
				Frequency: m.Frequency, Route: m.Route, Duration: m.Duration,
			})
		}
		out = append(out, p)
	}
	return out, nil
}

// ---- utilitários numéricos ----

// parseLimit lê um limite de score_level, que é texto livre no banco ("15", "1.003", " 40 ").
func parseLimit(s *string) (float64, bool) {
	if s == nil {
		return 0, false
	}
	t := strings.TrimSpace(*s)
	if t == "" {
		return 0, false
	}
	t = strings.ReplaceAll(t, ",", ".")
	v, err := strconv.ParseFloat(t, 64)
	if err != nil {
		return 0, false
	}
	return v, true
}

// formatNumberPT — número como o paciente lê: sem casa decimal inútil e com vírgula decimal.
func formatNumberPT(v float64) string {
	if math.Abs(v-math.Round(v)) < 1e-9 {
		return strconv.FormatInt(int64(math.Round(v)), 10)
	}
	// 'f' e nunca 'g': com 'g' um CK de 12345,6 sairia como "1,235e+04" no PDF do paciente, e um
	// triglicerídeo de 1234,5 perderia a casa decimal em silêncio. Arredonda em 4 casas (o banco
	// guarda decimal(12,4)) e deixa o Go escolher a representação mais curta que volta ao mesmo
	// número, sem zero à toa.
	r := math.Round(v*1e4) / 1e4
	return strings.Replace(strconv.FormatFloat(r, 'f', -1, 64), ".", ",", 1)
}

// resultDisplayText preserva a grafia do laudo ("3,00" continua "3,00", não vira "3") quando ela
// bate com o valor numérico gravado; caso contrário formata o número.
func resultDisplayText(raw string, numeric float64) string {
	t := strings.TrimSpace(raw)
	if t != "" {
		cleaned := strings.TrimLeft(t, "<>≤≥ ")
		if v, err := strconv.ParseFloat(strings.Replace(cleaned, ",", ".", 1), 64); err == nil {
			if math.Abs(v-numeric) < 1e-9 {
				return strings.Replace(cleaned, ".", ",", 1)
			}
		}
	}
	return formatNumberPT(numeric)
}

func roundTo(v float64, places int) float64 {
	p := math.Pow(10, float64(places))
	return math.Round(v*p) / p
}

func clampFloat(v, lo, hi float64) float64 {
	return math.Max(lo, math.Min(hi, v))
}

func clampInt(v, lo, hi int) int {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}
