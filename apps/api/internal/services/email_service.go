package services

import (
	"bytes"
	"context"
	"crypto/rand"
	"embed"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"io"
	"log"
	"mime"
	"net/http"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/emersion/go-imap/v2"
	"github.com/emersion/go-imap/v2/imapclient"
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
)

// XPlenyaSourceHeader é o header marker que identifica emails originados pelo EMR.
// O worker IMAP IDLE (Bloco 1) usa esse header pra pular emails que aparecem no Sent Items
// e já foram registrados como LeadActivity pelo SendLeadReply (evita dedup duplo).
const (
	XPlenyaSourceHeader   = "X-Plenya-Source"
	XPlenyaSourceEMR      = "emr"
	imapAppendTimeout     = 15 * time.Second
	defaultSentFolderName = "Sent Items"
)

// ErrLeadOptedOut é retornado quando se tenta enviar email pra Lead com EmailOptIn=false.
var ErrLeadOptedOut = errors.New("email: lead opt-out de email")

// ErrLeadNoEmail é retornado quando o Lead não tem endereço de email cadastrado.
var ErrLeadNoEmail = errors.New("email: lead sem endereço de email cadastrado")

// ErrEmailAttachmentTooLarge → soma dos anexos passou do limite Resend (40MB).
var ErrEmailAttachmentTooLarge = errors.New("email: soma de anexos excede 40MB (limite Resend)")

// ErrEmailAttachmentInvalid → path traversal, arquivo não existe, extensão fora da whitelist.
var ErrEmailAttachmentInvalid = errors.New("email: anexo inválido")

// emailAttachmentRoot é a raiz onde upload handler grava arquivos. Anexos com Path
// fora desse prefixo são rejeitados (defesa contra path traversal).
const emailAttachmentRoot = "/app/uploads"

// resendMaxTotalBytes é o limite duro do Resend pra soma dos anexos por email.
// Validamos antes de fazer base64 + POST pra falhar rápido com erro semântico.
const resendMaxTotalBytes = 40 * 1024 * 1024

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
//
// Opcional: db é usado pelo SendLeadReply pra (a) registrar LeadActivity outbound,
// (b) descobrir Subject anterior pra montar "Re: ...". Se nil, SendLeadReply falha.
type EmailService struct {
	cfg    *config.Config
	client *http.Client
	db     *gorm.DB
}

// NewEmailService cria um novo EmailService.
// db é opcional — só obrigatório pra SendLeadReply.
func NewEmailService(cfg *config.Config) *EmailService {
	return &EmailService{
		cfg:    cfg,
		client: &http.Client{Timeout: 10 * time.Second},
	}
}

