export const locales = ['pt', 'en', 'es'] as const;
export type Locale = (typeof locales)[number];

export const defaultLocale: Locale = 'pt';

export const localeLabels: Record<Locale, string> = {
  pt: 'Português',
  en: 'English',
  es: 'Español',
};

export const localeFlags: Record<Locale, string> = {
  pt: 'BR',
  en: 'EN',
  es: 'ES',
};

export function isLocale(value: string): value is Locale {
  return (locales as readonly string[]).includes(value);
}
