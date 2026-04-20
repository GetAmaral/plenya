'use client'

import { useMemo, useState } from 'react'

// ── Tipos ──────────────────────────────────────────────────────────────
export type RadarLetter = {
  code: string
  name: string
  score: number
  color: string
}

export type RadarPillar = {
  letter: string
  name: string
  score: number
}

interface RadarAgirProps {
  letters: RadarLetter[]
  pillars: RadarPillar[]
  globalScore?: number
  ariaLabel?: string
}

// ── Geometria ──────────────────────────────────────────────────────────
const VIEWBOX = 400
const RADAR_CX = 200
const RADAR_CY = 200
const RADAR_MAX = 150
const ARC_RADIUS = RADAR_MAX + 18
const LETTER_LABEL_RADIUS = ARC_RADIUS + 18

const round1 = (n: number) => Math.round(n * 10) / 10

function polar(angleDeg: number, radius: number) {
  const rad = ((angleDeg - 90) * Math.PI) / 180
  return {
    x: round1(RADAR_CX + radius * Math.cos(rad)),
    y: round1(RADAR_CY + radius * Math.sin(rad)),
  }
}

function arcPath(startAngle: number, endAngle: number, radius: number) {
  const s = polar(startAngle, radius)
  const e = polar(endAngle, radius)
  return `M ${s.x} ${s.y} A ${radius} ${radius} 0 0 1 ${e.x} ${e.y}`
}

// ── State ──────────────────────────────────────────────────────────────
type Hovered =
  | { type: 'none' }
  | { type: 'letter'; code: string }
  | { type: 'pillar'; index: number }

function tooltipPlacementForCardinal(midAngle: number) {
  // -45..45 = topo → tooltip embaixo
  //  45..135 = direita → tooltip à esquerda
  // 135..225 = base → tooltip em cima
  // 225..315 = esquerda → tooltip à direita
  const a = ((midAngle % 360) + 360) % 360
  if (a < 45 || a >= 315) return { horizontal: 'center' as const, vertical: 'bottom' as const }
  if (a < 135) return { horizontal: 'left' as const, vertical: 'center' as const }
  if (a < 225) return { horizontal: 'center' as const, vertical: 'top' as const }
  return { horizontal: 'right' as const, vertical: 'center' as const }
}

