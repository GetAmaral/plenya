import { apiClient } from '../api-client'
// LabRequest vem do GERADO (Go model → swag → openapi-typescript), via @plenya/types.
import type { LabRequest } from '@plenya/types'

export type { LabRequest }

export interface CreateLabRequestInput {
  patientId?: string // Optional - backend preenche automaticamente do selectedPatient
  date: string
  exams: string
  notes?: string
  doctorId?: string // Optional - backend preenche automaticamente do currentUser
  labRequestTemplateId?: string
}

export interface UpdateLabRequestInput {
  patientId?: string
  date?: string
  exams?: string
  notes?: string
  doctorId?: string
  labRequestTemplateId?: string
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

// Generate PDF for lab request
export async function generateLabRequestPdf(id: string): Promise<LabRequest> {
  return apiClient.post<LabRequest>(`/api/v1/lab-requests/${id}/generate-pdf`, {})
}

/** Baixa o PDF (autenticado) e abre numa nova aba. A rota estática /uploads foi removida (H1). */
export async function openLabRequestPdf(id: string) {
  const blob = await apiClient.getBlob(`/api/v1/lab-requests/${id}/pdf`)
  const url = URL.createObjectURL(blob)
  window.open(url, '_blank')
  setTimeout(() => URL.revokeObjectURL(url), 60_000)
}

/** Exame extraído de um pedido externo (foto/PDF), com o match no nosso catálogo. */
export interface ImportedExam {
  raw: string
  matched: boolean
  code?: string
  name?: string
  tussCode?: string
}

/** Importa um pedido de exames externo (foto/PDF), extrai e casa os exames no catálogo (dedup). */
export async function importLabRequestExams(file: File): Promise<{ items: ImportedExam[] }> {
  const form = new FormData()
  form.append('file', file)
  return apiClient.post<{ items: ImportedExam[] }>('/api/v1/lab-requests/extract-exams', form)
}
