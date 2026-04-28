export function PersonSchema() {
  const data = {
    '@context': 'https://schema.org',
    '@graph': [
      {
        '@type': ['Person', 'Physician'],
        '@id': 'https://drgetulioamaralfilho.com.br/#person',
        name: 'Dr. Getúlio José Mattos do Amaral Filho',
        url: 'https://drgetulioamaralfilho.com.br',
        image: 'https://drgetulioamaralfilho.com.br/images/getulio-square.jpg',
        jobTitle: 'Médico Nefrologista, Professor, Autor',
        medicalSpecialty: ['Nephrology', 'InternalMedicine'],
        identifier: [
          { '@type': 'PropertyValue', propertyID: 'CRM', value: 'CRM-PR 21.876' },
          { '@type': 'PropertyValue', propertyID: 'RQE', value: '16.038' },
        ],
        alumniOf: [
          { '@type': 'CollegeOrUniversity', name: 'Universidade Estadual de Londrina' },
          { '@type': 'Organization', name: 'Santa Casa de Londrina' },
          { '@type': 'Organization', name: 'Associação Brasileira de Medicina Funcional Integrativa (ABMFI)' },
        ],
        memberOf: [
          { '@type': 'Organization', name: 'Sociedade Brasileira de Nefrologia' },
          { '@type': 'Organization', name: 'Sociedade Paranaense de Nefrologia' },
        ],
        worksFor: [
          {
            '@type': 'MedicalOrganization',
            name: 'Plenya',
            url: 'https://plenyasaude.com.br',
            sameAs: 'https://instagram.com/plenyaSaude',
          },
          { '@type': 'MedicalOrganization', name: 'Nefroclínica Londrina', url: 'https://nefroclinica.com' },
        ],
        address: {
          '@type': 'PostalAddress',
          addressLocality: 'Londrina',
          addressRegion: 'PR',
          addressCountry: 'BR',
        },
        areaServed: [
          { '@type': 'Country', name: 'Brasil' },
          { '@type': 'AdministrativeArea', name: 'Paraná' },
          { '@type': 'City', name: 'Londrina' },
        ],
        knowsLanguage: ['pt-BR'],
        knowsAbout: [
          'Medicina Funcional Integrativa',
          'Nefrologia Preventiva',
          'Longevidade',
          'Healthspan',
          'Hipertensão arterial',
          'Doença renal crônica',
        ],
        sameAs: [
          'https://instagram.com/drGetulioAmaralFilho',
          'https://plenyasaude.com.br',
          'https://instagram.com/plenyaSaude',
        ],
      },
      {
        '@type': 'Book',
        '@id': 'https://drgetulioamaralfilho.com.br/livro#book',
        name: 'ANTES — A Janela Silenciosa entre o Normal e o Ótimo',
        author: { '@id': 'https://drgetulioamaralfilho.com.br/#person' },
        isbn: '978-65-02-06742-0',
        inLanguage: 'pt-BR',
        datePublished: '2026',
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
