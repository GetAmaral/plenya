// Package services — AppointmentNotificationService dispara confirmações,
// reminders, cancelamentos e reagendamentos de consultas (Calendar V1, Bloco E).
//
// Reusa a infra existente:
//   - EmailService.send (Resend) — pra mandar email transacional com .ics anexo
//   - WhatsAppService.SendTemplate — pra templates UTILITY aprovados no Meta
//
// Decisões:
//
//  1. Templates WhatsApp (appointment_confirmation, appointment_reminder_24h,
//     appointment_cancelled, appointment_rescheduled) AINDA NÃO foram submetidos
//     ao Meta — Bloco H. Quando WhatsAppService.IsConfigured()=false ou template
//     ainda não aprovado, log + degrade silencioso (não bloqueia confirmação).
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

// Templates Meta WhatsApp (UTILITY pt_BR). Nomes podem virar configuráveis no
// futuro (cfg.WhatsApp.TemplateAppointmentX) — por ora, hardcoded constants.
const (
	waTemplateAppointmentConfirmation = "appointment_confirmation"
	waTemplateAppointmentReminder24h  = "appointment_reminder_24h"
	waTemplateAppointmentCancelled    = "appointment_cancelled"
	waTemplateAppointmentRescheduled  = "appointment_rescheduled"

	icsAttachmentName = "consulta-plenya.ics"
)

// ErrAppointmentMissingPatient → preload do appointment não trouxe Patient.
// Sinaliza bug no caller (deve preload Patient + Doctor antes de chamar).
var ErrAppointmentMissingPatient = errors.New("appointment notif: patient not loaded")

