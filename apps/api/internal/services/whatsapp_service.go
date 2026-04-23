package services

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/plenya/api/internal/config"
)

// WhatsAppService é um cliente HTTP para a Meta WhatsApp Cloud API.
//
// Sem SDK terceiro: a Meta não publica SDK Go oficial e libs comunitárias trazem dependências
// desnecessárias para o que precisamos na Fase 1 (1 template + validação de webhook).
//
// Em dev (cfg.WhatsApp.PhoneNumberID vazio), todos os envios caem em log no stdout —
// mesmo padrão de fallback do EmailService.
type WhatsAppService struct {
	cfg    *config.Config
	client *http.Client
}

func NewWhatsAppService(cfg *config.Config) *WhatsAppService {
	return &WhatsAppService{
		cfg:    cfg,
		client: &http.Client{Timeout: 10 * time.Second},
	}
}

// E.164 regex (RFC 5733): + seguido de 1-15 dígitos. Meta exige sem o +.
var e164Pattern = regexp.MustCompile(`^\+[1-9]\d{1,14}$`)

// NormalizeE164 retorna o telefone sem o + e sem caracteres não-dígito.
// A Meta API espera "5511999998888", não "+5511999998888".
func NormalizeE164(phone string) (string, error) {
	cleaned := strings.TrimSpace(phone)
	if !e164Pattern.MatchString(cleaned) {
		return "", fmt.Errorf("whatsapp: telefone não está em E.164 (%q)", phone)
	}
	return cleaned[1:], nil
}

// SendMagicLink envia o template "magic_link" (categoria utility, pt_BR).
// O template aprovado na Meta deve ter o formato:
//
//	"Olá! Aqui é da Plenya. Seu resultado do Escore Plenya Light está pronto.
//	 Acesse pelo link seguro: {{1}}. Esse link é único e expira em 7 dias.
//	 Se não foi você, ignore esta mensagem."
//
// IMPORTANTE: o "7 dias" precisa bater com o JWT TTL em anonymous_score_service.go.
// Se mudar um, mudar o outro. Versão atual: 7 dias (alinhado em 2026-04-23).
// O nome do template é configurável via WHATSAPP_TEMPLATE_MAGIC_LINK (default: "magic_link").
func (s *WhatsAppService) SendMagicLink(toE164, magicLink string) error {
	templateName := s.cfg.WhatsApp.TemplateMagicLink
	if templateName == "" {
		templateName = "magic_link"
	}
	return s.SendTemplate(toE164, templateName, "pt_BR", []string{magicLink})
}

// SendTemplate envia uma mensagem de template aprovado pela Meta.
func (s *WhatsAppService) SendTemplate(toE164, templateName, langCode string, bodyParams []string) error {
	if s.cfg.WhatsApp.PhoneNumberID == "" || s.cfg.WhatsApp.AccessToken == "" {
		log.Printf("📱 [WHATSAPP DEV] SendTemplate to=%s template=%s lang=%s params=%d (sem credenciais — log apenas)",
			toE164, templateName, langCode, len(bodyParams))
		return nil
	}

	to, err := NormalizeE164(toE164)
	if err != nil {
		return err
	}

	parameters := make([]map[string]string, 0, len(bodyParams))
	for _, p := range bodyParams {
		parameters = append(parameters, map[string]string{"type": "text", "text": p})
	}

	components := []map[string]any{}
	if len(parameters) > 0 {
		components = append(components, map[string]any{
			"type":       "body",
			"parameters": parameters,
		})
	}

	payload := map[string]any{
		"messaging_product": "whatsapp",
		"recipient_type":    "individual",
		"to":                to,
		"type":              "template",
		"template": map[string]any{
			"name":       templateName,
			"language":   map[string]string{"code": langCode},
			"components": components,
		},
	}

	raw, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("whatsapp: marshal payload: %w", err)
	}

	apiVersion := s.cfg.WhatsApp.GraphAPIVersion
	if apiVersion == "" {
		apiVersion = "v18.0"
	}
	url := fmt.Sprintf("https://graph.facebook.com/%s/%s/messages", apiVersion, s.cfg.WhatsApp.PhoneNumberID)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(raw))
	if err != nil {
		return fmt.Errorf("whatsapp: build request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+s.cfg.WhatsApp.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("whatsapp: http: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(io.LimitReader(resp.Body, 2048))

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		// Não logar o link nem o conteúdo do template em prod.
		// Logamos apenas o message_id retornado pela Meta (útil pra correlacionar webhooks).
		var meta struct {
			Messages []struct {
				ID string `json:"id"`
			} `json:"messages"`
		}
		_ = json.Unmarshal(body, &meta)
		var msgID string
		if len(meta.Messages) > 0 {
			msgID = meta.Messages[0].ID
		}
		log.Printf("📱 [WHATSAPP] template=%s sent to=%s message_id=%s", templateName, to, msgID)
		return nil
	}

	return fmt.Errorf("whatsapp: status %d: %s", resp.StatusCode, string(body))
}

// VerifyWebhookSignature valida o header X-Hub-Signature-256 enviado pela Meta.
// Retorna nil se válido. Body deve ser o raw bytes do request, antes de qualquer parsing.
func (s *WhatsAppService) VerifyWebhookSignature(signatureHeader string, body []byte) error {
	if s.cfg.WhatsApp.AppSecret == "" {
		return fmt.Errorf("whatsapp: APP_SECRET não configurado — webhook não pode ser validado")
	}
	const prefix = "sha256="
	if !strings.HasPrefix(signatureHeader, prefix) {
		return fmt.Errorf("whatsapp: signature header inválido")
	}
	provided, err := hex.DecodeString(signatureHeader[len(prefix):])
	if err != nil {
		return fmt.Errorf("whatsapp: signature hex inválido: %w", err)
	}
	mac := hmac.New(sha256.New, []byte(s.cfg.WhatsApp.AppSecret))
	mac.Write(body)
	expected := mac.Sum(nil)
	if !hmac.Equal(provided, expected) {
		return fmt.Errorf("whatsapp: signature mismatch")
	}
	return nil
}

// IsConfigured indica se temos credenciais suficientes para falar com a Meta.
// Usado por handlers que só fazem dispatch real em prod.
func (s *WhatsAppService) IsConfigured() bool {
	return s.cfg.WhatsApp.PhoneNumberID != "" && s.cfg.WhatsApp.AccessToken != ""
}
