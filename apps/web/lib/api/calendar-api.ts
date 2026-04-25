/**
 * Calendar V1 API hooks (Bloco G frontend).
 *
 * Cobre:
 *  - Google OAuth integration (status, auth-url, disconnect)
 *  - Slot picker (calendar/slots)
 *  - Appointments CRUD + confirm/cancel
 *  - Working hours CRUD
 *  - Doctor absences CRUD
 *  - Doctors list (pra dropdowns)
 *
 * Padrão TanStack Query: queryKeys hierárquicos, invalidate em mutations,
 * polling 15-30s pra dados real-time, refetchOnWindowFocus default.
 *
 * Timezone: backend trabalha em UTC. UI mostra em America/Sao_Paulo.
 * Sempre envia ISO UTC pro backend (toISOString()).
 */
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { apiClient } from '../api-client';

// =====================================================
// Types
// =====================================================

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

export interface Appointment {
  id: string;
  patientId: string;
  doctorId: string;
  scheduledAt: string; // ISO UTC
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
  dailyRoomUrl?: string;
  dailyRoomName?: string;
  confirmationSentAt?: string;
  reminderSentAt?: string;
  displayTitle?: string;
  createdAt: string;
  updatedAt: string;
  // Possíveis relacionamentos preenchidos pelo backend (preload)
  patient?: { id: string; name: string; email?: string; phone?: string };
  doctor?: { id: string; name: string; email?: string };
}

export interface AppointmentListResponse {
  data: Appointment[];
  total: number;
}

export interface CreateAppointmentPayload {
  patientId: string;
  doctorId: string;
  scheduledAt: string; // ISO UTC RFC3339
  durationMinutes: number;
  type: AppointmentType;
  reason: string;
  patientNotes?: string;
}

export interface UpdateAppointmentPayload {
  scheduledAt?: string;
  status?: AppointmentStatus;
  doctorNotes?: string;
  diagnosis?: string;
}

// Google OAuth
export interface GoogleStatus {
  connected: boolean;
  configured: boolean;
  googleEmail?: string;
  lastSyncAt?: string;
  revokedAt?: string;
  calendarId?: string;
}

export interface GoogleAuthURL {
  url: string;
  configured?: boolean;
}

// Slot picker
export interface CalendarSlot {
  startUtc: string; // ISO UTC
  endUtc: string;
}

export interface CalendarSlotsResponse {
  doctorId: string;
  date: string;
  type: AppointmentType;
  durationMin: number;
  slots: CalendarSlot[];
}

// Working hours
export interface WorkingHours {
  id: string;
  doctorId: string;
  weekday: number; // 0=Dom, 1=Seg, ...
  startMinute: number;
  endMinute: number;
  slotDuration: number;
  createdAt: string;
  updatedAt: string;
}

export interface CreateWorkingHoursPayload {
  weekday: number;
  startMinute: number;
  endMinute: number;
  slotDuration: number;
}

// Doctor absences
export interface DoctorAbsence {
  id: string;
  doctorId: string;
  startAt: string; // ISO UTC
  endAt: string;
  reason: string;
  createdAt: string;
}

export interface CreateAbsencePayload {
  startAt: string;
  endAt: string;
  reason: string;
}

// Doctors list
export interface DoctorOption {
  id: string;
  name: string;
  email: string;
  specialty?: string;
}

// =====================================================
// Query keys
// =====================================================

export const calendarKeys = {
  all: ['calendar'] as const,
  googleStatus: () => [...calendarKeys.all, 'google', 'status'] as const,
  googleAuthURL: () => [...calendarKeys.all, 'google', 'auth-url'] as const,
  slots: (doctorId: string, date: string, type: AppointmentType, override?: number) =>
    [...calendarKeys.all, 'slots', doctorId, date, type, override ?? null] as const,
  appointments: (filter: AppointmentFilter) =>
    [...calendarKeys.all, 'appointments', filter] as const,
  appointment: (id: string) => [...calendarKeys.all, 'appointment', id] as const,
  workingHours: (doctorId: string) =>
    [...calendarKeys.all, 'working-hours', doctorId] as const,
  absences: (doctorId: string) => [...calendarKeys.all, 'absences', doctorId] as const,
  doctors: () => [...calendarKeys.all, 'doctors'] as const,
};

// =====================================================
// Google OAuth
// =====================================================

export function useGoogleStatus() {
  return useQuery({
    queryKey: calendarKeys.googleStatus(),
    queryFn: () => apiClient.get<GoogleStatus>('/api/v1/integrations/google/status'),
    staleTime: 30 * 1000,
    refetchInterval: 30 * 1000,
  });
}

/**
 * Resolve a auth-url on demand (não fica fazendo polling — é one-shot).
 * Use a mutation pattern pra carregar quando usuário clica "Conectar".
 */
export function useFetchGoogleAuthURL() {
  return useMutation({
    mutationFn: () => apiClient.get<GoogleAuthURL>('/api/v1/integrations/google/auth-url'),
  });
}

