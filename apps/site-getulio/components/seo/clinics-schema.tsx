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
    street: 'Av. Duque de Caxias, 1371',
    city: 'Londrina',
    state: 'PR',
    role: 'medicalDirector',
  },
  {
    name: 'Nefroclínica Londrina',
    description:
      'Nefrologia clínica em Londrina há quatro décadas. Doença renal crônica, hipertensão, distúrbios eletrolíticos, transplante renal e acompanhamento pré-diálise.',
    descriptionEn:
      'Clinical nephrology in Londrina for four decades. Chronic kidney disease, hypertension, electrolyte disorders, kidney transplant and pre-dialysis follow-up.',
    url: 'https://nefroclinica.com',
    street: 'Av. Duque de Caxias, 1371',
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
    name: 'DaVita Londrina',
    description:
      'Hemodiálise ambulatorial e hospitalar em Londrina. Acompanhamento crônico de pacientes em terapia renal substitutiva.',
    descriptionEn:
      'Outpatient and in-hospital hemodialysis in Londrina. Long-term follow-up of patients on renal replacement therapy.',
    street: 'Av. Duque de Caxias, 1371',
    city: 'Londrina',
    state: 'PR',
    role: 'medicalDirector',
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
