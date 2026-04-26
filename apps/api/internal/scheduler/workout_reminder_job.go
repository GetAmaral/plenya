// Package scheduler — WorkoutReminderJob.
//
// Roda a cada 15min. Pra cada NotificationPreferences com WorkoutReminder=true
// cujo WorkoutReminderTime cai na janela [now-7min, now+8min] em hora local
// do servidor, dispara push pra um plano ativo do paciente.
//
// Janela de 15min com tolerância de ±7min cobre o tick do scheduler sem perder
// usuários nem duplicar (cada userID-dia é processado uma vez via mapa em memória,
// resetado à meia-noite).
//
// Edge cases:
//   - Paciente sem plano ativo: não enviamos (nada útil pra lembrar).
//   - Múltiplos planos: usa o mais recente.
//   - Fuso: assume hora local do servidor (Brazil/SP em produção). Refinar
//     quando paciente puder configurar fuso.
package scheduler

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
)

const (
	workoutReminderJobInterval = 15 * time.Minute
	workoutReminderJobTimeout  = 5 * time.Minute
)

type WorkoutReminderJob struct {
	db        *gorm.DB
	push      *services.PushService
	notifPref *services.NotificationPreferencesService

	// Dedup intra-day: userID -> "YYYY-MM-DD" do último envio.
	// Reset implícito quando muda a chave de data.
	mu       sync.Mutex
	lastSent map[uuid.UUID]string
}

func NewWorkoutReminderJob(
	db *gorm.DB,
	push *services.PushService,
	notifPref *services.NotificationPreferencesService,
) *WorkoutReminderJob {
	return &WorkoutReminderJob{
		db:        db,
		push:      push,
		notifPref: notifPref,
		lastSent:  make(map[uuid.UUID]string),
	}
}

func (j *WorkoutReminderJob) Run() error {
	ctx, cancel := context.WithTimeout(context.Background(), workoutReminderJobTimeout)
	defer cancel()

	now := time.Now()
	dayKey := now.Format("2006-01-02")
	nowMin := now.Hour()*60 + now.Minute()

	var prefs []models.NotificationPreferences
	err := j.db.WithContext(ctx).
		Where("workout_reminder = ? AND deleted_at IS NULL", true).
		Find(&prefs).Error
	if err != nil {
		log.Printf("⚠️  [WORKOUT REMINDER] query falhou: %v", err)
		return err
	}

	sent, skipped := 0, 0
	for _, p := range prefs {
		min, ok := parseHHMM(p.WorkoutReminderTime)
		if !ok {
			continue
		}
		// Tolerância ±7min em torno do horário configurado
		if diff := absInt(nowMin - min); diff > 7 {
			continue
		}
		// Dedup intra-day
		j.mu.Lock()
		if last := j.lastSent[p.UserID]; last == dayKey {
			j.mu.Unlock()
			skipped++
			continue
		}
		j.mu.Unlock()

		// Verifica se tem plano ativo (sem plano = nada a lembrar)
		var planCount int64
		j.db.WithContext(ctx).
			Table("workout_plans AS wp").
			Joins("JOIN patients pat ON pat.id = wp.patient_id").
			Where("pat.user_id = ? AND wp.is_active = true AND wp.deleted_at IS NULL", p.UserID).
			Count(&planCount)
		if planCount == 0 {
			skipped++
			continue
		}

		err := j.push.Send(p.UserID, services.PushPayload{
			Title: "Hora do treino 💪",
			Body:  "Bora? Toca pra abrir seu plano.",
			URL:   "/treino",
			Data:  map[string]any{"kind": "workout_reminder"},
		})
		if err != nil {
			log.Printf("⚠️  [WORKOUT REMINDER] user=%s: %v", p.UserID, err)
			continue
		}
		j.mu.Lock()
		j.lastSent[p.UserID] = dayKey
		j.mu.Unlock()
		sent++
	}
	if sent > 0 || skipped > 0 {
		log.Printf("✅ [WORKOUT REMINDER] enviados=%d skipped=%d", sent, skipped)
	}
	return nil
}

func (j *WorkoutReminderJob) Start() {
	log.Printf("💪 [WORKOUT REMINDER] iniciado (interval=%s)", workoutReminderJobInterval)
	go func() {
		if err := j.Run(); err != nil {
			log.Printf("⚠️  [WORKOUT REMINDER] erro inicial: %v", err)
		}
		ticker := time.NewTicker(workoutReminderJobInterval)
		defer ticker.Stop()
		for range ticker.C {
			if err := j.Run(); err != nil {
				log.Printf("⚠️  [WORKOUT REMINDER] erro: %v", err)
			}
		}
	}()
}

func parseHHMM(s string) (int, bool) {
	if len(s) != 5 || s[2] != ':' {
		return 0, false
	}
	h := int(s[0]-'0')*10 + int(s[1]-'0')
	m := int(s[3]-'0')*10 + int(s[4]-'0')
	if h < 0 || h > 23 || m < 0 || m > 59 {
		return 0, false
	}
	return h*60 + m, true
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
