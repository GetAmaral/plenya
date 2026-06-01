package handlers

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/middleware"
	"github.com/plenya/api/internal/services"
)

type RecallHandler struct {
	recallService *services.RecallService
	validator     *validator.Validate
}

func NewRecallHandler(recallService *services.RecallService) *RecallHandler {
	return &RecallHandler{
		recallService: recallService,
		validator:     validator.New(),
	}
}

// Create POST /api/v1/recalls
func (h *RecallHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateRecallRequest
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
	resp, err := h.recallService.Create(userID, &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(resp)
}

// List GET /api/v1/recalls?status=pending|all|scheduled|dismissed&dueBefore=YYYY-MM-DD
func (h *RecallHandler) List(c *fiber.Ctx) error {
	var dueBefore *time.Time
	if v := c.Query("dueBefore"); v != "" {
		if t, err := time.Parse("2006-01-02", v); err == nil {
			dueBefore = &t
		}
	}
	resp, err := h.recallService.List(c.Query("status"), dueBefore)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}
	return c.JSON(resp)
}

// Update PUT /api/v1/recalls/:id
func (h *RecallHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid recall id",
			Message: "id must be a valid UUID",
		})
	}
	var req dto.UpdateRecallRequest
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
	resp, err := h.recallService.Update(id, &req)
	if err != nil {
		if errors.Is(err, services.ErrRecallNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "recall not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}
	return c.JSON(resp)
}

// Delete DELETE /api/v1/recalls/:id
func (h *RecallHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid recall id",
			Message: "id must be a valid UUID",
		})
	}
	if err := h.recallService.Delete(id); err != nil {
		if errors.Is(err, services.ErrRecallNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "recall not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
