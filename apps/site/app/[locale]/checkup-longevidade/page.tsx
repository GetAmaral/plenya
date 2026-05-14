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
  const t = await getTranslations({ locale, namespace: 'checkup' });
  return {
    title: t('metaTitle'),
    description: t('metaDescription'),
    alternates: {
      canonical: '/checkup-longevidade',
      languages: {
        'pt-BR': '/checkup-longevidade', pt: '/checkup-longevidade',
        en: '/en/longevity-checkup',
        'x-default': '/checkup-longevidade',
      },
    },
    openGraph: {
      title: t('ogTitle'),
      description: t('ogDescription'),
      type: 'article',
    },
  };
}

export default async function CheckupPage({ params }: { params: Params }) {
  const { locale } = await params;
  setRequestLocale(locale);
  const t = await getTranslations('checkup');

  const faq = [1, 2, 3, 4, 5, 6].map((n) => ({
    q: t(`faq${n}Q` as 'faq1Q'),
    a: t(`faq${n}A` as 'faq1A'),
  }));

  const layer1 = [1, 2, 3, 4, 5].map((n) => t(`aside1_${n}` as 'aside1_1'));
  const layer2 = [1, 2, 3].map((n) => t(`aside2_${n}` as 'aside2_1'));
  const layer3 = [1, 2, 3].map((n) => t(`aside3_${n}` as 'aside3_1'));
  const markers = Array.from({ length: 12 }, (_, i) => ({
    name: t(`marker${i + 1}Name` as 'marker1Name'),
    line: t(`marker${i + 1}Line` as 'marker1Line'),
  }));

  return (
    <>
      <FaqSchema items={faq} />
      <MedicalWebPageSchema
        name={t('schemaName')}
        description={t('schemaDescription')}
        path="/checkup-longevidade"
        about="Preventive medical checkup"
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
          <p className="text-cream/75 text-lg leading-relaxed mt-8 max-w-2xl">{t('heroP')}</p>
        </div>
      </section>

      <section className="bg-cream">
        <div className="site-container section grid lg:grid-cols-[2fr_1fr] gap-16">
          <div className="space-y-6 text-petrol/85 text-lg leading-relaxed">
            <p className="label-upper text-gold">{t('aboutLabel')}</p>
            <p>{t('aboutP1')}</p>
            <p>{t('aboutP2')}</p>
            <p>{t('aboutP3')}</p>
            <p>{t('aboutP4')}</p>
          </div>

          <aside className="space-y-6">
            <div className="border border-petrol/10 p-6 space-y-3">
              <p className="label-upper text-gold">{t('aside1Label')}</p>
              <ul className="text-petrol/75 text-sm space-y-1.5 leading-relaxed">
                {layer1.map((m) => (
                  <li key={m}>· {m}</li>
                ))}
              </ul>
            </div>
            <div className="border border-petrol/10 p-6 space-y-3">
              <p className="label-upper text-gold">{t('aside2Label')}</p>
              <ul className="text-petrol/75 text-sm space-y-1.5 leading-relaxed">
                {layer2.map((m) => (
                  <li key={m}>· {m}</li>
                ))}
              </ul>
            </div>
            <div className="border border-petrol/10 p-6 space-y-3">
              <p className="label-upper text-gold">{t('aside3Label')}</p>
              <ul className="text-petrol/75 text-sm space-y-1.5 leading-relaxed">
                {layer3.map((m) => (
                  <li key={m}>· {m}</li>
                ))}
              </ul>
            </div>
          </aside>
        </div>
      </section>

      <section className="bg-paper">
        <div className="site-container section">
          <p className="label-upper text-gold mb-4">{t('panelLabel')}</p>
          <h2 className="heading-section text-petrol text-3xl md:text-4xl max-w-2xl mb-8">
            {t('panelTitle')}
          </h2>
          <p className="text-petrol/75 max-w-3xl leading-relaxed mb-12">{t('panelIntro')}</p>
          <ul className="grid md:grid-cols-2 gap-x-12 gap-y-8">
            {markers.map((m, i) => (
              <li
                key={i}
                className="border-t border-petrol/15 pt-4 space-y-2"
              >
                <p className="font-serif text-lg text-petrol leading-snug">
                  <span className="text-gold tabular-nums mr-2">{String(i + 1).padStart(2, '0')}</span>
                  {m.name}
                </p>
                <p
                  className="text-petrol/70 text-sm leading-relaxed"
                  dangerouslySetInnerHTML={{ __html: m.line }}
                />
              </li>
            ))}
          </ul>
          <p className="text-petrol/55 text-sm mt-12 max-w-3xl leading-relaxed">
            {t('panelDisclaimer')}
          </p>
          <p className="mt-6 text-petrol/75">
            <span className="text-petrol/55">{t('panelReadMoreLabel')}</span>{' '}
            <Link
              href={{ pathname: '/blog/[slug]', params: { slug: t('panelBlogSlug') } }}
              className="text-gold underline-offset-4 hover:underline"
            >
              {t('panelReadMoreLink')}
            </Link>
          </p>
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
        pillars={['gestao-metabolica', 'longevidade']}
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
              href="/healthspan"
              className="border border-petrol/10 p-6 hover:bg-petrol hover:text-cream transition group"
            >
              <p className="label-upper text-gold mb-2">{t('crossHealthspanLabel')}</p>
              <p className="text-petrol/85 group-hover:text-cream/85 text-sm">
                {t('crossHealthspanDesc')}
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
