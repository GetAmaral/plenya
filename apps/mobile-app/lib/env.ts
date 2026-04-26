import Constants from 'expo-constants';

type Extra = {
  apiBaseUrl?: string;
  sentryDsn?: string;
  posthogApiKey?: string;
  posthogHost?: string;
};

const extra = (Constants.expoConfig?.extra ?? {}) as Extra;

function firstNonEmpty(...values: (string | undefined | null)[]): string | undefined {
  for (const v of values) {
    if (typeof v === 'string' && v.length > 0) return v;
  }
  return undefined;
}

export const env = {
  apiBaseUrl:
    firstNonEmpty(
      process.env.EXPO_PUBLIC_API_BASE_URL,
      process.env.API_BASE_URL,
      extra.apiBaseUrl,
    ) ?? 'http://localhost:3001',
  sentryDsn: firstNonEmpty(process.env.EXPO_PUBLIC_SENTRY_DSN, extra.sentryDsn),
  posthogApiKey: firstNonEmpty(process.env.EXPO_PUBLIC_POSTHOG_API_KEY, extra.posthogApiKey),
  posthogHost: firstNonEmpty(process.env.EXPO_PUBLIC_POSTHOG_HOST, extra.posthogHost),
  appVariant: (process.env.APP_VARIANT ?? 'development') as
    | 'development'
    | 'preview'
    | 'production',
} as const;
