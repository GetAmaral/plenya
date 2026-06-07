import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { apiClient } from '../api-client'
import type { ScoreItem } from './score-api'

// ============================================================
// Types — espelham apps/api/internal/models/score_version*.go + DTOs
// ============================================================

export interface ScoreVersionItem {
  id: string
  versionId: string
  scoreItemId: string
  displayOrder: number
  scoreItem?: ScoreItem
  createdAt: string
  updatedAt: string
}

export interface ScoreVersion {
  id: string
  name: string
  slug: string
  description?: string
  siteIntro?: string
  order: number
  active: boolean
  context: 'public' | 'patient_prep' // public = Triagem/Light; patient_prep = forms de preparação
  items?: ScoreVersionItem[] // só vem no GET /:id (Preload ordenado por displayOrder)
  createdAt: string
  updatedAt: string
}

export interface CreateScoreVersionDTO {
  name: string
  slug: string
  description?: string
  siteIntro?: string
  order?: number
  active?: boolean
}

export interface UpdateScoreVersionDTO {
  name?: string
  slug?: string
  description?: string
  siteIntro?: string
  order?: number
  active?: boolean
}

// Item do payload de SetItems — replace total da lista de itens da versão.
export interface SetVersionItemDTO {
  scoreItemId: string
  displayOrder: number
}

// ============================================================
// Query keys
// ============================================================

export const scoreVersionKeys = {
  all: ['score-versions'] as const,
  lists: () => [...scoreVersionKeys.all, 'list'] as const,
  details: () => [...scoreVersionKeys.all, 'detail'] as const,
  detail: (id: string) => [...scoreVersionKeys.details(), id] as const,
}

// ============================================================
// Queries
// ============================================================

export function useScoreVersions() {
  return useQuery({
    queryKey: scoreVersionKeys.lists(),
    queryFn: () => apiClient.get<ScoreVersion[]>('/api/v1/score-versions'),
    staleTime: 60 * 1000,
  })
}

export function useScoreVersion(id: string) {
  return useQuery({
    queryKey: scoreVersionKeys.detail(id),
    queryFn: () => apiClient.get<ScoreVersion>(`/api/v1/score-versions/${id}`),
    enabled: !!id,
  })
}

// ============================================================
// Mutations (todas exigem RequireAdmin no backend)
// ============================================================

export function useCreateScoreVersion() {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: (data: CreateScoreVersionDTO) =>
      apiClient.post<ScoreVersion>('/api/v1/score-versions', data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: scoreVersionKeys.lists() })
    },
  })
}

export function useUpdateScoreVersion() {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: UpdateScoreVersionDTO }) =>
      apiClient.put<ScoreVersion>(`/api/v1/score-versions/${id}`, data),
    onSuccess: (_, { id }) => {
      queryClient.invalidateQueries({ queryKey: scoreVersionKeys.detail(id) })
      queryClient.invalidateQueries({ queryKey: scoreVersionKeys.lists() })
    },
  })
}

export function useDeleteScoreVersion() {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: (id: string) => apiClient.delete(`/api/v1/score-versions/${id}`),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: scoreVersionKeys.all })
    },
  })
}

// SetItems — substitui (replace total) os itens da versão. Backend responde 204.
export function useSetVersionItems() {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: ({ id, items }: { id: string; items: SetVersionItemDTO[] }) =>
      apiClient.put(`/api/v1/score-versions/${id}/items`, { items }),
    onSuccess: (_, { id }) => {
      queryClient.invalidateQueries({ queryKey: scoreVersionKeys.detail(id) })
      queryClient.invalidateQueries({ queryKey: scoreVersionKeys.lists() })
    },
  })
}
