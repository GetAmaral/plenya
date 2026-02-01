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

type AnamnesisHandler struct {
	anamnesisService *services.AnamnesisService
	validator        *validator.Validate
}

func NewAnamnesisHandler(anamnesisService *services.AnamnesisService) *AnamnesisHandler {
	return &AnamnesisHandler{
		anamnesisService: anamnesisService,
		validator:        validator.New(),
	}
}

// Create cria uma nova anamnese
func (h *AnamnesisHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateAnamnesisRequest
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

	authorID := middleware.GetUserID(c)

	resp, err := h.anamnesisService.Create(authorID, &req)
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

// GetByID busca uma anamnese por ID
func (h *AnamnesisHandler) GetByID(c *fiber.Ctx) error {
	anamnesisID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid anamnesis id",
			Message: "anamnesis id must be a valid UUID",
		})
	}

	userID := middleware.GetUserID(c)
	userRole := middleware.GetUserRole(c)

	resp, err := h.anamnesisService.GetByID(anamnesisID, userID, userRole)
	if err != nil {
		if errors.Is(err, services.ErrAnamnesisNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "anamnesis not found",
				Message: "no anamnesis found with the given id",
			})
		}
		if errors.Is(err, services.ErrUnauthorized) {
			return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
				Error:   "unauthorized",
				Message: "you do not have permission to view this anamnesis",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.JSON(resp)
}

// List lista anamneses
func (h *AnamnesisHandler) List(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "20"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))
	if limit > 100 {
		limit = 100
	}

	userID := middleware.GetUserID(c)
	userRole := middleware.GetUserRole(c)

	// Parse optional patient ID filter
	var patientID *uuid.UUID
	if patientIDStr := c.Query("patientId"); patientIDStr != "" {
		pid, err := uuid.Parse(patientIDStr)
		if err == nil {
			patientID = &pid
		}
	}

	resp, err := h.anamnesisService.List(userID, userRole, patientID, limit, offset)
	if err != nil {
		if errors.Is(err, services.ErrUnauthorized) {
			return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
				Error:   "unauthorized",
				Message: "you do not have permission to list anamneses",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.JSON(resp)
}

// Update atualiza uma anamnese
func (h *AnamnesisHandler) Update(c *fiber.Ctx) error {
	anamnesisID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid anamnesis id",
			Message: "anamnesis id must be a valid UUID",
		})
	}

	var req dto.UpdateAnamnesisRequest
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

	authorID := middleware.GetUserID(c)
	userRole := middleware.GetUserRole(c)

	resp, err := h.anamnesisService.Update(anamnesisID, authorID, userRole, &req)
	if err != nil {
		if errors.Is(err, services.ErrAnamnesisNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "anamnesis not found",
				Message: "no anamnesis found with the given id",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.JSON(resp)
}

// Delete deleta uma anamnese
func (h *AnamnesisHandler) Delete(c *fiber.Ctx) error {
	anamnesisID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid anamnesis id",
			Message: "anamnesis id must be a valid UUID",
		})
	}

	authorID := middleware.GetUserID(c)
	userRole := middleware.GetUserRole(c)

	err = h.anamnesisService.Delete(anamnesisID, authorID, userRole)
	if err != nil {
		if errors.Is(err, services.ErrAnamnesisNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "anamnesis not found",
				Message: "no anamnesis found with the given id",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
