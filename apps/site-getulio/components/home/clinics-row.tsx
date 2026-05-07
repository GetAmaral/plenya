import { useTranslations } from 'next-intl';

type Clinic = {
  name: string;
  keyPrefix: 'plenya' | 'nefroclinica' | 'davitaIntra' | 'davitaLondrina';
  href: string | null;
};

// Logos foram padronizados como wordmark serif (texto): mistura de PNG raster
// (Nefroclínica) + SVG (Plenya) + texto (DaVita) destoava. Tipografia serif
// uniforme reforça o tom editorial/livraria do site.
const clinics: Clinic[] = [
  { name: 'Plenya', keyPrefix: 'plenya', href: 'https://plenyasaude.com.br' },
  { name: 'Nefroclínica Londrina', keyPrefix: 'nefroclinica', href: 'https://nefroclinica.com' },
  { name: 'DaVita Intra Hospitalar', keyPrefix: 'davitaIntra', href: null },
  { name: 'DaVita Londrina', keyPrefix: 'davitaLondrina', href: null },
];

export function ClinicsRow() {
  const t = useTranslations('home.clinics');
  const tCommon = useTranslations('common');
  return (
    <section className="border-t border-rule">
      <div className="editorial-container py-20 md:py-28">
        <p className="label-meta-lg mb-12">{t('sectionLabel')}</p>

        <div className="grid sm:grid-cols-2 lg:grid-cols-4 gap-12 md:gap-10">
          {clinics.map((c) => (
            <div key={c.name} className="space-y-5 border-t border-rule pt-6">
              <p className="label-meta text-bordo">{t(`${c.keyPrefix}Role`)}</p>
              <h3 className="font-serif text-2xl text-ink leading-tight">{c.name}</h3>
              <p className="font-serif text-ink-soft leading-relaxed">{t(`${c.keyPrefix}Body`)}</p>
              {c.href && (
                <a
                  href={c.href}
                  target="_blank"
                  rel="noreferrer"
                  className="link-text inline-block font-sans text-sm"
                >
                  {tCommon('visitSite')}
                </a>
              )}
            </div>
          ))}
        </div>
      </div>
    </section>
  );
}
