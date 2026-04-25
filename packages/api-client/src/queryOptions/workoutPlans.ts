import { queryOptions } from '@tanstack/react-query';
import { api } from '../fetcher';
import { queryKeys } from '../queryKeys';

export interface WorkoutPlanSummary {
  id: string;
  patientId: string;
  name: string;
  startDate: string;
  endDate?: string;
  status: 'draft' | 'active' | 'completed' | 'archived';
}

export interface WorkoutSessionExercise {
  exerciseId: string;
  exerciseName: string;
  sets: number;
  reps: string;
  rest: string;
  notes?: string;
  gifUrl?: string;
}

export interface WorkoutPlanSession {
  id: string;
  name: string;
  order: number;
  exercises: WorkoutSessionExercise[];
}

export interface WorkoutPlanDetail extends WorkoutPlanSummary {
  sessions: WorkoutPlanSession[];
  publicCode?: string;
  notes?: string;
}

/**
 * Lista planos de treino do paciente atualmente selecionado.
 * Backend escopa via User.SelectedPatientID; patientId compõe queryKey.
 */
export const patientWorkoutPlansOptions = (patientId: string) =>
  queryOptions({
    queryKey: queryKeys.patients.workoutPlans(patientId),
    queryFn: ({ signal }) =>
      api.get<WorkoutPlanSummary[]>('/api/v1/workout-plans', { signal }),
    enabled: Boolean(patientId),
  });

export const workoutPlanOptions = (id: string) =>
  queryOptions({
    queryKey: queryKeys.workoutPlans.detail(id),
    queryFn: ({ signal }) =>
      api.get<WorkoutPlanDetail>(`/api/v1/workout-plans/${id}`, { signal }),
    enabled: Boolean(id),
  });

export const publicWorkoutPlanOptions = (code: string) =>
  queryOptions({
    queryKey: queryKeys.workoutPlans.public(code),
    queryFn: ({ signal }) =>
      api.get<WorkoutPlanDetail>(`/api/v1/workout-plans/public/${code}`, {
        signal,
        skipAuth: true,
      }),
    enabled: Boolean(code),
  });

export interface CreateWorkoutPlanInput {
  name: string;
  status?: WorkoutPlanSummary['status'];
  startDate?: string;
  endDate?: string;
  notes?: string;
  sessions?: Array<{
    name: string;
    order: number;
    exercises: Array<{
      exerciseId: string;
      sets: number;
      reps: string;
      rest: string;
      notes?: string;
    }>;
  }>;
}

export const workoutPlanMutations = {
  create: (body: CreateWorkoutPlanInput) =>
    api.post<WorkoutPlanDetail>('/api/v1/workout-plans', body),
  update: (id: string, body: Partial<CreateWorkoutPlanInput>) =>
    api.put<WorkoutPlanDetail>(`/api/v1/workout-plans/${id}`, body),
};
