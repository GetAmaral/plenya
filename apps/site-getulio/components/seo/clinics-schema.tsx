type Clinic = {
  name: string;
  description: string;
  descriptionEn: string;
  url?: string;
  street?: string;
  neighborhood?: string;
  postalCode?: string;
  city: string;
  state: string;
  geo?: { latitude: number; longitude: number };
  role:
    | 'medicalDirector'
    | 'physician'
    | 'employee';
};

const BASE = 'https://drgetulioamaralfilho.com.br';

// Nomes próprios mantidos em PT (são razões sociais brasileiras).
// Apenas as descrições ganham versão EN.
const clinics: Clinic[] = [
  {
    name: 'Plenya — Saúde, Performance e Longevidade',
    description:
      'Clínica de medicina funcional integrativa com equipe multidisciplinar. Programa Continuum Plenya — saúde preventiva e longevidade. Atendimento presencial em Londrina ou online por telemedicina.',
    descriptionEn:
      'Integrative functional medicine clinic with a multidisciplinary team. Continuum Plenya program — preventive health and longevity. In-person in Londrina or online via telemedicine.',
    url: 'https://plenyasaude.com.br',
    street: 'Av. Ayrton Senna da Silva, 500 — Edifício Torre Pietra, sala 1402',
    neighborhood: 'Gleba Palhano',
    postalCode: '86050-460',
    city: 'Londrina',
    state: 'PR',
    geo: { latitude: -23.3296924, longitude: -51.1779253 },
    role: 'medicalDirector',
  },
  {
    name: 'Nefroclínica Londrina',
    description:
      'Nefrologia clínica em Londrina há quatro décadas. Doença renal crônica em todos os estágios, hipertensão, distúrbios eletrolíticos, glomerulopatias, doença renal do diabético, transplante renal (pré e pós), acompanhamento pré-diálise e seguimento longitudinal de pacientes em hemodiálise e diálise peritoneal.',
    descriptionEn:
      'Clinical nephrology in Londrina for four decades. Chronic kidney disease at all stages, hypertension, electrolyte disorders, glomerular diseases, diabetic kidney disease, kidney transplant (pre and post), pre-dialysis follow-up and longitudinal care of patients on hemodialysis and peritoneal dialysis.',
    url: 'https://nefroclinica.com',
    street: 'Av. Duque de Caxias, 1371',
    neighborhood: 'Jardim Petrópolis',
    postalCode: '86015-000',
    city: 'Londrina',
    state: 'PR',
    role: 'physician',
  },
  {
    name: 'Santa Casa de Londrina',
    description:
      'Hospital geral em Londrina. Dr. Getúlio Amaral Filho atua na nefrologia hospitalar — pareceres e seguimento de pacientes internados.',
    descriptionEn:
      'General hospital in Londrina. Dr. Getúlio Amaral Filho practices in hospital nephrology — inpatient consultations and follow-up.',
    url: 'https://www.iscal.com.br',
    street: 'Rua Espírito Santo, 523',
    neighborhood: 'Centro',
    postalCode: '86010-510',
    city: 'Londrina',
    state: 'PR',
    role: 'physician',
  },
  {
    name: 'Hospital Araucária',
    description:
      'Hospital geral em Londrina. Dr. Getúlio Amaral Filho atua na nefrologia hospitalar — pareceres e seguimento de pacientes internados.',
    descriptionEn:
      'General hospital in Londrina. Dr. Getúlio Amaral Filho practices in hospital nephrology — inpatient consultations and follow-up.',
    url: 'https://hospitalaraucaria.com.br',
    street: 'Rua Campo Grande, 211',
    neighborhood: 'Jardim Colina Verde',
    postalCode: '86050-550',
    city: 'Londrina',
    state: 'PR',
    role: 'physician',
  },
  {
    name: 'Hospital Unimed Londrina',
    description:
      'Hospital próprio da Unimed Londrina. Dr. Getúlio Amaral Filho atua na nefrologia hospitalar — pareceres e seguimento de pacientes internados.',
    descriptionEn:
      'Unimed Londrina own hospital. Dr. Getúlio Amaral Filho practices in hospital nephrology — inpatient consultations and follow-up.',
    url: 'https://www.unimedlondrina.com.br/unidades-e-servicos/hospital-unimed/',
    street: 'Av. dos Expedicionários, 750',
    city: 'Londrina',
    state: 'PR',
    role: 'physician',
  },
  {
    name: 'DaVita Intra Hospitalar Londrina',
    description:
      'Hemodiálise hospitalar para pacientes internados em estágio avançado de doença renal ou em lesão renal aguda. Dr. Getúlio Amaral Filho atua como responsável técnico, em conjunto com as equipes médicas dos hospitais.',
    descriptionEn:
      'In-hospital hemodialysis for inpatients in advanced kidney disease or acute kidney injury. Dr. Getúlio Amaral Filho serves as technical director, working alongside the medical teams of the partner hospitals.',
    city: 'Londrina',
    state: 'PR',
    role: 'medicalDirector',
  },
  {
    name: 'DaVita Londrina',
    description:
      'Hemodiálise ambulatorial e diálise peritoneal em Londrina. Seguimento longitudinal de pacientes em terapia renal substitutiva crônica. Dr. Getúlio Amaral Filho atende como médico nefrologista.',
    descriptionEn:
      'Outpatient hemodialysis and peritoneal dialysis in Londrina. Long-term follow-up of patients on chronic renal replacement therapy. Dr. Getúlio Amaral Filho practices as attending nephrologist.',
    street: 'Av. Duque de Caxias, 1371',
    neighborhood: 'Jardim Petrópolis',
    postalCode: '86015-000',
    city: 'Londrina',
    state: 'PR',
    role: 'physician',
  },
];

export function ClinicsSchema({ locale = 'pt' }: { locale?: string } = {}) {
  const isEn = locale === 'en';
  const data = {
    '@context': 'https://schema.org',
    '@graph': clinics.map((c, i) => ({
      '@type': 'MedicalOrganization',
      '@id': `${BASE}/onde-atendo#clinic-${i}`,
      name: c.name,
      description: isEn ? c.descriptionEn : c.description,
      inLanguage: isEn ? 'en' : 'pt-BR',
      ...(c.url ? { url: c.url, sameAs: [c.url] } : {}),
      address: {
        '@type': 'PostalAddress',
        ...(c.street ? { streetAddress: c.street } : {}),
        ...(c.neighborhood ? { addressDistrict: c.neighborhood } : {}),
        ...(c.postalCode ? { postalCode: c.postalCode } : {}),
        addressLocality: c.city,
        addressRegion: c.state,
        addressCountry: 'BR',
      },
      ...(c.geo
        ? { geo: { '@type': 'GeoCoordinates', latitude: c.geo.latitude, longitude: c.geo.longitude } }
        : {}),
      areaServed: { '@type': 'City', name: c.city },
      ...(c.role === 'medicalDirector'
        ? { medicalDirector: { '@id': `${BASE}/#person` } }
        : { physician: { '@id': `${BASE}/#person` } }),
    })),
  };
  return (
    <script
      type="application/ld+json"
      dangerouslySetInnerHTML={{ __html: JSON.stringify(data) }}
    />
  );
}
