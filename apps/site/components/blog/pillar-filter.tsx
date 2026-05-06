import { getTranslations } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { pillars, getPillarLabels, type Pillar } from '@/lib/blog';
import { defaultLocale, type Locale } from '@/lib/i18n/config';

const base = 'label-upper px-5 py-2 border transition-colors duration-300';
const active = `${base} border-gold bg-gold text-petrol`;
const inactive = `${base} border-petrol/15 text-petrol/60 hover:border-petrol/40`;

export async function PillarFilter({ active: current, locale = defaultLocale }: { active?: Pillar; locale?: Locale }) {
  const labels = getPillarLabels(locale);
  const t = await getTranslations({ locale, namespace: 'blogIndex' });
  return (
    <nav className="flex flex-wrap gap-2">
      <Link href="/blog" className={!current ? active : inactive}>
        {t('filterAll')}
      </Link>
      {pillars.map((p) => (
        <Link
          key={p}
          href={{ pathname: '/blog/categoria/[pilar]', params: { pilar: p } }}
          className={current === p ? active : inactive}
        >
          {labels[p]}
        </Link>
      ))}
    </nav>
  );
}
