const BASE = 'https://drgetulioamaralfilho.com.br';

const NAV_PT = [
  { name: 'Início', path: '' },
  { name: 'O médico', path: '/o-medico' },
  { name: 'Livros', path: '/livros' },
  { name: 'Palestras', path: '/palestras' },
  { name: 'Ensino', path: '/ensino' },
  { name: 'Onde atendo', path: '/onde-atendo' },
  { name: 'Artigos', path: '/artigos' },
  { name: 'Contato', path: '/contato' },
];

const NAV_EN = [
  { name: 'Home', path: '/en' },
  { name: 'The physician', path: '/en/the-physician' },
  { name: 'Books', path: '/en/books' },
  { name: 'Lectures', path: '/en/lectures' },
  { name: 'Teaching', path: '/en/teaching' },
  { name: 'Where I practice', path: '/en/where-i-practice' },
  { name: 'Articles', path: '/en/articles' },
  { name: 'Contact', path: '/en/contact' },
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
