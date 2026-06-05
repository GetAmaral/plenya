import type { AgirRadarData, RadarLetter, RadarPillar } from './types'

export type MockLetterInput = { code: string; name: string; pillars: string[]; displayCode?: string }

// Scores fictícios DETERMINÍSTICOS (sem Math.random — evita hydration mismatch SSR/CSR).
// Usado só pelos radares de MARKETING do site (sem dado de paciente).
const LETTER_BASE: Record<string, number> = { A: 80, G: 76, I: 72, R: 84 }

function mockScore(code: string, idx: number): number {
  const base = LETTER_BASE[code] ?? 76
  const wave = Math.round(Math.sin((idx + 1) * 1.3) * 8)
  return Math.max(58, Math.min(94, base + wave))
}

/**
 * Adapter de MOCK para o radar do site (marketing): recebe a estrutura de letras+pilares
 * (de agir-structure) e devolve letras+pilares com scores fictícios, no formato que o
 * RadarAgir compartilhado consome. Mesma renderização do EMR, com dados de exemplo.
 */
export function buildAgirMock(letters: MockLetterInput[]): AgirRadarData {
  const outLetters: RadarLetter[] = []
  const outPillars: RadarPillar[] = []
  for (const l of letters) {
    let sum = 0
    l.pillars.forEach((name, i) => {
      const s = mockScore(l.code, i)
      outPillars.push({ letter: l.code, name, score: s })
      sum += s
    })
    const avg = l.pillars.length ? sum / l.pillars.length : 0
    outLetters.push({ code: l.code, name: l.name, score: avg, color: '', label: l.displayCode })
  }
  return { letters: outLetters, pillars: outPillars }
}
