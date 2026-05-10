import type { MetadataRoute } from 'next';
import { brand } from '@plenya/brand';
import { locales, type Locale } from '@/lib/i18n/config';
import { pathnames } from '@/lib/i18n/navigation';
import { getAllPosts, pillars } from '@/lib/blog';
import { getAllDoctors } from '@/lib/team';

// Resolve a chave PT do mapa `pathnames` para a URL localizada (com prefixo /en quando aplicável).
function resolvePath(key: keyof typeof pathnames, locale: Locale): string {
  const entry = pathnames[key];
  const path = typeof entry === 'string' ? entry : entry[locale];
  const prefix = locale === 'pt' ? '' : `/${locale}`;
  // Em "/" o prefix sozinho não pode terminar com / quando locale=pt → sempre OK pois retornamos brand.url + ''.
  return path === '/' ? `${brand.url}${prefix || ''}` : `${brand.url}${prefix}${path}`;
}

function localizedAlternates(key: keyof typeof pathnames) {
  const out: Record<string, string> = {};
  for (const l of locales) {
    if (l === 'pt') {
      const ptUrl = resolvePath(key, l);
      out['pt-BR'] = ptUrl;
      out['pt'] = ptUrl; // genérico — cobre Portugal e demais lusófonos
    } else {
      out[l] = resolvePath(key, l);
    }
  }
  out['x-default'] = resolvePath(key, 'pt');
  return out;
}

const staticRouteKeys: (keyof typeof pathnames)[] = [
  '/',
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
  '/diagnostico',
  '/casos',
  '/como-funciona',
  '/boletim',
  '/privacidade',
  '/termos',
];

export default async function sitemap(): Promise<MetadataRoute.Sitemap> {
  const now = new Date();

  const staticEntries = staticRouteKeys.flatMap((key) =>
    locales.map((locale) => ({
      url: resolvePath(key, locale),
      lastModified: now,
      changeFrequency: 'weekly' as const,
      priority: key === '/' ? 1 : 0.7,
      alternates: { languages: localizedAlternates(key) },
    })),
  );

  // Categorias de blog (rota dinâmica por pilar; resolvedPath não suporta
  // params — montamos manualmente a partir do mapa).
  const pillarEntries = pillars.flatMap((pilar) =>
    locales.map((locale) => {
      const ptPath = `/blog/categoria/${pilar}`;
      const enPath = `/en/blog/category/${pilar}`;
      const url = locale === 'pt' ? `${brand.url}${ptPath}` : `${brand.url}${enPath}`;
      return {
        url,
        lastModified: now,
        changeFrequency: 'weekly' as const,
        priority: 0.5,
        alternates: {
          languages: {
            'pt-BR': `${brand.url}${ptPath}`,
            en: `${brand.url}${enPath}`,
            'x-default': `${brand.url}${ptPath}`,
          },
        },
      };
    }),
  );

  const postEntries: MetadataRoute.Sitemap = [];
  for (const locale of locales) {
    const posts = await getAllPosts(locale);
    for (const post of posts) {
      const ptUrl = `${brand.url}/blog/${post.slug}`;
      const enUrl = `${brand.url}/en/blog/${post.slug}`;
      postEntries.push({
        url: locale === 'pt' ? ptUrl : enUrl,
        lastModified: new Date(post.updated ?? post.date),
        changeFrequency: 'monthly' as const,
        priority: 0.8,
        alternates: {
          languages: {
            'pt-BR': ptUrl,
            en: enUrl,
            'x-default': ptUrl,
          },
        },
      });
    }
  }

  const doctors = await getAllDoctors();
  // Dr. Getúlio tem URL canônica em /dr-getulio (já listada em staticRoutes).
  const doctorEntries = doctors
    .filter((doc) => doc.slug !== 'getulio-amaral')
    .flatMap((doc) =>
      locales.map((locale) => {
        const ptUrl = `${brand.url}/equipe/${doc.slug}`;
        const enUrl = `${brand.url}/en/team/${doc.slug}`;
        return {
          url: locale === 'pt' ? ptUrl : enUrl,
          lastModified: now,
          changeFrequency: 'monthly' as const,
          priority: doc.featured ? 0.9 : 0.6,
          alternates: {
            languages: {
              'pt-BR': ptUrl,
              en: enUrl,
              'x-default': ptUrl,
            },
          },
        };
      }),
    );

  return [...staticEntries, ...pillarEntries, ...postEntries, ...doctorEntries];
}
