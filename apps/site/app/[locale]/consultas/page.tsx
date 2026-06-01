import type { Metadata } from 'next';
import Image from 'next/image';
import { setRequestLocale, getTranslations } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { getAllDoctors } from '@/lib/team';
import { FaqAccordion } from '@/components/marketing/faq-accordion';
import { FaqSchema } from '@/components/seo/faq-schema';
import { TestimonialsInline } from '@/components/testimonials/testimonials-inline';

type Params = Promise<{ locale: string }>;

export async function generateMetadata({ params }: { params: Params }): Promise<Metadata> {
  const { locale } = await params;
  const t = await getTranslations({ locale, namespace: 'consultas' });
  return {
    title: t('metaTitle'),
    description: t('metaDescription'),
    alternates: {
      canonical: locale === 'en' ? '/en/consultations' : '/consultas',
      languages: {
        'pt-BR': '/consultas', pt: '/consultas',
        en: '/en/consultations',
        'x-default': '/consultas',
      },
    },
  };
}

export default async function ConsultasPage({ params }: { params: Params }) {
  const { locale } = await params;
  setRequestLocale(locale);
  const t = await getTranslations('consultas');

  const consultasFaq = [1, 2, 3, 4, 5, 6].map((n) => ({
    q: t(`faq${n}Q` as 'faq1Q'),
    a: t(`faq${n}A` as 'faq1A'),
  }));

  const doctors = await getAllDoctors();
  const getulio = doctors.find((d) => d.slug === 'getulio-amaral');
  const outros = doctors.filter((d) => d.category === 'medico' && d.slug !== 'getulio-amaral');

  return (
    <>
      <FaqSchema items={consultasFaq.map((f) => ({ q: f.q, a: f.a }))} />

      {/* Hero */}
      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-24 md:pt-40 md:pb-32 max-w-3xl">
          <p className="label-upper text-gold mb-6">{t('heroLabel')}</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,4.5rem)] text-cream">
            {t('heroTitle')}
          </h1>
          <p className="text-cream/75 text-lg leading-relaxed mt-8 max-w-2xl">{t('heroP')}</p>
        </div>
      </section>

      {/* Sobre a consulta */}
      <section className="bg-cream">
        <div className="site-container section grid lg:grid-cols-[2fr_1fr] gap-16">
          <div className="space-y-6 text-petrol/85 text-lg leading-relaxed">
            <p className="label-upper text-gold">{t('aboutLabel')}</p>
            <p>{t('aboutP1')}</p>
            <p>{t('aboutP2')}</p>
            <p>
              {t('aboutP3Part1')}
              <Link
                href="/continuum"
                className="underline decoration-gold/60 underline-offset-4 hover:text-petrol transition"
              >
                {t('aboutP3Link')}
              </Link>
              {t('aboutP3Part2')}
            </p>
          </div>

          <aside className="space-y-8">
            <div className="space-y-3">
              <p className="label-upper text-gold">{t('modalitiesLabel')}</p>
              <p className="text-petrol/80 leading-relaxed">{t('modalitiesText')}</p>
            </div>

            <div className="space-y-3">
              <p className="label-upper text-gold">{t('paymentLabel')}</p>
              <p className="text-petrol/80 leading-relaxed">{t('paymentText')}</p>
            </div>

            <div className="space-y-3">
              <p className="label-upper text-gold">{t('schedulingLabel')}</p>
              <p className="text-petrol/80 leading-relaxed">{t('schedulingText')}</p>
            </div>
          </aside>
        </div>
      </section>

      {/* Modalidades — visual */}
      <section className="bg-paper">
        <div className="site-container section-sm">
          <div className="grid md:grid-cols-2 gap-6 md:gap-8">
            <figure className="space-y-3">
              <div className="relative aspect-4/5 overflow-hidden">
                <Image
                  src="/images/getulio-consulta.jpg"
                  alt={t('photoInPersonAlt')}
                  fill
                  className="object-cover"
                  sizes="(min-width: 768px) 540px, 100vw"
                />
              </div>
              <figcaption className="label-upper text-gold">{t('photoInPersonCaption')}</figcaption>
            </figure>
            <figure className="space-y-3">
              <div className="relative aspect-4/5 overflow-hidden">
                <Image
                  src="/images/getulio-consulta-online.jpg"
                  alt={t('photoOnlineAlt')}
                  fill
                  className="object-cover"
                  sizes="(min-width: 768px) 540px, 100vw"
                />
              </div>
              <figcaption className="label-upper text-gold">{t('photoOnlineCaption')}</figcaption>
            </figure>
          </div>
        </div>
      </section>

      {/* Quem atende */}
      <section className="bg-petrol text-cream">
        <div className="site-container section">
          <div className="grid lg:grid-cols-12 gap-12 mb-20">
            <div className="lg:col-span-5">
              <p className="label-upper text-gold mb-6">{t('doctorsLabel')}</p>
              <h2 className="heading-section text-cream text-3xl md:text-5xl">
                {t('doctorsTitle')}
              </h2>
            </div>
            <p className="lg:col-span-7 text-cream/75 text-lg leading-relaxed self-end max-w-2xl">
              {t('doctorsDesc')}
            </p>
          </div>

          <div className="grid lg:grid-cols-2 gap-12">
            {getulio && (
              <article className="border-t border-gold pt-8 space-y-5">
                <p className="label-upper text-gold">{t('directorLabel')}</p>
                <h3 className="heading-section text-cream text-3xl md:text-4xl">{getulio.name}</h3>
                <p className="label-upper text-cream/50">
                  {getulio.credentials}
                  {getulio.rqe ? ` · RQE ${getulio.rqe}` : ''}
                </p>
                <p className="text-cream/75 leading-relaxed">{getulio.shortBio}</p>
                <p>
                  <Link
                    href="/dr-getulio"
                    className="text-gold hover:text-cream transition border-b border-gold/40 pb-0.5"
                  >
                    {t('directorAboutLink')}
                  </Link>
                </p>
              </article>
            )}

            <article className="border-t border-cream/30 pt-8 space-y-5">
              <p className="label-upper text-cream/60">{t('otherDoctorsLabel')}</p>
              <h3 className="heading-section text-cream text-3xl md:text-4xl">
                {t('otherDoctorsTitle')}
              </h3>
              <p className="text-cream/75 leading-relaxed">{t('otherDoctorsDesc')}</p>
              {outros.length > 0 && (
                <ul className="space-y-3 pt-2">
                  {outros.map((d) => (
                    <li key={d.slug}>
                      <Link
                        href={{ pathname: '/equipe/[slug]', params: { slug: d.slug } }}
                        className="block group"
                      >
                        <span className="block text-cream group-hover:text-gold transition">
                          {d.name}
                        </span>
                        <span className="block text-cream/50 text-sm mt-0.5">{d.role}</span>
                      </Link>
                    </li>
                  ))}
                </ul>
              )}
              <p>
                <Link
                  href="/equipe"
                  className="text-gold hover:text-cream transition border-b border-gold/40 pb-0.5"
                >
                  {t('fullTeamLink')}
                </Link>
              </p>
            </article>
          </div>
        </div>
      </section>

      <TestimonialsInline
        bg="bg-cream"
        label={t('testimonialsLabel')}
        title={t('testimonialsTitle')}
        limit={3}
      />

      <FaqAccordion title={t('faqTitle')} items={consultasFaq} />

      {/* Agendamento + cross-links */}
      <section className="bg-cream">
        <div className="site-container section grid md:grid-cols-12 gap-12">
          <div className="md:col-span-6 space-y-6">
            <p className="label-upper text-gold">{t('scheduleLabel')}</p>
            <h2 className="heading-section text-petrol text-3xl md:text-4xl">
              {t('scheduleTitle')}
            </h2>
            <p className="text-petrol/75 leading-relaxed max-w-md">{t('scheduleDesc')}</p>
            <p>
              <Link href="/contato" className="btn-gold">
                {t('scheduleCta')}
              </Link>
            </p>
          </div>

          <div className="md:col-span-6 grid sm:grid-cols-2 gap-6 content-end">
            <Link href="/continuum" className="border-t border-petrol/15 pt-6 space-y-2 group">
              <p className="label-upper text-gold">{t('crossContinuityLabel')}</p>
              <p className="heading-section text-petrol text-lg group-hover:text-gold transition">
                {t('crossContinuityTitle')}
              </p>
            </Link>
            <Link href="/escore-plenya" className="border-t border-petrol/15 pt-6 space-y-2 group">
              <p className="label-upper text-gold">{t('crossDiagnosticLabel')}</p>
              <p className="heading-section text-petrol text-lg group-hover:text-gold transition">
                {t('crossDiagnosticTitle')}
              </p>
            </Link>
          </div>
        </div>
      </section>
    </>
  );
}
