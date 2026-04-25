import { AppState, AppStateStatus } from 'react-native';
import { useEffect, useRef } from 'react';

const DEFAULT_TIMEOUT_MS = 5 * 60 * 1000; // 5 minutes

export function useAutoLogout(onTimeout: () => void, timeoutMs: number = DEFAULT_TIMEOUT_MS) {
  const lastActive = useRef(Date.now());
  const timer = useRef<ReturnType<typeof setTimeout> | null>(null);

  useEffect(() => {
    const reset = () => {
      lastActive.current = Date.now();
      if (timer.current) clearTimeout(timer.current);
      timer.current = setTimeout(() => {
        onTimeout();
      }, timeoutMs);
    };

    const handleAppState = (status: AppStateStatus) => {
      if (status === 'active') {
        const idle = Date.now() - lastActive.current;
        if (idle >= timeoutMs) {
          onTimeout();
        } else {
          reset();
        }
      } else if (status === 'background' || status === 'inactive') {
        lastActive.current = Date.now();
      }
    };

    const sub = AppState.addEventListener('change', handleAppState);
    reset();
    return () => {
      sub.remove();
      if (timer.current) clearTimeout(timer.current);
    };
  }, [onTimeout, timeoutMs]);
}
