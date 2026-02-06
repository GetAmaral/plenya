import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { apiClient } from '../api-client'

// Types based on Go models
export interface ScoreLevel {
  id: string
  level: number
  name: string
  lowerLimit?: string
  upperLimit?: string
  operator: '=' | '>' | '>=' | '<' | '<=' | 'between'
  clinicalRelevance?: string
  patientExplanation?: string
  conduct?: string
  lastReview?: string
  itemId: string
  createdAt: string
  updatedAt: string
}

export interface ScoreItem {
  id: string
  name: string
  unit?: string
  unitConversion?: string
  clinicalRelevance?: string
  patientExplanation?: string
  conduct?: string
  lastReview?: string
  points?: number
  order: number
  subgroupId: string
  parentItemId?: string
  subgroup?: ScoreSubgroup
  levels?: ScoreLevel[]
  childItems?: ScoreItem[]
  createdAt: string
  updatedAt: string
}

export interface ScoreSubgroup {
  id: string
  name: string
  order: number
  groupId: string
  group?: ScoreGroup
  items?: ScoreItem[]
  createdAt: string
  updatedAt: string
}

export interface ScoreGroup {
  id: string
  name: string
  order: number
  subgroups?: ScoreSubgroup[]
  createdAt: string
  updatedAt: string
}

// DTOs for create/update operations
export interface CreateScoreGroupDTO {
  name: string
  order?: number
}

export interface UpdateScoreGroupDTO {
  name?: string
  order?: number
}

export interface CreateScoreSubgroupDTO {
  name: string
  groupId: string
  order?: number
}

export interface UpdateScoreSubgroupDTO {
  name?: string
  order?: number
}

export interface CreateScoreItemDTO {
  name: string
  unit?: string
  unitConversion?: string
  clinicalRelevance?: string
  patientExplanation?: string
  conduct?: string
  points?: number
  subgroupId: string
  parentItemId?: string
  order?: number
}

export interface UpdateScoreItemDTO {
  name?: string
  unit?: string
  unitConversion?: string
  clinicalRelevance?: string
  patientExplanation?: string
  conduct?: string
  points?: number
  order?: number
  subgroupId?: string
  parentItemId?: string | null
}

export interface CreateScoreLevelDTO {
  level: number
  name: string
  lowerLimit?: string
  upperLimit?: string
  operator: '=' | '>' | '>=' | '<' | '<=' | 'between'
  clinicalRelevance?: string
  patientExplanation?: string
  conduct?: string
  itemId: string
}

export interface UpdateScoreLevelDTO {
  level?: number
  name?: string
  lowerLimit?: string
  upperLimit?: string
  operator?: '=' | '>' | '>=' | '<' | '<=' | 'between'
  clinicalRelevance?: string
  patientExplanation?: string
  conduct?: string
}

// Query keys
export const scoreKeys = {
  all: ['scores'] as const,
  groups: () => [...scoreKeys.all, 'groups'] as const,
  group: (id: string) => [...scoreKeys.groups(), id] as const,
  groupTree: (id: string) => [...scoreKeys.group(id), 'tree'] as const,
  allGroupTrees: () => [...scoreKeys.groups(), 'trees'] as const,
  subgroups: () => [...scoreKeys.all, 'subgroups'] as const,
  subgroup: (id: string) => [...scoreKeys.subgroups(), id] as const,
  subgroupsByGroup: (groupId: string) => [...scoreKeys.subgroups(), 'group', groupId] as const,
  items: () => [...scoreKeys.all, 'items'] as const,
  item: (id: string) => [...scoreKeys.items(), id] as const,
  itemsBySubgroup: (subgroupId: string) => [...scoreKeys.items(), 'subgroup', subgroupId] as const,
  levels: () => [...scoreKeys.all, 'levels'] as const,
  level: (id: string) => [...scoreKeys.levels(), id] as const,
  levelsByItem: (itemId: string) => [...scoreKeys.levels(), 'item', itemId] as const,
}

// ============================================================
// Score Groups
// ============================================================

export function useScoreGroups() {
  return useQuery({
    queryKey: scoreKeys.groups(),
    queryFn: () => apiClient.get<ScoreGroup[]>('/api/v1/score-groups'),
    staleTime: 5 * 60 * 1000, // 5 minutes
  })
}

export function useScoreGroup(id: string) {
  return useQuery({
    queryKey: scoreKeys.group(id),
    queryFn: () => apiClient.get<ScoreGroup>(`/api/v1/score-groups/${id}`),
    enabled: !!id,
  })
}

export function useScoreGroupTree(id: string) {
  return useQuery({
    queryKey: scoreKeys.groupTree(id),
    queryFn: () => apiClient.get<ScoreGroup>(`/api/v1/score-groups/${id}/tree`),
    enabled: !!id,
  })
}

