package handlers

import (
	"errors"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/middleware"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
)

type PrescriptionHandler struct {
	prescriptionService *services.PrescriptionService
	validator           *validator.Validate
}

func NewPrescriptionHandler(prescriptionService *services.PrescriptionService) *PrescriptionHandler {
	return &PrescriptionHandler{
		prescriptionService: prescriptionService,
		validator:           validator.New(),
	}
}

// Create cria uma nova prescrição
func (h *PrescriptionHandler) Create(c *fiber.Ctx) error {
	var req dto.CreatePrescriptionRequest
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

	resp, err := h.prescriptionService.Create(doctorID, &req)
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

// GetByID busca uma prescrição por ID
func (h *PrescriptionHandler) GetByID(c *fiber.Ctx) error {
	prescriptionID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid prescription id",
			Message: "prescription id must be a valid UUID",
		})
	}

	userID := middleware.GetUserID(c)
	userRole := middleware.GetUserRole(c)

	resp, err := h.prescriptionService.GetByID(prescriptionID, userID, userRole)
	if err != nil {
		if errors.Is(err, services.ErrPrescriptionNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "prescription not found",
				Message: "no prescription found with the given id",
			})
		}
		if errors.Is(err, services.ErrUnauthorized) {
			return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
				Error:   "unauthorized",
				Message: "you do not have permission to view this prescription",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.JSON(resp)
}

// List lista prescrições
func (h *PrescriptionHandler) List(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "20"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))
	if limit > 100 {
		limit = 100
	}

	userID := middleware.GetUserID(c)
	userRole := middleware.GetUserRole(c)

	// Parse optional filters
	var patientID *uuid.UUID
	var status *models.PrescriptionStatus

	if patientIDStr := c.Query("patientId"); patientIDStr != "" {
		pid, err := uuid.Parse(patientIDStr)
		if err == nil {
			patientID = &pid
		}
	}
	if statusStr := c.Query("status"); statusStr != "" {
		s := models.PrescriptionStatus(statusStr)
		status = &s
	}

	resp, err := h.prescriptionService.List(userID, userRole, patientID, status, limit, offset)
	if err != nil {
		if errors.Is(err, services.ErrUnauthorized) {
			return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
				Error:   "unauthorized",
				Message: "you do not have permission to list prescriptions",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.JSON(resp)
}

// Update atualiza uma prescrição
func (h *PrescriptionHandler) Update(c *fiber.Ctx) error {
	prescriptionID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid prescription id",
			Message: "prescription id must be a valid UUID",
		})
	}

	var req dto.UpdatePrescriptionRequest
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
	userRole := middleware.GetUserRole(c)

	resp, err := h.prescriptionService.Update(prescriptionID, doctorID, userRole, &req)
	if err != nil {
		if errors.Is(err, services.ErrPrescriptionNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "prescription not found",
				Message: "no prescription found with the given id",
			})
		}
		if errors.Is(err, services.ErrUnauthorized) {
			return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
				Error:   "unauthorized",
				Message: "you do not have permission to update this prescription",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.JSON(resp)
}

// Delete deleta uma prescrição
func (h *PrescriptionHandler) Delete(c *fiber.Ctx) error {
	prescriptionID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid prescription id",
			Message: "prescription id must be a valid UUID",
		})
	}

	doctorID := middleware.GetUserID(c)
	userRole := middleware.GetUserRole(c)

	err = h.prescriptionService.Delete(prescriptionID, doctorID, userRole)
	if err != nil {
		if errors.Is(err, services.ErrPrescriptionNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "prescription not found",
				Message: "no prescription found with the given id",
			})
		}
		if errors.Is(err, services.ErrUnauthorized) {
			return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
				Error:   "unauthorized",
				Message: "you do not have permission to delete this prescription",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
