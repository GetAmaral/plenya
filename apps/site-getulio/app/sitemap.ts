import type { MetadataRoute } from 'next';
import { getAllPlenyaPostsFull } from '@/lib/plenya-blog';
import { getAllLectures } from '@/lib/lectures';
import { getAllBooks } from '@/lib/books';

const BASE = 'https://drgetulioamaralfilho.com.br';

// URL paths por locale (espelha o pathnames mapping em lib/i18n/config.ts).
const PATHS_PT: Record<string, string> = {
  home: '',
  oMedico: '/o-medico',
  livros: '/livros',
  palestras: '/palestras',
  artigos: '/artigos',
  ondeAtendo: '/onde-atendo',
  ensino: '/ensino',
  contato: '/contato',
};

const PATHS_EN: Record<string, string> = {
  home: '',
  oMedico: '/the-physician',
  livros: '/books',
  palestras: '/lectures',
  artigos: '/articles',
  ondeAtendo: '/where-i-practice',
  ensino: '/teaching',
  contato: '/contact',
};

const STATIC_KEYS: { key: keyof typeof PATHS_PT; priority: number; changeFrequency: 'weekly' | 'monthly' | 'yearly' }[] = [
  { key: 'home', priority: 1.0, changeFrequency: 'weekly' },
  { key: 'oMedico', priority: 0.9, changeFrequency: 'monthly' },
  { key: 'livros', priority: 0.9, changeFrequency: 'monthly' },
  { key: 'palestras', priority: 0.8, changeFrequency: 'monthly' },
  { key: 'artigos', priority: 0.8, changeFrequency: 'weekly' },
  { key: 'ondeAtendo', priority: 0.7, changeFrequency: 'monthly' },
  { key: 'ensino', priority: 0.6, changeFrequency: 'monthly' },
  { key: 'contato', priority: 0.6, changeFrequency: 'yearly' },
];

function ptUrl(path: string): string {
  return `${BASE}${path || '/'}`;
}
function enUrl(path: string): string {
  return `${BASE}/en${path}`;
}

function alternatesPair(pathPt: string, pathEn: string): { languages: Record<string, string> } {
  const pt = ptUrl(pathPt);
  return {
    languages: {
      'pt-BR': pt,
      pt, // genérico — cobre Portugal e demais lusófonos
      en: enUrl(pathEn),
      'x-default': pt,
    },
  };
}

export default async function sitemap(): Promise<MetadataRoute.Sitemap> {
  const now = new Date();

  const staticEntries: MetadataRoute.Sitemap = STATIC_KEYS.map((r) => {
    const pathPt = PATHS_PT[r.key];
    const pathEn = PATHS_EN[r.key];
    return {
      url: ptUrl(pathPt),
      lastModified: now,
      changeFrequency: r.changeFrequency,
      priority: r.priority,
      alternates: alternatesPair(pathPt, pathEn),
    };
  });

  const posts = await getAllPlenyaPostsFull('pt');
  const articleEntries: MetadataRoute.Sitemap = posts.map((p) => {
    const pathPt = `/artigos/${p.slug}`;
    const pathEn = `/articles/${p.slug}`;
    return {
      url: ptUrl(pathPt),
      lastModified: new Date(p.updated ?? p.date),
      changeFrequency: 'monthly' as const,
      priority: 0.5,
      alternates: alternatesPair(pathPt, pathEn),
      images: p.cover ? [`${BASE}${p.cover.startsWith('/') ? p.cover : `/${p.cover}`}`] : undefined,
    };
  });

  const lectures = await getAllLectures();
  const lectureEntries: MetadataRoute.Sitemap = lectures.map((l) => {
    const pathPt = `/palestras/${l.slug}`;
    const pathEn = `/lectures/${l.slug}`;
    return {
      url: ptUrl(pathPt),
      lastModified: now,
      changeFrequency: 'monthly' as const,
      priority: 0.6,
      alternates: alternatesPair(pathPt, pathEn),
    };
  });

  const books = await getAllBooks();
  const bookEntries: MetadataRoute.Sitemap = books.flatMap((b) => {
    const detailPt = `/livros/${b.slug}`;
    const detailEn = `/books/${b.slug}`;
    const excertosPt = `/livros/${b.slug}/excertos`;
    const excertosEn = `/books/${b.slug}/excerpts`;
    return [
      {
        url: ptUrl(detailPt),
        lastModified: now,
        changeFrequency: 'monthly' as const,
        priority: 0.85,
        alternates: alternatesPair(detailPt, detailEn),
        images: b.cover ? [`${BASE}${b.cover}`] : undefined,
      },
      {
        url: ptUrl(excertosPt),
        lastModified: now,
        changeFrequency: 'monthly' as const,
        priority: 0.55,
        alternates: alternatesPair(excertosPt, excertosEn),
      },
    ];
  });

  return [...staticEntries, ...articleEntries, ...lectureEntries, ...bookEntries];
}
