// Package scheduler — ConversationAutoReplyJob (Atendimento IA).
//
// Tick curto (15s). A cada tick, para cada conversa de WhatsApp com inbound do lead aguardando
// resposta (dentro da janela de 24h), resolve o modo EFETIVO (flag da conversa → global, com a
// janela de horário) e age conforme a máquina de estados do plano
// docs/emr/plano-atendimento-ia-global.md:
//
//   Off      → se Off por herança (sem flag explícita) e passou Z sem resposta: alerta o admin.
//   Copiloto → após X (debounce): gera/atualiza o rascunho (mostrado no compositor). Se passar Y
//              sem resposta: escala e envia (Auto).
//   Auto     → após X: gera o rascunho (janela de preview X/3 no frontend). Após X+X/3: envia.
//
// Handoff (dúvida clínica / pedido de humano): não envia, alerta e pausa a conversa.
// Flag por conversa expira 24h após a última mensagem (sweep no início do tick).
//
// Kill switch global: cfg.ReceptionBot.Enabled. Com false, o job não roda.
package scheduler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
)

const (
	autoReplyJobInterval = 5 * time.Second // tick curto pra honrar X pequeno e dar janela de preview
	autoReplyJobTimeout  = 4 * time.Minute
	autoReplyWindowHours = 24 * time.Hour
	handoffPauseDuration = 24 * time.Hour
	sendFailCooldown     = 15 * time.Minute
)

type ConversationAutoReplyJob struct {
	db          *gorm.DB
	convSvc     *services.ConversationService
	autoSvc     *services.ConversationAutomationService
	settingsSvc *services.ReceptionSettingsService
	notifSvc    *services.NotificationService
	cfg         *config.Config
}

func NewConversationAutoReplyJob(
	db *gorm.DB,
	convSvc *services.ConversationService,
	autoSvc *services.ConversationAutomationService,
	settingsSvc *services.ReceptionSettingsService,
	notifSvc *services.NotificationService,
	cfg *config.Config,
) *ConversationAutoReplyJob {
	return &ConversationAutoReplyJob{db: db, convSvc: convSvc, autoSvc: autoSvc, settingsSvc: settingsSvc, notifSvc: notifSvc, cfg: cfg}
}

type ownerRef struct {
	OwnerType string
	OwnerID   uuid.UUID
}

func (j *ConversationAutoReplyJob) Run() error {
	if !j.cfg.ReceptionBot.Enabled {
		return nil // auto-envio desligado globalmente (kill switch)
	}
	ctx, cancel := context.WithTimeout(context.Background(), autoReplyJobTimeout)
	defer cancel()

	j.expireStaleFlags(ctx)

	settings, err := j.settingsSvc.Get(ctx)
	if err != nil {
		log.Printf("⚠️  [ATENDIMENTO IA] settings: %v", err)
		return err
	}

	owners, err := j.listPendingOwners(ctx)
	if err != nil {
		log.Printf("⚠️  [ATENDIMENTO IA] query candidatos: %v", err)
		return err
	}
	for _, o := range owners {
		j.handleConversation(ctx, o.OwnerType, o.OwnerID, settings)
	}
	return nil
}

// expireStaleFlags apaga overrides por conversa sem nenhuma atividade nas últimas 24h
// (regra: a flag vale por 24h após a última mensagem; depois volta a seguir o global).
func (j *ConversationAutoReplyJob) expireStaleFlags(ctx context.Context) {
	cutoff := time.Now().UTC().Add(-autoReplyWindowHours)
	err := j.db.WithContext(ctx).Exec(`
		DELETE FROM conversation_automations ca
		WHERE NOT EXISTS (
			SELECT 1 FROM lead_activities la
			WHERE ((ca.owner_type = 'lead' AND la.lead_id = ca.owner_id)
			    OR (ca.owner_type = 'patient' AND la.patient_id = ca.owner_id))
			  AND la.created_at > ?
		)`, cutoff).Error
	if err != nil {
		log.Printf("⚠️  [ATENDIMENTO IA] expira flags: %v", err)
	}
}

