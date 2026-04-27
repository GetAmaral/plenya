import type { MetadataRoute } from 'next';
import { brand } from '@plenya/brand';
import { locales } from '@/lib/i18n/config';
import { getAllPosts, pillars } from '@/lib/blog';
import { getAllDoctors } from '@/lib/team';

const staticRoutes = [
  '',
  '/dr-getulio',
  '/a-plenya',
  '/metodo-agir',
  '/equipe',
  '/consultas',
  '/continuum',
  '/blog',
  '/contato',
  '/escore-plenya',
  '/depoimentos',
  '/healthspan',
  '/checkup-longevidade',
  '/avaliacao-renal-preventiva',
  '/medicina-funcional-integrativa',
  '/privacidade',
  '/termos',
];

function localized(path: string, locale: string) {
  return `${brand.url}${locale === 'pt' ? '' : `/${locale}`}${path}`;
}

export default async function sitemap(): Promise<MetadataRoute.Sitemap> {
  const now = new Date();

  const staticEntries = staticRoutes.flatMap((route) =>
    locales.map((locale) => ({
      url: localized(route, locale),
      lastModified: now,
      changeFrequency: 'weekly' as const,
      priority: route === '' ? 1 : 0.7,
      alternates: {
        languages: Object.fromEntries(
          locales.map((l) => [l === 'pt' ? 'pt-BR' : l, localized(route, l)]),
        ),
      },
    })),
  );

  const pillarEntries = pillars.flatMap((pilar) =>
    locales.map((locale) => ({
      url: localized(`/blog/categoria/${pilar}`, locale),
      lastModified: now,
      changeFrequency: 'weekly' as const,
      priority: 0.5,
    })),
  );

  const postEntries: MetadataRoute.Sitemap = [];
  for (const locale of locales) {
    const posts = await getAllPosts(locale);
    for (const post of posts) {
      postEntries.push({
        url: localized(`/blog/${post.slug}`, locale),
        lastModified: new Date(post.updated ?? post.date),
        changeFrequency: 'monthly' as const,
        priority: 0.8,
      });
    }
  }

  const doctors = await getAllDoctors();
  const doctorEntries = doctors.flatMap((doc) =>
    locales.map((locale) => ({
      url: localized(`/equipe/${doc.slug}`, locale),
      lastModified: now,
      changeFrequency: 'monthly' as const,
      priority: doc.featured ? 0.9 : 0.6,
    })),
  );

  return [...staticEntries, ...pillarEntries, ...postEntries, ...doctorEntries];
}
