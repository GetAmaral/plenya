import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { apiClient } from '../api-client'

export interface FitnessTest {
  id: string
  patientId: string
  createdById: string
  assessmentDate: string
  abdominalReps?: number
  pushupReps?: number
  plankSeconds?: number
  burpeeCycles?: number
  frtReps?: number
  abdominalLevel?: string
  pushupLevel?: string
  plankLevel?: string
  burpeeLevel?: string
  frtLevel?: string
  overallScore: number
  overallClassification: string
  notes?: string
  createdAt: string
  updatedAt: string
}

export interface CreateFitnessTestDTO {
  patientId?: string
  assessmentDate: string
  abdominalReps?: number
  pushupReps?: number
  plankSeconds?: number
  burpeeCycles?: number
  frtReps?: number
  notes?: string
}

// Query key factory
export const fitnessTestKeys = {
  all: ['fitness-tests'] as const,
  lists: () => [...fitnessTestKeys.all, 'list'] as const,
  list: (params?: { limit?: number; offset?: number }) => [...fitnessTestKeys.lists(), params] as const,
  details: () => [...fitnessTestKeys.all, 'detail'] as const,
  detail: (id: string) => [...fitnessTestKeys.details(), id] as const,
}

export function useFitnessTests(limit = 20, offset = 0) {
  return useQuery({
    queryKey: fitnessTestKeys.list({ limit, offset }),
    queryFn: () => apiClient.get<FitnessTest[]>(`/api/v1/fitness-tests?limit=${limit}&offset=${offset}`),
  })
}

export function useFitnessTest(id: string) {
  return useQuery({
    queryKey: fitnessTestKeys.detail(id),
    queryFn: () => apiClient.get<FitnessTest>(`/api/v1/fitness-tests/${id}`),
    enabled: !!id,
  })
}

export function useCreateFitnessTest() {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: (data: CreateFitnessTestDTO) =>
      apiClient.post<FitnessTest>('/api/v1/fitness-tests', data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: fitnessTestKeys.all })
    },
  })
}
