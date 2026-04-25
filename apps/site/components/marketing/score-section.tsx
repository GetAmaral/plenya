import { useTranslations } from 'next-intl';
import { Link } from '@/lib/i18n/navigation';
import { RadarAgir } from '@/components/escore/RadarAgir';

export function ScoreSection() {
  const t = useTranslations('home');
  return (
    <section className="bg-cream">
      <div className="site-container section grid gap-16 lg:grid-cols-2 items-center">
        <div className="space-y-8">
          <p className="label-upper text-gold">Diagnóstico</p>
          <h2 className="heading-section text-petrol text-3xl md:text-5xl">{t('scoreTitle')}</h2>
          <p className="text-petrol/80 text-lg leading-relaxed max-w-lg">{t('scoreSubtitle')}</p>

          <ul className="space-y-4 text-petrol/70">
            <li className="flex gap-4 items-start">
              <span className="text-gold text-lg leading-none mt-0.5">—</span>
              <span>
                <strong className="text-petrol">Triagem</strong> — sua primeira
                leitura: cerca de 35 perguntas{' '}
                <Link
                  href="/escore-plenya/avaliar"
                  className="text-gold underline decoration-gold/50 underline-offset-4 hover:decoration-gold transition"
                >
                  online e gratuitas
                </Link>
                , em poucos minutos.
              </span>
            </li>
            <li className="flex gap-4 items-start">
              <span className="text-gold text-lg leading-none mt-0.5">—</span>
              <span>
                <strong className="text-petrol">Consulta Plenya</strong> —
                aprofundamento clínico com o médico, presencial ou online, com
                base no painel ampliado de exames.{' '}
                <Link
                  href="/consultas"
                  className="text-gold underline decoration-gold/50 underline-offset-4 hover:decoration-gold transition"
                >
                  Agendar
                </Link>
                .
              </span>
            </li>
            <li className="flex gap-4 items-start">
              <span className="text-gold text-lg leading-none mt-0.5">—</span>
              <span>
                <strong className="text-petrol">Continuum Plenya</strong> —
                versão completa, aplicada pela equipe ao longo do{' '}
                <Link
                  href="/continuum"
                  className="text-gold underline decoration-gold/50 underline-offset-4 hover:decoration-gold transition"
                >
                  programa
                </Link>
                , com plano único e reavaliação a cada ciclo.
              </span>
            </li>
          </ul>

          <div className="flex flex-wrap gap-4">
            <Link href="/escore-plenya/avaliar" className="btn-gold">
              Fazer a Triagem
            </Link>
            <Link href="/escore-plenya" className="btn-outline-dark">
              Entender o Escore
            </Link>
          </div>
        </div>

        <div className="flex items-center justify-center">
          <RadarAgir />
        </div>
      </div>
    </section>
  );
}
