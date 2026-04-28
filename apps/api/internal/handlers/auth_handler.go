package handlers

import (
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/middleware"
	"github.com/plenya/api/internal/services"
)

type AuthHandler struct {
	authService *services.AuthService
	validator   *validator.Validate
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		validator:   validator.New(),
	}
}

// Register godoc
// @Summary Registrar novo usuário
// @Description Cria um novo usuário no sistema
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "Dados de registro"
// @Success 201 {object} dto.AuthResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 409 {object} dto.ErrorResponse
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req dto.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid request body",
			Message: err.Error(),
		})
	}

	// Validar request
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "validation failed",
			Details: formatValidationErrors(err),
		})
	}

	// Registrar usuário
	resp, err := h.authService.Register(&req)
	if err != nil {
		if errors.Is(err, services.ErrUserAlreadyExists) {
			return c.Status(fiber.StatusConflict).JSON(dto.ErrorResponse{
				Error:   "user already exists",
				Message: "a user with this email already exists",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}

// Login godoc
// @Summary Login
// @Description Autentica um usuário e retorna tokens JWT
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Credenciais"
// @Success 200 {object} dto.AuthResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid request body",
			Message: err.Error(),
		})
	}

	// Validar request
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "validation failed",
			Details: formatValidationErrors(err),
		})
	}

	// Login (pode retornar MFA challenge em vez de tokens)
	attempt, err := h.authService.Login(&req)
	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) {
			// Audit: login failure (sem userID válido — log por email + IP)
			log.Printf("[AUDIT] auth.login.failure email=%s ip=%s ua=%s", req.Email, c.IP(), c.Get("User-Agent"))
			return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{
				Error:   "invalid credentials",
				Message: "email or password is incorrect",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	if attempt.MFARequired {
		log.Printf("[AUDIT] auth.login.mfa_challenge email=%s ip=%s", req.Email, c.IP())
		return c.JSON(fiber.Map{
			"mfaRequired": true,
			"mfaToken":    attempt.MFAToken,
		})
	}

	log.Printf("[AUDIT] auth.login.success email=%s ip=%s ua=%s", req.Email, c.IP(), c.Get("User-Agent"))
	resp := *attempt.Auth
	if attempt.MFAEnrollmentMissing {
		// Compartilhamos via fiber.Map pra preservar struct existente
		out := fiber.Map{
			"accessToken":            resp.AccessToken,
			"refreshToken":           resp.RefreshToken,
			"user":                   resp.User,
			"mfaEnrollmentRequired":  true,
		}
		return c.JSON(out)
	}
	return c.JSON(resp)
}

// VerifyMFARequest — payload do POST /auth/login/verify-2fa
type VerifyMFARequest struct {
	MFAToken string `json:"mfaToken" validate:"required"`
	Code     string `json:"code" validate:"required,len=6"`
}

// VerifyMFA godoc
// @Summary Validar código 2FA após login
// @Description Recebe MFA challenge token + código TOTP, retorna tokens de auth
// @Tags auth
// @Accept json
// @Produce json
// @Param request body VerifyMFARequest true "MFA challenge + code"
// @Success 200 {object} dto.AuthResponse
// @Failure 401 {object} dto.ErrorResponse
// @Router /auth/login/verify-2fa [post]
func (h *AuthHandler) VerifyMFA(c *fiber.Ctx) error {
	var req VerifyMFARequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "invalid request body", Message: err.Error(),
		})
	}
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "validation failed", Details: formatValidationErrors(err),
		})
	}
	resp, err := h.authService.VerifyMFAChallenge(req.MFAToken, req.Code)
	if err != nil {
		log.Printf("[AUDIT] auth.login.mfa_failure ip=%s", c.IP())
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{
			Error: "invalid mfa", Message: "código inválido ou challenge expirado",
		})
	}
	log.Printf("[AUDIT] auth.login.mfa_success ip=%s", c.IP())
	return c.JSON(resp)
}

// Enable2FA godoc
// @Summary Inicia enrollment de 2FA — gera secret + provisioning URL
// @Tags auth
// @Security BearerAuth
// @Success 200 {object} services.Enable2FAResult
// @Router /auth/2fa/enable [post]
func (h *AuthHandler) Enable2FA(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)
	out, err := h.authService.StartEnable2FA(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: "failed to start 2fa enrollment", Message: err.Error(),
		})
	}
	return c.JSON(out)
}

// Verify2FARequest — confirma o primeiro código TOTP, habilita 2FA.
type Verify2FARequest struct {
	Code string `json:"code" validate:"required,len=6"`
}

