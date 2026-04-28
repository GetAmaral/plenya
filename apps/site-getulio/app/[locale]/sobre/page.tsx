import type { Metadata } from 'next';
import Image from 'next/image';
import { setRequestLocale } from 'next-intl/server';

export const metadata: Metadata = {
  title: 'Sobre',
  description:
    'Médico nefrologista formado pela UEL em 2004. Vinte anos de prática em medicina interna e nefrologia, com pós-graduação em medicina funcional integrativa.',
  alternates: { canonical: '/sobre' },
};

const formacao = [
  { year: '2004', label: 'Medicina — Universidade Estadual de Londrina' },
  { year: '2006', label: 'Especialização em Clínica Médica — Santa Casa de Londrina' },
  { year: '2008', label: 'Especialização em Nefrologia — Santa Casa de Londrina' },
  { year: '2008', label: 'Título de Especialista em Nefrologia — Sociedade Brasileira de Nefrologia' },
  { year: '2026', label: 'Pós-graduação em Medicina Funcional Integrativa — ABMFI' },
];

const atuacao = [
  'Coordenador da Residência Médica em Nefrologia — Santa Casa de Londrina',
  'Fundador da Residência Médica em Clínica Médica — Santa Casa de Londrina',
  'Sócio — Nefroclínica Londrina',
  'Responsável Técnico — DaVita Intra Hospitalar de Londrina',
  'Professor do curso de Medicina — PUC Londrina (2013–2014)',
  'Direção Clínica — Plenya',
];

const sociedades = [
  'Sociedade Brasileira de Nefrologia',
  'Sociedade Paranaense de Nefrologia',
  'Associação Brasileira de Medicina Funcional Integrativa',
];

export default async function SobrePage({
  params,
}: {
  params: Promise<{ locale: string }>;
}) {
  const { locale } = await params;
  setRequestLocale(locale);

  return (
    <article>
      {/* Cabeçalho */}
      <header className="editorial-container pt-16 md:pt-24 pb-12">
        <p className="label-meta mb-6">Sobre</p>
        <h1 className="heading-display text-[clamp(2.2rem,5vw,3.8rem)] max-w-3xl">
          Comecei na nefrologia, dentro do hospital.
        </h1>
      </header>

      {/* Corpo + aside */}
      <div className="editorial-container pb-24 grid lg:grid-cols-[1fr_360px] gap-16">
        <div className="prose-body">
          <Image
            src="/images/getulio-about-color.jpg"
            alt="Dr. Getúlio Amaral Filho"
            width={1200}
            height={800}
            className="w-full h-auto mb-12"
            priority
          />

          {/* Ato 1 — Origem */}
          <p>
            Sou de Londrina. Formei-me em medicina pela Universidade Estadual de Londrina em 2004
            e fiz residência em clínica médica e nefrologia na Santa Casa de Londrina, onde também
            obtive o título de especialista pela Sociedade Brasileira de Nefrologia em 2008.
          </p>
          <p>
            Naqueles anos, aprendi a medicina dentro do ambiente hospitalar — onde a doença
            renal e cardiovascular se mostra em estágios graves, com pouca margem para mudar
            o curso.
          </p>

          {/* Ato 2 — Vinte anos no hospital */}
          <p>
            Nas duas décadas seguintes, esse foi o meu dia a dia. Coordenei a residência médica
            em nefrologia da Santa Casa, fundei a residência em clínica médica da mesma
            instituição, e sou responsável técnico pela hemodiálise hospitalar da DaVita Intra
            Hospitalar de Londrina. Acompanhei centenas de pacientes em estágios avançados
            de doença — quadros em que, com frequência, a trajetória poderia ter sido
            modificada anos antes.
          </p>
          <p>
            Vi pacientes chegarem ao consultório com exames "normais" no ano anterior. Vi
            check-ups que registraram tudo dentro da faixa — e perderam o sinal de tudo o
            que estava se montando em silêncio. O laboratório dizia "normal". O corpo
            estava adoecendo há oito, dez, vinte anos.
          </p>
          <p>
            Foi essa observação repetida que me incomodou. Não era falha dos exames — era
            falha da pergunta. O check-up convencional procura doença instalada. Não procura
            o intervalo entre o normal e o ótimo, onde a longevidade é construída ou perdida.
          </p>

          {/* Ato 3 — A virada */}
          <p>
            Em 2026, conclui pós-graduação em medicina funcional integrativa pela Associação
            Brasileira de Medicina Funcional Integrativa. Não foi uma troca de paradigma. Foi
            uma extensão para trás — aplicar o mesmo rigor clínico que se exige no ambiente
            hospitalar, com exames e dados e decisão estruturada, ao período que precede
            a doença.
          </p>
          <p>
            É essa medicina que conduzo hoje na Plenya, ao lado da equipe que se reuniu para
            isso. E é sobre essa janela silenciosa entre o normal e o ótimo que escrevi
            o livro <em>Antes</em>, publicado em 2026.
          </p>

          {/* Ato 4 — O presente */}
          <p>
            Sigo atendendo nefrologia clínica na Nefroclínica Londrina, coordenando a residência
            na Santa Casa, respondendo tecnicamente pela hemodiálise hospitalar, e dirigindo
            clinicamente a Plenya. As quatro coisas são a mesma medicina, em momentos diferentes
            do tempo da pessoa.
          </p>

          {/* Linha humana — fechamento */}
          <p className="text-ink-muted italic">
            Vive em Londrina com a família. Corre quando o dia permite.
          </p>
        </div>

        {/* Aside — currículo. Eyebrows label-meta-lg + items text-base
           pra subir hierarquia e contraste (era ilegível em paper #f0e8d8). */}
        <aside className="space-y-12 lg:sticky lg:top-12 self-start">
          <div>
            <p className="label-meta-lg text-bordo mb-5">Formação</p>
            <ul className="space-y-3">
              {formacao.map((f) => (
                <li key={`${f.year}-${f.label}`} className="grid grid-cols-[64px_1fr] gap-3 border-b border-rule pb-3">
                  <span className="font-sans text-sm tracking-widest text-bordo font-medium pt-1">{f.year}</span>
                  <span className="font-serif text-base text-ink leading-snug">{f.label}</span>
                </li>
              ))}
            </ul>
          </div>

          <div>
            <p className="label-meta-lg text-bordo mb-5">Atuação</p>
            <ul className="space-y-3">
              {atuacao.map((a) => (
                <li key={a} className="font-serif text-base text-ink leading-relaxed">— {a}</li>
              ))}
            </ul>
          </div>

          <div>
            <p className="label-meta-lg text-bordo mb-5">Sociedades</p>
            <ul className="space-y-3">
              {sociedades.map((s) => (
                <li key={s} className="font-serif text-base text-ink leading-relaxed">— {s}</li>
              ))}
            </ul>
          </div>

          <div>
            <p className="label-meta-lg text-bordo mb-5">Credenciais</p>
            <p className="font-sans text-base text-ink">CRM-PR 21.876</p>
            <p className="font-sans text-base text-ink">RQE 16.038</p>
          </div>
        </aside>
      </div>
    </article>
  );
}
