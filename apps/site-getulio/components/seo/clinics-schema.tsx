type Clinic = {
  name: string;
  description: string;
  descriptionEn: string;
  url?: string;
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
      'Clínica de medicina funcional integrativa com equipe multidisciplinar. Programa Continuum Plenya — saúde preventiva e longevidade.',
    descriptionEn:
      'Integrative functional medicine clinic with a multidisciplinary team. Continuum Plenya program — preventive health and longevity.',
    url: 'https://plenyasaude.com.br',
    city: 'Londrina',
    state: 'PR',
    role: 'medicalDirector',
  },
  {
    name: 'Nefroclínica Londrina',
    description:
      'Nefrologia clínica em Londrina há quatro décadas. Doença renal crônica, hipertensão, distúrbios eletrolíticos, acompanhamento pré-diálise.',
    descriptionEn:
      'Clinical nephrology in Londrina for four decades. Chronic kidney disease, hypertension, electrolyte disorders, pre-dialysis follow-up.',
    url: 'https://nefroclinica.com',
    city: 'Londrina',
    state: 'PR',
    role: 'physician',
  },
  {
    name: 'DaVita — Hemodiálise Intra Hospitalar Santa Casa de Londrina',
    description:
      'Hemodiálise hospitalar em pacientes internados em estágio avançado de doença renal.',
    descriptionEn:
      'In-hospital hemodialysis for inpatients in advanced stages of kidney disease.',
    city: 'Londrina',
    state: 'PR',
    role: 'medicalDirector',
  },
  {
    name: 'DaVita Londrina',
    description:
      'Unidade ambulatorial de hemodiálise. Acompanhamento crônico de pacientes em terapia renal substitutiva.',
    descriptionEn:
      'Outpatient hemodialysis unit. Long-term follow-up of patients on renal replacement therapy.',
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
