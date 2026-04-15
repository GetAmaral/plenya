import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { apiClient } from '../api-client'

export interface PosturalAssessment {
  id: string
  patientId: string
  createdById: string
  assessmentDate: string
  physicalAssessmentId?: string
  viewType: string
  shoulderDeviation?: number
  hipDeviation?: number
  headLateralDeviation?: number
  fhp?: number
  thoracicKyphosis?: number
  lumbarLordosis?: number
  kneeAngle?: number
  photoUrl?: string
  posturalScore: number
  posturalClassification: string
  severeDeviations: number
  notes?: string
  createdAt: string
  updatedAt: string
}

export interface CreatePosturalAssessmentDTO {
  patientId?: string
  assessmentDate: string
  physicalAssessmentId?: string
  viewType: 'front' | 'side_left' | 'side_right' | 'back'
  shoulderDeviation?: number
  hipDeviation?: number
  headLateralDeviation?: number
  fhp?: number
  thoracicKyphosis?: number
  lumbarLordosis?: number
  kneeAngle?: number
  photoUrl?: string
  notes?: string
}

// Query key factory
export const posturalAssessmentKeys = {
  all: ['postural-assessments'] as const,
  lists: () => [...posturalAssessmentKeys.all, 'list'] as const,
  list: (params?: { limit?: number; offset?: number }) => [...posturalAssessmentKeys.lists(), params] as const,
  details: () => [...posturalAssessmentKeys.all, 'detail'] as const,
  detail: (id: string) => [...posturalAssessmentKeys.details(), id] as const,
}

export function usePosturalAssessments(limit = 20, offset = 0) {
  return useQuery({
    queryKey: posturalAssessmentKeys.list({ limit, offset }),
    queryFn: () => apiClient.get<PosturalAssessment[]>(`/api/v1/postural-assessments?limit=${limit}&offset=${offset}`),
  })
}

export function usePosturalAssessment(id: string) {
  return useQuery({
    queryKey: posturalAssessmentKeys.detail(id),
    queryFn: () => apiClient.get<PosturalAssessment>(`/api/v1/postural-assessments/${id}`),
    enabled: !!id,
  })
}

export function useCreatePosturalAssessment() {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: (data: CreatePosturalAssessmentDTO) =>
      apiClient.post<PosturalAssessment>('/api/v1/postural-assessments', data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: posturalAssessmentKeys.all })
    },
  })
}
