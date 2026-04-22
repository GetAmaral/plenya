package services

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"

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

// EmailService envia emails transacionais via SMTP. Em desenvolvimento,
// se SMTP_HOST não estiver configurado, faz log no stdout (sem enviar).
type EmailService struct {
	cfg *config.Config
}

// NewEmailService cria um novo EmailService
func NewEmailService(cfg *config.Config) *EmailService {
	return &EmailService{cfg: cfg}
}

// SendMagicLink envia o magic link do Escore Light por email.
// O link aponta pra /escore-plenya/claim/:token no site público.
func (s *EmailService) SendMagicLink(toEmail, magicLink string) error {
	subject := "Acesse seu Escore Plenya Light"
	siteURL := s.cfg.Site.PublicURL
	if siteURL == "" {
		siteURL = "https://plenyasaude.com.br"
	}
	body := fmt.Sprintf(`Olá,

Você solicitou guardar seu resultado do Escore Plenya Light.
Clique no link abaixo para acessar e salvar seu radar:

%s

O link expira em 15 minutos.

Se você não solicitou, ignore este email.

— Equipe Plenya
%s

────────────────────────────────────────
LGPD · Seus direitos
Você pode solicitar acesso, correção, portabilidade ou exclusão dos seus
dados a qualquer momento em: %s/lgpd/direitos
Encarregado de Proteção de Dados (DPO): dpo@plenyasaude.com.br
`, magicLink, siteURL, siteURL)

	return s.send(toEmail, subject, body)
}

// send é o envio low-level via SMTP. Usa fallback console em dev.
func (s *EmailService) send(toEmail, subject, body string) error {
	host := s.cfg.Email.SMTPHost
	from := s.cfg.Email.FromAddress
	if from == "" {
		from = "no-reply@plenyasaude.com.br"
	}

	if host == "" {
		// Dev fallback: só loga metadados (sem body) e o magic link, que é o único campo
		// que o dev precisa para clicar e testar. Não logamos o body completo (LGPD —
		// embora o magic link em si não seja dado sensível, evita vazar conteúdo
		// futuro de outros emails que adicionarmos).
		log.Printf("📧 [EMAIL DEV] To: %s | Subject: %s | Body: %d bytes (%d linhas)",
			toEmail, subject, len(body), countLines(body))
		// Para magic links, é útil ver o link no console em dev:
		if link := extractMagicLink(body); link != "" {
			log.Printf("📧 [EMAIL DEV] Magic link: %s", link)
		}
		return nil
	}

	addr := fmt.Sprintf("%s:%s", host, s.cfg.Email.SMTPPort)
	auth := smtp.PlainAuth("", s.cfg.Email.SMTPUser, s.cfg.Email.SMTPPass, host)

	msg := []byte(fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/plain; charset=utf-8\r\n\r\n%s",
		from, toEmail, subject, body,
	))

	return smtp.SendMail(addr, auth, from, []string{toEmail}, msg)
}
