import type { Metadata } from 'next';
import { setRequestLocale, getTranslations } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { HowItWorksSteps } from '@/components/marketing/how-it-works-steps';
import { FaqAccordion } from '@/components/marketing/faq-accordion';

type Params = Promise<{ locale: string }>;

export async function generateMetadata({ params }: { params: Params }): Promise<Metadata> {
  const { locale } = await params;
  const t = await getTranslations({ locale, namespace: 'howItWorks' });
  return {
    title: t('metaTitle'),
    description: t('metaDescription'),
    alternates: {
      canonical: locale === 'en' ? '/en/how-it-works' : '/como-funciona',
      languages: {
        'pt-BR': '/como-funciona',
        en: '/en/how-it-works',
        'x-default': '/como-funciona',
      },
    },
  };
}

export default async function ComoFuncionaPage({ params }: { params: Params }) {
  const { locale } = await params;
  setRequestLocale(locale);
  const t = await getTranslations('howItWorks');

  const faq = [
    {
      q: t('faq1Q'),
      a: (
        <>
          {t('faq1ABefore')}
          <Link
            href="/escore-plenya/avaliar"
            className="underline decoration-gold/60 underline-offset-4"
          >
            {t('faq1ATriageLink')}
          </Link>
          {t('faq1AAfter')}
        </>
      ),
    },
    { q: t('faq2Q'), a: t('faq2A') },
    { q: t('faq3Q'), a: t('faq3A') },
    { q: t('faq4Q'), a: t('faq4A') },
    { q: t('faq5Q'), a: t('faq5A') },
    { q: t('faq6Q'), a: t('faq6A') },
  ];

  return (
    <>
      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-24 md:pt-40 md:pb-32 max-w-3xl">
          <p className="label-upper text-gold mb-6">{t('heroLabel')}</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,4.5rem)] text-cream">
            {t('heroTitle')}
          </h1>
          <p className="text-cream/75 text-lg leading-relaxed mt-8 max-w-2xl">{t('heroP')}</p>
        </div>
      </section>

      <HowItWorksSteps bg="bg-cream" />

      <FaqAccordion title={t('faqTitle')} items={faq} />

      <section className="bg-petrol text-cream">
        <div className="site-container section text-center space-y-6">
          <p className="label-upper text-gold">{t('exitLabel')}</p>
          <h2 className="heading-section text-cream text-3xl md:text-4xl max-w-2xl mx-auto">
            {t('exitTitle')}
          </h2>
          <div className="flex flex-wrap justify-center gap-4 pt-4">
            <Link href="/escore-plenya/avaliar" className="btn-gold">
              {t('exitCtaTriage')}
            </Link>
            <Link href="/contato" className="btn-outline-dark border-cream/40 text-cream">
              {t('exitCtaContact')}
            </Link>
          </div>
        </div>
      </section>
    </>
  );
}
