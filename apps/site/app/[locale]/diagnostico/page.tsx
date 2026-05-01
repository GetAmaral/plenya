import type { Metadata } from 'next';
import { setRequestLocale, getTranslations } from 'next-intl/server';
import { DiagnosticoQuiz } from '@/components/marketing/diagnostico-quiz';

type Params = Promise<{ locale: string }>;

export async function generateMetadata({ params }: { params: Params }): Promise<Metadata> {
  const { locale } = await params;
  const t = await getTranslations({ locale, namespace: 'diagnostic' });
  return {
    title: t('metaTitle'),
    description: t('metaDescription'),
    alternates: {
      canonical: locale === 'en' ? '/en/assessment' : '/diagnostico',
      languages: {
        'pt-BR': '/diagnostico',
        en: '/en/assessment',
        'x-default': '/diagnostico',
      },
    },
  };
}

export default async function DiagnosticoPage({ params }: { params: Params }) {
  const { locale } = await params;
  setRequestLocale(locale);
  const t = await getTranslations('diagnostic');

  return (
    <>
      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-16 md:pt-40 md:pb-20">
          <p className="label-upper text-gold mb-6">{t('heroLabel')}</p>
          <h1 className="heading-hero text-[clamp(2.25rem,5vw,4rem)] text-cream max-w-3xl">
            {t('heroTitleLine1')}
            <br />
            {t('heroTitleLine2')}
          </h1>
          <p className="text-cream/70 text-lg mt-6 max-w-2xl leading-relaxed">{t('heroP')}</p>
        </div>
      </section>

      <section className="bg-cream">
        <div className="site-container section">
          <DiagnosticoQuiz />
        </div>
      </section>
    </>
  );
}
