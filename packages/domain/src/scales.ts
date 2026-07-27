/**
 * Registro compartilhado de escalas clínicas (instrumentos somados).
 *
 * Fonte única consumida pelo EMR (apps/web) e pelo site. Aqui ficam só os dados que NÃO dão
 * para derivar do banco: perguntas, opções de resposta e a fórmula de score. A classificação
 * (total → nível) continua vindo dos `levels` do score item no banco (via `matchLevel` deste
 * pacote, ou o `detectLevel` do EMR — mesma semântica de operador/limite).
 *
 * Chave do registro = `anamnese_item_code` do score item.
 *
 * Mecânicas (`kind`):
 *  - `sum`          : Σ das respostas. Tier A (PHQ-9, GAD-7, Epworth, IIEF-5).
 *  - `administered` : teste aplicado ao vivo (Dubois, Span). Tier B.
 *  - `custom`       : fórmula própria (FSS média, FSFI domínios, PSQI componentes). Tier C.
 */

export type ScaleKind = 'sum' | 'administered' | 'custom';

export interface ScaleOption {
  value: number;
  label: string;
}

export interface ScaleQuestion {
  text: string;
  /** Opções desta pergunta. Se omitido, usa `defaultOptions` da escala. */
  options?: ScaleOption[];
}

/** Uma palavra escolhida (categoria + termo) numa aplicação do teste de evocação. */
export interface ChosenWord {
  category: string;
  word: string;
}

export type AdministrationSpec =
  | {
      type: 'word_recall';
      /** Pools por categoria: a cada aplicação sorteia-se 1 palavra/categoria (formas
       *  alternativas reduzem o efeito de aprendizagem na reaplicação no mesmo paciente). */
      categories: { category: string; words: string[] }[];
    }
  | { type: 'digit_span'; direction: 'forward' | 'backward'; sequencesByLength: Record<number, number[][]> };

/** Sorteia 1 palavra por categoria (forma alternativa). A aleatoriedade fica na camada de UI;
 *  esta função recebe um seletor por índice para manter o domínio puro/testável. */
export function pickWordRecallSet(
  spec: Extract<AdministrationSpec, { type: 'word_recall' }>,
  pickIndex: (size: number, categoryIndex: number) => number,
): ChosenWord[] {
  return spec.categories.map((c, i) => ({
    category: c.category,
    word: c.words[Math.min(Math.max(pickIndex(c.words.length, i), 0), c.words.length - 1)] ?? c.words[0],
  }));
}

export interface ScaleDef {
  /** anamnese_item_code (chave do registro). */
  code: string;
  kind: ScaleKind;
  title: string;
  /** Enunciado mostrado acima das perguntas. */
  instructions?: string;
  /** Opções aplicadas às perguntas sem `options` próprias. */
  defaultOptions?: ScaleOption[];
  /** Perguntas (kind = 'sum'). Ausente em kind = 'administered'. */
  questions?: ScaleQuestion[];
  maxScore: number;
  /** Spec de aplicação guiada (kind = 'administered'). */
  administration?: AdministrationSpec;
  /** Score custom; default = soma das respostas. */
  score?: (answers: Record<number, number>) => number;
  /** Classificação custom a partir das respostas (não do total) — ex.: ASEX, cujo nível
   *  depende de regras por item. Retorna o `level` do score item. Quando ausente, o widget
   *  classifica o total pelos níveis do banco. */
  classifyAnswers?: (answers: Record<number, number>) => number | undefined;
}

/** Opções efetivas de uma pergunta (próprias ou o default da escala). */
export function questionOptions(def: ScaleDef, index: number): ScaleOption[] {
  return def.questions?.[index]?.options ?? def.defaultOptions ?? [];
}

/** Total a partir das respostas {índiceDaPergunta: valor}. Default = soma. */
export function scaleTotal(def: ScaleDef, answers: Record<number, number>): number {
  if (def.score) return def.score(answers);
  return Object.values(answers).reduce((sum, v) => sum + (v ?? 0), 0);
}

export function scaleAnsweredCount(answers: Record<number, number>): number {
  return Object.values(answers).filter((v) => v !== undefined && v !== null).length;
}

