import { useTranslations } from 'next-intl';
import { PlenyaSymbol } from '@plenya/brand/logo';
import { Link } from '@/lib/i18n/navigation';

const offers = [
  {
    label: 'Avulso',
    title: 'Consulta Plenya',
    desc: 'Presencial em Londrina ou online, com a equipe médica Plenya. Escuta longa e leitura cuidadosa dos dados.',
    href: '/consultas',
  },
  {
    label: 'Programa contínuo',
    title: 'Continuum Plenya',
    desc: 'Programa semestral ou anual, 100% online. Médico, nutricionista, psicólogo e educador físico em um plano único, mensurado pelo Escore Plenya.',
    href: '/continuum',
    highlight: true,
  },
];

export function PlansPreview() {
  const tCta = useTranslations('cta');
  return (
    <section className="bg-paper relative overflow-hidden">
      {/* P watermark — selo discreto da marca no canto superior direito */}
      <PlenyaSymbol
        aria-hidden="true"
        focusable="false"
        className="hidden md:block absolute -top-6 right-8 lg:right-20 h-40 lg:h-56 w-auto text-petrol/[0.06] pointer-events-none"
      />
      <div className="relative site-container section">
        <div className="flex items-center gap-4 mb-6">
          <PlenyaSymbol aria-hidden="true" className="h-7 w-auto text-gold" />
          <p className="label-upper text-gold">Como começar</p>
        </div>
        <h2 className="heading-section text-petrol text-3xl md:text-5xl max-w-2xl mb-16">
          Dois caminhos para entrar na medicina Plenya.
        </h2>

        <div className="grid md:grid-cols-2 gap-10">
          {offers.map((o) => (
            <Link
              key={o.title}
              href={o.href}
              className={`group block pt-8 space-y-5 border-t-2 transition ${
                o.highlight ? 'border-gold' : 'border-petrol/20'
              }`}
            >
              <p className="label-upper text-gold">{o.label}</p>
              <h3 className="heading-section text-petrol text-3xl md:text-4xl group-hover:text-gold transition">
                {o.title} →
              </h3>
              <p className="text-petrol/75 text-lg leading-snug">{o.desc}</p>
            </Link>
          ))}
        </div>

        <div className="mt-16">
          <Link href="/continuum" className="btn-gold">{tCta('knowPlans')}</Link>
        </div>
      </div>
    </section>
  );
}
