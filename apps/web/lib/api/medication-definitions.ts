import { apiClient } from './client'

export interface MedicationDefinition {
  id: string
  commonName: string
  activeIngredient: string
  category: 'simple' | 'c1' | 'c5' | 'antibiotic' | 'glp1'
  validityDays: number
  maxPerPrescription: number
  maxTreatmentDays: number
  requiresDigitalSignature: boolean
  requiresSNCR: boolean
  anvisaCode?: string
  createdAt: string
  updatedAt: string
}

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
  const response = await apiClient.get<MedicationDefinition[]>('/medication-definitions/search', {
    params: { q: query, limit },
  })
  return response.data
}

/**
 * List all medication definitions
 */
export async function listMedicationDefinitions(params?: {
  category?: string
  limit?: number
  offset?: number
}): Promise<MedicationDefinitionListResponse> {
  const response = await apiClient.get<MedicationDefinitionListResponse>(
    '/medication-definitions',
    { params }
  )
  return response.data
}

/**
 * Get medication definition by ID
 */
export async function getMedicationDefinition(id: string): Promise<MedicationDefinition> {
  const response = await apiClient.get<MedicationDefinition>(`/medication-definitions/${id}`)
  return response.data
}

/**
 * Create medication definition (admin only)
 */
export async function createMedicationDefinition(
  data: CreateMedicationDefinitionRequest
): Promise<MedicationDefinition> {
  const response = await apiClient.post<MedicationDefinition>('/medication-definitions', data)
  return response.data
}

/**
 * Update medication definition (admin only)
 */
export async function updateMedicationDefinition(
  id: string,
  data: Partial<CreateMedicationDefinitionRequest>
): Promise<MedicationDefinition> {
  const response = await apiClient.put<MedicationDefinition>(
    `/medication-definitions/${id}`,
    data
  )
  return response.data
}

/**
 * Delete medication definition (admin only)
 */
export async function deleteMedicationDefinition(id: string): Promise<void> {
  await apiClient.delete(`/medication-definitions/${id}`)
}
