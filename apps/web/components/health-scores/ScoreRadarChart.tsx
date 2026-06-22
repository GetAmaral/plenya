"use client"

import { PatientScoreSnapshot } from "@/lib/api/health-score-api"
import { Card, CardContent } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { Activity } from "lucide-react"
import { formatDate } from "@/lib/format-date"
import { RadarAgir, resolveColor } from "./RadarAgir"
import { buildAgir } from "./build-agir"

interface ScoreRadarChartProps {
  snapshot: PatientScoreSnapshot
}

function scoreLabel(p: number) {
  if (p >= 86) return "Ótimo"
  if (p >= 71) return "Bom"
  if (p >= 51) return "Regular"
  if (p >= 31) return "Atenção"
  return "Crítico"
}
function scoreColor(p: number) {
  if (p >= 86) return "hsl(142, 76%, 36%)"
  if (p >= 71) return "hsl(217, 91%, 60%)"
  if (p >= 51) return "hsl(48, 96%, 53%)"
  if (p >= 31) return "hsl(25, 95%, 53%)"
  return "hsl(0, 84%, 60%)"
}

/**
 * Card PRIMÁRIA do escore. O radar AGIR ocupa a ALTURA da área útil da tela
 * (tamanho = min(largura disponível, altura útil) — quadrado), e as infos de apoio
 * ficam à DIREITA (no mobile, empilham abaixo). Fonte única (buildAgir).
 */
export function ScoreRadarChart({ snapshot }: ScoreRadarChartProps) {
  const agir = buildAgir(snapshot)
  const pct = snapshot.totalScorePercentage
  const color = scoreColor(pct)

  // pilares visíveis por letra (pra mostrar "—" em letra vazia)
  const visibleByLetter = new Map<string, number>()
  agir?.pillars.forEach((p: any) => visibleByLetter.set(p.letter, (visibleByLetter.get(p.letter) ?? 0) + 1))

  if (!agir) {
    return (
      <Card>
        <CardContent className="flex h-[300px] items-center justify-center text-muted-foreground">
          Este escore não tem mapeamento de pilares AGIR.
        </CardContent>
      </Card>
    )
  }

  return (
    <Card>
      <CardContent className="p-4 lg:p-6">
        <div className="flex flex-col gap-6 lg:flex-row lg:items-center">
          {/* Radar — altura = área útil da tela; quadrado. No mobile, largura cheia. */}
          <div
            className="mx-auto w-full lg:mx-0 lg:shrink-0"
            style={{ width: "min(100%, calc(100svh - 14.5rem))" }}
          >
            <RadarAgir
              letters={agir.letters}
              pillars={agir.pillars}
              globalScore={pct}
              maxWidthClass="max-w-none"
              showLegend={false}
            />
          </div>

          {/* Infos — à direita (desktop) / abaixo (mobile) */}
          <div className="flex flex-1 flex-col justify-center gap-5">
            <div>
              <div className="flex flex-wrap items-center gap-2">
                <Activity className="h-5 w-5 text-primary" />
                <h2 className="text-lg font-semibold">Escore Plenya</h2>
                <Badge
                  variant="outline"
                  style={{ backgroundColor: color + "20", color, borderColor: color }}
                >
                  {scoreLabel(pct)}
                </Badge>
              </div>
              <p className="mt-1 text-sm text-muted-foreground">
                Método AGIR · calculado em{" "}
                {formatDate(snapshot.calculatedAt, "dd/MM/yyyy 'às' HH:mm")}
              </p>
            </div>

            <div className="border-t pt-4">
              <div className="text-4xl font-bold" style={{ color }}>
                {pct.toFixed(1)}%
              </div>
              <div className="text-sm text-muted-foreground">escore geral</div>
            </div>

            <dl className="grid grid-cols-3 gap-4 border-t pt-4 text-sm">
              <div>
                <dt className="text-muted-foreground">Pontos</dt>
                <dd className="font-medium text-foreground">
                  {snapshot.totalActualPoints.toFixed(0)} / {snapshot.totalPossiblePoints.toFixed(0)}
                </dd>
              </div>
              <div>
                <dt className="text-muted-foreground">Avaliados</dt>
                <dd className="font-medium text-foreground">{snapshot.itemsEvaluatedCount}</dd>
              </div>
              <div>
                <dt className="text-muted-foreground">Sem dados</dt>
                <dd className="font-medium text-foreground">{snapshot.itemsNotEvaluatedCount}</dd>
              </div>
            </dl>

            {/* Letras (movidas da legenda de baixo do radar) */}
            <div className="flex flex-wrap gap-x-4 gap-y-2 border-t pt-4 font-mono text-[11px] uppercase tracking-[0.18em]">
              {agir.letters.map((l: any) => {
                const vis = visibleByLetter.get(l.code) ?? 0
                const c = resolveColor(l.code, l.color)
                return (
                  <span key={l.code} className="flex items-center gap-2 text-foreground/70">
                    <span className="h-2 w-2 rounded-full" style={{ background: c, opacity: vis === 0 ? 0.5 : 1 }} />
                    <span className={vis === 0 ? "text-foreground/40" : undefined}>
                      {l.code} {vis === 0 ? "—" : l.score.toFixed(0)}
                    </span>
                  </span>
                )
              })}
            </div>
          </div>
        </div>
      </CardContent>
    </Card>
  )
}
