import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { apiClient } from '../api-client'

export interface HRZone {
  zone: number
  name: string
  minBpm: number
  maxBpm: number
  percent: string
}

export interface PhysicalAssessment {
  id: string
  patientId: string
  createdById: string
  assessmentDate: string

  // Antropometria
  weight?: number
  height?: number
  waistCircumference?: number

  // Composição corporal
  bmi?: number
  bri?: number
  bodyFatPercent?: number
  leanMass?: number

  // Cardiovascular
  systolicBp?: number
  diastolicBp?: number
  restingHeartRate?: number

  // Laboratorial
  ldl?: number
  hdl?: number
  totalCholesterol?: number
  triglycerides?: number
  fastingGlucose?: number
  hbA1c?: number

  // Histórico
  familyHistory?: boolean
  smokingStatus?: 'never' | 'former' | 'current'
  physicalActivityLevel?: 'sedentary' | 'insufficient' | 'moderate' | 'active'

  // Condições
  cardiovascularDisease?: boolean
  diabetesType?: string
  symptoms?: string
  clinicalAlert?: boolean

  // ACSM
  acsmRiskLevel?: string
  acsmRiskFactorsCount: number
  acsmRiskFactors: string[]
  acsmProtectiveFactors: string[]
  acsmRecommendation?: string
  acsmTags: string[]

  // Fotos
  frontPhotoUrl?: string
  sidePhotoUrl?: string

  // IA
  aiRecommendation?: string

  // FC (calculado)
  maxHr?: number
  hrZones?: HRZone[]

  displayTitle: string
  createdAt: string
  updatedAt: string
}

export interface CreatePhysicalAssessmentDTO {
  patientId?: string
  assessmentDate: string

  // Antropometria
  weight?: number
  height?: number
  waistCircumference?: number

  // Composição corporal
  bodyFatPercent?: number
  leanMass?: number

  // Cardiovascular
  systolicBp?: number
  diastolicBp?: number
  restingHeartRate?: number

  // Laboratorial
  ldl?: number
  hdl?: number
  totalCholesterol?: number
  triglycerides?: number
  fastingGlucose?: number
  hbA1c?: number

  // Histórico
  familyHistory?: boolean
  smokingStatus?: 'never' | 'former' | 'current'
  physicalActivityLevel?: 'sedentary' | 'insufficient' | 'moderate' | 'active'

  // Condições
  cardiovascularDisease?: boolean
  diabetesType?: string
  symptoms?: string
  clinicalAlert?: boolean

  // Fotos
  frontPhotoUrl?: string
  sidePhotoUrl?: string
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
