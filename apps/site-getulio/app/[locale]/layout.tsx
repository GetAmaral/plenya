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
import { NavigationSchema } from '@/components/seo/navigation-schema';
import '../globals.css';

const PLAUSIBLE_DOMAIN =
  process.env.NEXT_PUBLIC_PLAUSIBLE_DOMAIN || 'drgetulioamaralfilho.com.br';

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

const META_BY_LOCALE = {
  pt: {
    title: 'Dr. Getúlio Amaral Filho — Medicina guiada por raciocínio clínico',
    template: '%s · Dr. Getúlio Amaral',
    description:
      'Medicina guiada por raciocínio clínico. Nefrologista (CRM-PR 21.876 · RQE 16.038), professor, autor e diretor clínico da Plenya. Londrina-PR.',
    twitterDescription:
      'Nefrologista (CRM-PR 21.876 · RQE 16.038), medicina funcional integrativa, autor do livro ANTES.',
    ogLocale: 'pt_BR',
    ogAlt: 'Dr. Getúlio Amaral · Medicina guiada por raciocínio clínico',
  },
  en: {
    title: 'Dr. Getúlio Amaral Filho — Medicine guided by clinical reasoning',
    template: '%s · Dr. Getúlio Amaral',
    description:
      'Medicine guided by clinical reasoning. Nephrologist (CRM-PR 21,876 · RQE 16,038), professor, author and clinical director at Plenya. Londrina, Brazil.',
    twitterDescription:
      'Nephrologist (CRM-PR 21,876 · RQE 16,038), integrative functional medicine, author of the book ANTES.',
    ogLocale: 'en_US',
    ogAlt: 'Dr. Getúlio Amaral · Medicine guided by clinical reasoning',
  },
} as const;

export async function generateMetadata({
  params,
}: {
  params: Promise<{ locale: string }>;
}): Promise<Metadata> {
  const { locale } = await params;
  const meta = isLocale(locale) ? META_BY_LOCALE[locale] : META_BY_LOCALE.pt;
  const path = locale === 'en' ? '/en' : '/';

  return {
    metadataBase: new URL('https://drgetulioamaralfilho.com.br'),
    title: { default: meta.title, template: meta.template },
    description: meta.description,
    authors: [{ name: 'Dr. Getúlio Amaral Filho', url: 'https://drgetulioamaralfilho.com.br' }],
    creator: 'Dr. Getúlio Amaral Filho',
    publisher: 'Dr. Getúlio Amaral Filho',
    alternates: {
      canonical: path,
      languages: {
        'pt-BR': '/',
        en: '/en',
      },
    },
    openGraph: {
      type: 'website',
      locale: meta.ogLocale,
      alternateLocale: locale === 'en' ? ['pt_BR'] : ['en_US'],
      url: `https://drgetulioamaralfilho.com.br${path === '/' ? '' : path}`,
      siteName: 'Dr. Getúlio Amaral Filho',
      images: [
        { url: '/images/og-default.webp', width: 1024, height: 1024, alt: meta.ogAlt },
      ],
    },
    twitter: {
      card: 'summary_large_image',
      title: meta.title,
      description: meta.twitterDescription,
      images: ['/images/og-default.webp'],
    },
    robots: {
      index: true,
      follow: true,
      googleBot: { index: true, follow: true, 'max-image-preview': 'large', 'max-snippet': -1 },
    },
    icons: {
      icon: [
        { url: '/favicon.svg', type: 'image/svg+xml' },
        { url: '/favicon-32x32.png', sizes: '32x32', type: 'image/png' },
        { url: '/favicon-16x16.png', sizes: '16x16', type: 'image/png' },
      ],
      apple: [{ url: '/apple-touch-icon.png', sizes: '180x180', type: 'image/png' }],
      other: [{ rel: 'mask-icon', url: '/favicon.svg', color: '#0E2A2E' }],
    },
    manifest: '/manifest.webmanifest',
    verification: {
      google: process.env.GOOGLE_SITE_VERIFICATION || undefined,
      other: {
        'msvalidate.01': process.env.BING_SITE_VERIFICATION || '',
      },
    },
  };
}

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
    <html lang={locale === 'en' ? 'en' : 'pt-BR'} className={`${serif.variable} ${sans.variable}`}>
      <head>
        <meta name="geo.region" content="BR-PR" />
        <meta name="geo.placename" content="Londrina" />
        <meta name="ICBM" content="-23.3045, -51.1696" />
        <link
          rel="alternate"
          type="application/rss+xml"
          title="Dr. Getúlio Amaral · Artigos"
          href="https://drgetulioamaralfilho.com.br/artigos.xml"
        />
        <Script
          defer
          data-domain={PLAUSIBLE_DOMAIN}
          src="https://plausible.io/js/script.js"
          strategy="afterInteractive"
        />
      </head>
      <body>
        <NextIntlClientProvider messages={messages}>
          <PersonSchema locale={locale} />
          <WebSiteSchema locale={locale} />
          <NavigationSchema locale={locale} />
          <SiteHeader />
          <main>{children}</main>
          <SiteFooter />
        </NextIntlClientProvider>
      </body>
    </html>
  );
}
