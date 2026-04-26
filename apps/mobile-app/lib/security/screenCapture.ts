import { useEffect } from 'react';
import * as ScreenCapture from 'expo-screen-capture';

/**
 * Blocks screenshots and screen recording while the component is mounted.
 * Apply to every screen that renders patient data (anamnesis, prescriptions,
 * lab results, prontuário, etc.).
 */
export function useScreenCaptureProtection(enabled: boolean = true): void {
  useEffect(() => {
    if (!enabled) return;
    let active = true;
    ScreenCapture.preventScreenCaptureAsync().catch(() => {});
    return () => {
      if (!active) return;
      active = false;
      ScreenCapture.allowScreenCaptureAsync().catch(() => {});
    };
  }, [enabled]);
}
