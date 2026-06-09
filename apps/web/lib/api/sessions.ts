import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { apiClient } from '../api-client'

export interface Session {
  id: string
  userAgent?: string
  ipAddress?: string
  remember: boolean
  createdAt: string
  lastUsedAt?: string
  expiresAt: string
}

export const sessionKeys = {
  all: ['sessions'] as const,
}

export function useSessions() {
  return useQuery({
    queryKey: sessionKeys.all,
    queryFn: () => apiClient.get<Session[]>('/api/v1/auth/sessions'),
    staleTime: 30 * 1000,
  })
}

export function useRevokeSession() {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: (id: string) => apiClient.delete(`/api/v1/auth/sessions/${id}`),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: sessionKeys.all })
    },
  })
}
