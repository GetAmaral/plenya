package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/pdfdoc"
)

// O assistente da devolutiva.
//
// Um turno faz, em ordem e numa transação: lê o dossiê CONGELADO, monta o índice numérico, chama o
// modelo, TRIA cada operação contra o vocabulário e contra o índice, aplica o que é seguro, grava
// uma revisão com autoria da ferramenta, e transforma o resto em sugestões pendentes com a origem
// de cada número anexada.
//
// A chamada é síncrona. Não usa a fila de `processing_jobs` porque aquela tabela tem FK obrigatória
// para lote de exame — adaptá-la custaria mais que a feature — e porque o retry automático de um
// turno que já aplicou operação as re-aplicaria. O retry certo aqui é o médico reenviar, protegido
// por `clientMessageId`.

var (
	ErrPlanAssistantDuplicate = errors.New("este turno já foi processado")
	ErrPlanSuggestionStale    = errors.New("o slide mudou depois que esta sugestão foi criada")
)

type PatientPlanAssistantService struct {
	db        *gorm.DB
	ai        *AIService
	plans     *PatientPlanService
	dossiers  *PatientPlanDossierService
	revisions *PatientPlanRevisionService

	model         string
	promptVersion string
}

func NewPatientPlanAssistantService(
	db *gorm.DB, ai *AIService, plans *PatientPlanService,
	dossiers *PatientPlanDossierService, model, promptVersion string,
) *PatientPlanAssistantService {
	if promptVersion == "" {
		promptVersion = "p1"
	}
	return &PatientPlanAssistantService{
		db: db, ai: ai, plans: plans, dossiers: dossiers,
		revisions: NewPatientPlanRevisionService(db),
		model:     model, promptVersion: promptVersion,
	}
}

type SendPlanMessageInput struct {
	PlanID, PatientID, UserID uuid.UUID
	Body                      string
	ClientMessageID           string
	ExpectedRevision          *int
}

