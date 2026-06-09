package services

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
)

// LeadService gerencia leads (captura, listagem, conversão em Patient).
type LeadService struct {
	db                  *gorm.DB
	notificationService *NotificationService
	whatsappService     *WhatsAppService
	emailService        *EmailService
	cfg                 *config.Config

	// patientInboundWAHook é um callback opcional disparado APÓS registrar
	// uma mensagem inbound de paciente já existente (não Lead). Calendar V1
	// (Bloco F) usa pra rodar IA Claude e detectar intent confirm/cancel/
	// reschedule de consultas. Setado via SetPatientInboundWAHook após DI
	// — evita import cycle entre LeadService ↔ ConversationService/AppointmentService.
	patientInboundWAHook PatientInboundWAHook

	// Serviços de mídia (WhatsApp Fase 2), injetados via SetMediaServices após DI.
	// patientDocs salva arquivo de paciente nas mídias do prontuário; waMedia
	// guarda mídia de conversa (áudio/lead) cifrada.
	patientDocs *PatientDocumentsService
	waMedia     *WhatsAppMediaService
	transcriber *TranscriptionService
}

// PatientInboundWAHook é o callback chamado quando paciente cadastrado envia
// mensagem WhatsApp inbound. Roda async (já dentro de goSafe pelo LeadService),
// não retorna erro pra não bloquear o webhook Meta.
type PatientInboundWAHook func(patient *models.Patient, message string, waMessageID string)

func NewLeadService(
	db *gorm.DB,
	notificationService *NotificationService,
	whatsappService *WhatsAppService,
	emailService *EmailService,
	cfg *config.Config,
) *LeadService {
	return &LeadService{
		db:                  db,
		notificationService: notificationService,
		whatsappService:     whatsappService,
		emailService:        emailService,
		cfg:                 cfg,
	}
}

// SetPatientInboundWAHook registra o callback pós-inbound de paciente. main.go
// chama isso após criar ConversationService + AppointmentService, fechando o
// ciclo sem dep direta nos services.
func (s *LeadService) SetPatientInboundWAHook(h PatientInboundWAHook) {
	s.patientInboundWAHook = h
}

// SetMediaServices injeta os serviços de mídia (WhatsApp Fase 2) após DI em main.go.
func (s *LeadService) SetMediaServices(docs *PatientDocumentsService, media *WhatsAppMediaService) {
	s.patientDocs = docs
	s.waMedia = media
}

// SetTranscriber injeta o serviço de transcrição de áudio (Fase 2.2).
func (s *LeadService) SetTranscriber(t *TranscriptionService) {
	s.transcriber = t
}

// HashIP retorna SHA-256 hex do IP — usado pra registrar consentimento sem armazenar IP plano.
func HashIP(ip string) string {
	h := sha256.Sum256([]byte(strings.TrimSpace(ip)))
	return hex.EncodeToString(h[:])
}

// ============================================================
// Captura — Light claim (multi-canal, idempotente)
// ============================================================

// CreateOrUpdateFromLightClaimInput é o payload pra registrar/atualizar Lead vindo do claim.
type CreateOrUpdateFromLightClaimInput struct {
	SessionID        uuid.UUID
	Email            *string
	Phone            *string
	NewsletterOptIn  bool
	ConsentVersion   string
	ConsentTimestamp time.Time
	ConsentIPHash    string // SHA-256 do IP — não passar IP plano

	// UTM atribuição (propagado da AnonymousScoreSession capturada no site).
	UTMSource   *string
	UTMMedium   *string
	UTMCampaign *string
	UTMTerm     *string
}

// CreateOrUpdateFromLightClaim é idempotente por SessionID — se já existe Lead vinculado,
// atualiza opt-ins e contatos. Caso contrário, cria novo.
// Retorna o Lead resultante (com ID).
func (s *LeadService) CreateOrUpdateFromLightClaim(in CreateOrUpdateFromLightClaimInput) (*models.Lead, error) {
	if in.SessionID == uuid.Nil {
		return nil, errors.New("lead: sessionID obrigatório")
	}
	if in.Email == nil && in.Phone == nil {
		return nil, errors.New("lead: pelo menos um de email/phone obrigatório")
	}

	var lead models.Lead
	err := s.db.Where("anonymous_score_session_id = ?", in.SessionID).First(&lead).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("lead: lookup: %w", err)
	}

	consentTS := in.ConsentTimestamp
	if consentTS.IsZero() {
		consentTS = time.Now().UTC()
	}
	consentVer := in.ConsentVersion
	consentIP := in.ConsentIPHash

	if errors.Is(err, gorm.ErrRecordNotFound) {
		newLead := models.Lead{
			Source:                  models.LeadSourceLightClaim,
			Status:                  models.LeadStatusNew,
			Email:                   normalizeEmailPtr(in.Email),
			Phone:                   in.Phone,
			EmailOptIn:              in.Email != nil,
			WhatsAppOptIn:           in.Phone != nil,
			NewsletterOptIn:         in.NewsletterOptIn,
			ConsentVersion:          &consentVer,
			ConsentTimestamp:        &consentTS,
			ConsentIPHash:           &consentIP,
			AnonymousScoreSessionID: &in.SessionID,
			UTMSource:               in.UTMSource,
			UTMMedium:               in.UTMMedium,
			UTMCampaign:             in.UTMCampaign,
			UTMTerm:                 in.UTMTerm,
		}
		if err := s.db.Create(&newLead).Error; err != nil {
			// Race: outra request criou lead pra mesma sessão entre nosso SELECT e INSERT.
			// O uniqueIndex pega; refazemos lookup e seguimos pro update path.
			if isUniqueViolation(err) {
				if err := s.db.Where("anonymous_score_session_id = ?", in.SessionID).First(&lead).Error; err != nil {
					return nil, fmt.Errorf("lead: race recovery lookup: %w", err)
				}
				// fall through para o update path abaixo
			} else {
				return nil, fmt.Errorf("lead: create: %w", err)
			}
		} else {
			newLeadID := newLead.ID
			_ = s.RecordActivity(RecordActivityInput{
				LeadID:  &newLeadID,
				Type:    models.LeadActivityCreated,
				Channel: models.LeadChannelInternal,
				Content: ptr("Lead criado via claim do Escore Light"),
			})
			// Notifica equipe — async, só na criação (re-claim do mesmo session_id não notifica)
			leadCopy := newLead
			goSafe("notify_team", func() { s.NotifyTeamOfNewLead(&leadCopy) })
			return &newLead, nil
		}
	}

	// Update existente — preserva valores que não vieram, mas amplia opt-ins
	updates := map[string]any{
		"updated_at":        time.Now().UTC(),
		"newsletter_opt_in": lead.NewsletterOptIn || in.NewsletterOptIn,
		"consent_version":   consentVer,
		"consent_timestamp": consentTS,
		"consent_ip_hash":   consentIP,
	}
	// UTM — preencher só se ainda não houver atribuição (first-touch wins, evita
	// que um re-claim sobrescreva a atribuição original da sessão).
	if lead.UTMSource == nil && in.UTMSource != nil {
		updates["utm_source"] = *in.UTMSource
	}
	if lead.UTMMedium == nil && in.UTMMedium != nil {
		updates["utm_medium"] = *in.UTMMedium
	}
	if lead.UTMCampaign == nil && in.UTMCampaign != nil {
		updates["utm_campaign"] = *in.UTMCampaign
	}
	if lead.UTMTerm == nil && in.UTMTerm != nil {
		updates["utm_term"] = *in.UTMTerm
	}
	if in.Email != nil {
		updates["email"] = strings.ToLower(strings.TrimSpace(*in.Email))
		updates["email_opt_in"] = true
	}
	if in.Phone != nil {
		updates["phone"] = *in.Phone
		updates["whats_app_opt_in"] = true
	}
	if err := s.db.Model(&lead).Updates(updates).Error; err != nil {
		return nil, fmt.Errorf("lead: update: %w", err)
	}
	// Refetch
	if err := s.db.First(&lead, "id = ?", lead.ID).Error; err != nil {
		return nil, err
	}
	return &lead, nil
}

// ============================================================
// Captura — Contact form (público)
// ============================================================

type CreateFromContactFormInput struct {
	Name             string
	Email            string
	Phone            string
	Message          string
	Metadata         map[string]any // ex: { reason, window }
	ConsentVersion   string
	ConsentTimestamp time.Time
	ConsentIPHash    string
	NewsletterOptIn  bool
}

func (s *LeadService) CreateFromContactForm(in CreateFromContactFormInput) (*models.Lead, error) {
	if strings.TrimSpace(in.Email) == "" && strings.TrimSpace(in.Phone) == "" {
		return nil, errors.New("lead: email ou telefone obrigatório")
	}
	if in.ConsentVersion == "" {
		return nil, errors.New("lead: consentVersion obrigatório (LGPD)")
	}

	consentTS := in.ConsentTimestamp
	if consentTS.IsZero() {
		consentTS = time.Now().UTC()
	}

	var metaJSON datatypes.JSON
	if len(in.Metadata) > 0 {
		raw, err := json.Marshal(in.Metadata)
		if err != nil {
			return nil, fmt.Errorf("lead: metadata marshal: %w", err)
		}
		metaJSON = raw
	}

	name := strings.TrimSpace(in.Name)
	email := strings.ToLower(strings.TrimSpace(in.Email))
	phone := strings.TrimSpace(in.Phone)

	lead := models.Lead{
		Source:           models.LeadSourceContactForm,
		Status:           models.LeadStatusNew,
		EmailOptIn:       email != "",
		WhatsAppOptIn:    phone != "", // form de contato pede telefone como WhatsApp
		NewsletterOptIn:  in.NewsletterOptIn,
		Metadata:         metaJSON,
		ConsentVersion:   &in.ConsentVersion,
		ConsentTimestamp: &consentTS,
		ConsentIPHash:    &in.ConsentIPHash,
	}
	if name != "" {
		lead.Name = &name
	}
	if email != "" {
		lead.Email = &email
	}
	if phone != "" {
		lead.Phone = &phone
	}
	if msg := strings.TrimSpace(in.Message); msg != "" {
		lead.Message = &msg
	}

	if err := s.db.Create(&lead).Error; err != nil {
		return nil, fmt.Errorf("lead: create from contact form: %w", err)
	}

	leadID := lead.ID
	_ = s.RecordActivity(RecordActivityInput{
		LeadID:  &leadID,
		Type:    models.LeadActivityCreated,
		Channel: models.LeadChannelInternal,
		Content: ptr("Lead criado via formulário /contato"),
	})

	// Notifica equipe — async, best-effort
	leadCopy := lead
	goSafe("notify_team", func() { s.NotifyTeamOfNewLead(&leadCopy) })

	return &lead, nil
}

// ============================================================
// WhatsApp inbound — primeiro contato vira Lead
// ============================================================

type InboundWhatsAppInput struct {
	PhoneE164   string // sem +, formato Meta (ex: 5511999998888)
	Name        *string
	Text        string
	WAMessageID string
	ReceivedAt  time.Time
}

