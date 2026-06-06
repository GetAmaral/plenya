'use client'

import { useMemo, useState } from 'react'
import type { CSSProperties } from 'react'
import type { RadarLetter, RadarPillar } from './types'

export type RadarLabels = {
  /** Tag da letra no tooltip. count=0 → "sem dados". */
  letterTag?: (code: string, count: number) => string
  /** Tag do pilar no tooltip. */
  pillarTag?: (letterCode: string) => string
}

interface RadarAgirProps {
  letters: RadarLetter[]
  pillars: RadarPillar[]
  globalScore?: number
  ariaLabel?: string
  /** Classe Tailwind de largura máxima (responsivo: w-full + cap). */
  maxWidthClass?: string
  /** Override de largura via style (ex.: min(100%, 40rem, calc(100svh - 14.5rem))). Precede maxWidthClass. */
  widthStyle?: CSSProperties
  /** Mostra a legenda de letras abaixo do radar. */
  showLegend?: boolean
  /** Textos do tooltip (i18n). Default PT. */
  labels?: RadarLabels
}

// ── Paleta Plenya (= tokens de @plenya/brand: sage/gold/ocean/petrol/cream) ──
const PLENYA_PALETTE: Record<string, string> = {
  A: '#92b8b4', // sage
  G: '#b38645', // gold
  I: '#caa56b', // gold suave
  R: '#417e8e', // ocean
}
const POLYGON_COLOR = '#b38645' // gold
const CENTER_FILL = '#fbfaf6' // cream claro
const INK = '#063b4f' // petrol
const INK_70 = 'rgba(6,59,79,0.70)'
const INK_60 = 'rgba(6,59,79,0.60)'
const INK_40 = 'rgba(6,59,79,0.40)'
const INK_35 = 'rgba(6,59,79,0.35)'
const INK_12 = 'rgba(6,59,79,0.12)'
const CARD_BG = '#ffffff'

export function resolveColor(code: string, fallback?: string): string {
  return PLENYA_PALETTE[code] ?? fallback ?? '#94a3b8'
}

const DEFAULT_LABELS: Required<RadarLabels> = {
  letterTag: (code, count) =>
    count === 0 ? `Letra ${code} · sem dados` : `Letra ${code} · ${count} ${count === 1 ? 'pilar' : 'pilares'}`,
  pillarTag: (code) => `Pilar · letra ${code}`,
}

// ── Geometria ──────────────────────────────────────────────────────────
const VIEWBOX = 400
const RADAR_CX = 200
const RADAR_CY = 200
const RADAR_MAX = 150
const ARC_RADIUS = RADAR_MAX + 18
const LETTER_LABEL_RADIUS = ARC_RADIUS + 20
const ARC_GAP = 2.5
const MIN_LETTER_DEG = 10

const round1 = (n: number) => Math.round(n * 10) / 10

function polar(angleDeg: number, radius: number) {
  const rad = ((angleDeg - 90) * Math.PI) / 180
  return {
    x: round1(RADAR_CX + radius * Math.cos(rad)),
    y: round1(RADAR_CY + radius * Math.sin(rad)),
  }
}

// large-arc-flag dinâmico — essencial p/ letras que ocupam mais de 180° (G com muitos pilares).
function arcPath(startAngle: number, endAngle: number, radius: number) {
  const s = polar(startAngle, radius)
  const e = polar(endAngle, radius)
  const largeArc = (((endAngle - startAngle) % 360) + 360) % 360 > 180 ? 1 : 0
  return `M ${s.x} ${s.y} A ${radius} ${radius} 0 ${largeArc} 1 ${e.x} ${e.y}`
}

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

const TAG_STYLE: CSSProperties = { fontSize: 9, textTransform: 'uppercase', letterSpacing: '0.1em' }
const TITLE_STYLE: CSSProperties = { color: INK, fontWeight: 500, lineHeight: 1.15, maxWidth: '24ch' }

/**
 * Radar AGIR — fonte ÚNICA de geração (EMR + site). Largura de cada letra proporcional ao
 * nº de pilares visíveis (piso 10°/letra), A no topo, paleta de marca, SVG puro + tooltip.
 * Portável: cores hex (= @plenya/brand) + estilos inline; só usa utilitários Tailwind padrão.
 */
