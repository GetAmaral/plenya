package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/plenya/api/internal/services"
)

type NotificationHandler struct {
	service *services.NotificationService
}

func NewNotificationHandler(service *services.NotificationService) *NotificationHandler {
	return &NotificationHandler{service: service}
}

// GetUserNotifications godoc
// @Summary List user notifications
// @Description Returns all notifications for the authenticated user
// @Tags Notifications
// @Produce json
// @Param limit query int false "Limit results (default: 50)"
// @Success 200 {array} models.Notification
// @Failure 500 {object} map[string]string
// @Router /api/v1/notifications [get]
// @Security BearerAuth
func (h *NotificationHandler) GetUserNotifications(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID).String()

	limit := c.QueryInt("limit", 50)

	notifications, err := h.service.GetByUserID(userID, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch notifications",
		})
	}

	return c.JSON(notifications)
}

// GetUnreadNotifications godoc
// @Summary List unread notifications
// @Description Returns unread notifications for the authenticated user
// @Tags Notifications
// @Produce json
// @Success 200 {array} models.Notification
// @Failure 500 {object} map[string]string
// @Router /api/v1/notifications/unread [get]
// @Security BearerAuth
func (h *NotificationHandler) GetUnreadNotifications(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID).String()

	notifications, err := h.service.GetUnreadByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch unread notifications",
		})
	}

	return c.JSON(notifications)
}

// GetUnreadCount godoc
// @Summary Get unread count
// @Description Returns the count of unread notifications
// @Tags Notifications
// @Produce json
// @Success 200 {object} map[string]int64
// @Failure 500 {object} map[string]string
// @Router /api/v1/notifications/unread/count [get]
// @Security BearerAuth
func (h *NotificationHandler) GetUnreadCount(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID).String()

	count, err := h.service.CountUnreadByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to count unread notifications",
		})
	}

	return c.JSON(fiber.Map{
		"count": count,
	})
}

// MarkAsRead godoc
// @Summary Mark notification as read
// @Description Marks a notification as read
// @Tags Notifications
// @Param id path string true "Notification ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/notifications/{id}/read [post]
// @Security BearerAuth
func (h *NotificationHandler) MarkAsRead(c *fiber.Ctx) error {
	notificationID := c.Params("id")

	if err := h.service.MarkAsRead(notificationID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to mark notification as read",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// MarkAllAsRead godoc
// @Summary Mark all notifications as read
// @Description Marks all user notifications as read
// @Tags Notifications
// @Success 204
// @Failure 500 {object} map[string]string
// @Router /api/v1/notifications/read-all [post]
// @Security BearerAuth
func (h *NotificationHandler) MarkAllAsRead(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID).String()

	if err := h.service.MarkAllAsRead(userID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to mark all notifications as read",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// DeleteNotification godoc
// @Summary Delete notification
// @Description Soft deletes a notification
// @Tags Notifications
// @Param id path string true "Notification ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/notifications/{id} [delete]
// @Security BearerAuth
func (h *NotificationHandler) DeleteNotification(c *fiber.Ctx) error {
	notificationID := c.Params("id")

	if err := h.service.Delete(notificationID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete notification",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// DeleteAllNotifications godoc
// @Summary Delete all notifications
// @Description Deletes all user notifications
// @Tags Notifications
// @Success 204
// @Failure 500 {object} map[string]string
// @Router /api/v1/notifications [delete]
// @Security BearerAuth
func (h *NotificationHandler) DeleteAllNotifications(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID).String()

	if err := h.service.DeleteAll(userID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete all notifications",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
