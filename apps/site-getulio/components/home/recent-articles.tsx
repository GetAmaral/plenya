import { Link } from '@/lib/i18n/navigation';
import { getRecentArticles } from '@/lib/articles';

export async function RecentArticles() {
  const articles = await getRecentArticles(3);
  if (articles.length === 0) return null;

  return (
    <section className="border-t border-rule">
      <div className="editorial-container py-20 md:py-28">
        <div className="flex items-baseline justify-between mb-12">
          <p className="label-meta">Escritos recentes</p>
          <Link href="/escritos" className="link-text font-sans text-sm">Ver todos →</Link>
        </div>

        <ul className="divide-y divide-rule border-t border-rule">
          {articles.map((a) => (
            <li key={a.slug}>
              <Link
                href={`/escritos/${a.slug}`}
                className="grid md:grid-cols-[120px_1fr_180px] gap-6 py-6 group items-baseline"
              >
                <span className="font-sans text-sm text-ink-muted">{formatDate(a.date)}</span>
                <h3 className="font-serif text-xl text-ink group-hover:text-bordo transition-colors">
                  {a.title}
                </h3>
                <span className="label-meta text-bordo md:text-right">{a.tag}</span>
              </Link>
            </li>
          ))}
        </ul>
      </div>
    </section>
  );
}

function formatDate(d: string) {
  const dt = new Date(d);
  return dt.toLocaleDateString('pt-BR', { year: 'numeric', month: 'short' });
}
