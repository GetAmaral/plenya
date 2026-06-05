package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/services"
)

// NotificationEmailHandler expõe o bucket de e-mails automáticos (aba "Notificações").
type NotificationEmailHandler struct {
	service *services.NotificationEmailService
}

func NewNotificationEmailHandler(service *services.NotificationEmailService) *NotificationEmailHandler {
	return &NotificationEmailHandler{service: service}
}

// List GET /conversations/notifications?unread_only=&limit=
func (h *NotificationEmailHandler) List(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "100"))
	unreadOnly := c.Query("unread_only") == "true"

	items, err := h.service.List(limit, unreadOnly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: "failed to load notifications", Message: err.Error(),
		})
	}
	unread, err := h.service.UnreadCount()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: "failed to count notifications", Message: err.Error(),
		})
	}
	return c.JSON(fiber.Map{"items": items, "unreadCount": unread})
}

// MarkRead POST /conversations/notifications/:id/read
func (h *NotificationEmailHandler) MarkRead(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid id"})
	}
	if err := h.service.MarkRead(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: "failed to mark read", Message: err.Error(),
		})
	}
	return c.JSON(fiber.Map{"ok": true})
}
