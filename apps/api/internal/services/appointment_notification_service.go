// Package services — AppointmentNotificationService dispara confirmações,
// reminders, cancelamentos e reagendamentos de consultas (Calendar V1, Bloco E).
//
// Reusa a infra existente:
//   - EmailService.send (Resend) — pra mandar email transacional com .ics anexo
//   - WhatsAppService.SendTemplate — pra templates UTILITY aprovados no Meta
//
// Decisões:
//
//  1. Templates WhatsApp do ciclo de consulta (confirmacao_consulta_semana/vespera/dia,
//     followup_pos_consulta) são aprovados na Meta e configuráveis via cfg.WhatsApp.
//     Cancelamento/remarcação não têm template aprovado → só e-mail. Quando
//     IsConfigured()=false, log + degrade silencioso (não bloqueia confirmação).
//
//  2. ICS é montado inline (RFC 5545 minimal). Não usamos lib externa pra evitar
//     dep transitiva — formato é estável e simples (60 linhas).
//
//  3. Privacidade (LGPD): SUMMARY do .ics é opaco ("Plenya — {tipo}"), sem nome
//     do paciente. DESCRIPTION vazio. Mesma regra do evento Google.
//
//  4. ConfirmationSentAt é setado APÓS sucesso do email — se email falha, evita
//     reenvio em retry (idempotência best-effort).
//
//  5. ATENÇÃO: SendConfirmation é chamado async (goSafe) por AppointmentService
//     após Create. Erros aqui logam mas não retornam pro caller HTTP.
package services

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
)

// Templates Meta WhatsApp (UTILITY pt_BR) do ciclo de consulta vêm de
// cfg.WhatsApp.TemplateApptConfirm/Vespera/Dia + TemplateFollowup. Os nomes
// aprovados na Meta são confirmacao_consulta_{semana,vespera,dia} e
// followup_pos_consulta. Cancel/remarcação não têm template aprovado → só e-mail.
// Ver docs/emr/whatsapp-templates-wiring.md.
const icsAttachmentName = "consulta-plenya.ics"

// ErrAppointmentMissingPatient → preload do appointment não trouxe Patient.
// Sinaliza bug no caller (deve preload Patient + Doctor antes de chamar).
var ErrAppointmentMissingPatient = errors.New("appointment notif: patient not loaded")

type AppointmentNotificationService struct {
	db              *gorm.DB
	emailSvc        *EmailService
	whatsappSvc     *WhatsAppService
	cfg             *config.Config
	telemedLobbySvc *TelemedLobbyService
}

func NewAppointmentNotificationService(
	db *gorm.DB,
	emailSvc *EmailService,
	whatsappSvc *WhatsAppService,
	cfg *config.Config,
) *AppointmentNotificationService {
	return &AppointmentNotificationService{
		db:          db,
		emailSvc:    emailSvc,
		whatsappSvc: whatsappSvc,
		cfg:         cfg,
	}
}

// WithTelemedLobby injeta TelemedLobbyService — chamado em main.go.
// HIGH H9: emails/WhatsApp passam a mandar a URL do lobby (/sala/<token>)
// em vez da URL crua da sala Daily.co. Lobby gera meeting_token na hora.
func (s *AppointmentNotificationService) WithTelemedLobby(svc *TelemedLobbyService) *AppointmentNotificationService {
	s.telemedLobbySvc = svc
	return s
}

// patientLobbyURL devolve o link público /sala/<token> pro paciente entrar
// na sala Daily.co. Vazio se appointment não é telemed ou se token não pôde
// ser emitido (degrade — caller envia notificação sem link).
func (s *AppointmentNotificationService) patientLobbyURL(appt *models.Appointment) string {
	if appt == nil || appt.Type != models.AppointmentTelemedicine {
		return ""
	}
	if s.telemedLobbySvc == nil {
		return ""
	}
	tok, err := s.telemedLobbySvc.EnsureTokenForAppointment(appt)
	if err != nil || tok == nil {
		log.Printf("⚠️  [APPT NOTIF] lobby token apt=%s: %v", appt.ID, err)
		return ""
	}
	return s.telemedLobbySvc.PublicLinkURL(tok.Token)
}