// InboundResult descreve o destino da activity criada por ProcessInbound{Email,WhatsApp}.
//
// Exatamente um de Lead|Patient é não-nil:
//   - Patient setado → activity foi anexada a um cliente cadastrado (sem criar Lead novo).
//   - Lead setado    → activity foi anexada a um Lead (existente OU recém-criado).
//
// IsNew = true SOMENTE quando Lead foi criado nesta chamada (pra notificação diferenciada).
type InboundResult struct {
	Lead    *models.Lead
	Patient *models.Patient
	IsNew   bool
}

// ProcessInboundWhatsApp registra mensagem inbound. Hierarquia de roteamento:
//
//  1. Procura Patient pelo phone → se acha, activity vai pro Patient. FIM.
//  2. Procura Lead ativo (não converted/lost/unsubscribed) → se acha, activity vai pro Lead.
//  3. Cria Lead novo com source=whatsapp_inbound (consent implícito).
//
// Diferente do comportamento anterior, NÃO cria mais Lead "fantasma" pra Patient
// já cadastrado — paciente é cliente, conversa fica anexada nele direto.
// waInboundAlreadyProcessed indica se já registramos um inbound WhatsApp com este
// wa_message_id (idempotência — a Meta reentrega webhooks).
func (s *LeadService) waInboundAlreadyProcessed(waMessageID string) bool {
	if strings.TrimSpace(waMessageID) == "" {
		return false
	}
	var n int64
	s.db.Model(&models.LeadActivity{}).
		Where("type = ? AND channel = ? AND metadata->>'wa_message_id' = ?",
			models.LeadActivityMessageReceived, models.LeadChannelWhatsApp, waMessageID).
		Count(&n)
	return n > 0
}

func (s *LeadService) ProcessInboundWhatsApp(in InboundWhatsAppInput) (*InboundResult, error) {
	if strings.TrimSpace(in.PhoneE164) == "" {
		return nil, errors.New("lead: phone vazio em inbound WhatsApp")
	}
	if s.waInboundAlreadyProcessed(in.WAMessageID) {
		return &InboundResult{}, nil // já processado (reentrega)
	}

	// Normaliza para formato armazenado em Lead.Phone (com +).
	// Patient.Phone também é normalizado pra +55... no BeforeSave hook (NormalizePhoneBR),
	// então comparar com phone aqui é seguro.
	phone := strings.TrimSpace(in.PhoneE164)
	if !strings.HasPrefix(phone, "+") {
		phone = "+" + phone
	}

	now := in.ReceivedAt
	if now.IsZero() {
		now = time.Now().UTC()
	}

	// 1. Patient first — cliente cadastrado tem prioridade.
	var existingPatient models.Patient
	patientErr := s.db.Where("phone = ?", phone).First(&existingPatient).Error
	if patientErr == nil {
		patientID := existingPatient.ID
		_ = s.RecordActivity(RecordActivityInput{
			PatientID: &patientID,
			Type:      models.LeadActivityMessageReceived,
			Channel:   models.LeadChannelWhatsApp,
			Content:   &in.Text,
			Metadata: map[string]any{
				"wa_message_id": in.WAMessageID,
				"received_at":   in.ReceivedAt,
			},
		})
		patientCopy := existingPatient
		snippet := in.Text
		goSafe("notify_inbound_wa_patient", func() {
			s.NotifyTeamOfInboundMessage(ConversationOwner{Patient: &patientCopy}, models.LeadChannelWhatsApp, snippet)
		})
		// Calendar V1 Bloco F: dispara IA pra detectar intent de
		// confirm/cancel/reschedule. Hook é opcional (nil = sem IA).
		if s.patientInboundWAHook != nil {
			waMsgID := in.WAMessageID
			goSafe("appt_intent_detect", func() {
				s.patientInboundWAHook(&patientCopy, snippet, waMsgID)
			})
		}
		return &InboundResult{Patient: &existingPatient}, nil
	}
	if !errors.Is(patientErr, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("lead: inbound WA patient lookup: %w", patientErr)
	}

	// 2. Lead ativo existente
	var existing models.Lead
	err := s.db.
		Where("phone = ? AND status NOT IN (?)", phone,
			[]models.LeadStatus{models.LeadStatusConverted, models.LeadStatusLost, models.LeadStatusUnsubscribed}).
		Order("created_at DESC").
		First(&existing).Error

	if err == nil {
		existingID := existing.ID
		_ = s.RecordActivity(RecordActivityInput{
			LeadID:  &existingID,
			Type:    models.LeadActivityMessageReceived,
			Channel: models.LeadChannelWhatsApp,
			Content: &in.Text,
			Metadata: map[string]any{
				"wa_message_id": in.WAMessageID,
				"received_at":   in.ReceivedAt,
			},
		})
		_ = s.db.Model(&existing).Update("last_inbound_at", now).Error
		leadCopy := existing
		snippet := in.Text
		goSafe("notify_inbound_wa", func() {
			s.NotifyTeamOfInboundMessage(ConversationOwner{Lead: &leadCopy}, models.LeadChannelWhatsApp, snippet)
		})
		return &InboundResult{Lead: &existing}, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("lead: inbound WA lookup: %w", err)
	}

	// 3. Cria Lead novo (Patient já checado acima, não existe)
	newLead, err := s.createWALeadFromInbound(phone, in.Name, in.Text, now)
	if err != nil {
		return nil, err
	}

	newLeadID := newLead.ID
	_ = s.RecordActivity(RecordActivityInput{
		LeadID:  &newLeadID,
		Type:    models.LeadActivityMessageReceived,
		Channel: models.LeadChannelWhatsApp,
		Content: &in.Text,
		Metadata: map[string]any{
			"wa_message_id": in.WAMessageID,
			"received_at":   now,
		},
	})

	// Notifica equipe — async, best-effort. Só dispara em Lead novo.
	leadCopy := *newLead
	goSafe("notify_team", func() { s.NotifyTeamOfNewLead(&leadCopy) })

	return &InboundResult{Lead: newLead, IsNew: true}, nil
}

// createWALeadFromInbound cria um Lead a partir do primeiro contato inbound do
// WhatsApp (consent implícito) + registra a activity "created". O caller registra
// em seguida a mensagem (texto ou mídia) e dispara a notificação de Lead novo.
func (s *LeadService) createWALeadFromInbound(phone string, name *string, firstMessage string, now time.Time) (*models.Lead, error) {
	consentVer := "wa_inbound_v1" // 13 chars, respeita varchar(20) do schema
	consentIP := ""               // inbound WA não tem IP do cliente
	msg := firstMessage

	newLead := models.Lead{
		Source:           models.LeadSourceWhatsAppInbound,
		Status:           models.LeadStatusNew,
		Phone:            &phone,
		Name:             name,
		Message:          &msg,
		WhatsAppOptIn:    true, // cliente iniciou conversa = consent implícito
		ConsentVersion:   &consentVer,
		ConsentTimestamp: &now,
		ConsentIPHash:    &consentIP,
		LastInboundAt:    &now,
	}
	if err := s.db.Create(&newLead).Error; err != nil {
		return nil, fmt.Errorf("lead: create from WA inbound: %w", err)
	}
	newLeadID := newLead.ID
	_ = s.RecordActivity(RecordActivityInput{
		LeadID:  &newLeadID,
		Type:    models.LeadActivityCreated,
		Channel: models.LeadChannelInternal,
		Content: ptr("Lead criado via primeira mensagem inbound do WhatsApp"),
	})
	return &newLead, nil
}

// ============================================================
// WhatsApp inbound — MÍDIA (Fase 2)
// ============================================================

// InboundWhatsAppMediaInput é o payload de uma mensagem inbound que carrega mídia
// (foto, áudio/nota de voz, documento, figurinha, vídeo).
type InboundWhatsAppMediaInput struct {
	PhoneE164   string // sem +, formato Meta
	Name        *string
	MediaID     string // id da Meta para baixar o binário
	Kind        string // image|audio|voice|video|document|sticker
	MIME        string // MIME informado pela Meta (fallback do detectado no download)
	Filename    string // nome do arquivo (documents)
	Caption     string // legenda (image/video/document)
	WAMessageID string
	ReceivedAt  time.Time
}

// ProcessInboundWhatsAppMedia baixa a mídia da Meta e a roteia:
//   - Paciente cadastrado + arquivo (foto/PDF) → PatientDocument (mídias do prontuário).
//   - Paciente + áudio/voz/sticker/vídeo, ou Lead (sem prontuário) → bucket cifrado da conversa.
//
// Mesma hierarquia de dono do ProcessInboundWhatsApp (patient → lead ativo → lead novo).
func (s *LeadService) ProcessInboundWhatsAppMedia(in InboundWhatsAppMediaInput) (*InboundResult, error) {
	if s.whatsappService == nil || s.waMedia == nil || s.patientDocs == nil {
		return nil, errors.New("lead: serviços de mídia WhatsApp não configurados")
	}
	if strings.TrimSpace(in.PhoneE164) == "" {
		return nil, errors.New("lead: phone vazio em inbound de mídia WhatsApp")
	}
	if s.waInboundAlreadyProcessed(in.WAMessageID) {
		return &InboundResult{}, nil // já processado — evita download + doc/activity duplicados
	}
	phone := strings.TrimSpace(in.PhoneE164)
	if !strings.HasPrefix(phone, "+") {
		phone = "+" + phone
	}
	now := in.ReceivedAt
	if now.IsZero() {
		now = time.Now().UTC()
	}

	// Baixa o binário antes de rotear (precisamos do MIME real pra decidir prontuário).
	dl, err := s.whatsappService.DownloadMedia(in.MediaID)
	if err != nil {
		return nil, fmt.Errorf("lead: download de mídia WA: %w", err)
	}
	mime := dl.MIME
	if mime == "" {
		mime = in.MIME
	}
	filename := strings.TrimSpace(in.Filename)
	if filename == "" {
		filename = mediaLabel(in.Kind) + extFromMime(mime)
	}

	// 1. Paciente cadastrado
	var patient models.Patient
	perr := s.db.Where("phone = ?", phone).First(&patient).Error
	if perr == nil {
		return s.recordPatientMedia(&patient, in, dl, mime, filename, now)
	}
	if !errors.Is(perr, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("lead: inbound WA media patient lookup: %w", perr)
	}

	// 2. Lead ativo existente
	var lead models.Lead
	lerr := s.db.
		Where("phone = ? AND status NOT IN (?)", phone,
			[]models.LeadStatus{models.LeadStatusConverted, models.LeadStatusLost, models.LeadStatusUnsubscribed}).
		Order("created_at DESC").
		First(&lead).Error
	if lerr == nil {
		return s.recordLeadMedia(&lead, in, dl, mime, filename, now, false)
	}
	if !errors.Is(lerr, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("lead: inbound WA media lead lookup: %w", lerr)
	}

	// 3. Cria Lead novo
	first := in.Caption
	if first == "" {
		first = "[" + mediaLabel(in.Kind) + " recebido por WhatsApp]"
	}
	newLead, err := s.createWALeadFromInbound(phone, in.Name, first, now)
	if err != nil {
		return nil, err
	}
	return s.recordLeadMedia(newLead, in, dl, mime, filename, now, true)
}

