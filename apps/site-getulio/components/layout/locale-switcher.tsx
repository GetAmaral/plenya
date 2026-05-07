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

  const baseInactive = tone === 'dark' ? 'text-paper/50 hover:text-paper' : 'text-ink-muted hover:text-ink';
  const baseActive = tone === 'dark' ? 'text-paper underline underline-offset-4 decoration-gold decoration-2' : 'text-bordo underline underline-offset-4 decoration-bordo decoration-2';

  return (
    <div className="flex items-center gap-3 font-sans text-[11px] uppercase tracking-widest">
      {locales.map((l) => (
        <button
          key={l}
          disabled={isPending}
          onClick={() => switchTo(l)}
          className={l === current ? baseActive : baseInactive}
          aria-label={`Switch to ${l.toUpperCase()}`}
        >
          {localeFlags[l]}
        </button>
      ))}
    </div>
  );
}
