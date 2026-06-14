// Core domain types for Plenya EMR
// Method/Letter/Pillar e Score Group/Subgroup/Item/Level migraram pro alias gerado (./models).
import type { Patient, Method } from './models'

export interface SubscriptionPlan {
  id: string
  name: string
  description: string
  features: string // JSON string
  price: number
  currency: string
  billingCycle: 'monthly' | 'quarterly' | 'yearly' | 'one_time'
  methodId?: string
  method?: Method
  isActive: boolean
  trialPeriodDays: number
  order: number
  createdAt: string
  updatedAt: string
}

export interface PatientSubscription {
  id: string
  patientId: string
  subscriptionPlanId: string
  planSnapshot: string // JSON string with plan data at subscription time
  status: 'active' | 'inactive' | 'cancelled' | 'expired' | 'suspended' | 'trial'
  autoRenew: boolean
  startDate: string
  endDate?: string
  trialEndDate?: string
  nextBillingDate?: string
  cancelledAt?: string
  discountPercent: number
  discountReason?: string
  customPrice?: number
  customTrialDays?: number
  cancellationReason?: string
  notes?: string
  renewalCount: number
  createdAt: string
  updatedAt: string
  patient?: Patient
  subscriptionPlan?: SubscriptionPlan
}

export interface Notification {
  id: string
  userId: string
  patientId?: string
  subscriptionId?: string
  type: 'trial_expiring' | 'renewal_upcoming' | 'subscription_expired' | 'payment_pending' | 'general'
  title: string
  message: string
  actionUrl?: string
  actionText?: string
  read: boolean
  readAt?: string
  createdAt: string
  updatedAt: string
  patient?: Patient
  subscription?: PatientSubscription
}

// Tipos gerados (openapi-typescript a partir dos Go models). `components['schemas'][...]`
// dá acesso a qualquer schema; os aliases ergonômicos ficam em ./models.
export type { components, paths, operations } from './generated/api-types'

// Aliases de model sobre os tipos gerados (ver models.ts).
export * from './models'
