'use client'

/**
 * Visão do escore: radar à esquerda, detalhamento por pilar à direita, lado a lado e sempre
 * visível. É o que abre assim que o escore é gerado — sem botão, sem modo, sem navegação.
 *
 * O radar é a navegação: clicar num pilar abre o detalhe dele no painel ao lado. O painel
 * também lista os pilares, separados em "precisam de atenção" e "estão bem", para o paciente
 * ver num relance o que é o quê.
 */

import { useMemo, useState } from 'react'
import type { PatientScoreSnapshot } from '@/lib/api/health-score-api'
import { RadarAgir } from '@/components/health-scores/RadarAgir'
import { buildAgir } from '@/components/health-scores/build-agir'
import { ChevronLeft, ChevronRight } from 'lucide-react'

// ── Faixas de leitura ────────────────────────────────────────────────────────
const BANDS = [
  { min: 85, label: 'Ótimo', color: '#4f8f6d', soft: '#e8f2ec' },
  { min: 70, label: 'Bom', color: '#417e8e', soft: '#e6f0f3' },
  { min: 50, label: 'Atenção', color: '#c9922f', soft: '#faf1de' },
  { min: 0, label: 'Prioridade', color: '#b4523f', soft: '#fae9e5' },
] as const

const ATENCAO_ABAIXO_DE = 70

function band(pct: number) {
  return BANDS.find((b) => pct >= b.min) ?? BANDS[BANDS.length - 1]
}

/** Item isolado é lido pelo NÍVEL. A razão de pontos jogaria um "Raramente" (N2) no vermelho. */
function levelBand(level: number) {
  if (level >= 5) return BANDS[0]
  if (level >= 3) return BANDS[1]
  if (level === 2) return BANDS[2]
  return BANDS[3]
}

type ItemRow = {
  id: string
  name: string
  group: string
  subgroup: string
  level: number
  levelName: string
  explanation?: string
  conduct?: string
  actual: number
  max: number
}

type Pillar = {
  key: string
  name: string
  letter: string
  letterName: string
  pct: number
  evaluated: number
  total: number
  items: ItemRow[]
}

const LETTER_ORDER = ['A', 'G', 'I', 'R']

export function ScoreOverview({ snapshot }: { snapshot: PatientScoreSnapshot }) {
  const [aberto, setAberto] = useState<string | null>(null)
  const [verBons, setVerBons] = useState(false)

  const agir = useMemo(() => buildAgir(snapshot as never), [snapshot])

  const { pillars, semDados, completude } = useMemo(() => {
    const all = snapshot.itemResults ?? []
    const groupName = new Map(
      (snapshot.groupResults ?? []).map((g) => [g.groupId, g.group?.name ?? '']),
    )
    const map = new Map<string, Pillar & { actual: number; max: number }>()

    for (const r of all) {
      const item = r.item as never as {
        name: string
        subgroup?: { name?: string }
        methodPillars?: Array<{
          id: string
          name: string
          letter?: { code: string; name: string }
        }>
      } | undefined
      if (!item) continue

      const evaluated = r.status === 'evaluated'
      const lvl = (r as never as {
        levelMatched?: { name?: string; patientExplanation?: string; conduct?: string }
      }).levelMatched

      const row: ItemRow | null = evaluated
        ? {
            id: r.id,
            name: item.name,
            group: groupName.get(r.groupId) ?? '',
            subgroup: item.subgroup?.name ?? '',
            level: r.levelNumber ?? 5,
            levelName: lvl?.name ?? '',
            explanation: lvl?.patientExplanation,
            conduct: lvl?.conduct,
            actual: r.actualPoints,
            max: r.maxPoints,
          }
        : null

      for (const mp of item.methodPillars ?? []) {
        if (!mp.letter) continue
        const p =
          map.get(mp.id) ??
          ({
            key: mp.id,
            name: mp.name,
            letter: mp.letter.code,
            letterName: mp.letter.name,
            pct: 0,
            evaluated: 0,
            total: 0,
            items: [],
            actual: 0,
            max: 0,
          } as Pillar & { actual: number; max: number })
        p.total += 1
        if (row) {
          p.actual += row.actual
          p.max += row.max
          p.evaluated += 1
          p.items.push(row)
        }
        map.set(mp.id, p)
      }
    }

    const todos = [...map.values()]
    // max 0 é "não medido", não zero: zero na tela lê como péssimo.
    const medidos = todos
      .filter((p) => p.max > 0)
      .map((p) => ({ ...p, pct: (p.actual / p.max) * 100 }))
      .sort((a, b) => a.pct - b.pct)
    const naoMedidos = todos
      .filter((p) => p.max === 0)
      .sort(
        (a, b) =>
          LETTER_ORDER.indexOf(a.letter) - LETTER_ORDER.indexOf(b.letter) ||
          a.name.localeCompare(b.name),
      )

    const ev = snapshot.itemsEvaluatedCount ?? 0
    const ne = snapshot.itemsNotEvaluatedCount ?? 0

    return {
      pillars: medidos,
      semDados: naoMedidos,
      completude: ev + ne > 0 ? (ev / (ev + ne)) * 100 : 0,
    }
  }, [snapshot])

  if (!agir) return null

  const pilarAberto = pillars.find((p) => p.name === aberto) ?? null
  const atencao = pillars.filter((p) => p.pct < ATENCAO_ABAIXO_DE)
  const bons = pillars.filter((p) => p.pct >= ATENCAO_ABAIXO_DE)

  return (
    <div className="grid gap-6 rounded-xl border bg-card p-5 lg:grid-cols-[minmax(0,1fr)_minmax(0,26rem)]">
      {/* ── Radar ───────────────────────────────────────────────────────── */}
      <div className="flex flex-col items-center justify-center gap-3">
        <RadarAgir
          letters={agir.letters}
          pillars={agir.pillars}
          globalScore={snapshot.totalScorePercentage}
          maxWidthClass="max-w-[30rem]"
          onPillarClick={(p) => setAberto(p.name)}
          selectedPillar={pilarAberto?.name}
        />
        <p className="text-xs text-muted-foreground">
          Clique num ponto do radar para abrir o pilar ao lado.
        </p>
      </div>

      {/* ── Painel ──────────────────────────────────────────────────────── */}
      <div className="flex max-h-[38rem] flex-col overflow-hidden rounded-lg border bg-background">
        {pilarAberto ? (
          <PilarDetalhe pilar={pilarAberto} onVoltar={() => setAberto(null)} />
        ) : (
          <ListaPilares
            atencao={atencao}
            bons={bons}
            verBons={verBons}
            setVerBons={setVerBons}
            semDados={semDados}
            completude={completude}
            onAbrir={setAberto}
          />
        )}
      </div>
    </div>
  )
}

