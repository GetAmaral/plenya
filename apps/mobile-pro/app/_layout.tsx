import '../global.css';

import { useEffect, useState } from 'react';
import { useFonts } from 'expo-font';
import { router, Slot, SplashScreen } from 'expo-router';
import { GestureHandlerRootView } from 'react-native-gesture-handler';
import { SafeAreaProvider } from 'react-native-safe-area-context';
import { QueryClientProvider } from '@tanstack/react-query';
import { persistQueryClient } from '@tanstack/react-query-persist-client';
import { ToastProvider } from '@plenya/ui-mobile';
import { StatusBar } from 'expo-status-bar';

import { createQueryClient } from '../lib/queryClient';
import { getEncryptedStorage } from '../lib/storage/mmkv';
import { configureApiClient } from '../features/auth/configureApiClient';
import { useAuthStore } from '../features/auth/authStore';
import { runBootSecurity } from '../lib/security/boot';
import { initSentry, Sentry } from '../lib/observability/sentry';
import { getPostHog } from '../lib/observability/posthog';

SplashScreen.preventAutoHideAsync().catch(() => {});

const queryClient = createQueryClient();

function RootLayout() {
  const [ready, setReady] = useState(false);
  const hydrate = useAuthStore((s) => s.hydrate);
  const [fontsLoaded] = useFonts({});

  useEffect(() => {
    (async () => {
      initSentry();
      getPostHog(); // lazy init — no-op se não tem API key
      configureApiClient();
      const mmkv = await getEncryptedStorage();
      persistQueryClient({
        queryClient,
        persister: {
          persistClient: async (client) => {
            mmkv.set('rq-cache', JSON.stringify(client));
          },
          restoreClient: async () => {
            const raw = mmkv.getString('rq-cache');
            return raw ? JSON.parse(raw) : undefined;
          },
          removeClient: async () => {
            mmkv.delete('rq-cache');
          },
        },
        maxAge: 24 * 60 * 60 * 1000,
      });
      await hydrate();

      const gate = await runBootSecurity({ skipInDev: true });
      if (gate.kind === 'kill_switch') {
        router.replace({ pathname: '/(blocked)/kill-switch', params: { message: gate.message } });
      } else if (gate.kind === 'min_version') {
        router.replace({
          pathname: '/(blocked)/min-version',
          params: { required: gate.required, current: gate.current },
        });
      } else if (gate.kind === 'unreachable') {
        router.replace('/(blocked)/unreachable');
      }

      setReady(true);
    })();
  }, [hydrate]);

  useEffect(() => {
    if (ready && fontsLoaded) SplashScreen.hideAsync().catch(() => {});
  }, [ready, fontsLoaded]);

  if (!ready || !fontsLoaded) return null;

  return (
    <GestureHandlerRootView style={{ flex: 1 }}>
      <SafeAreaProvider>
        <QueryClientProvider client={queryClient}>
          <ToastProvider>
            <StatusBar style="auto" />
            <Slot />
          </ToastProvider>
        </QueryClientProvider>
      </SafeAreaProvider>
    </GestureHandlerRootView>
  );
}

export default Sentry.wrap(RootLayout);
