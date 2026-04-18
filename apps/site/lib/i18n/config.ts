// EN/ES temporariamente fora do ar — copy de marca ainda não traduzido.
// Quando o conteúdo institucional for traduzido (manifesto, AGIR pillars,
// propósito, etc.), reintroduzir 'en' e 'es' aqui.
export const locales = ['pt'] as const;
export type Locale = (typeof locales)[number];

export const defaultLocale: Locale = 'pt';

export const localeLabels: Record<Locale, string> = {
  pt: 'Português',
};

export const localeFlags: Record<Locale, string> = {
  pt: 'BR',
};

export function isLocale(value: string): value is Locale {
  return (locales as readonly string[]).includes(value);
}
