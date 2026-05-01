import { useTranslations } from 'next-intl';
import { PlenyaInfinity } from '@plenya/brand/logo';
import { Link } from '@/lib/i18n/navigation';

export function ContinuumSpotlight() {
  const t = useTranslations('home');
  const tCta = useTranslations('cta');

  const pilares = [
    { label: t('continuumPillarsModalitiesLabel'), value: t('continuumPillarsModalitiesValue') },
    { label: t('continuumPillarsFormatLabel'), value: t('continuumPillarsFormatValue') },
    { label: t('continuumPillarsTeamLabel'), value: t('continuumPillarsTeamValue') },
    { label: t('continuumPillarsCadenceLabel'), value: t('continuumPillarsCadenceValue') },
    { label: t('continuumPillarsDiagnosisLabel'), value: t('continuumPillarsDiagnosisValue') },
    { label: t('continuumPillarsMethodLabel'), value: t('continuumPillarsMethodValue') },
  ];

  return (
    <section className="bg-petrol text-cream relative overflow-hidden">
      <PlenyaInfinity
        aria-hidden="true"
        focusable="false"
        className="hidden md:block absolute -right-20 top-16 w-[420px] h-auto text-gold/[0.07] pointer-events-none"
      />

      <div className="relative site-container section">
        <div className="grid lg:grid-cols-[1.1fr_1fr] gap-12 lg:gap-20 mb-16">
          <div className="space-y-6">
            <p className="label-upper text-gold">{t('continuumLabel')}</p>
            <h2 className="heading-section text-cream text-4xl md:text-6xl">
              {t('continuumTitle')}
            </h2>
            <p className="heading-section text-gold text-2xl md:text-4xl">
              {t('continuumSubtitle')}
            </p>
          </div>

          <div className="space-y-5 text-cream/85 text-lg leading-relaxed self-end">
            <p>
              {t('continuumP1Part1')}<strong className="text-cream">{t('continuumP1Strong1')}</strong>{t('continuumP1Part2')}<strong className="text-cream">{t('continuumP1Strong2')}</strong>{t('continuumP1Part3')}
            </p>
            <p className="text-cream/70">
              {t('continuumP2')}
            </p>
          </div>
        </div>

        <div className="grid sm:grid-cols-2 lg:grid-cols-3 gap-px bg-cream/10 border-y border-cream/15 mb-12">
          {pilares.map((p) => (
            <div key={p.label} className="bg-petrol p-6 space-y-3">
              <p className="label-upper text-gold/85">{p.label}</p>
              <p className="text-cream text-xl leading-snug">{p.value}</p>
            </div>
          ))}
        </div>

        <div className="flex flex-wrap gap-5">
          <Link href="/continuum" className="btn-gold">
            {tCta('knowPlans')}
          </Link>
          <Link
            href="/diagnostico"
            className="btn-outline-light"
          >
            {tCta('isItForMe')}
          </Link>
        </div>
      </div>
    </section>
  );
}
