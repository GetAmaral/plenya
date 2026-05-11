import { brand } from '@plenya/brand';

export function PhysicianSchema({ locale = 'pt' }: { locale?: string } = {}) {
  const isEn = locale === 'en';
  // Nomes oficiais de instituições brasileiras: PT como nome canônico,
  // EN incluído como tradução paralela para clareza em buscas internacionais.
  const crmCouncil = isEn
    ? 'Regional Medical Council of Paraná (CRM-PR)'
    : 'Conselho Regional de Medicina do Paraná';
  const rqeCredentialName = isEn
    ? 'Specialist Registration — Nephrology (RQE 16.038)'
    : 'Registro de Qualificação de Especialista — Nefrologia (RQE 16.038)';
  const sbnName = isEn ? 'Brazilian Society of Nephrology' : 'Sociedade Brasileira de Nefrologia';
  const abmfiName = isEn
    ? 'Brazilian Association of Integrative Functional Medicine'
    : 'Associação Brasileira de Medicina Funcional Integrativa';
  const country = isEn ? 'Brazil' : 'Brasil';
  const procedureType = isEn ? 'Private' : 'Particular';
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
        recognizedBy: { '@type': 'Organization', name: crmCouncil },
        name: 'CRM-PR 21.876',
      },
      {
        '@type': 'EducationalOccupationalCredential',
        credentialCategory: 'specialty',
        name: rqeCredentialName,
      },
    ],
    memberOf: [
      { '@type': 'Organization', name: sbnName },
      { '@type': 'Organization', name: abmfiName },
    ],
    worksFor: { '@type': 'MedicalClinic', '@id': `${brand.url}/#clinic` },
    areaServed: [
      { '@type': 'Country', name: country },
      { '@type': 'AdministrativeArea', name: 'Paraná' },
      { '@type': 'City', name: 'Londrina' },
    ],
    availableService: [
      { '@type': 'MedicalProcedure', name: 'Consulta Plenya', procedureType },
      { '@type': 'MedicalProcedure', name: 'Continuum Plenya', procedureType },
    ],
    sameAs: [
      'https://drgetulioamaralfilho.com.br',
      'https://instagram.com/drGetulioAmaralFilho',
      'https://www.linkedin.com/in/getulio-amaral-filho-951981404',
      'http://lattes.cnpq.br/2492350974849886',
      'https://www.doctoralia.com.br/getulio-amaral-filho/nefrologista/londrina',
    ],
  };
  return (
    <script
      type="application/ld+json"
      dangerouslySetInnerHTML={{ __html: JSON.stringify(data) }}
    />
  );
}
