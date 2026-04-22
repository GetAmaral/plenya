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

/** Classifica um valor numérico contra os levels do item. Retorna o level matchado ou null. */
export function classifyNumeric(
  value: number,
  levels: LightLevelConfig[],
): LightLevelConfig | null {
  // Ordena por level ASC (worst → best) para que ">=" mais extremo seja avaliado primeiro
  const sorted = [...levels].sort((a, b) => a.level - b.level);
  for (const lv of sorted) {
    const lo = lv.lowerLimit ? Number(lv.lowerLimit) : null;
    const hi = lv.upperLimit ? Number(lv.upperLimit) : null;
    switch (lv.operator) {
      case '=':
        if (lo !== null && value === lo) return lv;
        break;
      case '>':
        if (lo !== null && value > lo) return lv;
        break;
      case '>=':
        if (lo !== null && value >= lo) return lv;
        break;
      case '<':
        if (lo !== null && value < lo) return lv;
        break;
      case '<=':
        if (lo !== null && value <= lo) return lv;
        break;
      case 'between':
        if (lo !== null && hi !== null && value >= lo && value <= hi) return lv;
        break;
    }
  }
  return null;
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

function levelBand(level: number): string {
  if (level === 0) return 'Faixa de risco alto';
  if (level === 1) return 'Faixa de risco';
  if (level === 2) return 'Limítrofe';
  if (level === 3) return 'Aceitável';
  if (level === 4) return 'Bom';
  return 'Ótimo';
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
}: {
  item: LightItemConfig;
  value: number | undefined;
  onChange: (v: number | undefined) => void;
}) {
  const matched = typeof value === 'number' ? classifyNumeric(value, item.levels) : null;

  return (
    <div className="space-y-3">
      <div className="flex items-center gap-3">
        <input
          type="number"
          inputMode="decimal"
          step="any"
          value={value ?? ''}
          onChange={(e) => {
            if (e.target.value === '') {
              onChange(undefined);
            } else {
              const n = Number(e.target.value);
              if (!Number.isNaN(n)) onChange(n);
            }
          }}
          className="flex-1 border border-petrol/20 bg-cream px-4 py-3 text-petrol text-lg tabular-nums focus:border-gold focus:outline-none"
          placeholder={placeholderFor(item)}
        />
      </div>

      {matched && (
        <div className={`flex items-center justify-between gap-3 px-4 py-2.5 border rounded-md ${levelColorClass(matched.level)}`}>
          <div>
            <p className="label-upper text-[10px] opacity-70">Sua faixa</p>
            <p className="text-base mt-0.5">
              {levelBand(matched.level)} <span className="opacity-60">· {levelRangeSummary(matched)}</span>
            </p>
          </div>
        </div>
      )}

      {typeof value === 'number' && !matched && (
        <p className="text-petrol/50 text-sm">Valor fora das faixas conhecidas — confira a unidade.</p>
      )}
    </div>
  );
}
