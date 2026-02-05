import { apiClient } from '../api-client'

export interface LabTestDefinition {
  id: string
  name: string
  code: string
}

export interface LabResultViewItem {
  id: string
  labResultViewId: string
  labTestDefinitionId: string
  order: number
  labTestDefinition?: LabTestDefinition
  createdAt?: string
  updatedAt?: string
}

export interface LabResultView {
  id: string
  name: string
  description?: string
  isActive: boolean
  displayOrder: number
  items?: LabResultViewItem[]
  createdAt?: string
  updatedAt?: string
}

export interface CreateLabResultViewRequest {
  name: string
  description?: string
  displayOrder: number
}

export interface UpdateLabResultViewRequest {
  name?: string
  description?: string
  isActive?: boolean
  displayOrder?: number
}

export interface UpdateLabResultViewItemsRequest {
  items: {
    labTestDefinitionId: string
    order: number
  }[]
}

export const labResultViewApi = {
  /**
   * Lista todas as views
   */
  list: async (
    withItems = false,
    includeInactive = false
  ): Promise<LabResultView[]> => {
    const params = new URLSearchParams()
    if (withItems) params.append('withItems', 'true')
    if (includeInactive) params.append('includeInactive', 'true')

    return apiClient.get<LabResultView[]>(
      `/api/v1/lab-result-views?${params}`
    )
  },

  /**
   * Busca uma view por ID
   */
  getById: async (id: string): Promise<LabResultView> => {
    return apiClient.get<LabResultView>(
      `/api/v1/lab-result-views/${id}`
    )
  },

  /**
   * Cria uma nova view
   */
  create: async (data: CreateLabResultViewRequest): Promise<LabResultView> => {
    return apiClient.post<LabResultView>(
      '/api/v1/lab-result-views',
      data
    )
  },

  /**
   * Atualiza uma view
   */
  update: async (
    id: string,
    data: UpdateLabResultViewRequest
  ): Promise<LabResultView> => {
    return apiClient.put<LabResultView>(
      `/api/v1/lab-result-views/${id}`,
      data
    )
  },

  /**
   * Deleta uma view
   */
  delete: async (id: string): Promise<void> => {
    return apiClient.delete(`/api/v1/lab-result-views/${id}`)
  },

  /**
   * Atualiza os items de uma view
   */
  updateItems: async (
    id: string,
    data: UpdateLabResultViewItemsRequest
  ): Promise<void> => {
    return apiClient.put(`/api/v1/lab-result-views/${id}/items`, data)
  },

  /**
   * Busca views por nome ou descrição
   */
  search: async (
    query: string,
    withItems = false
  ): Promise<LabResultView[]> => {
    const params = new URLSearchParams({ q: query })
    if (withItems) params.append('withItems', 'true')

    return apiClient.get<LabResultView[]>(
      `/api/v1/lab-result-views/search?${params}`
    )
  },
}
