package handlers

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/middleware"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
)

// MeMobileHandler agrupa endpoints sob /api/v1/me/* relacionados ao app mobile:
// device tokens, sessões ativas e consentimento LGPD.
type MeMobileHandler struct {
	deviceTokens *services.DeviceTokenService
	consent      *services.LGPDConsentService
	validator    *validator.Validate
}

// NewMeMobileHandler instancia o handler.
func NewMeMobileHandler(
	deviceTokens *services.DeviceTokenService,
	consent *services.LGPDConsentService,
) *MeMobileHandler {
	return &MeMobileHandler{
		deviceTokens: deviceTokens,
		consent:      consent,
		validator:    validator.New(),
	}
}

// RegisterDeviceTokenRequest é o body de POST /me/device-tokens.
type RegisterDeviceTokenRequest struct {
	Platform    string  `json:"platform" validate:"required,oneof=ios android"`
	AppVariant  string  `json:"appVariant" validate:"required,oneof=pro app"`
	Token       string  `json:"token" validate:"required,max=512"`
	DeviceLabel *string `json:"deviceLabel,omitempty" validate:"omitempty,max=200"`
	AppVersion  *string `json:"appVersion,omitempty" validate:"omitempty,max=40"`
}

// SessionDTO é a forma como expomos um device token em /me/sessions.
type SessionDTO struct {
	ID         uuid.UUID `json:"id"`
	Platform   string    `json:"platform"`
	AppVariant string    `json:"appVariant"`
	Device     string    `json:"device"`
	AppVersion string    `json:"appVersion,omitempty"`
	LastSeenAt string    `json:"lastSeenAt"`
	Current    bool      `json:"current"`
}

// RegisterDeviceToken godoc
// @Summary Registrar device token para push notifications
// @Description Cria ou atualiza (idempotente) o token Expo/FCM/APNs do device autenticado
// @Tags mobile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body handlers.RegisterDeviceTokenRequest true "Token info"
// @Success 201 {object} models.DeviceToken
// @Failure 400 {object} dto.ErrorResponse
// @Router /me/device-tokens [post]
func (h *MeMobileHandler) RegisterDeviceToken(c *fiber.Ctx) error {
	var req RegisterDeviceTokenRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid request body",
			Message: err.Error(),
		})
	}

	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "validation failed",
			Details: formatValidationErrors(err),
		})
	}

	userID := middleware.GetUserID(c)

	dt, err := h.deviceTokens.Register(services.RegisterTokenInput{
		UserID:      userID,
		Platform:    models.DevicePlatform(req.Platform),
		AppVariant:  models.AppVariant(req.AppVariant),
		Token:       req.Token,
		DeviceLabel: req.DeviceLabel,
		AppVersion:  req.AppVersion,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(dt)
}

// RemoveDeviceToken godoc
// @Summary Revogar device token
// @Description Remove o device token (logout do device)
// @Tags mobile
// @Security BearerAuth
// @Param id path string true "Token ID"
// @Success 204
// @Failure 404 {object} dto.ErrorResponse
// @Router /me/device-tokens/{id} [delete]
func (h *MeMobileHandler) RemoveDeviceToken(c *fiber.Ctx) error {
	tokenID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid id",
			Message: "id must be a valid UUID",
		})
	}

	userID := middleware.GetUserID(c)
	if err := h.deviceTokens.Revoke(userID, tokenID); err != nil {
		if errors.Is(err, services.ErrDeviceTokenNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error: "device token not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// ListSessions godoc
// @Summary Listar sessões/devices ativos do usuário
// @Description Cada device token é considerado uma sessão. Útil para LGPD (visibilidade) e segurança (revogação remota)
// @Tags mobile
// @Produce json
// @Security BearerAuth
// @Success 200 {array} handlers.SessionDTO
// @Router /me/sessions [get]
func (h *MeMobileHandler) ListSessions(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)

	tokens, err := h.deviceTokens.ListByUser(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	currentToken := c.Get("X-Device-Token-Id")

	sessions := make([]SessionDTO, 0, len(tokens))
	for _, t := range tokens {
		device := ""
		if t.DeviceLabel != nil {
			device = *t.DeviceLabel
		}
		appVersion := ""
		if t.AppVersion != nil {
			appVersion = *t.AppVersion
		}
		sessions = append(sessions, SessionDTO{
			ID:         t.ID,
			Platform:   string(t.Platform),
			AppVariant: string(t.AppVariant),
			Device:     device,
			AppVersion: appVersion,
			LastSeenAt: t.LastSeenAt.UTC().Format("2006-01-02T15:04:05Z07:00"),
			Current:    currentToken != "" && currentToken == t.ID.String(),
		})
	}

	return c.JSON(sessions)
}

// RevokeSession godoc
// @Summary Revogar sessão específica
// @Description Mesma operação que DELETE /me/device-tokens/:id (alias semântico para "logout deste device")
// @Tags mobile
// @Security BearerAuth
// @Param id path string true "Session/Token ID"
// @Success 204
// @Router /me/sessions/{id} [delete]
func (h *MeMobileHandler) RevokeSession(c *fiber.Ctx) error {
	return h.RemoveDeviceToken(c)
}

// GetLGPDConsent godoc
// @Summary Status do consentimento LGPD do usuário
// @Tags mobile
// @Produce json
// @Security BearerAuth
// @Success 200 {object} services.LGPDConsentStatus
// @Router /me/consent/lgpd [get]
func (h *MeMobileHandler) GetLGPDConsent(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)

	status, err := h.consent.Status(userID)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error: "user not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}
	return c.JSON(status)
}

// AcceptLGPDConsent godoc
// @Summary Aceitar termo LGPD
// @Description Registra o aceite. Idempotente — chamadas subsequentes mantêm timestamp original.
// @Tags mobile
// @Produce json
// @Security BearerAuth
// @Success 200 {object} services.LGPDConsentStatus
// @Router /me/consent/lgpd [post]
func (h *MeMobileHandler) AcceptLGPDConsent(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)

	status, err := h.consent.Accept(userID)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error: "user not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}
	return c.JSON(status)
}
