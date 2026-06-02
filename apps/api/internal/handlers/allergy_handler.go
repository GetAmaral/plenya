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

type AllergyHandler struct {
	service   *services.AllergyService
	validator *validator.Validate
}

func NewAllergyHandler(service *services.AllergyService) *AllergyHandler {
	return &AllergyHandler{service: service, validator: validator.New()}
}

// ListByPatient GET /api/v1/patients/:id/allergies?includeInactive=true
func (h *AllergyHandler) ListByPatient(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid patient id"})
	}
	includeInactive := c.Query("includeInactive") == "true"
	resp, err := h.service.ListByPatient(patientID, includeInactive)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "internal server error", Message: err.Error()})
	}
	return c.JSON(resp)
}

// Create POST /api/v1/patients/:id/allergies
func (h *AllergyHandler) Create(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid patient id"})
	}
	var req dto.CreatePatientAllergyRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid request body", Message: err.Error()})
	}
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "validation failed", Details: formatValidationErrors(err)})
	}
	resp, err := h.service.Create(patientID, middleware.GetUserID(c), &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "could not create allergy", Message: err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(resp)
}

// Update PUT /api/v1/patients/:id/allergies/:allergyId
func (h *AllergyHandler) Update(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid patient id"})
	}
	allergyID, err := uuid.Parse(c.Params("allergyId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid allergy id"})
	}
	var req dto.UpdatePatientAllergyRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid request body", Message: err.Error()})
	}
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "validation failed", Details: formatValidationErrors(err)})
	}
	resp, err := h.service.Update(allergyID, patientID, &req)
	if err != nil {
		if errors.Is(err, services.ErrAllergyNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "allergy not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "internal server error", Message: err.Error()})
	}
	return c.JSON(resp)
}

// Delete DELETE /api/v1/patients/:id/allergies/:allergyId
func (h *AllergyHandler) Delete(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid patient id"})
	}
	allergyID, err := uuid.Parse(c.Params("allergyId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid allergy id"})
	}
	if err := h.service.Delete(allergyID, patientID); err != nil {
		if errors.Is(err, services.ErrAllergyNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "allergy not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "internal server error", Message: err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
