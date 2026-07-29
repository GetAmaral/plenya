'use client'

/**
 * Detalhamento do escore, no dashboard (logo abaixo do radar do snapshot mais recente).
 *
 * A hierarquia é a mesma do radar:
 *   PILAR    = a letra do método (A, G, I, R) → 4 barras, uma por card
 *   SUBPILAR = cada ponto do radar (Cardiovascular, Função Cognitiva, …) → uma barra cada,
 *              visíveis por padrão e na mesma ordem em que aparecem no radar
 *   abrir um subpilar mostra os itens que o compõem
 *
 * Decisões de visualização (skill dataviz):
 * - Cada pilar tem UMA cor, a mesma do arco e do glifo no radar. Validada em light e dark:
 *   banda de luminosidade, piso de croma, visão normal (pior par ΔE 19,3) e contraste >= 3:1.
 *   A separação CVD fica na banda 6-8, legal porque a letra sempre vem com o glifo e o nome —
 *   a cor nunca é o único canal.
 * - Sem semântica de bom/ruim: nenhuma barra vermelha ou verde por desempenho. O comprimento
 *   carrega a magnitude.
 * - Barra do pilar cheia e mais grossa; subpilar um degrau mais claro e mais fino.
 */

import { useMemo, useState } from 'react'
import type React from 'react'
import type { PatientScoreSnapshot } from '@/lib/api/health-score-api'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { BarChart3, ChevronDown, FlaskConical } from 'lucide-react'
import { levelStyle } from '@/lib/score-level-styles'

type Item = {
  id: string
  name: string
  bloco: string
  level: number | null
  levelName: string
  actual: number
  max: number
  fonteLab: boolean
  valor?: number
  unidade?: string
}
type Subpilar = {
  key: string
  name: string
  order: number
  pct: number
  avaliados: number
  total: number
  itens: Item[]
}
type Pilar = {
  code: string
  name: string
  pct: number
  subpilares: Subpilar[]
  semDados: string[]
}

const LETTER_ORDER = ['A', 'G', 'I', 'R']

/**
 * Cor de cada pilar — as MESMAS do radar (packages/ui/src/score-radar/RadarAgir.tsx).
 * A paleta antiga reprovava: G #b38645 e I #caa56b eram o mesmo dourado, ΔE 9,3 em visão
 * normal contra piso 15.
 */
const LETTER_COLOR: Record<string, { light: string; dark: string }> = {
  A: { light: '#0f8f5f', dark: '#2b9c6d' },
  G: { light: '#c07520', dark: '#c9822f' },
  I: { light: '#b04a8a', dark: '#b85e93' },
  R: { light: '#2a6fb0', dark: '#4585bc' },
}
const CINZA = { light: '#6b7280', dark: '#9ca3af' }

/** Publica a cor do pilar como CSS var, para as classes Tailwind lerem em light e dark. */
function cssVars(code: string) {
  const c = LETTER_COLOR[code] ?? CINZA
  return { '--letra': c.light, '--letra-dark': c.dark } as React.CSSProperties
}

/**
 * Nomes de grupo/subgrupo carregam roteiro de anamnese dentro de parênteses
 * ("Doenças crônicas (Questionar ativamente as doenças crônicas mais comuns, …)").
 * Isso é instrução para o profissional, não rótulo de tela.
 */
function limpaRotulo(s: string) {
  return s
    .replace(/\s*\(([^)]{21,})\)/g, '')
    .replace(/[\s:·-]+$/, '')
    .trim()
}

