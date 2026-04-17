import type { Metadata } from 'next';
import Image from 'next/image';
import { setRequestLocale } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';

export const metadata: Metadata = {
  title: 'Dr. Getúlio Amaral Filho — Direção Clínica',
  description:
    'Médico nefrologista, clínico e especialista em medicina funcional integrativa. Direção clínica da Plenya.',
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
  'Responsável Técnico — DaVita Intra Hospitalar de Londrina',
  'Professor do curso de Medicina — PUC Londrina (2013-2014)',
  'Direção Clínica — Plenya',
];

const memberships = [
  'Sociedade Brasileira de Nefrologia',
  'Sociedade Paranaense de Nefrologia',
  'Associação Brasileira de Medicina Funcional Integrativa',
];

export default async function DrGetulioPage({ params }: { params: Promise<{ locale: string }> }) {
  const { locale } = await params;
  setRequestLocale(locale);

  return (
    <>
      {/* Hero */}
      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-24 md:pt-40 md:pb-32 grid lg:grid-cols-2 gap-16 items-center">
          <div className="space-y-6">
            <p className="label-upper text-gold">Direção Clínica Plenya</p>
            <h1 className="heading-hero text-[clamp(2.5rem,6vw,4.5rem)] text-cream">
              Dr. Getúlio José Mattos<br />do Amaral Filho
            </h1>
            <p className="label-upper text-cream/60">
              CRM-PR 21.876 · RQE 16.038
            </p>
            <p className="text-cream/85 text-lg leading-relaxed max-w-lg">
              Médico desde 2004. Especialista em nefrologia e clínica médica pela Santa Casa
              de Londrina. Pós-graduado em medicina funcional integrativa pela ABMFI.
            </p>
            <div className="flex flex-wrap gap-5 pt-4">
              <Link href="/contato" className="btn-gold">Agendar consulta</Link>
              <a
                href="https://instagram.com/drGetulioAmaralFilho"
                target="_blank"
                rel="noreferrer"
                className="btn-outline-dark border-cream/40 text-cream"
              >
                @drGetulioAmaralFilho
              </a>
            </div>
          </div>
          <div className="relative aspect-[3/4] overflow-hidden">
            <Image
              src="/images/dr-getulio.jpg"
              alt="Dr. Getúlio Amaral Filho — Direção Clínica Plenya"
              fill
              priority
              className="object-cover"
              sizes="(min-width: 1024px) 540px, 100vw"
            />
          </div>
        </div>
      </section>

      {/* Sobre a prática */}
      <section className="bg-cream">
        <div className="site-container section grid lg:grid-cols-[2fr_1fr] gap-16">
          <div className="space-y-6 text-petrol/85 text-lg leading-relaxed">
            <p className="label-upper text-gold">Sobre a prática</p>
            <p>
              Trabalho com medicina interna e nefrologia há 22 anos, parte significativa em
              ambiente hospitalar. Atendi muitos pacientes em estágios avançados de doença renal
              e cardiovascular — quadros em que boa parte do que se vê poderia ter sido modificado
              anos antes, em decisões de rotina e em monitoramento mais cuidadoso.
            </p>
            <p>
              A medicina funcional integrativa não é uma alternativa à medicina convencional. É uma
              forma de organizar o cuidado preventivo com o mesmo rigor que se aplica ao cuidado
              hospitalar — exames, dados, decisão estruturada — mas voltada para o que vem antes
              da doença.
            </p>
            <p>
              É essa medicina que ofereço na Plenya, junto com a equipe que reunimos para isso.
            </p>
          </div>

          <aside className="space-y-8">
            <div className="space-y-4">
              <p className="label-upper text-gold">Atuação</p>
              <ul className="space-y-3 text-petrol/80 text-sm">
                {atuacao.map((item) => (
                  <li key={item} className="flex gap-3">
                    <span className="text-gold mt-0.5">—</span>
                    <span>{item}</span>
                  </li>
                ))}
              </ul>
            </div>

            <div className="space-y-4">
              <p className="label-upper text-gold">Sociedades</p>
              <ul className="space-y-3 text-petrol/80 text-sm">
                {memberships.map((m) => (
                  <li key={m} className="flex gap-3">
                    <span className="text-gold mt-0.5">—</span>
                    <span>{m}</span>
                  </li>
                ))}
              </ul>
            </div>
          </aside>
        </div>
      </section>

      {/* Formação acadêmica */}
      <section className="bg-paper">
        <div className="site-container section">
          <p className="label-upper text-gold mb-10">Formação acadêmica</p>

          <div className="border-t border-petrol/15">
            {formacao.map((item) => (
              <div
                key={`${item.year}-${item.label}`}
                className="grid grid-cols-[80px_1fr] md:grid-cols-[120px_1fr] gap-6 py-6 border-b border-petrol/10"
              >
                <span className="label-upper text-gold pt-1">{item.year}</span>
                <p className="text-petrol/80">{item.label}</p>
              </div>
            ))}
          </div>

        </div>
      </section>

      {/* Cross-links */}
      <section className="bg-cream">
        <div className="site-container section grid md:grid-cols-3 gap-8">
          <Link href="/metodo-agir" className="border-t border-petrol/15 pt-8 space-y-3 group">
            <p className="label-upper text-gold">Método</p>
            <p className="heading-section text-petrol text-xl group-hover:text-gold transition">Método AGIR →</p>
          </Link>
          <Link href="/equipe" className="border-t border-petrol/15 pt-8 space-y-3 group">
            <p className="label-upper text-gold">Equipe</p>
            <p className="heading-section text-petrol text-xl group-hover:text-gold transition">Equipe Plenya →</p>
          </Link>
          <Link href="/planos" className="border-t border-petrol/15 pt-8 space-y-3 group">
            <p className="label-upper text-gold">Planos</p>
            <p className="heading-section text-petrol text-xl group-hover:text-gold transition">Planos e acompanhamento →</p>
          </Link>
        </div>
      </section>
    </>
  );
}
