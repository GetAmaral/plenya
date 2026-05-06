import { getTranslations } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { getAllPosts, type Pillar } from '@/lib/blog';
import { defaultLocale, type Locale } from '@/lib/i18n/config';

type Props = {
  /** Pillar(s) to filter by. If omitted, takes the most recent posts overall. */
  pillars?: Pillar[];
  /** Optional list of slugs to prioritise (in order). */
  preferred?: string[];
  limit?: number;
  title?: string;
  locale?: Locale;
};

export async function RelatedBlogPosts({
  pillars,
  preferred,
  limit = 3,
  title,
  locale = defaultLocale,
}: Props) {
  const t = await getTranslations({ locale, namespace: 'blogIndex' });
  const headingLabel = title ?? t('relatedReading');
  const all = await getAllPosts(locale);
  let pool = pillars?.length ? all.filter((p) => pillars.includes(p.pillar)) : all;

  if (preferred?.length) {
    const map = new Map(pool.map((p) => [p.slug, p]));
    const ordered = preferred.flatMap((s) => (map.get(s) ? [map.get(s)!] : []));
    const rest = pool.filter((p) => !preferred.includes(p.slug));
    pool = [...ordered, ...rest];
  }

  const posts = pool.slice(0, limit);
  if (!posts.length) return null;

  const dateTag = locale === 'en' ? 'en-US' : 'pt-BR';

  return (
    <section className="bg-paper">
      <div className="site-container section">
        <p className="label-upper text-gold mb-8">{headingLabel}</p>
        <ul className="grid md:grid-cols-3 gap-8">
          {posts.map((p) => (
            <li key={p.slug}>
              <Link
                href={{ pathname: '/blog/[slug]', params: { slug: p.slug } }}
                className="block group space-y-3"
              >
                <p className="label-upper text-petrol/55">
                  {new Date(p.date).toLocaleDateString(dateTag, { month: 'short', year: 'numeric' })}
                </p>
                <h3 className="font-serif text-lg text-petrol group-hover:text-gold transition-colors leading-snug">
                  {p.title}
                </h3>
                <p className="text-petrol/65 text-sm leading-relaxed line-clamp-3">{p.excerpt}</p>
              </Link>
            </li>
          ))}
        </ul>
      </div>
    </section>
  );
}
