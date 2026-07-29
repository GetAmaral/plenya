/**
 * Cores por nível do escore — padrão Plenya, o mesmo da anamnese, da página de detalhe do
 * snapshot e do acordeão de metodologia.
 *
 * Estava duplicado em três arquivos; passou a viver aqui para não divergir.
 */
export const LEVEL_STYLES = {
  0: { bg: 'bg-red-100', text: 'text-red-900', border: 'border-red-500' },
  1: { bg: 'bg-orange-100', text: 'text-orange-900', border: 'border-orange-500' },
  2: { bg: 'bg-yellow-100', text: 'text-yellow-900', border: 'border-yellow-500' },
  3: { bg: 'bg-blue-100', text: 'text-blue-900', border: 'border-blue-500' },
  4: { bg: 'bg-green-100', text: 'text-green-900', border: 'border-green-500' },
  5: { bg: 'bg-emerald-100', text: 'text-emerald-900', border: 'border-emerald-500' },
  6: { bg: 'bg-gray-100', text: 'text-gray-900', border: 'border-gray-500' },
} as const

/** Nível fora da tabela cai no cinza (6), como já acontecia nas telas antigas. */
export function levelStyle(level: number | null | undefined) {
  if (level === null || level === undefined) return null
  return LEVEL_STYLES[level as keyof typeof LEVEL_STYLES] ?? LEVEL_STYLES[6]
}
