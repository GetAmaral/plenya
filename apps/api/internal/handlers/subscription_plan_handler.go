package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/plenya/api/internal/services"
)

type SubscriptionPlanHandler struct {
	service *services.SubscriptionPlanService
}

func NewSubscriptionPlanHandler(service *services.SubscriptionPlanService) *SubscriptionPlanHandler {
	return &SubscriptionPlanHandler{service: service}
}

// GetAllPlans godoc
// @Summary List all subscription plans
// @Description Returns all subscription plans (admin only)
// @Tags Subscription Plans
// @Produce json
// @Success 200 {array} models.SubscriptionPlan
// @Failure 500 {object} map[string]string
// @Router /api/v1/subscription-plans [get]
// @Security BearerAuth
func (h *SubscriptionPlanHandler) GetAllPlans(c *fiber.Ctx) error {
	plans, err := h.service.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch subscription plans",
		})
	}

	return c.JSON(plans)
}

// GetActivePlans godoc
// @Summary List active subscription plans
// @Description Returns all active subscription plans
// @Tags Subscription Plans
// @Produce json
// @Success 200 {array} models.SubscriptionPlan
// @Failure 500 {object} map[string]string
// @Router /api/v1/subscription-plans/active [get]
// @Security BearerAuth
func (h *SubscriptionPlanHandler) GetActivePlans(c *fiber.Ctx) error {
	plans, err := h.service.GetActive()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch active plans",
		})
	}

	return c.JSON(plans)
}

// GetPlanByID godoc
// @Summary Get subscription plan by ID
// @Description Returns a subscription plan by ID
// @Tags Subscription Plans
// @Produce json
// @Param id path string true "Plan ID"
// @Success 200 {object} models.SubscriptionPlan
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/subscription-plans/{id} [get]
// @Security BearerAuth
func (h *SubscriptionPlanHandler) GetPlanByID(c *fiber.Ctx) error {
	planID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid plan ID",
		})
	}

	plan, err := h.service.GetByID(planID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Plan not found",
		})
	}

	return c.JSON(plan)
}

// CreatePlan godoc
// @Summary Create a new subscription plan
// @Description Creates a new subscription plan (admin only)
// @Tags Subscription Plans
// @Accept json
// @Produce json
// @Param plan body services.CreatePlanDTO true "Plan data"
// @Success 201 {object} models.SubscriptionPlan
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/subscription-plans [post]
// @Security BearerAuth
func (h *SubscriptionPlanHandler) CreatePlan(c *fiber.Ctx) error {
	var dto services.CreatePlanDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	plan, err := h.service.Create(dto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create plan",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(plan)
}

// UpdatePlan godoc
// @Summary Update a subscription plan
// @Description Updates an existing subscription plan (admin only)
// @Tags Subscription Plans
// @Accept json
// @Produce json
// @Param id path string true "Plan ID"
// @Param plan body services.UpdatePlanDTO true "Plan data"
// @Success 200 {object} models.SubscriptionPlan
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/subscription-plans/{id} [put]
// @Security BearerAuth
func (h *SubscriptionPlanHandler) UpdatePlan(c *fiber.Ctx) error {
	planID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid plan ID",
		})
	}

	var dto services.UpdatePlanDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	plan, err := h.service.Update(planID, dto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update plan",
		})
	}

	return c.JSON(plan)
}

// DeletePlan godoc
// @Summary Delete a subscription plan
// @Description Soft deletes a subscription plan (admin only)
// @Tags Subscription Plans
// @Param id path string true "Plan ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/subscription-plans/{id} [delete]
// @Security BearerAuth
func (h *SubscriptionPlanHandler) DeletePlan(c *fiber.Ctx) error {
	planID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid plan ID",
		})
	}

	if err := h.service.Delete(planID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete plan",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// ActivatePlan godoc
// @Summary Activate a subscription plan
// @Description Sets a subscription plan as active (admin only)
// @Tags Subscription Plans
// @Accept json
// @Produce json
// @Param id path string true "Plan ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/subscription-plans/{id}/activate [post]
// @Security BearerAuth
func (h *SubscriptionPlanHandler) ActivatePlan(c *fiber.Ctx) error {
	planID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid plan ID",
		})
	}

	if err := h.service.Activate(planID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to activate plan",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Plan activated successfully",
	})
}

// DeactivatePlan godoc
// @Summary Deactivate a subscription plan
// @Description Sets a subscription plan as inactive (admin only)
// @Tags Subscription Plans
// @Accept json
// @Produce json
// @Param id path string true "Plan ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/subscription-plans/{id}/deactivate [post]
// @Security BearerAuth
func (h *SubscriptionPlanHandler) DeactivatePlan(c *fiber.Ctx) error {
	planID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid plan ID",
		})
	}

	if err := h.service.Deactivate(planID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to deactivate plan",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Plan deactivated successfully",
	})
}
