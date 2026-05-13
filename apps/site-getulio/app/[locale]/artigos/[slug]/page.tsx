import type { Metadata } from 'next';
import { notFound } from 'next/navigation';
import { setRequestLocale, getTranslations } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { Breadcrumbs } from '@/components/layout/breadcrumbs';
import { MdxContent } from '@/components/blog/mdx-content';
import {
  getAllPostsFull,
  getPost,
  pillarLabels,
  plenyaMirrorUrl,
} from '@/lib/blog';
import { ArticleSchema } from '@/components/seo/article-schema';
import { BreadcrumbSchema } from '@/components/seo/breadcrumb-schema';
import { EducationalNotice } from '@/components/legal/educational-notice';

export async function generateStaticParams() {
  // Geramos slug params a partir do PT (fonte mais ampla); EN espelha o
  // mesmo conjunto. Quando um post EN não existir, o helper getPost
  // faz fallback pra PT silenciosamente.
  const all = await getAllPostsFull('pt');
  return all.map((p) => ({ slug: p.slug }));
}

export async function generateMetadata({
  params,
}: {
  params: Promise<{ locale: string; slug: string }>;
}): Promise<Metadata> {
  const { locale, slug } = await params;
  const post = await getPost(slug, locale);
  if (!post) return {};
  // Canonical aponta pra si próprio — versão editorial em 1a pessoa.
  // Cisão por audiência: Plenya tem versão clínica/institucional do mesmo tema.
  const localPath = locale === 'en' ? `/en/articles/${slug}` : `/artigos/${slug}`;
  const localUrl = `https://drgetulioamaralfilho.com.br${localPath}`;
  return {
    title: post.title,
    description: post.excerpt,
    alternates: {
      canonical: localPath,
      languages: {
        'pt-BR': `/artigos/${slug}`,
        pt: `/artigos/${slug}`,
        en: `/en/articles/${slug}`,
      },
    },
    openGraph: {
      type: 'article',
      url: localUrl,
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
  const post = await getPost(slug, locale);
  if (!post) notFound();
  const t = await getTranslations({ locale, namespace: 'escritos' });
  const labels = pillarLabels(locale);
  const dateLocale = t('dateLocale');

  const all = await getAllPostsFull(locale);
  const related = all
    .filter((p) => p.slug !== post.slug && p.pillar === post.pillar)
    .slice(0, 3);
  // URL local (canonical agora aponta pra si próprio).
  const localPath = locale === 'en' ? `/en/articles/${post.slug}` : `/artigos/${post.slug}`;
  const canonicalUrl = `https://drgetulioamaralfilho.com.br${localPath}`;
  // Cross-link recíproco pra versão Plenya (clínica/institucional).
  const plenyaUrl = plenyaMirrorUrl(post.slug, locale);

  return (
    <article>
      <ArticleSchema
        title={post.title}
        description={post.excerpt}
        slug={post.slug}
        date={post.date}
        tag={labels[post.pillar]}
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
          </p>
        </div>
      </section>

      <aside className="editorial-narrow py-10 border-t border-bordo/20 text-center">
        <p className="font-serif text-lg md:text-xl text-ink-soft leading-relaxed">
          {t('crossLinkPrompt')}{' '}
          <a
            href={plenyaUrl}
            target="_blank"
            rel="noreferrer"
            className="text-bordo font-medium underline decoration-bordo/40 underline-offset-[6px] decoration-2 hover:decoration-bordo hover:bg-bordo/5 transition-colors px-0.5"
          >
            {t('crossLinkInvite')}
            <span aria-hidden> →</span>
          </a>
        </p>
      </aside>

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
