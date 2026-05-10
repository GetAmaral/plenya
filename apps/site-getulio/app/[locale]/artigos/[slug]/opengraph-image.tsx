import { ImageResponse } from 'next/og';
import { getPlenyaPost, pillarLabels } from '@/lib/blog';

export const runtime = 'nodejs';
export const alt = 'Dr. Getúlio Amaral Filho — Artigo';
export const size = { width: 1200, height: 630 };
export const contentType = 'image/png';

export default async function Image({
  params,
}: {
  params: { locale: string; slug: string };
}) {
  const { locale, slug } = params;
  const post = await getPlenyaPost(slug, locale);
  const labels = pillarLabels(locale);

  const title = post?.title ?? 'Dr. Getúlio Amaral Filho';
  const pillar = post ? labels[post.pillar] : 'Medicina guiada por raciocínio clínico';
  const author = locale === 'en' ? 'Dr. Getúlio Amaral Filho · Nephrologist' : 'Dr. Getúlio Amaral Filho · Nefrologista';

  return new ImageResponse(
    (
      <div
        style={{
          width: '100%',
          height: '100%',
          display: 'flex',
          flexDirection: 'column',
          backgroundColor: '#0E2A2E',
          color: '#F5EFE6',
          padding: '72px 88px',
          fontFamily: 'Georgia, serif',
        }}
      >
        <div
          style={{
            fontSize: 22,
            letterSpacing: 6,
            textTransform: 'uppercase',
            color: '#C8A96A',
            fontFamily: 'system-ui, sans-serif',
          }}
        >
          {pillar}
        </div>

        <div
          style={{
            display: 'flex',
            fontSize: 64,
            lineHeight: 1.12,
            marginTop: 32,
            color: '#F5EFE6',
            fontWeight: 500,
            maxWidth: 1024,
          }}
        >
          {title}
        </div>

        <div style={{ flex: 1 }} />

        <div
          style={{
            display: 'flex',
            justifyContent: 'space-between',
            alignItems: 'flex-end',
            borderTop: '1px solid rgba(245, 239, 230, 0.25)',
            paddingTop: 28,
          }}
        >
          <div
            style={{
              fontSize: 22,
              fontFamily: 'system-ui, sans-serif',
              color: '#F5EFE6',
              opacity: 0.9,
            }}
          >
            {author}
          </div>
          <div
            style={{
              fontSize: 18,
              fontFamily: 'system-ui, sans-serif',
              color: '#C8A96A',
              letterSpacing: 4,
              textTransform: 'uppercase',
            }}
          >
            drgetulioamaralfilho.com.br
          </div>
        </div>
      </div>
    ),
    { ...size },
  );
}
