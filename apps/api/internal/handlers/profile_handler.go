package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/middleware"
	"github.com/plenya/api/internal/services"
)

type ProfileHandler struct {
	service *services.ProfileService
}

func NewProfileHandler(service *services.ProfileService) *ProfileHandler {
	return &ProfileHandler{service: service}
}

// GetProfile retorna o perfil do usuário logado
// @Summary Get current user profile
// @Tags profile
// @Security BearerAuth
// @Success 200 {object} dto.UserResponse
// @Router /api/v1/profile [get]
func (h *ProfileHandler) GetProfile(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)

	user, err := h.service.GetProfile(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get profile",
		})
	}

	return c.JSON(user)
}

// UpdateProfile atualiza o perfil do usuário logado
// @Summary Update current user profile
// @Tags profile
// @Security BearerAuth
// @Param body body dto.UpdateProfileRequest true "Profile data"
// @Success 200 {object} dto.UserResponse
// @Router /api/v1/profile [put]
func (h *ProfileHandler) UpdateProfile(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)

	var req dto.UpdateProfileRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	user, err := h.service.UpdateProfile(userID, &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(user)
}
