import { setRequestLocale } from 'next-intl/server';
import { Hero } from '@/components/home/hero';
import { Pillars } from '@/components/home/pillars';
import { BookStrip } from '@/components/home/book-strip';
import { ClinicsRow } from '@/components/home/clinics-row';
import { RecentArticles } from '@/components/home/recent-articles';

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
      <Pillars />
      <BookStrip />
      <ClinicsRow />
      <RecentArticles />
    </>
  );
}
