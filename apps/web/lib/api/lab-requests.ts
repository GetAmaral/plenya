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
  const response = await apiClient.post('/lab-requests', data)
  return response.data
}

// Get lab request by ID
export async function getLabRequestById(id: string): Promise<LabRequest> {
  const response = await apiClient.get(`/lab-requests/${id}`)
  return response.data
}

// Get lab requests by patient ID
export async function getLabRequestsByPatientId(patientId: string): Promise<LabRequest[]> {
  const response = await apiClient.get(`/patients/${patientId}/lab-requests`)
  return response.data
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
  const response = await apiClient.get('/lab-requests', { params })
  return response.data
}

// Get lab requests by date
export async function getLabRequestsByDate(date: string): Promise<LabRequest[]> {
  const response = await apiClient.get('/lab-requests/by-date', {
    params: { date }
  })
  return response.data
}

// Get lab requests by date range
export async function getLabRequestsByDateRange(startDate: string, endDate: string): Promise<LabRequest[]> {
  const response = await apiClient.get('/lab-requests/by-date-range', {
    params: { startDate, endDate }
  })
  return response.data
}

// Update lab request
export async function updateLabRequest(id: string, data: UpdateLabRequestInput): Promise<LabRequest> {
  const response = await apiClient.put(`/lab-requests/${id}`, data)
  return response.data
}

// Delete lab request
export async function deleteLabRequest(id: string): Promise<void> {
  await apiClient.delete(`/lab-requests/${id}`)
}
