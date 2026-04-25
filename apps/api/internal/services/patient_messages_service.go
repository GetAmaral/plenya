package services

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// PatientMessagesService — caixa de mensagens do paciente no portal.
//
// Modelo: reusa LeadActivity (com PatientID setado) — a equipe vê a mesma
// conversa na Central de Conversas do EMR. Direção:
//   - Equipe → paciente: ActorUserID != nil (ou Channel email/whatsapp via SendMessage staff)
//   - Paciente → equipe: Channel=portal, ActorUserID = userID do paciente
//
// O paciente NÃO escolhe canal de saída — a mensagem fica no portal e a equipe
// responde dali. Notificação in-app pra staff via NotificationService.
type PatientMessagesService struct {
	db           *gorm.DB
	notification *NotificationService
}

func NewPatientMessagesService(db *gorm.DB, notification *NotificationService) *PatientMessagesService {
	return &PatientMessagesService{db: db, notification: notification}
}

// MessageView é o que o paciente vê: sem campos internos da equipe.
type MessageView struct {
	ID        uuid.UUID  `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	From      string     `json:"from"` // "patient" | "team"
	AuthorName string    `json:"authorName,omitempty"`
	Channel   string     `json:"channel"`
	Content   string     `json:"content"`
}

// List retorna mensagens cronológicas DESC (mais nova primeiro).
// Filtra: só Channel email/whatsapp/portal/internal-com-content visível ao paciente
// (eventos de status/mudança ficam ocultos).
func (s *PatientMessagesService) List(patientID uuid.UUID, limit int, before *time.Time) ([]MessageView, error) {
	if limit <= 0 || limit > 200 {
		limit = 100
	}

	q := s.db.Model(&models.LeadActivity{}).
		Where("patient_id = ?", patientID).
		Where("channel IN ?", []string{"email", "whatsapp", "portal"}).
		Where("type IN ?", []string{
			string(models.LeadActivityMessageSent),
			string(models.LeadActivityMessageReceived),
		}).
		Preload("Actor").
		Order("created_at DESC").
		Limit(limit)

	if before != nil {
		q = q.Where("created_at < ?", *before)
	}

	var rows []models.LeadActivity
	if err := q.Find(&rows).Error; err != nil {
		return nil, err
	}

	out := make([]MessageView, 0, len(rows))
	for i := range rows {
		out = append(out, *toMessageView(&rows[i]))
	}
	return out, nil
}

// Send cria uma mensagem do paciente (Channel=portal, Type=message_received,
// ActorUserID = userID do paciente).
//
// "received" da perspectiva da equipe (entrada inbound), igual outras mensagens
// inbound (WA/email do cliente).
func (s *PatientMessagesService) Send(patientID, userID uuid.UUID, content string) (*MessageView, error) {
	content = strings.TrimSpace(content)
	if content == "" {
		return nil, errors.New("conteúdo vazio")
	}
	if len(content) > 5000 {
		return nil, errors.New("mensagem muito longa (máx 5000 caracteres)")
	}

	pid := patientID
	uid := userID
	activity := &models.LeadActivity{
		PatientID:   &pid,
		Type:        models.LeadActivityMessageReceived,
		Channel:     models.LeadChannelPortal,
		Content:     &content,
		ActorUserID: &uid,
	}
	if err := s.db.Create(activity).Error; err != nil {
		return nil, fmt.Errorf("failed to create message: %w", err)
	}

	// Recarrega com Actor pra toMessageView() classificar from=patient corretamente.
	if err := s.db.Preload("Actor").First(activity, "id = ?", activity.ID).Error; err != nil {
		// Não fatal — retorna sem Actor (UI marcará como team mas msg foi salva)
		_ = err
	}

	// Notifica equipe (admin/manager/secretary) — best-effort, não bloqueia retorno.
	go s.notifyStaff(patientID, content)

	return toMessageView(activity), nil
}

// UnreadCount conta mensagens da equipe (Channel=email/whatsapp/internal/portal,
// Type=message_sent, ActorUserID != nil) criadas após o último read do paciente.
//
// O cursor de leitura do PACIENTE fica em ConversationRead com OwnerType=patient
// e UserID = userID do paciente (reusa a tabela existente).
func (s *PatientMessagesService) UnreadCount(patientID, userID uuid.UUID) (int, error) {
	var read models.ConversationRead
	cursor := time.Time{}
	err := s.db.Where("user_id = ? AND owner_type = ? AND owner_id = ?",
		userID, "patient", patientID).First(&read).Error
	if err == nil {
		cursor = read.LastReadAt
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}

	var count int64
	q := s.db.Model(&models.LeadActivity{}).
		Where("patient_id = ?", patientID).
		Where("type = ?", models.LeadActivityMessageSent).
		Where("channel IN ?", []string{"email", "whatsapp", "internal"})
	if !cursor.IsZero() {
		q = q.Where("created_at > ?", cursor)
	}
	if err := q.Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

// MarkRead atualiza cursor LastReadAt = now (upsert).
// Reusa o mesmo modelo ConversationRead (UserID = paciente, OwnerType=patient).
func (s *PatientMessagesService) MarkRead(patientID, userID uuid.UUID) error {
	now := time.Now().UTC()
	return s.db.Exec(`
		INSERT INTO conversation_reads (id, user_id, owner_type, owner_id, last_read_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
		ON CONFLICT (user_id, owner_type, owner_id)
		DO UPDATE SET last_read_at = EXCLUDED.last_read_at, updated_at = EXCLUDED.updated_at
	`, uuid.Must(uuid.NewV7()), userID, "patient", patientID, now, now).Error
}

// ============================================================
// helpers
// ============================================================

func (s *PatientMessagesService) notifyStaff(patientID uuid.UUID, preview string) {
	if s.notification == nil {
		return
	}
	var patient models.Patient
	if err := s.db.Select("name").First(&patient, "id = ?", patientID).Error; err != nil {
		return
	}

	var staff []models.User
	_ = s.db.Where("roles::text LIKE ? OR roles::text LIKE ? OR roles::text LIKE ?",
		"%admin%", "%manager%", "%secretary%").
		Limit(20).
		Find(&staff).Error

	excerpt := preview
	if len(excerpt) > 120 {
		excerpt = excerpt[:117] + "…"
	}

	title := fmt.Sprintf("Nova mensagem de %s no portal", patient.Name)
	actionURL := fmt.Sprintf("/conversas?owner=patient&id=%s", patientID.String())
	actionText := "Abrir conversa"

	for i := range staff {
		_ = s.notification.CreateNotification(
			staff[i].ID,
			models.NotificationGeneral,
			title,
			excerpt,
			&patientID,
			nil,
			&actionURL,
			&actionText,
		)
	}
}

func toMessageView(a *models.LeadActivity) *MessageView {
	// Regra simples: se Channel=portal, foi enviada pelo paciente.
	// Outros canais (email/whatsapp/internal) representam mensagem da equipe.
	from := "team"
	if a.Channel == models.LeadChannelPortal {
		from = "patient"
	}

	authorName := ""
	if a.Actor != nil {
		authorName = a.Actor.Name
	}
	content := ""
	if a.Content != nil {
		content = *a.Content
	}

	return &MessageView{
		ID:         a.ID,
		CreatedAt:  a.CreatedAt,
		From:       from,
		AuthorName: authorName,
		Channel:    string(a.Channel),
		Content:    content,
	}
}

var _ = context.Background
