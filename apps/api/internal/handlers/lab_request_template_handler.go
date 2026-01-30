package handlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
)

// LabRequestTemplateHandler handles HTTP requests for lab request templates
type LabRequestTemplateHandler struct {
	service *services.LabRequestTemplateService
}

// NewLabRequestTemplateHandler creates a new lab request template handler
func NewLabRequestTemplateHandler(service *services.LabRequestTemplateService) *LabRequestTemplateHandler {
	return &LabRequestTemplateHandler{service: service}
}

// CreateLabRequestTemplate creates a new lab request template
// @Summary Create lab request template
// @Tags LabRequestTemplates
// @Accept json
// @Produce json
// @Param template body models.LabRequestTemplate true "Lab request template data"
// @Success 201 {object} models.LabRequestTemplate
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/lab-request-templates [post]
func (h *LabRequestTemplateHandler) CreateLabRequestTemplate(c *fiber.Ctx) error {
	var template models.LabRequestTemplate
	if err := c.BodyParser(&template); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	if err := h.service.CreateLabRequestTemplate(&template); err != nil {
		// Check if it's a duplicate key error (unique constraint violation)
		if strings.Contains(err.Error(), "duplicate key") || strings.Contains(err.Error(), "SQLSTATE 23505") {
			return c.Status(fiber.StatusConflict).JSON(dto.ErrorResponse{
				Error: "Já existe um template com este nome",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(template)
}

// GetLabRequestTemplateByID retrieves a lab request template by ID with all associated lab tests
// @Summary Get lab request template by ID
// @Tags LabRequestTemplates
// @Produce json
// @Param id path string true "Lab request template ID"
// @Success 200 {object} models.LabRequestTemplate
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/lab-request-templates/{id} [get]
func (h *LabRequestTemplateHandler) GetLabRequestTemplateByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	template, err := h.service.GetLabRequestTemplateByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(template)
}

// GetAllLabRequestTemplates retrieves all active lab request templates
// @Summary Get all lab request templates
// @Tags LabRequestTemplates
// @Produce json
// @Param withTests query bool false "Include lab tests" default(false)
// @Success 200 {array} models.LabRequestTemplate
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/lab-request-templates [get]
func (h *LabRequestTemplateHandler) GetAllLabRequestTemplates(c *fiber.Ctx) error {
	withTests := c.QueryBool("withTests", false)

	var templates []models.LabRequestTemplate
	var err error

	if withTests {
		templates, err = h.service.GetAllLabRequestTemplatesWithTests()
	} else {
		templates, err = h.service.GetAllLabRequestTemplates()
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(templates)
}

// UpdateLabRequestTemplate updates an existing lab request template
// @Summary Update lab request template
// @Tags LabRequestTemplates
// @Accept json
// @Produce json
// @Param id path string true "Lab request template ID"
// @Param template body models.LabRequestTemplate true "Updated lab request template data"
// @Success 200 {object} models.LabRequestTemplate
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/lab-request-templates/{id} [put]
func (h *LabRequestTemplateHandler) UpdateLabRequestTemplate(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	var template models.LabRequestTemplate
	if err := c.BodyParser(&template); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	if err := h.service.UpdateLabRequestTemplate(id, &template); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	updated, err := h.service.GetLabRequestTemplateByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(updated)
}

// UpdateLabRequestTemplateTestsRequest represents the request body for updating template tests
type UpdateLabRequestTemplateTestsRequest struct {
	TestIDs []uuid.UUID `json:"testIds" validate:"required"`
}

// UpdateLabRequestTemplateTests updates the lab tests associated with a template
// @Summary Update lab request template tests
// @Tags LabRequestTemplates
// @Accept json
// @Produce json
// @Param id path string true "Lab request template ID"
// @Param tests body UpdateLabRequestTemplateTestsRequest true "Test IDs"
// @Success 200 {object} models.LabRequestTemplate
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/lab-request-templates/{id}/tests [put]
func (h *LabRequestTemplateHandler) UpdateLabRequestTemplateTests(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	var req UpdateLabRequestTemplateTestsRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	if err := h.service.UpdateLabRequestTemplateTests(id, req.TestIDs); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	updated, err := h.service.GetLabRequestTemplateByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(updated)
}

// AddLabTestToTemplateRequest represents the request body for adding a single test
type AddLabTestToTemplateRequest struct {
	TestID uuid.UUID `json:"testId" validate:"required"`
}

// AddLabTestToTemplate adds a single lab test to a template
// @Summary Add lab test to template
// @Tags LabRequestTemplates
// @Accept json
// @Produce json
// @Param id path string true "Lab request template ID"
// @Param test body AddLabTestToTemplateRequest true "Test ID"
// @Success 200 {object} models.LabRequestTemplate
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/lab-request-templates/{id}/tests [post]
func (h *LabRequestTemplateHandler) AddLabTestToTemplate(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	var req AddLabTestToTemplateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	if err := h.service.AddLabTestToTemplate(id, req.TestID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	updated, err := h.service.GetLabRequestTemplateByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(updated)
}

// RemoveLabTestFromTemplate removes a single lab test from a template
// @Summary Remove lab test from template
// @Tags LabRequestTemplates
// @Param id path string true "Lab request template ID"
// @Param testId path string true "Test ID to remove"
// @Success 200 {object} models.LabRequestTemplate
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/lab-request-templates/{id}/tests/{testId} [delete]
func (h *LabRequestTemplateHandler) RemoveLabTestFromTemplate(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	testID, err := uuid.Parse(c.Params("testId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid test ID format",
		})
	}

	if err := h.service.RemoveLabTestFromTemplate(id, testID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	updated, err := h.service.GetLabRequestTemplateByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(updated)
}

// DeleteLabRequestTemplate soft deletes a lab request template
// @Summary Delete lab request template
// @Tags LabRequestTemplates
// @Param id path string true "Lab request template ID"
// @Success 204
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/lab-request-templates/{id} [delete]
func (h *LabRequestTemplateHandler) DeleteLabRequestTemplate(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	if err := h.service.DeleteLabRequestTemplate(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// SearchLabRequestTemplates searches templates by name
// @Summary Search lab request templates
// @Tags LabRequestTemplates
// @Produce json
// @Param q query string true "Search term"
// @Success 200 {array} models.LabRequestTemplate
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/lab-request-templates/search [get]
func (h *LabRequestTemplateHandler) SearchLabRequestTemplates(c *fiber.Ctx) error {
	searchTerm := c.Query("q")
	if searchTerm == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Search term is required",
		})
	}

	templates, err := h.service.SearchLabRequestTemplates(searchTerm)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(templates)
}