// recordPatientMedia anexa mídia a um paciente. Arquivo clínico (foto/PDF) vai
// pro prontuário (PatientDocument); o resto vai pro bucket cifrado da conversa.
func (s *LeadService) recordPatientMedia(patient *models.Patient, in InboundWhatsAppMediaInput, dl *MediaDownload, mime, filename string, now time.Time) (*InboundResult, error) {
	patientID := patient.ID
	rec := RecordActivityInput{
		PatientID:      &patientID,
		Type:           models.LeadActivityMessageReceived,
		Channel:        models.LeadChannelWhatsApp,
		Metadata:       map[string]any{"wa_message_id": in.WAMessageID, "received_at": now},
		MediaType:      ptr(in.Kind),
		MediaMIME:      ptr(mime),
		MediaFilename:  ptr(filename),
		MediaSizeBytes: ptr(dl.FileSize),
	}
	if in.Caption != "" {
		rec.Content = ptr(in.Caption)
	}

	if isProntuarioFile(in.Kind, mime) {
		doc, err := s.patientDocs.CreateFromBytes(CreateFromBytesInput{
			PatientID:         patientID,
			Bytes:             dl.Bytes,
			MIME:              mime,
			Filename:          filename,
			Title:             prontuarioTitle(in.Caption, filename),
			Type:              models.DocumentTypeOther,
			Source:            models.DocumentSourceWhatsApp,
			OriginWAMessageID: ptr(in.WAMessageID),
		})
		if err != nil {
			// Prontuário rejeitou (ex: magic bytes não batem com o MIME) — não perde
			// o arquivo: guarda no bucket cifrado da conversa.
			log.Printf("⚠️  [WA MEDIA] prontuário rejeitou (%v) — fallback bucket conversa", err)
			if err2 := s.storeConvMedia(&rec, "patient", patientID, dl, mime); err2 != nil {
				return nil, err2
			}
		} else {
			rec.PatientDocumentID = &doc.ID
		}
	} else {
		if err := s.storeConvMedia(&rec, "patient", patientID, dl, mime); err != nil {
			return nil, err
		}
	}

	act, err := s.recordActivityReturning(rec)
	if err != nil {
		return nil, err
	}
	s.maybeTranscribe(act.ID, in.Kind, dl.Bytes, mime)

	patientCopy := *patient
	snippet := mediaSnippet(in.Kind, in.Caption)
	goSafe("notify_inbound_wa_patient_media", func() {
		s.NotifyTeamOfInboundMessage(ConversationOwner{Patient: &patientCopy}, models.LeadChannelWhatsApp, snippet)
	})
	return &InboundResult{Patient: patient}, nil
}

// recordLeadMedia anexa mídia a um Lead (sem prontuário → sempre bucket cifrado).
func (s *LeadService) recordLeadMedia(lead *models.Lead, in InboundWhatsAppMediaInput, dl *MediaDownload, mime, filename string, now time.Time, isNew bool) (*InboundResult, error) {
	leadID := lead.ID
	rec := RecordActivityInput{
		LeadID:         &leadID,
		Type:           models.LeadActivityMessageReceived,
		Channel:        models.LeadChannelWhatsApp,
		Metadata:       map[string]any{"wa_message_id": in.WAMessageID, "received_at": now},
		MediaType:      ptr(in.Kind),
		MediaMIME:      ptr(mime),
		MediaFilename:  ptr(filename),
		MediaSizeBytes: ptr(dl.FileSize),
	}
	if in.Caption != "" {
		rec.Content = ptr(in.Caption)
	}
	if err := s.storeConvMedia(&rec, "lead", leadID, dl, mime); err != nil {
		return nil, err
	}
	act, err := s.recordActivityReturning(rec)
	if err != nil {
		return nil, err
	}
	s.maybeTranscribe(act.ID, in.Kind, dl.Bytes, mime)
	if !isNew {
		_ = s.db.Model(lead).Update("last_inbound_at", now).Error
	}

	leadCopy := *lead
	snippet := mediaSnippet(in.Kind, in.Caption)
	if isNew {
		goSafe("notify_team", func() { s.NotifyTeamOfNewLead(&leadCopy) })
	} else {
		goSafe("notify_inbound_wa_media", func() {
			s.NotifyTeamOfInboundMessage(ConversationOwner{Lead: &leadCopy}, models.LeadChannelWhatsApp, snippet)
		})
	}
	return &InboundResult{Lead: lead, IsNew: isNew}, nil
}

// maybeTranscribe transcreve áudio inbound (async) e grava na coluna transcription
// (cifrada via BeforeSave). No-op se transcrição desabilitada ou não for áudio.
func (s *LeadService) maybeTranscribe(activityID uuid.UUID, kind string, data []byte, mime string) {
	if s.transcriber == nil || !s.transcriber.Enabled() {
		return
	}
	if kind != "audio" && kind != "voice" {
		return
	}
	goSafe("transcribe_wa_audio", func() {
		text, err := s.transcriber.Transcribe(context.Background(), data, mime)
		if err != nil {
			log.Printf("⚠️  [WA TRANSCRIBE] %v", err)
			return
		}
		if strings.TrimSpace(text) == "" {
			return
		}
		var act models.LeadActivity
		if err := s.db.First(&act, "id = ?", activityID).Error; err != nil {
			return
		}
		act.Transcription = &text
		if err := s.db.Save(&act).Error; err != nil { // BeforeSave cifra a transcrição
			log.Printf("⚠️  [WA TRANSCRIBE] save: %v", err)
		}
	})
}

// RecordOutboundWhatsAppEcho registra no histórico uma mensagem que a EQUIPE enviou
// pelo APP do celular (coexistence — o webhook recebe o eco). Regras:
//   - NÃO cria Lead (é outbound; criar lead a partir disso seria errado).
//   - NÃO notifica a equipe e NÃO mexe na janela 24h (LastInboundAt).
//   - Dedup por wa_message_id (cobre reentrega e o eco do nosso próprio envio via EMR).
//
// Se o contato (cliente) não existe como Patient nem Lead ativo, ignora.
func (s *LeadService) RecordOutboundWhatsAppEcho(customerPhoneE164, text, waMessageID string, ts time.Time) {
	phone := strings.TrimSpace(customerPhoneE164)
	if phone == "" {
		return
	}
	if !strings.HasPrefix(phone, "+") {
		phone = "+" + phone
	}

	if waMessageID != "" {
		var n int64
		s.db.Model(&models.LeadActivity{}).
			Where("type = ? AND channel = ? AND metadata->>'wa_message_id' = ?",
				models.LeadActivityMessageSent, models.LeadChannelWhatsApp, waMessageID).Count(&n)
		if n > 0 {
			return // já registrado (eco reentregue ou enviado pelo próprio EMR)
		}
	}

	rec := RecordActivityInput{
		Type:    models.LeadActivityMessageSent,
		Channel: models.LeadChannelWhatsApp,
		Metadata: map[string]any{
			"wa_message_id": waMessageID,
			"origin":        "phone_app",
			"sent_at":       ts,
		},
	}
	if strings.TrimSpace(text) != "" {
		rec.Content = &text
	}

	var patient models.Patient
	if err := s.db.Where("phone = ?", phone).First(&patient).Error; err == nil {
		pid := patient.ID
		rec.PatientID = &pid
		_ = s.RecordActivity(rec)
		return
	}
	var lead models.Lead
	if err := s.db.
		Where("phone = ? AND status NOT IN (?)", phone,
			[]models.LeadStatus{models.LeadStatusConverted, models.LeadStatusLost, models.LeadStatusUnsubscribed}).
		Order("created_at DESC").First(&lead).Error; err == nil {
		lid := lead.ID
		rec.LeadID = &lid
		_ = s.RecordActivity(rec)
		return
	}
	log.Printf("📱 [WA ECHO] eco do app pra contato sem lead/patient (%s) — ignorado", maskPhone(phone))
}

// maskPhone mostra só os últimos 4 dígitos pra log (LGPD): ****8899.
func maskPhone(p string) string {
	if len(p) <= 4 {
		return "****"
	}
	return "****" + p[len(p)-4:]
}

// storeConvMedia cifra e grava a mídia no bucket da conversa, preenchendo os
// campos de storage no RecordActivityInput.
func (s *LeadService) storeConvMedia(rec *RecordActivityInput, ownerType string, ownerID uuid.UUID, dl *MediaDownload, mime string) error {
	key, size, err := s.waMedia.StoreInbound(ownerType, ownerID, dl.Bytes, mime)
	if err != nil {
		return fmt.Errorf("lead: store media conversa: %w", err)
	}
	rec.MediaStorageKey = &key
	rec.MediaSizeBytes = &size
	return nil
}

// isProntuarioFile decide se a mídia vai pro prontuário: só arquivos (foto/documento)
// com MIME clínico (PDF/JPG/PNG). Áudio/voz/sticker/vídeo ficam na conversa.
func isProntuarioFile(kind, mime string) bool {
	if kind != "image" && kind != "document" {
		return false
	}
	return allowedDocContentTypes[mime]
}

func mediaLabel(kind string) string {
	switch kind {
	case "image":
		return "imagem"
	case "audio", "voice":
		return "áudio"
	case "document":
		return "documento"
	case "video":
		return "vídeo"
	case "sticker":
		return "figurinha"
	default:
		return "arquivo"
	}
}

func mediaSnippet(kind, caption string) string {
	if strings.TrimSpace(caption) != "" {
		return caption
	}
	return "[" + mediaLabel(kind) + "]"
}

func prontuarioTitle(caption, filename string) string {
	if c := strings.TrimSpace(caption); c != "" {
		return c
	}
	if f := strings.TrimSpace(filename); f != "" {
		return f
	}
	return "Arquivo recebido por WhatsApp"
}

// ============================================================
// Email inbound — primeiro contato vira Lead
// ============================================================

// InboundEmailInput é o payload para um email inbound capturado pelo worker IMAP.
type InboundEmailInput struct {
	FromEmail   string    // já normalizado (lower-case)
	FromName    *string   // nome amigável do header From
	Subject     string    // assunto (decodificado)
	BodyText    string    // corpo plain (preferido)
	BodyHTML    string    // corpo HTML (fallback / referência)
	MessageID   string    // RFC 5322 Message-ID
	InReplyTo   string    // header In-Reply-To
	References  []string  // header References (split)
	ReceivedAt  time.Time // header Date (fallback: now)
	Folder      string    // pasta IMAP de origem (ex: "INBOX")
	Attachments []EmailAttachmentRef
}

// EmailAttachmentRef descreve um anexo que JÁ foi salvo em disco pelo worker.
// Não carrega bytes; só metadados pra registrar em LeadActivity.metadata.
type EmailAttachmentRef struct {
	Filename    string `json:"filename"`
	Path        string `json:"path"`
	ContentType string `json:"content_type"`
	SizeBytes   int    `json:"size_bytes"`
}

