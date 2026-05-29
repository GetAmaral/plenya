import { brand } from '@plenya/brand';

export function MedicalClinicSchema({ locale = 'pt' }: { locale?: string } = {}) {
  const isEn = locale === 'en';
  const data = {
    '@context': 'https://schema.org',
    '@type': 'MedicalClinic',
    '@id': `${brand.url}/#clinic`,
    name: brand.legalName,
    legalName: brand.companyName,
    alternateName: brand.name,
    taxID: brand.cnpj,
    url: brand.url,
    logo: `${brand.url}/logo.svg`,
    image: `${brand.url}/og-default.jpg`,
    telephone: '+5543999748899',
    medicalDirector: {
      '@type': 'Physician',
      '@id': `${brand.url}/dr-getulio#physician`,
      name: 'Dr. Getúlio José Mattos do Amaral Filho',
      identifier: [
        { '@type': 'PropertyValue', propertyID: 'CRM', value: 'CRM-PR 21.876' },
        { '@type': 'PropertyValue', propertyID: 'RQE', value: '16.038' },
      ],
    },
    description: isEn
      ? 'Premium integrative functional medicine clinic in Londrina, Brazil. Plenya Consultation and Continuum Plenya — health, performance, and longevity.'
      : 'Clínica premium de medicina funcional integrativa em Londrina-PR. Consulta Plenya e Continuum Plenya — saúde, performance e longevidade.',
    medicalSpecialty: ['FunctionalMedicine', 'PreventiveMedicine', 'Nephrology', 'InternalMedicine'],
    address: {
      '@type': 'PostalAddress',
      streetAddress: `${brand.address.street} — ${brand.address.complement}`,
      addressLocality: brand.address.city,
      addressRegion: brand.address.state,
      postalCode: brand.address.postalCode,
      addressCountry: brand.address.country,
    },
    geo: {
      '@type': 'GeoCoordinates',
      latitude: brand.address.geo.latitude,
      longitude: brand.address.geo.longitude,
    },
    hasMap: `https://www.google.com/maps/search/?api=1&query=${brand.address.geo.latitude},${brand.address.geo.longitude}`,
    areaServed: [
      { '@type': 'Country', name: isEn ? 'Brazil' : 'Brasil' },
      { '@type': 'AdministrativeArea', name: isEn ? 'Paraná (Brazil)' : 'Paraná' },
      { '@type': 'City', name: 'Londrina' },
    ],
    inLanguage: isEn ? 'en' : 'pt-BR',
    knowsLanguage: ['pt-BR', 'en'],
    sameAs: [
      brand.social.instagram,
      'https://drgetulioamaralfilho.com.br',
      'https://instagram.com/drGetulioAmaralFilho',
    ],
    priceRange: '$$$',
    paymentAccepted: isEn ? ['Self-pay', 'PIX', 'Card'] : ['Particular', 'PIX', 'Cartão'],
    openingHoursSpecification: [
      {
        '@type': 'OpeningHoursSpecification',
        dayOfWeek: ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday'],
        opens: '08:00',
        closes: '18:00',
      },
    ],
    availableService: isEn
      ? [
          {
            '@type': 'MedicalProcedure',
            name: 'Plenya Consultation',
            description:
              'Private medical consultation with a functional reading of labs, in person in Londrina or online.',
          },
          {
            '@type': 'MedicalProcedure',
            name: 'Continuum Plenya',
            description:
              'Six-month or annual program with multidisciplinary team — physician, nutritionist, psychologist, and exercise physiologist.',
          },
          {
            '@type': 'DiagnosticProcedure',
            name: 'Plenya Score',
            description:
              'Measurement instrument of The ACTS Method. Over 800 items synthesized into a clear, evolving, personalized score.',
          },
        ]
      : [
          {
            '@type': 'MedicalProcedure',
            name: 'Consulta Plenya',
            description:
              'Consulta médica particular com leitura funcional dos exames, presencial em Londrina ou online.',
          },
          {
            '@type': 'MedicalProcedure',
            name: 'Continuum Plenya',
            description:
              'Programa semestral ou anual com equipe multidisciplinar — médico, nutricionista, psicólogo e educador físico.',
          },
          {
            '@type': 'DiagnosticProcedure',
            name: 'Escore Plenya',
            description:
              'Instrumento de medida do Método AGIR. Mais de 800 itens em uma pontuação clara, evolutiva e personalizada.',
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
