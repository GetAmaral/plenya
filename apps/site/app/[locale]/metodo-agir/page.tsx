import type { Metadata } from 'next';
import Image from 'next/image';
import { setRequestLocale, getTranslations } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { getAgirLetters } from '@/lib/agir-structure';
import { getAgirPillars } from '@plenya/brand';
import { BookReferenceBlock } from '@/components/marketing/book-reference';
import { RelatedBlogPosts } from '@/components/marketing/related-blog-posts';
import { defaultLocale, isLocale } from '@/lib/i18n/config';

type Params = Promise<{ locale: string }>;

export async function generateMetadata({ params }: { params: Params }): Promise<Metadata> {
  const { locale } = await params;
  const t = await getTranslations({ locale, namespace: 'method' });
  return {
    title: t('metaTitle'),
    description: t('metaDescription'),
    alternates: {
      canonical: locale === 'en' ? '/en/acts-method' : '/metodo-agir',
      languages: {
        'pt-BR': '/metodo-agir', pt: '/metodo-agir',
        en: '/en/acts-method',
        'x-default': '/metodo-agir',
      },
    },
  };
}

export default async function MethodPage({ params }: { params: Params }) {
  const { locale } = await params;
  setRequestLocale(locale);
  const t = await getTranslations('method');
  const tCta = await getTranslations('cta');
  const letters = getAgirLetters(locale);
  const pillarsBrand = getAgirPillars(locale);

  // Mapeia os 4 pilares longos por slug — combina dados do brand (code/name)
  // com prosa específica desta página (territory/intro/quote/monitorados).
  const pilares = pillarsBrand.map((p, idx) => {
    const n = idx + 1;
    return {
      code: p.code,
      slug: p.slug,
      name: t(`p${n}Name` as 'p1Name'),
      territory: t(`p${n}Territory` as 'p1Territory'),
      intro: t(`p${n}Intro` as 'p1Intro'),
      quote: t(`p${n}Quote` as 'p1Quote'),
      monitoramos: getMonitored(t, n),
    };
  });

  const conviccoes = [1, 2, 3, 4, 5].map((n) => ({
    n: `0${n}`,
    text: t(`c${n}Text` as 'c1Text'),
    sub: t(`c${n}Sub` as 'c1Sub'),
  }));

  return (
    <>
      {/* Hero — fotográfico */}
      <section className="relative bg-petrol text-cream overflow-hidden">
        <Image
          src="/images/metodo-agir-hero.jpg"
          alt=""
          fill
          priority
          className="object-cover object-top opacity-40"
          sizes="100vw"
        />
        <div className="absolute inset-0 bg-linear-to-b from-petrol/70 via-petrol/60 to-petrol" />
        <div className="relative site-container pt-32 pb-24 md:pt-40 md:pb-32">
          <p className="label-upper text-gold mb-6">{t('heroLabel')}</p>
          <h1 className="heading-hero text-[clamp(2.8rem,7vw,5.5rem)] text-cream max-w-3xl">
            {t('heroTitle')}
          </h1>
          <p className="text-cream/80 text-lg md:text-xl mt-8 max-w-2xl leading-relaxed">
            {t('heroSubtitle')}
          </p>
        </div>
      </section>

      {/* Premissa */}
      <section className="bg-cream">
        <div className="site-container section">
          <div className="grid lg:grid-cols-[1fr_2fr] gap-16 items-start">
            <p className="label-upper text-gold lg:pt-3">{t('premiseLabel')}</p>
            <div className="space-y-6 text-petrol/85 text-lg leading-relaxed">
              <p className="heading-section text-petrol text-2xl md:text-4xl leading-tight">
                {t('premiseTitle')}
              </p>
              <p>{t('premiseP1')}</p>
              <p>{t('premiseP2')}</p>
            </div>
          </div>
        </div>
      </section>

      {/* Os 4 pilares */}
      <section className="bg-paper">
        <div className="site-container section space-y-24 md:space-y-32">
          {pilares.map((pilar, idx) => (
            <article
              key={pilar.slug}
              className="grid lg:grid-cols-[260px_1fr] gap-8 lg:gap-20 border-t border-petrol/15 pt-12"
            >
              <div className="space-y-3">
                <p className="label-upper text-gold">
                  {t('pillarLabel')} 0{idx + 1}
                </p>
                <span className="heading-hero text-gold text-[clamp(8rem,20vw,17rem)] leading-[0.82] block">
                  {pilar.code}
                </span>
              </div>

              <div className="space-y-10">
                <header className="space-y-3">
                  <h2 className="heading-section text-petrol text-3xl md:text-4xl lg:text-5xl">
                    {pilar.name}
                  </h2>
                  <p className="label-upper text-petrol/55">{pilar.territory}</p>
                </header>

                <p className="text-petrol/85 text-lg leading-relaxed max-w-2xl">{pilar.intro}</p>

                <blockquote className="border-l-2 border-gold pl-6 py-2 text-petrol/80 text-lg italic max-w-2xl">
                  &ldquo;{pilar.quote}&rdquo;
                </blockquote>

                {(() => {
                  const letter = letters.find((l) => l.code === pilar.code);
                  if (!letter) return null;
                  return (
                    <div className="space-y-4">
                      <p className="label-upper text-petrol/55">
                        {t('pillarsClinicalCount', { count: letter.pillarCount })}
                      </p>
                      <div className="space-y-5">
                        {letter.groups.map((group, gi) => (
                          <div key={gi} className="space-y-2">
                            {group.label && (
                              <p className="label-upper text-petrol/40 text-[10px]">
                                {group.label}
                              </p>
                            )}
                            <div className="flex flex-wrap gap-2">
                              {group.pillars.map((sp) => (
                                <span
                                  key={sp}
                                  className="inline-flex items-center px-3.5 py-1.5 rounded-full border border-petrol/20 text-petrol/85 text-[13px]"
                                >
                                  {sp}
                                </span>
                              ))}
                            </div>
                          </div>
                        ))}
                      </div>
                    </div>
                  );
                })()}

                <div className="space-y-4">
                  <p className="label-upper text-petrol/55">{t('monitoredLabel')}</p>
                  <ul className="grid sm:grid-cols-2 gap-x-10 gap-y-3 max-w-3xl">
                    {pilar.monitoramos.map((item) => (
                      <li key={item} className="flex gap-3 text-petrol/80 text-base">
                        <span className="text-gold mt-1.5 leading-none">—</span>
                        <span className="leading-relaxed">{item}</span>
                      </li>
                    ))}
                  </ul>
                  <p className="text-petrol/50 text-sm pt-2 max-w-3xl">
                    {t('monitoredFootnote')}
                  </p>
                </div>
              </div>
            </article>
          ))}
        </div>
      </section>

      {/* Convicções */}
      <section className="bg-petrol text-cream">
        <div className="site-container section">
          <div className="max-w-2xl space-y-4 mb-16">
            <p className="label-upper text-gold">{t('convictionsLabel')}</p>
            <h2 className="heading-section text-cream text-3xl md:text-5xl">
              {t('convictionsTitle')}
            </h2>
          </div>

          <div className="border-t border-cream/15">
            {conviccoes.map((c) => (
              <div
                key={c.n}
                className="grid md:grid-cols-[80px_1fr] gap-6 md:gap-12 py-10 border-b border-cream/10"
              >
                <span className="label-upper text-gold pt-1">{c.n}</span>
                <div className="space-y-3 max-w-3xl">
                  <p className="heading-section text-cream text-2xl md:text-3xl leading-tight">
                    {c.text}
                  </p>
                  <p className="text-cream/70 leading-relaxed">{c.sub}</p>
                </div>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Obra de referência — livro Antes */}
      <BookReferenceBlock
        locale={locale}
        contextLabel={locale === 'en' ? 'Reference work' : 'Obra de referência'}
        contextLine={
          locale === 'en'
            ? 'The book that traces the silent decade where the ACTS Method intervenes — written by the clinical director.'
            : 'O livro que descreve a década silenciosa onde o Método AGIR atua — escrito pelo diretor clínico.'
        }
        tone="cream"
      />

      {/* Como aplicamos */}
      <section className="bg-cream">
        <div className="site-container section grid lg:grid-cols-2 gap-16 items-start">
          <div className="space-y-6">
            <p className="label-upper text-gold">{t('applyLabel')}</p>
            <h2 className="heading-section text-petrol text-3xl md:text-4xl">
              {t('applyTitle')}
            </h2>
            <p className="text-petrol/80 text-lg leading-relaxed">{t('applyP1')}</p>
            <p className="text-petrol/80 text-lg leading-relaxed">{t('applyP2')}</p>
          </div>

          <div className="space-y-4 lg:pt-12">
            <Link href="/escore-plenya" className="btn-outline-dark w-full sm:w-auto">
              {tCta('understandScore')}
            </Link>
            <Link href="/continuum" className="btn-gold w-full sm:w-auto sm:ml-3">
              {tCta('knowPlans')}
            </Link>
          </div>
        </div>
      </section>

      <RelatedBlogPosts limit={3} locale={isLocale(locale) ? locale : defaultLocale} />

      {/* Cross-links */}
      <section className="bg-paper">
        <div className="site-container section grid md:grid-cols-3 gap-8">
          <Link href="/dr-getulio" className="border-t border-petrol/15 pt-8 space-y-3 group">
            <p className="label-upper text-gold">{t('crossDirectorLabel')}</p>
            <p className="heading-section text-petrol text-xl group-hover:text-gold transition">
              {t('crossDirectorTitle')}
            </p>
          </Link>
          <Link href="/equipe" className="border-t border-petrol/15 pt-8 space-y-3 group">
            <p className="label-upper text-gold">{t('crossTeamLabel')}</p>
            <p className="heading-section text-petrol text-xl group-hover:text-gold transition">
              {t('crossTeamTitle')}
            </p>
          </Link>
          <Link href="/a-plenya" className="border-t border-petrol/15 pt-8 space-y-3 group">
            <p className="label-upper text-gold">{t('crossPlenyaLabel')}</p>
            <p className="heading-section text-petrol text-xl group-hover:text-gold transition">
              {t('crossPlenyaTitle')}
            </p>
          </Link>
        </div>
      </section>
    </>
  );
}

// Helper: agrupa as 4-5 chaves de monitoramento por pilar.
function getMonitored(
  t: (key: string) => string,
  n: number,
): string[] {
  const counts: Record<number, number> = { 1: 4, 2: 7, 3: 5, 4: 4 };
  const max = counts[n] ?? 4;
  return Array.from({ length: max }, (_, i) => t(`p${n}M${i + 1}`));
}
