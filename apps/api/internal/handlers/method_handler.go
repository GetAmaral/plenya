package handlers

import (
	"github.com/plenya/api/internal/services"

	"github.com/gofiber/fiber/v2"
)

type MethodHandler struct {
	service *services.MethodService
}

func NewMethodHandler(service *services.MethodService) *MethodHandler {
	return &MethodHandler{service: service}
}

// GetAllMethods godoc
// @Summary List all methods
// @Description Returns all methods without relationships
// @Tags Methods
// @Produce json
// @Success 200 {array} models.Method
// @Failure 500 {object} map[string]string
// @Router /api/v1/methods [get]
// @Security BearerAuth
func (h *MethodHandler) GetAllMethods(c *fiber.Ctx) error {
	methods, err := h.service.GetAllMethods()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch methods",
		})
	}

	return c.JSON(methods)
}

// GetMethodByID godoc
// @Summary Get method by ID
// @Description Returns a method by ID without relationships
// @Tags Methods
// @Produce json
// @Param id path string true "Method ID"
// @Success 200 {object} models.Method
// @Failure 404 {object} map[string]string
// @Router /api/v1/methods/{id} [get]
// @Security BearerAuth
func (h *MethodHandler) GetMethodByID(c *fiber.Ctx) error {
	methodID := c.Params("id")

	method, err := h.service.GetMethodByID(methodID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Method not found",
		})
	}

	return c.JSON(method)
}

// GetMethodTree godoc
// @Summary Get method with full hierarchy
// @Description Returns method with letters, pillars, score items, and levels
// @Tags Methods
// @Produce json
// @Param id path string true "Method ID"
// @Success 200 {object} models.Method
// @Failure 404 {object} map[string]string
// @Router /api/v1/methods/{id}/tree [get]
// @Security BearerAuth
func (h *MethodHandler) GetMethodTree(c *fiber.Ctx) error {
	methodID := c.Params("id")

	method, err := h.service.GetMethodWithTree(methodID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Method not found",
		})
	}

	return c.JSON(method)
}

// GetAllMethodsWithTree godoc
// @Summary Get all methods with full hierarchy
// @Description Returns all methods with letters, pillars, score items, and levels
// @Tags Methods
// @Produce json
// @Success 200 {array} models.Method
// @Failure 500 {object} map[string]string
// @Router /api/v1/methods/tree [get]
// @Security BearerAuth
func (h *MethodHandler) GetAllMethodsWithTree(c *fiber.Ctx) error {
	methods, err := h.service.GetAllMethodsWithTree()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch methods",
		})
	}

	return c.JSON(methods)
}

// CreateMethod godoc
// @Summary Create a new method
// @Description Creates a new method
// @Tags Methods
// @Accept json
// @Produce json
// @Param method body services.CreateMethodDTO true "Method data"
// @Success 201 {object} models.Method
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/methods [post]
// @Security BearerAuth
func (h *MethodHandler) CreateMethod(c *fiber.Ctx) error {
	var dto services.CreateMethodDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	method, err := h.service.CreateMethod(&dto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create method",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(method)
}

// UpdateMethod godoc
// @Summary Update a method
// @Description Updates an existing method
// @Tags Methods
// @Accept json
// @Produce json
// @Param id path string true "Method ID"
// @Param method body services.UpdateMethodDTO true "Method data"
// @Success 200 {object} models.Method
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/methods/{id} [put]
// @Security BearerAuth
func (h *MethodHandler) UpdateMethod(c *fiber.Ctx) error {
	methodID := c.Params("id")

	var dto services.UpdateMethodDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	method, err := h.service.UpdateMethod(methodID, &dto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update method",
		})
	}

	return c.JSON(method)
}

// DeleteMethod godoc
// @Summary Delete a method
// @Description Soft deletes a method
// @Tags Methods
// @Param id path string true "Method ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/methods/{id} [delete]
// @Security BearerAuth
func (h *MethodHandler) DeleteMethod(c *fiber.Ctx) error {
	methodID := c.Params("id")

	if err := h.service.DeleteMethod(methodID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete method",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// ===== Method Letter Handlers =====

// GetMethodLetters godoc
// @Summary List letters for a method
// @Description Returns all letters for a specific method
// @Tags Methods
// @Produce json
// @Param id path string true "Method ID"
// @Success 200 {array} models.MethodLetter
// @Failure 500 {object} map[string]string
// @Router /api/v1/methods/{id}/letters [get]
// @Security BearerAuth
func (h *MethodHandler) GetMethodLetters(c *fiber.Ctx) error {
	methodID := c.Params("id")

	letters, err := h.service.GetMethodLetters(methodID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch letters",
		})
	}

	return c.JSON(letters)
}

