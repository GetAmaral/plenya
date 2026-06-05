'use client';

import { useState } from 'react';
import { useLocale } from 'next-intl';
import { agirLettersRaw } from '@/lib/agir-structure';

// ── Geometria ──────────────────────────────────────────────────────────
const RADAR_CX = 200;
const RADAR_CY = 200;
const RADAR_MAX = 150;
const ARC_RADIUS = RADAR_MAX + 18;
const LETTER_LABEL_RADIUS = ARC_RADIUS + 18;
const VIEWBOX = 400;

type LetterCode = 'A' | 'G' | 'I' | 'R'; // chave estável, locale-agnóstica
type DisplayCode = 'A' | 'G' | 'I' | 'R' | 'C' | 'T' | 'S';

const letterColors: Record<LetterCode, string> = {
  A: '#92b8b4', // sage
  G: '#b38645', // gold
  I: '#caa56b', // gold suave
  R: '#417e8e', // ocean
};

// Score exemplo por letra (mock visual). Ordem A-G-I-R independe do locale.
const letterScores: Record<LetterCode, number> = { A: 80, G: 76, I: 72, R: 84 };

// Territory por locale.
const letterTerritoryPt: Record<LetterCode, string> = {
  A: 'Os motores que nenhuma medicação substitui',
  G: 'O painel de controle interno',
  I: 'O eixo psicologia, sistema imune e inflamação',
  R: 'O substrato sobre o qual os outros pilares operam',
};
const letterTerritoryEn: Record<LetterCode, string> = {
  A: 'The engines no medication can replace',
  G: 'The internal control panel',
  I: 'The axis between psychology, immunity, and inflammation',
  R: 'The substrate every other pillar runs on',
};

// Display code (A/C/T/S em EN; A/G/I/R em PT).
function displayCode(letter: LetterCode, isEn: boolean): DisplayCode {
  if (!isEn) return letter;
  return letter === 'G' ? 'C' : letter === 'I' ? 'T' : letter === 'R' ? 'S' : 'A';
}

// Scores-semente (radar é exemplo visual). Índices além do array recebem um
// valor derivado determinístico — assim o radar NUNCA gera NaN quando a
// taxonomia em agir-structure cresce (ex.: pilar G expandido para 14 áreas).
const pillarScoreSeed: Record<LetterCode, number[]> = {
  A: [82, 78, 86, 74],
  G: [76, 70, 80, 72, 84, 78, 71, 82, 76, 73],
  I: [68, 72, 75, 70, 74],
  R: [84, 88, 80],
};

function pillarScore(letter: LetterCode, idx: number): number {
  const seed = pillarScoreSeed[letter];
  if (idx < seed.length) return seed[idx];
  // Onda suave determinística (±8) em torno do score-base da letra — sem
  // Math.random (evita hydration mismatch SSR/CSR) e sempre numérica.
  const wave = Math.round(Math.sin(idx * 1.7) * 8);
  return Math.max(60, Math.min(92, letterScores[letter] + wave));
}

// Setor de 90° por letra (mesmo do arco externo). Os pontos são distribuídos
// uniformemente DENTRO do setor conforme a contagem REAL de pilares — dinâmico,
// nunca desalinha com agir-structure.
const SECTOR_MID: Record<LetterCode, number> = { A: 0, G: 90, I: 180, R: 270 };
const SECTOR_HALF = 45;
const SECTOR_MARGIN = 11.25; // respiro nas bordas do setor

function pillarAngle(letter: LetterCode, idx: number, count: number): number {
  const mid = SECTOR_MID[letter];
  if (count <= 1) return mid;
  const start = mid - SECTOR_HALF + SECTOR_MARGIN;
  const span = 2 * (SECTOR_HALF - SECTOR_MARGIN);
  return start + (span * (idx + 0.5)) / count;
}

type Pillar = { letter: LetterCode; angle: number; score: number; name: string };

function buildRadarPillars(isEn: boolean): Pillar[] {
  const pillars: Pillar[] = [];
  for (const letter of ['A', 'G', 'I', 'R'] as LetterCode[]) {
    const raw = agirLettersRaw.find((l) => l.code === letter)!;
    const names = raw.groups.flatMap((g) => (isEn ? g.pillarsEn : g.pillars));
    names.forEach((name, idx) => {
      pillars.push({
        letter,
        angle: pillarAngle(letter, idx, names.length),
        score: pillarScore(letter, idx),
        name,
      });
    });
  }
  return pillars;
}

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

function tooltipPlacementForLetter(letter: LetterCode): {
  horizontal: 'left' | 'right' | 'center';
  vertical: 'top' | 'bottom' | 'center';
} {
  switch (letter) {
    case 'A': return { horizontal: 'center', vertical: 'bottom' };
    case 'G': return { horizontal: 'left',   vertical: 'center' };
    case 'I': return { horizontal: 'center', vertical: 'top'    };
    case 'R': return { horizontal: 'right',  vertical: 'center' };
  }
}

