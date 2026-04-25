package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/plenya/api/internal/services"
)

// MobileConfigHandler serve config pública usada pelos apps mobile no cold start
// (versão mínima, kill switch, SSL pins, feature flags).
type MobileConfigHandler struct {
	service *services.MobileConfigService
}

// NewMobileConfigHandler instancia o handler.
func NewMobileConfigHandler(service *services.MobileConfigService) *MobileConfigHandler {
	return &MobileConfigHandler{service: service}
}

// Get godoc
// @Summary Config pública dos apps mobile
// @Description Consultado em cold start; sem auth. Determina kill switch, SSL pins e versão mínima suportada.
// @Tags mobile
// @Produce json
// @Success 200 {object} services.MobileConfig
// @Router /mobile/config [get]
func (h *MobileConfigHandler) Get(c *fiber.Ctx) error {
	return c.JSON(h.service.Get())
}
