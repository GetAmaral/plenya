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

    return { chartData, letterPositions }
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
          /* Methodology View - Um único radar com labels de letras externas */
          <div className="w-full h-[500px] px-4 relative">
            <ResponsiveContainer width="100%" height="100%">
              <RadarChart
                data={methodologyChartData.chartData}
                margin={{ top: 60, right: 120, bottom: 60, left: 120 }}
              >
                <PolarGrid stroke="#e5e7eb" />

                {/* Labels dos pilares (alinhamento dinâmico) */}
                <PolarAngleAxis
                  dataKey="axis"
                  tick={(props) => {
                    const { x, y, payload, index, cx, cy } = props
                    const data = methodologyChartData.chartData[index]

                    // Calcular ângulo do pilar em relação ao centro
                    const dx = x - cx
                    const dy = y - cy
                    const angle = Math.atan2(dy, dx) * (180 / Math.PI)

                    // Normalizar ângulo para 0-360
                    const normalizedAngle = (angle + 360) % 360

                    // Estimar largura do texto (fontSize 14 = ~8px por char)
                    const textLength = payload.value.length
                    const rectWidth = Math.max(textLength * 8 + 16, 60)
                    const rectHeight = 28

                    // Determinar posição do badge baseado no ângulo
                    // O badge deve ficar FORA do radar, não sobrepor
                    let badgeCenterX = x
                    let badgeCenterY = y

                    if (normalizedAngle >= 45 && normalizedAngle < 135) {
                      // Embaixo - badge.top alinhado ao ponto
                      badgeCenterY = y + rectHeight / 2 + 8
                    } else if (normalizedAngle >= 135 && normalizedAngle < 225) {
                      // Esquerda - badge.right alinhado ao ponto
                      badgeCenterX = x - rectWidth / 2 - 8
                    } else if (normalizedAngle >= 225 && normalizedAngle < 315) {
                      // Topo - badge.bottom alinhado ao ponto
                      badgeCenterY = y - rectHeight / 2 - 8
                    } else {
                      // Direita - badge.left alinhado ao ponto
                      badgeCenterX = x + rectWidth / 2 + 8
                    }

                    const rectX = badgeCenterX - rectWidth / 2
                    const rectY = badgeCenterY - rectHeight / 2

                    return (
                      <g>
                        {/* Fundo com mesma cor, transparente */}
                        <rect
                          x={rectX}
                          y={rectY}
                          width={rectWidth}
                          height={rectHeight}
                          fill={data.letterColor}
                          opacity={0.15}
                          stroke={data.letterColor}
                          strokeWidth={2}
                          rx={4}
                        />
                        {/* Texto sempre centralizado no badge */}
                        <text
                          x={badgeCenterX}
                          y={badgeCenterY + 1}
                          textAnchor="middle"
                          dominantBaseline="middle"
                          fill={data.letterColor}
                          fontSize={14}
                          fontWeight="600"
                        >
                          {payload.value}
                        </text>
                        {/* Área invisível por cima para capturar hover (deve ser último) */}
                        <rect
                          x={rectX - 4}
                          y={rectY - 4}
                          width={rectWidth + 8}
                          height={rectHeight + 8}
                          fill="transparent"
                          style={{ cursor: 'default' }}
                          onMouseEnter={(e) => {
                            const rect = e.currentTarget.getBoundingClientRect()
                            setHoveredPillar({
                              name: payload.value,
                              score: data.pillarScore,
                              x: rect.left + rect.width / 2,
                              y: rect.top - 10
                            })
                          }}
                          onMouseLeave={() => setHoveredPillar(null)}
                        />
                      </g>
                    )
                  }}
                  tickLine={false}
                />

                <PolarRadiusAxis
                  angle={90}
                  domain={[0, 100]}
                  tick={{ fill: "#6b7280", fontSize: 10 }}
                  tickLine={false}
                />

                {/* Um único radar com os scores dos pilares */}
                <Radar
                  name="Score do Pilar"
                  dataKey="pillarScore"
                  stroke={getScoreColor(avgScore)}
                  fill={getScoreColor(avgScore)}
                  fillOpacity={0.5}
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

            {/* SVG Overlay para badges das letras e áreas coloridas */}
            <svg
              className="absolute inset-0 pointer-events-none"
              style={{ width: '100%', height: '100%' }}
              viewBox="0 0 100 100"
              preserveAspectRatio="xMidYMid meet"
            >
              {/* Áreas coloridas de fundo para cada letra */}
              {Array.from(methodologyChartData.letterPositions.entries()).map(([letterCode, letterData]) => {
                const centerX = 50
                const centerY = 50
                const innerRadius = 35 // Raio interno (toca o grid do radar)
                const outerRadius = 52 // Raio externo (bem além do radar)
                const startAngle = -Math.PI / 2
                const anglePerItem = (2 * Math.PI) / methodologyChartData.chartData.length

                // Calcular ângulos de início e fim desta letra
                const minIndex = Math.min(...letterData.indices)
                const maxIndex = Math.max(...letterData.indices)

                // Ângulo do centro do primeiro pilar
                const startPillarAngle = startAngle + (minIndex * anglePerItem)
                // Ângulo do centro do último pilar
                const endPillarAngle = startAngle + (maxIndex * anglePerItem)

                // Expandir um pouco além do centro dos pilares (meio ângulo para cada lado)
                const segmentStartAngle = startPillarAngle - (anglePerItem / 2)
                const segmentEndAngle = endPillarAngle + (anglePerItem / 2)

                // Calcular pontos do arco
                const x1Inner = centerX + innerRadius * Math.cos(segmentStartAngle)
                const y1Inner = centerY + innerRadius * Math.sin(segmentStartAngle)
                const x1Outer = centerX + outerRadius * Math.cos(segmentStartAngle)
                const y1Outer = centerY + outerRadius * Math.sin(segmentStartAngle)

                const x2Outer = centerX + outerRadius * Math.cos(segmentEndAngle)
                const y2Outer = centerY + outerRadius * Math.sin(segmentEndAngle)
                const x2Inner = centerX + innerRadius * Math.cos(segmentEndAngle)
                const y2Inner = centerY + innerRadius * Math.sin(segmentEndAngle)

                // Determinar se o arco é maior que 180 graus
                const angleDiff = segmentEndAngle - segmentStartAngle
                const largeArcFlag = angleDiff > Math.PI ? 1 : 0

                // Criar path do segmento (anel parcial)
                const pathData = [
                  `M ${x1Inner} ${y1Inner}`, // Mover para ponto interno inicial
                  `L ${x1Outer} ${y1Outer}`, // Linha para ponto externo inicial
                  `A ${outerRadius} ${outerRadius} 0 ${largeArcFlag} 1 ${x2Outer} ${y2Outer}`, // Arco externo
                  `L ${x2Inner} ${y2Inner}`, // Linha para ponto interno final
                  `A ${innerRadius} ${innerRadius} 0 ${largeArcFlag} 0 ${x1Inner} ${y1Inner}`, // Arco interno (reverso)
                  'Z' // Fechar path
                ].join(' ')

                return (
                  <path
                    key={`area-${letterCode}`}
                    d={pathData}
                    fill={letterData.letterColor}
                    opacity={0.15}
                    stroke={letterData.letterColor}
                    strokeWidth={0.3}
                  />
                )
              })}

              {/* Círculos e letras por cima das áreas - no centro da área colorida */}
              {Array.from(methodologyChartData.letterPositions.entries()).map(([letterCode, letterData]) => {
                const angle = calculateLetterAngle(letterData.indices, methodologyChartData.chartData.length)

                // Centro do container (50%, 50%)
                const centerX = 50
                const centerY = 50
                const innerRadius = 35
                const outerRadius = 52
                const letterRadius = (innerRadius + outerRadius) / 2 // Centro da área colorida = 43.5

                // Calcular posição x, y baseado no ângulo
                const x = centerX + letterRadius * Math.cos(angle)
                const y = centerY + letterRadius * Math.sin(angle)

                return (
                  <g key={letterCode} style={{ zIndex: 1000 }}>
                    {/* Círculo de fundo - branco semi-opaco para garantir legibilidade */}
                    <circle
                      cx={x}
                      cy={y}
                      r="4.5"
                      fill="white"
                      opacity={0.9}
                      stroke={letterData.letterColor}
                      strokeWidth={0.5}
                    />
                    {/* Letra */}
                    <text
                      x={x}
                      y={y}
                      textAnchor="middle"
                      dominantBaseline="middle"
                      fill={letterData.letterColor}
                      fontSize="3.5"
                      fontWeight="bold"
                      style={{ paintOrder: 'stroke', stroke: 'white', strokeWidth: '0.3px' }}
                    >
                      {letterData.letterCode}
                    </text>
                  </g>
                )
              })}
            </svg>

            {/* Tooltip customizado - mesmo CSS do Recharts */}
            {hoveredPillar && (
              <div
                className="recharts-default-tooltip"
                style={{
                  position: 'fixed',
                  left: hoveredPillar.x,
                  top: hoveredPillar.y,
                  transform: 'translate(-50%, -100%)',
                  backgroundColor: 'white',
                  border: '1px solid #e5e7eb',
                  borderRadius: '8px',
                  padding: '8px 12px',
                  pointerEvents: 'none',
                  zIndex: 1000,
                  whiteSpace: 'nowrap'
                }}
              >
                <p className="recharts-tooltip-label" style={{ margin: 0 }}>
                  {hoveredPillar.name}
                </p>
                <ul className="recharts-tooltip-item-list" style={{ margin: 0, padding: 0 }}>
                  <li
                    className="recharts-tooltip-item"
                    style={{ color: getScoreColor(hoveredPillar.score), display: 'block', paddingTop: 4, paddingBottom: 4 }}
                  >
                    <span className="recharts-tooltip-item-name">Score</span>
                    <span className="recharts-tooltip-item-separator"> : </span>
                    <span className="recharts-tooltip-item-value">{hoveredPillar.score.toFixed(1)}%</span>
                  </li>
                </ul>
              </div>
            )}
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
