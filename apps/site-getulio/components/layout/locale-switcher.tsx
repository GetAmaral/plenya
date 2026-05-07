'use client';

import { useTransition } from 'react';
import { useLocale } from 'next-intl';
import { useParams } from 'next/navigation';
import { usePathname, useRouter } from '@/lib/i18n/navigation';
import { locales, localeFlags, type Locale } from '@/lib/i18n/config';

export function LocaleSwitcher({ tone = 'light' }: { tone?: 'light' | 'dark' }) {
  const router = useRouter();
  const pathname = usePathname();
  const params = useParams();
  const current = useLocale() as Locale;
  const [isPending, startTransition] = useTransition();

  if (locales.length <= 1) return null;

  function switchTo(locale: Locale) {
    startTransition(() => {
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      router.replace({ pathname, params: params as any } as any, { locale });
    });
  }

  const inactive =
    tone === 'dark'
      ? 'text-paper/55 hover:text-paper'
      : 'text-ink-muted hover:text-ink';
  const active =
    tone === 'dark'
      ? 'text-paper'
      : 'text-bordo';

  return (
    <div
      className="flex items-center gap-2 font-sans text-[11px] uppercase tracking-widest"
      aria-label="Language"
    >
      {locales.map((l, i) => (
        <span key={l} className="flex items-center gap-2">
          {i > 0 && <span aria-hidden="true" className={tone === 'dark' ? 'text-paper/30' : 'text-ink-muted/40'}>·</span>}
          <button
            disabled={isPending}
            onClick={() => switchTo(l)}
            className={`px-1.5 py-0.5 rounded-sm transition-colors ${
              l === current
                ? `${active} font-semibold ring-1 ring-current/20`
                : inactive
            }`}
            aria-label={`Switch to ${l.toUpperCase()}`}
            aria-current={l === current ? 'true' : undefined}
          >
            {localeFlags[l]}
          </button>
        </span>
      ))}
    </div>
  );
}
