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

type LabResultHandler struct {
	labResultService *services.LabResultService
	validator        *validator.Validate
}

func NewLabResultHandler(labResultService *services.LabResultService) *LabResultHandler {
	return &LabResultHandler{
		labResultService: labResultService,
		validator:        validator.New(),
	}
}

// Create cria um novo resultado de exame
func (h *LabResultHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateLabResultRequest
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

	doctorID := middleware.GetUserID(c)

	resp, err := h.labResultService.Create(doctorID, &req)
	if err != nil {
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

	return c.Status(fiber.StatusCreated).JSON(resp)
}

// GetByID busca um resultado de exame por ID
func (h *LabResultHandler) GetByID(c *fiber.Ctx) error {
	labResultID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid lab result id",
			Message: "lab result id must be a valid UUID",
		})
	}

	userID := middleware.GetUserID(c)

	resp, err := h.labResultService.GetByID(labResultID, userID)
	if err != nil {
		if errors.Is(err, services.ErrLabResultNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "lab result not found",
				Message: "no lab result found with the given id",
			})
		}
		if errors.Is(err, services.ErrUnauthorized) {
			return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
				Error:   "unauthorized",
				Message: "you do not have permission to view this lab result",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.JSON(resp)
}

// List lista resultados de exames
func (h *LabResultHandler) List(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "20"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))
	if limit > 100 {
		limit = 100
	}

	userID := middleware.GetUserID(c)

	resp, err := h.labResultService.List(userID, limit, offset)
	if err != nil {
		if errors.Is(err, services.ErrUnauthorized) {
			return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
				Error:   "unauthorized",
				Message: "you do not have permission to list lab results",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.JSON(resp)
}

// Update atualiza um resultado de exame
func (h *LabResultHandler) Update(c *fiber.Ctx) error {
	labResultID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid lab result id",
			Message: "lab result id must be a valid UUID",
		})
	}

	var req dto.UpdateLabResultRequest
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

	resp, err := h.labResultService.Update(labResultID, userID, &req)
	if err != nil {
		if errors.Is(err, services.ErrLabResultNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "lab result not found",
				Message: "no lab result found with the given id",
			})
		}
		if errors.Is(err, services.ErrUnauthorized) {
			return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
				Error:   "unauthorized",
				Message: "you do not have permission to update this lab result",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.JSON(resp)
}

// Delete deleta um resultado de exame
func (h *LabResultHandler) Delete(c *fiber.Ctx) error {
	labResultID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid lab result id",
			Message: "lab result id must be a valid UUID",
		})
	}

	userID := middleware.GetUserID(c)
	userRole := string(middleware.GetPrimaryRole(c))

	err = h.labResultService.Delete(labResultID, userID, userRole)
	if err != nil {
		if errors.Is(err, services.ErrLabResultNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "lab result not found",
				Message: "no lab result found with the given id",
			})
		}
		if errors.Is(err, services.ErrUnauthorized) {
			return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
				Error:   "unauthorized",
				Message: "only admins can delete lab results",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
