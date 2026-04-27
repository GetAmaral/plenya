import type { Metadata } from 'next';
import Image from 'next/image';
import { setRequestLocale } from 'next-intl/server';

export const metadata: Metadata = {
  title: 'Antes — A Janela Silenciosa',
  description:
    'Antes — A Janela Silenciosa entre o Normal e o Ótimo. Livro do Dr. Getúlio Amaral Filho sobre os dez a vinte anos em que a longevidade é construída ou perdida.',
  alternates: { canonical: '/livro' },
};

const trechos = [
  {
    cap: 'Introdução',
    citacao:
      'O laboratório diz normal. O corpo diz: estou adoecendo há oito anos e ninguém está olhando.',
  },
  {
    cap: 'Capítulo 1 — O Homem que Morreu Saudável',
    citacao:
      'Medicina 2.0 contra Medicina 3.0. A diferença entre um alarme de incêndio e um detector de fumaça. O alarme toca quando já está queimando. O detector toca quando a fumaça começa.',
  },
  {
    cap: 'Capítulo 7 — Atividade Física',
    citacao:
      'Se existisse um fármaco que reduzisse o risco de morte por todas as causas em até 80%, a humanidade pagaria qualquer preço por ele. Esse fármaco existe. Chama-se exercício físico.',
  },
];

export default async function LivroPage({
  params,
}: {
  params: Promise<{ locale: string }>;
}) {
  const { locale } = await params;
  setRequestLocale(locale);

  return (
    <article>
      {/* Hero */}
      <section className="editorial-container pt-16 md:pt-24 pb-20">
        <div className="grid lg:grid-cols-[320px_1fr] gap-12 lg:gap-20 items-start">
          <div className="relative aspect-[2/3] w-full max-w-[320px] mx-auto lg:mx-0 shadow-2xl">
            <Image
              src="/images/livro-capa.jpg"
              alt="Capa do livro Antes — A Janela Silenciosa entre o Normal e o Ótimo"
              fill
              priority
              className="object-cover"
              sizes="(min-width: 1024px) 320px, 80vw"
            />
          </div>

          <div className="space-y-8">
            <p className="label-meta">Livro · 2026</p>
            <h1 className="heading-display text-[clamp(2.2rem,5vw,4rem)]">
              Antes
              <span className="block font-serif font-normal text-2xl md:text-3xl text-ink-muted mt-4 italic">
                A Janela Silenciosa entre o Normal e o Ótimo
              </span>
            </h1>

            <div className="prose-body max-w-xl">
              <p>
                A medicina brasileira se tornou excelente em tratar doença instalada — e
                ruim em prever o que está chegando. Entre a saúde que o laboratório
                reconhece e a doença que o médico trata existe uma janela silenciosa de dez
                a vinte anos, onde a longevidade é construída ou perdida.
              </p>
              <p>
                Baseado em vinte anos de prática clínica, este livro mostra o painel
                ampliado de biomarcadores que o check-up convencional ignora, o método AGIR
                de quatro pilares, e a Regra dos Dois — a única forma de mudança de hábito
                que produz resultado mensurável em três meses.
              </p>
            </div>

            <div className="space-y-1 font-sans text-sm text-ink-muted">
              <p>ISBN 978-65-02-06742-0 · 2026</p>
              <p>Edição em português · Brasil</p>
            </div>
          </div>
        </div>
      </section>

      {/* Trechos selecionados */}
      <section className="border-t border-rule bg-paper">
        <div className="editorial-container py-20 md:py-28">
          <p className="label-meta mb-14">Trechos</p>

          <div className="space-y-20 max-w-3xl">
            {trechos.map((t, i) => (
              <figure key={i} className="space-y-5">
                <p className="label-meta text-bordo">{t.cap}</p>
                <blockquote className="font-serif text-2xl md:text-3xl leading-snug text-ink italic border-l-2 border-bordo pl-6 md:pl-10">
                  {t.citacao}
                </blockquote>
              </figure>
            ))}
          </div>
        </div>
      </section>

      {/* Onde comprar — placeholder Sprint 2 com link real */}
      <section className="border-t border-rule">
        <div className="editorial-container py-20 md:py-24">
          <div className="grid md:grid-cols-[1fr_2fr] gap-12 items-start">
            <p className="label-meta">Onde comprar</p>
            <div className="space-y-4 font-serif text-ink-soft">
              <p>O livro <em>Antes</em> está disponível em formato físico e digital.</p>
              <p className="text-sm text-ink-muted">
                Links de compra serão atualizados em breve.
              </p>
            </div>
          </div>
        </div>
      </section>
    </article>
  );
}
