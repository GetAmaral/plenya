// Package scheduler — ConsultationPrepReminderJob.
//
// Roda a cada 1h e reenvia o convite de preparação pré-consulta (magic link) para consultas
// próximas que têm formulário atrelado e cuja preparação ainda NÃO foi enviada pelo paciente.
// Duas janelas: T-48h e T-24h. Idempotência por janela via prep_reminder_{48h,24h}_sent_at
// (margem de 2h absorve atraso do cron). Não lembra quem já submeteu a preparação.
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
	prepReminderJobInterval = 1 * time.Hour
	prepReminderJobTimeout  = 5 * time.Minute
)

type ConsultationPrepReminderJob struct {
	db       *gorm.DB
	notifier *services.ConsultationPrepNotifier
}

func NewConsultationPrepReminderJob(db *gorm.DB, notifier *services.ConsultationPrepNotifier) *ConsultationPrepReminderJob {
	return &ConsultationPrepReminderJob{db: db, notifier: notifier}
}

func (j *ConsultationPrepReminderJob) Run() error {
	ctx, cancel := context.WithTimeout(context.Background(), prepReminderJobTimeout)
	defer cancel()
	now := time.Now().UTC()
	j.runWindow(ctx, now.Add(47*time.Hour), now.Add(49*time.Hour), "prep_reminder_48h_sent_at")
	j.runWindow(ctx, now.Add(23*time.Hour), now.Add(25*time.Hour), "prep_reminder_24h_sent_at")
	return nil
}

func (j *ConsultationPrepReminderJob) runWindow(ctx context.Context, from, to time.Time, col string) {
	var ids []uuid.UUID
	err := j.db.WithContext(ctx).
		Model(&models.Appointment{}).
		Where("scheduled_at BETWEEN ? AND ?", from, to).
		Where("prep_form_version_id IS NOT NULL").
		Where("status IN ?", []models.AppointmentStatus{
			models.AppointmentScheduled,
			models.AppointmentConfirmed,
		}).
		Where(col+" IS NULL").
		Where("deleted_at IS NULL").
		Where("NOT EXISTS (SELECT 1 FROM consultation_preps cp WHERE cp.appointment_id = appointments.id AND cp.status = 'submitted' AND cp.deleted_at IS NULL)").
		Pluck("id", &ids).Error
	if err != nil {
		log.Printf("⚠️  [PREP REMINDER] query (%s) falhou: %v", col, err)
		return
	}
	if len(ids) == 0 {
		return
	}
	log.Printf("⏰ [PREP REMINDER] %d consultas na janela %s", len(ids), col)
	for _, id := range ids {
		if err := j.notifier.SendPrepInvite(ctx, id); err != nil {
			log.Printf("⚠️  [PREP REMINDER] apt=%s: %v", id, err)
			continue
		}
		if err := j.db.WithContext(ctx).Model(&models.Appointment{}).
			Where("id = ?", id).Update(col, time.Now().UTC()).Error; err != nil {
			log.Printf("⚠️  [PREP REMINDER] carimbar %s apt=%s: %v", col, id, err)
		}
	}
}

// Start arranca o ticker em goroutine. Primeira execução imediata (catch-up), depois a cada 1h.
func (j *ConsultationPrepReminderJob) Start() {
	log.Printf("⏰ [PREP REMINDER] iniciado (interval=%s)", prepReminderJobInterval)
	go func() {
		if err := j.Run(); err != nil {
			log.Printf("⚠️  [PREP REMINDER] erro inicial: %v", err)
		}
		ticker := time.NewTicker(prepReminderJobInterval)
		defer ticker.Stop()
		for range ticker.C {
			if err := j.Run(); err != nil {
				log.Printf("⚠️  [PREP REMINDER] erro: %v", err)
			}
		}
	}()
}
