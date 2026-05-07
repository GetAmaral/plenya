import { getTranslations } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';

export type HowStep = {
  n: string;
  title: string;
  body: string;
  detail?: string;
};

export async function HowItWorksSteps({
  steps,
  label,
  title,
  bg = 'bg-cream',
}: {
  steps?: HowStep[];
  label?: string;
  title?: string;
  bg?: string;
}) {
  const t = await getTranslations('howItWorks');
  const resolvedSteps: HowStep[] =
    steps ?? [
      { n: '01', title: t('step1Title'), body: t('step1Body'), detail: t('step1Detail') },
      { n: '02', title: t('step2Title'), body: t('step2Body'), detail: t('step2Detail') },
      { n: '03', title: t('step3Title'), body: t('step3Body'), detail: t('step3Detail') },
      { n: '04', title: t('step4Title'), body: t('step4Body'), detail: t('step4Detail') },
    ];
  const resolvedLabel = label ?? t('stepsLabel');
  const resolvedTitle = title ?? t('stepsTitle');

  return (
    <section className={bg}>
      <div className="site-container section">
        <div className="max-w-3xl mb-14 space-y-4">
          <p className="label-upper text-gold">{resolvedLabel}</p>
          <h2 className="heading-section text-petrol text-3xl md:text-5xl">{resolvedTitle}</h2>
        </div>

        <ol className="grid md:grid-cols-2 lg:grid-cols-4 gap-px bg-petrol/15 border-y border-petrol/15">
          {resolvedSteps.map((s) => (
            <li key={s.n} className="bg-paper p-8 md:p-10 space-y-5">
              <p className="text-gold text-5xl md:text-6xl font-light leading-none tabular-nums">
                {s.n}
              </p>
              <h3 className="heading-section text-petrol text-2xl">{s.title}</h3>
              <p className="text-petrol/75 leading-relaxed">{s.body}</p>
              {s.detail && (
                <p className="label-upper text-petrol/50 text-[10px] tracking-[0.2em] pt-3 border-t border-petrol/10">
                  {s.detail}
                </p>
              )}
            </li>
          ))}
        </ol>

        <div className="mt-12 flex flex-wrap gap-4">
          <Link href="/escore-plenya/avaliar" className="btn-gold">
            {t('stepsCtaTriage')}
          </Link>
          <Link href="/contato" className="btn-outline-dark">
            {t('stepsCtaContact')}
          </Link>
        </div>
      </div>
    </section>
  );
}
