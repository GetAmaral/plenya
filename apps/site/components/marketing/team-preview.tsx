import Image from 'next/image';
import { useTranslations } from 'next-intl';
import { Link } from '@/lib/i18n/navigation';

export function TeamPreview() {
  const t = useTranslations('home');
  const tCta = useTranslations('cta');
  return (
    <section className="bg-cream">
      <div className="site-container section grid lg:grid-cols-[1fr_auto] gap-12 lg:gap-20 items-center">
        <div className="space-y-6 max-w-xl">
          <p className="label-upper text-gold">{t('teamLabel')}</p>
          <h2 className="heading-section text-petrol text-3xl md:text-5xl">
            {t('teamTitle')}
          </h2>
          <p className="text-petrol/80 text-lg leading-relaxed">
            {t('teamP1')}
          </p>
          <p className="text-petrol/70 leading-relaxed">
            {t('teamP2')}
          </p>
          <div className="pt-4">
            <Link href="/equipe" className="btn-outline-dark">
              {tCta('knowTeamPlenya')}
            </Link>
          </div>
        </div>
        <div className="relative w-full max-w-md lg:w-[380px] lg:max-w-none aspect-1064/1891 overflow-hidden bg-petrol/5">
          <Image
            src="/images/team/equipe-formal.jpg"
            alt={t('teamLabel')}
            fill
            className="object-cover object-top"
            sizes="(min-width: 1024px) 380px, 100vw"
          />
        </div>
      </div>
    </section>
  );
}
