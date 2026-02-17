import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { apiClient } from '../api-client'
import type { Article, ScoreItem } from '@repo/types'

// Types para resultados de busca semântica
export interface ArticleSearchResult {
  article: Article
  similarity: number // 0.0 - 1.0 (cosine similarity)
  chunkText: string // Trecho do artigo que teve match
}

export interface ScoreItemSimilarity {
  scoreItem: ScoreItem
  similarity: number // 0.0 - 1.0 (cosine similarity)
}

// DTOs para requests
export interface SemanticSearchParams {
  q: string // Query de busca (min 3 chars)
  limit?: number // Default: 20
  minSimilarity?: number // Default: 0.4
}

export interface RecommendArticlesParams {
  scoreItemId: string // UUID do ScoreItem
  limit?: number // Default: 10
}

export interface DiscoverRelatedScoreItemsParams {
  articleId: string // UUID do Article
  limit?: number // Default: 20
}

// Response wrapper para semantic search
export interface SemanticSearchResponse {
  query: string
  count: number
  results: ArticleSearchResult[]
}

export interface RecommendArticlesResponse {
  scoreItemId: string
  count: number
  results: ArticleSearchResult[]
}

export interface DiscoverRelatedScoreItemsResponse {
  articleId: string
  count: number
  results: ScoreItemSimilarity[]
}

export interface EmbeddingStatsResponse {
  article_embeddings: number
  score_item_embeddings: number
  articles_with_embeddings: number
  total_articles: number
  coverage_percent: number
}

// Query Keys
export const semanticSearchKeys = {
  all: ['semantic-search'] as const,
  search: (query: string) => [...semanticSearchKeys.all, 'search', query] as const,
  recommend: (scoreItemId: string) => [...semanticSearchKeys.all, 'recommend', scoreItemId] as const,
  relatedScoreItems: (articleId: string) => [...semanticSearchKeys.all, 'related-score-items', articleId] as const,
  stats: () => [...semanticSearchKeys.all, 'stats'] as const,
}

// Hooks

/**
 * Busca semântica de artigos por query natural
 * Usa similaridade vetorial (RAG) para encontrar artigos relevantes
 *
 * @example
 * const { data, isLoading } = useSemanticSearch('diabetes treatment')
 */
export function useSemanticSearch(
  query: string,
  options?: {
    limit?: number
    minSimilarity?: number
    enabled?: boolean
  }
) {
  const { limit = 20, minSimilarity = 0.4, enabled = true } = options || {}

  return useQuery({
    queryKey: semanticSearchKeys.search(query),
    queryFn: async () => {
      const params = new URLSearchParams({
        q: query,
        limit: limit.toString(),
        minSimilarity: minSimilarity.toString(),
      })

      const response = await apiClient.get<SemanticSearchResponse>(
        `/api/v1/articles/semantic-search?${params}`
      )
      return response
    },
    enabled: enabled && query.length >= 3,
    staleTime: 5 * 60 * 1000, // 5 minutes
  })
}

/**
 * Recomenda artigos para um ScoreItem específico
 * Útil para sugerir evidências científicas ao editar parâmetros clínicos
 *
 * @example
 * const { data } = useRecommendedArticles(scoreItemId)
 */
export function useRecommendedArticles(
  scoreItemId: string | undefined,
  options?: {
    limit?: number
    enabled?: boolean
  }
) {
  const { limit = 10, enabled = true } = options || {}

  return useQuery({
    queryKey: semanticSearchKeys.recommend(scoreItemId || ''),
    queryFn: async () => {
      const params = new URLSearchParams({
        scoreItemId: scoreItemId!,
        limit: limit.toString(),
      })

      const response = await apiClient.get<RecommendArticlesResponse>(
        `/api/v1/articles/recommend?${params}`
      )
      return response
    },
    enabled: enabled && !!scoreItemId,
    staleTime: 5 * 60 * 1000,
  })
}

/**
 * Descobre quais ScoreItems um artigo cobre
 * Útil para sugerir vinculações automáticas
 *
 * @example
 * const { data } = useRelatedScoreItems(articleId)
 */
export function useRelatedScoreItems(
  articleId: string | undefined,
  options?: {
    limit?: number
    enabled?: boolean
  }
) {
  const { limit = 20, enabled = true } = options || {}

  return useQuery({
    queryKey: semanticSearchKeys.relatedScoreItems(articleId || ''),
    queryFn: async () => {
      const response = await apiClient.get<DiscoverRelatedScoreItemsResponse>(
        `/api/v1/articles/${articleId}/related-score-items?limit=${limit}`
      )
      return response
    },
    enabled: enabled && !!articleId,
    staleTime: 5 * 60 * 1000,
  })
}

/**
 * Retorna estatísticas sobre embeddings gerados
 *
 * @example
 * const { data } = useEmbeddingStats()
 */
export function useEmbeddingStats() {
  return useQuery({
    queryKey: semanticSearchKeys.stats(),
    queryFn: () => apiClient.get<EmbeddingStatsResponse>('/api/v1/articles/embedding-stats'),
    staleTime: 10 * 60 * 1000, // 10 minutes
  })
}
