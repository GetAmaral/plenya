'use client';

import { useEffect, useState } from 'react';

/**
 * Retorna o valor depois de `delayMs` ms sem mudar — útil pra debounce
 * de inputs de busca antes de disparar fetches.
 *
 * @example
 *   const [search, setSearch] = useState('');
 *   const debounced = useDebouncedValue(search, 300);
 *   // use `debounced` no queryKey/queryFn
 */
export function useDebouncedValue<T>(value: T, delayMs = 300): T {
  const [debounced, setDebounced] = useState(value);

  useEffect(() => {
    const handle = window.setTimeout(() => setDebounced(value), delayMs);
    return () => window.clearTimeout(handle);
  }, [value, delayMs]);

  return debounced;
}
