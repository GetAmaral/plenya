type ArticleSchemaProps = {
  title: string;
  description: string;
  slug: string;
  date: string;
  tag: string;
  image?: string;
  /**
   * If provided (mirror of an article whose canonical lives elsewhere),
   * mainEntityOfPage points to the canonical URL — Google attributes
   * authority there, not to this mirror.
   */
  canonicalUrl?: string;
  locale?: string;
};

const BASE = 'https://drgetulioamaralfilho.com.br';

export function ArticleSchema({
  title,
  description,
  slug,
  date,
  tag,
  image,
  canonicalUrl,
  locale = 'pt',
}: ArticleSchemaProps) {
  const isEn = locale === 'en';
  const localUrl = isEn ? `${BASE}/en/articles/${slug}` : `${BASE}/artigos/${slug}`;
  const mainUrl = canonicalUrl ?? localUrl;
  const data = {
    '@context': 'https://schema.org',
    '@type': ['BlogPosting', 'MedicalWebPage'],
    headline: title,
    description,
    url: localUrl,
    mainEntityOfPage: { '@type': 'WebPage', '@id': mainUrl },
    datePublished: date,
    dateModified: date,
    inLanguage: isEn ? 'en' : 'pt-BR',
    articleSection: tag,
    image: image ?? `${BASE}/images/getulio-square.jpg`,
    author: { '@id': `${BASE}/#person` },
    creator: { '@id': `${BASE}/#person` },
    publisher: canonicalUrl
      ? {
          '@type': 'MedicalOrganization',
          name: 'Plenya',
          url: 'https://plenyasaude.com.br',
        }
      : {
          '@type': 'Person',
          '@id': `${BASE}/#person`,
          name: 'Dr. Getúlio Amaral Filho',
        },
    reviewedBy: { '@id': `${BASE}/#person` },
    lastReviewed: date,
    isAccessibleForFree: true,
    medicalAudience: { '@type': 'MedicalAudience', audienceType: 'Patient' },
    isPartOf: { '@type': 'WebSite', '@id': `${BASE}/#website` },
    speakable: {
      '@type': 'SpeakableSpecification',
      cssSelector: ['h1', '.editorial-narrow > .space-y-8 > p'],
    },
    ...(canonicalUrl ? { isBasedOn: canonicalUrl } : {}),
  };
  return (
    <script
      type="application/ld+json"
      dangerouslySetInnerHTML={{ __html: JSON.stringify(data) }}
    />
  );
}
