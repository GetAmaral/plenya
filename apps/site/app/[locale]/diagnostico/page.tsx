import type { Metadata } from 'next';
import { setRequestLocale } from 'next-intl/server';
import { DiagnosticoQuiz } from '@/components/marketing/diagnostico-quiz';

export const metadata: Metadata = {
  title: 'Diagnóstico Plenya — o Continuum Plenya faz sentido para você?',
  description:
    'Cinco perguntas curtas para entender se o programa de acompanhamento contínuo da Plenya se encaixa no seu momento.',
  alternates: { canonical: '/diagnostico' },
};

export default async function DiagnosticoPage({
  params,
}: {
  params: Promise<{ locale: string }>;
}) {
  const { locale } = await params;
  setRequestLocale(locale);

  return (
    <>
      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-16 md:pt-40 md:pb-20">
          <p className="label-upper text-gold mb-6">Diagnóstico</p>
          <h1 className="heading-hero text-[clamp(2.25rem,5vw,4rem)] text-cream max-w-3xl">
            O Continuum Plenya<br />faz sentido para o seu momento?
          </h1>
          <p className="text-cream/70 text-lg mt-6 max-w-2xl leading-relaxed">
            Cinco perguntas curtas. Sem diagnóstico clínico, sem captura de
            dados. No fim, devolvemos quais retratos do programa mais se parecem
            com o que você descreve — e o caminho para conversar com a equipe.
          </p>
        </div>
      </section>

      <section className="bg-cream">
        <div className="site-container section">
          <DiagnosticoQuiz />
        </div>
      </section>
    </>
  );
}
