import { queryOptions } from '@tanstack/react-query';
import { api } from '../fetcher';
import { queryKeys } from '../queryKeys';

// =============================================================================
// Types — espelham os DTOs do Go (PatientPortalService + PatientWorkoutsService)
// =============================================================================

export interface PatientMeProfile {
  id: string;
  userId: string;
  name: string;
  email?: string;
  phone?: string;
  birthDate?: string;
  gender?: 'male' | 'female' | 'other';
}

export interface PatientAppointment {
  id: string;
  scheduledAt: string;
  durationMinutes: number;
  type: string;
  status: string;
  doctorId: string;
  doctorName: string;
  isTelemedicine: boolean;
  dailyRoomUrl?: string;
  notes?: string;
  patientConfirmedAt?: string;
  confirmedAt?: string;
  cancelledAt?: string;
  minutesUntilStart: number;
}

export interface PatientWorkoutPlanSummary {
  id: string;
  name: string;
  displayTitle: string;
  objective: string;
  intensity: string;
  weeklyFrequency: number;
  isActive: boolean;
  totalSessions: number;
  createdAt: string;
}

export interface PatientWorkoutSessionTemplate {
  id: string;
  name: string;
  order: number;
  exercises: Array<{
    id: string;
    exerciseId: string;
    phase: 'warmup' | 'main' | 'cooldown';
    order: number;
    sets: number;
    reps: string;
    cadence: string;
    restBetweenSetsSec: number;
    restBetweenExercisesSec: number;
    notes?: string;
    exercise?: {
      id: string;
      name: string;
      category?: string;
      gifUrl?: string;
      videoUrl?: string;
    };
  }>;
}

export interface PatientWorkoutPlanDetail extends PatientWorkoutPlanSummary {
  sessions: PatientWorkoutSessionTemplate[];
}

export interface PatientWorkoutLog {
  id: string;
  sessionId: string;
  planExerciseId: string;
  exerciseId: string;
  setNumber: number;
  reps?: number;
  weight?: number;
  durationSec?: number;
  rpe?: number;
  notes?: string;
  createdAt: string;
}

export interface PatientWorkoutSession {
  id: string;
  patientId: string;
  planId: string;
  planSessionId: string;
  scheduledDate: string;
  completedAt?: string;
  notes?: string;
  createdAt: string;
  logs?: PatientWorkoutLog[];
  planSession?: PatientWorkoutSessionTemplate;
}

export interface HealthCheckIn {
  id: string;
  patientId: string;
  energy: number;
  pain: number;
  painLocation?: string;
  mood: number;
  sleepHours: number;
  sleepQuality: number;
  stress: number;
  notes?: string;
  createdAt: string;
}

export interface NotificationPreferences {
  userId: string;
  appointmentReminder: boolean;
  messageAlert: boolean;
  workoutReminder: boolean;
  workoutReminderTime: string;
}

// =============================================================================
// Query options
// =============================================================================

export const patientMeProfileOptions = () =>
  queryOptions({
    queryKey: queryKeys.patientMe.profile(),
    queryFn: ({ signal }) =>
      api.get<PatientMeProfile>('/api/v1/patient/me', { signal }),
    staleTime: 60_000,
  });

export const patientMeAppointmentsOptions = (rangeKind: 'upcoming' | 'past' = 'upcoming') =>
  queryOptions({
    queryKey: [...queryKeys.patientMe.appointments(), rangeKind] as const,
    queryFn: ({ signal }) =>
      api.get<PatientAppointment[]>(
        `/api/v1/patient/me/appointments?range=${rangeKind}`,
        { signal },
      ),
    staleTime: 30_000,
  });

export const patientMeWorkoutPlansOptions = () =>
  queryOptions({
    queryKey: queryKeys.patientMe.workoutPlans(),
    queryFn: ({ signal }) =>
      api.get<PatientWorkoutPlanSummary[]>('/api/v1/patient/me/workout-plans', { signal }),
    staleTime: 5 * 60_000,
  });

