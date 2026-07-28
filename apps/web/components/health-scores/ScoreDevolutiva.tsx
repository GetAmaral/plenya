'use client'

/**
 * Modo Devolutiva — tela de apresentação do escore para o paciente na consulta de retorno.
 *
 * Cinco passos:
 *   1 Panorama   — radar AGIR + escore + completude + leitura rápida dos 3 pontos mais fracos
 *   2 Mapa       — todos os pilares como tiles coloridos, mais os que ainda não foram medidos
 *   3 Focos      — ranking por ganho potencial, separando o que é modificável do que é contexto
 *   4 Pilar      — um pilar por vez, subpilares em barras e itens com o nível que casou
 *   5 Plano      — os focos marcados viram plano, com explicação e conduta do nível
 *
 * Consome o mesmo snapshot da página de detalhe; nenhuma chamada de API própria.
 * Navegação por seta esquerda/direita para apresentar sem mouse.
 */

import { useEffect, useMemo, useState } from 'react'
import type { PatientScoreSnapshot } from '@/lib/api/health-score-api'
import { RadarAgir } from '@/components/health-scores/RadarAgir'
import { buildAgir } from '@/components/health-scores/build-agir'
import { Button } from '@/components/ui/button'
import { ArrowLeft, ArrowRight, Check, X } from 'lucide-react'

// ── Paleta Plenya ────────────────────────────────────────────────────────────
const INK = '#063b4f'
const GOLD = '#b38645'
const CREAM = '#fbfaf6'
const LINE = 'rgba(6,59,79,0.12)'
const MUTED = 'rgba(6,59,79,0.08)'

/**
 * Faixas de leitura. Recalibradas em 2026-07-28: com o corte antigo em 80, um paciente
 * sedentário e com queixa cognitiva era rotulado "Ótimo", que é exatamente o problema que a
 * devolutiva existe para não cometer.
 */
const BANDS = [
  { min: 85, label: 'Ótimo', color: '#4f8f6d', soft: '#e4f0e9' },
  { min: 70, label: 'Bom', color: '#417e8e', soft: '#e2eef1' },
  { min: 50, label: 'Atenção', color: '#c9922f', soft: '#f8efdb' },
  { min: 0, label: 'Prioridade', color: '#b4523f', soft: '#f8e5e1' },
] as const

function band(pct: number) {
  return BANDS.find((b) => pct >= b.min) ?? BANDS[BANDS.length - 1]
}

const NO_DATA = { label: 'Não medido', color: 'rgba(6,59,79,0.45)', soft: 'rgba(6,59,79,0.05)' }

/**
 * Cor de um item isolado. Vem do NÍVEL, não da razão de pontos: com a curva atual o nível 2
 * vale 40%, o que jogaria um "Raramente" na faixa vermelha e assustaria o paciente à toa.
 */
function levelBand(level: number) {
  if (level >= 5) return BANDS[0]
  if (level === 4) return BANDS[1]
  if (level === 3) return BANDS[1]
  if (level === 2) return BANDS[2]
  return BANDS[3]
}

/**
 * O que o paciente não muda daqui para frente: histórico familiar, genética, cirurgias já
 * feitas e o histórico de infância/juventude. Continuam pesando no escore, mas não podem
 * virar meta do plano — se entrarem no ranking de focos, o paciente sai da consulta com uma
 * tarefa impossível.
 */
function isContexto(group: string, subgroup: string) {
  if (group === 'Histórico Familiar de Doenças' || group === 'Genética') return true
  if (/^Cirurgias/i.test(subgroup)) return true
  if (/^Histórico/i.test(subgroup)) return true
  return false
}

type ItemRow = {
  id: string
  name: string
  subgroup: string
  group: string
  level: number
  levelName: string
  explanation?: string
  conduct?: string
  actual: number
  max: number
  lost: number
  contexto: boolean
}

type Pillar = {
  key: string
  name: string
  letter: string
  letterName: string
  actual: number
  max: number
  pct: number
  /** itens do pilar com dado × total de itens do pilar no escore */
  evaluated: number
  total: number
  items: ItemRow[]
}

const LETTER_ORDER = ['A', 'G', 'I', 'R']
const STEPS = ['Panorama', 'Mapa dos pilares', 'Onde focar', 'Detalhe do pilar', 'Plano']

