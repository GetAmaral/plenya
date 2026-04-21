import { PlenyaInfinity } from '@plenya/brand/logo';
import { Link } from '@/lib/i18n/navigation';

const pilares = [
  { label: 'Modalidades', value: 'Semestral ou anual' },
  { label: 'Formato', value: '100% online' },
  { label: 'Equipe', value: 'Médicos, nutri, psicóloga e EF' },
  { label: 'Cadência', value: 'Encontros semanais em rotação' },
  { label: 'Diagnóstico', value: 'Escore Plenya — 800+ itens' },
  { label: 'Método', value: 'AGIR — quatro pilares' },
];

export function ContinuumSpotlight() {
  return (
    <section className="bg-petrol text-cream relative overflow-hidden">
      <PlenyaInfinity
        aria-hidden="true"
        focusable="false"
        className="hidden md:block absolute -right-20 top-16 w-[420px] h-auto text-gold/[0.07] pointer-events-none"
      />

      <div className="relative site-container section">
        <div className="grid lg:grid-cols-[1.1fr_1fr] gap-12 lg:gap-20 mb-16">
          <div className="space-y-6">
            <p className="label-upper text-gold">Programa contínuo</p>
            <h2 className="heading-section text-cream text-3xl md:text-5xl">
              Continuum Plenya.
            </h2>
            <p className="heading-section text-gold text-xl md:text-2xl">
              Saúde que se sustenta no tempo.
            </p>
          </div>

          <div className="space-y-5 text-cream/85 text-lg leading-relaxed self-end">
            <p>
              O programa de acompanhamento contínuo da Plenya. Médicos,
              nutricionista, psicóloga e educador físico construindo um plano
              único, mensurado pelo <strong className="text-cream">Escore Plenya</strong> e
              organizado pelo <strong className="text-cream">Método AGIR</strong>.
            </p>
            <p className="text-cream/70">
              Não é uma consulta que se repete. É uma trajetória que se
              acompanha — com leitura ampliada, conduta integrada e ajuste
              contínuo a cada ciclo.
            </p>
          </div>
        </div>

        <div className="grid sm:grid-cols-2 lg:grid-cols-3 gap-px bg-cream/10 border-y border-cream/15 mb-12">
          {pilares.map((p) => (
            <div key={p.label} className="bg-petrol p-6 space-y-2">
              <p className="label-upper text-gold/80">{p.label}</p>
              <p className="text-cream text-lg">{p.value}</p>
            </div>
          ))}
        </div>

        <div className="flex flex-wrap gap-5">
          <Link href="/continuum" className="btn-gold">
            Conhecer o Continuum Plenya
          </Link>
          <Link
            href="/diagnostico"
            className="btn-outline-light"
          >
            É para mim?
          </Link>
        </div>
      </div>
    </section>
  );
}
