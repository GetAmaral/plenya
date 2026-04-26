import { ChevronRight } from 'lucide-react';
import { Link } from '@/lib/i18n/navigation';

type Crumb = { label: string; href?: string };

export function Breadcrumbs({ items }: { items: Crumb[] }) {
  return (
    <nav aria-label="Breadcrumb" className="font-sans text-xs text-ink-muted tracking-wide">
      <ol className="flex flex-wrap items-center gap-2">
        {items.map((item, i) => {
          const isLast = i === items.length - 1;
          return (
            <li key={`${item.label}-${i}`} className="flex items-center gap-2">
              {item.href && !isLast ? (
                <Link href={item.href} className="hover:text-bordo transition-colors">
                  {item.label}
                </Link>
              ) : (
                <span className={isLast ? 'text-ink' : ''}>{item.label}</span>
              )}
              {!isLast && <ChevronRight size={12} className="text-ink-muted/50" />}
            </li>
          );
        })}
      </ol>
    </nav>
  );
}
