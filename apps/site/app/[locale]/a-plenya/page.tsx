import type { Metadata } from 'next';
import Image from 'next/image';
import { setRequestLocale, getTranslations } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { getAgirPillars } from '@plenya/brand';
import { PlenyaInfinity } from '@plenya/brand/logo';

type Params = Promise<{ locale: string }>;

export async function generateMetadata({ params }: { params: Params }): Promise<Metadata> {
  const { locale } = await params;
  const t = await getTranslations({ locale, namespace: 'about' });
  return {
    title: t('metaTitle'),
    description: t('metaDescription'),
    alternates: {
      canonical: locale === 'en' ? '/en/about' : '/a-plenya',
      languages: {
        'pt-BR': '/a-plenya',
        en: '/en/about',
        'x-default': '/a-plenya',
      },
    },
  };
}

export default async function AboutPage({ params }: { params: Params }) {
  const { locale } = await params;
  setRequestLocale(locale);
  const t = await getTranslations('about');
  const pillars = getAgirPillars(locale);

  const manifesto = [1, 2, 3, 4, 5, 6].map((n) => t(`manifesto${n}` as 'manifesto1'));

  return (
    <>
      {/* Manifesto — com background fotográfico */}
      <section className="relative bg-petrol text-cream overflow-hidden">
        <Image
          src="/images/hero-about.jpg"
          alt=""
          fill
          priority
          className="object-cover opacity-50"
          sizes="100vw"
        />
        <div className="absolute inset-0 bg-gradient-to-r from-petrol via-petrol/80 to-petrol/40" />
        <div className="relative site-container pt-32 pb-24 md:pt-40 md:pb-32 max-w-3xl">
          <p className="label-upper text-gold mb-10">{t('manifestoLabel')}</p>
          <div className="space-y-5 heading-section text-3xl md:text-5xl text-cream/95">
            {manifesto.map((line, i) => (
              <p key={i}>{line}</p>
            ))}
            <p className="text-gold pt-8">{t('manifesto7')}</p>
          </div>
        </div>
      </section>

      {/* Origem */}
      <section className="bg-cream">
        <div className="site-container section">
          <div className="grid lg:grid-cols-[1fr_2fr] gap-12 lg:gap-20 items-start max-w-5xl">
            <div className="space-y-3">
              <p className="label-upper text-gold">{t('originLabel')}</p>
              <p className="font-sans text-petrol text-6xl md:text-7xl font-light leading-none tabular-nums">
                {t('originYears')}
              </p>
              <p className="label-upper text-petrol/60">{t('originYearsCaption')}</p>
            </div>
            <div className="space-y-6 text-petrol/80 text-lg leading-relaxed">
              <p>
                {t('originP1Part1')}
                <strong className="text-petrol">{t('originP1Strong')}</strong>
                {t('originP1Part2')}
              </p>
              <p>{t('originP2')}</p>
              <p>
                {t('originP3Part1')}
                <strong className="text-petrol">{t('originP3Strong')}</strong>
                {t('originP3Part2')}
              </p>
              <p className="heading-section text-petrol text-2xl md:text-3xl pt-2">
                {t('originP4')}
              </p>
            </div>
          </div>
        </div>
      </section>

      {/* Imagem da clínica */}
      <section className="bg-cream">
        <div className="site-container pt-4">
          <figure className="space-y-4">
            <div className="relative aspect-[3/2] overflow-hidden">
              <Image
                src="/images/clinic-interior.jpg"
                alt={t('clinicCaption')}
                fill
                className="object-cover"
                sizes="(min-width: 1024px) 1120px, 100vw"
              />
            </div>
            <figcaption className="flex items-baseline gap-3 text-petrol/60 text-sm">
              <span className="label-upper text-gold">{t('clinicCaptionLabel')}</span>
              <span>{t('clinicCaption')}</span>
            </figcaption>
          </figure>
        </div>
      </section>

      {/* Imagem da equipe */}
      <section className="bg-cream">
        <div className="site-container pt-12 md:pt-16">
          <figure className="space-y-4">
            <div className="relative aspect-[3/2] overflow-hidden">
              <Image
                src="/images/team/equipe-candid.jpg"
                alt={t('teamCaption')}
                fill
                className="object-cover"
                sizes="(min-width: 1024px) 1120px, 100vw"
              />
            </div>
            <figcaption className="flex items-baseline gap-3 text-petrol/60 text-sm">
              <span className="label-upper text-gold">{t('teamCaptionLabel')}</span>
              <span>{t('teamCaption')}</span>
            </figcaption>
          </figure>
        </div>
      </section>

      {/* Propósito */}
      <section className="bg-cream">
        <div className="site-container section">
          <div className="max-w-3xl space-y-6">
            <p className="label-upper text-gold">{t('purposeLabel')}</p>
            <h2 className="heading-section text-petrol text-3xl md:text-5xl">
              {t('purposeTitle')}
            </h2>
            <p className="text-petrol/80 text-lg leading-relaxed">{t('purposeP1')}</p>
          </div>
        </div>
      </section>

      {/* Círculo de Ouro */}
      <section className="bg-petrol text-cream">
        <div className="site-container section">
          <div className="grid lg:grid-cols-[auto_1fr] gap-12 lg:gap-20 items-start">
            <div className="space-y-4">
              <PlenyaInfinity className="h-14 w-auto text-gold" />
              <p className="label-upper text-gold">{t('goldenLabel')}</p>
              <h2 className="heading-section text-cream text-3xl md:text-5xl max-w-md">
                {t('goldenTitle')}
              </h2>
            </div>
            <dl className="space-y-10">
              <div>
                <dt className="heading-section text-gold text-xl mb-3">{t('goldenWhyLabel')}</dt>
                <dd className="text-cream/80 text-lg leading-relaxed">{t('goldenWhy')}</dd>
              </div>
              <div>
                <dt className="heading-section text-gold text-xl mb-3">{t('goldenHowLabel')}</dt>
                <dd className="text-cream/80 text-lg leading-relaxed">{t('goldenHow')}</dd>
              </div>
              <div>
                <dt className="heading-section text-gold text-xl mb-3">{t('goldenWhatLabel')}</dt>
                <dd className="text-cream/80 text-lg leading-relaxed">{t('goldenWhat')}</dd>
              </div>
            </dl>
          </div>
        </div>
      </section>

      {/* Método AGIR — preview com link para página completa */}
      <section className="bg-paper">
        <div className="site-container section">
          <p className="label-upper text-gold mb-6">{t('agirSectionLabel')}</p>
          <h2 className="heading-section text-petrol text-3xl md:text-5xl max-w-2xl mb-6">
            {t('agirSectionTitle')}
          </h2>
          <p className="text-petrol/70 text-lg max-w-2xl mb-16 leading-relaxed">
            {t('agirSectionDesc')}
          </p>

          <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-px bg-petrol/15 border-y border-petrol/15">
            {pillars.map((pillar) => (
              <div key={pillar.slug} className="bg-paper p-8 space-y-4">
                <span className="heading-section text-gold text-5xl block leading-none">
                  {pillar.code}
                </span>
                <h3 className="heading-section text-petrol text-xl">{pillar.name}</h3>
                <p className="text-petrol/70 text-sm leading-relaxed">{pillar.idea}</p>
                <p className="label-upper text-petrol/50 text-[10px] tracking-[0.25em] pt-2 border-t border-petrol/10">
                  {pillar.territory}
                </p>
              </div>
            ))}
          </div>

          <div className="mt-12">
            <Link href="/metodo-agir" className="btn-gold">
              {t('agirSectionCta')}
            </Link>
          </div>
        </div>
      </section>

      {/* Posicionamento */}
      <section className="bg-cream">
        <div className="site-container section">
          <div className="max-w-3xl space-y-6">
            <p className="label-upper text-gold">{t('positioningLabel')}</p>
            <h2 className="heading-section text-petrol text-3xl md:text-4xl">
              {t('positioningTitle')}
            </h2>
            <p className="text-petrol/80 text-lg leading-relaxed">{t('positioningDesc')}</p>
          </div>
        </div>
      </section>

      {/* Cross-links */}
      <section className="bg-paper">
        <div className="site-container section grid md:grid-cols-3 gap-8">
          <Link href="/dr-getulio" className="border-t border-petrol/15 pt-8 space-y-3 group">
            <p className="label-upper text-gold">{t('crossDirectorLabel')}</p>
            <p className="heading-section text-petrol text-xl group-hover:text-gold transition">
              {t('crossDirectorTitle')}
            </p>
          </Link>
          <Link href="/escore-plenya" className="border-t border-petrol/15 pt-8 space-y-3 group">
            <p className="label-upper text-gold">{t('crossScoreLabel')}</p>
            <p className="heading-section text-petrol text-xl group-hover:text-gold transition">
              {t('crossScoreTitle')}
            </p>
          </Link>
          <Link href="/continuum" className="border-t border-petrol/15 pt-8 space-y-3 group">
            <p className="label-upper text-gold">{t('crossContinuumLabel')}</p>
            <p className="heading-section text-petrol text-xl group-hover:text-gold transition">
              {t('crossContinuumTitle')}
            </p>
          </Link>
        </div>
      </section>
    </>
  );
}
