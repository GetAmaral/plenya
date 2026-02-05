import { apiClient } from "@/lib/api-client";

export interface AuthResponse {
  accessToken: string;
  refreshToken: string;
  user: {
    id: string;
    email: string;
    roles: string[];
    twoFactorEnabled: boolean;
    selectedPatientId?: string;
    createdAt: string;
  };
}

export async function loginWithGoogle(idToken: string): Promise<AuthResponse> {
  return apiClient.post<AuthResponse>("/api/v1/auth/oauth/google", {
    idToken,
  });
}

export async function loginWithApple(
  idToken: string,
  user?: { name?: { firstName: string; lastName: string } }
): Promise<AuthResponse> {
  return apiClient.post<AuthResponse>("/api/v1/auth/oauth/apple", {
    idToken,
    user,
  });
}
