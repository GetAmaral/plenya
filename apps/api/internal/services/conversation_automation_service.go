package services

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
)

// BotUserID é o UUID fixo do usuário "Assistente Plenya" (seed na migration 00017).
// Usado como ActorUserID das mensagens enviadas pelo recepcionista virtual.
var BotUserID = uuid.MustParse("019e9301-0000-7000-a000-0000000b0b01")

// ConversationAutomationService gerencia o modo do recepcionista virtual por conversa.
type ConversationAutomationService struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewConversationAutomationService(db *gorm.DB, cfg *config.Config) *ConversationAutomationService {
	return &ConversationAutomationService{db: db, cfg: cfg}
}

// EffectiveAutomation é a configuração resolvida (linha + defaults globais).
type EffectiveAutomation struct {
	OwnerType       string     `json:"ownerType"`
	OwnerID         string     `json:"ownerId"`
	Mode            string     `json:"mode"`
	FallbackMinutes int        `json:"fallbackMinutes"`
	PausedUntil     *time.Time `json:"pausedUntil,omitempty"`
	GloballyEnabled bool       `json:"globallyEnabled"` // RECEPTION_BOT_ENABLED (kill switch)
}

// Get resolve o modo efetivo de uma conversa (linha override ou default global).
func (s *ConversationAutomationService) Get(ctx context.Context, ownerType string, ownerID uuid.UUID) (*EffectiveAutomation, error) {
	if !isValidOwnerType(ownerType) || ownerID == uuid.Nil {
		return nil, ErrConversationOwnerInvalid
	}
	eff := &EffectiveAutomation{
		OwnerType:       ownerType,
		OwnerID:         ownerID.String(),
		Mode:            s.cfg.ReceptionBot.DefaultMode,
		FallbackMinutes: s.cfg.ReceptionBot.FallbackMinutes,
		GloballyEnabled: s.cfg.ReceptionBot.Enabled,
	}
	var row models.ConversationAutomation
	err := s.db.WithContext(ctx).
		Where("owner_type = ? AND owner_id = ?", ownerType, ownerID).
		First(&row).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return eff, nil
		}
		return nil, err
	}
	eff.Mode = row.Mode
	if row.FallbackMinutes > 0 {
		eff.FallbackMinutes = row.FallbackMinutes
	}
	eff.PausedUntil = row.PausedUntil
	return eff, nil
}

// Set faz upsert do modo/fallback de uma conversa (ação do atendente). Reativar (mode
// copilot/auto) limpa a pausa.
func (s *ConversationAutomationService) Set(ctx context.Context, ownerType string, ownerID uuid.UUID, mode string, fallbackMinutes int, updatedBy uuid.UUID) (*EffectiveAutomation, error) {
	if !isValidOwnerType(ownerType) || ownerID == uuid.Nil {
		return nil, ErrConversationOwnerInvalid
	}
	if mode != models.ConversationAutomationOff &&
		mode != models.ConversationAutomationCopilot &&
		mode != models.ConversationAutomationAuto {
		return nil, errors.New("automation: modo inválido")
	}

	var row models.ConversationAutomation
	err := s.db.WithContext(ctx).
		Where("owner_type = ? AND owner_id = ?", ownerType, ownerID).
		First(&row).Error
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		row = models.ConversationAutomation{
			OwnerType:       ownerType,
			OwnerID:         ownerID,
			Mode:            mode,
			FallbackMinutes: fallbackMinutes,
			UpdatedBy:       &updatedBy,
		}
		if err := s.db.WithContext(ctx).Create(&row).Error; err != nil {
			return nil, err
		}
	case err != nil:
		return nil, err
	default:
		row.Mode = mode
		row.FallbackMinutes = fallbackMinutes
		row.UpdatedBy = &updatedBy
		row.PausedUntil = nil // reativar limpa pausa
		if err := s.db.WithContext(ctx).Save(&row).Error; err != nil {
			return nil, err
		}
	}
	return s.Get(ctx, ownerType, ownerID)
}

