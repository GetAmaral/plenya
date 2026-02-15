import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { apiClient } from '../api-client'
import type { Method, MethodLetter, MethodPillar, ScoreItem } from '@plenya/types'

// Query keys
export const methodKeys = {
  all: ['methods'] as const,
  allTrees: () => [...methodKeys.all, 'trees'] as const,
  tree: (id: string) => [...methodKeys.all, id, 'tree'] as const,
  letters: (methodId: string) => [...methodKeys.all, methodId, 'letters'] as const,
  pillars: (letterId: string) => ['method-letters', letterId, 'pillars'] as const,
  pillar: (id: string) => ['method-pillars', id] as const,
  unassignedItems: () => ['score-items', 'unassigned'] as const,
  itemsWithPillars: () => ['score-items', 'with-pillars'] as const,
}

// ===== Method Queries =====

export function useAllMethods() {
  return useQuery({
    queryKey: methodKeys.all,
    queryFn: () => apiClient.get<Method[]>('/api/v1/methods'),
    staleTime: 10 * 60 * 1000, // 10 minutes
  })
}

export function useMethodTree(methodId: string | undefined) {
  return useQuery({
    queryKey: methodKeys.tree(methodId!),
    queryFn: () => apiClient.get<Method>(`/api/v1/methods/${methodId}/tree`),
    enabled: !!methodId,
    staleTime: 5 * 60 * 1000, // 5 minutes
  })
}

export function useAllMethodsWithTree() {
  return useQuery({
    queryKey: methodKeys.allTrees(),
    queryFn: () => apiClient.get<Method[]>('/api/v1/methods/tree'),
    staleTime: 5 * 60 * 1000,
  })
}

// ===== Method Letter Queries =====

export function useMethodLetters(methodId: string | undefined) {
  return useQuery({
    queryKey: methodKeys.letters(methodId!),
    queryFn: () => apiClient.get<MethodLetter[]>(`/api/v1/methods/${methodId}/letters`),
    enabled: !!methodId,
    staleTime: 10 * 60 * 1000,
  })
}

export function useMethodLetter(letterId: string | undefined) {
  return useQuery({
    queryKey: ['method-letters', letterId],
    queryFn: () => apiClient.get<MethodLetter>(`/api/v1/method-letters/${letterId}`),
    enabled: !!letterId,
    staleTime: 10 * 60 * 1000,
  })
}

// ===== Method Pillar Queries =====

export function useLetterPillars(letterId: string | undefined) {
  return useQuery({
    queryKey: methodKeys.pillars(letterId!),
    queryFn: () => apiClient.get<MethodPillar[]>(`/api/v1/method-letters/${letterId}/pillars`),
    enabled: !!letterId,
    staleTime: 10 * 60 * 1000,
  })
}

export function useMethodPillar(pillarId: string | undefined) {
  return useQuery({
    queryKey: methodKeys.pillar(pillarId!),
    queryFn: () => apiClient.get<MethodPillar>(`/api/v1/method-pillars/${pillarId}`),
    enabled: !!pillarId,
    staleTime: 5 * 60 * 1000,
  })
}

// ===== Score Item Queries =====

export function useUnassignedScoreItems() {
  return useQuery({
    queryKey: methodKeys.unassignedItems(),
    queryFn: () => apiClient.get<ScoreItem[]>('/api/v1/score-items/unassigned'),
    staleTime: 2 * 60 * 1000, // 2 minutes
  })
}

export function useAllScoreItemsWithPillars() {
  return useQuery({
    queryKey: methodKeys.itemsWithPillars(),
    queryFn: () => apiClient.get<ScoreItem[]>('/api/v1/score-items/with-pillars'),
    staleTime: 5 * 60 * 1000,
  })
}

// ===== Method Mutations =====

export interface CreateMethodDTO {
  name: string
  shortName: string
  description?: string
  version?: string
  color?: string
  order: number
  isDefault: boolean
}

export interface UpdateMethodDTO {
  name?: string
  shortName?: string
  description?: string
  version?: string
  color?: string
  order?: number
  isDefault?: boolean
}

