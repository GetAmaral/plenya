package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/services"
)

type ClinicalTimelineHandler struct {
	service *services.ClinicalTimelineService
}

func NewClinicalTimelineHandler(service *services.ClinicalTimelineService) *ClinicalTimelineHandler {
	return &ClinicalTimelineHandler{service: service}
}

// GetItemHistory godoc
// @Summary      Histórico longitudinal de um item clínico
// @Description  Retorna o histórico de um ScoreItem para o paciente, fundindo Escore Light, pré-consulta e anamnese, ordenado do mais recente ao mais antigo.
// @Tags         clinical-timeline
// @Produce      json
// @Param        id           path      string  true  "ID do paciente"
// @Param        scoreItemId  path      string  true  "ID do ScoreItem"
// @Success      200  {object}  dto.ClinicalItemHistoryResponse
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /patients/{id}/score-items/{scoreItemId}/history [get]
func (h *ClinicalTimelineHandler) GetItemHistory(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid patient id",
			Message: "patient id must be a valid UUID",
		})
	}

	scoreItemID, err := uuid.Parse(c.Params("scoreItemId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid score item id",
			Message: "score item id must be a valid UUID",
		})
	}

	entries, err := h.service.GetItemHistory(patientID, scoreItemID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.JSON(dto.ClinicalItemHistoryResponse{
		ScoreItemID: scoreItemID.String(),
		Entries:     entries,
	})
}
