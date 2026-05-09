/**
 * Person/Physician schema do Dr. Getúlio.
 *
 * Estratégia bilíngue: nomes próprios brasileiros (UEL, Santa Casa, ABMFI,
 * SBN, SPN) ficam canônicos em PT mesmo no payload EN — schema.org não
 * exige tradução de razões sociais. O que muda em EN: jobTitle, knowsAbout
 * (conceitos), areaServed.country, inLanguage e o título do livro.
 */

const BASE = 'https://drgetulioamaralfilho.com.br';

export function PersonSchema({ locale = 'pt' }: { locale?: string } = {}) {
  const isEn = locale === 'en';

  const jobTitle = isEn
    ? 'Nephrologist · Professor · Author'
    : 'Médico Nefrologista, Professor, Autor';

  const knowsAbout = isEn
    ? [
        'Integrative Functional Medicine',
        'Preventive Nephrology',
        'Longevity',
        'Healthspan',
        'Hypertension',
        'Chronic Kidney Disease',
      ]
    : [
        'Medicina Funcional Integrativa',
        'Nefrologia Preventiva',
        'Longevidade',
        'Healthspan',
        'Hipertensão arterial',
        'Doença renal crônica',
      ];

  const country = isEn ? 'Brazil' : 'Brasil';

  // Nomes paralelos PT (canônico) + EN entre parênteses para clareza.
  const uelName = isEn
    ? 'Universidade Estadual de Londrina (State University of Londrina)'
    : 'Universidade Estadual de Londrina';
  const santaCasa = isEn
    ? 'Santa Casa de Londrina (Santa Casa Hospital, Londrina)'
    : 'Santa Casa de Londrina';
  const abmfiName = isEn
    ? 'Associação Brasileira de Medicina Funcional Integrativa (Brazilian Association of Integrative Functional Medicine, ABMFI)'
    : 'Associação Brasileira de Medicina Funcional Integrativa (ABMFI)';
  const sbnName = isEn
    ? 'Sociedade Brasileira de Nefrologia (Brazilian Society of Nephrology)'
    : 'Sociedade Brasileira de Nefrologia';
  const spnName = isEn
    ? 'Sociedade Paranaense de Nefrologia (Paraná Society of Nephrology)'
    : 'Sociedade Paranaense de Nefrologia';
  const nefroclinica = isEn ? 'Nefroclínica Londrina (clinical nephrology)' : 'Nefroclínica Londrina';

  const bookName = isEn
    ? 'Antes: The Silent Window Between Normal and Optimal — Where Health Is Decided (Brazilian Portuguese edition)'
    : 'Antes: A Janela Silenciosa entre o Normal e o Ótimo — onde a saúde é decidida';

  const data = {
    '@context': 'https://schema.org',
    '@graph': [
      {
        '@type': ['Person', 'Physician'],
        '@id': `${BASE}/#person`,
        name: 'Dr. Getúlio José Mattos do Amaral Filho',
        url: BASE,
        image: `${BASE}/images/getulio-square.jpg`,
        jobTitle,
        medicalSpecialty: ['Nephrology', 'InternalMedicine'],
        identifier: [
          { '@type': 'PropertyValue', propertyID: 'CRM', value: 'CRM-PR 21.876' },
          { '@type': 'PropertyValue', propertyID: 'RQE', value: '16.038' },
        ],
        alumniOf: [
          { '@type': 'CollegeOrUniversity', name: uelName },
          { '@type': 'Organization', name: santaCasa },
          { '@type': 'Organization', name: abmfiName },
        ],
        memberOf: [
          { '@type': 'Organization', name: sbnName },
          { '@type': 'Organization', name: spnName },
        ],
        worksFor: [
          {
            '@type': 'MedicalOrganization',
            name: 'Plenya',
            url: 'https://plenyasaude.com.br',
            sameAs: 'https://instagram.com/plenyaSaude',
          },
          { '@type': 'MedicalOrganization', name: nefroclinica, url: 'https://nefroclinica.com' },
        ],
        address: {
          '@type': 'PostalAddress',
          addressLocality: 'Londrina',
          addressRegion: 'PR',
          addressCountry: 'BR',
        },
        areaServed: [
          { '@type': 'Country', name: country },
          { '@type': 'AdministrativeArea', name: 'Paraná' },
          { '@type': 'City', name: 'Londrina' },
        ],
        knowsLanguage: ['pt-BR', 'en'],
        knowsAbout,
        sameAs: [
          'https://instagram.com/drGetulioAmaralFilho',
          'https://plenyasaude.com.br',
          'https://instagram.com/plenyaSaude',
        ],
      },
      {
        '@type': 'Book',
        '@id': `${BASE}/livros/antes#book`,
        name: bookName,
        author: { '@id': `${BASE}/#person` },
        isbn: '978-65-02-06742-0',
        inLanguage: 'pt-BR',
        datePublished: '2026',
        offers: [
          {
            '@type': 'Offer',
            url: 'https://a.co/d/0fxsmomI',
            availability: 'https://schema.org/InStock',
            seller: { '@type': 'Organization', name: 'Amazon' },
          },
          {
            '@type': 'Offer',
            url: 'https://go.hotmart.com/J105758923K',
            availability: 'https://schema.org/InStock',
            seller: { '@type': 'Organization', name: 'Hotmart' },
          },
        ],
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
