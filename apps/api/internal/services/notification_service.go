package services

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/repository"
)

// PushSender é a interface mínima que NotificationService usa pra disparar
// push em paralelo ao registro in-app. Permite injetar PushService sem ciclo
// de import.
type PushSender interface {
	Send(userID uuid.UUID, payload PushPayload) error
}

type NotificationService struct {
	repo             *repository.NotificationRepository
	subscriptionRepo *repository.SubscriptionRepository
	push             PushSender
}

func NewNotificationService(repo *repository.NotificationRepository, subscriptionRepo *repository.SubscriptionRepository) *NotificationService {
	return &NotificationService{
		repo:             repo,
		subscriptionRepo: subscriptionRepo,
	}
}

// SetPushSender injeta o sender pós-construção (resolve ordem de wiring em main.go).
func (s *NotificationService) SetPushSender(sender PushSender) {
	s.push = sender
}

// dispatchPush é fire-and-forget: roda em goroutine pra não bloquear o
// fluxo da chamada original (criar notif in-app deve continuar mesmo se
// a Expo Push API estiver lenta/offline).
func (s *NotificationService) dispatchPush(n *models.Notification) {
	if s.push == nil {
		return
	}
	url := ""
	if n.ActionURL != nil {
		url = *n.ActionURL
	}
	go func() {
		_ = s.push.Send(n.UserID, PushPayload{
			Title: n.Title,
			Body:  SanitizeBody(n.Message),
			URL:   url,
		})
	}()
}

// planSnapshotData extracts plan data from snapshot JSON
type planSnapshotData struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// getPlanDataFromSnapshot extracts plan name and price from subscription's plan snapshot
func getPlanDataFromSnapshot(subscription *models.PatientSubscription) (string, float64) {
	var snapshot planSnapshotData
	if err := json.Unmarshal([]byte(subscription.PlanSnapshot), &snapshot); err != nil {
		// Fallback values if parsing fails
		return "Plano", 0.0
	}
	return snapshot.Name, snapshot.Price
}

// GetByUserID busca notificações de um usuário
func (s *NotificationService) GetByUserID(userID string, limit int) ([]models.Notification, error) {
	return s.repo.GetByUserID(userID, limit)
}

// GetUnreadByUserID busca notificações não lidas
func (s *NotificationService) GetUnreadByUserID(userID string) ([]models.Notification, error) {
	return s.repo.GetUnreadByUserID(userID)
}

// CountUnreadByUserID conta notificações não lidas
func (s *NotificationService) CountUnreadByUserID(userID string) (int64, error) {
	return s.repo.CountUnreadByUserID(userID)
}

// MarkAsRead marca como lida
func (s *NotificationService) MarkAsRead(id string) error {
	return s.repo.MarkAsRead(id)
}

// MarkAllAsRead marca todas como lidas
func (s *NotificationService) MarkAllAsRead(userID string) error {
	return s.repo.MarkAllAsReadByUserID(userID)
}

// Delete deleta uma notificação
func (s *NotificationService) Delete(id string) error {
	return s.repo.Delete(id)
}

// DeleteAll deleta todas as notificações de um usuário
func (s *NotificationService) DeleteAll(userID string) error {
	return s.repo.DeleteAllByUserID(userID)
}

// CreateNotification cria uma notificação genérica
func (s *NotificationService) CreateNotification(
	userID uuid.UUID,
	notifType models.NotificationType,
	title, message string,
	patientID, subscriptionID *uuid.UUID,
	actionURL, actionText *string,
) error {
	notification := &models.Notification{
		UserID:         userID,
		PatientID:      patientID,
		SubscriptionID: subscriptionID,
		Type:           notifType,
		Title:          title,
		Message:        message,
		ActionURL:      actionURL,
		ActionText:     actionText,
		Read:           false,
	}

	if err := s.repo.Create(notification); err != nil {
		return err
	}
	s.dispatchPush(notification)
	return nil
}

