import type { Metadata } from 'next';
import Image from 'next/image';
import { setRequestLocale, getTranslations } from 'next-intl/server';
import { ClinicsSchema } from '@/components/seo/clinics-schema';
import { BreadcrumbSchema } from '@/components/seo/breadcrumb-schema';

type Clinic = {
  name: string;
  role: string;
  body: string;
  address: string;
  href: string | null;
};

export async function generateMetadata({
  params,
}: {
  params: Promise<{ locale: string }>;
}): Promise<Metadata> {
  const { locale } = await params;
  const t = await getTranslations({ locale, namespace: 'ondeAtendo' });
  return {
    title: t('metaTitle'),
    description: t('metaDescription'),
    alternates: {
      canonical: locale === 'en' ? '/en/onde-atendo' : '/onde-atendo',
      languages: { 'pt-BR': '/onde-atendo', en: '/en/onde-atendo' },
    },
  };
}

export default async function OndeAtendoPage({
  params,
}: {
  params: Promise<{ locale: string }>;
}) {
  const { locale } = await params;
  setRequestLocale(locale);
  const t = await getTranslations({ locale, namespace: 'ondeAtendo' });
  const tCommon = await getTranslations({ locale, namespace: 'common' });
  const clinicas = t.raw('clinicas') as Clinic[];

  return (
    <article>
      <ClinicsSchema locale={locale} />
      <BreadcrumbSchema
        items={[
          { name: t('breadcrumbHome'), url: '/' },
          { name: t('breadcrumbCurrent') },
        ]}
      />
      <header className="editorial-container pt-16 md:pt-24 pb-12">
        <p className="label-meta mb-6">{t('kicker')}</p>
        <h1 className="heading-display text-[clamp(2.2rem,5vw,3.8rem)] max-w-3xl">
          {t('h1')}
        </h1>
        <p className="prose-body mt-8 max-w-2xl">{t('lead')}</p>
      </header>

      <section className="editorial-container pb-12">
        <div className="relative aspect-[16/10] w-full overflow-hidden">
          <Image
            src="/images/getulio-clinico.jpg"
            alt={t('portraitAlt')}
            fill
            className="object-cover object-top"
            sizes="(min-width: 1024px) 1100px, 100vw"
          />
        </div>
      </section>

      <section className="editorial-container pb-24 space-y-16">
        {clinicas.map((c) => (
          <div
            key={c.name}
            className="border-t border-rule pt-12 grid md:grid-cols-[260px_1fr] gap-10 md:gap-16"
          >
            <div className="space-y-4">
              <p className="label-meta-lg text-bordo">{c.role}</p>
            </div>
            <div className="space-y-5">
              <h2 className="heading-section text-3xl md:text-4xl text-ink">{c.name}</h2>
              <p className="font-serif text-lg text-ink-soft leading-relaxed max-w-2xl">{c.body}</p>
              <p className="font-sans text-sm text-ink-muted">{c.address}</p>
              {c.href && (
                <a
                  href={c.href}
                  target="_blank"
                  rel="noreferrer"
                  className="link-text inline-block font-sans text-sm"
                >
                  {tCommon('visitSite')}
                </a>
              )}
            </div>
          </div>
        ))}
      </section>

      <section className="border-t border-rule bg-paper">
        <div className="editorial-container py-20 grid md:grid-cols-[1fr_2fr] gap-12 items-start">
          <p className="label-meta">{t('scheduleKicker')}</p>
          <div className="space-y-4 font-serif text-ink-soft">
            <p>{t('scheduleBody')}</p>
            <p className="font-sans text-sm">
              <a href="mailto:contato@drgetulioamaralfilho.com.br" className="link-text">
                contato@drgetulioamaralfilho.com.br
              </a>
            </p>
          </div>
        </div>
      </section>
    </article>
  );
}
