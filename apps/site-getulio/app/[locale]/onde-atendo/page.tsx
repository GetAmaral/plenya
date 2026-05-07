import type { Metadata } from 'next';
import Image from 'next/image';
import { setRequestLocale, getTranslations } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { ClinicsSchema } from '@/components/seo/clinics-schema';
import { BreadcrumbSchema } from '@/components/seo/breadcrumb-schema';

type ConsultClinic = {
  name: string;
  role: string;
  tagline: string;
  body: string;
  address: string;
  href: string;
};

type HospitalUnit = {
  name: string;
  role: string;
  body: string;
  address: string;
};

export async function generateMetadata({
  params,
}: {
  params: Promise<{ locale: string }>;
}): Promise<Metadata> {
  const { locale } = await params;
  const t = await getTranslations({ locale, namespace: 'ondeAtendo' });
  return {
    title: t('metaTitle'),
    description: t('metaDescription'),
    alternates: {
      canonical: locale === 'en' ? '/en/onde-atendo' : '/onde-atendo',
      languages: { 'pt-BR': '/onde-atendo', en: '/en/onde-atendo' },
    },
  };
}

export default async function OndeAtendoPage({
  params,
}: {
  params: Promise<{ locale: string }>;
}) {
  const { locale } = await params;
  setRequestLocale(locale);
  const t = await getTranslations({ locale, namespace: 'ondeAtendo' });
  const consultas = t.raw('consultas') as ConsultClinic[];
  const hospital = t.raw('hospital') as HospitalUnit[];

  return (
    <article>
      <ClinicsSchema locale={locale} />
      <BreadcrumbSchema
        items={[
          { name: t('breadcrumbHome'), url: '/' },
          { name: t('breadcrumbCurrent') },
        ]}
      />

      <header className="editorial-container pt-16 md:pt-24 pb-12">
        <p className="label-meta mb-6">{t('kicker')}</p>
        <h1 className="heading-display text-[clamp(2.2rem,5vw,3.8rem)] max-w-3xl">
          {t('h1')}
        </h1>
        <p className="prose-body mt-8 max-w-2xl">{t('lead')}</p>
      </header>

      <section className="editorial-container pb-12">
        <div className="relative aspect-[16/10] w-full overflow-hidden">
          <Image
            src="/images/getulio-clinico.jpg"
            alt={t('portraitAlt')}
            fill
            className="object-cover object-top"
            sizes="(min-width: 1024px) 1100px, 100vw"
          />
        </div>
      </section>

      {/* Atendimento ambulatorial — duas vias de consulta, prominente */}
      <section className="editorial-container pb-16">
        <div className="border-t border-rule pt-12 md:pt-16">
          <p className="label-meta-lg text-bordo mb-4">{t('consultsKicker')}</p>
          <p className="font-serif text-lg md:text-xl text-ink-soft leading-relaxed max-w-2xl mb-12">
            {t('consultsLead')}
          </p>

          <div className="grid md:grid-cols-2 gap-10 lg:gap-14">
            {consultas.map((c) => (
              <a
                key={c.name}
                href={c.href}
                target="_blank"
                rel="noreferrer"
                className="group block bg-paper border border-rule p-8 md:p-10 hover:border-bordo transition-colors"
              >
                <div className="space-y-5">
                  <p className="label-meta text-bordo">{c.role}</p>
                  <h2 className="heading-section text-2xl md:text-3xl text-ink group-hover:text-bordo transition-colors">
                    {c.name}
                  </h2>
                  <p className="font-serif italic text-ink-muted text-lg leading-snug">
                    {c.tagline}
                  </p>
                  <p className="font-serif text-ink-soft leading-relaxed">{c.body}</p>
                  <p className="font-sans text-sm text-ink-muted">{c.address}</p>
                  <p className="font-sans text-sm text-bordo group-hover:underline pt-2">
                    {t('scheduleHere')}
                  </p>
                </div>
              </a>
            ))}
          </div>
        </div>

        {/* Em dúvida sobre qual via? */}
        <div className="border-t border-rule pt-10 mt-16 grid md:grid-cols-[1fr_2fr] gap-6 md:gap-12">
          <p className="label-meta text-bordo">{t('patientHelpKicker')}</p>
          <p className="font-serif text-ink-soft leading-relaxed max-w-2xl">
            {t('patientHelpBody')}
          </p>
        </div>
      </section>

      {/* Atuação hospitalar — secundário, transparência sobre escopo */}
      <section className="border-t border-rule bg-paper">
        <div className="editorial-container py-16 md:py-20">
          <p className="label-meta-lg text-bordo mb-4">{t('hospitalKicker')}</p>
          <p className="font-serif text-ink-soft text-base leading-relaxed max-w-2xl mb-10">
            {t('hospitalLead')}
          </p>

          <ul className="grid md:grid-cols-2 gap-x-10 gap-y-8">
            {hospital.map((h) => (
              <li key={h.name} className="space-y-2">
                <p className="label-meta text-ink-muted">{h.role}</p>
                <h3 className="font-serif text-xl text-ink leading-tight">{h.name}</h3>
                <p className="font-serif text-sm text-ink-soft leading-relaxed">{h.body}</p>
                <p className="font-sans text-xs text-ink-muted">{h.address}</p>
              </li>
            ))}
          </ul>
        </div>
      </section>

      {/* Outros canais */}
      <section className="border-t border-rule">
        <div className="editorial-container py-12 md:py-16">
          <div className="grid md:grid-cols-[1fr_2fr] gap-6 md:gap-12 items-baseline">
            <p className="label-meta text-bordo">{t('outrosCanaisKicker')}</p>
            <p className="font-serif text-ink-soft text-sm">
              {t('outrosCanaisBody')}
              <Link href="/contato" className="link-text">
                {t('outrosCanaisLink')}
              </Link>
              .
            </p>
          </div>
        </div>
      </section>
    </article>
  );
}
