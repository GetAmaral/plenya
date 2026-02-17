package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pgvector/pgvector-go"
	"github.com/plenya/api/internal/models"
	"gorm.io/gorm"
)

// ArticleVectorRepository gerencia operações de busca vetorial
// Separado do ArticleRepository para manter responsabilidades claras
type ArticleVectorRepository struct {
	db *gorm.DB
}

// NewArticleVectorRepository cria nova instância
func NewArticleVectorRepository(db *gorm.DB) *ArticleVectorRepository {
	return &ArticleVectorRepository{db: db}
}

// ArticleSearchResult representa um resultado de busca semântica
type ArticleSearchResult struct {
	Article    *models.Article `json:"article"`
	Similarity float64         `json:"similarity"` // Cosine similarity (0.0 - 1.0)
	ChunkText  string          `json:"chunkText"`  // Texto do chunk que teve match
}

// ChunkSearchResult representa um chunk individual (não artigo completo)
// Usado para enrichment científico de alta precisão
type ChunkSearchResult struct {
	ArticleID     uuid.UUID `json:"articleId"`
	ArticleTitle  string    `json:"articleTitle"`
	ArticleYear   int       `json:"articleYear"`
	Journal       string    `json:"journal"`
	ChunkIndex    int       `json:"chunkIndex"`
	ChunkText     string    `json:"chunkText"`
	Section       string    `json:"section"`       // results, discussion, methods, introduction, etc
	Similarity    float64   `json:"similarity"`    // 0.0-1.0 (weighted with section boost)
	WordCount     int       `json:"wordCount"`
	PageRange     *string   `json:"pageRange,omitempty"`
}

// SemanticSearch busca artigos por similaridade vetorial
// queryEmbedding: vetor de embedding da query (1024 dims)
// limit: quantidade máxima de resultados
// minSimilarity: threshold de similaridade (0.0 - 1.0), recomendado: 0.4
func (r *ArticleVectorRepository) SemanticSearch(
	queryEmbedding []float32,
	limit int,
	minSimilarity float64,
) ([]ArticleSearchResult, error) {
	// Validar entrada
	if len(queryEmbedding) != 1024 {
		return nil, fmt.Errorf("invalid embedding dimensions: got %d, expected 1024", len(queryEmbedding))
	}

	if limit <= 0 || limit > 100 {
		limit = 20 // Default
	}

	if minSimilarity < 0 || minSimilarity > 1 {
		minSimilarity = 0.4 // Default
	}

	// Converter []float32 para pgvector.Vector
	vector := pgvector.NewVector(queryEmbedding)

	// Query SQL:
	// 1. Calcular similaridade cosine (1 - cosine_distance)
	// 2. Filtrar por threshold
	// 3. Agrupar por article_id (DISTINCT ON - pega o chunk mais similar)
	// 4. Ordenar por similaridade DESC
	// 5. Limitar resultados
	type queryResult struct {
		ArticleID  uuid.UUID
		Similarity float64
		ChunkText  string
	}

	var results []queryResult

	// Raw SQL query usando pgvector <=> operator (cosine distance)
	// DISTINCT ON garante apenas 1 resultado por artigo (o chunk mais similar)
	// Usa CTE para primeiro fazer DISTINCT ON e depois ordenar por similaridade
	err := r.db.Raw(`
		WITH ranked_chunks AS (
			SELECT DISTINCT ON (article_id)
				article_id,
				1 - (embedding <=> ?::vector) AS similarity,
				chunk_text
			FROM article_embeddings
			WHERE 1 - (embedding <=> ?::vector) >= ?
			ORDER BY article_id, similarity DESC
		)
		SELECT article_id, similarity, chunk_text
		FROM ranked_chunks
		ORDER BY similarity DESC
		LIMIT ?
	`, vector, vector, minSimilarity, limit).Scan(&results).Error

	if err != nil {
		return nil, fmt.Errorf("semantic search query failed: %w", err)
	}

	// Se nenhum resultado, retornar slice vazio
	if len(results) == 0 {
		return []ArticleSearchResult{}, nil
	}

	// Buscar artigos completos
	articleIDs := make([]uuid.UUID, len(results))
	for i, r := range results {
		articleIDs[i] = r.ArticleID
	}

	var articles []models.Article
	err = r.db.
		Preload("ScoreItems.Subgroup.Group").
		Where("id IN ?", articleIDs).
		Find(&articles).Error

	if err != nil {
		return nil, fmt.Errorf("failed to fetch articles: %w", err)
	}

	// Criar mapa article_id -> article para lookup eficiente
	articleMap := make(map[uuid.UUID]*models.Article)
	for i := range articles {
		articleMap[articles[i].ID] = &articles[i]
	}

	// Montar resultados finais
	searchResults := make([]ArticleSearchResult, 0, len(results))
	for _, result := range results {
		if article, ok := articleMap[result.ArticleID]; ok {
			searchResults = append(searchResults, ArticleSearchResult{
				Article:    article,
				Similarity: result.Similarity,
				ChunkText:  result.ChunkText,
			})
		}
	}

	return searchResults, nil
}

