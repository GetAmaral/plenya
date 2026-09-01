package dto

import "github.com/plenya/api/internal/pdfdoc"

// Dossiê do plano de paciente — o insumo DERIVADO do prontuário para montar a devolutiva
// (o "deck" de resultados). Substitui o `reguas.json` que era montado à mão a cada paciente.
//
// A régua é o átomo visual da devolutiva: uma barra horizontal onde a faixa inteira é a escala
// Plenya do exame (nível 0 = pior, nível 5 = ótimo) e o ponto marca onde o paciente está. Ela
// existe porque barra colorida com rótulo avaliativo é o formato que o paciente entende melhor —
// melhor que tabela e que número cru — e porque mostrar a faixa-meta NO LUGAR da faixa de
// referência do laboratório (e não somada a ela) melhora o entendimento.
//
// O formato de PlanRuler espelha 1:1 o `reguas.json` legado, de propósito: os decks já feitos
// (`pacs/<NOME>/deck/build.py`) consomem esse JSON sem adaptação.

// PlanRulerSegment — uma faixa de nível dentro da régua. `a`/`b` já vêm resolvidos contra o eixo,
// então segmento aberto ("≤15", ">300") chega fechado na ponta do eixo e é desenhável direto.
type PlanRulerSegment struct {
	Level int     `json:"level"`
	A     float64 `json:"a"`
	B     float64 `json:"b"`
	// Label legível da faixa ("≤15", "15-30", ">300"), com vírgula decimal PT-BR.
	Label string `json:"label"`
}

// PlanRulerPoint — um resultado do paciente na linha do tempo da régua.
type PlanRulerPoint struct {
	Date  string  `json:"date"`  // AAAA-MM-DD da coleta
	Value float64 `json:"value"` // valor já convertido para a unidade padrão
	Level *int    `json:"level"` // nível do escore; nil quando não classificado
	Text  string  `json:"text"`  // valor formatado PT-BR ("1,023")
	Ref   string  `json:"ref"`   // faixa de referência impressa pelo laboratório
	Name  string  `json:"name"`  // nome do exame como veio no laudo
	Over  bool    `json:"over"`  // resultado censurado para cima (">1000")
	Under bool    `json:"under"` // resultado censurado para baixo ("<0,10")
	Plot  float64 `json:"plot"`  // valor de plotagem (igual a Value, preso ao eixo)
}

// PlanRuler — a régua de um exame para ESTE paciente: a escala (segmentos) vem do catálogo do
// escore, filtrada por sexo/idade/menopausa; o histórico vem dos resultados dele.
type PlanRuler struct {
	Code string `json:"code"` // lab_test_definitions.code (ex.: "PLNCEFB97FD")
	// ScoreItemID — QUAL variante do item foi usada. Um mesmo código pode ter várias (por sexo,
	// faixa etária, menopausa); sem isto não dá para saber de qual escala esta régua veio.
	ScoreItemID string  `json:"scoreItemId"`
	Name        string  `json:"name"`   // nome do score_item aplicável ("Ferritina - Homens")
	Unit        string  `json:"unit"`   // unidade padrão do item
	Points      float64 `json:"points"` // peso do item no escore

	// Axis — [mínimo, máximo] do eixo desenhado. Sempre 2 elementos.
	Axis []float64 `json:"axis"`
	// Edges — fronteiras finitas entre níveis, ordenadas.
	Edges    []float64          `json:"edges"`
	Segments []PlanRulerSegment `json:"segments"`
	History  []PlanRulerPoint   `json:"history"` // ordenado da coleta mais antiga para a mais recente
}

// PlanFindingKind — em qual lista da devolutiva o achado entra.
type PlanFindingKind string

const (
	// PlanFindingStrong — nível 4 ou 5: entra em "o que está bem".
	PlanFindingStrong PlanFindingKind = "strong"
	// PlanFindingMoving — nível ≤2, ou piora consistente no histórico: entra em
	// "o que está se movendo".
	PlanFindingMoving PlanFindingKind = "moving"
)

// PlanFindingTrend — direção da mudança entre o penúltimo e o último resultado. A direção importa
// mais do que a magnitude para quem lê, e é onde o leitor mais erra, então ela vem explícita em vez
// de ficar implícita no desenho.
type PlanFindingTrend string

const (
	PlanTrendSingle    PlanFindingTrend = "single"    // só um resultado, não dá para falar em direção
	PlanTrendStable    PlanFindingTrend = "stable"    // mesmo nível
	PlanTrendImproving PlanFindingTrend = "improving" // subiu de nível
	PlanTrendWorsening PlanFindingTrend = "worsening" // caiu de nível
)

// PlanFindingSource — de onde o achado veio. O plano não se monta só com exame: a leitura da
// consulta e a anamnese carregam metade do que vira slide.
type PlanFindingSource string

const (
	PlanSourceLab       PlanFindingSource = "lab"       // resultado de exame
	PlanSourceAnamnesis PlanFindingSource = "anamnesis" // resposta de anamnese
)

