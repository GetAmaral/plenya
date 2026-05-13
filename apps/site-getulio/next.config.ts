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
      { source: '/en/livro', destination: '/en/books/antes', permanent: true },
      { source: '/en/livro/excertos', destination: '/en/books/antes/excerpts', permanent: true },

      // Rename 2026-05-10: /sobre → /o-medico, /escritos → /artigos (PT)
      { source: '/sobre', destination: '/o-medico', permanent: true },
      { source: '/escritos', destination: '/artigos', permanent: true },
      { source: '/escritos/:slug', destination: '/artigos/:slug', permanent: true },

      // Slugs EN traduzidos (antes serviam PT em /en/)
      { source: '/en/sobre', destination: '/en/the-physician', permanent: true },
      { source: '/en/escritos', destination: '/en/articles', permanent: true },
      { source: '/en/escritos/:slug', destination: '/en/articles/:slug', permanent: true },
      { source: '/en/livros', destination: '/en/books', permanent: true },
      { source: '/en/livros/:slug', destination: '/en/books/:slug', permanent: true },
      { source: '/en/livros/:slug/excertos', destination: '/en/books/:slug/excerpts', permanent: true },
      { source: '/en/palestras', destination: '/en/lectures', permanent: true },
      { source: '/en/palestras/:slug', destination: '/en/lectures/:slug', permanent: true },
      { source: '/en/ensino', destination: '/en/teaching', permanent: true },
      { source: '/en/onde-atendo', destination: '/en/where-i-practice', permanent: true },
      { source: '/en/contato', destination: '/en/contact', permanent: true },

      // Renomeação 2026-05-13: artigo "12 exames..." (remoção de tom polêmico)
      {
        source: '/artigos/12-exames-que-valem-cada-centavo-e-12-que-sao-desperdicio',
        destination: '/artigos/12-exames-que-um-checkup-de-longevidade-pede',
        permanent: true,
      },
      {
        source: '/en/articles/12-tests-worth-every-penny-and-12-that-are-wasted',
        destination: '/en/articles/12-tests-a-longevity-checkup-orders',
        permanent: true,
      },
      // Renomeação 2026-05-13: artigo "Suplementos depois dos 40" (remoção de tom polêmico, foco em pilares + caso a caso)
      {
        source: '/artigos/suplementos-que-fazem-diferenca-depois-dos-40',
        destination: '/artigos/suplementacao-depois-dos-40-o-que-faz-diferenca',
        permanent: true,
      },
      {
        source: '/en/articles/supplements-that-make-a-difference-after-40',
        destination: '/en/articles/supplementation-after-40-what-makes-a-difference',
        permanent: true,
      },
    ];
  },
};

export default withNextIntl(nextConfig);
