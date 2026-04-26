import Image from 'next/image';

type Clinic = {
  name: string;
  logo: { src: string; height: number } | null;
  role: string;
  body: string;
  href: string | null;
};

const clinics: Clinic[] = [
  {
    name: 'Plenya',
    logo: { src: '/logos/plenya.svg', height: 28 },
    role: 'Direção clínica',
    body: 'Medicina funcional integrativa e longevidade.',
    href: 'https://plenyasaude.com.br',
  },
  {
    name: 'Nefroclínica',
    logo: { src: '/logos/nefroclinica.png', height: 56 },
    role: 'Sócio',
    body: 'Nefrologia clínica em Londrina.',
    href: 'https://nefroclinica.com',
  },
  {
    name: 'DaVita Intra Hospitalar',
    logo: null,
    role: 'Responsável técnico',
    body: 'Hemodiálise hospitalar — Santa Casa de Londrina.',
    href: null,
  },
  {
    name: 'DaVita Londrina',
    logo: null,
    role: 'Responsável técnico',
    body: 'Unidade de hemodiálise ambulatorial em Londrina.',
    href: null,
  },
];

export function ClinicsRow() {
  return (
    <section className="border-t border-rule">
      <div className="editorial-container py-20 md:py-28">
        <p className="label-meta mb-12">Onde atendo</p>

        <div className="grid sm:grid-cols-2 lg:grid-cols-4 gap-12 md:gap-10">
          {clinics.map((c) => (
            <div key={c.name} className="space-y-6">
              <div className="h-14 flex items-center">
                {c.logo ? (
                  <Image
                    src={c.logo.src}
                    alt={c.name}
                    width={200}
                    height={c.logo.height}
                    style={{ height: c.logo.height, width: 'auto' }}
                    className="object-contain object-left"
                  />
                ) : (
                  <span className="font-serif text-ink text-2xl tracking-tight">{c.name}</span>
                )}
              </div>
              <div className="space-y-2">
                <p className="label-meta text-bordo">{c.role}</p>
                {c.logo && <h3 className="heading-section text-xl">{c.name}</h3>}
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
