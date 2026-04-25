import { ChevronRight } from 'lucide-react';
import { Link } from '@/lib/i18n/navigation';

export type Crumb = { label: string; href?: string };

export function Breadcrumbs({ items, dark = false }: { items: Crumb[]; dark?: boolean }) {
  const baseColor = dark ? 'text-cream/60' : 'text-petrol/55';
  const linkHover = dark ? 'hover:text-cream' : 'hover:text-petrol';
  const sep = dark ? 'text-cream/30' : 'text-petrol/30';

  return (
    <nav aria-label="Breadcrumb" className={`label-upper text-xs ${baseColor}`}>
      <ol className="flex flex-wrap items-center gap-2">
        {items.map((it, i) => {
          const last = i === items.length - 1;
          return (
            <li key={`${it.label}-${i}`} className="flex items-center gap-2">
              {it.href && !last ? (
                <Link href={it.href} className={`underline-offset-4 hover:underline transition ${linkHover}`}>
                  {it.label}
                </Link>
              ) : (
                <span aria-current={last ? 'page' : undefined}>{it.label}</span>
              )}
              {!last && <ChevronRight size={12} className={sep} aria-hidden />}
            </li>
          );
        })}
      </ol>
    </nav>
  );
}
