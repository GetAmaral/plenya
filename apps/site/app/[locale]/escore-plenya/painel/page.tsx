import type { Metadata } from 'next';
import { setRequestLocale } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { EscoreLightForm } from '@/components/escore/escore-light-form';
import config from '@/content/data/score-light-config.json';
import type { LightConfig } from '@/lib/score-light/types';

export const metadata: Metadata = {
  title: 'Painel ampliado — Escore Plenya',
  description:
    'Versão expandida do Escore Plenya. Painel com mais de 80 itens — exames, sintomas, hábitos e marcadores — para uma leitura aprofundada antes da consulta.',
};

export default async function PainelPage({
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
            O Painel do Escore Plenya está sendo finalizado.
          </h1>
          <p className="text-petrol/80 leading-relaxed">
            A equipe clínica está concluindo a curadoria do painel ampliado. Em breve, esta
            página estará disponível.
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

  return <EscoreLightForm config={cfg} locale={locale} tierLabel="Escore Plenya Painel" />;
}