/** Alias compatível: mesmo hook, nome mais curto. Usar via .mutateAsync(). */
export const useGoogleAuthURL = useFetchGoogleAuthURL;

export function useGoogleDisconnect() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: () =>
      apiClient.post<{ disconnected: boolean }>('/api/v1/integrations/google/disconnect'),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: calendarKeys.googleStatus() });
    },
  });
}

// =====================================================
// Slot picker
// =====================================================

export function useCalendarSlots(
  doctorId: string | undefined,
  date: string | undefined, // YYYY-MM-DD
  type: AppointmentType,
  overrideDuration?: number,
) {
  return useQuery({
    queryKey: calendarKeys.slots(doctorId ?? '', date ?? '', type, overrideDuration),
    queryFn: () => {
      const params = new URLSearchParams({
        doctorId: doctorId ?? '',
        date: date ?? '',
        type,
      });
      if (overrideDuration) params.set('durationMin', String(overrideDuration));
      return apiClient.get<CalendarSlotsResponse>(`/api/v1/calendar/slots?${params.toString()}`);
    },
    enabled: !!doctorId && !!date,
    staleTime: 30 * 1000,
  });
}

// =====================================================
// Appointments
// =====================================================

export interface AppointmentFilter {
  status?: AppointmentStatus;
  doctorId?: string;
  patientId?: string;
  /** Janela [from, to) no formato ISO. Filtragem client-side por enquanto. */
  dateFrom?: string;
  dateTo?: string;
  limit?: number;
  offset?: number;
}

function buildAppointmentQuery(filter: AppointmentFilter): string {
  const params = new URLSearchParams();
  if (filter.status) params.set('status', filter.status);
  if (filter.doctorId) params.set('doctorId', filter.doctorId);
  if (filter.patientId) params.set('patientId', filter.patientId);
  params.set('limit', String(filter.limit ?? 100));
  params.set('offset', String(filter.offset ?? 0));
  return params.toString();
}

/**
 * Lista appointments. Backend pode retornar array direto OU {data, total}.
 * Normalizamos pra sempre devolver Appointment[].
 */
export function useAppointments(filter: AppointmentFilter = {}) {
  return useQuery({
    queryKey: calendarKeys.appointments(filter),
    queryFn: async () => {
      const result = await apiClient.get<Appointment[] | AppointmentListResponse>(
        `/api/v1/appointments?${buildAppointmentQuery(filter)}`,
      );
      const list = Array.isArray(result) ? result : result.data;
      // Filtragem client-side por janela de data (backend ainda não suporta)
      return list.filter((a) => {
        const at = new Date(a.scheduledAt).getTime();
        if (filter.dateFrom && at < new Date(filter.dateFrom).getTime()) return false;
        if (filter.dateTo && at >= new Date(filter.dateTo).getTime()) return false;
        return true;
      });
    },
    staleTime: 15 * 1000,
    refetchInterval: 15 * 1000,
    refetchOnWindowFocus: true,
  });
}

export function useAppointment(id: string | undefined) {
  return useQuery({
    queryKey: calendarKeys.appointment(id ?? ''),
    queryFn: () => apiClient.get<Appointment>(`/api/v1/appointments/${id}`),
    enabled: !!id,
    staleTime: 15 * 1000,
    refetchInterval: 15 * 1000,
    refetchOnWindowFocus: true,
  });
}

export function useCreateAppointment() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (payload: CreateAppointmentPayload) =>
      apiClient.post<Appointment>('/api/v1/appointments', payload),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: calendarKeys.all });
    },
  });
}

export function useUpdateAppointment(id: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (payload: UpdateAppointmentPayload) =>
      apiClient.put<Appointment>(`/api/v1/appointments/${id}`, payload),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: calendarKeys.appointment(id) });
      qc.invalidateQueries({ queryKey: calendarKeys.all });
    },
  });
}

export function useCancelAppointment(id: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (payload: { reason: string }) =>
      apiClient.post<Appointment>(`/api/v1/appointments/${id}/cancel`, payload),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: calendarKeys.appointment(id) });
      qc.invalidateQueries({ queryKey: calendarKeys.all });
    },
  });
}

export function useConfirmAppointment(id: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: () => apiClient.post<void>(`/api/v1/appointments/${id}/confirm`),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: calendarKeys.appointment(id) });
      qc.invalidateQueries({ queryKey: calendarKeys.all });
    },
  });
}

// =====================================================
// Working hours
// =====================================================

export function useWorkingHours(doctorId: string | undefined) {
  return useQuery({
    queryKey: calendarKeys.workingHours(doctorId ?? ''),
    queryFn: () =>
      apiClient.get<WorkingHours[]>(`/api/v1/doctors/${doctorId}/working-hours`),
    enabled: !!doctorId,
    staleTime: 60 * 1000,
  });
}

