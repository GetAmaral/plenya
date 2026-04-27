import type { MetadataRoute } from 'next';

const BASE = 'https://drgetulioamaralfilho.com.br';

export default function robots(): MetadataRoute.Robots {
  const isStaging =
    process.env.NEXT_PUBLIC_SITE_URL?.includes('staging') ||
    process.env.VERCEL_ENV === 'preview';
  if (isStaging) {
    return { rules: [{ userAgent: '*', disallow: '/' }] };
  }
  return {
    rules: [{ userAgent: '*', allow: '/', disallow: ['/api/'] }],
    sitemap: `${BASE}/sitemap.xml`,
    host: BASE,
  };
}