// PlanFinding — um achado candidato a virar slide, já classificado e pontuado. É uma PROPOSTA
// ordenada: quem decide o que entra no deck, em que ordem e com que título é o médico.
type PlanFinding struct {
	Source PlanFindingSource `json:"source"`
	Code   string            `json:"code"`
	Name   string            `json:"name"`
	Unit   string            `json:"unit"`
	Level  int               `json:"level"`
	Value  float64           `json:"value"`
	Text   string            `json:"text"`
	Date   string            `json:"date"`
	Points float64           `json:"points"`

	Kind  PlanFindingKind  `json:"kind"`
	Trend PlanFindingTrend `json:"trend"`

	// DaysAgo — idade da medida mais recente. Um achado de dois anos atrás não é o retrato de hoje.
	DaysAgo *int `json:"daysAgo,omitempty"`
	// Stale — a medida é antiga demais para sustentar conduta. Comparado com os decks aprovados,
	// um HOMA-IR de 2024 liderava o ranking do Ricardo enquanto o próprio deck o tratava como
	// exame A REFAZER, não como achado atual.
	Stale bool `json:"stale,omitempty"`

	// PointsLost — quanto do peso do item o paciente está deixando na mesa. É o critério de
	// ordenação: o que mais pesa aparece primeiro.
	PointsLost float64 `json:"pointsLost"`
	// Reason — por que este achado entrou nesta lista, em uma frase.
	Reason string `json:"reason"`
}

// PlanDossierPatient — identificação mínima; o dossiê não carrega dado sensível que a devolutiva
// não use.
type PlanDossierPatient struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}

// PlanDossierSnapshot — o escore vigente, quando existe.
type PlanDossierSnapshot struct {
	ID              string  `json:"id"`
	CalculatedAt    string  `json:"calculatedAt"`
	TotalPercentage float64 `json:"totalPercentage"`
}

// PlanDossierLabRequest — o último pedido de exames, para o slide "os exames que faltam".
type PlanDossierLabRequest struct {
	ID       string  `json:"id"`
	Date     string  `json:"date"`
	Exams    string  `json:"exams"`
	SignedAt *string `json:"signedAt,omitempty"`
}

// PlanDossierPrescription — receita vigente, para o slide "para levar".
type PlanDossierPrescription struct {
	ID       string  `json:"id"`
	Type     string  `json:"type"`
	Date     string  `json:"date"`
	Status   string  `json:"status"`
	SignedAt *string `json:"signedAt,omitempty"`
}

// PlanDossierVitals — a medida da consulta. O deck cita esses números direto ("a pressão está em
// 120 por 70, sem remédio nenhum"), e eles não vêm de exame nenhum: vêm de quem mediu no
// consultório.
type PlanDossierVitals struct {
	MeasuredAt  string   `json:"measuredAt"`
	SystolicBP  *int     `json:"systolicBp,omitempty"`
	DiastolicBP *int     `json:"diastolicBp,omitempty"`
	HeartRate   *int     `json:"heartRate,omitempty"`
	Weight      *float64 `json:"weight,omitempty"`
	Height      *float64 `json:"height,omitempty"`
	Waist       *float64 `json:"waistCircumference,omitempty"`
	BMI         *float64 `json:"bmi,omitempty"`
}

// PlanDossierResponse — tudo que a montagem do plano consegue derivar sozinha do prontuário.
//
// O que NÃO está aqui, e continua sendo escrito à mão, é justamente o que é julgamento clínico:
// a leitura dos achados, o arco narrativo, os títulos em voz de paciente e as condutas.
type PlanDossierResponse struct {
	Patient  PlanDossierPatient   `json:"patient"`
	Snapshot *PlanDossierSnapshot `json:"snapshot,omitempty"`
	Rulers   map[string]PlanRuler `json:"rulers"`
	// Vitals — a medida mais recente da consulta e a anterior, para dar direção ao número.
	Vitals      []PlanDossierVitals       `json:"vitals"`
	Strong      []PlanFinding             `json:"strong"`
	Moving      []PlanFinding             `json:"moving"`
	CarePlan    []CarePlanItemResponse    `json:"carePlan"`
	LabRequest  *PlanDossierLabRequest    `json:"labRequest,omitempty"`
	Medications []PlanDossierPrescription `json:"prescriptions"`
	GeneratedAt string                    `json:"generatedAt"`
}

// ---- Plano de devolutiva (patient_plans) ----

// SavePatientPlanRequest — cria ou reescreve o plano. `content` é a lista de slides.
type SavePatientPlanRequest struct {
	Title            string             `json:"title" validate:"omitempty,max=300"`
	Content          []pdfdoc.DeckSlide `json:"content"`
	SourceSnapshotID *string            `json:"sourceSnapshotId,omitempty" validate:"omitempty,uuid"`
}

// PatientPlanResponse — o plano como o EMR o mostra.
type PatientPlanResponse struct {
	ID        string             `json:"id"`
	PatientID string             `json:"patientId"`
	Title     string             `json:"title"`
	Status    string             `json:"status"`
	Version   int                `json:"version"`
	Content   []pdfdoc.DeckSlide `json:"content"`

	SourceSnapshotID *string `json:"sourceSnapshotId,omitempty"`
	AuthorUserID     string  `json:"authorUserId"`
	PublishedAt      *string `json:"publishedAt,omitempty"`
	// Documentos publicados no portal: o 16:9 para ver e mandar, o A4 para imprimir.
	Document16x9ID *string `json:"document16x9Id,omitempty"`
	DocumentA4ID   *string `json:"documentA4Id,omitempty"`

	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// PatientPlanOverflowResponse — a resposta de recusa quando algum slide transborda.
type PatientPlanOverflowResponse struct {
	Error  string                `json:"error"`
	Slides []pdfdoc.DeckOverflow `json:"slides"`
}
