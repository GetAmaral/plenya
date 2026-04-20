import type { Metadata } from 'next';
import { setRequestLocale } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';

export const metadata: Metadata = {
  title: 'Plano de Acompanhamento Plenya',
  description:
    'Programa estruturado de acompanhamento clínico contínuo com equipe multidisciplinar e o Método AGIR.',
};

const agir = [
  { name: 'Trimestral', period: '3 meses', price: 'Sob consulta' },
  { name: 'Semestral', period: '6 meses', price: 'Sob consulta', highlight: true },
  { name: 'Anual', period: '12 meses', price: 'Sob consulta' },
];

export default async function PlanosPage({ params }: { params: Promise<{ locale: string }> }) {
  const { locale } = await params;
  setRequestLocale(locale);

  return (
    <>
      {/* Hero */}
      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-24 md:pt-40 md:pb-32">
          <p className="label-upper text-gold mb-6">Acompanhamento</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,5rem)] text-cream max-w-2xl">
            Plano de Acompanhamento Plenya.
          </h1>
          <p className="text-cream/70 text-lg mt-6 max-w-xl">
            Cuidado integrado e contínuo, com equipe multidisciplinar e método. Para quem quer ir além
            de uma consulta avulsa e construir saúde no tempo.
          </p>
        </div>
      </section>

      {/* AGIR tiers */}
      <section className="bg-petrol text-cream border-t border-cream/10">
        <div className="site-container section">
          <p className="label-upper text-gold mb-6">Modalidades</p>
          <h2 className="heading-section text-cream text-3xl md:text-5xl max-w-2xl mb-16">
            Três horizontes de compromisso.
          </h2>

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
          <Link href="/consultas" className="border-t border-petrol/15 pt-8 space-y-3 group">
            <p className="label-upper text-gold">Avulso</p>
            <p className="heading-section text-petrol text-xl group-hover:text-gold transition">
              Consultas médicas →
            </p>
          </Link>
          <Link href="/metodo-agir" className="border-t border-petrol/15 pt-8 space-y-3 group">
            <p className="label-upper text-gold">Método</p>
            <p className="heading-section text-petrol text-xl group-hover:text-gold transition">
              Conheça o Método AGIR →
            </p>
          </Link>
          <Link href="/escore-plenya" className="border-t border-petrol/15 pt-8 space-y-3 group">
            <p className="label-upper text-gold">Diagnóstico</p>
            <p className="heading-section text-petrol text-xl group-hover:text-gold transition">
              Escore Plenya →
            </p>
          </Link>
        </div>
      </section>
    </>
  );
}
