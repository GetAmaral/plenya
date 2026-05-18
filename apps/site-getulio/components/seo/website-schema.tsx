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
  };
  return (
    <script
      type="application/ld+json"
      dangerouslySetInnerHTML={{ __html: JSON.stringify(data) }}
    />
  );
}
