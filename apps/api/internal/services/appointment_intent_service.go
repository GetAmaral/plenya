// Package services — DetectAppointmentIntent (Calendar V1, Bloco F).
//
// Usa Claude Haiku (rápido, barato) pra classificar mensagens inbound de
// pacientes em 4 categorias quando há appointments próximos:
//
//   - "confirm"     → paciente confirma presença
//   - "cancel"      → quer desmarcar
//   - "reschedule"  → quer mudar horário
//   - "none"        → mensagem normal (não relacionada à consulta)
//
// Contrato com a IA:
//
//  1. Prompt enxuto (Haiku) — máx ~300 tokens entrada, ~100 saída.
//  2. Resposta DEVE ser JSON estrito; parsing tolerante (extrai primeiro {…}).
//  3. confidence < 0.7 vira "none" (humano-no-loop).
//  4. Em ambiguidade, IA é instruída a preferir "none".
//
// LGPD: NÃO loga conteúdo da mensagem. Só intent + confidence agregados.
package services

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/plenya/api/internal/models"
)

// AppointmentIntent classifica intenção de mensagens inbound.
type AppointmentIntent string

const (
	IntentConfirm    AppointmentIntent = "confirm"
	IntentCancel     AppointmentIntent = "cancel"
	IntentReschedule AppointmentIntent = "reschedule"
	IntentNone       AppointmentIntent = "none"
)

// IntentMinConfidence é o threshold abaixo do qual IA fallback pra "none".
const IntentMinConfidence = 0.7

// AppointmentIntentResult é a estrutura devolvida pelo classifier.
type AppointmentIntentResult struct {
	Intent         AppointmentIntent `json:"intent"`
	Confidence     float64           `json:"confidence"`
	Reasoning      string            `json:"reasoning"`
	AppointmentRef *uuid.UUID        `json:"appointment_ref,omitempty"`
}

// Modelo Haiku pra classificação rápida.
const aiModelIntent = "claude-haiku-4-5-20251001"

// DetectAppointmentIntent classifica a mensagem usando contexto dos appointments.
//
// recentAppointments: lista curta (3-5) dos appointments do paciente em janela
// próxima (últimos 7d + futuros). IA recebe o contexto mas é instruída a só
// preencher appointment_ref quando tiver alta confiança.
func (s *ConversationService) DetectAppointmentIntent(
	ctx context.Context,
	message string,
	recentAppointments []models.Appointment,
) (*AppointmentIntentResult, error) {
	if s.aiService == nil {
		return &AppointmentIntentResult{Intent: IntentNone, Confidence: 0, Reasoning: "ai service não configurado"}, nil
	}
	msg := strings.TrimSpace(message)
	if msg == "" {
		return &AppointmentIntentResult{Intent: IntentNone, Confidence: 0, Reasoning: "mensagem vazia"}, nil
	}
	// Truncate (proteção contra prompt-injection abusiva).
	if len(msg) > 500 {
		msg = msg[:500] + "…"
	}

	prompt := buildIntentPrompt(msg, recentAppointments)

	raw, err := s.aiService.CompleteText(ctx, prompt, CompleteTextOptions{
		Model:       aiModelIntent,
		MaxTokens:   300,
		Temperature: 0.1, // determinismo alto
		Timeout:     15 * time.Second,
	})
	if err != nil {
		return nil, fmt.Errorf("intent detect: %w", err)
	}

	res := parseIntentResponse(raw)

	// Threshold confidence. Se baixa OU intent inválido → none.
	if res.Confidence < IntentMinConfidence || !isValidIntent(res.Intent) {
		return &AppointmentIntentResult{
			Intent:     IntentNone,
			Confidence: res.Confidence,
			Reasoning:  res.Reasoning,
		}, nil
	}

	return res, nil
}

// buildIntentPrompt monta o prompt incluindo lista de appointments próximos.
func buildIntentPrompt(message string, appts []models.Appointment) string {
	var apptLines strings.Builder
	if len(appts) == 0 {
		apptLines.WriteString("(nenhuma consulta agendada próxima)")
	} else {
		loc, err := time.LoadLocation("America/Sao_Paulo")
		if err != nil {
			loc = time.UTC
		}
		for _, a := range appts {
			local := a.ScheduledAt.In(loc)
			apptLines.WriteString(fmt.Sprintf("- id=%s | %s | tipo=%s | status=%s\n",
				a.ID.String(),
				local.Format("02/01/2006 15:04"),
				a.Type,
				a.Status,
			))
		}
	}

	return fmt.Sprintf(`Você é assistente da clínica brasileira Plenya. Classifique a intenção da mensagem WhatsApp do paciente em relação às consultas agendadas.

MENSAGEM DO PACIENTE:
"%s"

CONSULTAS AGENDADAS PRÓXIMAS:
%s

Responda APENAS com JSON válido (sem markdown, sem comentário) no formato:
{
  "intent": "confirm" | "cancel" | "reschedule" | "none",
  "confidence": 0.0-1.0,
  "reasoning": "explicação curta em PT-BR (máx 120 chars)",
  "appointment_ref": "uuid se tiver alta certeza qual consulta, senão omita"
}

REGRAS:
- "confirm": "ok", "sim", "confirmo", "estarei lá", "confirmado", "tudo certo".
- "cancel": "cancela", "desmarca", "não posso ir", "preciso cancelar".
- "reschedule": "outro horário", "remarca", "tem outro dia", "podemos mudar".
- "none": saudação, dúvida geral, agradecimento, ou qualquer coisa fora do escopo das consultas listadas.
- Em caso de ambiguidade, prefira "none" (humano vai analisar).
- confidence < 0.7 obrigatoriamente vira "none".
- Só preencha appointment_ref se tiver alta certeza qual consulta o paciente cita.`, message, apptLines.String())
}

// parseIntentResponse extrai JSON do output do modelo. Tolerante a wrappers
// (ex: "```json{...}```", texto adicional). Falha → intent=none.
func parseIntentResponse(raw string) *AppointmentIntentResult {
	fallback := &AppointmentIntentResult{
		Intent:     IntentNone,
		Confidence: 0,
		Reasoning:  "ai response parse failed",
	}

	// Heurística: extrair primeiro bloco { ... } balanceado.
	start := strings.Index(raw, "{")
	end := strings.LastIndex(raw, "}")
	if start < 0 || end <= start {
		return fallback
	}
	candidate := raw[start : end+1]

	var parsed struct {
		Intent         string  `json:"intent"`
		Confidence     float64 `json:"confidence"`
		Reasoning      string  `json:"reasoning"`
		AppointmentRef string  `json:"appointment_ref"`
	}
	if err := json.Unmarshal([]byte(candidate), &parsed); err != nil {
		return fallback
	}

	res := &AppointmentIntentResult{
		Intent:     AppointmentIntent(strings.ToLower(strings.TrimSpace(parsed.Intent))),
		Confidence: parsed.Confidence,
		Reasoning:  truncateReasoning(parsed.Reasoning),
	}
	if parsed.AppointmentRef != "" {
		if id, err := uuid.Parse(strings.TrimSpace(parsed.AppointmentRef)); err == nil {
			res.AppointmentRef = &id
		}
	}
	return res
}

func truncateReasoning(s string) string {
	s = strings.TrimSpace(s)
	if len(s) > 200 {
		return s[:200]
	}
	return s
}

func isValidIntent(i AppointmentIntent) bool {
	switch i {
	case IntentConfirm, IntentCancel, IntentReschedule, IntentNone:
		return true
	}
	return false
}
