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
  async list(
    withItems = false,
    includeInactive = false
  ): Promise<LabResultView[]> {
    const params = new URLSearchParams()
    if (withItems) params.append('withItems', 'true')
    if (includeInactive) params.append('includeInactive', 'true')

    const response = await apiClient.get<LabResultView[]>(
      `/lab-result-views?${params}`
    )
    return response.data
  },

  /**
   * Busca uma view por ID
   */
  async getById(id: string): Promise<LabResultView> {
    const response = await apiClient.get<LabResultView>(
      `/lab-result-views/${id}`
    )
    return response.data
  },

  /**
   * Cria uma nova view
   */
  async create(data: CreateLabResultViewRequest): Promise<LabResultView> {
    const response = await apiClient.post<LabResultView>(
      '/lab-result-views',
      data
    )
    return response.data
  },

  /**
   * Atualiza uma view
   */
  async update(
    id: string,
    data: UpdateLabResultViewRequest
  ): Promise<LabResultView> {
    const response = await apiClient.put<LabResultView>(
      `/lab-result-views/${id}`,
      data
    )
    return response.data
  },

  /**
   * Deleta uma view
   */
  async delete(id: string): Promise<void> {
    await apiClient.delete(`/lab-result-views/${id}`)
  },

  /**
   * Atualiza os items de uma view
   */
  async updateItems(
    id: string,
    data: UpdateLabResultViewItemsRequest
  ): Promise<void> {
    await apiClient.put(`/lab-result-views/${id}/items`, data)
  },

  /**
   * Busca views por nome ou descrição
   */
  async search(
    query: string,
    withItems = false
  ): Promise<LabResultView[]> {
    const params = new URLSearchParams({ q: query })
    if (withItems) params.append('withItems', 'true')

    const response = await apiClient.get<LabResultView[]>(
      `/lab-result-views/search?${params}`
    )
    return response.data
  },
}
