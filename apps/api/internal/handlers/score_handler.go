package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/services"
)

// ScoreHandler handles HTTP requests for score operations
type ScoreHandler struct {
	service   *services.ScoreService
	validator *validator.Validate
}

// NewScoreHandler creates a new score handler instance
func NewScoreHandler(service *services.ScoreService, validator *validator.Validate) *ScoreHandler {
	return &ScoreHandler{
		service:   service,
		validator: validator,
	}
}

// ============================================================
// ScoreGroup Handlers
// ============================================================

// GetAllScoreGroups godoc
// @Summary List all score groups
// @Description Get all score groups ordered by order field
// @Tags Score Groups
// @Accept json
// @Produce json
// @Success 200 {array} models.ScoreGroup
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/v1/score-groups [get]
func (h *ScoreHandler) GetAllScoreGroups(c *fiber.Ctx) error {
	groups, err := h.service.GetAllGroups()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch score groups",
		})
	}

	return c.JSON(groups)
}

// GetScoreGroupByID godoc
// @Summary Get score group by ID
// @Description Get a single score group by ID
// @Tags Score Groups
// @Accept json
// @Produce json
// @Param id path string true "Score Group ID (UUID)"
// @Success 200 {object} models.ScoreGroup
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/v1/score-groups/{id} [get]
func (h *ScoreHandler) GetScoreGroupByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid group ID",
		})
	}

	group, err := h.service.GetGroupByID(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(group)
}

// GetScoreGroupTree godoc
// @Summary Get score group with nested data
// @Description Get a score group with all subgroups, items, and levels
// @Tags Score Groups
// @Accept json
// @Produce json
// @Param id path string true "Score Group ID (UUID)"
// @Success 200 {object} models.ScoreGroup
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/v1/score-groups/{id}/tree [get]
func (h *ScoreHandler) GetScoreGroupTree(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid group ID",
		})
	}

	group, err := h.service.GetGroupTree(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(group)
}

// GetAllScoreGroupTrees godoc
// @Summary Get all score groups with nested data
// @Description Get all score groups with subgroups, items, and levels (full hierarchy for mindmap)
// @Tags Score Groups
// @Accept json
// @Produce json
// @Success 200 {array} models.ScoreGroup
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/v1/score-groups/tree [get]
func (h *ScoreHandler) GetAllScoreGroupTrees(c *fiber.Ctx) error {
	groups, err := h.service.GetAllGroupTrees()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch score group trees",
		})
	}

	return c.JSON(groups)
}

// CreateScoreGroup godoc
// @Summary Create a new score group
// @Description Create a new score group (admin only)
// @Tags Score Groups
// @Accept json
// @Produce json
// @Param body body services.CreateScoreGroupDTO true "Score Group Data"
// @Success 201 {object} models.ScoreGroup
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/v1/score-groups [post]
func (h *ScoreHandler) CreateScoreGroup(c *fiber.Ctx) error {
	var dto services.CreateScoreGroupDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.validator.Struct(dto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	group, err := h.service.CreateGroup(dto)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create score group",
		})
	}

	return c.Status(http.StatusCreated).JSON(group)
}

// UpdateScoreGroup godoc
// @Summary Update a score group
// @Description Update an existing score group (admin only)
// @Tags Score Groups
// @Accept json
// @Produce json
// @Param id path string true "Score Group ID (UUID)"
// @Param body body services.UpdateScoreGroupDTO true "Score Group Update Data"
// @Success 200 {object} models.ScoreGroup
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/v1/score-groups/{id} [put]
func (h *ScoreHandler) UpdateScoreGroup(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid group ID",
		})
	}

	var dto services.UpdateScoreGroupDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.validator.Struct(dto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	group, err := h.service.UpdateGroup(id, dto)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(group)
}

// DeleteScoreGroup godoc
// @Summary Delete a score group
// @Description Soft delete a score group (admin only)
// @Tags Score Groups
// @Accept json
// @Produce json
// @Param id path string true "Score Group ID (UUID)"
// @Success 204
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/v1/score-groups/{id} [delete]
func (h *ScoreHandler) DeleteScoreGroup(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid group ID",
		})
	}

	if err := h.service.DeleteGroup(id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(http.StatusNoContent)
}

// ============================================================
// ScoreSubgroup Handlers
// ============================================================

// GetSubgroupsByGroupID godoc
// @Summary Get subgroups by group ID
// @Description Get all subgroups for a specific group
// @Tags Score Subgroups
// @Accept json
// @Produce json
// @Param groupId path string true "Score Group ID (UUID)"
// @Success 200 {array} models.ScoreSubgroup
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/v1/score-groups/{groupId}/subgroups [get]
func (h *ScoreHandler) GetSubgroupsByGroupID(c *fiber.Ctx) error {
	groupID, err := uuid.Parse(c.Params("groupId"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid group ID",
		})
	}

	subgroups, err := h.service.GetSubgroupsByGroupID(groupID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch subgroups",
		})
	}

	return c.JSON(subgroups)
}