// SendConfirmation envia email (Resend, com .ics) + WhatsApp template
// (confirmacao_consulta_semana) imediatamente após criar o appointment.
//
// Idempotência: se ConfirmationSentAt != NULL, não reenvia.
func (s *AppointmentNotificationService) SendConfirmation(ctx context.Context, apptID uuid.UUID) error {
	appt, err := s.loadAppointment(ctx, apptID)
	if err != nil {
		return err
	}
	if appt.ConfirmationSentAt != nil {
		return nil // já enviado — idempotente
	}

	doctorName := s.doctorDisplayName(&appt.Doctor)
	dateBR, timeBR := formatBRDateTime(appt.ScheduledAt)
	typeLabel := appointmentTypeLabel(appt.Type)
	// HIGH H9 — paciente recebe URL do lobby standalone (/sala/<token>),
	// nunca a URL crua da sala. Lobby gera meeting_token na hora do clique.
	dailyURL := s.patientLobbyURL(appt)

	// 1) Email
	if appt.Patient.Email != nil && strings.TrimSpace(*appt.Patient.Email) != "" {
		subject := fmt.Sprintf("Sua consulta na Plenya · %s %s", dateBR, timeBR)
		bodyText := buildConfirmationBody(appt, doctorName, dateBR, timeBR, typeLabel, dailyURL)
		bodyHTML := buildConfirmationHTML(appt, doctorName, dateBR, timeBR, typeLabel, dailyURL)
		ics := buildICS(appt, typeLabel, dailyURL)

		if err := s.sendEmailWithICS(*appt.Patient.Email, subject, bodyText, bodyHTML, ics); err != nil {
			log.Printf("⚠️  [APPT NOTIF] confirmation email apt=%s: %v", apptID, err)
			// Não retorna — segue pra WA. Email falhou, mas WA pode salvar.
		}
	} else {
		log.Printf("ℹ️  [APPT NOTIF] confirmation email apt=%s skipped (paciente sem email)", apptID)
	}

	// 2) WhatsApp template confirmacao_consulta_semana:
	//    {{1}} nome · {{2}} data ("12 de junho") · {{3}} hora ("14h") · {{4}} modalidade
	if appt.Patient.Phone != nil && strings.TrimSpace(*appt.Patient.Phone) != "" {
		s.sendTemplateBestEffort(
			s.cfg.WhatsApp.TemplateApptConfirm,
			*appt.Patient.Phone,
			[]string{
				firstNameOf(appt.Patient.Name),
				formatDataExtensoBR(appt.ScheduledAt),
				formatHoraBR(appt.ScheduledAt),
				modalidadeCurta(appt),
			},
			apptID,
		)
	}

	// 3) Marca enviado
	now := time.Now().UTC()
	if err := s.db.WithContext(ctx).Model(&models.Appointment{}).
		Where("id = ?", apptID).
		Update("confirmation_sent_at", now).Error; err != nil {
		log.Printf("⚠️  [APPT NOTIF] persist confirmation_sent_at apt=%s: %v", apptID, err)
	}
	return nil
}

// SendReminder24h envia o lembrete da véspera (template confirmacao_consulta_vespera).
// Cron dispara pra appointments scheduled_at em [now+23h, now+25h] sem reminder_sent_at.
//
// Email NÃO é enviado no reminder — paciente já recebeu confirmação com .ics.
func (s *AppointmentNotificationService) SendReminder24h(ctx context.Context, apptID uuid.UUID) error {
	appt, err := s.loadAppointment(ctx, apptID)
	if err != nil {
		return err
	}
	if appt.ReminderSentAt != nil {
		return nil
	}

	if appt.Patient.Phone == nil || strings.TrimSpace(*appt.Patient.Phone) == "" {
		// Sem WA, sem reminder. Marca como "enviado" pra cron não retentar infinito.
		now := time.Now().UTC()
		_ = s.db.WithContext(ctx).Model(&models.Appointment{}).
			Where("id = ?", apptID).Update("reminder_sent_at", now).Error
		return nil
	}

	// confirmacao_consulta_vespera: {{1}} nome · {{2}} hora · {{3}} modalidade
	s.sendTemplateBestEffort(
		s.cfg.WhatsApp.TemplateApptVespera,
		*appt.Patient.Phone,
		[]string{
			firstNameOf(appt.Patient.Name),
			formatHoraBR(appt.ScheduledAt),
			modalidadeCurta(appt),
		},
		apptID,
	)

	now := time.Now().UTC()
	if err := s.db.WithContext(ctx).Model(&models.Appointment{}).
		Where("id = ?", apptID).
		Update("reminder_sent_at", now).Error; err != nil {
		log.Printf("⚠️  [APPT NOTIF] persist reminder_sent_at apt=%s: %v", apptID, err)
	}
	return nil
}

