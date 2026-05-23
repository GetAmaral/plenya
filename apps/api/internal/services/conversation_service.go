package services

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/crypto"
	"github.com/plenya/api/internal/models"
)

// ============================================================
// ConversationService — Bloco B da Central de Conversas
// ============================================================
//
// Endpoints:
//   GET  /conversations                     → List
//   GET  /conversations/:type/:id/messages  → Messages (+ marca lido)
//   POST /conversations/:type:/:id/read     → MarkRead
//   POST /conversations/:type/:id/email     → SendMessage(channel=email)
//   POST /conversations/:type/:id/whatsapp  → SendMessage(channel=whatsapp)
//
// Conversa = sequência de LeadActivity ligadas ao mesmo owner (Lead OU Patient).
// Bloco A já garantiu schema dual; aqui consolidamos a leitura.

// ConversationService é o serviço da Central de Conversas. Reusa LeadService,
// EmailService, WhatsAppService e NotificationService — não duplica lógica.
type ConversationService struct {
	db                  *gorm.DB
	leadService         *LeadService
	emailService        *EmailService
	whatsappService     *WhatsAppService
	notificationService *NotificationService
	aiService           *AIService

	// Serviços de mídia (WhatsApp Fase 2), injetados via SetMediaServices.
	waMedia     *WhatsAppMediaService
	patientDocs *PatientDocumentsService
	uploadsRoot string // raiz do storage local (/app/uploads), pra resolver anexos de e-mail

	// Interpretador de exames (2.0b), injetado via SetExamInterpreter.
	labBatches     *LabResultBatchService
	processingJobs *ProcessingJobService
}

// SetMediaServices injeta os serviços de mídia (WhatsApp Fase 2) após DI em main.go.
func (s *ConversationService) SetMediaServices(waMedia *WhatsAppMediaService, patientDocs *PatientDocumentsService) {
	s.waMedia = waMedia
	s.patientDocs = patientDocs
	if patientDocs != nil {
		s.uploadsRoot = patientDocs.uploadsRoot
	}
}

// activityAttachment espelha o shape persistido em metadata.attachments
// (e-mail inbound/outbound): { filename, path, content_type, size_bytes }.
type activityAttachment struct {
	Filename    string `json:"filename"`
	Path        string `json:"path"`
	ContentType string `json:"content_type"`
	SizeBytes   int    `json:"size_bytes"`
}

// loadActivityForOwner carrega a atividade e valida que pertence ao dono (type/id).
func (s *ConversationService) loadActivityForOwner(ownerType string, ownerID, activityID uuid.UUID) (*models.LeadActivity, error) {
	var act models.LeadActivity
	if err := s.db.First(&act, "id = ?", activityID).Error; err != nil {
		return nil, ErrConversationOwnerInvalid
	}
	switch ownerType {
	case "lead":
		if act.LeadID == nil || *act.LeadID != ownerID {
			return nil, ErrConversationOwnerInvalid
		}
	case "patient":
		if act.PatientID == nil || *act.PatientID != ownerID {
			return nil, ErrConversationOwnerInvalid
		}
	default:
		return nil, ErrConversationOwnerInvalid
	}
	return &act, nil
}

// activityAttachments extrai os anexos (e-mail) do metadata da atividade.
func activityAttachments(act *models.LeadActivity) []activityAttachment {
	if len(act.Metadata) == 0 {
		return nil
	}
	var m struct {
		Attachments []activityAttachment `json:"attachments"`
	}
	_ = json.Unmarshal(act.Metadata, &m)
	return m.Attachments
}

// resolveUploadPath resolve um path relativo a uploadsRoot com defesa anti
// path-traversal (mesma proteção do PatientDocumentsService).
func (s *ConversationService) resolveUploadPath(rel string) (string, error) {
	clean := strings.TrimPrefix(strings.TrimLeft(rel, "/"), "uploads/")
	full := filepath.Join(s.uploadsRoot, clean)
	abs, _ := filepath.Abs(full)
	rootAbs, _ := filepath.Abs(s.uploadsRoot)
	if !strings.HasPrefix(abs, rootAbs+string(os.PathSeparator)) {
		return "", ErrConversationAttachmentInvalid
	}
	if _, err := os.Stat(full); err != nil {
		return "", ErrConversationAttachmentInvalid
	}
	return full, nil
}

// GetAttachmentFile resolve um anexo de e-mail (por índice) pra download autenticado.
func (s *ConversationService) GetAttachmentFile(ownerType string, ownerID, activityID uuid.UUID, idx int) (*ActivityMediaResult, error) {
	act, err := s.loadActivityForOwner(ownerType, ownerID, activityID)
	if err != nil {
		return nil, err
	}
	atts := activityAttachments(act)
	if idx < 0 || idx >= len(atts) {
		return nil, ErrConversationAttachmentInvalid
	}
	att := atts[idx]
	full, err := s.resolveUploadPath(att.Path)
	if err != nil {
		return nil, err
	}
	return &ActivityMediaResult{FilePath: full, MIME: att.ContentType, Filename: att.Filename}, nil
}

// SaveActivityMediaToProntuario salva uma mídia/anexo (e-mail ou WhatsApp) de uma
// conversa de PACIENTE nas mídias do prontuário (PatientDocument). Idempotente.
//
//   - attachmentIndex != nil → anexo de e-mail (metadata.attachments[idx]).
//   - attachmentIndex == nil → mídia WhatsApp da própria atividade (colunas media_*).
//
// Só PDF/JPG/PNG entram no prontuário (whitelist clínica do PatientDocumentsService);
// áudio e outros tipos são rejeitados.
func (s *ConversationService) SaveActivityMediaToProntuario(ownerID, activityID uuid.UUID, attachmentIndex *int) (*models.PatientDocument, error) {
	if s.patientDocs == nil {
		return nil, errors.New("conversation: serviço de documentos não configurado")
	}
	act, err := s.loadActivityForOwner("patient", ownerID, activityID)
	if err != nil {
		return nil, err
	}

	// Anexo de e-mail.
	if attachmentIndex != nil {
		atts := activityAttachments(act)
		if *attachmentIndex < 0 || *attachmentIndex >= len(atts) {
			return nil, ErrConversationAttachmentInvalid
		}
		att := atts[*attachmentIndex]
		full, err := s.resolveUploadPath(att.Path)
		if err != nil {
			return nil, err
		}
		data, err := os.ReadFile(full)
		if err != nil {
			return nil, ErrConversationAttachmentInvalid
		}
		origin := fmt.Sprintf("act:%s:att:%d", activityID, *attachmentIndex)
		return s.patientDocs.CreateFromBytes(CreateFromBytesInput{
			PatientID:         ownerID,
			Bytes:             data,
			MIME:              att.ContentType,
			Filename:          att.Filename,
			Title:             att.Filename,
			Type:              models.DocumentTypeOther,
			Source:            models.DocumentSourceEmail,
			OriginWAMessageID: &origin,
		})
	}

	// Mídia WhatsApp da atividade.
	if act.PatientDocumentID != nil {
		// Já está no prontuário — devolve o existente.
		var doc models.PatientDocument
		if err := s.db.First(&doc, "id = ?", *act.PatientDocumentID).Error; err != nil {
			return nil, ErrConversationAttachmentInvalid
		}
		return &doc, nil
	}
	if act.MediaStorageKey == nil || s.waMedia == nil {
		return nil, errors.New("conversation: mensagem sem mídia pra salvar no prontuário")
	}
	data, err := s.waMedia.OpenDecrypted(*act.MediaStorageKey)
	if err != nil {
		return nil, ErrConversationAttachmentInvalid
	}
	mime := ""
	if act.MediaMIME != nil {
		mime = *act.MediaMIME
	}
	filename := "arquivo"
	if act.MediaFilename != nil {
		filename = *act.MediaFilename
	}
	origin := fmt.Sprintf("act:%s:media", activityID)
	doc, err := s.patientDocs.CreateFromBytes(CreateFromBytesInput{
		PatientID:         ownerID,
		Bytes:             data,
		MIME:              mime,
		Filename:          filename,
		Title:             filename,
		Type:              models.DocumentTypeOther,
		Source:            models.DocumentSourceWhatsApp,
		OriginWAMessageID: &origin,
	})
	if err != nil {
		return nil, err
	}
	// Linka a atividade ao documento criado (UI mostra "salvo no prontuário").
	_ = s.db.Model(&models.LeadActivity{}).Where("id = ?", activityID).
		Update("patient_document_id", doc.ID).Error
	return doc, nil
}

