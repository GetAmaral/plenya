import { useCallback } from 'react';
import { useQueryClient } from '@tanstack/react-query';
import { router } from 'expo-router';
import { options } from '@plenya/api-client';
import { useAuthStore } from './authStore';

export function useLogout() {
  const clear = useAuthStore((s) => s.clear);
  const queryClient = useQueryClient();

  return useCallback(async () => {
    try {
      await options.authMutations.logout();
    } catch {
      /* even if server logout fails, clear local state */
    }
    await clear();
    queryClient.clear();
    router.replace('/(auth)/login');
  }, [clear, queryClient]);
}
