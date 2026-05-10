import type { Metadata } from 'next';
import { setRequestLocale, getTranslations } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';

type Params = Promise<{ locale: string }>;

export async function generateMetadata({ params }: { params: Params }): Promise<Metadata> {
  const { locale } = await params;
  const t = await getTranslations({ locale, namespace: 'cases' });
  return {
    title: t('metaTitle'),
    description: t('metaDescription'),
    alternates: {
      canonical: locale === 'en' ? '/en/cases' : '/casos',
      languages: {
        'pt-BR': '/casos', pt: '/casos',
        en: '/en/cases',
        'x-default': '/casos',
      },
    },
  };
}

export default async function CasosPage({ params }: { params: Params }) {
  const { locale } = await params;
  setRequestLocale(locale);
  const t = await getTranslations('cases');

  type Pillar = 'A' | 'G' | 'I' | 'R';
  const pillarLabels: Record<Pillar, string> = {
    A: t('pillarA'),
    G: t('pillarG'),
    I: t('pillarI'),
    R: t('pillarR'),
  };

  const buildCase = (n: 1 | 2 | 3, achadosCount: number, pillar: Pillar) => ({
    slug: `case-${n}`,
    titulo: t(`c${n}Title` as 'c1Title'),
    resumo: t(`c${n}Resumo` as 'c1Resumo'),
    paciente: t(`c${n}Paciente` as 'c1Paciente'),
    contexto: t(`c${n}Contexto` as 'c1Contexto'),
    achados: Array.from({ length: achadosCount }, (_, i) => t(`c${n}A${i + 1}` as 'c1A1')),
    conduta: [1, 2, 3, 4].map((i) => t(`c${n}Co${i}` as 'c1Co1')),
    evolucao: [1, 2, 3, 4].map((i) => t(`c${n}E${i}` as 'c1E1')),
    tempo: t(`c${n}Tempo` as 'c1Tempo'),
    pillar,
  });

  const casos = [buildCase(1, 4, 'G'), buildCase(2, 5, 'A'), buildCase(3, 4, 'I')];

  return (
    <>
      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-24 md:pt-40 md:pb-32 max-w-3xl">
          <p className="label-upper text-gold mb-6">{t('heroLabel')}</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,4.5rem)] text-cream">
            {t('heroTitle')}
          </h1>
          <p className="text-cream/75 text-lg leading-relaxed mt-8 max-w-2xl">{t('heroP')}</p>
        </div>
      </section>

      <section className="bg-cream">
        <div className="site-container section space-y-20">
          {casos.map((c, i) => (
            <article
              key={c.slug}
              className="grid lg:grid-cols-12 gap-10 border-t border-petrol/15 pt-10"
            >
              <header className="lg:col-span-4 space-y-4">
                <p className="label-upper text-gold">
                  {t('caseLabel')} 0{i + 1} · {pillarLabels[c.pillar]}
                </p>
                <h2 className="heading-section text-petrol text-2xl md:text-3xl leading-tight">
                  {c.titulo}
                </h2>
                <p className="text-petrol/70 leading-relaxed">{c.resumo}</p>
                <p className="label-upper text-petrol/55 pt-2 border-t border-petrol/10">
                  {c.paciente}
                </p>
                <p className="label-upper text-gold">{c.tempo}</p>
              </header>

              <div className="lg:col-span-8 grid md:grid-cols-3 gap-8">
                <div className="space-y-3">
                  <p className="label-upper text-petrol/55">{t('contextLabel')}</p>
                  <p className="text-petrol/80 leading-relaxed text-sm">{c.contexto}</p>
                </div>
                <div className="space-y-3">
                  <p className="label-upper text-petrol/55">{t('findingsLabel')}</p>
                  <ul className="space-y-2">
                    {c.achados.map((a) => (
                      <li key={a} className="flex gap-2 text-petrol/80 text-sm leading-relaxed">
                        <span className="text-gold mt-1.5 leading-none">—</span>
                        <span>{a}</span>
                      </li>
                    ))}
                  </ul>
                </div>
                <div className="space-y-3">
                  <p className="label-upper text-petrol/55">{t('conductLabel')}</p>
                  <ul className="space-y-2">
                    {c.conduta.map((a) => (
                      <li key={a} className="flex gap-2 text-petrol/80 text-sm leading-relaxed">
                        <span className="text-gold mt-1.5 leading-none">—</span>
                        <span>{a}</span>
                      </li>
                    ))}
                  </ul>
                </div>

                <div className="md:col-span-3 bg-paper p-6 md:p-8 mt-2 space-y-3">
                  <p className="label-upper text-gold">{t('evolutionLabel')}</p>
                  <ul className="grid sm:grid-cols-2 gap-x-8 gap-y-2">
                    {c.evolucao.map((a) => (
                      <li key={a} className="flex gap-2 text-petrol/85 text-sm leading-relaxed">
                        <span className="text-gold mt-1.5 leading-none">→</span>
                        <span>{a}</span>
                      </li>
                    ))}
                  </ul>
                </div>
              </div>
            </article>
          ))}
        </div>
      </section>

      <section className="bg-paper">
        <div className="site-container section max-w-3xl space-y-4">
          <p className="label-upper text-petrol/55">{t('noteLabel')}</p>
          <p className="text-petrol/75 leading-relaxed text-sm">{t('noteText')}</p>
        </div>
      </section>

      <section className="bg-petrol text-cream">
        <div className="site-container section text-center space-y-6 max-w-2xl mx-auto">
          <p className="label-upper text-gold">{t('ctaLabel')}</p>
          <p className="heading-section text-cream text-2xl md:text-3xl">{t('ctaTitle')}</p>
          <div className="flex flex-wrap justify-center gap-4 pt-2">
            <Link href="/escore-plenya/avaliar" className="btn-gold">
              {t('ctaButton1')}
            </Link>
            <Link href="/contato" className="btn-outline-dark border-cream/40 text-cream">
              {t('ctaButton2')}
            </Link>
          </div>
        </div>
      </section>
    </>
  );
}
