import { create } from "zustand";
import { persist, createJSONStorage } from "zustand/middleware";

export type UserRole = "admin" | "doctor" | "nurse" | "patient";

export interface User {
  id: string;
  email: string;
  role: UserRole;
  twoFactorEnabled: boolean;
  createdAt: string;
}

interface AuthState {
  user: User | null;
  accessToken: string | null;
  refreshToken: string | null;
  setAuth: (user: User, accessToken: string, refreshToken: string) => void;
  clearAuth: () => void;
  updateAccessToken: (accessToken: string) => void;
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
    }),
    {
      name: "plenya-auth",
      storage: createJSONStorage(() => localStorage),
    }
  )
);
