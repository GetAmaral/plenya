import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { apiClient } from '../api-client'
import type { Exercise } from './exercise-api'

export type WorkoutObjective = 'hypertrophy' | 'strength' | 'endurance' | 'weight_loss' | 'conditioning' | 'rehabilitation'
export type WorkoutIntensity = 'very_light' | 'light' | 'moderate' | 'high' | 'very_high'
export type ExercisePhase = 'warmup' | 'main' | 'cooldown'
export type ExerciseCadence = 'normal' | 'slow' | 'paused' | 'explosive' | 'free'

export interface WorkoutSessionExercise {
  id: string
  sessionId: string
  exerciseId: string
  phase: ExercisePhase
  order: number
  sets: number
  reps: string
  cadence: ExerciseCadence
  restBetweenSetsSec: number
  restBetweenExercisesSec: number
  notes?: string
  exercise?: Exercise
}

export interface WorkoutPlanSession {
  id: string
  planId: string
  name: string
  order: number
  exercises: WorkoutSessionExercise[]
}

export interface WorkoutPlan {
  id: string
  patientId: string
  createdById: string
  name: string
  objective: WorkoutObjective
  intensity: WorkoutIntensity
  durationMinutes: number
  weeklyFrequency: number
  publicCode: string
  htmlContent?: string
  isActive: boolean
  displayTitle: string
  sessions: WorkoutPlanSession[]
  createdAt: string
  updatedAt: string
}

export interface CreateSessionExerciseDTO {
  exerciseId: string
  phase: ExercisePhase
  order: number
  sets: number
  reps: string
  cadence: ExerciseCadence
  restBetweenSetsSec: number
  restBetweenExercisesSec: number
  notes?: string
}

export interface CreateSessionDTO {
  name: string
  order: number
  exercises: CreateSessionExerciseDTO[]
}

export interface CreateWorkoutPlanDTO {
  patientId?: string
  name: string
  objective: WorkoutObjective
  intensity: WorkoutIntensity
  durationMinutes: number
  weeklyFrequency: number
  sessions: CreateSessionDTO[]
}

export interface UpdateWorkoutPlanDTO {
  name?: string
  objective?: WorkoutObjective
  intensity?: WorkoutIntensity
  durationMinutes?: number
  weeklyFrequency?: number
  isActive?: boolean
}

// Query key factory
export const workoutPlanKeys = {
  all: ['workout-plans'] as const,
  lists: () => [...workoutPlanKeys.all, 'list'] as const,
  list: (params?: { limit?: number; offset?: number }) => [...workoutPlanKeys.lists(), params] as const,
  details: () => [...workoutPlanKeys.all, 'detail'] as const,
  detail: (id: string) => [...workoutPlanKeys.details(), id] as const,
  public: (code: string) => [...workoutPlanKeys.all, 'public', code] as const,
}

export function useWorkoutPlans(limit = 20, offset = 0) {
  return useQuery({
    queryKey: workoutPlanKeys.list({ limit, offset }),
    queryFn: () => apiClient.get<WorkoutPlan[]>(`/api/v1/workout-plans?limit=${limit}&offset=${offset}`),
  })
}

export function useWorkoutPlan(id: string) {
  return useQuery({
    queryKey: workoutPlanKeys.detail(id),
    queryFn: () => apiClient.get<WorkoutPlan>(`/api/v1/workout-plans/${id}`),
    enabled: !!id,
  })
}

export function usePublicWorkoutPlan(code: string) {
  return useQuery({
    queryKey: workoutPlanKeys.public(code),
    queryFn: () => apiClient.get<WorkoutPlan>(`/api/v1/workout-plans/public/${code}`),
    enabled: !!code,
  })
}

export function useCreateWorkoutPlan() {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: (data: CreateWorkoutPlanDTO) =>
      apiClient.post<WorkoutPlan>('/api/v1/workout-plans', data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: workoutPlanKeys.all })
    },
  })
}

export function useUpdateWorkoutPlan() {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: UpdateWorkoutPlanDTO }) =>
      apiClient.put<WorkoutPlan>(`/api/v1/workout-plans/${id}`, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: workoutPlanKeys.all })
    },
  })
}

export function useGenerateHtml() {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: (id: string) =>
      apiClient.post<{ html: string }>(`/api/v1/workout-plans/${id}/generate-html`, {}),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: workoutPlanKeys.all })
    },
  })
}