type AppointmentNotificationService struct {
	db          *gorm.DB
	emailSvc    *EmailService
	whatsappSvc *WhatsAppService
	cfg         *config.Config
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

// SendConfirmation envia email (Resend, com .ics) + WhatsApp template
// (appointment_confirmation) imediatamente após criar o appointment.
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
	dailyURL := ""
	if appt.DailyRoomURL != nil {
		dailyURL = *appt.DailyRoomURL
	}

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

	// 2) WhatsApp template
	if appt.Patient.Phone != nil && strings.TrimSpace(*appt.Patient.Phone) != "" {
		s.sendTemplateBestEffort(
			waTemplateAppointmentConfirmation,
			*appt.Patient.Phone,
			[]string{doctorName, dateBR, timeBR, dailyURL},
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

// SendReminder24h envia template WA appointment_reminder_24h. Cron dispara isso
// pra appointments scheduled_at em [now+23h, now+25h] sem reminder_sent_at.
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

	doctorName := s.doctorDisplayName(&appt.Doctor)
	dateBR, timeBR := formatBRDateTime(appt.ScheduledAt)
	dailyURL := ""
	if appt.DailyRoomURL != nil {
		dailyURL = *appt.DailyRoomURL
	}

	s.sendTemplateBestEffort(
		waTemplateAppointmentReminder24h,
		*appt.Patient.Phone,
		[]string{doctorName, dateBR, timeBR, dailyURL},
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
		bodyText := fmt.Sprintf(`Sua consulta foi cancelada:

Tipo: %s
Data: %s às %s
Profissional: %s

Se quiser remarcar, é só responder por WhatsApp ou email.

— Equipe Plenya`, typeLabel, dateBR, timeBR, doctorName)

		bodyHTML := buildSimpleHTML(bodyText)
		if err := s.emailSvc.send(*appt.Patient.Email, subject, bodyText, bodyHTML); err != nil {
			log.Printf("⚠️  [APPT NOTIF] cancellation email apt=%s: %v", apptID, err)
		}
	}

	// WA template
	if appt.Patient.Phone != nil && strings.TrimSpace(*appt.Patient.Phone) != "" {
		s.sendTemplateBestEffort(
			waTemplateAppointmentCancelled,
			*appt.Patient.Phone,
			[]string{doctorName, dateBR, timeBR},
			apptID,
		)
	}
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
	dailyURL := ""
	if appt.DailyRoomURL != nil {
		dailyURL = *appt.DailyRoomURL
	}

	if appt.Patient.Email != nil && strings.TrimSpace(*appt.Patient.Email) != "" {
		subject := fmt.Sprintf("Consulta remarcada · %s %s", dateBR, timeBR)
		bodyText := fmt.Sprintf(`Sua consulta foi remarcada:

Tipo: %s
De: %s às %s
Para: %s às %s
Profissional: %s
%s

— Equipe Plenya`,
			typeLabel,
			oldDateBR, oldTimeBR,
			dateBR, timeBR,
			doctorName,
			dailyLine(dailyURL),
		)
		bodyHTML := buildSimpleHTML(bodyText)
		ics := buildICS(appt, typeLabel, dailyURL)

		if err := s.sendEmailWithICS(*appt.Patient.Email, subject, bodyText, bodyHTML, ics); err != nil {
			log.Printf("⚠️  [APPT NOTIF] reschedule email apt=%s: %v", apptID, err)
		}
	}

	if appt.Patient.Phone != nil && strings.TrimSpace(*appt.Patient.Phone) != "" {
		s.sendTemplateBestEffort(
			waTemplateAppointmentRescheduled,
			*appt.Patient.Phone,
			[]string{doctorName, dateBR, timeBR, dailyURL},
			apptID,
		)
	}
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

// doctorDisplayName retorna "Dr. {Name}" / "Dra. {Name}" quando possível,
// ou só Name. Heurística simples — User.Name pode vir com prefixo já.
func (s *AppointmentNotificationService) doctorDisplayName(doctor *models.User) string {
	name := strings.TrimSpace(doctor.Name)
	if name == "" {
		return "seu profissional"
	}
	low := strings.ToLower(name)
	if strings.HasPrefix(low, "dr.") || strings.HasPrefix(low, "dra.") || strings.HasPrefix(low, "dr ") {
		return name
	}
	return "Dr(a). " + name
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

	payload := map[string]any{
		"from":    from,
		"to":      []string{to},
		"subject": subject,
		"text":    bodyText,
		"html":    bodyHTML,
		"headers": map[string]string{
			XPlenyaSourceHeader: XPlenyaSourceEMR,
		},
		"attachments": []map[string]string{
			{
				"filename":     icsAttachmentName,
				"content":      base64.StdEncoding.EncodeToString(ics),
				"content_type": "text/calendar; charset=utf-8; method=REQUEST",
			},
		},
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

// formatBRDateTime devolve ("23/04/2026", "14:30") em horário local de São Paulo.
func formatBRDateTime(t time.Time) (string, string) {
	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		loc = time.UTC
	}
	local := t.In(loc)
	return local.Format("02/01/2006"), local.Format("15:04")
}

func dailyLine(url string) string {
	if url == "" {
		return ""
	}
	return "Link teleconsulta: " + url
}

func buildConfirmationBody(appt *models.Appointment, doctor, dateBR, timeBR, typeLabel, dailyURL string) string {
	durLabel := fmt.Sprintf("%d minutos", appt.DurationMinutes)
	dailyPart := ""
	if dailyURL != "" {
		dailyPart = "\n\nLink teleconsulta:\n" + dailyURL
	}
	return fmt.Sprintf(`Sua consulta na Plenya está confirmada:

Tipo: %s
Data: %s
Hora: %s
Duração: %s
Profissional: %s%s

Adicionamos um arquivo de calendário em anexo (consulta-plenya.ics) — abrir no celular adiciona automaticamente na sua agenda.

Se precisar cancelar ou remarcar, é só responder esta mensagem ou nos chamar no WhatsApp.

— Equipe Plenya`, typeLabel, dateBR, timeBR, durLabel, doctor, dailyPart)
}

func buildConfirmationHTML(appt *models.Appointment, doctor, dateBR, timeBR, typeLabel, dailyURL string) string {
	dailyPart := ""
	if dailyURL != "" {
		dailyPart = fmt.Sprintf(`<p><strong>Link teleconsulta:</strong><br><a href="%s">%s</a></p>`, dailyURL, dailyURL)
	}
	return fmt.Sprintf(`<div style="font-family:Arial,Helvetica,sans-serif;font-size:14px;line-height:1.6;color:#222">
<p>Sua consulta na Plenya está confirmada.</p>
<p>
<strong>Tipo:</strong> %s<br>
<strong>Data:</strong> %s<br>
<strong>Hora:</strong> %s<br>
<strong>Duração:</strong> %d minutos<br>
<strong>Profissional:</strong> %s
</p>
%s
<p>Adicionamos um arquivo de calendário em anexo — abrir no celular adiciona automaticamente na sua agenda.</p>
<p>Se precisar cancelar ou remarcar, é só responder esta mensagem ou nos chamar no WhatsApp.</p>
<p>— Equipe Plenya</p>
</div>`, typeLabel, dateBR, timeBR, appt.DurationMinutes, doctor, dailyPart)
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
	end := appt.ScheduledAt.Add(time.Duration(appt.DurationMinutes)*time.Minute).
		UTC().Format("20060102T150405Z")
	dtstamp := time.Now().UTC().Format("20060102T150405Z")

	location := "Plenya Saúde"
	if dailyURL != "" {
		location = dailyURL
	}
	summary := fmt.Sprintf("Plenya — %s", typeLabel)

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
		"SUMMARY:" + escapeICSText(summary),
		"LOCATION:" + escapeICSText(location),
		"DESCRIPTION:",
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
