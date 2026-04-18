'use client';

import { useEffect, useState } from 'react';
import { useTranslations } from 'next-intl';
import { Menu, X } from 'lucide-react';
import { brand } from '@plenya/brand';
import { PlenyaWordmark } from '@plenya/brand/logo';
import { Link, usePathname } from '@/lib/i18n/navigation';
import { LocaleSwitcher } from './locale-switcher';
import { cn } from '@/lib/cn';

const headerNav = [
  { href: '/', key: 'home' },
  { href: '/a-plenya', key: 'about' },
  { href: '/planos', key: 'plans' },
  { href: '/contato', key: 'contact' },
] as const;

const fullNav = [
  { href: '/', key: 'home' },
  { href: '/a-plenya', key: 'about' },
  { href: '/dr-getulio', key: 'drGetulio' },
  { href: '/equipe', key: 'team' },
  { href: '/metodo-agir', label: 'Método AGIR' },
  { href: '/escore-plenya', label: 'Escore Plenya' },
  { href: '/planos', key: 'plans' },
  { href: '/depoimentos', label: 'Depoimentos' },
  { href: '/blog', key: 'blog' },
  { href: '/contato', key: 'contact' },
] as const;

export function SiteHeader() {
  const t = useTranslations('nav');
  const [open, setOpen] = useState(false);
  const [scrolled, setScrolled] = useState(false);
  const pathname = usePathname();

  useEffect(() => setOpen(false), [pathname]);
  useEffect(() => {
    document.body.style.overflow = open ? 'hidden' : '';
    return () => { document.body.style.overflow = ''; };
  }, [open]);
  useEffect(() => {
    const onScroll = () => setScrolled(window.scrollY > 40);
    onScroll();
    window.addEventListener('scroll', onScroll, { passive: true });
    return () => window.removeEventListener('scroll', onScroll);
  }, []);

  function navLabel(item: { key?: string; label?: string }) {
    if (item.label) return item.label;
    return t(item.key as Parameters<typeof t>[0]);
  }

  return (
    <>
      <header
        className={cn(
          'fixed top-0 left-0 right-0 z-40 transition-all duration-300',
          scrolled && 'bg-petrol/95 backdrop-blur-md shadow-sm',
        )}
      >
        <div className="site-container flex items-center justify-between py-5">
          {/* Wordmark — vector from official brandbook */}
          <Link href="/" aria-label="Plenya" className="block text-cream">
            <PlenyaWordmark className="h-5 w-auto md:h-6" />
          </Link>

          {/* 4 pills — desktop only */}
          <nav className="hidden md:flex items-center gap-3">
            {headerNav.map((item) => {
              const active = item.href === '/' ? pathname === '/' : pathname.startsWith(item.href);
              return (
                <Link
                  key={item.key}
                  href={item.href}
                  className={cn('nav-pill', active && 'bg-cream/15 border-cream/80 text-cream')}
                >
                  {t(item.key)}
                </Link>
              );
            })}
          </nav>

          {/* Hamburger — sempre visível (desktop e mobile) */}
          <button
            aria-label={open ? 'Fechar menu' : 'Menu'}
            aria-expanded={open}
            className="text-cream p-2 hover:text-gold transition"
            onClick={() => setOpen((o) => !o)}
          >
            {open ? <X size={22} /> : <Menu size={22} />}
          </button>
        </div>
      </header>

      {/* Overlay panel — slide from right */}
      {open && (
        <>
          {/* Backdrop */}
          <div
            className="fixed inset-0 z-50 bg-petrol/40 backdrop-blur-sm"
            onClick={() => setOpen(false)}
          />

          {/* Panel */}
          <aside className="fixed top-0 right-0 bottom-0 z-50 w-full max-w-sm bg-petrol shadow-2xl overflow-y-auto">
            <div className="flex items-center justify-between p-6 border-b border-cream/10">
              <PlenyaWordmark className="h-5 w-auto text-cream" />
              <button onClick={() => setOpen(false)} className="text-cream hover:text-gold transition">
                <X size={22} />
              </button>
            </div>

            <nav className="p-8 space-y-1">
              {fullNav.map((item) => {
                const active = item.href === '/' ? pathname === '/' : pathname.startsWith(item.href);
                return (
                  <Link
                    key={item.href}
                    href={item.href}
                    onClick={() => setOpen(false)}
                    className={cn(
                      'block py-3 heading-section text-xl text-cream/80 hover:text-cream transition border-b border-cream/5',
                      active && 'text-gold',
                    )}
                  >
                    {navLabel(item)}
                  </Link>
                );
              })}
            </nav>

            <div className="px-8 pt-4 pb-6 space-y-6 border-t border-cream/10 mt-4">
              <a
                href={brand.appUrl}
                className="btn-gold w-full text-center"
                onClick={() => setOpen(false)}
              >
                {t('patientArea')}
              </a>
              <div className="flex items-center justify-center">
                <LocaleSwitcher />
              </div>
            </div>
          </aside>
        </>
      )}
    </>
  );
}
