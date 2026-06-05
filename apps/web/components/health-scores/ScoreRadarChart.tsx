"use client"

import { PatientScoreSnapshot } from "@/lib/api/health-score-api"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { RadarAgir } from "./RadarAgir"
import { buildAgir } from "./build-agir"

interface ScoreRadarChartProps {
  snapshot: PatientScoreSnapshot
}

/**
 * Painel de escore = MESMO radar AGIR da capa do paciente. Fonte única (buildAgir
 * a partir do snapshot). Não existe "radar por grupos" — o AGIR é o radar do EMR.
 */
export function ScoreRadarChart({ snapshot }: ScoreRadarChartProps) {
  const agir = buildAgir(snapshot)

  return (
    <Card>
      <CardHeader>
        <CardTitle>Visão Geral por Pilar</CardTitle>
        <CardDescription>
          Distribuição dos escores entre os pilares do Método AGIR
        </CardDescription>
      </CardHeader>
      <CardContent className="pt-6">
        {agir ? (
          <div className="w-full px-4 py-6 flex justify-center">
            <RadarAgir
              letters={agir.letters}
              pillars={agir.pillars}
              globalScore={snapshot.totalScorePercentage}
            />
          </div>
        ) : (
          <div className="flex items-center justify-center h-[300px] text-muted-foreground">
            Este escore não tem mapeamento de pilares AGIR.
          </div>
        )}
      </CardContent>
    </Card>
  )
}
