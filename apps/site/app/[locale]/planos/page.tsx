import type { Metadata } from 'next';
import { setRequestLocale } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';

export const metadata: Metadata = {
  title: 'Planos e Acompanhamento Plenya',
  description: 'Consultas particulares presenciais e online, e o acompanhamento Plenya com método AGIR.',
};

const offers = [
  { title: 'Consulta Presencial', desc: 'Avaliação clínica completa em nossa clínica em Londrina.', price: 'Sob consulta' },
  { title: 'Consulta Online', desc: 'Atendimento por telemedicina, mesma profundidade clínica.', price: 'Sob consulta' },
  { title: 'Consulta com Dr. Getúlio', desc: 'Atendimento direto com a direção clínica Plenya.', price: 'Sob consulta' },
];

const agir = [
  { name: 'Trimestral', period: '3 meses', price: 'Sob consulta' },
  { name: 'Semestral', period: '6 meses', price: 'Sob consulta', highlight: true },
  { name: 'Anual', period: '12 meses', price: 'Sob consulta' },
];

export default async function PlansPage({ params }: { params: Promise<{ locale: string }> }) {
  const { locale } = await params;
  setRequestLocale(locale);

  return (
    <>
      {/* Hero */}
      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-24 md:pt-40 md:pb-32">
          <p className="label-upper text-gold mb-6">Planos</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,5rem)] text-cream max-w-xl">
            Como você pode começar.
          </h1>
          <p className="text-cream/70 text-lg mt-6 max-w-lg">
            Atendimento particular, presencial e online. Consultas sob demanda ou acompanhamento estruturado com o Método AGIR.
          </p>
        </div>
      </section>

      {/* Consultas */}
      <section className="bg-cream">
        <div className="site-container section">
          <p className="label-upper text-gold mb-6">Consultas</p>
          <div className="grid md:grid-cols-3 gap-8">
            {offers.map((o) => (
              <div key={o.title} className="border-t border-petrol/15 pt-8 space-y-4">
                <h3 className="heading-section text-petrol text-2xl">{o.title}</h3>
                <p className="text-petrol/70 leading-relaxed">{o.desc}</p>
                <p className="label-upper text-petrol/50">{o.price}</p>
                <Link href="/contato" className="btn-outline-dark">Solicitar agendamento</Link>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* AGIR */}
      <section className="bg-petrol text-cream">
        <div className="site-container section">
          <p className="label-upper text-gold mb-6">Acompanhamento AGIR</p>
          <h2 className="heading-section text-cream text-3xl md:text-5xl max-w-2xl mb-6">
            Cuidado integrado e contínuo, com método.
          </h2>
          <p className="text-cream/70 text-lg max-w-xl mb-16">
            Programa estruturado de acompanhamento clínico com equipe multidisciplinar,
            monitoramento por dados e plano personalizado segundo o Método AGIR.
          </p>

          <div className="grid md:grid-cols-3 gap-8">
            {agir.map((tier) => (
              <div
                key={tier.name}
                className={
                  tier.highlight
                    ? 'border-t-2 border-gold pt-8 space-y-4'
                    : 'border-t border-cream/20 pt-8 space-y-4'
                }
              >
                <h3 className="heading-section text-cream text-2xl">AGIR {tier.name}</h3>
                <p className="text-cream/60">{tier.period} de acompanhamento</p>
                <p className="heading-section text-gold text-3xl">{tier.price}</p>
                <Link href="/contato" className="btn-gold">Conversar com a equipe</Link>
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
            <p className="heading-section text-petrol text-xl group-hover:text-gold transition">Conheça o Método AGIR →</p>
          </Link>
          <Link href="/escore-plenya" className="border-t border-petrol/15 pt-8 space-y-3 group">
            <p className="label-upper text-gold">Diagnóstico</p>
            <p className="heading-section text-petrol text-xl group-hover:text-gold transition">Escore Plenya →</p>
          </Link>
          <Link href="/depoimentos" className="border-t border-petrol/15 pt-8 space-y-3 group">
            <p className="label-upper text-gold">Pacientes</p>
            <p className="heading-section text-petrol text-xl group-hover:text-gold transition">Depoimentos →</p>
          </Link>
        </div>
      </section>
    </>
  );
}
