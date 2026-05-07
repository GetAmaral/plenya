/**
 * WebSite schema com SearchAction apontando para /escritos.
 * Habilita sitelinks searchbox do Google quando aplicável.
 */
const BASE = 'https://drgetulioamaralfilho.com.br';

export function WebSiteSchema({ locale = 'pt' }: { locale?: string } = {}) {
  const isEn = locale === 'en';
  const data = {
    '@context': 'https://schema.org',
    '@type': 'WebSite',
    '@id': `${BASE}/#website`,
    url: BASE,
    name: 'Dr. Getúlio Amaral Filho',
    description: isEn
      ? 'Medicine guided by clinical reasoning — nephrologist, professor, author.'
      : 'Medicina guiada por raciocínio clínico — nefrologista, professor, autor.',
    inLanguage: isEn ? 'en' : 'pt-BR',
    publisher: { '@id': `${BASE}/#person` },
    potentialAction: {
      '@type': 'SearchAction',
      target: {
        '@type': 'EntryPoint',
        urlTemplate: `${BASE}/escritos?q={search_term_string}`,
      },
      'query-input': 'required name=search_term_string',
    },
  };
  return (
    <script
      type="application/ld+json"
      dangerouslySetInnerHTML={{ __html: JSON.stringify(data) }}
    />
  );
}
