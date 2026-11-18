import { getTranslations } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { DoctorPortraits } from '@/components/team/doctor-portraits';

export async function TeamPreview() {
  const t = await getTranslations('home');
  const tCta = await getTranslations('cta');
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
        <DoctorPortraits
          className="w-full max-w-md lg:w-[420px] lg:max-w-none text-petrol"
          sizes="(min-width: 1024px) 210px, 50vw"
        />
      </div>
    </section>
  );
}