export function RadarAgir() {
  const locale = useLocale();
  const isEn = locale === 'en';
  const [hovered, setHovered] = useState<Hovered>({ type: 'none' });

  // Strings locale-aware do tooltip e caption.
  const i18n = isEn
    ? {
        letterLabel: (l: DisplayCode) => `Letter ${l}`,
        pillarLabel: (l: DisplayCode) => `Pillar · letter ${l}`,
        caption: 'Example · 22 pillars · scale 0–100 · hover for details',
        ariaLabel: 'Plenya Score example: 22 pillars organized into the four ACTS letters. Overall score 78.',
      }
    : {
        letterLabel: (l: DisplayCode) => `Letra ${l}`,
        pillarLabel: (l: DisplayCode) => `Pilar · letra ${l}`,
        caption: 'Exemplo · 22 pilares · escala 0–100 · passe o mouse para detalhes',
        ariaLabel: 'Exemplo de Escore Plenya: 22 pilares organizados em 4 letras AGIR. Score global 78.',
      };

  const letterMeta: Record<LetterCode, { full: string; score: number; territory: string }> = {
    A: {
      full: isEn ? agirLettersRaw[0].nameEn : agirLettersRaw[0].name,
      score: letterScores.A,
      territory: (isEn ? letterTerritoryEn : letterTerritoryPt).A,
    },
    G: {
      full: isEn ? agirLettersRaw[1].nameEn : agirLettersRaw[1].name,
      score: letterScores.G,
      territory: (isEn ? letterTerritoryEn : letterTerritoryPt).G,
    },
    I: {
      full: isEn ? agirLettersRaw[2].nameEn : agirLettersRaw[2].name,
      score: letterScores.I,
      territory: (isEn ? letterTerritoryEn : letterTerritoryPt).I,
    },
    R: {
      full: isEn ? agirLettersRaw[3].nameEn : agirLettersRaw[3].name,
      score: letterScores.R,
      territory: (isEn ? letterTerritoryEn : letterTerritoryPt).R,
    },
  };

  const radarPillars = buildRadarPillars(isEn);
  const radarPoints = radarPillars.map((p) => {
    const r = (p.score / 100) * RADAR_MAX;
    return { ...p, ...polar(p.angle, r) };
  });
  const polygonStr = radarPoints.map((p) => `${p.x},${p.y}`).join(' ');

  const isLetterActive = (l: LetterCode) =>
    (hovered.type === 'letter' && hovered.letter === l) ||
    (hovered.type === 'pillar' && radarPoints[hovered.index].letter === l);

  function tooltipPosition(h: Hovered) {
    if (h.type === 'letter') {
      const arc = letterArcs.find((a) => a.letter === h.letter)!;
      return polar(arc.midAngle, LETTER_LABEL_RADIUS);
    }
    if (h.type === 'pillar') {
      const p = radarPoints[h.index];
      return { x: p.x, y: p.y };
    }
    return null;
  }

  const ttPos = tooltipPosition(hovered);
  const hoveredLetter: LetterCode | null =
    hovered.type === 'letter' ? hovered.letter :
    hovered.type === 'pillar' ? radarPoints[hovered.index].letter :
    null;
  const placement = hoveredLetter ? tooltipPlacementForLetter(hoveredLetter) : null;

  const tooltipContent = (() => {
    if (hovered.type === 'letter') {
      const m = letterMeta[hovered.letter];
      const code = displayCode(hovered.letter, isEn);
      return (
        <>
          <div className="flex items-center gap-2 mb-1">
            <span className="w-2 h-2 rounded-full" style={{ background: letterColors[hovered.letter] }} />
            <span className="label-upper text-[9px]" style={{ color: letterColors[hovered.letter] }}>
              {i18n.letterLabel(code)}
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
      const code = displayCode(p.letter, isEn);
      return (
        <>
          <div className="flex items-center gap-2 mb-1">
            <span className="w-2 h-2 rounded-full" style={{ background: letterColors[p.letter] }} />
            <span className="label-upper text-[9px]" style={{ color: letterColors[p.letter] }}>
              {i18n.pillarLabel(code)}
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
        className="relative w-80 h-80 md:w-104 md:h-104"
        onMouseLeave={() => setHovered({ type: 'none' })}
      >
        <svg
          viewBox={`0 0 ${VIEWBOX} ${VIEWBOX}`}
          className="w-full h-full"
          role="img"
          aria-label={i18n.ariaLabel}
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

          {/* Labels A G I R (ou A C T S em EN) */}
          {letterArcs.map((arc) => {
            const p = polar(arc.midAngle, LETTER_LABEL_RADIUS);
            const active = isLetterActive(arc.letter);
            const code = displayCode(arc.letter, isEn);
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
                {code}
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

        {/* Tooltip flutuante */}
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
            <span>{displayCode(l, isEn)} {letterMeta[l].score}</span>
          </span>
        ))}
      </div>

      <p className="label-upper text-petrol/40 text-center text-[10px]">
        {i18n.caption}
      </p>
    </figure>
  );
}