export function ScoreDevolutiva({
  snapshot,
  patientName,
  onExit,
}: {
  snapshot: PatientScoreSnapshot
  patientName: string
  onExit: () => void
}) {
  const [step, setStep] = useState(0)
  const [focusPillar, setFocusPillar] = useState<string | null>(null)
  const [chosen, setChosen] = useState<Set<string>>(new Set())
  const [verTudo, setVerTudo] = useState(false)

  const agir = useMemo(() => buildAgir(snapshot as never), [snapshot])

  const { pillars, semDados, byLetter, focos, contexto, total, completude } = useMemo(() => {
    const all = snapshot.itemResults ?? []
    const groupName = new Map(
      (snapshot.groupResults ?? []).map((g) => [g.groupId, g.group?.name ?? '']),
    )

    const map = new Map<string, Pillar>()
    const rows: ItemRow[] = []

    for (const r of all) {
      const item = r.item as never as {
        name: string
        subgroup?: { name?: string }
        methodPillars?: Array<{
          id: string
          name: string
          order?: number
          letter?: { code: string; name: string }
        }>
      } | undefined
      if (!item) continue

      const grupo = groupName.get(r.groupId) ?? ''
      const sub = item.subgroup?.name ?? '—'
      const evaluated = r.status === 'evaluated'

      let row: ItemRow | null = null
      if (evaluated) {
        const lvl = (r as never as {
          levelMatched?: { name?: string; patientExplanation?: string; conduct?: string }
        }).levelMatched
        row = {
          id: r.id,
          name: item.name,
          subgroup: sub,
          group: grupo,
          level: r.levelNumber ?? 5,
          levelName: lvl?.name ?? '',
          explanation: lvl?.patientExplanation,
          conduct: lvl?.conduct,
          actual: r.actualPoints,
          max: r.maxPoints,
          lost: r.maxPoints - r.actualPoints,
          contexto: isContexto(grupo, sub),
        }
        rows.push(row)
      }

      for (const mp of item.methodPillars ?? []) {
        if (!mp.letter) continue
        const p =
          map.get(mp.id) ??
          ({
            key: mp.id,
            name: mp.name,
            letter: mp.letter.code,
            letterName: mp.letter.name,
            actual: 0,
            max: 0,
            pct: 0,
            evaluated: 0,
            total: 0,
            items: [],
          } as Pillar)
        p.total += 1
        if (evaluated && row) {
          p.actual += row.actual
          p.max += row.max
          p.evaluated += 1
          p.items.push(row)
        }
        map.set(mp.id, p)
      }
    }

    const todos = [...map.values()]
    // Pilar sem peso medido (max 0) é "não medido", não zero. Zero na tela lê como péssimo.
    const medidos = todos
      .filter((p) => p.max > 0)
      .map((p) => ({ ...p, pct: (p.actual / p.max) * 100 }))
      .sort(
        (a, b) =>
          LETTER_ORDER.indexOf(a.letter) - LETTER_ORDER.indexOf(b.letter) || a.pct - b.pct,
      )
    const naoMedidos = todos
      .filter((p) => p.max === 0)
      .sort(
        (a, b) =>
          LETTER_ORDER.indexOf(a.letter) - LETTER_ORDER.indexOf(b.letter) ||
          a.name.localeCompare(b.name),
      )

    const grouped = LETTER_ORDER.map((code) => ({
      code,
      name: medidos.find((p) => p.letter === code)?.letterName ?? '',
      pillars: medidos.filter((p) => p.letter === code),
    })).filter((l) => l.pillars.length > 0)

    const perdidos = rows.filter((r) => r.lost > 0).sort((a, b) => b.lost - a.lost)

    const evaluatedCount = snapshot.itemsEvaluatedCount ?? 0
    const notEvaluated = snapshot.itemsNotEvaluatedCount ?? 0

    return {
      pillars: medidos,
      semDados: naoMedidos,
      byLetter: grouped,
      focos: perdidos.filter((r) => !r.contexto),
      contexto: perdidos.filter((r) => r.contexto).slice(0, 4),
      total: snapshot.totalScorePercentage ?? 0,
      completude:
        evaluatedCount + notEvaluated > 0
          ? (evaluatedCount / (evaluatedCount + notEvaluated)) * 100
          : 0,
    }
  }, [snapshot])

  const active = pillars.find((p) => p.key === focusPillar) ?? pillars[0]

  // Navegação por teclado — apresentar sem tirar a mão do teclado.
  useEffect(() => {
    const onKey = (e: KeyboardEvent) => {
      if (e.key === 'ArrowRight') setStep((s) => Math.min(STEPS.length - 1, s + 1))
      if (e.key === 'ArrowLeft') setStep((s) => Math.max(0, s - 1))
      if (e.key === 'Escape') onExit()
    }
    window.addEventListener('keydown', onKey)
    return () => window.removeEventListener('keydown', onKey)
  }, [onExit])

  const abrirPilar = (key: string) => {
    setFocusPillar(key)
    setStep(3)
  }

  return (
    <div className="fixed inset-0 z-50 overflow-y-auto" style={{ background: CREAM, color: INK }}>
      <div
        className="sticky top-0 z-10 flex items-center justify-between gap-4 border-b px-6 py-3"
        style={{ background: CREAM, borderColor: LINE }}
      >
        <div className="flex items-baseline gap-3">
          <span className="text-sm font-semibold tracking-wide" style={{ color: GOLD }}>
            DEVOLUTIVA
          </span>
          <span className="text-sm opacity-70">{patientName}</span>
        </div>

        <div className="hidden items-center gap-1.5 md:flex">
          {STEPS.map((s, i) => (
            <button
              key={s}
              onClick={() => setStep(i)}
              className="rounded-full px-3 py-1 text-xs transition-colors"
              style={{
                background: i === step ? INK : 'transparent',
                color: i === step ? CREAM : 'rgba(6,59,79,0.55)',
              }}
            >
              {s}
            </button>
          ))}
        </div>

        <Button variant="ghost" size="sm" onClick={onExit} style={{ color: INK }}>
          <X className="mr-1 h-4 w-4" /> Sair
        </Button>
      </div>

      <div className="mx-auto max-w-6xl px-6 py-8">
        {step === 0 && (
          <Panorama
            agir={agir}
            total={total}
            completude={completude}
            snapshot={snapshot}
            fracos={pillars.slice(0, 3)}
            onPick={abrirPilar}
          />
        )}
        {step === 1 && <Mapa byLetter={byLetter} semDados={semDados} onPick={abrirPilar} />}
        {step === 2 && (
          <Focos
            focos={focos}
            contexto={contexto}
            verTudo={verTudo}
            setVerTudo={setVerTudo}
            chosen={chosen}
            toggle={(id) =>
              setChosen((prev) => {
                const n = new Set(prev)
                if (n.has(id)) n.delete(id)
                else n.add(id)
                return n
              })
            }
          />
        )}
        {step === 3 && active && (
          <PilarDetalhe pilar={active} pillars={pillars} onPick={setFocusPillar} />
        )}
        {step === 4 && <Plano focos={focos.filter((f) => chosen.has(f.id))} />}
      </div>

      <div
        className="sticky bottom-0 flex items-center justify-between border-t px-6 py-3"
        style={{ background: CREAM, borderColor: LINE }}
      >
        <Button
          variant="ghost"
          disabled={step === 0}
          onClick={() => setStep((s) => s - 1)}
          style={{ color: INK }}
        >
          <ArrowLeft className="mr-1 h-4 w-4" /> Voltar
        </Button>
        <span className="text-xs opacity-50">
          {step + 1} de {STEPS.length}
        </span>
        <Button
          disabled={step === STEPS.length - 1}
          onClick={() => setStep((s) => s + 1)}
          style={{ background: INK, color: CREAM }}
        >
          Avançar <ArrowRight className="ml-1 h-4 w-4" />
        </Button>
      </div>
    </div>
  )
}

