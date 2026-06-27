'use client'

import { use, useEffect, useMemo } from 'react'
import { useQuery } from '@tanstack/react-query'
import { Printer, Loader2 } from 'lucide-react'
import { getAnamnesisTemplateById } from '@/lib/api/anamnesis-templates'
import type { ScoreItem, ScoreLevel } from '@/lib/api/score-api'
import { Button } from '@/components/ui/button'
import { useRequireAuth } from '@/lib/use-auth'
import { getScaleDef } from '@plenya/domain'
import { cn } from '@/lib/utils'

// Cores por nível (espelha o pôster do score).
const LEVEL_BG: Record<number, string> = {
  0: 'bg-red-100 text-red-900 border-red-300',
  1: 'bg-orange-100 text-orange-900 border-orange-300',
  2: 'bg-yellow-100 text-yellow-900 border-yellow-300',
  3: 'bg-blue-100 text-blue-900 border-blue-300',
  4: 'bg-green-100 text-green-900 border-green-300',
  5: 'bg-emerald-100 text-emerald-900 border-emerald-300',
  6: 'bg-gray-100 text-gray-800 border-gray-300',
}

function levelRange(l: ScoreLevel): string {
  const lo = l.lowerLimit ?? ''
  const hi = l.upperLimit ?? ''
  switch (l.operator) {
    case 'between': return `${lo}–${hi}`
    case '>': return `>${lo}`
    case '>=': return `≥${lo}`
    case '<': return `<${hi}`
    case '<=': return `≤${hi}`
    case '=': return `${lo}`
    default: return ''
  }
}

function applicability(si: ScoreItem): string {
  const out: string[] = []
  if (si.gender && si.gender !== 'not_applicable') {
    out.push(si.gender === 'male' ? '♂' : si.gender === 'female' ? '♀' : String(si.gender))
  }
  if (si.ageRangeMin !== undefined || si.ageRangeMax !== undefined) {
    out.push(`${si.ageRangeMin ?? 0}–${si.ageRangeMax ?? '+'}a`)
  }
  if (si.postMenopause === true) out.push('pós-menop.')
  return out.join(' ')
}

interface OrgItem { templateOrder: number; scoreItem: ScoreItem; depth: number }
interface OrgSubgroup { name: string; order: number; items: OrgItem[] }
interface OrgGroup { name: string; order: number; subgroups: OrgSubgroup[] }

// Aninha os itens de um subgrupo por parentItemId (espelha o pôster do score):
// o pai é mantido e os filhos vêm logo abaixo, com profundidade crescente.
function nestByParent(items: OrgItem[]): OrgItem[] {
  const byId = new Map<string, OrgItem>()
  items.forEach((it) => byId.set(it.scoreItem.id, it))
  const childrenOf = new Map<string, OrgItem[]>()
  const roots: OrgItem[] = []
  for (const it of items) {
    const pid = it.scoreItem.parentItemId
    if (pid && byId.has(pid)) {
      const list = childrenOf.get(pid) ?? []
      list.push(it)
      childrenOf.set(pid, list)
    } else {
      roots.push(it)
    }
  }
  const byOrder = (a: OrgItem, b: OrgItem) => a.templateOrder - b.templateOrder
  const out: OrgItem[] = []
  const visit = (it: OrgItem, depth: number) => {
    out.push({ ...it, depth })
    ;(childrenOf.get(it.scoreItem.id) ?? []).slice().sort(byOrder).forEach((k) => visit(k, depth + 1))
  }
  roots.sort(byOrder).forEach((r) => visit(r, 0))
  return out
}

