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
  | 'checked_in'
  | 'in_progress'
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
  checkedInAt?: string;
  startedAt?: string;
  completedAt?: string;
  cancelledAt?: string;
  cancellationReason?: string;
  externalCalendarEventId?: string;
  // HIGH H9 — DailyRoomURL não é mais exposto via /api/v1/appointments/:id.
  // Pra abrir a sala, chame POST /api/v1/appointments/:id/telemed-token via
  // useTelemedToken() — retorna {joinUrl} com meeting_token escopado.
  dailyRoomName?: string;
  // Consentimento de telemedicina (CFM 2.314/2022) — só type=telemedicine.
  telemedConsentAt?: string;
  telemedConsentMode?: string;
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
  continuumItemId?: string; // Continuum (Fase 3): ancora a marco do programa.
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
  /** Filtro singular legado. Prefira `doctorIds` para multi-médico. */
  doctorId?: string;
  /** Lista de médicos (calendário multi-doctor). Mandado como CSV ?doctorIds=. */
  doctorIds?: string[];
  patientId?: string;
  /** Janela [from, to) no formato ISO — agora filtrada server-side. */
  dateFrom?: string;
  dateTo?: string;
  limit?: number;
  offset?: number;
}

function buildAppointmentQuery(filter: AppointmentFilter): string {
  const params = new URLSearchParams();
  if (filter.status) params.set('status', filter.status);
  if (filter.doctorIds && filter.doctorIds.length > 0) {
    params.set('doctorIds', filter.doctorIds.join(','));
  } else if (filter.doctorId) {
    params.set('doctorId', filter.doctorId);
  }
  if (filter.patientId) params.set('patientId', filter.patientId);
  if (filter.dateFrom) params.set('dateFrom', filter.dateFrom);
  if (filter.dateTo) params.set('dateTo', filter.dateTo);
  params.set('limit', String(filter.limit ?? 200));
  params.set('offset', String(filter.offset ?? 0));
  return params.toString();
}

/**
 * Lista appointments. Backend pode retornar array direto OU {data, total}.
 * Normalizamos pra sempre devolver Appointment[].
 */
