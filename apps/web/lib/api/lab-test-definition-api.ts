import { apiClient } from '../api-client'

export interface LabTestDefinition {
  id: string
  code: string
  name: string
  shortName?: string
  altNames?: string[]
  category: string
  isRequestable: boolean
  unit?: string
  resultType: string
  specimenType?: string
  fastingHours?: number
  description?: string
}

export const labTestDefinitionApi = {
  // Busca TODOS os exames (para edição de resultados)
  getAll: async () =>
    apiClient.get<LabTestDefinition[]>('/api/v1/lab-tests/definitions'),

  // Busca apenas exames que podem ser solicitados (para formulários)
  getRequestable: async () =>
    apiClient.get<LabTestDefinition[]>('/api/v1/lab-tests/requestable'),

  // Busca por termo (nome, código, categoria)
  search: async (query: string) =>
    apiClient.get<LabTestDefinition[]>('/api/v1/lab-tests/definitions/search', {
      params: { q: query },
    }),

  // Busca por ID
  getById: async (id: string) =>
    apiClient.get<LabTestDefinition>(`/api/v1/lab-tests/definitions/${id}`),
}