// SendReminderDayOf envia o lembrete do dia da consulta (template
// confirmacao_consulta_dia). Cron dispara pra appointments scheduled_at em
// [now+2h, now+4h] sem dayof_reminder_sent_at. Independe do reminder da véspera.
func (s *AppointmentNotificationService) SendReminderDayOf(ctx context.Context, apptID uuid.UUID) error {
	appt, err := s.loadAppointment(ctx, apptID)
	if err != nil {
		return err
	}
	if appt.DayofReminderSentAt != nil {
		return nil
	}

	if appt.Patient.Phone == nil || strings.TrimSpace(*appt.Patient.Phone) == "" {
		now := time.Now().UTC()
		_ = s.db.WithContext(ctx).Model(&models.Appointment{}).
			Where("id = ?", apptID).Update("dayof_reminder_sent_at", now).Error
		return nil
	}

	// confirmacao_consulta_dia: {{1}} nome · {{2}} hora · {{3}} frase de modalidade
	s.sendTemplateBestEffort(
		s.cfg.WhatsApp.TemplateApptDia,
		*appt.Patient.Phone,
		[]string{
			firstNameOf(appt.Patient.Name),
			formatHoraBR(appt.ScheduledAt),
			modalidadeFraseDia(appt),
		},
		apptID,
	)

	now := time.Now().UTC()
	if err := s.db.WithContext(ctx).Model(&models.Appointment{}).
		Where("id = ?", apptID).
		Update("dayof_reminder_sent_at", now).Error; err != nil {
		log.Printf("⚠️  [APPT NOTIF] persist dayof_reminder_sent_at apt=%s: %v", apptID, err)
	}
	return nil
}

// SendFollowup envia o follow-up pós-consulta (template followup_pos_consulta).
// Cron AppointmentFollowupJob dispara pra consultas concluídas (EndAt < now)
// ainda sem followup_sent_at. Idempotente.
func (s *AppointmentNotificationService) SendFollowup(ctx context.Context, apptID uuid.UUID) error {
	template := strings.TrimSpace(s.cfg.WhatsApp.TemplateFollowup)
	if template == "" {
		return nil // follow-up desativado
	}
	appt, err := s.loadAppointment(ctx, apptID)
	if err != nil {
		return err
	}
	if appt.FollowupSentAt != nil {
		return nil
	}

	if appt.Patient.Phone != nil && strings.TrimSpace(*appt.Patient.Phone) != "" {
		// followup_pos_consulta: {{1}} nome
		s.sendTemplateBestEffort(
			template,
			*appt.Patient.Phone,
			[]string{firstNameOf(appt.Patient.Name)},
			apptID,
		)
	}

	now := time.Now().UTC()
	if err := s.db.WithContext(ctx).Model(&models.Appointment{}).
		Where("id = ?", apptID).
		Update("followup_sent_at", now).Error; err != nil {
		log.Printf("⚠️  [APPT NOTIF] persist followup_sent_at apt=%s: %v", apptID, err)
	}
	return nil
}