// ProcessInboundEmail registra um email inbound. Hierarquia de roteamento (mesma do WA):
//
//  1. Procura Patient pelo LOWER(email) → activity vai pro Patient. FIM.
//  2. Procura Lead ativo (não converted/lost/unsubscribed) → activity vai pro Lead.
//  3. Cria Lead novo com source=email_inbound (consent implícito: cliente escreveu pra nós).
//
// Idempotente por Message-ID — chamadas duplicadas não criam activities repetidas
// (cobre re-runs do worker no resume após restart).
func (s *LeadService) ProcessInboundEmail(in InboundEmailInput) (*InboundResult, error) {
	from := strings.ToLower(strings.TrimSpace(in.FromEmail))
	if from == "" {
		return nil, errors.New("lead: from vazio em inbound email")
	}
	now := in.ReceivedAt
	if now.IsZero() {
		now = time.Now().UTC()
	}

	// Dedup por Message-ID — checa em todas activities (Lead OU Patient).
	if in.MessageID != "" {
		var dup int64
		_ = s.db.Model(&models.LeadActivity{}).
			Where("metadata->>'message_id' = ? AND type = ?",
				in.MessageID, models.LeadActivityMessageReceived).
			Count(&dup).Error
		if dup > 0 {
			// Já processado — devolve owner pro caller saber, sem duplicar.
			// Tenta achar Patient primeiro (consistente com nova ordem de roteamento).
			var existingPatient models.Patient
			if err := s.db.Where("LOWER(email) = ?", from).First(&existingPatient).Error; err == nil {
				return &InboundResult{Patient: &existingPatient}, nil
			}
			var existing models.Lead
			if err := s.db.Where("LOWER(email) = ?", from).Order("created_at DESC").First(&existing).Error; err == nil {
				return &InboundResult{Lead: &existing}, nil
			}
			return &InboundResult{}, nil
		}
	}

	bodyContent := in.BodyText
	if bodyContent == "" {
		bodyContent = in.BodyHTML
	}

	metadata := map[string]any{
		"message_id":  in.MessageID,
		"in_reply_to": in.InReplyTo,
		"references":  in.References,
		"subject":     in.Subject,
		"folder":      in.Folder,
		"received_at": now,
	}
	if len(in.Attachments) > 0 {
		metadata["attachments"] = in.Attachments
	}

	// 1. Patient first — cliente cadastrado tem prioridade.
	var existingPatient models.Patient
	patientErr := s.db.Where("LOWER(email) = ?", from).First(&existingPatient).Error
	if patientErr == nil {
		patientID := existingPatient.ID
		_ = s.RecordActivity(RecordActivityInput{
			PatientID: &patientID,
			Type:      models.LeadActivityMessageReceived,
			Channel:   models.LeadChannelEmail,
			Content:   &bodyContent,
			Metadata:  metadata,
		})
		patientCopy := existingPatient
		snippet := bodyContent
		goSafe("notify_inbound_email_patient", func() {
			s.NotifyTeamOfInboundMessage(ConversationOwner{Patient: &patientCopy}, models.LeadChannelEmail, snippet)
		})
		return &InboundResult{Patient: &existingPatient}, nil
	}
	if !errors.Is(patientErr, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("lead: inbound email patient lookup: %w", patientErr)
	}

	// 2. Lead ativo existente
	var existing models.Lead
	err := s.db.
		Where("LOWER(email) = ? AND status NOT IN (?)", from,
			[]models.LeadStatus{models.LeadStatusConverted, models.LeadStatusLost, models.LeadStatusUnsubscribed}).
		Order("created_at DESC").
		First(&existing).Error

	if err == nil {
		existingID := existing.ID
		_ = s.RecordActivity(RecordActivityInput{
			LeadID:   &existingID,
			Type:     models.LeadActivityMessageReceived,
			Channel:  models.LeadChannelEmail,
			Content:  &bodyContent,
			Metadata: metadata,
		})
		_ = s.db.Model(&existing).Update("last_inbound_at", now).Error
		leadCopy := existing
		snippet := bodyContent
		goSafe("notify_inbound_email", func() {
			s.NotifyTeamOfInboundMessage(ConversationOwner{Lead: &leadCopy}, models.LeadChannelEmail, snippet)
		})
		return &InboundResult{Lead: &existing}, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("lead: inbound email lookup: %w", err)
	}

	// 3. Cria Lead novo (Patient já checado acima, não existe)
	consentVer := "email_inbound_v1" // 16 chars, dentro do varchar(20)
	consentIP := ""                  // inbound não tem IP de cliente

	emailVal := from
	newLead := models.Lead{
		Source:           models.LeadSourceEmailInbound,
		Status:           models.LeadStatusNew,
		Email:            &emailVal,
		Name:             in.FromName,
		Message:          &bodyContent,
		EmailOptIn:       true,
		ConsentVersion:   &consentVer,
		ConsentTimestamp: &now,
		ConsentIPHash:    &consentIP,
		LastInboundAt:    &now,
	}
	if err := s.db.Create(&newLead).Error; err != nil {
		return nil, fmt.Errorf("lead: create from email inbound: %w", err)
	}

	newLeadID := newLead.ID
	_ = s.RecordActivity(RecordActivityInput{
		LeadID:  &newLeadID,
		Type:    models.LeadActivityCreated,
		Channel: models.LeadChannelInternal,
		Content: ptr("Lead criado via primeira mensagem inbound de email"),
	})
	_ = s.RecordActivity(RecordActivityInput{
		LeadID:   &newLeadID,
		Type:     models.LeadActivityMessageReceived,
		Channel:  models.LeadChannelEmail,
		Content:  &bodyContent,
		Metadata: metadata,
	})

	leadCopy := newLead
	goSafe("notify_team", func() { s.NotifyTeamOfNewLead(&leadCopy) })

	return &InboundResult{Lead: &newLead, IsNew: true}, nil
}

// ============================================================
// Manual compose — vendedor inicia contato com email arbitrário
// ============================================================

// CreateForManualCompose cria um Lead novo a partir de um envio manual de email
// pelo vendedor (botão "Novo email" da Central de Conversas). Usa Source=manual,
// Status=new, EmailOptIn=true (consent operacional do vendedor que iniciou),
// AssignedToUserID=actorUserID. Se name não vier, deriva da parte local do email.
//
// Idempotência NÃO é responsabilidade desta função — caller (ConversationService.Compose)
// já fez lookup de Patient + Lead ativo antes de chamar.
func (s *LeadService) CreateForManualCompose(email, name string, actorUserID uuid.UUID) (*models.Lead, error) {
	normalized := strings.ToLower(strings.TrimSpace(email))
	if normalized == "" {
		return nil, errors.New("lead: email obrigatório no manual compose")
	}
	if actorUserID == uuid.Nil {
		return nil, errors.New("lead: actorUserID obrigatório")
	}

	displayName := strings.TrimSpace(name)
	if displayName == "" {
		// Fallback: parte antes do @ do email. Vendedor pode editar depois no /leads.
		if at := strings.Index(normalized, "@"); at > 0 {
			displayName = normalized[:at]
		}
	}

	now := time.Now().UTC()
	consentVer := "manual_compose_v1" // 17 chars, dentro do varchar(20)
	consentIP := ""                   // não há IP de cliente (vendedor que iniciou)
	emailVal := normalized
	actorCopy := actorUserID

	lead := models.Lead{
		Source:           models.LeadSourceManual,
		Status:           models.LeadStatusNew,
		Email:            &emailVal,
		EmailOptIn:       true,
		ConsentVersion:   &consentVer,
		ConsentTimestamp: &now,
		ConsentIPHash:    &consentIP,
		AssignedToUserID: &actorCopy,
	}
	if displayName != "" {
		lead.Name = &displayName
	}

	if err := s.db.Create(&lead).Error; err != nil {
		return nil, fmt.Errorf("lead: create manual compose: %w", err)
	}

	leadID := lead.ID
	_ = s.RecordActivity(RecordActivityInput{
		LeadID:      &leadID,
		Type:        models.LeadActivityCreated,
		Channel:     models.LeadChannelInternal,
		Content:     ptr("Lead criado via composer manual da Central de Conversas"),
		ActorUserID: &actorCopy,
	})

	return &lead, nil
}

// RecordOutboundEmailMirror registra LeadActivity outbound espelhada de "Sent Items".
// Usado quando o usuário responde pelo iPhone/cliente IMAP externo: o worker observa
// a mensagem na pasta enviada e cria activity pra ter histórico no CRM.
//
// Roteamento: Patient first (pelo To), depois Lead. Se nada bate, no-op silencioso
// (pode ser email pra contato externo). Idempotente por (Message-ID, recipient).
func (s *LeadService) RecordOutboundEmailMirror(toEmail string, in InboundEmailInput) error {
	to := strings.ToLower(strings.TrimSpace(toEmail))
	if to == "" {
		return nil
	}

	// Dedup por (message_id, recipient) — chave composta na metadata.
	if in.MessageID != "" {
		var dup int64
		_ = s.db.Model(&models.LeadActivity{}).
			Where("metadata->>'message_id' = ? AND metadata->>'recipient' = ? AND type = ?",
				in.MessageID, to, models.LeadActivityMessageSent).
			Count(&dup).Error
		if dup > 0 {
			return nil
		}
	}

	bodyContent := in.BodyText
	if bodyContent == "" {
		bodyContent = in.BodyHTML
	}

	metadata := map[string]any{
		"message_id":  in.MessageID,
		"recipient":   to,
		"in_reply_to": in.InReplyTo,
		"references":  in.References,
		"subject":     in.Subject,
		"folder":      in.Folder,
		"sent_at":     in.ReceivedAt,
		"mirrored":    true, // veio da pasta Sent (cliente externo), não do EMR
	}
	if len(in.Attachments) > 0 {
		metadata["attachments"] = in.Attachments
	}

	// 1. Patient first
	var patient models.Patient
	if err := s.db.Where("LOWER(email) = ?", to).First(&patient).Error; err == nil {
		patientID := patient.ID
		return s.RecordActivity(RecordActivityInput{
			PatientID: &patientID,
			Type:      models.LeadActivityMessageSent,
			Channel:   models.LeadChannelEmail,
			Content:   &bodyContent,
			Metadata:  metadata,
		})
	}

	// 2. Lead fallback
	var lead models.Lead
	if err := s.db.Where("LOWER(email) = ?", to).Order("created_at DESC").First(&lead).Error; err != nil {
		return nil // sem owner conhecido — no-op
	}
	leadID := lead.ID
	return s.RecordActivity(RecordActivityInput{
		LeadID:   &leadID,
		Type:     models.LeadActivityMessageSent,
		Channel:  models.LeadChannelEmail,
		Content:  &bodyContent,
		Metadata: metadata,
	})
}

// ============================================================
// LeadActivity helpers
// ============================================================

// RecordActivityInput permite registrar activity em Lead OU Patient — exatamente um
// de LeadID/PatientID deve ser fornecido. Espelha a CHECK constraint do schema.
type RecordActivityInput struct {
	LeadID      *uuid.UUID
	PatientID   *uuid.UUID
	Type        models.LeadActivityType
	Channel     models.LeadActivityChannel
	Content     *string
	Metadata    map[string]any
	ActorUserID *uuid.UUID

	// Mídia (WhatsApp Fase 2) — opcionais.
	MediaType         *string
	MediaStorageKey   *string
	MediaMIME         *string
	MediaFilename     *string
	MediaSizeBytes    *int64
	PatientDocumentID *uuid.UUID
	Transcription     *string
}

