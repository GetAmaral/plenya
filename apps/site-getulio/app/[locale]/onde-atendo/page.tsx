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
  mapsUrl: string;
  whatsappUrl: string;
  whatsappLabel: string;
  siteUrl: string;
  instagramUrl: string;
  instagramHandle: string;
  email: string;
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
      canonical: locale === 'en' ? '/en/where-i-practice' : '/onde-atendo',
      languages: { 'pt-BR': '/onde-atendo', pt: '/onde-atendo', en: '/en/where-i-practice' },
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
  const hospitalNefro = t.raw('hospitalNefro') as HospitalUnit[];
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
              <article
                key={c.name}
                className="bg-paper border border-rule p-8 md:p-10 flex flex-col"
              >
                <div className="space-y-5">
                  <p className="label-meta text-bordo">{c.role}</p>
                  <h2 className="heading-section text-2xl md:text-3xl text-ink">
                    {c.name}
                  </h2>
                  <p className="font-serif italic text-ink-muted text-lg leading-snug">
                    {c.tagline}
                  </p>
                  <p className="font-serif text-ink-soft leading-relaxed">{c.body}</p>
                </div>

                {/* Bloco de contato — endereço, WhatsApp, site, Instagram */}
                <dl className="mt-6 pt-6 border-t border-rule grid grid-cols-[110px_1fr] gap-x-4 gap-y-3 text-sm">
                  <dt className="label-meta text-ink-muted">{t('labelAddress')}</dt>
                  <dd className="font-serif text-ink-soft">
                    <p>{c.address}</p>
                    {c.mapsUrl && (
                      <a
                        href={c.mapsUrl}
                        target="_blank"
                        rel="noreferrer"
                        className="link-text font-sans text-xs"
                      >
                        {t('labelDirections')}
                      </a>
                    )}
                  </dd>

                  {c.whatsappUrl && c.whatsappLabel && (
                    <>
                      <dt className="label-meta text-ink-muted">{t('labelWhatsApp')}</dt>
                      <dd>
                        <a
                          href={c.whatsappUrl}
                          target="_blank"
                          rel="noreferrer"
                          className="link-text font-mono text-sm"
                        >
                          {c.whatsappLabel}
                        </a>
                      </dd>
                    </>
                  )}

                  <dt className="label-meta text-ink-muted">{t('labelSite')}</dt>
                  <dd>
                    <a
                      href={c.siteUrl}
                      target="_blank"
                      rel="noreferrer"
                      className="link-text font-sans"
                    >
                      {c.siteUrl.replace(/^https?:\/\//, '').replace(/\/$/, '')}
                    </a>
                  </dd>

                  {c.instagramUrl && c.instagramHandle && (
                    <>
                      <dt className="label-meta text-ink-muted">{t('labelInstagram')}</dt>
                      <dd>
                        <a
                          href={c.instagramUrl}
                          target="_blank"
                          rel="noreferrer"
                          className="link-text font-sans"
                        >
                          {c.instagramHandle}
                        </a>
                      </dd>
                    </>
                  )}

                  {c.email && (
                    <>
                      <dt className="label-meta text-ink-muted">{t('labelEmail')}</dt>
                      <dd>
                        <a
                          href={`mailto:${c.email}`}
                          className="link-text font-sans break-all"
                        >
                          {c.email}
                        </a>
                      </dd>
                    </>
                  )}
                </dl>

                {/* CTA primário: WhatsApp se existir, senão site */}
                <div className="mt-6">
                  {c.whatsappUrl ? (
                    <a
                      href={c.whatsappUrl}
                      target="_blank"
                      rel="noreferrer"
                      className="btn-gold w-full text-center block"
                    >
                      {t('ctaWhatsApp')}
                    </a>
                  ) : (
                    <a
                      href={c.siteUrl}
                      target="_blank"
                      rel="noreferrer"
                      className="btn-gold w-full text-center block"
                    >
                      {t('ctaSite')}
                    </a>
                  )}
                </div>
              </article>
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

      {/* Atendimento hospitalar de nefrologia — 3 hospitais */}
      <section className="border-t border-rule bg-paper">
        <div className="editorial-container py-16 md:py-20">
          <p className="label-meta-lg text-bordo mb-4">{t('hospitalNefroKicker')}</p>
          <p className="font-serif text-ink-soft text-base leading-relaxed max-w-2xl mb-10">
            {t('hospitalNefroLead')}
          </p>

          <ul className="grid md:grid-cols-3 gap-x-10 gap-y-8">
            {hospitalNefro.map((h) => (
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

      {/* Hemodiálise — DaVita Londrina, responsabilidade técnica */}
      <section className="border-t border-rule">
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
