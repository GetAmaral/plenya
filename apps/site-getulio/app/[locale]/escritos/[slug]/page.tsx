import type { Metadata } from 'next';
import { notFound } from 'next/navigation';
import { setRequestLocale } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { Breadcrumbs } from '@/components/layout/breadcrumbs';
import { getAllArticles, getArticle } from '@/lib/articles';

export async function generateStaticParams() {
  const all = await getAllArticles();
  return all.map((a) => ({ slug: a.slug }));
}

export async function generateMetadata({
  params,
}: {
  params: Promise<{ slug: string }>;
}): Promise<Metadata> {
  const { slug } = await params;
  const a = await getArticle(slug);
  if (!a) return {};
  return { title: a.title, description: a.excerpt };
}

export default async function ArticlePage({
  params,
}: {
  params: Promise<{ locale: string; slug: string }>;
}) {
  const { locale, slug } = await params;
  setRequestLocale(locale);
  const article = await getArticle(slug);
  if (!article) notFound();

  const all = await getAllArticles();
  const related = all.filter((a) => a.slug !== article.slug).slice(0, 3);

  return (
    <article>
      <header className="editorial-container pt-12 md:pt-16 pb-8">
        <Breadcrumbs
          items={[
            { label: 'Início', href: '/' },
            { label: 'Escritos', href: '/escritos' },
            { label: article.title },
          ]}
        />
      </header>

      <section className="editorial-narrow pb-12">
        <div className="space-y-8">
          <div className="flex items-center gap-4 flex-wrap">
            <span className="label-meta text-bordo">{article.tag}</span>
            <span className="label-meta text-ink-muted">
              {formatDate(article.date)} · {article.readingMinutes} min
            </span>
          </div>
          <h1 className="heading-display text-[clamp(2rem,4.5vw,3.4rem)]">{article.title}</h1>
          <p className="font-serif italic text-ink-muted text-xl leading-relaxed">
            {article.excerpt}
          </p>
        </div>
      </section>

      <section className="editorial-narrow pb-20">
        <div className="prose-body whitespace-pre-line">{article.content.trim()}</div>
        {article.source && (
          <p className="mt-12 font-sans text-xs text-ink-muted/80 border-t border-rule pt-6">
            Fonte: {article.source}.
          </p>
        )}
      </section>

      {related.length > 0 && (
        <section className="border-t border-rule">
          <div className="editorial-container py-16">
            <p className="label-meta mb-8">Outros escritos</p>
            <ul className="grid md:grid-cols-3 gap-8">
              {related.map((a) => (
                <li key={a.slug}>
                  <Link
                    href={`/escritos/${a.slug}`}
                    className="block group space-y-3"
                  >
                    <span className="label-meta text-bordo">{a.tag}</span>
                    <h3 className="font-serif text-lg text-ink group-hover:text-bordo transition-colors">
                      {a.title}
                    </h3>
                    <p className="font-serif text-sm text-ink-muted leading-relaxed">{a.excerpt}</p>
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