// SetExamInterpreter injeta o pipeline de interpretação de exames (2.0b).
func (s *ConversationService) SetExamInterpreter(batches *LabResultBatchService, jobs *ProcessingJobService) {
	s.labBatches = batches
	s.processingJobs = jobs
}

// InterpretExamResult informa o documento/batch/job criados pela interpretação.
type InterpretExamResult struct {
	DocumentID uuid.UUID `json:"documentId"`
	BatchID    uuid.UUID `json:"batchId"`
	JobID      uuid.UUID `json:"jobId"`
}

// InterpretActivityAsExam salva a mídia da mensagem no prontuário (se necessário) e
// a submete ao interpretador de exames existente: cria um LabResultBatch e enfileira
// um ProcessingJob (OCR → IA → resultados classificados). Unifica e-mail e WhatsApp.
//
//   - attachmentIndex != nil → anexo de e-mail; == nil → mídia WhatsApp da atividade.
//
// Sob demanda (decisão de produto): só roda quando a equipe pede, evitando criar
// batches-lixo de arquivos que não são exames.
func (s *ConversationService) InterpretActivityAsExam(ownerType string, ownerID, activityID uuid.UUID, attachmentIndex *int) (*InterpretExamResult, error) {
	if s.processingJobs == nil || s.patientDocs == nil {
		return nil, errors.New("conversation: interpretador de exames não configurado")
	}
	// Só paciente: lead não tem prontuário nem exames.
	if ownerType != "patient" {
		return nil, ErrConversationOwnerInvalid
	}

	// Garante o documento no prontuário (salva se ainda não estiver) — idempotente.
	doc, err := s.SaveActivityMediaToProntuario(ownerID, activityID, attachmentIndex)
	if err != nil {
		return nil, err
	}
	_, fullPath, err := s.patientDocs.GetForDownload(ownerID, doc.ID)
	if err != nil {
		return nil, ErrConversationAttachmentInvalid
	}
	// OCR cobre PDF (pdftotext) e imagem escaneada (Tesseract); rejeita outros.
	if doc.ContentType != "application/pdf" && doc.ContentType != "image/jpeg" && doc.ContentType != "image/png" {
		return nil, fmt.Errorf("conversation: tipo %q não interpretável como exame", doc.ContentType)
	}

	batch := &models.LabResultBatch{
		PatientID:      ownerID,
		LaboratoryName: "Recebido por WhatsApp",
		CollectionDate: time.Now().UTC(),
		Status:         models.LabResultBatchPending,
	}
	if err := s.db.Create(batch).Error; err != nil {
		return nil, fmt.Errorf("conversation: criar batch de exame: %w", err)
	}

	job, err := s.processingJobs.Create(batch.ID, fullPath)
	if err != nil {
		return nil, fmt.Errorf("conversation: enfileirar interpretação: %w", err)
	}
	return &InterpretExamResult{DocumentID: doc.ID, BatchID: batch.ID, JobID: job.ID}, nil
}

// ActivityMediaResult é o resultado de resolver a mídia de uma atividade pra download.
// Exatamente um de FilePath (arquivo de prontuário, ler do disco) ou Bytes (mídia
// cifrada já decifrada) é preenchido.
type ActivityMediaResult struct {
	FilePath string
	Bytes    []byte
	MIME     string
	Filename string
}

// GetActivityMedia resolve a mídia de uma LeadActivity após validar que ela
// pertence ao dono (type/id). Roteia entre arquivo de prontuário e bucket cifrado.
func (s *ConversationService) GetActivityMedia(ownerType string, ownerID, activityID uuid.UUID) (*ActivityMediaResult, error) {
	var act models.LeadActivity
	if err := s.db.First(&act, "id = ?", activityID).Error; err != nil {
		return nil, ErrConversationOwnerInvalid
	}
	switch ownerType {
	case "lead":
		if act.LeadID == nil || *act.LeadID != ownerID {
			return nil, ErrConversationOwnerInvalid
		}
	case "patient":
		if act.PatientID == nil || *act.PatientID != ownerID {
			return nil, ErrConversationOwnerInvalid
		}
	default:
		return nil, ErrConversationOwnerInvalid
	}

	mime := ""
	if act.MediaMIME != nil {
		mime = *act.MediaMIME
	}
	filename := ""
	if act.MediaFilename != nil {
		filename = *act.MediaFilename
	}

	// Arquivo no prontuário (paciente).
	if act.PatientDocumentID != nil {
		if s.patientDocs == nil || act.PatientID == nil {
			return nil, ErrConversationAttachmentInvalid
		}
		doc, full, err := s.patientDocs.GetForDownload(*act.PatientID, *act.PatientDocumentID)
		if err != nil {
			return nil, ErrConversationAttachmentInvalid
		}
		if mime == "" {
			mime = doc.ContentType
		}
		if filename == "" {
			filename = doc.FileName
		}
		return &ActivityMediaResult{FilePath: full, MIME: mime, Filename: filename}, nil
	}

	// Mídia no bucket cifrado da conversa.
	if act.MediaStorageKey != nil {
		if s.waMedia == nil {
			return nil, ErrConversationAttachmentInvalid
		}
		data, err := s.waMedia.OpenDecrypted(*act.MediaStorageKey)
		if err != nil {
			return nil, ErrConversationAttachmentInvalid
		}
		return &ActivityMediaResult{Bytes: data, MIME: mime, Filename: filename}, nil
	}

	return nil, ErrConversationAttachmentInvalid
}