// SendMessage roda um turno completo.
func (s *PatientPlanAssistantService) SendMessage(in SendPlanMessageInput) (*dto.PlanAssistantTurn, error) {
	if strings.TrimSpace(in.Body) == "" {
		return nil, errors.New("mensagem vazia")
	}

	// Idempotência antes de qualquer trabalho: reenviar depois de fechar a aba não pode duplicar.
	if in.ClientMessageID != "" {
		var ja models.PatientPlanMessage
		err := s.db.Where("plan_id = ? AND client_message_id = ?", in.PlanID, in.ClientMessageID).
			First(&ja).Error
		if err == nil {
			return nil, fmt.Errorf("%w (%s)", ErrPlanAssistantDuplicate, ja.ID)
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	plan, err := s.plans.load(in.PlanID, in.PatientID)
	if err != nil {
		return nil, err
	}
	if err := CheckExpected(plan, in.ExpectedRevision); err != nil {
		return nil, err
	}

	dossie, dossieRow, err := s.dossiers.Current(in.PlanID)
	if err != nil {
		return nil, err
	}
	indice := BuildNumericIndex(dossie)

	turnos, err := s.historico(in.PlanID)
	if err != nil {
		return nil, err
	}

	dossieJSON, err := podaDossieParaPrompt(dossie, plan.Content)
	if err != nil {
		return nil, err
	}
	slidesJSON, err := json.Marshal(plan.Content)
	if err != nil {
		return nil, err
	}

	resultado, meta, chamadaErr := s.ai.EditPatientPlan(PlanEditRequest{
		DossierJSON:   dossieJSON,
		SlidesJSON:    string(slidesJSON),
		Turns:         turnos,
		Instruction:   in.Body,
		PromptVersion: s.promptVersion,
		Model:         s.model,
	})

	// A mensagem do médico é gravada sempre, mesmo quando a chamada falha: a conversa é registro.
	var out dto.PlanAssistantTurn
	err = s.db.Transaction(func(tx *gorm.DB) error {
		// Trava a linha do plano ANTES de qualquer coisa.
		//
		// Resolve dois problemas com o mesmo cadeado. O primeiro: entre carregar o plano e chegar
		// aqui passaram de dez a vinte segundos de chamada ao modelo, e um PUT do médico (ou outra
		// aba) nesse intervalo seria sobrescrito em silêncio pelo `tx.Save(plan)` lá embaixo —
		// exatamente a corrida que `revision_seq` existe para impedir, e que a checagem feita
		// ANTES da chamada não cobre. O segundo: `proximoSeq` é um `MAX(seq)+1` sem trava, e dois
		// turnos simultâneos calculavam o mesmo número e batiam no índice único depois de as duas
		// chamadas já terem sido pagas.
		var atual models.PatientPlan
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("id = ?", in.PlanID).First(&atual).Error; err != nil {
			return err
		}
		planoMudouDurante := atual.RevisionSeq != plan.RevisionSeq

		seq, err := s.proximoSeq(tx, in.PlanID)
		if err != nil {
			return err
		}
		msgUsuario := models.PatientPlanMessage{
			PlanID: in.PlanID, Seq: seq, Role: "user", Body: in.Body,
			UserID: &in.UserID, ClientMessageID: in.ClientMessageID,
			DossierID: &dossieRow.ID,
		}
		if err := tx.Create(&msgUsuario).Error; err != nil {
			return err
		}

		msgIA := models.PatientPlanMessage{
			PlanID: in.PlanID, Seq: seq + 1, Role: "assistant",
			AIModel: meta.Model, AIPromptVersion: s.promptVersion,
			AIInputTokens: meta.InputTokens, AICacheReadTokens: meta.CacheReadTokens,
			AIOutputTokens: meta.OutputTokens, AIStopReason: meta.StopReason,
			LatencyMs: meta.LatencyMs, DossierID: &dossieRow.ID,
		}
		if chamadaErr != nil {
			msgIA.Status = "failed"
			msgIA.ErrorMessage = chamadaErr.Error()
			msgIA.Body = "A rodada falhou e nada foi alterado."
			if err := tx.Create(&msgIA).Error; err != nil {
				return err
			}
			out = dto.PlanAssistantTurn{Reply: msgIA.Body, Failed: true, Error: chamadaErr.Error()}
			return nil
		}
		msgIA.Status = "ok"
		msgIA.Body = resultado.Reply
		if err := tx.Create(&msgIA).Error; err != nil {
			return err
		}

		// Triagem: o servidor reclassifica, a classe declarada pelo modelo é só sinal.
		ops := converteOps(resultado.Operations)
		triadas := Triage(plan.Content, ops, indice)

		var aplicar []PlanOp
		for i, t := range triadas {
			if t.Decision == TriageApply {
				aplicar = append(aplicar, ops[i])
			}
		}

		// O plano mudou embaixo da chamada: nada é aplicado, e o turno inteiro vira sugestão nem
		// é opção, porque as ops foram calculadas sobre um conteúdo que não existe mais e o
		// `base_hash` seria tirado do slide NOVO, escondendo o conflito em vez de mostrá-lo. A
		// conversa fica registrada e o médico reenvia; perder a chamada é o preço honesto.
		if planoMudouDurante {
			for i := range triadas {
				triadas[i].Decision = TriageReject
				triadas[i].Reason = "o plano foi alterado enquanto esta resposta era gerada"
			}
			aplicar = nil
			msgIA.Body = "O plano mudou enquanto eu respondia, então não alterei nada. Reenvie o pedido."
			if err := tx.Model(&msgIA).Updates(map[string]any{"body": msgIA.Body}).Error; err != nil {
				return err
			}
			out.Reply = msgIA.Body
			out.Stale = true
		}

		if len(aplicar) > 0 {
			// Uma op de cada vez, e não o lote inteiro.
			//
			// `setPath` volta por `json.Unmarshal` num `DeckSlide` TIPADO, então um valor com o
			// tipo errado vindo do modelo (`"legend":"true"`, `"title":42`) faz a gravação falhar.
			// A triagem classifica por caminho, não por tipo, e deixa esses casos passarem como
			// texto autoral. Abortando o lote, o turno inteiro caía: a mensagem do médico não era
			// gravada (contra o que a linha acima promete) e a chamada paga se perdia por causa de
			// uma vírgula do modelo. Agora a op ruim vira recusada e o resto entra.
			novos := plan.Content
			var entraram []PlanOp
			for _, op := range aplicar {
				passo, err := ApplyOps(novos, []PlanOp{op})
				if err != nil {
					for i := range triadas {
						if triadas[i].Op.SlideID == op.SlideID && triadas[i].Op.Path == op.Path &&
							triadas[i].Decision == TriageApply {
							triadas[i].Decision = TriageReject
							triadas[i].Reason = "não foi possível gravar este campo: " + err.Error()
							break
						}
					}
					continue
				}
				novos = passo
				entraram = append(entraram, op)
			}
			aplicar = entraram
		}

		if len(aplicar) > 0 {
			novos, err := ApplyOps(plan.Content, aplicar)
			if err != nil {
				return fmt.Errorf("falha ao aplicar operação triada: %w", err)
			}
			plan.Content = novos
			plan.Status = models.PatientPlanDraft
			seqRev, err := s.revisions.Record(tx, RecordRevisionInput{
				Plan: plan, Title: plan.Title, Content: plan.Content,
				AuthorKind: models.PlanAuthorAssistant, CreatedByID: in.UserID,
				Reason: models.PlanRevisionAIApply, Ops: triadas,
				MessageID: &msgIA.ID, DossierID: &dossieRow.ID,
				AIModel: meta.Model, AIPromptVersion: s.promptVersion,
			})
			if err != nil {
				return err
			}
			plan.RevisionSeq = seqRev
			if err := tx.Save(plan).Error; err != nil {
				return err
			}
		}

		// O que não aplicou vira sugestão pendente, com a origem de cada número.
		for i, t := range triadas {
			switch t.Decision {
			case TriageSuggest:
				sug, err := s.gravaSugestao(tx, plan, msgIA.ID, ops[i], t)
				if err != nil {
					return err
				}
				out.Suggestions = append(out.Suggestions, sug)
			case TriageReject:
				out.Rejected = append(out.Rejected, dto.PlanRejectedOp{
					Op: string(ops[i].Op), Path: ops[i].Path, SlideID: ops[i].SlideID, Reason: t.Reason,
				})
			case TriageApply:
				out.Applied = append(out.Applied, dto.PlanAppliedOp{
					Op: string(ops[i].Op), Path: ops[i].Path, SlideID: ops[i].SlideID, Reason: t.Reason,
				})
			}
		}

		out.Reply = resultado.Reply
		out.RevisionSeq = plan.RevisionSeq
		out.Plan = toPatientPlanDTO(plan)
		out.CacheReadTokens = meta.CacheReadTokens
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (s *PatientPlanAssistantService) gravaSugestao(
	tx *gorm.DB, plan *models.PatientPlan, msgID uuid.UUID, op PlanOp, t TriagedOp,
) (dto.PlanSuggestion, error) {
	var out dto.PlanSuggestion

	// O hash do slide alvo no momento em que a sugestão nasce. Sem ele, aceitar vinte minutos
	// depois sobrescreveria o que o médico escreveu à mão nesse meio-tempo.
	baseHash := ""
	if i := indiceDoSlide(plan.Content, op.SlideID); i >= 0 {
		h, err := HashDeckContent([]pdfdoc.DeckSlide{plan.Content[i]})
		if err != nil {
			return out, err
		}
		baseHash = h
	}

	antes := ""
	if op.Op == OpEdit && op.SlideID != "" {
		if i := indiceDoSlide(plan.Content, op.SlideID); i >= 0 {
			antes = valorAtualComoTexto(&plan.Content[i], op.Path)
		}
	}

	// O que a sugestão INTRODUZ. Para `edit` é o valor do campo; para as estruturais é o slide
	// novo ou a ordem nova.
	//
	// Sem isso, `add` e `reorder` eram inaceitáveis por construção: a triagem manda TODA op
	// estrutural para sugestão, mas só `op.Value` era persistido, e `opDaSugestao` devolvia uma op
	// sem `Slide` e sem `Order`. Aceitar batia em "add sem slide" dentro da transação de
	// `ResolveSuggestions` e abortava o lote inteiro, derrubando junto as sugestões de texto
	// aceitas na mesma chamada.
	introduzido := op.Value
	switch op.Op {
	case OpAdd:
		introduzido = op.Slide
	case OpReorder:
		introduzido = op.Order
	}
	novoValor, _ := json.Marshal(introduzido)
	antigoValor, _ := json.Marshal(antes)
	proveniencia, _ := json.Marshal(t.Provenance)

	classe := "numeric"
	if op.Op != OpEdit {
		classe = "structural"
	}

	sug := models.PatientPlanSuggestion{
		PlanID: plan.ID, MessageID: msgID,
		Op: string(op.Op), SlideID: op.SlideID, AfterSlideID: op.AfterSlideID,
		FieldPath: op.Path,
		OldValue:  datatypes.JSON(antigoValor), NewValue: datatypes.JSON(novoValor),
		BaseHash: baseHash, Class: classe, Rationale: t.Reason,
		Provenance: datatypes.JSON(proveniencia),
		Status:     models.SuggestionPending,
	}

	// Uma sugestão nova no mesmo campo do mesmo slide substitui a anterior: duas pendentes
	// disputando o mesmo campo é ruído, e aceitar as duas em ordem daria resultado arbitrário.
	if op.Op == OpEdit {
		if err := tx.Model(&models.PatientPlanSuggestion{}).
			Where("plan_id = ? AND slide_id = ? AND field_path = ? AND status = ?",
				plan.ID, op.SlideID, op.Path, models.SuggestionPending).
			Update("status", models.SuggestionSuperseded).Error; err != nil {
			return out, err
		}
	}
	if err := tx.Create(&sug).Error; err != nil {
		return out, err
	}
	return toSuggestionDTO(&sug), nil
}

type ResolveSuggestionsInput struct {
	PlanID, PatientID, UserID uuid.UUID
	Action                    string // accept | reject
	SuggestionIDs             []uuid.UUID
	SlideID                   string
	ExpectedRevision          *int
}

// ResolveSuggestions aceita ou recusa sugestões.
//
// Resultado parcial é resposta legítima: uma sugestão cujo `base_hash` não bate mais NÃO é
// aplicada, e é isso que impede o painel de vinte minutos atrás de apagar o que o médico acabou de
// escrever. O que foi pulado volta dito, com o motivo.
func (s *PatientPlanAssistantService) ResolveSuggestions(in ResolveSuggestionsInput) (*dto.PlanResolveResult, error) {
	out := &dto.PlanResolveResult{}
	err := s.db.Transaction(func(tx *gorm.DB) error {
		plan, err := s.plans.load(in.PlanID, in.PatientID)
		if err != nil {
			return err
		}
		if err := CheckExpected(plan, in.ExpectedRevision); err != nil {
			return err
		}

		q := tx.Where("plan_id = ? AND status = ?", in.PlanID, models.SuggestionPending)
		if len(in.SuggestionIDs) > 0 {
			q = q.Where("id IN ?", in.SuggestionIDs)
		} else if in.SlideID != "" {
			q = q.Where("slide_id = ?", in.SlideID)
		}
		var sugestoes []models.PatientPlanSuggestion
		if err := q.Order("created_at").Find(&sugestoes).Error; err != nil {
			return err
		}

		agora := time.Now()
		if in.Action == "reject" {
			for i := range sugestoes {
				sugestoes[i].Status = models.SuggestionRejected
				sugestoes[i].ResolvedByID = &in.UserID
				sugestoes[i].ResolvedAt = &agora
				if err := tx.Save(&sugestoes[i]).Error; err != nil {
					return err
				}
				out.Resolved = append(out.Resolved, sugestoes[i].ID.String())
			}
			out.RevisionSeq = plan.RevisionSeq
			return nil
		}

		var aplicar []PlanOp
		var aplicadas []*models.PatientPlanSuggestion
		for i := range sugestoes {
			sug := &sugestoes[i]
			if sug.BaseHash != "" {
				j := indiceDoSlide(plan.Content, sug.SlideID)
				if j < 0 {
					out.Skipped = append(out.Skipped, dto.PlanSkipped{ID: sug.ID.String(), Reason: "o slide não existe mais"})
					continue
				}
				atual, err := HashDeckContent([]pdfdoc.DeckSlide{plan.Content[j]})
				if err != nil {
					return err
				}
				if atual != sug.BaseHash {
					sug.Status = models.SuggestionStale
					if err := tx.Save(sug).Error; err != nil {
						return err
					}
					out.Skipped = append(out.Skipped, dto.PlanSkipped{
						ID: sug.ID.String(), Reason: "o slide mudou depois que esta sugestão foi criada",
					})
					continue
				}
			}
			op, err := opDaSugestao(sug)
			if err != nil {
				out.Skipped = append(out.Skipped, dto.PlanSkipped{ID: sug.ID.String(), Reason: err.Error()})
				continue
			}
			aplicar = append(aplicar, op)
			aplicadas = append(aplicadas, sug)
		}

		if len(aplicar) == 0 {
			out.RevisionSeq = plan.RevisionSeq
			return nil
		}

		// Uma sugestão de cada vez, como no turno: uma que falhe ao aplicar (ordem incompleta num
		// `reorder`, valor de tipo errado) não pode derrubar as outras aceitas no mesmo lote.
		novos := plan.Content
		var entraram []PlanOp
		var aceitasOK []*models.PatientPlanSuggestion
		for i, op := range aplicar {
			passo, errPasso := ApplyOps(novos, []PlanOp{op})
			if errPasso != nil {
				out.Skipped = append(out.Skipped, dto.PlanSkipped{
					ID: aplicadas[i].ID.String(), Reason: errPasso.Error(),
				})
				continue
			}
			novos = passo
			entraram = append(entraram, op)
			aceitasOK = append(aceitasOK, aplicadas[i])
		}
		aplicar, aplicadas = entraram, aceitasOK
		if len(aplicar) == 0 {
			out.RevisionSeq = plan.RevisionSeq
			return nil
		}
		// Guardado ANTES de trocar o conteúdo: a classificação abaixo pergunta o que o caminho ERA
		// no slide sobre o qual a sugestão foi feita. Classificando contra o conteúdo já aplicado,
		// um `remove` aceito não encontrava mais o slide e virava "desconhecido".
		conteudoAntes := plan.Content
		plan.Content = novos
		plan.Status = models.PatientPlanDraft
		// As ops entram na revisão, e não só o conteúdo resultante.
		//
		// Sem elas, a revisão de aceite não declara caminho nenhum, e a sugestão numérica aceita
		// (o número que a ferramenta escreveu, o conteúdo de maior risco do fluxo inteiro) some da
		// trilha de `ai_touched_paths`. O erro caía do lado errado: um plano cheio de números
		// gerados apareceria como se nenhum tivesse vindo da ferramenta.
		//
		// Sem `decision`: aceitar JÁ é a decisão. O extrator trata a ausência como aplicada.
		triadas := make([]TriagedOp, 0, len(aplicar))
		for i, op := range aplicar {
			classe := FieldUnknown
			if j := indiceDoSlide(conteudoAntes, op.SlideID); j >= 0 && op.Path != "" {
				classe, _ = ClassifyPath(&conteudoAntes[j], op.Path)
			}
			triadas = append(triadas, TriagedOp{Op: op, Class: classe.String(),
				Reason: "sugestão aceita pelo clínico (" + aplicadas[i].Class + ")"})
		}
		seq, err := s.revisions.Record(tx, RecordRevisionInput{
			Plan: plan, Title: plan.Title, Content: plan.Content,
			AuthorKind: models.PlanAuthorAssistant, CreatedByID: in.UserID,
			Reason: models.PlanRevisionSuggestionAccept, Ops: triadas,
		})
		if err != nil {
			return err
		}
		plan.RevisionSeq = seq
		if err := tx.Save(plan).Error; err != nil {
			return err
		}

		var revID *uuid.UUID
		var rev models.PatientPlanRevision
		if err := tx.Where("plan_id = ? AND seq = ?", plan.ID, seq).First(&rev).Error; err == nil {
			revID = &rev.ID
		}
		for _, sug := range aplicadas {
			sug.Status = models.SuggestionAccepted
			sug.ResolvedByID = &in.UserID
			sug.ResolvedAt = &agora
			sug.RevisionID = revID
			if err := tx.Save(sug).Error; err != nil {
				return err
			}
			out.Resolved = append(out.Resolved, sug.ID.String())
		}
		out.RevisionSeq = seq
		out.Plan = toPatientPlanDTO(plan)
		return nil
	})
	return out, err
}

// History devolve a conversa do plano.
func (s *PatientPlanAssistantService) History(planID, patientID uuid.UUID) ([]dto.PlanMessage, error) {
	if _, err := s.plans.load(planID, patientID); err != nil {
		return nil, err
	}
	var rows []models.PatientPlanMessage
	if err := s.db.Where("plan_id = ?", planID).Order("seq").Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]dto.PlanMessage, 0, len(rows))
	for i := range rows {
		out = append(out, dto.PlanMessage{
			ID: rows[i].ID.String(), Seq: rows[i].Seq, Role: rows[i].Role,
			Body: rows[i].Body, Status: rows[i].Status,
			CreatedAt: rows[i].CreatedAt.In(saoPaulo()).Format(time.RFC3339),
		})
	}
	return out, nil
}

// ListSuggestions devolve as sugestões pendentes.
func (s *PatientPlanAssistantService) ListSuggestions(planID, patientID uuid.UUID) ([]dto.PlanSuggestion, error) {
	if _, err := s.plans.load(planID, patientID); err != nil {
		return nil, err
	}
	var rows []models.PatientPlanSuggestion
	if err := s.db.Where("plan_id = ? AND status = ?", planID, models.SuggestionPending).
		Order("created_at").Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]dto.PlanSuggestion, 0, len(rows))
	for i := range rows {
		out = append(out, toSuggestionDTO(&rows[i]))
	}
	return out, nil
}

// ---- auxiliares ----

func (s *PatientPlanAssistantService) proximoSeq(tx *gorm.DB, planID uuid.UUID) (int, error) {
	var max int
	err := tx.Model(&models.PatientPlanMessage{}).Where("plan_id = ?", planID).
		Select("COALESCE(MAX(seq), 0)").Scan(&max).Error
	return max + 1, err
}

func (s *PatientPlanAssistantService) historico(planID uuid.UUID) ([]PlanEditTurn, error) {
	var rows []models.PatientPlanMessage
	if err := s.db.Where("plan_id = ? AND status = ?", planID, "ok").
		Order("seq").Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]PlanEditTurn, 0, len(rows))
	for _, r := range rows {
		if strings.TrimSpace(r.Body) == "" {
			continue
		}
		out = append(out, PlanEditTurn{Role: r.Role, Body: r.Body})
	}
	return out, nil
}

