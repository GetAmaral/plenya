import type { AgirRadarData, AgirSnapshotInput, RadarLetter, RadarPillar } from './types'

/**
 * Fonte ÚNICA de dados do radar AGIR (EMR + site). Deriva letras + pilares DIRETO do
 * snapshot (`itemResults[].item.methodPillars`), agregando pontos por pilar e por letra.
 *
 * Regras (Dr.):
 * - "Preenchido" = item com status `evaluated`. "Não preenchido" = `no_data_available`/`not_applicable`.
 * - Pilar só aparece se tiver ≥1 item preenchido. Item avaliado com nota 0 APARECE zerado.
 * - TODAS as letras do método sempre aparecem (o RadarAgir dá a elas um arco mínimo).
 *
 * Retorna null quando o snapshot não tem mapeamento de método (ex.: escore por grupo).
 */
export function buildAgir(snapshot: AgirSnapshotInput): AgirRadarData | null {
  const withPillars = (snapshot.itemResults ?? []).filter(
    (ir) => (ir.item?.methodPillars?.length ?? 0) > 0,
  )
  if (withPillars.length === 0) return null

  const pillarMap = new Map<
    string,
    { name: string; letter: string; order: number; actual: number; max: number; filled: number }
  >()
  const letterMap = new Map<
    string,
    { name: string; color: string; order: number; actual: number; max: number }
  >()

  for (const ir of withPillars) {
    const evaluated = ir.status === 'evaluated'
    for (const mp of ir.item?.methodPillars ?? []) {
      const letter = mp.letter
      if (!letter) continue

      const p =
        pillarMap.get(mp.id) ??
        { name: mp.name, letter: letter.code, order: mp.order ?? 0, actual: 0, max: 0, filled: 0 }
      if (evaluated) {
        p.actual += ir.actualPoints
        p.max += ir.maxPoints
        p.filled += 1
      }
      pillarMap.set(mp.id, p)

      const l =
        letterMap.get(letter.code) ??
        { name: letter.name, color: letter.color || '#94a3b8', order: letter.order ?? 0, actual: 0, max: 0 }
      if (evaluated) {
        l.actual += ir.actualPoints
        l.max += ir.maxPoints
      }
      letterMap.set(letter.code, l)
    }
  }
  if (letterMap.size === 0) return null

  const letters: RadarLetter[] = [...letterMap.entries()]
    .sort((a, b) => a[1].order - b[1].order)
    .map(([code, l]) => ({ code, name: l.name, color: l.color, score: l.max > 0 ? (l.actual / l.max) * 100 : 0 }))

  const letterRank = new Map(letters.map((l, i) => [l.code, i]))
  const pillars: RadarPillar[] = [...pillarMap.values()]
    .filter((p) => p.filled > 0)
    .sort(
      (a, b) =>
        (letterRank.get(a.letter) ?? 0) - (letterRank.get(b.letter) ?? 0) || a.order - b.order,
    )
    .map((p) => ({ letter: p.letter, name: p.name, score: p.max > 0 ? (p.actual / p.max) * 100 : 0 }))

  return { letters, pillars }
}
