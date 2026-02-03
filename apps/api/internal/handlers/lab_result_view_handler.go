package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/repository"
	"github.com/plenya/api/internal/services"
)

type LabResultViewHandler struct {
	service *services.LabResultViewService
}

func NewLabResultViewHandler(service *services.LabResultViewService) *LabResultViewHandler {
	return &LabResultViewHandler{service: service}
}

// GetAll lista todas as views
// @Summary Lista todas as views de resultados laboratoriais
// @Tags lab-result-views
// @Produce json
// @Param withItems query bool false "Incluir items"
// @Param includeInactive query bool false "Incluir inativas"
// @Success 200 {array} models.LabResultView
// @Router /api/v1/lab-result-views [get]
func (h *LabResultViewHandler) GetAll(c *fiber.Ctx) error {
	withItems := c.QueryBool("withItems", false)
	includeInactive := c.QueryBool("includeInactive", false)

	views, err := h.service.GetAll(withItems, includeInactive)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get lab result views",
		})
	}

	return c.JSON(views)
}

// Search busca views por nome ou descrição
// @Summary Busca views por nome ou descrição
// @Tags lab-result-views
// @Produce json
// @Param q query string true "Query de busca"
// @Param withItems query bool false "Incluir items"
// @Success 200 {array} models.LabResultView
// @Router /api/v1/lab-result-views/search [get]
func (h *LabResultViewHandler) Search(c *fiber.Ctx) error {
	query := c.Query("q")
	if query == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Query parameter 'q' is required",
		})
	}

	withItems := c.QueryBool("withItems", false)

	views, err := h.service.Search(query, withItems)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to search lab result views",
		})
	}

	return c.JSON(views)
}

// GetByID busca uma view por ID
// @Summary Busca uma view por ID
// @Tags lab-result-views
// @Produce json
// @Param id path string true "View ID"
// @Success 200 {object} models.LabResultView
// @Router /api/v1/lab-result-views/{id} [get]
func (h *LabResultViewHandler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID format",
		})
	}

	view, err := h.service.GetByID(id, true)
	if err != nil {
		if err == repository.ErrLabResultViewNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"error": "Lab result view not found",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get lab result view",
		})
	}

	return c.JSON(view)
}

// Create cria uma nova view
// @Summary Cria uma nova view
// @Tags lab-result-views
// @Accept json
// @Produce json
// @Param body body dto.CreateLabResultViewRequest true "View data"
// @Success 201 {object} models.LabResultView
// @Router /api/v1/lab-result-views [post]
func (h *LabResultViewHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateLabResultViewRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	view, err := h.service.Create(&req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create lab result view",
		})
	}

	return c.Status(http.StatusCreated).JSON(view)
}

// Update atualiza uma view
// @Summary Atualiza uma view
// @Tags lab-result-views
// @Accept json
// @Produce json
// @Param id path string true "View ID"
// @Param body body dto.UpdateLabResultViewRequest true "View data"
// @Success 200 {object} models.LabResultView
// @Router /api/v1/lab-result-views/{id} [put]
func (h *LabResultViewHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID format",
		})
	}

	var req dto.UpdateLabResultViewRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	view, err := h.service.Update(id, &req)
	if err != nil {
		if err == repository.ErrLabResultViewNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"error": "Lab result view not found",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update lab result view",
		})
	}

	return c.JSON(view)
}

// UpdateItems atualiza os items de uma view
// @Summary Atualiza os items de uma view
// @Tags lab-result-views
// @Accept json
// @Produce json
// @Param id path string true "View ID"
// @Param body body dto.UpdateLabResultViewItemsRequest true "Items data"
// @Success 200 {object} map[string]string
// @Router /api/v1/lab-result-views/{id}/items [put]
func (h *LabResultViewHandler) UpdateItems(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID format",
		})
	}

	var req dto.UpdateLabResultViewItemsRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.service.UpdateItems(id, &req); err != nil {
		if err == repository.ErrLabResultViewNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"error": "Lab result view not found",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update lab result view items",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Items updated successfully",
	})
}

// Delete deleta uma view
// @Summary Deleta uma view
// @Tags lab-result-views
// @Produce json
// @Param id path string true "View ID"
// @Success 204
// @Router /api/v1/lab-result-views/{id} [delete]
func (h *LabResultViewHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID format",
		})
	}

	if err := h.service.Delete(id); err != nil {
		if err == repository.ErrLabResultViewNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"error": "Lab result view not found",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete lab result view",
		})
	}

	return c.SendStatus(http.StatusNoContent)
}