func converteOps(ops []PlanModelOp) []PlanOp {
	out := make([]PlanOp, 0, len(ops))
	for _, o := range ops {
		p := PlanOp{
			Op: PlanOpKind(o.Op), SlideID: o.SlideID, AfterSlideID: o.AfterSlideID,
			Path: o.Path, Order: o.Order,
		}
		if len(o.Value) > 0 {
			var v any
			if err := json.Unmarshal(o.Value, &v); err == nil {
				p.Value = v
			}
		}
		if len(o.Slide) > 0 {
			var s pdfdoc.DeckSlide
			if err := json.Unmarshal(o.Slide, &s); err == nil {
				p.Slide = &s
			}
		}
		out = append(out, p)
	}
	return out
}

// opDaSugestao reconstrói a operação a partir da linha gravada.
//
// `new_value` guarda o que a sugestão introduz, e o que isso É depende do verbo: valor de campo
// para `edit`, o slide inteiro para `add`, a ordem completa para `reorder`. Ler tudo como
// `op.Value` era o que tornava as estruturais inaceitáveis.
func opDaSugestao(s *models.PatientPlanSuggestion) (PlanOp, error) {
	op := PlanOp{
		Op: PlanOpKind(s.Op), SlideID: s.SlideID, AfterSlideID: s.AfterSlideID, Path: s.FieldPath,
	}
	if len(s.NewValue) == 0 {
		return op, nil
	}
	switch op.Op {
	case OpAdd:
		var slide pdfdoc.DeckSlide
		if err := json.Unmarshal(s.NewValue, &slide); err != nil {
			return op, errors.New("slide da sugestão ilegível")
		}
		op.Slide = &slide
	case OpReorder:
		var ordem []string
		if err := json.Unmarshal(s.NewValue, &ordem); err != nil {
			return op, errors.New("ordem da sugestão ilegível")
		}
		op.Order = ordem
	default:
		var v any
		if err := json.Unmarshal(s.NewValue, &v); err != nil {
			return op, errors.New("valor da sugestão ilegível")
		}
		op.Value = v
	}
	return op, nil
}

