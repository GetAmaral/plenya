import { useQuery } from '@tanstack/react-query'
import { apiClient } from '../api-client'

export interface Exercise {
  id: string
  externalId: string
  name: string
  namePt: string
  bodyParts: string[]
  bodyPartsPt: string[]
  targetMuscles: string[]
  targetMusclesPt: string[]
  equipments: string[]
  equipmentsPt: string[]
  secondaryMuscles: string[]
  instructions: string[]
  instructionsPt: string[]
  gifUrl: string
  gifUrlFallback: string
  isActive: boolean
}

export interface ExerciseListResponse {
  exercises: Exercise[]
  total: number
}

export interface ExerciseFilters {
  search?: string
  bodyParts?: string[]
  targetMuscles?: string[]
  equipments?: string[]
  limit?: number
  offset?: number
}

// Query key factory
export const exerciseKeys = {
  all: ['exercises'] as const,
  lists: () => [...exerciseKeys.all, 'list'] as const,
  list: (filters: ExerciseFilters) => [...exerciseKeys.lists(), filters] as const,
  details: () => [...exerciseKeys.all, 'detail'] as const,
  detail: (id: string) => [...exerciseKeys.details(), id] as const,
}

export function useExercises(filters: ExerciseFilters = {}) {
  const params = new URLSearchParams()
  if (filters.search) params.set('search', filters.search)
  if (filters.bodyParts?.length) params.set('bodyParts', filters.bodyParts.join(','))
  if (filters.targetMuscles?.length) params.set('targetMuscles', filters.targetMuscles.join(','))
  if (filters.equipments?.length) params.set('equipments', filters.equipments.join(','))
  params.set('limit', String(filters.limit || 20))
  params.set('offset', String(filters.offset || 0))

  return useQuery({
    queryKey: exerciseKeys.list(filters),
    queryFn: () => apiClient.get<ExerciseListResponse>(`/api/v1/exercises?${params.toString()}`),
  })
}

export function useExercise(id: string) {
  return useQuery({
    queryKey: exerciseKeys.detail(id),
    queryFn: () => apiClient.get<Exercise>(`/api/v1/exercises/${id}`),
    enabled: !!id,
  })
}
