import { useAuthStore } from "./auth-store";
import { toast } from "sonner";

const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:3001";

class APIClient {
  private baseURL: string;
  private isRefreshing = false;
  private refreshPromise: Promise<boolean> | null = null;

  constructor(baseURL: string) {
    this.baseURL = baseURL;
  }

  private async tryRefreshToken(): Promise<boolean> {
    // Se já está fazendo refresh, aguarda a promise existente
    if (this.isRefreshing && this.refreshPromise) {
      return this.refreshPromise;
    }

    this.isRefreshing = true;
    this.refreshPromise = this._doRefresh();

    try {
      const result = await this.refreshPromise;
      return result;
    } finally {
      this.isRefreshing = false;
      this.refreshPromise = null;
    }
  }

  private async _doRefresh(): Promise<boolean> {
    const { refreshToken } = useAuthStore.getState();

    if (!refreshToken) {
      return false;
    }

    try {
      const response = await fetch(`${this.baseURL}/api/v1/auth/refresh`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ refreshToken }),
      });

      if (response.ok) {
        const data = await response.json();
        useAuthStore
          .getState()
          .setAuth(data.user, data.accessToken, data.refreshToken);
        return true;
      }
    } catch (error) {
      console.error("Error refreshing token:", error);
    }

    return false;
  }

  private async fetchWithAuth(
    endpoint: string,
    options: RequestInit = {}
  ): Promise<Response> {
    const { accessToken } = useAuthStore.getState();

    const headers: Record<string, string> = {
      ...(options.headers as Record<string, string>),
    };

    // Don't set Content-Type for FormData - let browser handle it
    const isFormData = options.body instanceof FormData;
    if (!isFormData && !headers["Content-Type"]) {
      headers["Content-Type"] = "application/json";
    }

    if (accessToken) {
      headers.Authorization = `Bearer ${accessToken}`;
    }

    return fetch(`${this.baseURL}${endpoint}`, {
      ...options,
      headers,
    });
  }

  private async request<T>(
    endpoint: string,
    options: RequestInit = {}
  ): Promise<T> {
    // Primeira tentativa
    let response = await this.fetchWithAuth(endpoint, options);

    // Se 401 e não é o endpoint de refresh, tenta renovar token
    if (response.status === 401 && endpoint !== "/api/v1/auth/refresh") {
      const refreshed = await this.tryRefreshToken();

      if (refreshed) {
        // Retry com novo token
        response = await this.fetchWithAuth(endpoint, options);
      } else {
        // Refresh falhou - fazer logout e avisar usuário
        useAuthStore.getState().clearAuth();
        toast.error("Sua sessão expirou", {
          description: "Por favor, faça login novamente.",
        });

        // Redirecionar para login após um pequeno delay
        setTimeout(() => {
          if (typeof window !== "undefined") {
            window.location.href = "/login";
          }
        }, 1500);

        throw new Error("Session expired");
      }
    }

    if (!response.ok) {
      const error = await response.json().catch(() => ({
        error: "Unknown error",
        message: response.statusText,
      }));
      // Extract error message from various backend response formats
      const errorMessage = error.error || error.message || response.statusText || "Unknown error";

      // Create error object with the message
      const apiError = new Error(errorMessage);
      // Attach original error data for debugging
      (apiError as any).data = error;
      throw apiError;
    }

    // Handle 204 No Content (common for DELETE operations)
    if (response.status === 204) {
      return {} as T;
    }

    return response.json();
  }

  async get<T>(endpoint: string): Promise<T> {
    return this.request<T>(endpoint, { method: "GET" });
  }

  async post<T>(endpoint: string, data?: unknown, options?: RequestInit): Promise<T> {
    const isFormData = data instanceof FormData;
    return this.request<T>(endpoint, {
      method: "POST",
      body: isFormData ? data : (data ? JSON.stringify(data) : undefined),
      ...options,
    });
  }

  async put<T>(endpoint: string, data?: unknown, options?: RequestInit): Promise<T> {
    const isFormData = data instanceof FormData;
    return this.request<T>(endpoint, {
      method: "PUT",
      body: isFormData ? data : (data ? JSON.stringify(data) : undefined),
      ...options,
    });
  }

  async patch<T>(endpoint: string, data?: unknown, options?: RequestInit): Promise<T> {
    const isFormData = data instanceof FormData;
    return this.request<T>(endpoint, {
      method: "PATCH",
      body: isFormData ? data : (data ? JSON.stringify(data) : undefined),
      ...options,
    });
  }

  async delete<T>(endpoint: string): Promise<T> {
    return this.request<T>(endpoint, { method: "DELETE" });
  }
}

export const apiClient = new APIClient(API_URL);