// listPendingOwners retorna donos com algum inbound de WhatsApp nas últimas 24h.
func (j *ConversationAutoReplyJob) listPendingOwners(ctx context.Context) ([]ownerRef, error) {
	since := time.Now().UTC().Add(-autoReplyWindowHours)
	var rows []ownerRef
	err := j.db.WithContext(ctx).Raw(`
		SELECT DISTINCT 'lead'::text AS owner_type, lead_id AS owner_id FROM lead_activities
		 WHERE type = ? AND channel = ? AND lead_id IS NOT NULL AND created_at > ?
		UNION
		SELECT DISTINCT 'patient'::text, patient_id FROM lead_activities
		 WHERE type = ? AND channel = ? AND patient_id IS NOT NULL AND created_at > ?`,
		models.LeadActivityMessageReceived, models.LeadChannelWhatsApp, since,
		models.LeadActivityMessageReceived, models.LeadChannelWhatsApp, since,
	).Scan(&rows).Error
	return rows, err
}

// handleConversation aplica a máquina de estados a uma conversa.
func (j *ConversationAutoReplyJob) handleConversation(ctx context.Context, ownerType string, ownerID uuid.UUID, settings *models.ReceptionSettings) {
	ownerCol := "lead_id"
	if ownerType == string(models.ConversationOwnerPatient) {
		ownerCol = "patient_id"
	}

	// Lead opt-out / descadastro não recebe IA (LGPD).
	if ownerType == string(models.ConversationOwnerLead) {
		var l models.Lead
		if err := j.db.WithContext(ctx).Select("status", "whats_app_opt_in").First(&l, "id = ?", ownerID).Error; err != nil {
			return
		}
		if l.Status == models.LeadStatusUnsubscribed || !l.WhatsAppOptIn {
			return
		}
	}

	// Último inbound do lead.
	var lastIn models.LeadActivity
	if err := j.db.WithContext(ctx).
		Where(ownerCol+" = ?", ownerID).
		Where("type = ?", models.LeadActivityMessageReceived).
		Where("channel = ?", models.LeadChannelWhatsApp).
		Order("created_at DESC").First(&lastIn).Error; err != nil {
		return
	}
	if lastIn.Content != nil && services.IsUnsubscribeKeyword(*lastIn.Content) {
		return
	}

	age := time.Now().UTC().Sub(lastIn.CreatedAt)
	if age > autoReplyWindowHours {
		return // fora da janela de 24h (free-form bloqueado; precisa template)
	}

	// Já respondido (humano ou bot) depois do inbound? → limpa rascunho e sai.
	var answered int64
	j.db.WithContext(ctx).Model(&models.LeadActivity{}).
		Where(ownerCol+" = ?", ownerID).
		Where("type = ?", models.LeadActivityMessageSent).
		Where("created_at > ?", lastIn.CreatedAt).Count(&answered)
	if answered > 0 {
		// Fim do atendimento (alguém respondeu): captura o que sobrou na memória social.
		j.convSvc.MaintainRelationshipSummary(ctx, ownerType, ownerID, true)
		j.clearDraft(ctx, ownerType, ownerID)
		return
	}

	// Memória de longo prazo (social): regenera o resumo a cada +5 mensagens novas. Best-effort,
	// gatilho interno barato (só chama a IA quando acumulou o suficiente). Guardrail §3.1.
	j.convSvc.MaintainRelationshipSummary(ctx, ownerType, ownerID, false)

	mode, explicitFlag, paused := j.resolveMode(ctx, ownerType, ownerID, settings)
	if paused {
		return // handoff em andamento
	}

	x := time.Duration(settings.DebounceSeconds) * time.Second
	y := time.Duration(settings.CopilotFallbackMinutes) * time.Minute
	z := time.Duration(settings.OffAlertMinutes) * time.Minute

	switch mode {
	case models.ConversationAutomationOff:
		// Z só para Off por herança (sem flag explícita).
		if !explicitFlag && age >= z {
			j.maybeAlertOff(ctx, ownerType, ownerID, lastIn.CreatedAt)
		}
	case models.ConversationAutomationCopilot:
		if age >= y {
			j.sendAuto(ctx, ownerType, ownerID) // escala Copiloto → Auto
		} else if age >= x {
			j.ensureDraft(ctx, ownerType, ownerID, lastIn.ID)
		}
	case models.ConversationAutomationAuto:
		if age >= x {
			// Sempre rascunha primeiro (preenche o compositor + "digitando…"); só envia depois
			// que o rascunho ficou visível pela janela de preview (x/3), nunca no mesmo tick em
			// que foi criado. Garante que a mensagem aparece antes de sair, mesmo com X pequeno.
			j.convSvc.MarkConversationTyping(ctx, ownerType, ownerID)
			j.ensureDraft(ctx, ownerType, ownerID, lastIn.ID)
			preview := time.Duration(settings.PreviewSeconds) * time.Second // X2 (janela de preview)
			if preview <= 0 {
				preview = x / 3
			}
			if view, _ := j.autoSvc.GetSuggestedReply(ctx, ownerType, ownerID); view != nil &&
				view.BasedOnActivityID != nil && *view.BasedOnActivityID == lastIn.ID &&
				time.Since(view.UpdatedAt) >= preview {
				j.sendAuto(ctx, ownerType, ownerID)
			}
		}
	}
}

