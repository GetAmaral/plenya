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

type CarePlanHandler struct {
	service       *services.CarePlanService
	reportService *services.CarePlanReportService
	validator     *validator.Validate
}

func NewCarePlanHandler(service *services.CarePlanService, reportService *services.CarePlanReportService) *CarePlanHandler {
	return &CarePlanHandler{service: service, reportService: reportService, validator: validator.New()}
}

// ListByPatient GET /api/v1/patients/:id/care-plan-items?includeInactive=
func (h *CarePlanHandler) ListByPatient(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid patient id", Message: err.Error()})
	}
	includeInactive := c.Query("includeInactive") == "true"
	items, err := h.service.ListByPatient(patientID, includeInactive)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "failed to list care plan", Message: err.Error()})
	}
	return c.JSON(items)
}

// Create POST /api/v1/patients/:id/care-plan-items
func (h *CarePlanHandler) Create(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid patient id", Message: err.Error()})
	}
	var req dto.CreateCarePlanItemRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid request body", Message: err.Error()})
	}
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "validation failed", Details: formatValidationErrors(err)})
	}
	item, err := h.service.Create(patientID, middleware.GetUserID(c), string(middleware.GetPrimaryRole(c)), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "failed to create care plan item", Message: err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(item)
}

// Update PUT /api/v1/patients/:id/care-plan-items/:itemId
func (h *CarePlanHandler) Update(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid patient id", Message: err.Error()})
	}
	itemID, err := uuid.Parse(c.Params("itemId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid item id", Message: err.Error()})
	}
	var req dto.UpdateCarePlanItemRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid request body", Message: err.Error()})
	}
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "validation failed", Details: formatValidationErrors(err)})
	}
	item, err := h.service.Update(itemID, patientID, &req)
	if err != nil {
		if errors.Is(err, services.ErrCarePlanItemNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "care plan item not found", Message: err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "failed to update care plan item", Message: err.Error()})
	}
	return c.JSON(item)
}

// GenerateReport POST /api/v1/patients/:id/care-plan-report
// Gera + assina + publica o relatório longitudinal AGIR no portal do paciente.
func (h *CarePlanHandler) GenerateReport(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid patient id", Message: err.Error()})
	}
	docID, err := h.reportService.GenerateAndPublish(patientID, middleware.GetUserID(c))
	if err != nil {
		if errors.Is(err, services.ErrNoSnapshotForReport) {
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "sem escore", Message: err.Error()})
		}
		if errors.Is(err, services.ErrPatientNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "paciente não encontrado", Message: err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "falha ao gerar relatório", Message: err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"issuedDocumentId": docID})
}

// Delete DELETE /api/v1/patients/:id/care-plan-items/:itemId
func (h *CarePlanHandler) Delete(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid patient id", Message: err.Error()})
	}
	itemID, err := uuid.Parse(c.Params("itemId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid item id", Message: err.Error()})
	}
	if err := h.service.Delete(itemID, patientID); err != nil {
		if errors.Is(err, services.ErrCarePlanItemNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "care plan item not found", Message: err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "failed to delete care plan item", Message: err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