// CreateLeadNotification cria notificação in-app vinculada a Lead.
// Usada pelos triggers do CRM (lead novo, inbound WA, atribuição).
//
// DEPRECATED: usar CreateConversationNotification, que aceita owner Lead OU Patient.
// Mantida pra não quebrar callers legados.
func (s *NotificationService) CreateLeadNotification(
	userID uuid.UUID,
	leadID uuid.UUID,
	notifType models.NotificationType,
	title, message, actionURL string,
) error {
	// E-mails inbound não geram mais notificação (in-app nem push): o aviso fica só
	// no badge de não-lidas do link "E-mails" no navbar. Decisão do usuário (2026-06).
	if notifType == models.NotificationLeadEmailInbound {
		return nil
	}
	actionText := "Abrir lead"
	notification := &models.Notification{
		UserID:     userID,
		LeadID:     &leadID,
		Type:       notifType,
		Title:      title,
		Message:    message,
		ActionURL:  &actionURL,
		ActionText: &actionText,
		Read:       false,
	}
	if err := s.repo.Create(notification); err != nil {
		return err
	}
	s.dispatchPush(notification)
	return nil
}

// CreateConversationNotification cria notificação in-app vinculada a uma conversa
// que pode pertencer a Lead OU Patient. Substitui CreateLeadNotification pra fluxos
// pós-Bloco B (Central de Conversas) que precisam alcançar Patient sem gambiarras.
//
// Exatamente UM de leadID|patientID deve ser não-nulo. Se ambos forem nulos ou ambos
// setados, cria notificação sem vínculo (warn-friendly degradation, não panic).
//
// actionText é "Abrir conversa" para refletir a UX da Central (Bloco C).
func (s *NotificationService) CreateConversationNotification(
	userID uuid.UUID,
	leadID *uuid.UUID,
	patientID *uuid.UUID,
	notifType models.NotificationType,
	title, message, actionURL string,
) error {
	// E-mails inbound não geram mais notificação (in-app nem push): só o badge de
	// não-lidas do link "E-mails" no navbar. WhatsApp segue notificando. (2026-06)
	if notifType == models.NotificationLeadEmailInbound {
		return nil
	}
	actionText := "Abrir conversa"

	// Dedupe por contato+canal: em vez de uma notificação por mensagem (que empilha sem fim),
	// atualizamos a notificação existente do mesmo (usuário, lead/patient, tipo) com a última
	// mensagem, reabrindo como não lida e trazendo pro topo. Mantém UMA notificação por contato.
	if existing, ferr := s.repo.FindLatestConversationNotification(userID, leadID, patientID, notifType); ferr == nil && existing != nil {
		existing.Title = title
		existing.Message = message
		existing.ActionURL = &actionURL
		existing.ActionText = &actionText
		existing.Read = false
		existing.ReadAt = nil
		existing.CreatedAt = time.Now()
		if err := s.repo.Update(existing); err != nil {
			return err
		}
		s.dispatchPush(existing)
		return nil
	}

	notification := &models.Notification{
		UserID:     userID,
		Type:       notifType,
		Title:      title,
		Message:    message,
		ActionURL:  &actionURL,
		ActionText: &actionText,
		Read:       false,
	}
	if leadID != nil && *leadID != uuid.Nil {
		notification.LeadID = leadID
	}
	if patientID != nil && *patientID != uuid.Nil {
		notification.PatientID = patientID
	}
	if err := s.repo.Create(notification); err != nil {
		return err
	}
	s.dispatchPush(notification)
	return nil
}

// MarkReadByTypes marca como lidas as notificações não lidas do usuário cujos tipos estão na
// lista. Usado ao entrar numa superfície de conversa (ex.: a tela de WhatsApp limpa as de WA).
func (s *NotificationService) MarkReadByTypes(userID string, types []models.NotificationType) error {
	strs := make([]string, 0, len(types))
	for _, t := range types {
		strs = append(strs, string(t))
	}
	return s.repo.MarkReadByTypes(userID, strs)
}

// CheckAndCreateSubscriptionNotifications verifica subscriptions e cria notificações
func (s *NotificationService) CheckAndCreateSubscriptionNotifications() error {
	// 1. Check trial expiring (7 dias)
	if err := s.checkTrialExpiring(7); err != nil {
		return err
	}

	// 2. Check renewal upcoming (3 dias)
	if err := s.checkRenewalUpcoming(3); err != nil {
		return err
	}

	// 3. Check expired subscriptions
	if err := s.checkExpired(); err != nil {
		return err
	}

	return nil
}

