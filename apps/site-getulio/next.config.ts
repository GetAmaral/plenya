import type { NextConfig } from 'next';
import path from 'path';
import createNextIntlPlugin from 'next-intl/plugin';

const withNextIntl = createNextIntlPlugin('./lib/i18n/request.ts');

const nextConfig: NextConfig = {
  // Saída standalone: o build emite um server.js mínimo + node_modules tree-shaken,
  // rodado por `node server.js` (sem o wrapper `pnpm start`, ~98MB de RAM a menos por app).
  // outputFileTracingRoot = raiz do monorepo, senão o tracing de deps quebra.
  output: 'standalone',
  outputFileTracingRoot: path.join(__dirname, '../../'),
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

      // Slug EN dedicado pra palestra "Janela Silenciosa" (era servida com slug PT em /en/)
      {
        source: '/en/lectures/janela-silenciosa',
        destination: '/en/lectures/silent-window',
        permanent: true,
      },

      // 2026-05-18: Slugs EN servidos historicamente sob /artigos/ (path PT) → /en/articles/
      {
        source: '/artigos/the-inverted-pyramid-of-longevity',
        destination: '/en/articles/the-inverted-pyramid-of-longevity',
        permanent: true,
      },
      {
        source: '/artigos/lipoprotein-a-the-test-your-cardiologist-didnt-order',
        destination: '/en/articles/lipoprotein-a-the-test-your-cardiologist-didnt-order',
        permanent: true,
      },
      {
        source: '/artigos/meditation-that-works-in-8-minutes',
        destination: '/en/articles/meditation-that-works-in-8-minutes',
        permanent: true,
      },
      {
        source: '/artigos/ferritin-30-to-100-the-normality-that-drains-women',
        destination: '/en/articles/ferritin-30-to-100-the-normality-that-drains-women',
        permanent: true,
      },
      {
        source: '/artigos/supplementation-after-40-what-makes-a-difference',
        destination: '/en/articles/supplementation-after-40-what-makes-a-difference',
        permanent: true,
      },
      {
        source: '/artigos/what-the-annual-checkup-doesnt-show-about-your-heart',
        destination: '/en/articles/what-the-annual-checkup-doesnt-show-about-your-heart',
        permanent: true,
      },
      {
        source: '/artigos/normal-vs-optimal',
        destination: '/en/articles/normal-vs-optimal',
        permanent: true,
      },
      // 2026-05-18: variante antiga do "12-tests" sob /artigos/ (path PT) → versão renomeada em /en/
      {
        source: '/artigos/12-tests-worth-every-penny-and-12-that-are-wasted',
        destination: '/en/articles/12-tests-a-longevity-checkup-orders',
        permanent: true,
      },
      // 2026-05-18: slug PT sob /en/articles/ → slug EN traduzido
      {
        source: '/en/articles/quatro-profissionais-falando-a-mesma-lingua',
        destination: '/en/articles/four-professionals-speaking-the-same-language',
        permanent: true,
      },
    ];
  },
};

export default withNextIntl(nextConfig);
