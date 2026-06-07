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
	"mime/multipart"
	"net/http"
	"net/url"
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

// SendConsultationPrepInvite envia o convite de preparação pré-consulta com o magic link como
// parâmetro {{1}}. Nome do template via WHATSAPP_TEMPLATE_PREP_INVITE (default "magic_link" — já
// aprovado; trocar p/ "consultation_prep_invite" quando o dedicado for aprovado pela Meta).
func (s *WhatsAppService) SendConsultationPrepInvite(toE164, link string) error {
	templateName := s.cfg.WhatsApp.TemplatePrepInvite
	if templateName == "" {
		templateName = "magic_link"
	}
	return s.SendTemplate(toE164, templateName, "pt_BR", []string{link})
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

// SendLeadAlert envia template "lead_alert" pra um vendedor (utility, pt_BR).
// Template aprovado:
//
//	"Novo lead Plenya: {{1}} ({{2}}). Origem: {{3}}. Abra no admin: {{4}}"
//
// Configurável via WHATSAPP_TEMPLATE_LEAD_ALERT (default "lead_alert").
func (s *WhatsAppService) SendLeadAlert(toE164, leadName, contact, source, adminURL string) error {
	templateName := s.cfg.WhatsApp.TemplateLeadAlert
	if templateName == "" {
		templateName = "lead_alert"
	}
	return s.SendTemplate(toE164, templateName, "pt_BR", []string{leadName, contact, source, adminURL})
}

// SendTextMessage envia mensagem free-form (não-template).
// Só funciona dentro da janela de 24h após inbound do mesmo número (regra Meta).
// Caller deve checar a janela antes de chamar — Meta retorna erro se fora.
// Retorna o wa_message_id retornado pela Meta (pra correlacionar com status updates).
func (s *WhatsAppService) SendTextMessage(toE164, body string) (string, error) {
	if s.cfg.WhatsApp.PhoneNumberID == "" || s.cfg.WhatsApp.AccessToken == "" {
		log.Printf("📱 [WHATSAPP DEV] SendTextMessage to=%s body=%d bytes (sem credenciais — log apenas)",
			toE164, len(body))
		return "dev-" + toE164, nil
	}

	to, err := NormalizeE164(toE164)
	if err != nil {
		return "", err
	}

	payload := map[string]any{
		"messaging_product": "whatsapp",
		"recipient_type":    "individual",
		"to":                to,
		"type":              "text",
		"text":              map[string]string{"body": body},
	}

	raw, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("whatsapp: marshal payload: %w", err)
	}

	apiVersion := s.cfg.WhatsApp.GraphAPIVersion
	if apiVersion == "" {
		apiVersion = "v18.0"
	}
	url := fmt.Sprintf("https://graph.facebook.com/%s/%s/messages", apiVersion, s.cfg.WhatsApp.PhoneNumberID)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(raw))
	if err != nil {
		return "", fmt.Errorf("whatsapp: build request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+s.cfg.WhatsApp.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("whatsapp: http: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(io.LimitReader(resp.Body, 2048))
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("whatsapp: status %d: %s", resp.StatusCode, string(respBody))
	}

	var meta struct {
		Messages []struct {
			ID string `json:"id"`
		} `json:"messages"`
	}
	_ = json.Unmarshal(respBody, &meta)
	var msgID string
	if len(meta.Messages) > 0 {
		msgID = meta.Messages[0].ID
	}
	log.Printf("📱 [WHATSAPP] text sent to=%s message_id=%s body=%d bytes", to, msgID, len(body))
	return msgID, nil
}

