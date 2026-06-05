"use client"

import { PatientScoreSnapshot } from "@/lib/api/health-score-api"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { Activity } from "lucide-react"
import { format } from "date-fns"
import { ptBR } from "date-fns/locale"
import { RadarAgir } from "./RadarAgir"
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
 * Card PRIMÁRIA do escore: o radar AGIR é o elemento de destaque (grande, ~50%
 * em telas médias/grandes, 100% no mobile). O score total não vira uma card
 * separada — vai no centro do radar + os números de apoio no rodapé desta card.
 * Fonte única (buildAgir) — mesmo radar da capa e do portal.
 */
export function ScoreRadarChart({ snapshot }: ScoreRadarChartProps) {
  const agir = buildAgir(snapshot)
  const pct = snapshot.totalScorePercentage
  const color = scoreColor(pct)

  return (
    <Card>
      <CardHeader>
        <div className="flex items-start justify-between gap-3">
          <div>
            <CardTitle className="flex items-center gap-2">
              <Activity className="h-5 w-5 text-primary" />
              Escore Plenya
              <Badge
                variant="outline"
                style={{ backgroundColor: color + "20", color, borderColor: color }}
              >
                {scoreLabel(pct)}
              </Badge>
            </CardTitle>
            <CardDescription>
              Método AGIR · calculado em{" "}
              {format(new Date(snapshot.calculatedAt), "dd/MM/yyyy 'às' HH:mm", { locale: ptBR })}
            </CardDescription>
          </div>
        </div>
      </CardHeader>
      <CardContent>
        {agir ? (
          <div className="py-2">
            <RadarAgir
              letters={agir.letters}
              pillars={agir.pillars}
              globalScore={pct}
              widthStyle={{ width: "min(100%, 40rem, calc(100svh - 22rem))" }}
            />
          </div>
        ) : (
          <div className="flex items-center justify-center h-[300px] text-muted-foreground">
            Este escore não tem mapeamento de pilares AGIR.
          </div>
        )}

        {/* Rodapé — números de apoio (antes uma card de resumo redundante) */}
        <div className="mt-2 flex flex-wrap items-center justify-center gap-x-6 gap-y-1 border-t pt-4 text-sm text-muted-foreground">
          <span>
            <strong style={{ color }}>{pct.toFixed(1)}%</strong> geral
          </span>
          <span>
            {snapshot.totalActualPoints.toFixed(0)} / {snapshot.totalPossiblePoints.toFixed(0)} pontos
          </span>
          <span>{snapshot.itemsEvaluatedCount} itens avaliados</span>
          {snapshot.itemsNotEvaluatedCount > 0 && (
            <span>{snapshot.itemsNotEvaluatedCount} sem dados</span>
          )}
        </div>
      </CardContent>
    </Card>
  )
}
