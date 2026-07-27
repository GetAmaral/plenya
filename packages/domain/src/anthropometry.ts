/**
 * Métricas antropométricas derivadas.
 *
 * Vários itens de `Composição corporal > Medidas Objetivas` não são medidos: saem de outros
 * itens do mesmo bloco (IMC vem de peso e altura, ASMI de massa apendicular e altura, e assim
 * por diante). Aqui ficam essas fórmulas, uma única vez, keyed pelo `anamnese_item_code` do
 * item derivado — o EMR preenche o campo sozinho conforme o profissional digita as medidas.
 *
 * Regras do registro:
 *  - `inputs` lista os códigos de entrada. Cada entrada pode ser uma ALTERNATIVA (array de
 *    códigos): usa-se o primeiro com valor. Serve para os pares homem/mulher, já que o filtro
 *    demográfico só deixa um deles na tela.
 *  - `compute` só é chamado com todas as entradas presentes e finitas.
 *  - o valor derivado é uma sugestão: digitação manual sobrepõe (ver renderer).
 */

/** Código de entrada, ou lista de códigos alternativos (o primeiro preenchido vence). */
export type DerivedInput = string | string[];

export interface DerivedMetric {
  /** anamnese_item_code do item derivado. */
  code: string;
  inputs: DerivedInput[];
  /** Recebe os valores na ordem de `inputs`. Retorna undefined se não fizer sentido. */
  compute: (values: number[]) => number | undefined;
  /** Casas decimais para arredondar o resultado. */
  decimals: number;
}

// --- códigos de entrada reutilizados ---------------------------------------
const PESO = 'PESO';
const ALTURA = 'ALTURA';
const QUADRIL = 'QUADRIL';
/** Cintura: só um dos dois aparece por paciente (filtro de gênero). */
const CINTURA: string[] = ['ABDOMINAL_HOMEM', 'ABDOMINAL_MULHER'];
const PESCOCO: string[] = ['PESCOCO_HOMEM', 'PESCOCO_MULHER'];
const MASSA_GORDA = 'MASSA_GORDA_TOTAL';
const MME = 'MASSA_MUSCULAR_ESQUELETICA';
const MASSA_APENDICULAR = 'MASSA_APENDICULAR';
const ACT = 'AGUA_CORPORAL_TOTAL';
const AEC = 'AGUA_EXTRACELULAR';

/** Altura em metros a partir da altura em cm. */
const metros = (alturaCm: number) => alturaCm / 100;

/** Índice de massa X: massa (kg) / altura² (m²). Base do IMC, FMI, índice MME e ASMI. */
const indiceDeMassa = (massaKg: number, alturaCm: number): number | undefined => {
  const m = metros(alturaCm);
  if (m <= 0) return undefined;
  return massaKg / (m * m);
};

/** Percentual de uma parte sobre um total. */
const percentual = (parte: number, total: number): number | undefined =>
  total > 0 ? (parte / total) * 100 : undefined;

/** Razão simples. */
const razao = (a: number, b: number): number | undefined => (b > 0 ? a / b : undefined);

