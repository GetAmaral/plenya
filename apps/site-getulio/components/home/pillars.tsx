import { Link } from '@/lib/i18n/navigation';

const pillars = [
  {
    label: 'Médico',
    title: 'Atendimento clínico',
    body: 'Atende em duas clínicas em Londrina e responde tecnicamente pela hemodiálise hospitalar.',
    href: '/onde-atendo',
  },
  {
    label: 'Professor',
    title: 'Coordenação de residência',
    body: 'Coordena a residência médica em nefrologia da Santa Casa de Londrina há mais de uma década.',
    href: '/ensino',
  },
  {
    label: 'Autor',
    title: 'A Janela Silenciosa',
    body: '"Antes — A Janela Silenciosa entre o Normal e o Ótimo" (2026).',
    href: '/livro',
  },
  {
    label: 'Palestrante',
    title: 'Palestras nacionais',
    body: 'Palestras para médicos, residentes e plateias corporativas — sobre prevenção que começa antes.',
    href: '/palestras',
  },
];

export function Pillars() {
  return (
    <section className="border-t border-rule">
      <div className="editorial-container py-20 md:py-28">
        <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-10 md:gap-14">
          {pillars.map((p) => (
            <Link
              key={p.label}
              href={p.href}
              className="group block border-l-2 border-bordo/30 pl-6 hover:border-bordo transition-colors"
            >
              <p className="label-meta mb-4">{p.label}</p>
              <h3 className="heading-section text-2xl mb-3 group-hover:text-bordo transition-colors">
                {p.title}
              </h3>
              <p className="font-serif text-ink-soft text-base leading-relaxed">{p.body}</p>
            </Link>
          ))}
        </div>
      </div>
    </section>
  );
}
