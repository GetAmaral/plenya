import type { Metadata } from 'next';
import Image from 'next/image';
import { setRequestLocale } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import {
  absoluteCover,
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
  const [featured, ...rest] = posts;

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
                    <span className="text-bordo">{PLENYA_PILLAR_LABELS[featured.pillar]}</span>
                    <span className="text-ink-muted">·</span>
                    <time className="text-ink-muted" dateTime={featured.date}>
                      {formatDate(featured.date)}
                    </time>
                  </div>
                  <h2 className="font-serif text-3xl md:text-4xl leading-tight text-ink group-hover:text-bordo transition-colors">
                    {featured.title}
                  </h2>
                  <p className="font-serif text-lg text-ink-soft leading-relaxed">
                    {featured.excerpt}
                  </p>
                  <p className="label-meta text-bordo pt-2">Ler artigo →</p>
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
                      <span className="text-bordo">{PLENYA_PILLAR_LABELS[p.pillar]}</span>
                      <span className="text-ink-muted">·</span>
                      <time className="text-ink-muted" dateTime={p.date}>
                        {formatDate(p.date)}
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