export function useCreateMethod() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (data: CreateMethodDTO) =>
      apiClient.post<Method>('/api/v1/methods', data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: methodKeys.all })
    },
  })
}

export function useUpdateMethod() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: UpdateMethodDTO }) =>
      apiClient.put<Method>(`/api/v1/methods/${id}`, data),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: methodKeys.all })
      queryClient.invalidateQueries({ queryKey: methodKeys.tree(variables.id) })
    },
  })
}

export function useDeleteMethod() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (id: string) => apiClient.delete(`/api/v1/methods/${id}`),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: methodKeys.all })
    },
  })
}

// ===== Method Letter Mutations =====

export interface CreateMethodLetterDTO {
  code: string
  name: string
  description?: string
  clinicalRelevance?: string
  patientExplanation?: string
  conduct?: string
  color?: string
  icon?: string
  order: number
}

export interface UpdateMethodLetterDTO {
  code?: string
  name?: string
  description?: string
  clinicalRelevance?: string
  patientExplanation?: string
  conduct?: string
  color?: string
  icon?: string
  order?: number
}

export function useCreateMethodLetter() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({ methodId, data }: { methodId: string; data: CreateMethodLetterDTO }) =>
      apiClient.post<MethodLetter>(`/api/v1/methods/${methodId}/letters`, data),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: methodKeys.all })
      queryClient.invalidateQueries({ queryKey: methodKeys.tree(variables.methodId) })
      queryClient.invalidateQueries({ queryKey: methodKeys.letters(variables.methodId) })
    },
  })
}

export function useUpdateMethodLetter() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: UpdateMethodLetterDTO }) =>
      apiClient.put<MethodLetter>(`/api/v1/method-letters/${id}`, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: methodKeys.all })
    },
  })
}

export function useDeleteMethodLetter() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (id: string) => apiClient.delete(`/api/v1/method-letters/${id}`),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: methodKeys.all })
    },
  })
}

// ===== Method Pillar Mutations =====

export interface CreateMethodPillarDTO {
  name: string
  description?: string
  clinicalRelevance?: string
  patientExplanation?: string
  conduct?: string
  order: number
}

export interface UpdateMethodPillarDTO {
  name?: string
  description?: string
  clinicalRelevance?: string
  patientExplanation?: string
  conduct?: string
  order?: number
}

export function useCreateMethodPillar() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({ letterId, data }: { letterId: string; data: CreateMethodPillarDTO }) =>
      apiClient.post<MethodPillar>(`/api/v1/method-letters/${letterId}/pillars`, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: methodKeys.all })
    },
  })
}

export function useUpdateMethodPillar() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: UpdateMethodPillarDTO }) =>
      apiClient.put<MethodPillar>(`/api/v1/method-pillars/${id}`, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: methodKeys.all })
    },
  })
}

export function useDeleteMethodPillar() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (id: string) => apiClient.delete(`/api/v1/method-pillars/${id}`),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: methodKeys.all })
    },
  })
}

// ===== Score Item Assignment Mutations =====

export function useAssignItemToPillar() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({ itemId, pillarId }: { itemId: string; pillarId: string }) =>
      apiClient.post(`/api/v1/method-pillars/${pillarId}/assign-item`, { itemId }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: methodKeys.all })
      queryClient.invalidateQueries({ queryKey: methodKeys.unassignedItems() })
      queryClient.invalidateQueries({ queryKey: methodKeys.itemsWithPillars() })
    },
  })
}

export function useUnassignItemFromPillar() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({ itemId, pillarId }: { itemId: string; pillarId: string }) =>
      apiClient.delete(`/api/v1/method-pillars/${pillarId}/unassign-item`, {
        data: { itemId }
      }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: methodKeys.all })
      queryClient.invalidateQueries({ queryKey: methodKeys.unassignedItems() })
      queryClient.invalidateQueries({ queryKey: methodKeys.itemsWithPillars() })
    },
  })
}
