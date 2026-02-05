import { apiClient } from './client'
import type { components } from '@plenya/types'

// Types from OpenAPI schema
export type Prescription = components['schemas']['models.Prescription']
export type PrescriptionStatus = components['schemas']['models.PrescriptionStatus']
export type MedicationCategory = components['schemas']['models.MedicationCategory']

export interface CreatePrescriptionRequest {
  patientId: string
  medicationName: string
  activeIngredient: string
  category: MedicationCategory
  concentration: string
  dosage: string
  frequency: string
  route: string
  duration: number
  quantity: number
  quantityInWords: string
  instructions?: string
  prescriptionDate: string // ISO date
}

export interface UpdatePrescriptionRequest {
  medicationName?: string
  activeIngredient?: string
  category?: MedicationCategory
  concentration?: string
  dosage?: string
  frequency?: string
  route?: string
  duration?: number
  quantity?: number
  quantityInWords?: string
  instructions?: string
  status?: PrescriptionStatus
}

export interface PrescriptionListResponse {
  data: Prescription[]
  total: number
  page: number
  limit: number
}

export interface SignPrescriptionResponse {
  signedPdfUrl: string
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
  const response = await apiClient.post<Prescription>('/prescriptions', data)
  return response.data
}

/**
 * Assinar prescrição e gerar PDF
 */
export async function signPrescription(
  id: string
): Promise<SignPrescriptionResponse> {
  const response = await apiClient.post<SignPrescriptionResponse>(
    `/prescriptions/${id}/sign`
  )
  return response.data
}

/**
 * Listar prescrições
 */
export async function listPrescriptions(params?: {
  patientId?: string
  status?: PrescriptionStatus
  limit?: number
  offset?: number
}): Promise<PrescriptionListResponse> {
  const response = await apiClient.get<PrescriptionListResponse>('/prescriptions', {
    params,
  })
  return response.data
}

/**
 * Obter prescrição por ID
 */
export async function getPrescription(id: string): Promise<Prescription> {
  const response = await apiClient.get<Prescription>(`/prescriptions/${id}`)
  return response.data
}

/**
 * Atualizar prescrição
 */
export async function updatePrescription(
  id: string,
  data: UpdatePrescriptionRequest
): Promise<Prescription> {
  const response = await apiClient.put<Prescription>(`/prescriptions/${id}`, data)
  return response.data
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
  const response = await apiClient.get<ValidationResult>(
    `/prescriptions/validate/${id}`
  )
  return response.data
}
