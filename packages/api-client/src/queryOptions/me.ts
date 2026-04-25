import { queryOptions } from '@tanstack/react-query';
import { api } from '../fetcher';
import { queryKeys } from '../queryKeys';

export interface UserProfileSelectedPatient {
  id: string;
  name: string;
  avatarUrl?: string;
}

export interface UserProfile {
  id: string;
  email: string;
  name: string;
  role: string;
  phone?: string;
  avatarUrl?: string;
  has2FAEnabled: boolean;
  selectedPatientId?: string;
  selectedPatient?: UserProfileSelectedPatient;
}

export interface Session {
  id: string;
  platform: 'ios' | 'android' | 'web';
  appVariant: 'pro' | 'app' | 'web';
  device: string;
  appVersion?: string;
  lastSeenAt: string;
  current: boolean;
}

export const meOptions = () =>
  queryOptions({
    queryKey: queryKeys.me(),
    queryFn: ({ signal }) => api.get<UserProfile>('/api/v1/me', { signal }),
    staleTime: 60_000,
  });

export const meSessionsOptions = () =>
  queryOptions({
    queryKey: queryKeys.meSessions(),
    queryFn: ({ signal }) => api.get<Session[]>('/api/v1/me/sessions', { signal }),
  });

export interface DataExport {
  exportedAt: string;
  userId: string;
  user: Record<string, unknown>;
  sessions: Array<Record<string, unknown>>;
  notifications: Array<Record<string, unknown>>;
  preferences?: unknown;
}

export const meMutations = {
  revokeSession: (sessionId: string) =>
    api.delete<void>(`/api/v1/me/sessions/${sessionId}`),
  registerDeviceToken: (body: {
    platform: 'ios' | 'android';
    token: string;
    appVariant: 'pro' | 'app';
  }) => api.post<{ id: string }>('/api/v1/me/device-tokens', body),
  removeDeviceToken: (id: string) => api.delete<void>(`/api/v1/me/device-tokens/${id}`),
  acceptLgpdConsent: () => api.post<void>('/api/v1/me/consent/lgpd'),
  changePassword: (body: { currentPassword: string; newPassword: string }) =>
    api.post<void>('/api/v1/me/password', body),
  disable2FA: (body: { password: string }) =>
    api.post<void>('/api/v1/me/2fa/disable', body),
  exportData: () => api.get<DataExport>('/api/v1/me/export'),
};
