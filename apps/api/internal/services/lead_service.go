package services

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// LeadService gerencia leads (captura, listagem, conversão em Patient).
type LeadService struct {
	db *gorm.DB
}

func NewLeadService(db *gorm.DB) *LeadService {
	return &LeadService{db: db}
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
			_ = s.RecordActivity(RecordActivityInput{
				LeadID:  newLead.ID,
				Type:    models.LeadActivityCreated,
				Channel: models.LeadChannelInternal,
				Content: ptr("Lead criado via claim do Escore Light"),
			})
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

	_ = s.RecordActivity(RecordActivityInput{
		LeadID:  lead.ID,
		Type:    models.LeadActivityCreated,
		Channel: models.LeadChannelInternal,
		Content: ptr("Lead criado via formulário /contato"),
	})

	return &lead, nil
}

// ============================================================
// WhatsApp inbound — primeiro contato vira Lead
// ============================================================

type InboundWhatsAppInput struct {
	PhoneE164    string // sem +, formato Meta (ex: 5511999998888)
	Name         *string
	Text         string
	WAMessageID  string
	ReceivedAt   time.Time
}

// ProcessInboundWhatsApp registra mensagem inbound. Se já existe Lead ativo pra esse phone,
// só adiciona LeadActivity. Caso contrário, cria Lead novo com source=whatsapp_inbound.
// Consent implícito: o cliente iniciou a conversa.
func (s *LeadService) ProcessInboundWhatsApp(in InboundWhatsAppInput) (*models.Lead, error) {
	if strings.TrimSpace(in.PhoneE164) == "" {
		return nil, errors.New("lead: phone vazio em inbound WhatsApp")
	}

	// Normaliza para formato armazenado em Lead.Phone (com +)
	phone := strings.TrimSpace(in.PhoneE164)
	if !strings.HasPrefix(phone, "+") {
		phone = "+" + phone
	}

	// Procura Lead ativo (não converted/lost/unsubscribed) pra esse phone
	var existing models.Lead
	err := s.db.
		Where("phone = ? AND status NOT IN (?)", phone,
			[]models.LeadStatus{models.LeadStatusConverted, models.LeadStatusLost, models.LeadStatusUnsubscribed}).
		Order("created_at DESC").
		First(&existing).Error

	if err == nil {
		// Append atividade
		_ = s.RecordActivity(RecordActivityInput{
			LeadID:  existing.ID,
			Type:    models.LeadActivityMessageReceived,
			Channel: models.LeadChannelWhatsApp,
			Content: &in.Text,
			Metadata: map[string]any{
				"wa_message_id": in.WAMessageID,
				"received_at":   in.ReceivedAt,
			},
		})
		return &existing, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("lead: inbound WA lookup: %w", err)
	}

	// Cria novo Lead
	now := in.ReceivedAt
	if now.IsZero() {
		now = time.Now().UTC()
	}
	consentVer := "whatsapp_inbound_implicit"
	consentIP := "" // inbound WA não tem IP do cliente

	newLead := models.Lead{
		Source:           models.LeadSourceWhatsAppInbound,
		Status:           models.LeadStatusNew,
		Phone:            &phone,
		Name:             in.Name,
		Message:          &in.Text,
		WhatsAppOptIn:    true, // cliente iniciou conversa = consent implícito
		ConsentVersion:   &consentVer,
		ConsentTimestamp: &now,
		ConsentIPHash:    &consentIP,
	}
	if err := s.db.Create(&newLead).Error; err != nil {
		return nil, fmt.Errorf("lead: create from WA inbound: %w", err)
	}

	_ = s.RecordActivity(RecordActivityInput{
		LeadID:  newLead.ID,
		Type:    models.LeadActivityCreated,
		Channel: models.LeadChannelInternal,
		Content: ptr("Lead criado via primeira mensagem inbound do WhatsApp"),
	})
	_ = s.RecordActivity(RecordActivityInput{
		LeadID:  newLead.ID,
		Type:    models.LeadActivityMessageReceived,
		Channel: models.LeadChannelWhatsApp,
		Content: &in.Text,
		Metadata: map[string]any{
			"wa_message_id": in.WAMessageID,
			"received_at":   now,
		},
	})

	return &newLead, nil
}

// ============================================================
// LeadActivity helpers
// ============================================================

type RecordActivityInput struct {
	LeadID      uuid.UUID
	Type        models.LeadActivityType
	Channel     models.LeadActivityChannel
	Content     *string
	Metadata    map[string]any
	ActorUserID *uuid.UUID
}

func (s *LeadService) RecordActivity(in RecordActivityInput) error {
	if in.LeadID == uuid.Nil {
		return errors.New("lead activity: leadID obrigatório")
	}
	var metaJSON datatypes.JSON
	if len(in.Metadata) > 0 {
		raw, err := json.Marshal(in.Metadata)
		if err != nil {
			return fmt.Errorf("lead activity: metadata marshal: %w", err)
		}
		metaJSON = raw
	}
	activity := models.LeadActivity{
		LeadID:      in.LeadID,
		Type:        in.Type,
		Channel:     in.Channel,
		Content:     in.Content,
		Metadata:    metaJSON,
		ActorUserID: in.ActorUserID,
	}
	return s.db.Create(&activity).Error
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

	if statusChanged {
		_ = s.RecordActivity(RecordActivityInput{
			LeadID:      id,
			Type:        models.LeadActivityStatusChanged,
			Channel:     models.LeadChannelInternal,
			Content:     ptr(fmt.Sprintf("Status: %s → %s", lead.Status, *patch.Status)),
			ActorUserID: &actorUserID,
		})
	}
	if patch.AssignedToUserID != nil {
		_ = s.RecordActivity(RecordActivityInput{
			LeadID:      id,
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
		LeadID:      leadID,
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
	return s.RecordActivity(RecordActivityInput{
		LeadID:      leadID,
		Type:        models.LeadActivityConverted,
		Channel:     models.LeadChannelInternal,
		Content:     ptr(fmt.Sprintf("Lead convertido em paciente %s", patientID.String())),
		ActorUserID: actorUserID,
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
	birth := time.Now().UTC() // placeholder — admin pode editar depois no Patient
	if in.BirthDate != nil {
		birth = *in.BirthDate
	}

	var patient models.Patient
	err = s.db.Transaction(func(tx *gorm.DB) error {
		userEmail := email
		if userEmail == "" {
			// Sem email — gera placeholder determinístico baseado no phone
			userEmail = fmt.Sprintf("wa-%s@placeholder.plenyasaude.com.br",
				strings.TrimPrefix(phone, "+"))
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