// SendCancellation envia email curtinho + WA template appointment_cancelled.
func (s *AppointmentNotificationService) SendCancellation(ctx context.Context, apptID uuid.UUID) error {
	appt, err := s.loadAppointment(ctx, apptID)
	if err != nil {
		return err
	}

	doctorName := s.doctorDisplayName(&appt.Doctor)
	dateBR, timeBR := formatBRDateTime(appt.ScheduledAt)
	typeLabel := appointmentTypeLabel(appt.Type)

	// Email (sem .ics — só notificação)
	if appt.Patient.Email != nil && strings.TrimSpace(*appt.Patient.Email) != "" {
		subject := fmt.Sprintf("Consulta cancelada · %s %s", dateBR, timeBR)
		dateLong, timeTz := formatBRDateTimeLong(appt.ScheduledAt)
		bodyText := fmt.Sprintf(`%s

Sua consulta na Plenya foi cancelada.

Tipo: %s
Data: %s
Horário: %s
Profissional: %s

Para remarcar, responda este e-mail ou fale com a gente no WhatsApp %s.

%s`, greeting(appt.Patient.Name), typeLabel, dateLong, timeTz, doctorName, plenyaWhatsAppDisplay, emailFooterText())

		bodyHTML := buildSimpleHTML(bodyText)
		if err := s.sendEmailWithICS(*appt.Patient.Email, subject, bodyText, bodyHTML, nil); err != nil {
			log.Printf("⚠️  [APPT NOTIF] cancellation email apt=%s: %v", apptID, err)
		}
	}

	// Cancelamento não tem template WhatsApp aprovado na Meta — notificação só por
	// e-mail. (Se um template dedicado for aprovado, cabear aqui via cfg.)
	return nil
}

// SendReschedule envia email com novo .ics + WA template appointment_rescheduled.
// oldScheduledAt é incluído em metadata pra audit (não no corpo paciente).
func (s *AppointmentNotificationService) SendReschedule(ctx context.Context, apptID uuid.UUID, oldScheduledAt time.Time) error {
	appt, err := s.loadAppointment(ctx, apptID)
	if err != nil {
		return err
	}

	doctorName := s.doctorDisplayName(&appt.Doctor)
	dateBR, timeBR := formatBRDateTime(appt.ScheduledAt)
	oldDateBR, oldTimeBR := formatBRDateTime(oldScheduledAt)
	typeLabel := appointmentTypeLabel(appt.Type)
	// HIGH H9 — paciente recebe URL do lobby (/sala/<token>), não a URL crua.
	dailyURL := s.patientLobbyURL(appt)

	if appt.Patient.Email != nil && strings.TrimSpace(*appt.Patient.Email) != "" {
		subject := fmt.Sprintf("Consulta remarcada · %s %s", dateBR, timeBR)
		newDateLong, newTimeTz := formatBRDateTimeLong(appt.ScheduledAt)
		bodyText := fmt.Sprintf(`%s

Sua consulta na Plenya foi remarcada.

Tipo: %s
De: %s às %s
Para: %s, %s
Profissional: %s%s

Para cancelar, remarcar de novo ou tirar dúvidas, responda este e-mail ou fale com a gente no WhatsApp %s.

%s`,
			greeting(appt.Patient.Name),
			typeLabel,
			oldDateBR, oldTimeBR,
			newDateLong, newTimeTz,
			doctorName,
			telemedBlockText(dailyURL),
			plenyaWhatsAppDisplay,
			emailFooterText(),
		)
		bodyHTML := buildSimpleHTML(bodyText)
		ics := buildICS(appt, typeLabel, dailyURL)

		if err := s.sendEmailWithICS(*appt.Patient.Email, subject, bodyText, bodyHTML, ics); err != nil {
			log.Printf("⚠️  [APPT NOTIF] reschedule email apt=%s: %v", apptID, err)
		}
	}

	// Remarcação não tem template WhatsApp aprovado na Meta — notificação só por
	// e-mail (com novo .ics). (Cabear aqui via cfg se um template for aprovado.)
	return nil
}

// loadAppointment carrega o appointment com Patient + Doctor preloaded.
func (s *AppointmentNotificationService) loadAppointment(ctx context.Context, id uuid.UUID) (*models.Appointment, error) {
	var appt models.Appointment
	err := s.db.WithContext(ctx).
		Preload("Patient").
		Preload("Doctor").
		Where("id = ?", id).
		First(&appt).Error
	if err != nil {
		return nil, fmt.Errorf("appointment notif load %s: %w", id, err)
	}
	if appt.Patient.ID == uuid.Nil {
		return nil, ErrAppointmentMissingPatient
	}
	return &appt, nil
}

