import type { Metadata } from 'next';
import { setRequestLocale } from 'next-intl/server';
import { Hero } from '@/components/home/hero';
import { Pillars } from '@/components/home/pillars';
import { BookStrip } from '@/components/home/book-strip';
import { ClinicsRow } from '@/components/home/clinics-row';
import { RecentArticles } from '@/components/home/recent-articles';

export const metadata: Metadata = {
  title: 'Dr. Getúlio Amaral Filho — Nefrologista e Medicina Funcional Integrativa em Londrina',
  description:
    'Médico nefrologista (CRM-PR 21.876 · RQE 16.038), professor, autor do livro "Antes — A Janela Silenciosa". Direção clínica da Plenya. Atendimento em Londrina-PR.',
  alternates: { canonical: '/' },
  openGraph: {
    title: 'Dr. Getúlio Amaral Filho — Nefrologia, longevidade, ensino',
    description:
      'Nefrologista, professor, autor. Direção clínica da Plenya. Vinte anos de prática em medicina interna e nefrologia.',
    type: 'profile',
    images: ['/images/getulio-square.jpg'],
  },
};

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
