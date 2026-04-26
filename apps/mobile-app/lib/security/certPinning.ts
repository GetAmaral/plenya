import { initializeSslPinning } from 'react-native-ssl-public-key-pinning';

interface PinsConfig {
  apiHost: string;
  current: string;
  backup: string;
}

/**
 * Call once during boot, after fetching /api/v1/mobile/config.
 * Hostname should match the API URL (without protocol).
 */
export async function configureCertPinning(config: PinsConfig): Promise<void> {
  await initializeSslPinning({
    [config.apiHost]: {
      includeSubdomains: true,
      publicKeyHashes: [config.current, config.backup].filter(Boolean),
    },
  });
}
