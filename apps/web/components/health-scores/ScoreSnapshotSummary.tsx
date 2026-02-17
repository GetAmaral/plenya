"use client"

import { useMemo, useState } from "react"
import { PatientScoreSnapshot } from "@/lib/api/health-score-api"
import { useActivePatientSubscription } from "@/lib/api/subscription-api"
import { useRequireSelectedPatient } from "@/lib/use-require-selected-patient"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Progress } from "@/components/ui/progress"
import { Badge } from "@/components/ui/badge"
import { ToggleGroup, ToggleGroupItem } from "@/components/ui/toggle-group"
import { format } from "date-fns"
import { ptBR } from "date-fns/locale"
import { TrendingUp, TrendingDown, Minus, Activity } from "lucide-react"

interface ScoreSnapshotSummaryProps {
  snapshot: PatientScoreSnapshot
  showDetailCards?: boolean // Se false, mostra apenas o card principal
  children?: React.ReactNode // Conteúdo inserido entre o card principal e os cards de detalhes
}

export function ScoreSnapshotSummary({ snapshot, showDetailCards = true, children }: ScoreSnapshotSummaryProps) {
  const { selectedPatient } = useRequireSelectedPatient()
  const { data: activeSubscription } = useActivePatientSubscription(selectedPatient?.id)
  const [viewMode, setViewMode] = useState<'traditional' | 'methodology'>('methodology')
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

  // Organize by methodology (Method → Letter → Pillar)
  const pillarScores = useMemo(() => {
    if (!snapshot?.itemResults || !activeSubscription?.subscriptionPlan?.method) {
      return null
    }

    const method = activeSubscription.subscriptionPlan.method
    if (!method.letters || method.letters.length === 0) return null

    const pillars: Array<{
      id: string
      name: string
      letterCode: string
      letterName: string
      letterColor?: string
      score: number
      itemCount: number
      evaluatedCount: number
      actualPoints: number
      possiblePoints: number
    }> = []

    method.letters.forEach(letter => {
      letter.pillars?.forEach(pillar => {
        // Find items for this pillar
        const pillarItems = snapshot.itemResults.filter(ir =>
          ir.item?.methodPillars?.some(mp => mp.id === pillar.id)
        )

        const evaluatedItems = pillarItems.filter(item => item.status === 'evaluated')
        const actualPoints = evaluatedItems.reduce((sum, item) => sum + item.actualPoints, 0)
        const possiblePoints = evaluatedItems.reduce((sum, item) => sum + item.maxPoints, 0)
        const score = possiblePoints > 0 ? (actualPoints / possiblePoints) * 100 : 0

        if (pillarItems.length > 0) {
          pillars.push({
            id: pillar.id,
            name: pillar.name,
            letterCode: letter.code,
            letterName: letter.name,
            letterColor: (letter as any).color,
            score,
            itemCount: pillarItems.length,
            evaluatedCount: evaluatedItems.length,
            actualPoints,
            possiblePoints,
          })
        }
      })
    })

    return {
      method,
      pillars: pillars.sort((a, b) => a.letterCode.localeCompare(b.letterCode))
    }
  }, [snapshot, activeSubscription])

  const shouldShowMethodologyView = viewMode === 'methodology' && pillarScores !== null

  return (
    <>
      {/* Card Principal: Score Total */}
      <Card className="w-full">
        <CardHeader>
          <div className="flex items-center justify-between">
            <div>
              <CardTitle className="flex items-center gap-2">
                <Activity className="h-5 w-5" />
                Score Total de Saúde
                {pillarScores && (
                  <span className="text-sm font-normal text-muted-foreground">
                    • {pillarScores.method.name}
                  </span>
                )}
              </CardTitle>
              <p className="text-sm text-muted-foreground mt-1">
                Calculado em{" "}
                {format(new Date(snapshot.calculatedAt), "dd/MM/yyyy 'às' HH:mm", {
                  locale: ptBR,
                })}
              </p>
            </div>

            {/* Toggle only if methodology available */}
            {pillarScores && (
              <ToggleGroup
                type="single"
                value={viewMode}
                onValueChange={(v) => v && setViewMode(v as any)}
                className="hidden sm:flex"
              >
                <ToggleGroupItem value="methodology" aria-label="Visualização por Metodologia" size="sm">
                  <span className="text-xs">Por Metodologia</span>
                </ToggleGroupItem>
                <ToggleGroupItem value="traditional" aria-label="Visualização Tradicional" size="sm">
                  <span className="text-xs">Por Grupos</span>
                </ToggleGroupItem>
              </ToggleGroup>
            )}
          </div>
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

      {/* Children (e.g., Radar Chart) inserido entre card principal e cards de detalhes */}
      {children}

      {/* Cards de detalhes */}
      {showDetailCards && (
        <div className="grid gap-4 md:grid-cols-3">
          {shouldShowMethodologyView ? (
        /* Cards por Pilar (Metodologia) */
        pillarScores!.pillars.map((pillar) => (
          <Card key={pillar.id} className="border-l-4" style={{ borderLeftColor: pillar.letterColor || 'hsl(var(--border))' }}>
            <CardHeader>
              <div className="flex items-center gap-2">
                <Badge variant="outline" className="text-xs">
                  {pillar.letterCode}
                </Badge>
                <CardTitle className="text-base">{pillar.name}</CardTitle>
              </div>
              <p className="text-xs text-muted-foreground">{pillar.letterName}</p>
            </CardHeader>
            <CardContent>
              <div className="space-y-2">
                <div
                  className="text-3xl font-bold"
                  style={{ color: getScoreColor(pillar.score) }}
                >
                  {pillar.score.toFixed(1)}%
                </div>

                <Progress value={pillar.score} className="h-2" />

                <div className="flex items-center justify-between text-xs text-muted-foreground">
                  <span>{pillar.evaluatedCount} / {pillar.itemCount} itens</span>
                  <span>
                    {pillar.actualPoints.toFixed(0)} / {pillar.possiblePoints.toFixed(0)}{" "}
                    pts
                  </span>
                </div>
              </div>
            </CardContent>
          </Card>
        ))
      ) : (
        /* Cards por Grupo (Tradicional) */
        snapshot.groupResults?.map((groupResult) => (
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
        ))
          )}
        </div>
      )}
    </>
  )
}
