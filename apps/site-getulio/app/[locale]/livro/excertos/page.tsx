import type { Metadata } from 'next';
import { setRequestLocale, getTranslations } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { Breadcrumbs } from '@/components/layout/breadcrumbs';
import { EducationalNotice } from '@/components/legal/educational-notice';
import { BreadcrumbSchema } from '@/components/seo/breadcrumb-schema';

type Trecho = { cap: string; citacao: string };

export async function generateMetadata({
  params,
}: {
  params: Promise<{ locale: string }>;
}): Promise<Metadata> {
  const { locale } = await params;
  const t = await getTranslations({ locale, namespace: 'book' });
  return {
    title: t('excerptsMetaTitle'),
    description: t('excerptsMetaDescription'),
    alternates: {
      canonical: locale === 'en' ? '/en/livro/excertos' : '/livro/excertos',
      languages: { 'pt-BR': '/livro/excertos', en: '/en/livro/excertos' },
    },
  };
}

export default async function ExcertosPage({
  params,
}: {
  params: Promise<{ locale: string }>;
}) {
  const { locale } = await params;
  setRequestLocale(locale);
  const t = await getTranslations({ locale, namespace: 'book' });
  const trechos = t.raw('trechos') as Trecho[];
  const amazonUrl = locale === 'en' ? t('amazonUrlEn') : t('amazonUrl');
  const hotmartUrl = `${t('hotmartUrl')}?src=excertos`;

  return (
    <article>
      <BreadcrumbSchema
        items={[
          { name: locale === 'en' ? 'Home' : 'Início', url: '/' },
          { name: locale === 'en' ? 'Book' : 'Livro', url: '/livro' },
          { name: t('excerptsKicker') },
        ]}
      />

      <header className="editorial-container pt-12 md:pt-16 pb-8">
        <Breadcrumbs
          items={[
            { label: locale === 'en' ? 'Home' : 'Início', href: '/' },
            { label: locale === 'en' ? 'Book' : 'Livro', href: '/livro' },
            { label: t('excerptsKicker') },
          ]}
        />
      </header>

      <section className="editorial-narrow pb-16">
        <div className="space-y-8">
          <p className="label-meta-lg text-bordo">{t('excerptsKicker')}</p>
          <h1 className="heading-display text-[clamp(2.4rem,5vw,4rem)]">{t('excerptsH1')}</h1>
          <p className="font-serif italic text-ink-soft text-xl md:text-2xl leading-relaxed">
            {t('excerptsLead')}
          </p>
        </div>
      </section>

      <section className="border-t border-rule bg-paper">
        <div className="editorial-narrow py-20 md:py-28 space-y-24 md:space-y-32">
          {trechos.map((tr, i) => (
            <figure key={i} className="space-y-6">
              <p className="label-meta-lg text-bordo">{tr.cap}</p>
              <blockquote className="font-serif text-2xl md:text-3xl leading-snug text-ink italic border-l-2 border-bordo pl-6 md:pl-10">
                {tr.citacao}
              </blockquote>
            </figure>
          ))}
        </div>
      </section>

      <EducationalNotice />

      <section className="border-t border-rule bg-paper">
        <div className="editorial-container py-20 md:py-24">
          <div className="grid md:grid-cols-[1fr_2fr] gap-12 items-start">
            <p className="label-meta-lg text-bordo">{t('excerptsCloserKicker')}</p>
            <div className="space-y-6 max-w-xl">
              <h2 className="heading-section text-ink text-2xl md:text-3xl leading-snug">
                {t('excerptsCloserH2')}
              </h2>
              <p className="font-serif text-ink-soft leading-relaxed">{t('excerptsCloserBody')}</p>
              <div className="flex flex-wrap items-center gap-4 pt-2">
                <a
                  href={amazonUrl}
                  target="_blank"
                  rel="noreferrer"
                  className="btn-gold"
                >
                  {t('excerptsCtaAmazon')}
                </a>
                <a
                  href={hotmartUrl}
                  target="_blank"
                  rel="noreferrer"
                  className="font-sans text-sm link-text text-ink-soft hover:text-bordo"
                >
                  {t('excerptsCtaHotmart')}
                </a>
              </div>
              <p className="font-sans text-sm pt-4">
                <Link href="/livro" className="link-text text-ink-muted">
                  {t('excerptsBackToBook')}
                </Link>
              </p>
            </div>
          </div>
        </div>
      </section>
    </article>
  );
}
