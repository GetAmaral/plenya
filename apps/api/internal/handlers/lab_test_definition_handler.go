package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
)

// LabTestDefinitionHandler handles HTTP requests for lab test definitions
type LabTestDefinitionHandler struct {
	service *services.LabTestDefinitionService
}

// NewLabTestDefinitionHandler creates a new lab test definition handler
func NewLabTestDefinitionHandler(service *services.LabTestDefinitionService) *LabTestDefinitionHandler {
	return &LabTestDefinitionHandler{service: service}
}

// ============================================================
// LabTestDefinition Endpoints
// ============================================================

// CreateLabTestDefinition creates a new lab test definition
// @Summary Create lab test definition
// @Tags LabTests
// @Accept json
// @Produce json
// @Param definition body models.LabTestDefinition true "Lab test definition data"
// @Success 201 {object} models.LabTestDefinition
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/lab-tests/definitions [post]
func (h *LabTestDefinitionHandler) CreateLabTestDefinition(c *fiber.Ctx) error {
	var def models.LabTestDefinition
	if err := c.BodyParser(&def); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Invalid request body",
		})
	}

	if err := h.service.CreateLabTestDefinition(&def); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(def)
}

// GetLabTestDefinitionByID retrieves a lab test definition by ID
// @Summary Get lab test definition by ID
// @Tags LabTests
// @Produce json
// @Param id path string true "Lab test definition ID"
// @Success 200 {object} models.LabTestDefinition
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/lab-tests/definitions/{id} [get]
func (h *LabTestDefinitionHandler) GetLabTestDefinitionByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	def, err := h.service.GetLabTestDefinitionByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(def)
}

// GetLabTestDefinitionByCode retrieves a lab test definition by code
// @Summary Get lab test definition by code
// @Tags LabTests
// @Produce json
// @Param code path string true "Lab test code"
// @Success 200 {object} models.LabTestDefinition
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/lab-tests/definitions/code/{code} [get]
func (h *LabTestDefinitionHandler) GetLabTestDefinitionByCode(c *fiber.Ctx) error {
	code := c.Params("code")

	def, err := h.service.GetLabTestDefinitionByCode(code)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(def)
}

// GetAllLabTestDefinitions retrieves all lab test definitions
// @Summary Get all lab test definitions
// @Tags LabTests
// @Produce json
// @Success 200 {array} models.LabTestDefinition
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/lab-tests/definitions [get]
func (h *LabTestDefinitionHandler) GetAllLabTestDefinitions(c *fiber.Ctx) error {
	defs, err := h.service.GetAllLabTestDefinitions()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(defs)
}

// GetRequestableLabTests retrieves only requestable lab tests
// @Summary Get requestable lab tests
// @Tags LabTests
// @Produce json
// @Param category query string false "Filter by category"
// @Success 200 {array} models.LabTestDefinition
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/lab-tests/requestable [get]
func (h *LabTestDefinitionHandler) GetRequestableLabTests(c *fiber.Ctx) error {
	var category *models.LabTestCategory
	if categoryStr := c.Query("category"); categoryStr != "" {
		cat := models.LabTestCategory(categoryStr)
		category = &cat
	}

	defs, err := h.service.GetRequestableLabTests(category)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(defs)
}

// GetSubTests retrieves all sub-tests for a parent test
// @Summary Get sub-tests
// @Tags LabTests
// @Produce json
// @Param id path string true "Parent test ID"
// @Success 200 {array} models.LabTestDefinition
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/lab-tests/definitions/{id}/sub-tests [get]
func (h *LabTestDefinitionHandler) GetSubTests(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	defs, err := h.service.GetSubTests(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(defs)
}

// SearchLabTestDefinitions searches lab tests by name or code
// @Summary Search lab tests
// @Tags LabTests
// @Produce json
// @Param q query string true "Search term"
// @Success 200 {array} models.LabTestDefinition
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/lab-tests/definitions/search [get]
func (h *LabTestDefinitionHandler) SearchLabTestDefinitions(c *fiber.Ctx) error {
	searchTerm := c.Query("q")
	if searchTerm == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Search term is required",
		})
	}

	defs, err := h.service.SearchLabTestDefinitions(searchTerm)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(defs)
}

