import { useTranslations } from 'next-intl';
import { Link } from '@/lib/i18n/navigation';

const plans = [
  { title: 'Consulta', desc: 'Presencial ou online, com a equipe médica Plenya.' },
  { title: 'Consulta Dr. Getúlio', desc: 'Atendimento direto com a direção clínica.' },
  { title: 'Acompanhamento AGIR', desc: 'Trimestral, semestral ou anual. Cuidado integrado contínuo.' },
];

export function PlansPreview() {
  const tCta = useTranslations('cta');
  return (
    <section className="bg-paper">
      <div className="site-container section">
        <p className="label-upper text-gold mb-6">Planos</p>
        <h2 className="heading-section text-petrol text-3xl md:text-5xl max-w-xl mb-16">
          Como você pode começar.
        </h2>

        <div className="grid md:grid-cols-3 gap-8">
          {plans.map((p) => (
            <div key={p.title} className="border-t border-petrol/15 pt-8 space-y-4">
              <h3 className="heading-section text-petrol text-2xl">{p.title}</h3>
              <p className="text-petrol/70 leading-relaxed">{p.desc}</p>
            </div>
          ))}
        </div>

        <div className="mt-16">
          <Link href="/planos" className="btn-gold">{tCta('knowPlans')}</Link>
        </div>
      </div>
    </section>
  );
}
