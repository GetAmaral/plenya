import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { apiClient } from '../api-client'

export interface Patient {
  id: string
  name: string
  cpf: string
  email?: string
  phone?: string
  birthDate: string
  gender: 'male' | 'female' | 'other'
  address?: string
  city?: string
  state?: string
  zipCode?: string
  municipality?: string
  createdAt: string
  updatedAt: string
}

// Query keys
export const patientKeys = {
  all: ['patients'] as const,
  detail: (id: string) => [...patientKeys.all, id] as const,
}

// ===== Patient Queries =====

export function usePatients() {
  return useQuery({
    queryKey: patientKeys.all,
    queryFn: () => apiClient.get<Patient[]>('/api/v1/patients'),
    staleTime: 5 * 60 * 1000, // 5 minutes
  })
}

export function usePatient(id: string) {
  return useQuery({
    queryKey: patientKeys.detail(id),
    queryFn: () => apiClient.get<Patient>(`/api/v1/patients/${id}`),
    staleTime: 5 * 60 * 1000, // 5 minutes
    enabled: !!id, // Only run if id is provided
  })
}

// ===== Patient Mutations =====

export function useCreatePatient() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (data: Partial<Patient>) => apiClient.post<Patient>('/api/v1/patients', data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: patientKeys.all })
    },
  })
}

export function useUpdatePatient() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: Partial<Patient> }) =>
      apiClient.put<Patient>(`/api/v1/patients/${id}`, data),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: patientKeys.all })
      queryClient.invalidateQueries({ queryKey: patientKeys.detail(variables.id) })
    },
  })
}

export function useDeletePatient() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (id: string) => apiClient.delete(`/api/v1/patients/${id}`),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: patientKeys.all })
    },
  })
}