export function isScaleComplete(def: ScaleDef, answers: Record<number, number>): boolean {
  // Administered (Dubois/Span): completude é tratada no widget; aqui basta ter alguma resposta.
  if (def.kind === 'administered') return scaleAnsweredCount(answers) > 0;
  return (def.questions ?? []).every((_, i) => answers[i] !== undefined && answers[i] !== null);
}

/**
 * Formata o resultado de uma escala para LEITURA: `"22/30"` ou `"22/30 · Disfunção leve"`.
 * Usado por todas as telas que exibem uma escala já preenchida (chip da consulta, lista de
 * anamnese salva, histórico). Mantém uma única forma de mostrar total + classificação.
 */
export function formatScaleResult(total: number, maxScore: number, levelName?: string | null): string {
  const base = `${total}/${maxScore}`;
  return levelName ? `${base} · ${levelName}` : base;
}

export function getScaleDef(code?: string | null): ScaleDef | undefined {
  return code ? SCALE_REGISTRY[code] : undefined;
}

// ---------------------------------------------------------------------------
// Opções reutilizadas
// ---------------------------------------------------------------------------

/** Frequência PHQ-9 / GAD-7 (últimas 2 semanas). */
const FREQ_2W: ScaleOption[] = [
  { value: 0, label: 'Nenhuma vez' },
  { value: 1, label: 'Menos da metade dos dias' },
  { value: 2, label: 'Mais da metade dos dias' },
  { value: 3, label: 'Quase todos os dias' },
];

/** Chance de cochilar (Epworth). */
const DOZE_CHANCE: ScaleOption[] = [
  { value: 0, label: 'Nunca cochilaria' },
  { value: 1, label: 'Pequena chance' },
  { value: 2, label: 'Chance moderada' },
  { value: 3, label: 'Alta chance de cochilar' },
];

/** Frequência IIEF-5 (Q2/Q3/Q5), com "0 = sem atividade sexual". */
const IIEF_FREQ: ScaleOption[] = [
  { value: 0, label: 'Não tentei / sem atividade sexual' },
  { value: 1, label: 'Quase nunca ou nunca' },
  { value: 2, label: 'Poucas vezes (menos da metade)' },
  { value: 3, label: 'Às vezes (cerca da metade)' },
  { value: 4, label: 'Muitas vezes (mais da metade)' },
  { value: 5, label: 'Quase sempre ou sempre' },
];

// ---------------------------------------------------------------------------
// Tier A — escalas de soma simples
// ---------------------------------------------------------------------------

const PHQ9: ScaleDef = {
  code: 'ESCALA_PHQ_9_27',
  kind: 'sum',
  title: 'PHQ-9 (humor)',
  instructions:
    'Nas últimas 2 semanas, com que frequência você foi incomodado(a) por algum dos problemas abaixo?',
  defaultOptions: FREQ_2W,
  maxScore: 27,
  questions: [
    { text: 'Pouco interesse ou prazer em fazer as coisas' },
    { text: 'Sentir-se triste, deprimido(a) ou sem perspectiva' },
    { text: 'Dificuldade para pegar no sono ou continuar dormindo, ou dormir demais' },
    { text: 'Sentir-se cansado(a) ou com pouca energia' },
    { text: 'Falta de apetite ou comer demais' },
    {
      text: 'Sentir-se mal consigo mesmo(a), achar que é um fracasso ou que decepcionou a família ou a si mesmo(a)',
    },
    { text: 'Dificuldade para se concentrar nas coisas (como ler o jornal ou ver televisão)' },
    {
      text: 'Lentidão para se mover ou falar a ponto de outras pessoas perceberem; ou o oposto, estar tão agitado(a) que tem se mexido bem mais que o normal',
    },
    { text: 'Pensar em se ferir de alguma forma ou que seria melhor estar morto(a)' },
  ],
};

const GAD7: ScaleDef = {
  code: 'GAD_7_21',
  kind: 'sum',
  title: 'GAD-7 (ansiedade)',
  instructions:
    'Nas últimas 2 semanas, com que frequência você foi incomodado(a) pelos seguintes problemas?',
  defaultOptions: FREQ_2W,
  maxScore: 21,
  questions: [
    { text: 'Sentir-se nervoso(a), ansioso(a) ou muito tenso(a)' },
    { text: 'Não conseguir parar ou controlar as preocupações' },
    { text: 'Preocupar-se muito com diversas coisas' },
    { text: 'Dificuldade para relaxar' },
    { text: 'Ficar tão agitado(a) que se torna difícil permanecer sentado(a)' },
    { text: 'Ficar facilmente aborrecido(a) ou irritado(a)' },
    { text: 'Sentir medo como se algo terrível fosse acontecer' },
  ],
};

