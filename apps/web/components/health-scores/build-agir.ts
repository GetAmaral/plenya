import { PatientScoreSnapshot } from "@/lib/api/health-score-api"
import type { RadarLetter, RadarPillar } from "./RadarAgir"

export type AgirRadarData = { letters: RadarLetter[]; pillars: RadarPillar[] }

/**
 * Fonte ÚNICA do radar AGIR no EMR.
 *
 * Deriva letras + pilares DIRETO do snapshot (`itemResults[].item.methodPillars`),
 * agregando pontos por pilar e por letra. NÃO depende de assinatura/plano — o
 * mapeamento de método já vem materializado no snapshot, então a capa do paciente
 * e o painel de escores usam exatamente os mesmos números e o mesmo desenho.
 *
 * Retorna null quando o snapshot não tem mapeamento de método (ex.: escore Light,
 * que é por grupo e não tem pilares AGIR).
 */
export function buildAgir(snapshot: PatientScoreSnapshot): AgirRadarData | null {
  const items = (snapshot.itemResults ?? []).filter(
    (ir) => ir.status === "evaluated" && (ir.item?.methodPillars?.length ?? 0) > 0,
  )
  if (items.length === 0) return null

  const pillarMap = new Map<string, { name: string; letter: string; actual: number; max: number }>()
  const letterMap = new Map<string, { name: string; color: string; order: number; actual: number; max: number }>()

  for (const ir of items) {
    for (const mp of ir.item?.methodPillars ?? []) {
      const letter = mp.letter
      if (!letter) continue
      const p = pillarMap.get(mp.id) ?? { name: mp.name, letter: letter.code, actual: 0, max: 0 }
      p.actual += ir.actualPoints
      p.max += ir.maxPoints
      pillarMap.set(mp.id, p)

      const l = letterMap.get(letter.code) ?? {
        name: letter.name,
        color: letter.color || "#94a3b8",
        order: letter.order ?? 0,
        actual: 0,
        max: 0,
      }
      l.actual += ir.actualPoints
      l.max += ir.maxPoints
      letterMap.set(letter.code, l)
    }
  }
  if (letterMap.size === 0) return null

  const letters: RadarLetter[] = [...letterMap.entries()]
    .sort((a, b) => a[1].order - b[1].order)
    .map(([code, l]) => ({ code, name: l.name, color: l.color, score: l.max > 0 ? (l.actual / l.max) * 100 : 0 }))
  const pillars: RadarPillar[] = [...pillarMap.values()].map((p) => ({
    letter: p.letter,
    name: p.name,
    score: p.max > 0 ? (p.actual / p.max) * 100 : 0,
  }))
  return { letters, pillars }
}
