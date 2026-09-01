package services

import (
	"strings"
	"testing"

	"github.com/google/uuid"

	"github.com/plenya/api/internal/models"
)

func ptrF(f float64) *float64 { return &f }

// itemDeSedimento monta um item cuja escala está em células/campo, como os três itens de
// sedimento urinário do catálogo.
func itemDeSedimento() *models.ScoreItem {
	id := uuid.New()
	return &models.ScoreItem{
		ID:          id,
		Name:        "Leucócitos (sedimento urinário)",
		LabTestCode: ptrS("PLNLEUCOURI"),
		Unit:        ptrS("células/campo"),
		Points:      ptrF(10),
		Levels: []models.ScoreLevel{
			{Level: 5, Operator: "<=", UpperLimit: ptrS("10")},
			{Level: 0, Operator: ">", LowerLimit: ptrS("10")},
		},
	}
}

func resultado(nome, unidade string, valor float64) models.LabResult {
	return models.LabResult{
		ID:            uuid.New(),
		TestType:      "urina",
		TestName:      nome,
		Unit:          ptrS(unidade),
		ResultNumeric: ptrF(valor),
	}
}

// Um resultado de 0,5/µL caía na faixa "≤10 células/campo" e saía como nível ÓTIMO. São
// grandezas diferentes: concentração não se compara com contagem por campo do microscópio.
func TestEvaluateScoreItem_RecusaClassificarComUnidadeDeOutraGrandeza(t *testing.T) {
	s := &ScoreSnapshotService{}
	item := itemDeSedimento()
	labs := map[string]models.LabResult{
		"PLNLEUCOURI": resultado("Leucócitos", "/µL", 0.5),
	}

	got := s.evaluateScoreItem(&models.Patient{}, item, uuid.New(), labs, nil, catalogoDeExames{})

	if got.Status != models.EvaluationStatusNotApplicable {
		t.Fatalf("status = %v, queria not_applicable (0,5/µL não pode virar nível ótimo)", got.Status)
	}
	if got.LevelNumber != nil {
		t.Fatalf("nível %d atribuído; não deveria classificar nada", *got.LevelNumber)
	}
	if got.NotEvaluatedReason == nil || !strings.Contains(*got.NotEvaluatedReason, "células/campo") {
		t.Fatalf("motivo = %v, queria citar as duas unidades", got.NotEvaluatedReason)
	}
}

func TestEvaluateScoreItem_ClassificaQuandoUnidadeBate(t *testing.T) {
	s := &ScoreSnapshotService{}
	item := itemDeSedimento()
	labs := map[string]models.LabResult{
		"PLNLEUCOURI": resultado("Leucócitos", "células/campo", 3),
	}

	got := s.evaluateScoreItem(&models.Patient{}, item, uuid.New(), labs, nil, catalogoDeExames{})

	if got.Status != models.EvaluationStatusEvaluated {
		t.Fatalf("status = %v (motivo: %v), queria evaluated", got.Status, got.NotEvaluatedReason)
	}
	if got.LevelNumber == nil || *got.LevelNumber != 5 {
		t.Fatalf("nível = %v, queria 5", got.LevelNumber)
	}
}

// O exame sem unidade declarada não pode travar a avaliação: a guarda só vale quando as duas
// pontas dizem em que unidade estão.
func TestEvaluateScoreItem_UnidadeAusenteNaoBloqueia(t *testing.T) {
	s := &ScoreSnapshotService{}
	item := itemDeSedimento()
	r := resultado("Leucócitos", "", 3)
	r.Unit = nil
	labs := map[string]models.LabResult{"PLNLEUCOURI": r}

	got := s.evaluateScoreItem(&models.Patient{}, item, uuid.New(), labs, nil, catalogoDeExames{})

	if got.Status != models.EvaluationStatusEvaluated {
		t.Fatalf("status = %v (motivo: %v), queria evaluated", got.Status, got.NotEvaluatedReason)
	}
}

// itemComContexto é o %Free PSA, que só discrimina com PSA total entre 4 e 10 ng/mL.
func itemComContexto() *models.ScoreItem {
	return &models.ScoreItem{
		ID:              uuid.New(),
		Name:            "PSA Livre/Total (%Free PSA)",
		LabTestCode:     ptrS("PLNFREEPSA"),
		Unit:            ptrS("%"),
		Points:          ptrF(28),
		RequiresLabCode: ptrS("PLN1BFE6CA3"),
		RequiresMin:     ptrF(4),
		RequiresMax:     ptrF(10),
		Levels: []models.ScoreLevel{
			{Level: 0, Operator: "<=", UpperLimit: ptrS("10")},
			{Level: 5, Operator: ">", LowerLimit: ptrS("25")},
		},
	}
}

