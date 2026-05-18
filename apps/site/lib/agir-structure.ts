/**
 * Estrutura completa do Método AGIR (PT) ↔ The ACTS Method (EN).
 * Usada pelas páginas /escore-plenya e /metodo-agir. Reflete a hierarquia do EMR:
 * Method → MethodLetter → MethodPillar (com agrupamento visual quando houver muitos pilares).
 *
 * Mapping de letras: A→A, G→C, I→T, R→S.
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
  itemCount: number; // estimativa agregada
  /** Quando há muitos pilares (>6), agrupar visualmente. */
  groups: { label?: string; labelEn?: string; pillars: string[]; pillarsEn: string[] }[];
};

export const agirLettersRaw: AgirLetterRaw[] = [
  {
    code: 'A',
    codeEn: 'A',
    name: 'Atividade Física, Alimentação e Suplementação Inteligente',
    nameEn: 'Activity, Alimentation & Smart Adjuncts',
    pillarCount: 4,
    itemCount: 120,
    groups: [
      {
        pillars: [
          'Avaliação Nutricional',
          'Prescrição de Exercícios',
          'Composição Corporal',
          'Suplementação',
        ],
        pillarsEn: [
          'Nutritional Assessment',
          'Exercise Prescription',
          'Body Composition',
          'Supplementation',
        ],
      },
    ],
  },
  {
    code: 'G',
    codeEn: 'C',
    name: 'Gestão Clínica e Metabólica',
    nameEn: 'Clinical Optimization',
    pillarCount: 14,
    itemCount: 672,
    groups: [
      {
        pillars: [
          'Cardiovascular',
          'Gastrointestinal',
          'Genético',
          'Hematológico',
          'Hepático',
          'Hormonal',
          'Imune e Inflamatório',
          'Metabólico',
          'Neurológico',
          'Nutrologia e Micronutrientes',
          'Osteomuscular',
          'Pulmonar',
          'Rastreio Oncológico',
          'Renal',
        ],
        pillarsEn: [
          'Cancer Screening',
          'Cardiovascular',
          'Gastrointestinal',
          'Genetic',
          'Hematological',
          'Hepatic',
          'Hormonal',
          'Immune & Inflammatory',
          'Metabolic',
          'Musculoskeletal',
          'Neurological',
          'Nutritional Medicine & Micronutrients',
          'Pulmonary',
          'Renal',
        ],
      },
    ],
  },
  {
    code: 'I',
    codeEn: 'T',
    name: 'Integração Mente-Corpo',
    nameEn: 'Tending Mind, Body & Bonds',
    pillarCount: 5,
    itemCount: 140,
    groups: [
      {
        pillars: [
          'Avaliação Psicológica',
          'Técnicas de Relaxamento',
          'Função Cognitiva',
          'Vida Sexual',
          'Vínculos Sociais e Suporte',
        ],
        pillarsEn: [
          'Psychological Assessment',
          'Relaxation Techniques',
          'Cognitive Function',
          'Sexual Health',
          'Social Bonds & Support',
        ],
      },
    ],
  },
  {
    code: 'R',
    codeEn: 'S',
    name: 'Ritmo Circadiano e Repouso',
    nameEn: 'Sleep, Rhythm & Recovery',
    pillarCount: 3,
    itemCount: 50,
    groups: [
      {
        pillars: ['Qualidade do Sono', 'Cronobiologia', 'Exposição à Luz'],
        pillarsEn: ['Sleep Quality', 'Chronobiology', 'Light Exposure'],
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
