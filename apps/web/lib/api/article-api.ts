'use client'

import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { apiClient } from '../api-client'

// Types
export interface ScoreGroup {
  id: string
  name: string
  description?: string
}

export interface ScoreSubgroup {
  id: string
  name: string
  description?: string
  scoreGroup?: ScoreGroup
}

export interface ScoreItem {
  id: string
  name: string
  description?: string
  unit?: string
  points: number
  scoreSubgroup?: ScoreSubgroup
}

export interface Article {
  id: string
  title: string
  authors: string
  journal: string
  publishDate: string
  language: string
  doi?: string
  pmid?: string
  issn?: string
  abstract?: string
  fullContent?: string
  notes?: string
  originalLink?: string
  internalLink?: string
  articleType: ArticleType
  keywords?: string[]
  meshTerms?: string[]
  specialty?: string
  favorite: boolean
  rating?: number
  fileHash?: string
  fileSize?: number
  indexedAt: string
  lastAccessedAt?: string
  createdBy?: string
  updatedBy?: string
  scoreItems?: ScoreItem[]
  createdAt: string
  updatedAt: string
}

export type ArticleType =
  | 'research_article'
  | 'review'
  | 'meta_analysis'
  | 'case_study'
  | 'clinical_trial'
  | 'editorial'
  | 'letter'
  | 'protocol'
  | 'lecture'

export const ARTICLE_TYPES: { value: ArticleType; label: string }[] = [
  { value: 'research_article', label: 'Artigo de Pesquisa' },
  { value: 'review', label: 'Revisão' },
  { value: 'meta_analysis', label: 'Meta-análise' },
  { value: 'case_study', label: 'Estudo de Caso' },
  { value: 'clinical_trial', label: 'Ensaio Clínico' },
  { value: 'editorial', label: 'Editorial' },
  { value: 'letter', label: 'Carta' },
  { value: 'protocol', label: 'Protocolo' },
  { value: 'lecture', label: 'Palestra' },
]

export interface CreateArticleDTO {
  title: string
  authors: string
  journal: string
  publishDate: string
  language: string
  doi?: string
  pmid?: string
  issn?: string
  abstract?: string
  fullContent?: string
  notes?: string
  originalLink?: string
  internalLink?: string
  articleType: ArticleType
  keywords?: string[]
  meshTerms?: string[]
  specialty?: string
  favorite?: boolean
  rating?: number
}

export interface UpdateArticleDTO {
  title?: string
  authors?: string
  journal?: string
  publishDate?: string
  language?: string
  doi?: string
  pmid?: string
  issn?: string
  abstract?: string
  fullContent?: string
  notes?: string
  originalLink?: string
  articleType?: ArticleType
  keywords?: string[]
  meshTerms?: string[]
  specialty?: string
  favorite?: boolean
  rating?: number
}

export interface ArticleFilters {
  journal?: string
  specialty?: string
  articleType?: ArticleType
  favorite?: boolean
  rating?: number
  publishedAfter?: string
  publishedBefore?: string
}

export interface PaginatedArticles {
  articles: Article[]
  pagination: {
    page: number
    pageSize: number
    total: number
    totalPages: number
  }
}

