/**
 * Cores das letras AGIR — fonte ÚNICA (radar do EMR, radar do site, barras por pilar).
 *
 * A paleta anterior (sage/gold/gold-suave/ocean) reprovava no validador de paleta: G #b38645 e
 * I #caa56b eram praticamente o mesmo dourado, ΔE 9,3 em visão normal contra um piso de 15 —
 * indistinguíveis mesmo com visão perfeita.
 *
 * Estas passam em todas as checagens em light e dark: banda de luminosidade, piso de croma,
 * visão normal (pior par ΔE 19,3) e contraste >= 3:1. A separação para daltonismo cai na banda
 * 6-8, legal porque a letra sempre aparece com o glifo A/G/I/R e o nome por extenso — a cor
 * nunca é o único canal.
 *
 * Os tons originais da marca não passavam no piso de croma: cor dessaturada lê como cinza em
 * barra fina.
 */
export const LETTER_COLORS: Record<string, { light: string; dark: string }> = {
  A: { light: '#0f8f5f', dark: '#2b9c6d' }, // verde
  G: { light: '#c07520', dark: '#c9822f' }, // âmbar
  I: { light: '#b04a8a', dark: '#b85e93' }, // magenta
  R: { light: '#2a6fb0', dark: '#4585bc' }, // azul
}

export const LETTER_FALLBACK = { light: '#6b7280', dark: '#9ca3af' }

export const LETTER_ORDER = ['A', 'G', 'I', 'R']

export function letterColor(code: string) {
  return LETTER_COLORS[code] ?? LETTER_FALLBACK
}