export function ScorePillarBreakdown({ snapshot }: { snapshot: PatientScoreSnapshot }) {
  const [aberto, setAberto] = useState<string | null>(null)

  const pilares = useMemo<Pilar[]>(() => {
    const groupName = new Map(
      (snapshot.groupResults ?? []).map((g) => [g.groupId, g.group?.name ?? '']),
    )

    type SubAcc = {
      key: string
      name: string
      letter: string
      order: number
      actual: number
      max: number
      avaliados: number
      total: number
      itens: Item[]
    }
    const subs = new Map<string, SubAcc>()
    const letra = new Map<string, { name: string; actual: number; max: number }>()

    for (const r of snapshot.itemResults ?? []) {
      const item = r.item as never as {
        name: string
        unit?: string
        subgroup?: { name?: string }
        methodPillars?: Array<{
          id: string
          name: string
          order?: number
          letter?: { code: string; name: string }
        }>
      } | undefined
      if (!item) continue

      const evaluated = r.status === 'evaluated'
      const grupo = limpaRotulo(groupName.get(r.groupId) ?? '')
      const sub = limpaRotulo(item.subgroup?.name ?? '')
      const bloco = (sub && sub !== grupo ? [grupo, sub] : [grupo]).filter(Boolean).join(' · ')
      const levelName =
        (r as never as { levelMatched?: { name?: string } }).levelMatched?.name ?? ''

      for (const mp of item.methodPillars ?? []) {
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
            id: r.id,
            name: item.name,
            bloco,
            level: r.levelNumber ?? null,
            levelName,
            actual: r.actualPoints,
            max: r.maxPoints,
            fonteLab: r.dataSource === 'lab_result',
            valor: (r as never as { valueUsed?: number }).valueUsed,
            unidade: item.unit,
          })
        }
        subs.set(mp.id, s)

        // O escore do pilar agrega os mesmos pares item×subpilar que o radar usa na letra,
        // para o número bater com a legenda do radar.
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
        // max 0 = sem peso medido; não vira barra zerada, que leria como "péssimo"
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
          // mesma ordem do radar
          .sort((a, b) => a.order - b.order || a.name.localeCompare(b.name)),
        semDados: doPilar
          .filter((s) => s.max === 0)
          .map((s) => s.name)
          .sort((a, b) => a.localeCompare(b)),
      } as Pilar
    }).filter((p): p is Pilar => p !== null)
  }, [snapshot])

  if (pilares.length === 0) return null

  const totalSub = pilares.reduce((n, p) => n + p.subpilares.length, 0)

  return (
    <Card>
      <CardHeader>
        <CardTitle className="flex items-center gap-2 text-base">
          <BarChart3 className="h-5 w-5 text-primary" />
          Detalhamento por pilar
        </CardTitle>
        <CardDescription>
          Uma barra por pilar do método e, dentro dela, os {totalSub} subpilares que aparecem no
          radar. Clique num subpilar para ver os itens que o compõem.
        </CardDescription>
      </CardHeader>

      <CardContent className="space-y-4">
        {pilares.map((p) => (
          <CardPilar key={p.code} pilar={p} aberto={aberto} setAberto={setAberto} />
        ))}
      </CardContent>
    </Card>
  )
}