// GetScoreSubgroupByID godoc
// @Summary Get score subgroup by ID
// @Description Get a single score subgroup by ID
// @Tags Score Subgroups
// @Accept json
// @Produce json
// @Param id path string true "Score Subgroup ID (UUID)"
// @Success 200 {object} models.ScoreSubgroup
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/v1/score-subgroups/{id} [get]
func (h *ScoreHandler) GetScoreSubgroupByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid subgroup ID",
		})
	}

	subgroup, err := h.service.GetSubgroupByID(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(subgroup)
}

// CreateScoreSubgroup godoc
// @Summary Create a new score subgroup
// @Description Create a new score subgroup (admin only)
// @Tags Score Subgroups
// @Accept json
// @Produce json
// @Param body body services.CreateScoreSubgroupDTO true "Score Subgroup Data"
// @Success 201 {object} models.ScoreSubgroup
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/v1/score-subgroups [post]
func (h *ScoreHandler) CreateScoreSubgroup(c *fiber.Ctx) error {
	var dto services.CreateScoreSubgroupDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.validator.Struct(dto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	subgroup, err := h.service.CreateSubgroup(dto)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(subgroup)
}

// UpdateScoreSubgroup godoc
// @Summary Update a score subgroup
// @Description Update an existing score subgroup (admin only)
// @Tags Score Subgroups
// @Accept json
// @Produce json
// @Param id path string true "Score Subgroup ID (UUID)"
// @Param body body services.UpdateScoreSubgroupDTO true "Score Subgroup Update Data"
// @Success 200 {object} models.ScoreSubgroup
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/v1/score-subgroups/{id} [put]
func (h *ScoreHandler) UpdateScoreSubgroup(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid subgroup ID",
		})
	}

	var dto services.UpdateScoreSubgroupDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.validator.Struct(dto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	subgroup, err := h.service.UpdateSubgroup(id, dto)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(subgroup)
}

// DeleteScoreSubgroup godoc
// @Summary Delete a score subgroup
// @Description Soft delete a score subgroup (admin only)
// @Tags Score Subgroups
// @Accept json
// @Produce json
// @Param id path string true "Score Subgroup ID (UUID)"
// @Success 204
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/v1/score-subgroups/{id} [delete]
func (h *ScoreHandler) DeleteScoreSubgroup(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid subgroup ID",
		})
	}

	if err := h.service.DeleteSubgroup(id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(http.StatusNoContent)
}

// ============================================================
// ScoreItem Handlers
// ============================================================

// GetItemsBySubgroupID godoc
// @Summary Get items by subgroup ID
// @Description Get all items for a specific subgroup
// @Tags Score Items
// @Accept json
// @Produce json
// @Param subgroupId path string true "Score Subgroup ID (UUID)"
// @Success 200 {array} models.ScoreItem
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/v1/score-subgroups/{subgroupId}/items [get]
func (h *ScoreHandler) GetItemsBySubgroupID(c *fiber.Ctx) error {
	subgroupID, err := uuid.Parse(c.Params("subgroupId"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid subgroup ID",
		})
	}

	items, err := h.service.GetItemsBySubgroupID(subgroupID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch items",
		})
	}

	return c.JSON(items)
}

// GetAllScoreItems godoc
// @Summary Get all score items
// @Description Get all score items ordered by name
// @Tags Score Items
// @Accept json
// @Produce json
// @Success 200 {array} models.ScoreItem
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/v1/score-items [get]
func (h *ScoreHandler) GetAllScoreItems(c *fiber.Ctx) error {
	items, err := h.service.GetAllScoreItems()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve score items",
		})
	}

	return c.Status(http.StatusOK).JSON(items)
}

// GetScoreItemByID godoc
// @Summary Get score item by ID
// @Description Get a single score item by ID with levels
// @Tags Score Items
// @Accept json
// @Produce json
// @Param id path string true "Score Item ID (UUID)"
// @Success 200 {object} models.ScoreItem
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/v1/score-items/{id} [get]
func (h *ScoreHandler) GetScoreItemByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid item ID",
		})
	}

	item, err := h.service.GetItemByID(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(item)
}

