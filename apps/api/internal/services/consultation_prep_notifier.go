package services

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// ConsultationPrepNotifier dispara o convite de preparação pré-consulta: um magic link do portal
// com deep-link em /preparacao, entregue por email (Resend). O envio proativo por WhatsApp exige
// um template Meta APROVADO (out-of-band, como appointment_reminder_24h) — quando o template
// existir, plugar no ponto marcado abaixo. No-op silencioso quando a consulta não tem form atrelado.
type ConsultationPrepNotifier struct {
	db       *gorm.DB
	portal   *PatientPortalService
	email    *EmailService
	whatsapp *WhatsAppService
}

func NewConsultationPrepNotifier(db *gorm.DB, portal *PatientPortalService, email *EmailService, whatsapp *WhatsAppService) *ConsultationPrepNotifier {
	return &ConsultationPrepNotifier{db: db, portal: portal, email: email, whatsapp: whatsapp}
}

// SendPrepInvite envia o convite de preparação para a consulta. Best-effort: loga e segue em caso
// de falha de canal. A idempotência (não reenviar) fica a cargo do caller (job carimba timestamps).
func (n *ConsultationPrepNotifier) SendPrepInvite(ctx context.Context, apptID uuid.UUID) error {
	var appt models.Appointment
	if err := n.db.WithContext(ctx).Preload("Patient").First(&appt, "id = ?", apptID).Error; err != nil {
		return err
	}
	if appt.PrepFormVersionID == nil || appt.Patient.UserID == uuid.Nil {
		return nil
	}

	next := fmt.Sprintf("/preparacao?appointmentId=%s", apptID.String())
	url, err := n.portal.MintMagicLinkURL(appt.Patient.UserID, next)
	if err != nil {
		return err
	}

	// Email (Resend) — funciona imediatamente.
	if n.email != nil && appt.Patient.Email != nil &&
		strings.TrimSpace(*appt.Patient.Email) != "" && !models.IsPlaceholderEmail(*appt.Patient.Email) {
		if err := n.email.SendConsultationPrepInvite(*appt.Patient.Email, appt.Patient.Name, url); err != nil {
			log.Printf("⚠️  [PREP INVITE] email apt=%s: %v", apptID, err)
		}
	}

	// WhatsApp proativo: reusa o template aprovado "magic_link" (recebe o link como parâmetro).
	// Quando houver um template dedicado "consultation_prep_invite" aprovado na WABA, basta
	// trocar para s.whatsapp.SendTemplate(phone, "consultation_prep_invite", "pt_BR", []string{url}).
	// Em dev (sem credenciais) o WhatsAppService loga em vez de enviar.
	if n.whatsapp != nil && appt.Patient.Phone != nil && strings.TrimSpace(*appt.Patient.Phone) != "" {
		if err := n.whatsapp.SendMagicLink(strings.TrimSpace(*appt.Patient.Phone), url); err != nil {
			log.Printf("⚠️  [PREP INVITE] whatsapp apt=%s: %v", apptID, err)
		}
	}

	return nil
}