// GetMethodLetterByID godoc
// @Summary Get letter by ID
// @Description Returns a method letter by ID
// @Tags Methods
// @Produce json
// @Param id path string true "Letter ID"
// @Success 200 {object} models.MethodLetter
// @Failure 404 {object} map[string]string
// @Router /api/v1/method-letters/{id} [get]
// @Security BearerAuth
func (h *MethodHandler) GetMethodLetterByID(c *fiber.Ctx) error {
	letterID := c.Params("id")

	letter, err := h.service.GetMethodLetterByID(letterID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Letter not found",
		})
	}

	return c.JSON(letter)
}

// CreateMethodLetter godoc
// @Summary Create a new letter
// @Description Creates a new method letter
// @Tags Methods
// @Accept json
// @Produce json
// @Param id path string true "Method ID"
// @Param letter body services.CreateMethodLetterDTO true "Letter data"
// @Success 201 {object} models.MethodLetter
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/methods/{id}/letters [post]
// @Security BearerAuth
func (h *MethodHandler) CreateMethodLetter(c *fiber.Ctx) error {
	methodID := c.Params("id")

	var dto services.CreateMethodLetterDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	letter, err := h.service.CreateMethodLetter(methodID, &dto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create letter",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(letter)
}

// UpdateMethodLetter godoc
// @Summary Update a letter
// @Description Updates an existing method letter
// @Tags Methods
// @Accept json
// @Produce json
// @Param id path string true "Letter ID"
// @Param letter body services.UpdateMethodLetterDTO true "Letter data"
// @Success 200 {object} models.MethodLetter
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/method-letters/{id} [put]
// @Security BearerAuth
func (h *MethodHandler) UpdateMethodLetter(c *fiber.Ctx) error {
	letterID := c.Params("id")

	var dto services.UpdateMethodLetterDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	letter, err := h.service.UpdateMethodLetter(letterID, &dto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update letter",
		})
	}

	return c.JSON(letter)
}

// DeleteMethodLetter godoc
// @Summary Delete a letter
// @Description Soft deletes a method letter
// @Tags Methods
// @Param id path string true "Letter ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/method-letters/{id} [delete]
// @Security BearerAuth
func (h *MethodHandler) DeleteMethodLetter(c *fiber.Ctx) error {
	letterID := c.Params("id")

	if err := h.service.DeleteMethodLetter(letterID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete letter",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// ===== Method Pillar Handlers =====

// GetLetterPillars godoc
// @Summary List pillars for a letter
// @Description Returns all pillars for a specific letter
// @Tags Methods
// @Produce json
// @Param id path string true "Letter ID"
// @Success 200 {array} models.MethodPillar
// @Failure 500 {object} map[string]string
// @Router /api/v1/method-letters/{id}/pillars [get]
// @Security BearerAuth
func (h *MethodHandler) GetLetterPillars(c *fiber.Ctx) error {
	letterID := c.Params("id")

	pillars, err := h.service.GetLetterPillars(letterID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch pillars",
		})
	}

	return c.JSON(pillars)
}

// GetMethodPillarByID godoc
// @Summary Get pillar by ID
// @Description Returns a method pillar by ID with score items
// @Tags Methods
// @Produce json
// @Param id path string true "Pillar ID"
// @Success 200 {object} models.MethodPillar
// @Failure 404 {object} map[string]string
// @Router /api/v1/method-pillars/{id} [get]
// @Security BearerAuth
func (h *MethodHandler) GetMethodPillarByID(c *fiber.Ctx) error {
	pillarID := c.Params("id")

	pillar, err := h.service.GetMethodPillarByID(pillarID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Pillar not found",
		})
	}

	return c.JSON(pillar)
}

