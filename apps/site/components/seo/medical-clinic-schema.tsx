import { brand } from '@plenya/brand';

export function MedicalClinicSchema() {
  const data = {
    '@context': 'https://schema.org',
    '@type': 'MedicalClinic',
    '@id': `${brand.url}/#clinic`,
    name: brand.legalName,
    alternateName: brand.name,
    url: brand.url,
    logo: `${brand.url}/logo.svg`,
    image: `${brand.url}/og-default.jpg`,
    medicalDirector: {
      '@type': 'Physician',
      '@id': `${brand.url}/dr-getulio#physician`,
      name: 'Dr. Getúlio José Mattos do Amaral Filho',
      identifier: [
        { '@type': 'PropertyValue', propertyID: 'CRM', value: 'CRM-PR 21.876' },
        { '@type': 'PropertyValue', propertyID: 'RQE', value: '16.038' },
      ],
    },
    description:
      'Clínica premium de medicina funcional integrativa em Londrina-PR. Consulta Plenya e Continuum Plenya — saúde, performance e longevidade.',
    medicalSpecialty: ['FunctionalMedicine', 'PreventiveMedicine', 'Nephrology', 'InternalMedicine'],
    address: {
      '@type': 'PostalAddress',
      addressLocality: 'Londrina',
      addressRegion: 'PR',
      addressCountry: 'BR',
    },
    sameAs: [brand.social.instagram],
    priceRange: '$$$',
    paymentAccepted: ['Particular', 'PIX', 'Cartão'],
    openingHoursSpecification: [
      {
        '@type': 'OpeningHoursSpecification',
        dayOfWeek: ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday'],
        opens: '08:00',
        closes: '18:00',
      },
    ],
    availableService: [
      {
        '@type': 'MedicalProcedure',
        name: 'Consulta Plenya',
        description: 'Consulta médica particular com leitura funcional dos exames, presencial em Londrina ou online.',
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