// API Functions
export const articleApi = {
  // List articles with filters
  list: async (
    page = 1,
    pageSize = 20,
    filters?: ArticleFilters
  ): Promise<PaginatedArticles> => {
    const params = new URLSearchParams({
      page: page.toString(),
      pageSize: pageSize.toString(),
    })

    if (filters?.journal) params.append('journal', filters.journal)
    if (filters?.specialty) params.append('specialty', filters.specialty)
    if (filters?.articleType) params.append('articleType', filters.articleType)
    if (filters?.favorite !== undefined)
      params.append('favorite', filters.favorite.toString())
    if (filters?.rating !== undefined)
      params.append('rating', filters.rating.toString())
    if (filters?.publishedAfter)
      params.append('publishedAfter', filters.publishedAfter)
    if (filters?.publishedBefore)
      params.append('publishedBefore', filters.publishedBefore)

    return apiClient.get<PaginatedArticles>(`/api/v1/articles?${params.toString()}`)
  },

  // Get article by ID
  getById: async (id: string): Promise<Article> => {
    return apiClient.get<Article>(`/api/v1/articles/${id}`)
  },

  // Create article
  create: async (data: CreateArticleDTO): Promise<Article> => {
    return apiClient.post<Article>('/api/v1/articles', data)
  },

  // Update article
  update: async (id: string, data: UpdateArticleDTO): Promise<Article> => {
    return apiClient.put<Article>(`/api/v1/articles/${id}`, data)
  },

  // Delete article
  delete: async (id: string): Promise<void> => {
    await apiClient.delete(`/api/v1/articles/${id}`)
  },

  // Upload PDF
  upload: async (file: File): Promise<Article> => {
    const formData = new FormData()
    formData.append('file', file)

    return apiClient.post<Article>('/api/v1/articles/upload', formData)
  },

  // Search articles
  search: async (
    query: string,
    page = 1,
    pageSize = 20
  ): Promise<PaginatedArticles> => {
    const params = new URLSearchParams({
      q: query,
      page: page.toString(),
      pageSize: pageSize.toString(),
    })

    return apiClient.get<PaginatedArticles>(`/api/v1/articles/search?${params.toString()}`)
  },

  // Get favorites
  getFavorites: async (page = 1, pageSize = 20): Promise<PaginatedArticles> => {
    const params = new URLSearchParams({
      page: page.toString(),
      pageSize: pageSize.toString(),
    })

    return apiClient.get<PaginatedArticles>(`/api/v1/articles/favorites?${params.toString()}`)
  },

  // Toggle favorite
  toggleFavorite: async (id: string): Promise<Article> => {
    return apiClient.patch<Article>(`/api/v1/articles/${id}/favorite`)
  },

  // Set rating
  setRating: async (id: string, rating: number): Promise<Article> => {
    return apiClient.patch<Article>(`/api/v1/articles/${id}/rating`, { rating })
  },

  // Get download URL
  getDownloadUrl: (id: string): string => {
    return `${process.env.NEXT_PUBLIC_API_URL}/api/v1/articles/${id}/download`
  },
}

// React Query Hooks

// List articles
export function useArticles(
  page = 1,
  pageSize = 20,
  filters?: ArticleFilters
) {
  return useQuery({
    queryKey: ['articles', page, pageSize, filters],
    queryFn: () => articleApi.list(page, pageSize, filters),
  })
}

// Get article by ID
export function useArticle(id: string) {
  return useQuery({
    queryKey: ['articles', id],
    queryFn: () => articleApi.getById(id),
    enabled: !!id,
  })
}

// Search articles
export function useSearchArticles(query: string, page = 1, pageSize = 20) {
  return useQuery({
    queryKey: ['articles', 'search', query, page, pageSize],
    queryFn: () => articleApi.search(query, page, pageSize),
    enabled: query.length >= 2,
  })
}

// Get favorites
export function useFavoriteArticles(page = 1, pageSize = 20) {
  return useQuery({
    queryKey: ['articles', 'favorites', page, pageSize],
    queryFn: () => articleApi.getFavorites(page, pageSize),
  })
}

// Create article
export function useCreateArticle() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: articleApi.create,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['articles'] })
    },
  })
}

// Update article
export function useUpdateArticle() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: UpdateArticleDTO }) =>
      articleApi.update(id, data),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: ['articles'] })
      queryClient.invalidateQueries({ queryKey: ['articles', variables.id] })
    },
  })
}

// Delete article
export function useDeleteArticle() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: articleApi.delete,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['articles'] })
    },
  })
}

// Upload PDF
export function useUploadArticle() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: articleApi.upload,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['articles'] })
    },
  })
}

// Toggle favorite
export function useToggleFavorite() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: articleApi.toggleFavorite,
    onSuccess: (data) => {
      queryClient.invalidateQueries({ queryKey: ['articles'] })
      queryClient.invalidateQueries({ queryKey: ['articles', data.id] })
      queryClient.invalidateQueries({ queryKey: ['articles', 'favorites'] })
    },
  })
}

// Set rating
export function useSetRating() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({ id, rating }: { id: string; rating: number }) =>
      articleApi.setRating(id, rating),
    onSuccess: (data) => {
      queryClient.invalidateQueries({ queryKey: ['articles'] })
      queryClient.invalidateQueries({ queryKey: ['articles', data.id] })
    },
  })
}
