package handlers

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/services"
)

type AnamnesisTemplateHandler struct {
	service   *services.AnamnesisTemplateService
	validator *validator.Validate
}

func NewAnamnesisTemplateHandler(service *services.AnamnesisTemplateService) *AnamnesisTemplateHandler {
	return &AnamnesisTemplateHandler{
		service:   service,
		validator: validator.New(),
	}
}

// GetAll retrieves all anamnesis templates
// @Summary Get all anamnesis templates
// @Tags AnamnesisTemplates
// @Produce json
// @Param withItems query boolean false "Include items"
// @Success 200 {array} models.AnamnesisTemplate
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/anamnesis-templates [get]
func (h *AnamnesisTemplateHandler) GetAll(c *fiber.Ctx) error {
	withItems := c.Query("withItems", "false") == "true"

	templates, err := h.service.GetAll(withItems)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.JSON(templates)
}

// GetByID retrieves an anamnesis template by ID
// @Summary Get anamnesis template by ID
// @Tags AnamnesisTemplates
// @Produce json
// @Param id path string true "Template ID"
// @Success 200 {object} models.AnamnesisTemplate
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/anamnesis-templates/{id} [get]
func (h *AnamnesisTemplateHandler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid template id",
			Message: "template id must be a valid UUID",
		})
	}

	template, err := h.service.GetByID(id, true) // Always include items
	if err != nil {
		if errors.Is(err, services.ErrAnamnesisTemplateNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "template not found",
				Message: "no anamnesis template found with the given id",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.JSON(template)
}

// Create creates a new anamnesis template
// @Summary Create anamnesis template
// @Tags AnamnesisTemplates
// @Accept json
// @Produce json
// @Param request body dto.CreateAnamnesisTemplateRequest true "Template data"
// @Success 201 {object} models.AnamnesisTemplate
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/anamnesis-templates [post]
func (h *AnamnesisTemplateHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateAnamnesisTemplateRequest
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

	template, err := h.service.Create(req.Name, req.Area)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(template)
}

// Update updates an anamnesis template
// @Summary Update anamnesis template
// @Tags AnamnesisTemplates
// @Accept json
// @Produce json
// @Param id path string true "Template ID"
// @Param request body dto.UpdateAnamnesisTemplateRequest true "Template data"
// @Success 200 {object} models.AnamnesisTemplate
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/anamnesis-templates/{id} [put]
func (h *AnamnesisTemplateHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid template id",
			Message: "template id must be a valid UUID",
		})
	}

	var req dto.UpdateAnamnesisTemplateRequest
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

	template, err := h.service.Update(id, req.Name, req.Area)
	if err != nil {
		if errors.Is(err, services.ErrAnamnesisTemplateNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "template not found",
				Message: "no anamnesis template found with the given id",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.JSON(template)
}

// Delete deletes an anamnesis template
// @Summary Delete anamnesis template
// @Tags AnamnesisTemplates
// @Param id path string true "Template ID"
// @Success 204
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/anamnesis-templates/{id} [delete]
func (h *AnamnesisTemplateHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid template id",
			Message: "template id must be a valid UUID",
		})
	}

	if err := h.service.Delete(id); err != nil {
		if errors.Is(err, services.ErrAnamnesisTemplateNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "template not found",
				Message: "no anamnesis template found with the given id",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// UpdateItems updates all items of a template
// @Summary Update template items
// @Tags AnamnesisTemplates
// @Accept json
// @Produce json
// @Param id path string true "Template ID"
// @Param request body dto.UpdateAnamnesisTemplateItemsRequest true "Items data"
// @Success 200
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/anamnesis-templates/{id}/items [put]
func (h *AnamnesisTemplateHandler) UpdateItems(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid template id",
			Message: "template id must be a valid UUID",
		})
	}

	var req dto.UpdateAnamnesisTemplateItemsRequest
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

	// Convert to service format
	itemsData := make([]struct {
		ScoreItemID string
		Order       int
	}, len(req.Items))
	for i, item := range req.Items {
		itemsData[i].ScoreItemID = item.ScoreItemID
		itemsData[i].Order = item.Order
	}

	if err := h.service.UpdateItems(id, itemsData); err != nil {
		if errors.Is(err, services.ErrAnamnesisTemplateNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "template not found",
				Message: "no anamnesis template found with the given id",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// Search searches anamnesis templates
// @Summary Search anamnesis templates
// @Tags AnamnesisTemplates
// @Produce json
// @Param q query string true "Search query"
// @Success 200 {array} models.AnamnesisTemplate
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/anamnesis-templates/search [get]
func (h *AnamnesisTemplateHandler) Search(c *fiber.Ctx) error {
	query := c.Query("q")
	if query == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "missing query parameter",
			Message: "query parameter 'q' is required",
		})
	}

	templates, err := h.service.Search(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.JSON(templates)
}

