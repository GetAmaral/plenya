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
  /** CSV de UUIDs no backend; passamos array aqui pra ergonomia */
  doctorIds?: string[];
  /** RFC3339 — filtra appointments com scheduledAt >= dateFrom */
  dateFrom?: string;
  /** RFC3339 — filtra appointments com scheduledAt <= dateTo */
  dateTo?: string;
  status?: AppointmentStatus;
  limit?: number;
  offset?: number;
}

export const appointmentsListOptions = (params: AppointmentListParams = {}) =>
  queryOptions({
    queryKey: [...queryKeys.appointments.all(), 'list', params] as const,
    queryFn: ({ signal }) => {
      const qs = new URLSearchParams();
      if (params.patientId) qs.set('patientId', params.patientId);
      if (params.doctorIds && params.doctorIds.length > 0) {
        qs.set('doctorIds', params.doctorIds.join(','));
      } else if (params.doctorId) {
        qs.set('doctorId', params.doctorId);
      }
      if (params.dateFrom) qs.set('dateFrom', params.dateFrom);
      if (params.dateTo) qs.set('dateTo', params.dateTo);
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

export interface CreateAppointmentInput {
  patientId: string;
  doctorId: string;
  /** RFC3339 — ex: 2026-05-10T14:00:00-03:00 */
  scheduledAt: string;
  durationMinutes: number;
  type: AppointmentType;
  reason: string;
  patientNotes?: string;
}

export interface UpdateAppointmentInput {
  scheduledAt?: string;
  status?: AppointmentStatus;
  doctorNotes?: string;
  diagnosis?: string;
}

export const appointmentMutations = {
  create: (body: CreateAppointmentInput) =>
    api.post<Appointment>('/api/v1/appointments', body),
  update: (id: string, body: UpdateAppointmentInput) =>
    api.put<Appointment>(`/api/v1/appointments/${id}`, body),
  confirm: (id: string) => api.post<void>(`/api/v1/appointments/${id}/confirm`),
  cancel: (id: string, body: { reason?: string } = {}) =>
    api.post<void>(`/api/v1/appointments/${id}/cancel`, body),
};

/** @deprecated use appointmentsListOptions + filtro client-side. Mantido para compat. */
export const appointmentsByRangeOptions = (_from: string, _to: string) =>
  appointmentsListOptions({ limit: 200 });