// WithDB devolve a mesma instância configurada com handle GORM.
// Permite construção encadeada sem mudar a assinatura legada de NewEmailService.
func (s *EmailService) WithDB(db *gorm.DB) *EmailService {
	s.db = db
	return s
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

// SendPortalInvite envia convite de acesso à área do paciente (minha.plenyasaude.com.br).
// Disparado pela equipe (admin/manager/secretary) no perfil do Patient.
func (s *EmailService) SendPortalInvite(toEmail, patientName, link string) error {
	subject := "Bem-vindo à área do paciente Plenya"
	bodyText := fmt.Sprintf(`Olá, %s.

A equipe Plenya criou seu acesso à área do paciente.

Lá você acompanha seu Continuum, próximas consultas, exames, prescrições e
conversa diretamente com a equipe.

Acesse pelo link (válido por 14 dias):

%s

Se você não esperava este convite, ignore este email.

— Equipe Plenya
`, patientName, link)

	bodyHTML, err := s.renderTemplate("portal_invite", map[string]string{
		"NAME": patientName,
		"LINK": link,
	})
	if err != nil {
		// Template ainda não criado — segue só com texto plano
		bodyHTML = ""
	}
	return s.send(toEmail, subject, bodyText, bodyHTML)
}

// SendPatientMagicLink envia magic link de login pra paciente já cadastrado.
// Usado em "entrar com link mágico" e "esqueci a senha".
func (s *EmailService) SendPatientMagicLink(toEmail, patientName, link string) error {
	subject := "Seu acesso à área Plenya"
	bodyText := fmt.Sprintf(`Olá, %s.

Use o link abaixo pra entrar na sua área Plenya (válido por 15 minutos):

%s

Se você não solicitou este link, ignore este email — sua conta segue protegida.

— Equipe Plenya
`, patientName, link)

	bodyHTML, err := s.renderTemplate("patient_magic_link", map[string]string{
		"NAME": patientName,
		"LINK": link,
	})
	if err != nil {
		bodyHTML = ""
	}
	return s.send(toEmail, subject, bodyText, bodyHTML)
}

// SendFamilyInvite envia convite pra familiar acessar dados de outro paciente.
func (s *EmailService) SendFamilyInvite(toEmail, granteeLabel, patientName, link string) error {
	subject := fmt.Sprintf("%s te convidou pra acompanhar a saúde", patientName)
	bodyText := fmt.Sprintf(`Olá%s,

%s compartilhou acesso ao perfil de saúde Plenya com você.

Clique no link pra aceitar o convite (válido por 14 dias):

%s

Se não esperava receber este convite, ignore este email.

— Equipe Plenya
`, optLabel(granteeLabel), patientName, link)
	return s.send(toEmail, subject, bodyText, "")
}

func optLabel(label string) string {
	label = strings.TrimSpace(label)
	if label == "" {
		return ""
	}
	return ", " + label
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

// ============================================================
// SendLeadReply / SendConversationReply (Bloco 3 + B) — resposta outbound por email
// ============================================================

// ConversationReplyOwner identifica o destinatário de uma resposta de email — Lead OU
// Patient (mutuamente exclusivo). Espelha services.ConversationOwner mas mantida aqui
// pra evitar import cycle (lead_service.go declara ConversationOwner, ambos no mesmo
// package services). NOTA: como estão no mesmo package, podemos usar ConversationOwner
// diretamente — mantemos esse alias só pra clareza no payload.
type ConversationReplyOwner = ConversationOwner

// SendLeadReplyInput é o payload de SendLeadReply.
//
// Lead carrega o destinatário e o opt-in. ActorUserID é quem está respondendo (do JWT).
// Subject é opcional — se vazio, monta "Re: <thread anterior>" via lookup das activities.
// BodyHTML/BodyText são alternativos: se BodyHTML vazio, gera versão simples a partir do
// BodyText (escape + nl2br + autolink). Se BodyText vazio, deriva do BodyHTML via
// ExtractPlainBody (best-effort).
type SendLeadReplyInput struct {
	Lead        *models.Lead
	ActorUserID uuid.UUID
	Subject     string
	BodyHTML    string
	BodyText    string
	InReplyTo   string
	References  []string
}

// SendConversationReplyInput é o payload genérico (Lead OU Patient) usado pela
// Central de Conversas (Bloco B). SendLeadReply delega aqui.
type SendConversationReplyInput struct {
	Owner       ConversationOwner
	ActorUserID uuid.UUID
	Subject     string
	BodyHTML    string
	BodyText    string
	InReplyTo   string
	References  []string
	Attachments []ReplyAttachment
}

// ReplyAttachment é a referência a um arquivo em /app/uploads — service lê o
// conteúdo, encoda base64 e envia no payload Resend. Filename é o nome
// apresentado ao destinatário (e gravado em metadata).
type ReplyAttachment struct {
	Path     string
	Filename string
}

// SendLeadReply mantém compat com handlers/lead_handler.go (Bloco 3 anterior).
// Internamente delega pra SendConversationReply. Não suporta attachments —
// quem precisar passa pelo ConversationService.
func (s *EmailService) SendLeadReply(ctx context.Context, in SendLeadReplyInput) error {
	return s.SendConversationReply(ctx, SendConversationReplyInput{
		Owner:       ConversationOwner{Lead: in.Lead},
		ActorUserID: in.ActorUserID,
		Subject:     in.Subject,
		BodyHTML:    in.BodyHTML,
		BodyText:    in.BodyText,
		InReplyTo:   in.InReplyTo,
		References:  in.References,
	})
}

// SendConversationReply envia email transacional via Resend, faz IMAP APPEND (best-effort)
// no Sent Items do Stalwart e cria LeadActivity{message_sent, channel=email} anexada ao
// owner correto (Lead ou Patient).
//
// Header X-Plenya-Source: emr é setado pra que o worker IMAP IDLE pule a mensagem
// quando aparecer em Sent Items (evita registro duplicado).
//
// Errors:
//   - ErrLeadNoEmail: Owner sem email cadastrado
//   - ErrLeadOptedOut: Lead.EmailOptIn == false (Patient não tem flag — assume opt-in implícito)
//   - resend HTTP error → bloqueia (envio falhou)
//   - IMAP APPEND falha → log warn, NÃO bloqueia (best-effort, async via goSafe)
func (s *EmailService) SendConversationReply(ctx context.Context, in SendConversationReplyInput) error {
	if s.db == nil {
		return errors.New("email: SendConversationReply requer DB (use WithDB ao construir)")
	}
	if in.ActorUserID == uuid.Nil {
		return errors.New("email: actorUserID obrigatório")
	}
	hasLead := in.Owner.Lead != nil
	hasPatient := in.Owner.Patient != nil
	if hasLead == hasPatient {
		return errors.New("email: owner deve ter exatamente um de Lead|Patient")
	}

	// Resolve destinatário (email + nome) e identidade pra threading
	var (
		to         string
		toName     string
		ownerID    uuid.UUID
		leadIDPtr  *uuid.UUID
		patIDPtr   *uuid.UUID
	)
	if hasLead {
		lead := in.Owner.Lead
		if lead.Email == nil || strings.TrimSpace(*lead.Email) == "" {
			return ErrLeadNoEmail
		}
		if !lead.EmailOptIn {
			return ErrLeadOptedOut
		}
		to = strings.TrimSpace(*lead.Email)
		if lead.Name != nil {
			toName = strings.TrimSpace(*lead.Name)
		}
		ownerID = lead.ID
		id := lead.ID
		leadIDPtr = &id
	} else {
		patient := in.Owner.Patient
		if patient.Email == nil || strings.TrimSpace(*patient.Email) == "" {
			return ErrLeadNoEmail
		}
		// Patient não tem flag explícita de opt-in; assumimos consent implícito
		// (signup do EMR coleta consent geral).
		to = strings.TrimSpace(*patient.Email)
		toName = strings.TrimSpace(patient.Name)
		ownerID = patient.ID
		id := patient.ID
		patIDPtr = &id
	}

	// Resolve Subject: se vazio, busca último email da thread.
	subject := strings.TrimSpace(in.Subject)
	if subject == "" {
		subject = s.resolveReplySubjectForOwner(leadIDPtr, patIDPtr)
	}
	_ = ownerID

	// Resolve corpo: ao menos um de BodyText/BodyHTML
	bodyText := in.BodyText
	bodyHTML := in.BodyHTML
	if strings.TrimSpace(bodyText) == "" && strings.TrimSpace(bodyHTML) == "" {
		return errors.New("email: corpo vazio (bodyText ou bodyHTML obrigatório)")
	}
	if strings.TrimSpace(bodyHTML) == "" {
		bodyHTML = renderPlainToHTML(bodyText)
	}
	if strings.TrimSpace(bodyText) == "" {
		bodyText = ExtractPlainBody(bodyHTML)
	}

	// From: usa LeadReplyFromAddress (default contato@) — caixa monitorada que recebe
	// respostas do cliente e fecha o ciclo via worker IMAP IDLE. NÃO usa FromAddress
	// (noreply@) que é só pra transacional.
	fromAddr := s.cfg.Email.LeadReplyFromAddress
	if fromAddr == "" {
		fromAddr = s.cfg.Email.FromAddress
	}
	if fromAddr == "" {
		fromAddr = "contato@plenyasaude.com.br"
	}
	fromName := s.cfg.Email.FromName
	fromHeader := fromAddr
	if fromName != "" {
		fromHeader = fmt.Sprintf("%s <%s>", fromName, fromAddr)
	}

	// Message-ID novo. Domínio: pega do FromAddress (após @), fallback "plenyasaude.com.br".
	domain := domainFromAddress(fromAddr)
	messageID := generateMessageID(domain)

	// Resolve attachments: lê arquivos do disco, encoda base64, valida total < 40MB.
	// Falha precoce se algum path é inválido ou soma estoura — antes de bater Resend.
	resolvedAtts, err := s.resolveAttachments(in.Attachments)
	if err != nil {
		return err
	}

	// Monta MIME pra IMAP APPEND. Headers cobrem threading + X-Plenya-Source.
	// MIME atual não inclui anexos no espelho IMAP — Sent Items mostra só corpo.
	// Trade-off aceitável: o histórico do EMR (LeadActivity.metadata.attachments)
	// é a fonte oficial de auditoria; mailbox SMTP é só backup pra caixa do agente.
	rawMIME := buildReplyMIME(replyMIMEInput{
		FromHeader: fromHeader,
		ToAddress:  to,
		ToName:     toName,
		Subject:    subject,
		MessageID:  messageID,
		InReplyTo:  in.InReplyTo,
		References: in.References,
		BodyText:   bodyText,
		BodyHTML:   bodyHTML,
		Date:       time.Now().UTC(),
	})

	// 1) Envia via Resend (BLOQUEANTE — falha aqui aborta o registro de activity)
	resendID, err := s.sendReplyViaResend(ctx, sendReplyResendInput{
		FromHeader:  fromHeader,
		To:          to,
		Subject:     subject,
		BodyText:    bodyText,
		BodyHTML:    bodyHTML,
		MessageID:   messageID,
		InReplyTo:   in.InReplyTo,
		References:  in.References,
		Attachments: resolvedAtts,
	})
	if err != nil {
		return fmt.Errorf("email: resend: %w", err)
	}

	// 2) IMAP APPEND no Sent — best-effort, async, não bloqueia caller.
	mirroredCh := make(chan bool, 1)
	goSafe("email_imap_append", func() {
		ok := s.appendToSentMailbox(rawMIME)
		mirroredCh <- ok
	})
	// Não esperamos a goroutine — ela só atualiza metadata posteriormente?
	// Pra simplicidade, NÃO bloqueia activity. Marcamos mirrored_to_imap=null e o worker
	// IMAP eventualmente vai ver o email (com X-Plenya-Source) e skip → activity já existe.
	// Drenamos o canal pra evitar leak (goroutine sempre escreve).
	go func() { <-mirroredCh }()

	// 3) Cria LeadActivity message_sent
	metadata := map[string]any{
		"subject":            subject,
		"message_id":         messageID,
		"recipient":          to,
		"resend_id":          resendID,
		"sent_at":            time.Now().UTC(),
		"source":             XPlenyaSourceEMR,
	}
	if in.InReplyTo != "" {
		metadata["in_reply_to"] = in.InReplyTo
	}
	if len(in.References) > 0 {
		metadata["references"] = in.References
	}
	if len(resolvedAtts) > 0 {
		// Mesmo schema dos inbound (EmailAttachmentRef): {filename, path, content_type, size_bytes}.
		// Frontend lê msg.metadata.attachments na timeline e renderiza chips clicáveis.
		attsMeta := make([]map[string]any, len(resolvedAtts))
		for i, a := range resolvedAtts {
			attsMeta[i] = map[string]any{
				"filename":     a.Filename,
				"path":         a.Path,
				"content_type": a.ContentType,
				"size_bytes":   a.SizeBytes,
			}
		}
		metadata["attachments"] = attsMeta
	}

	metaJSON, err := json.Marshal(metadata)
	if err != nil {
		return fmt.Errorf("email: marshal metadata: %w", err)
	}
	activity := models.LeadActivity{
		LeadID:      leadIDPtr,
		PatientID:   patIDPtr,
		Type:        models.LeadActivityMessageSent,
		Channel:     models.LeadChannelEmail,
		Content:     &bodyText,
		Metadata:    datatypes.JSON(metaJSON),
		ActorUserID: &in.ActorUserID,
	}
	if err := s.db.Create(&activity).Error; err != nil {
		return fmt.Errorf("email: persist activity: %w", err)
	}
	return nil
}

// resolveReplySubject busca a última activity de email no Lead (sent ou received) e
// retorna "Re: <subject anterior>" (ou apenas o subject se já começa com "Re:").
// Sem histórico, devolve fallback genérico.
func (s *EmailService) resolveReplySubject(leadID uuid.UUID) string {
	return s.resolveReplySubjectForOwner(&leadID, nil)
}

// resolveReplySubjectForOwner é a versão genérica que busca na thread de Lead OU Patient.
func (s *EmailService) resolveReplySubjectForOwner(leadID, patientID *uuid.UUID) string {
	const fallback = "Sobre seu contato com Plenya"
	if s.db == nil {
		return fallback
	}
	q := s.db.Model(&models.LeadActivity{}).
		Where("channel = ? AND type IN ?",
			models.LeadChannelEmail,
			[]models.LeadActivityType{models.LeadActivityMessageReceived, models.LeadActivityMessageSent})
	switch {
	case leadID != nil && *leadID != uuid.Nil:
		q = q.Where("lead_id = ?", *leadID)
	case patientID != nil && *patientID != uuid.Nil:
		q = q.Where("patient_id = ?", *patientID)
	default:
		return fallback
	}
	var act models.LeadActivity
	if err := q.Order("created_at DESC").First(&act).Error; err != nil {
		return fallback
	}
	if len(act.Metadata) == 0 {
		return fallback
	}
	var meta struct {
		Subject string `json:"subject"`
	}
	if err := json.Unmarshal(act.Metadata, &meta); err != nil {
		return fallback
	}
	subj := strings.TrimSpace(meta.Subject)
	if subj == "" {
		return fallback
	}
	low := strings.ToLower(subj)
	if strings.HasPrefix(low, "re:") || strings.HasPrefix(low, "res:") {
		return subj
	}
	return "Re: " + subj
}

// resolvedAttachment é o anexo já lido do disco e pronto pra Resend.
// Path/Filename/ContentType/SizeBytes preservados pra metadata da activity.
type resolvedAttachment struct {
	Path        string // relativo a /app/uploads — preservado em metadata
	Filename    string // nome mostrado ao destinatário
	ContentType string // content-type detectado por extensão
	SizeBytes   int    // tamanho original (não base64)
	Base64      string // conteúdo base64 padrão (sem URL-safe)
}

// resolveAttachments lê cada anexo do disco, detecta content-type por extensão,
// soma o total e valida contra o limite Resend (40MB). Retorna ErrEmailAttachmentInvalid
// para path traversal/arquivo ausente; ErrEmailAttachmentTooLarge para soma estourada.
//
// Nota: cleanPath impede path traversal — exigimos que o path resolva DENTRO de
// emailAttachmentRoot. Anexos vindos do upload handler já estão lá; rejeição é
// defesa contra payload mal-formado.
func (s *EmailService) resolveAttachments(in []ReplyAttachment) ([]resolvedAttachment, error) {
	if len(in) == 0 {
		return nil, nil
	}
	out := make([]resolvedAttachment, 0, len(in))
	var total int64
	for _, a := range in {
		clean, err := safeAttachmentPath(a.Path)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrEmailAttachmentInvalid, a.Path)
		}
		info, err := os.Stat(clean)
		if err != nil || info.IsDir() {
			return nil, fmt.Errorf("%w: %s não encontrado", ErrEmailAttachmentInvalid, a.Path)
		}
		total += info.Size()
		if total > resendMaxTotalBytes {
			return nil, ErrEmailAttachmentTooLarge
		}
		raw, err := os.ReadFile(clean)
		if err != nil {
			return nil, fmt.Errorf("%w: ler %s: %v", ErrEmailAttachmentInvalid, a.Path, err)
		}
		ct := mime.TypeByExtension(strings.ToLower(filepath.Ext(clean)))
		if ct == "" {
			ct = "application/octet-stream"
		}
		filename := strings.TrimSpace(a.Filename)
		if filename == "" {
			filename = filepath.Base(clean)
		}
		out = append(out, resolvedAttachment{
			Path:        a.Path,
			Filename:    filename,
			ContentType: ct,
			SizeBytes:   int(info.Size()),
			Base64:      base64.StdEncoding.EncodeToString(raw),
		})
	}
	return out, nil
}

