// Package scheduler — AppointmentReminderJob (Calendar V1, Bloco E).
//
// Roda a cada 1h checando appointments cuja scheduled_at está em [now+23h, now+25h]
// e que ainda não receberam reminder (reminder_sent_at IS NULL).
//
// Pra cada um, dispara AppointmentNotificationService.SendReminder24h, que envia
// o template WhatsApp `appointment_reminder_24h` e marca reminder_sent_at=now().
//
// Filtros:
//   - status IN (scheduled, confirmed) — não lembra cancelled/completed/no_show
//   - deleted_at IS NULL                — soft-delete
//
// Janela de 2h (23-25h) garante margem se o cron atrasar 1h. Idempotência via
// reminder_sent_at evita duplo envio.
package scheduler

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
)

const (
	reminderJobInterval = 1 * time.Hour
	reminderJobTimeout  = 5 * time.Minute
)

type AppointmentReminderJob struct {
	db       *gorm.DB
	notifSvc *services.AppointmentNotificationService
}

func NewAppointmentReminderJob(db *gorm.DB, notifSvc *services.AppointmentNotificationService) *AppointmentReminderJob {
	return &AppointmentReminderJob{db: db, notifSvc: notifSvc}
}

// Run executa uma passada do job. Loga sumário (OK | erros).
func (j *AppointmentReminderJob) Run() error {
	ctx, cancel := context.WithTimeout(context.Background(), reminderJobTimeout)
	defer cancel()

	now := time.Now().UTC()
	from := now.Add(23 * time.Hour)
	to := now.Add(25 * time.Hour)

	var ids []uuid.UUID
	err := j.db.WithContext(ctx).
		Model(&models.Appointment{}).
		Where("scheduled_at BETWEEN ? AND ?", from, to).
		Where("status IN ?", []models.AppointmentStatus{
			models.AppointmentScheduled,
			models.AppointmentConfirmed,
		}).
		Where("reminder_sent_at IS NULL").
		Where("deleted_at IS NULL").
		Pluck("id", &ids).Error
	if err != nil {
		log.Printf("⚠️  [REMINDER JOB] query falhou: %v", err)
		return err
	}

	if len(ids) > 0 {
		log.Printf("⏰ [REMINDER JOB] %d appointments na janela véspera (24h)", len(ids))
		ok, ko := 0, 0
		for _, id := range ids {
			if err := j.notifSvc.SendReminder24h(ctx, id); err != nil {
				log.Printf("⚠️  [REMINDER JOB] véspera apt=%s: %v", id, err)
				ko++
				continue
			}
			ok++
		}
		log.Printf("✅ [REMINDER JOB] véspera enviados=%d erros=%d", ok, ko)
	}

	// Segunda janela: lembrete no DIA da consulta (~2-4h antes), template
	// confirmacao_consulta_dia. Idempotência por dayof_reminder_sent_at.
	dayFrom := now.Add(2 * time.Hour)
	dayTo := now.Add(4 * time.Hour)
	var dayIDs []uuid.UUID
	if err := j.db.WithContext(ctx).
		Model(&models.Appointment{}).
		Where("scheduled_at BETWEEN ? AND ?", dayFrom, dayTo).
		Where("status IN ?", []models.AppointmentStatus{
			models.AppointmentScheduled,
			models.AppointmentConfirmed,
		}).
		Where("dayof_reminder_sent_at IS NULL").
		Where("deleted_at IS NULL").
		Pluck("id", &dayIDs).Error; err != nil {
		log.Printf("⚠️  [REMINDER JOB] query dia falhou: %v", err)
		return err
	}

	if len(dayIDs) > 0 {
		log.Printf("⏰ [REMINDER JOB] %d appointments na janela do dia (2-4h)", len(dayIDs))
		ok, ko := 0, 0
		for _, id := range dayIDs {
			if err := j.notifSvc.SendReminderDayOf(ctx, id); err != nil {
				log.Printf("⚠️  [REMINDER JOB] dia apt=%s: %v", id, err)
				ko++
				continue
			}
			ok++
		}
		log.Printf("✅ [REMINDER JOB] dia enviados=%d erros=%d", ok, ko)
	}

	return nil
}

// Start arranca o ticker em goroutine. Primeira execução é imediata, depois
// a cada reminderJobInterval.
func (j *AppointmentReminderJob) Start() {
	log.Printf("⏰ [REMINDER JOB] iniciado (interval=%s)", reminderJobInterval)
	go func() {
		// Run imediatamente (catch-up após restart).
		if err := j.Run(); err != nil {
			log.Printf("⚠️  [REMINDER JOB] erro inicial: %v", err)
		}
		ticker := time.NewTicker(reminderJobInterval)
		defer ticker.Stop()
		for range ticker.C {
			if err := j.Run(); err != nil {
				log.Printf("⚠️  [REMINDER JOB] erro: %v", err)
			}
		}
	}()
}
