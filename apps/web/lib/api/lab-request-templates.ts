import { apiClient } from '../api-client'

export interface LabTestDefinition {
  id: string
  code: string
  name: string
  shortName?: string
  altNames?: string[]
  tussCode?: string
  category: string
  isRequestable: boolean
  unit?: string
  resultType: string
  isActive: boolean
}

export interface LabRequestTemplate {
  id: string
  name: string
  description?: string
  displayOrder: number
  isActive: boolean
  createdAt: string
  updatedAt: string
  labTests?: LabTestDefinition[]
}

export interface CreateLabRequestTemplateInput {
  name: string
  description?: string
  displayOrder?: number
}

export interface UpdateLabRequestTemplateInput {
  name?: string
  description?: string
  displayOrder?: number
  isActive?: boolean
}

// Create lab request template
export async function createLabRequestTemplate(data: CreateLabRequestTemplateInput): Promise<LabRequestTemplate> {
  return apiClient.post<LabRequestTemplate>('/api/v1/lab-request-templates', data)
}

// Get lab request template by ID
export async function getLabRequestTemplateById(id: string): Promise<LabRequestTemplate> {
  return apiClient.get<LabRequestTemplate>(`/api/v1/lab-request-templates/${id}`)
}

// Get all lab request templates
export async function getAllLabRequestTemplates(): Promise<LabRequestTemplate[]> {
  return apiClient.get<LabRequestTemplate[]>('/api/v1/lab-request-templates')
}

// Search lab request templates
export async function searchLabRequestTemplates(searchTerm: string): Promise<LabRequestTemplate[]> {
  return apiClient.get<LabRequestTemplate[]>(`/api/v1/lab-request-templates/search?q=${encodeURIComponent(searchTerm)}`)
}

// Update lab request template
export async function updateLabRequestTemplate(
  id: string,
  data: UpdateLabRequestTemplateInput
): Promise<LabRequestTemplate> {
  return apiClient.put<LabRequestTemplate>(`/api/v1/lab-request-templates/${id}`, data)
}

// Update lab request template tests (replace all)
export async function updateLabRequestTemplateTests(
  id: string,
  testIds: string[]
): Promise<LabRequestTemplate> {
  return apiClient.put<LabRequestTemplate>(`/api/v1/lab-request-templates/${id}/tests`, {
    testIds
  })
}

// Add lab test to template
export async function addLabTestToTemplate(id: string, testId: string): Promise<LabRequestTemplate> {
  return apiClient.post<LabRequestTemplate>(`/api/v1/lab-request-templates/${id}/tests`, {
    testId
  })
}

// Remove lab test from template
export async function removeLabTestFromTemplate(id: string, testId: string): Promise<void> {
  return apiClient.delete<void>(`/api/v1/lab-request-templates/${id}/tests/${testId}`)
}

// Delete lab request template
export async function deleteLabRequestTemplate(id: string): Promise<void> {
  return apiClient.delete<void>(`/api/v1/lab-request-templates/${id}`)
}

// Get requestable lab tests (for template editor)
export async function getRequestableLabTests(): Promise<LabTestDefinition[]> {
  return apiClient.get<LabTestDefinition[]>('/api/v1/lab-tests/requestable')
}
