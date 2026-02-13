'use client'

import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { apiClient } from '../api-client'
import { toast } from 'sonner'

// Types
export interface PatientScoreSnapshot {
  id: string
  patientId: string
  calculatedByUserId: string
  calculatedAt: string
  totalActualPoints: number
  totalPossiblePoints: number
  totalScorePercentage: number
  itemsEvaluatedCount: number
  itemsNotEvaluatedCount: number
  notes?: string
  groupResults?: PatientScoreGroupResult[]
  itemResults?: PatientScoreItemResult[]
  calculatedByUser?: {
    id: string
    name: string
  }
  patient?: {
    id: string
    name: string
  }
  createdAt: string
  updatedAt: string
}

export interface PatientScoreGroupResult {
  id: string
  snapshotId: string
  groupId: string
  actualPoints: number
  possiblePoints: number
  scorePercentage: number
  itemsEvaluatedCount: number
  group?: {
    id: string
    name: string
    description?: string
  }
  createdAt: string
  updatedAt: string
}

export interface PatientScoreItemResult {
  id: string
  snapshotId: string
  itemId: string
  groupId: string
  status: 'evaluated' | 'not_applicable' | 'no_data_available'
  dataSource?: 'lab_result' | 'anamnesis_item'
  labResultId?: string
  anamnesisItemId?: string
  valueUsed?: number
  levelMatchedId?: string
  levelNumber?: number
  maxPoints: number
  actualPoints: number
  notEvaluatedReason?: string
  item?: {
    id: string
    name: string
    unit?: string
  }
  levelMatched?: {
    id: string
    level: number
    name: string
  }
  createdAt: string
  updatedAt: string
}

export interface CalculateSnapshotDTO {
  patientId: string
  notes?: string
}

// API Functions
export const healthScoreApi = {
  // Calculate new snapshot
  calculateSnapshot: async (
    patientId: string,
    notes?: string
  ): Promise<PatientScoreSnapshot> => {
    return apiClient.post<PatientScoreSnapshot>(
      `/api/v1/patients/${patientId}/score-snapshots`,
      { patientId, notes }
    )
  },

  // Get all snapshots for a patient
  getSnapshotsByPatientId: async (patientId: string): Promise<PatientScoreSnapshot[]> => {
    return apiClient.get<PatientScoreSnapshot[]>(
      `/api/v1/patients/${patientId}/score-snapshots`
    )
  },

  // Get latest snapshot for a patient
  getLatestSnapshot: async (patientId: string): Promise<PatientScoreSnapshot> => {
    return apiClient.get<PatientScoreSnapshot>(
      `/api/v1/patients/${patientId}/score-snapshots/latest`
    )
  },

  // Get snapshot by ID
  getSnapshotById: async (id: string): Promise<PatientScoreSnapshot> => {
    return apiClient.get<PatientScoreSnapshot>(`/api/v1/score-snapshots/${id}`)
  },

  // Delete snapshot
  deleteSnapshot: async (id: string): Promise<void> => {
    return apiClient.delete(`/api/v1/score-snapshots/${id}`)
  },
}

// React Query Hooks

// Query keys factory
export const healthScoreKeys = {
  all: ['health-scores'] as const,
  lists: () => [...healthScoreKeys.all, 'list'] as const,
  list: (patientId: string) => [...healthScoreKeys.lists(), patientId] as const,
  details: () => [...healthScoreKeys.all, 'detail'] as const,
  detail: (id: string) => [...healthScoreKeys.details(), id] as const,
  latest: (patientId: string) => [...healthScoreKeys.all, 'latest', patientId] as const,
}

// Get all snapshots for a patient
export function useHealthScores(patientId: string) {
  return useQuery({
    queryKey: healthScoreKeys.list(patientId),
    queryFn: () => healthScoreApi.getSnapshotsByPatientId(patientId),
    enabled: !!patientId,
  })
}

// Get latest snapshot for a patient
export function useLatestHealthScore(patientId: string) {
  return useQuery({
    queryKey: healthScoreKeys.latest(patientId),
    queryFn: () => healthScoreApi.getLatestSnapshot(patientId),
    enabled: !!patientId,
    retry: false, // Don't retry if no snapshots found
  })
}

// Get snapshot by ID
export function useHealthScoreDetail(id: string) {
  return useQuery({
    queryKey: healthScoreKeys.detail(id),
    queryFn: () => healthScoreApi.getSnapshotById(id),
    enabled: !!id,
  })
}

// Calculate new snapshot
export function useCalculateHealthScore() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({ patientId, notes }: { patientId: string; notes?: string }) =>
      healthScoreApi.calculateSnapshot(patientId, notes),
    onSuccess: (data, variables) => {
      // Invalidate related queries
      queryClient.invalidateQueries({ queryKey: healthScoreKeys.list(variables.patientId) })
      queryClient.invalidateQueries({ queryKey: healthScoreKeys.latest(variables.patientId) })

      toast.success('Escore calculado com sucesso!')
    },
    onError: (error: any) => {
      toast.error(error?.message || 'Erro ao calcular escore')
    },
  })
}

// Delete snapshot
export function useDeleteHealthScore() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (id: string) => healthScoreApi.deleteSnapshot(id),
    onSuccess: () => {
      // Invalidate all lists (we don't know which patient this belongs to)
      queryClient.invalidateQueries({ queryKey: healthScoreKeys.lists() })
      toast.success('Snapshot deletado com sucesso!')
    },
    onError: (error: any) => {
      toast.error(error?.message || 'Erro ao deletar snapshot')
    },
  })
}
