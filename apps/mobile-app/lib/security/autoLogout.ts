import { AppState, AppStateStatus } from 'react-native';
import { useEffect, useRef } from 'react';

// App paciente: 30min ociosidade (vs 5min do mobile-pro). Justificativa: paciente
// não manipula dados de outros pacientes — janela maior reduz fricção.
const DEFAULT_TIMEOUT_MS = 30 * 60 * 1000; // 30 minutes

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
