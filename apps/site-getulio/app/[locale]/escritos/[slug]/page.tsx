import type { Metadata } from 'next';
import { notFound } from 'next/navigation';
import { setRequestLocale } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { Breadcrumbs } from '@/components/layout/breadcrumbs';
import { MdxContent } from '@/components/blog/mdx-content';
import {
  getAllPlenyaPostsFull,
  getPlenyaPost,
  PLENYA_BLOG_BASE,
  PLENYA_PILLAR_LABELS,
} from '@/lib/plenya-blog';
import { ArticleSchema } from '@/components/seo/article-schema';
import { BreadcrumbSchema } from '@/components/seo/breadcrumb-schema';

const BASE = 'https://drgetulioamaralfilho.com.br';

export async function generateStaticParams() {
  const all = await getAllPlenyaPostsFull();
  return all.map((p) => ({ slug: p.slug }));
}

export async function generateMetadata({
  params,
}: {
  params: Promise<{ slug: string }>;
}): Promise<Metadata> {
  const { slug } = await params;
  const post = await getPlenyaPost(slug);
  if (!post) return {};
  const canonical = `${PLENYA_BLOG_BASE}/${slug}`;
  return {
    title: post.title,
    description: post.excerpt,
    // Canonical aponta para a Plenya — fonte oficial. Google atribui authority lá.
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
  const post = await getPlenyaPost(slug);
  if (!post) notFound();

  const all = await getAllPlenyaPostsFull();
  const related = all
    .filter((p) => p.slug !== post.slug && p.pillar === post.pillar)
    .slice(0, 3);
  const canonicalUrl = `${PLENYA_BLOG_BASE}/${post.slug}`;

  return (
    <article>
      <ArticleSchema
        title={post.title}
        description={post.excerpt}
        slug={post.slug}
        date={post.date}
        tag={PLENYA_PILLAR_LABELS[post.pillar]}
        canonicalUrl={canonicalUrl}
        image={post.cover ? `https://plenyasaude.com.br${post.cover}` : undefined}
      />
      <BreadcrumbSchema
        items={[
          { name: 'Início', url: '/' },
          { name: 'Escritos', url: '/escritos' },
          { name: post.title },
        ]}
      />
      <header className="editorial-container pt-12 md:pt-16 pb-8">
        <Breadcrumbs
          items={[
            { label: 'Início', href: '/' },
            { label: 'Escritos', href: '/escritos' },
            { label: post.title },
          ]}
        />
      </header>

      {/* Banner: este texto vive originalmente no blog da Plenya */}
      <div className="editorial-narrow pb-2">
        <p className="font-sans text-xs text-ink-muted bg-paper border-l-2 border-bordo px-4 py-3">
          Publicado originalmente no <strong>Blog Plenya</strong>.{' '}
          <a
            href={canonicalUrl}
            target="_blank"
            rel="noreferrer"
            className="link-text"
          >
            Ler na fonte ↗
          </a>
        </p>
      </div>

      <section className="editorial-narrow pb-12">
        <div className="space-y-8">
          <div className="flex items-center gap-4 flex-wrap">
            <span className="label-meta text-bordo">{PLENYA_PILLAR_LABELS[post.pillar]}</span>
            <span className="label-meta text-ink-muted">
              {formatDate(post.date)} · {post.readingMinutes} min
            </span>
          </div>
          <h1 className="heading-display text-[clamp(2rem,4.5vw,3.4rem)]">{post.title}</h1>
          <p className="font-serif italic text-ink-muted text-xl leading-relaxed">
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
          <p className="label-meta text-bordo mb-4">Referências</p>
          <ol className="list-decimal pl-6 space-y-2 font-sans text-sm text-ink-soft">
            {post.references.map((ref, i) => (
              <li key={ref.url ?? `${i}-${ref.label}`}>
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
        </section>
      )}

      <section className="border-t border-rule">
        <div className="editorial-narrow py-8">
          <p className="font-sans text-sm text-ink-muted">
            <strong className="text-ink">Revisão clínica.</strong> Conteúdo médico de autoria do
            Dr. Getúlio Amaral Filho · CRM-PR 21.876 · RQE 16.038 (Nefrologia). Publicado
            originalmente em{' '}
            <a href={canonicalUrl} target="_blank" rel="noreferrer" className="link-text">
              plenyasaude.com.br/blog
            </a>
            .
          </p>
        </div>
      </section>

      {related.length > 0 && (
        <section className="border-t border-rule">
          <div className="editorial-container py-16">
            <p className="label-meta mb-8">Outros escritos no mesmo pilar</p>
            <ul className="grid md:grid-cols-3 gap-8">
              {related.map((p) => (
                <li key={p.slug}>
                  <Link href={`/escritos/${p.slug}`} className="block group space-y-3">
                    <span className="label-meta text-bordo">
                      {PLENYA_PILLAR_LABELS[p.pillar]}
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

function formatDate(d: string) {
  const dt = new Date(d);
  return dt.toLocaleDateString('pt-BR', { year: 'numeric', month: 'long' });
}
