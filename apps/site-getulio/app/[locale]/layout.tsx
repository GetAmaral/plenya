import type { Metadata } from 'next';
import Script from 'next/script';
import { Cormorant_Garamond, Inter } from 'next/font/google';
import { NextIntlClientProvider } from 'next-intl';
import { getMessages, setRequestLocale } from 'next-intl/server';
import { notFound } from 'next/navigation';
import { isLocale, locales } from '@/lib/i18n/config';
import { SiteHeader } from '@/components/layout/site-header';
import { SiteFooter } from '@/components/layout/site-footer';
import { PersonSchema } from '@/components/seo/person-schema';
import { WebSiteSchema } from '@/components/seo/website-schema';
import '../globals.css';

const PLAUSIBLE_DOMAIN =
  process.env.NEXT_PUBLIC_PLAUSIBLE_DOMAIN || 'drgetulioamaralfilho.com.br';

// Cormorant Garamond — serifa display de alto contraste, fiel ao
// wordmark oficial do branding. Pesos: 300 italic (tagline),
// 400 (body+wordmark), 500/600 (subtítulos).
const serif = Cormorant_Garamond({
  subsets: ['latin'],
  weight: ['300', '400', '500', '600', '700'],
  style: ['normal', 'italic'],
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
    default: 'Dr. Getúlio Amaral Filho — Medicina guiada por raciocínio clínico',
    template: '%s · Dr. Getúlio Amaral',
  },
  description:
    'Medicina guiada por raciocínio clínico. Nefrologista (CRM-PR 21.876 · RQE 16.038), professor, autor e diretor clínico da Plenya. Londrina-PR.',
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
        url: '/images/og-default.webp',
        width: 1024,
        height: 1024,
        alt: 'Dr. Getúlio Amaral · Medicina guiada por raciocínio clínico',
      },
    ],
  },
  twitter: {
    card: 'summary_large_image',
    title: 'Dr. Getúlio Amaral Filho — Medicina guiada por raciocínio clínico',
    description:
      'Nefrologista (CRM-PR 21.876 · RQE 16.038), medicina funcional integrativa, autor do livro ANTES.',
    images: ['/images/og-default.webp'],
  },
  robots: {
    index: true,
    follow: true,
    googleBot: { index: true, follow: true, 'max-image-preview': 'large', 'max-snippet': -1 },
  },
  icons: {
    icon: [{ url: '/favicon.svg', type: 'image/svg+xml' }],
    apple: '/apple-touch-icon.svg',
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
        <meta name="geo.region" content="BR-PR" />
        <meta name="geo.placename" content="Londrina" />
        <meta name="ICBM" content="-23.3045, -51.1696" />
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
          <WebSiteSchema />
          <SiteHeader />
          <main>{children}</main>
          <SiteFooter />
        </NextIntlClientProvider>
      </body>
    </html>
  );
}
