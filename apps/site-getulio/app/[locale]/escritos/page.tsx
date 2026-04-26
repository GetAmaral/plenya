import type { Metadata } from 'next';
import { setRequestLocale } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { getAllArticles, TAGS } from '@/lib/articles';

export const metadata: Metadata = {
  title: 'Escritos',
  description:
    'Artigos do Dr. Getúlio Amaral Filho sobre nefrologia preventiva, longevidade e medicina funcional integrativa.',
};

export default async function EscritosPage({
  params,
  searchParams,
}: {
  params: Promise<{ locale: string }>;
  searchParams: Promise<{ tag?: string }>;
}) {
  const { locale } = await params;
  const { tag: tagParam } = await searchParams;
  setRequestLocale(locale);
  const all = await getAllArticles();
  const activeTag = TAGS.find((t) => t === tagParam);
  const articles = activeTag ? all.filter((a) => a.tag === activeTag) : all;

  return (
    <article>
      <header className="editorial-container pt-16 md:pt-24 pb-12">
        <p className="label-meta mb-6">Escritos</p>
        <h1 className="heading-display text-[clamp(2.2rem,5vw,3.8rem)] max-w-3xl">
          Artigos sobre o que vem antes.
        </h1>
        <p className="prose-body mt-8 max-w-2xl">
          Notas clínicas, recortes do livro <em>Antes</em> e ensaios curtos sobre nefrologia
          preventiva, longevidade e medicina funcional integrativa.
        </p>
      </header>

      {/* Filtros */}
      <section className="editorial-container pb-8">
        <div className="flex flex-wrap items-center gap-x-6 gap-y-3 border-y border-rule py-4">
          <Link
            href="/escritos"
            className={`font-sans text-sm tracking-wide ${!activeTag ? 'text-bordo' : 'text-ink-muted hover:text-ink'} transition-colors`}
          >
            Todos
          </Link>
          {TAGS.map((t) => (
            <Link
              key={t}
              href={`/escritos?tag=${encodeURIComponent(t)}`}
              className={`font-sans text-sm tracking-wide ${activeTag === t ? 'text-bordo' : 'text-ink-muted hover:text-ink'} transition-colors`}
            >
              {t}
            </Link>
          ))}
        </div>
      </section>

      <section className="editorial-container pb-24">
        {articles.length === 0 ? (
          <p className="font-serif text-ink-muted py-8">Nenhum artigo nesse tema ainda.</p>
        ) : (
          <ul className="border-t border-rule">
            {articles.map((a) => (
              <li key={a.slug} className="border-b border-rule">
                <Link
                  href={`/escritos/${a.slug}`}
                  className="block py-8 grid md:grid-cols-[120px_1fr_180px] gap-6 group items-baseline"
                >
                  <span className="font-sans text-sm text-ink-muted">
                    {formatDate(a.date)}
                  </span>
                  <div className="space-y-2">
                    <h2 className="font-serif text-xl md:text-2xl text-ink group-hover:text-bordo transition-colors">
                      {a.title}
                    </h2>
                    <p className="font-serif text-ink-soft leading-relaxed">{a.excerpt}</p>
                  </div>
                  <span className="label-meta text-bordo md:text-right">{a.tag}</span>
                </Link>
              </li>
            ))}
          </ul>
        )}
      </section>
    </article>
  );
}

function formatDate(d: string) {
  const dt = new Date(d);
  return dt.toLocaleDateString('pt-BR', { year: 'numeric', month: 'short' });
}
