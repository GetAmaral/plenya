import { useTranslations } from 'next-intl';
import { ChevronRight } from 'lucide-react';
import { Link } from '@/lib/i18n/navigation';

type Crumb = { label: string; href?: string };

export function Breadcrumbs({ items }: { items: Crumb[] }) {
  const t = useTranslations('common');
  return (
    <nav aria-label={t('breadcrumbAriaLabel')} className="font-sans text-sm text-ink-soft tracking-wide">
      <ol className="flex flex-wrap items-center gap-2.5">
        {items.map((item, i) => {
          const isLast = i === items.length - 1;
          return (
            <li key={`${item.label}-${i}`} className="flex items-center gap-2.5">
              {item.href && !isLast ? (
                // eslint-disable-next-line @typescript-eslint/no-explicit-any
                <Link href={item.href as any} className="hover:text-bordo transition-colors">
                  {item.label}
                </Link>
              ) : (
                <span className={isLast ? 'text-ink font-medium' : ''}>{item.label}</span>
              )}
              {!isLast && <ChevronRight size={14} className="text-bordo/60" />}
            </li>
          );
        })}
      </ol>
    </nav>
  );
}
