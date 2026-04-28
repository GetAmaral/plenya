import { brand } from '@plenya/brand';

/**
 * WebSite + SearchAction schema. Habilita o sitelinks searchbox do Google
 * (caixa de busca direto na SERP). publisher referencia a MedicalOrganization.
 */
export function WebSiteSchema() {
  const data = {
    '@context': 'https://schema.org',
    '@type': 'WebSite',
    '@id': `${brand.url}/#website`,
    url: brand.url,
    name: brand.name,
    alternateName: brand.legalName,
    description: brand.tagline,
    inLanguage: 'pt-BR',
    publisher: { '@id': `${brand.url}/#organization` },
    potentialAction: {
      '@type': 'SearchAction',
      target: {
        '@type': 'EntryPoint',
        urlTemplate: `${brand.url}/blog?q={search_term_string}`,
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