// checkTrialExpiring verifica trials expirando
func (s *NotificationService) checkTrialExpiring(daysUntilExpiry int) error {
	subscriptions, err := s.subscriptionRepo.GetExpiringTrials(daysUntilExpiry)
	if err != nil {
		return err
	}

	for _, subscription := range subscriptions {
		// Check if notification already exists for today
		exists, err := s.notificationExistsToday(
			subscription.PatientID,
			subscription.ID,
			models.NotificationTrialExpiring,
		)
		if err != nil || exists {
			continue
		}

		daysLeft := int(time.Until(*subscription.TrialEndDate).Hours() / 24)

		planName, _ := getPlanDataFromSnapshot(&subscription)
		title := fmt.Sprintf("Trial expirando em %d dias", daysLeft)
		message := fmt.Sprintf(
			"Seu período trial do plano '%s' expira em %d dias. Renove agora para continuar aproveitando os benefícios.",
			planName,
			daysLeft,
		)
		actionURL := fmt.Sprintf("/patients/%s/subscriptions", subscription.PatientID.String())
		actionText := "Ver Subscription"

		// Get patient's user_id (assuming patient has a user_id)
		// You'll need to query this from the patient table
		// For now, using a placeholder - you should fetch the actual user_id
		userID := subscription.PatientID // Placeholder - should be patient.UserID

		if err := s.CreateNotification(
			userID,
			models.NotificationTrialExpiring,
			title,
			message,
			&subscription.PatientID,
			&subscription.ID,
			&actionURL,
			&actionText,
		); err != nil {
			continue // Log error but continue with others
		}
	}

	return nil
}

// checkRenewalUpcoming verifica renovações próximas
func (s *NotificationService) checkRenewalUpcoming(daysUntilRenewal int) error {
	subscriptions, err := s.subscriptionRepo.GetUpcomingRenewals(daysUntilRenewal)
	if err != nil {
		return err
	}

	for _, subscription := range subscriptions {
		exists, err := s.notificationExistsToday(
			subscription.PatientID,
			subscription.ID,
			models.NotificationRenewalUpcoming,
		)
		if err != nil || exists {
			continue
		}

		daysLeft := int(time.Until(*subscription.NextBillingDate).Hours() / 24)

		planName, planPrice := getPlanDataFromSnapshot(&subscription)
		// Calculate final price: use custom_price if set, otherwise plan price with discount
		finalPrice := planPrice
		if subscription.CustomPrice != nil {
			finalPrice = *subscription.CustomPrice
		} else if subscription.DiscountPercent > 0 {
			finalPrice = planPrice * (1 - subscription.DiscountPercent/100)
		}

		title := fmt.Sprintf("Renovação em %d dias", daysLeft)
		message := fmt.Sprintf(
			"Seu subscription '%s' será renovado automaticamente em %d dias. Valor: R$ %.2f",
			planName,
			daysLeft,
			finalPrice,
		)
		actionURL := fmt.Sprintf("/patients/%s/subscriptions", subscription.PatientID.String())
		actionText := "Ver Detalhes"

		userID := subscription.PatientID

		if err := s.CreateNotification(
			userID,
			models.NotificationRenewalUpcoming,
			title,
			message,
			&subscription.PatientID,
			&subscription.ID,
			&actionURL,
			&actionText,
		); err != nil {
			continue
		}
	}

	return nil
}

// checkExpired verifica subscriptions expirados
func (s *NotificationService) checkExpired() error {
	subscriptions, err := s.subscriptionRepo.GetExpired()
	if err != nil {
		return err
	}

	for _, subscription := range subscriptions {
		exists, err := s.notificationExistsToday(
			subscription.PatientID,
			subscription.ID,
			models.NotificationSubscriptionExpired,
		)
		if err != nil || exists {
			continue
		}

		planName, _ := getPlanDataFromSnapshot(&subscription)
		title := "Subscription Expirado"
		message := fmt.Sprintf(
			"Seu subscription '%s' expirou. Renove agora para continuar com o acompanhamento.",
			planName,
		)
		actionURL := fmt.Sprintf("/patients/%s/subscriptions", subscription.PatientID.String())
		actionText := "Renovar Agora"

		userID := subscription.PatientID

		if err := s.CreateNotification(
			userID,
			models.NotificationSubscriptionExpired,
			title,
			message,
			&subscription.PatientID,
			&subscription.ID,
			&actionURL,
			&actionText,
		); err != nil {
			continue
		}
	}

	return nil
}

// notificationExistsToday verifica se já existe notificação hoje
func (s *NotificationService) notificationExistsToday(
	patientID uuid.UUID,
	subscriptionID uuid.UUID,
	notifType models.NotificationType,
) (bool, error) {
	// This would require a new repository method
	// For now, return false to always create notifications
	// You should implement proper duplicate checking
	return false, nil
}