// NewConversationService monta o serviço com DI.
func NewConversationService(
	db *gorm.DB,
	lead *LeadService,
	email *EmailService,
	wa *WhatsAppService,
	notif *NotificationService,
	ai *AIService,
) *ConversationService {
	return &ConversationService{
		db:                  db,
		leadService:         lead,
		emailService:        email,
		whatsappService:     wa,
		notificationService: notif,
		aiService:           ai,
	}
}

// ============================================================
// Errors públicos (handlers mapeiam pra HTTP semântico)
// ============================================================

var (
	// ErrConversationOwnerInvalid → owner type fora do enum ou ID inválido (404).
	ErrConversationOwnerInvalid = errors.New("conversation: owner inválido")

	// ErrConversationNoChannel → owner sem email/phone (422).
	ErrConversationNoChannel = errors.New("conversation: owner sem canal disponível")

	// ErrConversationOptedOut → Lead.EmailOptIn=false ou WhatsAppOptIn=false (422).
	ErrConversationOptedOut = errors.New("conversation: opt-out de canal")

	// ErrConversationWindowClosed → WhatsApp fora janela 24h (422).
	ErrConversationWindowClosed = errors.New("conversation: janela de 24h expirou — use template")

	// ErrConversationAttachmentTooLarge → soma de anexos excede limite Resend (422).
	ErrConversationAttachmentTooLarge = errors.New("conversation: anexos excedem 40MB")

	// ErrConversationAttachmentInvalid → path traversal, arquivo não encontrado, ou tipo inválido (422).
	ErrConversationAttachmentInvalid = errors.New("conversation: anexo inválido")
)

// ============================================================
// List — lista de conversas (UNION Lead + Patient)
// ============================================================

// ConversationItem é o resumo de cada linha da lista esquerda do /conversas.
type ConversationItem struct {
	OwnerType        string     `json:"ownerType"` // "lead" | "patient"
	OwnerID          uuid.UUID  `json:"ownerId"`
	Name             string     `json:"name"`
	Email            *string    `json:"email,omitempty"`
	Phone            *string    `json:"phone,omitempty"`
	LastChannel      string     `json:"lastChannel"`   // "email" | "whatsapp" | "internal"
	LastDirection    string     `json:"lastDirection"` // "in" | "out"
	LastSnippet      string     `json:"lastSnippet"`   // primeiros 120 chars
	LastAt           time.Time  `json:"lastAt"`
	UnreadCount      int        `json:"unreadCount"`                // activities depois de last_read_at
	AssignedToUserID *uuid.UUID `json:"assignedToUserId,omitempty"` // só Lead
	Channels         []string   `json:"channels"`                   // ["email","whatsapp"]
	LeadStatus       *string    `json:"leadStatus,omitempty"`       // só Lead
	// Campos pra UI desabilitar composer (consistência client-side com regras do backend).
	// Pra Patient: ambos opt-in default true; LastInboundAt nil (sem janela 24h).
	EmailOptIn    bool       `json:"emailOptIn"`
	WhatsAppOptIn bool       `json:"whatsAppOptIn"`
	LastInboundAt *time.Time `json:"lastInboundAt,omitempty"`
}

// ListConversationsParams agrupa filtros + paginação.
type ListConversationsParams struct {
	UserID        uuid.UUID
	AssignedToMe  bool
	UnreadOnly    bool
	ChannelFilter string // "" | "email" | "whatsapp"
	Search        string // matches name/email/phone (LIKE)
	Limit         int    // default 50, max 200
	OffsetCursor  string // base64 of (last_at, owner_type, owner_id)
}

// ListConversationsResult devolve items + cursor pra próxima página.
type ListConversationsResult struct {
	Items      []ConversationItem `json:"items"`
	NextCursor string             `json:"nextCursor,omitempty"`
}

// conversationCursor é o que codificamos em base64 pra paginação.
type conversationCursor struct {
	LastAt    time.Time `json:"t"`
	OwnerType string    `json:"o"`
	OwnerID   uuid.UUID `json:"i"`
}

func encodeCursor(c conversationCursor) string {
	raw, _ := json.Marshal(c)
	return base64.RawURLEncoding.EncodeToString(raw)
}

func decodeCursor(s string) (*conversationCursor, error) {
	if s == "" {
		return nil, nil
	}
	raw, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return nil, fmt.Errorf("cursor inválido: %w", err)
	}
	var c conversationCursor
	if err := json.Unmarshal(raw, &c); err != nil {
		return nil, fmt.Errorf("cursor inválido: %w", err)
	}
	return &c, nil
}

// listRow espelha o resultado da query SQL raw — campos crus do PG.
type listRow struct {
	OwnerType        string
	OwnerID          uuid.UUID
	Name             *string
	Email            *string
	Phone            *string
	AssignedToUserID *uuid.UUID
	LeadStatus       *string
	LastAt           time.Time
	LastChannel      string
	LastDirection    string
	LastContent      *string // pode estar criptografado se channel=email
	Channels         string  // "email,whatsapp" (string_agg)
	UnreadCount      int
	EmailOptIn       bool
	WhatsAppOptIn    bool
	LastInboundAt    *time.Time
}

