import * as Application from 'expo-application';
import { configureFetcher, options } from '@plenya/api-client';
import { Platform } from 'react-native';
import { env } from '../../lib/env';
import { useAuthStore } from './authStore';

let refreshPromise: Promise<string | null> | null = null;

async function refreshAccessToken(): Promise<string | null> {
  const state = useAuthStore.getState();
  if (!state.refreshToken) return null;
  if (refreshPromise) return refreshPromise;

  refreshPromise = (async () => {
    try {
      const res = await options.authMutations.refresh({ refreshToken: state.refreshToken! });
      await state.setTokens(res.accessToken, res.refreshToken);
      return res.accessToken;
    } catch {
      await state.clear();
      return null;
    } finally {
      refreshPromise = null;
    }
  })();

  return refreshPromise;
}

export function configureApiClient(): void {
  configureFetcher({
    baseUrl: env.apiBaseUrl,
    getAccessToken: async () => {
      const token = useAuthStore.getState().accessToken;
      return token;
    },
    onUnauthorized: async () => {
      await refreshAccessToken();
    },
    extraHeaders: () => ({
      'X-App-Variant': 'pro',
      'X-App-Version': Application.nativeApplicationVersion ?? '0.0.0',
      'X-Platform': Platform.OS,
    }),
  });
}
