import { brand } from '@plenya/brand';

/**
 * BookSchema do livro ANTES (Dr. Getúlio Amaral Filho).
 *
 * A Plenya não tem página própria do livro — a venda mora em
 * drgetulioamaralfilho.com.br/livros/antes. Este schema é emitido na
 * página /dr-getulio (onde a capa já aparece visualmente) para
 * tornar a obra legível para Google Knowledge Graph mesmo nesta
 * superfície, com `sameAs` ligando o nó Wikidata Q139762971.
 *
 * O autor é referenciado por `@id` apontando para o Physician
 * emitido em PhysicianSchema (`${brand.url}/dr-getulio#physician`),
 * fechando o grafo Pessoa → Obra dentro do site Plenya.
 */

const BOOK_URL_PT = 'https://drgetulioamaralfilho.com.br/livros/antes';
const BOOK_URL_EN = 'https://drgetulioamaralfilho.com.br/en/books/antes';
const COVER_URL = `${brand.url}/images/livro-capa.jpg`;
const AMAZON_PT = 'https://a.co/d/0fxsmomI';
const AMAZON_EN = 'https://a.co/d/00Jgudq4';
const HOTMART_PT = 'https://go.hotmart.com/J105758923K';

export function BookSchema({ locale = 'pt' }: { locale?: string } = {}) {
  const isEn = locale === 'en';
  const bookUrl = isEn ? BOOK_URL_EN : BOOK_URL_PT;
  const name = isEn
    ? 'BEFORE — The silent window between normal and optimal — a decade where health is decided'
    : 'Antes — A Janela Silenciosa entre o Normal e o Ótimo — onde a saúde é decidida';
  const description = isEn
    ? '2026 book by Dr. Getúlio Amaral Filho on preventive medicine, longevity and the silent decade between normal and optimal health.'
    : 'Livro de 2026 do Dr. Getúlio Amaral Filho sobre medicina preventiva, longevidade e a década silenciosa entre o normal e o ótimo.';

  const authorRef = { '@id': `${brand.url}/dr-getulio#physician` };

  const data = {
    '@context': 'https://schema.org',
    '@type': 'Book',
    '@id': `${bookUrl}#book`,
    name,
    description,
    isbn: '978-65-02-06742-0',
    inLanguage: ['pt-BR', 'en'],
    datePublished: '2026',
    numberOfPages: 410,
    url: bookUrl,
    image: COVER_URL,
    author: authorRef,
    publisher: { '@type': 'Organization', name: 'Plenya', url: brand.url },
    sameAs: 'https://www.wikidata.org/wiki/Q139762971',
    workExample: [
      {
        '@type': 'Book',
        '@id': `${bookUrl}#pt-ebook`,
        name: 'Antes — A Janela Silenciosa entre o Normal e o Ótimo',
        isbn: '978-65-02-06742-0',
        bookFormat: 'https://schema.org/EBook',
        inLanguage: 'pt-BR',
        bookEdition: 'Edição digital (EPUB) · português',
        author: authorRef,
        offers: {
          '@type': 'Offer',
          url: AMAZON_PT,
          availability: 'https://schema.org/InStock',
          seller: { '@type': 'Organization', name: 'Amazon' },
        },
      },
      {
        '@type': 'Book',
        '@id': `${bookUrl}#pt-paperback`,
        name: 'Antes — A Janela Silenciosa entre o Normal e o Ótimo',
        isbn: '978-65-02-07691-0',
        bookFormat: 'https://schema.org/Paperback',
        inLanguage: 'pt-BR',
        bookEdition: 'Edição impressa · português',
        author: authorRef,
        offers: {
          '@type': 'Offer',
          url: AMAZON_PT,
          availability: 'https://schema.org/InStock',
          seller: { '@type': 'Organization', name: 'Amazon' },
        },
      },
      {
        '@type': 'Book',
        '@id': `${bookUrl}#en-ebook`,
        name: 'BEFORE — The silent window between normal and optimal — a decade where health is decided',
        isbn: '978-65-975814-0-5',
        bookFormat: 'https://schema.org/EBook',
        inLanguage: 'en',
        bookEdition: 'Ebook edition · English',
        author: authorRef,
        offers: {
          '@type': 'Offer',
          url: AMAZON_EN,
          availability: 'https://schema.org/InStock',
          seller: { '@type': 'Organization', name: 'Amazon' },
        },
      },
      {
        '@type': 'Book',
        '@id': `${bookUrl}#en-paperback`,
        name: 'BEFORE — The silent window between normal and optimal — a decade where health is decided',
        isbn: '978-65-975814-1-2',
        bookFormat: 'https://schema.org/Paperback',
        inLanguage: 'en',
        bookEdition: 'Print edition · English',
        author: authorRef,
        offers: {
          '@type': 'Offer',
          url: AMAZON_EN,
          availability: 'https://schema.org/InStock',
          seller: { '@type': 'Organization', name: 'Amazon' },
        },
      },
    ],
    offers: [
      {
        '@type': 'Offer',
        url: isEn ? AMAZON_EN : AMAZON_PT,
        availability: 'https://schema.org/InStock',
        seller: { '@type': 'Organization', name: 'Amazon' },
      },
      ...(isEn
        ? []
        : [
            {
              '@type': 'Offer',
              url: HOTMART_PT,
              availability: 'https://schema.org/InStock',
              seller: { '@type': 'Organization', name: 'Hotmart' },
            },
          ]),
    ],
  };

  return (
    <script
      type="application/ld+json"
      dangerouslySetInnerHTML={{ __html: JSON.stringify(data) }}
    />
  );
}
