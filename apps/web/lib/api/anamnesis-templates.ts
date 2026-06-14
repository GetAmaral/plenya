import { apiClient } from '../api-client'
import { ScoreItem as ScoreItemType } from './score-api'
// AnamnesisTemplate(+Item) vêm do GERADO (Go model → swag → openapi-typescript), via @plenya/types.
import type { AnamnesisTemplate, AnamnesisTemplateItem } from '@plenya/types'

export type { AnamnesisTemplate, AnamnesisTemplateItem }

// Re-export ScoreItem from score-api
export type ScoreItem = ScoreItemType

// Types
export type AnamnesisTemplateArea = 'Medicina' | 'Nutricao' | 'Psicologia' | 'Educacao Fisica'

// AnamnesisTemplate e AnamnesisTemplateItem: ver import de @plenya/types no topo.

export interface CreateAnamnesisTemplateRequest {
  name: string
  area: AnamnesisTemplateArea
}

export interface UpdateAnamnesisTemplateRequest {
  name?: string
  area?: AnamnesisTemplateArea
}

export interface UpdateAnamnesisTemplateItemsRequest {
  items: {
    scoreItemId: string
    order: number
  }[]
}

// API functions

/**
 * Get all anamnesis templates
 */
export async function getAllAnamnesisTemplates(withItems = false): Promise<AnamnesisTemplate[]> {
  return apiClient.get<AnamnesisTemplate[]>(
    `/api/v1/anamnesis-templates?withItems=${withItems}`
  )
}

/**
 * Get anamnesis template by ID
 */
export async function getAnamnesisTemplateById(id: string, withItems = true): Promise<AnamnesisTemplate> {
  return apiClient.get<AnamnesisTemplate>(
    `/api/v1/anamnesis-templates/${id}?withItems=${withItems}`
  )
}

/**
 * Create anamnesis template
 */
export async function createAnamnesisTemplate(
  data: CreateAnamnesisTemplateRequest
): Promise<AnamnesisTemplate> {
  return apiClient.post<AnamnesisTemplate>('/api/v1/anamnesis-templates', data)
}

/**
 * Update anamnesis template info (name, area)
 */
export async function updateAnamnesisTemplate(
  id: string,
  data: UpdateAnamnesisTemplateRequest
): Promise<AnamnesisTemplate> {
  return apiClient.put<AnamnesisTemplate>(`/api/v1/anamnesis-templates/${id}`, data)
}

/**
 * Update anamnesis template items (replace all)
 */
export async function updateAnamnesisTemplateItems(
  id: string,
  data: UpdateAnamnesisTemplateItemsRequest
): Promise<void> {
  await apiClient.put(`/api/v1/anamnesis-templates/${id}/items`, data)
}

/**
 * Delete anamnesis template
 */
export async function deleteAnamnesisTemplate(id: string): Promise<void> {
  await apiClient.delete(`/api/v1/anamnesis-templates/${id}`)
}

/**
 * Search anamnesis templates
 */
export async function searchAnamnesisTemplates(query: string): Promise<AnamnesisTemplate[]> {
  return apiClient.get<AnamnesisTemplate[]>(
    `/api/v1/anamnesis-templates/search?q=${encodeURIComponent(query)}`
  )
}

/**
 * Get all score items (for template creation)
 */
export async function getAllScoreItems(): Promise<ScoreItem[]> {
  return apiClient.get<ScoreItem[]>('/api/v1/score-items')
}
