'use client';

import { useState } from 'react';

// ── Geometria ──────────────────────────────────────────────────────────
const RADAR_CX = 200;
const RADAR_CY = 200;
const RADAR_MAX = 150;
const ARC_RADIUS = RADAR_MAX + 18;
const LETTER_LABEL_RADIUS = ARC_RADIUS + 18;
const VIEWBOX = 400;

type LetterCode = 'A' | 'G' | 'I' | 'R';

const letterColors: Record<LetterCode, string> = {
  A: '#92b8b4', // sage
  G: '#b38645', // gold
  I: '#caa56b', // gold suave
  R: '#417e8e', // ocean
};

const letterMeta: Record<LetterCode, { full: string; score: number; territory: string }> = {
  A: { full: 'Alimentação e Atividade Física', score: 80, territory: 'Os motores que nenhuma medicação substitui' },
  G: { full: 'Gestão Metabólica', score: 76, territory: 'O painel de controle interno' },
  I: { full: 'Integração Mente-Corpo', score: 72, territory: 'O eixo psicologia, sistema imune e inflamação' },
  R: { full: 'Ritmo Circadiano', score: 84, territory: 'O substrato sobre o qual os outros pilares operam' },
};

type Pillar = { letter: LetterCode; angle: number; score: number; name: string };

const radarPillars: Pillar[] = [
  // A — top quadrant (-45° → +45°), 4 pilares
  { letter: 'A', angle: -33.75, score: 82, name: 'Avaliação Nutricional' },
  { letter: 'A', angle: -11.25, score: 78, name: 'Prescrição de Exercícios' },
  { letter: 'A', angle:  11.25, score: 86, name: 'Composição Corporal' },
  { letter: 'A', angle:  33.75, score: 74, name: 'Suplementação' },
  // G — right quadrant (45° → 135°), 10 pilares
  { letter: 'G', angle:  49.5, score: 76, name: 'Controle Glicêmico' },
  { letter: 'G', angle:  58.5, score: 70, name: 'Perfil Lipídico' },
  { letter: 'G', angle:  67.5, score: 80, name: 'Função Hepática' },
  { letter: 'G', angle:  76.5, score: 72, name: 'Função Renal' },
  { letter: 'G', angle:  85.5, score: 84, name: 'Risco Cardiovascular' },
  { letter: 'G', angle:  94.5, score: 78, name: 'Painel Hormonal' },
  { letter: 'G', angle: 103.5, score: 71, name: 'Inflamação e Imunidade' },
  { letter: 'G', angle: 112.5, score: 82, name: 'Vitaminas, Minerais e Micronutrientes' },
  { letter: 'G', angle: 121.5, score: 76, name: 'Hematologia' },
  { letter: 'G', angle: 130.5, score: 73, name: 'Rastreamento Oncológico' },
  // I — bottom quadrant (135° → 225°), 5 pilares
  { letter: 'I', angle: 144, score: 68, name: 'Avaliação Psicológica' },
  { letter: 'I', angle: 162, score: 72, name: 'Técnicas de Relaxamento' },
  { letter: 'I', angle: 180, score: 75, name: 'Função Cognitiva' },
  { letter: 'I', angle: 198, score: 70, name: 'Vida Sexual' },
  { letter: 'I', angle: 216, score: 74, name: 'Vínculos Sociais e Suporte' },
  // R — left quadrant (225° → 315°), 3 pilares
  { letter: 'R', angle: 240, score: 84, name: 'Qualidade do Sono' },
  { letter: 'R', angle: 270, score: 88, name: 'Cronobiologia' },
  { letter: 'R', angle: 300, score: 80, name: 'Exposição à Luz' },
];

// round1: arredonda para 1 decimal e devolve number "estável" (evita
// hydration mismatch entre SSR/CSR causado por floats com muitos dígitos).
const round1 = (n: number) => Math.round(n * 10) / 10;

function polar(angleDeg: number, radius: number) {
  const rad = ((angleDeg - 90) * Math.PI) / 180;
  return {
    x: round1(RADAR_CX + radius * Math.cos(rad)),
    y: round1(RADAR_CY + radius * Math.sin(rad)),
  };
}

