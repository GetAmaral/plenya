import { apiClient } from '../api-client'

// Types for Lab Result Batch
export interface LabResultBatchResponse {
  id: string
  patientId: string
  labRequestId?: string
  requestingDoctorId?: string
  laboratoryName: string
  collectionDate: string
  resultDate?: string
  status: 'pending' | 'partial' | 'completed'
  observations?: string
  attachments?: string
  resultCount: number
  createdAt: string
  updatedAt: string
}

export interface LabResultInBatchResponse {
  id: string
  labResultBatchId: string
  labTestDefinitionId?: string
  labTestDefinition?: {
    id: string
    name: string
    code: string
    category: string
    unit?: string
  }
  testName: string
  testType: string
  resultText?: string
  resultNumeric?: number
  unit?: string
  interpretation?: string
  level?: number
  createdAt: string
  updatedAt: string
}

export interface LabResultBatchDetailResponse extends LabResultBatchResponse {
  labResults: LabResultInBatchResponse[]
}

export interface CreateLabResultInBatchRequest {
  labTestDefinitionId?: string
  testName: string
  testType: string
  resultText?: string
  resultNumeric?: number
  unit?: string
  interpretation?: string
  level?: number
  matched?: boolean
}

export interface CreateLabResultBatchRequest {
  labRequestId?: string
  requestingDoctorId?: string
  laboratoryName: string
  collectionDate: string
  resultDate?: string
  status: 'pending' | 'partial' | 'completed'
  observations?: string
  attachments?: string
  labResults: CreateLabResultInBatchRequest[]
}

export interface UpdateLabResultBatchRequest {
  labRequestId?: string
  requestingDoctorId?: string
  laboratoryName?: string
  collectionDate?: string
  resultDate?: string
  status?: 'pending' | 'partial' | 'completed'
  observations?: string
  attachments?: string
}

export interface UpdateLabResultInBatchRequest {
  labTestDefinitionId?: string
  testName?: string
  testType?: string
  resultText?: string
  resultNumeric?: number
  unit?: string
  interpretation?: string
  level?: number
}

export const labResultBatchApi = {
  list: async (params?: { status?: string; limit?: number; offset?: number }): Promise<LabResultBatchDetailResponse[]> => {
    const queryParams = new URLSearchParams()
    if (params?.status) queryParams.append('status', params.status)
    if (params?.limit) queryParams.append('limit', params.limit.toString())
    if (params?.offset) queryParams.append('offset', params.offset.toString())
    const query = queryParams.toString()
    return apiClient.get<LabResultBatchDetailResponse[]>(`/api/v1/lab-result-batches${query ? `?${query}` : ''}`)
  },

  getById: async (id: string): Promise<LabResultBatchDetailResponse> => {
    return apiClient.get<LabResultBatchDetailResponse>(`/api/v1/lab-result-batches/${id}`)
  },

  create: async (data: CreateLabResultBatchRequest): Promise<LabResultBatchDetailResponse> => {
    return apiClient.post<LabResultBatchDetailResponse>('/api/v1/lab-result-batches', data)
  },

  update: async (id: string, data: UpdateLabResultBatchRequest): Promise<LabResultBatchResponse> => {
    return apiClient.put<LabResultBatchResponse>(`/api/v1/lab-result-batches/${id}`, data)
  },

  delete: async (id: string): Promise<void> => {
    await apiClient.delete(`/api/v1/lab-result-batches/${id}`)
  },

  addResult: async (batchId: string, data: CreateLabResultInBatchRequest): Promise<LabResultInBatchResponse> => {
    return apiClient.post<LabResultInBatchResponse>(`/api/v1/lab-result-batches/${batchId}/results`, data)
  },

  updateResult: async (batchId: string, resultId: string, data: UpdateLabResultInBatchRequest): Promise<LabResultInBatchResponse> => {
    return apiClient.put<LabResultInBatchResponse>(`/api/v1/lab-result-batches/${batchId}/results/${resultId}`, data)
  },

  deleteResult: async (batchId: string, resultId: string): Promise<void> => {
    await apiClient.delete(`/api/v1/lab-result-batches/${batchId}/results/${resultId}`)
  },
}
