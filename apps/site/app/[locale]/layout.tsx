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
import { MedicalClinicSchema } from '@/components/seo/medical-clinic-schema';
import { PhysicianSchema } from '@/components/seo/physician-schema';
import { WebSiteSchema } from '@/components/seo/website-schema';
import { NavigationSchema } from '@/components/seo/navigation-schema';

const heading = localFont({
  src: '../../node_modules/@plenya/brand/src/fonts/nalieta/Nalieta-Regular.otf',
  weight: '400',
  style: 'normal',
  variable: '--font-heading',
  display: 'swap',
  // Nalieta não inclui glyphs de dígitos (0-9). Excluindo o range
  // U+0030-0039 do unicode-range, o browser cai no próximo font da
  // cadeia (Cormorant Garamond) automaticamente, per-glyph.
  declarations: [
    { prop: 'unicode-range', value: 'U+0020-002F, U+003A-FFFF' },
  ],
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
      default: isEn
        ? `${brand.name} · Longevity & Functional Medicine Clinic in Londrina`
        : `${brand.name} · Clínica de Longevidade e Medicina Funcional em Londrina`,
      template: `%s · ${brand.name}`,
    },
    description: isEn
      ? 'Premium functional integrative medicine clinic. Personalized care for health, performance and longevity.'
      : isPt
        ? 'Clínica premium de saúde funcional integrativa. Cuidado personalizado em saúde, performance e longevidade.'
        : 'Clínica premium de medicina funcional integrativa. Cuidado personalizado en salud, rendimiento y longevidad.',
    alternates: {
      canonical: locale === 'pt' ? '/' : `/${locale}`,
      languages: {
        'pt-BR': '/', pt: '/',
        en: '/en',
        'x-default': '/',
      },
    },
    openGraph: {
      type: 'website',
      siteName: brand.name,
      locale: isEn ? 'en_US' : 'pt_BR',
      alternateLocale: isEn ? ['pt_BR'] : ['en_US'],
      images: [{ url: '/og-default.jpg', width: 1200, height: 630, alt: brand.name }],
    },
    twitter: {
      card: 'summary_large_image',
      site: '@plenyaSaude',
      creator: '@plenyaSaude',
      images: ['/og-default.jpg'],
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
        <meta name="geo.region" content="BR-PR" />
        <meta name="geo.placename" content="Londrina" />
        <meta name="ICBM" content="-23.3045, -51.1696" />
        <link
          rel="alternate"
          type="application/rss+xml"
          title="Plenya · Blog"
          href="https://plenyasaude.com.br/blog/rss.xml"
        />
      </head>
      <body className="min-h-screen flex flex-col font-sans">
        <NextIntlClientProvider locale={locale} messages={messages}>
          <SmoothScrollProvider>
            <OrganizationSchema locale={locale} />
            <MedicalClinicSchema locale={locale} />
            <PhysicianSchema locale={locale} />
            <WebSiteSchema locale={locale} />
            <NavigationSchema locale={locale} />
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
