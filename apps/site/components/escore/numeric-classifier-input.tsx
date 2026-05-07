/**
 * Input numérico com classificação automática em tempo real.
 * Para itens do Light que têm levels com faixas numéricas (lower/upper limits):
 * - Composição: Cintura, Pescoço, Prancha (segundos)
 * - Bioimpedância: ASMI, Gordura visceral, % Gordura
 * - Todos os labs (ApoB, LDL, HDL, TSH, ferritina, etc.)
 *
 * Usuário digita o valor → componente classifica → exibe a faixa correspondente.
 * O valor numérico é enviado ao backend, que confirma o cálculo via EvaluatesTrue.
 */
'use client';

import { useTranslations } from 'next-intl';
import type { LightItemConfig, LightLevelConfig } from '@/lib/score-light/types';

// Items de medida que precisam de input numérico mesmo quando o campo `unit` no DB vem vazio.
const FORCE_NUMERIC_PATTERNS: RegExp[] = [
  /Abdominal \(cintura/i,
  /^Pescoço - /i,
  /ASMI \(kg\/m/i,
  /Gordura visceral \(cm/i,
  /% Gordura corporal/i,
  /^IMC \(kg\/m/i,
  /Razão cintura\/altura/i,
];

/** Detecta se um item deve usar input numérico com classificação. */
export function isNumericClassifiable(item: LightItemConfig): boolean {
  if (!item.levels || item.levels.length === 0) return false;
  const hasRanges = item.levels.some(
    (lv) => (lv.lowerLimit && lv.lowerLimit !== '') || (lv.upperLimit && lv.upperLimit !== ''),
  );
  if (!hasRanges) return false;

  // Com unit explícito, sempre numérico
  if (item.unit && item.unit.trim() !== '') return true;

  // Medidas sem unit no DB mas que são numéricas (fallback por nome)
  return FORCE_NUMERIC_PATTERNS.some((re) => re.test(item.name));
}

/** Placeholder tailored por item/unit. */
export function placeholderFor(item: LightItemConfig): string {
  const name = item.name;

  // Matches específicos por nome (quando unit no DB não é descritivo)
  if (/Abdominal \(cintura/i.test(name)) return 'Cintura em centímetros (cm)';
  if (/^Pescoço - /i.test(name)) return 'Pescoço em centímetros (cm)';
  if (/ASMI \(kg\/m/i.test(name)) return 'ASMI em kg/m²';
  if (/Gordura visceral/i.test(name)) return 'Gordura visceral em cm²';
  if (/% Gordura corporal/i.test(name)) return 'Percentual de gordura (%)';
  if (/^IMC \(kg\/m/i.test(name)) return 'IMC em kg/m²';
  if (/Razão cintura\/altura/i.test(name)) return 'Razão cintura/altura';
  if (/Prancha/i.test(name)) return 'Tempo em segundos';
  if (/25-hidroxivitamina D/i.test(name)) return 'Vitamina D em ng/mL';
  if (/Apolipoproteína B/i.test(name)) return 'ApoB em mg/dL';
  if (/Lipoproteína A/i.test(name)) return 'Lp(a) em nmol/L';
  if (/PCR ultrassensível/i.test(name)) return 'PCR-us em mg/L';
  if (/HDL Colesterol/i.test(name)) return 'HDL em mg/dL';
  if (/LDL Colesterol/i.test(name)) return 'LDL em mg/dL';
  if (/^Triglicerídeos$/i.test(name)) return 'Triglicerídeos em mg/dL';
  if (/Relação Triglicerídeos\/HDL/i.test(name)) return 'Razão TG/HDL (número)';
  if (/NT-proBNP/i.test(name)) return 'NT-proBNP em pg/mL';
  if (/Hemoglobina glicada/i.test(name)) return 'HbA1c em %';
  if (/^HOMA-IR$/i.test(name)) return 'HOMA-IR (número)';
  if (/Homocisteína/i.test(name)) return 'Homocisteína em µmol/L';
  if (/Ferritina - /i.test(name)) return 'Ferritina em ng/mL';
  if (/^TSH$/i.test(name)) return 'TSH em mIU/L';
  if (/^T3 Livre$/i.test(name)) return 'T3 Livre em pg/mL';
  if (/Testosterona Total/i.test(name)) return 'Testosterona total em ng/dL';
  if (/Estradiol -/i.test(name)) return 'Estradiol em pg/mL';
  if (/Microalbuminúria/i.test(name)) return 'Microalbuminúria em mg/g';
  if (/TC coração.*escore de cálcio/i.test(name)) return 'Escore de cálcio (Agatston)';
  if (/Densitometria.*T-Score/i.test(name)) return 'T-score (desvios-padrão)';

  // Fallback por unit
  const unit = item.unit?.trim();
  if (unit) return `Valor em ${unit}`;
  return 'Valor numérico';
}

function placeholderForLocalized(item: LightItemConfig, locale: string): string {
  if (locale !== 'en') return placeholderFor(item);
  // EN fallback — placeholders específicos seguem PT pois referenciam nomes
  // de itens armazenados em PT no banco; isso é coerente com o nome do
  // item visível ao lado. Apenas o fallback genérico é traduzido.
  const unit = item.unit?.trim();
  if (unit) return `Value in ${unit}`;
  return 'Numeric value';
}

/**
 * Classifica um valor numérico contra os levels do item.
 *
 * Estratégia em duas passadas:
 * 1) Match exato: percorre levels (ASC, worst → best) e retorna o primeiro
 *    cujo intervalo contenha o valor.
 * 2) Fallback "nearest-range": se nenhum level cobrir exatamente o valor
 *    (gaps de borda decimal, faixas U-shape com lacunas), calcula a
 *    distância do valor a cada level e retorna o mais próximo.
 *
 * Isso garante que QUALQUER valor numérico válido recebe uma
 * classificação — eliminando "fora da faixa" causados por
 * imprecisão de configuração nos limites.
 */
export function classifyNumeric(
  value: number,
  levels: LightLevelConfig[],
): LightLevelConfig | null {
  if (!levels || levels.length === 0) return null;
  const sorted = [...levels].sort((a, b) => a.level - b.level);

  // 1) Match exato
  for (const lv of sorted) {
    if (matchesLevel(value, lv)) return lv;
  }

  // 2) Fallback: nearest range
  let best: { lv: LightLevelConfig; dist: number } | null = null;
  for (const lv of sorted) {
    const d = distanceToLevel(value, lv);
    if (d === null) continue;
    if (best === null || d < best.dist) {
      best = { lv, dist: d };
    }
  }
  return best?.lv ?? null;
}

function matchesLevel(value: number, lv: LightLevelConfig): boolean {
  const lo = lv.lowerLimit !== undefined && lv.lowerLimit !== '' ? Number(lv.lowerLimit) : null;
  const hi = lv.upperLimit !== undefined && lv.upperLimit !== '' ? Number(lv.upperLimit) : null;
  switch (lv.operator) {
    case '=':
      return lo !== null && value === lo;
    case '>':
      return lo !== null && value > lo;
    case '>=':
      return lo !== null && value >= lo;
    case '<':
      return lo !== null && value < lo;
    case '<=':
      return lo !== null && value <= lo;
    case 'between':
      return lo !== null && hi !== null && value >= lo && value <= hi;
    default:
      return false;
  }
}

/**
 * Distância (não-negativa) entre `value` e o intervalo do `lv`.
 * 0 = valor está dentro do intervalo. Quanto maior, mais longe.
 * Retorna null se o level não tem limites numéricos válidos.
 */
function distanceToLevel(value: number, lv: LightLevelConfig): number | null {
  const lo = lv.lowerLimit !== undefined && lv.lowerLimit !== '' ? Number(lv.lowerLimit) : null;
  const hi = lv.upperLimit !== undefined && lv.upperLimit !== '' ? Number(lv.upperLimit) : null;
  if (lo === null && hi === null) return null;
  switch (lv.operator) {
    case '=':
      return lo === null ? null : Math.abs(value - lo);
    case '>':
    case '>=':
      // intervalo [lo, +∞) — distância só se value < lo
      return lo === null ? null : Math.max(0, lo - value);
    case '<':
    case '<=':
      // intervalo (-∞, lo] — distância só se value > lo
      return lo === null ? null : Math.max(0, value - lo);
    case 'between':
      if (lo === null || hi === null) return null;
      if (value < lo) return lo - value;
      if (value > hi) return value - hi;
      return 0;
    default:
      return null;
  }
}

/** Cor por nível (0 = pior → 5 = melhor). */
function levelColorClass(level: number): string {
  if (level === 0) return 'text-red-700 bg-red-50 border-red-200';
  if (level === 1) return 'text-orange-700 bg-orange-50 border-orange-200';
  if (level === 2) return 'text-amber-700 bg-amber-50 border-amber-200';
  if (level === 3) return 'text-yellow-700 bg-yellow-50 border-yellow-200';
  if (level === 4) return 'text-emerald-700 bg-emerald-50 border-emerald-200';
  return 'text-emerald-800 bg-emerald-100 border-emerald-300';
}

function levelBand(level: number, t: (key: string) => string): string {
  if (level === 0) return t('classifierBand0');
  if (level === 1) return t('classifierBand1');
  if (level === 2) return t('classifierBand2');
  if (level === 3) return t('classifierBand3');
  if (level === 4) return t('classifierBand4');
  return t('classifierBand5');
}

/** Resumo legível da faixa do level (ex: "≥ 190", "70 a 99", "< 70"). */
export function levelRangeSummary(lv: LightLevelConfig): string {
  const lo = lv.lowerLimit ?? '';
  const hi = lv.upperLimit ?? '';
  switch (lv.operator) {
    case 'between':
      return `${lo} a ${hi}`;
    case '=':
      return `= ${lo}`;
    case '>':
      return `> ${lo}`;
    case '>=':
      return `≥ ${lo}`;
    case '<':
      return `< ${lo}`;
    case '<=':
      return `≤ ${lo}`;
    default:
      return lv.name ?? '';
  }
}

export function NumericClassifierInput({
  item,
  value,
  onChange,
  readOnly,
  derivedNote,
}: {
  item: LightItemConfig;
  value: number | undefined;
  onChange: (v: number | undefined) => void;
  readOnly?: boolean;
  derivedNote?: string;
}) {
  const t = useTranslations('escoreLight');
  // Locale só pra escolher placeholder; useLocale do next-intl não funciona em
  // todos os contextos client-side, então deduzimos via document.documentElement.lang.
  const locale = typeof document !== 'undefined' ? document.documentElement.lang : 'pt';
  const matched = typeof value === 'number' ? classifyNumeric(value, item.levels) : null;

  return (
    <div className="space-y-3">
      {derivedNote && (
        <p className="text-petrol/55 text-xs italic">{derivedNote}</p>
      )}
      <div className="flex items-center gap-3">
        <input
          type="number"
          inputMode="decimal"
          step="any"
          value={value ?? ''}
          readOnly={readOnly}
          aria-readonly={readOnly}
          onChange={(e) => {
            if (readOnly) return;
            if (e.target.value === '') {
              onChange(undefined);
            } else {
              const n = Number(e.target.value);
              if (!Number.isNaN(n)) onChange(n);
            }
          }}
          className={`flex-1 border px-4 py-3 text-petrol text-lg tabular-nums focus:outline-none ${
            readOnly
              ? 'border-petrol/10 bg-petrol/[0.03] cursor-not-allowed'
              : 'border-petrol/20 bg-cream focus:border-gold'
          }`}
          placeholder={readOnly ? t('classifierWaiting') : placeholderForLocalized(item, locale)}
        />
      </div>

      {matched && (
        <div className={`flex items-center justify-between gap-3 px-4 py-2.5 border rounded-md ${levelColorClass(matched.level)}`}>
          <div>
            <p className="label-upper text-[10px] opacity-70">{t('classifierYourBand')}</p>
            <p className="text-base mt-0.5">
              {levelBand(matched.level, t)} <span className="opacity-60">· {levelRangeSummary(matched)}</span>
            </p>
          </div>
        </div>
      )}

      {typeof value === 'number' && !matched && (
        <p className="text-petrol/50 text-sm">{t('classifierOutOfRange')}</p>
      )}
    </div>
  );
}
