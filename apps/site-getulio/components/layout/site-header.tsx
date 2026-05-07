'use client';

import { useState } from 'react';
import { useTranslations } from 'next-intl';
import { Menu, X } from 'lucide-react';
import { Link, usePathname } from '@/lib/i18n/navigation';
import { cn } from '@/lib/cn';
import { Wordmark } from '@/components/brand/wordmark';
import { LocaleSwitcher } from '@/components/layout/locale-switcher';

const nav = [
  { href: '/sobre', key: 'sobre' },
  { href: '/livro', key: 'livro' },
  { href: '/palestras', key: 'palestras' },
  { href: '/ensino', key: 'ensino' },
  { href: '/onde-atendo', key: 'ondeAtendo' },
  { href: '/escritos', key: 'escritos' },
  { href: '/contato', key: 'contato' },
] as const;

export function SiteHeader() {
  const t = useTranslations('nav');
  const tHeader = useTranslations('header');
  const pathname = usePathname();
  const [open, setOpen] = useState(false);

  return (
    <header className="border-b border-rule">
      <div className="editorial-container flex items-center justify-between py-5">
        <Link href="/" aria-label={tHeader('homeLabel')} className="block hover:opacity-80 transition-opacity">
          <Wordmark size="sm" tagline={false} />
        </Link>

        {/* Desktop — links de texto puros */}
        <nav className="hidden md:flex items-center gap-7">
          {nav.map((item) => {
            const active = pathname.startsWith(item.href);
            return (
              <Link
                key={item.key}
                href={item.href}
                className={cn(
                  'font-sans text-sm tracking-wide text-ink-muted hover:text-ink transition-colors',
                  active && 'text-bordo',
                )}
              >
                {t(item.key)}
              </Link>
            );
          })}
          <LocaleSwitcher />
        </nav>

        {/* Mobile toggle */}
        <button
          aria-label={open ? tHeader('menuClose') : tHeader('menuOpen')}
          aria-expanded={open}
          className="md:hidden text-ink p-2"
          onClick={() => setOpen((o) => !o)}
        >
          {open ? <X size={20} /> : <Menu size={20} />}
        </button>
      </div>

      {/* Mobile menu */}
      {open && (
        <nav className="md:hidden border-t border-rule bg-paper">
          <div className="editorial-container py-6 flex flex-col gap-4">
            {nav.map((item) => (
              <Link
                key={item.key}
                href={item.href}
                onClick={() => setOpen(false)}
                className="font-serif text-lg text-ink-soft hover:text-bordo transition-colors"
              >
                {t(item.key)}
              </Link>
            ))}
            <div className="pt-4 border-t border-rule">
              <LocaleSwitcher />
            </div>
          </div>
        </nav>
      )}
    </header>
  );
}
