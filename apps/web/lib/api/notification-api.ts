import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { apiClient } from '../api-client'
import type { Notification } from '@plenya/types'

// Query keys
export const notificationKeys = {
  all: ['notifications'] as const,
  unread: () => [...notificationKeys.all, 'unread'] as const,
  unreadCount: () => [...notificationKeys.all, 'unread', 'count'] as const,
}

// ===== Notification Queries =====

export function useNotifications(limit?: number) {
  return useQuery({
    queryKey: [...notificationKeys.all, limit],
    queryFn: () => {
      const params = limit ? `?limit=${limit}` : ''
      return apiClient.get<Notification[]>(`/api/v1/notifications${params}`)
    },
    staleTime: 1 * 60 * 1000, // 1 minute
    refetchInterval: 2 * 60 * 1000, // Refetch every 2 minutes
  })
}

export function useUnreadNotifications() {
  return useQuery({
    queryKey: notificationKeys.unread(),
    queryFn: () => apiClient.get<Notification[]>('/api/v1/notifications/unread'),
    staleTime: 30 * 1000, // 30 seconds
    refetchInterval: 1 * 60 * 1000, // Refetch every 1 minute
  })
}

export function useUnreadCount() {
  return useQuery({
    queryKey: notificationKeys.unreadCount(),
    queryFn: async () => {
      const data = await apiClient.get<{ count: number }>('/api/v1/notifications/unread/count')
      return data.count
    },
    staleTime: 30 * 1000, // 30 seconds
    refetchInterval: 1 * 60 * 1000, // Refetch every 1 minute
  })
}

// ===== Notification Mutations =====

export function useMarkAsRead() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (id: string) => apiClient.post(`/api/v1/notifications/${id}/read`, {}),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: notificationKeys.all })
      queryClient.invalidateQueries({ queryKey: notificationKeys.unread() })
      queryClient.invalidateQueries({ queryKey: notificationKeys.unreadCount() })
    },
  })
}

export function useMarkAllAsRead() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: () => apiClient.post('/api/v1/notifications/read-all', {}),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: notificationKeys.all })
      queryClient.invalidateQueries({ queryKey: notificationKeys.unread() })
      queryClient.invalidateQueries({ queryKey: notificationKeys.unreadCount() })
    },
  })
}

export function useDeleteNotification() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (id: string) => apiClient.delete(`/api/v1/notifications/${id}`),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: notificationKeys.all })
      queryClient.invalidateQueries({ queryKey: notificationKeys.unread() })
      queryClient.invalidateQueries({ queryKey: notificationKeys.unreadCount() })
    },
  })
}

export function useDeleteAllNotifications() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: () => apiClient.delete('/api/v1/notifications'),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: notificationKeys.all })
      queryClient.invalidateQueries({ queryKey: notificationKeys.unread() })
      queryClient.invalidateQueries({ queryKey: notificationKeys.unreadCount() })
    },
  })
}
