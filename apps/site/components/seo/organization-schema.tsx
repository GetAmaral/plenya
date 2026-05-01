import { brand } from '@plenya/brand';

export function OrganizationSchema({ locale = 'pt' }: { locale?: string } = {}) {
  const inLanguage = locale === 'en' ? 'en' : 'pt-BR';
  const data = {
    '@context': 'https://schema.org',
    '@type': 'MedicalOrganization',
    '@id': `${brand.url}/#organization`,
    name: brand.legalName,
    alternateName: brand.name,
    url: brand.url,
    logo: `${brand.url}/logo.svg`,
    image: `${brand.url}/og-default.jpg`,
    description: brand.tagline,
    inLanguage,
    sameAs: [
      brand.social.instagram,
      'https://drgetulioamaralfilho.com.br',
      'https://instagram.com/drGetulioAmaralFilho',
    ],
    medicalSpecialty: ['FunctionalMedicine', 'PreventiveMedicine', 'Nephrology', 'InternalMedicine'],
    areaServed: { '@type': 'Country', name: locale === 'en' ? 'Brazil' : 'Brasil' },
    knowsLanguage: ['pt-BR', 'en'],
  };
  return (
    <script
      type="application/ld+json"
      dangerouslySetInnerHTML={{ __html: JSON.stringify(data) }}
    />
  );
}
