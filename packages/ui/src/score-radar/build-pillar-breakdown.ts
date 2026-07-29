import { LETTER_ORDER } from './letter-colors'

/**
 * Monta o detalhamento do escore na MESMA hierarquia do radar, a partir de um snapshot:
 *
 *   PILAR    = letra do método (A, G, I, R)
 *   SUBPILAR = cada ponto do radar (Cardiovascular, Função Cognitiva, …)
 *   ITEM     = o que compõe o subpilar
 *
 * Fonte única do EMR e do site: o site tem menos campos por item (o payload público é enxuto),
 * então tudo que não é essencial é opcional e a expansão simplesmente mostra menos.
 *
 * O escore do pilar agrega os mesmos pares item×subpilar que `buildAgir` usa na letra, para o
 * número bater com a legenda do radar.
 */

export type BreakdownItemResult = {
  id?: string
  status: string
  actualPoints: number
  maxPoints: number
  levelNumber?: number | null
  levelName?: string
  /** rótulo de origem do item (ex.: "Sono · Atual"); usado só como agrupador na expansão */
  bloco?: string
  valueUsed?: number
  isLab?: boolean
  item?: {
    name?: string
    unit?: string
    methodPillars?: Array<{
      id: string
      name: string
      order?: number
      letter?: { code: string; name: string }
    }>
  }
}

export type BreakdownItem = {
  id: string
  name: string
  bloco: string
  level: number | null
  levelName: string
  actual: number
  max: number
  isLab: boolean
  valueUsed?: number
  unit?: string
}

export type BreakdownSubpilar = {
  key: string
  name: string
  order: number
  pct: number
  avaliados: number
  total: number
  itens: BreakdownItem[]
}

export type BreakdownPilar = {
  code: string
  name: string
  pct: number
  subpilares: BreakdownSubpilar[]
  /** subpilares tocados mas sem peso medido — não viram barra zerada */
  semDados: string[]
}

export function buildPillarBreakdown(results: BreakdownItemResult[]): BreakdownPilar[] {
  type SubAcc = {
    key: string
    name: string
    letter: string
    order: number
    actual: number
    max: number
    avaliados: number
    total: number
    itens: BreakdownItem[]
  }
  const subs = new Map<string, SubAcc>()
  const letra = new Map<string, { name: string; actual: number; max: number }>()

  for (const r of results) {
    const pillars = r.item?.methodPillars ?? []
    if (pillars.length === 0) continue
    const evaluated = r.status === 'evaluated'

    for (const mp of pillars) {
      const l = mp.letter
      if (!l) continue

      const s =
        subs.get(mp.id) ??
        ({
          key: mp.id,
          name: mp.name,
          letter: l.code,
          order: mp.order ?? 0,
          actual: 0,
          max: 0,
          avaliados: 0,
          total: 0,
          itens: [],
        } as SubAcc)
      s.total += 1
      if (evaluated) {
        s.actual += r.actualPoints
        s.max += r.maxPoints
        s.avaliados += 1
        s.itens.push({
          id: r.id ?? `${mp.id}-${s.itens.length}`,
          name: r.item?.name ?? '',
          bloco: r.bloco ?? '',
          level: r.levelNumber ?? null,
          levelName: r.levelName ?? '',
          actual: r.actualPoints,
          max: r.maxPoints,
          isLab: r.isLab ?? false,
          valueUsed: r.valueUsed,
          unit: r.item?.unit,
        })
      }
      subs.set(mp.id, s)

      const acc = letra.get(l.code) ?? { name: l.name, actual: 0, max: 0 }
      if (evaluated) {
        acc.actual += r.actualPoints
        acc.max += r.maxPoints
      }
      letra.set(l.code, acc)
    }
  }

  const todos = [...subs.values()]

  return LETTER_ORDER.map((code) => {
    const acc = letra.get(code)
    if (!acc || acc.max <= 0) return null
    const doPilar = todos.filter((s) => s.letter === code)
    return {
      code,
      name: acc.name,
      pct: (acc.actual / acc.max) * 100,
      subpilares: doPilar
        .filter((s) => s.max > 0)
        .map((s) => ({
          key: s.key,
          name: s.name,
          order: s.order,
          pct: (s.actual / s.max) * 100,
          avaliados: s.avaliados,
          total: s.total,
          itens: s.itens.sort(
            (a, b) => a.bloco.localeCompare(b.bloco) || a.name.localeCompare(b.name),
          ),
        }))
        // mesma ordem em que o radar desenha os pontos
        .sort((a, b) => a.order - b.order || a.name.localeCompare(b.name)),
      semDados: doPilar
        .filter((s) => s.max === 0)
        .map((s) => s.name)
        .sort((a, b) => a.localeCompare(b)),
    } as BreakdownPilar
  }).filter((p): p is BreakdownPilar => p !== null)
}
