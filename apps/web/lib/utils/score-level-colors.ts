/**
 * Centralized Score Level Color System
 *
 * Used across the application for consistent color representation of score levels.
 * Each level (0-6) has a specific color scheme:
 * - 0: Red (worst/highest risk)
 * - 1: Orange
 * - 2: Yellow
 * - 3: Blue
 * - 4: Green
 * - 5: Emerald (best/lowest risk)
 * - 6: Gray (not applicable/neutral)
 */

export interface ScoreLevelColors {
  bg: string
  text: string
  border: string
  hover: string
  // Full class string (for components that need all-in-one)
  full: string
  // Dark mode variants
  bgDark: string
  textDark: string
}

export const SCORE_LEVEL_COLORS: Record<number, ScoreLevelColors> = {
  0: {
    bg: 'bg-red-100',
    text: 'text-red-900',
    border: 'border-red-500',
    hover: 'hover:bg-red-200',
    full: 'bg-red-100 dark:bg-red-950 border-red-500 text-red-900 dark:text-red-100',
    bgDark: 'dark:bg-red-950',
    textDark: 'dark:text-red-100',
  },
  1: {
    bg: 'bg-orange-100',
    text: 'text-orange-900',
    border: 'border-orange-500',
    hover: 'hover:bg-orange-200',
    full: 'bg-orange-100 dark:bg-orange-950 border-orange-500 text-orange-900 dark:text-orange-100',
    bgDark: 'dark:bg-orange-950',
    textDark: 'dark:text-orange-100',
  },
  2: {
    bg: 'bg-yellow-100',
    text: 'text-yellow-900',
    border: 'border-yellow-500',
    hover: 'hover:bg-yellow-200',
    full: 'bg-yellow-100 dark:bg-yellow-950 border-yellow-500 text-yellow-900 dark:text-yellow-100',
    bgDark: 'dark:bg-yellow-950',
    textDark: 'dark:text-yellow-100',
  },
  3: {
    bg: 'bg-blue-100',
    text: 'text-blue-900',
    border: 'border-blue-500',
    hover: 'hover:bg-blue-200',
    full: 'bg-blue-100 dark:bg-blue-950 border-blue-500 text-blue-900 dark:text-blue-100',
    bgDark: 'dark:bg-blue-950',
    textDark: 'dark:text-blue-100',
  },
  4: {
    bg: 'bg-green-100',
    text: 'text-green-900',
    border: 'border-green-500',
    hover: 'hover:bg-green-200',
    full: 'bg-green-100 dark:bg-green-950 border-green-500 text-green-900 dark:text-green-100',
    bgDark: 'dark:bg-green-950',
    textDark: 'dark:text-green-100',
  },
  5: {
    bg: 'bg-emerald-100',
    text: 'text-emerald-900',
    border: 'border-emerald-500',
    hover: 'hover:bg-emerald-200',
    full: 'bg-emerald-100 dark:bg-emerald-950 border-emerald-500 text-emerald-900 dark:text-emerald-100',
    bgDark: 'dark:bg-emerald-950',
    textDark: 'dark:text-emerald-100',
  },
  6: {
    bg: 'bg-gray-100',
    text: 'text-gray-900',
    border: 'border-gray-500',
    hover: 'hover:bg-gray-200',
    full: 'bg-gray-100 dark:bg-gray-950 border-gray-500 text-gray-900 dark:text-gray-100',
    bgDark: 'dark:bg-gray-950',
    textDark: 'dark:text-gray-100',
  },
}

/**
 * Get color classes for a specific score level
 */
export function getScoreLevelColors(level: number): ScoreLevelColors {
  return SCORE_LEVEL_COLORS[level] || SCORE_LEVEL_COLORS[6]
}

/**
 * Get simplified color classes (for select dropdowns, small indicators)
 */
export function getScoreLevelColorSimple(level: number): string {
  const colors = getScoreLevelColors(level)
  return `${colors.bg} ${colors.border}`
}

/**
 * Get full color classes with dark mode support (for larger UI elements)
 */
export function getScoreLevelColorFull(level: number): string {
  const colors = getScoreLevelColors(level)
  return colors.full
}