// ── 1. Panorama ──────────────────────────────────────────────────────────────
function Panorama({
  agir,
  total,
  completude,
  snapshot,
  fracos,
  onPick,
}: {
  agir: ReturnType<typeof buildAgir>
  total: number
  completude: number
  snapshot: PatientScoreSnapshot
  fracos: Pillar[]
  onPick: (key: string) => void
}) {
  const b = band(total)
  const parcial = completude < 60
  const medidos = snapshot.itemsEvaluatedCount ?? 0
  const universo = medidos + (snapshot.itemsNotEvaluatedCount ?? 0)

  return (
    <div className="grid items-start gap-8 md:grid-cols-[minmax(0,1fr)_20rem]">
      <div className="order-2 md:order-1">
        {agir && (
          <RadarAgir letters={agir.letters} pillars={agir.pillars} globalScore={total} showLegend />
        )}
      </div>

      <div className="order-1 space-y-7 md:order-2">
        <div>
          <p className="text-xs uppercase tracking-widest opacity-50">Escore Plenya</p>
          <div className="flex items-end gap-3">
            <span className="text-7xl font-light leading-none" style={{ color: b.color }}>
              {total.toFixed(0)}
              <span className="text-3xl">%</span>
            </span>
            <span
              className="mb-2 rounded-full px-3 py-1 text-xs font-medium"
              style={{ background: b.soft, color: b.color }}
            >
              {b.label}
            </span>
          </div>
          {parcial && (
            <p className="mt-2 text-xs leading-relaxed opacity-70">
              Resultado parcial. Vale só para o que já foi medido, e ainda falta a maior parte
              do escore.
            </p>
          )}
        </div>

        <div className="space-y-2">
          <div className="flex items-baseline justify-between text-sm">
            <span className="opacity-70">Completude da avaliação</span>
            <span className="font-medium">{completude.toFixed(0)}%</span>
          </div>
          <div className="h-2 w-full overflow-hidden rounded-full" style={{ background: MUTED }}>
            <div className="h-full rounded-full" style={{ width: `${completude}%`, background: GOLD }} />
          </div>
          <p className="text-xs leading-relaxed opacity-60">
            {medidos} de {universo} pontos do escore foram medidos.
          </p>
        </div>

        {fracos.length > 0 && (
          <div className="space-y-2">
            <p className="text-xs uppercase tracking-widest opacity-50">Leitura rápida</p>
            <ul className="space-y-1.5">
              {fracos.map((p) => {
                const pb = band(p.pct)
                return (
                  <li key={p.key}>
                    <button
                      onClick={() => onPick(p.key)}
                      className="flex w-full items-center justify-between gap-3 rounded-lg px-3 py-2 text-left transition-colors hover:brightness-95"
                      style={{ background: pb.soft }}
                    >
                      <span className="min-w-0 truncate text-sm">{p.name}</span>
                      <span className="text-sm font-medium" style={{ color: pb.color }}>
                        {p.pct.toFixed(0)}%
                      </span>
                    </button>
                  </li>
                )
              })}
            </ul>
            <p className="text-xs opacity-55">Os três pilares mais frágeis entre os medidos.</p>
          </div>
        )}
      </div>
    </div>
  )
}

