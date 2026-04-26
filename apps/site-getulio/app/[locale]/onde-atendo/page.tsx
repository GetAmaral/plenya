import type { Metadata } from 'next';
import Image from 'next/image';
import { setRequestLocale } from 'next-intl/server';

export const metadata: Metadata = {
  title: 'Onde atendo',
  description:
    'Atendimento clínico na Plenya (medicina funcional integrativa) e na Nefroclínica Londrina (nefrologia clínica). Responsabilidade técnica pela DaVita Intra Hospitalar e DaVita Londrina.',
};

type Clinic = {
  name: string;
  logo: { src: string; height: number } | null;
  role: string;
  body: string;
  address: string;
  href: string | null;
};

const clinicas: Clinic[] = [
  {
    name: 'Plenya',
    logo: { src: '/logos/plenya.svg', height: 36 },
    role: 'Direção clínica',
    body:
      'Programa de saúde preventiva e longevidade. Medicina funcional integrativa, com equipe multidisciplinar (médico, nutricionista, psicóloga, educador físico) trabalhando o mesmo paciente sob o mesmo plano.',
    address: 'Londrina · Paraná',
    href: 'https://plenyasaude.com.br',
  },
  {
    name: 'Nefroclínica Londrina',
    logo: { src: '/logos/nefroclinica.png', height: 70 },
    role: 'Sócio',
    body:
      'Nefrologia clínica em Londrina há quatro décadas. Atendimento de doença renal crônica, hipertensão de difícil controle, distúrbios eletrolíticos e acompanhamento pré-diálise.',
    address: 'Londrina · Paraná',
    href: 'https://nefroclinica.com',
  },
  {
    name: 'DaVita Intra Hospitalar',
    logo: null,
    role: 'Responsável técnico',
    body:
      'Hemodiálise hospitalar — Santa Casa de Londrina. Atendimento a pacientes internados em estágio avançado de doença renal, em conjunto com a equipe de nefrologia da instituição.',
    address: 'Santa Casa de Londrina · Paraná',
    href: null,
  },
  {
    name: 'DaVita Londrina',
    logo: null,
    role: 'Responsável técnico',
    body:
      'Unidade ambulatorial de hemodiálise em Londrina. Acompanhamento crônico de pacientes em terapia renal substitutiva.',
    address: 'Londrina · Paraná',
    href: null,
  },
];

export default async function OndeAtendoPage({
  params,
}: {
  params: Promise<{ locale: string }>;
}) {
  const { locale } = await params;
  setRequestLocale(locale);

  return (
    <article>
      <header className="editorial-container pt-16 md:pt-24 pb-12">
        <p className="label-meta mb-6">Onde atendo</p>
        <h1 className="heading-display text-[clamp(2.2rem,5vw,3.8rem)] max-w-3xl">
          Quatro frentes, a mesma medicina.
        </h1>
        <p className="prose-body mt-8 max-w-2xl">
          Atendo em duas clínicas em Londrina e respondo tecnicamente por duas unidades de
          hemodiálise — uma hospitalar, uma ambulatorial. As quatro frentes são a mesma
          medicina em momentos diferentes do tempo da pessoa: antecipar, acompanhar,
          tratar.
        </p>
      </header>

      <section className="editorial-container pb-24 space-y-16">
        {clinicas.map((c) => (
          <div
            key={c.name}
            className="border-t border-rule pt-12 grid md:grid-cols-[260px_1fr] gap-10 md:gap-16"
          >
            <div className="space-y-4">
              <div className="h-20 flex items-center">
                {c.logo ? (
                  <Image
                    src={c.logo.src}
                    alt={c.name}
                    width={260}
                    height={c.logo.height}
                    style={{ height: c.logo.height, width: 'auto' }}
                    className="object-contain object-left"
                  />
                ) : (
                  <span className="font-serif text-ink text-3xl tracking-tight">{c.name}</span>
                )}
              </div>
              <p className="label-meta text-bordo">{c.role}</p>
            </div>
            <div className="space-y-5">
              {c.logo && <h2 className="heading-section text-2xl md:text-3xl">{c.name}</h2>}
              <p className="font-serif text-ink-soft leading-relaxed max-w-2xl">{c.body}</p>
              <p className="font-sans text-sm text-ink-muted">{c.address}</p>
              {c.href && (
                <a
                  href={c.href}
                  target="_blank"
                  rel="noreferrer"
                  className="link-text inline-block font-sans text-sm"
                >
                  Visitar site ↗
                </a>
              )}
            </div>
          </div>
        ))}
      </section>

      <section className="border-t border-rule bg-paper">
        <div className="editorial-container py-20 grid md:grid-cols-[1fr_2fr] gap-12 items-start">
          <p className="label-meta">Agendar</p>
          <div className="space-y-4 font-serif text-ink-soft">
            <p>
              Para agendamento, fale diretamente com a clínica de referência. Para outras
              demandas (palestra, imprensa), use os canais correspondentes.
            </p>
            <p className="font-sans text-sm">
              <a href="mailto:contato@drgetulioamaralfilho.com.br" className="link-text">
                contato@drgetulioamaralfilho.com.br
              </a>
            </p>
          </div>
        </div>
      </section>
    </article>
  );
}
