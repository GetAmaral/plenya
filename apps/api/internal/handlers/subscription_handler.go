package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/plenya/api/internal/services"
)

type SubscriptionHandler struct {
	service *services.SubscriptionService
}

func NewSubscriptionHandler(service *services.SubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{service: service}
}

// GetAllSubscriptions godoc
// @Summary List all subscriptions
// @Description Returns all subscriptions (admin only)
// @Tags Subscriptions
// @Produce json
// @Success 200 {array} models.PatientSubscription
// @Failure 500 {object} map[string]string
// @Router /api/v1/subscriptions [get]
// @Security BearerAuth
func (h *SubscriptionHandler) GetAllSubscriptions(c *fiber.Ctx) error {
	subscriptions, err := h.service.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch subscriptions",
		})
	}

	return c.JSON(subscriptions)
}

// GetActiveSubscriptions godoc
// @Summary List active subscriptions
// @Description Returns all active subscriptions
// @Tags Subscriptions
// @Produce json
// @Success 200 {array} models.PatientSubscription
// @Failure 500 {object} map[string]string
// @Router /api/v1/subscriptions/active [get]
// @Security BearerAuth
func (h *SubscriptionHandler) GetActiveSubscriptions(c *fiber.Ctx) error {
	subscriptions, err := h.service.GetActive()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch active subscriptions",
		})
	}

	return c.JSON(subscriptions)
}

// GetSubscriptionByID godoc
// @Summary Get subscription by ID
// @Description Returns a subscription by ID
// @Tags Subscriptions
// @Produce json
// @Param id path string true "Subscription ID"
// @Success 200 {object} models.PatientSubscription
// @Failure 404 {object} map[string]string
// @Router /api/v1/subscriptions/{id} [get]
// @Security BearerAuth
func (h *SubscriptionHandler) GetSubscriptionByID(c *fiber.Ctx) error {
	subscriptionID := c.Params("id")

	subscription, err := h.service.GetByID(subscriptionID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Subscription not found",
		})
	}

	return c.JSON(subscription)
}

// GetPatientSubscriptions godoc
// @Summary List patient subscriptions
// @Description Returns all subscriptions for a specific patient
// @Tags Subscriptions
// @Produce json
// @Param id path string true "Patient ID"
// @Success 200 {array} models.PatientSubscription
// @Failure 500 {object} map[string]string
// @Router /api/v1/patients/{id}/subscriptions [get]
// @Security BearerAuth
func (h *SubscriptionHandler) GetPatientSubscriptions(c *fiber.Ctx) error {
	patientID := c.Params("id")

	subscriptions, err := h.service.GetByPatientID(patientID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch patient subscriptions",
		})
	}

	return c.JSON(subscriptions)
}

// CreateSubscription godoc
// @Summary Create a new subscription
// @Description Creates a new patient subscription
// @Tags Subscriptions
// @Accept json
// @Produce json
// @Param id path string true "Patient ID"
// @Param subscription body services.CreateSubscriptionDTO true "Subscription data"
// @Success 201 {object} models.PatientSubscription
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/patients/{id}/subscriptions [post]
// @Security BearerAuth
func (h *SubscriptionHandler) CreateSubscription(c *fiber.Ctx) error {
	patientID := c.Params("id")

	var dto services.CreateSubscriptionDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Override patientID from path
	dto.PatientID = patientID

	subscription, err := h.service.Create(&dto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create subscription",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(subscription)
}

// UpdateSubscription godoc
// @Summary Update a subscription
// @Description Updates an existing subscription
// @Tags Subscriptions
// @Accept json
// @Produce json
// @Param id path string true "Subscription ID"
// @Param subscription body services.UpdateSubscriptionDTO true "Subscription data"
// @Success 200 {object} models.PatientSubscription
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/subscriptions/{id} [put]
// @Security BearerAuth
func (h *SubscriptionHandler) UpdateSubscription(c *fiber.Ctx) error {
	subscriptionID := c.Params("id")

	var dto services.UpdateSubscriptionDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	subscription, err := h.service.Update(subscriptionID, &dto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update subscription",
		})
	}

	return c.JSON(subscription)
}

