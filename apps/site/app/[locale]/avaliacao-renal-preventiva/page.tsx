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
  const t = await getTranslations({ locale, namespace: 'renal' });
  return {
    title: t('metaTitle'),
    description: t('metaDescription'),
    alternates: {
      canonical: '/avaliacao-renal-preventiva',
      languages: {
        'pt-BR': '/avaliacao-renal-preventiva',
        en: '/en/preventive-kidney-assessment',
        'x-default': '/avaliacao-renal-preventiva',
      },
    },
    openGraph: {
      title: t('ogTitle'),
      description: t('ogDescription'),
      type: 'article',
    },
  };
}

export default async function RenalPage({ params }: { params: Params }) {
  const { locale } = await params;
  setRequestLocale(locale);
  const t = await getTranslations('renal');

  const faq = [1, 2, 3, 4, 5, 6].map((n) => ({
    q: t(`faq${n}Q` as 'faq1Q'),
    a: t(`faq${n}A` as 'faq1A'),
  }));

  const panel = [1, 2, 3, 4, 5, 6, 7].map((n) => t(`asidePanel${n}` as 'asidePanel1'));
  const who = [1, 2, 3, 4, 5].map((n) => t(`asideWho${n}` as 'asideWho1'));

  return (
    <>
      <FaqSchema items={faq} />
      <MedicalWebPageSchema
        name={t('schemaName')}
        description={t('schemaDescription')}
        path="/avaliacao-renal-preventiva"
        about="Chronic Kidney Disease"
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
            <p className="label-upper text-gold">{t('whyLabel')}</p>
            <p>{t('whyP1')}</p>
            <p>{t('whyP2')}</p>
            <p>
              {t('whyP3Part1')}
              <strong>{t('whyP3Strong')}</strong>
              {t('whyP3Part2')}
            </p>
          </div>

          <aside className="space-y-6">
            <div className="border border-petrol/10 p-6 space-y-3">
              <p className="label-upper text-gold">{t('asidePanelLabel')}</p>
              <ul className="text-petrol/75 text-sm space-y-1.5 leading-relaxed">
                {panel.map((p) => (
                  <li key={p}>· {p}</li>
                ))}
              </ul>
            </div>
            <div className="border border-petrol/10 p-6 space-y-3">
              <p className="label-upper text-gold">{t('asideWhoLabel')}</p>
              <ul className="text-petrol/75 text-sm space-y-1.5 leading-relaxed">
                {who.map((w) => (
                  <li key={w}>· {w}</li>
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

      <RelatedBlogPosts title={t('relatedTitle')} pillars={['gestao-metabolica']} limit={3} locale={isLocale(locale) ? locale : defaultLocale} />

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
              href="/dr-getulio"
              className="border border-petrol/10 p-6 hover:bg-petrol hover:text-cream transition group"
            >
              <p className="label-upper text-gold mb-2">{t('crossDrLabel')}</p>
              <p className="text-petrol/85 group-hover:text-cream/85 text-sm">
                {t('crossDrDesc')}
              </p>
            </Link>
            <Link
              href="/checkup-longevidade"
              className="border border-petrol/10 p-6 hover:bg-petrol hover:text-cream transition group"
            >
              <p className="label-upper text-gold mb-2">{t('crossCheckupLabel')}</p>
              <p className="text-petrol/85 group-hover:text-cream/85 text-sm">
                {t('crossCheckupDesc')}
              </p>
            </Link>
          </div>
        </div>
      </section>
    </>
  );
}
