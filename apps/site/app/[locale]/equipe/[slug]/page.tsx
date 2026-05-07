import type { Metadata } from 'next';
import Image from 'next/image';
import { notFound, redirect } from 'next/navigation';
import { setRequestLocale } from 'next-intl/server';
import { isLocale, defaultLocale, locales } from '@/lib/i18n/config';
import { getAllDoctors, getDoctor, localizedBio, localizedShortBio } from '@/lib/team';
import { Link } from '@/lib/i18n/navigation';
import { MdxContent } from '@/components/blog/mdx-content';
import { PhysicianSchema } from '@/components/team/physician-schema';
import { Breadcrumbs } from '@/components/layout/breadcrumbs';

export async function generateStaticParams() {
  const docs = await getAllDoctors();
  // Dr. Getúlio tem página dedicada em /dr-getulio — não duplicar em /equipe/[slug]
  return locales.flatMap((locale) =>
    docs.filter((d) => d.slug !== 'getulio-amaral').map((d) => ({ locale, slug: d.slug })),
  );
}

export async function generateMetadata({ params }: { params: Promise<{ locale: string; slug: string }> }): Promise<Metadata> {
  const { locale, slug } = await params;
  const doctor = await getDoctor(slug);
  if (!doctor) return {};
  return {
    title: doctor.name,
    description: `${doctor.name} — ${doctor.role}. ${localizedShortBio(doctor, locale)}`,
  };
}

export default async function DoctorPage({ params }: { params: Promise<{ locale: string; slug: string }> }) {
  const { locale: rawLocale, slug } = await params;
  const locale = isLocale(rawLocale) ? rawLocale : defaultLocale;

  // Dr. Getúlio sempre vai para a página dedicada
  if (slug === 'getulio-amaral') {
    redirect(locale === 'pt' ? '/dr-getulio' : `/${locale}/dr-getulio`);
  }

  setRequestLocale(locale);

  const doctor = await getDoctor(slug);
  if (!doctor) notFound();

  return (
    <>
      <PhysicianSchema doctor={doctor} />

      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-24 md:pt-40 md:pb-32 grid lg:grid-cols-2 gap-16 items-center">
          <div className="space-y-6">
            <Breadcrumbs
              dark
              items={[
                { label: 'Home', href: '/' },
                { label: 'Equipe', href: '/equipe' },
                { label: doctor.name },
              ]}
            />
            <p className="label-upper text-gold">{doctor.role}</p>
            <h1 className="heading-hero text-[clamp(2.2rem,5vw,4rem)] text-cream">{doctor.name}</h1>
            <p className="label-upper text-cream/60">
              {doctor.credentials}{doctor.rqe ? ` · RQE ${doctor.rqe}` : ''}
            </p>
            <p className="text-cream/85 text-lg leading-relaxed max-w-lg">{localizedShortBio(doctor, locale)}</p>
            <div className="flex flex-wrap gap-5 pt-4">
              <Link href="/contato" className="btn-gold">{locale === 'en' ? 'Book a consultation' : 'Agendar consulta'}</Link>
              {doctor.instagram && (
                <a href={doctor.instagram} target="_blank" rel="noreferrer"
                  className="btn-outline-dark border-cream/40 text-cream">Instagram</a>
              )}
            </div>
          </div>
          <div className="relative aspect-[3/4] overflow-hidden">
            {doctor.photo ? (
              <Image
                src={doctor.photo}
                alt={doctor.name}
                fill
                priority
                className="object-cover"
                sizes="(min-width: 1024px) 540px, 100vw"
              />
            ) : (
              <div className="absolute inset-0 bg-cream-200" />
            )}
          </div>
        </div>
      </section>

      {doctor.specialties.length > 0 && (
        <section className="bg-cream">
          <div className="site-container py-10 flex flex-wrap gap-3">
            {doctor.specialties.map((s) => (
              <span key={s} className="label-upper text-petrol/60 border border-petrol/15 px-5 py-2">{s}</span>
            ))}
          </div>
        </section>
      )}

      <section className="bg-cream">
        <div className="site-narrow section prose-plenya">
          <MdxContent source={localizedBio(doctor, locale)} />
        </div>
      </section>

      <section className="bg-paper">
        <div className="site-container section grid md:grid-cols-2 gap-8">
          <Link href="/metodo-agir" className="border-t border-petrol/15 pt-8 space-y-3 group">
            <p className="label-upper text-gold">{locale === 'en' ? 'Method' : 'Método'}</p>
            <p className="heading-section text-petrol text-xl group-hover:text-gold transition">{locale === 'en' ? 'Discover The ACTS Method →' : 'Conheça o Método AGIR →'}</p>
          </Link>
          <Link href="/continuum" className="border-t border-petrol/15 pt-8 space-y-3 group">
            <p className="label-upper text-gold">Continuum</p>
            <p className="heading-section text-petrol text-xl group-hover:text-gold transition">Continuum Plenya →</p>
          </Link>
        </div>
      </section>
    </>
  );
}
