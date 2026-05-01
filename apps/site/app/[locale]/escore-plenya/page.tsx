import type { Metadata } from 'next';
import Image from 'next/image';
import { setRequestLocale, getTranslations } from 'next-intl/server';
import { Link, type Href } from '@/lib/i18n/navigation';
import { RadarAgir } from '@/components/escore/RadarAgir';
import { FaqAccordion } from '@/components/marketing/faq-accordion';
import { FaqSchema } from '@/components/seo/faq-schema';

type Params = Promise<{ locale: string }>;

export async function generateMetadata({ params }: { params: Params }): Promise<Metadata> {
  const { locale } = await params;
  const t = await getTranslations({ locale, namespace: 'score' });
  return {
    title: t('metaTitle'),
    description: t('metaDescription'),
    alternates: {
      canonical: locale === 'en' ? '/en/plenya-score' : '/escore-plenya',
      languages: {
        'pt-BR': '/escore-plenya',
        en: '/en/plenya-score',
        'x-default': '/escore-plenya',
      },
    },
  };
}

export default async function ScorePage({ params }: { params: Params }) {
  const { locale } = await params;
  setRequestLocale(locale);
  const t = await getTranslations('score');
  const tCta = await getTranslations('cta');

  const escoreFaq = [1, 2, 3, 4, 5, 6].map((n) => ({
    q: t(`faq${n}Q` as 'faq1Q'),
    a: t(`faq${n}A` as 'faq1A'),
  }));

  const steps = [
    { n: '01', title: t('step1Title'), body: t('step1Body') },
    { n: '02', title: t('step2Title'), body: t('step2Body') },
    { n: '03', title: t('step3Title'), body: t('step3Body') },
  ];

  const tiers: Array<{ title: string; anchor: string; href?: Href; desc: string }> = [
    {
      title: t('tier1Title'),
      anchor: t('tier1Anchor'),
      href: '/escore-plenya/avaliar',
      desc: t('tier1Desc'),
    },
    {
      title: t('tier2Title'),
      anchor: t('tier2Anchor'),
      href: '/consultas',
      desc: t('tier2Desc'),
    },
    {
      title: t('tier3Title'),
      anchor: t('tier3Anchor'),
      href: '/continuum',
      desc: t('tier3Desc'),
    },
  ];

  const bands = [
    { label: t('band1Label'), range: '0–30', color: 'bg-[#c1542d]' },
    { label: t('band2Label'), range: '31–50', color: 'bg-[#d89345]' },
    { label: t('band3Label'), range: '51–70', color: 'bg-[#caa54a]' },
    { label: t('band4Label'), range: '71–85', color: 'bg-[#7ba07a]' },
    { label: t('band5Label'), range: '86–100', color: 'bg-[#3f7d6e]' },
  ];

  return (
    <>
      <FaqSchema items={escoreFaq} />

      {/* HERO */}
      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-24 md:pt-40 md:pb-32">
          <p className="label-upper text-gold mb-6">{t('heroLabel')}</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,5rem)] text-cream max-w-2xl">
            {t('heroTitle')}
          </h1>
          <p className="text-cream/75 text-xl md:text-2xl mt-8 max-w-2xl leading-relaxed font-light">
            {t('heroP')}
          </p>
        </div>
      </section>

      {/* O QUE É */}
      <section className="bg-cream">
        <div className="site-container section">
          <div className="max-w-3xl space-y-6">
            <p className="label-upper text-gold">{t('whatLabel')}</p>
            <p className="heading-section text-petrol text-2xl md:text-4xl leading-tight">
              {t('whatTitle')}
            </p>
            <p className="text-petrol/80 text-lg leading-relaxed">{t('whatP')}</p>
          </div>
        </div>
      </section>

      {/* COMO FUNCIONA */}
      <section className="bg-paper">
        <div className="site-container section">
          <div className="max-w-3xl mb-14 space-y-4">
            <p className="label-upper text-gold">{t('howLabel')}</p>
            <h2 className="heading-section text-petrol text-3xl md:text-4xl">
              {t('howTitle')}
            </h2>
          </div>

          <ol className="grid md:grid-cols-3 gap-10 max-w-5xl">
            {steps.map((s) => (
              <li key={s.n} className="space-y-4">
                <p className="font-sans text-gold text-7xl md:text-8xl font-light leading-none tabular-nums">
                  {s.n}
                </p>
                <h3 className="heading-section text-petrol text-2xl md:text-3xl">{s.title}</h3>
                <p className="text-petrol/75 leading-relaxed">{s.body}</p>
              </li>
            ))}
          </ol>
        </div>
      </section>

      {/* PONTUAÇÃO */}
      <section className="bg-cream">
        <div className="site-container section grid lg:grid-cols-[1fr_auto] gap-16 items-center">
          <div className="space-y-6 max-w-xl">
            <p className="label-upper text-gold">{t('scoreLabel')}</p>
            <h2 className="heading-section text-petrol text-3xl md:text-5xl leading-tight">
              {t('scoreTitlePart1')}
              <em className="not-italic text-gold">{t('scoreTitleEm')}</em>
            </h2>
            <p className="text-petrol/80 text-lg leading-relaxed">{t('scoreP1Pt1')}</p>
            <p className="text-petrol/70 leading-relaxed">{t('scoreP2')}</p>
          </div>

          <RadarAgir />
        </div>
      </section>

      {/* FAIXAS */}
      <section className="bg-paper border-y border-petrol/10">
        <div className="site-container py-16 md:py-20">
          <div className="max-w-3xl mb-10 space-y-3">
            <p className="label-upper text-gold">{t('bandsLabel')}</p>
            <h2 className="heading-section text-petrol text-2xl md:text-3xl">
              {t('bandsTitle')}
            </h2>
          </div>

          <div className="grid sm:grid-cols-2 lg:grid-cols-5 gap-3 md:gap-4 max-w-5xl">
            {bands.map((f) => (
              <div
                key={f.label}
                className="flex items-center gap-3 p-4 border border-petrol/10 rounded-sm"
              >
                <span className={`w-3 h-10 rounded-sm ${f.color} flex-shrink-0`} aria-hidden />
                <div>
                  <p className="heading-section text-petrol text-base leading-tight">{f.label}</p>
                  <p className="font-mono text-xs text-petrol/50 mt-0.5">{f.range}</p>
                </div>
              </div>
            ))}
          </div>

          <p className="text-petrol/60 leading-relaxed mt-8 max-w-3xl">
            {t('bandsCaptionPart1')}
            <strong className="text-petrol">{t('bandsCaptionStrong1')}</strong>
            {t('bandsCaptionPart2')}
            <strong className="text-petrol">{t('bandsCaptionStrong2')}</strong>
            {t('bandsCaptionPart3')}
          </p>
        </div>
      </section>

      {/* EVOLUÇÃO */}
      <section className="bg-cream">
        <div className="site-container section grid lg:grid-cols-[1fr_auto] gap-16 items-center">
          <div className="space-y-6 max-w-xl">
            <p className="label-upper text-gold">{t('evolutionLabel')}</p>
            <h2 className="heading-section text-petrol text-3xl md:text-4xl">
              {t('evolutionTitle')}
            </h2>
            <p className="text-petrol/80 text-lg leading-relaxed">{t('evolutionP')}</p>
          </div>

          <figure className="flex flex-col items-center gap-3">
            <svg
              viewBox="0 0 420 220"
              className="w-full max-w-md md:max-w-lg"
              aria-label={t('evolutionChartAria')}
            >
              {[40, 90, 140, 190].map((y) => (
                <line
                  key={y}
                  x1="40"
                  y1={y}
                  x2="400"
                  y2={y}
                  stroke="#063b4f"
                  strokeOpacity="0.06"
                  strokeWidth="1"
                />
              ))}
              {[
                { y: 40, v: '100' },
                { y: 90, v: '75' },
                { y: 140, v: '50' },
                { y: 190, v: '25' },
              ].map((tk) => (
                <text
                  key={tk.v}
                  x="32"
                  y={tk.y + 4}
                  textAnchor="end"
                  fontFamily="monospace"
                  fontSize="10"
                  fill="#063b4f"
                  fillOpacity="0.4"
                >
                  {tk.v}
                </text>
              ))}
              <polyline
                points="60,160 140,140 220,110 300,85 380,68"
                fill="none"
                stroke="#b38645"
                strokeWidth="2.5"
                strokeLinecap="round"
                strokeLinejoin="round"
              />
              {[
                { x: 60, y: 160 },
                { x: 140, y: 140 },
                { x: 220, y: 110 },
                { x: 300, y: 85 },
                { x: 380, y: 68 },
              ].map((p, i) => (
                <circle
                  key={i}
                  cx={p.x}
                  cy={p.y}
                  r="4"
                  fill="#063b4f"
                  stroke="#fbfaf6"
                  strokeWidth="2"
                />
              ))}
              {['T0', '3m', '6m', '9m', '12m'].map((label, i) => (
                <text
                  key={label}
                  x={60 + i * 80}
                  y="210"
                  textAnchor="middle"
                  fontFamily="monospace"
                  fontSize="10"
                  fill="#063b4f"
                  fillOpacity="0.4"
                >
                  {label}
                </text>
              ))}
            </svg>
            <p className="label-upper text-petrol/40 text-center text-[10px]">
              {t('evolutionChartCaption')}
            </p>
          </figure>
        </div>
      </section>

      {/* IMAGEM ÂNCORA */}
      <section className="bg-cream">
        <div className="site-container">
          <div className="relative aspect-[16/7] overflow-hidden">
            <Image
              src="/images/hero-score.jpg"
              alt={t('anchorImageAlt')}
              fill
              className="object-cover"
              sizes="100vw"
            />
          </div>
        </div>
      </section>

      {/* VERSÕES */}
      <section className="bg-cream">
        <div className="site-container section">
          <div className="max-w-3xl mb-12 space-y-4">
            <p className="label-upper text-gold">{t('tiersLabel')}</p>
            <h2 className="heading-section text-petrol text-3xl md:text-4xl">{t('tiersTitle')}</h2>
            <p className="text-petrol/70 leading-relaxed">{t('tiersDesc')}</p>
          </div>
          <div className="grid md:grid-cols-3 gap-10">
            {tiers.map((tier, i) => (
              <div key={tier.title} className="border-t-2 border-petrol/15 pt-8 space-y-5">
                <span className="font-sans text-gold text-6xl md:text-7xl font-light block leading-none tabular-nums">
                  0{i + 1}
                </span>
                <h3 className="heading-section text-petrol text-3xl md:text-4xl">{tier.title}</h3>
                {tier.href ? (
                  <Link
                    href={tier.href}
                    className="inline-flex items-center gap-1 text-gold text-lg font-medium underline decoration-gold/50 underline-offset-4 hover:decoration-gold transition"
                  >
                    {tier.anchor}
                    <span aria-hidden className="transition-transform group-hover:translate-x-1">
                      →
                    </span>
                  </Link>
                ) : (
                  <p className="text-petrol/50 text-base italic">{tier.anchor}</p>
                )}
                <p className="text-petrol/75 text-base leading-relaxed">{tier.desc}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      <FaqAccordion title={t('faqTitle')} items={escoreFaq} />

      {/* SAÍDA */}
      <section className="bg-petrol text-cream">
        <div className="site-container section text-center space-y-8">
          <div className="max-w-2xl mx-auto space-y-4">
            <p className="label-upper text-gold">{t('exitLabel')}</p>
            <h2 className="heading-section text-cream text-3xl md:text-4xl">
              {t('exitTitleLine1')} <br className="hidden md:block" />
              {t('exitTitleLine2')}
            </h2>
          </div>
          <Link href="/metodo-agir" className="btn-gold">
            {tCta('knowMethod')}
          </Link>
        </div>
      </section>
    </>
  );
}
