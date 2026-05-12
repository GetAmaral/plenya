type Edition = {
  id: string;
  format: 'EBook' | 'Paperback' | 'Hardcover' | 'AudiobookFormat';
  language: string;
  isbn: string;
  label: string;
  labelEn?: string;
  purchaseUrl?: string;
};

type BookSchemaProps = {
  title: string;
  description: string;
  isbn: string;
  slug: string;
  amazonUrl: string;
  hotmartUrl: string;
  coverUrl: string;
  locale?: string;
  editions?: Edition[];
};

const BASE = 'https://drgetulioamaralfilho.com.br';

export function BookSchema({
  title,
  description,
  isbn,
  slug,
  amazonUrl,
  hotmartUrl,
  coverUrl,
  locale = 'pt',
  editions,
}: BookSchemaProps) {
  const isEn = locale === 'en';
  const url = isEn ? `${BASE}/en/books/${slug}` : `${BASE}/livros/${slug}`;

  const workExample = (editions ?? []).map((ed) => ({
    '@type': 'Book',
    '@id': `${BASE}/livros/${slug}#${ed.id}`,
    name: title,
    isbn: ed.isbn,
    inLanguage: ed.language,
    bookFormat: `https://schema.org/${ed.format}`,
    bookEdition: isEn && ed.labelEn ? ed.labelEn : ed.label,
    image: coverUrl,
    author: { '@id': `${BASE}/#person` },
    ...(ed.purchaseUrl
      ? {
          offers: {
            '@type': 'Offer',
            url: ed.purchaseUrl,
            availability: 'https://schema.org/InStock',
            seller: { '@type': 'Organization', name: 'Amazon' },
          },
        }
      : {}),
  }));

  const data = {
    '@context': 'https://schema.org',
    '@type': 'Book',
    '@id': `${BASE}/livros/${slug}#book`,
    name: title,
    description,
    isbn,
    inLanguage: workExample.length > 0 ? Array.from(new Set(workExample.map((w) => w.inLanguage))) : 'pt-BR',
    url,
    image: coverUrl,
    author: { '@id': `${BASE}/#person` },
    datePublished: '2026',
    sameAs: 'https://www.wikidata.org/wiki/Q139762971',
    ...(workExample.length > 0 ? { workExample } : { bookFormat: 'https://schema.org/EBook' }),
    offers: [
      {
        '@type': 'Offer',
        url: amazonUrl,
        seller: { '@type': 'Organization', name: 'Amazon' },
        availability: 'https://schema.org/InStock',
      },
      {
        '@type': 'Offer',
        url: hotmartUrl,
        seller: { '@type': 'Organization', name: 'Hotmart' },
        availability: 'https://schema.org/InStock',
      },
    ],
  };
  return (
    <script
      type="application/ld+json"
      dangerouslySetInnerHTML={{ __html: JSON.stringify(data) }}
    />
  );
}
