package services

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// ============================================================
// AI features pra Central de Conversas — resumir + sugerir resposta
// ============================================================
//
// Endpoints (handlers em conversation_handler.go):
//   POST /conversations/:type/:id/ai/summary       → SummarizeConversation
//   POST /conversations/:type/:id/ai/suggest-reply → SuggestReply
//
// Estratégia de prompt:
//   - Pega últimas N activities (50 pra resumo, 10 pra sugestão).
//   - Trunca Content individual em MaxContentChars pra não estourar contexto.
//   - Prefixo [DENTRO]/[FORA] indica direção (inbound = "dentro de casa" do cliente
//     pra Plenya; outbound = "fora", Plenya respondendo). Compacto e PT-BR.
//
// LGPD: prompts contêm conteúdo sensível, então usamos o AIService.CompleteText
// que NÃO loga conteúdo. Aqui também não logamos — só metadata (owner type,
// total chars do prompt).

// ErrAIConversationEmpty → conversa sem mensagens elegíveis pra IA.
var ErrAIConversationEmpty = errors.New("ai: conversa sem mensagens para análise")

// AISummaryResult devolve resumo + timestamp de geração + flag de cache hit.
type AISummaryResult struct {
	Summary     string    `json:"summary"`
	GeneratedAt time.Time `json:"generatedAt"`
	Cached      bool      `json:"cached"`
}

// AISuggestionResult devolve texto sugerido + model usado (transparência pro vendedor).
type AISuggestionResult struct {
	Suggestion string `json:"suggestion"`
	Model      string `json:"model"`
}

// ReceptionReplyResult é a saída estruturada do recepcionista virtual. Reusada tanto pelo
// modo Copiloto (humano revisa o Reply e envia) quanto pelo Automático (Fase 2).
type ReceptionReplyResult struct {
	Reply         string `json:"reply"`
	Action        string `json:"action"`        // ask|answer|handle_objection|propose_schedule|handoff
	HandoffReason string `json:"handoffReason"` // preenchido quando Action == handoff
	DiscloseAI    bool   `json:"discloseAI"`    // true quando a resposta deve identificar-se como assistente
	Model         string `json:"model"`
}

// Constantes pra controle de contexto. Calibrados pra ficar < 8k tokens em pt-BR
// (1 token ≈ 3.5 chars em PT) com folga: 50 msgs * 500 chars = 25k chars ≈ 7k tokens.
const (
	aiSummaryMaxMessages    = 50
	aiSuggestionMaxMessages = 10
	// Janela curta do recepcionista (memória episódica). Subida de 14→40 para que a Lívia
	// "lembre" de toda a conversa recente e nunca repita pergunta já respondida. O clínico
	// mencionado fica SÓ aqui (curto prazo), nunca no rolling_summary (§3.1 do plano).
	// 40 msgs * 500 chars ≈ 20k chars ≈ 6k tokens; com o cérebro (~3k) cabe com folga.
	aiReceptionMaxMessages = 40
	aiMaxContentChars      = 500

	// Memória de longo prazo (social): regenera o rolling_summary a cada +5 mensagens novas
	// (ou no fim do atendimento), lendo até este teto de mensagens.
	relationshipSummaryEveryMessages = 5
	relationshipSummaryMaxMessages   = 60
	aiModelRelationship              = aiModelSummary // Haiku: resumo barato e rápido

	// Modelos. Haiku pra resumo (barato + rápido); Sonnet pra escrita (qualidade
	// do tom Plenya). Atualize aqui quando subir versão (claude-md doc).
	aiModelSummary    = "claude-haiku-4-5-20251001"
	aiModelSuggestion = "claude-sonnet-4-6"

	// Cache: resumo idêntico (mesmas msgs) reusado por 1h. Refresh manual via ?force=true.
	aiSummaryCacheTTL = 1 * time.Hour
)

