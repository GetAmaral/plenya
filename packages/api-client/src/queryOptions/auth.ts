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
