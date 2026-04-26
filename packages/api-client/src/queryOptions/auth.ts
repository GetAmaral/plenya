import { api } from '../fetcher';

export interface LoginRequest {
  email: string;
  password: string;
}

export interface LoginResponse {
  accessToken: string;
  refreshToken: string;
  requires2FA: boolean;
  user: {
    id: string;
    email: string;
    name: string;
    role: string;
  };
}

export interface TwoFactorVerifyRequest {
  email: string;
  code: string;
  pendingToken: string;
}

export interface RefreshRequest {
  refreshToken: string;
}

export const authMutations = {
  login: (body: LoginRequest) =>
    api.post<LoginResponse>('/api/v1/auth/login', body, { skipAuth: true }),
  verifyTwoFactor: (body: TwoFactorVerifyRequest) =>
    api.post<LoginResponse>('/api/v1/auth/2fa/verify', body, { skipAuth: true }),
  refresh: (body: RefreshRequest) =>
    api.post<Pick<LoginResponse, 'accessToken' | 'refreshToken'>>(
      '/api/v1/auth/refresh',
      body,
      { skipAuth: true },
    ),
  logout: () => api.post<void>('/api/v1/auth/logout'),
};

// Patient auth — endpoints separados em /auth/patient/*. Resposta sem
// requires2FA (paciente pode ativar 2FA mas não é obrigatório no login).
export interface PatientAuthResponse {
  accessToken: string;
  refreshToken: string;
  user: {
    id: string;
    email: string;
    name: string;
    roles: string[];
    twoFactorEnabled: boolean;
  };
}

export const patientAuthMutations = {
  loginPassword: (body: { email: string; password: string }) =>
    api.post<PatientAuthResponse>('/api/v1/auth/patient/login', body, { skipAuth: true }),
  requestMagicLink: (body: { email: string }) =>
    api.post<void>('/api/v1/auth/patient/magic-link', body, { skipAuth: true }),
  consumeMagicLink: (body: { token: string }) =>
    api.post<PatientAuthResponse>('/api/v1/auth/patient/magic-link/consume', body, {
      skipAuth: true,
    }),
  consumeInvite: (body: { token: string }) =>
    api.post<PatientAuthResponse>('/api/v1/auth/patient/invite/consume', body, {
      skipAuth: true,
    }),
};