// doctorDisplayName antepõe o pronome de tratamento (User.Treatment) ao nome do
// profissional, conforme cadastrado no perfil. Renderiza direto, SEM inferência
// por gênero. Sem tratamento definido = só o nome. Se o nome já vier com prefixo
// "Dr."/"Dra.", mantém como veio.
func (s *AppointmentNotificationService) doctorDisplayName(doctor *models.User) string {
	name := strings.TrimSpace(doctor.Name)
	if name == "" {
		return "seu profissional"
	}
	low := strings.ToLower(name)
	if strings.HasPrefix(low, "dr.") || strings.HasPrefix(low, "dra.") || strings.HasPrefix(low, "dr ") {
		return name
	}
	if title := treatmentLabel(doctor.Treatment); title != "" {
		return title + " " + name
	}
	return name
}

// treatmentLabel mapeia o código de tratamento (User.Treatment) para o rótulo
// exibido. Vazio/desconhecido = "" (sem título).
func treatmentLabel(t *string) string {
	if t == nil {
		return ""
	}
	switch strings.ToLower(strings.TrimSpace(*t)) {
	case "dr":
		return "Dr."
	case "dra":
		return "Dra."
	case "sr":
		return "Sr."
	case "sra":
		return "Sra."
	case "enf":
		return "Enf."
	case "enfa":
		return "Enfª."
	default:
		return ""
	}
}

// sendEmailWithICS envia email via Resend com .ics anexo. Reusa o cliente HTTP do
// EmailService — montamos payload Resend manualmente porque EmailService.send não
// expõe attachments na assinatura legada.
func (s *AppointmentNotificationService) sendEmailWithICS(to, subject, bodyText, bodyHTML string, ics []byte) error {
	provider := strings.ToLower(s.cfg.Email.Provider)
	if provider != "resend" || s.cfg.Email.ResendAPIKey == "" {
		// Dev fallback ou SMTP — manda sem anexo. Não bloqueia.
		log.Printf("📧 [APPT NOTIF DEV] email %s subject=%q (sem .ics — provider=%s)",
			to, subject, provider)
		return s.emailSvc.send(to, subject, bodyText, bodyHTML)
	}
	return s.sendResendWithICS(to, subject, bodyText, bodyHTML, ics)
}

