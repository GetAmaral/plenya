"use client"

import { useState, useMemo } from "react"
import { PatientScoreSnapshot } from "@/lib/api/health-score-api"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Checkbox } from "@/components/ui/checkbox"
import { Label } from "@/components/ui/label"
import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend,
  ResponsiveContainer,
} from "recharts"
import { format, subDays, subMonths, subYears, isAfter } from "date-fns"
import { ptBR } from "date-fns/locale"

interface ScoreEvolutionChartProps {
  snapshots: PatientScoreSnapshot[]
}

type TimeRange = "7d" | "30d" | "90d" | "1y" | "all"

const TIME_RANGE_LABELS: Record<TimeRange, string> = {
  "7d": "7 dias",
  "30d": "30 dias",
  "90d": "90 dias",
  "1y": "1 ano",
  "all": "Tudo",
}

// Cores para os grupos (paleta consistente)
const GROUP_COLORS = [
  "#3b82f6", // Azul
  "#22c55e", // Verde
  "#f97316", // Laranja
  "#a855f7", // Roxo
  "#ec4899", // Rosa
  "#14b8a6", // Teal
  "#f59e0b", // Âmbar
  "#6366f1", // Indigo
]

export function ScoreEvolutionChart({ snapshots }: ScoreEvolutionChartProps) {
  const [timeRange, setTimeRange] = useState<TimeRange>("30d")
  const [selectedGroups, setSelectedGroups] = useState<string[]>([])

  // Extract unique groups from all snapshots
  const uniqueGroups = useMemo(() => {
    const groupMap = new Map<string, { id: string; name: string; color: string }>()
    let colorIndex = 0

    snapshots.forEach((snapshot) => {
      snapshot.groupResults?.forEach((gr) => {
        if (gr.group && !groupMap.has(gr.group.id)) {
          groupMap.set(gr.group.id, {
            id: gr.group.id,
            name: gr.group.name,
            color: GROUP_COLORS[colorIndex % GROUP_COLORS.length],
          })
          colorIndex++
        }
      })
    })

    return Array.from(groupMap.values())
  }, [snapshots])

  // Filter snapshots by time range
  const filteredSnapshots = useMemo(() => {
    if (timeRange === "all") return snapshots

    const now = new Date()
    const cutoffDate = {
      "7d": subDays(now, 7),
      "30d": subDays(now, 30),
      "90d": subDays(now, 90),
      "1y": subYears(now, 1),
    }[timeRange]

    return snapshots.filter((s) => isAfter(new Date(s.calculatedAt), cutoffDate))
  }, [snapshots, timeRange])

  // Transform snapshots into chart data
  const chartData = useMemo(() => {
    return filteredSnapshots
      .map((snapshot) => {
        const dataPoint: any = {
          date: format(new Date(snapshot.calculatedAt), "dd/MM", { locale: ptBR }),
          fullDate: new Date(snapshot.calculatedAt),
          total: parseFloat(snapshot.totalScorePercentage.toFixed(1)),
        }

        // Add each group's score
        snapshot.groupResults?.forEach((gr) => {
          if (gr.group) {
            dataPoint[gr.group.id] = parseFloat(gr.scorePercentage.toFixed(1))
          }
        })

        return dataPoint
      })
      .sort((a, b) => a.fullDate.getTime() - b.fullDate.getTime())
  }, [filteredSnapshots])

  const toggleGroup = (groupId: string) => {
    setSelectedGroups((prev) =>
      prev.includes(groupId) ? prev.filter((id) => id !== groupId) : [...prev, groupId]
    )
  }

  const getGroupColor = (groupId: string) => {
    return uniqueGroups.find((g) => g.id === groupId)?.color || "#6b7280"
  }

  const getGroupName = (groupId: string) => {
    return uniqueGroups.find((g) => g.id === groupId)?.name || groupId
  }

  return (
    <Card>
      <CardHeader>
        <div className="flex items-center justify-between flex-wrap gap-4">
          <div>
            <CardTitle>Evolução Temporal</CardTitle>
            <CardDescription>
              Acompanhe a evolução dos escores ao longo do tempo
            </CardDescription>
          </div>

          {/* Time range selector */}
          <div className="flex gap-2">
            {(Object.keys(TIME_RANGE_LABELS) as TimeRange[]).map((range) => (
              <Button
                key={range}
                variant={timeRange === range ? "default" : "outline"}
                size="sm"
                onClick={() => setTimeRange(range)}
              >
                {TIME_RANGE_LABELS[range]}
              </Button>
            ))}
          </div>
        </div>

        {/* Group filters */}
        {uniqueGroups.length > 0 && (
          <div className="flex flex-wrap gap-4 mt-4 pt-4 border-t">
            {uniqueGroups.map((group) => (
              <div key={group.id} className="flex items-center space-x-2">
                <Checkbox
                  id={`group-${group.id}`}
                  checked={selectedGroups.includes(group.id)}
                  onCheckedChange={() => toggleGroup(group.id)}
                />
                <Label
                  htmlFor={`group-${group.id}`}
                  className="text-sm font-normal cursor-pointer flex items-center gap-2"
                >
                  <div
                    className="w-3 h-3 rounded-full"
                    style={{ backgroundColor: group.color }}
                  />
                  {group.name}
                </Label>
              </div>
            ))}
          </div>
        )}
      </CardHeader>

      <CardContent>
        {chartData.length === 0 ? (
          <div className="flex items-center justify-center h-[300px] text-muted-foreground">
            Nenhum dado disponível para o período selecionado
          </div>
        ) : (
          <ResponsiveContainer width="100%" height={300}>
            <LineChart data={chartData}>
              <CartesianGrid strokeDasharray="3 3" stroke="#e5e7eb" />
              <XAxis
                dataKey="date"
                tick={{ fill: "#6b7280", fontSize: 12 }}
                tickLine={false}
              />
              <YAxis
                domain={[0, 100]}
                tick={{ fill: "#6b7280", fontSize: 12 }}
                tickLine={false}
                label={{ value: "Score (%)", angle: -90, position: "insideLeft" }}
              />
              <Tooltip
                contentStyle={{
                  backgroundColor: "white",
                  border: "1px solid #e5e7eb",
                  borderRadius: "8px",
                  padding: "8px 12px",
                }}
                formatter={(value: number) => `${value.toFixed(1)}%`}
                labelFormatter={(label) => `Data: ${label}`}
              />
              <Legend />

              {/* Total line (always visible) */}
              <Line
                type="monotone"
                dataKey="total"
                stroke="#000"
                strokeWidth={3}
                name="Total"
                dot={{ r: 4 }}
                activeDot={{ r: 6 }}
              />

              {/* Group lines (based on selection) */}
              {selectedGroups.map((groupId) => (
                <Line
                  key={groupId}
                  type="monotone"
                  dataKey={groupId}
                  stroke={getGroupColor(groupId)}
                  strokeWidth={2}
                  name={getGroupName(groupId)}
                  dot={{ r: 3 }}
                  activeDot={{ r: 5 }}
                />
              ))}
            </LineChart>
          </ResponsiveContainer>
        )}
      </CardContent>
    </Card>
  )
}
