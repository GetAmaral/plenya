/**
 * API client da área do paciente (meu.plenyasaude.com.br).
 *
 * Auth público (login/magic-link/invite-consume) usa fetch direto pra evitar
 * o pipeline de refresh-token do apiClient (que não faz sentido pré-login).
 *
 * Endpoints autenticados (/patient/me/*) usam apiClient normal — o JWT do
 * paciente fica no mesmo localStorage do staff (o backend valida role).
 */

import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { apiClient } from "@/lib/api-client";
import type { User, Patient } from "@/lib/auth-store";

const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:3001";

export interface PatientAuthResponse {
  accessToken: string;
  refreshToken: string;
  user: User;
}

async function publicPost<T>(path: string, body: unknown): Promise<T> {
  const res = await fetch(`${API_URL}${path}`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(body),
  });
  if (!res.ok) {
    let message = `${res.status}`;
    try {
      const data = await res.json();
      if (data?.error) message = data.error;
    } catch {}
    throw new Error(message);
  }
  if (res.status === 204) return undefined as T;
  return res.json();
}

// ============================================================
// Auth público
// ============================================================

export const patientAuthApi = {
  loginPassword: (email: string, password: string) =>
    publicPost<PatientAuthResponse>("/api/v1/auth/patient/login", {
      email,
      password,
    }),
  requestMagicLink: (email: string) =>
    publicPost<void>("/api/v1/auth/patient/magic-link", { email }),
  consumeMagicLink: (token: string) =>
    publicPost<PatientAuthResponse>(
      "/api/v1/auth/patient/magic-link/consume",
      { token },
    ),
  consumeInvite: (token: string) =>
    publicPost<PatientAuthResponse & { patientId: string }>(
      "/api/v1/auth/patient/invite/consume",
      { token },
    ),
  forgotPassword: (email: string) =>
    publicPost<void>("/api/v1/auth/patient/forgot", { email }),
};

// ============================================================
// /patient/me (autenticado)
// ============================================================

export interface PatientMe {
  userId: string;
  patient: Patient;
}

export const patientMeApi = {
  me: () => apiClient.get<PatientMe>("/api/v1/patient/me"),
  setPassword: (password: string) =>
    apiClient.post<void>("/api/v1/patient/me/password", { password }),
  dashboard: () =>
    apiClient.get<PatientDashboard>("/api/v1/patient/me/dashboard"),
};

export function usePatientMe(enabled = true) {
  return useQuery({
    queryKey: ["patient-me"],
    queryFn: patientMeApi.me,
    enabled,
    staleTime: 60_000,
  });
}

// ============================================================
// Dashboard payload
// ============================================================

export interface PatientDashboard {
  nextAppointment?: {
    id: string;
    scheduledAt: string;
    durationMinutes: number;
    type: string;
    status: string;
    doctorId: string;
    doctorName: string;
    isTelemedicine: boolean;
    patientConfirmedAt?: string | null;
    minutesUntilStart: number;
  };
  continuum?: {
    id: string;
    templateName: string;
    startDate: string;
    endDate: string;
    currentWeek: number;
    totalWeeks: number;
    itemsCompleted: number;
    itemsTotal: number;
    itemsLate: number;
    nextItemTitle?: string | null;
    nextItemExpectedDate?: string | null;
  };
  lastScore?: {
    id: string;
    source: "light" | "complete";
    createdAt: string;
    totalScorePercentage: number;
    deltaVsPrevious?: number | null;
  };
  boxesCount: { preparing: number; shipped: number; delivered: number };
  unreadMessages: number;
  pendingActions: { kind: string; description: string; link: string }[];
}

export function usePatientDashboard(enabled = true) {
  return useQuery({
    queryKey: ["patient-dashboard"],
    queryFn: patientMeApi.dashboard,
    enabled,
    staleTime: 30_000,
  });
}

// ============================================================
// Meu Continuum (read-only)
// ============================================================

export type PatientContinuumItemStatus =
  | "pending"
  | "scheduled"
  | "completed"
  | "missed"
  | "cancelled"
  | "skipped";

export interface PatientContinuumItemView {
  id: string;
  type: "appointment" | "box" | "reassessment" | "milestone" | "custom";
  specialty?: string | null;
  title: string;
  description?: string;
  weekOffset: number;
  expectedDate: string;
  lateAfterDate: string;
  status: PatientContinuumItemStatus;
  appointmentId?: string | null;
  boxId?: string | null;
  completedAt?: string | null;
  position: number;
}

export interface MyContinuumView {
  id: string;
  templateName: string;
  status: "active" | "paused" | "completed" | "cancelled";
  startDate: string;
  endDate: string;
  integratedPlanMarkdown: string;
  integratedPlanUpdatedAt?: string | null;
  items: PatientContinuumItemView[];
  coordinatorDoctor?: { id: string; name: string } | null;
}

export const patientContinuumApi = {
  myContinuum: () =>
    apiClient.get<{ continuum: MyContinuumView | null }>(
      "/api/v1/patient/me/continuum",
    ),
};

export function useMyContinuum(enabled = true) {
  return useQuery({
    queryKey: ["patient-my-continuum"],
    queryFn: patientContinuumApi.myContinuum,
    enabled,
    staleTime: 60_000,
  });
}

// ============================================================
// Minhas consultas
// ============================================================

