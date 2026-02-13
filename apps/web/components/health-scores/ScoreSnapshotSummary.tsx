"use client"

import { PatientScoreSnapshot } from "@/lib/api/health-score-api"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Progress } from "@/components/ui/progress"
import { Badge } from "@/components/ui/badge"
import { format } from "date-fns"
import { ptBR } from "date-fns/locale"
import { TrendingUp, TrendingDown, Minus, Activity } from "lucide-react"

interface ScoreSnapshotSummaryProps {
  snapshot: PatientScoreSnapshot
}

export function ScoreSnapshotSummary({ snapshot }: ScoreSnapshotSummaryProps) {
  const getScoreColor = (percentage: number) => {
    if (percentage >= 86) return "hsl(142, 76%, 36%)" // Verde
    if (percentage >= 71) return "hsl(217, 91%, 60%)" // Azul
    if (percentage >= 51) return "hsl(48, 96%, 53%)" // Amarelo
    if (percentage >= 31) return "hsl(25, 95%, 53%)" // Laranja
    return "hsl(0, 84%, 60%)" // Vermelho
  }

  const getScoreLabel = (percentage: number) => {
    if (percentage >= 86) return "Ótimo"
    if (percentage >= 71) return "Bom"
    if (percentage >= 51) return "Regular"
    if (percentage >= 31) return "Atenção"
    return "Crítico"
  }

  return (
    <div className="grid gap-4 md:grid-cols-3">
      {/* Card Principal: Score Total */}
      <Card className="md:col-span-3">
        <CardHeader>
          <CardTitle className="flex items-center gap-2">
            <Activity className="h-5 w-5" />
            Score Total de Saúde
          </CardTitle>
          <p className="text-sm text-muted-foreground">
            Calculado em{" "}
            {format(new Date(snapshot.calculatedAt), "dd/MM/yyyy 'às' HH:mm", {
              locale: ptBR,
            })}
          </p>
        </CardHeader>
        <CardContent>
          <div className="space-y-4">
            <div className="flex items-baseline gap-2">
              <span
                className="text-5xl font-bold"
                style={{ color: getScoreColor(snapshot.totalScorePercentage) }}
              >
                {snapshot.totalScorePercentage.toFixed(1)}%
              </span>
              <Badge
                variant="outline"
                style={{
                  backgroundColor: getScoreColor(snapshot.totalScorePercentage) + "20",
                  color: getScoreColor(snapshot.totalScorePercentage),
                  borderColor: getScoreColor(snapshot.totalScorePercentage),
                }}
              >
                {getScoreLabel(snapshot.totalScorePercentage)}
              </Badge>
            </div>

            <Progress
              value={snapshot.totalScorePercentage}
              className="h-3"
              style={
                {
                  "--progress-background": getScoreColor(snapshot.totalScorePercentage),
                } as React.CSSProperties
              }
            />

            <div className="flex items-center gap-4 text-sm text-muted-foreground">
              <span>
                {snapshot.totalActualPoints.toFixed(1)} / {snapshot.totalPossiblePoints.toFixed(1)}{" "}
                pontos
              </span>
              <span>•</span>
              <span>{snapshot.itemsEvaluatedCount} itens avaliados</span>
              {snapshot.itemsNotEvaluatedCount > 0 && (
                <>
                  <span>•</span>
                  <span>{snapshot.itemsNotEvaluatedCount} sem dados</span>
                </>
              )}
            </div>

            {snapshot.notes && (
              <div className="mt-4 p-3 bg-muted rounded-lg">
                <p className="text-sm">
                  <strong>Observações:</strong> {snapshot.notes}
                </p>
              </div>
            )}
          </div>
        </CardContent>
      </Card>

      {/* Cards por Grupo */}
      {snapshot.groupResults?.map((groupResult) => (
        <Card key={groupResult.id}>
          <CardHeader>
            <CardTitle className="text-base">{groupResult.group?.name}</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="space-y-2">
              <div
                className="text-3xl font-bold"
                style={{ color: getScoreColor(groupResult.scorePercentage) }}
              >
                {groupResult.scorePercentage.toFixed(1)}%
              </div>

              <Progress value={groupResult.scorePercentage} className="h-2" />

              <div className="flex items-center justify-between text-xs text-muted-foreground">
                <span>{groupResult.itemsEvaluatedCount} itens</span>
                <span>
                  {groupResult.actualPoints.toFixed(0)} / {groupResult.possiblePoints.toFixed(0)}{" "}
                  pts
                </span>
              </div>
            </div>
          </CardContent>
        </Card>
      ))}
    </div>
  )
}
