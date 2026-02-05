import { apiClient } from '../api-client'
import type { ScoreItem } from './score-api'

export interface AnamnesisItem {
  id: string
  scoreItemId: string
  scoreItem?: ScoreItem
  textValue?: string
  numericValue?: number
  order: number
  createdAt: string
  updatedAt: string
}

export interface AnamnesisItemRequest {
  scoreItemId: string
  textValue?: string
  numericValue?: number
  order: number
}

export interface Anamnesis {
  id: string
  patientId: string
  authorId: string
  anamnesisTemplateId?: string
  consultationDate: string
  content?: string // Plain text for search/indexing
  contentHtml?: string // HTML for display
  summary?: string // Plain text for search/indexing
  summaryHtml?: string // HTML for display
  visibility: 'all' | 'medicalOnly' | 'psychOnly' | 'authorOnly'
  notes?: string
  items?: AnamnesisItem[]
  createdAt: string
  updatedAt: string
  anamnesisTemplate?: {
    id: string
    name: string
    area: string
  }
  author?: {
    id: string
    name: string
    email: string
    role: string
  }
}

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