// %Free PSA fora da zona cinzenta do PSA total: o motivo tem que nomear o exame e mostrar o
// valor medido, não cuspir o código PLN… que não diz nada a quem lê o escore.
func TestEvaluateScoreItem_MotivoDeContextoNomeiaOExame(t *testing.T) {
	s := &ScoreSnapshotService{}
	item := &models.ScoreItem{
		ID:              uuid.New(),
		Name:            "PSA Livre/Total (%Free PSA)",
		LabTestCode:     ptrS("PLNFREEPSA"),
		Unit:            ptrS("%"),
		Points:          ptrF(28),
		RequiresLabCode: ptrS("PLN1BFE6CA3"),
		RequiresMin:     ptrF(4),
		RequiresMax:     ptrF(10),
		Levels: []models.ScoreLevel{
			{Level: 0, Operator: "<=", UpperLimit: ptrS("10")},
			{Level: 5, Operator: ">", LowerLimit: ptrS("25")},
		},
	}
	labs := map[string]models.LabResult{
		"PLNFREEPSA":  resultado("PSA Livre/Total", "%", 8.8),
		"PLN1BFE6CA3": resultado("PSA total", "ng/mL", 1.81),
	}

	got := s.evaluateScoreItem(&models.Patient{}, item, uuid.New(), labs, nil, catalogoDeExames{})

	if got.Status != models.EvaluationStatusNotApplicable {
		t.Fatalf("status = %v, queria not_applicable", got.Status)
	}
	motivo := ""
	if got.NotEvaluatedReason != nil {
		motivo = *got.NotEvaluatedReason
	}
	for _, quer := range []string{"PSA total", "entre 4 e 10", "1,81"} {
		if !strings.Contains(motivo, quer) {
			t.Errorf("motivo %q não contém %q", motivo, quer)
		}
	}
	if strings.Contains(motivo, "PLN1BFE6CA3") {
		t.Errorf("motivo %q ainda mostra o código cru em vez do nome", motivo)
	}
}

// Sem o exame de referência medido, o item não se aplica — e o motivo tem que dizer isso, não
// fingir que o valor estava fora de faixa.
func TestEvaluateScoreItem_MotivoQuandoExameDeReferenciaNaoFoiMedido(t *testing.T) {
	s := &ScoreSnapshotService{}
	item := &models.ScoreItem{
		ID:              uuid.New(),
		Name:            "PSA Livre/Total (%Free PSA)",
		LabTestCode:     ptrS("PLNFREEPSA"),
		RequiresLabCode: ptrS("PLN1BFE6CA3"),
		RequiresMin:     ptrF(4),
		RequiresMax:     ptrF(10),
		Levels:          []models.ScoreLevel{{Level: 0, Operator: "<=", UpperLimit: ptrS("10")}},
	}
	labs := map[string]models.LabResult{
		"PLNFREEPSA": resultado("PSA Livre/Total", "%", 8.8),
	}

	cat := catalogoDeExames{nomes: map[string]string{"PLN1BFE6CA3": "PSA total"}}
	got := s.evaluateScoreItem(&models.Patient{}, item, uuid.New(), labs, nil, cat)

	if got.Status != models.EvaluationStatusNotApplicable {
		t.Fatalf("status = %v, queria not_applicable", got.Status)
	}
	motivo := ""
	if got.NotEvaluatedReason != nil {
		motivo = *got.NotEvaluatedReason
	}
	if !strings.Contains(motivo, "não foi medido") {
		t.Errorf("motivo %q não diz que o exame de referência não foi medido", motivo)
	}
	if !strings.Contains(motivo, "PSA total") {
		t.Errorf("motivo %q não usa o nome do catálogo", motivo)
	}
	if strings.Contains(motivo, "PLN1BFE6CA3") {
		t.Errorf("motivo %q ainda mostra o código cru", motivo)
	}
}