func (s *LeadService) RecordActivity(in RecordActivityInput) error {
	_, err := s.recordActivityReturning(in)
	return err
}

// recordActivityReturning cria a activity e devolve o ponteiro (precisamos do ID
// pra hooks pós-criação, ex: transcrição de áudio).
func (s *LeadService) recordActivityReturning(in RecordActivityInput) (*models.LeadActivity, error) {
	hasLead := in.LeadID != nil && *in.LeadID != uuid.Nil
	hasPatient := in.PatientID != nil && *in.PatientID != uuid.Nil
	if hasLead == hasPatient {
		return nil, models.ErrLeadActivityOwnerInvalid
	}
	var metaJSON datatypes.JSON
	if len(in.Metadata) > 0 {
		raw, err := json.Marshal(in.Metadata)
		if err != nil {
			return nil, fmt.Errorf("lead activity: metadata marshal: %w", err)
		}
		metaJSON = raw
	}
	activity := models.LeadActivity{
		LeadID:            in.LeadID,
		PatientID:         in.PatientID,
		Type:              in.Type,
		Channel:           in.Channel,
		Content:           in.Content,
		Metadata:          metaJSON,
		ActorUserID:       in.ActorUserID,
		MediaType:         in.MediaType,
		MediaStorageKey:   in.MediaStorageKey,
		MediaMIME:         in.MediaMIME,
		MediaFilename:     in.MediaFilename,
		MediaSizeBytes:    in.MediaSizeBytes,
		PatientDocumentID: in.PatientDocumentID,
		Transcription:     in.Transcription,
	}
	if err := s.db.Create(&activity).Error; err != nil {
		return nil, err
	}
	return &activity, nil
}

// ============================================================
// CRUD admin (autenticado)
// ============================================================

type LeadFilter struct {
	Source           *models.LeadSource
	Status           *models.LeadStatus
	Search           string // busca por name | email | phone
	HasEmailOptIn    *bool
	HasWhatsAppOptIn *bool
	AssignedToUserID *uuid.UUID
	UTMCampaign      *string
}

type LeadListResult struct {
	Items      []models.Lead `json:"items"`
	Total      int64         `json:"total"`
	PageIndex  int           `json:"pageIndex"`
	PageSize   int           `json:"pageSize"`
	TotalPages int           `json:"totalPages"`
}

func (s *LeadService) List(filter LeadFilter, pageIndex, pageSize int) (*LeadListResult, error) {
	if pageSize <= 0 {
		pageSize = 25
	}
	if pageIndex < 0 {
		pageIndex = 0
	}

	q := s.db.Model(&models.Lead{})
	if filter.Source != nil {
		q = q.Where("source = ?", *filter.Source)
	}
	if filter.Status != nil {
		q = q.Where("status = ?", *filter.Status)
	}
	if filter.HasEmailOptIn != nil {
		q = q.Where("email_opt_in = ?", *filter.HasEmailOptIn)
	}
	if filter.HasWhatsAppOptIn != nil {
		q = q.Where("whats_app_opt_in = ?", *filter.HasWhatsAppOptIn)
	}
	if filter.AssignedToUserID != nil {
		q = q.Where("assigned_to_user_id = ?", *filter.AssignedToUserID)
	}
	if filter.UTMCampaign != nil && strings.TrimSpace(*filter.UTMCampaign) != "" {
		q = q.Where("utm_campaign = ?", strings.TrimSpace(*filter.UTMCampaign))
	}
	if s := strings.TrimSpace(filter.Search); s != "" {
		like := "%" + strings.ToLower(s) + "%"
		q = q.Where("LOWER(COALESCE(name,'')) LIKE ? OR LOWER(COALESCE(email,'')) LIKE ? OR COALESCE(phone,'') LIKE ?", like, like, like)
	}

	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, err
	}

	var items []models.Lead
	if err := q.
		Order("created_at DESC").
		Limit(pageSize).
		Offset(pageIndex * pageSize).
		Find(&items).Error; err != nil {
		return nil, err
	}

	totalPages := int((total + int64(pageSize) - 1) / int64(pageSize))
	return &LeadListResult{
		Items:      items,
		Total:      total,
		PageIndex:  pageIndex,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}, nil
}

func (s *LeadService) GetByID(id uuid.UUID) (*models.Lead, error) {
	var lead models.Lead
	err := s.db.
		Preload("Activities", func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at DESC")
		}).
		Preload("Activities.Actor").
		Preload("AssignedTo").
		Preload("ConvertedPatient").
		First(&lead, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &lead, nil
}

type LeadPatch struct {
	Status           *models.LeadStatus
	AssignedToUserID *uuid.UUID
	Name             *string
}

func (s *LeadService) Update(id uuid.UUID, patch LeadPatch, actorUserID uuid.UUID) (*models.Lead, error) {
	var lead models.Lead
	if err := s.db.First(&lead, "id = ?", id).Error; err != nil {
		return nil, err
	}

	updates := map[string]any{"updated_at": time.Now().UTC()}
	statusChanged := false
	if patch.Status != nil && *patch.Status != lead.Status {
		updates["status"] = *patch.Status
		statusChanged = true
	}
	if patch.AssignedToUserID != nil {
		updates["assigned_to_user_id"] = *patch.AssignedToUserID
	}
	if patch.Name != nil {
		updates["name"] = strings.TrimSpace(*patch.Name)
	}

	if err := s.db.Model(&lead).Updates(updates).Error; err != nil {
		return nil, err
	}

	leadIDPtr := id
	if statusChanged {
		_ = s.RecordActivity(RecordActivityInput{
			LeadID:      &leadIDPtr,
			Type:        models.LeadActivityStatusChanged,
			Channel:     models.LeadChannelInternal,
			Content:     ptr(fmt.Sprintf("Status: %s → %s", lead.Status, *patch.Status)),
			ActorUserID: &actorUserID,
		})
	}
	if patch.AssignedToUserID != nil {
		_ = s.RecordActivity(RecordActivityInput{
			LeadID:      &leadIDPtr,
			Type:        models.LeadActivityAssigned,
			Channel:     models.LeadChannelInternal,
			Content:     ptr(fmt.Sprintf("Atribuído ao usuário %s", patch.AssignedToUserID.String())),
			ActorUserID: &actorUserID,
		})
	}

	return s.GetByID(id)
}

func (s *LeadService) AddNote(leadID, actorUserID uuid.UUID, note string) error {
	note = strings.TrimSpace(note)
	if note == "" {
		return errors.New("lead: nota vazia")
	}
	return s.RecordActivity(RecordActivityInput{
		LeadID:      &leadID,
		Type:        models.LeadActivityNoteAdded,
		Channel:     models.LeadChannelInternal,
		Content:     &note,
		ActorUserID: &actorUserID,
	})
}

// MarkConverted marca o Lead como converted, vinculando ao Patient criado.
// Chamado tanto pela conversão automática (claim do Light) quanto pela manual (admin).
func (s *LeadService) MarkConverted(leadID, patientID uuid.UUID, actorUserID *uuid.UUID) error {
	now := time.Now().UTC()
	updates := map[string]any{
		"status":               models.LeadStatusConverted,
		"converted_patient_id": patientID,
		"converted_at":         now,
		"updated_at":           now,
	}
	if actorUserID != nil {
		updates["converted_by_user_id"] = *actorUserID
	}
	if err := s.db.Model(&models.Lead{}).Where("id = ?", leadID).Updates(updates).Error; err != nil {
		return fmt.Errorf("lead: mark converted: %w", err)
	}

	// Repointa o dossiê de relacionamento (memória social da Lívia) do lead para o paciente,
	// para nada se perder na conversão (Fase D / §6 do plano). Best-effort: não falha a conversão.
	if err := repointRelationshipDossier(s.db, leadID, patientID); err != nil {
		log.Printf("⚠️  [DOSSIÊ] repoint lead %s → paciente %s: %v", leadID, patientID, err)
	}

	return s.RecordActivity(RecordActivityInput{
		LeadID:      &leadID,
		Type:        models.LeadActivityConverted,
		Channel:     models.LeadChannelInternal,
		Content:     ptr(fmt.Sprintf("Lead convertido em paciente %s", patientID.String())),
		ActorUserID: actorUserID,
	})
}

// repointRelationshipDossier move profile + facts + people + events do dono lead/<leadID> para
// patient/<patientID>, resolvendo conflitos de unicidade (1 profile por dono; 1 fato ativo por
// key). Só dado social (LGPD/CFM). Idempotente o suficiente para reexecução.
func repointRelationshipDossier(db *gorm.DB, leadID, patientID uuid.UUID) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// 1) Perfil: só move se o paciente ainda não tiver um; senão descarta o do lead.
		if err := tx.Exec(`
			UPDATE relationship_profiles SET owner_type='patient', owner_id=?, updated_at=now()
			WHERE owner_type='lead' AND owner_id=?
			  AND NOT EXISTS (SELECT 1 FROM relationship_profiles WHERE owner_type='patient' AND owner_id=?)`,
			patientID, leadID, patientID).Error; err != nil {
			return err
		}
		if err := tx.Exec(`DELETE FROM relationship_profiles WHERE owner_type='lead' AND owner_id=?`, leadID).Error; err != nil {
			return err
		}

		// 2) Fatos: fecha os ativos do lead cuja key já está ativa no paciente (evita violar o
		// índice único parcial), depois move todo o resto.
		if err := tx.Exec(`
			UPDATE relationship_facts SET valid_until=now()
			WHERE owner_type='lead' AND owner_id=? AND valid_until IS NULL
			  AND key IN (SELECT key FROM relationship_facts WHERE owner_type='patient' AND owner_id=? AND valid_until IS NULL)`,
			leadID, patientID).Error; err != nil {
			return err
		}
		if err := tx.Exec(`UPDATE relationship_facts SET owner_type='patient', owner_id=?, updated_at=now() WHERE owner_type='lead' AND owner_id=?`, patientID, leadID).Error; err != nil {
			return err
		}

		// 3) Pessoas e eventos: sem unicidade, move direto.
		if err := tx.Exec(`UPDATE important_people SET owner_type='patient', owner_id=?, updated_at=now() WHERE owner_type='lead' AND owner_id=?`, patientID, leadID).Error; err != nil {
			return err
		}
		if err := tx.Exec(`UPDATE relationship_events SET owner_type='patient', owner_id=?, updated_at=now() WHERE owner_type='lead' AND owner_id=?`, patientID, leadID).Error; err != nil {
			return err
		}
		return nil
	})
}

// ConvertToPatientInput permite override de campos do Lead na hora de criar Patient.
// Campos não fornecidos caem nos do Lead.
type ConvertToPatientInput struct {
	Name      *string
	Email     *string
	Phone     *string
	BirthDate *time.Time
	Gender    *string // "male" | "female" | "other"
}

