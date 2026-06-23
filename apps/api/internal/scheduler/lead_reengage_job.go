// Package scheduler — LeadReengageJob.
//
// Reengaja leads "frios" (status contacted/qualified, opt-in WhatsApp, sem inbound
// há 7+ dias) com o template reengajamento_lead, uma única vez por lead.
//
// DESLIGADO por default: só roda se cfg.WhatsApp.TemplateReengage != "" (env
// WHATSAPP_TEMPLATE_REENGAGE=reengajamento_lead). É outreach a lead, então é opt-in
// explícito. O envio passa por ConversationService.SendWhatsAppTemplate, que já
// bloqueia leads sem opt-in ou unsubscribed e registra a atividade na conversa.
// Opt-out "PARAR/SAIR/STOP/..." já é tratado no webhook (UnsubscribeByPhone).
//
// Recortes anti-blast: só leads criados nos últimos 90 dias; lote de no máximo
// reengageBatchLimit por passada; idempotência via reengaged_at.
package scheduler

import (
	"context"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
)

const (
	reengageJobInterval = 6 * time.Hour
	reengageJobTimeout  = 5 * time.Minute
	reengageColdDays    = 7
	reengageMaxAgeDays  = 90
	reengageBatchLimit  = 50
)

type LeadReengageJob struct {
	db       *gorm.DB
	convSvc  *services.ConversationService
	template string // cfg.WhatsApp.TemplateReengage ("" = desativado)
}

func NewLeadReengageJob(db *gorm.DB, convSvc *services.ConversationService, template string) *LeadReengageJob {
	return &LeadReengageJob{db: db, convSvc: convSvc, template: template}
}

func (j *LeadReengageJob) Run() error {
	if j.template == "" {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), reengageJobTimeout)
	defer cancel()

	now := time.Now().UTC()
	cold := now.Add(-reengageColdDays * 24 * time.Hour)
	maxAge := now.Add(-reengageMaxAgeDays * 24 * time.Hour)

	var leads []models.Lead
	err := j.db.WithContext(ctx).
		Where("status IN ?", []models.LeadStatus{
			models.LeadStatusContacted,
			models.LeadStatusQualified,
		}).
		Where("whats_app_opt_in = ?", true).
		Where("phone IS NOT NULL AND phone <> ''").
		Where("reengaged_at IS NULL").
		Where("converted_patient_id IS NULL").
		Where("created_at > ?", maxAge).
		Where("(last_inbound_at IS NULL AND created_at < ?) OR last_inbound_at < ?", cold, cold).
		Order("created_at ASC").
		Limit(reengageBatchLimit).
		Find(&leads).Error
	if err != nil {
		log.Printf("⚠️  [REENGAGE JOB] query falhou: %v", err)
		return err
	}
	if len(leads) == 0 {
		return nil
	}

	log.Printf("📨 [REENGAGE JOB] %d leads frios p/ reengajar", len(leads))
	ok, ko, skip := 0, 0, 0
	for i := range leads {
		l := &leads[i]
		name := "tudo bem"
		if l.Name != nil && *l.Name != "" {
			name = firstWord(*l.Name)
		}
		_, err := j.convSvc.SendWhatsAppTemplate(ctx, services.SendWhatsAppTemplateInput{
			UserID:       services.BotUserID,
			OwnerType:    string(models.ConversationOwnerLead),
			OwnerID:      l.ID,
			TemplateName: j.template,
			Language:     "pt_BR",
			Params:       []string{name},
		})
		if err != nil {
			// Opt-out / sem canal → marca como reengajado pra não retentar pra sempre.
			if errors.Is(err, services.ErrConversationOptedOut) || errors.Is(err, services.ErrConversationNoChannel) {
				j.markReengaged(ctx, l.ID, now)
				skip++
				continue
			}
			log.Printf("⚠️  [REENGAGE JOB] lead=%s: %v", l.ID, err)
			ko++
			continue
		}
		j.markReengaged(ctx, l.ID, now)
		ok++
	}
	log.Printf("✅ [REENGAGE JOB] enviados=%d pulados=%d erros=%d", ok, skip, ko)
	return nil
}

func (j *LeadReengageJob) markReengaged(ctx context.Context, id uuid.UUID, ts time.Time) {
	if err := j.db.WithContext(ctx).Model(&models.Lead{}).
		Where("id = ?", id).Update("reengaged_at", ts).Error; err != nil {
		log.Printf("⚠️  [REENGAGE JOB] persist reengaged_at lead=%s: %v", id, err)
	}
}

func (j *LeadReengageJob) Start() {
	if j.template == "" {
		log.Printf("📨 [REENGAGE JOB] desativado (WHATSAPP_TEMPLATE_REENGAGE vazio)")
		return
	}
	log.Printf("📨 [REENGAGE JOB] iniciado (interval=%s, template=%s)", reengageJobInterval, j.template)
	go func() {
		if err := j.Run(); err != nil {
			log.Printf("⚠️  [REENGAGE JOB] erro inicial: %v", err)
		}
		ticker := time.NewTicker(reengageJobInterval)
		defer ticker.Stop()
		for range ticker.C {
			if err := j.Run(); err != nil {
				log.Printf("⚠️  [REENGAGE JOB] erro: %v", err)
			}
		}
	}()
}

// firstWord devolve a primeira palavra (primeiro nome) de uma string.
func firstWord(s string) string {
	fields := strings.Fields(s)
	if len(fields) == 0 {
		return s
	}
	return fields[0]
}
