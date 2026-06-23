// Package scheduler — AppointmentFollowupJob.
//
// Roda a cada 1h checando consultas CONCLUÍDAS (status=completed) cuja janela já
// terminou (end_at < now) nas últimas 48h e que ainda não receberam follow-up
// (followup_sent_at IS NULL). Pra cada uma, dispara
// AppointmentNotificationService.SendFollowup (template followup_pos_consulta).
//
// O recorte end_at > now-48h evita blast histórico no primeiro deploy: só consultas
// recém-concluídas entram. Idempotência via followup_sent_at.
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
	followupJobInterval = 1 * time.Hour
	followupJobTimeout  = 5 * time.Minute
	followupLookback    = 48 * time.Hour
)

type AppointmentFollowupJob struct {
	db       *gorm.DB
	notifSvc *services.AppointmentNotificationService
	enabled  bool // cfg.WhatsApp.TemplateFollowup != ""
}

func NewAppointmentFollowupJob(db *gorm.DB, notifSvc *services.AppointmentNotificationService, templateName string) *AppointmentFollowupJob {
	return &AppointmentFollowupJob{db: db, notifSvc: notifSvc, enabled: templateName != ""}
}

func (j *AppointmentFollowupJob) Run() error {
	if !j.enabled {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), followupJobTimeout)
	defer cancel()

	now := time.Now().UTC()
	var ids []uuid.UUID
	err := j.db.WithContext(ctx).
		Model(&models.Appointment{}).
		Where("status = ?", models.AppointmentCompleted).
		Where("end_at < ?", now).
		Where("end_at > ?", now.Add(-followupLookback)).
		Where("followup_sent_at IS NULL").
		Where("deleted_at IS NULL").
		Pluck("id", &ids).Error
	if err != nil {
		log.Printf("⚠️  [FOLLOWUP JOB] query falhou: %v", err)
		return err
	}
	if len(ids) == 0 {
		return nil
	}

	log.Printf("🩺 [FOLLOWUP JOB] %d consultas concluídas p/ follow-up", len(ids))
	ok, ko := 0, 0
	for _, id := range ids {
		if err := j.notifSvc.SendFollowup(ctx, id); err != nil {
			log.Printf("⚠️  [FOLLOWUP JOB] apt=%s: %v", id, err)
			ko++
			continue
		}
		ok++
	}
	log.Printf("✅ [FOLLOWUP JOB] enviados=%d erros=%d", ok, ko)
	return nil
}

func (j *AppointmentFollowupJob) Start() {
	if !j.enabled {
		log.Printf("🩺 [FOLLOWUP JOB] desativado (sem template)")
		return
	}
	log.Printf("🩺 [FOLLOWUP JOB] iniciado (interval=%s)", followupJobInterval)
	go func() {
		if err := j.Run(); err != nil {
			log.Printf("⚠️  [FOLLOWUP JOB] erro inicial: %v", err)
		}
		ticker := time.NewTicker(followupJobInterval)
		defer ticker.Stop()
		for range ticker.C {
			if err := j.Run(); err != nil {
				log.Printf("⚠️  [FOLLOWUP JOB] erro: %v", err)
			}
		}
	}()
}
