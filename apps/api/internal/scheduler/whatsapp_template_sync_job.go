package scheduler

import (
	"context"
	"log"
	"time"

	"github.com/plenya/api/internal/services"
)

// WhatsAppTemplateSyncJob sincroniza o catálogo de templates WhatsApp com a Meta periodicamente,
// mantendo o status (PENDING/APPROVED/REJECTED) sempre atualizado sem ação manual.
type WhatsAppTemplateSyncJob struct {
	svc *services.WhatsAppTemplateService
}

func NewWhatsAppTemplateSyncJob(svc *services.WhatsAppTemplateService) *WhatsAppTemplateSyncJob {
	return &WhatsAppTemplateSyncJob{svc: svc}
}

const waTemplateSyncInterval = 12 * time.Hour

func (j *WhatsAppTemplateSyncJob) Start() {
	go func() {
		time.Sleep(30 * time.Second) // deixa o boot terminar antes do 1º sync
		j.run()
		ticker := time.NewTicker(waTemplateSyncInterval)
		defer ticker.Stop()
		for range ticker.C {
			j.run()
		}
	}()
}

func (j *WhatsAppTemplateSyncJob) run() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	n, err := j.svc.SyncFromMeta(ctx)
	if err != nil {
		log.Printf("[wa-template-sync] erro: %v", err)
		return
	}
	log.Printf("[wa-template-sync] %d templates sincronizados", n)
}
