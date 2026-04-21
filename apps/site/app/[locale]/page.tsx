import { setRequestLocale } from 'next-intl/server';
import { HomeHero } from '@/components/marketing/home-hero';
import { SymbolBridge } from '@/components/marketing/symbol-bridge';
import { LifestyleGrid } from '@/components/marketing/lifestyle-grid';
import { DrGetulioPreview } from '@/components/marketing/dr-getulio-preview';
import { TeamPreview } from '@/components/marketing/team-preview';
import { AgirPillarsSection } from '@/components/marketing/agir-pillars-section';
import { EstruturaSection } from '@/components/marketing/estrutura-section';
import { ScoreSection } from '@/components/marketing/score-section';
import { ContinuumSpotlight } from '@/components/marketing/continuum-spotlight';
import { PlansPreview } from '@/components/marketing/plans-preview';
import { DiagnosticoStrip } from '@/components/marketing/diagnostico-strip';

export default async function HomePage({ params }: { params: Promise<{ locale: string }> }) {
  const { locale } = await params;
  setRequestLocale(locale);
  return (
    <>
      {/* Abertura — desejo do visitante */}
      <HomeHero />
      <SymbolBridge />
      <LifestyleGrid />

      {/* Manifesto — o que nos move */}
      <EstruturaSection />

      {/* Como cuidamos — primeiro a leitura, depois o método */}
      <ScoreSection />
      <AgirPillarsSection />

      {/* Quem cuida — direção clínica e equipe */}
      <DrGetulioPreview />
      <TeamPreview />

      {/* Produto principal — Continuum em destaque */}
      <ContinuumSpotlight />

      {/* Como começar — comparativo Consultas vs Continuum */}
      <PlansPreview />

      {/* Qualificador final — quiz */}
      <DiagnosticoStrip />
    </>
  );
}
