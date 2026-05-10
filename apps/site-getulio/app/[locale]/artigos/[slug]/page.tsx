import type { Metadata } from 'next';
import { notFound } from 'next/navigation';
import { setRequestLocale, getTranslations } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { Breadcrumbs } from '@/components/layout/breadcrumbs';
import { MdxContent } from '@/components/blog/mdx-content';
import {
  getAllPlenyaPostsFull,
  getPlenyaPost,
  pillarLabels,
  plenyaBlogBase,
} from '@/lib/plenya-blog';
import { ArticleSchema } from '@/components/seo/article-schema';
import { BreadcrumbSchema } from '@/components/seo/breadcrumb-schema';
import { EducationalNotice } from '@/components/legal/educational-notice';

export async function generateStaticParams() {
  // Geramos slug params a partir do PT (fonte mais ampla); EN espelha o
  // mesmo conjunto. Quando um post EN não existir, o helper getPlenyaPost
  // faz fallback pra PT silenciosamente.
  const all = await getAllPlenyaPostsFull('pt');
  return all.map((p) => ({ slug: p.slug }));
}

export async function generateMetadata({
  params,
}: {
  params: Promise<{ locale: string; slug: string }>;
}): Promise<Metadata> {
  const { locale, slug } = await params;
  const post = await getPlenyaPost(slug, locale);
  if (!post) return {};
  const blogBase = plenyaBlogBase(locale);
  // Canonical aponta pra Plenya (fonte oficial). Google atribui authority lá.
  const canonical = `${blogBase}/${slug}`;
  return {
    title: post.title,
    description: post.excerpt,
    alternates: { canonical },
    openGraph: {
      type: 'article',
      url: canonical,
      title: post.title,
      description: post.excerpt,
      publishedTime: post.date,
      modifiedTime: post.updated ?? post.date,
      authors: ['Dr. Getúlio Amaral Filho'],
      tags: post.tags,
      images: post.cover ? [`https://plenyasaude.com.br${post.cover}`] : ['/images/getulio-square.jpg'],
    },
    twitter: {
      card: 'summary_large_image',
      title: post.title,
      description: post.excerpt,
      images: post.cover ? [`https://plenyasaude.com.br${post.cover}`] : ['/images/getulio-square.jpg'],
    },
  };
}

