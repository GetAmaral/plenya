import type { NextConfig } from 'next';
import createNextIntlPlugin from 'next-intl/plugin';

const withNextIntl = createNextIntlPlugin('./lib/i18n/request.ts');

const nextConfig: NextConfig = {
  reactStrictMode: true,
  poweredByHeader: false,
  images: {
    formats: ['image/avif', 'image/webp'],
    remotePatterns: [
      { protocol: 'https', hostname: 'plenyasaude.com.br', pathname: '/images/**' },
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
      // Restruturação 2026-05: /livro (singular) → /livros/antes
      { source: '/livro', destination: '/livros/antes', permanent: true },
      { source: '/livro/excertos', destination: '/livros/antes/excertos', permanent: true },
      { source: '/en/livro', destination: '/en/livros/antes', permanent: true },
      { source: '/en/livro/excertos', destination: '/en/livros/antes/excertos', permanent: true },
    ];
  },
};

export default withNextIntl(nextConfig);
