package services

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/plenya/api/internal/repository"
)

// ArticleSemanticService gerencia operações de busca semântica (RAG)
// Usa EmbeddingService + ArticleVectorRepository
type ArticleSemanticService struct {
	vectorRepo       *repository.ArticleVectorRepository
	embeddingService *EmbeddingService
}

// NewArticleSemanticService cria nova instância
func NewArticleSemanticService(
	vectorRepo *repository.ArticleVectorRepository,
	embeddingService *EmbeddingService,
) *ArticleSemanticService {
	return &ArticleSemanticService{
		vectorRepo:       vectorRepo,
		embeddingService: embeddingService,
	}
}

// SemanticSearchDTO representa parâmetros de busca semântica
type SemanticSearchDTO struct {
	Query         string  `json:"query" validate:"required,min=3"`
	Limit         int     `json:"limit" validate:"omitempty,gte=1,lte=100"`
	MinSimilarity float64 `json:"minSimilarity" validate:"omitempty,gte=0,lte=1"`
}

// SemanticSearch busca artigos por similaridade semântica
// Gera embedding da query e busca artigos similares
func (s *ArticleSemanticService) SemanticSearch(
	ctx context.Context,
	dto SemanticSearchDTO,
) ([]repository.ArticleSearchResult, error) {
	// Validar entrada
	if dto.Query == "" || len(dto.Query) < 3 {
		return nil, fmt.Errorf("query must be at least 3 characters")
	}

	// Defaults
	if dto.Limit <= 0 {
		dto.Limit = 20
	}
	if dto.MinSimilarity <= 0 {
		dto.MinSimilarity = 0.6 // Threshold padrão para busca geral
	}

	// 1. Gerar embedding da query
	embedding, err := s.embeddingService.GenerateEmbedding(ctx, dto.Query)
	if err != nil {
		return nil, fmt.Errorf("failed to generate query embedding: %w", err)
	}

	// 2. Buscar artigos similares
	results, err := s.vectorRepo.SemanticSearch(embedding, dto.Limit, dto.MinSimilarity)
	if err != nil {
		return nil, fmt.Errorf("semantic search failed: %w", err)
	}

	return results, nil
}

// RecommendArticlesForScoreItem recomenda artigos para um parâmetro clínico
// Busca artigos que cobrem o score item baseado em similaridade vetorial
func (s *ArticleSemanticService) RecommendArticlesForScoreItem(
	scoreItemID uuid.UUID,
	limit int,
) ([]repository.ArticleSearchResult, error) {
	// Defaults
	if limit <= 0 {
		limit = 10
	}

	minSimilarity := 0.7 // Threshold mais alto para recomendações

	// Buscar artigos similares ao score item
	results, err := s.vectorRepo.FindSimilarArticlesForScoreItem(scoreItemID, limit, minSimilarity)
	if err != nil {
		return nil, fmt.Errorf("failed to recommend articles: %w", err)
	}

	return results, nil
}

// DiscoverRelatedScoreItems descobre quais parâmetros clínicos um artigo cobre
// Útil para sugerir vinculações automáticas
func (s *ArticleSemanticService) DiscoverRelatedScoreItems(
	articleID uuid.UUID,
	limit int,
) ([]repository.ScoreItemSimilarity, error) {
	// Defaults
	if limit <= 0 {
		limit = 20
	}

	minSimilarity := 0.7 // Threshold alto para auto-linking

	// Buscar score items similares ao artigo
	results, err := s.vectorRepo.FindSimilarScoreItems(articleID, limit, minSimilarity)
	if err != nil {
		return nil, fmt.Errorf("failed to discover related score items: %w", err)
	}

	return results, nil
}

// GetEmbeddingStats retorna estatísticas sobre embeddings
func (s *ArticleSemanticService) GetEmbeddingStats() (map[string]interface{}, error) {
	return s.vectorRepo.GetEmbeddingStats()
}

// ValidateQuery valida query de busca semântica
func (s *ArticleSemanticService) ValidateQuery(query string) error {
	if query == "" {
		return fmt.Errorf("query cannot be empty")
	}

	if len(query) < 3 {
		return fmt.Errorf("query must be at least 3 characters")
	}

	if len(query) > 1000 {
		return fmt.Errorf("query too long (max 1000 characters)")
	}

	return nil
}