func toSuggestionDTO(s *models.PatientPlanSuggestion) dto.PlanSuggestion {
	out := dto.PlanSuggestion{
		ID: s.ID.String(), Op: s.Op, SlideID: s.SlideID, Path: s.FieldPath,
		Class: s.Class, Rationale: s.Rationale, Status: string(s.Status),
		CreatedAt: s.CreatedAt.In(saoPaulo()).Format(time.RFC3339),
	}
	_ = json.Unmarshal(s.OldValue, &out.OldValue)
	_ = json.Unmarshal(s.NewValue, &out.NewValue)
	_ = json.Unmarshal(s.Provenance, &out.Provenance)
	return out
}

// podaDossieParaPrompt reduz o dossiê ao que cabe num prompt sem perder o que decide.
//
// Um dossiê tem dezenas de réguas e passa fácil de cem mil tokens. Mandar tudo é caro e piora a
// atenção do modelo. A poda manda: as réguas COMPLETAS dos exames já citados no deck, e um catálogo
// de uma linha para todo o resto — assim o modelo sabe o que existe e pode pedir, sem carregar o
// histórico inteiro de cada exame.
func podaDossieParaPrompt(d *dto.PlanDossierResponse, slides []pdfdoc.DeckSlide) (string, error) {
	if d == nil {
		return "{}", nil
	}
	citados := map[string]bool{}
	for code := range codigosCitados(slides) {
		citados[code] = true
	}

	completas := map[string]dto.PlanRuler{}
	var catalogo []string
	for code, r := range d.Rulers {
		if citados[code] {
			completas[code] = r
			continue
		}
		ultimo := ""
		if n := len(r.History); n > 0 {
			ultimo = r.History[n-1].Text + " " + r.Unit + " em " + r.History[n-1].Date
		}
		catalogo = append(catalogo, fmt.Sprintf("%s | %s | %s", code, r.Name, ultimo))
	}
	sort.Strings(catalogo)

	// O nome do paciente não vai para o modelo: ele não é necessário para escrever a devolutiva, e
	// o que não precisa sair do prontuário não sai.
	podado := map[string]any{
		"paciente":         map[string]any{"sexo": d.Patient.Gender, "idade": d.Patient.Age},
		"reguasCitadas":    completas,
		"catalogoDeExames": catalogo,
		"seMovendo":        primeiros(d.Moving, 15),
		"estaBem":          primeiros(d.Strong, 10),
		"vitais":           d.Vitals,
		"condutas":         d.CarePlan,
	}
	if d.Snapshot != nil {
		podado["escore"] = d.Snapshot.TotalPercentage
	}
	b, err := json.Marshal(podado)
	return string(b), err
}

