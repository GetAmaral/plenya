import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { apiClient } from '../api-client'
import type { PatientSubscription } from '@plenya/types'

// Query keys
export const subscriptionKeys = {
  all: ['subscriptions'] as const,
  active: () => [...subscriptionKeys.all, 'active'] as const,
  byPatient: (patientId: string) => [...subscriptionKeys.all, 'patient', patientId] as const,
  detail: (id: string) => [...subscriptionKeys.all, id] as const,
}

// ===== Subscription Queries =====

export function useAllSubscriptions() {
  return useQuery({
    queryKey: subscriptionKeys.all,
    queryFn: () => apiClient.get<PatientSubscription[]>('/api/v1/subscriptions'),
    staleTime: 5 * 60 * 1000, // 5 minutes
  })
}

export function useActiveSubscriptions() {
  return useQuery({
    queryKey: subscriptionKeys.active(),
    queryFn: () => apiClient.get<PatientSubscription[]>('/api/v1/subscriptions/active'),
    staleTime: 2 * 60 * 1000, // 2 minutes
  })
}

export function useSubscription(id: string | undefined) {
  return useQuery({
    queryKey: subscriptionKeys.detail(id!),
    queryFn: () => apiClient.get<PatientSubscription>(`/api/v1/subscriptions/${id}`),
    enabled: !!id,
    staleTime: 5 * 60 * 1000,
  })
}

export function usePatientSubscriptions(patientId: string | undefined) {
  return useQuery({
    queryKey: subscriptionKeys.byPatient(patientId!),
    queryFn: () => apiClient.get<PatientSubscription[]>(`/api/v1/patients/${patientId}/subscriptions`),
    enabled: !!patientId,
    staleTime: 2 * 60 * 1000,
  })
}

export function useActivePatientSubscription(patientId: string | undefined) {
  return useQuery({
    queryKey: [...subscriptionKeys.byPatient(patientId!), 'active'],
    queryFn: async () => {
      try {
        const response = await apiClient.get<PatientSubscription>(
          `/api/v1/patients/${patientId}/subscriptions/active`
        )
        return response
      } catch (error: any) {
        if (error.status === 204 || error.response?.status === 204) {
          return null // No active subscription
        }
        throw error
      }
    },
    enabled: !!patientId,
    staleTime: 5 * 60 * 1000,
    retry: false,
  })
}

// ===== Subscription Mutations =====

export interface CreateSubscriptionDTO {
  patientId: string
  subscriptionPlanId: string
  status: 'active' | 'inactive' | 'cancelled' | 'expired' | 'suspended' | 'trial'
  autoRenew: boolean
  startDate: string // YYYY-MM-DD
  endDate?: string
  trialEndDate?: string
  discountPercent: number
  discountReason?: string
  customPrice?: number
  customTrialDays?: number
  notes?: string
}

export interface UpdateSubscriptionDTO {
  subscriptionPlanId?: string
  status?: 'active' | 'inactive' | 'cancelled' | 'expired' | 'suspended' | 'trial'
  autoRenew?: boolean
  endDate?: string
  trialEndDate?: string
  nextBillingDate?: string
  discountPercent?: number
  discountReason?: string
  customPrice?: number
  customTrialDays?: number
  cancellationReason?: string
  notes?: string
}

export interface CancelSubscriptionDTO {
  cancellationReason?: string
}

export interface RenewSubscriptionDTO {
  endDate?: string
  nextBillingDate?: string
}

export function useCreateSubscription() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({ patientId, data }: { patientId: string; data: CreateSubscriptionDTO }) =>
      apiClient.post<PatientSubscription>(`/api/v1/patients/${patientId}/subscriptions`, data),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: subscriptionKeys.all })
      queryClient.invalidateQueries({ queryKey: subscriptionKeys.byPatient(variables.patientId) })
    },
  })
}

export function useUpdateSubscription() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: UpdateSubscriptionDTO }) =>
      apiClient.put<PatientSubscription>(`/api/v1/subscriptions/${id}`, data),
    onSuccess: (data) => {
      queryClient.invalidateQueries({ queryKey: subscriptionKeys.all })
      queryClient.invalidateQueries({ queryKey: subscriptionKeys.detail(data.id) })
      if (data.patientId) {
        queryClient.invalidateQueries({ queryKey: subscriptionKeys.byPatient(data.patientId) })
      }
    },
  })
}

export function useDeleteSubscription() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (id: string) => apiClient.delete(`/api/v1/subscriptions/${id}`),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: subscriptionKeys.all })
    },
  })
}

export function useCancelSubscription() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({ id, data }: { id: string; data?: CancelSubscriptionDTO }) =>
      apiClient.post<PatientSubscription>(`/api/v1/subscriptions/${id}/cancel`, data || {}),
    onSuccess: (data) => {
      queryClient.invalidateQueries({ queryKey: subscriptionKeys.all })
      queryClient.invalidateQueries({ queryKey: subscriptionKeys.detail(data.id) })
      if (data.patientId) {
        queryClient.invalidateQueries({ queryKey: subscriptionKeys.byPatient(data.patientId) })
      }
    },
  })
}

export function useRenewSubscription() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({ id, data }: { id: string; data?: RenewSubscriptionDTO }) =>
      apiClient.post<PatientSubscription>(`/api/v1/subscriptions/${id}/renew`, data || {}),
    onSuccess: (data) => {
      queryClient.invalidateQueries({ queryKey: subscriptionKeys.all })
      queryClient.invalidateQueries({ queryKey: subscriptionKeys.detail(data.id) })
      if (data.patientId) {
        queryClient.invalidateQueries({ queryKey: subscriptionKeys.byPatient(data.patientId) })
      }
    },
  })
}

export function useSuspendSubscription() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (id: string) =>
      apiClient.post<PatientSubscription>(`/api/v1/subscriptions/${id}/suspend`, {}),
    onSuccess: (data) => {
      queryClient.invalidateQueries({ queryKey: subscriptionKeys.all })
      queryClient.invalidateQueries({ queryKey: subscriptionKeys.detail(data.id) })
      if (data.patientId) {
        queryClient.invalidateQueries({ queryKey: subscriptionKeys.byPatient(data.patientId) })
      }
    },
  })
}

export function useActivateSubscription() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (id: string) =>
      apiClient.post<PatientSubscription>(`/api/v1/subscriptions/${id}/activate`, {}),
    onSuccess: (data) => {
      queryClient.invalidateQueries({ queryKey: subscriptionKeys.all })
      queryClient.invalidateQueries({ queryKey: subscriptionKeys.detail(data.id) })
      if (data.patientId) {
        queryClient.invalidateQueries({ queryKey: subscriptionKeys.byPatient(data.patientId) })
      }
    },
  })
}
