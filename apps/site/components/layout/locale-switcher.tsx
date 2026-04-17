'use client';

import { useTransition } from 'react';
import { useLocale } from 'next-intl';
import { usePathname, useRouter } from '@/lib/i18n/navigation';
import { locales, localeFlags, type Locale } from '@/lib/i18n/config';

export function LocaleSwitcher() {
  const router = useRouter();
  const pathname = usePathname();
  const current = useLocale() as Locale;
  const [isPending, startTransition] = useTransition();

  return (
    <div className="flex items-center gap-2 font-mono text-[11px] uppercase tracking-widest">
      {locales.map((l) => (
        <button
          key={l}
          disabled={isPending}
          onClick={() => startTransition(() => router.replace(pathname, { locale: l }))}
          className={
            l === current
              ? 'text-petrol underline underline-offset-4 decoration-gold decoration-2'
              : 'text-petrol/50 hover:text-petrol'
          }
        >
          {localeFlags[l]}
        </button>
      ))}
    </div>
  );
}
