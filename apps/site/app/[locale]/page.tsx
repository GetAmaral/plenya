import { setRequestLocale } from 'next-intl/server';
import { HomeHero } from '@/components/marketing/home-hero';
import { SymbolBridge } from '@/components/marketing/symbol-bridge';
import { LifestyleGrid } from '@/components/marketing/lifestyle-grid';
import { DrGetulioPreview } from '@/components/marketing/dr-getulio-preview';
import { AgirPillarsSection } from '@/components/marketing/agir-pillars-section';
import { ScoreSection } from '@/components/marketing/score-section';
import { PlansPreview } from '@/components/marketing/plans-preview';
import { TestimonialsSection } from '@/components/testimonials/testimonials-section';

export default async function HomePage({ params }: { params: Promise<{ locale: string }> }) {
  const { locale } = await params;
  setRequestLocale(locale);
  return (
    <>
      <HomeHero />
      <SymbolBridge />
      <LifestyleGrid />
      <DrGetulioPreview />
      <AgirPillarsSection />
      <ScoreSection />
      <TestimonialsSection />
      <PlansPreview />
    </>
  );
}
