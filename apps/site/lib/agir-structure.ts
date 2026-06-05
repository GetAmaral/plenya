/**
 * Estrutura completa do Método AGIR (PT) ↔ The ACTS Method (EN).
 * Usada pelas páginas /escore-plenya e /metodo-agir. Reflete a hierarquia do EMR:
 * Method → MethodLetter → MethodPillar.
 *
 * Mapping de letras: A→A, G→C, I→T, R→S.
 *
 * ⚙️ Sincronizado do banco do EMR (42 pilares) — gerado a partir de method_letters +
 *    method_pillars. Os scores do radar são MOCK (dados fictícios de marketing; sem
 *    paciente real). Para re-sincronizar, repuxe os pilares do EMR e regenere este arquivo.
 *
 * Inlined em apps/site para evitar problemas de cache do Turbopack
 * com export indireto de @plenya/brand.
 */

export type AgirLetterRaw = {
  code: 'A' | 'G' | 'I' | 'R';
  codeEn: 'A' | 'C' | 'T' | 'S';
  name: string;
  nameEn: string;
  pillarCount: number;
  itemCount: number; // contagem real de itens do EMR (estimativa agregada)
  /** Quando há muitos pilares (>6), agrupar visualmente. */
  groups: { label?: string; labelEn?: string; pillars: string[]; pillarsEn: string[] }[];
};

export const agirLettersRaw: AgirLetterRaw[] = [
  {
    code: 'A',
    codeEn: 'A',
    name: 'Atividade Física, Alimentação e Suplementação Inteligente',
    nameEn: 'Activity, Alimentation & Smart Adjuncts',
    pillarCount: 7,
    itemCount: 290,
    groups: [
      {
        pillars: [
          'Alimentação (Qualidade e Padrão)',
          'Tolerância Alimentar e Suplementação',
          'Atividade Física (Hábito)',
          'Capacidade Física (Testes)',
          'Adiposidade e Risco Cardiometabólico',
          'Massa Muscular e Hidratação Celular',
          'Trajetória e Origens',
        ],
        pillarsEn: [
          'Nutrition (Quality & Pattern)',
          'Food Tolerance & Supplementation',
          'Physical Activity (Habit)',
          'Physical Capacity (Testing)',
          'Adiposity & Cardiometabolic Risk',
          'Muscle Mass & Cellular Hydration',
          'Trajectory & Origins',
        ],
      },
    ],
  },
  {
    code: 'G',
    codeEn: 'C',
    name: 'Gestão Clínica e Metabólica',
    nameEn: 'Clinical Optimization',
    pillarCount: 23,
    itemCount: 1027,
    groups: [
      {
        pillars: [
          'Cardiovascular',
          'Metabólico',
          'Renal',
          'Hepático',
          'Gastrointestinal',
          'Pulmonar',
          'Hematológico',
          'Hormonal',
          'Imune e Inflamatório',
          'Neurológico',
          'Osteomuscular',
          'Nutrologia e Micronutrientes',
          'Rastreio Oncológico',
          'Genético',
          'Infeccioso',
          'Toxicológico e Detoxificação',
          'Dermatológico',
          'Saúde Urogenital e Reprodutiva',
          'Órgãos dos Sentidos',
          'Odontológico',
          'Medicamentos',
          'História Cirúrgica',
          'Hábitos e Vícios',
        ],
        pillarsEn: [
          'Cardiovascular',
          'Metabolic',
          'Renal',
          'Hepatic',
          'Gastrointestinal',
          'Pulmonary',
          'Hematological',
          'Hormonal',
          'Immune & Inflammatory',
          'Neurological',
          'Musculoskeletal',
          'Nutritional Medicine & Micronutrients',
          'Cancer Screening',
          'Genetic',
          'Infectious',
          'Toxicology & Detoxification',
          'Dermatological',
          'Urogenital & Reproductive Health',
          'Sensory Organs',
          'Dental',
          'Medications',
          'Surgical History',
          'Habits & Addictions',
        ],
      },
    ],
  },
  {
    code: 'I',
    codeEn: 'T',
    name: 'Integração Mente-Corpo',
    nameEn: 'Tending Mind, Body & Bonds',
    pillarCount: 7,
    itemCount: 199,
    groups: [
      {
        pillars: [
          'Humor',
          'Função Cognitiva',
          'Vitalidade e Disposição',
          'Estresse, Trauma e Resiliência',
          'Vida Sexual',
          'Vínculos Sociais e Suporte',
          'Contexto e Determinantes de Vida',
        ],
        pillarsEn: [
          'Mood',
          'Cognitive Function',
          'Vitality & Energy',
          'Stress, Trauma & Resilience',
          'Sexual Health',
          'Social Bonds & Support',
          'Life Context & Determinants',
        ],
      },
    ],
  },
  {
    code: 'R',
    codeEn: 'S',
    name: 'Ritmo Circadiano e Repouso',
    nameEn: 'Sleep, Rhythm & Recovery',
    pillarCount: 5,
    itemCount: 71,
    groups: [
      {
        pillars: [
          'Qualidade e Arquitetura do Sono',
          'Cronobiologia e Ritmo',
          'Exposição à Luz',
          'Higiene e Ambiente do Sono',
          'Distúrbios do Sono',
        ],
        pillarsEn: [
          'Sleep Quality & Architecture',
          'Chronobiology & Rhythm',
          'Light Exposure',
          'Sleep Hygiene & Environment',
          'Sleep Disorders',
        ],
      },
    ],
  },
];

export type AgirLetter = {
  code: string;
  name: string;
  pillarCount: number;
  itemCount: number;
  groups: { label?: string; pillars: string[] }[];
};

/** Devolve as letras AGIR/ACTS já localizadas. */
export function getAgirLetters(locale: string): AgirLetter[] {
  const isEn = locale === 'en';
  return agirLettersRaw.map((l) => ({
    code: isEn ? l.codeEn : l.code,
    name: isEn ? l.nameEn : l.name,
    pillarCount: l.pillarCount,
    itemCount: l.itemCount,
    groups: l.groups.map((g) => ({
      label: isEn ? g.labelEn : g.label,
      pillars: isEn ? g.pillarsEn : g.pillars,
    })),
  }));
}

/** @deprecated Use `getAgirLetters(locale)` para suporte i18n. */
export const agirLetters = agirLettersRaw.map((l) => ({
  code: l.code,
  name: l.name,
  pillarCount: l.pillarCount,
  itemCount: l.itemCount,
  groups: l.groups.map((g) => ({ label: g.label, pillars: g.pillars })),
}));

export const agirTotals = {
  letters: agirLettersRaw.length,
  pillars: agirLettersRaw.reduce((sum, l) => sum + l.pillarCount, 0),
  items: agirLettersRaw.reduce((sum, l) => sum + l.itemCount, 0),
};