// SendTypingIndicator marca a última mensagem do cliente como lida e mostra o "digitando…"
// no WhatsApp dele (Cloud API: status=read + typing_indicator). Dura até ~25s ou até a próxima
// mensagem que enviarmos. Best-effort: o erro não deve abortar o fluxo de resposta.
func (s *WhatsAppService) SendTypingIndicator(inboundWAMessageID string) error {
	if inboundWAMessageID == "" {
		return nil
	}
	if s.cfg.WhatsApp.PhoneNumberID == "" || s.cfg.WhatsApp.AccessToken == "" {
		return nil // sem credenciais (dev) — no-op silencioso
	}

	payload := map[string]any{
		"messaging_product": "whatsapp",
		"status":            "read",
		"message_id":        inboundWAMessageID,
		"typing_indicator":  map[string]string{"type": "text"},
	}
	raw, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	apiVersion := s.cfg.WhatsApp.GraphAPIVersion
	if apiVersion == "" {
		apiVersion = "v18.0"
	}
	url := fmt.Sprintf("https://graph.facebook.com/%s/%s/messages", apiVersion, s.cfg.WhatsApp.PhoneNumberID)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(raw))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+s.cfg.WhatsApp.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := io.ReadAll(io.LimitReader(resp.Body, 1024))
		return fmt.Errorf("whatsapp: typing status %d: %s", resp.StatusCode, string(b))
	}
	return nil
}

// maxWhatsAppMediaBytes é o teto de tamanho de mídia inbound que aceitamos baixar.
// Áudio/vídeo do WhatsApp raramente passam disso; protege contra OOM/abuso.
const maxWhatsAppMediaBytes = 30 * 1024 * 1024 // 30MB

// MediaDownload é o resultado de baixar um arquivo de mídia inbound da Meta.
type MediaDownload struct {
	Bytes    []byte
	MIME     string
	FileSize int64
	SHA256   string
}

// DownloadMedia baixa um arquivo de mídia inbound da Meta Cloud API.
//
// Dois passos (regra da Meta):
//  1. GET /{media-id} → JSON { url, mime_type, file_size, sha256 } (URL de vida curta);
//  2. GET nessa url com o Bearer token (a Meta exige Authorization no fetch do binário).
//
// Sem credenciais retorna erro — não há fallback de log pra mídia (não dá pra forjar bytes).
func (s *WhatsAppService) DownloadMedia(mediaID string) (*MediaDownload, error) {
	if s.cfg.WhatsApp.AccessToken == "" {
		return nil, fmt.Errorf("whatsapp: AccessToken não configurado — download de mídia indisponível")
	}
	if strings.TrimSpace(mediaID) == "" {
		return nil, fmt.Errorf("whatsapp: mediaID vazio")
	}

	apiVersion := s.cfg.WhatsApp.GraphAPIVersion
	if apiVersion == "" {
		apiVersion = "v18.0"
	}
	token := s.cfg.WhatsApp.AccessToken

	// Cliente com timeout maior — binário pode ser grande (vs. 10s do client de texto).
	mediaClient := &http.Client{Timeout: 60 * time.Second}

	// Passo 1 — metadados.
	metaURL := fmt.Sprintf("https://graph.facebook.com/%s/%s", apiVersion, mediaID)
	metaReq, err := http.NewRequest(http.MethodGet, metaURL, nil)
	if err != nil {
		return nil, fmt.Errorf("whatsapp: build media meta request: %w", err)
	}
	metaReq.Header.Set("Authorization", "Bearer "+token)

	metaResp, err := mediaClient.Do(metaReq)
	if err != nil {
		return nil, fmt.Errorf("whatsapp: media meta http: %w", err)
	}
	defer metaResp.Body.Close()
	metaBody, _ := io.ReadAll(io.LimitReader(metaResp.Body, 4096))
	if metaResp.StatusCode < 200 || metaResp.StatusCode >= 300 {
		return nil, fmt.Errorf("whatsapp: media meta status %d: %s", metaResp.StatusCode, string(metaBody))
	}

	var meta struct {
		URL      string `json:"url"`
		MIMEType string `json:"mime_type"`
		SHA256   string `json:"sha256"`
		FileSize int64  `json:"file_size"`
	}
	if err := json.Unmarshal(metaBody, &meta); err != nil {
		return nil, fmt.Errorf("whatsapp: parse media meta: %w", err)
	}
	if meta.URL == "" {
		return nil, fmt.Errorf("whatsapp: media meta sem url")
	}
	if meta.FileSize > maxWhatsAppMediaBytes {
		return nil, fmt.Errorf("whatsapp: mídia %d bytes excede limite de %d", meta.FileSize, maxWhatsAppMediaBytes)
	}

	// Defesa SSRF/vazamento de token: só anexamos o Bearer se a URL de mídia
	// aponta pra um host conhecido da Meta (HTTPS). Caso contrário, recusamos —
	// uma resposta MITM/forjada do passo 1 não deve levar o token pra host arbitrário.
	if !isMetaMediaHost(meta.URL) {
		return nil, fmt.Errorf("whatsapp: host de mídia inesperado")
	}

	// Passo 2 — binário (Authorization também é exigido aqui).
	binReq, err := http.NewRequest(http.MethodGet, meta.URL, nil)
	if err != nil {
		return nil, fmt.Errorf("whatsapp: build media bin request: %w", err)
	}
	binReq.Header.Set("Authorization", "Bearer "+token)

	binResp, err := mediaClient.Do(binReq)
	if err != nil {
		return nil, fmt.Errorf("whatsapp: media bin http: %w", err)
	}
	defer binResp.Body.Close()
	if binResp.StatusCode < 200 || binResp.StatusCode >= 300 {
		errBody, _ := io.ReadAll(io.LimitReader(binResp.Body, 1024))
		return nil, fmt.Errorf("whatsapp: media bin status %d: %s", binResp.StatusCode, string(errBody))
	}

	// Lê com limite +1 pra detectar excesso mesmo sem file_size confiável no passo 1.
	data, err := io.ReadAll(io.LimitReader(binResp.Body, maxWhatsAppMediaBytes+1))
	if err != nil {
		return nil, fmt.Errorf("whatsapp: read media bin: %w", err)
	}
	if len(data) > maxWhatsAppMediaBytes {
		return nil, fmt.Errorf("whatsapp: mídia excede limite de %d bytes", maxWhatsAppMediaBytes)
	}

	mime := meta.MIMEType
	if i := strings.Index(mime, ";"); i >= 0 { // "audio/ogg; codecs=opus" → "audio/ogg"
		mime = strings.TrimSpace(mime[:i])
	}
	return &MediaDownload{Bytes: data, MIME: mime, FileSize: int64(len(data)), SHA256: meta.SHA256}, nil
}

