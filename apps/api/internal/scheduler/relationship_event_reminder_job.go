// Package scheduler — RelationshipEventReminderJob (Fase C do Atendimento IA / Lívia).
//
// Roda 1×/dia. Duas tarefas:
//  1. Sincroniza aniversários: para pacientes já engajados (com relationship_profile) que têm
//     birth_date, garante um evento recorrente de aniversário (Upsert, idempotente).
//  2. Avisa o time: varre os próximos eventos de relacionamento e, para cada um cuja próxima
//     ocorrência cai dentro da antecedência (lead_time_days), notifica a equipe UMA vez por
//     ocorrência (idempotência via LeadActivity marcador), para alguém interagir (parabéns,
//     mensagem de apoio, follow-up). Tom da marca; nada clínico (só social).
//
// Ver docs/emr/plano-livia-memoria-dossie-360.md §5.
package scheduler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
)

const (
	relEventReminderTimeout = 5 * time.Minute
	relEventScanWindowDays  = 60 // janela ampla; o filtro real é o lead_time de cada evento
)

// Horários (horário de Londrina/SP) em que o job roda. A 1ª passada pega a maioria; as outras
// duas pegam qualquer falha. A idempotência (marcador por evento+ocorrência) evita aviso dobrado.
var relEventReminderHours = []int{9, 12, 16}

type RelationshipEventReminderJob struct {
	db          *gorm.DB
	convSvc     *services.ConversationService
	notifSvc    *services.NotificationService
	settingsSvc *services.ReceptionSettingsService
	cfg         *config.Config
}

func NewRelationshipEventReminderJob(db *gorm.DB, convSvc *services.ConversationService, notifSvc *services.NotificationService, settingsSvc *services.ReceptionSettingsService, cfg *config.Config) *RelationshipEventReminderJob {
	return &RelationshipEventReminderJob{db: db, convSvc: convSvc, notifSvc: notifSvc, settingsSvc: settingsSvc, cfg: cfg}
}

func (j *RelationshipEventReminderJob) Run() error {
	ctx, cancel := context.WithTimeout(context.Background(), relEventReminderTimeout)
	defer cancel()

	j.syncPatientBirthdays(ctx)

	now := time.Now()
	items, err := j.convSvc.UpcomingEvents(ctx, relEventScanWindowDays, now)
	if err != nil {
		log.Printf("⚠️  [AVISOS] próximos eventos: %v", err)
		return err
	}
	// Settings globais (modo/janela) p/ decidir se o cumprimento sai automático.
	settings, _ := j.settingsSvc.Get(ctx)

	for _, it := range items {
		if it.Status != models.RelationshipEventPending {
			continue
		}
		if it.DaysUntil > it.LeadTimeDays {
			continue // ainda não chegou a janela de antecedência deste evento
		}
		occ := it.NextDate.Format("2006-01-02")
		if j.alreadyAlerted(ctx, it.OwnerType, it.OwnerID, it.ID, occ) {
			continue
		}
		// Tenta cumprimentar automático (só aniversário próprio, não-restrito, modo auto, template
		// configurado, kill switch ligado). Se enviar, marca a ocorrência. Senão, notifica o time.
		if j.tryAutoGreet(ctx, it, occ, settings) {
			continue
		}
		j.alertStaff(ctx, it, occ)
	}
	return nil
}