// ── 2. Mapa dos pilares ──────────────────────────────────────────────────────
function Mapa({
  byLetter,
  semDados,
  onPick,
}: {
  byLetter: Array<{ code: string; name: string; pillars: Pillar[] }>
  semDados: Pillar[]
  onPick: (key: string) => void
}) {
  return (
    <div className="space-y-8">
      <header>
        <h2 className="text-2xl font-light">Como cada pilar está hoje</h2>
        <p className="mt-1 text-sm opacity-60">
          Do mais frágil para o mais forte dentro de cada letra do método. Clique num pilar para
          abrir o detalhe.
        </p>
      </header>

      <div className="flex flex-wrap items-center gap-4 text-xs">
        {BANDS.map((b) => (
          <span key={b.label} className="flex items-center gap-2">
            <span className="h-3 w-3 rounded-sm" style={{ background: b.color }} />
            {b.label}
          </span>
        ))}
        <span className="flex items-center gap-2 opacity-60">
          <span className="h-3 w-3 rounded-sm" style={{ background: NO_DATA.color }} />
          {NO_DATA.label}
        </span>
      </div>

      {byLetter.map((l) => (
        <section key={l.code} className="space-y-3">
          <div className="flex items-baseline gap-3">
            <span
              className="flex h-8 w-8 items-center justify-center rounded-full text-sm font-semibold"
              style={{ background: INK, color: CREAM }}
            >
              {l.code}
            </span>
            <h3 className="text-lg font-medium">{l.name}</h3>
          </div>

          <div className="grid gap-2 sm:grid-cols-2 lg:grid-cols-3">
            {l.pillars.map((p) => {
              const b = band(p.pct)
              // Só sinaliza quando o número pode enganar: pouquíssimos itens sustentando a nota.
              const basePequena = p.evaluated <= 2
              return (
                <button
                  key={p.key}
                  onClick={() => onPick(p.key)}
                  className="group rounded-lg p-3 text-left transition-shadow hover:shadow-md"
                  style={{ background: b.soft }}
                >
                  <div className="flex items-start justify-between gap-2">
                    <span className="text-sm font-medium leading-tight">{p.name}</span>
                    <span className="text-lg font-light leading-none" style={{ color: b.color }}>
                      {p.pct.toFixed(0)}
                    </span>
                  </div>
                  <div
                    className="mt-2 h-1.5 w-full overflow-hidden rounded-full"
                    style={{ background: MUTED }}
                  >
                    <div className="h-full rounded-full" style={{ width: `${p.pct}%`, background: b.color }} />
                  </div>
                  <p className="mt-1.5 text-xs opacity-55">
                    {p.evaluated} de {p.total} itens medidos
                    {basePequena ? ' · base pequena' : ''}
                  </p>
                </button>
              )
            })}
          </div>
        </section>
      ))}

      {semDados.length > 0 && (
        <section className="space-y-3 border-t pt-6" style={{ borderColor: LINE }}>
          <h3 className="text-lg font-medium">Ainda não medidos</h3>
          <p className="text-sm opacity-60">
            Estes pilares não entram no escore porque ainda não há dado para eles. É onde os
            exames e as próximas etapas da avaliação vão preencher.
          </p>
          <div className="flex flex-wrap gap-1.5">
            {semDados.map((p) => (
              <span
                key={p.key}
                className="rounded-full px-2.5 py-1 text-xs"
                style={{ background: NO_DATA.soft, color: NO_DATA.color }}
              >
                {p.letter} · {p.name}
              </span>
            ))}
          </div>
        </section>
      )}
    </div>
  )
}

