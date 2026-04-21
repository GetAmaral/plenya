import type { Metadata } from 'next';
import { setRequestLocale } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { EscoreLightForm } from '@/components/escore/escore-light-form';
import config from '@/content/data/score-light-config.json';
import type { LightConfig } from '@/lib/score-light/types';

export const metadata: Metadata = {
  title: 'Faça sua avaliação — Escore Plenya Light',
  description:
    'Versão pública e gratuita do Escore Plenya. Responda 30 a 40 perguntas e receba um radar com sua pontuação por pilar AGIR. Anônimo, sem cadastro.',
};

export default async function AvaliarPage({
  params,
}: {
  params: Promise<{ locale: string }>;
}) {
  const { locale } = await params;
  setRequestLocale(locale);

  const cfg = config as LightConfig;

  if (!cfg || cfg.itemCount === 0) {
    return (
      <section className="bg-cream">
        <div className="site-container pt-32 pb-24 md:pt-40 md:pb-32 max-w-2xl space-y-8">
          <p className="label-upper text-gold">Em preparação</p>
          <h1 className="heading-section text-petrol text-3xl md:text-4xl">
            A versão Light do Escore Plenya está sendo finalizada.
          </h1>
          <p className="text-petrol/80 leading-relaxed">
            A equipe clínica está concluindo a curadoria das perguntas que comporão a
            versão pública. Em breve, esta página estará disponível para você fazer
            sua autoavaliação.
          </p>
          <div className="pt-4">
            <Link href="/contato" className="btn-gold">
              Falar com a equipe
            </Link>
          </div>
        </div>
      </section>
    );
  }

  return <EscoreLightForm config={cfg} locale={locale} />;
}
