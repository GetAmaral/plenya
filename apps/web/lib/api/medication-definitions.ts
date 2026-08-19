import { apiClient } from '../api-client'
// Tipo de model vem do GERADO (Go model → swag → openapi-typescript), via @plenya/types.
import type { MedicationDefinition } from '@plenya/types'

export type { MedicationDefinition }

export interface MedicationDefinitionListResponse {
  data: MedicationDefinition[]
  total: number
  page: number
  limit: number
}

export interface CreateMedicationDefinitionRequest {
  commonName: string
  activeIngredient: string
  category: 'simple' | 'c1' | 'c5' | 'antibiotic' | 'glp1'
  validityDays: number
  maxPerPrescription: number
  maxTreatmentDays: number
  requiresDigitalSignature: boolean
  requiresSNCR: boolean
  anvisaCode?: string
}

/**
 * Search medications (autocomplete)
 */
export async function searchMedications(query: string, limit = 10): Promise<MedicationDefinition[]> {
  const qs = new URLSearchParams({ q: query, limit: String(limit) })
  return apiClient.get<MedicationDefinition[]>(`/api/v1/medication-definitions/search?${qs}`)
}

/**
 * List all medication definitions. O endpoint devolve envelope { data, total, page, limit }.
 */
export async function listMedicationDefinitions(params?: {
  category?: string
  limit?: number
  offset?: number
}): Promise<MedicationDefinitionListResponse> {
  const qs = new URLSearchParams()
  if (params?.category) qs.set('category', params.category)
  if (params?.limit != null) qs.set('limit', String(params.limit))
  if (params?.offset != null) qs.set('offset', String(params.offset))
  const suffix = qs.toString() ? `?${qs}` : ''
  return apiClient.get<MedicationDefinitionListResponse>(`/api/v1/medication-definitions${suffix}`)
}

/**
 * Get medication definition by ID
 */
export async function getMedicationDefinition(id: string): Promise<MedicationDefinition> {
  return apiClient.get<MedicationDefinition>(`/api/v1/medication-definitions/${id}`)
}

/**
 * Create medication definition (admin only)
 */
export async function createMedicationDefinition(
  data: CreateMedicationDefinitionRequest
): Promise<MedicationDefinition> {
  return apiClient.post<MedicationDefinition>('/api/v1/medication-definitions', data)
}

/**
 * Update medication definition (admin only)
 */
export async function updateMedicationDefinition(
  id: string,
  data: Partial<CreateMedicationDefinitionRequest>
): Promise<MedicationDefinition> {
  return apiClient.put<MedicationDefinition>(`/api/v1/medication-definitions/${id}`, data)
}

/**
 * Delete medication definition (admin only)
 */
export async function deleteMedicationDefinition(id: string): Promise<void> {
  await apiClient.delete(`/api/v1/medication-definitions/${id}`)
}
