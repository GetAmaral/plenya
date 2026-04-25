import { MMKV } from 'react-native-mmkv';
import { secureStorage, SecureKeys } from './secure';

function generateEncryptionKey(): string {
  const bytes = new Uint8Array(32);
  if (typeof globalThis.crypto?.getRandomValues === 'function') {
    globalThis.crypto.getRandomValues(bytes);
  } else {
    for (let i = 0; i < bytes.length; i++) bytes[i] = Math.floor(Math.random() * 256);
  }
  return Array.from(bytes, (b) => b.toString(16).padStart(2, '0')).join('');
}

let cachedStorage: MMKV | null = null;

export async function getEncryptedStorage(): Promise<MMKV> {
  if (cachedStorage) return cachedStorage;

  let key = await secureStorage.get(SecureKeys.MmkvEncryptionKey);
  if (!key) {
    key = generateEncryptionKey();
    await secureStorage.set(SecureKeys.MmkvEncryptionKey, key);
  }

  cachedStorage = new MMKV({ id: 'plenya-pro', encryptionKey: key });
  return cachedStorage;
}

export function mmkvToQueryPersisterStorage(storage: MMKV) {
  return {
    getItem: (key: string) => storage.getString(key) ?? null,
    setItem: (key: string, value: string) => storage.set(key, value),
    removeItem: (key: string) => storage.delete(key),
  };
}
