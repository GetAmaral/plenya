"use client"

import { useMemo, useState } from "react"
import { PatientScoreSnapshot } from "@/lib/api/health-score-api"
import { useActivePatientSubscription } from "@/lib/api/subscription-api"
import { useRequireSelectedPatient } from "@/lib/use-require-selected-patient"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { ToggleGroup, ToggleGroupItem } from "@/components/ui/toggle-group"
import {
  Radar,
  RadarChart,
  PolarGrid,
  PolarAngleAxis,
  PolarRadiusAxis,
  ResponsiveContainer,
  Tooltip,
} from "recharts"
import { RadarAgir, type RadarLetter, type RadarPillar } from "./RadarAgir"

interface ScoreRadarChartProps {
  snapshot: PatientScoreSnapshot
}

export function ScoreRadarChart({ snapshot }: ScoreRadarChartProps) {
  const { selectedPatient } = useRequireSelectedPatient()
  const { data: activeSubscription } = useActivePatientSubscription(selectedPatient?.id)
  const [viewMode, setViewMode] = useState<'traditional' | 'methodology'>('methodology')
  const [hoveredPillar, setHoveredPillar] = useState<{ name: string; score: number; x: number; y: number } | null>(null)

  // Transform group results into chart data (Traditional)
  const groupChartData = snapshot.groupResults?.map((gr) => ({
    name: gr.group?.name || "Sem nome",
    score: parseFloat(gr.scorePercentage.toFixed(1)),
    fullMark: 100,
  })) || []

  // Transform pillar results into chart data (Methodology)
  const methodologyChartData = useMemo(() => {
    if (!snapshot?.itemResults || !activeSubscription?.subscriptionPlan?.method) {
      return null
    }

    const method = activeSubscription.subscriptionPlan.method
    if (!method.letters || method.letters.length === 0) return null

    // Estrutura para armazenar dados por letra
    const letterScores = new Map<string, {
      letterCode: string
      letterName: string
      letterColor: string
      pillars: Array<{
        pillarName: string
        score: number
        actualPoints: number
        possiblePoints: number
      }>
      totalActualPoints: number
      totalPossiblePoints: number
    }>()

    // Processar cada letra e seus pilares
    method.letters.forEach(letter => {
      const letterData = {
        letterCode: letter.code,
        letterName: letter.name,
        letterColor: (letter as any).color || '#94a3b8',
        pillars: [] as any[],
        totalActualPoints: 0,
        totalPossiblePoints: 0,
      }

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
          letterData.pillars.push({
            pillarName: pillar.name,
            score: parseFloat(score.toFixed(1)),
            actualPoints,
            possiblePoints,
          })
          letterData.totalActualPoints += actualPoints
          letterData.totalPossiblePoints += possiblePoints
        }
      })

      if (letterData.pillars.length > 0) {
        letterScores.set(letter.code, letterData)
      }
    })

    // Criar dados unificados com um ponto por pilar
    const chartData: Array<{
      axis: string
      pillarScore: number
      letterCode: string
      letterColor: string
      fullMark: number
    }> = []

    // Mapa para calcular posições das letras
    const letterPositions = new Map<string, {
      indices: number[]
      letterCode: string
      letterColor: string
    }>()

    let currentIndex = 0
    letterScores.forEach((letterData, letterCode) => {
      const indices: number[] = []

      letterData.pillars.forEach(pillar => {
        chartData.push({
          axis: pillar.pillarName,
          pillarScore: pillar.score,
          letterCode: letterData.letterCode,
          letterColor: letterData.letterColor,
          fullMark: 100,
        })
        indices.push(currentIndex)
        currentIndex++
      })

      letterPositions.set(letterCode, {
        indices,
        letterCode: letterData.letterCode,
        letterColor: letterData.letterColor,
      })
    })

    // Dados para o novo RadarAgir (SVG puro)
    const radarLetters: RadarLetter[] = []
    const radarPillars: RadarPillar[] = []
    letterScores.forEach((letterData) => {
      const letterScore = letterData.totalPossiblePoints > 0
        ? (letterData.totalActualPoints / letterData.totalPossiblePoints) * 100
        : 0
      radarLetters.push({
        code: letterData.letterCode,
        name: letterData.letterName,
        score: parseFloat(letterScore.toFixed(1)),
        color: letterData.letterColor,
      })
      letterData.pillars.forEach((p) => {
        radarPillars.push({
          letter: letterData.letterCode,
          name: p.pillarName,
          score: p.score,
        })
      })
    })

    return { chartData, letterPositions, radarLetters, radarPillars }
  }, [snapshot, activeSubscription])

  const isMethodologyView = viewMode === 'methodology' && methodologyChartData

  const getScoreColor = (percentage: number) => {
    if (percentage >= 86) return "#22c55e" // Verde
    if (percentage >= 71) return "#3b82f6" // Azul
    if (percentage >= 51) return "#eab308" // Amarelo
    if (percentage >= 31) return "#f97316" // Laranja
    return "#ef4444" // Vermelho
  }

  const avgScore = isMethodologyView && methodologyChartData
    ? methodologyChartData.chartData.reduce((sum, item) => sum + item.pillarScore, 0) / methodologyChartData.chartData.length
    : groupChartData.length > 0
    ? groupChartData.reduce((sum, item) => sum + item.score, 0) / groupChartData.length
    : 0

  // Função para calcular posição angular de uma letra baseada nos índices de seus pilares
  const calculateLetterAngle = (indices: number[], totalItems: number) => {
    if (indices.length === 0) return 0

    // Ângulo inicial é -90 graus (topo) em radianos
    const startAngle = -Math.PI / 2
    // Cada item ocupa 360/totalItems graus
    const anglePerItem = (2 * Math.PI) / totalItems

    // Para 1 pilar: usar o ângulo dele
    // Para 2+ pilares: usar a média dos ângulos
    const avgIndex = indices.reduce((sum, idx) => sum + idx, 0) / indices.length

    return startAngle + (avgIndex * anglePerItem)
  }

  return (
    <Card>
      <CardHeader>
        <div className="flex items-start justify-between">
          <div>
            <CardTitle>
              {viewMode === 'methodology' && methodologyChartData
                ? 'Visão Geral por Pilar'
                : 'Visão Geral por Grupo'}
            </CardTitle>
            <CardDescription>
              {viewMode === 'methodology' && methodologyChartData
                ? 'Distribuição dos escores entre os pilares da metodologia'
                : 'Distribuição dos escores entre os diferentes grupos clínicos'}
            </CardDescription>
          </div>

          {/* Toggle only if methodology available */}
          {methodologyChartData && (
            <ToggleGroup
              type="single"
              value={viewMode}
              onValueChange={(v) => v && setViewMode(v as any)}
              className="hidden sm:flex"
            >
              <ToggleGroupItem value="methodology" aria-label="Visualização por Metodologia" size="sm">
                <span className="text-xs">Metodologia</span>
              </ToggleGroupItem>
              <ToggleGroupItem value="traditional" aria-label="Visualização Tradicional" size="sm">
                <span className="text-xs">Grupos</span>
              </ToggleGroupItem>
            </ToggleGroup>
          )}
        </div>
      </CardHeader>
      <CardContent className="pt-6">
        {(isMethodologyView ? !methodologyChartData || methodologyChartData.chartData.length === 0 : groupChartData.length === 0) ? (
          <div className="flex items-center justify-center h-[500px] text-muted-foreground">
            Nenhum dado disponível
          </div>
        ) : isMethodologyView && methodologyChartData ? (
          /* Methodology View - RadarAgir (SVG puro, igual ao site) */
          <div className="w-full px-4 py-6 flex justify-center">
            <RadarAgir
              letters={methodologyChartData.radarLetters}
              pillars={methodologyChartData.radarPillars}
              globalScore={avgScore}
            />
          </div>
        ) : (
          /* Traditional View - Grupos */
          <div className="w-full h-[500px] px-4 relative">
            <ResponsiveContainer width="100%" height="100%">
              <RadarChart
                data={groupChartData}
                margin={{ top: 20, right: 80, bottom: 20, left: 80 }}
              >
                <PolarGrid stroke="#e5e7eb" />
                <PolarAngleAxis
                  dataKey="name"
                  tick={{ fill: "#6b7280", fontSize: 11 }}
                  tickLine={false}
                />
                <PolarRadiusAxis
                  angle={90}
                  domain={[0, 100]}
                  tick={{ fill: "#6b7280", fontSize: 10 }}
                  tickLine={false}
                />
                <Radar
                  name="Score"
                  dataKey="score"
                  stroke={getScoreColor(avgScore)}
                  fill={getScoreColor(avgScore)}
                  fillOpacity={0.6}
                  strokeWidth={2}
                />
                <Tooltip
                  contentStyle={{
                    backgroundColor: "white",
                    border: "1px solid #e5e7eb",
                    borderRadius: "8px",
                    padding: "8px 12px",
                  }}
                  formatter={(value: number) => [`${value.toFixed(1)}%`, "Score"]}
                />
              </RadarChart>
            </ResponsiveContainer>

            {/* Tooltip customizado para labels dos grupos */}
            {hoveredPillar && (
              <div
                style={{
                  position: 'fixed',
                  left: hoveredPillar.x,
                  top: hoveredPillar.y,
                  transform: 'translate(-50%, -100%)',
                  backgroundColor: 'white',
                  border: '1px solid #e5e7eb',
                  borderRadius: '8px',
                  padding: '8px 12px',
                  boxShadow: '0 2px 8px rgba(0,0,0,0.1)',
                  pointerEvents: 'none',
                  zIndex: 1000,
                  whiteSpace: 'nowrap'
                }}
              >
                <div style={{ fontSize: '12px', fontWeight: 600, marginBottom: '4px' }}>
                  {hoveredPillar.name}
                </div>
                <div style={{ fontSize: '14px', fontWeight: 'bold' }}>
                  {hoveredPillar.score.toFixed(1)}%
                </div>
              </div>
            )}
          </div>
        )}
      </CardContent>
    </Card>
  )
}
