'use client';

import { useLocale } from 'next-intl';
import {
  RadarAgir as SharedRadarAgir,
  buildAgirMock,
  type RadarLabels,
} from '@plenya/ui/score-radar';
import { agirLettersRaw } from '@/lib/agir-structure';

// Display code AGIR (PT) → ACTS (EN): A→A, G→C, I→T, R→S. O `code` interno segue A/G/I/R
// (paleta + casamento letra↔pilar); só o glifo exibido muda.
const EN_DISPLAY: Record<string, string> = { A: 'A', G: 'C', I: 'T', R: 'S' };

const EN_LABELS: RadarLabels = {
  letterTag: (code, count) =>
    count === 0
      ? `Letter ${code} · no data`
      : `Letter ${code} · ${count} ${count === 1 ? 'pillar' : 'pillars'}`,
  pillarTag: (code) => `Pillar · letter ${code}`,
};

/**
 * Radar de marketing do site (/escore-plenya, home). É o MESMO componente do EMR
 * (`@plenya/ui` — fonte única, layout proporcional), com dados MOCK gerados de
 * `agir-structure` (sem dado de paciente). Localizado PT/EN.
 */
export function RadarAgir() {
  const isEn = useLocale() === 'en';
  const mock = buildAgirMock(
    agirLettersRaw.map((l) => ({
      code: l.code,
      name: isEn ? l.nameEn : l.name,
      displayCode: isEn ? EN_DISPLAY[l.code] : l.code,
      pillars: l.groups.flatMap((g) => (isEn ? g.pillarsEn : g.pillars)),
    })),
  );

  return (
    <SharedRadarAgir
      letters={mock.letters}
      pillars={mock.pillars}
      globalScore={78}
      maxWidthClass="max-w-[26rem] md:max-w-[28rem]"
      labels={isEn ? EN_LABELS : undefined}
    />
  );
}
