// Package scheduler — ConversationAutoReplyJob (recepcionista virtual, Fase 2).
//
// Roda a cada minuto. Para cada conversa em modo "auto" (e não pausada), verifica se
// há uma mensagem do cliente sem resposta humana há >= fallback_minutes e ainda dentro
// da janela de 24h. Se sim, gera a resposta do recepcionista virtual e envia (ou faz
// handoff quando a IA pede ajuda humana). Idempotência natural: ao enviar, passa a
// existir um message_sent depois do inbound, então a conversa deixa de ser elegível.
//
// Kill switch global: cfg.ReceptionBot.Enabled. Com false, o job não envia nada.
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

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
)

const (
	autoReplyJobInterval = 1 * time.Minute
	autoReplyJobTimeout  = 4 * time.Minute
	autoReplyWindowHours = 24 * time.Hour
	handoffPauseDuration = 24 * time.Hour
	sendFailCooldown     = 15 * time.Minute
)

type ConversationAutoReplyJob struct {
	db       *gorm.DB
	convSvc  *services.ConversationService
	autoSvc  *services.ConversationAutomationService
	notifSvc *services.NotificationService
	cfg      *config.Config
}

func NewConversationAutoReplyJob(
	db *gorm.DB,
	convSvc *services.ConversationService,
	autoSvc *services.ConversationAutomationService,
	notifSvc *services.NotificationService,
	cfg *config.Config,
) *ConversationAutoReplyJob {
	return &ConversationAutoReplyJob{db: db, convSvc: convSvc, autoSvc: autoSvc, notifSvc: notifSvc, cfg: cfg}
}

func (j *ConversationAutoReplyJob) Run() error {
	if !j.cfg.ReceptionBot.Enabled {
		return nil // auto-envio desligado globalmente
	}
	ctx, cancel := context.WithTimeout(context.Background(), autoReplyJobTimeout)
	defer cancel()

	candidates, err := j.autoSvc.ListAutoCandidates(ctx)
	if err != nil {
		log.Printf("⚠️  [RECEPÇÃO IA] query candidatos: %v", err)
		return err
	}
	if len(candidates) == 0 {
		return nil
	}

	sent, handoffs := 0, 0
	for _, c := range candidates {
		fallback := c.FallbackMinutes
		if fallback <= 0 {
			fallback = j.cfg.ReceptionBot.FallbackMinutes
		}
		if !j.eligible(ctx, c.OwnerType, c.OwnerID, fallback) {
			continue
		}
		did, isHandoff := j.handleOne(ctx, c.OwnerType, c.OwnerID)
		if did {
			if isHandoff {
				handoffs++
			} else {
				sent++
			}
		}
	}
	if sent > 0 || handoffs > 0 {
		log.Printf("✅ [RECEPÇÃO IA] enviadas=%d handoffs=%d (de %d candidatos)", sent, handoffs, len(candidates))
	}
	return nil
}

// eligible decide se a conversa precisa de resposta da IA agora.
func (j *ConversationAutoReplyJob) eligible(ctx context.Context, ownerType string, ownerID uuid.UUID, fallbackMin int) bool {
	ownerCol := "lead_id"
	if ownerType == string(models.ConversationOwnerPatient) {
		ownerCol = "patient_id"
	}

	// Lead que pediu descadastro ou sem opt-in não recebe resposta da IA (LGPD).
	if ownerType == string(models.ConversationOwnerLead) {
		var l models.Lead
		if err := j.db.WithContext(ctx).Select("status", "whats_app_opt_in").First(&l, "id = ?", ownerID).Error; err != nil {
			return false
		}
		if l.Status == models.LeadStatusUnsubscribed || !l.WhatsAppOptIn {
			return false
		}
	}

	// Última mensagem recebida do cliente (WhatsApp).
	var lastIn models.LeadActivity
	err := j.db.WithContext(ctx).
		Where(ownerCol+" = ?", ownerID).
		Where("type = ?", models.LeadActivityMessageReceived).
		Where("channel = ?", models.LeadChannelWhatsApp).
		Order("created_at DESC").
		First(&lastIn).Error
	if err != nil {
		return false // sem inbound → nada a responder
	}

	// Pedido de parar/descadastrar (qualquer canal/owner): a IA não responde. Enforce,
	// não confia só no prompt. Cobre paciente (que não tem flag de opt-out como o lead).
	if lastIn.Content != nil && services.IsUnsubscribeKeyword(*lastIn.Content) {
		return false
	}

	now := time.Now().UTC()
	age := now.Sub(lastIn.CreatedAt)
	if age < time.Duration(fallbackMin)*time.Minute {
		return false // ainda dentro do tempo do humano responder
	}
	if age > autoReplyWindowHours {
		return false // fora da janela de 24h da Meta (free-form bloqueado)
	}

	// Já respondida (humano ou bot) depois do inbound?
	var answered int64
	j.db.WithContext(ctx).Model(&models.LeadActivity{}).
		Where(ownerCol+" = ?", ownerID).
		Where("type = ?", models.LeadActivityMessageSent).
		Where("created_at > ?", lastIn.CreatedAt).
		Count(&answered)
	if answered > 0 {
		return false
	}

	// Anti-spam: limite de mensagens da IA por hora nesta conversa.
	var botLastHour int64
	j.db.WithContext(ctx).Model(&models.LeadActivity{}).
		Where(ownerCol+" = ?", ownerID).
		Where("type = ?", models.LeadActivityMessageSent).
		Where("created_at > ?", now.Add(-1*time.Hour)).
		Where("metadata ->> 'ai_generated' = ?", "true").
		Count(&botLastHour)
	if botLastHour >= int64(j.cfg.ReceptionBot.MaxMsgsPerHour) {
		return false
	}

	return true
}

