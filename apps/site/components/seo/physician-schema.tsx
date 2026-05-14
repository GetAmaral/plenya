import { brand } from '@plenya/brand';

export function PhysicianSchema({ locale = 'pt' }: { locale?: string } = {}) {
  const isEn = locale === 'en';
  // Nomes oficiais de instituições brasileiras: PT como nome canônico,
  // EN incluído como tradução paralela para clareza em buscas internacionais.
  const crmCouncil = isEn
    ? 'Regional Medical Council of Paraná (CRM-PR)'
    : 'Conselho Regional de Medicina do Paraná';
  const rqeCredentialName = isEn
    ? 'Specialist Title in Nephrology (RQE 16.038)'
    : 'Título de Especialista em Nefrologia (RQE 16.038)';
  const sbnName = isEn ? 'Brazilian Society of Nephrology' : 'Sociedade Brasileira de Nefrologia';
  const ambName = isEn ? 'Brazilian Medical Association' : 'Associação Médica Brasileira';
  const abmfiName = isEn
    ? 'Brazilian Academy of Integrative Functional Medicine'
    : 'Academia Brasileira de Medicina Funcional Integrativa';
  const country = isEn ? 'Brazil' : 'Brasil';
  const procedureType = isEn ? 'Private' : 'Particular';
  const bookName = isEn
    ? 'BEFORE — The silent window between normal and optimal — a decade where health is decided'
    : 'Antes — A Janela Silenciosa entre o Normal e o Ótimo — onde a saúde é decidida';
  const bookUrl = isEn
    ? 'https://drgetulioamaralfilho.com.br/en/books/antes'
    : 'https://drgetulioamaralfilho.com.br/livros/antes';

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
      {
        '@type': 'PropertyValue',
        propertyID: 'ORCID',
        value: '0009-0009-2506-2455',
        url: 'https://orcid.org/0009-0009-2506-2455',
      },
      {
        '@type': 'PropertyValue',
        propertyID: 'Lattes',
        value: '2492350974849886',
        url: 'http://lattes.cnpq.br/2492350974849886',
      },
      {
        '@type': 'PropertyValue',
        propertyID: 'Wikidata',
        value: 'Q139746596',
        url: 'https://www.wikidata.org/wiki/Q139746596',
      },
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
        recognizedBy: [
          { '@type': 'Organization', name: sbnName, url: 'https://sbn.org.br' },
          { '@type': 'Organization', name: ambName, url: 'https://amb.org.br' },
        ],
        dateCreated: '2008',
        identifier: 'RQE 16.038',
        url: 'https://sbn.org.br/wp-content/uploads/2026/02/TE_Completo_por_Ano_2025.pdf',
        description: isEn
          ? 'Officially listed at entry 047 of the SBN/AMB Brazilian Nephrology Specialist Title roster (TE_Completo_por_Ano_2025.pdf)'
          : 'Listado oficialmente na entrada 047 do registro nacional de Portadores do Título de Especialista em Nefrologia SBN/AMB (TE_Completo_por_Ano_2025.pdf)',
      },
    ],
    memberOf: [
      {
        '@type': 'OrganizationRole',
        roleName: isEn ? 'Cardiorenal Committee Member' : 'Membro do Comitê Cardiorrenal',
        memberOf: {
          '@type': 'Organization',
          name: sbnName,
          url: 'https://sbn.org.br/medicos/a-sbn/comites/',
        },
      },
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
    subjectOf: {
      '@type': 'Book',
      '@id': `${bookUrl}#book`,
      name: bookName,
      url: bookUrl,
      datePublished: '2026',
      isbn: '978-65-02-06742-0',
      sameAs: 'https://www.wikidata.org/wiki/Q139762971',
    },
    sameAs: [
      'https://www.wikidata.org/wiki/Q139746596',
      'https://drgetulioamaralfilho.com.br',
      'https://instagram.com/drGetulioAmaralFilho',
      'https://www.linkedin.com/in/getulio-amaral-filho-951981404',
      'https://orcid.org/0009-0009-2506-2455',
      'http://lattes.cnpq.br/2492350974849886',
      'https://www.doctoralia.com.br/getulio-amaral-filho/nefrologista/londrina',
      'https://busca.abmfuncionalintegrativa.com.br/profissional/getulio-jose-mattos-do-amaral-filho-2d3ee592',
    ],
  };
  return (
    <script
      type="application/ld+json"
      dangerouslySetInnerHTML={{ __html: JSON.stringify(data) }}
    />
  );
}
