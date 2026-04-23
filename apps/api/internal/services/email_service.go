package services

import (
	"bytes"
	"embed"
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

// emailTemplates é o conjunto de templates pré-renderizados pelo packages/emails (React Email).
// Sync via `pnpm email:sync` (copia packages/emails/dist/*.html → templates/).
//
//go:embed templates/*.html
var emailTemplates embed.FS

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

	bodyHTML, err := s.renderTemplate("magic_link", map[string]string{
		"LINK":     magicLink,
		"SITE_URL": siteURL,
	})
	if err != nil {
		// Fallback: log e segue sem HTML (texto puro vai pro provider)
		log.Printf("⚠️  [EMAIL] template magic_link falhou: %v — usando texto plano", err)
		bodyHTML = ""
	}

	return s.send(toEmail, subject, bodyText, bodyHTML)
}

// renderTemplate carrega template pré-renderizado e substitui placeholders {{KEY}}.
// Valores são escapados via html.EscapeString (defesa XSS).
func (s *EmailService) renderTemplate(name string, vars map[string]string) (string, error) {
	raw, err := emailTemplates.ReadFile("templates/" + name + ".html")
	if err != nil {
		return "", fmt.Errorf("email template %q: %w", name, err)
	}
	out := string(raw)
	for k, v := range vars {
		out = strings.ReplaceAll(out, "{{"+k+"}}", html.EscapeString(v))
	}
	return out, nil
}

// SendBoasVindas envia email de boas-vindas pós-conversão (Lead → Patient).
func (s *EmailService) SendBoasVindas(toEmail, patientName string) error {
	subject := "Bem-vindo à Plenya"
	siteURL := s.cfg.Site.PublicURL
	if siteURL == "" {
		siteURL = "https://plenyasaude.com.br"
	}
	bodyText := fmt.Sprintf(`Olá, %s.

Sua conta na Plenya está pronta. Agora você pode acompanhar seu Escore Plenya Light,
refazer a avaliação a cada 3 meses e ver sua evolução ao longo do tempo.

Quando quiser conversar com a equipe — sobre uma avaliação completa no Continuum, ou
sobre os pontos do seu radar que mais chamaram atenção — estamos no WhatsApp ou no
email contato@plenyasaude.com.br.

— Equipe Plenya
%s
`, patientName, siteURL)

	bodyHTML, err := s.renderTemplate("boas_vindas", map[string]string{
		"NAME":     patientName,
		"SITE_URL": siteURL,
	})
	if err != nil {
		log.Printf("⚠️  [EMAIL] template boas_vindas falhou: %v", err)
		bodyHTML = ""
	}
	return s.send(toEmail, subject, bodyText, bodyHTML)
}

// SendFollowUp30Dias envia follow-up manual (admin dispara via UI).
func (s *EmailService) SendFollowUp30Dias(toEmail, patientName string) error {
	subject := "Refaça seu Escore — veja sua evolução"
	siteURL := s.cfg.Site.PublicURL
	if siteURL == "" {
		siteURL = "https://plenyasaude.com.br"
	}
	bodyText := fmt.Sprintf(`Olá, %s.

Faz 30 dias que você fez seu Escore Plenya Light. Pequenas mudanças no dia-a-dia
já podem aparecer no radar — sono, alimentação, atividade física são os primeiros a
mexer.

Vale refazer a avaliação? É grátis e leva 7 minutos:
%s/escore-plenya/avaliar

— Equipe Plenya
`, patientName, siteURL)

	bodyHTML, err := s.renderTemplate("follow_up_30_dias", map[string]string{
		"NAME":     patientName,
		"SITE_URL": siteURL,
	})
	if err != nil {
		log.Printf("⚠️  [EMAIL] template follow_up_30_dias falhou: %v", err)
		bodyHTML = ""
	}
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
