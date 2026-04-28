type Clinic = {
  name: string;
  role: string;
  body: string;
  href: string | null;
};

// Logos foram padronizados como wordmark serif (texto): mistura de PNG raster
// (Nefroclínica) + SVG (Plenya) + texto (DaVita) destoava. Tipografia serif
// uniforme reforça o tom editorial/livraria do site.
const clinics: Clinic[] = [
  {
    name: 'Plenya',
    role: 'Direção clínica',
    body: 'Medicina funcional integrativa e longevidade.',
    href: 'https://plenyasaude.com.br',
  },
  {
    name: 'Nefroclínica Londrina',
    role: 'Sócio',
    body: 'Nefrologia clínica em Londrina.',
    href: 'https://nefroclinica.com',
  },
  {
    name: 'DaVita Intra Hospitalar',
    role: 'Responsável técnico',
    body: 'Hemodiálise hospitalar — Santa Casa de Londrina.',
    href: null,
  },
  {
    name: 'DaVita Londrina',
    role: 'Responsável técnico',
    body: 'Unidade ambulatorial de hemodiálise.',
    href: null,
  },
];

export function ClinicsRow() {
  return (
    <section className="border-t border-rule">
      <div className="editorial-container py-20 md:py-28">
        <p className="label-meta-lg mb-12">Onde atendo</p>

        <div className="grid sm:grid-cols-2 lg:grid-cols-4 gap-12 md:gap-10">
          {clinics.map((c) => (
            <div key={c.name} className="space-y-5 border-t border-rule pt-6">
              <p className="label-meta text-bordo">{c.role}</p>
              <h3 className="font-serif text-2xl text-ink leading-tight">{c.name}</h3>
              <p className="font-serif text-ink-soft leading-relaxed">{c.body}</p>
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