// FindSimilarScoreItems descobre quais score items um artigo cobre
// Calcula embedding médio do artigo e compara com score_item_embeddings
func (r *ArticleVectorRepository) FindSimilarScoreItems(
	articleID uuid.UUID,
	limit int,
	minSimilarity float64,
) ([]ScoreItemSimilarity, error) {
	if limit <= 0 || limit > 100 {
		limit = 20
	}

	if minSimilarity < 0 || minSimilarity > 1 {
		minSimilarity = 0.7 // Threshold mais restritivo para auto-linking
	}

	type queryResult struct {
		ScoreItemID uuid.UUID
		Similarity  float64
	}

	var results []queryResult

	// Query SQL:
	// 1. Calcular embedding médio do artigo (AVG de todos chunks)
	// 2. Comparar com score_item_embeddings
	// 3. Retornar top N score items com similarity > threshold
	err := r.db.Raw(`
		WITH article_avg_embedding AS (
			SELECT AVG(embedding) AS avg_embedding
			FROM article_embeddings
			WHERE article_id = ?
		)
		SELECT
			score_item_id,
			1 - (sie.embedding <=> aae.avg_embedding) AS similarity
		FROM score_item_embeddings sie, article_avg_embedding aae
		WHERE 1 - (sie.embedding <=> aae.avg_embedding) >= ?
		ORDER BY similarity DESC
		LIMIT ?
	`, articleID, minSimilarity, limit).Scan(&results).Error

	if err != nil {
		return nil, fmt.Errorf("find similar score items query failed: %w", err)
	}

	if len(results) == 0 {
		return []ScoreItemSimilarity{}, nil
	}

	// Buscar score items completos
	scoreItemIDs := make([]uuid.UUID, len(results))
	for i, r := range results {
		scoreItemIDs[i] = r.ScoreItemID
	}

	var scoreItems []models.ScoreItem
	err = r.db.
		Preload("Subgroup.Group").
		Preload("Levels", func(db *gorm.DB) *gorm.DB {
			return db.Order("level ASC")
		}).
		Where("id IN ?", scoreItemIDs).
		Find(&scoreItems).Error

	if err != nil {
		return nil, fmt.Errorf("failed to fetch score items: %w", err)
	}

	// Criar mapa para lookup
	scoreItemMap := make(map[uuid.UUID]*models.ScoreItem)
	for i := range scoreItems {
		scoreItemMap[scoreItems[i].ID] = &scoreItems[i]
	}

	// Montar resultados
	similarities := make([]ScoreItemSimilarity, 0, len(results))
	for _, result := range results {
		if scoreItem, ok := scoreItemMap[result.ScoreItemID]; ok {
			similarities = append(similarities, ScoreItemSimilarity{
				ScoreItem:  scoreItem,
				Similarity: result.Similarity,
			})
		}
	}

	return similarities, nil
}

// ScoreItemSimilarity representa similaridade entre artigo e score item
type ScoreItemSimilarity struct {
	ScoreItem  *models.ScoreItem `json:"scoreItem"`
	Similarity float64           `json:"similarity"`
}

