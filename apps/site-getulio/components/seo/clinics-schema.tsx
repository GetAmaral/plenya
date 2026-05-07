type Clinic = {
  name: string;
  description: string;
  url?: string;
  city: string;
  state: string;
  role:
    | 'medicalDirector'
    | 'physician'
    | 'employee';
};

const BASE = 'https://drgetulioamaralfilho.com.br';

const clinics: Clinic[] = [
  {
    name: 'Plenya — Saúde, Performance e Longevidade',
    description:
      'Clínica de medicina funcional integrativa com equipe multidisciplinar. Programa Continuum Plenya — saúde preventiva e longevidade.',
    url: 'https://plenyasaude.com.br',
    city: 'Londrina',
    state: 'PR',
    role: 'medicalDirector',
  },
  {
    name: 'Nefroclínica Londrina',
    description:
      'Nefrologia clínica em Londrina há quatro décadas. Doença renal crônica, hipertensão, distúrbios eletrolíticos, acompanhamento pré-diálise.',
    url: 'https://nefroclinica.com',
    city: 'Londrina',
    state: 'PR',
    role: 'physician',
  },
  {
    name: 'DaVita — Hemodiálise Intra Hospitalar Santa Casa de Londrina',
    description:
      'Hemodiálise hospitalar em pacientes internados em estágio avançado de doença renal.',
    city: 'Londrina',
    state: 'PR',
    role: 'medicalDirector',
  },
  {
    name: 'DaVita Londrina',
    description:
      'Unidade ambulatorial de hemodiálise. Acompanhamento crônico de pacientes em terapia renal substitutiva.',
    city: 'Londrina',
    state: 'PR',
    role: 'medicalDirector',
  },
];

export function ClinicsSchema({ locale: _locale = 'pt' }: { locale?: string } = {}) {
  // Versão bilíngue completa virá na Fase 6 — estrutura PT como fallback.
  void _locale;
  const data = {
    '@context': 'https://schema.org',
    '@graph': clinics.map((c, i) => ({
      '@type': 'MedicalOrganization',
      '@id': `${BASE}/onde-atendo#clinic-${i}`,
      name: c.name,
      description: c.description,
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
