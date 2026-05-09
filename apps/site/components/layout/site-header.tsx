'use client';

import { useEffect, useState } from 'react';
import { useTranslations } from 'next-intl';
import { Menu, X } from 'lucide-react';
import { brand } from '@plenya/brand';
import { PlenyaWordmark } from '@plenya/brand/logo';
import { Link, usePathname, type Href } from '@/lib/i18n/navigation';
import { LocaleSwitcher } from './locale-switcher';
import { cn } from '@/lib/cn';

const headerNav = [
  { href: '/como-funciona', key: 'howItWorks' },
  { href: '/consultas', key: 'consultations' },
  { href: '/continuum', key: 'plans' },
  { href: '/escore-plenya', key: 'score' },
  { href: '/contato', key: 'contact' },
] as const;

type NavItem = { href: Href; key?: string; label?: string };
type NavGroup = { title: string; items: readonly NavItem[] };

// Estrutura constante das chaves de tradução por grupo. Os textos visíveis
// vêm de `messages/{locale}.json` via `useTranslations('navMenu')`.
const navGroups: readonly {
  titleKey: 'groupAbout' | 'groupCare' | 'groupStart' | 'groupLearn';
  items: readonly { href: Href; key?: string; menuKey?: string }[];
}[] = [
  {
    titleKey: 'groupAbout',
    items: [
      { href: '/a-plenya', key: 'about' },
      { href: '/dr-getulio', key: 'drGetulio' },
      { href: '/equipe', key: 'team' },
      { href: '/depoimentos', menuKey: 'testimonials' },
    ],
  },
  {
    titleKey: 'groupCare',
    items: [
      { href: '/como-funciona', key: 'howItWorks' },
      { href: '/metodo-agir', menuKey: 'method' },
      { href: '/escore-plenya', menuKey: 'scoreFull' },
    ],
  },
  {
    titleKey: 'groupStart',
    items: [
      { href: '/diagnostico', menuKey: 'diagnostic' },
      { href: '/consultas', key: 'consultations' },
      { href: '/continuum', key: 'plans' },
      { href: '/contato', key: 'contact' },
    ],
  },
  {
    titleKey: 'groupLearn',
    items: [
      { href: '/blog', key: 'blog' },
      { href: '/casos', menuKey: 'cases' },
      { href: '/boletim', menuKey: 'newsletter' },
    ],
  },
] as const;

export function SiteHeader() {
  const t = useTranslations('nav');
  const tMenu = useTranslations('navMenu');
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

  function navLabel(item: { key?: string; menuKey?: string }) {
    if (item.menuKey) return tMenu(item.menuKey as Parameters<typeof tMenu>[0]);
    return t(item.key as Parameters<typeof t>[0]);
  }

  return (
    <>
      <header
        className={cn(
          'fixed top-0 left-0 right-0 z-40 transition-all duration-300',
          scrolled
            ? 'bg-petrol/95 backdrop-blur-md shadow-sm'
            : 'bg-petrol/30 backdrop-blur-sm',
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
              const active = pathname.startsWith(item.href);
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

          <div className="flex items-center gap-4">
            {/* Locale switcher — discreto à direita das pills */}
            <div className="hidden md:flex">
              <LocaleSwitcher />
            </div>

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

            <nav className="p-8 space-y-8">
              <Link
                href="/"
                onClick={() => setOpen(false)}
                className={cn(
                  'block py-3 heading-section text-xl text-cream/80 hover:text-cream transition border-b border-cream/5',
                  pathname === '/' && 'text-gold',
                )}
              >
                {t('home')}
              </Link>

              {navGroups.map((group) => (
                <div key={group.titleKey} className="space-y-1">
                  <p className="label-upper text-gold mb-3">{tMenu(group.titleKey)}</p>
                  {group.items.map((item) => {
                    const itemPath = typeof item.href === 'string' ? item.href : item.href.pathname;
                    const active = pathname.startsWith(itemPath);
                    return (
                      <Link
                        key={itemPath}
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
                </div>
              ))}
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
