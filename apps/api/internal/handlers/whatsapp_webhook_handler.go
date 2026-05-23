package handlers

import (
	"encoding/json"
	"fmt"
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

// mediaObj é o objeto comum de mídia no payload Meta (image/audio/video/document/sticker).
type mediaObj struct {
	ID       string `json:"id"`
	MIMEType string `json:"mime_type"`
	SHA256   string `json:"sha256"`
	Filename string `json:"filename"` // só documents
	Caption  string `json:"caption"`  // image/video/document
	Voice    bool   `json:"voice"`    // audio: true = nota de voz
}

// metaLocation é o payload de uma mensagem de localização.
type metaLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Name      string  `json:"name"`
	Address   string  `json:"address"`
}

// metaContact é um contato compartilhado (vCard).
type metaContact struct {
	Name struct {
		FormattedName string `json:"formatted_name"`
	} `json:"name"`
	Phones []struct {
		Phone string `json:"phone"`
	} `json:"phones"`
}

// metaMessage é uma mensagem inbound do payload Meta (nomeada pra os helpers
// pickMedia/formatContacts poderem recebê-la por ponteiro).
type metaMessage struct {
	From      string `json:"from"`
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	Type      string `json:"type"`
	Text      *struct {
		Body string `json:"body"`
	} `json:"text,omitempty"`
	Image    *mediaObj     `json:"image,omitempty"`
	Audio    *mediaObj     `json:"audio,omitempty"`
	Video    *mediaObj     `json:"video,omitempty"`
	Document *mediaObj     `json:"document,omitempty"`
	Sticker  *mediaObj     `json:"sticker,omitempty"`
	Location *metaLocation `json:"location,omitempty"`
	Contacts []metaContact `json:"contacts,omitempty"`
	// Context indica resposta/encaminhamento; em coexistence, mensagens enviadas
	// pelo app do celular trazem context.from = número do negócio (Fase 2.3).
	Context *struct {
		From string `json:"from"`
		ID   string `json:"id"`
	} `json:"context,omitempty"`
}

// metaWebhookPayload é o subset do payload Meta que processamos.
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
				Messages []metaMessage `json:"messages,omitempty"`
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
				ts, _ := parseUnixTS(msg.Timestamp)
				var namePtr *string
				if name, ok := profileNames[msg.From]; ok && name != "" {
					n := name
					namePtr = &n
				}
				if h.leadService == nil {
					continue
				}

				switch msg.Type {
				case "image", "audio", "video", "document", "sticker":
					media := pickMedia(&msg)
					if media == nil || media.ID == "" {
						log.Printf("[WHATSAPP WEBHOOK] midia %s sem id (msg=%s)", msg.Type, msg.ID)
						continue
					}
					kind := msg.Type
					if msg.Type == "audio" && media.Voice {
						kind = "voice"
					}
					input := services.InboundWhatsAppMediaInput{
						PhoneE164:   msg.From,
						Name:        namePtr,
						MediaID:     media.ID,
						Kind:        kind,
						MIME:        media.MIMEType,
						Filename:    media.Filename,
						Caption:     media.Caption,
						WAMessageID: msg.ID,
						ReceivedAt:  ts,
					}
					// Download da Meta pode levar segundos: processa async pra o webhook
					// responder 200 rapido (Meta faz retry agressivo em timeout).
					leadSvc := h.leadService
					go func() {
						if _, err := leadSvc.ProcessInboundWhatsAppMedia(input); err != nil {
							log.Printf("[WHATSAPP WEBHOOK] processInboundMedia(%s): %v", input.Kind, err)
						}
					}()
					continue

				case "location":
					if msg.Location == nil {
						continue
					}
					txt := formatLocation(msg.Location.Latitude, msg.Location.Longitude, msg.Location.Name, msg.Location.Address)
					if _, err := h.leadService.ProcessInboundWhatsApp(services.InboundWhatsAppInput{
						PhoneE164: msg.From, Name: namePtr, Text: txt, WAMessageID: msg.ID, ReceivedAt: ts,
					}); err != nil {
						log.Printf("[WHATSAPP WEBHOOK] processInbound(location): %v", err)
					}
					continue

				case "contacts":
					txt := formatContacts(&msg)
					if txt == "" {
						continue
					}
					if _, err := h.leadService.ProcessInboundWhatsApp(services.InboundWhatsAppInput{
						PhoneE164: msg.From, Name: namePtr, Text: txt, WAMessageID: msg.ID, ReceivedAt: ts,
					}); err != nil {
						log.Printf("[WHATSAPP WEBHOOK] processInbound(contacts): %v", err)
					}
					continue

				case "text":
					// segue pro processamento de texto abaixo
				default:
					continue
				}

				// Auto-unsubscribe: detecta keywords antes de processar como inbound normal
				if services.IsUnsubscribeKeyword(msg.Text.Body) {
					if h.leadService != nil {
						h.leadService.UnsubscribeByPhone(msg.From, msg.Text.Body)
					}
					continue
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

			// Status updates (delivered/read/failed) — Phase 2: persiste como
			// LeadActivity message_status_changed, vinculado ao envio original.
			for _, st := range change.Value.Statuses {
				ts, _ := parseUnixTS(st.Timestamp)
				if h.leadService != nil {
					h.leadService.RecordWhatsAppStatus(st.ID, st.Status, st.RecipientID, ts)
				}
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

// pickMedia retorna o objeto de mídia correspondente ao Type da mensagem.
func pickMedia(msg *metaMessage) *mediaObj {
	switch msg.Type {
	case "image":
		return msg.Image
	case "audio":
		return msg.Audio
	case "video":
		return msg.Video
	case "document":
		return msg.Document
	case "sticker":
		return msg.Sticker
	default:
		return nil
	}
}

// formatLocation transforma uma localização em texto legível pra a conversa.
func formatLocation(lat, long float64, name, address string) string {
	var b strings.Builder
	b.WriteString("[Localização]")
	if name != "" {
		b.WriteString(" " + name)
	}
	if address != "" {
		b.WriteString(" — " + address)
	}
	b.WriteString(fmt.Sprintf(" (%.6f, %.6f)", lat, long))
	b.WriteString(fmt.Sprintf(" https://maps.google.com/?q=%.6f,%.6f", lat, long))
	return b.String()
}

// formatContacts transforma contatos compartilhados em texto legível.
func formatContacts(msg *metaMessage) string {
	if len(msg.Contacts) == 0 {
		return ""
	}
	var parts []string
	for _, c := range msg.Contacts {
		name := strings.TrimSpace(c.Name.FormattedName)
		var phone string
		if len(c.Phones) > 0 {
			phone = strings.TrimSpace(c.Phones[0].Phone)
		}
		switch {
		case name != "" && phone != "":
			parts = append(parts, name+" "+phone)
		case name != "":
			parts = append(parts, name)
		case phone != "":
			parts = append(parts, phone)
		}
	}
	if len(parts) == 0 {
		return ""
	}
	return "[Contato] " + strings.Join(parts, "; ")
}