// DeleteSubscription godoc
// @Summary Delete a subscription
// @Description Soft deletes a subscription
// @Tags Subscriptions
// @Param id path string true "Subscription ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/subscriptions/{id} [delete]
// @Security BearerAuth
func (h *SubscriptionHandler) DeleteSubscription(c *fiber.Ctx) error {
	subscriptionID := c.Params("id")

	if err := h.service.Delete(subscriptionID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete subscription",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// CancelSubscription godoc
// @Summary Cancel a subscription
// @Description Cancels an active subscription
// @Tags Subscriptions
// @Accept json
// @Produce json
// @Param id path string true "Subscription ID"
// @Param cancel body services.CancelSubscriptionDTO true "Cancellation data"
// @Success 200 {object} models.PatientSubscription
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/subscriptions/{id}/cancel [post]
// @Security BearerAuth
func (h *SubscriptionHandler) CancelSubscription(c *fiber.Ctx) error {
	subscriptionID := c.Params("id")

	var dto services.CancelSubscriptionDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	subscription, err := h.service.Cancel(subscriptionID, &dto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to cancel subscription",
		})
	}

	return c.JSON(subscription)
}

// RenewSubscription godoc
// @Summary Renew a subscription
// @Description Renews an existing subscription
// @Tags Subscriptions
// @Accept json
// @Produce json
// @Param id path string true "Subscription ID"
// @Param renew body services.RenewSubscriptionDTO true "Renewal data"
// @Success 200 {object} models.PatientSubscription
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/subscriptions/{id}/renew [post]
// @Security BearerAuth
func (h *SubscriptionHandler) RenewSubscription(c *fiber.Ctx) error {
	subscriptionID := c.Params("id")

	var dto services.RenewSubscriptionDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	subscription, err := h.service.Renew(subscriptionID, &dto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to renew subscription",
		})
	}

	return c.JSON(subscription)
}

// GetActivePatientSubscription godoc
// @Summary Get active subscription for patient
// @Description Returns the active subscription for a specific patient with full method hierarchy
// @Tags Subscriptions
// @Produce json
// @Param id path string true "Patient ID"
// @Success 200 {object} models.PatientSubscription
// @Success 204 "No active subscription"
// @Failure 500 {object} map[string]string
// @Router /api/v1/patients/{id}/subscriptions/active [get]
// @Security BearerAuth
func (h *SubscriptionHandler) GetActivePatientSubscription(c *fiber.Ctx) error {
	patientID := c.Params("id")

	subscription, err := h.service.GetActiveByPatientID(patientID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch active subscription",
		})
	}

	if subscription == nil {
		return c.SendStatus(fiber.StatusNoContent)
	}

	return c.JSON(subscription)
}

// SuspendSubscription godoc
// @Summary Suspend a subscription
// @Description Suspends an active subscription
// @Tags Subscriptions
// @Produce json
// @Param id path string true "Subscription ID"
// @Success 200 {object} models.PatientSubscription
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/subscriptions/{id}/suspend [post]
// @Security BearerAuth
func (h *SubscriptionHandler) SuspendSubscription(c *fiber.Ctx) error {
	subscriptionID := c.Params("id")

	subscription, err := h.service.Suspend(subscriptionID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to suspend subscription",
		})
	}

	return c.JSON(subscription)
}

// ActivateSubscription godoc
// @Summary Activate a subscription
// @Description Activates a suspended or inactive subscription
// @Tags Subscriptions
// @Produce json
// @Param id path string true "Subscription ID"
// @Success 200 {object} models.PatientSubscription
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/subscriptions/{id}/activate [post]
// @Security BearerAuth
func (h *SubscriptionHandler) ActivateSubscription(c *fiber.Ctx) error {
	subscriptionID := c.Params("id")

	subscription, err := h.service.Activate(subscriptionID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to activate subscription",
		})
	}

	return c.JSON(subscription)
}