// ── Lista de pilares ─────────────────────────────────────────────────────────
function ListaPilares({
  atencao,
  bons,
  verBons,
  setVerBons,
  semDados,
  completude,
  onAbrir,
}: {
  atencao: Pillar[]
  bons: Pillar[]
  verBons: boolean
  setVerBons: (v: boolean) => void
  semDados: Pillar[]
  completude: number
  onAbrir: (name: string) => void
}) {
  return (
    <>
      <div className="border-b px-4 py-3">
        <h3 className="text-sm font-semibold">Pilares</h3>
        <p className="mt-0.5 text-xs text-muted-foreground">
          Do mais frágil para o mais forte. {completude.toFixed(0)}% do escore foi medido.
        </p>
      </div>

      <div className="flex-1 overflow-y-auto px-4 py-3">
        {atencao.length > 0 && (
          <section className="mb-4">
            <p className="mb-2 text-xs font-semibold uppercase tracking-wide text-muted-foreground">
              Precisam de atenção ({atencao.length})
            </p>
            <ul className="space-y-1.5">
              {atencao.map((p) => (
                <LinhaPilar key={p.key} pilar={p} onAbrir={onAbrir} />
              ))}
            </ul>
          </section>
        )}

        {bons.length > 0 && (
          <section className="mb-4">
            <button
              onClick={() => setVerBons(!verBons)}
              className="mb-2 flex w-full items-center justify-between text-xs font-semibold uppercase tracking-wide text-muted-foreground hover:text-foreground"
            >
              <span>Estão bem ({bons.length})</span>
              <span className="text-[10px] normal-case tracking-normal">
                {verBons ? 'ocultar' : 'mostrar'}
              </span>
            </button>
            {verBons && (
              <ul className="space-y-1.5">
                {bons.map((p) => (
                  <LinhaPilar key={p.key} pilar={p} onAbrir={onAbrir} />
                ))}
              </ul>
            )}
          </section>
        )}

        {semDados.length > 0 && (
          <section>
            <p className="mb-2 text-xs font-semibold uppercase tracking-wide text-muted-foreground">
              Ainda não medidos ({semDados.length})
            </p>
            <p className="mb-2 text-xs text-muted-foreground">
              Não entram no escore por falta de dado. É o que os exames vão preencher.
            </p>
            <div className="flex flex-wrap gap-1">
              {semDados.map((p) => (
                <span
                  key={p.key}
                  className="rounded-full bg-muted px-2 py-0.5 text-[11px] text-muted-foreground"
                >
                  {p.name}
                </span>
              ))}
            </div>
          </section>
        )}
      </div>
    </>
  )
}

