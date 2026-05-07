import type { Metadata } from 'next';
import Image from 'next/image';
import { setRequestLocale, getTranslations } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import {
  absoluteCover,
  getPlenyaPostsByGetulio,
  pillarLabels,
  plenyaBlogBase,
} from '@/lib/plenya-blog';

export async function generateMetadata({
  params,
}: {
  params: Promise<{ locale: string }>;
}): Promise<Metadata> {
  const { locale } = await params;
  const t = await getTranslations({ locale, namespace: 'escritos' });
  return {
    title: t('metaTitle'),
    description: t('metaDescription'),
    alternates: {
      canonical: locale === 'en' ? '/en/escritos' : '/escritos',
      languages: { 'pt-BR': '/escritos', en: '/en/escritos' },
    },
  };
}

export default async function EscritosPage({
  params,
  searchParams,
}: {
  params: Promise<{ locale: string }>;
  searchParams: Promise<{ pilar?: string }>;
}) {
  const { locale } = await params;
  const { pilar: pilarParam } = await searchParams;
  setRequestLocale(locale);
  const t = await getTranslations({ locale, namespace: 'escritos' });
  const labels = pillarLabels(locale);
  const pillarFilters = Object.entries(labels) as [keyof typeof labels, string][];
  const all = await getPlenyaPostsByGetulio(locale);
  const activePilar = pillarFilters.find(([k]) => k === pilarParam)?.[0];
  const posts = activePilar ? all.filter((p) => p.pillar === activePilar) : all;
  const [featured, ...rest] = posts;
  const dateLocale = t('dateLocale');
  const blogBase = plenyaBlogBase(locale);

  return (
    <article>
      <header className="editorial-container pt-16 md:pt-24 pb-12">
        <p className="label-meta mb-6">{t('kicker')}</p>
        <h1 className="heading-display text-[clamp(2.2rem,5vw,3.8rem)] max-w-3xl">
          {t('h1')}
        </h1>
        <p className="prose-body mt-8 max-w-2xl">
          {t('leadPre')}<em>{t('leadEm')}</em>{t('leadPost')}
        </p>
      </header>

      <section className="editorial-container pb-8">
        <div className="flex flex-wrap items-center gap-x-6 gap-y-3 border-y border-rule py-4">
          <Link
            href="/escritos"
            className={`font-sans text-sm tracking-wide ${!activePilar ? 'text-bordo' : 'text-ink-muted hover:text-ink'} transition-colors`}
          >
            {t('filterAll')}
          </Link>
          {pillarFilters.map(([key, label]) => (
            <Link
              key={key}
              href={`/escritos?pilar=${encodeURIComponent(key)}`}
              className={`font-sans text-sm tracking-wide ${activePilar === key ? 'text-bordo' : 'text-ink-muted hover:text-ink'} transition-colors`}
            >
              {label}
            </Link>
          ))}
        </div>
      </section>

      <section className="editorial-container pb-24">
        {posts.length === 0 ? (
          <p className="font-serif text-ink-muted py-8">{t('emptyState')}</p>
        ) : (
          <div className="space-y-20 md:space-y-24">
            {featured && (
              <Link
                href={`/escritos/${featured.slug}`}
                className="group block grid md:grid-cols-2 gap-8 md:gap-12 items-center"
              >
                {featured.cover && (
                  <div className="relative aspect-[4/3] overflow-hidden bg-paper order-1">
                    <Image
                      src={absoluteCover(featured.cover)!}
                      alt={featured.title}
                      fill
                      priority
                      sizes="(min-width: 768px) 560px, 100vw"
                      className="object-cover transition-transform duration-700 group-hover:scale-[1.02]"
                    />
                  </div>
                )}
                <div className="space-y-4 order-2">
                  <div className="flex items-center gap-3 label-meta">
                    <span className="text-bordo">{labels[featured.pillar]}</span>
                    <span className="text-ink-muted">·</span>
                    <time className="text-ink-muted" dateTime={featured.date}>
                      {formatDate(featured.date, dateLocale)}
                    </time>
                  </div>
                  <h2 className="font-serif text-3xl md:text-4xl leading-tight text-ink group-hover:text-bordo transition-colors">
                    {featured.title}
                  </h2>
                  <p className="font-serif text-lg text-ink-soft leading-relaxed">
                    {featured.excerpt}
                  </p>
                  <p className="label-meta text-bordo pt-2">{t('ctaRead')}</p>
                </div>
              </Link>
            )}

            {rest.length > 0 && (
              <div className="grid sm:grid-cols-2 lg:grid-cols-3 gap-x-8 gap-y-14 border-t border-rule pt-12">
                {rest.map((p) => (
                  <Link
                    key={p.slug}
                    href={`/escritos/${p.slug}`}
                    className="group block space-y-4"
                  >
                    {p.cover && (
                      <div className="relative aspect-[16/10] overflow-hidden bg-paper">
                        <Image
                          src={absoluteCover(p.cover)!}
                          alt={p.title}
                          fill
                          sizes="(min-width: 1024px) 360px, (min-width: 640px) 50vw, 100vw"
                          className="object-cover transition-transform duration-700 group-hover:scale-[1.02]"
                        />
                      </div>
                    )}
                    <div className="flex items-center gap-3 label-meta">
                      <span className="text-bordo">{labels[p.pillar]}</span>
                      <span className="text-ink-muted">·</span>
                      <time className="text-ink-muted" dateTime={p.date}>
                        {formatDate(p.date, dateLocale)}
                      </time>
                    </div>
                    <h3 className="font-serif text-xl md:text-2xl leading-tight text-ink group-hover:text-bordo transition-colors">
                      {p.title}
                    </h3>
                    <p className="font-serif text-ink-soft leading-relaxed">{p.excerpt}</p>
                  </Link>
                ))}
              </div>
            )}
          </div>
        )}

        <p className="font-sans text-xs text-ink-muted mt-16 text-center">
          {t('footerSourcePre')}
          <a href={blogBase} target="_blank" rel="noreferrer" className="link-text">
            {t('footerSourceLink')}
          </a>
          {t('footerSourcePost')}
        </p>
      </section>
    </article>
  );
}

function formatDate(d: string, locale: string) {
  const dt = new Date(d);
  return dt.toLocaleDateString(locale, { year: 'numeric', month: 'short' });
}
