import type { Config } from 'tailwindcss';

const config: Config = {
  content: [
    './app/**/*.{ts,tsx,mdx}',
    './components/**/*.{ts,tsx}',
    './content/**/*.{md,mdx}',
  ],
  theme: {
    extend: {
      colors: {
        // Opção A — "Ensaio editorial"
        paper: '#f5f1e8',          // off-white papel
        ink: '#1a1a1a',            // tinta nanquim
        'ink-soft': '#3a3a3a',
        'ink-muted': '#6b6b6b',
        bordo: '#6b1f2a',          // acento bordô-vinho (capa de livro)
        'bordo-soft': '#8a3340',
        rule: '#dcd6c5',           // linhas/divisores
      },
      fontFamily: {
        serif: ['var(--font-serif)', 'Georgia', 'serif'],
        sans: ['var(--font-sans)', 'system-ui', 'sans-serif'],
      },
      letterSpacing: {
        widest: '0.2em',
      },
    },
  },
  plugins: [],
};

export default config;
