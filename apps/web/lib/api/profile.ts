import { apiClient } from '../api-client'
import type { User } from './users'

export interface UpdateProfileRequest {
  name?: string
  professionalPhone?: string | null
  professionalAddress?: string | null
  specialty?: string | null
  crm?: string | null
  crmUF?: string | null
  rqe?: string | null
}

export async function updateProfile(data: UpdateProfileRequest): Promise<User> {
  const response = await apiClient.put('/api/v1/profile', data)
  return response
}

export async function getProfile(): Promise<User> {
  const response = await apiClient.get('/api/v1/profile')
  return response
}
