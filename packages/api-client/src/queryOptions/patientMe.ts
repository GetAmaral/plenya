import { queryOptions } from '@tanstack/react-query';
import { api } from '../fetcher';
import { queryKeys } from '../queryKeys';

// =============================================================================
// Types — espelham os DTOs do Go (PatientPortalService + PatientWorkoutsService)
// =============================================================================

// GET /patient/me retorna { userId, patient } — wrapped pelo handler
// (handler: patient_portal_handler.go:629). `patient` é o models.Patient
// inteiro serializado — listamos só os campos que as telas consomem.
export interface PatientMePatient {
  id: string;
  name: string;
  email?: string;
  phone?: string;
  address?: string;
  municipality?: string;
  state?: string;
  emergencyPhone?: string;
  birthDate?: string;
  gender?: 'male' | 'female' | 'other';
}

export interface PatientMeResponse {
  userId: string;
  patient: PatientMePatient;
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

export interface PatientLabBatchSummary {
  id: string;
  laboratoryName: string;
  collectionDate: string;
  resultDate?: string;
  status: string;
  requestingDoctor?: string;
  resultsCount: number;
}

export interface PatientLabValue {
  id: string;
  testName: string;
  testCode: string;
  value?: number;
  valueText?: string;
  unit?: string;
}

export interface PatientLabBatchDetail extends PatientLabBatchSummary {
  observations?: string;
  values: PatientLabValue[];
}

export interface PatientPrescriptionSummary {
  id: string;
  issuedAt: string;
  status: string;
  doctorName?: string;
  validUntil: string;
  medsCount: number;
  signedAt?: string;
  pdfPath?: string;
}

export interface PatientPhysicalAssessmentSummary {
  id: string;
  assessmentDate: string;
  assessorName?: string;
  weight?: number;
  height?: number;
  bmi?: number;
}

export interface PatientScoreGroupResult {
  groupId: string;
  groupName: string;
  scorePercentage: number;
}

export interface PatientScoreEntry {
  id: string;
  source: 'light' | 'complete';
  createdAt: string;
  totalScorePercentage: number;
  itemsEvaluatedCount: number;
  publicCode?: string;
  groupResults: PatientScoreGroupResult[];
}

// Items do continuum — espelham models.PatientContinuumItem (apenas campos
// que o handler retorna via Preload). Lista incompleta de propósito: TODO
// expandir conforme telas precisarem.
export interface PatientContinuumItem {
  id: string;
  title: string;
  description?: string;
  status: string;
  weekOffset?: number;
  expectedDate?: string;
  lateAfterDate?: string;
  appointmentId?: string;
  boxId?: string;
  position?: number;
  specialty?: string;
}

export interface PatientContinuumCoordinator {
  id?: string;
  name?: string;
}

export interface PatientContinuum {
  id: string;
  templateName: string;
  status: string;
  startDate: string;
  endDate?: string;
  integratedPlanMarkdown?: string;
  integratedPlanUpdatedAt?: string;
  items: PatientContinuumItem[];
  coordinatorDoctor?: PatientContinuumCoordinator;
}

export interface PatientMessage {
  id: string;
  createdAt: string;
  from: 'patient' | 'team';
  authorName?: string;
  channel: string;
  content: string;
}

export interface PatientMeProfileUpdate {
  phone?: string;
  email?: string;
  address?: string;
  municipality?: string;
  state?: string;
  emergencyPhone?: string;
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
    queryFn: async ({ signal }) => {
      // Handler retorna { userId, patient: { ... } } — desembrulha aqui pra
      // expor um shape achatado nas telas (mantem compat com PatientMeProfile).
      const res = await api.get<PatientMeResponse>('/api/v1/patient/me', { signal });
      return {
        userId: res.userId,
        id: res.patient?.id ?? '',
        name: res.patient?.name ?? '',
        email: res.patient?.email,
        phone: res.patient?.phone,
        address: res.patient?.address,
        municipality: res.patient?.municipality,
        state: res.patient?.state,
        emergencyPhone: res.patient?.emergencyPhone,
        birthDate: res.patient?.birthDate,
        gender: res.patient?.gender,
      } satisfies PatientMeProfile;
    },
    staleTime: 60_000,
  });

// PatientMeProfile = shape achatado consumido pelas telas.
export interface PatientMeProfile {
  userId: string;
  id: string;
  name: string;
  email?: string;
  phone?: string;
  address?: string;
  municipality?: string;
  state?: string;
  emergencyPhone?: string;
  birthDate?: string;
  gender?: 'male' | 'female' | 'other';
}

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

