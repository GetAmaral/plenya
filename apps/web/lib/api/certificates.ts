import { apiClient } from '../api-client'

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
  // apiClient.get devolve o corpo JSON já parseado ({ data: [...] }), não um wrapper axios.
  const response = await apiClient.get<CertificateListResponse>('/api/v1/certificates')
  return response.data
}

/**
 * Upload de certificado A1 para um médico (admin only)
 */
export async function uploadCertificate(
  formData: FormData
): Promise<UploadCertificateResponse> {
  // Não setar Content-Type: o browser precisa gerar o boundary do multipart.
  // Com o header fixo sem boundary, o parser do Fiber não lê os campos e o
  // backend responde "doctorId and password are required".
  const response = await apiClient.post<UploadCertificateResponse>(
    '/api/v1/admin/certificates/upload',
    formData
  )
  return response
}

/**
 * Deletar certificado
 */
export async function deleteCertificate(userId: string): Promise<UploadCertificateResponse> {
  const response = await apiClient.delete<UploadCertificateResponse>(
    `/api/v1/certificates/${userId}`
  )
  return response
}

/**
 * Obter status do certificado do médico logado
 */
export async function getCertificateStatus(): Promise<CertificateStatus> {
  const response = await apiClient.get<CertificateStatus>('/api/v1/certificates/status')
  return response
}

/**
 * Ativar/desativar certificado (admin only)
 */
export async function toggleCertificateActive(userId: string): Promise<{
  success: boolean
  certificateActive: boolean
  message: string
}> {
  return apiClient.patch<{
    success: boolean
    certificateActive: boolean
    message: string
  }>(`/api/v1/admin/certificates/${userId}/toggle`)
}
