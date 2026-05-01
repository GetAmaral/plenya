import { useTranslations } from 'next-intl';
import { Link } from '@/lib/i18n/navigation';
import { RadarAgir } from '@/components/escore/RadarAgir';

export function ScoreSection() {
  const t = useTranslations('home');
  const tCta = useTranslations('cta');
  return (
    <section className="bg-cream">
      <div className="site-container section grid gap-16 lg:grid-cols-2 items-center">
        <div className="space-y-8">
          <p className="label-upper text-gold">{t('scoreLabel')}</p>
          <h2 className="heading-section text-petrol text-3xl md:text-5xl">{t('scoreTitle')}</h2>
          <p className="text-petrol/80 text-lg leading-relaxed max-w-lg">{t('scoreSubtitle')}</p>

          <ul className="space-y-4 text-petrol/70">
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

          <div className="flex flex-wrap gap-4">
            <Link href="/escore-plenya/avaliar" className="btn-gold">
              {tCta('takeTriage')}
            </Link>
            <Link href="/escore-plenya" className="btn-outline-dark">
              {tCta('understandScore')}
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