export default function AnamnesisTemplatePrintPage({
  params,
}: {
  params: Promise<{ id: string }>
}) {
  const { id } = use(params)
  useRequireAuth()

  const { data: template, isLoading } = useQuery({
    queryKey: ['anamnesis-template', id, 'print'],
    queryFn: () => getAnamnesisTemplateById(id, true),
  })

  useEffect(() => {
    if (template) document.title = `Template — ${template.name}`
  }, [template])

  const groups = useMemo<OrgGroup[]>(() => {
    const items = template?.items ?? []
    const gMap = new Map<string, OrgGroup>()
    for (const it of items) {
      const si = it.scoreItem as ScoreItem | undefined
      if (!si) continue
      const sg = si.subgroup
      const g = sg?.group
      const gName = g?.name ?? '(Sem grupo)'
      const sgName = sg?.name ?? ''
      let grp = gMap.get(gName)
      if (!grp) { grp = { name: gName, order: g?.order ?? 999, subgroups: [] }; gMap.set(gName, grp) }
      let sub = grp.subgroups.find((s) => s.name === sgName)
      if (!sub) { sub = { name: sgName, order: sg?.order ?? 999, items: [] }; grp.subgroups.push(sub) }
      sub.items.push({ templateOrder: it.order ?? 0, scoreItem: si, depth: 0 })
    }
    const arr = Array.from(gMap.values()).sort((a, b) => a.order - b.order)
    arr.forEach((g) => {
      g.subgroups.sort((a, b) => a.order - b.order)
      g.subgroups.forEach((s) => { s.items = nestByParent(s.items) })
    })
    return arr
  }, [template])

  const totalItems = groups.reduce((n, g) => n + g.subgroups.reduce((m, s) => m + s.items.length, 0), 0)

  if (isLoading || !template) {
    return <div className="flex min-h-screen items-center justify-center"><Loader2 className="h-10 w-10 animate-spin text-primary" /></div>
  }

  return (
    <div className="min-h-screen bg-white text-zinc-900">
      <div className="no-print sticky top-0 z-10 flex items-center justify-between border-b bg-white px-6 py-2">
        <span className="text-sm font-semibold">{template.name} · {totalItems} itens</span>
        <Button size="sm" onClick={() => window.print()} className="gap-2">
          <Printer className="h-4 w-4" /> Imprimir / Salvar PDF
        </Button>
      </div>

      <div className="doc mx-auto px-6 py-4 print:px-0 print:py-0">
        <header className="mb-2 border-b-2 border-zinc-800 pb-1">
          <p className="text-[8px] font-semibold uppercase tracking-widest text-zinc-400">Plenya · Template de Anamnese</p>
          <h2 className="text-[15px] font-bold leading-tight">{template.name}</h2>
          <p className="text-[9px] text-zinc-500">{template.area} · {totalItems} itens · {groups.length} grupos</p>
        </header>

        {groups.map((g) => (
          <section key={g.name} className="mb-1">
            <h3 className="mt-1.5 mb-0.5 border-b border-zinc-400 text-[10.5px] font-bold text-zinc-800">{g.name}</h3>
            {g.subgroups.map((sub) => (
              <div key={sub.name}>
                {sub.name && (
                  <p className="mt-0.5 text-[7px] font-semibold uppercase tracking-wide text-zinc-400">{sub.name}</p>
                )}
                {sub.items.map((oi) => <ItemRow key={oi.scoreItem.id} si={oi.scoreItem} depth={oi.depth} />)}
              </div>
            ))}
          </section>
        ))}
      </div>

      <style jsx global>{`
        @media print {
          .no-print { display: none !important; }
          @page { size: A4 portrait; margin: 7mm; }
          body { -webkit-print-color-adjust: exact; print-color-adjust: exact; }
          .doc { font-size: 8.5px; }
        }
      `}</style>
    </div>
  )
}

function ItemRow({ si, depth = 0 }: { si: ScoreItem; depth?: number }) {
  const levels = (si.levels ?? []).slice().sort((a, b) => a.level - b.level)
  const scaleDef = getScaleDef(si.anamneseItemCode)
  const apps = applicability(si)
  const isScale = !!scaleDef
  const nQ = scaleDef?.questions?.length ?? (scaleDef?.administration?.type === 'word_recall' ? scaleDef.administration.categories.length : undefined)

  return (
    <div className="flex items-baseline gap-1.5 border-b border-zinc-100 py-[1px] leading-[1.2] break-inside-avoid">
      <span
        className="w-[25%] shrink-0 text-[8.5px] font-medium text-zinc-900"
        style={depth > 0 ? { paddingLeft: depth * 9, borderLeft: '2px solid #d4d4d8', marginLeft: 1 } : undefined}
      >
        {depth > 0 && <span className="mr-0.5 text-zinc-300">└</span>}
        {si.name}
        {si.unit && <span className="text-zinc-400"> ({si.unit})</span>}
        {apps && <span className="ml-1 text-[7px] text-zinc-400">{apps}</span>}
      </span>
      <span className="flex flex-1 flex-wrap items-center gap-x-[3px] gap-y-[1px]">
        {isScale && (
          <span className="rounded border border-zinc-300 bg-zinc-50 px-[3px] text-[7px] text-zinc-500">
            escala{nQ ? ` · ${nQ}q` : ''}
          </span>
        )}
        {levels.length > 0 ? (
          levels.map((l) => (
            <span key={l.id} className={cn('rounded border px-[3px] text-[7px] leading-tight', LEVEL_BG[l.level] ?? LEVEL_BG[6])}>
              <b>N{l.level}</b> {l.name}{levelRange(l) && <span className="opacity-60"> {levelRange(l)}</span>}
            </span>
          ))
        ) : si.unit ? (
          <span className="text-[7.5px] text-zinc-500">numérico ({si.unit})</span>
        ) : !isScale ? (
          <span className="text-[7.5px] text-zinc-300">resposta livre</span>
        ) : null}
      </span>
    </div>
  )
}
