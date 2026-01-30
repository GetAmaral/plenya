import { apiClient } from '../api-client'

export interface LabRequest {
  id: string
  patientId: string
  date: string
  exams: string
  notes?: string
  doctorId?: string
  createdAt: string
  updatedAt: string
  patient?: {
    id: string
    name: string
  }
  doctor?: {
    id: string
    email: string
  }
}

export interface CreateLabRequestInput {
  patientId: string
  date: string
  exams: string
  notes?: string
  doctorId?: string
}

export interface UpdateLabRequestInput {
  patientId?: string
  date?: string
  exams?: string
  notes?: string
  doctorId?: string
}

// Create lab request
export async function createLabRequest(data: CreateLabRequestInput): Promise<LabRequest> {
  return apiClient.post<LabRequest>('/api/v1/lab-requests', data)
}

// Get lab request by ID
export async function getLabRequestById(id: string): Promise<LabRequest> {
  return apiClient.get<LabRequest>(`/api/v1/lab-requests/${id}`)
}

// Get lab requests by patient ID
export async function getLabRequestsByPatientId(patientId: string): Promise<LabRequest[]> {
  return apiClient.get<LabRequest[]>(`/api/v1/patients/${patientId}/lab-requests`)
}

// Get all lab requests (paginated)
export async function getAllLabRequests(params?: {
  limit?: number
  offset?: number
}): Promise<{
  data: LabRequest[]
  total: number
  limit: number
  offset: number
}> {
  return apiClient.get<{
    data: LabRequest[]
    total: number
    limit: number
    offset: number
  }>('/api/v1/lab-requests', { params } as any)
}

// Get lab requests by date
export async function getLabRequestsByDate(date: string): Promise<LabRequest[]> {
  return apiClient.get<LabRequest[]>(`/api/v1/lab-requests/by-date?date=${encodeURIComponent(date)}`)
}

// Get lab requests by date range
export async function getLabRequestsByDateRange(startDate: string, endDate: string): Promise<LabRequest[]> {
  return apiClient.get<LabRequest[]>(`/api/v1/lab-requests/by-date-range?startDate=${encodeURIComponent(startDate)}&endDate=${encodeURIComponent(endDate)}`)
}

// Update lab request
export async function updateLabRequest(id: string, data: UpdateLabRequestInput): Promise<LabRequest> {
  return apiClient.put<LabRequest>(`/api/v1/lab-requests/${id}`, data)
}

// Delete lab request
export async function deleteLabRequest(id: string): Promise<void> {
  return apiClient.delete<void>(`/api/v1/lab-requests/${id}`)
}
