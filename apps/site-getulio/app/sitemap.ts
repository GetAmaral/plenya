import type { MetadataRoute } from 'next';
import { getAllPlenyaPostsFull } from '@/lib/plenya-blog';
import { getAllLectures } from '@/lib/lectures';
import { getAllBooks } from '@/lib/books';

const BASE = 'https://drgetulioamaralfilho.com.br';

const staticRoutes: { path: string; priority: number; changeFrequency: 'weekly' | 'monthly' | 'yearly' }[] = [
  { path: '', priority: 1.0, changeFrequency: 'weekly' },
  { path: '/sobre', priority: 0.9, changeFrequency: 'monthly' },
  { path: '/livros', priority: 0.9, changeFrequency: 'monthly' },
  { path: '/palestras', priority: 0.8, changeFrequency: 'monthly' },
  { path: '/escritos', priority: 0.8, changeFrequency: 'weekly' },
  { path: '/onde-atendo', priority: 0.7, changeFrequency: 'monthly' },
  { path: '/ensino', priority: 0.6, changeFrequency: 'monthly' },
  { path: '/contato', priority: 0.6, changeFrequency: 'yearly' },
];

function ptUrl(path: string): string {
  return `${BASE}${path || '/'}`;
}
function enUrl(path: string): string {
  return `${BASE}/en${path}`;
}

function alternates(path: string): { languages: Record<string, string> } {
  return {
    languages: {
      'pt-BR': ptUrl(path),
      en: enUrl(path),
      'x-default': ptUrl(path),
    },
  };
}

export default async function sitemap(): Promise<MetadataRoute.Sitemap> {
  const now = new Date();

  const staticEntries: MetadataRoute.Sitemap = staticRoutes.map((r) => ({
    url: ptUrl(r.path),
    lastModified: now,
    changeFrequency: r.changeFrequency,
    priority: r.priority,
    alternates: alternates(r.path),
  }));

  // Posts mirrored from Plenya. Canonical points to Plenya, but
  // keep them in sitemap so Google still discovers them via this site.
  const posts = await getAllPlenyaPostsFull('pt');
  const articleEntries: MetadataRoute.Sitemap = posts.map((p) => {
    const path = `/escritos/${p.slug}`;
    return {
      url: ptUrl(path),
      lastModified: new Date(p.updated ?? p.date),
      changeFrequency: 'monthly' as const,
      priority: 0.5,
      alternates: alternates(path),
      images: p.cover ? [`${BASE}${p.cover.startsWith('/') ? p.cover : `/${p.cover}`}`] : undefined,
    };
  });

  const lectures = await getAllLectures();
  const lectureEntries: MetadataRoute.Sitemap = lectures.map((l) => {
    const path = `/palestras/${l.slug}`;
    return {
      url: ptUrl(path),
      lastModified: now,
      changeFrequency: 'monthly' as const,
      priority: 0.6,
      alternates: alternates(path),
    };
  });

  const books = await getAllBooks();
  const bookEntries: MetadataRoute.Sitemap = books.flatMap((b) => {
    const detailPath = `/livros/${b.slug}`;
    const excertosPath = `/livros/${b.slug}/excertos`;
    return [
      {
        url: ptUrl(detailPath),
        lastModified: now,
        changeFrequency: 'monthly' as const,
        priority: 0.85,
        alternates: alternates(detailPath),
        images: b.cover ? [`${BASE}${b.cover}`] : undefined,
      },
      {
        url: ptUrl(excertosPath),
        lastModified: now,
        changeFrequency: 'monthly' as const,
        priority: 0.55,
        alternates: alternates(excertosPath),
      },
    ];
  });

  return [...staticEntries, ...articleEntries, ...lectureEntries, ...bookEntries];
}
