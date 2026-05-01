'use client';

import { useTransition } from 'react';
import { useLocale } from 'next-intl';
import { useParams } from 'next/navigation';
import { usePathname, useRouter } from '@/lib/i18n/navigation';
import { locales, localeFlags, type Locale } from '@/lib/i18n/config';

export function LocaleSwitcher() {
  const router = useRouter();
  const pathname = usePathname();
  const params = useParams();
  const current = useLocale() as Locale;
  const [isPending, startTransition] = useTransition();

  // Não renderizar quando há apenas um idioma disponível.
  if (locales.length <= 1) return null;

  function switchTo(locale: Locale) {
    startTransition(() => {
      // Cast: pathname pode ser uma rota dinâmica como '/blog/[slug]';
      // os params atuais já fornecem os segmentos necessários.
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
      router.replace({ pathname, params: params as any } as any, { locale });
    });
  }

  return (
    <div className="flex items-center gap-3 font-mono text-[11px] uppercase tracking-widest">
      {locales.map((l) => (
        <button
          key={l}
          disabled={isPending}
          onClick={() => switchTo(l)}
          className={
            l === current
              ? 'text-cream underline underline-offset-4 decoration-gold decoration-2'
              : 'text-cream/50 hover:text-cream'
          }
          aria-label={`Trocar para ${l.toUpperCase()}`}
        >
          {localeFlags[l]}
        </button>
      ))}
    </div>
  );
}