// List retorna conversas ordenadas por última atividade DESC.
//
// Estratégia: UNION ALL de (Lead + activities) + (Patient + activities), agregando
// por owner. Em SQL raw — GORM aggregations produzem várias queries serializadas e
// ficariam ~5x mais lentas pra 10k conversas.
//
// Performance esperada: <200ms pra <10k conversas com índices existentes (idx_lead
// _activities_lead, idx_lead_activities_patient + composite created_at).
func (s *ConversationService) List(ctx context.Context, params ListConversationsParams) (*ListConversationsResult, error) {
	if params.Limit <= 0 {
		params.Limit = 50
	}
	if params.Limit > 200 {
		params.Limit = 200
	}

	cursor, err := decodeCursor(params.OffsetCursor)
	if err != nil {
		return nil, err
	}

	// Filtros aplicados nos JOINs LEFT lateral abaixo. ChannelFilter e Search são
	// AND'd direto. AssignedToMe é interpretado como:
	//   - Lead.assigned_to_user_id = $userID
	//   - Patient é INCLUÍDO sempre (paciente não tem assignment direto; conversa
	//     com paciente é responsabilidade do time todo).
	// UnreadOnly filtra unread_count > 0 no WHERE final.

	args := map[string]any{
		"userID": params.UserID,
		"limit":  params.Limit + 1, // +1 pra detectar próxima página
	}
	conds := []string{"la.channel IN ('email','whatsapp')"}
	if params.ChannelFilter != "" {
		conds = append(conds, "la.channel = @channel")
		args["channel"] = params.ChannelFilter
	}

	leadWhere := []string{"l.deleted_at IS NULL"}
	if params.AssignedToMe {
		leadWhere = append(leadWhere, "l.assigned_to_user_id = @userID")
	}
	if params.Search != "" {
		args["search"] = "%" + strings.ToLower(strings.TrimSpace(params.Search)) + "%"
		leadWhere = append(leadWhere, "(LOWER(COALESCE(l.name,'')) LIKE @search OR LOWER(COALESCE(l.email,'')) LIKE @search OR COALESCE(l.phone,'') LIKE @search)")
	}

	patientWhere := []string{"p.deleted_at IS NULL"}
	if params.Search != "" {
		patientWhere = append(patientWhere, "(LOWER(COALESCE(p.name,'')) LIKE @search OR LOWER(COALESCE(p.email,'')) LIKE @search OR COALESCE(p.phone,'') LIKE @search)")
	}

	cursorClause := ""
	if cursor != nil {
		cursorClause = " AND (last_at, owner_type, owner_id::text) < (@curT, @curO, @curID::text) "
		args["curT"] = cursor.LastAt
		args["curO"] = cursor.OwnerType
		args["curID"] = cursor.OwnerID
	}

	unreadClause := ""
	if params.UnreadOnly {
		unreadClause = " AND unread_count > 0 "
	}

	// SQL: pra cada owner, calculamos last_at via aggregate; depois LATERAL pra pegar
	// a última activity (channel/direction/content); LEFT JOIN com conversation_reads
	// pra unread_count. Last activity vem por type direção: actor_user_id IS NULL → "in".
	query := fmt.Sprintf(`
WITH lead_conv AS (
	SELECT
		'lead'::text AS owner_type,
		l.id AS owner_id,
		l.name,
		l.email,
		l.phone,
		l.assigned_to_user_id,
		l.status::text AS lead_status,
		l.email_opt_in,
		l.whats_app_opt_in,
		l.last_inbound_at,
		MAX(la.created_at) AS last_at,
		COUNT(la.*) FILTER (
			WHERE la.created_at > COALESCE(cr.last_read_at, '1970-01-01'::timestamptz)
			AND la.actor_user_id IS NULL
			AND la.type IN ('message_received')
		) AS unread_count,
		string_agg(DISTINCT la.channel::text, ',') AS channels
	FROM leads l
	JOIN lead_activities la ON la.lead_id = l.id
	LEFT JOIN conversation_reads cr
		ON cr.user_id = @userID
		AND cr.owner_type = 'lead'
		AND cr.owner_id = l.id
	WHERE %s AND %s
	GROUP BY l.id, cr.last_read_at, l.email_opt_in, l.whats_app_opt_in, l.last_inbound_at
),
patient_conv AS (
	SELECT
		'patient'::text AS owner_type,
		p.id AS owner_id,
		p.name,
		p.email,
		p.phone,
		NULL::uuid AS assigned_to_user_id,
		NULL::text AS lead_status,
		TRUE AS email_opt_in,
		TRUE AS whats_app_opt_in,
		NULL::timestamptz AS last_inbound_at,
		MAX(la.created_at) AS last_at,
		COUNT(la.*) FILTER (
			WHERE la.created_at > COALESCE(cr.last_read_at, '1970-01-01'::timestamptz)
			AND la.actor_user_id IS NULL
			AND la.type IN ('message_received')
		) AS unread_count,
		string_agg(DISTINCT la.channel::text, ',') AS channels
	FROM patients p
	JOIN lead_activities la ON la.patient_id = p.id
	LEFT JOIN conversation_reads cr
		ON cr.user_id = @userID
		AND cr.owner_type = 'patient'
		AND cr.owner_id = p.id
	WHERE %s AND %s
	GROUP BY p.id, cr.last_read_at
),
unioned AS (
	SELECT * FROM lead_conv
	UNION ALL
	SELECT * FROM patient_conv
),
enriched AS (
	SELECT
		u.*,
		(SELECT la2.channel::text FROM lead_activities la2
			WHERE (u.owner_type = 'lead' AND la2.lead_id = u.owner_id)
			   OR (u.owner_type = 'patient' AND la2.patient_id = u.owner_id)
			ORDER BY la2.created_at DESC LIMIT 1) AS last_channel,
		(SELECT CASE WHEN la2.actor_user_id IS NULL THEN 'in' ELSE 'out' END
			FROM lead_activities la2
			WHERE (u.owner_type = 'lead' AND la2.lead_id = u.owner_id)
			   OR (u.owner_type = 'patient' AND la2.patient_id = u.owner_id)
			ORDER BY la2.created_at DESC LIMIT 1) AS last_direction,
		(SELECT la2.content FROM lead_activities la2
			WHERE (u.owner_type = 'lead' AND la2.lead_id = u.owner_id)
			   OR (u.owner_type = 'patient' AND la2.patient_id = u.owner_id)
			ORDER BY la2.created_at DESC LIMIT 1) AS last_content
	FROM unioned u
)
SELECT * FROM enriched
WHERE 1=1 %s %s
ORDER BY last_at DESC, owner_type DESC, owner_id DESC
LIMIT @limit
`,
		strings.Join(conds, " AND "),
		strings.Join(leadWhere, " AND "),
		strings.Join(conds, " AND "),
		strings.Join(patientWhere, " AND "),
		cursorClause,
		unreadClause,
	)

	var rows []listRow
	if err := s.db.WithContext(ctx).Raw(query, args).Scan(&rows).Error; err != nil {
		return nil, fmt.Errorf("conversation: list query: %w", err)
	}

	// Paginação: pegamos limit+1; se veio mais que limit, sobra é a próxima página.
	hasMore := len(rows) > params.Limit
	if hasMore {
		rows = rows[:params.Limit]
	}

	items := make([]ConversationItem, 0, len(rows))
	for _, r := range rows {
		// Decrypt content se for email (mesma heurística que LeadActivity.AfterFind).
		snippet := ""
		if r.LastContent != nil {
			snippet = decryptIfNeeded(*r.LastContent)
		}
		if len(snippet) > 120 {
			snippet = snippet[:117] + "..."
		}
		name := ""
		if r.Name != nil {
			name = *r.Name
		}
		if name == "" {
			if r.Email != nil {
				name = *r.Email
			} else if r.Phone != nil {
				name = *r.Phone
			} else {
				name = "(sem nome)"
			}
		}
		channels := []string{}
		if r.Channels != "" {
			for _, c := range strings.Split(r.Channels, ",") {
				c = strings.TrimSpace(c)
				if c != "" {
					channels = append(channels, c)
				}
			}
		}
		item := ConversationItem{
			OwnerType:        r.OwnerType,
			OwnerID:          r.OwnerID,
			Name:             name,
			Email:            r.Email,
			Phone:            r.Phone,
			LastChannel:      r.LastChannel,
			LastDirection:    r.LastDirection,
			LastSnippet:      snippet,
			LastAt:           r.LastAt,
			UnreadCount:      r.UnreadCount,
			AssignedToUserID: r.AssignedToUserID,
			Channels:         channels,
			LeadStatus:       r.LeadStatus,
			EmailOptIn:       r.EmailOptIn,
			WhatsAppOptIn:    r.WhatsAppOptIn,
			LastInboundAt:    r.LastInboundAt,
		}
		items = append(items, item)
	}

	res := &ListConversationsResult{Items: items}
	if hasMore && len(items) > 0 {
		last := items[len(items)-1]
		res.NextCursor = encodeCursor(conversationCursor{
			LastAt:    last.LastAt,
			OwnerType: last.OwnerType,
			OwnerID:   last.OwnerID,
		})
	}
	return res, nil
}