// ConvertToPatient cria User+Patient a partir de um Lead e marca o Lead como converted.
// Idempotente: se já está converted, retorna o Patient existente.
//
// Requer name + birthDate + gender (do Lead ou override) — campos NOT NULL no Patient.
// Email é opcional no Patient mas obrigatório no User (gera placeholder se vazio).
func (s *LeadService) ConvertToPatient(leadID, actorUserID uuid.UUID, in ConvertToPatientInput) (*models.Patient, error) {
	lead, err := s.GetByID(leadID)
	if err != nil {
		return nil, err
	}
	if lead.Status == models.LeadStatusConverted && lead.ConvertedPatientID != nil {
		var existing models.Patient
		if err := s.db.First(&existing, "id = ?", *lead.ConvertedPatientID).Error; err != nil {
			return nil, fmt.Errorf("lead: load converted patient: %w", err)
		}
		return &existing, nil
	}

	// Resolve campos finais (override > lead)
	name := strings.TrimSpace(coalesceStr(in.Name, lead.Name))
	if name == "" {
		return nil, errors.New("lead: nome obrigatório para conversão")
	}
	email := strings.ToLower(strings.TrimSpace(coalesceStr(in.Email, lead.Email)))
	phone := strings.TrimSpace(coalesceStr(in.Phone, lead.Phone))
	if email == "" && phone == "" {
		return nil, errors.New("lead: email ou phone obrigatório para conversão")
	}
	gender := "other"
	if in.Gender != nil && *in.Gender != "" {
		gender = *in.Gender
	}
	if in.BirthDate == nil {
		return nil, errors.New("lead: birthDate obrigatório para conversão (data real, não placeholder)")
	}
	birth := *in.BirthDate

	var patient models.Patient
	err = s.db.Transaction(func(tx *gorm.DB) error {
		userEmail := email
		if userEmail == "" {
			// Sem email — gera placeholder determinístico baseado no phone
			userEmail = models.BuildPlaceholderEmail(phone)
		}
		newUser := models.User{Name: name, Email: userEmail}
		if err := newUser.SetRoles([]string{string(models.RolePatient)}); err != nil {
			return err
		}
		if err := tx.Where(models.User{Email: userEmail}).Attrs(newUser).FirstOrCreate(&newUser).Error; err != nil {
			return err
		}

		// Reaproveita Patient se User já tiver
		var existing models.Patient
		if err := tx.Where("user_id = ?", newUser.ID).First(&existing).Error; err == nil {
			patient = existing
			return nil
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		patient = models.Patient{
			UserID:    newUser.ID,
			Name:      name,
			BirthDate: birth,
			Gender:    models.Gender(gender),
			Source:    string(lead.Source),
		}
		if email != "" {
			e := email
			patient.Email = &e
		}
		if phone != "" {
			p := phone
			patient.Phone = &p
		}
		if err := tx.Create(&patient).Error; err != nil {
			return err
		}
		newUser.SelectedPatientID = &patient.ID
		return tx.Save(&newUser).Error
	})
	if err != nil {
		return nil, fmt.Errorf("lead: convert tx: %w", err)
	}

	if err := s.MarkConverted(leadID, patient.ID, &actorUserID); err != nil {
		return nil, err
	}

	// Vincula a sessão da Triagem (se o lead veio do Escore Light) ao paciente recém-criado,
	// para reaproveitar as respostas como prefill na preparação pré-consulta. Best-effort:
	// não derruba a conversão se falhar; só reivindica sessões ainda não reivindicadas.
	if lead.AnonymousScoreSessionID != nil {
		now := time.Now()
		if err := s.db.Model(&models.AnonymousScoreSession{}).
			Where("id = ? AND claimed_by_patient_id IS NULL", *lead.AnonymousScoreSessionID).
			Updates(map[string]interface{}{
				"claimed_by_patient_id": patient.ID,
				"claimed_at":            now,
				"expires_at":            nil,
			}).Error; err != nil {
			log.Printf("⚠️  [CONVERT] vincular sessão %s ao paciente %s: %v", *lead.AnonymousScoreSessionID, patient.ID, err)
		}
	}

	// Email boas-vindas (best-effort, async, só pra emails reais — não placeholders WA)
	if s.emailService != nil && patient.Email != nil && !models.IsPlaceholderEmail(*patient.Email) {
		emailAddr, pname := *patient.Email, patient.Name
		goSafe("send_boas_vindas_manual", func() {
			if err := s.emailService.SendBoasVindas(emailAddr, pname); err != nil {
				log.Printf("⚠️  [BOAS_VINDAS] envio falhou para %s: %v", emailAddr, err)
			}
		})
	}

	return &patient, nil
}

// coalesceStr retorna *override se não-nulo e não-vazio; senão *fallback se não-nulo; senão "".
func coalesceStr(override, fallback *string) string {
	if override != nil && strings.TrimSpace(*override) != "" {
		return *override
	}
	if fallback != nil {
		return *fallback
	}
	return ""
}

// ============================================================
// Stats — Dashboard de funil (Phase 2)
// ============================================================

// StatsPeriod resolve um período do query param ("7d", "30d", "90d", "custom")
// para um intervalo concreto (UTC).
type StatsPeriod struct {
	From  time.Time
	To    time.Time
	Label string
}

func ResolveStatsPeriod(period, fromStr, toStr string) (StatsPeriod, error) {
	now := time.Now().UTC()
	switch period {
	case "", "30d":
		return StatsPeriod{From: now.AddDate(0, 0, -30), To: now, Label: "Últimos 30 dias"}, nil
	case "7d":
		return StatsPeriod{From: now.AddDate(0, 0, -7), To: now, Label: "Últimos 7 dias"}, nil
	case "90d":
		return StatsPeriod{From: now.AddDate(0, 0, -90), To: now, Label: "Últimos 90 dias"}, nil
	case "custom":
		from, err := time.Parse("2006-01-02", fromStr)
		if err != nil {
			return StatsPeriod{}, fmt.Errorf("from inválido (use YYYY-MM-DD): %w", err)
		}
		to, err := time.Parse("2006-01-02", toStr)
		if err != nil {
			return StatsPeriod{}, fmt.Errorf("to inválido (use YYYY-MM-DD): %w", err)
		}
		return StatsPeriod{From: from, To: to.Add(24 * time.Hour), Label: fromStr + " — " + toStr}, nil
	default:
		return StatsPeriod{}, fmt.Errorf("period inválido (use 7d, 30d, 90d ou custom)")
	}
}

type StatsByDay struct {
	Date  string `json:"date"` // YYYY-MM-DD
	Count int64  `json:"count"`
}

type StatsResponse struct {
	Period struct {
		From  time.Time `json:"from"`
		To    time.Time `json:"to"`
		Label string    `json:"label"`
	} `json:"period"`
	Totals           map[string]int64 `json:"totals"`           // por status + "all"
	BySource         map[string]int64 `json:"bySource"`         // por source
	ConversionRate   float64          `json:"conversionRate"`   // converted / all
	AvgTimeToContact *float64         `json:"avgTimeToContact"` // segundos (created → primeira message_sent)
	AvgTimeToConvert *float64         `json:"avgTimeToConvert"` // segundos (created → converted_at)
	ByDay            []StatsByDay     `json:"byDay"`
}

func (s *LeadService) Stats(p StatsPeriod) (*StatsResponse, error) {
	resp := &StatsResponse{
		Totals:   map[string]int64{},
		BySource: map[string]int64{},
		ByDay:    []StatsByDay{},
	}
	resp.Period.From = p.From
	resp.Period.To = p.To
	resp.Period.Label = p.Label

	// Totals por status (+ "all")
	type statusCount struct {
		Status string
		Count  int64
	}
	var statusRows []statusCount
	if err := s.db.Model(&models.Lead{}).
		Select("status, COUNT(*) as count").
		Where("created_at >= ? AND created_at < ?", p.From, p.To).
		Group("status").
		Find(&statusRows).Error; err != nil {
		return nil, fmt.Errorf("stats: status counts: %w", err)
	}
	var allTotal int64
	for _, r := range statusRows {
		resp.Totals[r.Status] = r.Count
		allTotal += r.Count
	}
	resp.Totals["all"] = allTotal

	// Por source
	type sourceCount struct {
		Source string
		Count  int64
	}
	var sourceRows []sourceCount
	if err := s.db.Model(&models.Lead{}).
		Select("source, COUNT(*) as count").
		Where("created_at >= ? AND created_at < ?", p.From, p.To).
		Group("source").
		Find(&sourceRows).Error; err != nil {
		return nil, fmt.Errorf("stats: source counts: %w", err)
	}
	for _, r := range sourceRows {
		resp.BySource[r.Source] = r.Count
	}

	// Conversion rate
	if allTotal > 0 {
		resp.ConversionRate = float64(resp.Totals["converted"]) / float64(allTotal)
	}

	// Avg time to convert (created_at → converted_at)
	type avgRow struct {
		AvgSeconds *float64
	}
	var avgConvert avgRow
	if err := s.db.Raw(`
		SELECT AVG(EXTRACT(EPOCH FROM (converted_at - created_at))) as avg_seconds
		FROM leads
		WHERE created_at >= ? AND created_at < ? AND converted_at IS NOT NULL
	`, p.From, p.To).Scan(&avgConvert).Error; err == nil {
		resp.AvgTimeToConvert = avgConvert.AvgSeconds
	}

	// Avg time to first contact (created_at → MIN(message_sent activity created_at))
	var avgContact avgRow
	if err := s.db.Raw(`
		SELECT AVG(EXTRACT(EPOCH FROM (first_msg.created_at - leads.created_at))) as avg_seconds
		FROM leads
		JOIN LATERAL (
			SELECT created_at FROM lead_activities
			WHERE lead_activities.lead_id = leads.id
			  AND lead_activities.type = 'message_sent'
			ORDER BY created_at ASC
			LIMIT 1
		) first_msg ON TRUE
		WHERE leads.created_at >= ? AND leads.created_at < ?
	`, p.From, p.To).Scan(&avgContact).Error; err == nil {
		resp.AvgTimeToContact = avgContact.AvgSeconds
	}

	// By day (count por dia)
	type byDayRow struct {
		Date  time.Time
		Count int64
	}
	var byDayRows []byDayRow
	if err := s.db.Raw(`
		SELECT DATE_TRUNC('day', created_at) as date, COUNT(*) as count
		FROM leads
		WHERE created_at >= ? AND created_at < ?
		GROUP BY DATE_TRUNC('day', created_at)
		ORDER BY date ASC
	`, p.From, p.To).Scan(&byDayRows).Error; err != nil {
		return nil, fmt.Errorf("stats: by day: %w", err)
	}
	for _, r := range byDayRows {
		resp.ByDay = append(resp.ByDay, StatsByDay{
			Date:  r.Date.Format("2006-01-02"),
			Count: r.Count,
		})
	}

	return resp, nil
}

// IsUnsubscribeKeyword detecta palavras de unsubscribe no texto inbound.
// Lista mínima (PARAR/SAIR/STOP/CANCELAR/DESCADASTRAR), case+acento insensitivo,
// match exato após strip de whitespace e pontuação comum (.,!?).
// Não é substring — "parar de receber" não casa (intencional, evita falso positivo).
func IsUnsubscribeKeyword(text string) bool {
	t := strings.ToLower(strings.TrimSpace(text))
	t = removeAccents(t)
	// Strip pontuação comum no final
	t = strings.TrimRight(t, ".,!?;:")
	t = strings.TrimSpace(t)
	switch t {
	case "parar", "sair", "stop", "cancelar", "descadastrar":
		return true
	}
	return false
}

// removeAccents remove diacríticos comuns BR (á, é, ã, ç…).
// Implementação manual mínima — evita dependência externa.
func removeAccents(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for _, r := range s {
		switch r {
		case 'á', 'à', 'â', 'ã', 'ä':
			b.WriteRune('a')
		case 'é', 'è', 'ê', 'ë':
			b.WriteRune('e')
		case 'í', 'ì', 'î', 'ï':
			b.WriteRune('i')
		case 'ó', 'ò', 'ô', 'õ', 'ö':
			b.WriteRune('o')
		case 'ú', 'ù', 'û', 'ü':
			b.WriteRune('u')
		case 'ç':
			b.WriteRune('c')
		case 'ñ':
			b.WriteRune('n')
		default:
			b.WriteRune(r)
		}
	}
	return b.String()
}

// UnsubscribeByPhone marca todos os Leads ativos do phone como unsubscribed +
// WhatsAppOptIn=false. Registra LeadActivity de unsubscribe SÓ pra leads que de fato
// transicionaram (RowsAffected). Idempotente sob race — segundo chamador concorrente
// vê 0 rows affected e não duplica activities.
func (s *LeadService) UnsubscribeByPhone(phone, originalMessage string) {
	if phone == "" {
		return
	}
	if !strings.HasPrefix(phone, "+") {
		phone = "+" + phone
	}

	now := time.Now().UTC()

	// Carrega IDs dos leads ativos pra esse phone (precisamos pra activities depois)
	var leadIDs []uuid.UUID
	if err := s.db.Model(&models.Lead{}).
		Where("phone = ? AND status NOT IN ?", phone,
			[]models.LeadStatus{models.LeadStatusUnsubscribed, models.LeadStatusConverted}).
		Pluck("id", &leadIDs).Error; err != nil {
		log.Printf("⚠️  [UNSUBSCRIBE] lookup phone=%s falhou: %v", phone, err)
		return
	}
	if len(leadIDs) == 0 {
		return // já estava unsubscribed (ou não existe lead) — no-op
	}

	// Atualização atômica em UM UPDATE com WHERE status NOT IN (...).
	// RowsAffected nos diz quantos leads efetivamente mudaram nesta chamada;
	// chamadas concorrentes que chegarem depois encontram status=unsubscribed
	// e ficam com RowsAffected=0 (não criam activities duplicadas).
	res := s.db.Model(&models.Lead{}).
		Where("id IN ? AND status NOT IN ?", leadIDs,
			[]models.LeadStatus{models.LeadStatusUnsubscribed, models.LeadStatusConverted}).
		Updates(map[string]any{
			"status":           models.LeadStatusUnsubscribed,
			"whats_app_opt_in": false,
			"updated_at":       now,
		})
	if res.Error != nil {
		log.Printf("⚠️  [UNSUBSCRIBE] batch update falhou: %v", res.Error)
		return
	}
	if res.RowsAffected == 0 {
		return // race com outra goroutine que já marcou — silencia
	}

	// Re-busca pra confirmar quais leads de fato foram tocados (race-safe).
	var actuallyUnsubscribed []models.Lead
	if err := s.db.Where("id IN ? AND status = ?", leadIDs, models.LeadStatusUnsubscribed).
		Find(&actuallyUnsubscribed).Error; err != nil {
		log.Printf("⚠️  [UNSUBSCRIBE] refetch falhou: %v", err)
		return
	}

	keywordContent := strings.ToLower(strings.TrimSpace(originalMessage))
	for _, lead := range actuallyUnsubscribed {
		// Activity é registrada uma única vez por lead — checa se já existe
		// activity unsubscribed nos últimos 5 segundos pra esse lead (race window).
		var existingCount int64
		_ = s.db.Model(&models.LeadActivity{}).
			Where("lead_id = ? AND type = ? AND created_at > ?",
				lead.ID, models.LeadActivityUnsubscribed, now.Add(-5*time.Second)).
			Count(&existingCount).Error
		if existingCount > 0 {
			continue
		}
		leadID := lead.ID
		_ = s.RecordActivity(RecordActivityInput{
			LeadID:  &leadID,
			Type:    models.LeadActivityUnsubscribed,
			Channel: models.LeadChannelWhatsApp,
			Content: &originalMessage,
			Metadata: map[string]any{
				"trigger":   "auto_keyword",
				"keyword":   keywordContent,
				"timestamp": now,
			},
		})
		log.Printf("📱 [UNSUBSCRIBE] lead=%s phone=%s marcado como unsubscribed", lead.ID, phone)
	}
}

// RecordWhatsAppStatus persiste um status update (delivered/read/failed) recebido
// via webhook Meta. Faz lookup do LeadActivity outbound original via wa_message_id.
// Best-effort: se não acha (mensagem antiga ou de outro sistema), só loga e ignora.
func (s *LeadService) RecordWhatsAppStatus(waMessageID, status, recipientID string, ts time.Time) {
	if waMessageID == "" {
		return
	}
	var orig models.LeadActivity
	// JSONB: metadata->>'wa_message_id' = ? — restringe à mensagem ENVIADA. Sem o filtro de
	// type, os próprios eventos de status (que também guardam wa_message_id) eram encontrados,
	// fazendo o "delivered/read" apontar para o evento "sent" em vez da mensagem original.
	err := s.db.Where("metadata->>'wa_message_id' = ?", waMessageID).
		Where("type = ?", models.LeadActivityMessageSent).
		Order("created_at DESC").
		First(&orig).Error
	if err != nil {
		// Não acha: pode ser msg antiga ou template enviado fora do CRM
		log.Printf("📱 [WHATSAPP STATUS] wa_message_id=%s sem activity outbound (status=%s)", waMessageID, status)
		return
	}

	metaJSON, _ := json.Marshal(map[string]any{
		"wa_message_id":     waMessageID,
		"status":            status,
		"recipient":         recipientID,
		"timestamp":         ts,
		"original_activity": orig.ID,
	})
	// Propaga ownership da activity original — pode ser Lead OU Patient (Bloco A).
	activity := &models.LeadActivity{
		LeadID:    orig.LeadID,
		PatientID: orig.PatientID,
		Type:      models.LeadActivityMessageStatusChanged,
		Channel:   models.LeadChannelWhatsApp,
		Metadata:  metaJSON,
	}
	if err := s.db.Create(activity).Error; err != nil {
		log.Printf("⚠️  [WHATSAPP STATUS] persist falhou: %v", err)
	}
}

// SendWhatsAppText envia mensagem free-form (session message) via WhatsApp.
// Valida: lead tem phone, opt-in, e janela 24h ainda aberta.
// Persiste LeadActivity (message_sent, channel=whatsapp) com wa_message_id em metadata.
func (s *LeadService) SendWhatsAppText(leadID, actorUserID uuid.UUID, text string) (*models.LeadActivity, error) {
	text = strings.TrimSpace(text)
	if text == "" {
		return nil, errors.New("lead: mensagem vazia")
	}
	if len(text) > 4096 {
		return nil, errors.New("lead: mensagem excede 4096 caracteres")
	}

	lead, err := s.GetByID(leadID)
	if err != nil {
		return nil, err
	}
	if lead.Phone == nil || *lead.Phone == "" {
		return nil, errors.New("lead: sem telefone cadastrado")
	}
	if !lead.WhatsAppOptIn {
		return nil, errors.New("lead: sem opt-in de WhatsApp")
	}
	if lead.Status == models.LeadStatusUnsubscribed {
		return nil, errors.New("lead: cliente solicitou unsubscribe")
	}
	if lead.LastInboundAt == nil {
		return nil, errors.New("lead: sem mensagem inbound recente — fora da janela de 24h. Use template aprovado")
	}
	if time.Since(*lead.LastInboundAt) > 24*time.Hour {
		return nil, errors.New("lead: janela de 24h expirou — use template aprovado")
	}

	if s.whatsappService == nil {
		return nil, errors.New("lead: WhatsApp service não configurado")
	}
	wamid, err := s.whatsappService.SendTextMessage(*lead.Phone, text)
	if err != nil {
		return nil, fmt.Errorf("lead: envio WA falhou: %w", err)
	}

	activity := &models.LeadActivity{
		LeadID:      &leadID,
		Type:        models.LeadActivityMessageSent,
		Channel:     models.LeadChannelWhatsApp,
		Content:     &text,
		ActorUserID: &actorUserID,
	}
	metaJSON, _ := json.Marshal(map[string]any{
		"wa_message_id": wamid,
		"sent_at":       time.Now().UTC(),
	})
	activity.Metadata = metaJSON
	if err := s.db.Create(activity).Error; err != nil {
		return nil, fmt.Errorf("lead: persist activity: %w", err)
	}
	return activity, nil
}

// Delete remove um Lead (soft delete via gorm.DeletedAt). Atividades caem em cascade.
// Usado pelo direito LGPD de eliminação (art. 18 VI) e por housekeeping admin.
func (s *LeadService) Delete(id uuid.UUID) error {
	res := s.db.Delete(&models.Lead{}, "id = ?", id)
	if res.Error != nil {
		return fmt.Errorf("lead: delete: %w", res.Error)
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// ============================================================
// Notificação interna do CRM (Phase 2)
// ============================================================

// NotifyTeamOfNewLead dispara notificações pra equipe de vendas:
// - In-app (badge no sino) — sempre, pra todos os recipients ativos
// - WhatsApp Business (template lead_alert) — best-effort, só pra users com ProfessionalPhone
//
// Recipients são resolvidos por:
//  1. cfg.CRM.LeadNotifyUserIDs (csv hardcoded) se preenchido
//  2. caso contrário, todos users com role admin/secretary/manager
//
// Erros não bloqueiam — só loga. Idempotente: chamar N vezes pra mesmo lead pode duplicar
// notif (sem sentinel anti-dup ainda). Caller deve chamar 1×.
func (s *LeadService) NotifyTeamOfNewLead(lead *models.Lead) {
	if lead == nil || s.notificationService == nil {
		return
	}

	recipients, err := s.resolveLeadNotifyRecipients()
	if err != nil {
		log.Printf("⚠️  [LEAD NOTIFY] resolve recipients: %v", err)
		return
	}
	if len(recipients) == 0 {
		log.Printf("⚠️  [LEAD NOTIFY] nenhum recipient configurado (set CRM_LEAD_NOTIFY_USER_IDS ou crie users admin)")
		return
	}

	contact := ""
	if lead.Email != nil {
		contact = *lead.Email
	} else if lead.Phone != nil {
		contact = *lead.Phone
	}
	leadName := "(sem nome)"
	if lead.Name != nil && *lead.Name != "" {
		leadName = *lead.Name
	}

	sourceLabel := leadSourceLabel(lead.Source)
	title := fmt.Sprintf("Novo lead: %s", leadName)
	message := fmt.Sprintf("Origem: %s · %s", sourceLabel, contact)
	notifType := models.NotificationLeadNew
	if lead.Source == models.LeadSourceWhatsAppInbound {
		notifType = models.NotificationLeadWhatsAppInbound
	} else if lead.Source == models.LeadSourceEmailInbound {
		notifType = models.NotificationLeadEmailInbound
	}

	adminBase := strings.TrimRight(s.cfg.CRM.AdminURL, "/")
	// Lead vindo do WhatsApp abre direto na conversa de WA (não na ficha do lead) — mesma
	// regra de NotifyTeamOfInboundMessage. Bug reportado em 2026-06-08.
	leadURL := fmt.Sprintf("%s/leads/%s", adminBase, lead.ID.String())
	if lead.Source == models.LeadSourceWhatsAppInbound {
		leadURL = fmt.Sprintf("%s/conversas/whatsapp?owner=lead&id=%s", adminBase, lead.ID.String())
	}

	for _, user := range recipients {
		// 1. In-app
		if err := s.notificationService.CreateLeadNotification(
			user.ID, lead.ID, notifType, title, message, leadURL,
		); err != nil {
			log.Printf("⚠️  [LEAD NOTIFY] in-app pra %s falhou: %v", user.ID, err)
		}

		// 2. WhatsApp interno (best-effort, só se user tem ProfessionalPhone E não opt-out)
		if s.whatsappService != nil && user.ProfessionalPhone != nil && *user.ProfessionalPhone != "" &&
			userWantsWhatsAppNotification(user) {
			if err := s.whatsappService.SendLeadAlert(
				*user.ProfessionalPhone, leadName, contact, sourceLabel, leadURL,
			); err != nil {
				log.Printf("⚠️  [LEAD NOTIFY] WhatsApp pra %s falhou: %v", user.ID, err)
			}
		}
	}
}

// ConversationOwner identifica o dono de uma conversa (Lead OU Patient — mutuamente
// exclusivo). Usado por NotifyTeamOfInboundMessage pra decidir título/URL do notif.
type ConversationOwner struct {
	Lead    *models.Lead
	Patient *models.Patient
}

// NotifyTeamOfInboundMessage dispara notificação in-app quando uma conversa JÁ EXISTENTE
// (Lead ou Patient) recebe nova mensagem inbound. NotifyTeamOfNewLead trata só leads novos.
//
// owner: exatamente um de Lead|Patient setado. Patient ganha sufixo "(paciente)" no título
// e URL aponta pra /conversas/patient/:id (Bloco C). Lead → /leads/:id (admin legacy)
// ou /conversas/lead/:id se preferido — adminURL define o prefixo.
//
// channel: models.LeadChannelEmail | models.LeadChannelWhatsApp
// snippet: corpo da mensagem (truncado pra ~80 chars). Sem WhatsApp template — só sino.
//
// Pós-Bloco B: usa NotificationService.CreateConversationNotification que aceita owner
// Lead OU Patient nativamente (notifications.patient_id setado pra owner=Patient).
func (s *LeadService) NotifyTeamOfInboundMessage(owner ConversationOwner, channel models.LeadActivityChannel, snippet string) {
	if s.notificationService == nil {
		return
	}
	hasLead := owner.Lead != nil
	hasPatient := owner.Patient != nil
	if hasLead == hasPatient {
		return // 0 ou 2 owners — invariante violada, abort silencioso
	}
	recipients, err := s.resolveLeadNotifyRecipients()
	if err != nil || len(recipients) == 0 {
		return
	}

	if len(snippet) > 80 {
		snippet = snippet[:77] + "..."
	}
	channelLabel := "Email"
	notifType := models.NotificationLeadEmailInbound
	if channel == models.LeadChannelWhatsApp {
		channelLabel = "WhatsApp"
		notifType = models.NotificationLeadWhatsAppInbound
	}

	adminBase := strings.TrimRight(s.cfg.CRM.AdminURL, "/")

	var (
		ownerName string
		fallback  string
		notifURL  string
		leadID    *uuid.UUID
		patientID *uuid.UUID
	)

	// Inbound de WhatsApp abre direto na tela de WhatsApp (inbox de chat), com a conversa já
	// selecionada via query param; e-mail mantém a rota antiga. Antes o lead caía em /leads/:id
	// (página de cadastro do lead, não a conversa) — bug reportado em 2026-06-08.
	isWhatsApp := channel == models.LeadChannelWhatsApp

	if hasLead {
		lead := owner.Lead
		ownerName = "(sem nome)"
		if lead.Name != nil && *lead.Name != "" {
			ownerName = *lead.Name
		}
		if lead.Email != nil {
			fallback = *lead.Email
		} else if lead.Phone != nil {
			fallback = *lead.Phone
		}
		base := adminBase
		if base == "" {
			base = ""
		}
		if isWhatsApp {
			notifURL = fmt.Sprintf("%s/conversas/whatsapp?owner=lead&id=%s", base, lead.ID.String())
		} else {
			notifURL = fmt.Sprintf("%s/leads/%s", base, lead.ID.String())
		}
		id := lead.ID
		leadID = &id
	} else {
		patient := owner.Patient
		ownerName = patient.Name
		if ownerName == "" {
			ownerName = "(sem nome)"
		}
		ownerName += " (paciente)"
		if patient.Email != nil {
			fallback = *patient.Email
		} else if patient.Phone != nil {
			fallback = *patient.Phone
		}
		base := adminBase
		if isWhatsApp {
			notifURL = fmt.Sprintf("%s/conversas/whatsapp?owner=patient&id=%s", base, patient.ID.String())
		} else {
			// Rota oficial pós-Bloco C: /conversas/patient/:id
			notifURL = fmt.Sprintf("%s/conversas/patient/%s", base, patient.ID.String())
		}
		id := patient.ID
		patientID = &id
	}

	title := fmt.Sprintf("%s de %s", channelLabel, ownerName)
	message := snippet
	if message == "" {
		message = fallback
	}

	for _, user := range recipients {
		if err := s.notificationService.CreateConversationNotification(
			user.ID, leadID, patientID, notifType, title, message, notifURL,
		); err != nil {
			log.Printf("⚠️  [INBOUND NOTIFY] in-app pra %s falhou: %v", user.ID, err)
		}
	}
}

// resolveLeadNotifyRecipients retorna a lista de users que devem receber notificação de lead.
//
// Resolução em 2 etapas:
//  1. Lista candidatos:
//     - se CRM_LEAD_NOTIFY_USER_IDS preenchido, usa lista hardcoded
//     - senão, fallback pra todos com role admin/secretary/manager
//  2. Filtra por opt-out individual:
//     - se Preferences.leadNotifications.inApp == false, usuário é EXCLUÍDO
//     - default (sem preferência setada) é receber (opt-out, não opt-in)
func (s *LeadService) resolveLeadNotifyRecipients() ([]models.User, error) {
	var candidates []models.User

	// 1. Lista candidatos
	if len(s.cfg.CRM.LeadNotifyUserIDs) > 0 {
		ids := s.cfg.CRM.LeadNotifyUserIDs
		if err := s.db.Where("id IN ?", ids).Find(&candidates).Error; err != nil {
			return nil, err
		}
	} else {
		if err := s.db.Where(
			`(roles @> ?::jsonb OR roles @> ?::jsonb OR roles @> ?::jsonb) AND deleted_at IS NULL`,
			`["admin"]`, `["secretary"]`, `["manager"]`,
		).Find(&candidates).Error; err != nil {
			return nil, err
		}
	}

	// 2. Filtra por opt-out individual via Preferences.leadNotifications.inApp
	out := make([]models.User, 0, len(candidates))
	for _, u := range candidates {
		if !userWantsLeadNotifications(u) {
			continue
		}
		out = append(out, u)
	}
	return out, nil
}

// userWantsLeadNotifications lê User.Preferences.leadNotifications.inApp.
// Default true se chave inexistente — opt-OUT explícito é necessário pra parar.
// Shape esperado: {"leadNotifications":{"inApp":false,"whatsapp":false}}
func userWantsLeadNotifications(u models.User) bool {
	if len(u.Preferences) == 0 {
		return true
	}
	var prefs struct {
		LeadNotifications *struct {
			InApp *bool `json:"inApp,omitempty"`
		} `json:"leadNotifications,omitempty"`
	}
	if err := json.Unmarshal(u.Preferences, &prefs); err != nil {
		return true // default permissivo se parse falha
	}
	if prefs.LeadNotifications == nil || prefs.LeadNotifications.InApp == nil {
		return true
	}
	return *prefs.LeadNotifications.InApp
}

// userWantsWhatsAppNotification segue a mesma lógica pra canal WhatsApp.
func userWantsWhatsAppNotification(u models.User) bool {
	if len(u.Preferences) == 0 {
		return true
	}
	var prefs struct {
		LeadNotifications *struct {
			WhatsApp *bool `json:"whatsapp,omitempty"`
		} `json:"leadNotifications,omitempty"`
	}
	if err := json.Unmarshal(u.Preferences, &prefs); err != nil {
		return true
	}
	if prefs.LeadNotifications == nil || prefs.LeadNotifications.WhatsApp == nil {
		return true
	}
	return *prefs.LeadNotifications.WhatsApp
}

// leadSourceLabel retorna texto humano pro source (PT).
func leadSourceLabel(src models.LeadSource) string {
	switch src {
	case models.LeadSourceLightClaim:
		return "Escore Light"
	case models.LeadSourceContactForm:
		return "Formulário /contato"
	case models.LeadSourceWhatsAppInbound:
		return "WhatsApp inbound"
	case models.LeadSourceEmailInbound:
		return "Email inbound"
	case models.LeadSourceNewsletter:
		return "Newsletter"
	case models.LeadSourceManual:
		return "Manual"
	default:
		return string(src)
	}
}

// FindLeadBySession retorna o Lead vinculado a uma AnonymousScoreSession (se existir).
func (s *LeadService) FindLeadBySession(sessionID uuid.UUID) (*models.Lead, error) {
	var lead models.Lead
	err := s.db.Where("anonymous_score_session_id = ?", sessionID).First(&lead).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &lead, nil
}

// ============================================================
// Helpers
// ============================================================

func normalizeEmailPtr(p *string) *string {
	if p == nil {
		return nil
	}
	v := strings.ToLower(strings.TrimSpace(*p))
	if v == "" {
		return nil
	}
	return &v
}

func ptr[T any](v T) *T { return &v }

// goSafe roda fn em goroutine com defer recover() — panics em background
// não derrubam o processo. Logam stack trace pra investigação.
func goSafe(label string, fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("🛑 [PANIC %s] recovered: %v", label, r)
			}
		}()
		fn()
	}()
}

// isUniqueViolation detecta unique constraint violation do PostgreSQL.
// Verifica códigos SQLState 23505 e mensagens conhecidas (sem dependência do pgconn).
func isUniqueViolation(err error) bool {
	if err == nil {
		return false
	}
	msg := err.Error()
	return strings.Contains(msg, "23505") ||
		strings.Contains(msg, "duplicate key") ||
		strings.Contains(msg, "unique constraint")
}
