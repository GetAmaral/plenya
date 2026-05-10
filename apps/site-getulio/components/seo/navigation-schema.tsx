const BASE = 'https://drgetulioamaralfilho.com.br';

const NAV_PT = [
  { name: 'Início', path: '' },
  { name: 'O médico', path: '/sobre' },
  { name: 'Livros', path: '/livros' },
  { name: 'Palestras', path: '/palestras' },
  { name: 'Ensino', path: '/ensino' },
  { name: 'Onde atendo', path: '/onde-atendo' },
  { name: 'Artigos', path: '/escritos' },
  { name: 'Contato', path: '/contato' },
];

const NAV_EN = [
  { name: 'Home', path: '/en' },
  { name: 'The physician', path: '/en/sobre' },
  { name: 'Books', path: '/en/livros' },
  { name: 'Lectures', path: '/en/palestras' },
  { name: 'Teaching', path: '/en/ensino' },
  { name: 'Where I practice', path: '/en/onde-atendo' },
  { name: 'Articles', path: '/en/escritos' },
  { name: 'Contact', path: '/en/contato' },
];

export function NavigationSchema({ locale = 'pt' }: { locale?: string } = {}) {
  const items = locale === 'en' ? NAV_EN : NAV_PT;
  const data = items.map((it, i) => ({
    '@context': 'https://schema.org',
    '@type': 'SiteNavigationElement',
    position: i + 1,
    name: it.name,
    url: it.path === '' ? `${BASE}/` : `${BASE}${it.path}`,
  }));
  return (
    <script
      type="application/ld+json"
      dangerouslySetInnerHTML={{ __html: JSON.stringify(data) }}
    />
  );
}
