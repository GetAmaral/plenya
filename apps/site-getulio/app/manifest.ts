import type { MetadataRoute } from 'next';

export default function manifest(): MetadataRoute.Manifest {
  return {
    name: 'Dr. Getúlio Amaral Filho',
    short_name: 'Dr. Getúlio',
    description:
      'Medicina guiada por raciocínio clínico. Nefrologista, professor, autor e diretor clínico da Plenya.',
    start_url: '/',
    display: 'standalone',
    background_color: '#F5EFE6',
    theme_color: '#0E2A2E',
    lang: 'pt-BR',
    orientation: 'portrait',
    icons: [
      { src: '/favicon.svg', sizes: 'any', type: 'image/svg+xml', purpose: 'any' },
      { src: '/favicon-192x192.png', sizes: '192x192', type: 'image/png', purpose: 'any' },
      { src: '/favicon-512x512.png', sizes: '512x512', type: 'image/png', purpose: 'any' },
      { src: '/maskable-512x512.png', sizes: '512x512', type: 'image/png', purpose: 'maskable' },
    ],
  };
}
