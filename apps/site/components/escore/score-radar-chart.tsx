'use client';

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

  return (
    <svg
      viewBox={`0 0 ${size} ${size}`}
      width="100%"
      height="auto"
      style={{ maxWidth: size }}
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
      />
      {/* Pontos */}
      {valuePoints.map((p, i) => (
        <circle key={`pt-${i}`} cx={p.x} cy={p.y} r={4} fill="#C49A5A" />
      ))}
      {/* Labels */}
      {axisPoints.map((ap, i) => {
        const anchor =
          ap.labelX < cx - 4 ? 'end' : ap.labelX > cx + 4 ? 'start' : 'middle';
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
            {ap.label}
          </text>
        );
      })}
    </svg>
  );
}
