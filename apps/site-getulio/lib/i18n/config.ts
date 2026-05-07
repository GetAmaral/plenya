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
