import { ImageResponse } from 'next/og';
import { brand } from '@plenya/brand';

export const size = { width: 1200, height: 630 };
export const contentType = 'image/png';

export default async function OpengraphImage({ params }: { params: Promise<{ locale: string }> }) {
  const { locale } = await params;
  const claim =
    locale === 'en' ? 'Live well. Live more.' : locale === 'es' ? 'Vive bien. Vive más.' : 'Viva bem. Viva mais.';
  const tag =
    locale === 'en'
      ? 'Health · Performance · Longevity'
      : locale === 'es'
        ? 'Salud · Rendimiento · Longevidad'
        : 'Saúde · Performance · Longevidade';

  return new ImageResponse(
    (
      <div
        style={{
          width: '100%',
          height: '100%',
          display: 'flex',
          flexDirection: 'column',
          background: '#0E3B4F',
          color: '#E6E1D8',
          padding: '80px',
          justifyContent: 'space-between',
          fontFamily: 'serif',
        }}
      >
        <div
          style={{
            fontFamily: 'monospace',
            fontSize: 22,
            letterSpacing: 8,
            color: '#A08456',
            textTransform: 'uppercase',
          }}
        >
          {brand.name.toUpperCase()}
        </div>

        <div style={{ display: 'flex', flexDirection: 'column', gap: 24 }}>
          <div style={{ fontSize: 84, lineHeight: 1.05, fontWeight: 300, maxWidth: 900 }}>{claim}</div>
          <div
            style={{
              fontFamily: 'monospace',
              fontSize: 22,
              letterSpacing: 6,
              color: '#A08456',
              textTransform: 'uppercase',
            }}
          >
            {tag}
          </div>
        </div>

        <div
          style={{
            display: 'flex',
            justifyContent: 'space-between',
            alignItems: 'flex-end',
            fontFamily: 'monospace',
            fontSize: 18,
            color: '#E6E1D8aa',
          }}
        >
          <span>{brand.domain}</span>
          <span>{brand.legalName}</span>
        </div>
      </div>
    ),
    { ...size },
  );
}
