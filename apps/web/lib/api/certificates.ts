import { apiClient } from './client'

export interface Certificate {
  userId: string
  userName: string
  userEmail: string
  userCPF?: string
  validUntil: string
  certificateSerial: string
  certificateCPF?: string
  certificateName?: string
  certificateActive: boolean
  daysUntilExpiry: number
  needsRenewal: boolean
  isExpired: boolean
}

export interface CertificateListResponse {
  data: Certificate[]
}

export interface CertificateStatus {
  hasCertificate: boolean
  validUntil?: string
  certificateSerial?: string
  daysUntilExpiry?: number
  needsRenewal?: boolean
}

export interface UploadCertificateResponse {
  success: boolean
  message: string
}

/**
 * Listar certificados (admin vê todos, usuário comum vê apenas o próprio)
 */
export async function listCertificates(): Promise<Certificate[]> {
  const response = await apiClient.get<CertificateListResponse>('/api/v1/certificates')
  return response.data.data
}

/**
 * Upload de certificado A1 para um médico (admin only)
 */
export async function uploadCertificate(
  formData: FormData
): Promise<UploadCertificateResponse> {
  const response = await apiClient.post<UploadCertificateResponse>(
    '/api/v1/admin/certificates/upload',
    formData,
    {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    }
  )
  return response.data
}

/**
 * Deletar certificado
 */
export async function deleteCertificate(userId: string): Promise<UploadCertificateResponse> {
  const response = await apiClient.delete<UploadCertificateResponse>(
    `/api/v1/certificates/${userId}`
  )
  return response.data
}

/**
 * Obter status do certificado do médico logado
 */
export async function getCertificateStatus(): Promise<CertificateStatus> {
  const response = await apiClient.get<CertificateStatus>('/api/v1/certificates/status')
  return response.data
}

/**
 * Ativar/desativar certificado (admin only)
 */
export async function toggleCertificateActive(userId: string): Promise<{
  success: boolean
  certificateActive: boolean
  message: string
}> {
  const response = await apiClient.patch(`/api/v1/admin/certificates/${userId}/toggle`)
  return response.data
}