// tryAutoGreet envia o cumprimento de aniversário automático quando TODOS os guardrails passam.
// Retorna true se enviou (e marcou a ocorrência); false → cai no aviso ao time.
func (j *RelationshipEventReminderJob) tryAutoGreet(ctx context.Context, it services.UpcomingEventItem, occ string, settings *models.ReceptionSettings) bool {
	// Allowlist do "óbvio": só aniversário PRÓPRIO (não de pessoa da rede), não-restrito, só no dia.
	if it.Type != models.RelationshipEventBirthday || it.Restricted || it.RelatedPersonID != nil || it.DaysUntil != 0 {
		return false
	}
	if !j.cfg.ReceptionBot.Enabled || j.cfg.WhatsApp.TemplateBirthday == "" || settings == nil {
		return false
	}
	if j.settingsSvc.EffectiveGlobalMode(settings, time.Now()) != models.ConversationAutomationAuto {
		return false // só envia sozinho se o Atendimento IA está em auto
	}
	ownerID, err := uuid.Parse(it.OwnerID)
	if err != nil {
		return false
	}
	firstName := ""
	if fields := strings.Fields(strings.TrimSpace(it.OwnerName)); len(fields) > 0 {
		firstName = fields[0]
	}
	if firstName == "" {
		return false
	}
	if _, err := j.convSvc.SendWhatsAppTemplate(ctx, services.SendWhatsAppTemplateInput{
		UserID:       services.BotUserID,
		OwnerType:    it.OwnerType,
		OwnerID:      ownerID,
		TemplateName: j.cfg.WhatsApp.TemplateBirthday,
		Language:     "pt_BR",
		Params:       []string{firstName},
	}); err != nil {
		log.Printf("⚠️  [AVISOS] cumprimento auto falhou (cai p/ aviso ao time): %v", err)
		return false
	}
	// Marca a ocorrência (idempotência) e registra que foi enviado.
	j.writeEventMarker(ctx, it.OwnerType, ownerID, it.ID, occ,
		"🎂 Cumprimento de aniversário enviado automaticamente pela Lívia.", true)
	return true
}

// syncPatientBirthdays garante evento de aniversário recorrente para pacientes engajados com birth_date.
func (j *RelationshipEventReminderJob) syncPatientBirthdays(ctx context.Context) {
	type row struct {
		PatientID uuid.UUID
		Name      string
		BirthDate time.Time
	}
	var rows []row
	err := j.db.WithContext(ctx).
		Table("relationship_profiles rp").
		Select("rp.owner_id as patient_id, p.name, p.birth_date").
		Joins("JOIN patients p ON p.id = rp.owner_id").
		Where("rp.owner_type = ?", string(models.ConversationOwnerPatient)).
		Where("p.birth_date > ?", "1900-01-01").
		Where("p.deleted_at IS NULL").
		Scan(&rows).Error
	if err != nil {
		log.Printf("⚠️  [AVISOS] sync aniversários: %v", err)
		return
	}
	evSvc := services.NewRelationshipEventService(j.db)
	for _, r := range rows {
		if r.BirthDate.IsZero() {
			continue
		}
		_, _ = evSvc.Upsert(ctx, string(models.ConversationOwnerPatient), r.PatientID, nil,
			models.RelationshipEventBirthday, "Aniversário", r.BirthDate, true, 7,
			"system", "", nil, false)
	}
}

// alreadyAlerted verifica se já avisamos sobre este evento nesta ocorrência (idempotência).
func (j *RelationshipEventReminderJob) alreadyAlerted(ctx context.Context, ownerType, ownerID, eventID, occ string) bool {
	ownerCol := "lead_id"
	if ownerType == string(models.ConversationOwnerPatient) {
		ownerCol = "patient_id"
	}
	var n int64
	j.db.WithContext(ctx).Model(&models.LeadActivity{}).
		Where(ownerCol+" = ?", ownerID).
		Where("type = ?", models.LeadActivityNoteAdded).
		Where("metadata ->> 'rel_event_alert' = ?", eventID).
		Where("metadata ->> 'occ' = ?", occ).Count(&n)
	return n > 0
}

