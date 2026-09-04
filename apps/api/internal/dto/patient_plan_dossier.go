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
//
// Carregava só id, tipo, data e status, e com isso o slide "para levar" nascia sem insumo nenhum:
// o prompt mandava tirar dose de `prescriptions` e o dossiê não trazia dose. O conteúdo estruturado
// SEMPRE existiu no banco (`prescription_formulas` + `prescription_formula_components`); o que
// faltava era o `Preload` e estes campos.
type PlanDossierPrescription struct {
	ID       string  `json:"id"`
	Type     string  `json:"type"`
	Date     string  `json:"date"`
	Status   string  `json:"status"`
	SignedAt *string `json:"signedAt,omitempty"`

	// Formulas — as manipuladas. Cada uma é uma cápsula/creme com sua posologia.
	Formulas []PlanDossierFormula `json:"formulas,omitempty"`
	// Medications — as industrializadas.
	Medications []PlanDossierMedication `json:"medications,omitempty"`
}

// PlanDossierFormula — uma fórmula manipulada da receita.
type PlanDossierFormula struct {
	Name string `json:"name"`
	Form string `json:"form,omitempty"`
	// Posology é a frase que o paciente lê ("Tomar 1 cápsula de 12/12 horas"), e é ela que vira o
	// `dose` do "para levar".
	Posology   string                        `json:"posology,omitempty"`
	Route      string                        `json:"route,omitempty"`
	Duration   int                           `json:"durationDays,omitempty"`
	Quantity   string                        `json:"quantityToDispense,omitempty"`
	Components []PlanDossierFormulaComponent `json:"components,omitempty"`
}

// PlanDossierFormulaComponent — um ativo da fórmula, com a quantidade.
//
// `note` NÃO entra: ela sai impressa na receita assinada e é instrução de manipulação, não
// conteúdo para o paciente ler no deck.
type PlanDossierFormulaComponent struct {
	Substance string  `json:"substance"`
	Quantity  float64 `json:"quantity"`
	Unit      string  `json:"unit"`
}