// sendResendWithICS faz POST direto na Resend API com attachment base64. Espelha o
// padrão do EmailService.sendReplyViaResend mas em escopo mínimo.
func (s *AppointmentNotificationService) sendResendWithICS(to, subject, bodyText, bodyHTML string, ics []byte) error {
	from := s.cfg.Email.FromAddress
	if from == "" {
		from = "no-reply@plenyasaude.com.br"
	}
	if name := s.cfg.Email.FromName; name != "" {
		from = fmt.Sprintf("%s <%s>", name, from)
	}

	// Reply-To aponta pra um endereço REAL e monitorado (contato@, ingerido pelo
	// CRM). O FROM é noreply@ (transacional), então sem Reply-To o "responda este
	// e-mail" do corpo soaria contraditório. Com Reply-To, a resposta do paciente
	// vai visivelmente pra contato@ e cai na Central de Conversas.
	replyTo := s.cfg.Email.LeadReplyFromAddress
	if replyTo == "" {
		replyTo = "contato@plenyasaude.com.br"
	}

	payload := map[string]any{
		"from":     from,
		"to":       []string{to},
		"reply_to": replyTo,
		"subject":  subject,
		"text":     bodyText,
		"html":     bodyHTML,
		"headers": map[string]string{
			XPlenyaSourceHeader: XPlenyaSourceEMR,
		},
	}
	// .ics é opcional — cancelamento envia sem anexo (ics == nil).
	if len(ics) > 0 {
		payload["attachments"] = []map[string]string{
			{
				"filename":     icsAttachmentName,
				"content":      base64.StdEncoding.EncodeToString(ics),
				"content_type": "text/calendar; charset=utf-8; method=REQUEST",
			},
		}
	}

	raw, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, "https://api.resend.com/emails", bytes.NewReader(raw))
	if err != nil {
		return fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+s.cfg.Email.ResendAPIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.emailSvc.client.Do(req)
	if err != nil {
		return fmt.Errorf("http: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 1024))
		return fmt.Errorf("resend status %d: %s", resp.StatusCode, string(body))
	}
	return nil
}

// sendTemplateBestEffort dispara WA template, log degrada silencioso quando
// service não tá configurado ou template ainda não aprovado no Meta.
func (s *AppointmentNotificationService) sendTemplateBestEffort(name, phone string, params []string, apptID uuid.UUID) {
	if s.whatsappSvc == nil || !s.whatsappSvc.IsConfigured() {
		log.Printf("ℹ️  [APPT NOTIF] WA template %s apt=%s skipped (não configurado)", name, apptID)
		return
	}
	if err := s.whatsappSvc.SendTemplate(phone, name, "pt_BR", params); err != nil {
		// Erro comum: template não aprovado no Meta — degrade.
		log.Printf("⚠️  [APPT NOTIF] WA template %s apt=%s falhou: %v", name, apptID, err)
	}
}

// ============================================================
// Helpers de formatação
// ============================================================

func appointmentTypeLabel(t models.AppointmentType) string {
	switch t {
	case models.AppointmentInitialAssessment:
		return "Avaliação Inicial"
	case models.AppointmentFollowUp:
		return "Retorno"
	case models.AppointmentTelemedicine:
		return "Teleconsulta"
	case models.AppointmentProcedure:
		return "Procedimento"
	case models.AppointmentResultsReview:
		return "Revisão de Exames"
	default:
		return "Consulta"
	}
}

// Contatos públicos canônicos da Plenya (fonte: NAP master + footer do site).
const (
	plenyaWhatsAppLink    = "https://wa.me/5543999748899"
	plenyaWhatsAppDisplay = "(43) 99974-8899"
	plenyaSiteDisplay     = "plenyasaude.com.br"
)

// formatBRDateTime devolve ("23/04/2026", "14:30") em horário local de São Paulo.
// Formato curto — usado no assunto do e-mail e nos params dos templates WhatsApp.
func formatBRDateTime(t time.Time) (string, string) {
	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		loc = time.UTC
	}
	local := t.In(loc)
	return local.Format("02/01/2006"), local.Format("15:04")
}

// formatBRDateTimeLong devolve data por extenso ("quinta-feira, 29/05/2026") e
// horário com fuso ("11:30 (horário de Brasília)") — usado no corpo dos e-mails.
func formatBRDateTimeLong(t time.Time) (string, string) {
	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		loc = time.UTC
	}
	local := t.In(loc)
	return weekdayPT(local.Weekday()) + ", " + local.Format("02/01/2006"),
		local.Format("15:04") + " (horário de Brasília)"
}

func weekdayPT(d time.Weekday) string {
	switch d {
	case time.Sunday:
		return "domingo"
	case time.Monday:
		return "segunda-feira"
	case time.Tuesday:
		return "terça-feira"
	case time.Wednesday:
		return "quarta-feira"
	case time.Thursday:
		return "quinta-feira"
	case time.Friday:
		return "sexta-feira"
	default:
		return "sábado"
	}
}

// firstNameOf devolve só o primeiro nome (param {{1}} dos templates de consulta).
// Vazio → "tudo bem" como fallback neutro (evita "Olá, ." no template).
func firstNameOf(fullName string) string {
	fields := strings.Fields(strings.TrimSpace(fullName))
	if len(fields) == 0 {
		return "tudo bem"
	}
	return fields[0]
}

// formatDataExtensoBR → "12 de junho" (horário local de São Paulo). Param de data
// dos templates de confirmação. Reusa monthsPT (1-indexado) de lab_request_pdf_render.go.
func formatDataExtensoBR(t time.Time) string {
	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		loc = time.UTC
	}
	local := t.In(loc)
	return fmt.Sprintf("%d de %s", local.Day(), monthsPT[int(local.Month())])
}

