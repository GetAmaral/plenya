import * as LocalAuthentication from 'expo-local-authentication';

export type BiometricResult =
  | { ok: true }
  | { ok: false; reason: 'unavailable' | 'cancelled' | 'failed' };

export async function isBiometricAvailable(): Promise<boolean> {
  const compatible = await LocalAuthentication.hasHardwareAsync();
  if (!compatible) return false;
  const enrolled = await LocalAuthentication.isEnrolledAsync();
  return enrolled;
}

export async function authenticateBiometric(
  reason: string = 'Confirme sua identidade',
): Promise<BiometricResult> {
  const available = await isBiometricAvailable();
  if (!available) return { ok: false, reason: 'unavailable' };

  const res = await LocalAuthentication.authenticateAsync({
    promptMessage: reason,
    disableDeviceFallback: false,
    cancelLabel: 'Cancelar',
  });
  if (res.success) return { ok: true };
  return { ok: false, reason: res.error === 'user_cancel' ? 'cancelled' : 'failed' };
}
