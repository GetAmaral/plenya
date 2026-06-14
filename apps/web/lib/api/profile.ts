import { apiClient } from '../api-client'
import type { User } from './users'

export interface UpdateProfileRequest {
  name?: string
  professionalPhone?: string | null
  professionalAddress?: string | null
  specialty?: string | null
  gender?: string | null
  treatment?: string | null
  crm?: string | null
  crmUF?: string | null
  rqe?: string | null
}

export async function updateProfile(data: UpdateProfileRequest): Promise<User> {
  return apiClient.put<User>('/api/v1/profile', data)
}

export async function getProfile(): Promise<User> {
  return apiClient.get<User>('/api/v1/profile')
}
