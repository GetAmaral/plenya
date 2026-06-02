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

type VitalsHandler struct {
	service   *services.VitalsService
	validator *validator.Validate
}

func NewVitalsHandler(service *services.VitalsService) *VitalsHandler {
	return &VitalsHandler{service: service, validator: validator.New()}
}

// ListByPatient GET /api/v1/patients/:id/vitals?appointmentId=&limit=
func (h *VitalsHandler) ListByPatient(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid patient id"})
	}
	var appointmentID *uuid.UUID
	if v := c.Query("appointmentId"); v != "" {
		if aid, err := uuid.Parse(v); err == nil {
			appointmentID = &aid
		}
	}
	limit, _ := strconv.Atoi(c.Query("limit", "50"))
	resp, err := h.service.ListByPatient(patientID, appointmentID, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "internal server error", Message: err.Error()})
	}
	return c.JSON(resp)
}

// Create POST /api/v1/patients/:id/vitals
func (h *VitalsHandler) Create(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid patient id"})
	}
	var req dto.CreateConsultationVitalsRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid request body", Message: err.Error()})
	}
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "validation failed", Details: formatValidationErrors(err)})
	}
	resp, err := h.service.Create(patientID, middleware.GetUserID(c), &req)
	if err != nil {
		if errors.Is(err, services.ErrAppointmentNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "appointment not found"})
		}
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "could not save vitals", Message: err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(resp)
}
