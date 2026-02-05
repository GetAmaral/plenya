import { create } from "zustand";
import { persist, createJSONStorage } from "zustand/middleware";

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

export const useAuthStore = create<AuthState>()(
  persist(
    (set) => ({
      user: null,
      accessToken: null,
      refreshToken: null,
      setAuth: (user, accessToken, refreshToken) =>
        set({ user, accessToken, refreshToken }),
      clearAuth: () => set({ user: null, accessToken: null, refreshToken: null }),
      updateAccessToken: (accessToken) => set({ accessToken }),
      updateUser: (user) => set({ user }),
    }),
    {
      name: "plenya-auth",
      storage: createJSONStorage(() => localStorage),
    }
  )
);