// O exame de referência foi feito mas veio sem número (qualitativo/texto). Dizer "não foi
// medido" mandaria pedir de novo um exame que já está no prontuário.
func TestEvaluateScoreItem_ReferenciaSemValorNumerico(t *testing.T) {
	s := &ScoreSnapshotService{}
	item := itemComContexto()
	ref := resultado("PSA total", "ng/mL", 0)
	ref.ResultNumeric = nil
	labs := map[string]models.LabResult{
		"PLNFREEPSA":  resultado("PSA Livre/Total", "%", 8.8),
		"PLN1BFE6CA3": ref,
	}

	got := s.evaluateScoreItem(&models.Patient{}, item, uuid.New(), labs, nil, catalogoDeExames{})

	motivo := ""
	if got.NotEvaluatedReason != nil {
		motivo = *got.NotEvaluatedReason
	}
	if strings.Contains(motivo, "não foi medido") {
		t.Errorf("motivo %q afirma que não foi medido, mas o exame existe", motivo)
	}
	if !strings.Contains(motivo, "sem valor numérico") {
		t.Errorf("motivo %q não explica que faltou o número", motivo)
	}
}

// Só piso, só teto: os dois ramos da frase da faixa.
func TestEvaluateScoreItem_MotivoComApenasPisoOuTeto(t *testing.T) {
	s := &ScoreSnapshotService{}
	labs := map[string]models.LabResult{
		"PLNFREEPSA":  resultado("PSA Livre/Total", "%", 8.8),
		"PLN1BFE6CA3": resultado("PSA total", "ng/mL", 1.81),
	}

	soPiso := itemComContexto()
	soPiso.RequiresMax = nil
	if got := s.evaluateScoreItem(&models.Patient{}, soPiso, uuid.New(), labs, nil, catalogoDeExames{}); got.NotEvaluatedReason == nil ||
		!strings.Contains(*got.NotEvaluatedReason, "a partir de 4") {
		t.Errorf("só piso: motivo = %v", got.NotEvaluatedReason)
	}

	soTeto := itemComContexto()
	soTeto.RequiresMin = nil
	soTeto.RequiresMax = ptrF(1)
	if got := s.evaluateScoreItem(&models.Patient{}, soTeto, uuid.New(), labs, nil, catalogoDeExames{}); got.NotEvaluatedReason == nil ||
		!strings.Contains(*got.NotEvaluatedReason, "até 1") {
		t.Errorf("só teto: motivo = %v", got.NotEvaluatedReason)
	}
}

// Item sem unidade declarada + exame com unidade: a guarda não pode nem bloquear nem estourar.
// É o caso que prova a ausência de nil-deref no `*item.Unit` da mensagem.
func TestEvaluateScoreItem_ItemSemUnidadeNaoEstoura(t *testing.T) {
	s := &ScoreSnapshotService{}
	item := itemDeSedimento()
	item.Unit = nil
	labs := map[string]models.LabResult{
		"PLNLEUCOURI": resultado("Leucócitos", "/µL", 3),
	}

	got := s.evaluateScoreItem(&models.Patient{}, item, uuid.New(), labs, nil, catalogoDeExames{})

	if got.Status != models.EvaluationStatusEvaluated {
		t.Fatalf("status = %v (motivo: %v), queria evaluated", got.Status, got.NotEvaluatedReason)
	}
}

// Unidade incompatível descarta o EXAME, não o item: se o médico preencheu a anamnese para o
// mesmo item, esse dado é válido e tem que ser usado.
func TestEvaluateScoreItem_UnidadeIncompativelCaiParaAnamnese(t *testing.T) {
	s := &ScoreSnapshotService{}
	item := itemDeSedimento()
	labs := map[string]models.LabResult{
		"PLNLEUCOURI": resultado("Leucócitos", "/µL", 0.5),
	}
	nivel := 5
	anamnese := map[uuid.UUID]models.AnamnesisItem{
		item.ID: {ID: uuid.New(), SelectedLevel: &nivel},
	}

	got := s.evaluateScoreItem(&models.Patient{}, item, uuid.New(), labs, anamnese, catalogoDeExames{})

	if got.Status != models.EvaluationStatusEvaluated {
		t.Fatalf("status = %v (motivo: %v), queria evaluated pela anamnese", got.Status, got.NotEvaluatedReason)
	}
	if got.LevelNumber == nil || *got.LevelNumber != 5 {
		t.Fatalf("nível = %v, queria 5 (o que o médico registrou)", got.LevelNumber)
	}
}
