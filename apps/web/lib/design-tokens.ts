/**
 * Design Tokens - Single Source of Truth
 * Baseado nos padrões de janeiro 2026
 */

export const designTokens = {
  // Spacing Scale (4px base)
  spacing: {
    0: "0",
    0.5: "0.125rem", // 2px
    1: "0.25rem", // 4px
    1.5: "0.375rem", // 6px
    2: "0.5rem", // 8px
    2.5: "0.625rem", // 10px
    3: "0.75rem", // 12px
    3.5: "0.875rem", // 14px
    4: "1rem", // 16px
    5: "1.25rem", // 20px
    6: "1.5rem", // 24px
    7: "1.75rem", // 28px
    8: "2rem", // 32px
    9: "2.25rem", // 36px
    10: "2.5rem", // 40px
    12: "3rem", // 48px
    14: "3.5rem", // 56px
    16: "4rem", // 64px
    20: "5rem", // 80px
    24: "6rem", // 96px
    28: "7rem", // 112px
    32: "8rem", // 128px
  },

  // Typography Scale (Fluid Typography)
  typography: {
    fontFamily: {
      sans: "var(--font-inter), system-ui, sans-serif",
      mono: "var(--font-mono), ui-monospace, monospace",
    },
    fontSize: {
      xs: ["0.75rem", { lineHeight: "1rem" }],
      sm: ["0.875rem", { lineHeight: "1.25rem" }],
      base: ["1rem", { lineHeight: "1.5rem" }],
      lg: ["1.125rem", { lineHeight: "1.75rem" }],
      xl: ["1.25rem", { lineHeight: "1.75rem" }],
      "2xl": ["1.5rem", { lineHeight: "2rem" }],
      "3xl": ["1.875rem", { lineHeight: "2.25rem" }],
      "4xl": ["2.25rem", { lineHeight: "2.5rem" }],
      "5xl": ["3rem", { lineHeight: "1" }],
      "6xl": ["3.75rem", { lineHeight: "1" }],
    },
    fontWeight: {
      normal: "400",
      medium: "500",
      semibold: "600",
      bold: "700",
    },
  },

  // Border Radius (Modern, mais arredondado)
  borderRadius: {
    none: "0",
    sm: "0.25rem", // 4px
    DEFAULT: "0.5rem", // 8px - padrão moderno
    md: "0.75rem", // 12px
    lg: "1rem", // 16px
    xl: "1.25rem", // 20px
    "2xl": "1.5rem", // 24px
    "3xl": "2rem", // 32px
    full: "9999px",
  },

  // Shadows (Soft, modernos)
  boxShadow: {
    sm: "0 1px 2px 0 rgb(0 0 0 / 0.05)",
    DEFAULT:
      "0 1px 3px 0 rgb(0 0 0 / 0.1), 0 1px 2px -1px rgb(0 0 0 / 0.1)",
    md: "0 4px 6px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1)",
    lg: "0 10px 15px -3px rgb(0 0 0 / 0.1), 0 4px 6px -4px rgb(0 0 0 / 0.1)",
    xl: "0 20px 25px -5px rgb(0 0 0 / 0.1), 0 8px 10px -6px rgb(0 0 0 / 0.1)",
    "2xl": "0 25px 50px -12px rgb(0 0 0 / 0.25)",
    inner: "inset 0 2px 4px 0 rgb(0 0 0 / 0.05)",
  },

  // Animation Timing
  animation: {
    duration: {
      fast: "150ms",
      normal: "250ms",
      slow: "350ms",
      slower: "500ms",
    },
    easing: {
      default: "cubic-bezier(0.4, 0, 0.2, 1)",
      in: "cubic-bezier(0.4, 0, 1, 1)",
      out: "cubic-bezier(0, 0, 0.2, 1)",
      inOut: "cubic-bezier(0.4, 0, 0.2, 1)",
    },
  },

  // Medical-specific colors (semantic)
  medical: {
    patient: {
      stable: "hsl(142 76% 36%)", // Verde
      observation: "hsl(38 92% 50%)", // Amarelo
      critical: "hsl(0 84% 60%)", // Vermelho
      unknown: "hsl(220 13% 46%)", // Cinza
    },
    priority: {
      urgent: "hsl(0 84% 60%)",
      high: "hsl(25 95% 53%)",
      normal: "hsl(210 100% 50%)",
      low: "hsl(220 13% 46%)",
    },
    status: {
      success: "hsl(142 76% 36%)",
      warning: "hsl(38 92% 50%)",
      error: "hsl(0 84% 60%)",
      info: "hsl(199 89% 48%)",
    },
  },
} as const;

// Helper function para acessar tokens
export function token<T extends keyof typeof designTokens>(
  category: T
): (typeof designTokens)[T] {
  return designTokens[category];
}