// isMetaMediaHost valida que a URL de download de mídia é HTTPS e aponta pra um
// host conhecido da Meta/Facebook (CDN), antes de anexar o Bearer token.
func isMetaMediaHost(rawURL string) bool {
	u, err := url.Parse(rawURL)
	if err != nil || u.Scheme != "https" {
		return false
	}
	host := strings.ToLower(u.Hostname())
	for _, suffix := range []string{".fbcdn.net", ".facebook.com", ".fbsbx.com", ".whatsapp.net"} {
		if strings.HasSuffix(host, suffix) {
			return true
		}
	}
	return host == "lookaside.fbsbx.com"
}

// WhatsAppMediaTypeFromMIME mapeia um content-type pro tipo de mídia da Cloud API.
func WhatsAppMediaTypeFromMIME(mime string) string {
	switch {
	case strings.HasPrefix(mime, "image/"):
		return "image"
	case strings.HasPrefix(mime, "audio/"):
		return "audio"
	case strings.HasPrefix(mime, "video/"):
		return "video"
	default:
		return "document"
	}
}

// UploadMedia faz upload de um arquivo pra Meta e retorna o media_id (válido ~30 dias).
// POST /{phone-number-id}/media (multipart: messaging_product, type, file).
func (s *WhatsAppService) UploadMedia(data []byte, mime, filename string) (string, error) {
	if s.cfg.WhatsApp.PhoneNumberID == "" || s.cfg.WhatsApp.AccessToken == "" {
		log.Printf("📱 [WHATSAPP DEV] UploadMedia mime=%s bytes=%d (sem credenciais — log apenas)", mime, len(data))
		return "dev-media-id", nil
	}
	if filename == "" {
		filename = "arquivo"
	}

	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	_ = w.WriteField("messaging_product", "whatsapp")
	_ = w.WriteField("type", mime)
	part, err := w.CreateFormFile("file", filename)
	if err != nil {
		return "", fmt.Errorf("whatsapp: upload media form: %w", err)
	}
	if _, err := part.Write(data); err != nil {
		return "", fmt.Errorf("whatsapp: upload media write: %w", err)
	}
	if err := w.Close(); err != nil {
		return "", err
	}

	apiVersion := s.cfg.WhatsApp.GraphAPIVersion
	if apiVersion == "" {
		apiVersion = "v18.0"
	}
	url := fmt.Sprintf("https://graph.facebook.com/%s/%s/media", apiVersion, s.cfg.WhatsApp.PhoneNumberID)
	req, err := http.NewRequest(http.MethodPost, url, &buf)
	if err != nil {
		return "", fmt.Errorf("whatsapp: build upload request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+s.cfg.WhatsApp.AccessToken)
	req.Header.Set("Content-Type", w.FormDataContentType())

	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("whatsapp: upload media http: %w", err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(io.LimitReader(resp.Body, 2048))
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("whatsapp: upload media status %d: %s", resp.StatusCode, string(respBody))
	}
	var out struct {
		ID string `json:"id"`
	}
	if err := json.Unmarshal(respBody, &out); err != nil || out.ID == "" {
		return "", fmt.Errorf("whatsapp: upload media sem id: %s", string(respBody))
	}
	return out.ID, nil
}

// SendMediaMessage envia uma mensagem de mídia (image|document|audio|video) por
// media_id já enviado via UploadMedia. Só dentro da janela de 24h (regra Meta).
// Retorna o wa_message_id.
func (s *WhatsAppService) SendMediaMessage(toE164, waType, mediaID, caption, filename string) (string, error) {
	if s.cfg.WhatsApp.PhoneNumberID == "" || s.cfg.WhatsApp.AccessToken == "" {
		log.Printf("📱 [WHATSAPP DEV] SendMediaMessage to=%s type=%s (sem credenciais — log apenas)", toE164, waType)
		return "dev-" + toE164, nil
	}
	to, err := NormalizeE164(toE164)
	if err != nil {
		return "", err
	}

	media := map[string]any{"id": mediaID}
	// caption só vale pra image/video/document; audio ignora.
	if caption != "" && waType != "audio" {
		media["caption"] = caption
	}
	if waType == "document" && filename != "" {
		media["filename"] = filename
	}
	payload := map[string]any{
		"messaging_product": "whatsapp",
		"recipient_type":    "individual",
		"to":                to,
		"type":              waType,
		waType:              media,
	}

	raw, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("whatsapp: marshal media payload: %w", err)
	}
	apiVersion := s.cfg.WhatsApp.GraphAPIVersion
	if apiVersion == "" {
		apiVersion = "v18.0"
	}
	url := fmt.Sprintf("https://graph.facebook.com/%s/%s/messages", apiVersion, s.cfg.WhatsApp.PhoneNumberID)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(raw))
	if err != nil {
		return "", fmt.Errorf("whatsapp: build media msg request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+s.cfg.WhatsApp.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("whatsapp: media msg http: %w", err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(io.LimitReader(resp.Body, 2048))
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("whatsapp: media msg status %d: %s", resp.StatusCode, string(respBody))
	}
	var meta struct {
		Messages []struct {
			ID string `json:"id"`
		} `json:"messages"`
	}
	_ = json.Unmarshal(respBody, &meta)
	var msgID string
	if len(meta.Messages) > 0 {
		msgID = meta.Messages[0].ID
	}
	log.Printf("📱 [WHATSAPP] media (%s) sent to=%s message_id=%s", waType, to, msgID)
	return msgID, nil
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