const EPWORTH: ScaleDef = {
  code: 'ESCALA_DE_SONOLENCIA_DE_EPWORTH_24',
  kind: 'sum',
  title: 'Escala de Sonolência de Epworth',
  instructions:
    'Qual a probabilidade de você cochilar ou adormecer nas situações abaixo (não apenas se sentir cansado[a]), considerando seu modo de vida nos últimos tempos?',
  defaultOptions: DOZE_CHANCE,
  maxScore: 24,
  questions: [
    { text: 'Sentado(a) e lendo' },
    { text: 'Assistindo à televisão' },
    { text: 'Sentado(a), quieto(a), em lugar público (sala de espera, cinema, teatro, reunião)' },
    { text: 'Como passageiro(a) de carro por uma hora, sem parar' },
    { text: 'Deitado(a) para descansar à tarde, quando as circunstâncias permitem' },
    { text: 'Sentado(a) conversando com alguém' },
    { text: 'Sentado(a) calmamente após o almoço, sem ter bebido álcool' },
    { text: 'No carro, parado(a) por alguns minutos no trânsito' },
  ],
};

const IIEF5: ScaleDef = {
  code: 'IIEF_5_25',
  kind: 'sum',
  title: 'IIEF-5 (função erétil)',
  instructions: 'Nos últimos 6 meses:',
  maxScore: 25,
  questions: [
    {
      text: 'Como você classifica sua confiança em conseguir e manter uma ereção?',
      options: [
        { value: 1, label: 'Muito baixa' },
        { value: 2, label: 'Baixa' },
        { value: 3, label: 'Moderada' },
        { value: 4, label: 'Alta' },
        { value: 5, label: 'Muito alta' },
      ],
    },
    {
      text: 'Quando você teve ereções com estímulo sexual, com que frequência elas foram suficientes para a penetração?',
      options: IIEF_FREQ,
    },
    {
      text: 'Durante a relação sexual, com que frequência você conseguiu manter a ereção após a penetração?',
      options: IIEF_FREQ,
    },
    {
      text: 'Durante a relação sexual, quão difícil foi manter a ereção até o final?',
      options: [
        { value: 1, label: 'Extremamente difícil' },
        { value: 2, label: 'Muito difícil' },
        { value: 3, label: 'Difícil' },
        { value: 4, label: 'Pouco difícil' },
        { value: 5, label: 'Sem dificuldade' },
      ],
    },
    {
      text: 'Quando você tentou a relação sexual, com que frequência ela foi satisfatória?',
      options: IIEF_FREQ,
    },
  ],
};

// ---------------------------------------------------------------------------
// Tier B — testes de memória/cognição (administração guiada)
// ---------------------------------------------------------------------------

// PROPOSTA (a confirmar com o Dr. Getúlio): teste das 5 palavras de Dubois — 5 categorias
// semânticas clássicas, cada uma com um POOL de palavras PT-BR concretas e inequívocas.
// A cada aplicação sorteia-se 1 palavra por categoria → formas alternativas para reaplicar
// no mesmo paciente sem efeito de aprendizagem. Imediato e tardio usam o MESMO sorteio
// dentro de uma anamnese (persistido em scale_responses.words).
const DUBOIS_CATEGORIES = [
  { category: 'Bebida', words: ['Limonada', 'Café', 'Suco', 'Chá', 'Vinho', 'Refrigerante'] },
  { category: 'Edifício', words: ['Museu', 'Igreja', 'Hospital', 'Castelo', 'Estádio', 'Teatro'] },
  { category: 'Inseto', words: ['Gafanhoto', 'Borboleta', 'Formiga', 'Abelha', 'Besouro', 'Mosquito'] },
  { category: 'Utensílio de cozinha', words: ['Peneira', 'Panela', 'Concha', 'Garfo', 'Colher', 'Espátula'] },
  { category: 'Veículo', words: ['Caminhão', 'Trem', 'Navio', 'Bicicleta', 'Avião', 'Ônibus'] },
];