// SummarizeConversation gera (ou reusa do cache) resumo da conversa.
//
// Cache: salva resultado como LeadActivity{type=note_added, channel=internal, metadata
// {ai_summary:true, prompt_hash, generated_at}}. Reuso quando hash bate E < TTL.
// force=true pula o cache.
func (s *ConversationService) SummarizeConversation(
	ctx context.Context,
	userID uuid.UUID,
	ownerType string,
	ownerID uuid.UUID,
	force bool,
) (*AISummaryResult, error) {
	if s.aiService == nil {
		return nil, fmt.Errorf("conversation: ai service não configurado")
	}
	if !isValidOwnerType(ownerType) || ownerID == uuid.Nil {
		return nil, ErrConversationOwnerInvalid
	}

	activities, err := s.loadConversationActivities(ctx, ownerType, ownerID, aiSummaryMaxMessages)
	if err != nil {
		return nil, err
	}
	if len(activities) == 0 {
		return nil, ErrAIConversationEmpty
	}

	transcript := buildTranscript(activities)
	hash := promptHash("summary", transcript)

	// Cache lookup: última activity de resumo desse owner. Se hash bate e < TTL, reusa.
	if !force {
		cached, ok := s.lookupAISummaryCache(ctx, ownerType, ownerID, hash)
		if ok {
			return cached, nil
		}
	}

	prompt := buildSummaryPrompt(transcript)
	summary, err := s.aiService.CompleteText(ctx, prompt, CompleteTextOptions{
		Model:       aiModelSummary,
		MaxTokens:   600,
		Temperature: 0.3,
		Timeout:     20 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	res := &AISummaryResult{Summary: summary, GeneratedAt: now, Cached: false}

	// Persiste cache (best-effort — falha não derruba o fluxo).
	s.persistAISummaryCache(ctx, userID, ownerType, ownerID, summary, hash, now)

	return res, nil
}

// SuggestReply sugere texto de resposta no tom Plenya. Não persiste cache (sugestão é
// efêmera; vendedor pode regenerar quantas vezes quiser, sempre quer fresca).
//
// Restringido a canal email no MVP (UI só mostra botão pra email). Backend não bloqueia
// por canal — se chamado pra WhatsApp, gera mesmo assim com tom adequado.
func (s *ConversationService) SuggestReply(
	ctx context.Context,
	ownerType string,
	ownerID uuid.UUID,
	intent string,
) (*AISuggestionResult, error) {
	if s.aiService == nil {
		return nil, fmt.Errorf("conversation: ai service não configurado")
	}
	if !isValidOwnerType(ownerType) || ownerID == uuid.Nil {
		return nil, ErrConversationOwnerInvalid
	}

	activities, err := s.loadConversationActivities(ctx, ownerType, ownerID, aiSuggestionMaxMessages)
	if err != nil {
		return nil, err
	}
	if len(activities) == 0 {
		return nil, ErrAIConversationEmpty
	}

	transcript := buildTranscript(activities)
	prompt := buildSuggestionPrompt(transcript, strings.TrimSpace(intent))

	suggestion, err := s.aiService.CompleteText(ctx, prompt, CompleteTextOptions{
		Model:       aiModelSuggestion,
		MaxTokens:   1024,
		Temperature: 0.7,
		Timeout:     20 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	return &AISuggestionResult{
		Suggestion: suggestion,
		Model:      aiModelSuggestion,
	}, nil
}

// GenerateReceptionReply gera a próxima mensagem do recepcionista virtual, ancorada no
// cérebro (reception_brain.go: script + objeções + guardrails). Saída estruturada com a
// ação sugerida e se deve identificar-se como IA. Não persiste (efêmero no Copiloto; no
// Automático quem persiste é o job ao enviar).
func (s *ConversationService) GenerateReceptionReply(
	ctx context.Context,
	ownerType string,
	ownerID uuid.UUID,
) (*ReceptionReplyResult, error) {
	if s.aiService == nil {
		return nil, fmt.Errorf("conversation: ai service não configurado")
	}
	if !isValidOwnerType(ownerType) || ownerID == uuid.Nil {
		return nil, ErrConversationOwnerInvalid
	}

	activities, err := s.loadConversationActivities(ctx, ownerType, ownerID, aiReceptionMaxMessages)
	if err != nil {
		return nil, err
	}
	if len(activities) == 0 {
		return nil, ErrAIConversationEmpty
	}

	transcript := buildTranscript(activities)
	slotsText := ""
	if s.receptionSlots != nil {
		slotsText = s.receptionSlots(ctx)
	}
	bizHours := ""
	if s.receptionBusinessHours != nil {
		bizHours = s.receptionBusinessHours(ctx)
	}
	memory := s.buildReceptionMemory(ctx, ownerType, ownerID)
	prompt := buildReceptionPrompt(transcript, slotsText, bizHours, receptionNowLine(), memory)

	raw, err := s.aiService.CompleteText(ctx, prompt, CompleteTextOptions{
		Model:       aiModelSuggestion,
		MaxTokens:   900,
		Temperature: 0.5,
		Timeout:     20 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	res := parseReceptionReply(raw)
	res.Reply = sanitizeReceptionVoice(res.Reply)
	res.Model = aiModelSuggestion
	if strings.TrimSpace(res.Reply) == "" {
		return nil, ErrAIConversationEmpty
	}
	return res, nil
}

// sanitizeReceptionVoice remove maneirismos de IA que o modelo às vezes insere apesar do
// prompt — sobretudo o travessão (AI-tell em 2026, banido na voz Plenya). Troca por vírgula.
func sanitizeReceptionVoice(s string) string {
	out := s
	// Travessão usado como pontuação (com espaços) → vírgula. Inclui o em-dash colado
	// (AI-tell típico). Não mexe no en-dash colado a números (ex: "9–10h", "18,5–24,9").
	out = strings.ReplaceAll(out, " — ", ", ")
	out = strings.ReplaceAll(out, " – ", ", ")
	out = strings.ReplaceAll(out, "—", ", ")
	out = strings.ReplaceAll(out, " ,", ",")
	out = strings.ReplaceAll(out, ",,", ",")
	out = strings.ReplaceAll(out, "  ", " ")
	return strings.TrimSpace(out)
}

// parseReceptionReply extrai o JSON da resposta do modelo de forma tolerante (lida com
// markdown ou texto antes/depois). Se o parse falhar, trata o texto inteiro como Reply.
func parseReceptionReply(raw string) *ReceptionReplyResult {
	res := &ReceptionReplyResult{Action: "answer"}
	trimmed := strings.TrimSpace(raw)

	start := strings.Index(trimmed, "{")
	end := strings.LastIndex(trimmed, "}")
	if start >= 0 && end > start {
		var parsed ReceptionReplyResult
		if err := json.Unmarshal([]byte(trimmed[start:end+1]), &parsed); err == nil && strings.TrimSpace(parsed.Reply) != "" {
			if parsed.Action == "" {
				parsed.Action = "answer"
			}
			return &parsed
		}
	}

	// Fallback: modelo não devolveu JSON — usa o texto cru como resposta.
	res.Reply = trimmed
	return res
}

// loadConversationActivities lê as últimas N activities elegíveis (channel email|whatsapp,
// type message_sent|received) em ordem cronológica ASC.
func (s *ConversationService) loadConversationActivities(
	ctx context.Context,
	ownerType string,
	ownerID uuid.UUID,
	limit int,
) ([]models.LeadActivity, error) {
	q := s.db.WithContext(ctx).Model(&models.LeadActivity{}).
		Where("channel IN (?, ?)", models.LeadChannelEmail, models.LeadChannelWhatsApp).
		Where("type IN (?, ?)", models.LeadActivityMessageSent, models.LeadActivityMessageReceived)

	switch ownerType {
	case string(models.ConversationOwnerLead):
		q = q.Where("lead_id = ?", ownerID)
	case string(models.ConversationOwnerPatient):
		q = q.Where("patient_id = ?", ownerID)
	}

	// DESC + Limit pra pegar últimas N; depois reverte pra ASC pro prompt.
	var activities []models.LeadActivity
	if err := q.Order("created_at DESC").Limit(limit).Find(&activities).Error; err != nil {
		return nil, fmt.Errorf("conversation: ai load activities: %w", err)
	}
	// Reversa in-place pra ASC (cronológico).
	for i, j := 0, len(activities)-1; i < j; i, j = i+1, j-1 {
		activities[i], activities[j] = activities[j], activities[i]
	}
	return activities, nil
}

// buildTranscript monta texto pro prompt: "[DENTRO 24/04 14:30 email] conteúdo".
// Trunca cada content em aiMaxContentChars pra controlar tamanho do contexto.
func buildTranscript(activities []models.LeadActivity) string {
	var sb strings.Builder
	for _, a := range activities {
		direction := "DENTRO"
		if a.ActorUserID != nil {
			direction = "FORA"
		}
		content := ""
		if a.Content != nil {
			content = strings.TrimSpace(*a.Content)
		}
		if content == "" {
			continue
		}
		if len(content) > aiMaxContentChars {
			content = content[:aiMaxContentChars] + "…"
		}
		// Normaliza newlines internos pra evitar layout quebrado.
		content = strings.ReplaceAll(content, "\r\n", "\n")
		fmt.Fprintf(&sb, "[%s %s %s] %s\n",
			direction,
			a.CreatedAt.In(brLocation).Format("02/01 15:04"), // horário de Londrina/SP, não UTC
			string(a.Channel),
			content,
		)
	}
	return sb.String()
}

// receptionNowLine devolve o "agora" em horário de Londrina/SP (o container roda em UTC; sem
// isto a IA deduz a hora pelos carimbos do histórico e erra a saudação). Ex.: "sábado, 06/06 03:51".
func receptionNowLine() string {
	t := time.Now().In(brLocation)
	return fmt.Sprintf("%s, %s", ptWeekday[t.Weekday()], t.Format("02/01 15:04"))
}

// buildSummaryPrompt monta prompt pra resumo executivo em PT-BR.
func buildSummaryPrompt(transcript string) string {
	return fmt.Sprintf(`Você é assistente de uma clínica médica brasileira premium chamada Plenya.

Resuma a conversa abaixo em 3 a 5 bullets curtos em português, focando em:
- Necessidade ou interesse do cliente
- Última pendência aberta (dúvida, objeção, ou aguardo)
- Próximo passo sugerido para a equipe Plenya

A conversa segue ordem cronológica. Cada linha começa com [DENTRO] (mensagem do cliente) ou [FORA] (mensagem da Plenya), seguido de data/hora e canal.

CONVERSA:
%s

Responda APENAS com os bullets, sem cabeçalho ("Resumo:" etc.) e sem comentários extras. Cada bullet começa com "- " e tem no máximo 25 palavras.`, transcript)
}

// buildSuggestionPrompt monta prompt pra resposta. intent opcional permite ao vendedor
// dar pista ("agendar consulta", "recusar educadamente"). Quando vazio, IA infere.
func buildSuggestionPrompt(transcript, intent string) string {
	intentLine := "Sugira uma resposta apropriada para a última mensagem do cliente."
	if intent != "" {
		// Truncate intent pra evitar prompt injection brutal.
		if len(intent) > 200 {
			intent = intent[:200]
		}
		intentLine = fmt.Sprintf("Intenção do vendedor: %s", intent)
	}

	return fmt.Sprintf(`Você é vendedor da Plenya, clínica médica brasileira premium focada em longevidade, performance e estilo de vida saudável.

Tom: caloroso, profissional, direto. Português brasileiro do dia a dia. Sem jargões médicos quando falando com cliente leigo. Sem clichês comerciais ("Em primeiro lugar, gostaríamos de agradecer...").

Histórico recente da conversa (cronológico, [DENTRO] = cliente, [FORA] = Plenya):

%s

%s

REGRAS:
- Escreva APENAS o corpo da resposta.
- Sem assinatura (sem "Atenciosamente", sem nome).
- Sem cabeçalho (sem "Olá [Nome]", sem "Prezado").
- Sem placeholders entre colchetes ([Nome], [Data]).
- Máximo 4 parágrafos curtos.
- Se o canal mais recente for WhatsApp, mantenha tom mais informal e enxuto (1-2 parágrafos).`, transcript, intentLine)
}

// promptHash gera SHA-256 hex do prefixo+conteúdo. Usado pra cache idempotente: mesma
// conversa → mesmo hash → reusa resumo se ainda quente.
func promptHash(prefix, content string) string {
	h := sha256.Sum256([]byte(prefix + "|" + content))
	return hex.EncodeToString(h[:])
}

// lookupAISummaryCache busca a activity de cache mais recente que bate hash + TTL.
// Retorna (resultado, true) se hit; (nil, false) caso contrário.
func (s *ConversationService) lookupAISummaryCache(
	ctx context.Context,
	ownerType string,
	ownerID uuid.UUID,
	hash string,
) (*AISummaryResult, bool) {
	q := s.db.WithContext(ctx).Model(&models.LeadActivity{}).
		Where("channel = ?", models.LeadChannelInternal).
		Where("type = ?", models.LeadActivityNoteAdded).
		Where("metadata ->> 'ai_summary' = ?", "true").
		Where("metadata ->> 'prompt_hash' = ?", hash).
		Where("created_at > ?", time.Now().UTC().Add(-aiSummaryCacheTTL))

	switch ownerType {
	case string(models.ConversationOwnerLead):
		q = q.Where("lead_id = ?", ownerID)
	case string(models.ConversationOwnerPatient):
		q = q.Where("patient_id = ?", ownerID)
	}

	var act models.LeadActivity
	if err := q.Order("created_at DESC").First(&act).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			// LGPD: não logar conteúdo. Só sinaliza erro estrutural.
			fmt.Printf("⚠️  ai cache lookup error: %v\n", err)
		}
		return nil, false
	}
	if act.Content == nil {
		return nil, false
	}
	return &AISummaryResult{
		Summary:     *act.Content,
		GeneratedAt: act.CreatedAt,
		Cached:      true,
	}, true
}

// persistAISummaryCache grava o resumo como LeadActivity internal/note_added pra
// reuso futuro. Best-effort — erros logam mas não falham o request.
func (s *ConversationService) persistAISummaryCache(
	ctx context.Context,
	userID uuid.UUID,
	ownerType string,
	ownerID uuid.UUID,
	summary, hash string,
	generatedAt time.Time,
) {
	meta := map[string]any{
		"ai_summary":   true,
		"prompt_hash":  hash,
		"generated_at": generatedAt,
		"model":        aiModelSummary,
	}
	metaJSON, err := json.Marshal(meta)
	if err != nil {
		fmt.Printf("⚠️  ai cache marshal: %v\n", err)
		return
	}

	activity := &models.LeadActivity{
		Type:        models.LeadActivityNoteAdded,
		Channel:     models.LeadChannelInternal,
		Content:     &summary,
		Metadata:    datatypes.JSON(metaJSON),
		ActorUserID: &userID,
	}
	switch ownerType {
	case string(models.ConversationOwnerLead):
		id := ownerID
		activity.LeadID = &id
	case string(models.ConversationOwnerPatient):
		id := ownerID
		activity.PatientID = &id
	}

	if err := s.db.WithContext(ctx).Create(activity).Error; err != nil {
		// LGPD: não loga summary. Só estrutura do erro.
		fmt.Printf("⚠️  ai cache persist: %v\n", err)
	}
}

// ============================================================
// Memória de longo prazo (SOCIAL) da Lívia — relationship_profiles
// ============================================================
//
// Duas camadas no prompt: (1) janela curta = últimas N mensagens (memória episódica, pode
// conter menção clínica só pra não repetir pergunta na MESMA conversa); (2) rolling_summary =
// resumo rolante SOCIAL, persistente entre conversas. O clínico nunca entra no rolling_summary
// (§3.1 do plano). Custo: resumo + janela em vez de history bruto.

// buildReceptionMemory monta o bloco de memória injetado no prompt: flags derivadas (lead vs
// paciente) + resumo rolante social. Best-effort: "" se não houver nada útil ainda.
func (s *ConversationService) buildReceptionMemory(ctx context.Context, ownerType string, ownerID uuid.UUID) string {
	var sb strings.Builder

	// Flag derivada básica (Fase A): já é paciente ou ainda é lead. Continuum/frequente: Fase D.
	switch ownerType {
	case string(models.ConversationOwnerPatient):
		sb.WriteString("RELAÇÃO: esta pessoa já é paciente da Plenya (não trate como contato novo).\n")
	case string(models.ConversationOwnerLead):
		sb.WriteString("RELAÇÃO: ainda é um lead (não é paciente).\n")
	}

	// Fatos sociais conhecidos (memória semântica): lista compacta key: value por categoria.
	facts, _ := NewRelationshipFactService(s.db).ListActive(ctx, ownerType, ownerID)
	if len(facts) > 0 {
		sb.WriteString("FATOS CONHECIDOS (já registrados; não pergunte de novo o que já está aqui):\n")
		for _, f := range facts {
			sb.WriteString("- ")
			sb.WriteString(factLabel(f.Key))
			sb.WriteString(": ")
			sb.WriteString(f.Value)
			sb.WriteString("\n")
		}
	}

	rp := NewRelationshipProfileService(s.db)
	prof, err := rp.Get(ctx, ownerType, ownerID)
	if err == nil && prof != nil {
		if sum := strings.TrimSpace(prof.RollingSummary); sum != "" {
			sb.WriteString("O QUE JÁ SE SABE (resumo social acumulado de conversas anteriores; use para não repetir perguntas e personalizar, sem citar que tem um \"resumo\"):\n")
			sb.WriteString(sum)
			sb.WriteString("\n")
		}
	}
	return strings.TrimSpace(sb.String())
}

// MaintainRelationshipMemory regenera (best-effort) a memória SOCIAL de longo prazo da pessoa:
// o rolling_summary E os fatos atômicos (ADD/UPDATE/DELETE), numa única chamada de IA. Roda
// quando acumulou >= 5 mensagens novas (force=false) ou sempre que houver mensagens não resumidas
// (force=true, ex. fim do atendimento). Guardrail §3.1: SÓ social/relacional; nada clínico entra
// (prompt + backstop factLooksClinical). Não falha o fluxo do job: erros apenas retornam.
func (s *ConversationService) MaintainRelationshipMemory(ctx context.Context, ownerType string, ownerID uuid.UUID, force bool) {
	if s.aiService == nil || !isValidOwnerType(ownerType) || ownerID == uuid.Nil {
		return
	}
	total := s.countConversationMessages(ctx, ownerType, ownerID)
	if total == 0 {
		return
	}

	rp := NewRelationshipProfileService(s.db)
	prof, err := rp.GetOrCreate(ctx, ownerType, ownerID)
	if err != nil || prof == nil {
		return
	}

	newMsgs := total - prof.SummaryMsgCount
	if newMsgs <= 0 {
		return // nada novo desde o último resumo
	}
	if !force && newMsgs < relationshipSummaryEveryMessages {
		return // ainda não acumulou o gatilho de 5 mensagens
	}

	activities, err := s.loadConversationActivities(ctx, ownerType, ownerID, relationshipSummaryMaxMessages)
	if err != nil || len(activities) == 0 {
		return
	}
	transcript := buildTranscript(activities)

	factSvc := NewRelationshipFactService(s.db)
	current, _ := factSvc.ListActive(ctx, ownerType, ownerID)

	prompt := buildRelationshipMemoryPrompt(transcript, prof.RollingSummary, current)
	raw, err := s.aiService.CompleteText(ctx, prompt, CompleteTextOptions{
		Model:       aiModelRelationship,
		MaxTokens:   700,
		Temperature: 0.2,
		Timeout:     25 * time.Second,
	})
	if err != nil {
		return
	}
	ext := parseRelationshipExtraction(raw)

	// 1) Resumo social.
	summary := sanitizeRelationshipSummary(ext.Summary)
	if err := rp.SaveSummary(ctx, prof.ID, summary, total); err != nil {
		fmt.Printf("⚠️  relationship summary persist: %v\n", err)
	}

	// 2) Fatos sociais (ADD/UPDATE/DELETE) com backstop anti-clínico.
	conf := float32(0.7)
	for _, op := range ext.Facts {
		key := normalizeFactKey(op.Key)
		val := strings.TrimSpace(op.Value)
		cat := normalizeFactCategory(op.Category)
		switch strings.ToUpper(strings.TrimSpace(op.Op)) {
		case "ADD", "UPDATE":
			if key == "" || val == "" || factLooksClinical(key, val) {
				continue
			}
			_, _ = factSvc.SetFact(ctx, ownerType, ownerID, cat, key, val, models.RelationshipSourceAI, nil, &conf)
		case "DELETE":
			if key != "" {
				_ = factSvc.CloseFact(ctx, ownerType, ownerID, key)
			}
		}
	}
}

// sanitizeRelationshipSummary normaliza a saída do resumo. Se o modelo sinalizar que não há
// nada social a guardar, devolve "" (não polui o perfil com ruído).
func sanitizeRelationshipSummary(s string) string {
	out := strings.TrimSpace(s)
	upper := strings.ToUpper(out)
	if out == "" || upper == "NADA" || upper == "NENHUM" || strings.HasPrefix(upper, "NADA ") {
		return ""
	}
	return out
}

// countConversationMessages conta as activities elegíveis (mesmo filtro de loadConversationActivities)
// para decidir o gatilho de regeneração do resumo.
func (s *ConversationService) countConversationMessages(ctx context.Context, ownerType string, ownerID uuid.UUID) int {
	q := s.db.WithContext(ctx).Model(&models.LeadActivity{}).
		Where("channel IN (?, ?)", models.LeadChannelEmail, models.LeadChannelWhatsApp).
		Where("type IN (?, ?)", models.LeadActivityMessageSent, models.LeadActivityMessageReceived)
	switch ownerType {
	case string(models.ConversationOwnerLead):
		q = q.Where("lead_id = ?", ownerID)
	case string(models.ConversationOwnerPatient):
		q = q.Where("patient_id = ?", ownerID)
	}
	var n int64
	if err := q.Count(&n).Error; err != nil {
		return 0
	}
	return int(n)
}

// relationshipExtraction é a saída estruturada da manutenção de memória social: resumo + ops de fato.
type relationshipExtraction struct {
	Summary string               `json:"summary"`
	Facts   []relationshipFactOp `json:"facts"`
}

// relationshipFactOp é uma operação proposta pela IA sobre um fato social.
type relationshipFactOp struct {
	Op       string `json:"op"`       // ADD|UPDATE|DELETE|NOOP
	Category string `json:"category"` // taxonomia §3
	Key      string `json:"key"`
	Value    string `json:"value"`
}

// buildRelationshipMemoryPrompt monta o prompt que atualiza o resumo SOCIAL + extrai fatos atômicos
// sociais. Guardrail forte anti-clínico: sintomas/exames/diagnósticos/medicações NÃO entram (vão
// pro prontuário, não pro CRM). Recebe o resumo anterior e os fatos já conhecidos (p/ UPDATE/DELETE).
func buildRelationshipMemoryPrompt(transcript, previousSummary string, current []models.RelationshipFact) string {
	prevBlock := "(nenhum)"
	if p := strings.TrimSpace(previousSummary); p != "" {
		prevBlock = p
	}
	knownBlock := "(nenhum)"
	if len(current) > 0 {
		var sb strings.Builder
		for _, f := range current {
			fmt.Fprintf(&sb, "- %s = %s (categoria: %s)\n", f.Key, f.Value, f.Category)
		}
		knownBlock = strings.TrimRight(sb.String(), "\n")
	}
	return fmt.Sprintf(`Você mantém a MEMÓRIA SOCIAL de uma pessoa em contato com a Plenya, uma clínica brasileira. Objetivo: a recepção nunca esquecer quem é a pessoa nem repetir perguntas já respondidas.

Você produz DUAS coisas: (1) um resumo social curto e (2) uma lista de operações sobre FATOS sociais atômicos.

SÓ informação SOCIAL/RELACIONAL/ADMINISTRATIVA entra. Categorias válidas e exemplos de "key":
- identidade_social: apelido, profissao, cidade, idioma
- familia_rede: conjuge, filhos, indicado_por, acompanhante
- preferencias_atendimento: canal_preferido, formato_preferido (presencial/online), melhor_horario
- contexto_chegada: como_conheceu, motivo_contato (NÃO clínico, ex.: "quer cuidar da saúde com mais atenção")
- relacionamento: estagio (conversando/quer_agendar/agendou), sensibilidade (ex.: "não gosta de insistência")

PROIBIDO terminantemente (é do prontuário, com o médico — NÃO entra aqui): sintomas, queixas de saúde, resultados de exame, diagnósticos, doenças, medicações, condutas, ou qualquer dado clínico. Se a pessoa mencionou algo clínico, IGNORE completamente (não vire fato, não entre no resumo).

Operações de fato (op): "ADD" novo fato; "UPDATE" valor mudou; "DELETE" deixou de ser verdade; "NOOP" nada a fazer. Use a MESMA key dos fatos já conhecidos ao atualizar. Não invente: só o que aparece na conversa. key em snake_case, minúsculas.

RESUMO ANTERIOR (atualize fundindo o novo, sem perder o que já era verdade; "" se não houver nada social):
%s

FATOS JÁ CONHECIDOS:
%s

CONVERSA (cronológica; [DENTRO] = a pessoa, [FORA] = Plenya):
%s

Responda APENAS com um objeto JSON válido, sem texto fora dele:
{"summary": "<resumo social curto, máx 6 linhas, ou string vazia>", "facts": [{"op":"ADD|UPDATE|DELETE|NOOP","category":"<categoria>","key":"<snake_case>","value":"<valor curto>"}]}`, prevBlock, knownBlock, transcript)
}

// parseRelationshipExtraction extrai o JSON tolerante (lida com markdown/texto ao redor).
func parseRelationshipExtraction(raw string) relationshipExtraction {
	var ext relationshipExtraction
	trimmed := strings.TrimSpace(raw)
	start := strings.Index(trimmed, "{")
	end := strings.LastIndex(trimmed, "}")
	if start >= 0 && end > start {
		_ = json.Unmarshal([]byte(trimmed[start:end+1]), &ext)
	}
	return ext
}

// normalizeFactKey normaliza a key: minúsculas, espaços→_, só [a-z0-9_], teto de 60 chars.
func normalizeFactKey(k string) string {
	k = strings.ToLower(strings.TrimSpace(k))
	k = strings.ReplaceAll(k, " ", "_")
	var b strings.Builder
	for _, r := range k {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '_' {
			b.WriteRune(r)
		}
	}
	out := strings.Trim(b.String(), "_")
	if len(out) > 60 {
		out = out[:60]
	}
	return out
}

// normalizeFactCategory mapeia para uma das categorias válidas; default relacionamento.
func normalizeFactCategory(c string) string {
	switch strings.ToLower(strings.TrimSpace(c)) {
	case models.RelationshipCategoryIdentity:
		return models.RelationshipCategoryIdentity
	case models.RelationshipCategoryFamily:
		return models.RelationshipCategoryFamily
	case models.RelationshipCategoryPreferences:
		return models.RelationshipCategoryPreferences
	case models.RelationshipCategoryArrival:
		return models.RelationshipCategoryArrival
	default:
		return models.RelationshipCategoryRelationship
	}
}

// factClinicalTerms é um backstop simples: se a key/valor contém um termo clínico, o fato é
// descartado (o clínico fica só na janela curta da conversa; longo prazo é só social — §3.1).
var factClinicalTerms = []string{
	"sintoma", "dor ", "doença", "doenca", "diagnós", "diagnos", "exame", "exames",
	"medicament", "remédio", "remedio", "medicação", "medicacao", "pressão alta", "pressao alta",
	"diabet", "câncer", "cancer", "depress", "ansiedad", "insônia", "insonia", "colesterol",
	"glicose", "hipertens", "tireoide", "tireóide", "cirurgia", "tratament", "laudo", "receita",
}

// factLooksClinical retorna true se o fato parece conter dado clínico (backstop do guardrail).
func factLooksClinical(key, value string) bool {
	hay := strings.ToLower(key + " " + value)
	for _, term := range factClinicalTerms {
		if strings.Contains(hay, term) {
			return true
		}
	}
	return false
}

// factLabel devolve um rótulo legível para uma key de fato (para a injeção no prompt e a tela 360).
func factLabel(key string) string {
	if lbl, ok := factKeyLabels[key]; ok {
		return lbl
	}
	return strings.ReplaceAll(key, "_", " ")
}

// factKeyLabels mapeia keys comuns para rótulos em PT-BR (humaniza a tela 360 e o prompt).
var factKeyLabels = map[string]string{
	"apelido":           "como gosta de ser chamado",
	"profissao":         "profissão",
	"cidade":            "cidade",
	"idioma":            "idioma",
	"conjuge":           "cônjuge",
	"filhos":            "filhos",
	"indicado_por":      "indicado por",
	"acompanhante":      "quem acompanha",
	"canal_preferido":   "canal preferido",
	"formato_preferido": "formato preferido",
	"melhor_horario":    "melhor horário",
	"como_conheceu":     "como conheceu a Plenya",
	"motivo_contato":    "motivo do contato",
	"estagio":           "estágio",
	"sensibilidade":     "sensibilidade de abordagem",
}
