import { createContext, useCallback, useContext, useMemo, useState } from 'react';
import { Text, View } from 'react-native';
import { cn } from './cn';

export type ToastTone = 'default' | 'success' | 'error' | 'warning';

export interface ToastPayload {
  id: string;
  message: string;
  tone?: ToastTone;
}

interface ToastContextValue {
  show: (message: string, tone?: ToastTone) => void;
}

const ToastContext = createContext<ToastContextValue | null>(null);

const toneClass: Record<ToastTone, string> = {
  default: 'bg-foreground',
  success: 'bg-emerald-600',
  error: 'bg-destructive',
  warning: 'bg-amber-500',
};

export function ToastProvider({ children }: { children: React.ReactNode }) {
  const [toasts, setToasts] = useState<ToastPayload[]>([]);

  const show = useCallback((message: string, tone: ToastTone = 'default') => {
    const id = `${Date.now()}-${Math.random()}`;
    setToasts((prev) => [...prev, { id, message, tone }]);
    setTimeout(() => {
      setToasts((prev) => prev.filter((t) => t.id !== id));
    }, 3200);
  }, []);

  const value = useMemo(() => ({ show }), [show]);

  return (
    <ToastContext.Provider value={value}>
      {children}
      <View pointerEvents="none" className="absolute inset-x-0 top-12 items-center gap-2 px-4">
        {toasts.map((t) => (
          <View
            key={t.id}
            className={cn('rounded-full px-4 py-2 shadow-lg', toneClass[t.tone ?? 'default'])}
          >
            <Text className="text-sm font-medium text-white">{t.message}</Text>
          </View>
        ))}
      </View>
    </ToastContext.Provider>
  );
}

export function useToast(): ToastContextValue {
  const ctx = useContext(ToastContext);
  if (!ctx) throw new Error('useToast must be used within a ToastProvider');
  return ctx;
}
