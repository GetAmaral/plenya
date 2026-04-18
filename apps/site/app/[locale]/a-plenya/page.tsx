import type { Metadata } from 'next';
import Image from 'next/image';
import { setRequestLocale } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { agirPillars, brandEssence } from '@plenya/brand';
import { PlenyaInfinity } from '@plenya/brand/logo/PlenyaInfinity';

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
        <div className="relative site-narrow pt-32 pb-24 md:pt-40 md:pb-32">
          <p className="label-upper text-gold mb-10">Manifesto</p>
          <div className="space-y-6 heading-section text-2xl md:text-3xl text-cream/95">
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

      {/* Imagem da clínica */}
      <section className="bg-cream">
        <div className="site-container pt-16">
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
                <dd className="text-cream/80 text-lg leading-relaxed">{brandEssence.goldenCircle.why}</dd>
              </div>
              <div>
                <dt className="heading-section text-gold text-xl mb-3">Como</dt>
                <dd className="text-cream/80 text-lg leading-relaxed">{brandEssence.goldenCircle.how}</dd>
              </div>
              <div>
                <dt className="heading-section text-gold text-xl mb-3">O quê</dt>
                <dd className="text-cream/80 text-lg leading-relaxed">{brandEssence.goldenCircle.what}</dd>
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
          <Link href="/planos" className="border-t border-petrol/15 pt-8 space-y-3 group">
            <p className="label-upper text-gold">Planos</p>
            <p className="heading-section text-petrol text-xl group-hover:text-gold transition">Como começar →</p>
          </Link>
        </div>
      </section>
    </>
  );
}
