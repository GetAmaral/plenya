import { useTranslations } from 'next-intl';
import { Link } from '@/lib/i18n/navigation';
import { RadarAgir } from '@/components/escore/RadarAgir';

export function ScoreSection() {
  const t = useTranslations('home');
  const tCta = useTranslations('cta');
  return (
    <section className="bg-cream">
      <div className="site-container py-20 md:py-28">
        {/* Faixa gold + label de entrada — sinaliza visualmente "porta de entrada gratuita" */}
        <div className="border-t-4 border-gold pt-10 md:pt-14">
          <div className="grid gap-12 lg:gap-16 lg:grid-cols-[1.3fr_1fr] items-start">
            <div className="space-y-8">
              <div className="space-y-5">
                <p className="label-upper text-gold">{t('scoreLabel')}</p>
                <h2 className="heading-section text-petrol text-4xl md:text-6xl leading-[1.05]">
                  {t('scoreTitle')}
                </h2>
                <p className="text-petrol/80 text-lg leading-relaxed max-w-lg">
                  {t('scoreSubtitle')}
                </p>
              </div>

              {/* Feature pills — peso visual de "porta de entrada" */}
              <ul className="flex flex-wrap gap-3">
                <li className="px-4 py-2 border border-gold/50 text-gold label-upper text-[11px] tracking-[0.18em]">
                  {t('scoreFeatureFree')}
                </li>
                <li className="px-4 py-2 border border-gold/50 text-gold label-upper text-[11px] tracking-[0.18em]">
                  {t('scoreFeatureTime')}
                </li>
                <li className="px-4 py-2 border border-gold/50 text-gold label-upper text-[11px] tracking-[0.18em]">
                  {t('scoreFeatureEmail')}
                </li>
              </ul>

              <ul className="space-y-3 text-petrol/70 text-[15px]">
                <li className="flex gap-4 items-start">
                  <span className="text-gold text-lg leading-none mt-0.5">—</span>
                  <span>
                    <strong className="text-petrol">{t('scoreItem1Strong')}</strong>{t('scoreItem1Part1')}
                    <Link
                      href="/escore-plenya/avaliar"
                      className="text-gold underline decoration-gold/50 underline-offset-4 hover:decoration-gold transition"
                    >
                      {t('scoreItem1Link')}
                    </Link>
                    {t('scoreItem1Part2')}
                  </span>
                </li>
                <li className="flex gap-4 items-start">
                  <span className="text-gold text-lg leading-none mt-0.5">—</span>
                  <span>
                    <strong className="text-petrol">{t('scoreItem2Strong')}</strong>{t('scoreItem2Part1')}
                    <Link
                      href="/consultas"
                      className="text-gold underline decoration-gold/50 underline-offset-4 hover:decoration-gold transition"
                    >
                      {t('scoreItem2Link')}
                    </Link>
                    {t('scoreItem2Part2')}
                  </span>
                </li>
                <li className="flex gap-4 items-start">
                  <span className="text-gold text-lg leading-none mt-0.5">—</span>
                  <span>
                    <strong className="text-petrol">{t('scoreItem3Strong')}</strong>{t('scoreItem3Part1')}
                    <Link
                      href="/continuum"
                      className="text-gold underline decoration-gold/50 underline-offset-4 hover:decoration-gold transition"
                    >
                      {t('scoreItem3Link')}
                    </Link>
                    {t('scoreItem3Part2')}
                  </span>
                </li>
              </ul>

              <div className="flex flex-wrap gap-4 pt-2">
                <Link href="/escore-plenya/avaliar" className="btn-gold">
                  {tCta('takeTriage')}
                </Link>
                <Link href="/escore-plenya" className="btn-outline-dark">
                  {tCta('understandScore')}
                </Link>
              </div>
            </div>

            <div className="flex items-center justify-center lg:pt-4">
              <RadarAgir />
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
