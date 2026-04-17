import type { Metadata } from 'next';
import { NextIntlClientProvider, hasLocale } from 'next-intl';
import { getMessages, setRequestLocale } from 'next-intl/server';
import { notFound } from 'next/navigation';
import localFont from 'next/font/local';
import { IBM_Plex_Mono, Inter } from 'next/font/google';
import { brand } from '@plenya/brand';
import { locales } from '@/lib/i18n/config';
import { SiteHeader } from '@/components/layout/site-header';
import { SiteFooter } from '@/components/layout/site-footer';
import { SmoothScrollProvider } from '@/components/layout/smooth-scroll-provider';
import { WhatsAppBubble } from '@/components/marketing/whatsapp-bubble';
import { OrganizationSchema } from '@/components/seo/organization-schema';

const heading = localFont({
  src: '../../node_modules/@plenya/brand/src/fonts/nalieta/Nalieta-Regular.otf',
  weight: '400',
  style: 'normal',
  variable: '--font-heading',
  display: 'swap',
});
const mono = IBM_Plex_Mono({ subsets: ['latin'], weight: ['400', '500'], variable: '--font-mono', display: 'swap' });
const body = Inter({ subsets: ['latin'], variable: '--font-sans', display: 'swap' });

export function generateStaticParams() {
  return locales.map((locale) => ({ locale }));
}

export async function generateMetadata({ params }: { params: Promise<{ locale: string }> }): Promise<Metadata> {
  const { locale } = await params;
  const isPt = locale === 'pt';
  const isEn = locale === 'en';
  return {
    title: {
      default: `${brand.name} — ${isEn ? 'Health, Performance & Longevity' : brand.tagline}`,
      template: `%s | ${brand.name}`,
    },
    description: isEn
      ? 'Premium functional integrative medicine clinic. Personalized care for health, performance and longevity.'
      : isPt
        ? 'Clínica premium de saúde funcional integrativa. Cuidado personalizado em saúde, performance e longevidade.'
        : 'Clínica premium de medicina funcional integrativa. Cuidado personalizado en salud, rendimiento y longevidad.',
    alternates: {
      canonical: `/${locale === 'pt' ? '' : locale}`,
      languages: {
        'pt-BR': '/',
        en: '/en',
        es: '/es',
      },
    },
    openGraph: {
      type: 'website',
      siteName: brand.name,
      locale: locale === 'pt' ? 'pt_BR' : locale === 'en' ? 'en_US' : 'es_ES',
    },
    icons: {
      icon: [{ url: '/favicon.svg', type: 'image/svg+xml' }],
      apple: '/apple-touch-icon.svg',
    },
  };
}

export default async function LocaleLayout({
  children,
  params,
}: {
  children: React.ReactNode;
  params: Promise<{ locale: string }>;
}) {
  const { locale } = await params;
  if (!hasLocale(locales, locale)) notFound();
  setRequestLocale(locale);
  const messages = await getMessages();

  return (
    <html
      lang={locale}
      className={`${heading.variable} ${mono.variable} ${body.variable}`}
      suppressHydrationWarning
    >
      <head>
        <meta name="apple-itunes-app" content="app-id=000000000" />
      </head>
      <body className="min-h-screen flex flex-col font-sans">
        <NextIntlClientProvider locale={locale} messages={messages}>
          <SmoothScrollProvider>
            <OrganizationSchema />
            <SiteHeader />
            <main className="flex-1">{children}</main>
            <SiteFooter />
            <WhatsAppBubble />
          </SmoothScrollProvider>
        </NextIntlClientProvider>
      </body>
    </html>
  );
}
