'use client';

import { useState } from 'react';

/**
 * Radar SVG puro (sem recharts) — adequado para o site público.
 * Recebe lista de grupos com scorePercentage 0-100 e renderiza um polígono.
 *
 * Não confundir com o RadarAgir.tsx (radar marketing, dados mockados).
 */

type RadarPoint = {
  label: string;
  value: number; // 0-100
};

export function ScoreRadarChart({
  data,
  size = 360,
}: {
  data: RadarPoint[];
  size?: number;
}) {
  const points = data.length;
  if (points < 3) {
    return (
      <p className="text-petrol/60 text-sm">
        Radar precisa de pelo menos 3 pilares para renderizar.
      </p>
    );
  }

  const cx = size / 2;
  const cy = size / 2;
  const radius = size / 2 - 60;
  const angleStep = (Math.PI * 2) / points;
  // Padding extra no viewBox para acomodar labels que estouram o quadrado.
  const vbPad = 70;

  // Quebra labels longos em 2 linhas (corte na quebra de palavra mais central).
  const wrapLabel = (label: string): string[] => {
    if (label.length <= 14) return [label];
    const mid = label.length / 2;
    let bestIdx = -1;
    let bestDist = Infinity;
    for (let i = 0; i < label.length; i++) {
      if (label[i] === ' ') {
        const dist = Math.abs(i - mid);
        if (dist < bestDist) {
          bestDist = dist;
          bestIdx = i;
        }
      }
    }
    return bestIdx === -1
      ? [label]
      : [label.slice(0, bestIdx), label.slice(bestIdx + 1)];
  };

  const axisPoints = data.map((d, i) => {
    const angle = -Math.PI / 2 + i * angleStep;
    const x = cx + Math.cos(angle) * radius;
    const y = cy + Math.sin(angle) * radius;
    const labelX = cx + Math.cos(angle) * (radius + 28);
    const labelY = cy + Math.sin(angle) * (radius + 28);
    return { x, y, labelX, labelY, label: d.label, angle };
  });

  const valuePoints = data.map((d, i) => {
    const angle = -Math.PI / 2 + i * angleStep;
    const r = (Math.max(0, Math.min(100, d.value)) / 100) * radius;
    return { x: cx + Math.cos(angle) * r, y: cy + Math.sin(angle) * r };
  });

  const polygon = valuePoints.map((p) => `${p.x},${p.y}`).join(' ');

  // Anéis de referência (25/50/75/100)
  const rings = [0.25, 0.5, 0.75, 1].map((scale) => {
    const ringPoints = axisPoints
      .map((ap) => {
        const angle = ap.angle;
        const x = cx + Math.cos(angle) * radius * scale;
        const y = cy + Math.sin(angle) * radius * scale;
        return `${x},${y}`;
      })
      .join(' ');
    return ringPoints;
  });

  const [hovered, setHovered] = useState<number | null>(null);
  const vbWidth = size + vbPad * 2;
  const vbHeight = size + vbPad * 2;

  // Posição do tooltip em % do container (coords SVG → % do viewBox).
  const tooltipPos =
    hovered !== null
      ? {
          left: ((valuePoints[hovered].x + vbPad) / vbWidth) * 100,
          top: ((valuePoints[hovered].y + vbPad) / vbHeight) * 100,
          // Anchor depende do quadrante: lado direito tooltip à esquerda etc.
          horizontal: (valuePoints[hovered].x < cx ? 'right' : 'left') as 'left' | 'right',
          vertical: (valuePoints[hovered].y < cy ? 'bottom' : 'top') as 'top' | 'bottom',
          point: data[hovered],
        }
      : null;

  return (
    <figure
      className="relative inline-block select-none"
      style={{ width: '100%', maxWidth: vbWidth }}
      onMouseLeave={() => setHovered(null)}
    >
      <svg
        viewBox={`${-vbPad} ${-vbPad} ${vbWidth} ${vbHeight}`}
        width="100%"
        height="auto"
        role="img"
        aria-label="Radar com pontuação por pilar"
      >
        {/* Anéis de fundo */}
        {rings.map((r, i) => (
          <polygon
            key={i}
            points={r}
            fill="none"
            stroke="rgba(28,55,55,0.15)"
            strokeWidth={1}
          />
        ))}
        {/* Eixos */}
        {axisPoints.map((ap, i) => (
          <line
            key={`axis-${i}`}
            x1={cx}
            y1={cy}
            x2={ap.x}
            y2={ap.y}
            stroke="rgba(28,55,55,0.15)"
            strokeWidth={1}
          />
        ))}
        {/* Polígono dos valores */}
        <polygon
          points={polygon}
          fill="rgba(196,154,90,0.25)"
          stroke="#C49A5A"
          strokeWidth={2}
          pointerEvents="none"
        />
        {/* Pontos + hit area */}
        {valuePoints.map((p, i) => {
          const isHovered = hovered === i;
          return (
            <g key={`pt-${i}`}>
              {isHovered && (
                <line
                  x1={cx}
                  y1={cy}
                  x2={p.x}
                  y2={p.y}
                  stroke="#C49A5A"
                  strokeOpacity="0.45"
                  strokeWidth={1}
                  strokeDasharray="2 2"
                  pointerEvents="none"
                />
              )}
              <circle
                cx={p.x}
                cy={p.y}
                r={14}
                fill="transparent"
                onMouseEnter={() => setHovered(i)}
                onFocus={() => setHovered(i)}
                tabIndex={0}
                style={{ cursor: 'pointer' }}
              />
              <circle
                cx={p.x}
                cy={p.y}
                r={isHovered ? 6 : 4}
                fill="#C49A5A"
                stroke="#fbfaf6"
                strokeWidth={1.2}
                pointerEvents="none"
                style={{ transition: 'r 150ms' }}
              />
            </g>
          );
        })}
        {/* Labels (com quebra em 2 linhas quando longos) */}
        {axisPoints.map((ap, i) => {
          const anchor =
            ap.labelX < cx - 4 ? 'end' : ap.labelX > cx + 4 ? 'start' : 'middle';
          const lines = wrapLabel(ap.label);
          const lineHeight = 13;
          const startDy = -((lines.length - 1) * lineHeight) / 2;
          return (
            <text
              key={`lbl-${i}`}
              x={ap.labelX}
              y={ap.labelY}
              fontSize={11}
              fill="#1C3737"
              textAnchor={anchor}
              dominantBaseline="middle"
              style={{ fontFamily: 'inherit', letterSpacing: '0.02em' }}
            >
              {lines.map((line, j) => (
                <tspan
                  key={j}
                  x={ap.labelX}
                  dy={j === 0 ? startDy : lineHeight}
                >
                  {line}
                </tspan>
              ))}
            </text>
          );
        })}
      </svg>

      {/* Tooltip flutuante */}
      <div
        className="absolute z-10 pointer-events-none transition-opacity duration-150"
        style={{
          left: tooltipPos ? `${tooltipPos.left}%` : 0,
          top: tooltipPos ? `${tooltipPos.top}%` : 0,
          transform: tooltipPos
            ? `translate(${tooltipPos.horizontal === 'left' ? '-100%' : '0'}, ${tooltipPos.vertical === 'top' ? '-100%' : '0'})`
            : undefined,
          marginLeft: tooltipPos ? (tooltipPos.horizontal === 'left' ? '-12px' : '12px') : 0,
          marginTop: tooltipPos ? (tooltipPos.vertical === 'top' ? '-12px' : '12px') : 0,
          opacity: tooltipPos ? 1 : 0,
          visibility: tooltipPos ? 'visible' : 'hidden',
        }}
      >
        {tooltipPos && (
          <div className="bg-paper border border-petrol/15 shadow-xl shadow-petrol/10 rounded-md px-3 py-2 w-[180px]">
            <p className="heading-section text-petrol text-sm leading-tight">
              {tooltipPos.point.label}
            </p>
            <p className="font-mono text-petrol/60 text-xs mt-1">
              {Math.round(tooltipPos.point.value)}{' '}
              <span className="text-petrol/35">/ 100</span>
            </p>
          </div>
        )}
      </div>
    </figure>
  );
}
