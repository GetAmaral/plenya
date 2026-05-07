type LectureSchemaProps = {
  title: string;
  description: string;
  slug: string;
  duration: string;
  audience: string[];
  locale?: string;
};

const BASE = 'https://drgetulioamaralfilho.com.br';

export function LectureSchema({ title, description, slug, duration, audience, locale = 'pt' }: LectureSchemaProps) {
  const isEn = locale === 'en';
  const url = isEn ? `${BASE}/en/palestras/${slug}` : `${BASE}/palestras/${slug}`;
  const data = {
    '@context': 'https://schema.org',
    '@type': 'CreativeWork',
    name: title,
    description,
    url,
    inLanguage: isEn ? 'en' : 'pt-BR',
    timeRequired: duration,
    audience: audience.map((a) => ({ '@type': 'Audience', audienceType: a })),
    creator: { '@id': `${BASE}/#person` },
    author: { '@id': `${BASE}/#person` },
    learningResourceType: 'Lecture',
  };
  return (
    <script
      type="application/ld+json"
      dangerouslySetInnerHTML={{ __html: JSON.stringify(data) }}
    />
  );
}
