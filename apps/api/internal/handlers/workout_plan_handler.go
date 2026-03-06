package handlers

import (
	"errors"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/middleware"
	"github.com/plenya/api/internal/services"
)

type WorkoutPlanHandler struct {
	service    *services.WorkoutPlanService
	htmlService *services.WorkoutHtmlService
	validator  *validator.Validate
}

func NewWorkoutPlanHandler(service *services.WorkoutPlanService, htmlService *services.WorkoutHtmlService) *WorkoutPlanHandler {
	return &WorkoutPlanHandler{
		service:    service,
		htmlService: htmlService,
		validator:  validator.New(),
	}
}

// Create cria um plano de treino
func (h *WorkoutPlanHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateWorkoutPlanRequest
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
	resp, err := h.service.Create(userID, &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}

// GetByID busca um plano por ID
func (h *WorkoutPlanHandler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid plan id",
			Message: "plan id must be a valid UUID",
		})
	}

	userID := middleware.GetUserID(c)
	resp, err := h.service.GetByID(id, userID)
	if err != nil {
		if errors.Is(err, services.ErrWorkoutPlanNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error: "workout plan not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.JSON(resp)
}

// GetByPublicCode busca um plano por código público (sem auth)
func (h *WorkoutPlanHandler) GetByPublicCode(c *fiber.Ctx) error {
	code := c.Params("code")
	if code == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "public code is required",
		})
	}

	resp, err := h.service.GetByPublicCode(code)
	if err != nil {
		if errors.Is(err, services.ErrPublicCodeNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error: "workout plan not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.JSON(resp)
}

// List lista planos do paciente selecionado
func (h *WorkoutPlanHandler) List(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "20"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))
	if limit > 100 {
		limit = 100
	}

	userID := middleware.GetUserID(c)
	resp, err := h.service.List(userID, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.JSON(resp)
}

// Update atualiza um plano de treino
func (h *WorkoutPlanHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid plan id",
			Message: "plan id must be a valid UUID",
		})
	}

	var req dto.UpdateWorkoutPlanRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid request body",
			Message: err.Error(),
		})
	}

	userID := middleware.GetUserID(c)
	resp, err := h.service.Update(id, userID, &req)
	if err != nil {
		if errors.Is(err, services.ErrWorkoutPlanNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error: "workout plan not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.JSON(resp)
}

// GenerateHTML gera HTML do plano com GIFs base64 inline
func (h *WorkoutPlanHandler) GenerateHTML(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid plan id",
			Message: "plan id must be a valid UUID",
		})
	}

	html, err := h.htmlService.GenerateHTML(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "html generation failed",
			Message: err.Error(),
		})
	}

	return c.JSON(fiber.Map{"html": html})
}

// Delete deleta um plano de treino
func (h *WorkoutPlanHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid plan id",
			Message: "plan id must be a valid UUID",
		})
	}

	userRole := middleware.GetPrimaryRole(c)
	err = h.service.Delete(id, userRole)
	if err != nil {
		if errors.Is(err, services.ErrWorkoutPlanNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error: "workout plan not found",
			})
		}
		if errors.Is(err, services.ErrUnauthorized) {
			return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
				Error: "unauthorized",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