// Score de evocação = SCORE TOTAL PONDERADO do teste dos 5 palavras, por fase (0-10).
// Cotação canônica (Dubois 2002; Cowppli-Bony 2005; Croisile 2010; ficha INESSS/HAS):
//   evocação espontânea (rappel libre)   = 2 pontos
//   evocação com dica   (rappel indicé)  = 1 ponto
//   não evocada nem com dica             = 0
// Imediato (/10) + tardio (/10) = Score Total Ponderado /20 — 19-20 declínio pouco provável,
// ≤18 investigar. O ponderado dobra o rappel libre e é MAIS discriminante que o "score total"
// antigo (/5 por fase), que dava o mesmo peso à evocação com dica — era o que fazíamos antes.
const wordRecallScore = (answers: Record<number, number>): number =>
  Object.values(answers).reduce((sum, v) => sum + (v ?? 0), 0);

// Score de span: maior comprimento com pelo menos uma tentativa correta (valor 1).
const digitSpanScore = (answers: Record<number, number>): number => {
  let max = 0;
  for (const [len, ok] of Object.entries(answers)) if (ok === 1) max = Math.max(max, Number(len));
  return max;
};

// Sequências de dígitos por comprimento (2 tentativas cada). Arbitrárias por desenho —
// editáveis. Direto: 3–8. Inverso: 2–7.
const SPAN_FORWARD: Record<number, number[][]> = {
  3: [[3, 8, 6], [6, 1, 2]],
  4: [[3, 4, 9, 7], [6, 1, 5, 8]],
  5: [[5, 2, 1, 8, 6], [3, 8, 6, 4, 1]],
  6: [[8, 9, 2, 6, 4, 7], [4, 1, 7, 9, 3, 8]],
  7: [[1, 7, 4, 2, 8, 6, 3], [9, 4, 1, 7, 5, 2, 8]],
  8: [[5, 8, 1, 9, 2, 6, 4, 7], [3, 9, 6, 2, 8, 1, 5, 4]],
};
const SPAN_BACKWARD: Record<number, number[][]> = {
  2: [[5, 8], [2, 6]],
  3: [[6, 2, 9], [4, 1, 5]],
  4: [[3, 2, 7, 9], [4, 9, 6, 8]],
  5: [[1, 5, 2, 8, 6], [6, 1, 8, 4, 3]],
  6: [[5, 3, 9, 4, 1, 8], [7, 2, 4, 8, 5, 6]],
  7: [[8, 1, 4, 9, 3, 6, 2], [4, 7, 3, 9, 1, 2, 8]],
};

const DUBOIS_IMEDIATO: ScaleDef = {
  code: '5_PALAVRAS_DE_DUBOIS_IMEDIATO_5',
  kind: 'administered',
  title: '5 palavras de Dubois — evocação imediata',
  instructions:
    'Apresente as 5 palavras com a categoria de cada uma. Logo após, peça a evocação livre e, para as não lembradas, ofereça a dica (categoria). Marque como cada palavra foi recuperada. Espontânea vale 2 pontos, com dica 1 (score ponderado /10).',
  maxScore: 10,
  administration: { type: 'word_recall', categories: DUBOIS_CATEGORIES },
  score: wordRecallScore,
};

const DUBOIS_TARDIO: ScaleDef = {
  code: '5_PALAVRAS_DE_DUBOIS_TARDIO_5',
  kind: 'administered',
  title: '5 palavras de Dubois — evocação tardia',
  instructions:
    'Após o intervalo/interferência, peça novamente a evocação das 5 palavras (livre e, se preciso, com dica). Marque como cada palavra foi recuperada. Espontânea vale 2 pontos, com dica 1 (score ponderado /10).',
  maxScore: 10,
  administration: { type: 'word_recall', categories: DUBOIS_CATEGORIES },
  score: wordRecallScore,
};

const SPAN_DIRETO: ScaleDef = {
  code: 'SPAN_DE_DIGITOS_DIRETO_8',
  kind: 'administered',
  title: 'Span de dígitos — ordem direta',
  instructions:
    'Leia cada sequência em voz alta (1 dígito por segundo) e peça para repetir na mesma ordem. Marque a maior sequência repetida corretamente.',
  maxScore: 8,
  administration: { type: 'digit_span', direction: 'forward', sequencesByLength: SPAN_FORWARD },
  score: digitSpanScore,
};

