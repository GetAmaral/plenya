import type { Metadata } from 'next';
import Image from 'next/image';
import { setRequestLocale, getTranslations } from 'next-intl/server';
import { isLocale, defaultLocale } from '@/lib/i18n/config';
import { Link } from '@/lib/i18n/navigation';
import { getAllDoctors, localizedShortBio } from '@/lib/team';
import { DoctorCard } from '@/components/team/doctor-card';

type Params = Promise<{ locale: string }>;

export async function generateMetadata({ params }: { params: Params }): Promise<Metadata> {
  const { locale } = await params;
  const t = await getTranslations({ locale, namespace: 'team' });
  return {
    title: t('metaTitle'),
    description: t('metaDescription'),
    alternates: {
      canonical: locale === 'en' ? '/en/team' : '/equipe',
      languages: {
        'pt-BR': '/equipe', pt: '/equipe',
        en: '/en/team',
        'x-default': '/equipe',
      },
    },
  };
}

export default async function TeamPage({ params }: { params: Params }) {
  const { locale: rawLocale } = await params;
  const locale = isLocale(rawLocale) ? rawLocale : defaultLocale;
  setRequestLocale(locale);
  const t = await getTranslations('team');

  const all = await getAllDoctors();
  const direcao = all.find((d) => d.slug === 'getulio-amaral');
  const nucleoMedico = all.filter(
    (d) => d.category === 'medico' && d.slug !== 'getulio-amaral',
  );
  const multidisciplinar = all.filter((d) => d.category !== 'medico');

  return (
    <>
      {/* Hero */}
      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-16 md:pt-40 md:pb-20 grid lg:grid-cols-[1fr_auto] gap-12 lg:gap-20 items-center">
          <div className="max-w-xl">
            <p className="label-upper text-gold mb-6">{t('heroLabel')}</p>
            <h1 className="heading-hero text-[clamp(2.5rem,6vw,5rem)] text-cream">
              {t('heroTitle')}
            </h1>
            <p className="text-cream/80 text-lg mt-6 leading-relaxed">{t('heroP1')}</p>
            <p className="text-cream/60 mt-4 leading-relaxed">{t('heroP2')}</p>
          </div>
          <div className="relative w-full max-w-md lg:w-[360px] lg:max-w-none aspect-1064/1891 overflow-hidden bg-petrol/40">
            <Image
              src="/images/team/equipe-formal.jpg"
              alt={t('heroAlt')}
              fill
              priority
              className="object-cover object-top"
              sizes="(min-width: 1024px) 360px, 100vw"
            />
          </div>
        </div>
      </section>

      {/* Direção clínica */}
      {direcao && (
        <section className="bg-cream">
          <div className="site-container section">
            <div className="flex items-baseline gap-4 mb-12">
              <h2 className="heading-section text-petrol text-2xl md:text-3xl">
                {t('directionTitle')}
              </h2>
              <span className="label-upper text-petrol/40">{t('directionCaption')}</span>
            </div>
            <Link
              href="/dr-getulio"
              className="group grid md:grid-cols-[460px_1fr] gap-10 lg:gap-16 items-start bg-paper hover:bg-cream-100 transition-colors duration-300 p-6 md:p-10"
            >
              <div className="relative aspect-3/4 overflow-hidden">
                {direcao.photo && (
                  <Image
                    src={direcao.photo}
                    alt={direcao.name}
                    fill
                    className="object-cover object-top transition-transform duration-700 group-hover:scale-[1.02]"
                    sizes="(min-width: 768px) 460px, 100vw"
                  />
                )}
              </div>
              <div className="space-y-5 pt-4">
                <p className="label-upper text-gold">{direcao.role}</p>
                <h3 className="heading-section text-petrol text-4xl md:text-6xl group-hover:text-gold transition">
                  {direcao.name}
                </h3>
                <p className="label-upper-sm text-petrol/55">
                  {direcao.credentials}
                  {direcao.rqe ? ` · RQE ${direcao.rqe}` : ''}
                </p>
                <p className="text-petrol/80 leading-relaxed text-lg max-w-2xl">
                  {localizedShortBio(direcao, locale)}
                </p>
                <p className="label-upper text-gold pt-2 group-hover:underline underline-offset-4">
                  {t('directionCta')}
                </p>
              </div>
            </Link>
          </div>
        </section>
      )}

      {/* Núcleo médico */}
      {nucleoMedico.length > 0 && (
        <section className="bg-paper">
          <div className="site-container section">
            <div className="flex items-baseline gap-4 mb-12">
              <h2 className="heading-section text-petrol text-2xl md:text-3xl">
                {t('doctorsTitle')}
              </h2>
              <span className="label-upper text-petrol/40">{t('doctorsCaption')}</span>
            </div>
            <div className="grid gap-8 sm:grid-cols-2 lg:grid-cols-3">
              {nucleoMedico.map((d) => (
                <DoctorCard key={d.slug} doctor={d} locale={locale} />
              ))}
            </div>
          </div>
        </section>
      )}

      {/* Multidisciplinar */}
      {multidisciplinar.length > 0 && (
        <section className="bg-cream">
          <div className="site-container section">
            <div className="flex items-baseline gap-4 mb-12">
              <h2 className="heading-section text-petrol text-2xl md:text-3xl">
                {t('multiTitle')}
              </h2>
              <span className="label-upper text-petrol/40">{t('multiCaption')}</span>
            </div>
            <div className="grid gap-8 sm:grid-cols-2 lg:grid-cols-3">
              {multidisciplinar.map((d) => (
                <DoctorCard key={d.slug} doctor={d} locale={locale} />
              ))}
            </div>
          </div>
        </section>
      )}

      <section className="bg-petrol text-cream">
        <div className="site-container section text-center space-y-6">
          <p className="heading-section text-cream text-2xl md:text-3xl max-w-2xl mx-auto">
            {t('ctaTitle')}
          </p>
          <Link href="/contato" className="btn-gold">
            {t('ctaButton')}
          </Link>
        </div>
      </section>
    </>
  );
}
