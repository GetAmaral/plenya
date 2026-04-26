import Image from 'next/image';

const clinics = [
  {
    name: 'Plenya',
    logo: '/logos/plenya.svg',
    logoH: 28,
    role: 'Direção clínica',
    body: 'Medicina funcional integrativa e longevidade.',
    href: 'https://plenyasaude.com.br',
    external: true,
  },
  {
    name: 'Nefroclínica',
    logo: '/logos/nefroclinica.png',
    logoH: 56,
    role: 'Sócio',
    body: 'Nefrologia clínica em Londrina.',
    href: 'https://nefroclinica.com',
    external: true,
  },
  {
    name: 'DaVita Intra Hospitalar de Londrina',
    logo: '/logos/davita.svg',
    logoH: 36,
    role: 'Responsável técnico',
    body: 'Hemodiálise hospitalar — Santa Casa de Londrina.',
    href: null,
    external: false,
  },
];

export function ClinicsRow() {
  return (
    <section className="border-t border-rule">
      <div className="editorial-container py-20 md:py-28">
        <p className="label-meta mb-12">Onde atendo</p>

        <div className="grid md:grid-cols-3 gap-12 md:gap-16">
          {clinics.map((c) => (
            <div key={c.name} className="space-y-6">
              <div className="h-14 flex items-center">
                <Image
                  src={c.logo}
                  alt={c.name}
                  width={200}
                  height={c.logoH}
                  style={{ height: c.logoH, width: 'auto' }}
                  className="object-contain object-left"
                />
              </div>
              <div className="space-y-2">
                <p className="label-meta text-bordo">{c.role}</p>
                <h3 className="heading-section text-xl">{c.name}</h3>
                <p className="font-serif text-ink-soft leading-relaxed">{c.body}</p>
              </div>
              {c.href && (
                <a
                  href={c.href}
                  target="_blank"
                  rel="noreferrer"
                  className="link-text inline-block font-sans text-sm"
                >
                  Visitar site ↗
                </a>
              )}
            </div>
          ))}
        </div>
      </div>
    </section>
  );
}
