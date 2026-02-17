package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/services"
	"gorm.io/gorm"
)

// ScoreEnrichmentPreparationHandler gerencia endpoints de preparação de enrichment
type ScoreEnrichmentPreparationHandler struct {
	service *services.ScoreEnrichmentPreparationService
}

// NewScoreEnrichmentPreparationHandler cria nova instância
func NewScoreEnrichmentPreparationHandler(db *gorm.DB) *ScoreEnrichmentPreparationHandler {
	return &ScoreEnrichmentPreparationHandler{
		service: services.NewScoreEnrichmentPreparationService(db),
	}
}

// PrepareEnrichment prepara chunks para enrichment (ETAPA 1 - SEM IA)
// @Summary Preparar chunks científicos para enrichment
// @Description Busca top 20 chunks mais relevantes via RAG e salva no banco (sem IA)
// @Tags Score Items - Enrichment
// @Accept json
// @Produce json
// @Param id path string true "Score Item ID (UUID)"
// @Param request body dto.PrepareEnrichmentRequest false "Parâmetros opcionais (limit, minSimilarity)"
// @Success 200 {object} dto.PrepareEnrichmentResponse "Preparação criada com sucesso"
// @Failure 400 {object} ErrorResponse "Parâmetros inválidos"
// @Failure 404 {object} ErrorResponse "Score item não encontrado"
// @Failure 500 {object} ErrorResponse "Erro ao preparar chunks"
// @Router /api/v1/score-items/{id}/prepare-enrichment [post]
func (h *ScoreEnrichmentPreparationHandler) PrepareEnrichment(c *fiber.Ctx) error {
	// Parse score_item_id
	scoreItemID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid score item ID",
		})
	}

	// Parse request body (opcional)
	var req dto.PrepareEnrichmentRequest
	if err := c.BodyParser(&req); err != nil {
		// Body vazio é válido, usar defaults
		req = dto.PrepareEnrichmentRequest{}
	}

	// Defaults
	limit := 20
	minSimilarity := 0.65

	if req.Limit != nil {
		limit = *req.Limit
	}
	if req.MinSimilarity != nil {
		minSimilarity = *req.MinSimilarity
	}

	// Preparar chunks
	prep, err := h.service.PrepareChunks(scoreItemID, limit, minSimilarity)
	if err != nil {
		status := fiber.StatusInternalServerError
		if err.Error() == "score_item not found: "+scoreItemID.String() {
			status = fiber.StatusNotFound
		}
		return c.Status(status).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Converter para DTO
	response := dto.ToEnrichmentPreparationDTO(prep)

	return c.Status(fiber.StatusOK).JSON(response)
}

// GetPreparation busca preparação existente por score_item_id
// @Summary Obter preparação de enrichment
// @Description Retorna preparação de chunks existente para um score item
// @Tags Score Items - Enrichment
// @Produce json
// @Param id path string true "Score Item ID (UUID)"
// @Success 200 {object} dto.PrepareEnrichmentResponse "Preparação encontrada"
// @Failure 404 {object} ErrorResponse "Preparação não encontrada"
// @Failure 500 {object} ErrorResponse "Erro ao buscar preparação"
// @Router /api/v1/score-items/{id}/prepare-enrichment [get]
func (h *ScoreEnrichmentPreparationHandler) GetPreparation(c *fiber.Ctx) error {
	scoreItemID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid score item ID",
		})
	}

	prep, err := h.service.GetPreparation(scoreItemID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dto.ToEnrichmentPreparationDTO(prep)

	return c.Status(fiber.StatusOK).JSON(response)
}

// DeletePreparation remove preparação por score_item_id
// @Summary Deletar preparação de enrichment
// @Description Remove preparação de chunks de um score item
// @Tags Score Items - Enrichment
// @Produce json
// @Param id path string true "Score Item ID (UUID)"
// @Success 204 "Preparação deletada com sucesso"
// @Failure 400 {object} ErrorResponse "ID inválido"
// @Failure 500 {object} ErrorResponse "Erro ao deletar preparação"
// @Router /api/v1/score-items/{id}/prepare-enrichment [delete]
func (h *ScoreEnrichmentPreparationHandler) DeletePreparation(c *fiber.Ctx) error {
	scoreItemID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid score item ID",
		})
	}

	err = h.service.DeletePreparation(scoreItemID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// GetStats retorna estatísticas de preparações
// @Summary Estatísticas de preparações
// @Description Retorna contadores de preparações por status
// @Tags Score Items - Enrichment
// @Produce json
// @Success 200 {object} dto.EnrichmentPreparationStatsResponse "Estatísticas"
// @Failure 500 {object} ErrorResponse "Erro ao buscar estatísticas"
// @Router /api/v1/score-items/enrichment/stats [get]
func (h *ScoreEnrichmentPreparationHandler) GetStats(c *fiber.Ctx) error {
	stats, err := h.service.GetStats()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(stats)
}
