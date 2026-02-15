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

// Re-export generated schemas if needed
export * from './generated/api-schemas'
