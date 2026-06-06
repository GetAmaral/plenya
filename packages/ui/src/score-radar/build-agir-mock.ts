import type { AgirRadarData, RadarLetter, RadarPillar } from './types'

export type MockLetterInput = { code: string; name: string; pillars: string[]; displayCode?: string }

// Scores fictícios DETERMINÍSTICOS (sem Math.random — evita hydration mismatch SSR/CSR).
// Usado só pelos radares de MARKETING do site (sem dado de paciente).
// Objetivo: parecer um escore REAL — regiões fortes e fracas com ondulação orgânica — e não
// uma flor simétrica (o seno por índice-dentro-da-letra fazia todos os pilares oscilarem em
// fase). Combina: base por letra + onda de baixa frequência sobre o índice GLOBAL (cria
// "regiões") + jitter pequeno por pilar (irregularidade fina). Tudo determinístico.
const LETTER_BASE: Record<string, number> = { A: 77, G: 72, I: 80, R: 83 }

// FNV-1a → [0,1). Estável entre SSR e CSR (só aritmética de inteiros sobre o texto).
function hash01(s: string): number {
  let h = 2166136261
  for (let i = 0; i < s.length; i++) {
    h ^= s.charCodeAt(i)
    h = Math.imul(h, 16777619)
  }
  return ((h >>> 0) % 100000) / 100000
}

function mockScore(code: string, name: string, globalIdx: number): number {
  const base = LETTER_BASE[code] ?? 75
  // Onda multi-frequência (não-harmônica) sobre o índice global → vales e cristas largos.
  const region =
    Math.sin(globalIdx * 0.55 + 0.7) * 9 + Math.sin(globalIdx * 0.21 + 2.1) * 6
  // Jitter fino por pilar (±7), irregular e determinístico.
  const jitter = (hash01(`${code}:${name}:${globalIdx}`) - 0.5) * 14
  return Math.max(46, Math.min(94, Math.round(base + region + jitter)))
}

/**
 * Adapter de MOCK para o radar do site (marketing): recebe a estrutura de letras+pilares
 * (de agir-structure) e devolve letras+pilares com scores fictícios, no formato que o
 * RadarAgir compartilhado consome. Mesma renderização do EMR, com dados de exemplo.
 */
export function buildAgirMock(letters: MockLetterInput[]): AgirRadarData {
  const outLetters: RadarLetter[] = []
  const outPillars: RadarPillar[] = []
  let globalIdx = 0
  for (const l of letters) {
    let sum = 0
    l.pillars.forEach((name) => {
      const s = mockScore(l.code, name, globalIdx)
      globalIdx += 1
      outPillars.push({ letter: l.code, name, score: s })
      sum += s
    })
    const avg = l.pillars.length ? sum / l.pillars.length : 0
    outLetters.push({ code: l.code, name: l.name, score: avg, color: '', label: l.displayCode })
  }
  return { letters: outLetters, pillars: outPillars }
}
