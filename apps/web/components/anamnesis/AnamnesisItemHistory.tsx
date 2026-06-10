'use client'

import { TrendingUp, Loader2 } from 'lucide-react'
import { format } from 'date-fns'
import { ptBR } from 'date-fns/locale'
import { cn } from '@/lib/utils'
import { useItemHistory, type ClinicalTimelineSource } from '@/lib/api/clinical-timeline'
import type { ScoreLevel } from '@/lib/api/score-api'

// Cor do dot + badge por origem.
const SOURCE_STYLE: Record<ClinicalTimelineSource, { dot: string; badge: string }> = {
  escore_light: { dot: 'bg-violet-500', badge: 'bg-violet-100 text-violet-800' },
  pre_consulta: { dot: 'bg-sky-500', badge: 'bg-sky-100 text-sky-800' },
  anamnese: { dot: 'bg-amber-500', badge: 'bg-amber-100 text-amber-800' },
}

interface AnamnesisItemHistoryProps {
  patientId: string | null | undefined
  scoreItemId: string
  levels?: ScoreLevel[]
  unit?: string
  enabled: boolean
}

export function AnamnesisItemHistory({
  patientId,
  scoreItemId,
  levels = [],
  unit,
  enabled,
}: AnamnesisItemHistoryProps) {
  const { data, isLoading } = useItemHistory(patientId, scoreItemId, enabled)

  if (!patientId) return null

  const entries = data?.entries ?? []

  const levelName = (lvl?: number) =>
    lvl === undefined ? undefined : levels.find((l) => l.level === lvl)?.name

  // Valor exibido: "N2 · 1,4 L" / "N2 · Bom" / "1,4 L" — o que houver.
  const formatValue = (e: (typeof entries)[number]) => {
    const parts: string[] = []
    if (e.selectedLevel !== undefined) {
      const name = levelName(e.selectedLevel)
      parts.push(`N${e.selectedLevel}${name ? ` · ${name}` : ''}`)
    }
    if (e.numericValue !== undefined && e.numericValue !== null) {
      parts.push(`${e.numericValue}${unit ? ` ${unit}` : ''}`)
    }
    if (!parts.length && e.textValue) parts.push(e.textValue)
    return parts.join(' · ') || '—'
  }

  return (
    <div className="mt-3 border-t border-dashed pt-3">
      <div className="mb-2 flex items-center gap-1.5 text-[11px] font-semibold uppercase tracking-wide text-muted-foreground">
        <TrendingUp className="h-3.5 w-3.5" />
        Histórico deste item
      </div>

      {isLoading ? (
        <div className="flex items-center gap-2 py-1 text-xs text-muted-foreground">
          <Loader2 className="h-3.5 w-3.5 animate-spin" />
          Carregando…
        </div>
      ) : entries.length === 0 ? (
        <p className="py-1 text-xs text-muted-foreground">
          Sem registros anteriores para este item.
        </p>
      ) : (
        <div className="space-y-0.5">
          {entries.map((e, i) => {
            const style = SOURCE_STYLE[e.source] ?? SOURCE_STYLE.anamnese
            return (
              <div key={i} className="flex items-center gap-2 py-1">
                <span className={cn('h-2 w-2 shrink-0 rounded-full', style.dot)} />
                <span className={cn('shrink-0 rounded px-1.5 py-0.5 text-[10px] font-bold', style.badge)}>
                  {e.sourceLabel}
                </span>
                <span className="text-[11px] text-muted-foreground">
                  {format(new Date(e.date), 'dd/MM/yy', { locale: ptBR })}
                </span>
                <span className="ml-auto text-[11px] font-semibold text-foreground">
                  {formatValue(e)}
                </span>
              </div>
            )
          })}
        </div>
      )}
    </div>
  )
}
