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

// Constantes pra controle de contexto. Calibrados pra ficar < 8k tokens em pt-BR
// (1 token ≈ 3.5 chars em PT) com folga: 50 msgs * 500 chars = 25k chars ≈ 7k tokens.
const (
	aiSummaryMaxMessages    = 50
	aiSuggestionMaxMessages = 10
	aiMaxContentChars       = 500

	// Modelos. Haiku pra resumo (barato + rápido); Sonnet pra escrita (qualidade
	// do tom Plenya). Atualize aqui quando subir versão (claude-md doc).
	aiModelSummary    = "claude-haiku-4-5-20251001"
	aiModelSuggestion = "claude-sonnet-4-6-20251001"

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
			a.CreatedAt.Format("02/01 15:04"),
			string(a.Channel),
			content,
		)
	}
	return sb.String()
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
