import type { Metadata } from 'next';
import Image from 'next/image';
import { setRequestLocale, getTranslations } from 'next-intl/server';
import { BreadcrumbSchema } from '@/components/seo/breadcrumb-schema';
import { FaqSchema } from '@/components/seo/faq-schema';

type FormacaoItem = { year: string; label: string };
type FaqItem = { q: string; a: string };

export async function generateMetadata({
  params,
}: {
  params: Promise<{ locale: string }>;
}): Promise<Metadata> {
  const { locale } = await params;
  const t = await getTranslations({ locale, namespace: 'about' });
  return {
    title: t('metaTitle'),
    description: t('metaDescription'),
    alternates: {
      canonical: locale === 'en' ? '/en/the-physician' : '/o-medico',
      languages: { 'pt-BR': '/o-medico', pt: '/o-medico', en: '/en/the-physician' },
    },
  };
}

export default async function SobrePage({
  params,
}: {
  params: Promise<{ locale: string }>;
}) {
  const { locale } = await params;
  setRequestLocale(locale);
  const t = await getTranslations({ locale, namespace: 'about' });

  const formacao = t.raw('formacao') as FormacaoItem[];
  const atuacao = t.raw('atuacao') as string[];
  const areas = t.raw('areas') as string[];
  const sociedades = t.raw('sociedades') as string[];
  const faq = t.raw('faq') as FaqItem[];

  const homeLabel = locale === 'en' ? 'Home' : 'Início';

  return (
    <article>
      <BreadcrumbSchema
        items={[
          { name: homeLabel, url: '/' },
          { name: t('h1') },
        ]}
      />
      <FaqSchema items={faq} />
      {/* Cabeçalho */}
      <header className="editorial-container pt-16 md:pt-24 pb-12">
        <p className="label-meta mb-6">{t('kicker')}</p>
        <h1 className="heading-display text-[clamp(2.2rem,5vw,3.8rem)] max-w-3xl">
          {t('h1')}
        </h1>
      </header>

      {/* Corpo + aside */}
      <div className="editorial-container pb-24 grid lg:grid-cols-[1fr_360px] gap-16">
        <div className="prose-body">
          <Image
            src="/images/getulio-about-color.jpg"
            alt={t('portraitAlt')}
            width={1200}
            height={800}
            className="w-full h-auto mb-12"
            priority
          />

          {/* Ato 1 — Origem */}
          <p>{t('p1')}</p>
          <p>{t('p2')}</p>

          {/* Ato 2 — Vinte anos no hospital */}
          <p>{t('p3')}</p>
          <p>{t('p4')}</p>
          <p>{t('p5')}</p>

          {/* Ato 3 — A virada */}
          <p>{t('p6')}</p>
          <p>
            {t('p7Pre')}
            <em>{t('p7Em')}</em>
            {t('p7Post')}
          </p>

          {/* Ato 4 — O presente */}
          <p>{t('p8')}</p>
          <p>{t('p9')}</p>

          {/* Linha humana — fechamento */}
          <p className="text-ink-muted italic">{t('human')}</p>
        </div>

        <aside className="space-y-12 lg:sticky lg:top-12 self-start">
          <div>
            <p className="label-meta-lg text-bordo mb-5">{t('asideFormacao')}</p>
            <ul className="space-y-3">
              {formacao.map((f) => (
                <li key={`${f.year}-${f.label}`} className="grid grid-cols-[64px_1fr] gap-3 border-b border-rule pb-3">
                  <span className="font-sans text-sm tracking-widest text-bordo font-medium pt-1">{f.year}</span>
                  <span className="font-serif text-base text-ink leading-snug">{f.label}</span>
                </li>
              ))}
            </ul>
          </div>

          <div>
            <p className="label-meta-lg text-bordo mb-5">{t('asideAtuacao')}</p>
            <ul className="space-y-3">
              {atuacao.map((a) => (
                <li key={a} className="font-serif text-base text-ink leading-relaxed">— {a}</li>
              ))}
            </ul>
          </div>

          <div>
            <p className="label-meta-lg text-bordo mb-5">{t('asideAreas')}</p>
            <ul className="space-y-3">
              {areas.map((a) => (
                <li key={a} className="font-serif text-base text-ink leading-relaxed">— {a}</li>
              ))}
            </ul>
          </div>

          <div>
            <p className="label-meta-lg text-bordo mb-5">{t('asideSociedades')}</p>
            <ul className="space-y-3">
              {sociedades.map((s) => (
                <li key={s} className="font-serif text-base text-ink leading-relaxed">— {s}</li>
              ))}
            </ul>
          </div>

          <div>
            <p className="label-meta-lg text-bordo mb-5">{t('asideCredenciais')}</p>
            <p className="font-sans text-base text-ink">CRM-PR 21.876</p>
            <p className="font-sans text-base text-ink">RQE 16.038</p>
          </div>
        </aside>
      </div>

      {/* FAQ — visível + FAQPage schema */}
      <section className="border-t border-rule bg-paper">
        <div className="editorial-container py-16 md:py-24">
          <p className="label-meta-lg text-bordo mb-4">{t('faqKicker')}</p>
          <h2 className="heading-section text-2xl md:text-3xl text-ink max-w-3xl mb-12">
            {t('faqH2')}
          </h2>
          <div className="divide-y divide-rule">
            {faq.map((item) => (
              <details
                key={item.q}
                className="group py-5 md:py-6 [&_a]:link-text [&_p]:mb-3 last:[&_p]:mb-0 [&_p]:font-serif [&_p]:text-ink-soft [&_p]:leading-relaxed"
              >
                <summary className="cursor-pointer list-none flex items-baseline justify-between gap-4">
                  <h3 className="font-serif text-lg md:text-xl text-ink leading-snug">
                    {item.q}
                  </h3>
                  <span
                    aria-hidden
                    className="font-sans text-bordo text-xl shrink-0 transition-transform group-open:rotate-45"
                  >
                    +
                  </span>
                </summary>
                <div
                  className="mt-4 max-w-3xl"
                  dangerouslySetInnerHTML={{ __html: item.a }}
                />
              </details>
            ))}
          </div>
        </div>
      </section>
    </article>
  );
}
