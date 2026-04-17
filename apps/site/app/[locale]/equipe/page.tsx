import type { Metadata } from 'next';
import { setRequestLocale } from 'next-intl/server';
import { isLocale, defaultLocale } from '@/lib/i18n/config';
import { Link } from '@/lib/i18n/navigation';
import { categoryLabels, categoryOrder, getDoctorsByCategory } from '@/lib/team';
import { DoctorCard } from '@/components/team/doctor-card';

export const metadata: Metadata = {
  title: 'Equipe Plenya',
  description: 'Médicos e equipe multidisciplinar da Plenya — guardiões do Método AGIR.',
};

export default async function TeamPage({ params }: { params: Promise<{ locale: string }> }) {
  const { locale: rawLocale } = await params;
  const locale = isLocale(rawLocale) ? rawLocale : defaultLocale;
  setRequestLocale(locale);

  const grouped = await getDoctorsByCategory();

  return (
    <>
      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-24 md:pt-40 md:pb-32">
          <p className="label-upper text-gold mb-6">Guardiões do Método</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,5rem)] text-cream max-w-xl">
            Equipe Plenya
          </h1>
          <p className="text-cream/70 text-lg mt-6 max-w-lg">
            Médicos e profissionais multidisciplinares. Cuidado integrado com método, proximidade e
            responsabilidade compartilhada.
          </p>
        </div>
      </section>

      <section className="bg-cream">
        <div className="site-container section space-y-20">
          {categoryOrder.map((cat) => {
            const doctors = grouped[cat];
            if (!doctors.length) return null;
            return (
              <div key={cat}>
                <div className="flex items-baseline gap-4 mb-10">
                  <h2 className="heading-section text-petrol text-2xl md:text-3xl">{categoryLabels[cat]}</h2>
                  <span className="label-upper text-petrol/40">
                    {doctors.length} {doctors.length === 1 ? 'profissional' : 'profissionais'}
                  </span>
                </div>
                <div className="grid gap-8 sm:grid-cols-2 lg:grid-cols-3">
                  {doctors.map((d) => (
                    <DoctorCard key={d.slug} doctor={d} />
                  ))}
                </div>
              </div>
            );
          })}
        </div>
      </section>

      <section className="bg-paper">
        <div className="site-container section text-center space-y-6">
          <p className="heading-section text-petrol text-2xl md:text-3xl">
            Quer fazer parte da equipe Plenya?
          </p>
          <Link href="/contato" className="btn-gold">Entre em contato</Link>
        </div>
      </section>
    </>
  );
}
