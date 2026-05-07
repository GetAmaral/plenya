import { useTranslations } from 'next-intl';
import { Link } from '@/lib/i18n/navigation';

const pillarMeta = [
  { keyPrefix: 'doctor', href: '/onde-atendo' },
  { keyPrefix: 'professor', href: '/ensino' },
  { keyPrefix: 'author', href: '/livro' },
  { keyPrefix: 'speaker', href: '/palestras' },
] as const;

export function Pillars() {
  const t = useTranslations('home.pillars');
  return (
    <section className="border-t border-rule">
      <div className="editorial-container py-20 md:py-28">
        <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-10 md:gap-14">
          {pillarMeta.map((p) => (
            <Link
              key={p.keyPrefix}
              href={p.href}
              className="group block border-l-2 border-bordo/30 pl-6 hover:border-bordo transition-colors"
            >
              <p className="label-meta mb-4">{t(`${p.keyPrefix}Label`)}</p>
              <h3 className="heading-section text-2xl mb-3 group-hover:text-bordo transition-colors">
                {t(`${p.keyPrefix}Title`)}
              </h3>
              <p className="font-serif text-ink-soft text-base leading-relaxed">
                {t(`${p.keyPrefix}Body`)}
              </p>
            </Link>
          ))}
        </div>
      </div>
    </section>
  );
}