// ── 3. Onde focar ────────────────────────────────────────────────────────────
function Focos({
  focos,
  contexto,
  verTudo,
  setVerTudo,
  chosen,
  toggle,
}: {
  focos: ItemRow[]
  contexto: ItemRow[]
  verTudo: boolean
  setVerTudo: (v: boolean) => void
  chosen: Set<string>
  toggle: (id: string) => void
}) {
  const maxLost = Math.max(...focos.map((f) => f.lost), 1)
  const visiveis = verTudo ? focos : focos.slice(0, 6)

  return (
    <div className="space-y-6">
      <header>
        <h2 className="text-2xl font-light">Onde está o maior ganho</h2>
        <p className="mt-1 text-sm opacity-60">
          Ordenado pelo que cada ponto devolve ao escore se melhorar. Só entra aqui o que dá para
          mudar daqui para frente. Marque o que vai virar plano.
        </p>
      </header>

      <ul className="space-y-2">
        {visiveis.map((f) => {
          const b = levelBand(f.level)
          const on = chosen.has(f.id)
          return (
            <li key={f.id}>
              <button
                onClick={() => toggle(f.id)}
                className="flex w-full items-center gap-4 rounded-lg border p-4 text-left transition-colors"
                style={{
                  borderColor: on ? INK : LINE,
                  background: on ? 'rgba(6,59,79,0.04)' : '#fff',
                }}
              >
                <span
                  className="flex h-6 w-6 shrink-0 items-center justify-center rounded-md border"
                  style={{
                    background: on ? INK : 'transparent',
                    borderColor: on ? INK : 'rgba(6,59,79,0.25)',
                  }}
                >
                  {on && <Check className="h-4 w-4" style={{ color: CREAM }} />}
                </span>

                <div className="min-w-0 flex-1">
                  <p className="text-sm font-medium">{f.name}</p>
                  <p className="mt-0.5 text-xs opacity-60">
                    {f.group} · {f.subgroup}
                  </p>
                  {f.levelName && (
                    <span
                      className="mt-1.5 inline-block rounded-full px-2 py-0.5 text-xs"
                      style={{ background: b.soft, color: b.color }}
                    >
                      {f.levelName}
                    </span>
                  )}
                  {f.explanation && (
                    <p className="mt-1.5 text-xs leading-relaxed opacity-70">{f.explanation}</p>
                  )}
                </div>

                <div className="w-32 shrink-0">
                  <div className="h-2 w-full overflow-hidden rounded-full" style={{ background: MUTED }}>
                    <div
                      className="h-full rounded-full"
                      style={{ width: `${(f.lost / maxLost) * 100}%`, background: GOLD }}
                    />
                  </div>
                  <p className="mt-1 text-right text-xs opacity-60">+{f.lost.toFixed(0)} pts</p>
                </div>
              </button>
            </li>
          )
        })}
      </ul>

      {focos.length > 6 && (
        <button
          onClick={() => setVerTudo(!verTudo)}
          className="text-sm underline underline-offset-4 opacity-70"
        >
          {verTudo ? 'Mostrar só os 6 principais' : `Ver os outros ${focos.length - 6}`}
        </button>
      )}

      {contexto.length > 0 && (
        <section className="space-y-2 border-t pt-6" style={{ borderColor: LINE }}>
          <h3 className="text-sm font-semibold uppercase tracking-wide opacity-60">
            Contexto, não meta
          </h3>
          <p className="text-sm opacity-60">
            Pesa no escore e orienta o rastreio, mas não muda daqui para frente. Fica fora do
            plano de propósito.
          </p>
          <div className="flex flex-wrap gap-1.5 pt-1">
            {contexto.map((c) => (
              <span
                key={c.id}
                className="rounded-full px-2.5 py-1 text-xs"
                style={{ background: MUTED, color: 'rgba(6,59,79,0.65)' }}
              >
                {c.name}
                {c.levelName ? ` · ${c.levelName}` : ''}
              </span>
            ))}
          </div>
        </section>
      )}
    </div>
  )
}

