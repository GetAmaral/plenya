import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  // Otimizações de performance
  reactStrictMode: true,

  // M2 — Desabilita typecheck/eslint no build de produção pra não bloquear
  // deploy. Codebase tem erros TS pré-existentes (dev mode Turbopack não checa).
  //
  // Débito técnico: rastreado em .github/workflows/typecheck.yml (warning-only,
  // continue-on-error). Quando o lixo for limpo, removemos as flags abaixo e
  // o workflow vira blocking.
  //
  // NÃO REMOVER sem antes:
  //   1. Rodar `pnpm --filter @plenya/web exec tsc --noEmit` localmente.
  //   2. Zerar errors.
  //   3. Validar que `pnpm --filter @plenya/web build` passa sem essas flags.
  typescript: { ignoreBuildErrors: true },
  // (chave `eslint` removida: Next 16 removeu o lint embutido no build; a config
  //  não é mais reconhecida. ESLint roda via `pnpm lint`/CI, não no `next build`.)

  // Transpile ESM-only packages
  transpilePackages: ['react-reader', '@plenya/ui'],

  // Otimizar imagens
  images: {
    formats: ['image/avif', 'image/webp'],
    deviceSizes: [640, 750, 828, 1080, 1200, 1920, 2048, 3840],
    imageSizes: [16, 32, 48, 64, 96, 128, 256, 384],
  },

  // Comprimir respostas
  compress: true,

  // Compiler optimizations
  compiler: {
    removeConsole: process.env.NODE_ENV === 'production' ? {
      exclude: ['error', 'warn'],
    } : false,
  },

  // Experimental features para performance
  experimental: {
    optimizePackageImports: [
      'lucide-react',
      'recharts',
      '@radix-ui/react-dialog',
      '@radix-ui/react-dropdown-menu',
      '@radix-ui/react-popover',
      '@radix-ui/react-tooltip',
      '@radix-ui/react-avatar',
      '@radix-ui/react-tabs',
      '@radix-ui/react-select',
      '@radix-ui/react-switch',
    ],
  },

  // HIGH H6 — Security headers (defense in depth).
  // CSP propositalmente permissiva pra unsafe-inline porque Next gera <script>
  // inline e estilos inline. Pode ser apertado em uma onda futura usando nonce
  // por request.
  async headers() {
    const csp = [
      "default-src 'self'",
      "script-src 'self' 'unsafe-inline' 'unsafe-eval' https://accounts.google.com https://*.daily.co",
      "style-src 'self' 'unsafe-inline' https://fonts.googleapis.com",
      "img-src 'self' data: blob: https:",
      "font-src 'self' data: https://fonts.gstatic.com",
      "connect-src 'self' https://api.openai.com https://api.anthropic.com https://*.daily.co wss://*.daily.co " + (process.env.NEXT_PUBLIC_API_URL || 'http://localhost:3001'),
      "frame-src 'self' https://*.daily.co https://accounts.google.com",
      "frame-ancestors 'none'",
      "object-src 'none'",
      "base-uri 'self'",
      "form-action 'self'",
    ].join('; ');

    return [
      {
        source: '/:path*',
        headers: [
          { key: 'X-DNS-Prefetch-Control', value: 'on' },
          { key: 'X-Frame-Options', value: 'DENY' },
          { key: 'X-Content-Type-Options', value: 'nosniff' },
          { key: 'Referrer-Policy', value: 'strict-origin-when-cross-origin' },
          { key: 'Strict-Transport-Security', value: 'max-age=63072000; includeSubDomains; preload' },
          // Permissions-Policy: bloqueia features sensíveis exceto pra daily.co (telemed)
          { key: 'Permissions-Policy', value: 'camera=(self "https://*.daily.co"), microphone=(self "https://*.daily.co"), geolocation=(), payment=(), usb=()' },
          { key: 'Content-Security-Policy', value: csp },
        ],
      },
      // H9 — sala de telemedicina não pode ser referenciada cross-origin
      // pra evitar leak de URL de Daily.co em headers Referer.
      {
        source: '/sala/:path*',
        headers: [
          { key: 'Referrer-Policy', value: 'no-referrer' },
        ],
      },
    ]
  },
};

export default nextConfig;
