type BookSchemaProps = {
  title: string;
  description: string;
  isbn: string;
  amazonUrl: string;
  hotmartUrl: string;
  coverUrl: string;
  locale?: string;
};

const BASE = 'https://drgetulioamaralfilho.com.br';

export function BookSchema({
  title,
  description,
  isbn,
  amazonUrl,
  hotmartUrl,
  coverUrl,
  locale = 'pt',
}: BookSchemaProps) {
  const isEn = locale === 'en';
  const url = isEn ? `${BASE}/en/livro` : `${BASE}/livro`;
  const data = {
    '@context': 'https://schema.org',
    '@type': 'Book',
    name: title,
    description,
    isbn,
    inLanguage: 'pt-BR',
    url,
    image: coverUrl,
    author: { '@id': `${BASE}/#person` },
    bookFormat: 'https://schema.org/EBook',
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