// CreateMethodPillar godoc
// @Summary Create a new pillar
// @Description Creates a new method pillar
// @Tags Methods
// @Accept json
// @Produce json
// @Param id path string true "Letter ID"
// @Param pillar body services.CreateMethodPillarDTO true "Pillar data"
// @Success 201 {object} models.MethodPillar
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/method-letters/{id}/pillars [post]
// @Security BearerAuth
func (h *MethodHandler) CreateMethodPillar(c *fiber.Ctx) error {
	letterID := c.Params("id")

	var dto services.CreateMethodPillarDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	pillar, err := h.service.CreateMethodPillar(letterID, &dto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create pillar",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(pillar)
}

// UpdateMethodPillar godoc
// @Summary Update a pillar
// @Description Updates an existing method pillar
// @Tags Methods
// @Accept json
// @Produce json
// @Param id path string true "Pillar ID"
// @Param pillar body services.UpdateMethodPillarDTO true "Pillar data"
// @Success 200 {object} models.MethodPillar
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/method-pillars/{id} [put]
// @Security BearerAuth
func (h *MethodHandler) UpdateMethodPillar(c *fiber.Ctx) error {
	pillarID := c.Params("id")

	var dto services.UpdateMethodPillarDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	pillar, err := h.service.UpdateMethodPillar(pillarID, &dto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update pillar",
		})
	}

	return c.JSON(pillar)
}

// DeleteMethodPillar godoc
// @Summary Delete a pillar
// @Description Soft deletes a method pillar
// @Tags Methods
// @Param id path string true "Pillar ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/method-pillars/{id} [delete]
// @Security BearerAuth
func (h *MethodHandler) DeleteMethodPillar(c *fiber.Ctx) error {
	pillarID := c.Params("id")

	if err := h.service.DeleteMethodPillar(pillarID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete pillar",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// ===== Score Item Assignment Handlers =====

// AssignScoreItemToPillar godoc
// @Summary Assign score item to pillar
// @Description Adds a score item to a method pillar (M:N relationship)
// @Tags Methods
// @Accept json
// @Produce json
// @Param id path string true "Pillar ID"
// @Param assignment body services.AssignItemToPillarDTO true "Item ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/method-pillars/{id}/assign-item [post]
// @Security BearerAuth
func (h *MethodHandler) AssignScoreItemToPillar(c *fiber.Ctx) error {
	pillarID := c.Params("id")

	var dto services.AssignItemToPillarDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.service.AssignScoreItemToPillar(pillarID, &dto); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to assign item to pillar",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Item assigned successfully",
	})
}

// UnassignScoreItemFromPillar godoc
// @Summary Unassign score item from pillar
// @Description Removes a score item from a method pillar
// @Tags Methods
// @Accept json
// @Produce json
// @Param id path string true "Pillar ID"
// @Param assignment body services.AssignItemToPillarDTO true "Item ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/method-pillars/{id}/unassign-item [delete]
// @Security BearerAuth
func (h *MethodHandler) UnassignScoreItemFromPillar(c *fiber.Ctx) error {
	pillarID := c.Params("id")

	var dto services.AssignItemToPillarDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.service.UnassignScoreItemFromPillar(pillarID, &dto); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to unassign item from pillar",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Item unassigned successfully",
	})
}

// GetUnassignedScoreItems godoc
// @Summary Get unassigned score items
// @Description Returns all score items not assigned to any pillar
// @Tags Methods
// @Produce json
// @Success 200 {array} models.ScoreItem
// @Failure 500 {object} map[string]string
// @Router /api/v1/score-items/unassigned [get]
// @Security BearerAuth
func (h *MethodHandler) GetUnassignedScoreItems(c *fiber.Ctx) error {
	items, err := h.service.GetUnassignedScoreItems()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch unassigned items",
		})
	}

	return c.JSON(items)
}

// GetAllScoreItemsWithPillars godoc
// @Summary Get all score items with pillars
// @Description Returns all score items with their method pillar associations
// @Tags Methods
// @Produce json
// @Success 200 {array} models.ScoreItem
// @Failure 500 {object} map[string]string
// @Router /api/v1/score-items/with-pillars [get]
// @Security BearerAuth
func (h *MethodHandler) GetAllScoreItemsWithPillars(c *fiber.Ctx) error {
	items, err := h.service.GetAllScoreItemsWithPillars()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch items",
		})
	}

	return c.JSON(items)
}