export function RadarAgir({
  letters,
  pillars,
  globalScore,
  ariaLabel,
  maxWidthClass = 'max-w-[26rem]',
  widthStyle,
  showLegend = true,
  labels,
}: RadarAgirProps) {
  const [hovered, setHovered] = useState<Hovered>({ type: 'none' })
  const L10N = { ...DEFAULT_LABELS, ...labels }

  const { letterArcs, radarPoints, polygonStr, letterByCode } = useMemo(() => {
    const byLetter = new Map<string, RadarPillar[]>()
    for (const p of pillars) {
      if (!byLetter.has(p.letter)) byLetter.set(p.letter, [])
      byLetter.get(p.letter)!.push(p)
    }

    const total = pillars.length
    const Lc = letters.length || 1
    const reserved = Lc * MIN_LETTER_DEG
    const remaining = Math.max(0, 360 - reserved)
    const perPillar = total > 0 ? remaining / total : 0
    const spanOf = (vp: number) => (total > 0 ? MIN_LETTER_DEG + vp * perPillar : 360 / Lc)
    const innerPad = total > 0 ? MIN_LETTER_DEG / 2 : 0

    const firstVp = (byLetter.get(letters[0]?.code) ?? []).length
    const ROT = -spanOf(firstVp) / 2

    const points: Array<RadarPillar & { x: number; y: number; angle: number; color: string }> = []
    const arcs: Array<RadarLetter & { start: number; end: number; midAngle: number; count: number; color: string }> = []

    let cursor = 0
    for (const l of letters) {
      const ps = byLetter.get(l.code) ?? []
      const vp = ps.length
      const span = spanOf(vp)
      const start = cursor + ROT
      const end = cursor + span + ROT
      const color = resolveColor(l.code, l.color)

      ps.forEach((p, i) => {
        const angle = start + innerPad + perPillar * (i + 0.5)
        const r = (Math.max(0, Math.min(100, p.score)) / 100) * RADAR_MAX
        const { x, y } = polar(angle, r)
        points.push({ ...p, angle, x, y, color })
      })

      arcs.push({ ...l, color, start, end, midAngle: (start + end) / 2, count: vp })
      cursor += span
    }

    const poly = points.map((p) => `${p.x},${p.y}`).join(' ')
    const byCode = new Map(arcs.map((a) => [a.code, a]))
    return { letterArcs: arcs, radarPoints: points, polygonStr: poly, letterByCode: byCode }
  }, [letters, pillars])

  const isLetterActive = (code: string) =>
    (hovered.type === 'letter' && hovered.code === code) ||
    (hovered.type === 'pillar' && radarPoints[hovered.index]?.letter === code)

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
      const empty = l.count === 0
      return (
        <>
          <div className="flex items-center gap-2 mb-1">
            <span className="w-2 h-2 rounded-full" style={{ background: l.color }} />
            <span style={{ ...TAG_STYLE, color: l.color }}>{L10N.letterTag(l.label ?? l.code, l.count)}</span>
          </div>
          <p className="text-base" style={TITLE_STYLE}>{l.name}</p>
          {!empty && (
            <p className="font-mono text-xs mt-1.5" style={{ color: INK_60 }}>
              {l.score.toFixed(1)} <span style={{ color: INK_35 }}>/ 100</span>
            </p>
          )}
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
            <span style={{ ...TAG_STYLE, color: p.color }}>{L10N.pillarTag(p.letter)}</span>
          </div>
          <p className="text-base" style={TITLE_STYLE}>{p.name}</p>
          <p className="font-mono text-xs mt-1.5" style={{ color: INK_60 }}>
            {p.score.toFixed(1)} <span style={{ color: INK_35 }}>/ 100</span>
          </p>
        </>
      )
    }
    return null
  })()

  const tooltipStyle: CSSProperties = (() => {
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

    return { left: `${left}%`, top: `${top}%`, transform: `translate(${translateX}, ${translateY})`, marginLeft, marginTop }
  })()

  const computedGlobal =
    globalScore ?? (letters.length > 0 ? letters.reduce((s, l) => s + l.score, 0) / letters.length : 0)

  return (
    <figure
      className={`flex flex-col items-center gap-5 select-none mx-auto${widthStyle ? '' : ` w-full ${maxWidthClass}`}`}
      style={widthStyle}
    >
      <div className="relative w-full aspect-square" onMouseLeave={() => setHovered({ type: 'none' })}>
        <svg
          viewBox={`0 0 ${VIEWBOX} ${VIEWBOX}`}
          className="w-full h-full"
          role="img"
          aria-label={ariaLabel || `Escore Plenya: ${pillars.length} pilares em ${letters.length} letras. Score global ${computedGlobal.toFixed(0)}.`}
        >
          {[37.5, 75, 112.5, 150].map((r) => (
            <circle key={r} cx={RADAR_CX} cy={RADAR_CY} r={r} fill="none" stroke={INK} strokeOpacity="0.07" strokeWidth="1" />
          ))}

          <line x1={RADAR_CX} y1={RADAR_CY - RADAR_MAX} x2={RADAR_CX} y2={RADAR_CY + RADAR_MAX} stroke={INK} strokeOpacity="0.06" strokeWidth="1" />
          <line x1={RADAR_CX - RADAR_MAX} y1={RADAR_CY} x2={RADAR_CX + RADAR_MAX} y2={RADAR_CY} stroke={INK} strokeOpacity="0.06" strokeWidth="1" />

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

          {letterArcs.map((arc) => {
            const active = isLetterActive(arc.code)
            const gap = Math.min(ARC_GAP, (arc.end - arc.start) / 3)
            return (
              <path
                key={`arc-${arc.code}`}
                d={arcPath(arc.start + gap, arc.end - gap, ARC_RADIUS)}
                fill="none"
                stroke={arc.color}
                strokeWidth={active ? 10 : 7}
                strokeLinecap="round"
                opacity={hovered.type === 'none' ? (arc.count === 0 ? 0.5 : 0.85) : active ? 1 : 0.3}
                onMouseEnter={() => setHovered({ type: 'letter', code: arc.code })}
                style={{ transition: 'stroke-width 180ms, opacity 180ms', cursor: 'pointer' }}
              />
            )
          })}

          {radarPoints.map((p, i) => {
            const isPointHovered = hovered.type === 'pillar' && hovered.index === i
            const isInActiveLetter = hovered.type === 'letter' && hovered.code === p.letter
            const dimmed = hovered.type !== 'none' && !isPointHovered && !isInActiveLetter
            const radius = isPointHovered ? 6.5 : isInActiveLetter ? 5 : 3.5
            return (
              <g key={`pt-${i}`}>
                {isPointHovered && (
                  <line x1={RADAR_CX} y1={RADAR_CY} x2={p.x} y2={p.y} stroke={p.color} strokeOpacity="0.4" strokeWidth="1" strokeDasharray="2 2" pointerEvents="none" />
                )}
                <circle cx={p.x} cy={p.y} r="12" fill="transparent" onMouseEnter={() => setHovered({ type: 'pillar', index: i })} style={{ cursor: 'pointer' }} />
                <circle cx={p.x} cy={p.y} r={radius} fill={p.color} stroke={CENTER_FILL} strokeWidth="1.2" opacity={dimmed ? 0.35 : 1} pointerEvents="none" style={{ transition: 'r 180ms, opacity 180ms' }} />
              </g>
            )
          })}

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
                opacity={hovered.type === 'none' ? (arc.count === 0 ? 0.55 : 1) : active ? 1 : 0.35}
                onMouseEnter={() => setHovered({ type: 'letter', code: arc.code })}
                style={{ transition: 'fill 180ms, opacity 180ms', cursor: 'pointer' }}
              >
                {arc.label ?? arc.code}
              </text>
            )
          })}

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

        <div
          className="absolute z-10 pointer-events-none transition-opacity duration-150 rounded-md px-4 py-3 w-[200px]"
          style={{
            ...tooltipStyle,
            background: CARD_BG,
            border: `1px solid ${INK_12}`,
            boxShadow: `0 10px 25px ${INK_12}`,
            opacity: tooltipContent ? 1 : 0,
            visibility: tooltipContent ? 'visible' : 'hidden',
          }}
        >
          {tooltipContent}
        </div>
      </div>

      {showLegend && (
        <div className="flex flex-wrap justify-center gap-2 md:gap-2.5">
          {letterArcs.map((l) => {
            const empty = l.count === 0
            return (
              <span
                key={l.code}
                className="inline-flex items-center gap-2 rounded-full py-1 pl-1 pr-3"
                style={{
                  background: empty ? 'transparent' : `${l.color}14`,
                  border: `1px solid ${empty ? INK_12 : `${l.color}55`}`,
                }}
              >
                <span
                  className="grid place-items-center rounded-full"
                  style={{
                    width: 22,
                    height: 22,
                    background: empty ? INK_12 : l.color,
                    color: empty ? INK_40 : CENTER_FILL,
                    fontFamily: "var(--font-cormorant), 'Cormorant Garamond', serif",
                    fontSize: 15,
                    fontWeight: 600,
                    lineHeight: 1,
                  }}
                >
                  {l.label ?? l.code}
                </span>
                <span
                  className="font-mono text-[12px] tabular-nums"
                  style={{ color: empty ? INK_40 : INK_70 }}
                >
                  {empty ? '—' : l.score.toFixed(0)}
                </span>
              </span>
            )
          })}
        </div>
      )}
    </figure>
  )
}
