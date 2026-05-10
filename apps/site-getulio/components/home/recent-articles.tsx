import { getLocale, getTranslations } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { getPlenyaPostsByGetulio, pillarLabels } from '@/lib/plenya-blog';

export async function RecentArticles() {
  const locale = await getLocale();
  const t = await getTranslations('home.articles');
  const labels = pillarLabels(locale);
  const all = await getPlenyaPostsByGetulio();
  const articles = all.slice(0, 3);
  if (articles.length === 0) return null;

  const dateLocale = t('dateLocale');

  return (
    <section className="border-t border-rule">
      <div className="editorial-container py-20 md:py-28">
        <div className="flex items-baseline justify-between mb-12">
          <p className="label-meta">{t('sectionLabel')}</p>
          <Link href="/artigos" className="link-text font-sans text-sm">{t('viewAll')}</Link>
        </div>

        <ul className="divide-y divide-rule border-t border-rule">
          {articles.map((a) => (
            <li key={a.slug}>
              <Link
                href={{ pathname: '/artigos/[slug]', params: { slug: a.slug } }}
                className="grid md:grid-cols-[120px_1fr_180px] gap-6 py-6 group items-baseline"
              >
                <span className="font-sans text-sm text-ink-muted">{formatDate(a.date, dateLocale)}</span>
                <h3 className="font-serif text-xl text-ink group-hover:text-bordo transition-colors">
                  {a.title}
                </h3>
                <span className="label-meta text-bordo md:text-right">
                  {labels[a.pillar]}
                </span>
              </Link>
            </li>
          ))}
        </ul>
      </div>
    </section>
  );
}

function formatDate(d: string, locale: string) {
  const dt = new Date(d);
  return dt.toLocaleDateString(locale, { year: 'numeric', month: 'short' });
}
