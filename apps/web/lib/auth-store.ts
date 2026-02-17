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
    }
  )
);