// Verify2FA godoc
// @Summary Confirma enrollment 2FA — valida primeiro código e habilita
// @Tags auth
// @Security BearerAuth
// @Param request body Verify2FARequest true "Código TOTP de 6 dígitos"
// @Success 204
// @Router /auth/2fa/verify [post]
func (h *AuthHandler) Verify2FA(c *fiber.Ctx) error {
	var req Verify2FARequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "invalid request body", Message: err.Error(),
		})
	}
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "validation failed", Details: formatValidationErrors(err),
		})
	}
	userID := middleware.GetUserID(c)
	if err := h.authService.ConfirmEnable2FA(userID, req.Code); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{
			Error: "invalid code", Message: "código TOTP inválido",
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// Refresh godoc
// @Summary Refresh token
// @Description Gera um novo access token a partir do refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.RefreshRequest true "Refresh token"
// @Success 200 {object} dto.AuthResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Router /auth/refresh [post]
func (h *AuthHandler) Refresh(c *fiber.Ctx) error {
	var req dto.RefreshRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid request body",
			Message: err.Error(),
		})
	}

	// Validar request
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "validation failed",
			Details: formatValidationErrors(err),
		})
	}

	// Refresh token
	resp, err := h.authService.RefreshToken(req.RefreshToken)
	if err != nil {
		if errors.Is(err, services.ErrInvalidToken) {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{
				Error:   "invalid token",
				Message: "refresh token is invalid or expired",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.JSON(resp)
}

// Logout godoc
// @Summary Logout
// @Description Revoga TODOS os refresh tokens ativos do user (força re-login em todos devices)
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Router /auth/logout [post]
func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)
	if err := h.authService.Logout(userID); err != nil {
		log.Printf("[AUDIT] auth.logout.error userId=%s err=%v", userID, err)
		// Não falha o request — logout deve ser idempotente do POV do client
	}
	log.Printf("[AUDIT] auth.logout.success userId=%s ip=%s", userID, c.IP())
	return c.JSON(fiber.Map{
		"message": "logged out successfully",
	})
}

// GetMe godoc
// @Summary Get current user
// @Description Retorna informações do usuário autenticado incluindo paciente selecionado
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.UserDTO
// @Failure 401 {object} dto.ErrorResponse
// @Router /users/me [get]
func (h *AuthHandler) GetMe(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)

	user, err := h.authService.GetUserByID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.JSON(user)
}

// UpdateSelectedPatient godoc
// @Summary Update selected patient
// @Description Atualiza o paciente selecionado para o usuário autenticado
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.UpdateSelectedPatientRequest true "Patient ID"
// @Success 200 {object} dto.UserDTO
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /users/me/selected-patient [put]
func (h *AuthHandler) UpdateSelectedPatient(c *fiber.Ctx) error {
	var req dto.UpdateSelectedPatientRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid request body",
			Message: err.Error(),
		})
	}

	// Validar request
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "validation failed",
			Details: formatValidationErrors(err),
		})
	}

	// Parse patient ID
	patientID, err := uuid.Parse(req.PatientID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid patient id",
			Message: "patient id must be a valid UUID",
		})
	}

	userID := middleware.GetUserID(c)

	user, err := h.authService.UpdateSelectedPatient(userID, patientID)
	if err != nil {
		// Log the error for debugging
		log.Printf("[ERROR] UpdateSelectedPatient failed: %v", err)

		if errors.Is(err, services.ErrPatientNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "patient not found",
				Message: "no patient found with the given id",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.JSON(user)
}

// UpdatePreferences godoc
// @Summary Update user preferences
// @Description Atualiza as preferências do usuário autenticado (viewport do mindmap, etc.)
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body map[string]interface{} true "Preferences JSON"
// @Success 200 {object} dto.UserDTO
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /users/me/preferences [patch]
func (h *AuthHandler) UpdatePreferences(c *fiber.Ctx) error {
	var preferences map[string]interface{}
	if err := c.BodyParser(&preferences); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid request body",
			Message: err.Error(),
		})
	}

	userID := middleware.GetUserID(c)

	user, err := h.authService.UpdatePreferences(userID, preferences)
	if err != nil {
		log.Printf("[ERROR] UpdatePreferences failed: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.JSON(user)
}

// formatValidationErrors formata erros de validação
func formatValidationErrors(err error) map[string]string {
	errors := make(map[string]string)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			errors[fieldError.Field()] = fieldError.Tag()
		}
	}
	return errors
}
