import type { MetadataRoute } from 'next';
import { getAllArticles } from '@/lib/articles';
import { getAllLectures } from '@/lib/lectures';

const BASE = 'https://drgetulioamaralfilho.com.br';

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

export default async function sitemap(): Promise<MetadataRoute.Sitemap> {
  const now = new Date();

  const staticEntries: MetadataRoute.Sitemap = staticRoutes.map((r) => ({
    url: `${BASE}${r.path}`,
    lastModified: now,
    changeFrequency: r.changeFrequency,
    priority: r.priority,
  }));

  const articles = await getAllArticles();
  const articleEntries: MetadataRoute.Sitemap = articles.map((a) => ({
    url: `${BASE}/escritos/${a.slug}`,
    lastModified: new Date(a.date),
    changeFrequency: 'monthly',
    priority: 0.7,
  }));

  const lectures = await getAllLectures();
  const lectureEntries: MetadataRoute.Sitemap = lectures.map((l) => ({
    url: `${BASE}/palestras/${l.slug}`,
    lastModified: now,
    changeFrequency: 'monthly',
    priority: 0.6,
  }));

  return [...staticEntries, ...articleEntries, ...lectureEntries];
}
