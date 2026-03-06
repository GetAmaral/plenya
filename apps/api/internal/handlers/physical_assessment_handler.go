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

type PhysicalAssessmentHandler struct {
	service   *services.PhysicalAssessmentService
	validator *validator.Validate
}

func NewPhysicalAssessmentHandler(service *services.PhysicalAssessmentService) *PhysicalAssessmentHandler {
	return &PhysicalAssessmentHandler{
		service:   service,
		validator: validator.New(),
	}
}

// Create cria uma avaliação física
func (h *PhysicalAssessmentHandler) Create(c *fiber.Ctx) error {
	var req dto.CreatePhysicalAssessmentRequest
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
		if errors.Is(err, services.ErrPatientNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error: "patient not found",
			})
		}
		if errors.Is(err, services.ErrAnamnesisNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "anamnesis not found",
				Message: "anamnesis not found or does not belong to patient",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}

// GetByID busca uma avaliação por ID
func (h *PhysicalAssessmentHandler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid assessment id",
			Message: "assessment id must be a valid UUID",
		})
	}

	userID := middleware.GetUserID(c)
	resp, err := h.service.GetByID(id, userID)
	if err != nil {
		if errors.Is(err, services.ErrPhysicalAssessmentNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error: "physical assessment not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.JSON(resp)
}

// List lista avaliações físicas do paciente selecionado
func (h *PhysicalAssessmentHandler) List(c *fiber.Ctx) error {
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

// Delete deleta uma avaliação física
func (h *PhysicalAssessmentHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid assessment id",
			Message: "assessment id must be a valid UUID",
		})
	}

	userRole := middleware.GetPrimaryRole(c)
	err = h.service.Delete(id, userRole)
	if err != nil {
		if errors.Is(err, services.ErrPhysicalAssessmentNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error: "physical assessment not found",
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