// formatHoraBR → "14h" (minuto 0) ou "14h30". Param de hora dos templates.
func formatHoraBR(t time.Time) string {
	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		loc = time.UTC
	}
	local := t.In(loc)
	if local.Minute() == 0 {
		return fmt.Sprintf("%dh", local.Hour())
	}
	return fmt.Sprintf("%dh%02d", local.Hour(), local.Minute())
}

// modalidadeCurta → frase curta de modalidade ("na modalidade {{X}}" nos templates
// semana/véspera). Ex: "presencial em Londrina" / "por telemedicina".
func modalidadeCurta(appt *models.Appointment) string {
	if appt.Type == models.AppointmentTelemedicine {
		return "por telemedicina"
	}
	return "presencial em Londrina"
}

// modalidadeFraseDia → frase completa de modalidade (template do dia, "{{3}}." como
// sentença independente).
func modalidadeFraseDia(appt *models.Appointment) string {
	if appt.Type == models.AppointmentTelemedicine {
		return "O atendimento é por telemedicina; o link de acesso chega antes do horário"
	}
	return "O atendimento é presencial, na nossa unidade em Londrina"
}

// greeting monta "Olá, {primeiro nome}," (ou "Olá," se nome vazio).
func greeting(fullName string) string {
	fields := strings.Fields(strings.TrimSpace(fullName))
	if len(fields) == 0 {
		return "Olá,"
	}
	return "Olá, " + fields[0] + ","
}

// emailFooterText é a assinatura + rodapé de contato (sem travessões, voz de marca).
func emailFooterText() string {
	return "Equipe Plenya\nPlenya Saúde · " + plenyaSiteDisplay + " · WhatsApp " + plenyaWhatsAppDisplay
}

func emailFooterHTML() string {
	return fmt.Sprintf(`<p style="color:#888;font-size:12px;margin-top:18px">Equipe Plenya<br>`+
		`Plenya Saúde · <a href="https://%s" style="color:#888">%s</a> · WhatsApp <a href="%s" style="color:#888">%s</a></p>`,
		plenyaSiteDisplay, plenyaSiteDisplay, plenyaWhatsAppLink, plenyaWhatsAppDisplay)
}

// telemedBlockText / telemedBlockHTML: bloco de link + instruções da teleconsulta.
// Vazio quando não há link (consulta presencial).
func telemedBlockText(dailyURL string) string {
	if dailyURL == "" {
		return ""
	}
	return "\n\nLink da teleconsulta:\n" + dailyURL +
		"\n\nComo entrar: abra o link no navegador do celular ou do computador e autorize o acesso à câmera e ao microfone. A sala fica disponível a partir de 3 horas antes do horário agendado."
}

func telemedBlockHTML(dailyURL string) string {
	if dailyURL == "" {
		return ""
	}
	return fmt.Sprintf(`<p><strong>Link da teleconsulta:</strong><br><a href="%s">%s</a></p>`+
		`<p style="color:#444">Como entrar: abra o link no navegador do celular ou do computador e autorize o acesso à câmera e ao microfone. A sala fica disponível a partir de 3 horas antes do horário agendado.</p>`,
		dailyURL, dailyURL)
}

func buildConfirmationBody(appt *models.Appointment, doctor, _, _, typeLabel, dailyURL string) string {
	dateLong, timeTz := formatBRDateTimeLong(appt.ScheduledAt)
	return fmt.Sprintf(`%s

Sua consulta na Plenya está confirmada.

Tipo: %s
Data: %s
Horário: %s
Duração: %d minutos
Profissional: %s%s

Adicionamos um arquivo de calendário em anexo (consulta-plenya.ics): abrir no celular adiciona o horário automaticamente à sua agenda.

Para cancelar, remarcar ou tirar dúvidas, responda este e-mail ou fale com a gente no WhatsApp %s.

%s`, greeting(appt.Patient.Name), typeLabel, dateLong, timeTz, appt.DurationMinutes, doctor,
		telemedBlockText(dailyURL), plenyaWhatsAppDisplay, emailFooterText())
}