// PlanDossierMedication — um medicamento industrializado da receita.
type PlanDossierMedication struct {
	Name          string `json:"name"`
	Concentration string `json:"concentration,omitempty"`
	Dosage        string `json:"dosage,omitempty"`
	Frequency     string `json:"frequency,omitempty"`
	Route         string `json:"route,omitempty"`
	Duration      int    `json:"durationDays,omitempty"`
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

// ---- Dossiê congelado ----

// PlanDossierStaleness — se o prontuário andou desde o congelamento, e no quê.
//
// Devolve os motivos e não só um booleano porque a decisão de refrescar depende deles: exame novo
// importa num deck que fala de exame; aferição nova, num deck que não cita pressão, não.
type PlanDossierStaleness struct {
	DossierID string   `json:"dossierId,omitempty"`
	FrozenAt  string   `json:"frozenAt,omitempty"`
	Stale     bool     `json:"stale"`
	Reasons   []string `json:"reasons,omitempty"`
}

// PlanDossierCitation — onde no deck um exame é citado.
type PlanDossierCitation struct {
	SlideID string `json:"slideId"`
	Index   int    `json:"index"`
	Title   string `json:"title,omitempty"`
}

// PlanDossierChange — um exame citado no deck cujo valor mudou entre dois congelamentos.
type PlanDossierChange struct {
	Code    string                `json:"code"`
	Name    string                `json:"name"`
	Unit    string                `json:"unit,omitempty"`
	Was     string                `json:"was"`
	Now     string                `json:"now"`
	CitedIn []PlanDossierCitation `json:"citedIn"`
}

// PlanDossierRefreshResponse — o resultado de refrescar.
//
// `Changed` é restrito ao que o deck CITA: um dossiê tem dezenas de réguas e dizer "mudou" sobre o
// conjunto não ajuda. O que decide é "destes que você citou, estes mudaram, e estão nestes slides".
type PlanDossierRefreshResponse struct {
	DossierID  string              `json:"dossierId"`
	Changed    []PlanDossierChange `json:"changed"`
	Unaffected int                 `json:"unaffected"`
}

// ---- Plano de devolutiva (patient_plans) ----

// SavePatientPlanRequest — cria ou reescreve o plano. `content` é a lista de slides.
type SavePatientPlanRequest struct {
	// ExpectedRevision — a revisão que o cliente acha ser a corrente. Quando vem e não bate, a
	// gravação é recusada com 409 em vez de sobrescrever escrita que o cliente não viu. Ausente
	// significa "não sei", e passa: cliente antigo não é bloqueado.
	ExpectedRevision *int `json:"expectedRevision,omitempty"`

	Title            string             `json:"title" validate:"omitempty,max=300"`
	Content          []pdfdoc.DeckSlide `json:"content"`
	SourceSnapshotID *string            `json:"sourceSnapshotId,omitempty" validate:"omitempty,uuid"`
}

// PatientPlanResponse — o plano como o EMR o mostra.
type PatientPlanResponse struct {
	ID        string `json:"id"`
	PatientID string `json:"patientId"`
	Title     string `json:"title"`
	Status    string `json:"status"`
	Version   int    `json:"version"`
	// RevisionSeq — a última edição do rascunho, e o token que o cliente devolve em
	// `expectedRevision` para não sobrescrever escrita que não viu. Conta edições; `Version`
	// conta publicações.
	RevisionSeq int                `json:"revisionSeq"`
	Content     []pdfdoc.DeckSlide `json:"content"`

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

// ---- Assistente da devolutiva ----

// PlanMessage — um turno da conversa.
type PlanMessage struct {
	ID        string `json:"id"`
	Seq       int    `json:"seq"`
	Role      string `json:"role"`
	Body      string `json:"body"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
}

// PlanNumeralMatch — uma origem candidata para um número.
type PlanNumeralMatch struct {
	Value  float64 `json:"value"`
	Unit   string  `json:"unit,omitempty"`
	Source string  `json:"source"`
	Label  string  `json:"label"`
}

// PlanNumeralProof — um número que a sugestão escreve, e de onde ele pode ter vindo.
//
// É o que aparece ao lado do botão de aceitar. Sem isto o médico estaria aceitando prosa, e a
// evidência é clara sobre o que acontece quando a revisão é só leitura.
type PlanNumeralProof struct {
	Numeral string             `json:"numeral"`
	Found   bool               `json:"found"`
	Matches []PlanNumeralMatch `json:"matches,omitempty"`
}

// PlanSuggestion — uma alteração proposta esperando aceite.
type PlanSuggestion struct {
	ID         string             `json:"id"`
	Op         string             `json:"op"`
	SlideID    string             `json:"slideId,omitempty"`
	Path       string             `json:"path,omitempty"`
	Class      string             `json:"class"`
	Rationale  string             `json:"rationale"`
	Status     string             `json:"status"`
	OldValue   any                `json:"oldValue,omitempty"`
	NewValue   any                `json:"newValue,omitempty"`
	Provenance []PlanNumeralProof `json:"provenance,omitempty"`
	CreatedAt  string             `json:"createdAt"`
}

// PlanAppliedOp — uma operação que entrou direto, ou foi recusada.
type PlanAppliedOp struct {
	Op      string `json:"op"`
	SlideID string `json:"slideId,omitempty"`
	Path    string `json:"path,omitempty"`
	Reason  string `json:"reason"`
}

// PlanRejectedOp — mesma forma, outro significado.
type PlanRejectedOp = PlanAppliedOp

// PlanAssistantTurn — o resultado de um turno.
type PlanAssistantTurn struct {
	Reply       string               `json:"reply"`
	Applied     []PlanAppliedOp      `json:"applied,omitempty"`
	Suggestions []PlanSuggestion     `json:"suggestions,omitempty"`
	Rejected    []PlanRejectedOp     `json:"rejected,omitempty"`
	Plan        *PatientPlanResponse `json:"plan,omitempty"`
	RevisionSeq int                  `json:"revisionSeq"`
	Failed      bool                 `json:"failed,omitempty"`
	Error       string               `json:"error,omitempty"`
	// Stale — o plano foi alterado por outra mão enquanto o modelo respondia, então nada foi
	// aplicado. A conversa fica registrada e o médico reenvia.
	Stale bool `json:"stale,omitempty"`
	// CacheReadTokens — a partir do segundo turno tem que ser maior que zero. Zero significa que
	// algo volátil entrou antes do ponto de cache e o prefixo está sendo reenviado inteiro.
	CacheReadTokens int `json:"cacheReadTokens,omitempty"`
}

// PlanSkipped — sugestão que não foi aplicada, e por quê.
type PlanSkipped struct {
	ID     string `json:"id"`
	Reason string `json:"reason"`
}

// PlanResolveResult — resultado de aceitar ou recusar. Parcial é resposta legítima.
type PlanResolveResult struct {
	Resolved    []string             `json:"resolved"`
	Skipped     []PlanSkipped        `json:"skipped,omitempty"`
	RevisionSeq int                  `json:"revisionSeq"`
	Plan        *PatientPlanResponse `json:"plan,omitempty"`
}

// SendPlanMessageRequest — o corpo do turno.
type SendPlanMessageRequest struct {
	Body             string `json:"body" validate:"required,max=4000"`
	ClientMessageID  string `json:"clientMessageId" validate:"omitempty,max=64"`
	ExpectedRevision *int   `json:"expectedRevision"`
}

// ResolveSuggestionsRequest — aceitar ou recusar, por id ou por slide.
type ResolveSuggestionsRequest struct {
	Action           string   `json:"action" validate:"required,oneof=accept reject"`
	SuggestionIDs    []string `json:"suggestionIds"`
	SlideID          string   `json:"slideId"`
	ExpectedRevision *int     `json:"expectedRevision"`
}

// PlanGenUsage — o que uma geração consumiu, somando as chamadas todas (arco, seções e reparo).
//
// Sai na resposta em vez de ficar só no log: no painel da Anthropic uma geração não se distingue de
// um turno de conversa nem de uma leitura de laudo, e sem separar não dá para responder quanto
// custa gerar um plano.
type PlanGenUsage struct {
	InputTokens int `json:"inputTokens"`
	// CacheReadTokens custa ~0,1x a entrada; CacheWriteTokens ~1,25x.
	CacheReadTokens  int `json:"cacheReadTokens"`
	CacheWriteTokens int `json:"cacheWriteTokens"`
	OutputTokens     int `json:"outputTokens"`
	LatencyMs        int `json:"latencyMs"`
	// Chamadas — quantas idas ao modelo: 1 do arco, uma por seção, e 1 se houve reparo.
	Chamadas int `json:"chamadas"`
}

// PlanGenWarningKind — a natureza do aviso da geração.
type PlanGenWarningKind string

const (
	// PlanGenWarningNumeral — número escrito que não existe no dossiê.
	PlanGenWarningNumeral PlanGenWarningKind = "numeral"
	// PlanGenWarningRuler — régua ou linha de resumo cujo exame não está no dossiê.
	PlanGenWarningRuler PlanGenWarningKind = "regua"
	// PlanGenWarningEstilo — desvio do padrão medido nos decks aprovados (punch, contagem de
	// régua, travessão). Não é erro de dado: é o deck fora da forma.
	PlanGenWarningEstilo PlanGenWarningKind = "estilo"
	// PlanGenWarningLacuna — o que o deck NÃO pôde ter porque o prontuário não tinha o dado. Não é
	// defeito da geração, e a distinção importa: o médico precisa saber se o remédio é editar o
	// deck ou registrar a consulta.
	PlanGenWarningLacuna PlanGenWarningKind = "lacuna"
)

// PlanGenWarning — um número que o modelo escreveu e o servidor NÃO encontrou no dossiê.
//
// Não bloqueia a geração: bloquear jogaria fora o deck inteiro por causa de um número. Vira aviso
// no slide exato, para o médico olhar aquela frase em vez de reler vinte slides.
type PlanGenWarning struct {
	// Kind separa "número que não confere" de "régua que saiu": a tela dizia "N número(s) que não
	// encontrei" para os dois, e um deck onde só caiu uma régua acusava número não verificado.
	Kind       PlanGenWarningKind `json:"kind,omitempty"`
	SlideIndex int                `json:"slideIndex"`
	SlideID    string             `json:"slideId,omitempty"`
	Title      string             `json:"title,omitempty"`
	Numeral    string             `json:"numeral"`
	Reason     string             `json:"reason"`
}

// GeneratePlanRequest — o corpo (opcional) do pedido de geração.
type GeneratePlanRequest struct {
	Title string `json:"title,omitempty"`
	// Instruction deixa o médico dirigir a geração ("foque no ferro e no sono", "sem o bloco de
	// condutas"). Vazio = o arco padrão.
	Instruction string `json:"instruction,omitempty"`
}
