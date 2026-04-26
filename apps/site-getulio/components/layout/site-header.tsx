'use client';

import { useState } from 'react';
import { useTranslations } from 'next-intl';
import { Menu, X } from 'lucide-react';
import { Link, usePathname } from '@/lib/i18n/navigation';
import { cn } from '@/lib/cn';

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
  const pathname = usePathname();
  const [open, setOpen] = useState(false);

  return (
    <header className="border-b border-rule">
      <div className="editorial-container flex items-center justify-between py-6">
        <Link href="/" className="font-serif text-lg text-ink hover:text-bordo transition-colors">
          Dr. Getúlio Amaral Filho
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
        </nav>

        {/* Mobile toggle */}
        <button
          aria-label={open ? 'Fechar menu' : 'Menu'}
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
          </div>
        </nav>
      )}
    </header>
  );
}
