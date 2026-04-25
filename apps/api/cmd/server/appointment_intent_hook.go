// appointment_intent_hook.go — Calendar V1 (Bloco F) integration glue.
//
// Liga LeadService.ProcessInboundWhatsApp → ConversationService.DetectAppointmentIntent
// → AppointmentService.Confirm OU NotificationService.CreateConversationNotification.
//
// Mantido aqui (cmd/server) ao invés de em services/ pra evitar import cycle:
// services/lead_service.go declara o tipo PatientInboundWAHook (assinatura),
// e essa wiring monta a closure que usa ConversationService + AppointmentService
// sem que LeadService precise conhecer ambos.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
)

const (
	// Janela de busca por appointments próximos do paciente (passados + futuros).
	intentLookbackHours = 7 * 24 * time.Hour
	intentLookaheadDays = 14
	intentMaxAppts      = 5
	intentDetectTimeout = 30 * time.Second
)

// buildAppointmentIntentHook devolve uma PatientInboundWAHook que:
//
//  1. Carrega últimos appointments do paciente (passados 7d + futuros 14d, máx 5).
//  2. Se não houver appointments → no-op (não chama IA pra economizar tokens).
//  3. Roda DetectAppointmentIntent → branches:
//     - confirm (≥0.7): AppointmentService.Confirm(actor=nil) + audit
//     - cancel | reschedule: cria notificação in-app pra secretária
//     - none: nada (segue fluxo CRM normal)
//  4. Sempre cria LeadActivity{note_added, internal} com metadata da decisão IA.
func buildAppointmentIntentHook(
	db *gorm.DB,
	convSvc *services.ConversationService,
	apptSvc *services.AppointmentService,
	notifSvc *services.NotificationService,
	cfg *config.Config,
) services.PatientInboundWAHook {
	return func(patient *models.Patient, message string, waMessageID string) {
		ctx, cancel := context.WithTimeout(context.Background(), intentDetectTimeout)
		defer cancel()

		// 1) Carrega appointments próximos do paciente.
		now := time.Now().UTC()
		from := now.Add(-intentLookbackHours)
		to := now.AddDate(0, 0, intentLookaheadDays)

		var appts []models.Appointment
		err := db.WithContext(ctx).
			Where("patient_id = ?", patient.ID).
			Where("scheduled_at BETWEEN ? AND ?", from, to).
			Where("status NOT IN ?", []models.AppointmentStatus{
				models.AppointmentCancelled,
				models.AppointmentNoShow,
				models.AppointmentCompleted,
			}).
			Order("scheduled_at ASC").
			Limit(intentMaxAppts).
			Find(&appts).Error
		if err != nil {
			log.Printf("⚠️  [INTENT HOOK] load appts patient=%s: %v", patient.ID, err)
			return
		}
		if len(appts) == 0 {
			return
		}

		// 2) Roda IA classifier.
		res, err := convSvc.DetectAppointmentIntent(ctx, message, appts)
		if err != nil {
			log.Printf("⚠️  [INTENT HOOK] detect patient=%s: %v", patient.ID, err)
			return
		}
		if res == nil || res.Intent == services.IntentNone {
			// Não logamos conteúdo da mensagem — só intent agregado.
			log.Printf("ℹ️  [INTENT HOOK] patient=%s intent=none confidence=%.2f", patient.ID, res.Confidence)
			return
		}

		// Resolve appointment alvo: pref appointment_ref, fallback ao mais próximo.
		targetID := pickAppointmentTarget(res.AppointmentRef, appts)
		var target *models.Appointment
		for i := range appts {
			if appts[i].ID == targetID {
				target = &appts[i]
				break
			}
		}
		if target == nil {
			log.Printf("ℹ️  [INTENT HOOK] patient=%s sem appt resolvável (intent=%s)", patient.ID, res.Intent)
			return
		}

		// 3) Branches por intent.
		switch res.Intent {
		case services.IntentConfirm:
			if err := apptSvc.Confirm(ctx, target.ID, nil); err != nil {
				log.Printf("⚠️  [INTENT HOOK] auto-confirm apt=%s: %v", target.ID, err)
				return
			}
			log.Printf("✅ [INTENT HOOK] auto-confirmou apt=%s patient=%s confidence=%.2f",
				target.ID, patient.ID, res.Confidence)

		case services.IntentCancel, services.IntentReschedule:
			notifyStaffAboutIntent(ctx, db, notifSvc, cfg, patient, target, res, message)

		default:
			// any unknown intent → ignore.
		}

		// 4) Audit: LeadActivity{note_added, internal} com metadata IA.
		recordIntentAudit(ctx, db, patient.ID, target.ID, res, message, waMessageID)
	}
}

