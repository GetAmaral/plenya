import { openServerPdf } from '@/lib/open-pdf'
import { apiClient } from '../api-client'
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

// Espelha o fiber.Map montado em PrescriptionHandler.ValidatePublic (resposta pública do QR Code).
export interface ValidationResult {
  valid: boolean
  pdfIntact: boolean
  signatureMode?: string
  prescription: {
    id: string
    prescriptionDate: string
    validUntil: string
    isExpired: boolean
    isUsed: boolean
    sncrNumber?: string
    medicationCount: number
  }
  patient: {
    name: string
    cpf: string // Masked: ***.***. 789-00
  }
  doctor: {
    name: string
    crm: string
  }
  medications: Array<{
    id: string
    name: string
    activeIngredient: string
    category: MedicationCategory
    concentration: string
    quantity: number
    quantityInWords: string
  }>
  signature: {
    mode?: string
    signedAt: string
    certificateSerial?: string
  }
}

/**
 * Criar nova prescrição
 */
export async function createPrescription(
  data: CreatePrescriptionRequest
): Promise<Prescription> {
  return apiClient.post<Prescription>('/api/v1/prescriptions', data)
}

/**
 * Assinar prescrição e gerar PDF
 */
export async function signPrescription(
  id: string
): Promise<SignPrescriptionResponse> {
  return apiClient.post<SignPrescriptionResponse>(`/api/v1/prescriptions/${id}/sign`)
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
  return apiClient.get<Prescription[]>(`/api/v1/prescriptions${suffix}`)
}

/**
 * Obter prescrição por ID
 */
export async function getPrescription(id: string): Promise<Prescription> {
  return apiClient.get<Prescription>(`/api/v1/prescriptions/${id}`)
}

/**
 * Atualizar prescrição
 */
export async function updatePrescription(
  id: string,
  data: UpdatePrescriptionRequest
): Promise<Prescription> {
  return apiClient.put<Prescription>(`/api/v1/prescriptions/${id}`, data)
}

/**
 * Deletar prescrição (soft delete)
 */
export async function deletePrescription(id: string): Promise<void> {
  await apiClient.delete(`/api/v1/prescriptions/${id}`)
}

/**
 * Validar prescrição publicamente (sem autenticação)
 * Usado por farmácias via QR Code
 */
export async function validatePublic(id: string): Promise<ValidationResult> {
  return apiClient.get<ValidationResult>(`/api/v1/prescriptions/validate/${id}`)
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
 * Baixa o PDF da prescrição com o token (Bearer) e entrega com o nome do servidor
 * ("Ana-Cláudia_Receita_2026-08-31_01a0592b.pdf"). Ver lib/open-pdf.ts.
 */
export async function openPrescriptionPdf(id: string): Promise<void> {
  await openServerPdf(`/api/v1/prescriptions/${id}/download`, `Receita_${id.slice(0, 8)}.pdf`)
}

/**
 * Abre o RASCUNHO da receita, antes de assinar.
 *
 * Sai sempre no layout de assinatura manual, com o campo de assinatura em branco, mesmo quando o
 * médico tem certificado ativo: rascunho com selo ICP-Brasil é indistinguível de receita válida
 * para quem receber o arquivo solto. O servidor não grava nada e não marca a receita como
 * assinada.
 */
export async function openPrescriptionDraftPdf(id: string): Promise<void> {
  await openServerPdf(
    `/api/v1/prescriptions/${id}/preview`,
    `RascunhoReceita_${id.slice(0, 8)}.pdf`,
  )
}
