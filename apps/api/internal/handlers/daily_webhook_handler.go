package handlers

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/services"
)

// DailyWebhookHandler recebe os eventos do domínio Daily.co (gravação + transcrição).
//
//   - POST /webhooks/daily — recording.started/ready-to-download/error +
//     transcript.started/ready-to-download/error.
//
// Segurança: assinatura HMAC (X-Webhook-Signature + X-Webhook-Timestamp) com o
// segredo retornado ao registrar o webhook (DAILY_CO_WEBHOOK_SECRET). Fail-closed
// em prod. O ping de criação ({"test":"test"}) é respondido 200 sem assinatura.
type DailyWebhookHandler struct {
	cfg           *config.Config
	dailyCoSvc    *services.DailyCoService
	telemedRecSvc *services.TelemedRecordingService
}

func NewDailyWebhookHandler(cfg *config.Config, dailyCoSvc *services.DailyCoService, telemedRecSvc *services.TelemedRecordingService) *DailyWebhookHandler {
	return &DailyWebhookHandler{cfg: cfg, dailyCoSvc: dailyCoSvc, telemedRecSvc: telemedRecSvc}
}

type dailyWebhookEvent struct {
	Version string          `json:"version"`
	Type    string          `json:"type"`
	ID      string          `json:"id"`
	Payload json.RawMessage `json:"payload"`
	EventTS float64         `json:"event_ts"`
}

type dailyRecordingPayload struct {
	Action      string `json:"action"`
	RecordingID string `json:"recording_id"`
	RoomName    string `json:"room_name"`
	Status      string `json:"status"`
	Duration    int    `json:"duration"`
	S3Key       string `json:"s3_key"`
	StartTS     int64  `json:"start_ts"`
	ErrorMsg    string `json:"error_msg"`
}

type dailyTranscriptPayload struct {
	ID       string `json:"id"` // transcript id (≠ event id)
	RoomName string `json:"room_name"`
	Status   string `json:"status"`
	StartTS  int64  `json:"start_ts"`
	ErrorMsg string `json:"error_msg"`
}

// Receive — POST /webhooks/daily
func (h *DailyWebhookHandler) Receive(c *fiber.Ctx) error {
	body := c.Body() // raw — necessário pra HMAC

	// Ping de criação do webhook: Daily faz POST {"test":"test"} e exige 200 rápido
	// ANTES de existir assinatura. Respondemos sem verificar.
	var probe struct {
		Test string `json:"test"`
	}
	if err := json.Unmarshal(body, &probe); err == nil && probe.Test != "" {
		return c.SendStatus(fiber.StatusOK)
	}

	// Verificação de assinatura. Fail-closed em prod; em dev, loga e segue.
	if h.cfg.DailyCo.WebhookSecret == "" {
		if h.cfg.Server.Environment != "development" {
			log.Printf("🚫 [DAILY WEBHOOK] DAILY_CO_WEBHOOK_SECRET ausente em %q — rejeitando", h.cfg.Server.Environment)
			return c.SendStatus(fiber.StatusServiceUnavailable)
		}
		log.Printf("⚠️  [DAILY WEBHOOK] segredo não configurado (dev) — assinatura NÃO validada")
	} else if h.dailyCoSvc != nil {
		sig := c.Get("X-Webhook-Signature")
		ts := c.Get("X-Webhook-Timestamp")
		if err := h.dailyCoSvc.VerifyWebhookSignature(sig, ts, body); err != nil {
			log.Printf("❌ [DAILY WEBHOOK] assinatura inválida: %v", err)
			return c.SendStatus(fiber.StatusUnauthorized)
		}
	}

	var ev dailyWebhookEvent
	if err := json.Unmarshal(body, &ev); err != nil {
		log.Printf("❌ [DAILY WEBHOOK] parse falhou: %v", err)
		return c.SendStatus(fiber.StatusOK) // não devolve 4xx (evita circuit-breaker)
	}
	if h.telemedRecSvc == nil || ev.Type == "" {
		return c.SendStatus(fiber.StatusOK)
	}

	// Processa async (download do VTT pode levar segundos) e responde 200 já,
	// pra não tropeçar no circuit-breaker do Daily.
	evType := ev.Type
	payload := append([]byte(nil), ev.Payload...)
	eventTS := ev.EventTS
	svc := h.telemedRecSvc
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		h.dispatch(ctx, svc, evType, payload, eventTS)
	}()

	return c.SendStatus(fiber.StatusOK)
}

func (h *DailyWebhookHandler) dispatch(ctx context.Context, svc *services.TelemedRecordingService, evType string, payload []byte, eventTS float64) {
	tsTime := func(unix int64) time.Time {
		if unix > 0 {
			return time.Unix(unix, 0).UTC()
		}
		if eventTS > 0 {
			return time.Unix(int64(eventTS), 0).UTC()
		}
		return time.Now().UTC()
	}

	var err error
	switch evType {
	case "recording.started":
		var p dailyRecordingPayload
		_ = json.Unmarshal(payload, &p)
		if p.RoomName == "" {
			return
		}
		err = svc.MarkRecordingStarted(ctx, p.RoomName, p.RecordingID, tsTime(p.StartTS))
	case "recording.ready-to-download":
		var p dailyRecordingPayload
		_ = json.Unmarshal(payload, &p)
		if p.RoomName == "" {
			return
		}
		err = svc.MarkRecordingReady(ctx, p.RoomName, p.RecordingID, p.Duration, p.S3Key, tsTime(0))
	case "recording.error":
		var p dailyRecordingPayload
		_ = json.Unmarshal(payload, &p)
		if p.RoomName == "" {
			return
		}
		err = svc.MarkRecordingError(ctx, p.RoomName, p.ErrorMsg)
	case "transcript.started":
		var p dailyTranscriptPayload
		_ = json.Unmarshal(payload, &p)
		if p.RoomName == "" {
			return
		}
		err = svc.MarkTranscriptStarted(ctx, p.RoomName, p.ID)
	case "transcript.ready-to-download":
		var p dailyTranscriptPayload
		_ = json.Unmarshal(payload, &p)
		if p.RoomName == "" {
			return
		}
		err = svc.MarkTranscriptReady(ctx, p.RoomName, p.ID, tsTime(0))
	case "transcript.error":
		var p dailyTranscriptPayload
		_ = json.Unmarshal(payload, &p)
		if p.RoomName == "" {
			return
		}
		err = svc.MarkTranscriptError(ctx, p.RoomName, p.ErrorMsg)
	default:
		return
	}
	if err != nil {
		log.Printf("⚠️  [DAILY WEBHOOK] %s: %v", evType, err)
	}
}