const METRICS: DerivedMetric[] = [
  {
    code: 'IMC',
    inputs: [PESO, ALTURA],
    compute: ([peso, altura]) => indiceDeMassa(peso, altura),
    decimals: 1,
  },
  {
    // Body Roundness Index — Thomas DM et al., Obesity 2013;21(11):2264-71.
    // BRI = 364,2 − 365,5 · √(1 − ((cintura / 2π)² / (0,5 · altura)²))
    code: 'BRI',
    inputs: [CINTURA, ALTURA],
    compute: ([cintura, altura]) => {
      const raioCintura = cintura / (2 * Math.PI);
      const semiAltura = 0.5 * altura;
      if (semiAltura <= 0) return undefined;
      const excentricidade = 1 - (raioCintura / semiAltura) ** 2;
      if (excentricidade < 0) return undefined; // cintura maior que a altura: fora do domínio
      return 364.2 - 365.5 * Math.sqrt(excentricidade);
    },
    decimals: 2,
  },
  {
    code: 'RAZAO_CINTURA_ALTURA',
    inputs: [CINTURA, ALTURA],
    compute: ([cintura, altura]) => razao(cintura, altura),
    decimals: 2,
  },
  {
    code: 'RAZAO_CINTURA_QUADRIL_HOMEM',
    inputs: ['ABDOMINAL_HOMEM', QUADRIL],
    compute: ([cintura, quadril]) => razao(cintura, quadril),
    decimals: 2,
  },
  {
    code: 'RAZAO_CINTURA_QUADRIL_MULHER',
    inputs: ['ABDOMINAL_MULHER', QUADRIL],
    compute: ([cintura, quadril]) => razao(cintura, quadril),
    decimals: 2,
  },
  {
    // cm/m: pescoço em cm sobre a altura em METROS.
    code: 'RELACAO_PESCOCO_ALTURA_HOMEM',
    inputs: ['PESCOCO_HOMEM', ALTURA],
    compute: ([pescoco, altura]) => razao(pescoco, metros(altura)),
    decimals: 1,
  },
  {
    code: 'RELACAO_PESCOCO_ALTURA_MULHER',
    inputs: ['PESCOCO_MULHER', ALTURA],
    compute: ([pescoco, altura]) => razao(pescoco, metros(altura)),
    decimals: 1,
  },
  {
    code: 'GORDURA_CORPORAL_HOMEM',
    inputs: [MASSA_GORDA, PESO],
    compute: ([gordura, peso]) => percentual(gordura, peso),
    decimals: 1,
  },
  {
    code: 'GORDURA_CORPORAL_MULHER',
    inputs: [MASSA_GORDA, PESO],
    compute: ([gordura, peso]) => percentual(gordura, peso),
    decimals: 1,
  },
  {
    code: 'FMI_FAT_MASS_INDEX_HOMEM',
    inputs: [MASSA_GORDA, ALTURA],
    compute: ([gordura, altura]) => indiceDeMassa(gordura, altura),
    decimals: 1,
  },
  {
    code: 'FMI_FAT_MASS_INDEX_MULHER',
    inputs: [MASSA_GORDA, ALTURA],
    compute: ([gordura, altura]) => indiceDeMassa(gordura, altura),
    decimals: 1,
  },
  {
    code: 'MME_PESO',
    inputs: [MME, PESO],
    compute: ([mme, peso]) => percentual(mme, peso),
    decimals: 1,
  },
  {
    code: 'INDICE_MME',
    inputs: [MME, ALTURA],
    compute: ([mme, altura]) => indiceDeMassa(mme, altura),
    decimals: 2,
  },
  {
    code: 'ASMI_HOMEM',
    inputs: [MASSA_APENDICULAR, ALTURA],
    compute: ([apendicular, altura]) => indiceDeMassa(apendicular, altura),
    decimals: 2,
  },
  {
    code: 'ASMI_MULHER',
    inputs: [MASSA_APENDICULAR, ALTURA],
    compute: ([apendicular, altura]) => indiceDeMassa(apendicular, altura),
    decimals: 2,
  },
  {
    code: 'AGUA_CORPORAL_TOTAL_HOMEM',
    inputs: [ACT, PESO],
    compute: ([act, peso]) => percentual(act, peso),
    decimals: 1,
  },
  {
    code: 'AGUA_CORPORAL_TOTAL_MULHER',
    inputs: [ACT, PESO],
    compute: ([act, peso]) => percentual(act, peso),
    decimals: 1,
  },
  {
    code: 'RAZAO_AEC_ACT',
    inputs: [AEC, ACT],
    compute: ([aec, act]) => percentual(aec, act),
    decimals: 1,
  },
];

/** Registro keyed pelo código do item derivado. */
export const DERIVED_METRICS: Record<string, DerivedMetric> = Object.fromEntries(
  METRICS.map((m) => [m.code, m]),
);

export function getDerivedMetric(code?: string | null): DerivedMetric | undefined {
  return code ? DERIVED_METRICS[code] : undefined;
}

/** Todos os códigos que alimentam algum derivado (para saber o que dispara recálculo). */
export const DERIVED_INPUT_CODES: Set<string> = new Set(
  METRICS.flatMap((m) => m.inputs.flatMap((i) => (Array.isArray(i) ? i : [i]))),
);

/**
 * Calcula um derivado a partir de um mapa `código → valor medido`.
 * Retorna undefined quando falta alguma entrada ou a fórmula não se aplica.
 */
export function computeDerived(
  metric: DerivedMetric,
  measured: Record<string, number | undefined>,
): number | undefined {
  const values: number[] = [];
  for (const input of metric.inputs) {
    const codes = Array.isArray(input) ? input : [input];
    const found = codes.map((c) => measured[c]).find((v) => v !== undefined && Number.isFinite(v));
    if (found === undefined) return undefined;
    values.push(found);
  }
  const raw = metric.compute(values);
  if (raw === undefined || !Number.isFinite(raw)) return undefined;
  const factor = 10 ** metric.decimals;
  return Math.round(raw * factor) / factor;
}
