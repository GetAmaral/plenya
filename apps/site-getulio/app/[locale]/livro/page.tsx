import type { Metadata } from 'next';
import Image from 'next/image';
import { setRequestLocale } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';

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

      {/* Autor */}
      <section className="border-t border-rule">
        <div className="editorial-container py-20 md:py-24">
          <div className="grid md:grid-cols-[1fr_320px] gap-12 lg:gap-20 items-center">
            <div className="space-y-5 max-w-xl order-2 md:order-1">
              <p className="label-meta-lg text-bordo">Sobre o autor</p>
              <p className="font-serif text-lg md:text-xl text-ink-soft leading-relaxed">
                Dr. Getúlio Amaral Filho é médico nefrologista (CRM-PR 21.876 · RQE 16.038).
                Há vinte anos atende em hospital, ambulatório e diálise. Dirige a Plenya —
                programa de longevidade que inverte a sequência do cuidado: começa antes da
                doença instalada, na janela silenciosa onde o exame ainda diz "normal".
              </p>
              <p className="font-sans text-sm">
                <Link href="/sobre" className="link-text">Ler biografia completa →</Link>
              </p>
            </div>
            <div className="relative aspect-[3/4] w-full max-w-[320px] mx-auto md:mx-0 order-1 md:order-2">
              <Image
                src="/images/getulio-autor.jpg"
                alt="Dr. Getúlio Amaral Filho"
                fill
                className="object-cover"
                sizes="(min-width: 768px) 320px, 70vw"
              />
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

      {/* "Onde comprar" — escondido até ter os links reais (Amazon/Saraiva).
         Placeholder anterior ("links em breve") parecia obra inacabada.
         Substituído por bloco "Ficar sabendo" com captura editorial:
         lead time vira lead capture. */}
      <section className="border-t border-rule bg-paper">
        <div className="editorial-container py-20 md:py-24">
          <div className="grid md:grid-cols-[1fr_2fr] gap-12 items-start">
            <p className="label-meta-lg text-bordo">Lançamento</p>
            <div className="space-y-5 max-w-xl">
              <h2 className="heading-section text-ink text-2xl md:text-3xl leading-snug">
                Quer ser avisado quando o livro estiver disponível?
              </h2>
              <p className="font-serif text-ink-soft leading-relaxed">
                <em>Antes — A Janela Silenciosa entre o Normal e o Ótimo</em> entra em
                pré-venda em 2026. Para ser notificado da data, escreva para o endereço
                abaixo com o assunto "Antes — pré-venda".
              </p>
              <p className="font-sans text-base">
                <a href="mailto:contato@drgetulioamaralfilho.com.br?subject=Antes%20%E2%80%94%20pr%C3%A9-venda" className="link-text">
                  contato@drgetulioamaralfilho.com.br
                </a>
              </p>
            </div>
          </div>
        </div>
      </section>
    </article>
  );
}
