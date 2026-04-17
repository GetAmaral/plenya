import { brand } from '@plenya/brand';

export function OrganizationSchema() {
  const data = {
    '@context': 'https://schema.org',
    '@type': 'MedicalOrganization',
    name: brand.legalName,
    url: brand.url,
    logo: `${brand.url}/logo.svg`,
    sameAs: [brand.social.instagram],
    medicalSpecialty: ['FunctionalMedicine', 'PreventiveMedicine', 'Nephrology'],
  };
  return (
    <script
      type="application/ld+json"
      dangerouslySetInnerHTML={{ __html: JSON.stringify(data) }}
    />
  );
}
