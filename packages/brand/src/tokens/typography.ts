export const fontFamilies = {
  heading: ['Nalieta', 'Cormorant Garamond', 'Georgia', 'serif'],
  mono: ['PolarbandWOO-Front', 'IBM Plex Mono', 'ui-monospace', 'monospace'],
  body: ['Inter', 'Geist', 'system-ui', '-apple-system', 'sans-serif'],
};

/**
 * Body type scale — modular, gentle 1.2 ratio for reading comfort.
 */
export const fontSizes: Record<string, [string, { lineHeight: string; letterSpacing?: string }]> = {
  '2xs': ['0.6875rem', { lineHeight: '1rem', letterSpacing: '0.08em' }],
  xs: ['0.75rem', { lineHeight: '1.125rem', letterSpacing: '0.04em' }],
  sm: ['0.875rem', { lineHeight: '1.375rem' }],
  base: ['1rem', { lineHeight: '1.65rem' }],
  lg: ['1.125rem', { lineHeight: '1.75rem' }],
  xl: ['1.25rem', { lineHeight: '1.85rem' }],
  '2xl': ['1.5rem', { lineHeight: '2rem', letterSpacing: '-0.005em' }],
  '3xl': ['1.875rem', { lineHeight: '2.25rem', letterSpacing: '-0.01em' }],
  '4xl': ['2.25rem', { lineHeight: '2.5rem', letterSpacing: '-0.015em' }],
  '5xl': ['3rem', { lineHeight: '1.05', letterSpacing: '-0.02em' }],
  '6xl': ['3.75rem', { lineHeight: '1.02', letterSpacing: '-0.025em' }],
  '7xl': ['4.5rem', { lineHeight: '1', letterSpacing: '-0.03em' }],
  '8xl': ['6rem', { lineHeight: '0.98', letterSpacing: '-0.035em' }],

  /**
   * Editorial display scale — 1.333 ratio (perfect fourth).
   * Used by .heading-display class for premium hero typography.
   * Negative letter-spacing tightens display text → signal of refinement.
   */
  'display-xs': ['2rem', { lineHeight: '1.1', letterSpacing: '-0.02em' }],
  'display-sm': ['2.667rem', { lineHeight: '1.05', letterSpacing: '-0.025em' }],
  'display-md': ['3.553rem', { lineHeight: '1.02', letterSpacing: '-0.03em' }],
  'display-lg': ['4.737rem', { lineHeight: '1', letterSpacing: '-0.035em' }],
  'display-xl': ['6.314rem', { lineHeight: '0.98', letterSpacing: '-0.04em' }],
  'display-2xl': ['8.418rem', { lineHeight: '0.96', letterSpacing: '-0.045em' }],
};

export const letterSpacing = {
  tighter: '-0.05em',
  tight: '-0.025em',
  snug: '-0.01em',
  normal: '0em',
  wide: '0.025em',
  wider: '0.05em',
  widest: '0.1em',
  brand: '0.3em',
  brandWide: '0.5em',
};
