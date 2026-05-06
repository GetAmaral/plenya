import { ImageResponse } from 'next/og';
import { brand } from '@plenya/brand';
import { getPost, getPillarLabels } from '@/lib/blog';
import { isLocale, defaultLocale, type Locale } from '@/lib/i18n/config';

export const size = { width: 1200, height: 630 };
export const contentType = 'image/png';

export default async function PostOg({ params }: { params: Promise<{ locale: string; slug: string }> }) {
  const { locale: rawLocale, slug } = await params;
  const locale: Locale = isLocale(rawLocale) ? rawLocale : defaultLocale;
  const post = await getPost(locale, slug);
  const pillarLabels = getPillarLabels(locale);
  const fallback = locale === 'en' ? 'Plenya Blog' : 'Blog Plenya';

  return new ImageResponse(
    (
      <div
        style={{
          width: '100%',
          height: '100%',
          display: 'flex',
          flexDirection: 'column',
          background: '#FAFAF7',
          color: '#063b4f',
          padding: '72px',
          justifyContent: 'space-between',
          fontFamily: 'serif',
        }}
      >
        <div
          style={{
            fontFamily: 'monospace',
            fontSize: 18,
            letterSpacing: 6,
            color: '#b38645',
            textTransform: 'uppercase',
          }}
        >
          {post ? pillarLabels[post.pillar] : fallback}
        </div>

        <div style={{ display: 'flex', flexDirection: 'column', gap: 20 }}>
          <div style={{ fontSize: 64, lineHeight: 1.1, fontWeight: 400, maxWidth: 1000 }}>
            {post?.title ?? 'Plenya'}
          </div>
        </div>

        <div
          style={{
            display: 'flex',
            justifyContent: 'space-between',
            alignItems: 'flex-end',
            fontFamily: 'monospace',
            fontSize: 16,
            color: '#063b4faa',
            textTransform: 'uppercase',
            letterSpacing: 4,
          }}
        >
          <span>{brand.domain}</span>
          <span>{brand.name.toUpperCase()}</span>
        </div>
      </div>
    ),
    { ...size },
  );
}
