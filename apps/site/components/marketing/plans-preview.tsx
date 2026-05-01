import { useTranslations } from 'next-intl';
import { PlenyaSymbol } from '@plenya/brand/logo';
import { Link, type Href } from '@/lib/i18n/navigation';

export function PlansPreview() {
  const t = useTranslations('plans');
  const tCta = useTranslations('cta');

  const offers: Array<{
    label: string;
    title: string;
    desc: string;
    href: Href;
    highlight?: boolean;
  }> = [
    {
      label: t('consultLabel'),
      title: t('consultTitle'),
      desc: t('consultDesc'),
      href: '/consultas',
    },
    {
      label: t('continuumLabel'),
      title: t('continuumTitle'),
      desc: t('continuumDesc'),
      href: '/continuum',
      highlight: true,
    },
  ];

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
          <p className="label-upper text-gold">{t('sectionLabel')}</p>
        </div>
        <h2 className="heading-section text-petrol text-3xl md:text-5xl max-w-2xl mb-16">
          {t('sectionTitle')}
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