// decryptIfNeeded descriptografa snippet quando bate o formato AES-GCM base64 do
// crypto.EncryptWithDefaultKey. Como List() faz Raw scan (sem GORM hook AfterFind),
// precisamos descriptografar manualmente pra UX da lista mostrar preview legível.
// Em corrupção, retorna placeholder.
func decryptIfNeeded(s string) string {
	if s == "" {
		return ""
	}
	// Heurística rápida: ciphertext do nosso crypto é base64 puro (sem espaço/newline)
	// e tipicamente >40 chars. Texto plano normal tem espaços e/ou é mais curto.
	trimmed := strings.TrimSpace(s)
	if len(trimmed) < 40 || strings.ContainsAny(trimmed, " \n\r\t") {
		return s
	}
	// Tenta base64 decode + decrypt. Falha → assume plain.
	if _, err := base64.StdEncoding.DecodeString(trimmed); err != nil {
		return s
	}
	decrypted, err := crypto.DecryptWithDefaultKey(trimmed)
	if err != nil {
		// Falha silenciosa: pode ser texto plain que coincidiu com formato base64.
		// Devolve original em vez de placeholder pra não esconder mensagens válidas.
		return s
	}
	return decrypted
}

// ============================================================
// Messages — histórico de uma conversa
// ============================================================

// GetMessagesParams agrupa filtros pra Messages().
type GetMessagesParams struct {
	UserID    uuid.UUID
	OwnerType string
	OwnerID   uuid.UUID
	Limit     int        // default 100, max 500
	Before    *time.Time // pra paginação infinite scroll
}

// Messages retorna activities dum owner em ordem cronológica DESC (mais recente primeiro)
// — UI scrolls reverso. Com Before, pega activities CRIADAS antes daquele timestamp.
//
// Side-effect: marca conversa como lida (LastReadAt = now). Quem abre, leu — UX direta,
// sem precisar de "Marcar como lido" explícito (mantemos o endpoint POST /read mesmo
// pra UI marcar sem fetch).
func (s *ConversationService) Messages(ctx context.Context, p GetMessagesParams) ([]models.LeadActivity, error) {
	if p.Limit <= 0 {
		p.Limit = 100
	}
	if p.Limit > 500 {
		p.Limit = 500
	}
	if !isValidOwnerType(p.OwnerType) || p.OwnerID == uuid.Nil {
		return nil, ErrConversationOwnerInvalid
	}

	q := s.db.WithContext(ctx).Model(&models.LeadActivity{}).
		Preload("Actor")
	switch p.OwnerType {
	case string(models.ConversationOwnerLead):
		q = q.Where("lead_id = ?", p.OwnerID)
	case string(models.ConversationOwnerPatient):
		q = q.Where("patient_id = ?", p.OwnerID)
	}
	if p.Before != nil {
		q = q.Where("created_at < ?", *p.Before)
	}

	var activities []models.LeadActivity
	if err := q.Order("created_at DESC").Limit(p.Limit).Find(&activities).Error; err != nil {
		return nil, fmt.Errorf("conversation: messages query: %w", err)
	}

	// Side-effect: marca lido. Não-bloqueante; erros logam mas não quebram fetch.
	if err := s.MarkRead(ctx, p.UserID, p.OwnerType, p.OwnerID); err != nil {
		// não retornar erro — fetch deve sempre funcionar
		_ = err
	}
	return activities, nil
}

// ============================================================
// MarkRead — upsert ConversationRead.LastReadAt = now
// ============================================================

// MarkRead atualiza LastReadAt = now pro (user, owner). Idempotente (upsert).
func (s *ConversationService) MarkRead(ctx context.Context, userID uuid.UUID, ownerType string, ownerID uuid.UUID) error {
	if !isValidOwnerType(ownerType) || ownerID == uuid.Nil || userID == uuid.Nil {
		return ErrConversationOwnerInvalid
	}
	now := time.Now().UTC()

	// Upsert via ON CONFLICT (idx_conv_reads_user_owner é UNIQUE composto).
	// GORM Clauses pra evitar SELECT+UPDATE sequencial.
	read := models.ConversationRead{
		UserID:     userID,
		OwnerType:  ownerType,
		OwnerID:    ownerID,
		LastReadAt: now,
	}
	if err := s.db.WithContext(ctx).Exec(`
		INSERT INTO conversation_reads (id, user_id, owner_type, owner_id, last_read_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
		ON CONFLICT (user_id, owner_type, owner_id)
		DO UPDATE SET last_read_at = EXCLUDED.last_read_at, updated_at = EXCLUDED.updated_at
	`, uuid.Must(uuid.NewV7()), read.UserID, read.OwnerType, read.OwnerID, read.LastReadAt, now).Error; err != nil {
		return fmt.Errorf("conversation: mark read: %w", err)
	}
	return nil
}

// ============================================================
// SendMessage — outbound (email ou whatsapp)
// ============================================================

// OutboundAttachment é um ref de anexo já uploadado (sem bytes — service lê do disco).
// Path é relativo ao uploadRoot (/app/uploads). Filename é o nome mostrado ao destinatário.
type OutboundAttachment struct {
	Path     string
	Filename string
}

// SendMessageInput agrupa o payload da rota POST /conversations/:type/:id/email ou /whatsapp.
type SendMessageInput struct {
	UserID      uuid.UUID
	OwnerType   string
	OwnerID     uuid.UUID
	Channel     models.LeadActivityChannel // email | whatsapp
	Subject     string                     // email; backend resolve se vazio
	BodyText    string
	BodyHTML    string // só email
	InReplyTo   string // só email
	References  []string
	Attachments []OutboundAttachment // só email
}

// SendMessage rota mensagem outbound pro canal escolhido. Reusa EmailService /
// WhatsAppService — não duplica lógica de envio.
//
// Validações:
//   - Email:    Owner.Email != nil; pra Lead também checa EmailOptIn
//   - WhatsApp: Owner.Phone != nil; pra Lead também checa WhatsAppOptIn + janela 24h
//
// Patient não tem flags granulares de opt-in (consent foi coletado no signup); checamos
// só presença do canal. Se requisitos LGPD apertarem, fácil estender.
//
// Returna a LeadActivity criada (com decryption já aplicado via AfterFind).
func (s *ConversationService) SendMessage(ctx context.Context, in SendMessageInput) (*models.LeadActivity, error) {
	if !isValidOwnerType(in.OwnerType) || in.OwnerID == uuid.Nil {
		return nil, ErrConversationOwnerInvalid
	}
	if in.UserID == uuid.Nil {
		return nil, errors.New("conversation: actorUserID obrigatório")
	}

	owner, err := s.loadOwner(ctx, in.OwnerType, in.OwnerID)
	if err != nil {
		return nil, err
	}

	switch in.Channel {
	case models.LeadChannelEmail:
		return s.sendEmail(ctx, in, owner)
	case models.LeadChannelWhatsApp:
		return s.sendWhatsApp(ctx, in, owner)
	default:
		return nil, fmt.Errorf("conversation: canal inválido %q", in.Channel)
	}
}

