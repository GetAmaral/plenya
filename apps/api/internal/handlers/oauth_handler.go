package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/services"
)

type OAuthHandler struct {
	service  *services.OAuthService
	validate *validator.Validate
}

func NewOAuthHandler(service *services.OAuthService) *OAuthHandler {
	return &OAuthHandler{
		service:  service,
		validate: validator.New(),
	}
}

// GoogleCallback processa login via Google OAuth
// @Summary Login com Google
// @Description Autentica usuário via Google OAuth 2.0
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.GoogleOAuthRequest true "Google ID Token"
// @Success 200 {object} dto.AuthResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 409 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/auth/oauth/google [post]
func (h *OAuthHandler) GoogleCallback(c *fiber.Ctx) error {
	var req dto.GoogleOAuthRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "validation failed",
			"details": err.Error(),
		})
	}

	// Processar OAuth
	response, err := h.service.GoogleCallback(c.Context(), req.IDToken)
	if err != nil {
		// Email conflict (409)
		if err == services.ErrEmailConflict {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": "email_conflict",
				"message": err.Error(),
			})
		}

		// Invalid token (401)
		if err == services.ErrInvalidIDToken {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid_token",
				"message": "ID token inválido ou expirado",
			})
		}

		// Server error (500)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal_server_error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// AppleCallback processa login via Apple Sign In
// @Summary Login com Apple
// @Description Autentica usuário via Apple Sign In
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.AppleOAuthRequest true "Apple ID Token"
// @Success 200 {object} dto.AuthResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 409 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/auth/oauth/apple [post]
func (h *OAuthHandler) AppleCallback(c *fiber.Ctx) error {
	var req dto.AppleOAuthRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "validation failed",
			"details": err.Error(),
		})
	}

	// Extrair nome se fornecido (apenas no primeiro login da Apple)
	var userName *dto.AppleUserName
	if req.User != nil && req.User.Name != nil {
		userName = req.User.Name
	}

	// Processar OAuth
	response, err := h.service.AppleCallback(c.Context(), req.IDToken, userName)
	if err != nil {
		// Email conflict (409)
		if err == services.ErrEmailConflict {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": "email_conflict",
				"message": err.Error(),
			})
		}

		// Invalid token (401)
		if err == services.ErrInvalidIDToken {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid_token",
				"message": "ID token inválido ou expirado",
			})
		}

		// Server error (500)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "internal_server_error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