// writeEventMarker grava o marcador de ocorrência (idempotência + auditoria). greeted=true quando
// o cumprimento foi enviado automaticamente.
func (j *RelationshipEventReminderJob) writeEventMarker(ctx context.Context, ownerType string, ownerID uuid.UUID, eventID, occ, note string, greeted bool) {
	var leadID, patientID *uuid.UUID
	if ownerType == string(models.ConversationOwnerPatient) {
		patientID = &ownerID
	} else {
		leadID = &ownerID
	}
	meta, _ := json.Marshal(map[string]any{"rel_event_alert": eventID, "occ": occ, "greeted": greeted})
	botID := services.BotUserID
	_ = j.db.WithContext(ctx).Create(&models.LeadActivity{
		LeadID:      leadID,
		PatientID:   patientID,
		Type:        models.LeadActivityNoteAdded,
		Channel:     models.LeadChannelInternal,
		Content:     &note,
		Metadata:    meta,
		ActorUserID: &botID,
	}).Error
}

// alertStaff registra o marcador (idempotência/auditoria) e notifica o time.
func (j *RelationshipEventReminderJob) alertStaff(ctx context.Context, it services.UpcomingEventItem, occ string) {
	ownerID, err := uuid.Parse(it.OwnerID)
	if err != nil {
		return
	}
	var leadID, patientID *uuid.UUID
	if it.OwnerType == string(models.ConversationOwnerPatient) {
		patientID = &ownerID
	} else {
		leadID = &ownerID
	}

	whenLabel := "hoje"
	if it.DaysUntil == 1 {
		whenLabel = "amanhã"
	} else if it.DaysUntil > 1 {
		whenLabel = fmt.Sprintf("em %d dias", it.DaysUntil)
	}
	title := relEventTitle(it.Type)
	who := it.OwnerName
	if who == "" {
		who = "um contato"
	}
	msg := fmt.Sprintf("%s — %s (%s). Um bom momento para uma mensagem no tom da Plenya.", it.Title, who, whenLabel)

	j.writeEventMarker(ctx, it.OwnerType, ownerID, it.ID, occ, title+": "+msg, false)

	if j.notifSvc == nil {
		return
	}
	actionURL := fmt.Sprintf("/conversas?owner=%s&id=%s", it.OwnerType, it.OwnerID)
	var users []models.User
	if err := j.db.WithContext(ctx).
		Where(`roles::jsonb ?| array['admin','secretary','manager']`).
		Where("id <> ?", services.BotUserID).
		Find(&users).Error; err != nil {
		log.Printf("⚠️  [AVISOS] lookup staff: %v", err)
		return
	}
	for _, u := range users {
		_ = j.notifSvc.CreateConversationNotification(
			u.ID, leadID, patientID, models.NotificationLeadAssigned, title, msg, actionURL)
	}
}

func relEventTitle(t string) string {
	switch t {
	case models.RelationshipEventBirthday:
		return "Aniversário chegando"
	case models.RelationshipEventGraduation:
		return "Formatura chegando"
	case models.RelationshipEventBirth:
		return "Nascimento na família"
	case models.RelationshipEventWedding:
		return "Casamento chegando"
	case models.RelationshipEventLoss:
		return "Data sensível"
	case models.RelationshipEventFollowup:
		return "Follow-up"
	default:
		return "Data importante"
	}
}

func (j *RelationshipEventReminderJob) Start() {
	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		loc = time.UTC
	}
	log.Printf("🎂 [AVISOS] RelationshipEventReminderJob iniciado (roda às %v, horário de SP)", relEventReminderHours)
	go func() {
		for {
			next := nextReminderRun(time.Now().In(loc), relEventReminderHours, loc)
			time.Sleep(time.Until(next))
			if err := j.Run(); err != nil {
				log.Printf("⚠️  [AVISOS] erro: %v", err)
			}
		}
	}()
}

// nextReminderRun devolve o próximo horário de execução (entre os horários do dia, senão o
// primeiro horário de amanhã).
func nextReminderRun(now time.Time, hours []int, loc *time.Location) time.Time {
	for _, h := range hours {
		cand := time.Date(now.Year(), now.Month(), now.Day(), h, 0, 0, 0, loc)
		if cand.After(now) {
			return cand
		}
	}
	t := now.AddDate(0, 0, 1)
	return time.Date(t.Year(), t.Month(), t.Day(), hours[0], 0, 0, 0, loc)
}
