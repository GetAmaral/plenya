import { brand } from '@plenya/brand';
import type { Doctor } from '@/lib/team';

export function PhysicianSchema({ doctor }: { doctor: Doctor }) {
  const isMedico = doctor.category === 'medico';
  const data = {
    '@context': 'https://schema.org',
    '@type': isMedico ? 'Physician' : 'Person',
    name: doctor.name,
    jobTitle: doctor.role,
    worksFor: {
      '@type': 'MedicalOrganization',
      name: brand.legalName,
      url: brand.url,
    },
    medicalSpecialty: isMedico ? doctor.specialties : undefined,
    identifier: doctor.credentials,
    sameAs: [doctor.instagram, doctor.linkedin, doctor.website].filter(Boolean),
    url: `${brand.url}/equipe/${doctor.slug}`,
  };
  return <script type="application/ld+json" dangerouslySetInnerHTML={{ __html: JSON.stringify(data) }} />;
}
