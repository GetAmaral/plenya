import JailMonkey from 'jail-monkey';

export interface DeviceIntegrityReport {
  jailBroken: boolean;
  onExternalStorage: boolean;
  canMockLocation: boolean;
  isDebuggedMode: boolean;
  isDevelopmentSettingsMode: boolean;
  adbEnabled: boolean;
  trustworthy: boolean;
}

export async function checkDeviceIntegrity(): Promise<DeviceIntegrityReport> {
  const jailBroken = JailMonkey.isJailBroken();
  const onExternalStorage = JailMonkey.isOnExternalStorage();
  const canMockLocation = await JailMonkey.canMockLocation();
  const isDebuggedMode = await JailMonkey.isDebuggedMode();
  const isDevelopmentSettingsMode = await JailMonkey.isDevelopmentSettingsMode();
  const adbEnabled = JailMonkey.AdbEnabled();

  const trustworthy = !jailBroken && !onExternalStorage;

  return {
    jailBroken,
    onExternalStorage,
    canMockLocation,
    isDebuggedMode,
    isDevelopmentSettingsMode,
    adbEnabled,
    trustworthy,
  };
}
