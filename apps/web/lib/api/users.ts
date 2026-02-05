import { apiClient } from '../api-client'

export interface User {
  id: string
  name: string
  email: string
  cpf?: string
  roles: string[]
  twoFactorEnabled: boolean
  createdAt: string
  updatedAt: string
}

export interface CreateUserRequest {
  name: string
  email: string
  password: string
  cpf?: string
  roles: string[]
}

export interface UpdateUserRequest {
  name?: string
  email?: string
  password?: string
  cpf?: string | null
  roles?: string[]
}

export interface UserListResponse {
  data: User[]
  total: number
  page: number
  limit: number
}

export async function getAllUsers(params: {
  role?: string
  limit?: number
  offset?: number
}): Promise<UserListResponse> {
  const { role, limit = 20, offset = 0 } = params
  const queryParams = new URLSearchParams({
    limit: limit.toString(),
    offset: offset.toString(),
  })

  if (role) {
    queryParams.append('role', role)
  }

  return apiClient.get<UserListResponse>(
    `/api/v1/admin/users?${queryParams.toString()}`
  )
}

export async function getUserById(id: string): Promise<User> {
  return apiClient.get<User>(`/api/v1/admin/users/${id}`)
}

export async function createUser(data: CreateUserRequest): Promise<User> {
  return apiClient.post<User>('/api/v1/admin/users', data)
}

export async function updateUser(
  id: string,
  data: UpdateUserRequest
): Promise<User> {
  return apiClient.put<User>(`/api/v1/admin/users/${id}`, data)
}

export async function deleteUser(id: string): Promise<void> {
  return apiClient.delete(`/api/v1/admin/users/${id}`)
}

/**
 * List users by role (helper function)
 */
export async function listByRole(role: string): Promise<User[]> {
  const response = await getAllUsers({ role, limit: 100 })
  return response.data
}
