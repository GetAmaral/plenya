package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/services"
)

type WebPushHandler struct {
	service *services.WebPushService
	notif   *services.NotificationService
}

func NewWebPushHandler(service *services.WebPushService, notif *services.NotificationService) *WebPushHandler {
	return &WebPushHandler{service: service, notif: notif}
}

// VAPIDPublicKey godoc
// @Summary Chave pública VAPID
// @Description Devolve a chave pública VAPID usada pelo navegador para se inscrever em Web Push
// @Tags WebPush
// @Produce json
// @Success 200 {object} map[string]any
// @Router /api/v1/web-push/vapid-public-key [get]
// @Security BearerAuth
func (h *WebPushHandler) VAPIDPublicKey(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"publicKey": h.service.PublicKey(),
		"enabled":   h.service.Enabled(),
	})
}

type webPushSubscribeRequest struct {
	Endpoint string `json:"endpoint"`
	Keys     struct {
		P256dh string `json:"p256dh"`
		Auth   string `json:"auth"`
	} `json:"keys"`
	DeviceLabel string `json:"deviceLabel"`
}

// Subscribe godoc
// @Summary Inscrever navegador em Web Push
// @Description Registra a PushSubscription do navegador para o usuário autenticado
// @Tags WebPush
// @Accept json
// @Produce json
// @Param body body webPushSubscribeRequest true "PushSubscription"
// @Success 200 {object} map[string]any
// @Failure 400 {object} map[string]string
// @Router /api/v1/web-push/subscribe [post]
// @Security BearerAuth
func (h *WebPushHandler) Subscribe(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	var req webPushSubscribeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "corpo inválido"})
	}

	sub, err := h.service.Subscribe(userID, services.WebPushSubscribeInput{
		Endpoint: req.Endpoint,
		P256dh:   req.Keys.P256dh,
		Auth:     req.Keys.Auth,
		Label:    req.DeviceLabel,
		UA:       c.Get("User-Agent"),
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"id": sub.ID, "ok": true})
}

type webPushUnsubscribeRequest struct {
	Endpoint string `json:"endpoint"`
}

// Unsubscribe godoc
// @Summary Cancelar inscrição Web Push
// @Description Remove a inscrição de um endpoint (ao desativar avisos no navegador)
// @Tags WebPush
// @Accept json
// @Produce json
// @Param body body webPushUnsubscribeRequest true "Endpoint"
// @Success 204
// @Router /api/v1/web-push/unsubscribe [post]
// @Security BearerAuth
func (h *WebPushHandler) Unsubscribe(c *fiber.Ctx) error {
	var req webPushUnsubscribeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "corpo inválido"})
	}
	if err := h.service.Unsubscribe(req.Endpoint); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "falha ao cancelar"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// Test godoc
// @Summary Enviar push de teste
// @Description Dispara uma notificação Web Push de teste para o próprio usuário (verificação)
// @Tags WebPush
// @Produce json
// @Success 200 {object} map[string]any
// @Router /api/v1/web-push/test [post]
// @Security BearerAuth
func (h *WebPushHandler) Test(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	err := h.service.Send(userID, services.PushPayload{
		Title: "Plenya — avisos ativados",
		Body:  "Se você está vendo isto, as notificações web estão funcionando neste aparelho.",
		URL:   "/conversas",
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"ok": true})
}