// ── 4. Detalhe do pilar ──────────────────────────────────────────────────────
function PilarDetalhe({
  pilar,
  pillars,
  onPick,
}: {
  pilar: Pillar
  pillars: Pillar[]
  onPick: (key: string) => void
}) {
  const b = band(pilar.pct)

  const subs = useMemo(() => {
    const m = new Map<string, { name: string; actual: number; max: number; items: ItemRow[] }>()
    for (const it of pilar.items) {
      // "Atual" e "Histórico" se repetem em vários grupos; sozinhos não dizem nada.
      const k = `${it.group} · ${it.subgroup}`
      const s = m.get(k) ?? { name: k, actual: 0, max: 0, items: [] }
      s.actual += it.actual
      s.max += it.max
      s.items.push(it)
      m.set(k, s)
    }
    return [...m.values()]
      .map((s) => ({ ...s, pct: s.max > 0 ? (s.actual / s.max) * 100 : 0 }))
      .sort((a, b2) => a.pct - b2.pct)
  }, [pilar])

  const navPorLetra = LETTER_ORDER.map((code) => ({
    code,
    pillars: pillars.filter((p) => p.letter === code),
  })).filter((l) => l.pillars.length > 0)

  return (
    <div className="space-y-8">
      <header className="flex flex-wrap items-end justify-between gap-4">
        <div>
          <p className="text-xs uppercase tracking-widest opacity-50">
            Letra {pilar.letter} · {pilar.letterName}
          </p>
          <h2 className="text-3xl font-light">{pilar.name}</h2>
          <p className="mt-1 text-sm opacity-60">
            {pilar.evaluated} de {pilar.total} itens do pilar têm dado
          </p>
        </div>
        <div className="text-right">
          <span className="text-5xl font-light" style={{ color: b.color }}>
            {pilar.pct.toFixed(0)}%
          </span>
          <p className="text-sm" style={{ color: b.color }}>
            {b.label}
          </p>
        </div>
      </header>

      {/* régua de faixa com marcador */}
      <div>
        <div className="flex h-3 w-full overflow-hidden rounded-full">
          {BANDS.slice()
            .reverse()
            .map((x, i, arr) => (
              <div
                key={x.label}
                className="h-full"
                style={{
                  background: x.soft,
                  flexGrow: i === arr.length - 1 ? 100 - x.min : arr[i + 1].min - x.min,
                }}
              />
            ))}
        </div>
        <div className="relative h-8">
          <div
            className="absolute -translate-x-1/2"
            style={{ left: `${Math.min(98, Math.max(2, pilar.pct))}%` }}
          >
            <div className="mx-auto h-3 w-0.5" style={{ background: INK }} />
            <span className="block whitespace-nowrap text-xs font-medium" style={{ color: INK }}>
              {pilar.pct.toFixed(0)}%
            </span>
          </div>
        </div>
      </div>

      <section className="space-y-3">
        <h3 className="text-sm font-semibold uppercase tracking-wide opacity-60">Subpilares</h3>
        {subs.map((s) => {
          const sb = band(s.pct)
          return (
            <div
              key={s.name}
              className="rounded-lg border p-4"
              style={{ borderColor: LINE, background: '#fff' }}
            >
              <div className="flex items-baseline justify-between gap-3">
                <span className="text-sm font-medium">
                  {s.name}
                  <span className="ml-2 text-xs font-normal opacity-50">
                    {s.items.length} {s.items.length === 1 ? 'item' : 'itens'}
                  </span>
                </span>
                <span className="text-sm" style={{ color: sb.color }}>
                  {s.pct.toFixed(0)}%
                </span>
              </div>
              <div className="mt-2 h-2 w-full overflow-hidden rounded-full" style={{ background: MUTED }}>
                <div className="h-full rounded-full" style={{ width: `${s.pct}%`, background: sb.color }} />
              </div>

              <ul className="mt-3 space-y-1.5">
                {s.items
                  .slice()
                  .sort((a, b2) => a.level - b2.level)
                  .map((it) => {
                    const ib = levelBand(it.level)
                    return (
                      <li key={it.id} className="flex items-baseline justify-between gap-3">
                        <span className="min-w-0 flex-1 truncate text-sm">{it.name}</span>
                        <span
                          className="shrink-0 rounded-full px-2 py-0.5 text-xs"
                          style={{ background: ib.soft, color: ib.color }}
                        >
                          {it.levelName || `N${it.level}`}
                        </span>
                      </li>
                    )
                  })}
              </ul>
            </div>
          )
        })}
      </section>

      <section className="space-y-2 border-t pt-6" style={{ borderColor: LINE }}>
        <h3 className="text-sm font-semibold uppercase tracking-wide opacity-60">
          Ir para outro pilar
        </h3>
        {navPorLetra.map((l) => (
          <div key={l.code} className="flex flex-wrap items-center gap-1.5">
            <span
              className="mr-1 flex h-5 w-5 items-center justify-center rounded-full text-xs font-semibold"
              style={{ background: INK, color: CREAM }}
            >
              {l.code}
            </span>
            {l.pillars.map((p) => (
              <button
                key={p.key}
                onClick={() => onPick(p.key)}
                className="rounded-full px-2.5 py-1 text-xs"
                style={{
                  background: p.key === pilar.key ? INK : band(p.pct).soft,
                  color: p.key === pilar.key ? CREAM : band(p.pct).color,
                }}
              >
                {p.name}
              </button>
            ))}
          </div>
        ))}
      </section>
    </div>
  )
}