// UpdateLabTestDefinition updates an existing lab test definition
// @Summary Update lab test definition
// @Tags LabTests
// @Accept json
// @Produce json
// @Param id path string true "Lab test definition ID"
// @Param definition body models.LabTestDefinition true "Updated lab test definition data"
// @Success 200 {object} models.LabTestDefinition
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/lab-tests/definitions/{id} [put]
func (h *LabTestDefinitionHandler) UpdateLabTestDefinition(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	var def models.LabTestDefinition
	if err := c.BodyParser(&def); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Invalid request body",
		})
	}

	if err := h.service.UpdateLabTestDefinition(id, &def); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	updated, _ := h.service.GetLabTestDefinitionByID(id)
	return c.JSON(updated)
}

// DeleteLabTestDefinition soft deletes a lab test definition
// @Summary Delete lab test definition
// @Tags LabTests
// @Param id path string true "Lab test definition ID"
// @Success 204
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/lab-tests/definitions/{id} [delete]
func (h *LabTestDefinitionHandler) DeleteLabTestDefinition(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	if err := h.service.DeleteLabTestDefinition(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// ============================================================
// LabTestScoreMapping Endpoints
// ============================================================

// CreateLabTestScoreMapping creates a new mapping
// @Summary Create lab test score mapping
// @Tags LabTests
// @Accept json
// @Produce json
// @Param mapping body models.LabTestScoreMapping true "Mapping data"
// @Success 201 {object} models.LabTestScoreMapping
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/lab-tests/score-mappings [post]
func (h *LabTestDefinitionHandler) CreateLabTestScoreMapping(c *fiber.Ctx) error {
	var mapping models.LabTestScoreMapping
	if err := c.BodyParser(&mapping); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Invalid request body",
		})
	}

	if err := h.service.CreateLabTestScoreMapping(&mapping); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(mapping)
}

// GetLabTestScoreMappingByID retrieves a mapping by ID
// @Summary Get lab test score mapping by ID
// @Tags LabTests
// @Produce json
// @Param id path string true "Mapping ID"
// @Success 200 {object} models.LabTestScoreMapping
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/lab-tests/score-mappings/{id} [get]
func (h *LabTestDefinitionHandler) GetLabTestScoreMappingByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	mapping, err := h.service.GetLabTestScoreMappingByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(mapping)
}

// GetMappingsForLabTest retrieves all score mappings for a lab test
// @Summary Get mappings for lab test
// @Tags LabTests
// @Produce json
// @Param id path string true "Lab test ID"
// @Success 200 {array} models.LabTestScoreMapping
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/lab-tests/definitions/{id}/score-mappings [get]
func (h *LabTestDefinitionHandler) GetMappingsForLabTest(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	mappings, err := h.service.GetMappingsForLabTest(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(mappings)
}

// UpdateLabTestScoreMapping updates a mapping
// @Summary Update lab test score mapping
// @Tags LabTests
// @Accept json
// @Produce json
// @Param id path string true "Mapping ID"
// @Param mapping body models.LabTestScoreMapping true "Updated mapping data"
// @Success 200 {object} models.LabTestScoreMapping
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/lab-tests/score-mappings/{id} [put]
func (h *LabTestDefinitionHandler) UpdateLabTestScoreMapping(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	var mapping models.LabTestScoreMapping
	if err := c.BodyParser(&mapping); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Invalid request body",
		})
	}

	if err := h.service.UpdateLabTestScoreMapping(id, &mapping); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	updated, _ := h.service.GetLabTestScoreMappingByID(id)
	return c.JSON(updated)
}

// DeleteLabTestScoreMapping soft deletes a mapping
// @Summary Delete lab test score mapping
// @Tags LabTests
// @Param id path string true "Mapping ID"
// @Success 204
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/lab-tests/score-mappings/{id} [delete]
func (h *LabTestDefinitionHandler) DeleteLabTestScoreMapping(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	if err := h.service.DeleteLabTestScoreMapping(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