// resolveMode devolve o modo efetivo da conversa: flag (se houver linha; stale já foi varrida) →
// senão o global efetivo (com janela de horário).
func (j *ConversationAutoReplyJob) resolveMode(ctx context.Context, ownerType string, ownerID uuid.UUID, settings *models.ReceptionSettings) (mode string, explicit bool, paused bool) {
	var row models.ConversationAutomation
	err := j.db.WithContext(ctx).Where("owner_type = ? AND owner_id = ?", ownerType, ownerID).First(&row).Error
	if err == nil {
		if row.PausedUntil != nil && row.PausedUntil.After(time.Now().UTC()) {
			return row.Mode, true, true
		}
		return row.Mode, true, false
	}
	return j.settingsSvc.EffectiveGlobalMode(settings, time.Now()), false, false
}

// ensureDraft gera (ou atualiza) o rascunho da conversa, debounced por inbound. Se o cérebro
// pedir handoff, alerta e pausa (não fica oferecendo resposta clínica).
func (j *ConversationAutoReplyJob) ensureDraft(ctx context.Context, ownerType string, ownerID uuid.UUID, basedOn uuid.UUID) {
	var existing models.ConversationSuggestedReply
	err := j.db.WithContext(ctx).Where("owner_type = ? AND owner_id = ?", ownerType, ownerID).First(&existing).Error
	if err == nil && existing.BasedOnActivityID != nil && *existing.BasedOnActivityID == basedOn {
		return // já tem rascunho fresco para este inbound
	}

	res, gErr := j.convSvc.GenerateReceptionReply(ctx, ownerType, ownerID)
	if gErr != nil {
		log.Printf("⚠️  [ATENDIMENTO IA] rascunho owner=%s/%s: %v", ownerType, ownerID, gErr)
		return
	}

	rec := models.ConversationSuggestedReply{
		OwnerType:         ownerType,
		OwnerID:           ownerID,
		Reply:             res.Reply,
		Action:            res.Action,
		Model:             res.Model,
		BasedOnActivityID: &basedOn,
	}
	if err := j.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "owner_type"}, {Name: "owner_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"reply", "action", "model", "based_on_activity_id", "updated_at"}),
	}).Create(&rec).Error; err != nil {
		log.Printf("⚠️  [ATENDIMENTO IA] salvar rascunho: %v", err)
	}

	if res.Action == "handoff" {
		j.doHandoff(ctx, ownerType, ownerID, res.HandoffReason)
	}
}

