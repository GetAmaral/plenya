// Package scheduler — ContinuumLateJob roda 1×/dia checando items pendentes
// cuja LateAfterDate já passou. Marca como missed e cria notificação in-app
// pro coordenador da inscrição (ou admin se não houver coordenador).
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
	continuumLateInterval = 24 * time.Hour
	continuumLateTimeout  = 5 * time.Minute
)

type ContinuumLateJob struct {
	db       *gorm.DB
	notifSvc *services.NotificationService
}

func NewContinuumLateJob(db *gorm.DB, notifSvc *services.NotificationService) *ContinuumLateJob {
	return &ContinuumLateJob{db: db, notifSvc: notifSvc}
}

// Run executa uma passada: pega items pending com late_after_date < now,
// marca como missed, e cria notification pro coordenador.
func (j *ContinuumLateJob) Run() error {
	ctx, cancel := context.WithTimeout(context.Background(), continuumLateTimeout)
	defer cancel()

	now := time.Now().UTC()

	// Pega items que viraram atrasados nesta passada.
	type itemRow struct {
		ID                  uuid.UUID
		ContinuumID         uuid.UUID
		Title               string
		ExpectedDate        time.Time
		PatientID           uuid.UUID
		PatientName         string
		CoordinatorDoctorID *uuid.UUID
	}
	var rows []itemRow
	err := j.db.WithContext(ctx).
		Table("patient_continuum_items pci").
		Select(`pci.id, pci.continuum_id, pci.title, pci.expected_date,
			pc.patient_id, p.name as patient_name, pc.coordinator_doctor_id`).
		Joins("JOIN patient_continuums pc ON pc.id = pci.continuum_id").
		Joins("JOIN patients p ON p.id = pc.patient_id").
		Where("pci.status = ?", models.ContinuumItemPending).
		Where("pci.late_after_date < ?", now).
		Where("pc.status = ?", models.ContinuumActive).
		Where("pc.deleted_at IS NULL").
		Scan(&rows).Error
	if err != nil {
		log.Printf("⚠️  [CONTINUUM LATE JOB] query falhou: %v", err)
		return err
	}
	if len(rows) == 0 {
		return nil
	}
	log.Printf("⏰ [CONTINUUM LATE JOB] %d items virando missed", len(rows))

	updated := 0
	notified := 0
	for _, r := range rows {
		// Atualiza status numa única query atômica (evita race se 2 jobs concorrerem).
		res := j.db.WithContext(ctx).
			Model(&models.PatientContinuumItem{}).
			Where("id = ? AND status = ?", r.ID, models.ContinuumItemPending).
			Update("status", models.ContinuumItemMissed)
		if res.Error != nil {
			log.Printf("⚠️  [CONTINUUM LATE JOB] update %s: %v", r.ID, res.Error)
			continue
		}
		if res.RowsAffected == 0 {
			continue // alguém ancorou no meio do caminho
		}
		updated++

		// Notifica coordenador (skip se não houver — admin pode rodar dashboard depois).
		if r.CoordinatorDoctorID == nil {
			continue
		}
		title := "[Continuum] Marco atrasado"
		msg := fmt.Sprintf("%s — %s: esperado em %s.",
			r.PatientName, r.Title, r.ExpectedDate.Format("02/01/2006"))
		actionURL := fmt.Sprintf("/patients/%s/continuum", r.PatientID)
		actionText := "Abrir timeline"
		patientID := r.PatientID
		if err := j.notifSvc.CreateNotification(
			*r.CoordinatorDoctorID,
			models.NotificationGeneral,
			title,
			msg,
			&patientID,
			nil,
			&actionURL,
			&actionText,
		); err != nil {
			log.Printf("⚠️  [CONTINUUM LATE JOB] notif: %v", err)
			continue
		}
		notified++
	}
	log.Printf("✅ [CONTINUUM LATE JOB] missed=%d notif=%d", updated, notified)
	return nil
}

// Start arranca o ticker em goroutine. Roda imediatamente após start (catch-up
// de período sem servidor) e depois a cada continuumLateInterval.
func (j *ContinuumLateJob) Start() {
	log.Printf("⏰ [CONTINUUM LATE JOB] iniciado (interval=%s)", continuumLateInterval)
	go func() {
		if err := j.Run(); err != nil {
			log.Printf("⚠️  [CONTINUUM LATE JOB] erro inicial: %v", err)
		}
		ticker := time.NewTicker(continuumLateInterval)
		defer ticker.Stop()
		for range ticker.C {
			if err := j.Run(); err != nil {
				log.Printf("⚠️  [CONTINUUM LATE JOB] erro: %v", err)
			}
		}
	}()
}