// loadOwner carrega o Lead OU Patient indicado. Não preload Activities (custoso e não usado).
func (s *ConversationService) loadOwner(ctx context.Context, ownerType string, ownerID uuid.UUID) (ConversationOwner, error) {
	switch ownerType {
	case string(models.ConversationOwnerLead):
		var lead models.Lead
		if err := s.db.WithContext(ctx).First(&lead, "id = ?", ownerID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ConversationOwner{}, ErrConversationOwnerInvalid
			}
			return ConversationOwner{}, fmt.Errorf("conversation: load lead: %w", err)
		}
		return ConversationOwner{Lead: &lead}, nil
	case string(models.ConversationOwnerPatient):
		var patient models.Patient
		if err := s.db.WithContext(ctx).First(&patient, "id = ?", ownerID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ConversationOwner{}, ErrConversationOwnerInvalid
			}
			return ConversationOwner{}, fmt.Errorf("conversation: load patient: %w", err)
		}
		return ConversationOwner{Patient: &patient}, nil
	}
	return ConversationOwner{}, ErrConversationOwnerInvalid
}

// sendEmail despacha pro EmailService.SendConversationReply e devolve a LeadActivity criada.
func (s *ConversationService) sendEmail(ctx context.Context, in SendMessageInput, owner ConversationOwner) (*models.LeadActivity, error) {
	if s.emailService == nil {
		return nil, errors.New("conversation: email service não configurado")
	}

	// Pre-checks específicos pra retornar erros semânticos antes de chamar Resend.
	if owner.Lead != nil {
		l := owner.Lead
		if l.Email == nil || strings.TrimSpace(*l.Email) == "" {
			return nil, ErrConversationNoChannel
		}
		if !l.EmailOptIn {
			return nil, ErrConversationOptedOut
		}
	}
	if owner.Patient != nil {
		p := owner.Patient
		if p.Email == nil || strings.TrimSpace(*p.Email) == "" {
			return nil, ErrConversationNoChannel
		}
	}

	if strings.TrimSpace(in.BodyText) == "" && strings.TrimSpace(in.BodyHTML) == "" {
		return nil, errors.New("conversation: corpo vazio")
	}

	// Snapshot do timestamp pré-envio pra recuperar a activity criada pelo EmailService.
	preSend := time.Now().UTC().Add(-1 * time.Second)

	atts := make([]ReplyAttachment, 0, len(in.Attachments))
	for _, a := range in.Attachments {
		atts = append(atts, ReplyAttachment{Path: a.Path, Filename: a.Filename})
	}

	if err := s.emailService.SendConversationReply(ctx, SendConversationReplyInput{
		Owner:       owner,
		ActorUserID: in.UserID,
		Subject:     in.Subject,
		BodyText:    in.BodyText,
		BodyHTML:    in.BodyHTML,
		InReplyTo:   in.InReplyTo,
		References:  in.References,
		Attachments: atts,
	}); err != nil {
		switch {
		case errors.Is(err, ErrLeadOptedOut):
			return nil, ErrConversationOptedOut
		case errors.Is(err, ErrLeadNoEmail):
			return nil, ErrConversationNoChannel
		case errors.Is(err, ErrEmailAttachmentTooLarge):
			return nil, ErrConversationAttachmentTooLarge
		case errors.Is(err, ErrEmailAttachmentInvalid):
			return nil, ErrConversationAttachmentInvalid
		default:
			return nil, err
		}
	}

	// Recupera a activity criada (mais recente do owner pós-preSend, type=message_sent).
	return s.findLatestActivity(ctx, in.OwnerType, in.OwnerID, models.LeadActivityMessageSent, preSend)
}

// sendWhatsApp valida janela 24h + opt-in (pra Lead) e dispara WhatsAppService.SendTextMessage.
func (s *ConversationService) sendWhatsApp(ctx context.Context, in SendMessageInput, owner ConversationOwner) (*models.LeadActivity, error) {
	if s.whatsappService == nil {
		return nil, errors.New("conversation: whatsapp service não configurado")
	}
	body := strings.TrimSpace(in.BodyText)
	hasMedia := len(in.Attachments) > 0
	if body == "" && !hasMedia {
		return nil, errors.New("conversation: corpo vazio")
	}
	if len(body) > 4096 {
		return nil, errors.New("conversation: mensagem excede 4096 caracteres")
	}

	var (
		phone     string
		leadID    *uuid.UUID
		patientID *uuid.UUID
	)
	switch {
	case owner.Lead != nil:
		l := owner.Lead
		if l.Phone == nil || *l.Phone == "" {
			return nil, ErrConversationNoChannel
		}
		if !l.WhatsAppOptIn {
			return nil, ErrConversationOptedOut
		}
		if l.Status == models.LeadStatusUnsubscribed {
			return nil, ErrConversationOptedOut
		}
		// Janela 24h: WhatsApp Business só permite session messages dentro disso.
		if l.LastInboundAt == nil || time.Since(*l.LastInboundAt) > 24*time.Hour {
			return nil, ErrConversationWindowClosed
		}
		phone = *l.Phone
		id := l.ID
		leadID = &id
	case owner.Patient != nil:
		p := owner.Patient
		if p.Phone == nil || *p.Phone == "" {
			return nil, ErrConversationNoChannel
		}
		// Janela 24h pra Patient: derivamos da última inbound activity do paciente.
		// Sem index dedicado, é OK porque a query usa idx_lead_activities_patient.
		var lastInbound models.LeadActivity
		err := s.db.WithContext(ctx).
			Where("patient_id = ? AND channel = ? AND type = ?",
				p.ID, models.LeadChannelWhatsApp, models.LeadActivityMessageReceived).
			Order("created_at DESC").
			First(&lastInbound).Error
		if err != nil {
			return nil, ErrConversationWindowClosed
		}
		if time.Since(lastInbound.CreatedAt) > 24*time.Hour {
			return nil, ErrConversationWindowClosed
		}
		phone = *p.Phone
		id := p.ID
		patientID = &id
	}

	var (
		wamid           string
		mediaType       *string
		mediaMime       *string
		mediaFilename   *string
		mediaStorageKey *string
		mediaSize       *int64
		err             error
	)

	if hasMedia {
		att := in.Attachments[0] // WhatsApp envia uma mídia por mensagem
		full, rerr := s.resolveUploadPath(att.Path)
		if rerr != nil {
			return nil, rerr
		}
		data, rerr := os.ReadFile(full)
		if rerr != nil {
			return nil, ErrConversationAttachmentInvalid
		}
		mime := DetectBytesMime(data).ContentType
		waType := WhatsAppMediaTypeFromMIME(mime)

		mediaID, uerr := s.whatsappService.UploadMedia(data, mime, att.Filename)
		if uerr != nil {
			return nil, fmt.Errorf("conversation: whatsapp upload media: %w", uerr)
		}
		wamid, err = s.whatsappService.SendMediaMessage(phone, waType, mediaID, body, att.Filename)
		if err != nil {
			return nil, fmt.Errorf("conversation: whatsapp send media: %w", err)
		}
		// Guarda a mídia cifrada pra renderizar no histórico (mesmo bucket do inbound).
		if s.waMedia != nil {
			ownerKind := "lead"
			oid := uuid.Nil
			if patientID != nil {
				ownerKind, oid = "patient", *patientID
			} else if leadID != nil {
				oid = *leadID
			}
			if key, size, serr := s.waMedia.StoreInbound(ownerKind, oid, data, mime); serr == nil {
				mediaStorageKey, mediaSize = &key, &size
			}
		}
		mediaType, mediaMime, mediaFilename = &waType, &mime, &att.Filename
	} else {
		wamid, err = s.whatsappService.SendTextMessage(phone, body)
		if err != nil {
			return nil, fmt.Errorf("conversation: whatsapp send: %w", err)
		}
	}

	// Persiste LeadActivity outbound — espelha lead_service.go::SendWhatsAppText.
	metaJSON, _ := json.Marshal(map[string]any{
		"wa_message_id": wamid,
		"sent_at":       time.Now().UTC(),
		"recipient":     phone,
	})
	var contentPtr *string
	if body != "" {
		contentPtr = &body
	}
	activity := &models.LeadActivity{
		LeadID:          leadID,
		PatientID:       patientID,
		Type:            models.LeadActivityMessageSent,
		Channel:         models.LeadChannelWhatsApp,
		Content:         contentPtr,
		Metadata:        datatypes.JSON(metaJSON),
		ActorUserID:     &in.UserID,
		MediaType:       mediaType,
		MediaMIME:       mediaMime,
		MediaFilename:   mediaFilename,
		MediaStorageKey: mediaStorageKey,
		MediaSizeBytes:  mediaSize,
	}
	if err := s.db.WithContext(ctx).Create(activity).Error; err != nil {
		return nil, fmt.Errorf("conversation: persist activity: %w", err)
	}
	// Refetch pra rodar AfterFind (decrypt content).
	return s.fetchActivity(ctx, activity.ID)
}