function CardPilar({
  pilar,
  aberto,
  setAberto,
}: {
  pilar: Pilar
  aberto: string | null
  setAberto: (k: string | null) => void
}) {
  return (
    <section
      style={cssVars(pilar.code)}
      className="rounded-lg border border-[var(--letra)] p-3 dark:border-[var(--letra-dark)] sm:p-4"
    >
      <div className="mb-1.5 flex flex-wrap items-baseline gap-x-2 gap-y-1">
        <span className="flex h-6 w-6 shrink-0 items-center justify-center rounded-full bg-[var(--letra)] text-[11px] font-semibold text-white dark:bg-[var(--letra-dark)]">
          {pilar.code}
        </span>
        <h3 className="min-w-0 flex-1 text-sm font-semibold">{pilar.name}</h3>
        <span className="shrink-0 text-xs text-muted-foreground">
          <span className="text-sm font-semibold text-foreground">{pilar.pct.toFixed(0)}%</span>
          <span className="ml-2">
            {pilar.subpilares.length}{' '}
            {pilar.subpilares.length === 1 ? 'subpilar' : 'subpilares'}
          </span>
        </span>
      </div>

      {/* Barra do PILAR — a mais grossa, no tom cheio */}
      <span className="mb-3 block h-3.5 w-full overflow-hidden rounded-l-[2px] rounded-r bg-[var(--letra)]/12 dark:bg-[var(--letra-dark)]/18">
        <span
          className="block h-full rounded-l-[2px] rounded-r-[4px] bg-[var(--letra)] transition-[width] duration-300 dark:bg-[var(--letra-dark)]"
          style={{ width: `${Math.max(1.5, pilar.pct)}%` }}
        />
      </span>

      {/* Barras dos SUBPILARES — os pontos do radar, na mesma ordem, fechadas */}
      <ul className="gap-x-6 space-y-[7px] xl:columns-2">
        {pilar.subpilares.map((s) => {
          const isOpen = aberto === s.key
          return (
            <li key={s.key} className="break-inside-avoid">
              <button
                onClick={() => setAberto(isOpen ? null : s.key)}
                aria-expanded={isOpen}
                className="group w-full text-left"
              >
                <span className="flex items-baseline justify-between gap-2">
                  <span className="flex min-w-0 items-center gap-1">
                    <ChevronDown
                      className={`h-3 w-3 shrink-0 text-muted-foreground transition-transform ${isOpen ? '' : '-rotate-90'}`}
                    />
                    <span className="truncate text-xs text-muted-foreground group-hover:text-foreground">
                      {s.name}
                    </span>
                  </span>
                  <span className="shrink-0 text-xs tabular-nums text-muted-foreground">
                    {s.pct.toFixed(0)}%
                  </span>
                </span>

                <span className="mt-1 block h-2 w-full overflow-hidden rounded-l-[2px] rounded-r bg-[var(--letra)]/10 dark:bg-[var(--letra-dark)]/15">
                  <span
                    className="block h-full rounded-l-[2px] rounded-r-[4px] bg-[var(--letra)]/70 transition-[width] duration-300 dark:bg-[var(--letra-dark)]/75"
                    style={{ width: `${Math.max(1.5, s.pct)}%` }}
                  />
                </span>
              </button>

              {isOpen && <ItensDoSubpilar itens={s.itens} />}
            </li>
          )
        })}
      </ul>

      {pilar.semDados.length > 0 && (
        <div className="mt-3 flex flex-wrap items-center gap-1 border-t pt-3">
          <span className="mr-1 text-[11px] text-muted-foreground">Sem dados ainda:</span>
          {pilar.semDados.map((n) => (
            <span
              key={n}
              className="rounded-full bg-muted px-2 py-0.5 text-[11px] text-muted-foreground"
            >
              {n}
            </span>
          ))}
        </div>
      )}
    </section>
  )
}

function ItensDoSubpilar({ itens }: { itens: Item[] }) {
  // Agrupa pelo bloco de origem (grupo · subgrupo) só como rótulo, sem mais um nível de clique.
  const blocos = useMemo(() => {
    const m = new Map<string, Item[]>()
    for (const it of itens) {
      const arr = m.get(it.bloco) ?? []
      arr.push(it)
      m.set(it.bloco, arr)
    }
    return [...m.entries()]
  }, [itens])

  return (
    <div className="mt-2 space-y-2 border-l pl-2 sm:pl-3">
      {blocos.map(([bloco, lista]) => (
        <div key={bloco}>
          <p className="text-[10px] uppercase tracking-wide text-muted-foreground">{bloco}</p>
          <ul className="mt-1 space-y-1.5">
            {lista.map((it) => {
              const ls = levelStyle(it.level)
              return (
                <li
                  key={it.id}
                  className="flex flex-col gap-1.5 rounded-md border bg-card p-2.5 text-sm sm:flex-row sm:items-start sm:gap-2"
                >
                  <span className="min-w-0 flex-1">
                    <span className="font-semibold">{it.name}</span>
                    {it.max > 0 && (
                      <span className="mt-1 block text-xs text-muted-foreground">
                        {it.actual.toFixed(1)} / {it.max.toFixed(1)} pts (
                        {((it.actual / it.max) * 100).toFixed(0)}%)
                      </span>
                    )}
                  </span>

                  {ls && (
                    <span
                      className={`self-start rounded-full border-2 px-2 py-0.5 text-xs font-bold sm:shrink-0 sm:whitespace-nowrap ${ls.bg} ${ls.text} ${ls.border}`}
                    >
                      {it.fonteLab ? (
                        <span className="flex items-center gap-1">
                          <FlaskConical className="h-3 w-3" />
                          {it.valor?.toFixed(2)} {it.unidade}
                        </span>
                      ) : (
                        <>
                          N{it.level}: {it.levelName}
                        </>
                      )}
                    </span>
                  )}
                </li>
              )
            })}
          </ul>
        </div>
      ))}
    </div>
  )
}
