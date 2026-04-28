import type { Metadata } from 'next';
import Image from 'next/image';
import { setRequestLocale } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { agirPillars } from '@plenya/brand';
import { PlenyaInfinity } from '@plenya/brand/logo';

const goldenCircle = {
  why: 'Porque as pessoas estão vivendo mais, mas não necessariamente melhor — e seguem sendo tratadas em partes, não como um todo. Ninguém deveria chegar à fase mais importante da vida limitado pelas escolhas que fez antes.',
  how: 'Através de um acompanhamento integrado, contínuo e personalizado, que conecta diferentes áreas da saúde em um único plano estruturado, com dados, proximidade e responsabilidade compartilhada.',
  what: 'Um programa de saúde e longevidade que combina equipe multidisciplinar, método próprio (AGIR) e acompanhamento próximo para transformar hábitos, corpo e mente de forma sustentável.',
} as const;

export const metadata: Metadata = {
  title: 'A Plenya — Manifesto, Propósito e Método',
  description:
    'Conheça o manifesto, o propósito e o método AGIR da Plenya. Saúde funcional integrativa, contínua e personalizada.',
};

export default async function AboutPage({ params }: { params: Promise<{ locale: string }> }) {
  const { locale } = await params;
  setRequestLocale(locale);

  return (
    <>
      {/* Manifesto — com background fotográfico */}
      <section className="relative bg-petrol text-cream overflow-hidden">
        <Image
          src="/images/hero-about.jpg"
          alt=""
          fill
          priority
          className="object-cover opacity-50"
          sizes="100vw"
        />
        <div className="absolute inset-0 bg-gradient-to-r from-petrol via-petrol/80 to-petrol/40" />
        <div className="relative site-container pt-32 pb-24 md:pt-40 md:pb-32 max-w-3xl">
          <p className="label-upper text-gold mb-10">Manifesto</p>
          <div className="space-y-5 heading-section text-3xl md:text-5xl text-cream/95">
            <p>Tudo está conectado.</p>
            <p>Corpo, mente, tempo e escolhas.</p>
            <p>Viver bem é o reflexo do que você escolhe todos os dias.</p>
            <p>É consciência. É liberdade.</p>
            <p>Plenitude não é um ponto de chegada.</p>
            <p>É uma linha contínua.</p>
            <p className="text-gold pt-8">Plenya. Viva bem, viva mais.</p>
          </div>
        </div>
      </section>

      {/* Origem — paixão pela vida e pela longevidade */}
      <section className="bg-cream">
        <div className="site-container section">
          <div className="grid lg:grid-cols-[1fr_2fr] gap-12 lg:gap-20 items-start max-w-5xl">
            <div className="space-y-3">
              <p className="label-upper text-gold">Origem</p>
              {/* font-sans (Inter) — Nalieta tem cmap pra dígitos mas glyphs renderizam invisíveis */}
              <p className="font-sans text-petrol text-6xl md:text-7xl font-light leading-none tabular-nums">20+</p>
              <p className="label-upper text-petrol/60">anos de prática clínica</p>
            </div>
            <div className="space-y-6 text-petrol/80 text-lg leading-relaxed">
              <p>
                A Plenya nasce da paixão do <strong className="text-petrol">Dr. Getúlio Amaral Filho</strong>
                {' '}e da equipe que se reuniu em torno dele — médicos, nutricionista, psicólogo
                e educador físico — pela vida em sua forma plena e pela longevidade vivida com
                lucidez, energia e propósito.
              </p>
              <p>
                Em mais de duas décadas entre hospital, hemodiálise e consultório, o Dr. Getúlio
                viu a medicina salvar vidas em quadros graves — e viu, em paralelo, pessoas
                vivendo mais e vivendo pior, fragmentadas entre cinco, seis profissionais que
                não conversavam entre si. Pacientes com exames &ldquo;normais&rdquo; que já
                estavam adoecendo em silêncio.
              </p>
              <p>
                Plenya é a resposta clínica a essa observação. Um modelo desenhado para enxergar
                antes, acompanhar de perto e agir com clareza — com foco em
                <strong className="text-petrol"> saúde, performance e longevidade</strong>.
              </p>
              <p className="heading-section text-petrol text-2xl md:text-3xl pt-2">
                Porque viver mais já não é suficiente. O que importa é como você vive.
              </p>
            </div>
          </div>
        </div>
      </section>

      {/* Imagem da clínica */}
      <section className="bg-cream">
        <div className="site-container pt-4">
          <figure className="space-y-4">
            <div className="relative aspect-[3/2] overflow-hidden">
              <Image
                src="/images/clinic-interior.jpg"
                alt="Projeto arquitetônico do interior da futura sede Plenya"
                fill
                className="object-cover"
                sizes="(min-width: 1024px) 1120px, 100vw"
              />
            </div>
            <figcaption className="flex items-baseline gap-3 text-petrol/60 text-sm">
              <span className="label-upper text-gold">Projeto</span>
              <span>Render do interior da futura sede Plenya. Recepção e ambiente projetados para acolhimento e cuidado contínuo.</span>
            </figcaption>
          </figure>
        </div>
      </section>

      {/* Propósito */}
      <section className="bg-cream">
        <div className="site-container section">
          <div className="max-w-3xl space-y-6">
            <p className="label-upper text-gold">Propósito</p>
            <h2 className="heading-section text-petrol text-3xl md:text-5xl">
              Guiar pessoas a viverem mais e melhor, com autonomia sobre a própria saúde.
            </h2>
            <p className="text-petrol/80 text-lg leading-relaxed">
              Em um cenário marcado por excesso de informação e abordagens fragmentadas, a Plenya oferece um
              cuidado integrado, personalizado e sustentável ao longo do tempo. Conectamos diferentes áreas da
              saúde em um único plano estruturado, com dados, proximidade e responsabilidade compartilhada.
            </p>
          </div>
        </div>
      </section>

      {/* Círculo de Ouro — por quê / como / o quê (brandbook page 2) */}
      <section className="bg-petrol text-cream">
        <div className="site-container section">
          <div className="grid lg:grid-cols-[auto_1fr] gap-12 lg:gap-20 items-start">
            <div className="space-y-4">
              <PlenyaInfinity className="h-14 w-auto text-gold" />
              <p className="label-upper text-gold">Círculo de Ouro</p>
              <h2 className="heading-section text-cream text-3xl md:text-5xl max-w-md">
                Por que existimos. Como cuidamos. O que entregamos.
              </h2>
            </div>
            <dl className="space-y-10">
              <div>
                <dt className="heading-section text-gold text-xl mb-3">Por que</dt>
                <dd className="text-cream/80 text-lg leading-relaxed">{goldenCircle.why}</dd>
              </div>
              <div>
                <dt className="heading-section text-gold text-xl mb-3">Como</dt>
                <dd className="text-cream/80 text-lg leading-relaxed">{goldenCircle.how}</dd>
              </div>
              <div>
                <dt className="heading-section text-gold text-xl mb-3">O quê</dt>
                <dd className="text-cream/80 text-lg leading-relaxed">{goldenCircle.what}</dd>
              </div>
            </dl>
          </div>
        </div>
      </section>

      {/* Método AGIR — preview com link para página completa */}
      <section className="bg-paper">
        <div className="site-container section">
          <p className="label-upper text-gold mb-6">Método AGIR</p>
          <h2 className="heading-section text-petrol text-3xl md:text-5xl max-w-2xl mb-6">
            Quatro pilares interdependentes para uma saúde que se sustenta no tempo.
          </h2>
          <p className="text-petrol/70 text-lg max-w-2xl mb-16 leading-relaxed">
            O framework clínico que organiza o cuidado Plenya. Não é um protocolo — é uma forma
            de estruturar acompanhamento médico em torno do que sustenta saúde no longo prazo.
          </p>

          <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-px bg-petrol/15 border-y border-petrol/15">
            {agirPillars.map((pillar) => (
              <div key={pillar.code} className="bg-paper p-8 space-y-4">
                <span className="heading-section text-gold text-5xl block leading-none">{pillar.code}</span>
                <h3 className="heading-section text-petrol text-xl">{pillar.name}</h3>
                <p className="text-petrol/70 text-sm leading-relaxed">{pillar.idea}</p>
                <p className="label-upper text-petrol/50 text-[10px] tracking-[0.25em] pt-2 border-t border-petrol/10">
                  {pillar.territory}
                </p>
              </div>
            ))}
          </div>

          <div className="mt-12">
            <Link href="/metodo-agir" className="btn-gold">Conhecer o Método AGIR completo</Link>
          </div>
        </div>
      </section>

      {/* Posicionamento */}
      <section className="bg-cream">
        <div className="site-container section">
          <div className="max-w-3xl space-y-6">
            <p className="label-upper text-gold">Posicionamento</p>
            <h2 className="heading-section text-petrol text-3xl md:text-4xl">
              A fragmentação do cuidado é o problema. A integração é a resposta.
            </h2>
            <p className="text-petrol/80 text-lg leading-relaxed">
              Hoje, o paciente transita entre diferentes profissionais, informações e orientações desconectados,
              sem clareza, continuidade e direcionamento único. Na Plenya, integramos clínica, dados e
              acompanhamento próximo para transformar o cuidado em um processo contínuo, personalizado e mensurável.
            </p>
          </div>
        </div>
      </section>

      {/* Cross-links */}
      <section className="bg-paper">
        <div className="site-container section grid md:grid-cols-3 gap-8">
          <Link href="/dr-getulio" className="border-t border-petrol/15 pt-8 space-y-3 group">
            <p className="label-upper text-gold">Direção</p>
            <p className="heading-section text-petrol text-xl group-hover:text-gold transition">Conheça o Dr. Getúlio →</p>
          </Link>
          <Link href="/escore-plenya" className="border-t border-petrol/15 pt-8 space-y-3 group">
            <p className="label-upper text-gold">Diagnóstico</p>
            <p className="heading-section text-petrol text-xl group-hover:text-gold transition">Escore Plenya →</p>
          </Link>
          <Link href="/continuum" className="border-t border-petrol/15 pt-8 space-y-3 group">
            <p className="label-upper text-gold">Continuum</p>
            <p className="heading-section text-petrol text-xl group-hover:text-gold transition">Continuum Plenya →</p>
          </Link>
        </div>
      </section>
    </>
  );
}
