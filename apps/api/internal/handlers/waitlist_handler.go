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

type WaitlistHandler struct {
	waitlistService *services.WaitlistService
	validator       *validator.Validate
}

func NewWaitlistHandler(waitlistService *services.WaitlistService) *WaitlistHandler {
	return &WaitlistHandler{
		waitlistService: waitlistService,
		validator:       validator.New(),
	}
}

// Create POST /api/v1/waitlist
func (h *WaitlistHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateWaitlistEntryRequest
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
	userID := middleware.GetUserID(c)
	resp, err := h.waitlistService.Create(userID, &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(resp)
}

// List GET /api/v1/waitlist?status=waiting|all|scheduled|cancelled
func (h *WaitlistHandler) List(c *fiber.Ctx) error {
	resp, err := h.waitlistService.List(c.Query("status"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}
	return c.JSON(resp)
}

// Update PUT /api/v1/waitlist/:id
func (h *WaitlistHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid waitlist id",
			Message: "id must be a valid UUID",
		})
	}
	var req dto.UpdateWaitlistEntryRequest
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
	resp, err := h.waitlistService.Update(id, &req)
	if err != nil {
		if errors.Is(err, services.ErrWaitlistEntryNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "waitlist entry not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}
	return c.JSON(resp)
}

// Delete DELETE /api/v1/waitlist/:id
func (h *WaitlistHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid waitlist id",
			Message: "id must be a valid UUID",
		})
	}
	if err := h.waitlistService.Delete(id); err != nil {
		if errors.Is(err, services.ErrWaitlistEntryNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "waitlist entry not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
