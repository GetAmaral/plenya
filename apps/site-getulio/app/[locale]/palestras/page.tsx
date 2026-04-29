import type { Metadata } from 'next';
import Image from 'next/image';
import { setRequestLocale } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { getAllLectures, getAudienceLabel } from '@/lib/lectures';

export const metadata: Metadata = {
  title: 'Palestras',
  description:
    'Palestras de Dr. Getúlio Amaral Filho sobre prevenção que começa antes do diagnóstico, longevidade prática e nefrologia preventiva.',
  alternates: { canonical: '/palestras' },
};

export default async function PalestrasPage({
  params,
}: {
  params: Promise<{ locale: string }>;
}) {
  const { locale } = await params;
  setRequestLocale(locale);
  const lectures = await getAllLectures();

  return (
    <article>
      <header className="editorial-container pt-16 md:pt-24 pb-12">
        <p className="label-meta mb-6">Palestras</p>
        <h1 className="heading-display text-[clamp(2.2rem,5vw,3.8rem)] max-w-3xl">
          Sobre prevenção que começa antes.
        </h1>
        <p className="prose-body mt-8 max-w-2xl">
          Palestras para médicos, residentes, equipes corporativas e plateias abertas — sobre
          a janela silenciosa entre o normal e o ótimo, e o método clínico para encurtá-la.
        </p>
      </header>

      <section className="editorial-container pb-16 md:pb-20">
        <div className="relative aspect-[2/3] sm:aspect-[3/4] md:aspect-[4/5] lg:aspect-[3/2] w-full overflow-hidden bg-paper">
          <Image
            src="/images/getulio-palestrante-arvore.jpg"
            alt="Dr. Getúlio Amaral Filho durante palestra — slide com diagrama da árvore das doenças crônicas (cardiovasculares, câncer, neurodegenerativas, metabólicas)"
            fill
            priority
            className="object-cover object-[center_top] lg:object-[center_25%]"
            sizes="(min-width: 1024px) 1100px, 100vw"
          />
        </div>
      </section>

      <section className="editorial-container pb-24">
        <ul className="border-t border-rule">
          {lectures.map((l) => (
            <li key={l.slug} className="border-b border-rule">
              <Link
                href={`/palestras/${l.slug}`}
                className="block py-10 grid md:grid-cols-[1fr_280px] gap-8 group"
              >
                <div className="space-y-4">
                  <div className="flex items-center gap-3 flex-wrap">
                    {l.anchor && <span className="label-meta text-bordo">Palestra-âncora</span>}
                    <span className="label-meta text-ink-muted">{l.duration}</span>
                  </div>
                  <h2 className="heading-section text-2xl md:text-3xl group-hover:text-bordo transition-colors">
                    {l.title}
                  </h2>
                  {l.subtitle && (
                    <p className="font-serif italic text-ink-muted text-lg">{l.subtitle}</p>
                  )}
                  <p className="font-serif text-ink-soft leading-relaxed max-w-2xl">{l.excerpt}</p>
                </div>
                <div className="space-y-3 md:text-right">
                  <p className="label-meta">Público</p>
                  <ul className="space-y-1">
                    {l.audience.map((a) => (
                      <li key={a} className="font-serif text-sm text-ink-soft">
                        {getAudienceLabel(a)}
                      </li>
                    ))}
                  </ul>
                  <p className="label-meta pt-3">Formato</p>
                  <p className="font-serif text-sm text-ink-soft">{l.format}</p>
                </div>
              </Link>
            </li>
          ))}
        </ul>
      </section>

      <section className="border-t border-rule bg-paper">
        <div className="editorial-container py-20 grid md:grid-cols-[1fr_2fr] gap-12 items-start">
          <p className="label-meta">Convidar para palestrar</p>
          <div className="space-y-4 font-serif text-ink-soft">
            <p>
              Envie proposta com nome do evento, data, local, público estimado, formato
              (presencial ou remoto) e duração desejada.
            </p>
            <p className="font-sans text-sm">
              <a href="mailto:palestras@drgetulioamaralfilho.com.br" className="link-text">
                palestras@drgetulioamaralfilho.com.br
              </a>
              <br />
              <span className="text-ink-muted text-xs">Resposta em até 48h úteis.</span>
            </p>
          </div>
        </div>
      </section>
    </article>
  );
}