function LinhaPilar({ pilar, onAbrir }: { pilar: Pillar; onAbrir: (name: string) => void }) {
  const b = band(pilar.pct)
  return (
    <li>
      <button
        onClick={() => onAbrir(pilar.name)}
        className="group flex w-full items-center gap-3 rounded-md px-2 py-2 text-left transition-colors hover:bg-muted/60"
      >
        <span
          className="flex h-6 w-6 shrink-0 items-center justify-center rounded-full text-[10px] font-semibold"
          style={{ background: b.soft, color: b.color }}
        >
          {pilar.letter}
        </span>

        <span className="min-w-0 flex-1">
          <span className="block truncate text-sm">{pilar.name}</span>
          <span className="mt-1 block h-1.5 w-full overflow-hidden rounded-full bg-muted">
            <span
              className="block h-full rounded-full"
              style={{ width: `${pilar.pct}%`, background: b.color }}
            />
          </span>
        </span>

        <span className="shrink-0 text-right">
          <span className="block text-sm font-medium" style={{ color: b.color }}>
            {pilar.pct.toFixed(0)}%
          </span>
          <span className="block text-[10px] text-muted-foreground">
            {pilar.evaluated}/{pilar.total} itens
          </span>
        </span>

        <ChevronRight className="h-4 w-4 shrink-0 text-muted-foreground opacity-0 transition-opacity group-hover:opacity-100" />
      </button>
    </li>
  )
}

// ── Detalhe de um pilar ──────────────────────────────────────────────────────
function PilarDetalhe({ pilar, onVoltar }: { pilar: Pillar; onVoltar: () => void }) {
  const b = band(pilar.pct)

  const { ruins, ok } = useMemo(() => {
    const ordenado = [...pilar.items].sort((a, c) => a.level - c.level)
    return {
      ruins: ordenado.filter((i) => i.level <= 3),
      ok: ordenado.filter((i) => i.level > 3),
    }
  }, [pilar])

  return (
    <>
      <div className="border-b px-4 py-3">
        <button
          onClick={onVoltar}
          className="mb-2 flex items-center gap-1 text-xs text-muted-foreground hover:text-foreground"
        >
          <ChevronLeft className="h-3.5 w-3.5" /> Todos os pilares
        </button>
        <div className="flex items-start justify-between gap-3">
          <div className="min-w-0">
            <h3 className="truncate text-sm font-semibold">{pilar.name}</h3>
            <p className="mt-0.5 text-xs text-muted-foreground">
              Letra {pilar.letter} · {pilar.evaluated} de {pilar.total} itens medidos
            </p>
          </div>
          <div className="shrink-0 text-right">
            <span className="text-2xl font-semibold leading-none" style={{ color: b.color }}>
              {pilar.pct.toFixed(0)}%
            </span>
            <span className="mt-0.5 block text-[11px]" style={{ color: b.color }}>
              {b.label}
            </span>
          </div>
        </div>
      </div>

      <div className="flex-1 overflow-y-auto px-4 py-3">
        {ruins.length > 0 && (
          <section className="mb-4">
            <p className="mb-2 text-xs font-semibold uppercase tracking-wide text-muted-foreground">
              O que puxa para baixo ({ruins.length})
            </p>
            <ul className="space-y-2">
              {ruins.map((it) => (
                <ItemDetalhe key={it.id} item={it} expandido />
              ))}
            </ul>
          </section>
        )}

        {ok.length > 0 && (
          <section>
            <p className="mb-2 text-xs font-semibold uppercase tracking-wide text-muted-foreground">
              O que está bem ({ok.length})
            </p>
            <ul className="space-y-1">
              {ok.map((it) => (
                <ItemDetalhe key={it.id} item={it} />
              ))}
            </ul>
          </section>
        )}

        {pilar.items.length === 0 && (
          <p className="text-sm text-muted-foreground">Nenhum item deste pilar tem dado ainda.</p>
        )}
      </div>
    </>
  )
}

function ItemDetalhe({ item, expandido = false }: { item: ItemRow; expandido?: boolean }) {
  const b = levelBand(item.level)
  return (
    <li
      className={expandido ? 'rounded-md border p-3' : 'flex items-center gap-2 px-1 py-1'}
      style={expandido ? { borderColor: b.soft, background: b.soft } : undefined}
    >
      {expandido ? (
        <>
          <div className="flex items-start justify-between gap-2">
            <span className="text-sm font-medium">{item.name}</span>
            <span
              className="shrink-0 rounded-full px-2 py-0.5 text-[11px] font-medium"
              style={{ background: '#fff', color: b.color }}
            >
              {item.levelName || `N${item.level}`}
            </span>
          </div>
          {item.explanation && (
            <p className="mt-1.5 text-xs leading-relaxed text-foreground/75">{item.explanation}</p>
          )}
          {item.conduct && (
            <p className="mt-1.5 text-xs font-medium leading-relaxed" style={{ color: b.color }}>
              {item.conduct}
            </p>
          )}
        </>
      ) : (
        <>
          <span className="h-1.5 w-1.5 shrink-0 rounded-full" style={{ background: b.color }} />
          <span className="min-w-0 flex-1 truncate text-xs">{item.name}</span>
          <span className="shrink-0 text-[11px] text-muted-foreground">{item.levelName}</span>
        </>
      )}
    </li>
  )
}
