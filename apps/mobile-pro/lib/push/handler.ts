import { useEffect, useRef } from 'react';
import * as Notifications from 'expo-notifications';
import { router } from 'expo-router';
import * as Linking from 'expo-linking';

/**
 * Roteia uma URL recebida via push notification para a rota correta do app.
 * Aceita tanto deep links absolutos (plenyapro://patients/abc) quanto
 * paths relativos (/patients/abc).
 */
function routeFromUrl(url: string): void {
  if (!url) return;
  if (url.startsWith('plenyapro://') || url.startsWith('http')) {
    Linking.openURL(url).catch(() => {});
    return;
  }
  router.push(url as never);
}

/**
 * Hook que escuta notificações enquanto o app está aberto e quando o usuário
 * toca numa notificação para abrir o app. Usar em (tabs)/_layout.tsx pra que
 * só dispare em sessão autenticada.
 */
export function usePushNotificationRouter(): void {
  const lastTappedRef = useRef<string | null>(null);

  useEffect(() => {
    const handleResponse = (
      response: Notifications.NotificationResponse,
    ) => {
      const url =
        (response.notification.request.content.data as { url?: string } | undefined)?.url;
      if (!url) return;
      if (lastTappedRef.current === url) return;
      lastTappedRef.current = url;
      routeFromUrl(url);
    };

    const sub = Notifications.addNotificationResponseReceivedListener(handleResponse);

    Notifications.getLastNotificationResponseAsync()
      .then((res) => {
        if (res) handleResponse(res);
      })
      .catch(() => {});

    return () => sub.remove();
  }, []);
}
