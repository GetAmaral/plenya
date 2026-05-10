export const locales = ['pt', 'en'] as const;
export type Locale = (typeof locales)[number];

export const defaultLocale: Locale = 'pt';

export const localeLabels: Record<Locale, string> = {
  pt: 'Português',
  en: 'English',
};

export const localeFlags: Record<Locale, string> = {
  pt: 'PT',
  en: 'EN',
};

export function isLocale(value: string): value is Locale {
  return (locales as readonly string[]).includes(value);
}

// Mapeamento de URL por locale (next-intl pathnames).
// CHAVE = caminho canônico interno (= nome da pasta em app/[locale]/...).
// Esses são os hrefs que o código usa no <Link> e nos esquemas internos.
// VALOR = URL pública renderizada por locale.
//
// Regras:
//   - Folders renomeados: /o-medico, /artigos (PT slug = canônico)
//   - EN traduz TUDO (sem misturar PT em URL EN)
//   - Slugs dinâmicos preservam [slug] em ambos os locais
export const pathnames = {
  '/': '/',
  '/o-medico': { pt: '/o-medico', en: '/the-physician' },
  '/livros': { pt: '/livros', en: '/books' },
  '/livros/[slug]': { pt: '/livros/[slug]', en: '/books/[slug]' },
  '/livros/[slug]/excertos': {
    pt: '/livros/[slug]/excertos',
    en: '/books/[slug]/excerpts',
  },
  '/palestras': { pt: '/palestras', en: '/lectures' },
  '/palestras/[slug]': { pt: '/palestras/[slug]', en: '/lectures/[slug]' },
  '/ensino': { pt: '/ensino', en: '/teaching' },
  '/onde-atendo': { pt: '/onde-atendo', en: '/where-i-practice' },
  '/artigos': { pt: '/artigos', en: '/articles' },
  '/artigos/[slug]': { pt: '/artigos/[slug]', en: '/articles/[slug]' },
  '/contato': { pt: '/contato', en: '/contact' },
} as const;

export type AppPathname = keyof typeof pathnames;
