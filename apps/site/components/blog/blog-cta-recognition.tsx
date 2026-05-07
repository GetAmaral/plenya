import { getTranslations } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';

export async function BlogCTARecognition() {
  const t = await getTranslations('blogCta');
  return (
    <div className="my-16 border-l-2 border-gold pl-8 py-6 space-y-6">
      <p className="label-upper text-gold">{t('recognitionLabel')}</p>
      <p className="heading-section text-petrol text-2xl max-w-xl">
        {t('recognitionTitle')}
      </p>
      <p className="text-petrol/75 max-w-xl leading-relaxed">
        {t('recognitionDesc')}
      </p>
      <div className="flex flex-wrap gap-3">
        <Link href="/diagnostico" className="btn-gold">
          {t('ctaDiagnose')}
        </Link>
        <Link
          href="/contato"
          className="inline-flex items-center px-6 py-3 border border-petrol/30 text-petrol hover:bg-petrol hover:text-cream transition label-upper"
        >
          {t('ctaContact')}
        </Link>
      </div>
    </div>
  );
}