export function RadarAgir({ letters, pillars, globalScore, ariaLabel }: RadarAgirProps) {
  const [hovered, setHovered] = useState<Hovered>({ type: 'none' })

  // Distribui pilares dentro do quadrante de cada letra.
  // Cada letra ocupa um setor de 360/N graus, centrado em -90 + (i * 360/N).
  // Para AGIR (N=4): A=0°, G=90°, I=180°, R=270° (topo, direita, base, esquerda).
  const { letterArcs, radarPoints, polygonStr, letterByCode } = useMemo(() => {
    const N = letters.length || 1
    const sectorSize = 360 / N

    const letterArcs = letters.map((l, i) => {
      const midAngle = i * sectorSize // 0, 90, 180, 270 para 4
      const start = midAngle - sectorSize / 2
      const end = midAngle + sectorSize / 2
      return { ...l, midAngle, start, end }
    })

    // Calcular ângulo de cada pilar dentro do setor da sua letra
    const pillarsByLetter = new Map<string, RadarPillar[]>()
    pillars.forEach((p) => {
      if (!pillarsByLetter.has(p.letter)) pillarsByLetter.set(p.letter, [])
      pillarsByLetter.get(p.letter)!.push(p)
    })

    const points: Array<RadarPillar & { x: number; y: number; angle: number; color: string }> = []
    letterArcs.forEach((arc) => {
      const ps = pillarsByLetter.get(arc.code) || []
      const n = ps.length
      // Margem dentro do setor para não colar nas bordas
      const innerStart = arc.start + sectorSize * 0.08
      const innerEnd = arc.end - sectorSize * 0.08
      const innerSpan = innerEnd - innerStart
      ps.forEach((p, i) => {
        // Distribuição uniforme: i+0.5 dentro do span
        const angle = n === 1
          ? arc.midAngle
          : innerStart + (innerSpan * (i + 0.5)) / n
        const r = (Math.max(0, Math.min(100, p.score)) / 100) * RADAR_MAX
        const { x, y } = polar(angle, r)
        points.push({ ...p, angle, x, y, color: arc.color })
      })
    })

    const polygonStr = points.map((p) => `${p.x},${p.y}`).join(' ')
    const letterByCode = new Map(letters.map((l) => [l.code, l]))

    return { letterArcs, radarPoints: points, polygonStr, letterByCode }
  }, [letters, pillars])

  const isLetterActive = (code: string) =>
    (hovered.type === 'letter' && hovered.code === code) ||
    (hovered.type === 'pillar' && radarPoints[hovered.index]?.letter === code)

  // Posição e placement do tooltip
  const ttData = (() => {
    if (hovered.type === 'letter') {
      const arc = letterArcs.find((a) => a.code === hovered.code)
      if (!arc) return null
      const p = polar(arc.midAngle, LETTER_LABEL_RADIUS)
      return { pos: p, placement: tooltipPlacementForCardinal(arc.midAngle) }
    }
    if (hovered.type === 'pillar') {
      const p = radarPoints[hovered.index]
      if (!p) return null
      const arc = letterArcs.find((a) => a.code === p.letter)
      const placement = arc
        ? tooltipPlacementForCardinal(arc.midAngle)
        : { horizontal: 'center' as const, vertical: 'top' as const }
      return { pos: { x: p.x, y: p.y }, placement }
    }
    return null
  })()

  const tooltipContent = (() => {
    if (hovered.type === 'letter') {
      const l = letterByCode.get(hovered.code)
      if (!l) return null
      return (
        <>
          <div className="flex items-center gap-2 mb-1">
            <span className="w-2 h-2 rounded-full" style={{ background: l.color }} />
            <span className="label-upper text-[9px]" style={{ color: l.color }}>
              Letra {l.code}
            </span>
          </div>
          <p className="heading-section text-foreground text-base leading-tight max-w-[24ch]">{l.name}</p>
          <p className="font-mono text-foreground/60 text-xs mt-1.5">
            {l.score.toFixed(1)} <span className="text-foreground/35">/ 100</span>
          </p>
        </>
      )
    }
    if (hovered.type === 'pillar') {
      const p = radarPoints[hovered.index]
      if (!p) return null
      return (
        <>
          <div className="flex items-center gap-2 mb-1">
            <span className="w-2 h-2 rounded-full" style={{ background: p.color }} />
            <span className="label-upper text-[9px]" style={{ color: p.color }}>
              Pilar · letra {p.letter}
            </span>
          </div>
          <p className="heading-section text-foreground text-base leading-tight max-w-[24ch]">{p.name}</p>
          <p className="font-mono text-foreground/60 text-xs mt-1.5">
            {p.score.toFixed(1)} <span className="text-foreground/35">/ 100</span>
          </p>
        </>
      )
    }
    return null
  })()

  const tooltipStyle: React.CSSProperties = (() => {
    if (!ttData) return {}
    const { pos, placement } = ttData
    const left = (pos.x / VIEWBOX) * 100
    const top = (pos.y / VIEWBOX) * 100
    const offsetPx = 14

    let translateX = '-50%'
    let translateY = '-50%'
    let marginLeft = '0'
    let marginTop = '0'

    if (placement.horizontal === 'left') {
      translateX = '-100%'
      marginLeft = `-${offsetPx}px`
    } else if (placement.horizontal === 'right') {
      translateX = '0'
      marginLeft = `${offsetPx}px`
    }
    if (placement.vertical === 'top') {
      translateY = '-100%'
      marginTop = `-${offsetPx}px`
    } else if (placement.vertical === 'bottom') {
      translateY = '0'
      marginTop = `${offsetPx}px`
    }

    return {
      left: `${left}%`,
      top: `${top}%`,
      transform: `translate(${translateX}, ${translateY})`,
      marginLeft,
      marginTop,
    }
  })()

  const computedGlobal =
    globalScore ??
    (letters.length > 0
      ? letters.reduce((sum, l) => sum + l.score, 0) / letters.length
      : 0)

  return (
    <figure className="flex flex-col items-center gap-5 select-none">
      <div
        className="relative w-full max-w-[26rem] aspect-square"
        onMouseLeave={() => setHovered({ type: 'none' })}
      >
        <svg
          viewBox={`0 0 ${VIEWBOX} ${VIEWBOX}`}
          className="w-full h-full"
          role="img"
          aria-label={ariaLabel || `Escore Plenya: ${pillars.length} pilares organizados em ${letters.length} letras. Score global ${computedGlobal.toFixed(0)}.`}
        >
          {/* Anéis concêntricos */}
          {[37.5, 75, 112.5, 150].map((r) => (
            <circle key={r} cx={RADAR_CX} cy={RADAR_CY} r={r} fill="none" stroke="hsl(var(--foreground))" strokeOpacity="0.07" strokeWidth="1" />
          ))}

          {/* Eixos cardinais sutis */}
          <line x1={RADAR_CX} y1={RADAR_CY - RADAR_MAX} x2={RADAR_CX} y2={RADAR_CY + RADAR_MAX} stroke="hsl(var(--foreground))" strokeOpacity="0.06" strokeWidth="1" />
          <line x1={RADAR_CX - RADAR_MAX} y1={RADAR_CY} x2={RADAR_CX + RADAR_MAX} y2={RADAR_CY} stroke="hsl(var(--foreground))" strokeOpacity="0.06" strokeWidth="1" />

          {/* Polígono ligando os vértices */}
          {radarPoints.length >= 3 && (
            <polygon
              points={polygonStr}
              fill="hsl(var(--primary))"
              fillOpacity="0.16"
              stroke="hsl(var(--primary))"
              strokeOpacity="0.85"
              strokeWidth="1.5"
              strokeLinejoin="round"
              pointerEvents="none"
            />
          )}

          {/* Anel externo colorido por letra — interativo */}
          {letterArcs.map((arc) => {
            const active = isLetterActive(arc.code)
            return (
              <path
                key={`arc-${arc.code}`}
                d={arcPath(arc.start + 2, arc.end - 2, ARC_RADIUS)}
                fill="none"
                stroke={arc.color}
                strokeWidth={active ? 10 : 7}
                strokeLinecap="round"
                opacity={hovered.type === 'none' ? 0.85 : active ? 1 : 0.3}
                onMouseEnter={() => setHovered({ type: 'letter', code: arc.code })}
                style={{ transition: 'stroke-width 180ms, opacity 180ms', cursor: 'pointer' }}
              />
            )
          })}

          {/* Pontos por pilar */}
          {radarPoints.map((p, i) => {
            const isPointHovered = hovered.type === 'pillar' && hovered.index === i
            const isInActiveLetter = hovered.type === 'letter' && hovered.code === p.letter
            const dimmed = hovered.type !== 'none' && !isPointHovered && !isInActiveLetter
            const radius = isPointHovered ? 6.5 : isInActiveLetter ? 5 : 3.5
            return (
              <g key={`pt-${i}`}>
                {isPointHovered && (
                  <line
                    x1={RADAR_CX}
                    y1={RADAR_CY}
                    x2={p.x}
                    y2={p.y}
                    stroke={p.color}
                    strokeOpacity="0.4"
                    strokeWidth="1"
                    strokeDasharray="2 2"
                    pointerEvents="none"
                  />
                )}
                <circle
                  cx={p.x}
                  cy={p.y}
                  r="14"
                  fill="transparent"
                  onMouseEnter={() => setHovered({ type: 'pillar', index: i })}
                  style={{ cursor: 'pointer' }}
                />
                <circle
                  cx={p.x}
                  cy={p.y}
                  r={radius}
                  fill={p.color}
                  stroke="hsl(var(--background))"
                  strokeWidth="1.2"
                  opacity={dimmed ? 0.35 : 1}
                  pointerEvents="none"
                  style={{ transition: 'r 180ms, opacity 180ms' }}
                />
              </g>
            )
          })}

          {/* Labels das letras */}
          {letterArcs.map((arc) => {
            const p = polar(arc.midAngle, LETTER_LABEL_RADIUS)
            const active = isLetterActive(arc.code)
            return (
              <text
                key={`label-${arc.code}`}
                x={p.x}
                y={p.y + 8}
                textAnchor="middle"
                fontFamily="var(--font-cormorant), 'Cormorant Garamond', serif"
                fontSize="26"
                fontWeight="500"
                fill={active ? arc.color : 'hsl(var(--foreground))'}
                opacity={hovered.type === 'none' ? 1 : active ? 1 : 0.35}
                onMouseEnter={() => setHovered({ type: 'letter', code: arc.code })}
                style={{ transition: 'fill 180ms, opacity 180ms', cursor: 'pointer' }}
              >
                {arc.code}
              </text>
            )
          })}

          {/* Score global no centro */}
          <circle cx={RADAR_CX} cy={RADAR_CY} r="34" fill="hsl(var(--background))" stroke="hsl(var(--foreground))" strokeOpacity="0.18" strokeWidth="1.5" pointerEvents="none" />
          <text
            x={RADAR_CX}
            y={RADAR_CY + 11}
            textAnchor="middle"
            fontFamily="var(--font-cormorant), 'Cormorant Garamond', serif"
            fontSize="34"
            fill="hsl(var(--foreground))"
            letterSpacing="-1"
            pointerEvents="none"
          >
            {computedGlobal.toFixed(0)}
          </text>
        </svg>

        {/* Tooltip flutuante */}
        <div
          className="absolute z-10 pointer-events-none transition-opacity duration-150"
          style={{
            ...tooltipStyle,
            opacity: tooltipContent ? 1 : 0,
            visibility: tooltipContent ? 'visible' : 'hidden',
          }}
        >
          <div className="bg-card border border-border shadow-xl shadow-foreground/10 rounded-md px-4 py-3 w-[200px]">
            {tooltipContent}
          </div>
        </div>
      </div>

      {/* Legenda das letras (sempre visível) */}
      <div className="flex flex-wrap justify-center gap-4 md:gap-6 font-mono text-[11px] uppercase tracking-[0.2em]">
        {letters.map((l) => (
          <span key={l.code} className="flex items-center gap-2 text-foreground/70">
            <span className="w-2 h-2 rounded-full" style={{ background: l.color }} />
            <span>{l.code} {l.score.toFixed(0)}</span>
          </span>
        ))}
      </div>
    </figure>
  )
}
