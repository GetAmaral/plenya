import type { Metadata } from 'next';
import { setRequestLocale } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import {
  getPlenyaPostsByGetulio,
  PLENYA_BLOG_BASE,
  PLENYA_PILLAR_LABELS,
} from '@/lib/plenya-blog';

export const metadata: Metadata = {
  title: 'Escritos',
  description:
    'Artigos do Dr. Getúlio Amaral Filho sobre nefrologia preventiva, longevidade e medicina funcional integrativa. Espelho fiel do blog Plenya.',
  alternates: { canonical: '/escritos' },
};

const PILLAR_FILTERS = Object.entries(PLENYA_PILLAR_LABELS) as [
  keyof typeof PLENYA_PILLAR_LABELS,
  string,
][];

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
  const all = await getPlenyaPostsByGetulio();
  const activePilar = PILLAR_FILTERS.find(([k]) => k === pilarParam)?.[0];
  const posts = activePilar ? all.filter((p) => p.pillar === activePilar) : all;

  return (
    <article>
      <header className="editorial-container pt-16 md:pt-24 pb-12">
        <p className="label-meta mb-6">Escritos</p>
        <h1 className="heading-display text-[clamp(2.2rem,5vw,3.8rem)] max-w-3xl">
          Artigos sobre o que vem antes.
        </h1>
        <p className="prose-body mt-8 max-w-2xl">
          Notas clínicas, recortes do livro <em>Antes</em> e ensaios curtos sobre nefrologia
          preventiva, longevidade e medicina funcional integrativa. Conteúdo publicado
          originalmente no blog da Plenya, espelhado aqui na íntegra.
        </p>
      </header>

      <section className="editorial-container pb-8">
        <div className="flex flex-wrap items-center gap-x-6 gap-y-3 border-y border-rule py-4">
          <Link
            href="/escritos"
            className={`font-sans text-sm tracking-wide ${!activePilar ? 'text-bordo' : 'text-ink-muted hover:text-ink'} transition-colors`}
          >
            Todos
          </Link>
          {PILLAR_FILTERS.map(([key, label]) => (
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
          <p className="font-serif text-ink-muted py-8">Nenhum artigo nesse pilar ainda.</p>
        ) : (
          <ul className="border-t border-rule">
            {posts.map((p) => (
              <li key={p.slug} className="border-b border-rule">
                <Link
                  href={`/escritos/${p.slug}`}
                  className="block py-8 grid md:grid-cols-[120px_1fr_180px] gap-6 group items-baseline"
                >
                  <span className="font-sans text-sm text-ink-muted">
                    {formatDate(p.date)}
                  </span>
                  <div className="space-y-2">
                    <h2 className="font-serif text-xl md:text-2xl text-ink group-hover:text-bordo transition-colors">
                      {p.title}
                    </h2>
                    <p className="font-serif text-ink-soft leading-relaxed">{p.excerpt}</p>
                  </div>
                  <span className="label-meta text-bordo md:text-right">
                    {PLENYA_PILLAR_LABELS[p.pillar]}
                  </span>
                </Link>
              </li>
            ))}
          </ul>
        )}

        <p className="font-sans text-xs text-ink-muted mt-12 text-center">
          Conteúdo publicado originalmente no{' '}
          <a href={PLENYA_BLOG_BASE} target="_blank" rel="noreferrer" className="link-text">
            Blog Plenya
          </a>
          . Cada artigo aqui é uma versão fiel do original, com canonical apontando para a fonte.
        </p>
      </section>
    </article>
  );
}

function formatDate(d: string) {
  const dt = new Date(d);
  return dt.toLocaleDateString('pt-BR', { year: 'numeric', month: 'short' });
}
