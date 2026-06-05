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

// ── Paleta Plenya (mesma do radar do site /escore-plenya) ───────────────
// A cor do método no banco é genérica (#10B981, #3B82F6…); aqui cravamos a
// paleta de marca por letra para o radar ter o design elegante do site.
// Fallback para a cor da prop (e cinza) caso surja uma letra fora de AGIR.
const PLENYA_PALETTE: Record<string, string> = {
  A: '#92b8b4', // sage
  G: '#b38645', // gold
  I: '#caa56b', // gold suave
  R: '#417e8e', // ocean
}
const POLYGON_COLOR = '#b38645' // gold — silhueta única (AGIR como um todo)
const CENTER_FILL = '#fbfaf6' // cream
const INK = '#063b4f' // petrol

function resolveColor(code: string, fallback?: string): string {
  return PLENYA_PALETTE[code] ?? fallback ?? '#94a3b8'
}

// ── Geometria ──────────────────────────────────────────────────────────
const VIEWBOX = 400
const RADAR_CX = 200
const RADAR_CY = 200
const RADAR_MAX = 150
const ARC_RADIUS = RADAR_MAX + 18
const LETTER_LABEL_RADIUS = ARC_RADIUS + 20
const ARC_GAP = 2.5 // respiro angular entre os arcos de letras

const round1 = (n: number) => Math.round(n * 10) / 10

// angleDeg: 0 = topo, cresce no sentido horário (90 = direita, 180 = base…).
function polar(angleDeg: number, radius: number) {
  const rad = ((angleDeg - 90) * Math.PI) / 180
  return {
    x: round1(RADAR_CX + radius * Math.cos(rad)),
    y: round1(RADAR_CY + radius * Math.sin(rad)),
  }
}

// Arco entre dois ângulos. large-arc-flag dinâmico — essencial para letras
// que ocupam mais de 180° (ex.: G com 23 pilares ≈ 197°).
function arcPath(startAngle: number, endAngle: number, radius: number) {
  const s = polar(startAngle, radius)
  const e = polar(endAngle, radius)
  const largeArc = (((endAngle - startAngle) % 360) + 360) % 360 > 180 ? 1 : 0
  return `M ${s.x} ${s.y} A ${radius} ${radius} 0 ${largeArc} 1 ${e.x} ${e.y}`
}

// ── State ──────────────────────────────────────────────────────────────
type Hovered =
  | { type: 'none' }
  | { type: 'letter'; code: string }
  | { type: 'pillar'; index: number }

function tooltipPlacementForAngle(angle: number) {
  const a = ((angle % 360) + 360) % 360
  if (a < 45 || a >= 315) return { horizontal: 'center' as const, vertical: 'bottom' as const }
  if (a < 135) return { horizontal: 'left' as const, vertical: 'center' as const }
  if (a < 225) return { horizontal: 'center' as const, vertical: 'top' as const }
  return { horizontal: 'right' as const, vertical: 'center' as const }
}