const radarPoints = radarPillars.map((p) => {
  const r = (p.score / 100) * RADAR_MAX;
  return { ...p, ...polar(p.angle, r) };
});

const polygonStr = radarPoints.map((p) => `${p.x},${p.y}`).join(' ');

function arcPath(startAngle: number, endAngle: number, radius: number) {
  const s = polar(startAngle, radius);
  const e = polar(endAngle, radius);
  return `M ${s.x} ${s.y} A ${radius} ${radius} 0 0 1 ${e.x} ${e.y}`;
}

const letterArcs: { letter: LetterCode; start: number; end: number; midAngle: number }[] = [
  { letter: 'A', start: -45, end:  45, midAngle:   0 },
  { letter: 'G', start:  45, end: 135, midAngle:  90 },
  { letter: 'I', start: 135, end: 225, midAngle: 180 },
  { letter: 'R', start: 225, end: 315, midAngle: 270 },
];

// ── State ──────────────────────────────────────────────────────────────
type Hovered =
  | { type: 'none' }
  | { type: 'letter'; letter: LetterCode }
  | { type: 'pillar'; index: number };

// Posição do tooltip em coordenadas SVG (0..400) → percentual no container
function tooltipPosition(hovered: Hovered) {
  if (hovered.type === 'letter') {
    const arc = letterArcs.find((a) => a.letter === hovered.letter)!;
    const p = polar(arc.midAngle, LETTER_LABEL_RADIUS);
    return { x: p.x, y: p.y };
  }
  if (hovered.type === 'pillar') {
    const p = radarPoints[hovered.index];
    return { x: p.x, y: p.y };
  }
  return null;
}

// Placement do tooltip determinado pela LETRA (não pela posição em pixels) —
// estável para todos os pilares dentro do mesmo setor, evita flips quando
// o mouse pula entre pontos adjacentes próximos do limite de quadrante.
function tooltipPlacementForLetter(letter: LetterCode): {
  horizontal: 'left' | 'right' | 'center';
  vertical: 'top' | 'bottom' | 'center';
} {
  switch (letter) {
    case 'A': return { horizontal: 'center', vertical: 'bottom' }; // topo → tooltip abaixo
    case 'G': return { horizontal: 'left',   vertical: 'center' }; // direita → tooltip à esquerda
    case 'I': return { horizontal: 'center', vertical: 'top'    }; // base → tooltip acima
    case 'R': return { horizontal: 'right',  vertical: 'center' }; // esquerda → tooltip à direita
  }
}

