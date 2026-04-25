import { queryOptions } from '@tanstack/react-query';
import { api } from '../fetcher';
import { queryKeys } from '../queryKeys';

export interface MobileConfig {
  minVersion: {
    ios: string;
    android: string;
  };
  killSwitch: {
    enabled: boolean;
    message?: string;
  };
  sslPins: {
    current: string;
    backup: string;
  };
  featureFlags: Record<string, boolean>;
  supportContact: {
    email: string;
    whatsapp?: string;
  };
}

export const mobileConfigOptions = () =>
  queryOptions({
    queryKey: queryKeys.mobileConfig(),
    queryFn: ({ signal }) =>
      api.get<MobileConfig>('/api/v1/mobile/config', { signal, skipAuth: true }),
    staleTime: 5 * 60_000,
  });
