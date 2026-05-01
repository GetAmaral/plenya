import { useTranslations } from 'next-intl';
import { Link } from '@/lib/i18n/navigation';

export function DiagnosticoStrip() {
  const t = useTranslations('home');
  const tCta = useTranslations('cta');
  return (
    <section className="bg-paper border-b border-petrol/5">
      <div className="site-container py-8 flex flex-col md:flex-row md:items-center md:justify-between gap-4">
        <div className="space-y-1">
          <p className="label-upper text-gold">{t('diagnosticoLabel')}</p>
          <p className="text-petrol/85 leading-relaxed">
            {t('diagnosticoLine1')}
            <span className="text-petrol/60">{t('diagnosticoLine2')}</span>
          </p>
        </div>
        <Link
          href="/diagnostico"
          className="inline-flex items-center self-start md:self-auto px-5 py-2.5 border border-petrol/30 text-petrol hover:bg-petrol hover:text-cream transition label-upper whitespace-nowrap"
        >
          {tCta('doDiagnostic')}
        </Link>
      </div>
    </section>
  );
}
