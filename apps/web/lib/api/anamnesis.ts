import { apiClient } from '../api-client'
// Anamnesis e AnamnesisItem vêm do GERADO (Go model → swag → openapi-typescript), via @plenya/types.
import type { Anamnesis, AnamnesisItem } from '@plenya/types'

export type { Anamnesis, AnamnesisItem }

export interface AnamnesisItemRequest {
  scoreItemId: string
  textValue?: string
  numericValue?: number
  selectedLevel?: number
  order: number
}

// Anamnesis: ver import de @plenya/types no topo (tipo gerado).

export interface CreateAnamnesisRequest {
  patientId?: string // Optional - backend uses selectedPatient from JWT
  anamnesisTemplateId?: string
  consultationDate: string // RFC3339 format
  content?: string // Plain text
  contentHtml?: string // HTML
  summary?: string // Plain text
  summaryHtml?: string // HTML
  visibility: 'all' | 'medicalOnly' | 'psychOnly' | 'authorOnly'
  notes?: string
  items?: AnamnesisItemRequest[]
}

export interface UpdateAnamnesisRequest {
  anamnesisTemplateId?: string
  consultationDate?: string
  content?: string // Plain text
  contentHtml?: string // HTML
  summary?: string // Plain text
  summaryHtml?: string // HTML
  visibility?: 'all' | 'medicalOnly' | 'psychOnly'
  notes?: string
  items?: AnamnesisItemRequest[]
}

export interface AnamnesisListResponse {
  data: Anamnesis[]
  total: number
  page: number
  limit: number
}

export async function getAllAnamnesis(params: {
  limit?: number
  offset?: number
  patientId?: string
}): Promise<AnamnesisListResponse> {
  const { limit = 20, offset = 0, patientId } = params
  const queryParams = new URLSearchParams({
    limit: limit.toString(),
    offset: offset.toString(),
  })

  if (patientId) {
    queryParams.append('patientId', patientId)
  }

  const result = await apiClient.get<Anamnesis[] | AnamnesisListResponse>(
    `/api/v1/anamnesis?${queryParams.toString()}`
  )

  // Handle both array and paginated response formats
  if (Array.isArray(result)) {
    return {
      data: result,
      total: result.length,
      page: 1,
      limit: result.length,
    }
  }

  return result
}

export async function getAnamnesisById(id: string): Promise<Anamnesis> {
  return apiClient.get<Anamnesis>(`/api/v1/anamnesis/${id}`)
}

export async function createAnamnesis(data: CreateAnamnesisRequest): Promise<Anamnesis> {
  return apiClient.post<Anamnesis>('/api/v1/anamnesis', data)
}

export async function updateAnamnesis(
  id: string,
  data: UpdateAnamnesisRequest
): Promise<Anamnesis> {
  return apiClient.put<Anamnesis>(`/api/v1/anamnesis/${id}`, data)
}

export async function deleteAnamnesis(id: string): Promise<void> {
  return apiClient.delete(`/api/v1/anamnesis/${id}`)
}