// ── 5. Plano ─────────────────────────────────────────────────────────────────
function Plano({ focos }: { focos: ItemRow[] }) {
  const ganho = focos.reduce((s, f) => s + f.lost, 0)

  return (
    <div className="space-y-6">
      <header>
        <h2 className="text-2xl font-light">O plano até o próximo retorno</h2>
        <p className="mt-1 text-sm opacity-60">
          {focos.length === 0
            ? 'Nenhum foco marcado ainda. Volte para "Onde focar" e escolha os pontos.'
            : `${focos.length} ${focos.length === 1 ? 'foco' : 'focos'}, somando ${ganho.toFixed(0)} pontos de ganho possível.`}
        </p>
      </header>

      <ol className="space-y-3">
        {focos.map((f, i) => (
          <li
            key={f.id}
            className="flex gap-4 rounded-lg border p-4"
            style={{ borderColor: LINE, background: '#fff' }}
          >
            <span
              className="flex h-8 w-8 shrink-0 items-center justify-center rounded-full text-sm"
              style={{ background: GOLD, color: '#fff' }}
            >
              {i + 1}
            </span>
            <div className="min-w-0">
              <p className="text-sm font-medium">{f.name}</p>
              {f.levelName && <p className="mt-0.5 text-xs opacity-60">Hoje: {f.levelName}</p>}
              {f.explanation && (
                <p className="mt-2 text-sm leading-relaxed opacity-80">{f.explanation}</p>
              )}
              {f.conduct && (
                <p className="mt-2 text-sm font-medium leading-relaxed" style={{ color: GOLD }}>
                  {f.conduct}
                </p>
              )}
            </div>
          </li>
        ))}
      </ol>
    </div>
  )
}
