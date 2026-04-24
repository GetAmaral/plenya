'use client';

import { useRouter, useSearchParams, usePathname } from 'next/navigation';
import { useCallback, useMemo } from 'react';

/**
 * Lê e escreve filtros via URL search params (?status=new,contacted&...).
 *
 * - Persiste em refresh / colável.
 * - Não causa scroll na navegação.
 * - Substitui (replace) ao invés de push, pra não poluir histórico.
 *
 * @example
 *   const { params, setParam, setParams, clearAll } = useUrlFilters();
 *   const status = params.get('status')?.split(',') ?? [];
 *   setParam('status', status.length ? status.join(',') : null);
 */
export function useUrlFilters() {
  const router = useRouter();
  const pathname = usePathname();
  const searchParams = useSearchParams();

  const params = useMemo(
    () => new URLSearchParams(searchParams.toString()),
    [searchParams],
  );

  const replace = useCallback(
    (next: URLSearchParams) => {
      const qs = next.toString();
      const url = qs ? `${pathname}?${qs}` : pathname;
      router.replace(url, { scroll: false });
    },
    [router, pathname],
  );

  /** Define um único param. `null`/'' remove. */
  const setParam = useCallback(
    (key: string, value: string | null | undefined) => {
      const next = new URLSearchParams(searchParams.toString());
      if (value === null || value === undefined || value === '') {
        next.delete(key);
      } else {
        next.set(key, value);
      }
      replace(next);
    },
    [searchParams, replace],
  );

  /** Define vários params de uma vez. Atomico — uma única navegação. */
  const setParams = useCallback(
    (updates: Record<string, string | null | undefined>) => {
      const next = new URLSearchParams(searchParams.toString());
      for (const [key, value] of Object.entries(updates)) {
        if (value === null || value === undefined || value === '') {
          next.delete(key);
        } else {
          next.set(key, value);
        }
      }
      replace(next);
    },
    [searchParams, replace],
  );

  const clearAll = useCallback(() => {
    replace(new URLSearchParams());
  }, [replace]);

  return { params, setParam, setParams, clearAll };
}

/** Helpers pra serializar/deserializar arrays separados por vírgula. */
export function parseList(raw: string | null): string[] {
  if (!raw) return [];
  return raw.split(',').filter(Boolean);
}

export function serializeList(values: ReadonlyArray<string>): string | null {
  return values.length ? values.join(',') : null;
}