// CreateScoreItem godoc
// @Summary Create a new score item
// @Description Create a new score item (admin only)
// @Tags Score Items
// @Accept json
// @Produce json
// @Param body body services.CreateScoreItemDTO true "Score Item Data"
// @Success 201 {object} models.ScoreItem
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/v1/score-items [post]
func (h *ScoreHandler) CreateScoreItem(c *fiber.Ctx) error {
	var dto services.CreateScoreItemDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.validator.Struct(dto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	item, err := h.service.CreateItem(dto)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(item)
}

// UpdateScoreItem godoc
// @Summary Update a score item
// @Description Update an existing score item (admin only)
// @Tags Score Items
// @Accept json
// @Produce json
// @Param id path string true "Score Item ID (UUID)"
// @Param body body services.UpdateScoreItemDTO true "Score Item Update Data"
// @Success 200 {object} models.ScoreItem
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/v1/score-items/{id} [put]
func (h *ScoreHandler) UpdateScoreItem(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid item ID",
		})
	}

	var dto services.UpdateScoreItemDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.validator.Struct(dto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	item, err := h.service.UpdateItem(id, dto)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(item)
}

// DeleteScoreItem godoc
// @Summary Delete a score item
// @Description Soft delete a score item (admin only)
// @Tags Score Items
// @Accept json
// @Produce json
// @Param id path string true "Score Item ID (UUID)"
// @Success 204
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/v1/score-items/{id} [delete]
func (h *ScoreHandler) DeleteScoreItem(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid item ID",
		})
	}

	if err := h.service.DeleteItem(id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(http.StatusNoContent)
}

// ============================================================
// ScoreLevel Handlers
// ============================================================

// GetLevelsByItemID godoc
// @Summary Get levels by item ID
// @Description Get all levels for a specific item
// @Tags Score Levels
// @Accept json
// @Produce json
// @Param itemId path string true "Score Item ID (UUID)"
// @Success 200 {array} models.ScoreLevel
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/v1/score-items/{itemId}/levels [get]
func (h *ScoreHandler) GetLevelsByItemID(c *fiber.Ctx) error {
	itemID, err := uuid.Parse(c.Params("itemId"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid item ID",
		})
	}

	levels, err := h.service.GetLevelsByItemID(itemID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch levels",
		})
	}

	return c.JSON(levels)
}

// GetScoreLevelByID godoc
// @Summary Get score level by ID
// @Description Get a single score level by ID
// @Tags Score Levels
// @Accept json
// @Produce json
// @Param id path string true "Score Level ID (UUID)"
// @Success 200 {object} models.ScoreLevel
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/v1/score-levels/{id} [get]
func (h *ScoreHandler) GetScoreLevelByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid level ID",
		})
	}

	level, err := h.service.GetLevelByID(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(level)
}

// CreateScoreLevel godoc
// @Summary Create a new score level
// @Description Create a new score level (admin only)
// @Tags Score Levels
// @Accept json
// @Produce json
// @Param body body services.CreateScoreLevelDTO true "Score Level Data"
// @Success 201 {object} models.ScoreLevel
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/v1/score-levels [post]
func (h *ScoreHandler) CreateScoreLevel(c *fiber.Ctx) error {
	var dto services.CreateScoreLevelDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.validator.Struct(dto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	level, err := h.service.CreateLevel(dto)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(level)
}

// UpdateScoreLevel godoc
// @Summary Update a score level
// @Description Update an existing score level (admin only)
// @Tags Score Levels
// @Accept json
// @Produce json
// @Param id path string true "Score Level ID (UUID)"
// @Param body body services.UpdateScoreLevelDTO true "Score Level Update Data"
// @Success 200 {object} models.ScoreLevel
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/v1/score-levels/{id} [put]
func (h *ScoreHandler) UpdateScoreLevel(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid level ID",
		})
	}

	var dto services.UpdateScoreLevelDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.validator.Struct(dto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	level, err := h.service.UpdateLevel(id, dto)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(level)
}

// DeleteScoreLevel godoc
// @Summary Delete a score level
// @Description Soft delete a score level (admin only)
// @Tags Score Levels
// @Accept json
// @Produce json
// @Param id path string true "Score Level ID (UUID)"
// @Success 204
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/v1/score-levels/{id} [delete]
func (h *ScoreHandler) DeleteScoreLevel(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid level ID",
		})
	}

	if err := h.service.DeleteLevel(id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(http.StatusNoContent)
}

// ============================================================
// PDF Generation
// ============================================================

// GeneratePosterPDF godoc
// @Summary Generate score poster PDF
// @Description Generate a PDF poster (60cm x 300cm) with all scores
// @Tags Score Groups
// @Produce application/pdf
// @Success 200 {file} binary
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /api/v1/score-groups/poster-pdf [get]
func (h *ScoreHandler) GeneratePosterPDF(c *fiber.Ctx) error {
	// Generate PDF
	pdfBytes, err := h.service.GeneratePosterPDF()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to generate PDF: %v", err),
		})
	}

	// Set headers for PDF download
	c.Set("Content-Type", "application/pdf")
	c.Set("Content-Disposition", "attachment; filename=escore-plenya-poster.pdf")

	return c.Send(pdfBytes)
}