export function useAllScoreGroupTrees() {
  return useQuery({
    queryKey: scoreKeys.allGroupTrees(),
    queryFn: () => apiClient.get<ScoreGroup[]>('/api/v1/score-groups/tree'),
    staleTime: 5 * 60 * 1000, // 5 minutes
    notifyOnChangeProps: 'all', // Força re-render em qualquer mudança
    structuralSharing: false, // Desabilita otimização que pode impedir re-render
  })
}

export function useCreateScoreGroup() {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: (data: CreateScoreGroupDTO) =>
      apiClient.post<ScoreGroup>('/api/v1/score-groups', data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: scoreKeys.groups() })
      queryClient.invalidateQueries({ queryKey: scoreKeys.allGroupTrees() })
    },
  })
}

export function useUpdateScoreGroup() {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: UpdateScoreGroupDTO }) =>
      apiClient.put<ScoreGroup>(`/api/v1/score-groups/${id}`, data),
    onSuccess: (_, { id }) => {
      queryClient.invalidateQueries({ queryKey: scoreKeys.group(id) })
      queryClient.invalidateQueries({ queryKey: scoreKeys.groups() })
      queryClient.invalidateQueries({ queryKey: scoreKeys.allGroupTrees() })
    },
  })
}

export function useDeleteScoreGroup() {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: (id: string) => apiClient.delete(`/api/v1/score-groups/${id}`),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: scoreKeys.groups() })
      queryClient.invalidateQueries({ queryKey: scoreKeys.allGroupTrees() })
    },
  })
}

// ============================================================
// Score Subgroups
// ============================================================

export function useScoreSubgroup(id: string) {
  return useQuery({
    queryKey: scoreKeys.subgroup(id),
    queryFn: () => apiClient.get<ScoreSubgroup>(`/api/v1/score-subgroups/${id}`),
    enabled: !!id,
  })
}

export function useSubgroupsByGroup(groupId: string) {
  return useQuery({
    queryKey: scoreKeys.subgroupsByGroup(groupId),
    queryFn: () => apiClient.get<ScoreSubgroup[]>(`/api/v1/score-groups/${groupId}/subgroups`),
    enabled: !!groupId,
  })
}

export function useCreateScoreSubgroup() {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: async (data: CreateScoreSubgroupDTO) => {
      const response = await apiClient.post<ScoreSubgroup>('/api/v1/score-subgroups', data)
      return response as any
    },
    onSuccess: (subgroup) => {
      queryClient.invalidateQueries({ queryKey: scoreKeys.subgroupsByGroup(subgroup.groupId) })
      queryClient.invalidateQueries({ queryKey: scoreKeys.groups() })
      queryClient.invalidateQueries({ queryKey: scoreKeys.allGroupTrees() })
    },
  })
}

export function useUpdateScoreSubgroup() {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: async ({ id, data }: { id: string; data: UpdateScoreSubgroupDTO }) => {
      const response = await apiClient.put<ScoreSubgroup>(`/api/v1/score-subgroups/${id}`, data)
      return response as any
    },
    onSuccess: (subgroup, { id }) => {
      queryClient.invalidateQueries({ queryKey: scoreKeys.subgroup(id) })
      queryClient.invalidateQueries({ queryKey: scoreKeys.subgroupsByGroup(subgroup.groupId) })
      queryClient.invalidateQueries({ queryKey: scoreKeys.groups() })
      queryClient.invalidateQueries({ queryKey: scoreKeys.allGroupTrees() })
    },
  })
}

export function useDeleteScoreSubgroup() {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: (id: string) => apiClient.delete(`/api/v1/score-subgroups/${id}`),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: scoreKeys.subgroups() })
      queryClient.invalidateQueries({ queryKey: scoreKeys.groups() })
      queryClient.invalidateQueries({ queryKey: scoreKeys.allGroupTrees() })
    },
  })
}

// ============================================================
// Score Items
// ============================================================

export function useScoreItem(id: string) {
  return useQuery({
    queryKey: scoreKeys.item(id),
    queryFn: () => apiClient.get<ScoreItem>(`/api/v1/score-items/${id}`),
    enabled: !!id,
  })
}

export function useItemsBySubgroup(subgroupId: string) {
  return useQuery({
    queryKey: scoreKeys.itemsBySubgroup(subgroupId),
    queryFn: () => apiClient.get<ScoreItem[]>(`/api/v1/score-subgroups/${subgroupId}/items`),
    enabled: !!subgroupId,
  })
}

