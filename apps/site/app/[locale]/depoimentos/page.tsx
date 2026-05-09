import type { Metadata } from 'next';
import { setRequestLocale, getTranslations } from 'next-intl/server';
import { isLocale, defaultLocale } from '@/lib/i18n/config';
import { Link } from '@/lib/i18n/navigation';
import { getAllTestimonials } from '@/lib/testimonials';
import { TestimonialCard } from '@/components/testimonials/testimonial-card';

type Params = Promise<{ locale: string }>;

export async function generateMetadata({ params }: { params: Params }): Promise<Metadata> {
  const { locale } = await params;
  const t = await getTranslations({ locale, namespace: 'testimonials' });
  return {
    title: t('metaTitle'),
    description: t('metaDescription'),
    alternates: {
      canonical: locale === 'en' ? '/en/testimonials' : '/depoimentos',
      languages: {
        'pt-BR': '/depoimentos',
        en: '/en/testimonials',
        'x-default': '/depoimentos',
      },
    },
  };
}

export default async function TestimonialsPage({ params }: { params: Params }) {
  const { locale: rawLocale } = await params;
  const locale = isLocale(rawLocale) ? rawLocale : defaultLocale;
  setRequestLocale(locale);
  const t = await getTranslations('testimonials');

  const items = await getAllTestimonials();

  return (
    <>
      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-24 md:pt-40 md:pb-32">
          <p className="label-upper text-gold mb-6">{t('heroLabel')}</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,5rem)] text-cream max-w-xl">
            {t('heroTitle')}
          </h1>
          <p className="text-cream/70 text-lg mt-6 max-w-lg">{t('heroP')}</p>
        </div>
      </section>

      <section className="bg-cream">
        <div className="site-container section">
          {items.length ? (
            <div className="grid gap-8 md:grid-cols-2 lg:grid-cols-3">
              {items.map((tm) => (
                <TestimonialCard key={tm.slug} testimonial={tm} />
              ))}
            </div>
          ) : (
            <p className="text-petrol/50 text-center label-upper">{t('emptyState')}</p>
          )}
        </div>
      </section>

      <section className="bg-paper">
        <div className="site-container section max-w-3xl space-y-4">
          <p className="label-upper text-petrol/55">{t('noteLabel')}</p>
          <p className="text-petrol/75 leading-relaxed text-sm">{t('noteText')}</p>
        </div>
      </section>

      <section className="bg-cream border-t border-petrol/10">
        <div className="site-container section text-center space-y-6">
          <p className="heading-section text-petrol text-2xl md:text-3xl">{t('ctaTitle')}</p>
          <Link href="/contato" className="btn-gold">
            {t('ctaButton')}
          </Link>
        </div>
      </section>
    </>
  );
}
