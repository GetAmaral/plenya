// Package scheduler — AppointmentPushReminderJob (App Paciente Sprint 0).
//
// Roda a cada 15min checando appointments cuja scheduled_at está em [now+45min, now+75min]
// e que ainda não receberam push reminder (push_reminder_1h_sent_at IS NULL).
//
// Idempotência via push_reminder_1h_sent_at evita duplo envio quando o cron
// roda 4× na mesma janela. Honra NotificationPreferences.AppointmentReminder
// — paciente que desligou nas prefs não recebe.
package scheduler

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
)

const (
	pushReminderJobInterval = 15 * time.Minute
	pushReminderJobTimeout  = 5 * time.Minute
)

type AppointmentPushReminderJob struct {
	db        *gorm.DB
	push      *services.PushService
	notifPref *services.NotificationPreferencesService
}

func NewAppointmentPushReminderJob(
	db *gorm.DB,
	push *services.PushService,
	notifPref *services.NotificationPreferencesService,
) *AppointmentPushReminderJob {
	return &AppointmentPushReminderJob{db: db, push: push, notifPref: notifPref}
}

func (j *AppointmentPushReminderJob) Run() error {
	ctx, cancel := context.WithTimeout(context.Background(), pushReminderJobTimeout)
	defer cancel()

	now := time.Now().UTC()
	from := now.Add(45 * time.Minute)
	to := now.Add(75 * time.Minute)

	type row struct {
		ID          string
		ScheduledAt time.Time
		PatientUser string
		DoctorName  string
	}
	var rows []row
	err := j.db.WithContext(ctx).
		Table("appointments AS a").
		Select("a.id, a.scheduled_at, p.user_id AS patient_user, COALESCE(u.name,'') AS doctor_name").
		Joins("JOIN patients p ON p.id = a.patient_id").
		Joins("LEFT JOIN users u ON u.id = a.doctor_id").
		Where("a.scheduled_at BETWEEN ? AND ?", from, to).
		Where("a.status IN ?", []models.AppointmentStatus{
			models.AppointmentScheduled,
			models.AppointmentConfirmed,
		}).
		Where("a.push_reminder_1h_sent_at IS NULL").
		Where("a.deleted_at IS NULL").
		Scan(&rows).Error
	if err != nil {
		log.Printf("⚠️  [PUSH REMINDER 1h] query falhou: %v", err)
		return err
	}
	if len(rows) == 0 {
		return nil
	}

	log.Printf("📲 [PUSH REMINDER 1h] %d appointments na janela", len(rows))
	sent, skipped, errs := 0, 0, 0
	for _, r := range rows {
		userID, err := uuid.Parse(r.PatientUser)
		if err != nil {
			errs++
			continue
		}
		prefs, err := j.notifPref.Get(userID)
		if err != nil || !prefs.AppointmentReminder {
			skipped++
			j.markSent(ctx, r.ID, now)
			continue
		}
		body := fmt.Sprintf("Sua consulta com %s é em ~1h.", r.DoctorName)
		if r.DoctorName == "" {
			body = "Sua consulta começa em ~1h."
		}
		if err := j.push.Send(userID, services.PushPayload{
			Title: "Lembrete de consulta",
			Body:  body,
			URL:   "/appointments/" + r.ID,
			Data:  map[string]any{"appointmentId": r.ID, "kind": "appointment_reminder_1h"},
		}); err != nil {
			log.Printf("⚠️  [PUSH REMINDER 1h] apt=%s: %v", r.ID, err)
			errs++
			continue
		}
		j.markSent(ctx, r.ID, now)
		sent++
	}
	log.Printf("✅ [PUSH REMINDER 1h] enviados=%d skipped=%d erros=%d", sent, skipped, errs)
	return nil
}

func (j *AppointmentPushReminderJob) markSent(ctx context.Context, id string, t time.Time) {
	if err := j.db.WithContext(ctx).
		Model(&models.Appointment{}).
		Where("id = ?", id).
		Update("push_reminder_1h_sent_at", t).Error; err != nil {
		log.Printf("⚠️  [PUSH REMINDER 1h] markSent apt=%s: %v", id, err)
	}
}

func (j *AppointmentPushReminderJob) Start() {
	log.Printf("📲 [PUSH REMINDER 1h] iniciado (interval=%s)", pushReminderJobInterval)
	go func() {
		if err := j.Run(); err != nil {
			log.Printf("⚠️  [PUSH REMINDER 1h] erro inicial: %v", err)
		}
		ticker := time.NewTicker(pushReminderJobInterval)
		defer ticker.Stop()
		for range ticker.C {
			if err := j.Run(); err != nil {
				log.Printf("⚠️  [PUSH REMINDER 1h] erro: %v", err)
			}
		}
	}()
}