export function RadarAgir({ letters, pillars, globalScore, ariaLabel }: RadarAgirProps) {
  const [hovered, setHovered] = useState<Hovered>({ type: 'none' })

  // Largura angular de cada letra PROPORCIONAL ao seu nº de pilares.
  // Cada pilar ocupa um "slot" de 360/total graus; os pilares de uma letra
  // formam um bloco contíguo. Os vértices ficam igualmente espaçados no
  // círculo inteiro (G não fica mais espremido em 90°). Começa no topo.
  const { letterArcs, radarPoints, polygonStr, letterByCode } = useMemo(() => {
    // Pilares agrupados por letra, na ordem das letras (já ordenadas por order).
    const byLetter = new Map<string, RadarPillar[]>()
    for (const p of pillars) {
      if (!byLetter.has(p.letter)) byLetter.set(p.letter, [])
      byLetter.get(p.letter)!.push(p)
    }

    const total = pillars.length || 1
    const slot = 360 / total
    // Centraliza a PRIMEIRA letra (A, por order) no topo (12h): rotaciona
    // todos os ângulos para que o meio do arco de A caia em 0°.
    const firstCount = (byLetter.get(letters[0]?.code) ?? []).length
    const ROT = -(firstCount * slot) / 2

    const points: Array<RadarPillar & { x: number; y: number; angle: number; color: string }> = []
    const letterArcs: Array<RadarLetter & { start: number; end: number; midAngle: number; count: number; color: string }> = []

    let cursor = 0 // índice de slot acumulado
    for (const l of letters) {
      const ps = byLetter.get(l.code) ?? []
      const count = ps.length
      const start = cursor * slot + ROT
      const end = (cursor + count) * slot + ROT
      const color = resolveColor(l.code, l.color)

      ps.forEach((p, i) => {
        const angle = (cursor + i + 0.5) * slot + ROT // centro do slot
        const r = (Math.max(0, Math.min(100, p.score)) / 100) * RADAR_MAX
        const { x, y } = polar(angle, r)
        points.push({ ...p, angle, x, y, color })
      })

      if (count > 0) {
        letterArcs.push({ ...l, color, start, end, midAngle: (start + end) / 2, count })
      }
      cursor += count
    }

    const polygonStr = points.map((p) => `${p.x},${p.y}`).join(' ')
    const letterByCode = new Map(letterArcs.map((l) => [l.code, l]))

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
      return { pos: p, placement: tooltipPlacementForAngle(arc.midAngle) }
    }
    if (hovered.type === 'pillar') {
      const p = radarPoints[hovered.index]
      if (!p) return null
      return { pos: { x: p.x, y: p.y }, placement: tooltipPlacementForAngle(p.angle) }
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
              Letra {l.code} · {l.count} {l.count === 1 ? 'pilar' : 'pilares'}
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
            <circle key={r} cx={RADAR_CX} cy={RADAR_CY} r={r} fill="none" stroke={INK} strokeOpacity="0.07" strokeWidth="1" />
          ))}

          {/* Eixos cardinais sutis */}
          <line x1={RADAR_CX} y1={RADAR_CY - RADAR_MAX} x2={RADAR_CX} y2={RADAR_CY + RADAR_MAX} stroke={INK} strokeOpacity="0.06" strokeWidth="1" />
          <line x1={RADAR_CX - RADAR_MAX} y1={RADAR_CY} x2={RADAR_CX + RADAR_MAX} y2={RADAR_CY} stroke={INK} strokeOpacity="0.06" strokeWidth="1" />

          {/* Polígono ligando os vértices — silhueta dourada única */}
          {radarPoints.length >= 3 && (
            <polygon
              points={polygonStr}
              fill={POLYGON_COLOR}
              fillOpacity="0.16"
              stroke={POLYGON_COLOR}
              strokeOpacity="0.85"
              strokeWidth="1.5"
              strokeLinejoin="round"
              pointerEvents="none"
            />
          )}

          {/* Anel externo colorido por letra — largura proporcional, interativo */}
          {letterArcs.map((arc) => {
            const active = isLetterActive(arc.code)
            return (
              <path
                key={`arc-${arc.code}`}
                d={arcPath(arc.start + ARC_GAP, arc.end - ARC_GAP, ARC_RADIUS)}
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
                  r="12"
                  fill="transparent"
                  onMouseEnter={() => setHovered({ type: 'pillar', index: i })}
                  style={{ cursor: 'pointer' }}
                />
                <circle
                  cx={p.x}
                  cy={p.y}
                  r={radius}
                  fill={p.color}
                  stroke={CENTER_FILL}
                  strokeWidth="1.2"
                  opacity={dimmed ? 0.35 : 1}
                  pointerEvents="none"
                  style={{ transition: 'r 180ms, opacity 180ms' }}
                />
              </g>
            )
          })}

          {/* Labels das letras — no meio do arco proporcional */}
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
                fill={active ? arc.color : INK}
                opacity={hovered.type === 'none' ? 1 : active ? 1 : 0.35}
                onMouseEnter={() => setHovered({ type: 'letter', code: arc.code })}
                style={{ transition: 'fill 180ms, opacity 180ms', cursor: 'pointer' }}
              >
                {arc.code}
              </text>
            )
          })}

          {/* Score global no centro */}
          <circle cx={RADAR_CX} cy={RADAR_CY} r="34" fill={CENTER_FILL} stroke={INK} strokeOpacity="0.18" strokeWidth="1.5" pointerEvents="none" />
          <text
            x={RADAR_CX}
            y={RADAR_CY + 11}
            textAnchor="middle"
            fontFamily="var(--font-cormorant), 'Cormorant Garamond', serif"
            fontSize="34"
            fill={INK}
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

      {/* Legenda das letras (sempre visível) — score + nº de pilares */}
      <div className="flex flex-wrap justify-center gap-4 md:gap-6 font-mono text-[11px] uppercase tracking-[0.2em]">
        {letterArcs.map((l) => (
          <span key={l.code} className="flex items-center gap-2 text-foreground/70">
            <span className="w-2 h-2 rounded-full" style={{ background: l.color }} />
            <span>{l.code} {l.score.toFixed(0)}</span>
          </span>
        ))}
      </div>
    </figure>
  )
}
