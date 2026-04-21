package services

import (
	"fmt"
	"log"
	"net/smtp"

	"github.com/plenya/api/internal/config"
)

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
	body := fmt.Sprintf(`Olá,

Você solicitou guardar seu resultado do Escore Plenya Light.
Clique no link abaixo para acessar e salvar seu radar:

%s

O link expira em 15 minutos.

Se você não solicitou, ignore este email.

— Equipe Plenya
plenyasaude.com.br
`, magicLink)

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
		// Dev fallback: só loga
		log.Printf("📧 [EMAIL DEV] To: %s | Subject: %s\n---\n%s\n---", toEmail, subject, body)
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
