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
  /** Classe Tailwind de largura máxima do radar (responsivo: w-full + cap). */
  maxWidthClass?: string
  /** Override de largura via style (ex.: min(100%, 40rem, calc(100svh - 18rem))
   *  pra o radar quadrado caber na altura útil da tela). Tem precedência sobre maxWidthClass. */
  widthStyle?: React.CSSProperties
  /** Mostra a legenda de letras abaixo do radar. Off quando a legenda vai pro lado. */
  showLegend?: boolean
}

// ── Paleta Plenya (mesma do radar do site /escore-plenya) ───────────────
// A cor do método no banco é genérica (#10B981, #3B82F6…); aqui cravamos a
// paleta de marca por letra para o radar ter o design elegante do site.
const PLENYA_PALETTE: Record<string, string> = {
  A: '#92b8b4', // sage
  G: '#b38645', // gold
  I: '#caa56b', // gold suave
  R: '#417e8e', // ocean
}
const POLYGON_COLOR = '#b38645' // gold — silhueta única (AGIR como um todo)
const CENTER_FILL = '#fbfaf6' // cream
const INK = '#063b4f' // petrol

export function resolveColor(code: string, fallback?: string): string {
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
// Cada LETRA ocupa no mínimo este arco, mesmo sem nenhum pilar visível.
const MIN_LETTER_DEG = 10

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
// que ocupam mais de 180° (ex.: G com muitos pilares).
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

export function RadarAgir({ letters, pillars, globalScore, ariaLabel, maxWidthClass = 'max-w-[26rem]', widthStyle, showLegend = true }: RadarAgirProps) {
  const [hovered, setHovered] = useState<Hovered>({ type: 'none' })

  // Largura angular de cada letra PROPORCIONAL ao seu nº de pilares VISÍVEIS,
  // com piso de MIN_LETTER_DEG por letra. Cada pilar visível ocupa a mesma
  // largura (perPillar) em todo o círculo. Letras vazias aparecem só com o
  // arco mínimo + rótulo (sem pontos). A primeira letra (A) centrada no topo.
  const { letterArcs, radarPoints, polygonStr, letterByCode } = useMemo(() => {
    const byLetter = new Map<string, RadarPillar[]>()
    for (const p of pillars) {
      if (!byLetter.has(p.letter)) byLetter.set(p.letter, [])
      byLetter.get(p.letter)!.push(p)
    }

    const total = pillars.length
    const L = letters.length || 1
    const reserved = L * MIN_LETTER_DEG
    const remaining = Math.max(0, 360 - reserved)
    const perPillar = total > 0 ? remaining / total : 0
    // span de uma letra com vp pilares visíveis (degenerado: divide igual).
    const spanOf = (vp: number) => (total > 0 ? MIN_LETTER_DEG + vp * perPillar : 360 / L)
    // padding interno que centraliza os pilares dentro do arco da letra.
    const innerPad = total > 0 ? MIN_LETTER_DEG / 2 : 0

    // Centraliza a 1ª letra (A, por order) no topo (12h).
    const firstVp = (byLetter.get(letters[0]?.code) ?? []).length
    const ROT = -spanOf(firstVp) / 2

    const points: Array<RadarPillar & { x: number; y: number; angle: number; color: string }> = []
    const letterArcs: Array<RadarLetter & { start: number; end: number; midAngle: number; count: number; color: string }> = []

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

      letterArcs.push({ ...l, color, start, end, midAngle: (start + end) / 2, count: vp })
      cursor += span
    }

    const polygonStr = points.map((p) => `${p.x},${p.y}`).join(' ')
    const letterByCode = new Map(letterArcs.map((a) => [a.code, a]))

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
      const empty = l.count === 0
      return (
        <>
          <div className="flex items-center gap-2 mb-1">
            <span className="w-2 h-2 rounded-full" style={{ background: l.color }} />
            <span className="label-upper text-[9px]" style={{ color: l.color }}>
              Letra {l.code} · {empty ? 'sem dados' : `${l.count} ${l.count === 1 ? 'pilar' : 'pilares'}`}
            </span>
          </div>
          <p className="heading-section text-foreground text-base leading-tight max-w-[24ch]">{l.name}</p>
          {!empty && (
            <p className="font-mono text-foreground/60 text-xs mt-1.5">
              {l.score.toFixed(1)} <span className="text-foreground/35">/ 100</span>
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
    <figure
      className={`flex flex-col items-center gap-5 select-none mx-auto${widthStyle ? '' : ` w-full ${maxWidthClass}`}`}
      style={widthStyle}
    >
      <div
        className="relative w-full aspect-square"
        onMouseLeave={() => setHovered({ type: 'none' })}
      >
        <svg
          viewBox={`0 0 ${VIEWBOX} ${VIEWBOX}`}
          className="w-full h-full"
          role="img"
          aria-label={ariaLabel || `Escore Plenya: ${pillars.length} pilares em ${letters.length} letras. Score global ${computedGlobal.toFixed(0)}.`}
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

          {/* Anel externo colorido por letra — largura proporcional, interativo.
              Letras vazias mostram só o arco mínimo + rótulo. */}
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

          {/* Labels das letras — no meio do arco (proporcional) */}
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

      {/* Legenda das letras — letra vazia mostra "—". Off quando vai pro painel lateral. */}
      {showLegend && (
        <div className="flex flex-wrap justify-center gap-4 md:gap-6 font-mono text-[11px] uppercase tracking-[0.2em]">
          {letterArcs.map((l) => (
            <span key={l.code} className="flex items-center gap-2 text-foreground/70">
              <span className="w-2 h-2 rounded-full" style={{ background: l.color, opacity: l.count === 0 ? 0.5 : 1 }} />
              <span className={l.count === 0 ? 'text-foreground/40' : undefined}>
                {l.code} {l.count === 0 ? '—' : l.score.toFixed(0)}
              </span>
            </span>
          ))}
        </div>
      )}
    </figure>
  )
}
