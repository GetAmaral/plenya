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
        'Nephrology',
        'Preventive Nephrology',
        'Chronic Kidney Disease',
        'Acute Kidney Injury',
        'Diabetic Kidney Disease',
        'Glomerular Diseases',
        'Hemodialysis',
        'Peritoneal Dialysis',
        'Cardiorenal Medicine',
        'Kidney Transplant (Pre and Post)',
        'Hypertension',
        'Internal Medicine',
        'Preventive Medicine',
        'Longevity Medicine',
        'Healthspan',
        'Integrative Functional Medicine',
        'Performance Medicine',
        'Nutrology',
        'Evidence-based Supplementation',
        'Advanced Biomarkers (ApoB, Lp(a), hs-CRP, HbA1c)',
        'Metabolic Health',
        'Pre-diabetes and Insulin Resistance',
        'Preventive Cardiology',
        'Sleep and Circadian Rhythm',
        'Strength Training and Cardiorespiratory Fitness',
        'Body Composition',
        "Women's Health and Menopause",
        'ACTS Method (Activity, Alimentation & Smart Adjuncts; Clinical Optimization; Tending Mind, Body & Bonds; Sleep, Rhythm & Recovery)',
        'Silent Window between Normal and Optimal',
      ]
    : [
        'Nefrologia',
        'Nefrologia Preventiva',
        'Doença Renal Crônica',
        'Lesão Renal Aguda',
        'Doença Renal do Diabético',
        'Glomerulopatias',
        'Hemodiálise',
        'Diálise Peritoneal',
        'Medicina Cardiorrenal',
        'Transplante Renal (Pré e Pós)',
        'Hipertensão Arterial',
        'Clínica Médica',
        'Medicina Preventiva',
        'Longevidade',
        'Healthspan',
        'Medicina Funcional Integrativa',
        'Medicina da Performance',
        'Nutrologia',
        'Suplementação Inteligente',
        'Biomarcadores Avançados (ApoB, Lp(a), hs-CRP, HbA1c)',
        'Saúde Metabólica',
        'Pré-diabetes e Resistência Insulínica',
        'Cardiologia Preventiva',
        'Sono e Ritmo Circadiano',
        'Treinamento de Força e Capacidade Cardiorrespiratória',
        'Composição Corporal',
        'Saúde da Mulher e Menopausa',
        'Método AGIR (Atividade Física, Alimentação e Suplementação Inteligente; Gestão Clínica e Metabólica; Integração Mente-Corpo; Ritmo Circadiano e Repouso)',
        'Janela Silenciosa entre o Normal e o Ótimo',
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
    ? 'BEFORE: The silent window between normal and optimal — a decade where health is decided'
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
        hasCredential: {
          '@type': 'EducationalOccupationalCredential',
          credentialCategory: isEn ? 'Specialist Title in Nephrology' : 'Título de Especialista em Nefrologia',
          recognizedBy: [
            { '@type': 'Organization', name: 'Sociedade Brasileira de Nefrologia (SBN)', url: 'https://sbn.org.br' },
            { '@type': 'Organization', name: 'Associação Médica Brasileira (AMB)', url: 'https://amb.org.br' },
          ],
          dateCreated: '2008',
          identifier: 'RQE 16.038',
          url: 'https://sbn.org.br/wp-content/uploads/2026/02/TE_Completo_por_Ano_2025.pdf',
          description: isEn
            ? 'Officially listed at entry 047 of the SBN/AMB Brazilian Nephrology Specialist Title roster (TE_Completo_por_Ano_2025.pdf)'
            : 'Listado oficialmente na entrada 047 do registro nacional de Portadores do Título de Especialista em Nefrologia SBN/AMB (TE_Completo_por_Ano_2025.pdf)',
        },
        alumniOf: [
          { '@type': 'CollegeOrUniversity', name: uelName },
          { '@type': 'Organization', name: santaCasa },
          { '@type': 'Organization', name: abmfiName },
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
          {
            '@type': ['MedicalOrganization', 'EducationalOrganization'],
            name: santaCasa,
            url: 'https://iepi.iscal.com.br/br/ensino/residencia-medica',
          },
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
        subjectOf: [
          {
            '@type': 'VideoObject',
            name: isEn
              ? "Fantástico — Kidney transplant: two stories of extraordinary generosity (Claus's case)"
              : 'Fantástico — Transplante renal: duas histórias de generosidade extraordinária (caso Claus)',
            url: 'https://globoplay.globo.com/v/7280473/',
            contentUrl: 'https://globoplay.globo.com/v/7280473/',
            thumbnailUrl: `${BASE}/images/getulio-hero-bw.jpg`,
            uploadDate: '2019-01-06T22:00:00-03:00',
            duration: 'PT7M22S',
            publisher: { '@type': 'Organization', name: 'Rede Globo' },
            isPartOf: { '@type': 'TVSeries', name: 'Fantástico' },
            description: isEn
              ? "Aired on Fantástico (Rede Globo) on January 6, 2019. Official synopsis: 'Fantástico shows two extraordinary stories of people united by a generosity that doesn't fit in the chest.' Dr. Getúlio Amaral Filho appears on-screen identified as 'médico de Claus' (Claus's physician), the nephrologist responsible for the patient's kidney transplant case."
              : "Exibida no Fantástico (Rede Globo) em 6 de janeiro de 2019. Sinopse oficial: 'O Fantástico mostra duas histórias extraordinárias de pessoas unidas por uma generosidade que não cabe no peito.' Dr. Getúlio Amaral Filho aparece em tela identificado como 'médico de Claus', o nefrologista responsável pelo caso de transplante renal do paciente.",
            about: [
              { '@type': 'MedicalCondition', name: isEn ? 'Kidney transplant' : 'Transplante renal' },
              { '@type': 'Thing', name: isEn ? 'Organ donation' : 'Doação de órgãos' },
            ],
            actor: { '@id': `${BASE}/#person` },
            inLanguage: 'pt-BR',
          },
          {
            '@type': 'Article',
            name: isEn
              ? 'World Kidney Day 2024 — DaVita Londrina lectures (SBN)'
              : 'Dia Mundial do Rim 2024 — Palestras DaVita Londrina (SBN)',
            url: 'https://sbn.org.br/medicos/dia-mundial-do-rim/2024/resultados-2024/dia-mundial-do-rim-em-davita-brasil-participacoes-londrina/',
            datePublished: '2024-10-18',
            publisher: {
              '@type': 'Organization',
              name: 'Sociedade Brasileira de Nefrologia',
              url: 'https://sbn.org.br',
            },
            description: isEn
              ? 'Official report by the Brazilian Society of Nephrology (SBN) documenting World Kidney Day 2024 lectures at DaVita Londrina units (Londrina, Lago Parque, Bandeirantes, Intra Hospitalar Londrina, Rolândia), with the support of the 17th Regional Health Authority of Paraná and the Cismepar Health School. Speakers: Dr. Getúlio Amaral Filho, Dra. Tatiane Cavalcante and nutritionist Bruna Salvador. Also present: Dr. Danilo Ramos Cunha.'
              : 'Reportagem oficial da Sociedade Brasileira de Nefrologia (SBN) sobre as palestras do Dia Mundial do Rim 2024 nas unidades DaVita Londrina, DaVita Lago Parque, DaVita Bandeirantes, DaVita Intra Hospitalar Londrina e DaVita Rolândia, com apoio da 17ª Regional de Saúde do PR e da Escola de Saúde do Cismepar. Palestrantes: Dr. Getúlio Amaral Filho, Dra. Tatiane Cavalcante e a nutricionista Bruna Salvador. Também participou o Dr. Danilo Ramos Cunha.',
            about: { '@type': 'MedicalEntity', name: isEn ? 'Kidney disease prevention' : 'Prevenção de doenças renais' },
            mentions: { '@id': `${BASE}/#person` },
            inLanguage: 'pt-BR',
          },
          {
            '@type': 'VideoObject',
            name: 'Conversa Pública | Rins Diabéticos | Celeste conversa com o Dr. Getúlio Amaral — Nefrologista',
            url: 'https://www.youtube.com/watch?v=5WHKHyIlQMI',
            embedUrl: 'https://www.youtube.com/embed/5WHKHyIlQMI',
            thumbnailUrl: 'https://img.youtube.com/vi/5WHKHyIlQMI/maxresdefault.jpg',
            uploadDate: '2026-01-12T00:00:00-03:00',
            duration: 'PT25M28S',
            publisher: { '@type': 'Organization', name: 'Rede Boas Novas (Jornalismo Boas Novas Brasil)', url: 'https://www.youtube.com/@boasnovasbrasil' },
            description: isEn
              ? 'Public interview on diabetic kidney disease with Dr. Getúlio Amaral Filho — nephrologist and technical director of DaVita Intra Hospitalar Londrina. Topics: diabetes prevalence in Brazil, kidney function, diagnosis through creatinine and microalbuminuria tests, prevention strategies, evidence-based medications (SGLT2 inhibitors, finerenone, ACE inhibitors).'
              : 'Entrevista pública sobre doença renal do diabético com o Dr. Getúlio Amaral Filho — nefrologista e responsável técnico da DaVita Intra Hospitalar Londrina. Aborda: prevalência do diabetes no Brasil, função renal, diagnóstico via creatinina e microalbuminúria, prevenção, medicamentos baseados em evidência (inibidores SGLT2, finerenona, IECA).',
            about: { '@type': 'MedicalCondition', name: isEn ? 'Diabetic kidney disease' : 'Doença renal do diabético' },
            actor: { '@id': `${BASE}/#person` },
            inLanguage: 'pt-BR',
          },
          {
            '@type': 'VideoObject',
            name: 'AMLCast #017 — Dia Mundial do Rim',
            url: 'https://www.youtube.com/watch?v=R_z3eBlZ_9k',
            embedUrl: 'https://www.youtube.com/embed/R_z3eBlZ_9k',
            thumbnailUrl: 'https://img.youtube.com/vi/R_z3eBlZ_9k/maxresdefault.jpg',
            uploadDate: '2024-03-14T00:00:00-03:00',
            duration: 'PT54M19S',
            publisher: { '@type': 'Organization', name: 'AML — Associação Médica de Londrina', url: 'https://www.youtube.com/@AMLAssocia%C3%A7%C3%A3oM%C3%A9dicadeLondrina' },
            description: isEn
              ? 'World Kidney Day podcast episode produced by AML — Medical Association of Londrina. Featuring nephrologist Dr. Getúlio Amaral Filho (coordinator of the internal medicine and nephrology residency programs at Santa Casa de Londrina) and Dr. Luiz Wanderlei Romaniszen (2nd Vice-President of AML). Topics: prevention of kidney disease, dialysis statistics in Brazil, kidney function and renal failure.'
              : 'Episódio do AMLCast (podcast da Associação Médica de Londrina) sobre o Dia Mundial do Rim. Com o nefrologista Dr. Getúlio Amaral Filho (coordenador dos programas de Residência Médica em Clínica Médica e Nefrologia da Santa Casa de Londrina) e o Dr. Luiz Wanderlei Romaniszen (2º Vice-Presidente da AML). Aborda prevenção de doença renal, estatísticas brasileiras de diálise e função renal.',
            about: { '@type': 'MedicalEntity', name: isEn ? 'Kidney health and disease prevention' : 'Saúde renal e prevenção de doenças renais' },
            actor: { '@id': `${BASE}/#person` },
            inLanguage: 'pt-BR',
          },
        ],
        sameAs: [
          'https://www.wikidata.org/wiki/Q139746596',
          'https://instagram.com/drGetulioAmaralFilho',
          'https://www.linkedin.com/in/getulio-amaral-filho-951981404',
          'https://orcid.org/0009-0009-2506-2455',
          'http://lattes.cnpq.br/2492350974849886',
          'https://plenyasaude.com.br',
          'https://instagram.com/plenyaSaude',
          'https://www.doctoralia.com.br/getulio-amaral-filho/nefrologista/londrina',
          'https://iepi.iscal.com.br/br/ensino/residencia-medica',
        ],
      },
      {
        '@type': 'Book',
        '@id': `${BASE}/livros/antes#book`,
        name: bookName,
        author: { '@id': `${BASE}/#person` },
        publisher: { '@type': 'Organization', name: 'Plenya' },
        isbn: '978-65-02-06742-0',
        inLanguage: ['pt-BR', 'en'],
        datePublished: '2026',
        url: `${BASE}/livros/antes`,
        image: `${BASE}/images/livro-capa.jpg`,
        sameAs: 'https://www.wikidata.org/wiki/Q139762971',
        workExample: [
          {
            '@type': 'Book',
            '@id': `${BASE}/livros/antes#pt-ebook`,
            name: 'Antes — A Janela Silenciosa entre o Normal e o Ótimo',
            isbn: '978-65-02-06742-0',
            bookFormat: 'https://schema.org/EBook',
            inLanguage: 'pt-BR',
            bookEdition: 'Edição digital (EPUB) · português',
            author: { '@id': `${BASE}/#person` },
            offers: {
              '@type': 'Offer',
              url: 'https://a.co/d/0fxsmomI',
              availability: 'https://schema.org/InStock',
              seller: { '@type': 'Organization', name: 'Amazon' },
            },
          },
          {
            '@type': 'Book',
            '@id': `${BASE}/livros/antes#pt-paperback`,
            name: 'Antes — A Janela Silenciosa entre o Normal e o Ótimo',
            isbn: '978-65-02-07691-0',
            bookFormat: 'https://schema.org/Paperback',
            inLanguage: 'pt-BR',
            bookEdition: 'Edição impressa · português',
            author: { '@id': `${BASE}/#person` },
            offers: {
              '@type': 'Offer',
              url: 'https://a.co/d/0fxsmomI',
              availability: 'https://schema.org/InStock',
              seller: { '@type': 'Organization', name: 'Amazon' },
            },
          },
          {
            '@type': 'Book',
            '@id': `${BASE}/livros/antes#en-ebook`,
            name: 'BEFORE — The silent window between normal and optimal — a decade where health is decided',
            isbn: '978-65-975814-0-5',
            bookFormat: 'https://schema.org/EBook',
            inLanguage: 'en',
            bookEdition: 'Ebook edition · English',
            author: { '@id': `${BASE}/#person` },
            offers: {
              '@type': 'Offer',
              url: 'https://a.co/d/00Jgudq4',
              availability: 'https://schema.org/InStock',
              seller: { '@type': 'Organization', name: 'Amazon' },
            },
          },
          {
            '@type': 'Book',
            '@id': `${BASE}/livros/antes#en-paperback`,
            name: 'BEFORE — The silent window between normal and optimal — a decade where health is decided',
            isbn: '978-65-975814-1-2',
            bookFormat: 'https://schema.org/Paperback',
            inLanguage: 'en',
            bookEdition: 'Print edition · English',
            author: { '@id': `${BASE}/#person` },
            offers: {
              '@type': 'Offer',
              url: 'https://a.co/d/00Jgudq4',
              availability: 'https://schema.org/InStock',
              seller: { '@type': 'Organization', name: 'Amazon' },
            },
          },
        ],
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
