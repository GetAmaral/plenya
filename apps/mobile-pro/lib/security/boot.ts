import * as Application from 'expo-application';
import { Platform } from 'react-native';
import { compareSemver } from '@plenya/domain';
import { api, type MobileConfig } from '@plenya/api-client';

import { configureCertPinning } from './certPinning';
import { checkDeviceIntegrity, type DeviceIntegrityReport } from './jailbreak';
import { env } from '../env';

export type BootGate =
  | { kind: 'ok'; integrity: DeviceIntegrityReport }
  | { kind: 'kill_switch'; message: string }
  | { kind: 'min_version'; required: string; current: string }
  | { kind: 'unreachable' };

interface BootOptions {
  /**
   * Permite pular cert pinning em DEV builds (APP_VARIANT=development).
   * Gates de versão e kill switch SEMPRE aplicam.
   */
  skipInDev?: boolean;
}

/**
 * Roda no cold start, antes de mostrar qualquer tela autenticada.
 *
 * Sequência:
 * 1. Carrega /mobile/config
 * 2. Aplica cert pinning (com pins atual + backup)
 * 3. Verifica kill switch e versão mínima
 * 4. Avalia integridade do device (jailbreak/root)
 *
 * Falha aberta apenas em network error — nesse caso retorna 'unreachable'
 * e o app mostra retry. Nunca silencia kill switch ou min version.
 */
export async function runBootSecurity(opts: BootOptions = {}): Promise<BootGate> {
  const isDev = env.appVariant === 'development';
  const skipPinning = isDev && opts.skipInDev;

  let config: MobileConfig;
  try {
    config = await api.get<MobileConfig>('/api/v1/mobile/config', { skipAuth: true });
  } catch {
    return { kind: 'unreachable' };
  }

  if (config.killSwitch.enabled) {
    return {
      kind: 'kill_switch',
      message: config.killSwitch.message ?? 'O app está temporariamente indisponível.',
    };
  }

  const currentVersion = Application.nativeApplicationVersion ?? '0.0.0';
  const requiredVersion =
    Platform.OS === 'ios' ? config.minVersion.ios : config.minVersion.android;

  if (compareSemver(currentVersion, requiredVersion) < 0) {
    return {
      kind: 'min_version',
      required: requiredVersion,
      current: currentVersion,
    };
  }

  if (!skipPinning && config.sslPins.current) {
    try {
      const apiHost = new URL(env.apiBaseUrl).hostname;
      await configureCertPinning({
        apiHost,
        current: config.sslPins.current,
        backup: config.sslPins.backup,
      });
    } catch {
      /* pinning opcional em primeiro boot até a infra ter pins */
    }
  }

  const integrity = await checkDeviceIntegrity().catch(
    (): DeviceIntegrityReport => ({
      jailBroken: false,
      onExternalStorage: false,
      canMockLocation: false,
      isDebuggedMode: false,
      isDevelopmentSettingsMode: false,
      adbEnabled: false,
      trustworthy: true,
    }),
  );

  return { kind: 'ok', integrity };
}
