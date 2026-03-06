package handlers

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/services"
)

type TrainingAIHandler struct {
	service   *services.TrainingAIService
	validator *validator.Validate
}

func NewTrainingAIHandler(service *services.TrainingAIService) *TrainingAIHandler {
	return &TrainingAIHandler{
		service:   service,
		validator: validator.New(),
	}
}

// Chat handles AI chat for training context
func (h *TrainingAIHandler) Chat(c *fiber.Ctx) error {
	var req services.ChatRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid request body",
			Message: err.Error(),
		})
	}

	if len(req.Messages) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "at least one message is required",
		})
	}

	resp, err := h.service.Chat(context.Background(), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "ai chat failed",
			Message: err.Error(),
		})
	}

	return c.JSON(resp)
}

// Recommendations generates AI workout recommendations
func (h *TrainingAIHandler) Recommendations(c *fiber.Ctx) error {
	var req struct {
		PatientContext *services.PatientContext `json:"patientContext"`
		Objective      string                  `json:"objective" validate:"required"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid request body",
			Message: err.Error(),
		})
	}

	if req.Objective == "" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "objective is required",
		})
	}

	recommendation, err := h.service.GenerateWorkoutRecommendation(
		context.Background(),
		req.PatientContext,
		req.Objective,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "recommendation generation failed",
			Message: err.Error(),
		})
	}

	return c.JSON(fiber.Map{"recommendation": recommendation})
}
