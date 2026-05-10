import type { Metadata } from 'next';
import { setRequestLocale, getTranslations } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { FaqAccordion } from '@/components/marketing/faq-accordion';
import { FaqSchema } from '@/components/seo/faq-schema';
import { BreadcrumbSchema } from '@/components/seo/breadcrumb-schema';
import { MedicalWebPageSchema } from '@/components/seo/medical-webpage-schema';
import { ClinicalReviewBadge } from '@/components/marketing/clinical-review-badge';
import { RelatedBlogPosts } from '@/components/marketing/related-blog-posts';

type Params = Promise<{ locale: string }>;

export async function generateMetadata({ params }: { params: Params }): Promise<Metadata> {
  const { locale } = await params;
  const t = await getTranslations({ locale, namespace: 'mfi' });
  return {
    title: t('metaTitle'),
    description: t('metaDescription'),
    alternates: {
      canonical: '/medicina-funcional-integrativa',
      languages: {
        'pt-BR': '/medicina-funcional-integrativa', pt: '/medicina-funcional-integrativa',
        en: '/en/integrative-functional-medicine',
        'x-default': '/medicina-funcional-integrativa',
      },
    },
    openGraph: {
      title: t('ogTitle'),
      description: t('ogDescription'),
      type: 'article',
    },
  };
}

export default async function MfiPage({ params }: { params: Params }) {
  const { locale } = await params;
  setRequestLocale(locale);
  const t = await getTranslations('mfi');

  const faq = [1, 2, 3, 4, 5, 6].map((n) => ({
    q: t(`faq${n}Q` as 'faq1Q'),
    a: t(`faq${n}A` as 'faq1A'),
  }));

  const cascade = [1, 2, 3, 4].map((n) => t(`asideCascade${n}` as 'asideCascade1'));
  const helps = [1, 2, 3, 4, 5, 6, 7].map((n) => t(`asideHelps${n}` as 'asideHelps1'));

  return (
    <>
      <FaqSchema items={faq} />
      <MedicalWebPageSchema
        name={t('schemaName')}
        description={t('schemaDescription')}
        path="/medicina-funcional-integrativa"
        about="Functional Medicine"
      />
      <BreadcrumbSchema
        items={[{ name: 'Home', url: '/' }, { name: t('breadcrumbName') }]}
      />

      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-24 md:pt-40 md:pb-32 max-w-3xl">
          <p className="label-upper text-gold mb-6">{t('heroLabel')}</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,4.5rem)] text-cream">
            {t('heroTitle')}
          </h1>
          <p className="text-cream/75 text-lg leading-relaxed mt-8 max-w-2xl">
            {t('heroPPart1')}
            <em>{t('heroPEm1')}</em>
            {t('heroPPart2')}
            <em>{t('heroPEm2')}</em>
            {t('heroPPart3')}
          </p>
        </div>
      </section>

      <section className="bg-cream">
        <div className="site-container section grid lg:grid-cols-[2fr_1fr] gap-16">
          <div className="space-y-6 text-petrol/85 text-lg leading-relaxed">
            <p className="label-upper text-gold">{t('approachLabel')}</p>
            <p>{t('approachP1')}</p>
            <p>{t('approachP2')}</p>
            <p>{t('approachP3Part1')}</p>
            <p>{t('approachP4')}</p>
          </div>

          <aside className="space-y-6">
            <div className="border border-petrol/10 p-6 space-y-3">
              <p className="label-upper text-gold">{t('asideCascadeLabel')}</p>
              <ol className="text-petrol/75 text-sm space-y-1.5 leading-relaxed list-decimal pl-5">
                {cascade.map((c) => (
                  <li key={c}>{c}</li>
                ))}
              </ol>
            </div>
            <div className="border border-petrol/10 p-6 space-y-3">
              <p className="label-upper text-gold">{t('asideHelpsLabel')}</p>
              <ul className="text-petrol/75 text-sm space-y-1.5 leading-relaxed">
                {helps.map((h) => (
                  <li key={h}>· {h}</li>
                ))}
              </ul>
            </div>
          </aside>
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
        pillars={['gestao-metabolica', 'integracao-corpo-mente', 'longevidade']}
        limit={3}
      />

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
            <Link
              href="/metodo-agir"
              className="border border-petrol/10 p-6 hover:bg-petrol hover:text-cream transition group"
            >
              <p className="label-upper text-gold mb-2">{t('crossMethodLabel')}</p>
              <p className="text-petrol/85 group-hover:text-cream/85 text-sm">
                {t('crossMethodDesc')}
              </p>
            </Link>
            <Link
              href="/continuum"
              className="border border-petrol/10 p-6 hover:bg-petrol hover:text-cream transition group"
            >
              <p className="label-upper text-gold mb-2">{t('crossContinuumLabel')}</p>
              <p className="text-petrol/85 group-hover:text-cream/85 text-sm">
                {t('crossContinuumDesc')}
              </p>
            </Link>
          </div>
        </div>
      </section>
    </>
  );
}
