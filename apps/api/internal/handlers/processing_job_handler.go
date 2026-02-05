package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/plenya/api/internal/services"
)

type ProcessingJobHandler struct {
	processingJobService *services.ProcessingJobService
}

func NewProcessingJobHandler(service *services.ProcessingJobService) *ProcessingJobHandler {
	return &ProcessingJobHandler{processingJobService: service}
}

// GetByID busca status do job de processamento
// @Summary Buscar status de job de processamento
// @Description Busca informações e status atual de um job de processamento assíncrono (ex: extração de PDF)
// @Tags ProcessingJobs
// @Accept json
// @Produce json
// @Param id path string true "Job ID (UUID)"
// @Success 200 {object} dto.ProcessingJobResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /api/v1/processing-jobs/{id} [get]
func (h *ProcessingJobHandler) GetByID(c *fiber.Ctx) error {
	jobID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid job id",
		})
	}

	job, err := h.processingJobService.GetByID(jobID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "job not found",
		})
	}

	return c.JSON(job)
}
