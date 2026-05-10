import type { Metadata } from 'next';
import { setRequestLocale, getTranslations } from 'next-intl/server';
import { Hero } from '@/components/home/hero';
import { CareManager } from '@/components/home/care-manager';
import { Pillars } from '@/components/home/pillars';
import { BookStrip } from '@/components/home/book-strip';
import { ClinicsRow } from '@/components/home/clinics-row';
import { RecentArticles } from '@/components/home/recent-articles';

export async function generateMetadata({
  params,
}: {
  params: Promise<{ locale: string }>;
}): Promise<Metadata> {
  const { locale } = await params;
  const t = await getTranslations({ locale, namespace: 'home' });
  return {
    title: t('metaTitle'),
    description: t('metaDescription'),
    alternates: {
      canonical: locale === 'en' ? '/en' : '/',
      languages: { 'pt-BR': '/', pt: '/', en: '/en' },
    },
    openGraph: {
      title: t('ogTitle'),
      description: t('ogDescription'),
      type: 'profile',
      images: ['/images/getulio-square.jpg'],
    },
  };
}

export default async function HomePage({
  params,
}: {
  params: Promise<{ locale: string }>;
}) {
  const { locale } = await params;
  setRequestLocale(locale);

  return (
    <>
      <Hero />
      <CareManager />
      <Pillars />
      <BookStrip />
      <ClinicsRow />
      <RecentArticles />
    </>
  );
}
