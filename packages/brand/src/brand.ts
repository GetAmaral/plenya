export const brand = {
  name: 'Plenya',
  legalName: 'Plenya Saúde',
  tagline: 'Saúde, Performance & Longevidade',
  claim: 'Viva bem, viva mais.',
  domain: 'plenyasaude.com.br',
  url: 'https://plenyasaude.com.br',
  appUrl: 'https://app.plenyasaude.com.br',
  apiUrl: 'https://api.plenyasaude.com.br',
  email: 'contato@plenyasaude.com.br',
  social: {
    instagram: 'https://instagram.com/plenyaSaude',
  },
} as const;

/**
 * Brand essence — pulled verbatim from the official Plenya brandbook.
 * Source of truth for all institutional copy on the site.
 */
export const brandEssence = {
  purpose: 'Guiar pessoas a viverem mais e melhor, com autonomia sobre a própria saúde.',
  positioning:
    'Integramos ciência, dados e acompanhamento próximo para transformar o cuidado em um processo contínuo, personalizado e mensurável.',
  symbol:
    'P + ∞ — a união entre as letras P, associada ao símbolo do infinito, traduz a continuidade do cuidado, a integridade e o ciclo de evolução contínua. Vida em corpo e mente, processo sem fim.',
  goldenCircle: {
    why: 'Porque as pessoas estão vivendo mais, mas não necessariamente melhor — e seguem sendo tratadas em partes, não como um todo. Ninguém deveria chegar à fase mais importante da vida limitado pelas escolhas que fez antes.',
    how: 'Através de um acompanhamento integrado, contínuo e personalizado, que conecta diferentes áreas da saúde em um único plano estruturado, com dados, proximidade e responsabilidade compartilhada.',
    what: 'Um programa de saúde e longevidade que combina equipe multidisciplinar, método próprio (AGIR) e acompanhamento próximo para transformar hábitos, corpo e mente de forma sustentável.',
  },
} as const;

export const manifesto = [
  'Tudo está conectado.',
  'Corpo, mente, tempo e escolhas.',
  'Viver bem é o reflexo do que você escolhe todos os dias.',
  'É consciência. É liberdade.',
  'Plenitude não é um ponto de chegada.',
  'É uma linha contínua.',
  'Plenya. Viva bem, viva mais.',
] as const;

/**
 * Tom de voz — princípios oficiais de comunicação da marca.
 */
export const toneOfVoice = [
  {
    title: 'Clareza',
    body: 'A Plenya traduz complexidade em direção. Evita excesso de termos técnicos, mas também não simplifica a ponto de perder profundidade.',
  },
  {
    title: 'Autoridade',
    body: 'A marca se posiciona como especialista, mas sem distanciamento. Fala com segurança, embasamento e consistência, sem soar impositiva ou inacessível.',
  },
  {
    title: 'Proximidade',
    body: 'Existe acolhimento, mas não informalidade excessiva. A Plenya entende o momento do paciente, mas não banaliza o tema saúde. O tom é humano, mas sempre comprometido com seriedade e consistência.',
  },
  {
    title: 'Constância',
    body: 'A marca não se apoia em promessas rápidas ou soluções imediatas. Sua comunicação reforça processo e continuidade, evitando sensacionalismo ou atalhos.',
  },
  {
    title: 'Inspiração',
    body: 'A Plenya inspira mudança, mas sempre conectada à realidade. Evita frases vazias ou motivacionais genéricas. Toda inspiração parte de um princípio concreto.',
  },
] as const;

/**
 * Arquétipos de relação — os três círculos de pessoas com quem a marca se conecta.
 */
export const relationArchetypes = [
  {
    title: 'Os Pacientes Protagonistas',
    body:
      'O paciente é o centro de origem de tudo. É a partir das suas escolhas que o cuidado ganha forma. Vive um momento de movimento e momento em que algo precisa de mudança, mais consciência sobre o próprio corpo, mente e ciência.',
  },
  {
    title: 'Os Familiares Catalisadores',
    body:
      'A transformação individual não acontece no isolamento. Família, parceiros e entorno influenciam consistência. Eles são parte do tecido que sustenta o cuidado.',
  },
  {
    title: 'Time Plenya — Guardiões do Método',
    body:
      'O método só acontece na medida em que a união, dão profundidade ao cuidado. São o time multidisciplinar que dá forma à Plenya em cada vida que toca.',
  },
] as const;

/**
 * Pilares do Método AGIR (PT) ↔ The ACTS Method (EN).
 *
 * Mapping:
 *   A → A — Atividade Física, Alimentação e Suplementação Inteligente
 *           ↔ Activity, Alimentation & Smart Adjuncts
 *   G → C — Gestão Clínica e Metabólica
 *           ↔ Clinical Optimization
 *   I → T — Integração Mente-Corpo
 *           ↔ Tending Mind, Body & Bonds
 *   R → S — Ritmo Circadiano e Repouso
 *           ↔ Sleep, Rhythm & Recovery
 *
 * Slug é estável (em PT) — usado como chave de roteamento no banco e nos
 * filtros de blog. Os campos textuais são localizados via `getAgirPillars(locale)`.
 *
 * Fonte de verdade: TRANSLATION-GUIDE.md + glossario.yaml do livro ANTES.
 */
