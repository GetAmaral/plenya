import { create } from 'zustand';
import { secureStorage, SecureKeys } from '../../lib/storage/secure';
import type { options } from '@plenya/api-client';

type UserProfile = Awaited<ReturnType<ReturnType<typeof options.meOptions>['queryFn']>>;

interface AuthState {
  user: UserProfile | null;
  accessToken: string | null;
  refreshToken: string | null;
  biometricUnlocked: boolean;
  hydrated: boolean;

  hydrate: () => Promise<void>;
  setTokens: (access: string, refresh: string) => Promise<void>;
  setUser: (user: UserProfile | null) => void;
  markBiometricUnlocked: () => void;
  clear: () => Promise<void>;
}

export const useAuthStore = create<AuthState>((set) => ({
  user: null,
  accessToken: null,
  refreshToken: null,
  biometricUnlocked: false,
  hydrated: false,

  hydrate: async () => {
    const [access, refresh] = await Promise.all([
      secureStorage.get(SecureKeys.AccessToken),
      secureStorage.get(SecureKeys.RefreshToken),
    ]);
    set({ accessToken: access, refreshToken: refresh, hydrated: true });
  },

  setTokens: async (access, refresh) => {
    await Promise.all([
      secureStorage.set(SecureKeys.AccessToken, access),
      secureStorage.set(SecureKeys.RefreshToken, refresh),
    ]);
    set({ accessToken: access, refreshToken: refresh });
  },

  setUser: (user) => set({ user }),

  markBiometricUnlocked: () => set({ biometricUnlocked: true }),

  clear: async () => {
    await Promise.all([
      secureStorage.remove(SecureKeys.AccessToken),
      secureStorage.remove(SecureKeys.RefreshToken),
    ]);
    set({
      user: null,
      accessToken: null,
      refreshToken: null,
      biometricUnlocked: false,
    });
  },
}));

export function getAccessToken(): string | null {
  return useAuthStore.getState().accessToken;
}
