import type { NextConfig } from 'next';
import createNextIntlPlugin from 'next-intl/plugin';

const withNextIntl = createNextIntlPlugin('./lib/i18n/request.ts');

const nextConfig: NextConfig = {
  reactStrictMode: true,
  poweredByHeader: false,
  transpilePackages: ['@plenya/brand', '@plenya/ui'],
  // Expose server-only env (Turbopack filtra process.env não-NEXT_PUBLIC_ no Server Components).
  // INTERNAL_API_URL aponta pro nome do serviço Docker quando rodando em container,
  // pra evitar que SSR tente alcançar localhost (que é o próprio container, não o host).
  env: {
    INTERNAL_API_URL: process.env.INTERNAL_API_URL ?? 'http://api:3001',
  },
  images: {
    formats: ['image/avif', 'image/webp'],
    remotePatterns: [
      { protocol: 'https', hostname: 'plenyasaude.com.br' },
      { protocol: 'https', hostname: 'cdn.plenyasaude.com.br' },
    ],
  },
  experimental: {
    optimizePackageImports: ['lucide-react'],
  },
  async headers() {
    return [
      {
        source: '/(.*)',
        headers: [
          { key: 'X-Content-Type-Options', value: 'nosniff' },
          { key: 'X-Frame-Options', value: 'SAMEORIGIN' },
          { key: 'Referrer-Policy', value: 'strict-origin-when-cross-origin' },
          { key: 'Permissions-Policy', value: 'camera=(), microphone=(), geolocation=()' },
        ],
      },
    ];
  },
  async redirects() {
    return [
      {
        source: '/blog/12-exames-que-valem-cada-centavo-e-12-que-sao-desperdicio',
        destination: '/blog/12-exames-que-um-checkup-de-longevidade-pede',
        permanent: true,
      },
      {
        source: '/en/blog/12-tests-worth-every-penny-and-12-that-are-wasted',
        destination: '/en/blog/12-tests-a-longevity-checkup-orders',
        permanent: true,
      },
      // Renomeação 2026-05-13: artigo "Suplementos depois dos 40" (remoção de tom polêmico, foco em pilares + caso a caso)
      {
        source: '/blog/suplementos-que-fazem-diferenca-depois-dos-40',
        destination: '/blog/suplementacao-depois-dos-40-o-que-faz-diferenca',
        permanent: true,
      },
      {
        source: '/en/blog/supplements-that-make-a-difference-after-40',
        destination: '/en/blog/supplementation-after-40-what-makes-a-difference',
        permanent: true,
      },
    ];
  },
};

export default withNextIntl(nextConfig);
