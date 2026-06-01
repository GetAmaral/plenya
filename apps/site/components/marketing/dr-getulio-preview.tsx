import Image from 'next/image';
import { useTranslations } from 'next-intl';
import { PlenyaSymbol } from '@plenya/brand/logo';
import { Link } from '@/lib/i18n/navigation';

export function DrGetulioPreview() {
  const t = useTranslations('home');
  const tCta = useTranslations('cta');

  return (
    <section className="bg-paper">
      <div className="site-container section grid gap-16 lg:grid-cols-2 items-center">
        <div className="relative aspect-3/4 overflow-hidden">
          <Image
            src="/images/dr-getulio.jpg"
            alt={t('drGetulioName')}
            fill
            className="object-cover"
            sizes="(min-width: 1024px) 540px, 100vw"
          />
        </div>

        <div className="space-y-8">
          <div className="flex items-center gap-4">
            <PlenyaSymbol aria-hidden="true" className="h-7 w-auto text-gold" />
            <p className="label-upper text-gold">{t('drGetulioLabel')}</p>
          </div>

          <h2 className="heading-section text-petrol text-4xl md:text-5xl">
            {t('drGetulioName')}
          </h2>

          <div className="space-y-1.5">
            <p className="label-upper-sm text-petrol/55">
              {t('drGetulioRole')}
            </p>
            <p className="label-upper-sm text-petrol/40 text-[10px]">
              {t('drGetulioCredentials')}
            </p>
          </div>

          <p className="text-petrol/80 text-lg leading-relaxed max-w-lg">
            {t('drGetulioP1')}
          </p>
          <p className="text-petrol/70 leading-relaxed max-w-lg">
            {t('drGetulioP2')}
          </p>

          <div className="flex flex-wrap gap-5 pt-2">
            <Link href="/dr-getulio" className="btn-gold">
              {tCta('knowDrGetulio')}
            </Link>
            <Link href="/contato" className="btn-outline-dark">
              {tCta('scheduleConsultation')}
            </Link>
          </div>
        </div>
      </div>
    </section>
  );
}