export const patientMeWorkoutPlanOptions = (id: string) =>
  queryOptions({
    queryKey: queryKeys.patientMe.workoutPlanDetail(id),
    queryFn: ({ signal }) =>
      api.get<PatientWorkoutPlanDetail>(`/api/v1/patient/me/workout-plans/${id}`, { signal }),
    staleTime: 5 * 60_000,
    enabled: Boolean(id),
  });

export const patientMeWorkoutSessionsOptions = (limit = 30) =>
  queryOptions({
    queryKey: [...queryKeys.patientMe.workoutSessions(), { limit }] as const,
    queryFn: ({ signal }) =>
      api.get<PatientWorkoutSession[]>(
        `/api/v1/patient/me/workout-sessions?limit=${limit}`,
        { signal },
      ),
    staleTime: 30_000,
  });

export const patientMeWorkoutSessionOptions = (id: string) =>
  queryOptions({
    queryKey: queryKeys.patientMe.workoutSessionDetail(id),
    queryFn: ({ signal }) =>
      api.get<PatientWorkoutSession>(`/api/v1/patient/me/workout-sessions/${id}`, { signal }),
    enabled: Boolean(id),
  });

export const patientMeCheckInsOptions = (limit = 30) =>
  queryOptions({
    queryKey: [...queryKeys.patientMe.checkIns(), { limit }] as const,
    queryFn: ({ signal }) =>
      api.get<HealthCheckIn[]>(`/api/v1/patient/me/check-ins?limit=${limit}`, { signal }),
    staleTime: 30_000,
  });

export const patientMeCheckInTodayOptions = () =>
  queryOptions({
    queryKey: queryKeys.patientMe.checkInToday(),
    queryFn: ({ signal }) =>
      api.get<{ checkIn: HealthCheckIn | null }>('/api/v1/patient/me/check-ins/today', {
        signal,
      }),
    staleTime: 60_000,
  });

export const patientMeNotificationPreferencesOptions = () =>
  queryOptions({
    queryKey: queryKeys.patientMe.notificationPreferences(),
    queryFn: ({ signal }) =>
      api.get<NotificationPreferences>('/api/v1/patient/me/notification-preferences', {
        signal,
      }),
    staleTime: 5 * 60_000,
  });

// =============================================================================
// Mutations
// =============================================================================

export const patientMeMutations = {
  confirmAppointment: (id: string) =>
    api.post<void>(`/api/v1/patient/me/appointments/${id}/confirm`),

  startWorkoutSession: (body: {
    planId: string;
    planSessionId: string;
    scheduledDate: string;
  }) => api.post<PatientWorkoutSession>('/api/v1/patient/me/workout-sessions', body),

  logSet: (
    sessionId: string,
    body: {
      planExerciseId: string;
      exerciseId: string;
      setNumber: number;
      reps?: number;
      weight?: number;
      durationSec?: number;
      rpe?: number;
      notes?: string;
    },
  ) =>
    api.post<PatientWorkoutLog>(
      `/api/v1/patient/me/workout-sessions/${sessionId}/logs`,
      body,
    ),

  completeWorkoutSession: (sessionId: string, body: { notes?: string } = {}) =>
    api.post<PatientWorkoutSession>(
      `/api/v1/patient/me/workout-sessions/${sessionId}/complete`,
      body,
    ),

  createCheckIn: (body: {
    energy: number;
    pain: number;
    painLocation?: string;
    mood: number;
    sleepHours: number;
    sleepQuality: number;
    stress: number;
    notes?: string;
  }) => api.post<HealthCheckIn>('/api/v1/patient/me/check-ins', body),

  patchNotificationPreferences: (body: Partial<Omit<NotificationPreferences, 'userId'>>) =>
    api.patch<NotificationPreferences>(
      '/api/v1/patient/me/notification-preferences',
      body,
    ),
};
