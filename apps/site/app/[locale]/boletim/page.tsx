import type { Metadata } from 'next';
import { setRequestLocale, getTranslations } from 'next-intl/server';
import { NewsletterInline } from '@/components/blog/newsletter-inline';
import { Link } from '@/lib/i18n/navigation';

type Params = Promise<{ locale: string }>;

export async function generateMetadata({ params }: { params: Params }): Promise<Metadata> {
  const { locale } = await params;
  const t = await getTranslations({ locale, namespace: 'newsletter' });
  return {
    title: t('metaTitle'),
    description: t('metaDescription'),
    alternates: {
      canonical: locale === 'en' ? '/en/newsletter' : '/boletim',
      languages: {
        'pt-BR': '/boletim',
        en: '/en/newsletter',
        'x-default': '/boletim',
      },
    },
  };
}

export default async function BoletimPage({ params }: { params: Params }) {
  const { locale } = await params;
  setRequestLocale(locale);
  const t = await getTranslations('newsletter');

  const promessas = [1, 2, 3, 4].map((n) => ({
    label: t(`p${n}Label` as 'p1Label'),
    title: t(`p${n}Title` as 'p1Title'),
    body: t(`p${n}Body` as 'p1Body'),
  }));

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

      <section className="bg-cream">
        <div className="site-container section grid md:grid-cols-2 gap-x-12 gap-y-10">
          {promessas.map((p) => (
            <div key={p.label} className="border-t border-petrol/15 pt-6 space-y-2">
              <p className="label-upper text-gold">{p.label}</p>
              <h2 className="heading-section text-petrol text-xl md:text-2xl">{p.title}</h2>
              <p className="text-petrol/75 leading-relaxed">{p.body}</p>
            </div>
          ))}
        </div>
      </section>

      <section className="bg-paper">
        <div className="site-container section max-w-2xl">
          <NewsletterInline />
        </div>
      </section>

      <section className="bg-cream">
        <div className="site-container section text-center space-y-6 max-w-2xl mx-auto">
          <p className="label-upper text-gold">{t('ctaLabel')}</p>
          <p className="heading-section text-petrol text-2xl md:text-3xl">{t('ctaTitle')}</p>
          <p>
            <Link href="/blog" className="btn-gold">
              {t('ctaButton')}
            </Link>
          </p>
        </div>
      </section>
    </>
  );
}
