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
};

export function usePatientMe(enabled = true) {
  return useQuery({
    queryKey: ["patient-me"],
    queryFn: patientMeApi.me,
    enabled,
    staleTime: 60_000,
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