export interface MyAppointment {
  id: string;
  scheduledAt: string;
  durationMinutes: number;
  type: "initial_assessment" | "follow_up" | "telemedicine" | "procedure" | "results_review";
  status: "scheduled" | "confirmed" | "completed" | "cancelled" | "no_show";
  doctorId: string;
  doctorName: string;
  isTelemedicine: boolean;
  dailyRoomUrl?: string | null;
  notes?: string;
  patientConfirmedAt?: string | null;
  confirmedAt?: string | null;
  cancelledAt?: string | null;
  minutesUntilStart: number;
}

export const patientAppointmentApi = {
  list: (range: "upcoming" | "past" = "upcoming") =>
    apiClient.get<MyAppointment[]>(`/api/v1/patient/me/appointments?range=${range}`),
  get: (id: string) => apiClient.get<MyAppointment>(`/api/v1/patient/me/appointments/${id}`),
  confirm: (id: string) =>
    apiClient.post<void>(`/api/v1/patient/me/appointments/${id}/confirm`, {}),
  requestCancel: (id: string, reason: string) =>
    apiClient.post<void>(`/api/v1/patient/me/appointments/${id}/request-cancel`, { reason }),
  requestReschedule: (id: string, reason: string) =>
    apiClient.post<void>(
      `/api/v1/patient/me/appointments/${id}/request-reschedule`,
      { reason },
    ),
};

export function useMyAppointments(range: "upcoming" | "past" = "upcoming") {
  return useQuery({
    queryKey: ["patient-appointments", range],
    queryFn: () => patientAppointmentApi.list(range),
    staleTime: 30_000,
  });
}

export function useMyAppointment(id: string | undefined) {
  return useQuery({
    queryKey: ["patient-appointment", id],
    queryFn: () => patientAppointmentApi.get(id!),
    enabled: !!id,
    refetchInterval: (q) => {
      // Polling se telemed dentro de janela: server pode liberar dailyRoomUrl em qualquer minuto
      const data = q.state.data;
      if (data?.isTelemedicine && Math.abs(data.minutesUntilStart) < 30) return 15_000;
      return false;
    },
  });
}

export function useConfirmAppointment() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: patientAppointmentApi.confirm,
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: ["patient-appointments"] });
      qc.invalidateQueries({ queryKey: ["patient-appointment"] });
      qc.invalidateQueries({ queryKey: ["patient-dashboard"] });
    },
  });
}

export function useRequestAppointmentAction() {
  return useMutation({
    mutationFn: ({ id, action, reason }: { id: string; action: "cancel" | "reschedule"; reason: string }) =>
      action === "cancel"
        ? patientAppointmentApi.requestCancel(id, reason)
        : patientAppointmentApi.requestReschedule(id, reason),
  });
}

// ============================================================
// Mensagens
// ============================================================

export interface PatientMessage {
  id: string;
  createdAt: string;
  from: "patient" | "team";
  authorName?: string;
  channel: string;
  content: string;
}

export const patientMessagesApi = {
  list: (before?: string) => {
    const qs = before ? `?before=${encodeURIComponent(before)}` : "";
    return apiClient.get<PatientMessage[]>(`/api/v1/patient/me/messages${qs}`);
  },
  send: (content: string) =>
    apiClient.post<PatientMessage>("/api/v1/patient/me/messages", { content }),
  unreadCount: () =>
    apiClient.get<{ unread: number }>("/api/v1/patient/me/messages/unread-count"),
  markRead: () => apiClient.post<void>("/api/v1/patient/me/messages/read", {}),
};

export function useMyMessages() {
  return useQuery({
    queryKey: ["patient-messages"],
    queryFn: () => patientMessagesApi.list(),
    refetchInterval: 30_000, // pull a cada 30s pra capturar respostas da equipe
    staleTime: 10_000,
  });
}

export function useMyMessagesUnread() {
  return useQuery({
    queryKey: ["patient-messages-unread"],
    queryFn: patientMessagesApi.unreadCount,
    refetchInterval: 60_000,
  });
}

export function useSendMessage() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: patientMessagesApi.send,
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: ["patient-messages"] });
    },
  });
}

export function useMarkMessagesRead() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: patientMessagesApi.markRead,
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: ["patient-messages-unread"] });
      qc.invalidateQueries({ queryKey: ["patient-dashboard"] });
    },
  });
}

export function useSetPatientPassword() {
  return useMutation({
    mutationFn: patientMeApi.setPassword,
  });
}

// ============================================================
// Convite de portal (uso staff — disparado no perfil do Patient)
// ============================================================

export interface PortalInviteStatus {
  status: "none" | "pending" | "accepted" | "expired";
  id?: string;
  createdAt?: string;
  expiresAt?: string;
  acceptedAt?: string | null;
}

export interface CreatePortalInviteResult {
  inviteId: string;
  link: string;
  expiresAt: string;
  sentEmail: boolean;
  sentWA: boolean;
}

export const portalInviteStaffApi = {
  status: (patientId: string) =>
    apiClient.get<PortalInviteStatus>(
      `/api/v1/patients/${patientId}/portal-invite`,
    ),
  create: (
    patientId: string,
    payload: { sendEmail: boolean; sendWA: boolean },
  ) =>
    apiClient.post<CreatePortalInviteResult>(
      `/api/v1/patients/${patientId}/portal-invite`,
      payload,
    ),
};

export function usePortalInviteStatus(patientId: string | undefined) {
  return useQuery({
    queryKey: ["portal-invite-status", patientId],
    queryFn: () => portalInviteStaffApi.status(patientId!),
    enabled: !!patientId,
  });
}

export function useCreatePortalInvite(patientId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (payload: { sendEmail: boolean; sendWA: boolean }) =>
      portalInviteStaffApi.create(patientId, payload),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: ["portal-invite-status", patientId] });
    },
  });
}
