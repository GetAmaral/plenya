import { apiClient } from '../api-client'
import { useAuthStore } from '../auth-store'
import type { components } from '@plenya/types'
// Prescription (= dto.PrescriptionResponse refinado: medications sempre presente) vem do GERADO.
import type { Prescription } from '@plenya/types'

const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:3001'

// Tipos do schema gerado (Go → swag → openapi-typescript).
export type { Prescription }
export type PrescriptionStatus = components['schemas']['models.PrescriptionStatus']
export type MedicationCategory = components['schemas']['models.MedicationCategory']
// Requests = DTOs gerados (medications[] é a forma real; o flat antigo estava obsoleto).
export type CreatePrescriptionRequest = components['schemas']['dto.CreatePrescriptionRequest']
export type UpdatePrescriptionRequest = components['schemas']['dto.UpdatePrescriptionRequest']

export interface SignPrescriptionResponse {
  signedPdfUrl: string
  signatureMode?: 'digital' | 'manual'
  message?: string
  sncrNumber?: string
}

export interface ValidationResult {
  valid: boolean
  pdfIntact: boolean
  prescription: {
    id: string
    prescriptionDate: string
    validUntil: string
    isExpired: boolean
    isUsed: boolean
    sncrNumber?: string
    category: MedicationCategory
  }
  patient: {
    name: string
    cpf: string // Masked: ***.***. 789-00
  }
  doctor: {
    name: string
    crm: string
  }
  medication: {
    name: string
    activeIngredient: string
    concentration: string
    quantity: number
    quantityInWords: string
  }
  signature: {
    signedAt: string
    certificateSerial?: string
    signedPdfUrl?: string
  }
}

/**
 * Criar nova prescrição
 */
export async function createPrescription(
  data: CreatePrescriptionRequest
): Promise<Prescription> {
  return apiClient.post<Prescription>('/prescriptions', data)
}

/**
 * Assinar prescrição e gerar PDF
 */
export async function signPrescription(
  id: string
): Promise<SignPrescriptionResponse> {
  return apiClient.post<SignPrescriptionResponse>(`/prescriptions/${id}/sign`)
}

/**
 * Listar prescrições. O endpoint devolve um array puro (`[]dto.PrescriptionResponse`).
 */
export async function listPrescriptions(params?: {
  patientId?: string
  status?: PrescriptionStatus
  limit?: number
  offset?: number
}): Promise<Prescription[]> {
  const qs = new URLSearchParams()
  if (params?.patientId) qs.set('patientId', params.patientId)
  if (params?.status) qs.set('status', params.status)
  if (params?.limit != null) qs.set('limit', String(params.limit))
  if (params?.offset != null) qs.set('offset', String(params.offset))
  const suffix = qs.toString() ? `?${qs}` : ''
  return apiClient.get<Prescription[]>(`/prescriptions${suffix}`)
}

/**
 * Obter prescrição por ID
 */
export async function getPrescription(id: string): Promise<Prescription> {
  return apiClient.get<Prescription>(`/prescriptions/${id}`)
}

/**
 * Atualizar prescrição
 */
export async function updatePrescription(
  id: string,
  data: UpdatePrescriptionRequest
): Promise<Prescription> {
  return apiClient.put<Prescription>(`/prescriptions/${id}`, data)
}

/**
 * Deletar prescrição (soft delete)
 */
export async function deletePrescription(id: string): Promise<void> {
  await apiClient.delete(`/prescriptions/${id}`)
}

/**
 * Validar prescrição publicamente (sem autenticação)
 * Usado por farmácias via QR Code
 */
export async function validatePublic(id: string): Promise<ValidationResult> {
  return apiClient.get<ValidationResult>(`/prescriptions/validate/${id}`)
}

/**
 * URL do download autenticado do PDF da prescrição.
 * O servir estático de /uploads foi removido (vazava PDFs sem auth); o PDF agora é
 * entregue como PatientDocument por endpoint autenticado.
 */
export function prescriptionDownloadURL(id: string): string {
  return `${API_URL}/api/v1/prescriptions/${id}/download`
}

/**
 * Baixa o PDF da prescrição com o token (Bearer) e abre numa nova aba.
 * window.open direto não funciona porque o endpoint exige autenticação.
 */
export async function openPrescriptionPdf(id: string): Promise<void> {
  const token = useAuthStore.getState().accessToken
  const res = await fetch(prescriptionDownloadURL(id), {
    headers: token ? { Authorization: `Bearer ${token}` } : {},
  })
  if (!res.ok) throw new Error('Não foi possível baixar o PDF da prescrição')
  const blob = await res.blob()
  const url = URL.createObjectURL(blob)
  window.open(url, '_blank')
  setTimeout(() => URL.revokeObjectURL(url), 60_000)
}