const SPAN_INVERSO: ScaleDef = {
  code: 'SPAN_DE_DIGITOS_INVERSO_7',
  kind: 'administered',
  title: 'Span de dígitos — ordem inversa',
  instructions:
    'Leia cada sequência em voz alta e peça para repetir na ORDEM INVERSA. Marque a maior sequência repetida corretamente.',
  maxScore: 7,
  administration: { type: 'digit_span', direction: 'backward', sequencesByLength: SPAN_BACKWARD },
  score: digitSpanScore,
};

// ---------------------------------------------------------------------------
// Tier C — ASEX (Arizona Sexual Experience Scale)
// ---------------------------------------------------------------------------
// 5 itens, cada um 1–6 (1 = melhor função, 6 = pior). Total 5–30 (NÃO 0–25).
// Fonte: McGahuey et al., 2000 (J Sex Marital Ther) — Arizona Sexual Experience Scale,
// validação original. Rastreio positivo para disfunção sexual se total ≥19, OU qualquer
// item ≥5, OU ≥3 itens ≥4. A classificação NÃO é pelo total puro: usa essa regra.
// Mapeamento de níveis (banco): L0 total>18 · L1 ≥3 itens≥4 · L2 1 item≥5 · L5 normal.

const ASEX_EASE: ScaleOption[] = [
  { value: 1, label: 'Extremamente fácil' },
  { value: 2, label: 'Muito fácil' },
  { value: 3, label: 'Moderadamente fácil' },
  { value: 4, label: 'Difícil' },
  { value: 5, label: 'Muito difícil' },
  { value: 6, label: 'Nunca' },
];

const ASEX: ScaleDef = {
  code: 'ASEX_25',
  kind: 'custom',
  title: 'ASEX (função sexual)',
  instructions: 'Considerando a sua experiência sexual na última semana:',
  maxScore: 30,
  questions: [
    {
      text: 'Como está o seu desejo sexual (vontade de ter atividade sexual)?',
      options: [
        { value: 1, label: 'Extremamente forte' },
        { value: 2, label: 'Muito forte' },
        { value: 3, label: 'Moderado' },
        { value: 4, label: 'Pouco' },
        { value: 5, label: 'Muito pouco' },
        { value: 6, label: 'Nenhum' },
      ],
    },
    { text: 'Com que facilidade você fica sexualmente excitado(a)?', options: ASEX_EASE },
    {
      text: 'Sua resposta física acontece com que facilidade (lubrificação vaginal, nas mulheres; ereção, nos homens)?',
      options: ASEX_EASE,
    },
    { text: 'Com que facilidade você atinge o orgasmo?', options: ASEX_EASE },
    {
      text: 'Quão satisfatórios são os seus orgasmos?',
      options: [
        { value: 1, label: 'Extremamente satisfatórios' },
        { value: 2, label: 'Muito satisfatórios' },
        { value: 3, label: 'Moderadamente satisfatórios' },
        { value: 4, label: 'Pouco satisfatórios' },
        { value: 5, label: 'Muito pouco satisfatórios' },
        { value: 6, label: 'Não consigo ter orgasmo' },
      ],
    },
  ],
  score: (answers) => Object.values(answers).reduce((s, v) => s + (v ?? 0), 0),
  classifyAnswers: (answers) => {
    const vals = Object.values(answers);
    const total = vals.reduce((s, v) => s + (v ?? 0), 0);
    const anyGE5 = vals.some((v) => v >= 5);
    const countGE4 = vals.filter((v) => v >= 4).length;
    if (total > 18) return 0; // L0
    if (countGE4 >= 3) return 1; // L1
    if (anyGE5) return 2; // L2
    return 5; // L5 — sem indícios de disfunção
  },
};

export const SCALE_REGISTRY: Record<string, ScaleDef> = {
  [PHQ9.code]: PHQ9,
  [GAD7.code]: GAD7,
  [EPWORTH.code]: EPWORTH,
  [IIEF5.code]: IIEF5,
  [DUBOIS_IMEDIATO.code]: DUBOIS_IMEDIATO,
  [DUBOIS_TARDIO.code]: DUBOIS_TARDIO,
  [SPAN_DIRETO.code]: SPAN_DIRETO,
  [SPAN_INVERSO.code]: SPAN_INVERSO,
  [ASEX.code]: ASEX,
};
