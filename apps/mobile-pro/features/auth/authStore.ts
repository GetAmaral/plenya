import { create } from 'zustand';
import { secureStorage, SecureKeys } from '../../lib/storage/secure';
import type { UserProfile } from '@plenya/api-client';

interface AuthState {
  user: UserProfile | null;
  accessToken: string | null;
  refreshToken: string | null;
  biometricUnlocked: boolean;
  lgpdAccepted: boolean;
  hydrated: boolean;

  hydrate: () => Promise<void>;
  setTokens: (access: string, refresh: string) => Promise<void>;
  setUser: (user: UserProfile | null) => void;
  markBiometricUnlocked: () => void;
  markLgpdAccepted: () => Promise<void>;
  clear: () => Promise<void>;
}

export const useAuthStore = create<AuthState>((set) => ({
  user: null,
  accessToken: null,
  refreshToken: null,
  biometricUnlocked: false,
  lgpdAccepted: false,
  hydrated: false,

  hydrate: async () => {
    const [access, refresh, consent] = await Promise.all([
      secureStorage.get(SecureKeys.AccessToken),
      secureStorage.get(SecureKeys.RefreshToken),
      secureStorage.get(SecureKeys.ConsentLgpd),
    ]);
    set({
      accessToken: access,
      refreshToken: refresh,
      lgpdAccepted: consent === '1',
      hydrated: true,
    });
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

  markLgpdAccepted: async () => {
    await secureStorage.set(SecureKeys.ConsentLgpd, '1');
    set({ lgpdAccepted: true });
  },

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
