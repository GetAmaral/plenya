import { brand } from '@plenya/brand';

type MedicalWebPageProps = {
  name: string;
  description: string;
  path: string;
  about?: string;
  audience?: 'Patient' | 'MedicalProfessional';
  lastReviewed?: string;
};

export function MedicalWebPageSchema({
  name,
  description,
  path,
  about,
  audience = 'Patient',
  lastReviewed,
}: MedicalWebPageProps) {
  const url = `${brand.url}${path}`;
  const data = {
    '@context': 'https://schema.org',
    '@type': 'MedicalWebPage',
    name,
    description,
    url,
    inLanguage: 'pt-BR',
    isPartOf: { '@id': `${brand.url}/#clinic` },
    publisher: { '@id': `${brand.url}/#clinic` },
    reviewedBy: { '@id': `${brand.url}/dr-getulio#physician` },
    audience: { '@type': 'MedicalAudience', audienceType: audience },
    ...(about ? { about: { '@type': 'MedicalCondition', name: about } } : {}),
    ...(lastReviewed ? { lastReviewed } : { lastReviewed: new Date().toISOString().slice(0, 10) }),
  };
  return (
    <script
      type="application/ld+json"
      dangerouslySetInnerHTML={{ __html: JSON.stringify(data) }}
    />
  );
}