// safeAttachmentPath resolve o path relativo dentro de emailAttachmentRoot e
// rejeita qualquer escape (.., absoluto que aponta fora, links simbólicos não checados — confiamos
// no fato do upload handler nunca cria symlinks).
func safeAttachmentPath(rel string) (string, error) {
	if rel == "" {
		return "", errors.New("empty path")
	}
	// Aceita tanto "conversation-outbound/..." quanto "/uploads/conversation-outbound/..."
	// — frontend já manda o que o upload handler retornou.
	rel = strings.TrimPrefix(rel, "/")
	rel = strings.TrimPrefix(rel, "uploads/")
	cleaned := filepath.Clean("/" + rel) // garante leading slash pra Clean resolver ".."
	abs := filepath.Join(emailAttachmentRoot, cleaned)
	// Confirma containment (defesa em profundidade).
	rootAbs, err := filepath.Abs(emailAttachmentRoot)
	if err != nil {
		return "", err
	}
	finalAbs, err := filepath.Abs(abs)
	if err != nil {
		return "", err
	}
	if !strings.HasPrefix(finalAbs+string(filepath.Separator), rootAbs+string(filepath.Separator)) &&
		finalAbs != rootAbs {
		return "", errors.New("path escape")
	}
	return finalAbs, nil
}

