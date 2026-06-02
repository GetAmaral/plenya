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

type ProblemHandler struct {
	service   *services.ProblemService
	validator *validator.Validate
}

func NewProblemHandler(service *services.ProblemService) *ProblemHandler {
	return &ProblemHandler{service: service, validator: validator.New()}
}

// ListByPatient GET /api/v1/patients/:id/problems?includeResolved=true
func (h *ProblemHandler) ListByPatient(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid patient id"})
	}
	resp, err := h.service.ListByPatient(patientID, c.Query("includeResolved") == "true")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "internal server error", Message: err.Error()})
	}
	return c.JSON(resp)
}

// Create POST /api/v1/patients/:id/problems
func (h *ProblemHandler) Create(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid patient id"})
	}
	var req dto.CreatePatientProblemRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid request body", Message: err.Error()})
	}
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "validation failed", Details: formatValidationErrors(err)})
	}
	resp, err := h.service.Create(patientID, middleware.GetUserID(c), &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "could not create problem", Message: err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(resp)
}

// Update PUT /api/v1/patients/:id/problems/:problemId
func (h *ProblemHandler) Update(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid patient id"})
	}
	problemID, err := uuid.Parse(c.Params("problemId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid problem id"})
	}
	var req dto.UpdatePatientProblemRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid request body", Message: err.Error()})
	}
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "validation failed", Details: formatValidationErrors(err)})
	}
	resp, err := h.service.Update(problemID, patientID, &req)
	if err != nil {
		if errors.Is(err, services.ErrProblemNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "problem not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "internal server error", Message: err.Error()})
	}
	return c.JSON(resp)
}

// Delete DELETE /api/v1/patients/:id/problems/:problemId
func (h *ProblemHandler) Delete(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid patient id"})
	}
	problemID, err := uuid.Parse(c.Params("problemId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid problem id"})
	}
	if err := h.service.Delete(problemID, patientID); err != nil {
		if errors.Is(err, services.ErrProblemNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "problem not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "internal server error", Message: err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
