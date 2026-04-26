import PostHog from 'posthog-react-native';
import { env } from '../env';

let client: PostHog | null = null;

/**
 * Inicializa PostHog. Lazy — só cria a instância na primeira chamada,
 * pra que o config do app possa rodar antes (env vars vêm de Constants).
 *
 * No-op se EXPO_PUBLIC_POSTHOG_API_KEY não estiver setado.
 */
export function getPostHog(): PostHog | null {
  if (client) return client;
  if (!env.posthogApiKey) return null;

  client = new PostHog(env.posthogApiKey, {
    host: env.posthogHost ?? 'https://app.posthog.com',
    enableSessionReplay: false,
    captureNativeAppLifecycleEvents: true,
    flushAt: 20,
    flushInterval: 30_000,
  });

  return client;
}

export function trackEvent(event: string, properties?: Record<string, unknown>): void {
  const ph = getPostHog();
  if (!ph) return;
  ph.capture(event, properties as Parameters<typeof ph.capture>[1]);
}
