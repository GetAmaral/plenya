"use client"

import { PatientScoreSnapshot } from "@/lib/api/health-score-api"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import {
  Radar,
  RadarChart,
  PolarGrid,
  PolarAngleAxis,
  PolarRadiusAxis,
  ResponsiveContainer,
  Tooltip,
} from "recharts"

interface ScoreRadarChartProps {
  snapshot: PatientScoreSnapshot
}

export function ScoreRadarChart({ snapshot }: ScoreRadarChartProps) {
  // Transform group results into chart data
  const chartData = snapshot.groupResults?.map((gr) => ({
    group: gr.group?.name || "Sem nome",
    score: parseFloat(gr.scorePercentage.toFixed(1)),
    fullMark: 100,
  })) || []

  const getScoreColor = (percentage: number) => {
    if (percentage >= 86) return "#22c55e" // Verde
    if (percentage >= 71) return "#3b82f6" // Azul
    if (percentage >= 51) return "#eab308" // Amarelo
    if (percentage >= 31) return "#f97316" // Laranja
    return "#ef4444" // Vermelho
  }

  const avgScore = chartData.length > 0
    ? chartData.reduce((sum, item) => sum + item.score, 0) / chartData.length
    : 0

  return (
    <Card>
      <CardHeader>
        <CardTitle>Visão Geral por Grupo</CardTitle>
        <CardDescription>
          Distribuição dos escores entre os diferentes grupos clínicos
        </CardDescription>
      </CardHeader>
      <CardContent>
        {chartData.length === 0 ? (
          <div className="flex items-center justify-center h-[400px] text-muted-foreground">
            Nenhum dado disponível
          </div>
        ) : (
          <ResponsiveContainer width="100%" height={400}>
            <RadarChart data={chartData}>
              <PolarGrid stroke="#e5e7eb" />
              <PolarAngleAxis
                dataKey="group"
                tick={{ fill: "#6b7280", fontSize: 12 }}
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
        )}
      </CardContent>
    </Card>
  )
}