export default async function EscritoPage({
  params,
}: {
  params: Promise<{ locale: string; slug: string }>;
}) {
  const { locale, slug } = await params;
  setRequestLocale(locale);
  const post = await getPlenyaPost(slug, locale);
  if (!post) notFound();
  const t = await getTranslations({ locale, namespace: 'escritos' });
  const labels = pillarLabels(locale);
  const blogBase = plenyaBlogBase(locale);
  const dateLocale = t('dateLocale');

  const all = await getAllPlenyaPostsFull(locale);
  const related = all
    .filter((p) => p.slug !== post.slug && p.pillar === post.pillar)
    .slice(0, 3);
  const canonicalUrl = `${blogBase}/${post.slug}`;

  return (
    <article>
      <ArticleSchema
        title={post.title}
        description={post.excerpt}
        slug={post.slug}
        date={post.date}
        tag={labels[post.pillar]}
        canonicalUrl={canonicalUrl}
        image={post.cover ? `https://plenyasaude.com.br${post.cover}` : undefined}
        locale={locale}
      />
      <BreadcrumbSchema
        items={[
          { name: t('detailBreadcrumbHome'), url: '/' },
          { name: t('detailBreadcrumbList'), url: locale === 'en' ? '/en/articles' : '/artigos' },
          { name: post.title },
        ]}
      />
      <header className="editorial-container pt-12 md:pt-16 pb-8">
        <Breadcrumbs
          items={[
            { label: t('detailBreadcrumbHome'), href: '/' },
            { label: t('detailBreadcrumbList'), href: '/artigos' },
            { label: post.title },
          ]}
        />
      </header>

      <div className="editorial-narrow pb-2">
        <p className="font-sans text-xs text-ink-muted bg-paper border-l-2 border-bordo px-4 py-3">
          {t('detailBannerPre')}<strong>{t('detailBlogStrong')}</strong>.{' '}
          <a
            href={canonicalUrl}
            target="_blank"
            rel="noreferrer"
            className="link-text"
          >
            {t('detailBannerLink')}
          </a>
        </p>
      </div>

      <section className="editorial-narrow pb-12">
        <div className="space-y-8">
          <div className="flex items-center gap-4 flex-wrap">
            <span className="label-meta-lg text-bordo">{labels[post.pillar]}</span>
            <span className="label-meta-lg text-ink-muted">
              {formatDate(post.date, dateLocale)} · {post.readingMinutes} {t('detailReadingSuffix')}
            </span>
          </div>
          <h1 className="heading-display text-[clamp(2.6rem,5.5vw,4.4rem)]">{post.title}</h1>
          <p className="font-serif italic text-ink-soft text-2xl leading-relaxed">
            {post.excerpt}
          </p>
        </div>
      </section>

      {post.cover && (
        <section className="editorial-narrow pb-8">
          {/* eslint-disable-next-line @next/next/no-img-element */}
          <img
            src={`https://plenyasaude.com.br${post.cover}`}
            alt={post.title}
            className="w-full h-auto"
          />
        </section>
      )}

      <section className="editorial-narrow pb-16">
        <MdxContent source={post.content} />
      </section>

      {post.references.length > 0 && (
        <section className="editorial-narrow pb-12">
          <div className="border-t-2 border-bordo/40 pt-8">
            <p className="label-meta-lg text-bordo mb-6">{t('detailReferencesKicker')}</p>
            <ol className="list-decimal pl-6 space-y-3 font-serif text-base text-ink-soft leading-relaxed">
              {post.references.map((ref, i) => (
                <li key={ref.url ?? `${i}-${ref.label}`} className="pl-1">
                  {ref.url ? (
                    <a href={ref.url} target="_blank" rel="noreferrer" className="link-text">
                      {ref.label}
                    </a>
                  ) : (
                    <span>{ref.label}</span>
                  )}
                </li>
              ))}
            </ol>
          </div>
        </section>
      )}

      <section className="border-t border-rule">
        <div className="editorial-narrow py-8">
          <p className="font-sans text-sm text-ink-muted">
            <strong className="text-ink">{t('detailReviewStrong')}</strong> {t('detailReviewBody')}
            <a href={canonicalUrl} target="_blank" rel="noreferrer" className="link-text">
              {locale === 'en' ? 'plenyasaude.com.br/en/blog' : 'plenyasaude.com.br/blog'}
            </a>
            .
          </p>
        </div>
      </section>

      <EducationalNotice />

      {related.length > 0 && (
        <section className="border-t border-rule">
          <div className="editorial-container py-16">
            <p className="label-meta mb-8">{t('detailRelatedKicker')}</p>
            <ul className="grid md:grid-cols-3 gap-8">
              {related.map((p) => (
                <li key={p.slug}>
                  <Link
                    href={{ pathname: '/artigos/[slug]', params: { slug: p.slug } }}
                    rel="related"
                    className="block group space-y-3"
                  >
                    <span className="label-meta text-bordo">
                      {labels[p.pillar]}
                    </span>
                    <h3 className="font-serif text-lg text-ink group-hover:text-bordo transition-colors">
                      {p.title}
                    </h3>
                    <p className="font-serif text-sm text-ink-muted leading-relaxed">
                      {p.excerpt}
                    </p>
                  </Link>
                </li>
              ))}
            </ul>
          </div>
        </section>
      )}
    </article>
  );
}

function formatDate(d: string, locale: string) {
  const dt = new Date(d);
  return dt.toLocaleDateString(locale, { year: 'numeric', month: 'long' });
}
