import '../global.css';

import { useEffect, useState } from 'react';
import { useFonts } from 'expo-font';
import { Slot, SplashScreen } from 'expo-router';
import { GestureHandlerRootView } from 'react-native-gesture-handler';
import { SafeAreaProvider } from 'react-native-safe-area-context';
import { QueryClientProvider } from '@tanstack/react-query';
import { persistQueryClient } from '@tanstack/react-query-persist-client';
import { ToastProvider } from '@plenya/ui-mobile';
import { StatusBar } from 'expo-status-bar';

import { createQueryClient } from '../lib/queryClient';
import { getEncryptedStorage, mmkvToQueryPersisterStorage } from '../lib/storage/mmkv';
import { configureApiClient } from '../features/auth/configureApiClient';
import { useAuthStore } from '../features/auth/authStore';

SplashScreen.preventAutoHideAsync().catch(() => {});

const queryClient = createQueryClient();

export default function RootLayout() {
  const [ready, setReady] = useState(false);
  const hydrate = useAuthStore((s) => s.hydrate);
  const [fontsLoaded] = useFonts({});

  useEffect(() => {
    (async () => {
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
      setReady(true);
    })();
  }, [hydrate]);

  useEffect(() => {
    if (ready && fontsLoaded) SplashScreen.hideAsync().catch(() => {});
  }, [ready, fontsLoaded]);

  if (!ready || !fontsLoaded) return null;

  // suppress unused var linter: mmkvToQueryPersisterStorage kept as helper
  void mmkvToQueryPersisterStorage;

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