export function useAppointments(filter: AppointmentFilter = {}) {
  // Quando o caller pediu multi-doctor mas a lista veio vazia, evitamos chamar
  // o backend (que faria SELECT em todos os médicos como se fosse "todos").
  const enabled =
    filter.doctorIds === undefined || filter.doctorIds.length > 0 || !!filter.patientId;

  return useQuery({
    queryKey: calendarKeys.appointments(filter),
    queryFn: async () => {
      const result = await apiClient.get<Appointment[] | AppointmentListResponse>(
        `/api/v1/appointments?${buildAppointmentQuery(filter)}`,
      );
      return Array.isArray(result) ? result : result.data;
    },
    enabled,
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

/**
 * Reagendamento otimista — usado pelo drag-and-drop do calendário.
 *
 * Diferente de useUpdateAppointment(id), que fixa o id na criação do hook,
 * este recebe {id, scheduledAt} por chamada — necessário no grid, onde a
 * consulta arrastada só é conhecida no drop. Aplica update otimista em TODAS
 * as listas de appointments em cache (o bloco "salta" na hora) e faz rollback
 * se o backend recusar (ex.: 409 de conflito de horário do médico).
 */
export function useRescheduleAppointment() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: ({ id, scheduledAt }: { id: string; scheduledAt: string }) =>
      apiClient.put<Appointment>(`/api/v1/appointments/${id}`, { scheduledAt }),
    onMutate: async ({ id, scheduledAt }) => {
      const listKey = [...calendarKeys.all, 'appointments'];
      // Cancela refetches em voo (polling 15s) pra não sobrescrever o otimista.
      await qc.cancelQueries({ queryKey: listKey });
      const previous = qc.getQueriesData<Appointment[]>({ queryKey: listKey });
      qc.setQueriesData<Appointment[]>({ queryKey: listKey }, (old) =>
        old ? old.map((a) => (a.id === id ? { ...a, scheduledAt } : a)) : old,
      );
      return { previous };
    },
    onError: (_err, _vars, context) => {
      // Rollback: restaura todas as listas tocadas pro estado pré-otimista.
      context?.previous?.forEach(([key, data]) => {
        qc.setQueryData(key, data);
      });
    },
    onSettled: () => {
      // Reconcilia com a verdade do servidor (sucesso ou erro).
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

/**
 * HIGH H9 — gera meeting_token de owner pra abrir a sala Daily.co.
 *
 * Sala é privacy=private no Daily — URL crua não funciona. Cada vez que
 * o médico/staff abre a sala, gera um token novo (sem caching) escopado a
 * ele (is_owner=true, screenshare=true, exp=closesAt).
 *
 * Uso (mutation, não query — token é fresh a cada call):
 *   const tokenMutation = useTelemedToken(appointmentId);
 *   const { joinUrl } = await tokenMutation.mutateAsync();
 */
export interface TelemedTokenResponse {
  joinUrl: string;
}
export function useTelemedToken(id: string) {
  return useMutation({
    mutationFn: () =>
      apiClient.post<TelemedTokenResponse>(`/api/v1/appointments/${id}/telemed-token`),
  });
}

// Termo de consentimento de telemedicina (texto canônico, do backend).
export interface TelemedConsentTerm {
  text: string;
}
export function useTelemedConsentTerm() {
  return useQuery({
    queryKey: ['telemed-consent-term'],
    queryFn: () => apiClient.get<TelemedConsentTerm>('/api/v1/appointments/telemed-consent-term'),
    staleTime: 60 * 60 * 1000, // termo muda raramente
  });
}

// Registra o consentimento de telemedicina (CFM 2.314/2022). Mode default 'verbal'.
export function useRegisterTelemedConsent(id: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (mode: 'verbal' | 'written' = 'verbal') =>
      apiClient.post<Appointment>(`/api/v1/appointments/${id}/telemed-consent`, { mode }),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: calendarKeys.appointment(id) });
      qc.invalidateQueries({ queryKey: calendarKeys.all });
    },
  });
}

// Gravação + transcrição da teleconsulta (Daily.co cloud recording + Deepgram nova-3).
// Os artefatos chegam de forma assíncrona via webhook do Daily, então fazemos poll
// enquanto algo ainda está em processamento. 404 = ainda não há gravação (null).
export interface TelemedRecording {
  id: string;
  appointmentId?: string;
  patientId?: string;
  recordingStatus: 'pending' | 'started' | 'finished' | 'error';
  hasRecording: boolean;
  recordingReadyAt?: string;
  recordingDurationSeconds?: number;
  recordingError?: string;
  transcriptStatus: 'none' | 'in_progress' | 'finished' | 'failed';
  hasTranscript: boolean;
  transcriptReadyAt?: string;
  transcriptText?: string;
  transcriptError?: string;
  // Nota clínica gerada por IA (AI scribe) — rascunho revisável, não assinável.
  generatedNoteStatus: 'none' | 'generating' | 'done' | 'failed';
  generatedNoteFormat?: string;
  generatedNoteModel?: string;
  generatedNoteAt?: string;
  generatedNoteError?: string;
  generatedNote?: TelemedGeneratedNote;
  updatedAt: string;
}

export type TelemedNoteFormat = 'anamnese' | 'soap';
export type SoapTarget = 'subjective' | 'objective' | 'assessment' | 'plan';
export interface TelemedGeneratedNoteSection {
  chave: string;
  titulo: string;
  texto: string;
  soapTarget: SoapTarget;
}
export interface TelemedGeneratedNote {
  format: string;
  sections: TelemedGeneratedNoteSection[];
  itensAmbiguos?: string[];
  papeis?: Record<string, string>;
}