export const patientMeLabBatchesOptions = () =>
  queryOptions({
    queryKey: queryKeys.patientMe.labBatches(),
    queryFn: ({ signal }) =>
      api.get<PatientLabBatchSummary[]>('/api/v1/patient/me/lab-batches', { signal }),
    staleTime: 60_000,
  });

export const patientMeLabBatchOptions = (id: string) =>
  queryOptions({
    queryKey: [...queryKeys.patientMe.labBatches(), id] as const,
    queryFn: ({ signal }) =>
      api.get<PatientLabBatchDetail>(`/api/v1/patient/me/lab-batches/${id}`, { signal }),
    enabled: Boolean(id),
  });

export const patientMePrescriptionsOptions = () =>
  queryOptions({
    queryKey: [...queryKeys.patientMe.all(), 'prescriptions'] as const,
    queryFn: ({ signal }) =>
      api.get<PatientPrescriptionSummary[]>('/api/v1/patient/me/prescriptions', { signal }),
    staleTime: 60_000,
  });

export const patientMePhysicalAssessmentsOptions = () =>
  queryOptions({
    queryKey: [...queryKeys.patientMe.all(), 'physical-assessments'] as const,
    queryFn: ({ signal }) =>
      api.get<PatientPhysicalAssessmentSummary[]>(
        '/api/v1/patient/me/physical-assessments',
        { signal },
      ),
    staleTime: 60_000,
  });

export const patientMeScoresOptions = () =>
  queryOptions({
    queryKey: queryKeys.patientMe.scores(),
    queryFn: ({ signal }) =>
      api.get<PatientScoreEntry[]>('/api/v1/patient/me/scores', { signal }),
    staleTime: 60_000,
  });

export const patientMeContinuumOptions = () =>
  queryOptions({
    queryKey: [...queryKeys.patientMe.all(), 'continuum'] as const,
    queryFn: async ({ signal }) => {
      // Handler retorna { continuum: ... | null } — desembrulha.
      const res = await api.get<{ continuum: PatientContinuum | null }>(
        '/api/v1/patient/me/continuum',
        { signal },
      );
      return res.continuum;
    },
    staleTime: 60_000,
  });

export const patientMeMessagesOptions = () =>
  queryOptions({
    queryKey: queryKeys.patientMe.messages(),
    queryFn: ({ signal }) =>
      api.get<PatientMessage[]>('/api/v1/patient/me/messages', { signal }),
    staleTime: 15_000,
  });

export const patientMeMessagesUnreadOptions = () =>
  queryOptions({
    queryKey: [...queryKeys.patientMe.messages(), 'unread'] as const,
    queryFn: ({ signal }) =>
      api.get<{ unread: number }>('/api/v1/patient/me/messages/unread-count', {
        signal,
      }),
    refetchInterval: 30_000,
  });

export const patientMeAppointmentOptions = (id: string) =>
  queryOptions({
    queryKey: [...queryKeys.patientMe.appointments(), id] as const,
    queryFn: ({ signal }) =>
      api.get<PatientAppointment>(`/api/v1/patient/me/appointments/${id}`, { signal }),
    enabled: Boolean(id),
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
  requestCancelAppointment: (id: string, body: { reason?: string } = {}) =>
    api.post<void>(`/api/v1/patient/me/appointments/${id}/request-cancel`, body),
  requestRescheduleAppointment: (id: string, body: { reason?: string } = {}) =>
    api.post<void>(`/api/v1/patient/me/appointments/${id}/request-reschedule`, body),

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

  // Mensagens
  sendMessage: (body: { content: string }) =>
    api.post<PatientMessage>('/api/v1/patient/me/messages', body),
  markMessagesRead: () => api.post<void>('/api/v1/patient/me/messages/read'),

  // Perfil
  updateProfile: (body: PatientMeProfileUpdate) =>
    api.put<void>('/api/v1/patient/me/profile', body),
  // Handler aceita só { password } — sem currentPassword (paciente acabou de
  // entrar via magic-link/invite e está definindo senha).
  setPassword: (body: { password: string }) =>
    api.post<void>('/api/v1/patient/me/password', body),

  // LGPD
  lgpdRequestDelete: (body: { reason?: string } = {}) =>
    api.post<void>('/api/v1/patient/me/lgpd/account-delete-request', body),
};

// LGPD export retorna JSON arbitrário — usar fetch direto (resposta não-cacheada
// por design, é um download pontual). Endpoint: GET /api/v1/patient/me/lgpd/export.

