import type { Metadata } from 'next';
import { setRequestLocale, getTranslations } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { ComparatorVsConvenio } from '@/components/marketing/comparator-vs-convenio';
import { FaqAccordion } from '@/components/marketing/faq-accordion';
import { FaqSchema } from '@/components/seo/faq-schema';
import { TestimonialsInline } from '@/components/testimonials/testimonials-inline';

type Params = Promise<{ locale: string }>;

export async function generateMetadata({ params }: { params: Params }): Promise<Metadata> {
  const { locale } = await params;
  const t = await getTranslations({ locale, namespace: 'continuum' });
  return {
    title: t('metaTitle'),
    description: t('metaDescription'),
    alternates: {
      canonical: '/continuum',
      languages: {
        'pt-BR': '/continuum',
        en: '/en/continuum',
        'x-default': '/continuum',
      },
    },
  };
}

export default async function ContinuumPage({ params }: { params: Params }) {
  const { locale } = await params;
  setRequestLocale(locale);
  const t = await getTranslations('continuum');

  const continuumFaq = [1, 2, 3, 4, 5, 6, 7].map((n) => ({
    q: t(`faq${n}Q` as 'faq1Q'),
    a: t(`faq${n}A` as 'faq1A'),
  }));

  const continuumNaoE = [1, 2, 3].map((n) => ({
    titulo: t(`notForT${n}` as 'notForT1'),
    body: t(`notForB${n}` as 'notForB1'),
  }));

  const cobertura = [
    { code: locale === 'en' ? 'A' : 'A', name: t('covAName'), quem: t('covAQuem') },
    { code: locale === 'en' ? 'C' : 'G', name: t('covGName'), quem: t('covGQuem') },
    { code: locale === 'en' ? 'T' : 'I', name: t('covIName'), quem: t('covIQuem') },
    { code: locale === 'en' ? 'S' : 'R', name: t('covRName'), quem: t('covRQuem') },
  ];

  const jornada = [1, 2, 3, 4, 5, 6].map((n) => ({
    fase: t(`j${n}Phase` as 'j1Phase'),
    titulo: t(`j${n}Title` as 'j1Title'),
    body: t(`j${n}Body` as 'j1Body'),
  }));

  const modalidades = [
    {
      name: t('modSemestralName'),
      period: t('modSemestralPeriod'),
      body: t('modSemestralBody'),
    },
    {
      name: t('modAnualName'),
      period: t('modAnualPeriod'),
      body: t('modAnualBody'),
      highlight: true,
    },
  ];

  const entra = [1, 2, 3, 4, 5, 6, 7].map((n) => t(`in${n}` as 'in1'));

  const fwCards = [1, 2, 3, 4, 5, 6].map((n) => ({
    title: t(`fw${n}Title` as 'fw1Title'),
    body: t(`fw${n}Body` as 'fw1Body'),
  }));

  return (
    <>
      <FaqSchema items={continuumFaq} />

      {/* HERO */}
      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-24 md:pt-40 md:pb-32">
          <p className="label-upper text-gold mb-6">{t('heroLabel')}</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,5rem)] text-cream max-w-3xl">
            {t('heroTitle')}
          </h1>
          <p className="text-cream/80 text-lg md:text-xl mt-8 max-w-2xl leading-relaxed">
            {t('heroPPart1')}
            <Link
              href="/equipe"
              className="underline decoration-gold/60 underline-offset-4 hover:text-cream/100 transition"
            >
              {t('heroPLinkTeam')}
            </Link>
            {t('heroPPart2')}
            <Link
              href="/escore-plenya"
              className="underline decoration-gold/60 underline-offset-4 hover:text-cream/100 transition"
            >
              {t('heroPLinkScore')}
            </Link>
            {t('heroPPart3')}
            <Link
              href="/metodo-agir"
              className="underline decoration-gold/60 underline-offset-4 hover:text-cream/100 transition"
            >
              {t('heroPLinkMethod')}
            </Link>
            {t('heroPPart4')}
          </p>
        </div>
      </section>

      {/* PROBLEMA — fricção antes da solução */}
      <section className="bg-petrol-800 text-cream">
        <div className="site-container py-20 md:py-28">
          <div className="grid lg:grid-cols-12 gap-10">
            <div className="lg:col-span-4">
              <p className="label-upper text-gold mb-4">{t('problemLabel')}</p>
            </div>
            <div className="lg:col-span-8 space-y-6">
              <h2 className="heading-section text-cream text-3xl md:text-5xl leading-tight max-w-3xl">
                {t('problemTitle')}
              </h2>
              <p className="text-cream/80 text-lg leading-relaxed max-w-2xl">
                {t('problemP1')}
              </p>
              <p className="text-cream/80 text-lg leading-relaxed max-w-2xl">
                {t('problemP2')}
              </p>
            </div>
          </div>
        </div>
      </section>

      {/* SOBRE O PROGRAMA */}
      <section className="bg-cream">
        <div className="site-container section">
          <div className="max-w-3xl space-y-6">
            <p className="label-upper text-gold">{t('aboutLabel')}</p>
            <p className="heading-section text-petrol text-2xl md:text-4xl leading-tight">
              {t('aboutTitle')}
            </p>
            <p className="text-petrol/80 text-lg leading-relaxed">{t('aboutP1')}</p>
            <p className="text-petrol/75 leading-relaxed">{t('aboutP2')}</p>
          </div>
        </div>
      </section>

      {/* PARA QUEM */}
      <section className="bg-paper">
        <div className="site-container section">
          <div className="grid lg:grid-cols-12 gap-12 mb-16">
            <div className="lg:col-span-5">
              <p className="label-upper text-gold mb-6">{t('forWhomLabel')}</p>
              <h2 className="heading-section text-petrol text-3xl md:text-5xl">
                {t('forWhomTitle')}
              </h2>
            </div>
            <p className="lg:col-span-7 text-petrol/75 text-lg leading-relaxed self-end max-w-2xl">
              {t('forWhomDesc')}
            </p>
          </div>

          <div className="grid sm:grid-cols-2 lg:grid-cols-3 gap-x-12 gap-y-12">
            {fwCards.map((c) => (
              <div key={c.title} className="space-y-3">
                <h3 className="heading-section text-petrol text-xl">{c.title}</h3>
                <p className="text-petrol/75 leading-relaxed">{c.body}</p>
              </div>
            ))}
          </div>

          {/* NÃO É — qualificador honesto */}
          <div className="mt-20 grid lg:grid-cols-12 gap-10 border-t border-petrol/10 pt-12">
            <div className="lg:col-span-4 space-y-3">
              <p className="label-upper text-gold">{t('honestyLabel')}</p>
              <h3 className="heading-section text-petrol text-2xl md:text-3xl leading-tight">
                {t('honestyTitlePart1')}
                <em className="not-italic text-gold">{t('honestyTitleEm')}</em>
                {t('honestyTitlePart2')}
              </h3>
              <p className="text-petrol/70 leading-relaxed">{t('honestyDesc')}</p>
            </div>
            <ul className="lg:col-span-8 grid sm:grid-cols-3 gap-x-8 gap-y-8">
              {continuumNaoE.map((it) => (
                <li key={it.titulo} className="space-y-2">
                  <p className="heading-section text-petrol text-lg leading-snug">{it.titulo}</p>
                  <p className="text-petrol/70 text-sm leading-relaxed">{it.body}</p>
                </li>
              ))}
            </ul>
          </div>

          <div className="mt-20 grid lg:grid-cols-[2fr_1fr] gap-10 items-center border-t border-petrol/10 pt-10">
            <div className="space-y-3">
              <p className="label-upper text-gold">{t('doubtLabel')}</p>
              <p className="text-petrol text-xl md:text-2xl heading-section max-w-xl">
                {t('doubtTitle')}
              </p>
              <p className="text-petrol/70 leading-relaxed max-w-xl">{t('doubtDesc')}</p>
            </div>
            <div className="flex lg:justify-end">
              <Link href="/diagnostico" className="btn-gold">
                {t('ctaDoubtLink')}
              </Link>
            </div>
          </div>
        </div>
      </section>

      {/* EQUIPE */}
      <section className="bg-petrol text-cream">
        <div className="site-container section grid lg:grid-cols-2 gap-16 items-start">
          <div className="space-y-6">
            <p className="label-upper text-gold">{t('teamLabel')}</p>
            <h2 className="heading-section text-cream text-3xl md:text-5xl">
              {t('teamTitle')}
            </h2>
            <p className="text-cream/80 text-lg leading-relaxed max-w-xl">{t('teamDesc')}</p>
          </div>
          <div className="lg:pt-20">
            <Link href="/equipe" className="btn-outline-dark border-cream/40 text-cream">
              {t('teamCta')}
            </Link>
          </div>
        </div>
      </section>

      {/* COMO ACONTECE */}
      <section className="bg-paper">
        <div className="site-container section">
          <div className="max-w-3xl space-y-4 mb-16">
            <p className="label-upper text-gold">{t('howLabel')}</p>
            <h2 className="heading-section text-petrol text-3xl md:text-5xl">
              {t('howTitle')}
            </h2>
            <p className="text-petrol/70 text-lg leading-relaxed">{t('howDesc')}</p>
          </div>

          <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-px bg-petrol/15 border-y border-petrol/15 mb-20">
            {cobertura.map((c) => (
              <div key={c.code} className="bg-paper p-8 space-y-4">
                <span className="heading-section text-gold text-5xl block leading-none">{c.code}</span>
                <h3 className="heading-section text-petrol text-lg">{c.name}</h3>
                <p className="label-upper text-petrol/55 text-[10px] tracking-[0.2em] pt-2 border-t border-petrol/10">
                  {c.quem}
                </p>
              </div>
            ))}
          </div>

          <div className="max-w-3xl mb-12">
            <p className="label-upper text-gold mb-3">{t('timelineLabel')}</p>
            <h3 className="heading-section text-petrol text-2xl md:text-3xl">
              {t('timelineTitle')}
            </h3>
          </div>

          <div className="border-t border-petrol/15">
            {jornada.map((j, i) => (
              <article
                key={j.fase}
                className="grid md:grid-cols-[180px_1fr] gap-6 md:gap-12 py-10 border-b border-petrol/10"
              >
                <div className="space-y-2">
                  <p className="label-upper text-gold">{String(i + 1).padStart(2, '0')}</p>
                  <p className="label-upper text-petrol/60">{j.fase}</p>
                </div>
                <div className="space-y-3 max-w-2xl">
                  <h3 className="heading-section text-petrol text-2xl md:text-3xl leading-tight">
                    {j.titulo}
                  </h3>
                  <p className="text-petrol/75 leading-relaxed">{j.body}</p>
                </div>
              </article>
            ))}
          </div>
        </div>
      </section>

      {/* ENTRA · NÃO ENTRA */}
      <section className="bg-paper">
        <div className="site-container section">
          <div className="max-w-3xl space-y-4 mb-16">
            <p className="label-upper text-gold">{t('scopeLabel')}</p>
            <h2 className="heading-section text-petrol text-3xl md:text-5xl">{t('scopeTitle')}</h2>
          </div>

          <div className="grid md:grid-cols-2 gap-12 lg:gap-16">
            <div className="space-y-6">
              <p className="label-upper text-petrol/55">{t('scopeIncludesLabel')}</p>
              <ul className="space-y-4">
                {entra.map((item) => (
                  <li key={item} className="flex gap-4 text-petrol/85 leading-relaxed">
                    <span className="text-gold mt-1.5 leading-none">—</span>
                    <span>{item}</span>
                  </li>
                ))}
              </ul>
            </div>

            <div className="space-y-6 md:border-l md:border-petrol/15 md:pl-12">
              <p className="label-upper text-petrol/55">{t('scopeExcludesLabel')}</p>
              <ul className="space-y-4">
                <li className="flex gap-4 text-petrol/75 leading-relaxed">
                  <span className="text-petrol/40 mt-1.5 leading-none">—</span>
                  <span>{t('out1')}</span>
                </li>
                <li className="flex gap-4 text-petrol/75 leading-relaxed">
                  <span className="text-petrol/40 mt-1.5 leading-none">—</span>
                  <span>
                    {t('out2Part1')}
                    <Link
                      href="/consultas"
                      className="underline decoration-gold/60 underline-offset-4 hover:text-petrol transition"
                    >
                      {t('out2Link')}
                    </Link>
                  </span>
                </li>
                <li className="flex gap-4 text-petrol/75 leading-relaxed">
                  <span className="text-petrol/40 mt-1.5 leading-none">—</span>
                  <span>{t('out3')}</span>
                </li>
              </ul>
            </div>
          </div>
        </div>
      </section>

      <ComparatorVsConvenio title={t('comparatorTitle')} bg="bg-cream" />

      <TestimonialsInline
        bg="bg-paper"
        label={t('testimonialsLabel')}
        title={t('testimonialsTitle')}
        limit={3}
      />

      <FaqAccordion title={t('faqTitle')} items={continuumFaq} />

      {/* MODALIDADES — oferta logo antes do CTA */}
      <section className="bg-cream">
        <div className="site-container section">
          <div className="max-w-3xl space-y-4 mb-16">
            <p className="label-upper text-gold">{t('modLabel')}</p>
            <h2 className="heading-section text-petrol text-3xl md:text-5xl">{t('modTitle')}</h2>
          </div>

          <div className="grid md:grid-cols-2 gap-12 mb-12">
            {modalidades.map((m) => (
              <article
                key={m.name}
                className={
                  m.highlight
                    ? 'border-t-2 border-gold pt-8 space-y-5'
                    : 'border-t border-petrol/20 pt-8 space-y-5'
                }
              >
                <h3 className="heading-section text-petrol text-3xl md:text-4xl">{m.name}</h3>
                <p className="label-upper text-petrol/60">{m.period}</p>
                <p className="text-petrol/75 leading-relaxed">{m.body}</p>
              </article>
            ))}
          </div>

          <div className="border-t border-petrol/10 pt-8 max-w-3xl space-y-2">
            <p className="label-upper text-gold">{t('modPriceLabel')}</p>
            <p className="text-petrol/75 leading-relaxed">{t('modPriceText')}</p>
          </div>
        </div>
      </section>

      {/* CTA + cross-links */}
      <section className="bg-petrol text-cream">
        <div className="site-container section">
          <div className="max-w-2xl space-y-6 mb-20">
            <p className="label-upper text-gold">{t('ctaLabel')}</p>
            <h2 className="heading-section text-cream text-3xl md:text-4xl">{t('ctaTitle')}</h2>
            <p className="text-cream/75 leading-relaxed">{t('ctaDesc')}</p>
            <p>
              <Link href="/contato" className="btn-gold">
                {t('ctaButton')}
              </Link>
            </p>
            <p className="text-cream/60 text-sm pt-2 leading-relaxed">
              {t('ctaDoubtPart1')}
              <Link
                href="/diagnostico"
                className="underline decoration-gold/60 underline-offset-4 hover:text-cream transition"
              >
                {t('ctaDoubtLink')}
              </Link>
              {t('ctaDoubtPart2')}
            </p>
          </div>

          <div className="grid md:grid-cols-3 gap-8">
            <Link href="/escore-plenya" className="border-t border-cream/20 pt-8 space-y-3 group">
              <p className="label-upper text-gold">{t('crossInstrumentLabel')}</p>
              <p className="heading-section text-cream text-xl group-hover:text-gold transition">
                {t('crossInstrumentTitle')}
              </p>
            </Link>
            <Link href="/metodo-agir" className="border-t border-cream/20 pt-8 space-y-3 group">
              <p className="label-upper text-gold">{t('crossMethodLabel')}</p>
              <p className="heading-section text-cream text-xl group-hover:text-gold transition">
                {t('crossMethodTitle')}
              </p>
            </Link>
            <Link href="/consultas" className="border-t border-cream/20 pt-8 space-y-3 group">
              <p className="label-upper text-gold">{t('crossSingleLabel')}</p>
              <p className="heading-section text-cream text-xl group-hover:text-gold transition">
                {t('crossSingleTitle')}
              </p>
            </Link>
          </div>
        </div>
      </section>
    </>
  );
}