export function useTelemedRecording(
  id: string | undefined,
  opts?: { enabled?: boolean; expectRecording?: boolean },
) {
  const expect = opts?.expectRecording ?? false;
  return useQuery({
    queryKey: ['telemed-recording', id ?? ''],
    enabled: (opts?.enabled ?? true) && !!id,
    queryFn: async (): Promise<TelemedRecording | null> => {
      try {
        return await apiClient.get<TelemedRecording>(
          `/api/v1/appointments/${id}/telemed-recording`,
        );
      } catch (err) {
        if ((err as { status?: number }).status === 404) return null;
        throw err;
      }
    },
    refetchInterval: (query) => {
      const d = query.state.data as TelemedRecording | null | undefined;
      // Sem linha ainda: faz poll só se esperamos gravação (consentimento dado).
      if (!d) return expect ? 15000 : false;
      const pendingRec = d.recordingStatus === 'pending' || d.recordingStatus === 'started';
      const pendingTr = d.transcriptStatus === 'in_progress';
      return pendingRec || pendingTr ? 15000 : false;
    },
  });
}

// Link assinado de download do MP4, gerado sob demanda (a gravação fica no Daily).
export function useTelemedRecordingDownload(id: string) {
  return useMutation({
    mutationFn: () =>
      apiClient.get<{ downloadUrl: string }>(
        `/api/v1/appointments/${id}/telemed-recording/download`,
      ),
  });
}

// Gera (via IA) a nota clínica estruturada (anamnese|soap) a partir do transcript.
// Síncrono (~5-15s). Retorna o recording com generatedNote; invalida o cache.
export function useGenerateTelemedNote(id: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (format: TelemedNoteFormat = 'anamnese') =>
      apiClient.post<TelemedRecording>(
        `/api/v1/appointments/${id}/telemed-recording/generate-note`,
        { format },
      ),
    onSuccess: (data) => {
      qc.setQueryData(['telemed-recording', id], data);
      qc.invalidateQueries({ queryKey: ['telemed-recording', id] });
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

// Recepção registra a chegada do paciente (status -> checked_in).
export function useCheckInAppointment(id: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: () =>
      apiClient.post<Appointment>(`/api/v1/appointments/${id}/check-in`),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: calendarKeys.appointment(id) });
      qc.invalidateQueries({ queryKey: calendarKeys.all });
    },
  });
}

// Início do atendimento (status -> in_progress).
export function useStartAppointment(id: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: () =>
      apiClient.post<Appointment>(`/api/v1/appointments/${id}/start`),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: calendarKeys.appointment(id) });
      qc.invalidateQueries({ queryKey: calendarKeys.all });
    },
  });
}

// Tempo de espera em minutos (StartedAt - CheckedInAt), ou null se indisponível.
export function waitMinutes(appt: Appointment): number | null {
  if (!appt.checkedInAt) return null;
  const end = appt.startedAt ? new Date(appt.startedAt) : new Date();
  const start = new Date(appt.checkedInAt);
  const diff = Math.round((end.getTime() - start.getTime()) / 60000);
  return diff >= 0 ? diff : null;
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
  checked_in: 'Aguardando',
  in_progress: 'Em atendimento',
  completed: 'Concluída',
  cancelled: 'Cancelada',
  no_show: 'Não compareceu',
};

export const APPOINTMENT_STATUS_COLORS: Record<AppointmentStatus, string> = {
  scheduled: 'bg-blue-100 text-blue-900 border-blue-200',
  confirmed: 'bg-emerald-100 text-emerald-900 border-emerald-200',
  checked_in: 'bg-amber-100 text-amber-900 border-amber-200',
  in_progress: 'bg-violet-100 text-violet-900 border-violet-200',
  completed: 'bg-stone-100 text-stone-900 border-stone-200',
  cancelled: 'bg-rose-100 text-rose-900 border-rose-200',
  no_show: 'bg-orange-100 text-orange-900 border-orange-200',
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
