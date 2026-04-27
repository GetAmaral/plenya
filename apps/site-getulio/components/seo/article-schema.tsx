type ArticleSchemaProps = {
  title: string;
  description: string;
  slug: string;
  date: string;
  tag: string;
  image?: string;
};

const BASE = 'https://drgetulioamaralfilho.com.br';

export function ArticleSchema({ title, description, slug, date, tag, image }: ArticleSchemaProps) {
  const url = `${BASE}/escritos/${slug}`;
  const data = {
    '@context': 'https://schema.org',
    '@type': ['BlogPosting', 'MedicalWebPage'],
    headline: title,
    description,
    url,
    mainEntityOfPage: { '@type': 'WebPage', '@id': url },
    datePublished: date,
    dateModified: date,
    inLanguage: 'pt-BR',
    articleSection: tag,
    image: image ? `${BASE}${image}` : `${BASE}/images/getulio-square.jpg`,
    author: { '@id': `${BASE}/#person` },
    creator: { '@id': `${BASE}/#person` },
    publisher: {
      '@type': 'Person',
      '@id': `${BASE}/#person`,
      name: 'Dr. Getúlio Amaral Filho',
    },
    reviewedBy: { '@id': `${BASE}/#person` },
    isAccessibleForFree: true,
    medicalAudience: { '@type': 'MedicalAudience', audienceType: 'Patient' },
  };
  return (
    <script
      type="application/ld+json"
      dangerouslySetInnerHTML={{ __html: JSON.stringify(data) }}
    />
  );
}
