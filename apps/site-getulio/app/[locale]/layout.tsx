import type { Metadata } from 'next';
import { Source_Serif_4, Inter } from 'next/font/google';
import { NextIntlClientProvider } from 'next-intl';
import { getMessages, setRequestLocale } from 'next-intl/server';
import { notFound } from 'next/navigation';
import { isLocale, locales } from '@/lib/i18n/config';
import { SiteHeader } from '@/components/layout/site-header';
import { SiteFooter } from '@/components/layout/site-footer';
import { PersonSchema } from '@/components/seo/person-schema';
import '../globals.css';

const serif = Source_Serif_4({
  subsets: ['latin'],
  weight: ['300', '400', '500', '600', '700'],
  display: 'swap',
  variable: '--font-serif',
});

const sans = Inter({
  subsets: ['latin'],
  display: 'swap',
  variable: '--font-sans',
});

export const metadata: Metadata = {
  metadataBase: new URL('https://drgetulioamaralfilho.com.br'),
  title: {
    default: 'Dr. Getúlio Amaral Filho — Nefrologista, autor, professor',
    template: '%s · Dr. Getúlio Amaral Filho',
  },
  description:
    'Nefrologista, médico de medicina funcional integrativa, professor, coordenador de residência, autor e diretor clínico da Plenya. Londrina-PR.',
  authors: [{ name: 'Dr. Getúlio Amaral Filho' }],
  openGraph: {
    type: 'website',
    locale: 'pt_BR',
    url: 'https://drgetulioamaralfilho.com.br',
    siteName: 'Dr. Getúlio Amaral Filho',
    images: ['/images/getulio-square.jpg'],
  },
  robots: { index: true, follow: true },
};

export function generateStaticParams() {
  return locales.map((locale) => ({ locale }));
}

export default async function LocaleLayout({
  children,
  params,
}: {
  children: React.ReactNode;
  params: Promise<{ locale: string }>;
}) {
  const { locale } = await params;
  if (!isLocale(locale)) notFound();
  setRequestLocale(locale);
  const messages = await getMessages();

  return (
    <html lang={locale} className={`${serif.variable} ${sans.variable}`}>
      <body>
        <NextIntlClientProvider messages={messages}>
          <PersonSchema />
          <SiteHeader />
          <main>{children}</main>
          <SiteFooter />
        </NextIntlClientProvider>
      </body>
    </html>
  );
}
