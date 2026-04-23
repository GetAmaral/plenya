package handlers

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/services"
)

// WhatsAppWebhookHandler implementa os endpoints do webhook Meta WhatsApp Cloud API.
//
//   - GET  /webhooks/whatsapp — verificação inicial (Meta envia challenge)
//   - POST /webhooks/whatsapp — recebe mensagens inbound + status updates (delivered/read)
type WhatsAppWebhookHandler struct {
	cfg             *config.Config
	whatsappService *services.WhatsAppService
	leadService     *services.LeadService
}

func NewWhatsAppWebhookHandler(
	cfg *config.Config,
	whatsappService *services.WhatsAppService,
	leadService *services.LeadService,
) *WhatsAppWebhookHandler {
	return &WhatsAppWebhookHandler{
		cfg:             cfg,
		whatsappService: whatsappService,
		leadService:     leadService,
	}
}

// Verify — GET /webhooks/whatsapp
// Meta envia query params hub.mode=subscribe, hub.challenge, hub.verify_token.
// Devemos retornar o challenge se o verify_token bater com o nosso.
func (h *WhatsAppWebhookHandler) Verify(c *fiber.Ctx) error {
	mode := c.Query("hub.mode")
	token := c.Query("hub.verify_token")
	challenge := c.Query("hub.challenge")

	if mode == "subscribe" && token != "" && token == h.cfg.WhatsApp.WebhookVerifyToken {
		return c.SendString(challenge)
	}
	return c.SendStatus(fiber.StatusUnauthorized)
}

// metaWebhookPayload é o subset do payload Meta que processamos na Fase 1.
// Estrutura completa em https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/payload-examples
type metaWebhookPayload struct {
	Object string `json:"object"`
	Entry  []struct {
		ID      string `json:"id"`
		Changes []struct {
			Field string `json:"field"`
			Value struct {
				MessagingProduct string `json:"messaging_product"`
				Metadata         struct {
					DisplayPhoneNumber string `json:"display_phone_number"`
					PhoneNumberID      string `json:"phone_number_id"`
				} `json:"metadata"`
				Contacts []struct {
					Profile struct {
						Name string `json:"name"`
					} `json:"profile"`
					WaID string `json:"wa_id"` // E.164 sem +
				} `json:"contacts"`
				Messages []struct {
					From      string `json:"from"`
					ID        string `json:"id"`
					Timestamp string `json:"timestamp"`
					Type      string `json:"type"`
					Text      *struct {
						Body string `json:"body"`
					} `json:"text,omitempty"`
				} `json:"messages,omitempty"`
				Statuses []struct {
					ID          string `json:"id"`
					RecipientID string `json:"recipient_id"`
					Status      string `json:"status"` // sent | delivered | read | failed
					Timestamp   string `json:"timestamp"`
				} `json:"statuses,omitempty"`
			} `json:"value"`
		} `json:"changes"`
	} `json:"entry"`
}

// Receive — POST /webhooks/whatsapp
// Valida assinatura HMAC X-Hub-Signature-256 e processa mensagens inbound.
// Sempre retorna 200 (Meta retry agressivo se 5xx).
func (h *WhatsAppWebhookHandler) Receive(c *fiber.Ctx) error {
	body := c.Body() // raw bytes — necessário pra HMAC
	signature := c.Get("X-Hub-Signature-256")

	// Em PROD, AppSecret é obrigatório. Sem ele, qualquer atacante pode forjar
	// payloads e criar Leads. Falhamos fechado.
	if h.cfg.WhatsApp.AppSecret == "" {
		if h.cfg.Server.Environment != "development" {
			log.Printf("🚫 [WHATSAPP WEBHOOK] APP_SECRET ausente em ambiente %q — rejeitando", h.cfg.Server.Environment)
			return c.SendStatus(fiber.StatusServiceUnavailable)
		}
		log.Printf("⚠️  [WHATSAPP WEBHOOK] APP_SECRET não configurado (dev) — assinatura NÃO validada")
	} else {
		if err := h.whatsappService.VerifyWebhookSignature(signature, body); err != nil {
			log.Printf("❌ [WHATSAPP WEBHOOK] assinatura inválida: %v", err)
			return c.SendStatus(fiber.StatusUnauthorized)
		}
	}

	var payload metaWebhookPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		log.Printf("❌ [WHATSAPP WEBHOOK] parse falhou: %v", err)
		return c.SendStatus(fiber.StatusOK) // não devolve 4xx pra evitar retries
	}

	for _, entry := range payload.Entry {
		for _, change := range entry.Changes {
			// Mapa profile name por wa_id (mesmo número pode aparecer em multiplas msgs)
			profileNames := make(map[string]string, len(change.Value.Contacts))
			for _, contact := range change.Value.Contacts {
				if contact.WaID != "" && contact.Profile.Name != "" {
					profileNames[contact.WaID] = contact.Profile.Name
				}
			}

			for _, msg := range change.Value.Messages {
				if msg.Type != "text" || msg.Text == nil {
					// Fase 1: só processa text. Outros tipos (image/audio/document) ficam pra Fase 2.
					log.Printf("📱 [WHATSAPP WEBHOOK] tipo não processado: %s (id=%s)", msg.Type, msg.ID)
					continue
				}

				ts, _ := parseUnixTS(msg.Timestamp)
				var namePtr *string
				if name, ok := profileNames[msg.From]; ok && name != "" {
					n := name
					namePtr = &n
				}

				if h.leadService != nil {
					_, err := h.leadService.ProcessInboundWhatsApp(services.InboundWhatsAppInput{
						PhoneE164:   msg.From,
						Name:        namePtr,
						Text:        msg.Text.Body,
						WAMessageID: msg.ID,
						ReceivedAt:  ts,
					})
					if err != nil {
						log.Printf("❌ [WHATSAPP WEBHOOK] processInbound: %v", err)
					}
				}
			}

			// Status updates (delivered/read/failed) — Fase 1: apenas log
			for _, st := range change.Value.Statuses {
				log.Printf("📱 [WHATSAPP WEBHOOK] status=%s message_id=%s recipient=%s",
					st.Status, st.ID, st.RecipientID)
			}
		}
	}

	return c.SendStatus(fiber.StatusOK)
}

// parseUnixTS converte timestamp Meta (string Unix segs) para time.Time UTC.
func parseUnixTS(s string) (time.Time, bool) {
	s = strings.TrimSpace(s)
	if s == "" {
		return time.Time{}, false
	}
	var sec int64
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return time.Time{}, false
		}
		sec = sec*10 + int64(ch-'0')
	}
	return time.Unix(sec, 0).UTC(), true
}
