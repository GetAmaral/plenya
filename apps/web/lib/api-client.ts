import { useAuthStore } from "./auth-store";
import { toast } from "sonner";

const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:3001";

/**
 * Resultado de uma tentativa de renovar a sessão.
 * - "ok": sessão renovada.
 * - "invalid": o servidor RECUSOU o refresh (401/403) — a sessão morreu de verdade.
 * - "offline": não deu pra falar com o servidor (rede caiu, app voltou do background,
 *   5xx). A sessão continua válida; insistir depois resolve.
 */
type RefreshOutcome = "ok" | "invalid" | "offline";

class APIClient {
  private baseURL: string;
  private isRefreshing = false;
  private refreshPromise: Promise<RefreshOutcome> | null = null;

  constructor(baseURL: string) {
    this.baseURL = baseURL;
  }

  private async tryRefreshToken(): Promise<RefreshOutcome> {
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

  private async _doRefresh(): Promise<RefreshOutcome> {
    const { refreshToken } = useAuthStore.getState();

    if (!refreshToken) {
      return "invalid";
    }

    // Uma segunda tentativa cobre o caso mais comum no celular: o PWA volta do background
    // e dispara o refresh antes de a rede estar de pé.
    for (let attempt = 0; attempt < 2; attempt++) {
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
          return "ok";
        }

        // Só 401/403 significam "esta sessão não vale mais". Qualquer outro status é
        // problema do servidor/rede — derrubar o login aqui era o que fazia o PWA do
        // celular pedir senha do nada.
        if (response.status === 401 || response.status === 403) {
          return "invalid";
        }
      } catch (error) {
        console.error("Error refreshing token:", error);
      }

      if (attempt === 0) {
        await new Promise((resolve) => setTimeout(resolve, 1200));
      }
    }

    return "offline";
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
      const outcome = await this.tryRefreshToken();

      if (outcome === "ok") {
        // Retry com novo token
        response = await this.fetchWithAuth(endpoint, options);
      } else if (outcome === "invalid") {
        // O servidor recusou o refresh: a sessão acabou mesmo.
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
      } else {
        // Sem rede/servidor fora: mantém a sessão. Ao voltar a conexão o próximo request
        // renova sozinho — o usuário não perde o login por causa de um sinal ruim.
        toast.error("Sem conexão com o servidor", {
          description: "Tente de novo em instantes.",
        });
        throw new Error("Network unavailable");
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
      // Attach original error data + HTTP status (ex: 409 conflito de horario)
      // para o chamador decidir a mensagem amigavel.
      (apiError as any).data = error;
      (apiError as any).status = response.status;
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

  async delete<T>(endpoint: string, data?: unknown): Promise<T> {
    return this.request<T>(endpoint, {
      method: "DELETE",
      body: data !== undefined ? JSON.stringify(data) : undefined,
    });
  }

  /** Faz GET autenticado e devolve o body como Blob — usado pra binários (PNG, PDF). */
  async getBlob(endpoint: string): Promise<Blob> {
    return (await this.getBlobWithName(endpoint)).blob;
  }

  /**
   * Igual ao getBlob, mas devolve também o nome do arquivo que o servidor mandou no
   * Content-Disposition. Sem isso o arquivo vira "unknown.pdf" ao ser compartilhado do
   * celular: blob URL não carrega nome nenhum.
   */
  async getBlobWithName(
    endpoint: string
  ): Promise<{ blob: Blob; filename: string | null }> {
    let response = await this.fetchWithAuth(endpoint, { method: "GET" });
    if (response.status === 401 && endpoint !== "/api/v1/auth/refresh") {
      const outcome = await this.tryRefreshToken();
      if (outcome === "ok") response = await this.fetchWithAuth(endpoint, { method: "GET" });
    }
    if (!response.ok) {
      throw new Error(`Falha ao baixar arquivo (${response.status})`);
    }
    return {
      blob: await response.blob(),
      filename: parseContentDispositionFilename(
        response.headers.get("Content-Disposition")
      ),
    };
  }
}


/**
 * Lê o nome do arquivo de um Content-Disposition. Prefere `filename*=UTF-8''…` (acentos
 * preservados) e cai pro `filename="…"` clássico. Devolve null quando o header não veio —
 * o header só chega ao JS se a API expuser via CORS (ExposeHeaders).
 */
function parseContentDispositionFilename(header: string | null): string | null {
  if (!header) return null;

  const utf8 = header.match(/filename\*\s*=\s*UTF-8''([^;]+)/i);
  if (utf8?.[1]) {
    try {
      return decodeURIComponent(utf8[1].trim());
    } catch {
      /* header malformado: tenta o clássico abaixo */
    }
  }

  const plain = header.match(/filename\s*=\s*"([^"]+)"/i) ?? header.match(/filename\s*=\s*([^;]+)/i);
  return plain?.[1]?.trim() || null;
}

export const apiClient = new APIClient(API_URL);
