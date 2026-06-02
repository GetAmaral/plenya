package handlers

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/middleware"
	"github.com/plenya/api/internal/services"
)

type MedicationInUseHandler struct {
	service   *services.MedicationInUseService
	validator *validator.Validate
}

func NewMedicationInUseHandler(service *services.MedicationInUseService) *MedicationInUseHandler {
	return &MedicationInUseHandler{service: service, validator: validator.New()}
}

// ListByPatient GET /api/v1/patients/:id/medications?includeInactive=true
func (h *MedicationInUseHandler) ListByPatient(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid patient id"})
	}
	resp, err := h.service.ListByPatient(patientID, c.Query("includeInactive") == "true")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "internal server error", Message: err.Error()})
	}
	return c.JSON(resp)
}

// Create POST /api/v1/patients/:id/medications
func (h *MedicationInUseHandler) Create(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid patient id"})
	}
	var req dto.CreateMedicationInUseRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid request body", Message: err.Error()})
	}
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "validation failed", Details: formatValidationErrors(err)})
	}
	resp, err := h.service.Create(patientID, middleware.GetUserID(c), &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "could not create medication", Message: err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(resp)
}

// Update PUT /api/v1/patients/:id/medications/:medId
func (h *MedicationInUseHandler) Update(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid patient id"})
	}
	medID, err := uuid.Parse(c.Params("medId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid medication id"})
	}
	var req dto.UpdateMedicationInUseRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid request body", Message: err.Error()})
	}
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "validation failed", Details: formatValidationErrors(err)})
	}
	resp, err := h.service.Update(medID, patientID, &req)
	if err != nil {
		if errors.Is(err, services.ErrMedicationInUseNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "medication not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "internal server error", Message: err.Error()})
	}
	return c.JSON(resp)
}

// Delete DELETE /api/v1/patients/:id/medications/:medId
func (h *MedicationInUseHandler) Delete(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid patient id"})
	}
	medID, err := uuid.Parse(c.Params("medId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid medication id"})
	}
	if err := h.service.Delete(medID, patientID); err != nil {
		if errors.Is(err, services.ErrMedicationInUseNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "medication not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "internal server error", Message: err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