// sendAuto gera a resposta na hora, re-checa e envia (Auto / escalada do Copiloto). Trata handoff
// e anti-spam. Limpa o rascunho ao enviar.
func (j *ConversationAutoReplyJob) sendAuto(ctx context.Context, ownerType string, ownerID uuid.UUID) {
	// Anti-spam: limite de mensagens da IA por hora nesta conversa.
	ownerCol := "lead_id"
	if ownerType == string(models.ConversationOwnerPatient) {
		ownerCol = "patient_id"
	}
	var botLastHour int64
	j.db.WithContext(ctx).Model(&models.LeadActivity{}).
		Where(ownerCol+" = ?", ownerID).
		Where("type = ?", models.LeadActivityMessageSent).
		Where("created_at > ?", time.Now().UTC().Add(-1*time.Hour)).
		Where("metadata ->> 'ai_generated' = ?", "true").Count(&botLastHour)
	if botLastHour >= int64(j.cfg.ReceptionBot.MaxMsgsPerHour) {
		return
	}

	// Mostra "digitando…" pro cliente enquanto a IA gera a resposta (e marca a msg como lida).
	j.convSvc.MarkConversationTyping(ctx, ownerType, ownerID)

	res, err := j.convSvc.GenerateReceptionReply(ctx, ownerType, ownerID)
	if err != nil {
		log.Printf("⚠️  [ATENDIMENTO IA] gerar resposta owner=%s/%s: %v", ownerType, ownerID, err)
		return
	}

	if !j.stillSendable(ctx, ownerType, ownerID) {
		return // alguém respondeu / pausou / assumiu no meio
	}

	if res.Action == "handoff" {
		j.doHandoff(ctx, ownerType, ownerID, res.HandoffReason)
		// envia a cortesia do handoff e encerra
	}

	_, sErr := j.convSvc.SendMessage(ctx, services.SendMessageInput{
		UserID:      services.BotUserID,
		OwnerType:   ownerType,
		OwnerID:     ownerID,
		Channel:     models.LeadChannelWhatsApp,
		BodyText:    res.Reply,
		AIGenerated: true,
		AIModel:     res.Model,
	})
	if sErr != nil {
		log.Printf("⚠️  [ATENDIMENTO IA] enviar owner=%s/%s: %v", ownerType, ownerID, sErr)
		if res.Action != "handoff" {
			_ = j.autoSvc.Pause(ctx, ownerType, ownerID, time.Now().UTC().Add(sendFailCooldown))
		}
		return
	}
	j.autoSvc.TouchLastBot(ctx, ownerType, ownerID)
	j.clearDraft(ctx, ownerType, ownerID)
}

// stillSendable re-verifica imediatamente antes do envio: ainda há inbound sem resposta e a
// conversa não foi pausada. Fecha a corrida do intervalo do Claude.
func (j *ConversationAutoReplyJob) stillSendable(ctx context.Context, ownerType string, ownerID uuid.UUID) bool {
	ownerCol := "lead_id"
	if ownerType == string(models.ConversationOwnerPatient) {
		ownerCol = "patient_id"
	}
	var lastIn models.LeadActivity
	if err := j.db.WithContext(ctx).
		Where(ownerCol+" = ?", ownerID).
		Where("type = ?", models.LeadActivityMessageReceived).
		Where("channel = ?", models.LeadChannelWhatsApp).
		Order("created_at DESC").First(&lastIn).Error; err != nil {
		return false
	}
	var answered int64
	j.db.WithContext(ctx).Model(&models.LeadActivity{}).
		Where(ownerCol+" = ?", ownerID).
		Where("type = ?", models.LeadActivityMessageSent).
		Where("created_at > ?", lastIn.CreatedAt).Count(&answered)
	if answered > 0 {
		return false
	}
	var auto models.ConversationAutomation
	if err := j.db.WithContext(ctx).
		Where("owner_type = ? AND owner_id = ?", ownerType, ownerID).First(&auto).Error; err == nil {
		if auto.PausedUntil != nil && auto.PausedUntil.After(time.Now().UTC()) {
			return false
		}
	}
	return true
}

func (j *ConversationAutoReplyJob) clearDraft(ctx context.Context, ownerType string, ownerID uuid.UUID) {
	j.db.WithContext(ctx).Where("owner_type = ? AND owner_id = ?", ownerType, ownerID).
		Delete(&models.ConversationSuggestedReply{})
}