// stillSendable re-verifica, imediatamente antes do envio, que a conversa ainda precisa de
// resposta: há inbound recente sem resposta posterior e a automação não foi pausada nem
// trocada de modo no meio-tempo. Fecha a corrida do intervalo do Claude.
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
		if auto.Mode != models.ConversationAutomationAuto {
			return false
		}
		if auto.PausedUntil != nil && auto.PausedUntil.After(time.Now().UTC()) {
			return false
		}
	}
	return true
}

// handleOne gera e age sobre uma conversa. Retorna (agiu, foiHandoff).
func (j *ConversationAutoReplyJob) handleOne(ctx context.Context, ownerType string, ownerID uuid.UUID) (bool, bool) {
	res, err := j.convSvc.GenerateReceptionReply(ctx, ownerType, ownerID)
	if err != nil {
		log.Printf("⚠️  [RECEPÇÃO IA] gerar resposta owner=%s/%s: %v", ownerType, ownerID, err)
		return false, false
	}

	// Re-checa logo antes de enviar: o Claude leva ~segundos; nesse intervalo um humano
	// (ou o copiloto) pode ter respondido, ou a conversa pode ter sido pausada. Fecha a
	// janela de corrida pra IA não falar por cima de alguém da equipe.
	if !j.stillSendable(ctx, ownerType, ownerID) {
		return false, false
	}

	// Envia a mensagem (handoff também envia a cortesia "vou chamar a equipe").
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
		log.Printf("⚠️  [RECEPÇÃO IA] enviar owner=%s/%s: %v", ownerType, ownerID, sErr)
		// Mesmo sem conseguir enviar, se for handoff, pausa + notifica pra humano assumir.
		if res.Action == "handoff" {
			j.doHandoff(ctx, ownerType, ownerID, res.HandoffReason)
			return true, true
		}
		// Falha persistente (ex: Meta rejeitou, janela expirou na corrida): pausa por um
		// cooldown curto pra não re-chamar o Claude a cada minuto nessa conversa.
		_ = j.autoSvc.Pause(ctx, ownerType, ownerID, time.Now().UTC().Add(sendFailCooldown))
		return false, false
	}

	j.autoSvc.TouchLastBot(ctx, ownerType, ownerID)

	if res.Action == "handoff" {
		j.doHandoff(ctx, ownerType, ownerID, res.HandoffReason)
		return true, true
	}
	return true, false
}

// doHandoff pausa a IA e avisa a equipe (admin/secretary/manager).
func (j *ConversationAutoReplyJob) doHandoff(ctx context.Context, ownerType string, ownerID uuid.UUID, reason string) {
	_ = j.autoSvc.Pause(ctx, ownerType, ownerID, time.Now().UTC().Add(handoffPauseDuration))

	var leadID, patientID *uuid.UUID
	if ownerType == string(models.ConversationOwnerPatient) {
		id := ownerID
		patientID = &id
	} else {
		id := ownerID
		leadID = &id
	}

	// Registra o handoff como activity interna (auditoria + métrica).
	note := "Recepção IA encaminhou para humano"
	if reason != "" {
		note = note + ": " + reason
	}
	metaJSON, _ := json.Marshal(map[string]any{"ai_handoff": true})
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
	title := "Recepção IA pediu ajuda humana"
	msg := reason
	if msg == "" {
		msg = "A IA encaminhou esta conversa para atendimento humano."
	}
	actionURL := fmt.Sprintf("/conversas?owner=%s&id=%s", ownerType, ownerID.String())

	var users []models.User
	if err := j.db.WithContext(ctx).
		Where(`roles::jsonb ?| array['admin','secretary','manager']`).
		Where("id <> ?", services.BotUserID).
		Find(&users).Error; err != nil {
		log.Printf("⚠️  [RECEPÇÃO IA] lookup staff p/ handoff: %v", err)
		return
	}
	for _, u := range users {
		_ = j.notifSvc.CreateConversationNotification(
			u.ID, leadID, patientID,
			models.NotificationLeadAssigned,
			title, msg, actionURL,
		)
	}
}

func (j *ConversationAutoReplyJob) Start() {
	if !j.cfg.ReceptionBot.Enabled {
		log.Printf("ℹ️  [RECEPÇÃO IA] auto-envio DESLIGADO (RECEPTION_BOT_ENABLED=false); job não inicia")
		return
	}
	log.Printf("🤖 [RECEPÇÃO IA] iniciado (interval=%s, fallback=%dmin)", autoReplyJobInterval, j.cfg.ReceptionBot.FallbackMinutes)
	go func() {
		ticker := time.NewTicker(autoReplyJobInterval)
		defer ticker.Stop()
		for range ticker.C {
			if err := j.Run(); err != nil {
				log.Printf("⚠️  [RECEPÇÃO IA] erro: %v", err)
			}
		}
	}()
}
