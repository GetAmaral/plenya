import type { MetadataRoute } from 'next';
import { getAllPlenyaPostsFull } from '@/lib/plenya-blog';
import { getAllLectures } from '@/lib/lectures';

const BASE = 'https://drgetulioamaralfilho.com.br';
const LOCALES = ['pt', 'en'] as const;

const staticRoutes: { path: string; priority: number; changeFrequency: 'weekly' | 'monthly' | 'yearly' }[] = [
  { path: '', priority: 1.0, changeFrequency: 'weekly' },
  { path: '/sobre', priority: 0.9, changeFrequency: 'monthly' },
  { path: '/livro', priority: 0.9, changeFrequency: 'monthly' },
  { path: '/palestras', priority: 0.8, changeFrequency: 'monthly' },
  { path: '/escritos', priority: 0.8, changeFrequency: 'weekly' },
  { path: '/onde-atendo', priority: 0.7, changeFrequency: 'monthly' },
  { path: '/ensino', priority: 0.6, changeFrequency: 'monthly' },
  { path: '/contato', priority: 0.6, changeFrequency: 'yearly' },
];

function urlFor(locale: string, path: string): string {
  // PT é raiz (sem prefixo). EN ganha /en/. Path "" vira "/" no PT
  // e "/en" no EN.
  if (locale === 'pt') return `${BASE}${path || '/'}`;
  return `${BASE}/en${path}`;
}

export default async function sitemap(): Promise<MetadataRoute.Sitemap> {
  const now = new Date();

  const staticEntries: MetadataRoute.Sitemap = staticRoutes.flatMap((r) =>
    LOCALES.map((l) => ({
      url: urlFor(l, r.path),
      lastModified: now,
      changeFrequency: r.changeFrequency,
      priority: l === 'pt' ? r.priority : Math.max(0.3, r.priority - 0.1),
    })),
  );

  // Posts do blog Plenya espelhados em /escritos/{slug}.
  // Canonical aponta para plenyasaude.com.br/blog/{slug} (ou /en/blog/{slug}),
  // por isso prioridade baixa. Lista PT = lista canônica de slugs.
  const posts = await getAllPlenyaPostsFull('pt');
  const articleEntries: MetadataRoute.Sitemap = posts.flatMap((p) =>
    LOCALES.map((l) => ({
      url: urlFor(l, `/escritos/${p.slug}`),
      lastModified: new Date(p.updated ?? p.date),
      changeFrequency: 'monthly' as const,
      priority: 0.4,
    })),
  );

  const lectures = await getAllLectures();
  const lectureEntries: MetadataRoute.Sitemap = lectures.flatMap((l) =>
    LOCALES.map((loc) => ({
      url: urlFor(loc, `/palestras/${l.slug}`),
      lastModified: now,
      changeFrequency: 'monthly' as const,
      priority: 0.6,
    })),
  );

  return [...staticEntries, ...articleEntries, ...lectureEntries];
}