// pickAppointmentTarget escolhe o appointment alvo: usa appointment_ref da IA
// quando válido (existir na lista), senão pega o mais próximo (primeira posição
// da lista que veio ordenada ASC por scheduled_at).
func pickAppointmentTarget(ref *uuid.UUID, appts []models.Appointment) uuid.UUID {
	if ref != nil {
		for _, a := range appts {
			if a.ID == *ref {
				return *ref
			}
		}
	}
	if len(appts) > 0 {
		return appts[0].ID
	}
	return uuid.Nil
}

// notifyStaffAboutIntent cria notificação in-app pros usuários do CRM (admin/
// secretary/manager) avisando que paciente quer cancelar ou remarcar.
func notifyStaffAboutIntent(
	ctx context.Context,
	db *gorm.DB,
	notifSvc *services.NotificationService,
	cfg *config.Config,
	patient *models.Patient,
	appt *models.Appointment,
	res *services.AppointmentIntentResult,
	message string,
) {
	if notifSvc == nil {
		return
	}
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	if loc == nil {
		loc = time.UTC
	}
	when := appt.ScheduledAt.In(loc).Format("02/01 15:04")
	intentLabel := "cancelar"
	if res.Intent == services.IntentReschedule {
		intentLabel = "remarcar"
	}
	title := fmt.Sprintf("Paciente quer %s consulta", intentLabel)

	snippet := message
	if len(snippet) > 100 {
		snippet = snippet[:97] + "..."
	}
	msg := fmt.Sprintf("%s · consulta %s · \"%s\"", patient.Name, when, snippet)
	actionURL := fmt.Sprintf("/conversas?owner=patient&id=%s", patient.ID.String())

	// Resolve destinatários (admin/secretary/manager). Reusa lookup simples por roles.
	var users []models.User
	if err := db.WithContext(ctx).
		Where(`roles::jsonb ?| array['admin','secretary','manager']`).
		Find(&users).Error; err != nil {
		log.Printf("⚠️  [INTENT HOOK] lookup staff: %v", err)
		return
	}
	for _, u := range users {
		patientID := patient.ID
		if err := notifSvc.CreateConversationNotification(
			u.ID, nil, &patientID,
			models.NotificationLeadAssigned,
			title, msg, actionURL,
		); err != nil {
			log.Printf("⚠️  [INTENT HOOK] notify user=%s: %v", u.ID, err)
		}
	}
}

// recordIntentAudit grava trilha imutável da decisão IA (LeadActivity internal).
// Não armazena conteúdo bruto da mensagem (LGPD) — só intent + confidence + reasoning.
func recordIntentAudit(
	ctx context.Context,
	db *gorm.DB,
	patientID uuid.UUID,
	apptID uuid.UUID,
	res *services.AppointmentIntentResult,
	message string,
	waMessageID string,
) {
	meta := map[string]any{
		"ai_intent":      string(res.Intent),
		"ai_confidence":  res.Confidence,
		"ai_reasoning":   res.Reasoning,
		"appointment_id": apptID.String(),
		"wa_message_id":  waMessageID,
	}
	metaJSON, err := json.Marshal(meta)
	if err != nil {
		return
	}
	content := fmt.Sprintf("IA detectou intent=%s (confidence=%.2f) na mensagem WhatsApp do paciente",
		res.Intent, res.Confidence)
	pid := patientID
	activity := &models.LeadActivity{
		PatientID: &pid,
		Type:      models.LeadActivityNoteAdded,
		Channel:   models.LeadChannelInternal,
		Content:   &content,
		Metadata:  datatypes.JSON(metaJSON),
	}
	if err := db.WithContext(ctx).Create(activity).Error; err != nil {
		log.Printf("⚠️  [INTENT HOOK] audit patient=%s: %v", patientID, err)
	}
	// silence unused parm in case message turns into PHI sink in future.
	_ = message
}
