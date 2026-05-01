import { Link } from '@/lib/i18n/navigation';
import { pillars, pillarLabels, type Pillar } from '@/lib/blog';

const base = 'label-upper px-5 py-2 border transition-colors duration-300';
const active = `${base} border-gold bg-gold text-petrol`;
const inactive = `${base} border-petrol/15 text-petrol/60 hover:border-petrol/40`;

export function PillarFilter({ active: current }: { active?: Pillar }) {
  return (
    <nav className="flex flex-wrap gap-2">
      <Link href="/blog" className={!current ? active : inactive}>
        Todos
      </Link>
      {pillars.map((p) => (
        <Link
          key={p}
          href={{ pathname: '/blog/categoria/[pilar]', params: { pilar: p } }}
          className={current === p ? active : inactive}
        >
          {pillarLabels[p]}
        </Link>
      ))}
    </nav>
  );
}
