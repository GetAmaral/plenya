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
        // Paleta oficial — branding Dr. Getúlio Amaral.
        // Cream papel envelhecido (do wordmark sobre fundo claro).
        paper: '#f0e8d8',
        'paper-soft': '#e6dcc4',
        // Tinta sobre cream / texto sobre claro.
        ink: '#1a1a1a',
        'ink-soft': '#3a3a3a',
        'ink-muted': '#6b6b6b',
        // Acento dourado — filetes do "DR." e tagline.
        // Mantemos o token `bordo` resolvendo para o mesmo valor:
        // a paleta bordo era um equívoco antigo; agora "bordo" é o ouro
        // institucional. Evita renomear 45 ocorrências em 18 arquivos.
        gold: '#b48a4a',
        'gold-soft': '#c79f63',
        bordo: '#b48a4a',
        'bordo-soft': '#c79f63',
        // Navy quase preto — variação dark do branding (metade dos PNGs).
        navy: '#0d1f2e',
        'navy-soft': '#173248',
        // Linhas sutis sobre cream.
        rule: '#d4c8aa',
      },
      fontFamily: {
        // Cormorant Garamond — serifa de alto contraste, all-caps friendly,
        // bate exatamente com o wordmark do branding.
        serif: ['var(--font-serif)', 'Georgia', 'serif'],
        sans: ['var(--font-sans)', 'system-ui', 'sans-serif'],
      },
      letterSpacing: {
        widest: '0.2em',
        wordmark: '0.32em',
      },
    },
  },
  plugins: [],
};

export default config;