export const agirPillars = [
  {
    code: 'A',
    codeEn: 'A',
    slug: 'alimentacao-atividade-fisica',
    name: 'Atividade Física, Alimentação e Suplementação Inteligente',
    nameEn: 'Activity, Alimentation & Smart Adjuncts',
    nameEs: 'Actividad Física, Alimentación y Suplementación Inteligente',
    idea: 'O corpo responde ao que você sustenta, não ao que você tenta ocasionalmente.',
    ideaEn: 'The body responds to what you sustain — not to what you attempt occasionally.',
    territory: 'Rotina, disciplina e construção de hábito.',
    territoryEn: 'Routine, discipline, and habit formation.',
    keyMessages: [
      'Não é sobre intensidade. É sobre constância.',
      'Seu corpo reflete o que você faz todos os dias.',
      'Resultado não vem de esforço isolado.',
      'Movimento e alimentação são base de funcionamento.',
    ],
    keyMessagesEn: [
      'It is not about intensity. It is about consistency.',
      'Your body reflects what you do every day.',
      'Results do not come from isolated effort.',
      'Movement and food are the baseline of how you function.',
    ],
  },
  {
    code: 'G',
    codeEn: 'C',
    slug: 'gestao-metabolica',
    name: 'Gestão Clínica e Metabólica',
    nameEn: 'Clinical Optimization',
    nameEs: 'Gestión Clínica y Metabólica',
    idea: 'Você não melhora o que não acompanha.',
    ideaEn: 'You cannot improve what you do not track.',
    territory: 'Dados, consciência, controle.',
    territoryEn: 'Data, awareness, and control.',
    keyMessages: [
      'Saúde também é dado.',
      'O que não é monitorado, não evolui.',
      'Clareza gera direção e direção gera resultado.',
      'Seu corpo fala, e os dados ajudam a entender.',
    ],
    keyMessagesEn: [
      'Health is also data.',
      'What is not monitored does not evolve.',
      'Clarity creates direction, and direction creates results.',
      'Your body speaks, and the data helps you understand.',
    ],
  },
  {
    code: 'I',
    codeEn: 'T',
    slug: 'integracao-corpo-mente',
    name: 'Integração Mente-Corpo',
    nameEn: 'Tending Mind, Body & Bonds',
    nameEs: 'Integración Mente-Cuerpo',
    idea: 'Sem alinhamento interno, não existe constância externa.',
    ideaEn: 'Without inner alignment, there is no outer consistency.',
    territory: 'Consciência, comportamento, profundidade.',
    territoryEn: 'Awareness, behavior, depth.',
    keyMessages: [
      'Saúde mental é estrutura.',
      'Comportamento sustenta resultado.',
      'A mudança real começa na forma como você pensa e decide.',
      'Não existe disciplina sem equilíbrio emocional.',
    ],
    keyMessagesEn: [
      'Mental health is structure.',
      'Behavior sustains the result.',
      'Real change begins in the way you think and decide.',
      'There is no discipline without emotional balance.',
    ],
  },
  {
    code: 'R',
    codeEn: 'S',
    slug: 'ritmo-circadiano',
    name: 'Ritmo Circadiano e Repouso',
    nameEn: 'Sleep, Rhythm & Recovery',
    nameEs: 'Ritmo Circadiano y Descanso',
    idea: 'Recuperação é parte do progresso.',
    ideaEn: 'Recovery is part of progress.',
    territory: 'Energia, equilíbrio, sustentabilidade.',
    territoryEn: 'Energy, balance, sustainability.',
    keyMessages: [
      'Descanso também é estratégia.',
      'Seu corpo precisa de ritmo para funcionar bem.',
      'Energia não se cria, se regula.',
      'Dormir bem é inegociável.',
    ],
    keyMessagesEn: [
      'Rest is also strategy.',
      'Your body needs rhythm to function well.',
      'Energy is not created — it is regulated.',
      'Sleeping well is non-negotiable.',
    ],
  },
] as const;

/**
 * Devolve os pilares já localizados, com a forma curta do nome ("A — Activity, …").
 * Use sempre este helper em componentes — não acesse `agirPillars[i].nameEn` direto.
 */
export type AgirLocalizedPillar = {
  code: string;
  slug: string;
  name: string;
  idea: string;
  territory: string;
  keyMessages: readonly string[];
};

export function getAgirPillars(locale: string): AgirLocalizedPillar[] {
  const isEn = locale === 'en';
  return agirPillars.map((p) => ({
    code: isEn ? p.codeEn : p.code,
    slug: p.slug,
    name: isEn ? p.nameEn : p.name,
    idea: isEn ? p.ideaEn : p.idea,
    territory: isEn ? p.territoryEn : p.territory,
    keyMessages: isEn ? p.keyMessagesEn : p.keyMessages,
  }));
}

/**
 * Público-alvo da marca — referência para curadoria de copy e seleção de imagens.
 */
export const audience = {
  ageRange: '35–55 anos',
  description:
    'Homens e mulheres em um momento de vida marcado por alta demanda profissional, responsabilidades familiares e uma crescente preocupação com saúde, energia e longevidade.',
  pains: [
    'Sensação de cansaço constante e queda de energia',
    'Dificuldade em manter constância em hábitos saudáveis',
    'Acompanhamento médico confuso e desconectado',
    'Medo de repetir padrões de saúde da família no futuro',
  ],
  desires: [
    'Ter mais energia e disposição no dia a dia',
    'Sentir controle e clareza sobre a própria saúde',
    'Viver com qualidade ao longo do tempo',
    'Prevenir problemas antes que eles apareçam',
  ],
  insight: 'Mais do que saúde, esse público busca segurança, autonomia e tranquilidade para viver bem todas as fases da vida.',
} as const;
