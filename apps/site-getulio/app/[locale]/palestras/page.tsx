import type { Metadata } from 'next';
import Image from 'next/image';
import { setRequestLocale, getTranslations } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { getAllLectures, getAudienceLabel, localizedLecture, sortAudience } from '@/lib/lectures';
import { EducationalNotice } from '@/components/legal/educational-notice';
import { BreadcrumbSchema } from '@/components/seo/breadcrumb-schema';

export async function generateMetadata({
  params,
}: {
  params: Promise<{ locale: string }>;
}): Promise<Metadata> {
  const { locale } = await params;
  const t = await getTranslations({ locale, namespace: 'palestras' });
  return {
    title: t('metaTitle'),
    description: t('metaDescription'),
    alternates: {
      canonical: locale === 'en' ? '/en/palestras' : '/palestras',
      languages: { 'pt-BR': '/palestras', en: '/en/palestras' },
    },
  };
}

export default async function PalestrasPage({
  params,
}: {
  params: Promise<{ locale: string }>;
}) {
  const { locale } = await params;
  setRequestLocale(locale);
  const t = await getTranslations({ locale, namespace: 'palestras' });
  const lectures = await getAllLectures();

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

      <section className="editorial-container pb-16 md:pb-20">
        <div className="relative aspect-[2/3] sm:aspect-[3/4] md:aspect-[4/5] lg:aspect-[3/2] w-full overflow-hidden bg-paper">
          <Image
            src="/images/getulio-palestrante-arvore.jpg"
            alt={t('headerImageAlt')}
            fill
            priority
            className="object-cover object-[center_top] lg:object-[center_25%]"
            sizes="(min-width: 1024px) 1100px, 100vw"
          />
        </div>
      </section>

      <section className="editorial-container pb-24">
        <ul className="border-t border-rule">
          {lectures.map((l) => {
            const loc = localizedLecture(l, locale);
            return (
              <li key={l.slug} className="border-b border-rule">
                <Link
                  href={`/palestras/${l.slug}`}
                  className="block py-10 grid md:grid-cols-[1fr_280px] gap-8 group"
                >
                  <div className="space-y-4">
                    <div className="flex items-center gap-3 flex-wrap">
                      {l.anchor && <span className="label-meta text-bordo">{t('anchorBadge')}</span>}
                      <span className="label-meta text-ink-muted">{loc.duration}</span>
                    </div>
                    <h2 className="heading-section text-2xl md:text-3xl group-hover:text-bordo transition-colors">
                      {loc.title}
                    </h2>
                    {loc.subtitle && (
                      <p className="font-serif italic text-ink-muted text-lg">{loc.subtitle}</p>
                    )}
                    <p className="font-serif text-ink-soft leading-relaxed max-w-2xl">{loc.excerpt}</p>
                  </div>
                  <div className="space-y-3 md:text-right">
                    <p className="label-meta">{t('audienceLabel')}</p>
                    <ul className="space-y-1">
                      {sortAudience(l.audience).map((a) => (
                        <li key={a} className="font-serif text-sm text-ink-soft">
                          {getAudienceLabel(a, locale)}
                        </li>
                      ))}
                    </ul>
                    <p className="label-meta pt-3">{t('formatLabel')}</p>
                    <p className="font-serif text-sm text-ink-soft">{loc.format}</p>
                  </div>
                </Link>
              </li>
            );
          })}
        </ul>
      </section>

      <EducationalNotice />

      <section className="border-t border-rule bg-paper">
        <div className="editorial-container py-20 grid md:grid-cols-[1fr_2fr] gap-12 items-start">
          <p className="label-meta">{t('inviteKicker')}</p>
          <div className="space-y-4 font-serif text-ink-soft">
            <p>{t('inviteBody')}</p>
            <p className="font-sans text-sm">
              <a href="mailto:palestras@drgetulioamaralfilho.com.br" className="link-text">
                palestras@drgetulioamaralfilho.com.br
              </a>
              <br />
              <span className="text-ink-muted text-xs">{t('inviteResponse')}</span>
            </p>
          </div>
        </div>
      </section>
    </article>
  );
}