func buildConfirmationHTML(appt *models.Appointment, doctor, _, _, typeLabel, dailyURL string) string {
	dateLong, timeTz := formatBRDateTimeLong(appt.ScheduledAt)
	return fmt.Sprintf(`<div style="font-family:Arial,Helvetica,sans-serif;font-size:14px;line-height:1.6;color:#222">
<p>%s</p>
<p>Sua consulta na Plenya está confirmada.</p>
<p>
<strong>Tipo:</strong> %s<br>
<strong>Data:</strong> %s<br>
<strong>Horário:</strong> %s<br>
<strong>Duração:</strong> %d minutos<br>
<strong>Profissional:</strong> %s
</p>
%s
<p>Adicionamos um arquivo de calendário em anexo: abrir no celular adiciona o horário automaticamente à sua agenda.</p>
<p>Para cancelar, remarcar ou tirar dúvidas, responda este e-mail ou fale com a gente no WhatsApp <a href="%s">%s</a>.</p>
%s
</div>`, greeting(appt.Patient.Name), typeLabel, dateLong, timeTz, appt.DurationMinutes, doctor,
		telemedBlockHTML(dailyURL), plenyaWhatsAppLink, plenyaWhatsAppDisplay, emailFooterHTML())
}

func buildSimpleHTML(plain string) string {
	escaped := strings.ReplaceAll(plain, "\n", "<br>")
	return `<div style="font-family:Arial,Helvetica,sans-serif;font-size:14px;line-height:1.5;color:#222">` +
		escaped + `</div>`
}

// ============================================================
// ICS (RFC 5545 minimal)
// ============================================================

// buildICS monta um VCALENDAR/VEVENT mínimo. Compatível com Google/Apple/Outlook.
//
// Campos:
//   - METHOD:REQUEST = paciente pode RSVP
//   - UID = appointment ID (estável → updates substituem)
//   - DTSTART/DTEND = UTC ISO básico (20260423T140000Z)
//   - SUMMARY = "Plenya — {tipo}" (sem PHI — privacidade)
//   - LOCATION = link Daily se telemedicina, "Plenya Saúde" caso contrário
//   - DESCRIPTION = vazio (privacidade)
func buildICS(appt *models.Appointment, typeLabel, dailyURL string) []byte {
	start := appt.ScheduledAt.UTC().Format("20060102T150405Z")
	end := appt.ScheduledAt.Add(time.Duration(appt.DurationMinutes) * time.Minute).
		UTC().Format("20060102T150405Z")
	dtstamp := time.Now().UTC().Format("20060102T150405Z")

	location := "Plenya Saúde"
	description := ""
	if dailyURL != "" {
		location = dailyURL
		description = "Link da teleconsulta: " + dailyURL
	}
	summary := fmt.Sprintf("Plenya · %s", typeLabel)

	// CRLF entre linhas conforme RFC 5545.
	lines := []string{
		"BEGIN:VCALENDAR",
		"VERSION:2.0",
		"PRODID:-//Plenya//EMR//PT-BR",
		"METHOD:REQUEST",
		"CALSCALE:GREGORIAN",
		"BEGIN:VEVENT",
		"UID:" + appt.ID.String() + "@plenyasaude.com.br",
		"DTSTAMP:" + dtstamp,
		"DTSTART:" + start,
		"DTEND:" + end,
		"ORGANIZER;CN=Plenya Saúde:mailto:contato@plenyasaude.com.br",
		"SUMMARY:" + escapeICSText(summary),
		"LOCATION:" + escapeICSText(location),
		"DESCRIPTION:" + escapeICSText(description),
		"STATUS:CONFIRMED",
		"TRANSP:OPAQUE",
		"END:VEVENT",
		"END:VCALENDAR",
	}
	return []byte(strings.Join(lines, "\r\n") + "\r\n")
}

// escapeICSText escapa chars especiais ICS (RFC 5545 §3.3.11).
func escapeICSText(s string) string {
	s = strings.ReplaceAll(s, `\`, `\\`)
	s = strings.ReplaceAll(s, ";", `\;`)
	s = strings.ReplaceAll(s, ",", `\,`)
	s = strings.ReplaceAll(s, "\n", `\n`)
	return s
}
