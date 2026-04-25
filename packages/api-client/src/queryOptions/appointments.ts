import { queryOptions } from '@tanstack/react-query';
import { api } from '../fetcher';
import { queryKeys } from '../queryKeys';

export type AppointmentStatus =
  | 'scheduled'
  | 'confirmed'
  | 'completed'
  | 'cancelled'
  | 'no_show';

export type AppointmentType =
  | 'initial_assessment'
  | 'follow_up'
  | 'telemedicine'
  | 'procedure'
  | 'results_review';

export const appointmentTypeLabels: Record<AppointmentType, string> = {
  initial_assessment: 'Avaliação inicial',
  follow_up: 'Retorno',
  telemedicine: 'Teleconsulta',
  procedure: 'Procedimento',
  results_review: 'Revisão de exames',
};

export const appointmentTypeDefaultDuration: Record<AppointmentType, number> = {
  initial_assessment: 90,
  follow_up: 30,
  telemedicine: 30,
  procedure: 60,
  results_review: 45,
};

/**
 * Reflete o model Go `Appointment` (Calendar V1).
 * Campos opcionais correspondem a *time.Time / *string nos models.
 */
export interface Appointment {
  id: string;
  patientId: string;
  patientName?: string;
  doctorId: string;
  doctorName?: string;
  scheduledAt: string;
  durationMinutes: number;
  type: AppointmentType;
  status: AppointmentStatus;
  reason: string;
  patientNotes?: string;
  doctorNotes?: string;
  diagnosis?: string;
  anamnesisId?: string;
  confirmedAt?: string;
  completedAt?: string;
  cancelledAt?: string;
  cancellationReason?: string;
  externalCalendarEventId?: string;
  /** Sala Daily.co — somente preenchido em type=telemedicine */
  dailyRoomUrl?: string;
  dailyRoomName?: string;
  reminderSentAt?: string;
  confirmationSentAt?: string;
  displayTitle?: string;
  createdAt: string;
  updatedAt: string;
}

export interface AppointmentListParams {
  patientId?: string;
  doctorId?: string;
  status?: AppointmentStatus;
  limit?: number;
  offset?: number;
}

/**
 * Lista appointments. Backend não filtra por range — para Hoje/7-dias o
 * caller filtra client-side por scheduledAt usando o limit/offset alto.
 */
export const appointmentsListOptions = (params: AppointmentListParams = {}) =>
  queryOptions({
    queryKey: [...queryKeys.appointments.all(), 'list', params] as const,
    queryFn: ({ signal }) => {
      const qs = new URLSearchParams();
      if (params.patientId) qs.set('patientId', params.patientId);
      if (params.doctorId) qs.set('doctorId', params.doctorId);
      if (params.status) qs.set('status', params.status);
      qs.set('limit', String(params.limit ?? 100));
      if (params.offset) qs.set('offset', String(params.offset));
      return api.get<Appointment[]>(`/api/v1/appointments?${qs.toString()}`, { signal });
    },
  });

export const appointmentDetailOptions = (id: string) =>
  queryOptions({
    queryKey: [...queryKeys.appointments.all(), 'detail', id] as const,
    queryFn: ({ signal }) => api.get<Appointment>(`/api/v1/appointments/${id}`, { signal }),
    enabled: Boolean(id),
  });

export const appointmentMutations = {
  confirm: (id: string) => api.post<void>(`/api/v1/appointments/${id}/confirm`),
  cancel: (id: string, body: { reason?: string } = {}) =>
    api.post<void>(`/api/v1/appointments/${id}/cancel`, body),
};

/** @deprecated use appointmentsListOptions + filtro client-side. Mantido para compat. */
export const appointmentsByRangeOptions = (_from: string, _to: string) =>
  appointmentsListOptions({ limit: 200 });
