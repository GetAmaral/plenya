import { brand } from '@plenya/brand';

export function PhysicianSchema({ locale = 'pt' }: { locale?: string } = {}) {
  const isEn = locale === 'en';
  const data = {
    '@context': 'https://schema.org',
    '@type': 'Physician',
    '@id': `${brand.url}/dr-getulio#physician`,
    name: 'Dr. Getúlio José Mattos do Amaral Filho',
    url: `${brand.url}/dr-getulio`,
    image: `${brand.url}/images/team/getulio-amaral.jpg`,
    jobTitle: isEn
      ? 'Clinical Director · Nephrologist · Integrative Functional Medicine (Brazil)'
      : 'Diretor Clínico · Médico Nefrologista · Medicina Funcional Integrativa',
    inLanguage: isEn ? 'en' : 'pt-BR',
    knowsLanguage: ['pt-BR', 'en'],
    medicalSpecialty: ['Nephrology', 'InternalMedicine'],
    knowsAbout: isEn
      ? ['Integrative Functional Medicine', 'Preventive Nephrology', 'Longevity', 'Healthspan']
      : ['Medicina Funcional Integrativa', 'Nefrologia Preventiva', 'Longevidade', 'Healthspan'],
    identifier: [
      { '@type': 'PropertyValue', propertyID: 'CRM', value: 'CRM-PR 21.876' },
      { '@type': 'PropertyValue', propertyID: 'RQE', value: '16.038' },
    ],
    hasCredential: [
      {
        '@type': 'EducationalOccupationalCredential',
        credentialCategory: 'license',
        recognizedBy: { '@type': 'Organization', name: 'Conselho Regional de Medicina do Paraná' },
        name: 'CRM-PR 21.876',
      },
      {
        '@type': 'EducationalOccupationalCredential',
        credentialCategory: 'specialty',
        name: 'Registro de Qualificação de Especialista — Nefrologia (RQE 16.038)',
      },
    ],
    memberOf: [
      { '@type': 'Organization', name: 'Sociedade Brasileira de Nefrologia' },
      { '@type': 'Organization', name: 'Associação Brasileira de Medicina Funcional Integrativa' },
    ],
    worksFor: { '@type': 'MedicalClinic', '@id': `${brand.url}/#clinic` },
    areaServed: [
      { '@type': 'Country', name: 'Brasil' },
      { '@type': 'AdministrativeArea', name: 'Paraná' },
      { '@type': 'City', name: 'Londrina' },
    ],
    availableService: [
      { '@type': 'MedicalProcedure', name: 'Consulta Plenya', procedureType: 'Particular' },
      { '@type': 'MedicalProcedure', name: 'Continuum Plenya', procedureType: 'Particular' },
    ],
    sameAs: [
      'https://drgetulioamaralfilho.com.br',
      'https://instagram.com/drGetulioAmaralFilho',
    ],
  };
  return (
    <script
      type="application/ld+json"
      dangerouslySetInnerHTML={{ __html: JSON.stringify(data) }}
    />
  );
}