// Pause suspende a IA até `until` (handoff). Cria a linha se não existir.
func (s *ConversationAutomationService) Pause(ctx context.Context, ownerType string, ownerID uuid.UUID, until time.Time) error {
	var row models.ConversationAutomation
	err := s.db.WithContext(ctx).
		Where("owner_type = ? AND owner_id = ?", ownerType, ownerID).
		First(&row).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		row = models.ConversationAutomation{
			OwnerType:   ownerType,
			OwnerID:     ownerID,
			Mode:        models.ConversationAutomationAuto,
			PausedUntil: &until,
		}
		return s.db.WithContext(ctx).Create(&row).Error
	}
	if err != nil {
		return err
	}
	row.PausedUntil = &until
	return s.db.WithContext(ctx).Save(&row).Error
}

// TouchLastBot marca o último envio da IA (anti-spam/telemetria).
func (s *ConversationAutomationService) TouchLastBot(ctx context.Context, ownerType string, ownerID uuid.UUID) {
	now := time.Now().UTC()
	s.db.WithContext(ctx).Model(&models.ConversationAutomation{}).
		Where("owner_type = ? AND owner_id = ?", ownerType, ownerID).
		Update("last_bot_at", now)
}

// ReceptionMetrics resume a atuação do recepcionista virtual num período.
type ReceptionMetrics struct {
	PeriodDays           int              `json:"periodDays"`
	AutoReplies          int64            `json:"autoReplies"`
	Handoffs             int64            `json:"handoffs"`
	ConversationsWithBot int64            `json:"conversationsWithBot"`
	ConvertedAfterBot    int64            `json:"convertedAfterBot"`
	ByMode               map[string]int64 `json:"byMode"`
	GloballyEnabled      bool             `json:"globallyEnabled"`
}

// ComputeMetrics calcula as métricas do bot nos últimos `days` dias.
func (s *ConversationAutomationService) ComputeMetrics(ctx context.Context, days int) (*ReceptionMetrics, error) {
	if days <= 0 {
		days = 30
	}
	since := time.Now().UTC().AddDate(0, 0, -days)
	m := &ReceptionMetrics{PeriodDays: days, ByMode: map[string]int64{}, GloballyEnabled: s.cfg.ReceptionBot.Enabled}

	s.db.WithContext(ctx).Model(&models.LeadActivity{}).
		Where("type = ?", models.LeadActivityMessageSent).
		Where("created_at > ?", since).
		Where("metadata ->> 'ai_generated' = ?", "true").
		Count(&m.AutoReplies)

	s.db.WithContext(ctx).Model(&models.LeadActivity{}).
		Where("type = ?", models.LeadActivityNoteAdded).
		Where("created_at > ?", since).
		Where("metadata ->> 'ai_handoff' = ?", "true").
		Count(&m.Handoffs)

	// Nota: lead_activities é log imutável append-only (sem coluna deleted_at).
	s.db.WithContext(ctx).Raw(`
		SELECT count(DISTINCT coalesce(lead_id::text, patient_id::text))
		FROM lead_activities
		WHERE type = ? AND created_at > ? AND metadata ->> 'ai_generated' = 'true'`,
		models.LeadActivityMessageSent, since).Scan(&m.ConversationsWithBot)

	s.db.WithContext(ctx).Raw(`
		SELECT count(*) FROM leads
		WHERE status = 'converted' AND deleted_at IS NULL AND id IN (
			SELECT DISTINCT lead_id FROM lead_activities
			WHERE lead_id IS NOT NULL AND type = ? AND created_at > ? AND metadata ->> 'ai_generated' = 'true'
		)`, models.LeadActivityMessageSent, since).Scan(&m.ConvertedAfterBot)

	type modeCount struct {
		Mode  string
		Count int64
	}
	var rows []modeCount
	s.db.WithContext(ctx).Model(&models.ConversationAutomation{}).
		Select("mode, count(*) as count").Group("mode").Scan(&rows)
	for _, r := range rows {
		m.ByMode[r.Mode] = r.Count
	}
	return m, nil
}

// ListAutoCandidates retorna as conversas em modo auto não pausadas.
func (s *ConversationAutomationService) ListAutoCandidates(ctx context.Context) ([]models.ConversationAutomation, error) {
	var rows []models.ConversationAutomation
	now := time.Now().UTC()
	err := s.db.WithContext(ctx).
		Where("mode = ?", models.ConversationAutomationAuto).
		Where("paused_until IS NULL OR paused_until < ?", now).
		Find(&rows).Error
	return rows, err
}
