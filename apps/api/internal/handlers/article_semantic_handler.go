package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/plenya/api/internal/services"
)

// ArticleSemanticHandler gerencia endpoints de busca semântica (RAG)
type ArticleSemanticHandler struct {
	service   *services.ArticleSemanticService
	validator *validator.Validate
}

// NewArticleSemanticHandler cria nova instância
func NewArticleSemanticHandler(
	service *services.ArticleSemanticService,
	validator *validator.Validate,
) *ArticleSemanticHandler {
	return &ArticleSemanticHandler{
		service:   service,
		validator: validator,
	}
}

// SemanticSearch godoc
// @Summary Semantic search articles
// @Description Busca artigos usando similaridade vetorial (RAG). Retorna artigos ordenados por relevância semântica.
// @Tags Articles - RAG
// @Produce json
// @Param q query string true "Query de busca (min 3 chars)" minlength(3) maxlength(1000)
// @Param limit query int false "Limite de resultados (1-100)" minimum(1) maximum(100) default(20)
// @Param minSimilarity query number false "Threshold de similaridade (0.0-1.0)" minimum(0) maximum(1) default(0.6)
// @Success 200 {array} repository.ArticleSearchResult
// @Failure 400 {object} map[string]interface{} "Invalid query parameters"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Security BearerAuth
// @Router /articles/semantic-search [get]
func (h *ArticleSemanticHandler) SemanticSearch(c *fiber.Ctx) error {
	println("========================================")
	println("DEBUG: SemanticSearch handler called!")
	println("DEBUG: Request path:", c.Path())
	println("DEBUG: Request method:", c.Method())
	println("========================================")

	// Parse query parameters
	query := c.Query("q")
	println("DEBUG: Received query:", query, "len:", len(query))

	if query == "" {
		println("DEBUG: Query is empty")
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "query parameter 'q' is required",
		})
	}

	// Validar query
	if err := h.service.ValidateQuery(query); err != nil {
		println("DEBUG: ValidateQuery failed:", err.Error())
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Parse optional parameters
	limit := c.QueryInt("limit", 20)

	// Parse minSimilarity manually (QueryFloat doesn't exist in Fiber v2)
	minSimilarity := 0.6
	if minSimStr := c.Query("minSimilarity"); minSimStr != "" {
		if parsed, err := strconv.ParseFloat(minSimStr, 64); err == nil {
			minSimilarity = parsed
		}
	}

	// Criar DTO
	dto := services.SemanticSearchDTO{
		Query:         query,
		Limit:         limit,
		MinSimilarity: minSimilarity,
	}

	println("DEBUG: DTO created - query:", dto.Query, "limit:", dto.Limit, "minSim:", dto.MinSimilarity)

	// Validar DTO
	if err := h.validator.Struct(dto); err != nil {
		println("DEBUG: DTO validation failed:", err.Error())
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	println("DEBUG: DTO validation passed, calling service...")

	// Executar busca semântica
	ctx := context.Background()
	results, err := h.service.SemanticSearch(ctx, dto)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Retornar resultados
	return c.JSON(fiber.Map{
		"query":   query,
		"count":   len(results),
		"results": results,
	})
}

// RecommendForScoreItem godoc
// @Summary Recommend articles for score item
// @Description Recomenda artigos relevantes para um parâmetro clínico (ScoreItem). Usa similaridade vetorial para encontrar artigos que cobrem o parâmetro.
// @Tags Articles - RAG
// @Produce json
// @Param scoreItemId query string true "ID do ScoreItem (UUID)" format(uuid)
// @Param limit query int false "Limite de resultados (1-100)" minimum(1) maximum(100) default(10)
// @Success 200 {array} repository.ArticleSearchResult
// @Failure 400 {object} map[string]interface{} "Invalid parameters"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Security BearerAuth
// @Router /articles/recommend [get]
func (h *ArticleSemanticHandler) RecommendForScoreItem(c *fiber.Ctx) error {
	// Parse scoreItemId
	scoreItemIDStr := c.Query("scoreItemId")
	if scoreItemIDStr == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "query parameter 'scoreItemId' is required",
		})
	}

	scoreItemID, err := uuid.Parse(scoreItemIDStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid scoreItemId format (must be UUID)",
		})
	}

	// Parse limit
	limit := c.QueryInt("limit", 10)

	// Executar recomendação
	results, err := h.service.RecommendArticlesForScoreItem(scoreItemID, limit)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Retornar resultados
	return c.JSON(fiber.Map{
		"scoreItemId": scoreItemID,
		"count":       len(results),
		"results":     results,
	})
}

// DiscoverRelatedScoreItems godoc
// @Summary Find score items covered by article
// @Description Descobre quais parâmetros clínicos (ScoreItems) um artigo cobre. Útil para sugerir vinculações automáticas.
// @Tags Articles - RAG
// @Produce json
// @Param id path string true "Article ID (UUID)" format(uuid)
// @Param limit query int false "Limite de resultados (1-100)" minimum(1) maximum(100) default(20)
// @Success 200 {array} repository.ScoreItemSimilarity
// @Failure 400 {object} map[string]interface{} "Invalid parameters"
// @Failure 404 {object} map[string]interface{} "Article not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Security BearerAuth
// @Router /articles/{id}/related-score-items [get]
func (h *ArticleSemanticHandler) DiscoverRelatedScoreItems(c *fiber.Ctx) error {
	// Parse article ID
	articleIDStr := c.Params("id")
	articleID, err := uuid.Parse(articleIDStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid article ID format",
		})
	}

	// Parse limit
	limit := c.QueryInt("limit", 20)

	// Executar descoberta
	results, err := h.service.DiscoverRelatedScoreItems(articleID, limit)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Retornar resultados
	return c.JSON(fiber.Map{
		"articleId": articleID,
		"count":     len(results),
		"results":   results,
	})
}

// GetEmbeddingStats godoc
// @Summary Get embedding statistics
// @Description Retorna estatísticas sobre embeddings gerados (artigos e score items).
// @Tags Articles - RAG
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Security BearerAuth
// @Router /articles/embedding-stats [get]
func (h *ArticleSemanticHandler) GetEmbeddingStats(c *fiber.Ctx) error {
	stats, err := h.service.GetEmbeddingStats()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(stats)
}
