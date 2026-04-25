// Package handlers — TelemedLobbyHandler expõe o lobby standalone:
//
//   GET /api/v1/sala/:token        — retorna LobbyView (público, sem auth)
//   POST /api/v1/sala/:token/used  — marca primeira entrada (best-effort)
package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"github.com/plenya/api/internal/services"
)

type TelemedLobbyHandler struct {
	svc *services.TelemedLobbyService
}

func NewTelemedLobbyHandler(svc *services.TelemedLobbyService) *TelemedLobbyHandler {
	return &TelemedLobbyHandler{svc: svc}
}

func (h *TelemedLobbyHandler) Resolve(c *fiber.Ctx) error {
	token := c.Params("token")
	view, err := h.svc.Resolve(token)
	if err != nil {
		if errors.Is(err, services.ErrLobbyTokenNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "link não encontrado"})
		}
		if errors.Is(err, services.ErrLobbyTokenExpired) {
			return c.Status(fiber.StatusGone).JSON(fiber.Map{"error": "link expirado"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(view)
}

func (h *TelemedLobbyHandler) MarkUsed(c *fiber.Ctx) error {
	token := c.Params("token")
	_ = h.svc.MarkUsed(token)
	return c.SendStatus(fiber.StatusNoContent)
}
