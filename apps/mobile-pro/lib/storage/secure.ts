import * as SecureStore from 'expo-secure-store';

const PREFIX = 'plenya.pro.';

const options: SecureStore.SecureStoreOptions = {
  keychainAccessible: SecureStore.AFTER_FIRST_UNLOCK_THIS_DEVICE_ONLY,
};

export const secureStorage = {
  get(key: string): Promise<string | null> {
    return SecureStore.getItemAsync(PREFIX + key, options);
  },
  set(key: string, value: string): Promise<void> {
    return SecureStore.setItemAsync(PREFIX + key, value, options);
  },
  remove(key: string): Promise<void> {
    return SecureStore.deleteItemAsync(PREFIX + key, options);
  },
};

export const SecureKeys = {
  AccessToken: 'auth.access',
  RefreshToken: 'auth.refresh',
  MmkvEncryptionKey: 'mmkv.key',
  BiometricGate: 'auth.biometric_gate',
  ConsentLgpd: 'consent.lgpd',
} as const;