// FindSimilarArticlesForScoreItem encontra artigos relacionados a um score item
// Busca por similaridade entre score_item_embedding e article_embeddings
func (r *ArticleVectorRepository) FindSimilarArticlesForScoreItem(
	scoreItemID uuid.UUID,
	limit int,
	minSimilarity float64,
) ([]ArticleSearchResult, error) {
	if limit <= 0 || limit > 100 {
		limit = 10
	}

	if minSimilarity < 0 || minSimilarity > 1 {
		minSimilarity = 0.7 // Threshold mais alto para recomendações
	}

	type queryResult struct {
		ArticleID  uuid.UUID
		Similarity float64
		ChunkText  string
	}

	var results []queryResult

	// Query SQL:
	// 1. Buscar embedding do score item
	// 2. Comparar com article_embeddings
	// 3. Agrupar por article_id (DISTINCT ON - chunk mais similar)
	// 4. Retornar top N artigos ordenados por similaridade
	err := r.db.Raw(`
		WITH ranked_chunks AS (
			SELECT DISTINCT ON (ae.article_id)
				ae.article_id,
				1 - (ae.embedding <=> sie.embedding) AS similarity,
				ae.chunk_text
			FROM article_embeddings ae
			CROSS JOIN score_item_embeddings sie
			WHERE sie.score_item_id = ?
			  AND 1 - (ae.embedding <=> sie.embedding) >= ?
			ORDER BY ae.article_id, similarity DESC
		)
		SELECT article_id, similarity, chunk_text
		FROM ranked_chunks
		ORDER BY similarity DESC
		LIMIT ?
	`, scoreItemID, minSimilarity, limit).Scan(&results).Error

	if err != nil {
		return nil, fmt.Errorf("find similar articles query failed: %w", err)
	}

	if len(results) == 0 {
		return []ArticleSearchResult{}, nil
	}

	// Buscar artigos completos
	articleIDs := make([]uuid.UUID, len(results))
	for i, r := range results {
		articleIDs[i] = r.ArticleID
	}

	var articles []models.Article
	err = r.db.
		Preload("ScoreItems.Subgroup.Group").
		Where("id IN ?", articleIDs).
		Find(&articles).Error

	if err != nil {
		return nil, fmt.Errorf("failed to fetch articles: %w", err)
	}

	// Criar mapa
	articleMap := make(map[uuid.UUID]*models.Article)
	for i := range articles {
		articleMap[articles[i].ID] = &articles[i]
	}

	// Montar resultados
	searchResults := make([]ArticleSearchResult, 0, len(results))
	for _, result := range results {
		if article, ok := articleMap[result.ArticleID]; ok {
			searchResults = append(searchResults, ArticleSearchResult{
				Article:    article,
				Similarity: result.Similarity,
				ChunkText:  result.ChunkText,
			})
		}
	}

	return searchResults, nil
}

// GetEmbeddingStats retorna estatísticas sobre embeddings
func (r *ArticleVectorRepository) GetEmbeddingStats() (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	// Count total de article embeddings
	var articleEmbeddingCount int64
	r.db.Model(&models.ArticleEmbedding{}).Count(&articleEmbeddingCount)
	stats["article_embeddings"] = articleEmbeddingCount

	// Count total de score item embeddings
	var scoreItemEmbeddingCount int64
	r.db.Model(&models.ScoreItemEmbedding{}).Count(&scoreItemEmbeddingCount)
	stats["score_item_embeddings"] = scoreItemEmbeddingCount

	// Count de artigos com embeddings
	var articlesWithEmbeddings int64
	r.db.Model(&models.Article{}).Where("embedding_status = ?", "completed").Count(&articlesWithEmbeddings)
	stats["articles_with_embeddings"] = articlesWithEmbeddings

	// Count total de artigos
	var totalArticles int64
	r.db.Model(&models.Article{}).Count(&totalArticles)
	stats["total_articles"] = totalArticles

	// Percentage de cobertura
	if totalArticles > 0 {
		stats["coverage_percent"] = float64(articlesWithEmbeddings) / float64(totalArticles) * 100
	} else {
		stats["coverage_percent"] = 0
	}

	return stats, nil
}

