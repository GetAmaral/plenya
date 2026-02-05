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
  testName: string
  testType: string
  status: 'pending' | 'completed' | 'cancelled'
  resultText?: string
  resultNumeric?: number
  unit?: string
  referenceRange?: string
  interpretation?: string
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
  status: 'pending' | 'completed' | 'cancelled'
  resultText?: string
  resultNumeric?: number
  unit?: string
  referenceRange?: string
  interpretation?: string
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
  status?: 'pending' | 'completed' | 'cancelled'
  resultText?: string
  resultNumeric?: number
  unit?: string
  referenceRange?: string
  interpretation?: string
}

export const labResultBatchApi = {
  list: async (params?: { status?: string; limit?: number; offset?: number }) =>
    apiClient.get<LabResultBatchDetailResponse[]>('/api/v1/lab-result-batches', { params }),

  getById: async (id: string) =>
    apiClient.get<LabResultBatchDetailResponse>(`/api/v1/lab-result-batches/${id}`),

  create: async (data: CreateLabResultBatchRequest) =>
    apiClient.post<LabResultBatchDetailResponse>('/api/v1/lab-result-batches', data),

  update: async (id: string, data: UpdateLabResultBatchRequest) =>
    apiClient.put<LabResultBatchResponse>(`/api/v1/lab-result-batches/${id}`, data),

  delete: async (id: string) =>
    apiClient.delete(`/api/v1/lab-result-batches/${id}`),

  addResult: async (batchId: string, data: CreateLabResultInBatchRequest) =>
    apiClient.post<LabResultInBatchResponse>(`/api/v1/lab-result-batches/${batchId}/results`, data),

  updateResult: async (batchId: string, resultId: string, data: UpdateLabResultInBatchRequest) =>
    apiClient.put<LabResultInBatchResponse>(`/api/v1/lab-result-batches/${batchId}/results/${resultId}`, data),

  deleteResult: async (batchId: string, resultId: string) =>
    apiClient.delete(`/api/v1/lab-result-batches/${batchId}/results/${resultId}`),
}
