import { create } from "zustand";
import { persist, createJSONStorage } from "zustand/middleware";

const DEV_BYPASS_AUTH = process.env.NEXT_PUBLIC_DEV_BYPASS_AUTH === 'true';

export type UserRole = "admin" | "doctor" | "nurse" | "patient" | "nutritionist" | "psychologist" | "physicalEducator" | "secretary" | "manager";

export interface Patient {
  id: string;
  userId: string;
  name: string;
  birthDate: string;
  gender: "male" | "female" | "other";
  age?: number;
  menopause?: boolean;
  hormoneTherapy?: boolean;
  phone?: string;
  address?: string;
  municipality?: string;
  state?: string;
  motherName?: string;
  fatherName?: string;
  height?: number;
  weight?: number;
  createdAt: string;
  updatedAt: string;
}

export interface User {
  id: string;
  name?: string;
  email: string;
  roles: UserRole[];
  twoFactorEnabled: boolean;
  selectedPatientId?: string;
  selectedPatient?: Patient;
  preferences?: Record<string, any>;
  createdAt: string;
}

// Helper function para verificar se usuário tem role específico
export function isGranted(user: User | null, role: UserRole): boolean {
  if (!user) return false;
  return user.roles.includes(role);
}

interface AuthState {
  user: User | null;
  accessToken: string | null;
  refreshToken: string | null;
  /**
   * A store já leu o localStorage?
   *
   * Antes de ser true, "sem token" NÃO significa "sem sessão" — significa que a leitura ainda não
   * aconteceu. Confundir as duas coisas era o que derrubava o login do PWA: uma requisição saía
   * antes da hidratação, tomava 401, ia renovar, achava a store vazia e concluía que a sessão
   * tinha acabado — apagando um refresh token que estava válido no aparelho e no servidor.
   */
  hasHydrated: boolean;
  setAuth: (user: User, accessToken: string, refreshToken: string) => void;
  clearAuth: () => void;
  updateAccessToken: (accessToken: string) => void;
  updateUser: (user: User) => void;
}

const initialState = DEV_BYPASS_AUTH
  ? {
      user: {
        id: 'dev-admin-placeholder',
        email: 'admin@plenya.com',
        name: 'Dev Admin',
        roles: ['admin'] as UserRole[],
        twoFactorEnabled: false,
        createdAt: new Date().toISOString(),
      },
      accessToken: 'dev-bypass-token',
      refreshToken: 'dev-bypass-token',
    }
  : {
      user: null,
      accessToken: null,
      refreshToken: null,
    };

export const useAuthStore = create<AuthState>()(
  persist(
    (set) => ({
      ...initialState,
      // Com bypass de dev não há leitura de storage a esperar.
      hasHydrated: DEV_BYPASS_AUTH,
      setAuth: (user, accessToken, refreshToken) =>
        set({ user, accessToken, refreshToken }),
      clearAuth: () => {
        // Com bypass ativo, apenas recarregar (não limpar store)
        if (DEV_BYPASS_AUTH) {
          window.location.href = '/';
          return;
        }
        set({ user: null, accessToken: null, refreshToken: null });
      },
      updateAccessToken: (accessToken) => set({ accessToken }),
      updateUser: (user) => set({ user }),
    }),
    {
      name: "plenya-auth",
      storage: createJSONStorage(() => localStorage),
      // hasHydrated é estado de runtime, não da sessão: não vai para o storage.
      partialize: ({ user, accessToken, refreshToken }) => ({ user, accessToken, refreshToken }),
      // Marca a hidratação inclusive quando ela FALHA (storage bloqueado no iOS, por exemplo):
      // o app precisa sair do limbo de qualquer jeito, e aí "sem token" vira resposta legítima.
      onRehydrateStorage: () => () => {
        useAuthStore.setState({ hasHydrated: true });
      },
    }
  )
);

/**
 * Resolve quando a store terminou de ler o localStorage.
 *
 * Quem decide "esta sessão acabou" precisa esperar por isto. Sem a espera, o boot do PWA no
 * celular — mais lento que o do desktop — dispara requisição antes da leitura e conclui que não
 * há sessão, apagando a que existia.
 */
export function waitForAuthHydration(timeoutMs = 3000): Promise<void> {
  if (typeof window === "undefined") return Promise.resolve();
  if (useAuthStore.getState().hasHydrated) return Promise.resolve();

  return new Promise((resolve) => {
    const done = () => {
      clearTimeout(timer);
      unsubscribe();
      resolve();
    };
    // Teto de segurança: se a hidratação nunca sinalizar, o app não pode ficar preso.
    const timer = setTimeout(done, timeoutMs);
    const unsubscribe = useAuthStore.subscribe((state) => {
      if (state.hasHydrated) done();
    });
  });
}
