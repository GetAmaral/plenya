import type { Metadata } from 'next';
import { setRequestLocale } from 'next-intl/server';

export const metadata: Metadata = {
  title: 'Ensino',
  description:
    'Coordenador da Residência Médica em Nefrologia da Santa Casa de Londrina. Fundador da Residência em Clínica Médica. Professor PUC Londrina.',
  alternates: { canonical: '/ensino' },
};

const cargos = [
  {
    period: 'Desde 2015',
    title: 'Coordenador da Residência Médica em Nefrologia',
    org: 'Santa Casa de Londrina',
    body:
      'Responsável pela formação de residentes em nefrologia, coordenando o programa, a grade clínica e a integração com a hemodiálise hospitalar.',
  },
  {
    period: 'Fundador',
    title: 'Residência Médica em Clínica Médica',
    org: 'Santa Casa de Londrina',
    body:
      'Desenhei e implantei o programa de residência em clínica médica da instituição — porque sem clínica geral sólida não há especialização que valha.',
  },
  {
    period: '2013 – 2014',
    title: 'Professor do curso de Medicina',
    org: 'PUC Londrina',
    body:
      'Disciplinas de clínica médica e nefrologia para a graduação. Formação de futuros médicos com ênfase em raciocínio clínico e leitura integrada de exames.',
  },
];

export default async function EnsinoPage({
  params,
}: {
  params: Promise<{ locale: string }>;
}) {
  const { locale } = await params;
  setRequestLocale(locale);

  return (
    <article>
      <header className="editorial-container pt-16 md:pt-24 pb-12">
        <p className="label-meta mb-6">Ensino</p>
        <h1 className="heading-display text-[clamp(2.2rem,5vw,3.8rem)] max-w-3xl">
          Ensinar é a outra metade da clínica.
        </h1>
        <p className="prose-body mt-8 max-w-2xl">
          Coordeno residência médica há mais de uma década. O que se aprende preparando o
          próximo médico é diferente do que se aprende atendendo o próximo paciente — e os
          dois se sustentam.
        </p>
      </header>

      <section className="editorial-container pb-20">
        <ul className="border-t border-rule">
          {cargos.map((c) => (
            <li key={c.title} className="border-b border-rule py-10 grid md:grid-cols-[180px_1fr] gap-6 md:gap-12">
              <p className="label-meta text-bordo">{c.period}</p>
              <div className="space-y-3">
                <h2 className="heading-section text-xl md:text-2xl">{c.title}</h2>
                <p className="label-meta text-ink-muted">{c.org}</p>
                <p className="font-serif text-ink-soft leading-relaxed max-w-2xl">{c.body}</p>
              </div>
            </li>
          ))}
        </ul>
      </section>

      <section className="border-t border-rule bg-paper">
        <div className="editorial-container py-20 max-w-3xl">
          <p className="label-meta mb-6">Filosofia de ensino</p>
          <div className="prose-body space-y-6">
            <p>
              O bom residente aprende a ler o paciente antes do exame. O bom programa de
              residência ensina a duvidar do "normal" — a ver o intervalo entre o que está
              dentro da faixa e o que sustenta a longevidade do indivíduo à frente.
            </p>
            <p>
              O que tento passar é simples: rigor clínico hospitalar é compatível — e
              necessário — com medicina preventiva de precisão. Não são dois mundos. É a
              mesma medicina, em momentos diferentes do tempo da pessoa.
            </p>
          </div>
        </div>
      </section>
    </article>
  );
}
