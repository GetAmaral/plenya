import type { MetadataRoute } from 'next';
import { getAllPlenyaPostsFull } from '@/lib/plenya-blog';
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

  // Espelhamos os posts do blog Plenya em /escritos/{slug}, mas o canonical
  // aponta para plenyasaude.com.br/blog/{slug}. Por isso a prioridade aqui é
  // baixa — não queremos competir com a fonte canônica.
  const posts = await getAllPlenyaPostsFull();
  const articleEntries: MetadataRoute.Sitemap = posts.map((p) => ({
    url: `${BASE}/escritos/${p.slug}`,
    lastModified: new Date(p.updated ?? p.date),
    changeFrequency: 'monthly',
    priority: 0.4,
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