// FindTopChunksForScoreItem busca os chunks mais relevantes para enrichment científico
// Retorna múltiplos chunks por artigo (remove DISTINCT ON)
// Aplica boost por seção (results: +10%, discussion: +5%, methods: +2%, introduction: +1%)
// scoreItemID: ID do score item a enriquecer
// limit: Total de chunks (ex: 20)
// minSimilarity: Threshold mínimo (0.65 mais permissivo para pegar mais contexto)
func (r *ArticleVectorRepository) FindTopChunksForScoreItem(
	scoreItemID uuid.UUID,
	limit int,
	minSimilarity float64,
) ([]ChunkSearchResult, error) {
	if limit <= 0 || limit > 100 {
		limit = 20 // Default
	}

	if minSimilarity < 0 || minSimilarity > 1 {
		minSimilarity = 0.65 // Mais permissivo que FindSimilarArticlesForScoreItem (0.7)
	}

	type queryResult struct {
		ArticleID    uuid.UUID
		ArticleTitle string
		ArticleYear  int
		Journal      string
		ChunkIndex   int
		ChunkText    string
		Section      string
		Similarity   float64
		WordCount    int
		PageRange    *string
	}

	var results []queryResult

	// SQL com boost por seção (priorizar results e discussion)
	// Usa COALESCE para seção default 'other' se metadata não tiver section
	err := r.db.Raw(`
		WITH scored_chunks AS (
			SELECT
				ae.article_id,
				a.title AS article_title,
				EXTRACT(YEAR FROM a.publish_date)::int AS article_year,
				a.journal,
				ae.chunk_index,
				ae.chunk_text,
				COALESCE(ae.chunk_metadata->>'section', 'other') AS section,
				COALESCE((ae.chunk_metadata->>'word_count')::int, 0) AS word_count,
				ae.chunk_metadata->>'page_range' AS page_range,
				1 - (ae.embedding <=> sie.embedding) AS base_similarity,
				-- Boost por seção (priorizar results e discussion para dados científicos)
				CASE COALESCE(ae.chunk_metadata->>'section', 'other')
					WHEN 'results' THEN 1.10      -- +10% (dados quantitativos)
					WHEN 'discussion' THEN 1.05   -- +5% (interpretação clínica)
					WHEN 'methods' THEN 1.02      -- +2% (protocolos)
					WHEN 'introduction' THEN 1.01 -- +1% (contexto)
					ELSE 1.0
				END AS section_weight
			FROM article_embeddings ae
			JOIN articles a ON ae.article_id = a.id
			CROSS JOIN score_item_embeddings sie
			WHERE sie.score_item_id = ?
			  AND a.deleted_at IS NULL
			  AND 1 - (ae.embedding <=> sie.embedding) >= ?
		)
		SELECT
			article_id,
			article_title,
			article_year,
			journal,
			chunk_index,
			chunk_text,
			section,
			word_count,
			page_range,
			base_similarity * section_weight AS similarity
		FROM scored_chunks
		ORDER BY similarity DESC
		LIMIT ?
	`, scoreItemID, minSimilarity, limit).Scan(&results).Error

	if err != nil {
		return nil, fmt.Errorf("find top chunks failed: %w", err)
	}

	if len(results) == 0 {
		return []ChunkSearchResult{}, nil
	}

	// Converter para ChunkSearchResult
	chunks := make([]ChunkSearchResult, len(results))
	for i, r := range results {
		chunks[i] = ChunkSearchResult{
			ArticleID:    r.ArticleID,
			ArticleTitle: r.ArticleTitle,
			ArticleYear:  r.ArticleYear,
			Journal:      r.Journal,
			ChunkIndex:   r.ChunkIndex,
			ChunkText:    r.ChunkText,
			Section:      r.Section,
			Similarity:   r.Similarity,
			WordCount:    r.WordCount,
			PageRange:    r.PageRange,
		}
	}

	return chunks, nil
}