func primeiros(fs []dto.PlanFinding, n int) []dto.PlanFinding {
	if len(fs) <= n {
		return fs
	}
	return fs[:n]
}

// ---------------------------------------------------------------------------
// Geração do rascunho.

// GenerateDraftInput — o pedido de "escreva o rascunho deste paciente".
type GenerateDraftInput struct {
	PatientID uuid.UUID
	UserID    uuid.UUID
	Title     string
	// Instruction é opcional: o médico pode dirigir ("foque no ferro e no sono").
	Instruction string
}

// GenerateDraftResult — o plano criado, o que o modelo explicou e o que o servidor NÃO conseguiu
// confirmar.
type GenerateDraftResult struct {
	Plan     *dto.PatientPlanResponse `json:"plan"`
	Reply    string                   `json:"reply"`
	Warnings []dto.PlanGenWarning     `json:"warnings,omitempty"`
	Overflow []pdfdoc.DeckOverflow    `json:"overflow,omitempty"`
	Model    string                   `json:"model,omitempty"`
}

// GenerateDraft escreve o rascunho inteiro a partir do prontuário compilado.
//
// É o passo que faltava na feature: até aqui "novo plano" nascia com dois slides vazios e a
// conversa editava um documento que ninguém tinha escrito.
//
// A ordem importa e não é arbitrária:
//  1. cria o plano vazio — o dossiê é congelado POR PLANO, então precisa de um id;
//  2. congela o dossiê — é contra ele que todo número vai ser conferido, e ele tem que ser o mesmo
//     que a conversa vai usar depois (byte-idêntico, para o cache de prompt valer);
//  3. gera;
//  4. CONFERE cada número de cada slide contra o índice numérico do dossiê;
//  5. grava como revisão do assistente, não do médico: quem escreveu foi a ferramenta, e
//     `ai_touched_paths` na publicação depende disso ser verdade.
//
// O que a conferência NÃO faz, e está dito na tela: ela prova que o número EXISTE no dossiê, nunca
// que ele significa o que a frase diz. "Sua ferritina está em 96" e "seu colesterol está em 96"
// passam idênticas se 96 existir em algum lugar.
func (s *PatientPlanAssistantService) GenerateDraft(in GenerateDraftInput) (*GenerateDraftResult, error) {
	titulo := strings.TrimSpace(in.Title)
	if titulo == "" {
		titulo = "Seus exames"
	}

	// 1. o plano vazio, para haver id.
	criado, err := s.plans.Create(in.PatientID, in.UserID, &dto.SavePatientPlanRequest{
		Title: titulo, Content: []pdfdoc.DeckSlide{},
	})
	if err != nil {
		return nil, err
	}
	planID, err := uuid.Parse(criado.ID)
	if err != nil {
		return nil, err
	}

	// 2. o dossiê congelado.
	if err := s.db.Transaction(func(tx *gorm.DB) error {
		_, errF := s.dossiers.Freeze(tx, planID, in.PatientID, &in.UserID)
		return errF
	}); err != nil {
		return nil, err
	}
	dossie, dossieRow, err := s.dossiers.Current(planID)
	if err != nil {
		return nil, err
	}
	if dossieVazio(dossie) {
		return nil, errors.New("este paciente não tem exame nem anamnese suficientes para uma devolutiva")
	}

	// A poda com `nil` de slides devolve o dossiê inteiro: na geração ainda não há deck para dizer
	// quais códigos são citados.
	dossieJSON, err := podaDossieParaPrompt(dossie, nil)
	if err != nil {
		return nil, err
	}

	// 3. gera.
	res, meta, err := s.ai.GeneratePatientPlan(PlanGenerateRequest{
		DossierJSON:   dossieJSON,
		Instruction:   in.Instruction,
		PromptVersion: s.promptVersion,
		Model:         s.model,
	})
	if err != nil {
		return nil, err
	}

	// Os ids que o modelo inventa ("bem-1", "capa") são descartados: id de slide é opaco e é o
	// alvo de toda operação e sugestão depois. `EnsureSlideIDs` só preenche os VAZIOS, então o
	// zeramento tem que ser explícito.
	for i := range res.Slides {
		res.Slides[i].ID = ""
	}
	slides, _ := EnsureSlideIDs(res.Slides)

	// 3b. a escala e o histórico de cada régua vêm do dossiê, não do modelo.
	slides, avisos := hidrataReguas(slides, dossie)

	// 4. confere os números slide a slide.
	indice := BuildNumericIndex(dossie)
	for i := range slides {
		for _, p := range provaDoSlide(slides[i], indice) {
			if p.Found {
				continue
			}
			avisos = append(avisos, dto.PlanGenWarning{
				SlideIndex: i + 1,
				SlideID:    slides[i].ID,
				Title:      slides[i].Title,
				Numeral:    p.Numeral,
				Reason:     "este número não foi encontrado no prontuário compilado",
			})
		}
	}

	// 5. grava o conteúdo como revisão DO ASSISTENTE.
	plan, err := s.plans.load(planID, in.PatientID)
	if err != nil {
		return nil, err
	}
	if err := s.db.Transaction(func(tx *gorm.DB) error {
		plan.Content = slides
		plan.Status = models.PatientPlanDraft
		seq, errR := s.revisions.Record(tx, RecordRevisionInput{
			Plan: plan, Title: plan.Title, Content: plan.Content,
			AuthorKind: models.PlanAuthorAssistant, CreatedByID: in.UserID,
			Reason:    models.PlanRevisionAIApply,
			DossierID: &dossieRow.ID,
			AIModel:   meta.Model, AIPromptVersion: s.promptVersion,
		})
		if errR != nil {
			return errR
		}
		plan.RevisionSeq = seq
		return tx.Save(plan).Error
	}); err != nil {
		return nil, err
	}

	// A conferência geométrica no fim: o deck acabou de nascer e o médico precisa saber ANTES de
	// editar se algum slide já nasceu estourando.
	estouro, _ := pdfdoc.CheckDeckOverflow(pdfdoc.Deck{Title: plan.Title, Slides: plan.Content})

	final, err := s.plans.Get(planID, in.PatientID)
	if err != nil {
		return nil, err
	}
	return &GenerateDraftResult{
		Plan: final, Reply: res.Reply, Warnings: avisos, Overflow: estouro, Model: meta.Model,
	}, nil
}

// dossieVazio — paciente sem exame e sem anamnese não tem devolutiva a montar, e gerar em cima do
// vazio produziria um deck inventado. Melhor recusar com a razão.
func dossieVazio(d *dto.PlanDossierResponse) bool {
	if d == nil {
		return true
	}
	return len(d.Rulers) == 0 && len(d.Strong) == 0 && len(d.Moving) == 0
}