// SendWhatsAppTemplateInput é o payload de envio de template (reabre conversa >24h).
type SendWhatsAppTemplateInput struct {
	UserID       uuid.UUID
	OwnerType    string
	OwnerID      uuid.UUID
	TemplateName string
	Language     string
	Params       []string
}

// SendWhatsAppTemplate envia um template aprovado pela Meta. Diferente de sendWhatsApp,
// NÃO checa a janela de 24h — template é justamente o caminho pra reabrir conversa
// fora dela. Valida opt-in/canal e persiste a activity outbound.
func (s *ConversationService) SendWhatsAppTemplate(ctx context.Context, in SendWhatsAppTemplateInput) (*models.LeadActivity, error) {
	if s.whatsappService == nil {
		return nil, errors.New("conversation: whatsapp service não configurado")
	}
	if !isValidOwnerType(in.OwnerType) || in.OwnerID == uuid.Nil {
		return nil, ErrConversationOwnerInvalid
	}
	if in.UserID == uuid.Nil {
		return nil, errors.New("conversation: actorUserID obrigatório")
	}
	if strings.TrimSpace(in.TemplateName) == "" {
		return nil, errors.New("conversation: template obrigatório")
	}
	lang := in.Language
	if lang == "" {
		lang = "pt_BR"
	}

	owner, err := s.loadOwner(ctx, in.OwnerType, in.OwnerID)
	if err != nil {
		return nil, err
	}

	var (
		phone     string
		leadID    *uuid.UUID
		patientID *uuid.UUID
	)
	switch {
	case owner.Lead != nil:
		l := owner.Lead
		if l.Phone == nil || *l.Phone == "" {
			return nil, ErrConversationNoChannel
		}
		if !l.WhatsAppOptIn || l.Status == models.LeadStatusUnsubscribed {
			return nil, ErrConversationOptedOut
		}
		phone = *l.Phone
		id := l.ID
		leadID = &id
	case owner.Patient != nil:
		p := owner.Patient
		if p.Phone == nil || *p.Phone == "" {
			return nil, ErrConversationNoChannel
		}
		phone = *p.Phone
		id := p.ID
		patientID = &id
	}

	if err := s.whatsappService.SendTemplate(phone, in.TemplateName, lang, in.Params); err != nil {
		return nil, fmt.Errorf("conversation: whatsapp template: %w", err)
	}

	content := "[template: " + in.TemplateName + "]"
	metaJSON, _ := json.Marshal(map[string]any{
		"template":  in.TemplateName,
		"language":  lang,
		"sent_at":   time.Now().UTC(),
		"recipient": phone,
	})
	activity := &models.LeadActivity{
		LeadID:      leadID,
		PatientID:   patientID,
		Type:        models.LeadActivityMessageSent,
		Channel:     models.LeadChannelWhatsApp,
		Content:     &content,
		Metadata:    datatypes.JSON(metaJSON),
		ActorUserID: &in.UserID,
	}
	if err := s.db.WithContext(ctx).Create(activity).Error; err != nil {
		return nil, fmt.Errorf("conversation: persist activity: %w", err)
	}
	return s.fetchActivity(ctx, activity.ID)
}

// findLatestActivity busca a activity mais recente de um owner, do type esperado e
// criada DEPOIS de preSend. Usado por sendEmail pra recuperar a activity que o
// EmailService.SendConversationReply criou — evita refatorar pra retornar pointer.
func (s *ConversationService) findLatestActivity(
	ctx context.Context,
	ownerType string,
	ownerID uuid.UUID,
	actType models.LeadActivityType,
	since time.Time,
) (*models.LeadActivity, error) {
	q := s.db.WithContext(ctx).Model(&models.LeadActivity{}).
		Where("type = ? AND created_at >= ?", actType, since).
		Preload("Actor")
	switch ownerType {
	case string(models.ConversationOwnerLead):
		q = q.Where("lead_id = ?", ownerID)
	case string(models.ConversationOwnerPatient):
		q = q.Where("patient_id = ?", ownerID)
	}
	var act models.LeadActivity
	if err := q.Order("created_at DESC").First(&act).Error; err != nil {
		return nil, fmt.Errorf("conversation: refetch activity: %w", err)
	}
	return &act, nil
}

