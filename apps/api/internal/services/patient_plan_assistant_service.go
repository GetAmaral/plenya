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

		if len(aplicar) > 0 {
			novos, err := ApplyOps(plan.Content, aplicar)
			if err != nil {
				// Uma operação que passou na triagem e falha ao aplicar é defeito nosso, não do
				// modelo: aborta o turno inteiro em vez de gravar meio deck.
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

	novoValor, _ := json.Marshal(op.Value)
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

		novos, err := ApplyOps(plan.Content, aplicar)
		if err != nil {
			return err
		}
		plan.Content = novos
		plan.Status = models.PatientPlanDraft
		seq, err := s.revisions.Record(tx, RecordRevisionInput{
			Plan: plan, Title: plan.Title, Content: plan.Content,
			AuthorKind: models.PlanAuthorAssistant, CreatedByID: in.UserID,
			Reason: models.PlanRevisionSuggestionAccept,
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

func opDaSugestao(s *models.PatientPlanSuggestion) (PlanOp, error) {
	op := PlanOp{
		Op: PlanOpKind(s.Op), SlideID: s.SlideID, AfterSlideID: s.AfterSlideID, Path: s.FieldPath,
	}
	if len(s.NewValue) > 0 {
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
