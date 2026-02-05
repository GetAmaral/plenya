package handlers

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
)

type MedicationDefinitionHandler struct {
	service *services.MedicationDefinitionService
}

func NewMedicationDefinitionHandler(service *services.MedicationDefinitionService) *MedicationDefinitionHandler {
	return &MedicationDefinitionHandler{service: service}
}

// List godoc
// @Summary List medication definitions
// @Description Get all medication definitions with optional category filter
// @Tags MedicationDefinitions
// @Produce json
// @Param category query string false "Filter by category (simple, c1, c5, antibiotic, glp1)"
// @Param limit query int false "Limit" default(20)
// @Param offset query int false "Offset" default(0)
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "data, total, page, limit"
// @Failure 500 {object} dto.ErrorResponse
// @Router /medication-definitions [get]
func (h *MedicationDefinitionHandler) List(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "20"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))
	if limit > 100 {
		limit = 100
	}

	var category *models.MedicationCategory
	if categoryStr := c.Query("category"); categoryStr != "" {
		cat := models.MedicationCategory(categoryStr)
		category = &cat
	}

	medications, total, err := h.service.List(category, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data":  medications,
		"total": total,
		"page":  (offset / limit) + 1,
		"limit": limit,
	})
}

// Search godoc
// @Summary Search medication definitions
// @Description Search medications by name or active ingredient (autocomplete)
// @Tags MedicationDefinitions
// @Produce json
// @Param q query string true "Search query"
// @Param limit query int false "Limit" default(10)
// @Security BearerAuth
// @Success 200 {array} models.MedicationDefinition
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /medication-definitions/search [get]
func (h *MedicationDefinitionHandler) Search(c *fiber.Ctx) error {
	query := c.Query("q")
	if query == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "query parameter 'q' is required",
			Message: "provide a search query",
		})
	}

	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	if limit > 50 {
		limit = 50
	}

	medications, err := h.service.Search(query, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.JSON(medications)
}

// GetByID godoc
// @Summary Get medication definition by ID
// @Description Get a single medication definition
// @Tags MedicationDefinitions
// @Produce json
// @Param id path string true "Medication Definition UUID"
// @Security BearerAuth
// @Success 200 {object} models.MedicationDefinition
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /medication-definitions/{id} [get]
func (h *MedicationDefinitionHandler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid id",
			Message: "id must be a valid UUID",
		})
	}

	medication, err := h.service.GetByID(id)
	if err != nil {
		if errors.Is(err, services.ErrMedicationDefinitionNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "medication definition not found",
				Message: "no medication definition found with the given id",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.JSON(medication)
}

// Create godoc
// @Summary Create medication definition
// @Description Create a new medication definition (admin only)
// @Tags MedicationDefinitions
// @Accept json
// @Produce json
// @Param medication body models.MedicationDefinition true "Medication definition data"
// @Security BearerAuth
// @Success 201 {object} models.MedicationDefinition
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /medication-definitions [post]
func (h *MedicationDefinitionHandler) Create(c *fiber.Ctx) error {
	var medication models.MedicationDefinition
	if err := c.BodyParser(&medication); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid request body",
			Message: err.Error(),
		})
	}

	if err := h.service.Create(&medication); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(medication)
}

// Update godoc
// @Summary Update medication definition
// @Description Update an existing medication definition (admin only)
// @Tags MedicationDefinitions
// @Accept json
// @Produce json
// @Param id path string true "Medication Definition UUID"
// @Param medication body models.MedicationDefinition true "Medication definition data"
// @Security BearerAuth
// @Success 200 {object} models.MedicationDefinition
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /medication-definitions/{id} [put]
func (h *MedicationDefinitionHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid id",
			Message: "id must be a valid UUID",
		})
	}

	var medication models.MedicationDefinition
	if err := c.BodyParser(&medication); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid request body",
			Message: err.Error(),
		})
	}

	if err := h.service.Update(id, &medication); err != nil {
		if errors.Is(err, services.ErrMedicationDefinitionNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "medication definition not found",
				Message: "no medication definition found with the given id",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	// Fetch updated medication
	updated, _ := h.service.GetByID(id)
	return c.JSON(updated)
}

// Delete godoc
// @Summary Delete medication definition
// @Description Soft delete a medication definition (admin only)
// @Tags MedicationDefinitions
// @Param id path string true "Medication Definition UUID"
// @Security BearerAuth
// @Success 204
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /medication-definitions/{id} [delete]
func (h *MedicationDefinitionHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid id",
			Message: "id must be a valid UUID",
		})
	}

	if err := h.service.Delete(id); err != nil {
		if errors.Is(err, services.ErrMedicationDefinitionNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "medication definition not found",
				Message: "no medication definition found with the given id",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
