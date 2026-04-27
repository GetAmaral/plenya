import type { Metadata } from 'next';
import Script from 'next/script';
import { Source_Serif_4, Inter } from 'next/font/google';
import { NextIntlClientProvider } from 'next-intl';
import { getMessages, setRequestLocale } from 'next-intl/server';
import { notFound } from 'next/navigation';
import { isLocale, locales } from '@/lib/i18n/config';
import { SiteHeader } from '@/components/layout/site-header';
import { SiteFooter } from '@/components/layout/site-footer';
import { PersonSchema } from '@/components/seo/person-schema';
import '../globals.css';

const PLAUSIBLE_DOMAIN =
  process.env.NEXT_PUBLIC_PLAUSIBLE_DOMAIN || 'drgetulioamaralfilho.com.br';

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
    'Nefrologista (CRM-PR 21.876 · RQE 16.038), médico de medicina funcional integrativa, professor, coordenador de residência, autor e diretor clínico da Plenya. Londrina-PR.',
  authors: [{ name: 'Dr. Getúlio Amaral Filho', url: 'https://drgetulioamaralfilho.com.br' }],
  creator: 'Dr. Getúlio Amaral Filho',
  publisher: 'Dr. Getúlio Amaral Filho',
  alternates: { canonical: '/' },
  openGraph: {
    type: 'website',
    locale: 'pt_BR',
    url: 'https://drgetulioamaralfilho.com.br',
    siteName: 'Dr. Getúlio Amaral Filho',
    images: [
      {
        url: '/images/getulio-square.jpg',
        width: 1200,
        height: 1200,
        alt: 'Dr. Getúlio Amaral Filho',
      },
    ],
  },
  twitter: {
    card: 'summary_large_image',
    title: 'Dr. Getúlio Amaral Filho — Nefrologista, autor, professor',
    description:
      'Nefrologista (CRM-PR 21.876 · RQE 16.038), medicina funcional integrativa e diretor clínico da Plenya.',
    images: ['/images/getulio-square.jpg'],
  },
  robots: {
    index: true,
    follow: true,
    googleBot: { index: true, follow: true, 'max-image-preview': 'large', 'max-snippet': -1 },
  },
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
      <head>
        <Script
          defer
          data-domain={PLAUSIBLE_DOMAIN}
          src="https://plausible.io/js/script.js"
          strategy="afterInteractive"
        />
      </head>
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
