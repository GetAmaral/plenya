package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"net/smtp"
	"strings"
	"time"

	"github.com/plenya/api/internal/config"
)

// countLines conta linhas em string (helper para log resumido).
func countLines(s string) int {
	if s == "" {
		return 0
	}
	return strings.Count(s, "\n") + 1
}

// extractMagicLink procura uma URL https? no body. Útil só em dev para testar magic links.
func extractMagicLink(body string) string {
	for _, line := range strings.Split(body, "\n") {
		l := strings.TrimSpace(line)
		if strings.HasPrefix(l, "http://") || strings.HasPrefix(l, "https://") {
			return l
		}
	}
	return ""
}

// EmailService envia emails transacionais. Provider escolhe a implementação concreta:
// "resend" usa a API REST do Resend; "smtp" usa net/smtp; vazio cai no fallback dev (log no stdout).
type EmailService struct {
	cfg    *config.Config
	client *http.Client
}

// NewEmailService cria um novo EmailService
func NewEmailService(cfg *config.Config) *EmailService {
	return &EmailService{
		cfg:    cfg,
		client: &http.Client{Timeout: 10 * time.Second},
	}
}

// SendMagicLink envia o magic link do Escore Light.
// Gera versões text + HTML simples; provider decide o que usar.
func (s *EmailService) SendMagicLink(toEmail, magicLink string) error {
	subject := "Acesse seu Escore Plenya Light"
	siteURL := s.cfg.Site.PublicURL
	if siteURL == "" {
		siteURL = "https://plenyasaude.com.br"
	}
	bodyText := fmt.Sprintf(`Olá,

Você solicitou guardar seu resultado do Escore Plenya Light.
Clique no link abaixo para acessar e salvar seu radar:

%s

O link expira em 7 dias.

Se você não solicitou, ignore este email.

— Equipe Plenya
%s

────────────────────────────────────────
LGPD · Seus direitos
Você pode solicitar acesso, correção, portabilidade ou exclusão dos seus
dados a qualquer momento em: %s/lgpd/direitos
Encarregado de Proteção de Dados (DPO): dpo@plenyasaude.com.br
`, magicLink, siteURL, siteURL)

	// Escapa HTML em todos os valores interpolados — defesa em profundidade contra XSS
	// caso magicLink ou siteURL sejam manipulados em algum caminho futuro.
	escLink := html.EscapeString(magicLink)
	escSite := html.EscapeString(siteURL)
	bodyHTML := fmt.Sprintf(`<!DOCTYPE html>
<html><body style="font-family: Georgia, serif; color: #1f3640; max-width: 560px; margin: 32px auto; padding: 0 16px; line-height: 1.6;">
<p>Olá,</p>
<p>Você solicitou guardar seu resultado do <strong>Escore Plenya Light</strong>.<br/>
Clique no botão abaixo para acessar e salvar seu radar:</p>
<p style="margin: 32px 0;"><a href="%s" style="background:#c19a4a; color:#fff8eb; padding: 12px 24px; text-decoration: none; display: inline-block; letter-spacing: 0.5px;">Acessar meu resultado</a></p>
<p style="font-size: 14px; color: #4a6478;">Ou copie este link: <br/><a href="%s" style="color:#c19a4a; word-break: break-all;">%s</a></p>
<p style="font-size: 14px; color: #4a6478;">O link expira em <strong>7 dias</strong>. Se você não solicitou, ignore este email.</p>
<p style="margin-top: 40px;">— Equipe Plenya<br/><a href="%s" style="color:#4a6478;">%s</a></p>
<hr style="border:none; border-top: 1px solid #e6dfd1; margin: 32px 0;"/>
<p style="font-size: 12px; color: #6b7c8a;"><strong>LGPD · Seus direitos</strong><br/>
Acesso, correção, portabilidade ou exclusão dos seus dados em <a href="%s/lgpd/direitos" style="color:#4a6478;">%s/lgpd/direitos</a><br/>
Encarregado de Proteção de Dados (DPO): <a href="mailto:dpo@plenyasaude.com.br" style="color:#4a6478;">dpo@plenyasaude.com.br</a></p>
</body></html>`, escLink, escLink, escLink, escSite, escSite, escSite, escSite)

	return s.send(toEmail, subject, bodyText, bodyHTML)
}

// send despacha pro provider configurado. Em dev (sem credenciais) loga metadados + magic link.
func (s *EmailService) send(toEmail, subject, bodyText, bodyHTML string) error {
	provider := strings.ToLower(s.cfg.Email.Provider)

	switch provider {
	case "resend":
		if s.cfg.Email.ResendAPIKey == "" {
			s.devLog(toEmail, subject, bodyText, "resend (sem API key)")
			return nil
		}
		return s.sendViaResend(toEmail, subject, bodyText, bodyHTML)

	case "smtp", "":
		if s.cfg.Email.SMTPHost == "" {
			s.devLog(toEmail, subject, bodyText, "smtp (sem host)")
			return nil
		}
		return s.sendViaSMTP(toEmail, subject, bodyText)

	default:
		return fmt.Errorf("email: provider desconhecido %q (use 'resend' ou 'smtp')", provider)
	}
}

func (s *EmailService) devLog(toEmail, subject, bodyText, providerHint string) {
	log.Printf("📧 [EMAIL DEV - %s] To: %s | Subject: %s | Body: %d bytes (%d linhas)",
		providerHint, toEmail, subject, len(bodyText), countLines(bodyText))
	if link := extractMagicLink(bodyText); link != "" {
		log.Printf("📧 [EMAIL DEV] Magic link: %s", link)
	}
}

// sendViaSMTP — net/smtp (legacy/fallback).
func (s *EmailService) sendViaSMTP(toEmail, subject, bodyText string) error {
	from := s.cfg.Email.FromAddress
	if from == "" {
		from = "no-reply@plenyasaude.com.br"
	}
	addr := fmt.Sprintf("%s:%s", s.cfg.Email.SMTPHost, s.cfg.Email.SMTPPort)
	auth := smtp.PlainAuth("", s.cfg.Email.SMTPUser, s.cfg.Email.SMTPPass, s.cfg.Email.SMTPHost)
	msg := []byte(fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/plain; charset=utf-8\r\n\r\n%s",
		from, toEmail, subject, bodyText,
	))
	return smtp.SendMail(addr, auth, from, []string{toEmail}, msg)
}

// sendViaResend — POST https://api.resend.com/emails. ~30 linhas, sem SDK.
func (s *EmailService) sendViaResend(toEmail, subject, bodyText, bodyHTML string) error {
	from := s.cfg.Email.FromAddress
	if from == "" {
		from = "no-reply@plenyasaude.com.br"
	}
	if name := s.cfg.Email.FromName; name != "" {
		from = fmt.Sprintf("%s <%s>", name, from)
	}

	payload := map[string]any{
		"from":    from,
		"to":      []string{toEmail},
		"subject": subject,
		"text":    bodyText,
	}
	if bodyHTML != "" {
		payload["html"] = bodyHTML
	}

	raw, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("resend: marshal payload: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, "https://api.resend.com/emails", bytes.NewReader(raw))
	if err != nil {
		return fmt.Errorf("resend: build request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+s.cfg.Email.ResendAPIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("resend: http: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	body, _ := io.ReadAll(io.LimitReader(resp.Body, 1024))
	return fmt.Errorf("resend: status %d: %s", resp.StatusCode, string(body))
}
