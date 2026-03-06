import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { apiClient } from '../api-client'

export type PeriodizationFramework = 'bompa' | 'linear' | 'undulating' | 'block'
export type MesocyclePhase = 'accumulation' | 'transformation' | 'realization' | 'hypertrophy' | 'strength' | 'endurance' | 'power'

export interface Mesocycle {
  id: string
  periodizationId: string
  order: number
  phase: MesocyclePhase
  durationWeeks: number
  volumePercent: number
  intensityPercent: number
  physiologicalFocus: string
  startDate: string
  endDate: string
}

export interface Periodization {
  id: string
  patientId: string
  createdById: string
  framework: PeriodizationFramework
  startDate: string
  totalWeeks: number
  objective: string
  scientificJustification?: string
  displayTitle: string
  mesocycles: Mesocycle[]
  createdAt: string
  updatedAt: string
}

export interface CreatePeriodizationDTO {
  patientId?: string
  framework: PeriodizationFramework
  startDate: string
  totalWeeks: number
  objective: string
}

// Query key factory
export const periodizationKeys = {
  all: ['periodizations'] as const,
  lists: () => [...periodizationKeys.all, 'list'] as const,
  list: (params?: { limit?: number; offset?: number }) => [...periodizationKeys.lists(), params] as const,
  details: () => [...periodizationKeys.all, 'detail'] as const,
  detail: (id: string) => [...periodizationKeys.details(), id] as const,
}

export function usePeriodizations(limit = 20, offset = 0) {
  return useQuery({
    queryKey: periodizationKeys.list({ limit, offset }),
    queryFn: () => apiClient.get<Periodization[]>(`/api/v1/periodizations?limit=${limit}&offset=${offset}`),
  })
}

export function usePeriodization(id: string) {
  return useQuery({
    queryKey: periodizationKeys.detail(id),
    queryFn: () => apiClient.get<Periodization>(`/api/v1/periodizations/${id}`),
    enabled: !!id,
  })
}

export function useCreatePeriodization() {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: (data: CreatePeriodizationDTO) =>
      apiClient.post<Periodization>('/api/v1/periodizations', data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: periodizationKeys.all })
    },
  })
}
