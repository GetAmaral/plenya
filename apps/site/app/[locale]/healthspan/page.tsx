import type { Metadata } from 'next';
import { setRequestLocale, getTranslations } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { FaqAccordion } from '@/components/marketing/faq-accordion';
import { FaqSchema } from '@/components/seo/faq-schema';
import { BreadcrumbSchema } from '@/components/seo/breadcrumb-schema';
import { MedicalWebPageSchema } from '@/components/seo/medical-webpage-schema';
import { ClinicalReviewBadge } from '@/components/marketing/clinical-review-badge';
import { RelatedBlogPosts } from '@/components/marketing/related-blog-posts';
import { defaultLocale, isLocale } from '@/lib/i18n/config';

type Params = Promise<{ locale: string }>;

export async function generateMetadata({ params }: { params: Params }): Promise<Metadata> {
  const { locale } = await params;
  const t = await getTranslations({ locale, namespace: 'healthspan' });
  return {
    title: t('metaTitle'),
    description: t('metaDescription'),
    alternates: {
      canonical: '/healthspan',
      languages: {
        'pt-BR': '/healthspan',
        en: '/en/healthspan',
        'x-default': '/healthspan',
      },
    },
    openGraph: {
      title: t('ogTitle'),
      description: t('ogDescription'),
      type: 'article',
    },
  };
}

export default async function HealthspanPage({ params }: { params: Params }) {
  const { locale } = await params;
  setRequestLocale(locale);
  const t = await getTranslations('healthspan');

  const faq = [1, 2, 3, 4, 5, 6].map((n) => ({
    q: t(`faq${n}Q` as 'faq1Q'),
    a: t(`faq${n}A` as 'faq1A'),
  }));

  const measure = [1, 2, 3, 4, 5, 6, 7].map((n) => t(`asideMeasure${n}` as 'asideMeasure1'));
  const noList = [1, 2, 3, 4].map((n) => t(`no${n}` as 'no1'));
  const yesList = [1, 2, 3, 4].map((n) => t(`yes${n}` as 'yes1'));

  return (
    <>
      <FaqSchema items={faq} />
      <MedicalWebPageSchema
        name={t('schemaName')}
        description={t('schemaDescription')}
        path="/healthspan"
        about="Healthspan"
      />
      <BreadcrumbSchema
        items={[
          { name: 'Home', url: '/' },
          { name: 'Healthspan' },
        ]}
      />

      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-24 md:pt-40 md:pb-32 max-w-3xl">
          <p className="label-upper text-gold mb-6">{t('heroLabel')}</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,4.5rem)] text-cream">
            {t('heroTitle')}
          </h1>
          <p className="text-cream/75 text-lg leading-relaxed mt-8 max-w-2xl">
            {t('heroPPart1')}
            <strong className="text-cream">{t('heroPStrong')}</strong>
            {t('heroPPart2')}
          </p>
        </div>
      </section>

      <section className="bg-cream">
        <div className="site-container section grid lg:grid-cols-[2fr_1fr] gap-16">
          <div className="space-y-6 text-petrol/85 text-lg leading-relaxed">
            <p className="label-upper text-gold">{t('windowLabel')}</p>
            <p>{t('windowP1')}</p>
            <p>{t('windowP2')}</p>
            <p>{t('windowP3')}</p>
            <p>{t('windowP4')}</p>
          </div>

          <aside className="space-y-6 lg:sticky lg:top-28 self-start">
            <div className="bg-paper border-l-2 border-gold p-7 space-y-4">
              <p className="label-upper text-gold">{t('asideMeasureLabel')}</p>
              <ul className="text-petrol/85 text-base space-y-2 leading-relaxed">
                {measure.map((m) => (
                  <li key={m}>· {m}</li>
                ))}
              </ul>
            </div>
            <div className="bg-petrol text-cream p-7 space-y-4">
              <p className="label-upper text-gold">{t('asideContinuumLabel')}</p>
              <p className="heading-section text-cream text-2xl leading-tight">
                {t('asideContinuumTitle')}
              </p>
              <p className="text-cream/75 text-base leading-relaxed">{t('asideContinuumDesc')}</p>
              <Link href="/continuum" className="btn-outline-light w-full text-center">
                {t('asideContinuumCta')}
              </Link>
            </div>
          </aside>
        </div>
      </section>

      <section className="bg-paper">
        <div className="site-container section max-w-4xl space-y-8">
          <p className="label-upper text-gold">{t('promiseLabel')}</p>
          <h2 className="heading-section text-petrol text-3xl md:text-4xl max-w-2xl">
            {t('promiseTitlePart1')}
            <em>{t('promiseTitleEm')}</em>
            {t('promiseTitlePart2')}
          </h2>
          <div className="grid md:grid-cols-2 gap-8 text-petrol/80 leading-relaxed">
            <div className="space-y-3">
              <p className="label-upper text-petrol/60">{t('noLabel')}</p>
              <ul className="space-y-2 text-base">
                {noList.map((n) => (
                  <li key={n}>· {n}</li>
                ))}
              </ul>
            </div>
            <div className="space-y-3">
              <p className="label-upper text-petrol/60">{t('yesLabel')}</p>
              <ul className="space-y-2 text-base">
                {yesList.map((y) => (
                  <li key={y}>· {y}</li>
                ))}
              </ul>
            </div>
          </div>
        </div>
      </section>

      <section className="bg-cream">
        <div className="site-container">
          <ClinicalReviewBadge />
        </div>
      </section>

      <FaqAccordion title={t('faqTitle')} items={faq} />

      <RelatedBlogPosts
        title={t('relatedTitle')}
        pillars={['longevidade', 'gestao-metabolica']}
        limit={3}
        locale={isLocale(locale) ? locale : defaultLocale}
      />

      <section className="bg-cream">
        <div className="site-container section grid md:grid-cols-12 gap-12">
          <div className="md:col-span-6 space-y-6">
            <p className="label-upper text-gold">{t('ctaLabel')}</p>
            <h2 className="heading-section text-petrol text-3xl md:text-4xl">
              {t('ctaTitle')}
            </h2>
            <p className="text-petrol/75 leading-relaxed max-w-md">{t('ctaDesc')}</p>
            <div className="flex gap-4 flex-wrap pt-2">
              <Link href="/escore-plenya" className="btn-gold">
                {t('ctaButton1')}
              </Link>
              <Link href="/contato" className="btn-outline-petrol">
                {t('ctaButton2')}
              </Link>
            </div>
          </div>
          <div className="md:col-span-6 grid sm:grid-cols-2 gap-6 content-end">
            <Link
              href="/continuum"
              className="border border-petrol/10 p-6 hover:bg-petrol hover:text-cream transition group"
            >
              <p className="label-upper text-gold mb-2">{t('crossContinuumLabel')}</p>
              <p className="text-petrol/85 group-hover:text-cream/85 text-sm">
                {t('crossContinuumDesc')}
              </p>
            </Link>
            <Link
              href="/metodo-agir"
              className="border border-petrol/10 p-6 hover:bg-petrol hover:text-cream transition group"
            >
              <p className="label-upper text-gold mb-2">{t('crossMethodLabel')}</p>
              <p className="text-petrol/85 group-hover:text-cream/85 text-sm">
                {t('crossMethodDesc')}
              </p>
            </Link>
          </div>
        </div>
      </section>
    </>
  );
}
