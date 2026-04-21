/**
 * Estrutura completa do Método AGIR — usada pelas páginas /escore-plenya
 * e /metodo-agir. Reflete a hierarquia do EMR: Method → MethodLetter →
 * MethodPillar (com agrupamento visual quando houver muitos pilares).
 *
 * Inlined em apps/site para evitar problemas de cache do Turbopack
 * com export indireto de @plenya/brand.
 */

export type AgirLetter = {
  code: 'A' | 'G' | 'I' | 'R';
  name: string;
  pillarCount: number;
  itemCount: number; // estimativa agregada
  /** Quando há muitos pilares (>6), agrupar visualmente. */
  groups: { label?: string; pillars: string[] }[];
};

export const agirLetters: AgirLetter[] = [
  {
    code: 'A',
    name: 'Atividade Física, Alimentação e Suplementação Inteligente',
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
      },
    ],
  },
  {
    code: 'G',
    name: 'Gestão Clínica e Metabólica',
    pillarCount: 10,
    itemCount: 480,
    groups: [
      {
        label: 'Sistemas',
        pillars: [
          'Controle Glicêmico',
          'Perfil Lipídico',
          'Função Hepática',
          'Função Renal',
          'Risco Cardiovascular',
        ],
      },
      {
        label: 'Painéis bioquímicos',
        pillars: [
          'Painel Hormonal',
          'Inflamação e Imunidade',
          'Vitaminas, Minerais e Micronutrientes',
          'Hematologia',
          'Rastreamento Oncológico',
        ],
      },
    ],
  },
  {
    code: 'I',
    name: 'Integração Mente-Corpo',
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
      },
    ],
  },
  {
    code: 'R',
    name: 'Ritmo Circadiano e Repouso',
    pillarCount: 3,
    itemCount: 50,
    groups: [
      {
        pillars: ['Qualidade do Sono', 'Cronobiologia', 'Exposição à Luz'],
      },
    ],
  },
];

export const agirTotals = {
  letters: agirLetters.length,
  pillars: agirLetters.reduce((sum, l) => sum + l.pillarCount, 0),
  items: agirLetters.reduce((sum, l) => sum + l.itemCount, 0),
};
