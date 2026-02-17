// Core domain types for Plenya EMR

export interface Method {
  id: string
  name: string
  shortName: string
  description?: string
  version?: string
  color?: string
  order: number
  isDefault: boolean
  letters?: MethodLetter[]
  createdAt: string
  updatedAt: string
}

export interface MethodLetter {
  id: string
  methodId: string
  code: string
  name: string
  description?: string
  clinicalRelevance?: string
  patientExplanation?: string
  conduct?: string
  color?: string
  icon?: string
  order: number
  pillars?: MethodPillar[]
  createdAt: string
  updatedAt: string
}

export interface MethodPillar {
  id: string
  letterId: string
  name: string
  description?: string
  clinicalRelevance?: string
  patientExplanation?: string
  conduct?: string
  order: number
  scoreItems?: ScoreItem[]
  createdAt: string
  updatedAt: string
}

export interface ScoreItem {
  id: string
  groupId: string
  subgroupId: string
  name: string
  description?: string
  clinicalRelevance?: string
  patientExplanation?: string
  conduct?: string
  unit?: string
  labTestCode?: string
  order: number
  lastReview?: string
  gender?: string
  ageMin?: number
  ageMax?: number
  postMenopause?: boolean
  levels?: ScoreLevel[]
  methodPillars?: MethodPillar[]
  createdAt: string
  updatedAt: string
}

export interface ScoreLevel {
  id: string
  itemId: string
  level: string
  color: string
  operatorType: string
  minValue?: number
  maxValue?: number
  textValue?: string
  description?: string
  clinicalRelevance?: string
  patientExplanation?: string
  conduct?: string
  order: number
  lastReview?: string
  createdAt: string
  updatedAt: string
}

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

// Re-export generated schemas if needed
export * from './generated/api-schemas'