export function useCreateWorkingHours(doctorId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (payload: CreateWorkingHoursPayload) =>
      apiClient.post<WorkingHours>(`/api/v1/doctors/${doctorId}/working-hours`, payload),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: calendarKeys.workingHours(doctorId) });
    },
  });
}

export function useUpdateWorkingHours(doctorId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (input: { id: string; payload: CreateWorkingHoursPayload }) =>
      apiClient.put<WorkingHours>(
        `/api/v1/doctors/${doctorId}/working-hours/${input.id}`,
        input.payload,
      ),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: calendarKeys.workingHours(doctorId) });
    },
  });
}

export function useDeleteWorkingHours(doctorId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: string) =>
      apiClient.delete<void>(`/api/v1/doctors/${doctorId}/working-hours/${id}`),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: calendarKeys.workingHours(doctorId) });
    },
  });
}

// =====================================================
// Doctor absences
// =====================================================

export function useDoctorAbsences(doctorId: string | undefined) {
  return useQuery({
    queryKey: calendarKeys.absences(doctorId ?? ''),
    queryFn: () =>
      apiClient.get<DoctorAbsence[]>(`/api/v1/doctors/${doctorId}/absences`),
    enabled: !!doctorId,
    staleTime: 60 * 1000,
  });
}

export function useCreateAbsence(doctorId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (payload: CreateAbsencePayload) =>
      apiClient.post<DoctorAbsence>(`/api/v1/doctors/${doctorId}/absences`, payload),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: calendarKeys.absences(doctorId) });
    },
  });
}

export function useDeleteAbsence(doctorId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: string) =>
      apiClient.delete<void>(`/api/v1/doctors/${doctorId}/absences/${id}`),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: calendarKeys.absences(doctorId) });
    },
  });
}

// =====================================================
// Doctors list (dropdown)
// =====================================================

export function useDoctors() {
  return useQuery({
    queryKey: calendarKeys.doctors(),
    queryFn: () => apiClient.get<DoctorOption[]>('/api/v1/users/doctors'),
    staleTime: 5 * 60 * 1000,
  });
}

// =====================================================
// Visual helpers (labels, colors, etc)
// =====================================================

export const APPOINTMENT_TYPE_LABELS: Record<AppointmentType, string> = {
  initial_assessment: 'Avaliação Inicial',
  follow_up: 'Retorno',
  telemedicine: 'Teleconsulta',
  procedure: 'Procedimento',
  results_review: 'Revisão de Exames',
};

export const APPOINTMENT_TYPE_DEFAULT_DURATION: Record<AppointmentType, number> = {
  initial_assessment: 90,
  follow_up: 30,
  telemedicine: 30,
  procedure: 60,
  results_review: 45,
};

/** Tailwind classes pra background do bloco no calendário. */
export const APPOINTMENT_TYPE_COLORS: Record<AppointmentType, string> = {
  initial_assessment: 'bg-blue-100 border-blue-300 text-blue-900 hover:bg-blue-200',
  follow_up: 'bg-emerald-100 border-emerald-300 text-emerald-900 hover:bg-emerald-200',
  telemedicine: 'bg-purple-100 border-purple-300 text-purple-900 hover:bg-purple-200',
  procedure: 'bg-orange-100 border-orange-300 text-orange-900 hover:bg-orange-200',
  results_review: 'bg-amber-100 border-amber-300 text-amber-900 hover:bg-amber-200',
};

export const APPOINTMENT_STATUS_LABELS: Record<AppointmentStatus, string> = {
  scheduled: 'Agendada',
  confirmed: 'Confirmada',
  completed: 'Concluída',
  cancelled: 'Cancelada',
  no_show: 'Não compareceu',
};

export const APPOINTMENT_STATUS_COLORS: Record<AppointmentStatus, string> = {
  scheduled: 'bg-blue-100 text-blue-900 border-blue-200',
  confirmed: 'bg-emerald-100 text-emerald-900 border-emerald-200',
  completed: 'bg-stone-100 text-stone-900 border-stone-200',
  cancelled: 'bg-rose-100 text-rose-900 border-rose-200',
  no_show: 'bg-amber-100 text-amber-900 border-amber-200',
};

export const WEEKDAY_LABELS = [
  'Domingo',
  'Segunda',
  'Terça',
  'Quarta',
  'Quinta',
  'Sexta',
  'Sábado',
];

export const WEEKDAY_LABELS_SHORT = ['Dom', 'Seg', 'Ter', 'Qua', 'Qui', 'Sex', 'Sáb'];

/** "08:00" -> 480 */
export function timeStringToMinutes(t: string): number {
  const [h, m] = t.split(':').map(Number);
  return (h ?? 0) * 60 + (m ?? 0);
}

/** 480 -> "08:00" */
export function minutesToTimeString(m: number): string {
  const h = Math.floor(m / 60);
  const min = m % 60;
  return `${String(h).padStart(2, '0')}:${String(min).padStart(2, '0')}`;
}
