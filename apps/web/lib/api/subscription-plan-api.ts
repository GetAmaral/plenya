import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { apiClient } from '../api-client'
import type { SubscriptionPlan } from '@plenya/types'

// Query keys
export const subscriptionPlanKeys = {
  all: ['subscription-plans'] as const,
  active: () => [...subscriptionPlanKeys.all, 'active'] as const,
  detail: (id: string) => [...subscriptionPlanKeys.all, id] as const,
}

// ===== Subscription Plan Queries =====

/**
 * Get all subscription plans
 * @returns All subscription plans (active and inactive)
 */
export function useSubscriptionPlans() {
  return useQuery({
    queryKey: subscriptionPlanKeys.all,
    queryFn: () => apiClient.get<SubscriptionPlan[]>('/api/v1/subscription-plans'),
    staleTime: 5 * 60 * 1000, // 5 minutes
  })
}

/**
 * Get only active subscription plans
 * @returns Active subscription plans available for selection
 */
export function useActiveSubscriptionPlans() {
  return useQuery({
    queryKey: subscriptionPlanKeys.active(),
    queryFn: () => apiClient.get<SubscriptionPlan[]>('/api/v1/subscription-plans/active'),
    staleTime: 2 * 60 * 1000, // 2 minutes
  })
}

/**
 * Get a single subscription plan by ID
 * @param id - Subscription plan UUID
 */
export function useSubscriptionPlan(id: string | undefined) {
  return useQuery({
    queryKey: subscriptionPlanKeys.detail(id!),
    queryFn: () => apiClient.get<SubscriptionPlan>(`/api/v1/subscription-plans/${id}`),
    enabled: !!id,
    staleTime: 5 * 60 * 1000,
  })
}

// ===== Subscription Plan Mutations =====

export interface CreateSubscriptionPlanDTO {
  name: string
  description: string
  features: string // JSON string of features array
  price: number
  currency: string
  billingCycle: 'monthly' | 'quarterly' | 'yearly' | 'one_time'
  methodId?: string
  trialPeriodDays: number
  order: number
}

export interface UpdateSubscriptionPlanDTO {
  name?: string
  description?: string
  features?: string // JSON string of features array
  price?: number
  currency?: string
  billingCycle?: 'monthly' | 'quarterly' | 'yearly' | 'one_time'
  methodId?: string
  trialPeriodDays?: number
  order?: number
}

/**
 * Create a new subscription plan
 * Invalidates: all subscription plans cache
 */
export function useCreateSubscriptionPlan() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (data: CreateSubscriptionPlanDTO) =>
      apiClient.post<SubscriptionPlan>('/api/v1/subscription-plans', data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: subscriptionPlanKeys.all })
    },
  })
}

/**
 * Update an existing subscription plan
 * Invalidates: all plans cache + specific plan detail
 */
export function useUpdateSubscriptionPlan() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: UpdateSubscriptionPlanDTO }) =>
      apiClient.put<SubscriptionPlan>(`/api/v1/subscription-plans/${id}`, data),
    onSuccess: (data) => {
      queryClient.invalidateQueries({ queryKey: subscriptionPlanKeys.all })
      queryClient.invalidateQueries({ queryKey: subscriptionPlanKeys.detail(data.id) })
    },
  })
}

/**
 * Delete a subscription plan
 * Invalidates: all subscription plans cache
 */
export function useDeleteSubscriptionPlan() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (id: string) => apiClient.delete(`/api/v1/subscription-plans/${id}`),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: subscriptionPlanKeys.all })
    },
  })
}

/**
 * Activate a subscription plan
 * Invalidates: all plans cache + active plans cache + specific plan detail
 */
export function useActivateSubscriptionPlan() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (id: string) =>
      apiClient.post<SubscriptionPlan>(`/api/v1/subscription-plans/${id}/activate`, {}),
    onSuccess: (data) => {
      queryClient.invalidateQueries({ queryKey: subscriptionPlanKeys.all })
      queryClient.invalidateQueries({ queryKey: subscriptionPlanKeys.active() })
      queryClient.invalidateQueries({ queryKey: subscriptionPlanKeys.detail(data.id) })
    },
  })
}

/**
 * Deactivate a subscription plan
 * Invalidates: all plans cache + active plans cache + specific plan detail
 */
export function useDeactivateSubscriptionPlan() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (id: string) =>
      apiClient.post<SubscriptionPlan>(`/api/v1/subscription-plans/${id}/deactivate`, {}),
    onSuccess: (data) => {
      queryClient.invalidateQueries({ queryKey: subscriptionPlanKeys.all })
      queryClient.invalidateQueries({ queryKey: subscriptionPlanKeys.active() })
      queryClient.invalidateQueries({ queryKey: subscriptionPlanKeys.detail(data.id) })
    },
  })
}