// sendReplyResendInput agrupa o que vai pro POST /emails da Resend.
type sendReplyResendInput struct {
	FromHeader  string
	To          string
	Subject     string
	BodyText    string
	BodyHTML    string
	MessageID   string
	InReplyTo   string
	References  []string
	Attachments []resolvedAttachment
}

// sendReplyViaResend faz POST /emails na Resend com headers de threading + X-Plenya-Source.
// Retorna o id retornado pela API ("" em dev fallback).
func (s *EmailService) sendReplyViaResend(ctx context.Context, in sendReplyResendInput) (string, error) {
	// Em dev sem API key, loga e retorna OK.
	if s.cfg.Email.ResendAPIKey == "" {
		log.Printf("📧 [EMAIL DEV] reply pra %s | subject=%q | %d bytes",
			in.To, in.Subject, len(in.BodyText))
		return "", nil
	}

	headers := map[string]string{
		XPlenyaSourceHeader: XPlenyaSourceEMR,
		"Message-ID":        in.MessageID,
	}
	if in.InReplyTo != "" {
		headers["In-Reply-To"] = in.InReplyTo
	}
	if len(in.References) > 0 {
		headers["References"] = strings.Join(in.References, " ")
	}

	payload := map[string]any{
		"from":    in.FromHeader,
		"to":      []string{in.To},
		"subject": in.Subject,
		"text":    in.BodyText,
		"html":    in.BodyHTML,
		"headers": headers,
	}
	if len(in.Attachments) > 0 {
		// Resend espera [{filename, content (base64), content_type?}].
		atts := make([]map[string]string, len(in.Attachments))
		for i, a := range in.Attachments {
			atts[i] = map[string]string{
				"filename":     a.Filename,
				"content":      a.Base64,
				"content_type": a.ContentType,
			}
		}
		payload["attachments"] = atts
	}

	raw, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("marshal: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost,
		"https://api.resend.com/emails", bytes.NewReader(raw))
	if err != nil {
		return "", fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+s.cfg.Email.ResendAPIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("http: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("status %d: %s", resp.StatusCode, string(body))
	}
	var respBody struct {
		ID string `json:"id"`
	}
	_ = json.Unmarshal(body, &respBody)
	return respBody.ID, nil
}

// appendToSentMailbox conecta TLS na 993, login com creds do MailIngestConfig,
// e faz APPEND da MIME completa em "Sent Items" com flag \Seen. Retorna true em sucesso.
//
// Best-effort: qualquer erro é logado e false retornado. Não panickea.
func (s *EmailService) appendToSentMailbox(rawMIME []byte) bool {
	mc := s.cfg.MailIngest
	if !mc.Enabled || mc.IMAPHost == "" || mc.Username == "" || mc.Password == "" {
		log.Printf("📧 [EMAIL APPEND] mail ingest desabilitado/sem creds — skip mirror")
		return false
	}
	addr := fmt.Sprintf("%s:%d", mc.IMAPHost, mc.IMAPPort)

	// timeout local — se Stalwart estiver lento, não trava goroutine
	done := make(chan bool, 1)
	go func() {
		done <- s.doAppend(addr, mc.Username, mc.Password, rawMIME)
	}()
	select {
	case ok := <-done:
		return ok
	case <-time.After(imapAppendTimeout):
		log.Printf("⚠️  [EMAIL APPEND] timeout %s — skip", imapAppendTimeout)
		return false
	}
}

func (s *EmailService) doAppend(addr, user, pass string, raw []byte) bool {
	c, err := imapclient.DialTLS(addr, &imapclient.Options{
		WordDecoder: &mime.WordDecoder{},
	})
	if err != nil {
		log.Printf("⚠️  [EMAIL APPEND] dial %s: %v", addr, err)
		return false
	}
	defer c.Close()

	if err := c.Login(user, pass).Wait(); err != nil {
		log.Printf("⚠️  [EMAIL APPEND] login: %v", err)
		return false
	}
	defer func() { _ = c.Logout().Wait() }()

	folder := s.resolveSentFolder(c)
	cmd := c.Append(folder, int64(len(raw)), &imap.AppendOptions{
		Flags: []imap.Flag{imap.FlagSeen},
	})
	if _, err := cmd.Write(raw); err != nil {
		log.Printf("⚠️  [EMAIL APPEND] write: %v", err)
		_ = cmd.Close()
		return false
	}
	if err := cmd.Close(); err != nil {
		log.Printf("⚠️  [EMAIL APPEND] close: %v", err)
		return false
	}
	if _, err := cmd.Wait(); err != nil {
		log.Printf("⚠️  [EMAIL APPEND] wait: %v", err)
		return false
	}
	log.Printf("📧 [EMAIL APPEND] espelhado em %s (%d bytes)", folder, len(raw))
	return true
}

// resolveSentFolder tenta detectar o nome correto da pasta de enviados via LIST.
// Stalwart costuma usar "Sent Items"; outros servidores variam.
func (s *EmailService) resolveSentFolder(c *imapclient.Client) string {
	candidates := []string{"Sent Items", "Sent", "Enviados", "INBOX.Sent"}
	listCmd := c.List("", "*", nil)
	mailboxes, err := listCmd.Collect()
	if err != nil {
		return defaultSentFolderName
	}
	have := make(map[string]bool, len(mailboxes))
	for _, mb := range mailboxes {
		have[strings.ToLower(mb.Mailbox)] = true
	}
	for _, candidate := range candidates {
		if have[strings.ToLower(candidate)] {
			return candidate
		}
	}
	return defaultSentFolderName
}

// ============================================================
// MIME helpers
// ============================================================

// replyMIMEInput agrega tudo necessário pra montar a mensagem RFC822.
type replyMIMEInput struct {
	FromHeader string
	ToAddress  string
	ToName     string
	Subject    string
	MessageID  string
	InReplyTo  string
	References []string
	BodyText   string
	BodyHTML   string
	Date       time.Time
}

// buildReplyMIME monta uma MIME multipart/alternative simples (text + html)
// com headers de threading + X-Plenya-Source. CRLF terminado conforme RFC 5322.
func buildReplyMIME(in replyMIMEInput) []byte {
	var b bytes.Buffer

	toHeader := in.ToAddress
	if in.ToName != "" {
		toHeader = fmt.Sprintf("%s <%s>", mimeEncodeWord(in.ToName), in.ToAddress)
	}

	boundary := "plenya-" + randomToken(16)

	writeHeader := func(k, v string) {
		b.WriteString(k)
		b.WriteString(": ")
		b.WriteString(v)
		b.WriteString("\r\n")
	}

	writeHeader("Date", in.Date.Format(time.RFC1123Z))
	writeHeader("From", in.FromHeader)
	writeHeader("To", toHeader)
	writeHeader("Subject", mimeEncodeWord(in.Subject))
	writeHeader("Message-ID", in.MessageID)
	if in.InReplyTo != "" {
		writeHeader("In-Reply-To", in.InReplyTo)
	}
	if len(in.References) > 0 {
		writeHeader("References", strings.Join(in.References, " "))
	}
	writeHeader(XPlenyaSourceHeader, XPlenyaSourceEMR)
	writeHeader("MIME-Version", "1.0")
	writeHeader("Content-Type", `multipart/alternative; boundary="`+boundary+`"`)
	b.WriteString("\r\n")

	// Plain text part
	b.WriteString("--")
	b.WriteString(boundary)
	b.WriteString("\r\n")
	b.WriteString("Content-Type: text/plain; charset=utf-8\r\n")
	b.WriteString("Content-Transfer-Encoding: 8bit\r\n\r\n")
	b.WriteString(normalizeNewlines(in.BodyText))
	b.WriteString("\r\n")

	// HTML part
	b.WriteString("--")
	b.WriteString(boundary)
	b.WriteString("\r\n")
	b.WriteString("Content-Type: text/html; charset=utf-8\r\n")
	b.WriteString("Content-Transfer-Encoding: 8bit\r\n\r\n")
	b.WriteString(normalizeNewlines(in.BodyHTML))
	b.WriteString("\r\n")

	b.WriteString("--")
	b.WriteString(boundary)
	b.WriteString("--\r\n")

	return b.Bytes()
}

// renderPlainToHTML escapa o input plano e converte newlines em <br>.
// Não suporta markdown — usuário insere texto puro; UI também serve só plain.
func renderPlainToHTML(plain string) string {
	if plain == "" {
		return ""
	}
	escaped := html.EscapeString(plain)
	withBR := strings.ReplaceAll(escaped, "\n", "<br>\n")
	return `<div style="font-family:Arial,Helvetica,sans-serif;font-size:14px;line-height:1.5;color:#222">` +
		withBR +
		`</div>`
}

// normalizeNewlines garante CRLF (RFC 5322).
func normalizeNewlines(s string) string {
	s = strings.ReplaceAll(s, "\r\n", "\n")
	s = strings.ReplaceAll(s, "\r", "\n")
	return strings.ReplaceAll(s, "\n", "\r\n")
}

// generateMessageID monta "<<base64-token>@domain>" RFC 5322-compliant.
func generateMessageID(domain string) string {
	if domain == "" {
		domain = "plenyasaude.com.br"
	}
	tok := randomToken(18)
	return fmt.Sprintf("<%s@%s>", tok, domain)
}

// randomToken gera N bytes random base64 url-safe (sem padding).
func randomToken(n int) string {
	buf := make([]byte, n)
	if _, err := rand.Read(buf); err != nil {
		// Fallback determinístico — só cai aqui em ambiente sem entropia (raro).
		return strings.ReplaceAll(uuid.NewString(), "-", "")
	}
	return base64.RawURLEncoding.EncodeToString(buf)
}

// domainFromAddress extrai o domínio de "user@host" — vazio se não tem '@'.
func domainFromAddress(addr string) string {
	at := strings.LastIndex(addr, "@")
	if at < 0 || at == len(addr)-1 {
		return ""
	}
	return strings.TrimSpace(addr[at+1:])
}

// mimeEncodeWord codifica subject/name não-ASCII em RFC 2047 (Q ou B encoding).
// Pra simplicidade, sempre Q; passthrough quando puro ASCII.
func mimeEncodeWord(s string) string {
	if isASCII(s) {
		return s
	}
	return mime.QEncoding.Encode("utf-8", s)
}

func isASCII(s string) bool {
	for _, r := range s {
		if r > 127 {
			return false
		}
	}
	return true
}