export function useCreateScoreItem() {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: async (data: CreateScoreItemDTO) => {
      const response = await apiClient.post<ScoreItem>('/api/v1/score-items', data)
      return response as any
    },
    onSuccess: (item) => {
      queryClient.invalidateQueries({ queryKey: scoreKeys.itemsBySubgroup(item.subgroupId) })
      queryClient.invalidateQueries({ queryKey: scoreKeys.groups() })
      queryClient.invalidateQueries({ queryKey: scoreKeys.allGroupTrees() })
    },
  })
}

export function useUpdateScoreItem() {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: async ({ id, data }: { id: string; data: UpdateScoreItemDTO }) => {
      const response = await apiClient.put<ScoreItem>(`/api/v1/score-items/${id}`, data)
      return response as any
    },
    onSuccess: (updatedItem) => {
      queryClient.setQueryData<ScoreGroup[]>(
        scoreKeys.allGroupTrees(),
        (old) => {
          if (!old) return old

          return old.map(g => ({
            ...g,
            subgroups: g.subgroups?.map(sg => ({
              ...sg,
              items: sg.items?.map(item =>
                item.id === updatedItem.id ? { ...item, ...updatedItem } : item
              )
            }))
          }))
        }
      )
    },
  })
}

export function useDeleteScoreItem() {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: (id: string) => apiClient.delete(`/api/v1/score-items/${id}`),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: scoreKeys.items() })
      queryClient.invalidateQueries({ queryKey: scoreKeys.groups() })
      queryClient.invalidateQueries({ queryKey: scoreKeys.allGroupTrees() })
    },
  })
}

// ============================================================
// Score Levels
// ============================================================

export function useScoreLevel(id: string) {
  return useQuery({
    queryKey: scoreKeys.level(id),
    queryFn: () => apiClient.get<ScoreLevel>(`/api/v1/score-levels/${id}`),
    enabled: !!id,
  })
}

export function useLevelsByItem(itemId: string) {
  return useQuery({
    queryKey: scoreKeys.levelsByItem(itemId),
    queryFn: () => apiClient.get<ScoreLevel[]>(`/api/v1/score-items/${itemId}/levels`),
    enabled: !!itemId,
  })
}

export function useCreateScoreLevel() {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: async (data: CreateScoreLevelDTO) => {
      const response = await apiClient.post<ScoreLevel>('/api/v1/score-levels', data)
      return response as any
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: scoreKeys.allGroupTrees() })
      queryClient.invalidateQueries({ queryKey: scoreKeys.groups() })
    },
  })
}

export function useUpdateScoreLevel() {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: async ({ id, data }: { id: string; data: UpdateScoreLevelDTO }) => {
      const response = await apiClient.put<ScoreLevel>(`/api/v1/score-levels/${id}`, data)
      return response as any
    },
    onSuccess: (updatedLevel) => {
      queryClient.setQueryData<ScoreGroup[]>(
        scoreKeys.allGroupTrees(),
        (old) => {
          if (!old) return old

          // Atualizar apenas o level específico
          return old.map(g => ({
            ...g,
            subgroups: g.subgroups?.map(sg => ({
              ...sg,
              items: sg.items?.map(item => ({
                ...item,
                levels: item.levels?.map(lv =>
                  lv.id === updatedLevel.id ? updatedLevel : lv
                )
              }))
            }))
          }))
        }
      )
    },
  })
}

export function useDeleteScoreLevel() {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: (id: string) => apiClient.delete(`/api/v1/score-levels/${id}`),
    onMutate: async (id) => {
      // Cancel any outgoing refetches
      await queryClient.cancelQueries({ queryKey: scoreKeys.allGroupTrees() })

      // Snapshot the previous value
      const previousData = queryClient.getQueryData<ScoreGroup[]>(scoreKeys.allGroupTrees())

      // Optimistically update by removing the level
      queryClient.setQueryData<ScoreGroup[]>(
        scoreKeys.allGroupTrees(),
        (oldData) => {
          if (!oldData) return oldData

          return oldData.map(group => {
            if (!group.subgroups) return group

            return {
              ...group,
              subgroups: group.subgroups.map(subgroup => {
                if (!subgroup.items) return subgroup

                return {
                  ...subgroup,
                  items: subgroup.items.map(item => {
                    if (!item.levels) return item

                    return {
                      ...item,
                      levels: item.levels.filter(level => level.id !== id)
                    }
                  })
                }
              })
            }
          })
        }
      )

      return { previousData }
    },
    onError: (_err, _id, context) => {
      // Rollback on error
      if (context?.previousData) {
        queryClient.setQueryData(scoreKeys.allGroupTrees(), context.previousData)
      }
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: scoreKeys.levels() })
    },
  })
}
