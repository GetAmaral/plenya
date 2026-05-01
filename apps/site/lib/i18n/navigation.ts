import { createNavigation } from 'next-intl/navigation';
import { defineRouting } from 'next-intl/routing';
import { locales, defaultLocale } from './config';

// Mapa PT ↔ EN para todas as rotas do site.
// Adicione aqui qualquer rota nova — uma URL traduzida por locale.
export const pathnames = {
  '/': '/',
  '/a-plenya': { pt: '/a-plenya', en: '/about' },
  '/dr-getulio': { pt: '/dr-getulio', en: '/dr-getulio' },
  '/equipe': { pt: '/equipe', en: '/team' },
  '/equipe/[slug]': { pt: '/equipe/[slug]', en: '/team/[slug]' },
  '/depoimentos': { pt: '/depoimentos', en: '/testimonials' },
  '/casos': { pt: '/casos', en: '/cases' },
  '/como-funciona': { pt: '/como-funciona', en: '/how-it-works' },
  '/metodo-agir': { pt: '/metodo-agir', en: '/acts-method' },
  '/escore-plenya': { pt: '/escore-plenya', en: '/plenya-score' },
  '/escore-plenya/avaliar': { pt: '/escore-plenya/avaliar', en: '/plenya-score/take' },
  '/escore-plenya/painel': { pt: '/escore-plenya/painel', en: '/plenya-score/dashboard' },
  '/escore-plenya/resultado/[code]': {
    pt: '/escore-plenya/resultado/[code]',
    en: '/plenya-score/result/[code]',
  },
  '/escore-plenya/claim/[token]': {
    pt: '/escore-plenya/claim/[token]',
    en: '/plenya-score/claim/[token]',
  },
  '/diagnostico': { pt: '/diagnostico', en: '/assessment' },
  '/checkup-longevidade': { pt: '/checkup-longevidade', en: '/longevity-checkup' },
  '/medicina-funcional-integrativa': {
    pt: '/medicina-funcional-integrativa',
    en: '/integrative-functional-medicine',
  },
  '/avaliacao-renal-preventiva': {
    pt: '/avaliacao-renal-preventiva',
    en: '/preventive-kidney-assessment',
  },
  '/healthspan': '/healthspan',
  '/consultas': { pt: '/consultas', en: '/consultations' },
  '/continuum': '/continuum',
  '/contato': { pt: '/contato', en: '/contact' },
  '/boletim': { pt: '/boletim', en: '/newsletter' },
  '/blog': '/blog',
  '/blog/[slug]': '/blog/[slug]',
  '/blog/categoria/[pilar]': { pt: '/blog/categoria/[pilar]', en: '/blog/category/[pilar]' },
  '/privacidade': { pt: '/privacidade', en: '/privacy' },
  '/termos': { pt: '/termos', en: '/terms' },
  '/lgpd': { pt: '/lgpd', en: '/data-rights' },
  '/lgpd/direitos': { pt: '/lgpd/direitos', en: '/data-rights/your-rights' },
  '/lgpd/encarregado': { pt: '/lgpd/encarregado', en: '/data-rights/officer' },
} as const;

export const routing = defineRouting({
  locales,
  defaultLocale,
  localePrefix: 'as-needed',
  pathnames,
});

export const { Link, redirect, usePathname, useRouter, getPathname } = createNavigation(routing);

import type { ComponentProps } from 'react';
/** Tipo do href aceito pelo `<Link>` localizado — use ao tipar arrays de rotas. */
export type Href = ComponentProps<typeof Link>['href'];