// maybeAlertOff dispara, uma única vez por inbound, o alerta de "conversa parada no Off" para a
// equipe (idempotente via LeadActivity note com metadata off_alert).
func (j *ConversationAutoReplyJob) maybeAlertOff(ctx context.Context, ownerType string, ownerID uuid.UUID, sinceInbound time.Time) {
	ownerCol := "lead_id"
	if ownerType == string(models.ConversationOwnerPatient) {
		ownerCol = "patient_id"
	}
	var already int64
	j.db.WithContext(ctx).Model(&models.LeadActivity{}).
		Where(ownerCol+" = ?", ownerID).
		Where("type = ?", models.LeadActivityNoteAdded).
		Where("created_at > ?", sinceInbound).
		Where("metadata ->> 'off_alert' = ?", "true").Count(&already)
	if already > 0 {
		return
	}
	j.notifyStaff(ctx, ownerType, ownerID,
		"Conversa parada sem atendimento",
		"Uma mensagem está sem resposta e o Atendimento IA está em Off. Você pode entrar e ligar o Auto.",
		map[string]any{"off_alert": true})
}

// doHandoff pausa a IA e avisa a equipe (admin/secretary/manager).
func (j *ConversationAutoReplyJob) doHandoff(ctx context.Context, ownerType string, ownerID uuid.UUID, reason string) {
	_ = j.autoSvc.Pause(ctx, ownerType, ownerID, time.Now().UTC().Add(handoffPauseDuration))
	msg := reason
	if msg == "" {
		msg = "A IA encaminhou esta conversa para atendimento humano."
	}
	j.notifyStaff(ctx, ownerType, ownerID, "Recepção IA pediu ajuda humana", msg, map[string]any{"ai_handoff": true})
}

// notifyStaff registra a activity interna (idempotência/auditoria) e notifica o time.
func (j *ConversationAutoReplyJob) notifyStaff(ctx context.Context, ownerType string, ownerID uuid.UUID, title, msg string, meta map[string]any) {
	var leadID, patientID *uuid.UUID
	if ownerType == string(models.ConversationOwnerPatient) {
		id := ownerID
		patientID = &id
	} else {
		id := ownerID
		leadID = &id
	}

	note := title
	if msg != "" {
		note = title + ": " + msg
	}
	metaJSON, _ := json.Marshal(meta)
	botID := services.BotUserID
	_ = j.db.WithContext(ctx).Create(&models.LeadActivity{
		LeadID:      leadID,
		PatientID:   patientID,
		Type:        models.LeadActivityNoteAdded,
		Channel:     models.LeadChannelInternal,
		Content:     &note,
		Metadata:    datatypes.JSON(metaJSON),
		ActorUserID: &botID,
	}).Error

	if j.notifSvc == nil {
		return
	}
	actionURL := fmt.Sprintf("/conversas?owner=%s&id=%s", ownerType, ownerID.String())
	var users []models.User
	if err := j.db.WithContext(ctx).
		Where(`roles::jsonb ?| array['admin','secretary','manager']`).
		Where("id <> ?", services.BotUserID).
		Find(&users).Error; err != nil {
		log.Printf("⚠️  [ATENDIMENTO IA] lookup staff: %v", err)
		return
	}
	for _, u := range users {
		_ = j.notifSvc.CreateConversationNotification(
			u.ID, leadID, patientID, models.NotificationLeadAssigned, title, msg, actionURL)
	}
}

func (j *ConversationAutoReplyJob) Start() {
	if !j.cfg.ReceptionBot.Enabled {
		log.Printf("ℹ️  [ATENDIMENTO IA] auto-envio DESLIGADO (RECEPTION_BOT_ENABLED=false); job não inicia")
		return
	}
	log.Printf("🤖 [ATENDIMENTO IA] iniciado (interval=%s)", autoReplyJobInterval)
	go func() {
		ticker := time.NewTicker(autoReplyJobInterval)
		defer ticker.Stop()
		for range ticker.C {
			if err := j.Run(); err != nil {
				log.Printf("⚠️  [ATENDIMENTO IA] erro: %v", err)
			}
		}
	}()
}
