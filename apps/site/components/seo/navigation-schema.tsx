import { brand } from '@plenya/brand';

const NAV_PT = [
  { name: 'Início', path: '' },
  { name: 'Como funciona', path: '/como-funciona' },
  { name: 'Consultas', path: '/consultas' },
  { name: 'Continuum', path: '/continuum' },
  { name: 'Escore Plenya', path: '/escore-plenya' },
  { name: 'Método AGIR', path: '/metodo-agir' },
  { name: 'A Plenya', path: '/a-plenya' },
  { name: 'Dr. Getúlio', path: '/dr-getulio' },
  { name: 'Equipe', path: '/equipe' },
  { name: 'Depoimentos', path: '/depoimentos' },
  { name: 'Diagnóstico', path: '/diagnostico' },
  { name: 'Casos', path: '/casos' },
  { name: 'Blog', path: '/blog' },
  { name: 'Boletim', path: '/boletim' },
  { name: 'Contato', path: '/contato' },
];

const NAV_EN = [
  { name: 'Home', path: '/en' },
  { name: 'How it works', path: '/en/how-it-works' },
  { name: 'Consultations', path: '/en/consultations' },
  { name: 'Continuum', path: '/en/continuum' },
  { name: 'Plenya Score', path: '/en/plenya-score' },
  { name: 'AGIR Method', path: '/en/acts-method' },
  { name: 'About Plenya', path: '/en/about' },
  { name: 'Dr. Getúlio', path: '/en/dr-getulio' },
  { name: 'Team', path: '/en/team' },
  { name: 'Testimonials', path: '/en/testimonials' },
  { name: 'Diagnostic', path: '/en/assessment' },
  { name: 'Cases', path: '/en/cases' },
  { name: 'Blog', path: '/en/blog' },
  { name: 'Newsletter', path: '/en/newsletter' },
  { name: 'Contact', path: '/en/contact' },
];

export function NavigationSchema({ locale = 'pt' }: { locale?: string } = {}) {
  const items = locale === 'en' ? NAV_EN : NAV_PT;
  const data = items.map((it, i) => ({
    '@context': 'https://schema.org',
    '@type': 'SiteNavigationElement',
    position: i + 1,
    name: it.name,
    url: it.path === '' ? `${brand.url}/` : `${brand.url}${it.path}`,
  }));
  return (
    <script
      type="application/ld+json"
      dangerouslySetInnerHTML={{ __html: JSON.stringify(data) }}
    />
  );
}