export function RadarAgir() {
  const [hovered, setHovered] = useState<Hovered>({ type: 'none' });

  const isLetterActive = (l: LetterCode) =>
    (hovered.type === 'letter' && hovered.letter === l) ||
    (hovered.type === 'pillar' && radarPoints[hovered.index].letter === l);

  const ttPos = tooltipPosition(hovered);
  const hoveredLetter: LetterCode | null =
    hovered.type === 'letter' ? hovered.letter :
    hovered.type === 'pillar' ? radarPoints[hovered.index].letter :
    null;
  const placement = hoveredLetter ? tooltipPlacementForLetter(hoveredLetter) : null;

  // Conteúdo do tooltip
  const tooltipContent = (() => {
    if (hovered.type === 'letter') {
      const m = letterMeta[hovered.letter];
      return (
        <>
          <div className="flex items-center gap-2 mb-1">
            <span className="w-2 h-2 rounded-full" style={{ background: letterColors[hovered.letter] }} />
            <span className="label-upper text-[9px]" style={{ color: letterColors[hovered.letter] }}>
              Letra {hovered.letter}
            </span>
          </div>
          <p className="heading-section text-petrol text-base leading-tight max-w-[24ch]">{m.full}</p>
          <p className="font-mono text-petrol/60 text-xs mt-1.5">
            {m.score} <span className="text-petrol/35">/ 100</span>
          </p>
          <p className="text-petrol/50 text-[11px] italic mt-1.5 max-w-[28ch] leading-snug">{m.territory}</p>
        </>
      );
    }
    if (hovered.type === 'pillar') {
      const p = radarPoints[hovered.index];
      return (
        <>
          <div className="flex items-center gap-2 mb-1">
            <span className="w-2 h-2 rounded-full" style={{ background: letterColors[p.letter] }} />
            <span className="label-upper text-[9px]" style={{ color: letterColors[p.letter] }}>
              Pilar · letra {p.letter}
            </span>
          </div>
          <p className="heading-section text-petrol text-base leading-tight max-w-[24ch]">{p.name}</p>
          <p className="font-mono text-petrol/60 text-xs mt-1.5">
            {p.score} <span className="text-petrol/35">/ 100</span>
          </p>
        </>
      );
    }
    return null;
  })();

  // Estilo de posicionamento do tooltip
  const tooltipStyle: React.CSSProperties = (() => {
    if (!ttPos || !placement) return {};
    const left = (ttPos.x / VIEWBOX) * 100;
    const top = (ttPos.y / VIEWBOX) * 100;
    const offsetPx = 14;

    let translateX = '-50%';
    let translateY = '-50%';
    let marginLeft = '0';
    let marginTop = '0';

    if (placement.horizontal === 'left') {
      translateX = '-100%';
      marginLeft = `-${offsetPx}px`;
    } else if (placement.horizontal === 'right') {
      translateX = '0';
      marginLeft = `${offsetPx}px`;
    }

    if (placement.vertical === 'top') {
      translateY = '-100%';
      marginTop = `-${offsetPx}px`;
    } else if (placement.vertical === 'bottom') {
      translateY = '0';
      marginTop = `${offsetPx}px`;
    }

    return {
      left: `${left}%`,
      top: `${top}%`,
      transform: `translate(${translateX}, ${translateY})`,
      marginLeft,
      marginTop,
    };
  })();

  return (
    <figure className="flex flex-col items-center gap-5 select-none">
      <div
        className="relative w-80 h-80 md:w-[26rem] md:h-[26rem]"
        onMouseLeave={() => setHovered({ type: 'none' })}
      >
        <svg
          viewBox={`0 0 ${VIEWBOX} ${VIEWBOX}`}
          className="w-full h-full"
          role="img"
          aria-label="Exemplo de Escore Plenya: 22 pilares organizados em 4 letras AGIR. Score global 78."
        >
          {/* Anéis concêntricos de fundo */}
          {[37.5, 75, 112.5, 150].map((r) => (
            <circle key={r} cx={RADAR_CX} cy={RADAR_CY} r={r} fill="none" stroke="#063b4f" strokeOpacity="0.07" strokeWidth="1" />
          ))}

          {/* Eixos cardinais sutis */}
          <line x1={RADAR_CX} y1={RADAR_CY - RADAR_MAX} x2={RADAR_CX} y2={RADAR_CY + RADAR_MAX} stroke="#063b4f" strokeOpacity="0.06" strokeWidth="1" />
          <line x1={RADAR_CX - RADAR_MAX} y1={RADAR_CY} x2={RADAR_CX + RADAR_MAX} y2={RADAR_CY} stroke="#063b4f" strokeOpacity="0.06" strokeWidth="1" />

          {/* Polígono ligando os 22 vértices */}
          <polygon
            points={polygonStr}
            fill="#b38645"
            fillOpacity="0.16"
            stroke="#b38645"
            strokeOpacity="0.85"
            strokeWidth="1.5"
            strokeLinejoin="round"
            pointerEvents="none"
          />

          {/* Anel externo colorido por letra — interativo */}
          {letterArcs.map((arc) => {
            const active = isLetterActive(arc.letter);
            return (
              <path
                key={`arc-${arc.letter}`}
                d={arcPath(arc.start + 2, arc.end - 2, ARC_RADIUS)}
                fill="none"
                stroke={letterColors[arc.letter]}
                strokeWidth={active ? 10 : 7}
                strokeLinecap="round"
                opacity={hovered.type === 'none' ? 0.85 : active ? 1 : 0.3}
                onMouseEnter={() => setHovered({ type: 'letter', letter: arc.letter })}
                style={{ transition: 'stroke-width 180ms, opacity 180ms', cursor: 'pointer' }}
              />
            );
          })}

          {/* Pontos por pilar */}
          {radarPoints.map((p, i) => {
            const isPointHovered = hovered.type === 'pillar' && hovered.index === i;
            const isInActiveLetter = hovered.type === 'letter' && hovered.letter === p.letter;
            const dimmed = hovered.type !== 'none' && !isPointHovered && !isInActiveLetter;
            const radius = isPointHovered ? 6.5 : isInActiveLetter ? 5 : 3.5;
            return (
              <g key={`pt-${i}`}>
                {isPointHovered && (
                  <line
                    x1={RADAR_CX}
                    y1={RADAR_CY}
                    x2={p.x}
                    y2={p.y}
                    stroke={letterColors[p.letter]}
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
                  fill={letterColors[p.letter]}
                  stroke="#fbfaf6"
                  strokeWidth="1.2"
                  opacity={dimmed ? 0.35 : 1}
                  pointerEvents="none"
                  style={{ transition: 'r 180ms, opacity 180ms' }}
                />
              </g>
            );
          })}

          {/* Labels A G I R */}
          {letterArcs.map((arc) => {
            const p = polar(arc.midAngle, LETTER_LABEL_RADIUS);
            const active = isLetterActive(arc.letter);
            return (
              <text
                key={`label-${arc.letter}`}
                x={p.x}
                y={p.y + 8}
                textAnchor="middle"
                fontFamily="'Cormorant Garamond', serif"
                fontSize="26"
                fontWeight="500"
                fill={active ? letterColors[arc.letter] : '#063b4f'}
                opacity={hovered.type === 'none' ? 1 : active ? 1 : 0.35}
                onMouseEnter={() => setHovered({ type: 'letter', letter: arc.letter })}
                style={{ transition: 'fill 180ms, opacity 180ms', cursor: 'pointer' }}
              >
                {arc.letter}
              </text>
            );
          })}

          {/* Score global no centro */}
          <circle cx={RADAR_CX} cy={RADAR_CY} r="34" fill="#fbfaf6" stroke="#063b4f" strokeOpacity="0.18" strokeWidth="1.5" pointerEvents="none" />
          <text
            x={RADAR_CX}
            y={RADAR_CY + 11}
            textAnchor="middle"
            fontFamily="'Cormorant Garamond', serif"
            fontSize="34"
            fill="#063b4f"
            letterSpacing="-1"
            pointerEvents="none"
          >
            78
          </text>
        </svg>

        {/* Tooltip flutuante — sempre montado, opacity controla visibilidade
            (evita re-disparar animação a cada mudança de hover) */}
        <div
          className="absolute z-10 pointer-events-none transition-opacity duration-150"
          style={{
            ...tooltipStyle,
            opacity: tooltipContent ? 1 : 0,
            visibility: tooltipContent ? 'visible' : 'hidden',
          }}
        >
          <div className="bg-paper border border-petrol/15 shadow-xl shadow-petrol/10 rounded-md px-4 py-3 w-[200px]">
            {tooltipContent}
          </div>
        </div>
      </div>

      {/* Legenda das cores por letra (sempre visível) */}
      <div className="flex gap-4 md:gap-6 font-mono text-[11px] uppercase tracking-[0.2em]">
        {(['A', 'G', 'I', 'R'] as LetterCode[]).map((l) => (
          <span key={l} className="flex items-center gap-2 text-petrol/70">
            <span className="w-2 h-2 rounded-full" style={{ background: letterColors[l] }} />
            <span>{l} {letterMeta[l].score}</span>
          </span>
        ))}
      </div>

      <p className="label-upper text-petrol/40 text-center text-[10px]">
        Exemplo · 22 pilares · escala 0–100 · passe o mouse para detalhes
      </p>
    </figure>
  );
}
