import type { Metadata } from 'next';
import { setRequestLocale, getTranslations } from 'next-intl/server';
import { BreadcrumbSchema } from '@/components/seo/breadcrumb-schema';

type Cargo = { period: string; title: string; org: string; body: string };

export async function generateMetadata({
  params,
}: {
  params: Promise<{ locale: string }>;
}): Promise<Metadata> {
  const { locale } = await params;
  const t = await getTranslations({ locale, namespace: 'ensino' });
  return {
    title: t('metaTitle'),
    description: t('metaDescription'),
    alternates: {
      canonical: locale === 'en' ? '/en/teaching' : '/ensino',
      languages: { 'pt-BR': '/ensino', pt: '/ensino', en: '/en/teaching' },
    },
  };
}

export default async function EnsinoPage({
  params,
}: {
  params: Promise<{ locale: string }>;
}) {
  const { locale } = await params;
  setRequestLocale(locale);
  const t = await getTranslations({ locale, namespace: 'ensino' });
  const cargos = t.raw('cargos') as Cargo[];

  const homeLabel = locale === 'en' ? 'Home' : 'Início';

  return (
    <article>
      <BreadcrumbSchema
        items={[{ name: homeLabel, url: '/' }, { name: t('h1') }]}
      />
      <header className="editorial-container pt-16 md:pt-24 pb-12">
        <p className="label-meta mb-6">{t('kicker')}</p>
        <h1 className="heading-display text-[clamp(2.2rem,5vw,3.8rem)] max-w-3xl">
          {t('h1')}
        </h1>
        <p className="prose-body mt-8 max-w-2xl">{t('lead')}</p>
      </header>

      <section className="editorial-container pb-20">
        <ul className="border-t border-rule">
          {cargos.map((c) => (
            <li key={c.title} className="border-b border-rule py-10 grid md:grid-cols-[180px_1fr] gap-6 md:gap-12">
              <p className="label-meta text-bordo">{c.period}</p>
              <div className="space-y-3">
                <h2 className="heading-section text-xl md:text-2xl">{c.title}</h2>
                <p className="label-meta text-ink-muted">{c.org}</p>
                <p className="font-serif text-ink-soft leading-relaxed max-w-2xl">{c.body}</p>
              </div>
            </li>
          ))}
        </ul>
      </section>

      <section className="border-t border-rule bg-paper">
        <div className="editorial-container py-20 max-w-3xl">
          <p className="label-meta mb-6">{t('philosophyKicker')}</p>
          <div className="prose-body space-y-6">
            <p>{t('philosophy1')}</p>
            <p>{t('philosophy2')}</p>
          </div>
        </div>
      </section>
    </article>
  );
}