// fetchActivity refaz GET pra rodar AfterFind (decryption do Content).
func (s *ConversationService) fetchActivity(ctx context.Context, id uuid.UUID) (*models.LeadActivity, error) {
	var act models.LeadActivity
	if err := s.db.WithContext(ctx).Preload("Actor").First(&act, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &act, nil
}

// isValidOwnerType valida que o tipo é um dos enums esperados.
func isValidOwnerType(t string) bool {
	return t == string(models.ConversationOwnerLead) || t == string(models.ConversationOwnerPatient)
}

// ============================================================
// Compose — vendedor inicia email pra qualquer endereço
// ============================================================

// ErrComposeRecipientUnsubscribed → destinatário possui Lead com Status=unsubscribed.
// Backend rejeita compose pra evitar reativar contato que pediu pra parar.
var ErrComposeRecipientUnsubscribed = errors.New("conversation: destinatário descadastrado")

// ErrComposeInvalidEmail → formato do "to" não passou na validação básica.
var ErrComposeInvalidEmail = errors.New("conversation: email do destinatário inválido")

// composeEmailRegex é validação RFC 5322 ultra-simplificada (suficiente pra UI).
// Backend não precisa validar 100% — o provider (Resend) bounce-a se for inválido.
// Frontend usa a mesma regex pra feedback rápido.
var composeEmailRegex = regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`)

// ComposeInput é o payload do POST /conversations/compose.
type ComposeInput struct {
	UserID      uuid.UUID
	To          string
	Name        string // opcional; se Lead novo for criado, vira Lead.Name
	Subject     string
	BodyText    string
	BodyHTML    string
	Attachments []OutboundAttachment
}

// ComposeResult é o retorno do compose — frontend usa pra navegar pra conversa criada
// e pra exibir toast diferenciado quando Lead novo nasce.
type ComposeResult struct {
	OwnerType   string    `json:"ownerType"`
	OwnerID     uuid.UUID `json:"ownerId"`
	LeadCreated bool      `json:"leadCreated"`
	URL         string    `json:"url,omitempty"`
}

// Compose resolve o destinatário e dispara o email. Hierarquia de roteamento:
//
//  1. Patient com email igual → usa Patient como owner (sem criar Lead).
//  2. Lead ativo (status NOT IN converted/lost/unsubscribed) → usa esse Lead.
//  3. Lead descadastrado existente → REJEITA com ErrComposeRecipientUnsubscribed.
//  4. Caso contrário → cria Lead novo (source=manual) e usa como owner.
//
// O envio em si reusa SendMessage(channel=email) que já cuida de:
//   - Resend POST + headers de threading (X-Plenya-Source)
//   - LeadActivity{message_sent, channel=email} anexada ao owner correto
//   - IMAP APPEND best-effort no Sent Items
//
// Notification: só dispara NotifyTeamOfNewLead pra Lead novo (LeadCreated=true).
// Reusar Lead/Patient não notifica (já é conhecido).
func (s *ConversationService) Compose(ctx context.Context, in ComposeInput) (*ComposeResult, *models.LeadActivity, error) {
	if s.leadService == nil {
		return nil, nil, errors.New("conversation: lead service não configurado")
	}
	if in.UserID == uuid.Nil {
		return nil, nil, errors.New("conversation: actorUserID obrigatório")
	}
	to := strings.ToLower(strings.TrimSpace(in.To))
	if to == "" || !composeEmailRegex.MatchString(to) {
		return nil, nil, ErrComposeInvalidEmail
	}
	if strings.TrimSpace(in.Subject) == "" {
		return nil, nil, errors.New("conversation: assunto obrigatório no compose")
	}
	if strings.TrimSpace(in.BodyText) == "" && strings.TrimSpace(in.BodyHTML) == "" {
		return nil, nil, errors.New("conversation: corpo vazio")
	}

	var (
		ownerType   string
		ownerID     uuid.UUID
		leadCreated bool
	)

	// 1. Patient first
	var patient models.Patient
	if err := s.db.WithContext(ctx).Where("LOWER(email) = ?", to).First(&patient).Error; err == nil {
		ownerType = string(models.ConversationOwnerPatient)
		ownerID = patient.ID
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil, fmt.Errorf("conversation: compose patient lookup: %w", err)
	}

	// 2/3. Lead lookup (ativo OU unsubscribed pra rejeitar explicitamente)
	if ownerID == uuid.Nil {
		// Checa unsubscribed primeiro pra mensagem de erro semântica.
		var unsub models.Lead
		err := s.db.WithContext(ctx).
			Where("LOWER(email) = ? AND status = ?", to, models.LeadStatusUnsubscribed).
			Order("created_at DESC").
			First(&unsub).Error
		if err == nil {
			return nil, nil, ErrComposeRecipientUnsubscribed
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, fmt.Errorf("conversation: compose unsub lookup: %w", err)
		}

		var existing models.Lead
		err = s.db.WithContext(ctx).
			Where("LOWER(email) = ? AND status NOT IN (?)", to,
				[]models.LeadStatus{models.LeadStatusConverted, models.LeadStatusLost, models.LeadStatusUnsubscribed}).
			Order("created_at DESC").
			First(&existing).Error
		if err == nil {
			ownerType = string(models.ConversationOwnerLead)
			ownerID = existing.ID
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, fmt.Errorf("conversation: compose lead lookup: %w", err)
		}
	}

	// 4. Cria Lead novo
	if ownerID == uuid.Nil {
		newLead, err := s.leadService.CreateForManualCompose(to, in.Name, in.UserID)
		if err != nil {
			return nil, nil, fmt.Errorf("conversation: compose create lead: %w", err)
		}
		ownerType = string(models.ConversationOwnerLead)
		ownerID = newLead.ID
		leadCreated = true
	}

	// Envia email reusando SendMessage (já trata attachments, threading, persistência).
	activity, sendErr := s.SendMessage(ctx, SendMessageInput{
		UserID:      in.UserID,
		OwnerType:   ownerType,
		OwnerID:     ownerID,
		Channel:     models.LeadChannelEmail,
		Subject:     in.Subject,
		BodyText:    in.BodyText,
		BodyHTML:    in.BodyHTML,
		Attachments: in.Attachments,
	})
	if sendErr != nil {
		// Lead novo foi criado mas envio falhou — mantemos o Lead (vendedor pode tentar
		// de novo via composer normal da conversa). Erro chega ao handler com semântica
		// preservada (opt-out, attachment too large, etc).
		return nil, nil, sendErr
	}

	// Notifica equipe SÓ pra Lead novo (Lead/Patient pré-existente já é conhecido).
	if leadCreated {
		var leadCopy models.Lead
		if err := s.db.WithContext(ctx).First(&leadCopy, "id = ?", ownerID).Error; err == nil {
			leadRef := leadCopy
			goSafe("notify_compose", func() { s.leadService.NotifyTeamOfNewLead(&leadRef) })
		}
	}

	res := &ComposeResult{
		OwnerType:   ownerType,
		OwnerID:     ownerID,
		LeadCreated: leadCreated,
		URL:         fmt.Sprintf("/conversas?focus=%s:%s", ownerType, ownerID.String()),
	}
	return res, activity, nil
}
