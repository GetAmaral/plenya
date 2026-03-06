import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { apiClient } from '../api-client'

export interface HRZone {
  zone: number
  name: string
  minBpm: number
  maxBpm: number
  percent: string
}

export interface ResolvedData {
  weight?: number
  height?: number
  bmi?: number
  bri?: number
  waistCm?: number
  bodyFatPercent?: number
  ldl?: number
  hdl?: number
  totalChol?: number
  triglycerides?: number
  fastingGlucose?: number
  hbA1c?: number
  maxHr?: number
  hrZones?: HRZone[]
}

export interface PhysicalAssessment {
  id: string
  patientId: string
  createdById: string
  assessmentDate: string
  anamnesisId: string
  acsmRiskLevel?: string
  acsmRiskFactorsCount: number
  acsmRiskFactors: string[]
  acsmProtectiveFactors: string[]
  acsmRecommendation?: string
  acsmTags: string[]
  frontPhotoUrl?: string
  sidePhotoUrl?: string
  aiRecommendation?: string
  displayTitle: string
  resolvedData?: ResolvedData
  createdAt: string
  updatedAt: string
}

export interface CreatePhysicalAssessmentDTO {
  patientId?: string
  anamnesisId: string
  assessmentDate: string
}

// Query key factory
export const physicalAssessmentKeys = {
  all: ['physical-assessments'] as const,
  lists: () => [...physicalAssessmentKeys.all, 'list'] as const,
  list: (params?: { limit?: number; offset?: number }) => [...physicalAssessmentKeys.lists(), params] as const,
  details: () => [...physicalAssessmentKeys.all, 'detail'] as const,
  detail: (id: string) => [...physicalAssessmentKeys.details(), id] as const,
}

export function usePhysicalAssessments(limit = 20, offset = 0) {
  return useQuery({
    queryKey: physicalAssessmentKeys.list({ limit, offset }),
    queryFn: () => apiClient.get<PhysicalAssessment[]>(`/api/v1/physical-assessments?limit=${limit}&offset=${offset}`),
  })
}

export function usePhysicalAssessment(id: string) {
  return useQuery({
    queryKey: physicalAssessmentKeys.detail(id),
    queryFn: () => apiClient.get<PhysicalAssessment>(`/api/v1/physical-assessments/${id}`),
    enabled: !!id,
  })
}

export function useCreatePhysicalAssessment() {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: (data: CreatePhysicalAssessmentDTO) =>
      apiClient.post<PhysicalAssessment>('/api/v1/physical-assessments', data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: physicalAssessmentKeys.all })
    },
  })
}
